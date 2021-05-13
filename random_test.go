package gocommon

import (
	"reflect"
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 1; i < 100; i++ {
		r1 := RandBytes(i)
		r2 := RandBytes(i)
		if reflect.DeepEqual(r1, r2) {
			t.Errorf("Random failed! (%v)==(%v)", ToHex(r1), ToHex(r2))
		}
	}
}

func TestRandomHex(t *testing.T) {
	for i := 2; i < 100; i++ {
		r1 := RandHex(i)
		r2 := RandHex(i)
		if r1 == r2 {
			t.Errorf("Random failed! (%v)==(%v)", r1, r2)
		}
	}
}
