package main

import (
	"fmt"

	"github.com/dashjay/gobang/windows"
)

const (
	InputFormatError = "输入的格式有误，请重新输入 「x,y」"
)

func main() {
	// When change turn we can use this arr to help
	changeTurn := []int{0, 2, 1}

	// start a game width:10 height:10 player1 first TODO: player0 will undertake by AI
	tryGame := windows.NewGame(10, 10, 1)
	// player1 play
	for {

		fmt.Printf("轮到玩家%d落子\n", tryGame.Turn)
		fmt.Print("\r" + tryGame.Print())
		var (
			res int
			x   int
			y   int
		)

		// read x,y from stdin
		_, err := fmt.Scanf("%d,%d", &x, &y)
		if err != nil {
			fmt.Println(InputFormatError, err)
			continue
		}

		// point one on x,y
		err = tryGame.Point(x, y)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// check if someone win the game
		res = tryGame.WinGame(x, y)

		// 没有输赢
		if res == -1 {
			tryGame.Turn = changeTurn[tryGame.Turn]
			continue
		}
		fmt.Printf("玩家%d已经赢了", res)
		return
	}
}
