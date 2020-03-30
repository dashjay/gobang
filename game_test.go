package main

import (
	"fmt"
	"testing"

	"github.com/dashjay/gobang/windows"
)

func TestWindowGame(t *testing.T) {
	t.Log("test w > h")
	a := windows.NewGame(10, 20)
	fmt.Print(a.Print())

	t.Log("test w < h")
	b := windows.NewGame(20, 10)
	fmt.Print(b.Print())

	t.Log("test w > h and play some chess")
	c := windows.NewGame(10, 20)
	c.Point(5, 4)
	c.Point(4, 5)
	c.Point(2, 3)
	fmt.Print(c.Print())
	t.Log("it works!!")
}
