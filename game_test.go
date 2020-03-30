package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/dashjay/gobang/ai"
	"github.com/dashjay/gobang/windows"
)

func TestWindowGame(t *testing.T) {
	t.Log("test w > h")
	a := windows.NewGame(10, 20, 1)
	fmt.Print(a.Print())

	t.Log("test w < h")
	b := windows.NewGame(20, 10, 1)
	fmt.Print(b.Print())

	t.Log("test w > h and play some chess")
	c := windows.NewGame(10, 20, 1)
	c.Point(5, 4)
	c.Point(4, 5)
	c.Point(2, 3)
	fmt.Print(c.Print())
	t.Log("it works!!")
}

func TestWinGame(t *testing.T) {
	var res int
	// vertical win
	a := windows.NewGame(20, 10, 1)
	for _, i := range []int{5, 6, 7, 8, 9} {
		err := a.Point(1, i)
		if err != nil {
			
		}
	}

	fmt.Print(a.Print())
	fmt.Println(res)

	// Horizontal win
	b := windows.NewGame(10, 20, 1)
	b.Point(2, 5)
	b.Point(3, 5)
	b.Point(4, 5)
	b.Point(5, 5)
	b.Point(6, 5)
	fmt.Println(b.Print())
	fmt.Print(a.Print())
	fmt.Println(res)

	// left up to right down win
	c := windows.NewGame(20, 10, 1)
	for _, i := range []int{1, 2, 3, 4, 5} {
		res = c.Point(i, i)
	}
	fmt.Print(a.Print())
	fmt.Println(res)

	// right up to left down win
	d := windows.NewGame(20, 20, 1)
	d.Point(7, 2)
	d.Point(6, 3)
	d.Point(5, 4)
	d.Point(4, 5)
	res = d.Point(3, 6)

	fmt.Print(d.Print())
	fmt.Println(res)
}

func TestWinRegexp(t *testing.T) {
	win := "...*#*#*#...*****"
	res, _ := regexp.Compile(`\*{5}`)
	t.Log(res.MatchString(win))
}

func TestScoreRegexp(t *testing.T) {
	var score = []int{0, 0, 50, 200, 3000, 999999999}

	// line := "..@..@@@.@@@.@@@@@." // 1000003799
	// line := "..@..@@@.@@@.@@@@." // 3800
	line := "..@..@@.@@@.@@@." // 750
	getScore := 0
	var o int
	// 				  2, 3, 4, 5
	for i := 5; i >= 2; i-- {
		reg := regexp.MustCompile(fmt.Sprintf(`@{%d}`, i))
		temp := len(reg.FindAllStringSubmatch(line, -1))
		temp1 := temp - o
		getScore += temp1 * score[i]
		o += temp - temp1
	}
	fmt.Println(getScore)
}

func TestRunAI(t *testing.T) {
	g := windows.NewGame(10, 10, 1)
	g.Point(1, 1)
	g.Point(2, 2)
	g.Point(3, 3)
	g.Point(4, 4)
	g.Point(5, 5)
	ai.AI(*g)
}
