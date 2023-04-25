// generics/type-sets/begin/main.go
package main

import "fmt"

// create a numeric interface with a type set
type numeric interface {
	~int | ~float64
	grow()
}

// The benefit from doing it this way is that numeric is an interface like any other like
// we are accustomed to and go. We can add a method to it and require that any type numeric
// also have that method

// update sum function to use a numeric interface with a type set
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int

func (si specialInt) grow() {

}

func main() {
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))
}
