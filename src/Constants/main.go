package main

import "fmt"

const Department string = "Human Resources"
const city string = "Madrid"

func main() {
	fmt.Printf("%v - %T\n", Department, Department)
	fmt.Printf("%v - %T\n", city, city)

	fmt.Println();

	const name = "Jorge"
	const country string = "Spain"
	fmt.Printf("%v - %T\n", name, name)
	fmt.Printf("%v - %T\n", country, country)

	fmt.Println();

	const one, two, three = "1", 2, true
	fmt.Printf("%v - %T\n", one, one)
	fmt.Printf("%v - %T\n", two, two)
	fmt.Printf("%t - %T\n", three, three)

	fmt.Println();

	const (
		four = "Four"
		five = 5.1
	)
	fmt.Printf("%v - %T\n", four, four)
	fmt.Printf("%v - %T\n", five, five)

	fmt.Println();

	const value1 = 1
	const value2 = 2
	const value3 = 3
	const value4 int32 = 4

	result1 := value1 + value2
	fmt.Printf("%v - %T\n", result1, result1)
	result2 := result1 + value3
	fmt.Printf("%v - %T\n", result2, result2)
	//result3 := result1 + value4
	//fmt.Printf("%v - %T\n", result3, result3)
}