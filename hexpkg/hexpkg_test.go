package hexpkg

import (
	"encoding/hex"
	"fmt"
)

func Example_manualEncoding() {
	src := []byte("A") // Unicode Hex(41)
	dest := fmt.Sprintf("%X", src)
	fmt.Println(dest)

	// Output:
	// 41
}

func Example_hexPackage() {
	src := []byte("A")
	dest := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dest, src)
	fmt.Println(string(dest))

	// Output:
	// 41
}
