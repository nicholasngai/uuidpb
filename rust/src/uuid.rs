#[allow(clippy::derive_partial_eq_without_eq)]
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Uuid {
    #[prost(fixed64, tag = "1")]
    pub hi: u64,
    #[prost(fixed64, tag = "2")]
    pub lo: u64,
}
