// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i any = "hello"
	// var i any = 1 // This is bad news. Bad Assertions may cause a panic.
	fmt.Println(i.(string))
	// perform a type assertion and handle the error

	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an int\n", i)
	}
}
