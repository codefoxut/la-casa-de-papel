package main

import (
	"fmt"
	"strings"
)

func moveTheRRobot(inputStr string) (int, int) {
	var x, y int
	for _, move := range inputStr {
		switch strings.ToLower(string(move)) {
		case "u":
			y += 1
		case "d":
			y -= 1
		case "r":
			x += 1
		case "l":
			x -= 1
		}

	}
	return x, y
}

func main() {
	inputString := "UDLRRRR"
	x, y := moveTheRRobot(inputString)

	if x == 0 && y == 0 {
		fmt.Println("Robot has returned to the origin.")
	} else {
		fmt.Println("Robot has not returned to the origin")
	}

}
