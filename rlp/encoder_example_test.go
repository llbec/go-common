package rlp

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type MyCoolType struct {
	Name string
	A, B uint
}

// EncodeRLP writes x as RLP list [a, b] that omits the Name field.
func (x *MyCoolType) EncodeRLP(w io.Writer) (err error) {
	// Note: the receiver can be a nil pointer. This allows you to
	// control the encoding of nil, but it also means that you have to
	// check for a nil receiver.
	if x == nil {
		err = Encode(w, []uint{0, 0})
	} else {
		err = Encode(w, []uint{x.A, x.B})
	}
	return err
	/*if x == nil {
		err = Encode(w, "")
		if err != nil {
			return err
		}
		return Encode(w, []uint{0, 0})
	}
	err = Encode(w, x.Name)
	if err != nil {
		return err
	}
	return Encode(w, []uint{x.A, x.B})*/
}

// DecodeRLP writes x as RLP list [a, b] that omits the Name field.
func (x *MyCoolType) DecodeRLP(s *Stream) (err error) {
	var a []uint
	if err = s.Decode(&a); err != nil {
		return
	}
	if len(a) != 2 {
		return fmt.Errorf("Invalid elements number(%d)", len(a))
	}
	x.A = a[0]
	x.B = a[1]
	return
}

func decodefunc(input []byte) error {
	var t MyCoolType
	err := Decode(bytes.NewReader(input), &t)
	//err := DecodeBytes(input, &s)
	if err != nil {
		return fmt.Errorf("Decode Error: %v (%v)", err, input)
	}
	fmt.Printf("Decoded value: %#v\n", t)
	return nil
}

func exampleEncoder() error {
	var t *MyCoolType // t is nil pointer to MyCoolType
	bytes, _ := EncodeToBytes(t)
	fmt.Printf("%v → %X(%v)\n", t, bytes, bytes)

	t = &MyCoolType{Name: "foobar", A: 5, B: 6}
	bytes, _ = EncodeToBytes(t)
	fmt.Printf("%v → %X(%v)\n", t, bytes, bytes)

	// Output:
	// <nil> → C28080
	// &{foobar 5 6} → C20506

	return decodefunc(bytes)
}

func TestExample(t *testing.T) {
	e := exampleEncoder()
	if e != nil {
		t.Error(e)
	}
}
