package main

import (
	"github.com/dashjay/gobang/windows"
)

func main() {

	g := windows.NewGame(10, 10)
	g.Point(5, 4)
	g.Print()

}
