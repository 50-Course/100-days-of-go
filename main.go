package main

import (
	"fmt"
)

func main() {
	const conferenceName = "GoCon Nigeria"

	var userName string
	var userAge uint32

	fmt.Println("Please enter username: ")
	fmt.Scan(&userName)

	fmt.Println("Age: ")
	fmt.Scan(&userAge)

	fmt.Printf("Welcome %v! This is the %v 2025 special edition. \n", userName, conferenceName)
	fmt.Println("This is going to be awesome. Welcome to my conference app")
	fmt.Println(doSomething(1, 2))

	learnArrays()
}

const (
	englishPrefix = "Hello"
	frenchPrefix  = "Bonjour"
	spanishPrefix = "Hola"
	french        = "French"
	spanish       = "Spanish"
)

func greet(name, language string) string {
	if name == "" {
		name = "World!"
	}

	// defaults to english
	prefix := englishPrefix

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix

	}
	return prefix + " " + name
}

/*
This function is going to do something, take some abitrary arguments and return some signature

Method signature are basically the returned values expected of that function.

// x, y are example of arguments here || a, b are the corresponding returned values
*/
func doSomething(x, y int) (a, b int) {
	return 0, 0
}

// Arrays
func learnArrays() {
	fmt.Println("This is something on Arrays")

	// var arrayOfInts [5]int
	dynamicArray := []string{"Marsh", "mellow", "someone", "dawta"}

	fmt.Println("Result", dynamicArray)
	fmt.Printf("Size of the array %v\n, and its type %T\n", len(dynamicArray), dynamicArray)
}
