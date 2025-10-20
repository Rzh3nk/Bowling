package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("bowling game")
	var strikeLast = false
	var strikePreLast = false
	var spareCount = 0
	var round = 0
	var totalCount = 0
	for round < 10 {
		var points int
		var prevpoints int
		fmt.Println(totalCount)
		fmt.Println("Сколько очков за первый бросок?")
		fmt.Fscan(os.Stdin, &points)
		totalCount += points
		if strikeLast || spareCount == 1 {
			totalCount += points
		}
		if strikePreLast {
			totalCount += points
		}

		strikePreLast = strikeLast

		if spareCount == 1 {
			spareCount = 0
		}
		strikeLast = false
		if points == 10 {
			strikeLast = true
		}

		if points != 10 || round == 9 {
			fmt.Println("Сколько очков за второй бросок?")
			prevpoints = points
			fmt.Fscan(os.Stdin, &points)
			totalCount += points
			if strikeLast && round != 9 {
				totalCount += points
			}
			if strikePreLast {
				totalCount += points
			}

			if points+prevpoints == 10 {
				spareCount++
			}
			strikePreLast = strikeLast
		}
		if round == 9 && (spareCount == 1 || strikeLast || strikePreLast) {
			fmt.Println("Сколько очков за третий бросок?")
			prevpoints = points
			fmt.Fscan(os.Stdin, &points)
			totalCount += points
		}

		round++

	}
	println(totalCount)

}

