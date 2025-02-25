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

func (d *human) MarshalCBOR() ([]byte, error) {

	bday := d.birthday.Format(time.RFC3339)

	return cbor.Marshal(&struct {
		Firstname string
		Surname   string
		Birthday  string
		Address   address
	}{
		Firstname: d.firstname,
		Surname:   d.surname,
		Birthday:  bday,
		Address:   d.address,
	})
}

func (d *human) UnmarshalCBOR(data []byte) error {
	var aux struct {
		Firstname string
		Surname   string
		Birthday  string
		Address   address
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
	d.address = aux.Address
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

func Example_unmarshalInterfaceCBOR() {

	bday := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.Local).UTC()

	h := NewHuman("john", "doe", bday, address{1, "road", "SG"})
	b, err := cbor.Marshal(h)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("1. %x\n", b)

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
	// 1. a46946697273746e616d65646a6f686e675375726e616d6563646f6568426972746864617974323032352d30332d33315432333a30303a30305a6741646472657373a368486f7573654e756d0164526f616464726f616467436f756e747279625347
	// 2.1 john
	// 2.2 doe
	// 2.3 2025-03-31T23:00:00Z
	// 2.4 {1 road SG}
}
