package main

import (
	"testing"

	"github.com/dashjay/gobang/windows"
)

func TestWindowGame(t *testing.T) {
	t.Log("test w > h")
	g := windows.NewGame(10, 20)
	t.Log(g.Print())
}
