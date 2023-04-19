// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
//
type author struct {
	name string
}

// define a book type with a title and author
//
type book struct {
	title  string
	author author
}

// define a library type to track a list of books
//
type library map[string][]book

// define addBook to add a book to the library
//
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
//
func (l library) lookupByAuthor(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	//
	library := make(library)
	// add 2 books to the library by the same author
	//
	/*var author_1 = author{name: "JK Rowling"}
	var book_1 = book{
		title:  "Harry Potter: First One",
		author: author_1,
	}
	var book_2 = book{
		title:  "Harry Potter: Second One",
		author: author_1,
	}
	//fmt.Printf("%v and %v", book_1, book_2)

	library.addBook(book_1)*/
	author_1 := author{name: "JK Rowling"}
	library.addBook(book{
		title:  "Harry Potter: First One",
		author: author_1,
	})
	library.addBook(book{
		title:  "Harry Potter: Second One",
		author: author_1,
	})
	// add 1 book to the library by a different author
	//
	library.addBook(book{
		title:  "Lord of the Rings",
		author: author{name: "Johannes Oven"},
	})
	// dump the library with spew
	//
	spew.Dump(library)
	// lookup books by known author in the library
	//
	books := library.lookupByAuthor(author_1.name)
	//spew.Dump(books)
	// print out the first book's title and its author's name
	//
	if len(books) != 0 {
		b := books[0]
		fmt.Println(b.title, "by", b.author.name)
	}
}
