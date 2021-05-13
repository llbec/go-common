package gocommon

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"
)

func TestHash(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	k1 := make([]byte, 32)
	ShaHash(key, k1)
	k2 := Sha256(key)
	k3 := Sha2Sum(key)

	fmt.Printf("Original: %v\nShaHash: %v\nSha256 : %v\nSha2Sum: %v\n",
		HashHex(key), HashHex(k1), HashHex(k2), HashHex(k3))

	if !reflect.DeepEqual(k1, k3) {
		t.Errorf("sha hash256 deep equal false!")
	}

	h1 := make([]byte, 20)
	RimpHash(k1, h1)
	h2 := Rimp160(k1)

	if !reflect.DeepEqual(h1, h2) {
		t.Errorf("rimp hash160 deep equal false!")
	}
}
