package crypto

import (
	"sync"
)

var (
	cryptoInstances = make(map[string]Crypto)
)

var mutex sync.Mutex

//RegisterCrypto func
func RegisterCrypto(n string, c Crypto) {
	mutex.Lock()
	defer mutex.Unlock()
	if c == nil {
		panic("crypto: RegisterCrypto instance is nil")
	}
	if _, dup := cryptoInstances[n]; dup {
		panic("crypto: RegisterCrypto called twice for instance " + n)
	}
	cryptoInstances[n] = c
}

//NewCrypto get crypto instance by n
func NewCrypto(n string) (c Crypto) {
	mutex.Lock()
	defer mutex.Unlock()
	c, ok := cryptoInstances[n]
	if !ok {
		panic("unknown instance " + n)
	}
	return c
}
