package dleq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ed25519_Secp256k1_DLEq(t *testing.T) {
	proof, _, err := ed25519_secp256k1_prove()
	require.NoError(t, err)
	err = ed25519_secp256k1_verify(proof)
	require.NoError(t, err)
}
