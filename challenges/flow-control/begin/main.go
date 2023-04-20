// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
)

type count map[string]string

func handleFiles() {
	//data, err := os.ReadFile("challenges/data.txt")
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read the file: %s", err))
	}
	os.Stdout.Write(data)
	fmt.Println() // to print a new line for not having all the information within the same line tho 
}

func main() {
	// handle any panics that might occur anywhere in this function
	//
	fmt.Println("Starting the challenge #4")
	//os.Stderr.Read([]byte, os.Args[])
	handleFiles()
	// use os package to read the file specified as a command line argument
	//

	// convert the bytes to a string
	//

	// initialize a map to store the counts
	//
	/*counter := count {}
	counter := count{"aloha": "aloha"}
	fmt.Printf("\nThis is my map: %#v\n", counter)*/
	// use the standard library utility package that can help us split the string into words
	//

	// capture the length of the words slice
	//

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//

	// dump the map to the console using the spew package
	//
}
