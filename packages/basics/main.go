// packages/basics/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"

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

func main() {
	//fmt.Println(greet())
	//fmt.Println(greetWithName("Claudia"))
	//fmt.Println(greetWithNameAndAge("David", 29))
	//fmt.Println(divide(10, 2))
	//fmt.Println(divide(10, 0))
	fmt.Println(getRandomProverb())
}
