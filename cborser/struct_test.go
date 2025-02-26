package cborser

import (
	"fmt"
	"time"

	"github.com/fxamacker/cbor/v2"
)

type address struct {
	HouseNum int
	Road     string
	Country  string
}

func Example_adddress() {
	a := address{
		HouseNum: 1,
		Road:     "road",
		Country:  "SG",
	}

	b, err := cbor.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%X\n", b)

	var a1 address
	err = cbor.Unmarshal(b, &a1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", a1)

	// Output:
	// A368486F7573654E756D0164526F616464726F616467436F756E747279625347
	// {1 road SG}
}

type Human interface {
	Firstname() string
	Surname() string
	Birthday() time.Time
	Address() address
}

type human struct {
	firstname string
	surname   string
	birthday  time.Time
	address
}

func (d human) Firstname() string {
	return d.firstname
}

func (d human) Surname() string {
	return d.surname
}

func (d human) Birthday() time.Time {
	return d.birthday
}

func (d human) Address() address {
	return d.address
}

type auxAddr struct {
	HouseNum int    `cbor:"housenum"`
	Road     string `cbor:"road"`
	Country  string `cbor:"country"`
}

func (d *human) MarshalCBOR() ([]byte, error) {

	bday := d.birthday.Format(time.RFC3339)

	return cbor.Marshal(&struct {
		Firstname string  `cbor:"firstname"`
		Surname   string  `cbor:"surname"`
		Birthday  string  `cbor:"birthday"`
		Address   auxAddr `cbor:"address"`
	}{
		Firstname: d.firstname,
		Surname:   d.surname,
		Birthday:  bday,
		Address: auxAddr{
			HouseNum: d.address.HouseNum,
			Road:     d.address.Road,
			Country:  d.address.Country,
		},
	})
}

func (d *human) UnmarshalCBOR(data []byte) error {
	var aux struct {
		Firstname string  `cbor:"firstname"`
		Surname   string  `cbor:"surname"`
		Birthday  string  `cbor:"birthday"`
		Address   auxAddr `cbor:"address"`
	}
	if err := cbor.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.firstname = aux.Firstname
	d.surname = aux.Surname
	bday, err := time.Parse(time.RFC3339, aux.Birthday)
	if err != nil {
		return err
	}
	d.birthday = bday
	d.address.HouseNum = aux.Address.HouseNum
	d.address.Road = aux.Address.Road
	d.address.Country = aux.Address.Country
	return nil
}

func NewHuman(firstname, surname string, birthday time.Time, addr address) Human {
	return &human{
		firstname: firstname,
		surname:   surname,
		birthday:  birthday,
		address:   addr,
	}
}

func EmptyHuman() Human {
	return &human{}
}

func Example_interfaceCBOR() {

	bday := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.Local).UTC()

	h := NewHuman("john", "doe", bday, address{1, "road", "SG"})
	b, err := cbor.Marshal(h)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("1. %v %x\n", h, b)

	h1 := EmptyHuman()
	err = cbor.Unmarshal(b, &h1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("2.1 %v\n", h1.Firstname())
	fmt.Printf("2.2 %v\n", h1.Surname())
	fmt.Printf("2.3 %v\n", h1.Birthday().Format(time.RFC3339))
	fmt.Printf("2.4 %v\n", h1.Address())

	// Output:
	// 1. &{john doe {0 63879058800 <nil>} {1 road SG}} a46966697273746e616d65646a6f686e677375726e616d6563646f6568626972746864617974323032352d30332d33315432333a30303a30305a6761646472657373a368686f7573656e756d0164726f616464726f616467636f756e747279625347
	// 2.1 john
	// 2.2 doe
	// 2.3 2025-03-31T23:00:00Z
	// 2.4 {1 road SG}
}
