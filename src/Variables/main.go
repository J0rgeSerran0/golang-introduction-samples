package main

import "fmt"


var Nick string = "MyNick"

var avatar string = "MyAvatar"

func main() {
	var age = 21
	fmt.Printf("%v years old\n", age)

	department := "Administration"
	fmt.Printf("%v\n", department)

	var name string = "Jorge"
	var surname string 
	surname = "Serrano"
	fmt.Printf("%v %v\n", name, surname)
	
	var address, city, country string
	address = "Main Street 1"
	city = "Madrid"
	country = "Spain"
	//var address, city, country string = "Main Street 1", "Madrid", "Spain"
	fmt.Printf("%v - %v (%v)\n", address, city, country)

	var (
		teacher string = "John"
		class string = "6A"
	)
	fmt.Printf("%v (%v)\n", teacher, class)

	fmt.Printf("%v - %v\n", avatar, Nick)
}