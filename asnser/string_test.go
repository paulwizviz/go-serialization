package asnser

import (
	"encoding/asn1"
	"fmt"
	"unicode/utf8"
)

func Example_bitstring() {
	value := []byte("Hello world")
	input := asn1.BitString{
		Bytes:     value,
		BitLength: len(value) * 8, // Length in bits
	}

	enc, _ := asn1.Marshal(input)
	fmt.Printf("Tag: %x(%v) Value: %x\n", asn1.TagBitString, enc[0], enc)

	var bstring asn1.BitString
	rest, err := asn1.Unmarshal(enc, &bstring)
	fmt.Printf("Unmarshal: %s Rest: %v, Error: %v", bstring.Bytes, rest, err)

	// Output:
	// Tag: 3(3) Value: 030c0048656c6c6f20776f726c64
	// Unmarshal: Hello world Rest: [], Error: <nil>
}

func Example_octetstring() {
	utf8Char := make([]byte, 3)
	utf8.EncodeRune(utf8Char, '世')
	enc, _ := asn1.Marshal(utf8Char)
	fmt.Printf("Tag: %X(%v) Value: %x\n", asn1.TagOctetString, enc[0], enc)

	var octetString []byte
	asn1.Unmarshal(enc, &octetString)
	fmt.Printf("Unmarshal: %s\n", octetString)

	s := []byte("Hello world")
	enc, _ = asn1.Marshal(s)
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagOctetString, enc[0], enc)

	asn1.Unmarshal(enc, &octetString)
	fmt.Printf("Unmarshal: %s\n", octetString)

	// Output:
	// Tag: 4(4) Value: 0403e4b896
	// Unmarshal: 世
	// Tag: 4(4) Value: 040b48656c6c6f20776f726c64
	// Unmarshal: Hello world
}

func Example_utf8string() {
	enc, _ := asn1.Marshal("Hello world! 你好世界")
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagUTF8String, enc[0], enc)

	var utf8string string
	asn1.Unmarshal(enc, &utf8string)
	fmt.Printf("Unmarshal: %s", utf8string)

	// Output:
	// Tag: 12(12) Value: 0c1948656c6c6f20776f726c642120e4bda0e5a5bde4b896e7958c
	// Unmarshal: Hello world! 你好世界
}

func Example_printablestring() {
	enc, _ := asn1.Marshal("Ola world")
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagPrintableString, enc[0], enc)

	var printable string
	asn1.Unmarshal(enc, &printable)
	fmt.Printf("Unmarshal: %v", printable)

	// Output:
	// Tag: 19(19) Value: 13094f6c6120776f726c64
	// Unmarshal: Ola world
}
