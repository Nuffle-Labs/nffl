use mockall::automock;
use crate::safeclient::SafeClient;

#[automock]
pub trait MockSafeClient: SafeClient {}