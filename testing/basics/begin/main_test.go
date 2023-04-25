// testing/basics/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

// Tp execute it it needs a go test and name of the file
// write a test for sum
func TestSum(t *testing.T) {
	got := sum(1, 2, 3)
	want := 6
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("***SETUP***")
	m.Run()
	fmt.Println("***TEARDOWN***")
}

// write a TestMain for setup and teardown
