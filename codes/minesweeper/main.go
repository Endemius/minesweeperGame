package main

import (
	"fmt"
	. "minesweeper/board"
	. "minesweeper/game_logic"
)

func main() {
	var h, w int // Step 1: Read board size
	fmt.Printf("Enter height and width:\n")
	if _, err := fmt.Scanf("%d %d", &h, &w); err != nil {
		fmt.Printf("Error: invalid input.(not a digit)\n")
		return
	}
	if h < 3 || w < 3 {
		fmt.Printf("Error: invalid input.( the map is too small)\n")
		return
	}
	if h > 20 || w > 20 {
		fmt.Printf("Error: invalid input.( the map is too big)\n")
		return
	}

	grid := make([][]rune, h)
	bombs := 0
	fmt.Printf("Enter the content of the map(. or *):\n")
	for i := 0; i < h; i++ { // Step 2: Read the board layout
		var row string
		fmt.Scanf("%s", &row)
		if len(row) != w {
			fmt.Printf("Error: invalid input.(rows don't match the height or width)\n")
			return
		}
		grid[i] = make([]rune, w) // Create row and validate each cell
		for j, ch := range row {
			if ch == '*' {
				grid[i][j] = '*'
				bombs++
			} else if ch == '.' {
				grid[i][j] = '.'
			} else {
				fmt.Printf("Error: invalid input.(only '.' or '*' allowed)\n")
				return
			}
		}
	}
	if bombs < 2 { // Ensure there are at least 2 bombs to make the game playable
		fmt.Printf("Error: invalid input.(not enough bombs)\n")
		return
	}
	if bombs >= h*w { // Ensure there are at least 2 bombs to make the game playable
		fmt.Printf("Error: invalid input.(too many bombs)\n")
		return
	}

	revealed := make([][]bool, h) // Step 3: Start revealed matrix
	for i := range revealed {
		revealed[i] = make([]bool, w)
	}
	counts := ComputeCounts(h, w, grid) // Step 4: Start bomb neighbor counts

	DrawBoard(h, w, revealed, grid, counts) // Draw initial hidden board
	fmt.Printf("Enter coordinates:\n")

	var moves int // Game state tracking
	totalSafe := h*w - bombs
	for { // Step 5: Main game loop
		var x, y int
		if _, err := fmt.Scanf("%d %d", &x, &y); err != nil { //// Read player move
			fmt.Printf("Invalid input.\nEnter coordinates:\n")
			continue
		}
		x--
		y--
		if !InBounds(h, w, x, y) || revealed[x][y] {
			fmt.Printf("Invalid input. (already revealed or out of bounds)\nEnter coordinates:\n")
			continue
		}
		moves++
		if grid[x][y] == '*' { // Step 6: Bomb hit (Game Over)
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if grid[i][j] == '*' {
						revealed[i][j] = true
					}
				}
			}
			DrawBoard(h, w, revealed, grid, counts)
			fmt.Printf("Game Over!\nYour statistics:\n- Field size: %dx%d\n", h, w)
			fmt.Printf("- Number of bombs: %d\n", bombs)
			fmt.Printf("- Number of moves: %d\n", moves)
			return
		}

		FloodReveal(h, w, x, y, revealed, counts) // Step 7: Reveal safe cells

		// Step 8: Check for win condition
		revealedCount := 0
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if revealed[i][j] && grid[i][j] != '*' {
					revealedCount++
				}
			}
		}
		if revealedCount == totalSafe { // print stats
			DrawBoard(h, w, revealed, grid, counts)
			fmt.Printf("You Win!\nYour statistics:\n- Field size: %dx%d\n", h, w)
			fmt.Printf("- Number of bombs: %d\n", bombs)
			fmt.Printf("- Number of moves: %d\n", moves)
			return
		}
		DrawBoard(h, w, revealed, grid, counts) // Step 9: Continue game
		fmt.Printf("Enter coordinates:\n")
	}
}
