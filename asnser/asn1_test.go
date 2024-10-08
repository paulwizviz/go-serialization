package asnser

import (
	"encoding/asn1"
	"fmt"
	"time"
)

// The following examples demonstrate the marshaling/unmarshaling of
// primitive types

// NOTE: asn1.Marshal encode to DER format

func Example_null() {

	enc, _ := asn1.Marshal(asn1.NullRawValue)
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagNull, enc[0], enc)

	var rawVal asn1.RawValue
	asn1.Unmarshal(enc, &rawVal)

	if rawVal.Tag == int(asn1.TagNull) {
		fmt.Println("Null", rawVal)
	}

	// Output:
	// Tag: 5(5) Value: 0500
	// Null {0 5 false [] [5 0]}
}

func Example_oid() {
	var oid asn1.ObjectIdentifier = []int{1, 23, 4}
	enc, _ := asn1.Marshal(oid)
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagOID, enc[0], enc)

	var oid1 asn1.ObjectIdentifier
	asn1.Unmarshal(enc, &oid1)
	fmt.Printf("Unmarshal: %v", oid1)

	// Output:
	// Tag: 6(6) Value: 06023f04
	// Unmarshal: 1.23.4
}

func Example_enum() {
	var enum asn1.Enumerated = 1
	enc, _ := asn1.Marshal(enum)
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagEnum, enc[0], enc)

	var enum1 asn1.Enumerated
	asn1.Unmarshal(enc, &enum1)
	fmt.Printf("Unmarshal: %v", enum1)

	// Output:
	// Tag: 10(10) Value: 0a0101
	// Unmarshal: 1
}

func Example_utctime() {
	now := time.Date(2022, time.March, 1, 1, 1, 1, 1, time.Local)
	enc, _ := asn1.Marshal(now)
	fmt.Printf("Tag: %v(%v) Value: %x\n", asn1.TagUTCTime, enc[0], enc)

	var tm time.Time
	asn1.Unmarshal(enc, &tm)
	fmt.Printf("Time: %v", tm)

	// Output:
	// Tag: 23(23) Value: 170d3232303330313031303130315a
	// Time: 2022-03-01 01:01:01 +0000 UTC
}

// The following examples demonstrate the marshaling/unmarshaling of ASN.1 SEQUENCE/SET type.

type inStruct struct {
	TrackingNumber int
	Question       string `asn1:"ia5"`
}

func Example_inStruct() {

	is := inStruct{
		TrackingNumber: 5,
		Question:       "Anybody there?",
	}

	enc, _ := asn1.Marshal(is)
	for _, ch := range enc {
		fmt.Printf("%x(%d)\n", ch, ch)
	}

	// Meaning of output:
	// 30 — type tag indicating SEQUENCE
	// 13 — length in octets of value that follows
	//  02 — type tag indicating INTEGER
	//  01 — length in octets of value that follows
	//	05 — value (5)
	//  16 — type tag indicating IA5String
	//	 (IA5 means the full 7-bit ISO 646 set, including variants,
	//	  but is generally US-ASCII)
	//  0e — length in octets of value that follows
	//	41 6e 79 62 6f 64 79 20 74 68 65 72 65 3f — value ("Anybody there?")

	// Output:
	// 30(48)
	// 13(19)
	// 2(2)
	// 1(1)
	// 5(5)
	// 16(22)
	// e(14)
	// 41(65)
	// 6e(110)
	// 79(121)
	// 62(98)
	// 6f(111)
	// 64(100)
	// 79(121)
	// 20(32)
	// 74(116)
	// 68(104)
	// 65(101)
	// 72(114)
	// 65(101)
	// 3f(63)
}
