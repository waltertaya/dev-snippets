/**
* Solution by https://github.com/waltertaya
* isValidSudoku func:
*		param: [][]byte
* 		rtype: bool
* 	isValid:
* 		1. Each row must contain the digits 1-9 without repetition
*		2. Each column must contain the digits 1-9 without repetition
*		3. Each of the nine 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition
**/

package main

import "fmt"

func main() {
	testBoard1 := [][]byte{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if results := isValidSudoku(testBoard1); results {
		fmt.Println("Test case passed")
	} else {
		fmt.Println("Test case failed")
	}
}

func isValidSudoku(board [][]byte) bool {
	// Check rows
	for _, row := range board {
		seen := make(map[byte]bool)
		for _, val := range row {
			if val == 0 {
				continue
			}
			if seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	// Check columns
	for col := 0; col < 9; col++ {
		seen := make(map[byte]bool)
		for row := 0; row < 9; row++ {
			val := board[row][col]
			if val == 0 {
				continue
			}
			if seen[val] {
				return false
			}
			seen[val] = true
		}
	}

	// Check 3x3 sub-boxes
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := make(map[byte]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := board[boxRow*3+i][boxCol*3+j]
					if val == 0 {
						continue
					}
					if seen[val] {
						return false
					}
					seen[val] = true
				}
			}
		}
	}

	return true
}
