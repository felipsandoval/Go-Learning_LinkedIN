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

	// read a value from the map with a known key
	fmt.Println(authors["cb"])
	fmt.Println("Not exist: ", authors["jr"]) // Non existent value from the map
	a, ok := authors["jr"]
	fmt.Printf("a = %v, ok = %v\n", a, ok) // Check if a value is real and exists within the map. The OK is a boolean

	//update a value from the map
	authors["cb"] = author{name: "Charles Beck"}
	fmt.Println(authors["cb"])

	//delete a value from a map

	delete(authors, "cb")
	fmt.Printf("%v\n", authors)



}
