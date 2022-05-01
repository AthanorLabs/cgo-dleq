package dleq

/*
#cgo LDFLAGS: -L${SRCDIR}/lib ./lib/libdleq.so
#include "./lib/libdleq.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	Ed25519PrivateKeySize   = 32
	Ed25519PublicKeySize    = 32
	Secp256k1PrivateKeySize = 32
	Secp256k1PublicKeySize  = 33
)

// Proof represents a serialized DLEq proof
type Proof []byte

// PrivateKey represents a 32-byte private key
type PrivateKey []byte

// Ed25519Secp256k1Prove creates a DLEq proof of the returned private key
// for ed25519 and secp256k1.
func Ed25519Secp256k1Prove() (Proof, PrivateKey, error) {
	// TOOD: malloc instread of using "make" cause of GC?
	proofSize := C.ed25519_secp256k1_proof_size()
	dst := make([]byte, proofSize+Ed25519PrivateKeySize)
	ptr := unsafe.Pointer(&dst[0])

	ok := C.ed25519_secp256k1_prove((*C.char)(ptr))
	if byte(ok) == 0 {
		return nil, nil, fmt.Errorf("failed to generate proof")
	}

	// first 32 bytes are private key, rest is proof
	return Proof(dst[32:]), PrivateKey(dst[:32]), nil
}

// Ed25519Secp256k1Verify verifies the given DLEq proof.
// It returns an error if verification fails.
func Ed25519Secp256k1Verify(proof []byte) error {
	dst := make([]byte, Ed25519PublicKeySize+Secp256k1PublicKeySize)
	dstPtr := unsafe.Pointer(&dst[0])

	proofPtr := unsafe.Pointer(&proof[0])

	ok := C.ed25519_secp256k1_verify((*C.char)(proofPtr), (*C.char)(dstPtr))
	if byte(ok) == 0 {
		return fmt.Errorf("failed to verify proof")
	}

	return nil
}
