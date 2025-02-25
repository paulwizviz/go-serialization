package cborser

import (
	"fmt"
	"time"

	"github.com/fxamacker/cbor/v2"
)

func Example_int() {

	for i, v := range []int{1, 2, 10, 20, 23, 24, 70, 100} {
		b, _ := cbor.Marshal(v)
		fmt.Printf("%d. %v %x\n", i, v, b)
	}

	// Output:
	// 0. 1 01
	// 1. 2 02
	// 2. 10 0a
	// 3. 20 14
	// 4. 23 17
	// 5. 24 1818
	// 6. 70 1846
	// 7. 100 1864
}

func Example_float() {

	for i, v := range []float32{0, 1.1, 10.3, 20.5, 23.7, 71.4, 100.9} {
		b, _ := cbor.Marshal(v)
		fmt.Printf("%d. %v %x\n", i, v, b)
	}

	// Output:
	// 0. 0 fa00000000
	// 1. 1.1 fa3f8ccccd
	// 2. 10.3 fa4124cccd
	// 3. 20.5 fa41a40000
	// 4. 23.7 fa41bd999a
	// 5. 71.4 fa428ecccd
	// 6. 100.9 fa42c9cccd
}

func Example_string() {

	for i, v := range []string{"hello", "world"} {
		b, _ := cbor.Marshal(v)
		fmt.Printf("%d. %v %x\n", i, v, b)
	}

	// Output:
	// 0. hello 6568656c6c6f
	// 1. world 65776f726c64

}

func Example_slice() {
	s := []string{"hello", "world"}
	b, _ := cbor.Marshal(s)
	fmt.Printf("%v %x", s, b)

	// Output:
	// [hello world] 826568656c6c6f65776f726c64
}

func Example_map() {
	m := make(map[string]int)
	m["abc"] = 1
	m["efg"] = 2
	m["hij"] = 1
	b, _ := cbor.Marshal(m)
	fmt.Printf("%v %x\n", m, b)

	m1 := make(map[string]float32)
	m1["abc"] = 1.0
	m1["efg"] = 2.3
	m1["hij"] = 1.7
	b, _ = cbor.Marshal(m1)
	fmt.Printf("%v %x\n", m1, b)

	// Output:
	// a3636162630163656667026368696a01
	// a363616263fa3f80000063656667fa401333336368696afa3fd9999a
}

func Example_time() {

	tm := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.Local).Format(time.RFC3339)
	fmt.Println(tm)
	b, _ := cbor.Marshal(tm)
	fmt.Printf("%X\n", b)

	var tm1 time.Time
	cbor.Unmarshal(b, &tm1)
	fmt.Println(tm1)

	// Output:
	// 2025-04-01T00:00:00+01:00
	// 7819323032352D30342D30315430303A30303A30302B30313A3030
	// 2025-04-01 00:00:00 +0100 BST
}

func Example_timeRFC3339() {

	tm, _ := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")

	// Encode time as string in RFC3339 format with second precision.
	em, err := cbor.EncOptions{Time: cbor.TimeRFC3339}.EncMode()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	b, err := em.Marshal(tm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("1: %s\n", b)

	var tms string
	err = cbor.Unmarshal(b, &tms)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("2: %s", tms)

	// Output:
	// 1: t2013-03-21T20:04:00Z
	// 2: 2013-03-21T20:04:00Z

}

func Example_timeUnix() {

	tm, _ := time.Parse(time.RFC3339, "2013-03-21T20:04:00Z")

	// Encode time as numerical representation of seconds since January 1, 1970 UTC.
	em, err := cbor.EncOptions{Time: cbor.TimeUnix}.EncMode()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	b, err := em.Marshal(tm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("1. %x\n", b)

	var i int64
	cbor.Unmarshal(b, &i)
	tmu := time.Unix(i, 0)
	fmt.Printf("2. %d %v\n", i, tmu.Format(time.RFC3339))

	// Output:
	// 1. 1a514b67b0
	// 2. 1363896240 2013-03-21T20:04:00Z
}
