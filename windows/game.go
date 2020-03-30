package windows

import (
	"fmt"
	"strings"
)

type Game struct {
	str    []string
	Width  int
	Height int
	Board  [][]int
}

func NewGame(w, h int) *Game {

	// init board ptr
	var board [][]int
	// every Y
	for tempy := 0; tempy <= w+1; tempy++ {
		// for every line ,init point state to 0
		var line = make([]int, h+1)
		board = append(board, line)
	}

	g := Game{
		Width:  w,
		Height: h,
		Board:  board,
		str:    []string{" . ", " * ", " # "},
	}

	return &g
}

func (g *Game) Print() string {
	// header line like this
	//  x  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18 19 20x轴
	var header strings.Builder

	// output multi line like this
	// 1  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
	// 2  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
	// ...
	// 10  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
	var output strings.Builder
	// write a mount point to (0, 0)
	header.WriteString(" x ")

	// for numbers < 9, they takes just 1 grid.
	// but numbers > 9, they takes more than 1 grid
	// so I set a two side margin to numbers < 9
	// just left side margin to number > 10
	for i := 1; i <= g.Width; i++ {
		if i > 9 {
			header.WriteString(fmt.Sprintf(" %d", i))
		} else {
			header.WriteString(fmt.Sprintf(" %d ", i))
		}
	}

	// range the board
	for tempy := 1; tempy <= g.Height; tempy++ {
		// this is also like above
		if tempy > 9 {
			output.WriteString(fmt.Sprintf("%d ", tempy))
		} else {
			output.WriteString(fmt.Sprintf(" %d ", tempy))
		}
		// write state
		for tempx := 1; tempx <= g.Width; tempx++ {
			output.WriteString(g.str[g.Board[tempx][tempy]])
		}
		// line end
		output.WriteString("\n")
	}

	output.WriteString("y轴 \n")
	header.WriteString("x轴 \n")

	return header.String() + output.String()
}

func (g *Game) Point(x, y int) {
	if x > g.Width || y > g.Height || x < 0 || y < 0 {
		panic("x y error")
	} else {
		g.Board[x][y] = 1
	}
}
