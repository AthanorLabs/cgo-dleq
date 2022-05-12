build:
	git submodule update --init
	cd lib/dleq-rs && cargo build --release
	cp lib/dleq-rs/target/release/libdleq.so lib/

test: build
	go test
