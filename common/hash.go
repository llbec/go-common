package common

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

const (
	//HashLength is the expected length of the hash
	HashLength = 32
)

//Hash represents the 32 byte Keccak256 hash of arbitrary data.
type Hash [HashLength]byte

//BytesToHash sets b to hash.
// If b is larger than len(h), b will be cropped from the left.
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

//HexToHash hex -> hash
func HexToHash(s string) Hash {
	b, err := FromHex(s)
	if err != nil {
		panic(err)
	}
	return BytesToHash(b)
}

//Bytes Get the []byte representation of the underlying hash
func (h Hash) Bytes() []byte { return h[:] }

//SetBytes sets the hash to the value of b.
// If b is larger than len(h), b will be cropped from the left.
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

//HashHex []byte -> hex
func HashHex(d []byte) string {
	var buf [HashLength * 2]byte
	hex.Encode(buf[:], d)
	return string(buf[:])
}

//IsHex 是否是hex字符串
func IsHex(str string) bool {
	l := len(str)
	return l >= 4 && l%2 == 0 && str[0:2] == "0x"
}

/*===================================================*/

//Sha256 encryption
func Sha256(b []byte) []byte {
	data := sha256.Sum256(b)
	return data[:]
}

//ShaHash write hash(sha256) in out []byte
func ShaHash(b []byte, out []byte) {
	s := sha256.New()
	s.Write(b[:])
	tmp := s.Sum(nil)
	s.Reset()
	s.Write(tmp)
	copy(out[:], s.Sum(nil))
}

// Sha2Sum Returns SHA256 hash: SHA256( SHA256( data ) )
// Where possible, using ShaHash() should be a bit faster
func Sha2Sum(b []byte) []byte {
	tmp := sha256.Sum256(b)
	tmp = sha256.Sum256(tmp[:])
	return tmp[:]
}

//RimpHash func
func RimpHash(in []byte, out []byte) {
	sha := sha256.New()
	_, err := sha.Write(in)
	if err != nil {
		return
	}
	rim := ripemd160.New()
	_, err = rim.Write(sha.Sum(nil)[:])
	if err != nil {
		return
	}
	copy(out, rim.Sum(nil))
}

// Rimp160 Returns hash: RIMP160( SHA256( data ) )
// Where possible, using RimpHash() should be a bit faster
func Rimp160(b []byte) []byte {
	out := make([]byte, 20)
	RimpHash(b, out[:])
	return out[:]
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Keccak256Hash calculates and returns the Keccak256 hash of the input data,
// converting it to an internal Hash data structure.
func Keccak256Hash(data ...[]byte) (h Hash) {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:0])
	return h
}

// Keccak512 calculates and returns the Keccak512 hash of the input data.
func Keccak512(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak512()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
