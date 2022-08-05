# Berserker Network

Experimental blockchain network for transaction processing.

## Supported features:

- Persistent ledger (double-entry bookkeeping)
- Mempool blocks formation (inspired by Etherum)
- Multi-node p2p dynamic peering (via bootstrap nodes)
- Full and partial sync
- HTTP API interface

## Future considerations

Mempool is currently very simplified - thing to consider moving on ([inspirations](https://github.com/ethereum/go-ethereum/blob/7b32d2a47017570c44cd7f8a83612a29656c9857/core/tx_pool.go#L211)):
 - limit max size
 - workout what happens when full
 - processing fee dynamic/fixed (gas)
 - processing order FIFO/gas-priority/other
 - DDoS attack prevention

## Underway

Features I'm working on:

 - gRPC interface
 - Better transaction model
 - PoW
 - PoS
 - Gas fees
