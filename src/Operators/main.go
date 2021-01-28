package main

import "fmt"

func main() {
    // Boolean
	var isOk bool = true
	var isKo bool
	// Operator (&&)
	var boolOp1 = isOk && isKo
	fmt.Printf("%v => %T\n", boolOp1, boolOp1)
	// Operator (||)
	var boolOp2 = isOk || isKo
	fmt.Printf("%v => %T\n", boolOp2, boolOp2)

	fmt.Println()

	// Integers
	// Operator (+)
	var two, three = 2, 3
	var result = two + three
	fmt.Printf("%v => %T\n", result, result)
	// Increment and Decrement Operators (Increment by 1)
	three++
	fmt.Printf("%v => %T\n", three, three)
	// Assignment Operator (Decrement by 2)
	three -= 2
	fmt.Printf("%v => %T\n", three, three)
	// Comparison Operators
	var isGreaterAndEqualThan3 = three >= 3
	fmt.Printf("%v => %T\n", isGreaterAndEqualThan3, isGreaterAndEqualThan3)

	fmt.Println()

	// Complex Number
	var complexNumber1 = complex(1.23, 3.45)
	var complexNumber2 = complex(0.19, 0.51)
	// Operator (+)
	var complexNumber = complexNumber1 + complexNumber2
	fmt.Printf("%v => %T\n", complexNumber1, complexNumber1)
	fmt.Printf("%v => %T\n", complexNumber2, complexNumber2)
	fmt.Printf("%v => %T\n", complexNumber, complexNumber)
}