package bootcamp

import (
	"fmt"
	"github.com/alem-platform/ap"
)

func printTopBorder(w int) {
	for i := 0; i < w; i++ {
		ap.PutRune(' ')
		for j := 0; j < 7; j++ {
			ap.PutRune('_')
		}
	}
	ap.PutRune('\n')
}

func printMiddleRow(line []rune, w int) {
	for i := 0; i < w; i++ {
		ap.PutRune('|')
		switch line[i] {
		case '0': // wall
			for j := 0; j < 7; j++ {
				ap.PutRune('X')
			}
		case '2': // player
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune('>')
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune(' ')
		case '3': // award
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune('*')
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune(' ')
			ap.PutRune(' ')
		default: // '1' or anything else (free cell)
			for j := 0; j < 7; j++ {
				ap.PutRune(' ')
			}
		}
	}
	ap.PutRune('|')
	ap.PutRune('\n')
}

func printBottomBorder(w int) {
	for i := 0; i < w; i++ {
		ap.PutRune('|')
		for j := 0; j < 7; j++ {
			ap.PutRune('_')
		}
	}
	ap.PutRune('|')
	ap.PutRune('\n')
}

func main() {
	var h, w int
	fmt.Scanf("%d %d\n", &h, &w)

	var grid [100][100]rune

	// Read grid input
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Scanf("%c", &grid[i][j])
		}
		fmt.Scanf("\n")
	}

	// Draw map
	for i := 0; i < h; i++ {
		printTopBorder(w)
		printMiddleRow(grid[i][:], w)
		printBottomBorder(w)
	}
}

