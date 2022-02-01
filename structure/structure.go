package structure

import "fmt"

type person struct {
	name    string
	age     uint8 // can hold for 100 and memory usability
	address map[string]string
	married bool
}

func Init() {
	p := person{
		name:    "Bikram",
		age:     26,
		address: map[string]string{"USA": "Irving"},
		married: false,
	}
	fmt.Printf("%+v\n", p)

}
