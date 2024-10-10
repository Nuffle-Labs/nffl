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

#[cfg(test)]
mod tests {
    use super::*;
    use std::sync::Arc;
    use tokio::runtime::Runtime;

    #[test]
    fn test_notifier_subscribe_and_notify() {
        let rt = Runtime::new().unwrap();
        rt.block_on(async {
            let notifier = Notifier::new();
            let rollup_id = 1;

            // Subscribe to notifications
            let (mut rx, id) = notifier.subscribe(rollup_id, move |block| block.rollup_id == rollup_id);

            // Create a test block
            let test_block = BlockData {
                rollup_id,
                block: Default::default(), // You might need to adjust this based on your Block type
                transaction_id: [0; 32],
                commitment: [0; 32],
            };

            // Notify the subscribers
            notifier.notify(rollup_id, test_block.clone()).unwrap();

            // Check if we received the notification
            let received_block = rx.recv().await.unwrap();
            assert_eq!(received_block.rollup_id, test_block.rollup_id);

            // Unsubscribe
            notifier.unsubscribe(rollup_id, id);
        });
    }

    #[test]
    fn test_notifier_multiple_subscribers() {
        let rt = Runtime::new().unwrap();
        rt.block_on(async {
            let notifier = Notifier::new();
            let rollup_id = 1;

            // Subscribe two listeners
            let (mut rx1, id1) = notifier.subscribe(rollup_id, move |block| block.rollup_id == rollup_id);
            let (mut rx2, id2) = notifier.subscribe(rollup_id, move |block| block.rollup_id == rollup_id);

            // Create a test block
            let test_block = BlockData {
                rollup_id,
                block: Default::default(),
                transaction_id: [0; 32],
                commitment: [0; 32],
            };

            // Notify the subscribers
            notifier.notify(rollup_id, test_block.clone()).unwrap();

            // Check if both subscribers received the notification
            let received_block1 = rx1.recv().await.unwrap();
            let received_block2 = rx2.recv().await.unwrap();
            assert_eq!(received_block1.rollup_id, test_block.rollup_id);
            assert_eq!(received_block2.rollup_id, test_block.rollup_id);

            // Unsubscribe
            notifier.unsubscribe(rollup_id, id1);
            notifier.unsubscribe(rollup_id, id2);
        });
    }

    #[test]
    fn test_notifier_predicate_filtering() {
        let rt = Runtime::new().unwrap();
        rt.block_on(async {
            let notifier = Notifier::new();
            let rollup_id = 1;

            // Subscribe with a specific predicate
            let (mut rx, id) = notifier.subscribe(rollup_id, move |block| block.rollup_id == rollup_id && block.transaction_id[0] == 1);

            // Create test blocks
            let matching_block = BlockData {
                rollup_id,
                block: Default::default(),
                transaction_id: [1; 32],
                commitment: [0; 32],
            };
            let non_matching_block = BlockData {
                rollup_id,
                block: Default::default(),
                transaction_id: [0; 32],
                commitment: [0; 32],
            };

            // Notify the subscribers
            notifier.notify(rollup_id, matching_block.clone()).unwrap();
            notifier.notify(rollup_id, non_matching_block.clone()).unwrap();

            // Check if only the matching block was received
            let received_block = rx.recv().await.unwrap();
            assert_eq!(received_block.transaction_id[0], 1);

            // Ensure no more blocks were received
            assert!(rx.try_recv().is_err());

            // Unsubscribe
            notifier.unsubscribe(rollup_id, id);
        });
    }
}