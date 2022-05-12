package dleq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ed25519_Secp256k1_DLEq(t *testing.T) {
	proof, _, err := Ed25519Secp256k1Prove()
	require.NoError(t, err)
	err = Ed25519Secp256k1Verify(proof)
	require.NoError(t, err)
}

func Test_Ed25519_Secp256k1_DLEq_fail(t *testing.T) {
	proof, _, err := Ed25519Secp256k1Prove()
	require.NoError(t, err)
	proof[1] = 0
	err = Ed25519Secp256k1Verify(proof)
	require.NotNil(t, err)
}
