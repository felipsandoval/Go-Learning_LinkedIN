// type/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author. A collection of fields
// When you create and declare a struct type you are basing on a Known Go Type
type author struct {
	first, last string
}

// fullName returns the full name of the author. First and last Name os the author separated by the name.
// THIS IS A STRUCT METHOD
func (a author) fullName() string {
	return a.first + " " + a.last
}

// changeName changes the first and last name of the author
// IMPORTANT to notice the asterisk in order to access the memory value from our data struct
func (a *author) changeName(first, last string) {
	a.first = first
	a.last = last
}

// GO does not support inheritance just adage embedding

type person struct {
	first string
	last  string
}

func (p person) personFullName() string {
	return p.first + " " + p.last
}

type woman struct {
	person
	penName string
}

// override fullName method for author
func (w woman) fullName() string {
	return fmt.Sprintf("%s (%s)", w.person.personFullName(), w.penName)
}

func main() {
	// initialize
	a := author{
		first: "James",
		last:  "Baldwin",
	}
	b := author{
		first: "Carlos",
		last:  "Gaviria",
	}
	//fmt.Printf("%#v\n", a) // Print the author key main fields. Important debugging tool
	//fmt.Printf("%v\n", a) // Print the author values fields
	fmt.Println(a.fullName())
	fmt.Println(b.fullName())        // In Go you CANT declare a variable and not use it. Its forbidden.
	a.changeName("David", "Baldwin") // This does not work because it modify a copy of our author NOT the one we initialize
	fmt.Println(a.fullName())
	p := person{
		first: "Matthew",
		last:  "Phantom",
	}
	fmt.Println(p.personFullName())
	w := woman{
		person: person{
			first: "Jenna",
			last:  "Lopez",
		},
		penName: "J-LO",
	}
	fmt.Println(w.personFullName())
	fmt.Println(w.fullName())
	// Declare an arrays
	var arr [5]int
	// Print the array. YOU CANT RESIZE IT
	fmt.Println(arr)

	// declare a slice string the make function. Dynamic arrays index
	// No more worries about sizing concerns
	names := make([]string, 0)
	// Appending values to my array
	names = append(names, "Ali")
	names = append(names, "Charles")
	names = append(names, "Philippe")
	// If I go beneath the underlying array size a new one will be created with the original values transferred.
	// Usually double the size of the original
	// Index starts in 0
	fmt.Println(names)
	// Slice the slice using syntax slice[low:high] High bound is NOT inclusive
	fmt.Println(names[1:3])
	fmt.Println(names[1:])
	fmt.Println(names[:1])
	fmt.Println(names[:])
	fmt.Println(names[0:])
	fmt.Println(names[0:2])
	//fmt.Println(names[1:5])
}
