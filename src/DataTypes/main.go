package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Boolean
	fmt.Println("Boolean");
	var isRegistered bool = true
	var isEmailConfirmed bool

	var hasCompleteAccess bool = isRegistered && isEmailConfirmed
	var hasPartialAccess bool = isRegistered || isEmailConfirmed

	fmt.Printf("Is Registered - %v => %T\n", isRegistered, isRegistered)
	fmt.Printf("Is Email Confirmed - %v => %T\n", isEmailConfirmed, isEmailConfirmed)

	fmt.Printf("Complete Access - %v => %T\n", hasCompleteAccess, hasCompleteAccess)
	fmt.Printf("Partial Access - %v => %T\n", hasPartialAccess, hasPartialAccess)

	
	// Numerics
	
	// Signed Integers
	fmt.Println();
	fmt.Println("Signed Integers");
	var number int = -2021
	fmt.Printf("%v => %T => %d bytes\n", number, number, unsafe.Sizeof(number))

	// Unsigned Integers
	fmt.Println();
	fmt.Println("Unsigned Integers");
	var elements uint32 = 900 
	fmt.Printf("%v => %T\n", elements, elements)

	// Floating Point
	fmt.Println();
	fmt.Println("Floating Points");
	floatingPoint := 1234.567
	fmt.Printf("%v => %T\n", floatingPoint, floatingPoint)

	// Complex Number
	fmt.Println();
	fmt.Println("Complex Numbers");
	var complexValue = 2 + 5i
	fmt.Printf("%v => %T\n", complexValue, complexValue)
	// Using the complex function
	var value1 = 1.23
	var value2 = 4.56
	var complexResult = complex(value1, value2)
	fmt.Printf("%v => %T\n", complexResult, complexResult)

	// Byte 
	fmt.Println();
	fmt.Println("Byte");
	var byteLetter byte = 'A'
	fmt.Printf("%v => %T => %d bytes\n", byteLetter, byteLetter, unsafe.Sizeof(byteLetter))
	
	// Rune
	fmt.Println();
	fmt.Println("Rune");
	var runeLetter rune = 'µ'
	fmt.Printf("%v => %U => %T => %d bytes\n", runeLetter, runeLetter, runeLetter, unsafe.Sizeof(runeLetter))

	// Strings
	fmt.Println();
	fmt.Println("Strings");
	var name = "Jorge"
	var surname = `Serrano`
	fmt.Printf("%v => %T\n", name, name)
	fmt.Printf("%v => %T\n", surname, surname)
}