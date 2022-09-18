package solve

import (
	"bytes"
	"fmt"
	"os"
)

// Check tetros validity
// Check file format (tetros separated by only one \n)

func CheckTetrominos(tetrominos [][][]byte, validTetrominos [][][]byte) {
	for _, tetro := range tetrominos {
		validity := false
		for _, validTetro := range validTetrominos {
			if len(tetro) == len(validTetro) {
				counter := 0
				for i := 0; i < len(tetro); i++ {
					if bytes.Equal(tetro[i], validTetro[i]) {
						counter++
					}
					if counter == len(tetro) {
						validity = true
					}
				}
			}
			if validity { // optimization
				break
			}
		}
		if !validity {
			fmt.Println("\033[31m", "ERROR : bad tetrominos")
			PrintError(tetro)
			os.Exit(0)
		}
	}
}

func CheckFileFormat(tetrominos []string) {
	for _, tetro := range tetrominos {
		counter := 0
		for _, cara := range tetro {
			if cara == '\n' {
				counter++
			}
		}
		if counter != 4 {
			fmt.Println("\033[31m" + "ERROR : file is badly formatted")

			fmt.Print("\033[31m", tetro)
			os.Exit(0)
		}
	}
}

func PrintError(tetro [][]byte) {
	for _, line := range tetro {
		for _, cara := range line {
			fmt.Print("\033[31m", string(cara))
		}
		fmt.Println()
	}
}

func CheckEmptyTetrosError(tetro [][]byte) {
	if tetro == nil {
		fmt.Println("\033[31m" + "ERROR : Empty Tetromino")
		os.Exit(0)
	}
}

func AllTetrominos() [][][]byte {
	tab := [][][]byte{
		{{35, 35}, {35, 35}},
		{{35, 35, 35, 35}},
		{{35}, {35}, {35}, {35}},
		{{35, 35, 35}, {46, 35, 46}},
		{{35, 46}, {35, 35}, {35, 46}},
		{{46, 35}, {35, 35}, {46, 35}},
		{{46, 35, 46}, {35, 35, 35}},
		{{35, 35, 46}, {46, 35, 35}},
		{{46, 35, 35}, {35, 35, 46}},
		{{35, 46}, {35, 35}, {46, 35}},
		{{46, 35}, {35, 35}, {35, 46}},
		{{35, 35, 35}, {46, 46, 35}},
		{{35, 35, 35}, {35, 46, 46}},
		{{46, 46, 35}, {35, 35, 35}},
		{{35, 46, 46}, {35, 35, 35}},
		{{35, 35}, {46, 35}, {46, 35}},
		{{35, 35}, {35, 46}, {35, 46}},
		{{35, 46}, {35, 46}, {35, 35}},
		{{46, 35}, {46, 35}, {35, 35}},
	}
	return tab
}
