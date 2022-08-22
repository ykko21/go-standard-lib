package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//Zipcode   string   `json:"-"`
	Zipcode  string   `json:"zipcode"`
	Products []string `json:"products,omitempty"`
}

func main() {
	t3()
}

func t3() { //Read JSON,
	data := `
	[
		{
			"first_name":"Kal",
			"last_name": "Kakao",
			"zipcode": "10100",
			"products": ["Radio01","Pencel01"]
		},
		{
			"first_name":"Peter",
			"last_name": "Skpe",
			"zipcode": "20103",
			"products": ["Cellphone1","keyboard01"]
		}				
	]
	`
	data_byte := []byte(data)
	valid := json.Valid(data_byte)
	if valid {
		var custs []Customer
		err := json.Unmarshal(data_byte, &custs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(custs)
	}
}

func t2() { //decode
	data := []byte(`
	 	{
			"first_name":"Kal",
			"last_name": "Kakao",
			"zipcode": "10100",
			"products": ["Radio01","Pencel01"]
		}
	 `)

	// Convert to Json
	valid := json.Valid(data)
	if valid {
		var c Customer
		json.Unmarshal(data, &c)
		fmt.Printf("%v\n", c)
		fmt.Printf("%#v\n", c)
	}
	fmt.Println()
	//Convert to Map
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)
	for k, v := range m {
		fmt.Println(k, v)
	}

}
func t1() { //encode
	customers := []Customer{
		{"Peter", "Kim", "12345", nil},
		{"Juli", "Lala", "12346", []string{"Fan01", "TV01"}},
	}
	//r, err := json.Marshal(customers)
	r, err := json.MarshalIndent(customers, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", r)
}
