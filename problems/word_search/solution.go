package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	runeWord := []rune(word)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			flag := solutionFinder(board, runeWord, [][]int{}, i, j)
			fmt.Println("end", board[i][j], flag)
			if flag {
				return true
			}
		}
	}
	return false
}

func solutionFinder(board [][]byte, word []rune, res [][]int, x, y int) bool {
	res = append(res, []int{x, y})
	fmt.Println("resultset", res)
	if !compareRune(word, res, board) {
		return false
	} else {
		if len(word) == len(res) {
			return true
		}
	}

	temp := board[x][y]
	board[x][y] = 0
	movements := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, item := range movements {
		newX := x + item[0]
		newY := y + item[1]
		if checkValidIndexBoard(newX, newY, len(board), len(board[0]), board) {
			flag := solutionFinder(board, word, res, newX, newY)
			if flag {
				return flag
			}
		}

	}
	board[x][y] = temp
	return false

}

func checkValidIndex(x, y, rows, columns int, indexes [][]int) bool {
	if x < 0 || x >= rows || y < 0 || y >= columns {
		return false
	}
	for _, item := range indexes {
		if x == item[0] && y == item[1] {
			return false
		}
	}
	return true
}

func checkValidIndexBoard(x, y, rows, columns int, board [][]byte) bool {
	if x < 0 || x >= rows || y < 0 || y >= columns || board[x][y] == 0 {
		return false
	}

	return true
}

func compareRune(word []rune, res [][]int, board [][]byte) bool {
	if len(res) > len(word) {
		return false
	}
	lenCur := len(res)
	x := res[lenCur-1]

	return word[lenCur-1] == rune(board[x[0]][x[1]])
}

func main() {
	board := []string{"ABCE", "SFCS", "ADEE"}
	word := "ABCB"
	board1 := make([][]byte, 0)
	for _, x := range board {
		board1 = append(board1, []byte(x))
	}
	fmt.Println(board1, word)
	fmt.Println(exist(board1, word))
}
