package main

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//Zipcode   string   `json:"-"`
	Zipcode  string   `json:"zipcode"`
	Products []string `json:"products,omitempty"`
}

func main() {
	t2()
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
		var c customer
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
	customers := []customer{
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
