package main

import "fmt"

func main() {
	var student = [4]string{"John", "Mary", "Rose", "Peter"}
	fmt.Println("Values: ", student)

	people := [4] string {"Nuria", "Charles", "Arthur", "Rose"}
	// people = append(people, "George")
	fmt.Println("Values: ", people)
	
	names := [] string {"Frank", "Steve", "Olga", "Peter"}
	fmt.Println("Values: ", names)
	fmt.Println("Cap: ", cap(names))
	names = append(names, "George")
	fmt.Println("Values: ", names)
	fmt.Println("Len: ", len(names))
	fmt.Println("Cap: ", cap(names))
}