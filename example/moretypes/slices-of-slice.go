package main

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("============ 切片的切片 ==========")

	one := []string{"_", "_", "_"}
	board := [][]string{
		one,
		one,
		one,
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
