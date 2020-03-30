package windows

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//  Game 一个游戏
type Game struct {
	// 游戏所使用的字符集 例如 .代表空 @代表player1 #代表player2
	Str []string
	// 宽度 x 轴
	Width int
	// 高度 y 轴
	Height int
	// 棋盘
	Board [][]int
	// 当前游戏轮到谁
	Turn int
}

// scoreArr 当有0,1个子是毫无意义的,
// 当有2个连起来对应50分
//    3个连起来对应200分
// 当有4个连起来对应3000分
// 5 -> 999999999
var scoreArr = []int{0, 0, 50, 200, 3000, 2 << 28}

// start a game w: width
// h: height
// first 0 or 1
func NewGame(w, h, first int) *Game {
	if first != 1 && first != 2 {
		panic(fmt.Sprintf("先手只能为1或2你传入了%d", first))
	}

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
		Str:    []string{" . ", " @ ", " # "},
		Turn:   first, // who play first
	}

	return &g
}

// print the board
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
			output.WriteString(g.Str[g.Board[tempx][tempy]])
		}
		// line end
		output.WriteString("\n")
	}

	output.WriteString("y轴 \n")
	header.WriteString("x轴 \n")

	return header.String() + output.String()
}

var (
	ErrorOverFlow  = errors.New("point over the width or height")
	HasBeenTakenUp = errors.New("point has been taken up")
)

// Point point one at x, y
func (g *Game) Point(x, y int) error {
	// 超过了边界
	if x > g.Width || y > g.Height || x < 0 || y < 0 {
		return ErrorOverFlow
	}
	// 已经被占用
	if g.Board[x][y] != 0 {
		return HasBeenTakenUp
	}
	g.Board[x][y] = g.Turn
	return nil
}

// WinGame who wins the game ?
// pass in x,y current position
func (g *Game) WinGame(x, y int) int {

	// 赢的条件
	Win := regexp.MustCompile(fmt.Sprintf(`%s{5}`, strings.TrimSpace(g.Str[g.Turn])))

	var (
		// from left to right
		buf21 strings.Builder
		// from up to down
		buf14 strings.Builder
		// // from right up to left down
		buf13 strings.Builder
		// // from left up to right down
		buf24 strings.Builder
	)
	// pin y, range x from 1 to width
	for tempx := 1; tempx <= g.Width; tempx++ {
		buf21.WriteString(strings.TrimSpace(g.Str[g.Board[tempx][y]]))
	}

	if Win.MatchString(buf21.String()) {
		return g.Turn
	}

	for tempy := 1; tempy <= g.Height; tempy++ {
		buf14.WriteString(strings.TrimSpace(g.Str[g.Board[x][tempy]]))
	}

	if Win.MatchString(buf14.String()) {
		return g.Turn
	}
	var (
		tempx = x
		tempy = y
	)
	// 寻到这条线左上角的点
	for ; tempx > 1 && tempy > 1; {
		tempx--
		tempy--
	}
	// 获取所有字符
	for ; tempx < g.Width && tempy < g.Height; {
		buf24.WriteString(strings.TrimSpace(g.Str[g.Board[tempx][tempy]]))
		// 增加
		tempx++
		tempy++
	}
	// 判断输赢
	if Win.MatchString(buf24.String()) {
		return g.Turn
	}

	tempx = x
	tempy = y
	// 寻找到右上角的点
	for ; tempx < g.Width && tempy > 1; {
		tempx++
		tempy--
	}

	for ; tempx > 1 && tempy < g.Height; {
		buf13.WriteString(strings.TrimSpace(g.Str[g.Board[tempx][tempy]]))
		tempx--
		tempy++
	}
	if Win.MatchString(buf13.String()) {
		return g.Turn
	}

	return -1
}

// GetScore 得到某条线上的得分~
func (g *Game) GetScore(line string) int {

	var (
		MyScore       = 0
		OppositeScore = 0
		MyCount       = 0
		OppositeCount = 0
	)
	opposite := []int{0, 2, 1}

	for i := 5; i >= 2; i-- {

		// 我是的得分正则 如果我是 1那么对手就是 2
		MyWin := regexp.MustCompile(fmt.Sprintf(`%s{%d}`, strings.TrimSpace(g.Str[g.Turn]), i))

		// 寻找我的连续 n 颗棋子
		tempMyCount := len(MyWin.FindAllStringSubmatch(line, -1))
		tempMyCount1 := tempMyCount - MyCount
		MyScore += tempMyCount1 * scoreArr[i]
		MyCount += tempMyCount1

		// 寻找对手的连续 n 颗棋子
		OppositeWin := regexp.MustCompile(fmt.Sprintf(`%s{%d}`, strings.TrimSpace(g.Str[opposite[g.Turn]]), i))
		tempOppositeCount := len(OppositeWin.FindAllStringSubmatch(line, -1))
		tempOppositeCount1 := tempOppositeCount - OppositeCount
		OppositeScore += tempOppositeCount1 * scoreArr[i]
		OppositeCount += OppositeScore

	}

	return MyScore - OppositeScore
}
