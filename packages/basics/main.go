// packages/basics/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"

	//"unicode"
	"github.com/jboursiquot/go-proverbs"
)

// importing the fmt package from the pck go website
/*func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Today is %s", time.Now().Weekday())
	fmt.Println("")
}*/

// A simple greet function
func greet() string {
	return "Hi!"
}

// A simple greet function with the name of the person
func greetWithName(name string) string {
	return "Hi, " + name + "!"
}

// A simple greet function with the name of the person
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I am " + strconv.Itoa(age) +
		" years old."
	return
}

// Function that divides by zero two input numbers. If is divided by zero raise
// and exception
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil // nil is used to catch execution errors
	// nil es el valor cero para punteros, interfaces, mapas, slices, canales y funciones; y corresponde a la representación de un valor no inicializado
}

// Challenge #1
func getRandomProverb() (result string) {
	result = proverbs.Random().Saying
	return
}

// 2nd teaching lesson
var d, e, f int
var x, y, z = true, false, true

const pi = 3.14159
const a rune = 'a'

// a rune constast for the letter 'a' to print only UNICODE char
// Solamente acepta comillas simples y no dobles

func main() {
	//fmt.Println(greet())
	//fmt.Println(greetWithName("Claudia"))
	//fmt.Println(greetWithNameAndAge("David", 29))
	//fmt.Println(divide(10, 2))
	//fmt.Println(divide(10, 0))
	//fmt.Println(getRandomProverb())
	//d = 1
	//fmt.Println("d: %d\n", d)
	//x, y, z := 1, 1.25, "hello"
	//fmt.Printf("x: %x\n", x)
	//fmt.Println("x, y, z:", x, y, z)
	//printVars()
	//fmt.Printf("pi: %v - %T\n", pi, pi)
	//fmt.Printf("pi: %c - %T\n", a, a)
	//unicode.IsLetter(a)
	//a := 42.5
	//b := uint(a)
	// It lost precision when you convert
	//fmt.Printf("%v=%T, %v=%T\n", a, a, b, b)
	pointersExample()
}

// This will print the variables that defines as "global", not the variables within func
func printVars() {
	fmt.Println("x, y, z:", x, y, z)
}

func convertVars() {
	a := 42.5
	b := uint(a)
	// It lost precision when you convert
	fmt.Printf("%v=%T, %v=%T\n", a, a, b, b)
}

//Pointers function

func pointersExample() *int {
	// create a variable of type *T where T is an int
	var a *int

	// declare and assign b variable of type int
	b := 100

	// assign the address of b to a
	a = &b

	// print out the value of a which is the address of b
	// La ubicación en memoria
	fmt.Println(a)

	// print out the value at the address of b
	// el valor en memoria
	fmt.Println(*a)

	return a
}
