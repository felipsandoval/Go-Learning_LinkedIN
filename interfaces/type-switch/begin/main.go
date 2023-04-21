// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type

func whatAmI(arg interface{}) string {
	switch arg.(type) {
	case int:
		return fmt.Sprintf("%v is an integer", arg)
	case string:
		return fmt.Sprintf("%v is a string", arg)
	default:
		return fmt.Sprintf("%v is not a supported type", arg)
	}
}

func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
