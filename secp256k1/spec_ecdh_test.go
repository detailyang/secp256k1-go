package secp256k1

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpecEcdh(t *testing.T) {
	ctx, err := ContextCreate(ContextSign | ContextVerify)
	if err != nil {
		panic(err)
	}

	alice := []byte{0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41, 0x41}
	bob := []byte{0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40}
	expected, _ := hex.DecodeString("238c14f420887f8e9bfa78bc9bdded1975f0bb6384e33b4ebbf7a8c776844aec")

	r, Alice, err := EcPubkeyCreate(ctx, alice)
	spOK(t, r, err)

	r, Bob, err := EcPubkeyCreate(ctx, bob)
	spOK(t, r, err)

	// Test case: a*B == A*b == expected
	r, bobSecret, err := Ecdh(ctx, Alice, bob)
	spOK(t, r, err)

	r, aliceSecret, err := Ecdh(ctx, Bob, alice)
	spOK(t, r, err)

	assert.Equal(t, expected, aliceSecret)
	assert.Equal(t, expected, bobSecret)
}
