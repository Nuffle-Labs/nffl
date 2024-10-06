use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use tokio::sync::mpsc;
use anyhow::{Result, anyhow};

use crate::types::BlockData;

type BlockPredicate = Arc<dyn Fn(&BlockData) -> bool + Send + Sync>;

struct SubscriberData {
    predicate: BlockPredicate,
    notifier_tx: mpsc::Sender<BlockData>,
}

pub struct Notifier {
    rollup_ids_to_subscribers: Arc<Mutex<HashMap<u32, Vec<SubscriberData>>>>,
}

impl Notifier {
    pub fn new() -> Self {
        Self {
            rollup_ids_to_subscribers: Arc::new(Mutex::new(HashMap::new())),
        }
    }

    pub fn subscribe(&self, rollup_id: u32, predicate: impl Fn(&BlockData) -> bool + Send + Sync + 'static) -> (mpsc::Receiver<BlockData>, usize) {
        let mut subscribers = self.rollup_ids_to_subscribers.lock().unwrap();
        let subscribers_for_rollup = subscribers.entry(rollup_id).or_insert_with(Vec::new);

        let (tx, rx) = mpsc::channel(100);
        let subscriber = SubscriberData {
            predicate: Arc::new(predicate),
            notifier_tx: tx,
        };

        let id = subscribers_for_rollup.len();
        subscribers_for_rollup.push(subscriber);

        (rx, id)
    }

    pub fn notify(&self, rollup_id: u32, block: BlockData) -> Result<()> {
        let subscribers = self.rollup_ids_to_subscribers.lock().unwrap();
        let subscribers_for_rollup = subscribers.get(&rollup_id).ok_or_else(|| anyhow!("Unknown rollup ID"))?;

        for subscriber in subscribers_for_rollup {
            if (subscriber.predicate)(&block) {
                if let Err(e) = subscriber.notifier_tx.try_send(block.clone()) {
                    eprintln!("Failed to send block to subscriber: {:?}", e);
                }
            }
        }

        Ok(())
    }

    pub fn unsubscribe(&self, rollup_id: u32, id: usize) {
        let mut subscribers = self.rollup_ids_to_subscribers.lock().unwrap();
        if let Some(subscribers_for_rollup) = subscribers.get_mut(&rollup_id) {
            subscribers_for_rollup.remove(id);
        }
    }
}

impl Clone for Notifier {
    fn clone(&self) -> Self {
        Self {
            rollup_ids_to_subscribers: self.rollup_ids_to_subscribers.clone(),
        }
    }
}