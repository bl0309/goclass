package structure

import "fmt"

type base struct {
	des
	num int
}

type des interface {
	describe() string
}

type container struct {
	base
	// num int
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("Base num = %v", b.num)
}

func Init3() {
	con := container{
		base: base{
			num: 10,
		},
		str: "test",
	}
	fmt.Printf("Container = %v\n", con)

	fmt.Println(con.num, con.str)
	fmt.Println(con.base.num, con.str)

	fmt.Println(con.describe())
}
