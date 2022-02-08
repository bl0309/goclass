package examplejson

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Age       int                    `json:"age"`
	Married   bool                   `json:"married"`
	Address   map[string]interface{} `json:"address"`
	Phone     []map[string]string    `json:"phone_number"`
}

type response struct {
	Page int      `json:"page"`
	Data []Person `json:"data"`
}

func InitExampleJson() {

	// key, value
	// json object
	// json array
	//data type:= number, string, bool

	ex1, _ := json.Marshal(true)
	fmt.Println(string(ex1))

	ex2, _ := json.Marshal(1)
	fmt.Println(string(ex2))

	ex3, _ := json.Marshal("Golang")
	fmt.Println(string(ex3))

	exD := []string{"a", "b", "c", "d"}
	ex4, _ := json.Marshal(exD)
	fmt.Println(string(ex4))

	exD2 := map[string]int{"a": 653, "b": 34, "c": 45, "d": 5325}
	ex5, _ := json.Marshal(exD2)
	fmt.Println(string(ex5))

	exD3 := map[string]interface{}{"first_name": "Bikram", "last_name": "Lamsal", "age": 26, "married": false}
	ex6, _ := json.Marshal(exD3)
	fmt.Println(string(ex6))

	resnd1 := &response{
		Page: 1,
		Data: []Person{
			{
				FirstName: "Bikram",
				LastName:  "Lamsal",
				Age:       26,
				Married:   false,
				Address: map[string]interface{}{
					"Street": "3453 Creep Rd",
					"City":   "Forthworth",
					"County": "Dallas",
					"zip":    75073,
				},
				Phone: []map[string]string{
					{"home": "8376487639"},
					{"office": "2345634577"},
				},
			},
		},
	}
	fmt.Printf("resnd1=%v\n", resnd1)
	ex7, _ := json.Marshal(resnd1)
	fmt.Println(string(ex7))

	exJson := []byte(`{"a":1,"b":"abc","c":[10,11]}`)

	var ex8 map[string]interface{}

	err := json.Unmarshal(exJson, &ex8)
	if err != nil {
		panic((err))
	}
	fmt.Printf("%+v", ex8)
	// for _, value := ex8.Data {
	// 	if value.Age > 15 {
	// 		fmt.Printf("%s \n", value.FirstName, value.LastName)
	// 	}
	// }
}
