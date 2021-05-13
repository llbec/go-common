package gocommon

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_Big256(t *testing.T) {
	fmt.Println(Big1)
	fmt.Println(Big256)
	fmt.Println(big.NewInt(1000000000000000000))
	fmt.Println(big.NewInt(1e18))
	fmt.Println(big.NewInt(11e17))
	fmt.Println(Big256.Add(big.NewInt(1e18), big.NewInt(6e17)))
}
