package crypto

//Crypto defines an interface
type Crypto interface {
	MakeNewKey(compress ...bool) SecretKey
	GetSecretKeyFromString(s string) SecretKey
	GetPublicKeyFromBytes(p []byte) PublicKey
	GetAddressFromString(s string) (Address, error)
}

//SecretKey defines the interface of secret Key
type SecretKey interface {
	PublicKey() PublicKey
	String() string
	IsCompressed() bool
	Bytes() []byte
	Sign(bs []byte) ([]byte, error)
}

//PublicKey defines the interface of public key, from secret key
type PublicKey interface {
	Address() Address
	HexStr() string
	Bytes() []byte
	Verify(sig, hash []byte) error
}

//Address defines the interface of address, from public key
type Address interface {
	String() string
	Bytes() []byte
	Verify(sig, hash []byte) error
	Equal(hs []byte) bool
}
