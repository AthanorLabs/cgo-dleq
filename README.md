# cgo-dleq

This repository contains C and Go bindings for the library [dleq-rs](https://github.com/kayabaNerve/dleq-rs/tree/develop), specifically for DLEq proofs between ed25519 and secp256k1 points.

## Note

The `libdleq.so` library was compiled using the `x86_64-unknown-linux-gnu` Rust host, so if you want to use this for another platform, you will probably need to rebuild the library using `make build`.
