package board

import (
	"fmt"
)

// DrawBoard renders the Minesweeper board on the screen
func DrawBoard(h, w int, revealed [][]bool, grid [][]rune, counts [][]int) {
	cellW := 7

	// Column numbers
	fmt.Printf("   ")
	for j := 0; j < w; j++ {
		numStr := fmt.Sprintf("%d", j+1)
		fmt.Printf("%-*s", cellW, numStr) // left-aligned inside cell width
	}
	fmt.Printf("\n")

	// Top line
	fmt.Printf("    ")
	for i := 0; i < w*cellW+w; i++ { // +w for the '|' separators
		fmt.Printf("_")
	}
	fmt.Printf("\n")

	// Rows
	for i := 0; i < h; i++ {
		// top empty line
		fmt.Printf("   ")
		for j := 0; j < w; j++ {
			fmt.Printf("|")
			if revealed[i][j] {
				fmt.Printf("%-*s", cellW, "")
			} else {
				fmt.Printf("%-*s", cellW, "XXXXXXX")
			}
		}
		fmt.Printf("|\n")

		// content line
		fmt.Printf("%2d ", i+1) // row number
		for j := 0; j < w; j++ {
			fmt.Printf("|")
			if revealed[i][j] {
				if grid[i][j] == '*' {
					fmt.Printf("%-*s", cellW, "*")
				} else if counts[i][j] == 0 {
					fmt.Printf("%-*s", cellW, "")
				} else {
					fmt.Printf("%-*d", cellW, counts[i][j])
				}
			} else {
				fmt.Printf("%-*s", cellW, "XXXXXXX")
			}
		}
		fmt.Printf("|\n")

		// bottom line
		fmt.Printf("   ")
		for j := 0; j < w; j++ {
			fmt.Printf("|")
			for k := 0; k < cellW; k++ {
				if revealed[i][j] {
					fmt.Printf("_")
				} else {
					fmt.Printf("X")
				}
			}
		}
		fmt.Printf("|\n")
	}
}
