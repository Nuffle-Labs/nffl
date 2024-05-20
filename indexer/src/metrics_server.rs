use actix_web::{web, App, HttpResponse, HttpServer};
use prometheus::{Encoder, Registry, TextEncoder};
use std::{io::ErrorKind, net::SocketAddr, time::Duration};
use tracing::{error, info};

use crate::errors::Result;

const METRICS: &str = "metrics";

enum ServerState {
    WaitingForConnection,
    WaitingForReconnection,
    Shutdown,
    Exit(std::io::Error),
}

pub struct MetricsServer {
    metrics_addr: SocketAddr,
    registry: Registry,
    next_step: ServerState,
}

impl MetricsServer {
    pub fn new(metrics_addr: SocketAddr, registry: Registry) -> Self {
        Self {
            metrics_addr,
            registry,
            next_step: ServerState::WaitingForConnection,
        }
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

    async fn reconnect(&self) -> ServerState {
        let create_server = |registry: Registry| {
            let registry_data = web::Data::new(registry);
            let metrics_server = HttpServer::new(move || {
                App::new()
                    .app_data(registry_data.clone())
                    .service(web::resource("/metrics").route(web::get().to(Self::metrics)))
            })
            .bind(self.metrics_addr)?;

            Ok::<_, std::io::Error>(metrics_server)
        };

        let metrics_server = match create_server(self.registry.clone()) {
            Ok(v) => v,
            Err(err) => {
                if !can_recover(err.kind()) {
                    error!(target: METRICS, "Couldn't create server with code: {}", err.kind());
                    return ServerState::Exit(err);
                }

                return ServerState::WaitingForReconnection;
            }
        };

        // TODO: actix::spawn this and introduce ServerState::Active?
        match metrics_server.run().await {
            Ok(()) => {
                // Graceful shutdown
                info!(target: METRICS, "Server shutdown gracefully");
                ServerState::Shutdown
            }
            Err(err) => {
                if !can_recover(err.kind()) {
                    error!(target: METRICS, "Error while running server: {}", err.kind());
                    return ServerState::Exit(err);
                }

                ServerState::WaitingForReconnection
            }
        }
    }

    pub async fn run(mut self) -> Result<()> {
        const RECONNECTION_RETRIES: u32 = 3;
        const RECONNECTION_INTERVAL: Duration = Duration::from_secs(1);

        let mut retries = 0;
        loop {
            self.next_step = match self.next_step {
                ServerState::WaitingForConnection => self.reconnect().await,
                ServerState::Exit(err) => return Err(err.into()),
                ServerState::Shutdown => return Ok(()),
                ServerState::WaitingForReconnection => {
                    if retries == RECONNECTION_RETRIES {
                        info!(target: METRICS, "Reconnection attempts limit reached");
                        ServerState::Shutdown
                    } else {
                        info!(target: METRICS, "Reconnecting server");
                        retries += 1;
                        tokio::time::sleep(RECONNECTION_INTERVAL).await;
                        self.reconnect().await
                    }
                }
            };
        }
    }
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
