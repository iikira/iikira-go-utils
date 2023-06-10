package jwted25519_test

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"testing"
)

func TestGenEd25519Key(t *testing.T) {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	fmt.Println(publicKey, privateKey)
}
