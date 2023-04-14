// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	//var authors map[string]author

	// initialize the map with make. Its not needed to initialize a size for it
	/*authors = make(map[string]author)

	// add authors to the map
	authors["tm"] = author{name: "Toni Morrison"}
	authors["cb"] = author{name: "Carla Becker"}*/
	authors := map[string]author{
		"tm": {name: "Toni Morrison"},
		"cb": {name: "Carla Becker"},
	}

	// print the map with fmt.Printf
	// verbose representation of the map
	fmt.Printf("%#v\n", authors)
	fmt.Printf("%v\n", authors)
	fmt.Println(authors["tm"])
	fmt.Println("Not exist: ", authors["jr"])
	// read a value from the map with a known key
	//
}
