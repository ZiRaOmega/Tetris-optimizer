package solve

import (
	"fmt"
	"os"
)

func DisplayHead() {
	txt := ""
	for i := 7; i < len(os.Args[1])-4; i++ {
		txt += string(os.Args[1][i])
	}
	fmt.Println("\033[37m" + "=============================")
	fmt.Println("\033[0m", "###", txt, "###\n\nLoading ...")
}

func Color(x int, y int, field [][]byte) string {
	color := ""
	switch field[y][x] {
	case 'A':
		color = "\033[31m"
	case 'B':
		color = "\033[32m"
	case 'C':
		color = "\033[33m"
	case 'D':
		color = "\033[34m"
	case 'E':
		color = "\033[35m"
	case 'F':
		color = "\033[36m"
	case 'G':
		color = "\033[37m"
	case 'H':
		color = "\033[38;5;208m"
	case 'I':
		color = "\033[31m"
	case 'J':
		color = "\033[32m"
	case 'K':
		color = "\033[33m"
	case 'L':
		color = "\033[34m"
	case 'M':
		color = "\033[35m"
	case 'N':
		color = "\033[36m"
	case 'O':
		color = "\033[37m"
	case 'P':
		color = "\033[31m"
	}
	return color
}

// show solving details
func Showdetails(field [][]byte) {
	if len(os.Args) > 2 && os.Args[2] == "--details" {
		PrintSolution(field)
		fmt.Println()
	}
}
