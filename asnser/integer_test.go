package asnser

import (
	"encoding/asn1"
	"fmt"
)

func Example_int() {
	enc, _ := asn1.Marshal(10)
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagInteger, enc[0], enc)

	var i int
	asn1.Unmarshal(enc, &i)
	fmt.Printf("Unmarshal: %v\n", i)

	// Output:
	// Tag: 2(2) Value: 02010a
	// Unmarshal: 10
}

func Example_rune() {
	enc, _ := asn1.Marshal('a')
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagInteger, enc[0], enc)

	var i rune
	asn1.Unmarshal(enc, &i)
	if i == 'a' {
		fmt.Printf("Unmarshal: %v\n", i)
	}

	// Output:
	// Tag: 2(2) Value: 020161
	// Unmarshal: 97
}
