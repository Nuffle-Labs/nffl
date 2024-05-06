use actix_web::{web, App, HttpResponse, HttpServer};
use prometheus::{
    core::{AtomicF64, GenericCounter},
    Counter, Encoder, Histogram, HistogramOpts, Opts, Registry, TextEncoder,
};
use std::{io::ErrorKind, net::SocketAddr, time::Duration};
use tokio::task::JoinHandle;
use tracing::{error, info};

use crate::errors::Result;

const METRICS: &str = "metrics";
const INDEXER_NAMESPACE: &str = "sffl_indexer";
const CANDIDATES_SUBSYSTEM: &str = "candidates_validator";
const LISTENER_SUBSYSTEM: &str = "block_listener";
const PUBLISHER_SUBSYSTEM: &str = "rabbit_publisher";

#[derive(Clone)]
pub struct CandidatesListener {
    pub num_successful: GenericCounter<AtomicF64>,
    pub num_failed: GenericCounter<AtomicF64>,
}

pub struct BlockEventListener {
    pub num_candidates: GenericCounter<AtomicF64>,
}

#[derive(Clone)]
pub struct PublisherListener {
    pub num_published_blocks: GenericCounter<AtomicF64>,
    pub num_failed_publishes: GenericCounter<AtomicF64>,
    pub publish_duration_histogram: Histogram,
}

pub trait Metricable {
    fn enable_metrics(&mut self, registry: Registry) -> Result<()>;
}

pub(crate) fn make_candidates_validator_metrics(registry: Registry) -> Result<CandidatesListener> {
    let opts = Opts::new("num_of_successful_candidates", "Number of successful candidates")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(CANDIDATES_SUBSYSTEM);
    let num_successful = Counter::with_opts(opts)?;

    registry.register(Box::new(num_successful.clone()))?;

    let opts = Opts::new("num_of_failed_candidates", "Number of rejected candidates")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(CANDIDATES_SUBSYSTEM);
    let num_failed = Counter::with_opts(opts)?;

    registry.register(Box::new(num_failed.clone()))?;

    Ok(CandidatesListener {
        num_successful,
        num_failed,
    })
}

pub(crate) fn make_block_listener_metrics(registry: Registry) -> Result<BlockEventListener> {
    let opts = Opts::new("num_of_candidates", "Number of candidates indexed")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(LISTENER_SUBSYSTEM);
    let num_candidates = Counter::with_opts(opts)?;
    registry.register(Box::new(num_candidates.clone()))?;

    Ok(BlockEventListener { num_candidates })
}

pub(crate) fn make_publisher_metrics(registry: Registry) -> Result<PublisherListener> {
    let opts = Opts::new("num_of_published_blocks", "Number of published blocks to MQ")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(PUBLISHER_SUBSYSTEM);
    let num_published_blocks = Counter::with_opts(opts)?;

    registry.register(Box::new(num_published_blocks.clone()))?;

    let opts = Opts::new("num_failed_published", "Number of failed publishes to MQ")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(PUBLISHER_SUBSYSTEM);
    let num_failed_publishes = Counter::with_opts(opts)?;

    registry.register(Box::new(num_failed_publishes.clone()))?;

    let publish_duration_opts = HistogramOpts::new("publish_duration_seconds", "Time spent in publishing messages")
        .namespace(INDEXER_NAMESPACE)
        .subsystem(PUBLISHER_SUBSYSTEM)
        .buckets(vec![5., 10., 25., 50., 100., 250., 500., 1000., 2500., 5000., 10000.]); // ms
    let publish_duration_histogram = Histogram::with_opts(publish_duration_opts)?;

    registry.register(Box::new(publish_duration_histogram.clone()))?;

    Ok(PublisherListener {
        num_published_blocks,
        num_failed_publishes,
        publish_duration_histogram,
    })
}

async fn metrics(registry: web::Data<Registry>) -> HttpResponse {
    let metric_families = registry.gather();
    let mut buffer = Vec::new();
    let encoder = TextEncoder::new();

    match encoder.encode(&metric_families, &mut buffer) {
        Ok(_) => HttpResponse::Ok().content_type(encoder.format_type()).body(buffer),
        Err(e) => HttpResponse::InternalServerError().body(e.to_string()),
    }
}

async fn metrics_runner(metrics_addr: SocketAddr, registry: Registry) {
    const RECONNECTION_RETRIES: u32 = 3;
    const RECONNECTION_INTERVAL: Duration = Duration::from_secs(1);

    let create_server = |registry: Registry| {
        let registry_data = web::Data::new(registry);
        let metrics_server = HttpServer::new(move || {
            App::new()
                .app_data(registry_data.clone())
                .service(web::resource("/metrics").route(web::get().to(metrics)))
        })
        .bind(metrics_addr)?;

        Ok::<_, std::io::Error>(metrics_server)
    };

    let mut timeout = false;
    for _ in 0..RECONNECTION_RETRIES {
        if timeout {
            info!(target: METRICS, "reconnecting");
            tokio::time::sleep(RECONNECTION_INTERVAL).await;
        }

        let metrics_server = match create_server(registry.clone()) {
            Ok(v) => v,
            Err(err) => {
                if !can_recover(err.kind()) {
                    error!(target: METRICS, "Couldn't create server with code: {}", err.kind());
                    return;
                }

                // TODO: with the compilation issue with this
                // tokio::time::sleep(RECONNECTION_INTERVAL).await;\
                timeout = true;
                continue;
            }
        };

        match metrics_server.run().await {
            Ok(()) => {
                // Graceful shutdown
                info!(target: METRICS, "Server shutdown gracefully");
                return;
            }
            Err(err) => {
                if !can_recover(err.kind()) {
                    error!(target: METRICS, "Error while running server: {}", err.kind());
                    return;
                }

                timeout = true;
                continue;
            }
        }
    }
}

pub fn run_metrics_server(metrics_addr: SocketAddr, registry: Registry) -> JoinHandle<()> {
    tokio::spawn(metrics_runner(metrics_addr, registry))
}

fn can_recover(error_kind: ErrorKind) -> bool {
    match error_kind {
        // reconnect
        ErrorKind::ConnectionRefused
        | ErrorKind::ConnectionReset
        | ErrorKind::ConnectionAborted
        | ErrorKind::NotConnected
        | ErrorKind::BrokenPipe
        | ErrorKind::TimedOut
        | ErrorKind::Interrupted => true,
        // exit
        ErrorKind::AddrInUse
        | ErrorKind::AddrNotAvailable
        | ErrorKind::AlreadyExists
        | ErrorKind::WouldBlock
        | ErrorKind::InvalidInput
        | ErrorKind::InvalidData
        | _ => false,
    }
}
