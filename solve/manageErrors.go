package solve

import (
	"fmt"
	"os"
	"strings"
)

// Check tetros validity
// Check file format (tetros separated by only one \n)

func ManageErrors(tetrominos []string, validTetrosStr string) {
	validTetros := strings.Split(validTetrosStr, "\n\n")
	fmt.Println(validTetros[0])
	fmt.Println()
	fmt.Println(tetrominos[0])
	for _, tetro := range tetrominos {
		founded := false
		for _, validtetro := range validTetros {
			if validtetro == tetro {
				founded = true
				break
			}
		}
		if !founded {
			fmt.Println("ERROR : Tetrominos invalide")
			os.Exit(1)
		}
	}
}
