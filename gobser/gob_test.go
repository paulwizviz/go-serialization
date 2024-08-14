// GOB is a golang specific streams.
// See https://go.dev/blog/gob
package gobser

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Model
type Sender struct {
	X, Y, z int
}

type SenderP struct {
	X, Y, z *int
}

type Receiver struct {
	X, Y, Z *int8
}

// Examples

func Example_simpleStruct() {
	s := Sender{
		X: 10,
		Y: 8,
		z: 1,
	}

	var network bytes.Buffer // Stand-in for network
	encoder := gob.NewEncoder(&network)
	decoder := gob.NewDecoder(&network)

	err := encoder.Encode(s)
	fmt.Println(err)

	var r Receiver
	err = decoder.Decode(&r)

	fmt.Printf("Receiver: {%d,%d,%d}\n", *r.X, *r.Y, r.Z)

	// Output:
	// <nil>
	// Receiver: {10,8,0}

}

func Example_structPtr() {
	x := 2
	xPtr := &x
	y := 4
	yPtr := &y
	sp := SenderP{
		X: xPtr,
		Y: yPtr,
	}

	var network bytes.Buffer // Stand-in for network
	encoder := gob.NewEncoder(&network)
	decoder := gob.NewDecoder(&network)

	err := encoder.Encode(sp)
	fmt.Println(err)

	var r Receiver
	err = decoder.Decode(&r)
	fmt.Printf("Receiver: {%d,%d}\n", *r.X, *r.Y)

	// Output:
	// <nil>
	// Receiver: {2,4}

}
