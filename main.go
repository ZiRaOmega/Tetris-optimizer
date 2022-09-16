package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"TetrisOptimizer/solve"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("ERROR : Missing file")
		return
	}

	file_content := ReadFile(os.Args[1]) + "\n\n"
	tetrominos := FindTetrominos(file_content)
	tetrominos = SimplifyTetrominos(tetrominos)

	validTetrominos := ReadFile("ValidTetrosList.txt")
	validTetrominosTab := strings.Split(validTetrominos, "\n\n")
	// validTetrominosByte := CreateByteTab(validTetrominosTab)
	for _, elem := range validTetrominosTab {
		fmt.Println(elem)
		fmt.Println()
	}
	// solve.ManageErrors(tetrominos, validTetrominos)

	tetrominosByte := CreateByteTab(tetrominos)
	tetrominosByte = FindAndRemoveExtraDots(tetrominosByte)
	min_size := FindMinSize(tetrominos)
	field := solve.CreateField(min_size)

	solve.DisplayHead()
	solve.Solve(0, tetrominosByte, min_size, field)
}

func ReadFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		print(err.Error())
	}
	return string(file)
}

func FindTetrominos(file_content string) []string {
	var tetrominos []string
	piece := ""
	line_counter := 1
	for _, cara := range file_content {
		if line_counter != 5 {
			piece += string(cara)
		} else {
			tetrominos = append(tetrominos, piece)
			piece = ""
		}

		if line_counter == 5 {
			line_counter = 0
		}
		if cara == '\n' {
			line_counter++
		}
	}
	return tetrominos
}

func SimplifyTetrominos(tab []string) []string {
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

func FindAndRemoveExtraDots(tetrominos [][][]byte) [][][]byte {
	newTab := [][][]byte{}
	for i := range tetrominos {
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
