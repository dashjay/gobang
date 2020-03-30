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

	var board [][]int
	for tempy := 0; tempy <= w+1; tempy++ {
		var line []int
		for tempx := 0; tempx <= w+1; tempx++ {
			line = append(line, 0)
		}
		board = append(board, line)
	}
	g := Game{
		Width:  w,
		Height: h,
		Board:  board,
		str:    []string{" . ", " * ", " # "},
	}

	g.Board[1][2] = 1
	return &g
}

func (g *Game) Print() {
	var header strings.Builder
	var output strings.Builder
	header.WriteString(" x ")
	for tempy := 1; tempy <= g.Width; tempy++ {
		header.WriteString(fmt.Sprintf(" %d ", tempy))
		if tempy > 9 {
			output.WriteString(fmt.Sprintf("%d ", tempy))
		} else {
			output.WriteString(fmt.Sprintf(" %d ", tempy))
		}
		for tempx := 1; tempx <= g.Height; tempx++ {
			output.WriteString(g.str[g.Board[tempx][tempy]])
		}
		output.WriteString("\n")
	}

	output.WriteString("y轴 \n")
	header.WriteString("x轴 \n")
	fmt.Print(header.String())
	fmt.Print(output.String())
}

func (g *Game) Point(x, y int) {
	if x > g.Width || y > g.Height || x < 0 || y < 0 {
		panic("x y error")
	} else {
		g.Board[x][y] = 1
	}
}
