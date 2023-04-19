// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	// defer changes the flow-control within my program
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("Starting my script")
	// defer recovery
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("recover from panic:", err)
		}
	}()

	// panic
	panic("something bad happened")
}

// This might felt similar to a try-catch in other languages but go encourages to use and handle errors as values
// 