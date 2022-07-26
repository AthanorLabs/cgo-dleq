build:
	git submodule update --init
	cd lib/c-dleq && cargo build --release
	cp lib/c-dleq/target/release/libdleq.so lib/

test: build
	go test
