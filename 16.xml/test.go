package main

import (
	"encoding/xml"
	"fmt"
)

/*
type customer struct {
	FirstName string
	LastName  string
	//Zipcode   string   `json:"-"`
	Zipcode  string
	Products []string
}
*/

type customer struct {
	XMLName   xml.Name `xml:"customer"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastName"`
	Zipcode   string   `xml:"zipcode,attr"`
	Products  []string `xml:"products"`
}

func main() {
	t2()
}

func t2() { //decode
	xmldata := `
	<customer zipcode="12341">
		<firstname></firstname>
		<lastName></lastName>
		<products>AA01</products>
		<products>AA02</products>
	</customer>
	`
	var c customer
	xml.Unmarshal([]byte(xmldata), &c)
	fmt.Printf("%#v\n", c)
}

func t1() { //encode
	customers := []customer{
		{FirstName: "Peter", LastName: "Kim", Zipcode: "12345", Products: nil},
		{FirstName: "Juli", LastName: "Lala", Zipcode: "12346", Products: []string{"Fan01", "TV01"}},
	}
	r, err := xml.MarshalIndent(customers, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %s", xml.Header, r)
}
