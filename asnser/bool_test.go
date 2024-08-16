package asnser

import (
	"encoding/asn1"
	"fmt"
)

func Example_true() {
	t, _ := asn1.Marshal(true)
	fmt.Printf("Tag: %X(%v) Value: %x\n", asn1.TagBoolean, t[0], t)

	var b bool
	asn1.Unmarshal(t, &b)
	fmt.Printf("Unmarshal: %v\n", b)

	// Output:
	// Tag: 1(1) Value: 0101ff
	// Unmarshal: true
}

func Example_false() {

	f, _ := asn1.Marshal(false)
	fmt.Printf("Tag: %X(%v) Value: %x\n", asn1.TagBoolean, f[0], f)

	var b bool
	asn1.Unmarshal(f, &b)
	fmt.Printf("Unmarshal: %v\n", b)

	// Output:
	// Tag: 1(1) Value: 010100
	// Unmarshal: false
}
