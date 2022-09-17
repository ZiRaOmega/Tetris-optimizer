package solve

import (
	"fmt"
	"os"
)

func Solve(tetroIndex int, tetrominos [][][]byte, size int, field [][]byte) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if CanPut(x, y, tetroIndex, tetrominos, field) {
				Showdetails(field)
				if y == len(field)-1 || tetroIndex == len(tetrominos)-1 {
					fmt.Println("\033[0m" + "=====================")
					PrintSolution(field)
					fmt.Println("\033[0m" + "=====================")
					os.Exit(0)

				} else {
					Solve(tetroIndex+1, tetrominos, size, field) // continue solving ...
				}
				BackTrack(x, y, tetroIndex, tetrominos, field)
			}
		}
	}

	if tetroIndex == 0 {
		field := CreateField(size + 1)
		Solve(tetroIndex, tetrominos, size+1, field)
	}
}

// checks if we can put the tetro in the field
func CanPut(x int, y int, tetroIndex int, tetrominos [][][]byte, field [][]byte) bool {
	// does the tetro fit in the field ?
	if len(tetrominos[tetroIndex])+y > len(field) {
		return false
	}
	for i := 0; i < len(tetrominos[tetroIndex]); i++ {
		if len(tetrominos[tetroIndex][i])+x > len(field) {
			return false
		}
	}

	// is there enough empty space in the field in order to put the tetro in it ?
	for a := y; a < len(tetrominos[tetroIndex])+y; a++ {
		for b := x; b < len(tetrominos[tetroIndex][a-y])+x; b++ {

			if tetrominos[tetroIndex][a-y][b-x] == '.' {
				continue
			}

			if field[a][b] == '?' {
				field[a][b] = byte(tetroIndex + 'A')
			} else {
				BackTrack(x, y, tetroIndex, tetrominos, field)
				return false
			}
		}
	}
	return true
}

func CreateField(size int) [][]byte {
	var board [][]byte

	for i := 0; i < size; i++ {
		board = append(board, nil)
		for k := 0; k < size; k++ {
			board[i] = append(board[i], '?')
		}
	}
	return board
}

func BackTrack(x int, y int, tetroIndex int, tetrominos [][][]byte, field [][]byte) {
	for a := y; a < len(tetrominos[tetroIndex])+y; a++ {
		for b := x; b < len(tetrominos[tetroIndex][a-y])+x; b++ {

			if tetrominos[tetroIndex][a-y][b-x] == '.' {
				continue
			}

			if field[a][b] == byte(tetroIndex+'A') {
				field[a][b] = '?'
			}
		}
	}
}

func PrintSolution(field [][]byte) {
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			if field[y][x] == '?' {
				fmt.Print("\033[0m", ".")
			} else {
				fmt.Print(Color(x, y, field), string(field[y][x]))
			}
		}
		fmt.Println()
	}
}
