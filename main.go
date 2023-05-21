package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"

	"TetrisOptimizer/solve"
)

func main() {
	start := time.Now()

	if len(os.Args) == 1 {
		fmt.Println("ERROR : Missing file")
		return
	}
	solve.DisplayHead()

	// creating tetrominos tab
	file_content := ReadFile(os.Args[1])
	tetrominos := FindTetrominos(file_content)

	solve.CheckFileFormat(tetrominos)

	tetrominos = RemoveDotsLines(tetrominos)
	tetrominosByte := CreateByteTab(tetrominos)
	tetrominosByte = RemoveDotsColumns(tetrominosByte)

	// Error management (bad tetrominos or bad file format)
	validTetros := solve.AllTetrominos()
	solve.CheckTetrominos(tetrominosByte, validTetros)
	// set field
	min_size := FindMinSize(tetrominos)
	field := solve.CreateField(min_size)

	solve.Solve(0, tetrominosByte, min_size, field, start)
}

func ReadFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		print(err.Error())
	}
	return string(file)
}

func FindTetrominos(file_content string) []string {
	tab := strings.Split(file_content, "\n\n")

	tetrominos := []string{}
	for i, elem := range tab {
		if i != len(tetrominos)-1 {
			elem += "\n"
			tetrominos = append(tetrominos, elem)
		}
	}
	return tetrominos
}

func RemoveDotsLines(tab []string) []string {
	new_tab := []string{}
	// CHECK LINES
	for _, tetromino := range tab {
		pattern := regexp.MustCompile(`\.\.\.\.\n`)
		new_tab = append(new_tab, pattern.ReplaceAllString(tetromino, ``))
	}

	return new_tab
}

func FindMinSize(tetrominos []string) int {
	var hashtag_amount float64
	for _, tetromino := range tetrominos {
		for _, cara := range tetromino {
			if cara != '.' && cara != '\n' {
				hashtag_amount++
			}
		}
	}

	return RoundSup(math.Sqrt(hashtag_amount))
}

func RoundSup(x float64) int {
	x += 0.49
	return int(math.Round(x))
}

func CreateByteTab(tetrominos []string) [][][]byte {
	var tab [][][]byte
	for i, tetro := range tetrominos {
		tab = append(tab, nil)
		var line []byte

		for _, lettre := range tetro {
			if lettre != '\n' {
				line = append(line, byte(lettre))
			} else {
				tab[i] = append(tab[i], line)
				line = nil
			}
		}
	}
	return tab
}

func RemoveDotsColumns(tetrominos [][][]byte) [][][]byte {
	newTab := [][][]byte{}
	for i := range tetrominos {
		solve.CheckEmptyTetrosError(tetrominos[i])
		founded := false
		counter := 0
		indexsToRemove := []int{}
		for k := 0; k < len(tetrominos[i][0]); k++ {
			for _, line := range tetrominos[i] {
				if line[k] == '.' {
					counter++
				}
			}
			if counter == len(tetrominos[i]) {
				indexsToRemove = append(indexsToRemove, k)
				founded = true
			}
			counter = 0
		}
		if founded {
			newTab = append(newTab, nil)

			for _, line := range tetrominos[i] {
				newLine := RemoveExtraDots(line, indexsToRemove)
				newTab[i] = append(newTab[i], newLine)

			}
		} else {
			newTab = append(newTab, tetrominos[i])
		}
	}
	return newTab
}

func RemoveExtraDots(line []byte, indexs []int) []byte {
	newLine := []byte{}
	for i := range line {
		indexfounded := false
		for _, val := range indexs {
			if i == val {
				indexfounded = true
			}
		}
		if !indexfounded {
			newLine = append(newLine, line[i])
		}
	}

	return newLine
}
