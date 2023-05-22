package main

import (
	"fmt"
)

// pass by value
type Addres struct {
	City, Province, Country string
}

func main () {

	// pass by valu. add &addres1 for pass by reference
	// adress1 := Addres{"Sumedang","Jawa Barat", "Indonesia"}
	// adress2 := &adress1

	var adress1 Addres = Addres{"Sumedang","Jawa Barat", "Indonesia"}
	var adress2 *Addres = &adress1

	adress2.City = "Bandung"

	fmt.Println("Address1", adress1)
	fmt.Println("Address2", adress2)

}