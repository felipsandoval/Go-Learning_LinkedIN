// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func handleFile() (text string) {
	//data, err := os.ReadFile("challenges/data.txt")
	// use os package to read the file specified as a command line argument
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read the file: %s", err))
	}
	//os.Stdout.Write(data) //prints the value within my file
	// convert the bytes to a string
	text = string(data)
	fmt.Println() // to print a new line for not having all the information within the same line tho
	return text
}

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	fmt.Println("Starting the challenge #4")
	//os.Stderr.Read([]byte, os.Args[])
	// Handle the file and returns the text within its document
	text := handleFile()

	// initialize a map to store the counts
	count := make(map[string]int, 0)
	fmt.Printf("%v\n", count)
	// use the standard library utility package that can help us split the string into words
	words := strings.Split(text, " ")
	//words := strings.Fields(text) // it does the same as the upper func

	//fmt.Printf("(%T) (len=%v) \"words\": (%T) %v\n", "words", len("words"), len(words), len(words))
	// capture the length of the words slice
	count["words"] = len(words)
	//spew.Dump(count) //useful to show the len and values from a MAP
	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//var numbers int
	//var letters int
	//var symbols int
	for _, word := range words {
		for _, char := range word {
			if unicode.IsNumber(char) {
				//numbers += 1
				count["numbers"]++
			} else if unicode.IsLetter(char) {
				//letters += 1
				count["letters"]++
			}else{
				//symbols += 1
				count["symbols"]++
			}
		}
	}
	/*
	count["letters"] = letters
	count["symbols"] = symbols
	count["numbers"] = numbers*/
	// dump the map to the console using the spew package
	spew.Dump(count)
	// the order of the keys of the map is totally random. They are different from the order they are defined.
}
