package ai

import (
	"fmt"
	"strings"

	"github.com/dashjay/gobang/windows"
)

var (
	cutCount    int
	searchCount int
)

// 必须把当前整个游戏传进来
func AI(d windows.Game) (x, y int) {
	// 初始化

	cutCount = 0
	searchCount = 0

	// 0 曾搜索测试
	fmt.Println(NegMax(&d, 0, -99999999, 99999999))

	fmt.Println("剪枝", cutCount)
	fmt.Println("搜索", searchCount)
	return 0, 0
}

func NegMax(game *windows.Game, depth, alpha, beta int) int {
	if depth == 0 {
		fmt.Println("depth 0")
		return Evaluation(game)
	}

	for x := 0; x <= game.Width; x++ {
		for y := 0; y <= game.Height; y++ {
			res := game.Point(x, y)
		}
	}
	return 0
}

//
func Evaluation(game *windows.Game) int {

	var totalScore = 0

	for y := 1; y < game.Height; y++ {
		var line strings.Builder
		// 收集一条line
		for x := 1; x < game.Width; x++ {
			line.WriteString(strings.TrimSpace(game.Str[game.Board[x][y]]))
		}
		totalScore += game.GetScore(line.String())
	}

	for x := 1; x < game.Width; x++ {
		var line strings.Builder
		for y := 1; y < game.Height; y++ {
			line.WriteString(strings.TrimSpace(game.Str[game.Board[x][y]]))
		}
		totalScore += game.GetScore(line.String())
	}

	fmt.Println(game.Print())
	return totalScore
}
