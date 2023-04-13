// packages/basics/main.go
package main

import (
	"fmt"
)

// Declaring a package level variable val with a string value = "global"
var val string = "global"

func main() {
	// create a local variable "val" with a int value
	val := 42
	// Print type and value of my local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)
	// Print the value of my global variable "val"
	printVal()
	// Update the value of my global variable "val"
	updateVal()
	// Print the value of my global dereferenced variable "val"
	printVal()
	// Print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", &val, &val)
	// Assing a value directly to the pointer address of the local variable "val"
	*(&val) = 100
	fmt.Printf("local val = %v\n", val)
}

// Print out the type and value of the local val
func printVal() {
	fmt.Println("global val = ", val)
}

// Update the value of global value val
func updateVal() {
	val = "updated global"
}
