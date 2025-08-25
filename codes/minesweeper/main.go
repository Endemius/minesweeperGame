package main

import (
	"fmt"
	"strconv"
	"strings"
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

func InBounds(h, w, x, y int) bool { // InBounds checks whether the coordinates (x, y) are inside the board boundaries.
	return x >= 0 && x < h && y >= 0 && y < w
}

func ComputeCounts(h, w int, grid [][]rune) [][]int { // ComputeCounts calculates the number of bombs around each cell. It returns a 2D slice (same size as grid) where each cell contains the count of adjacent bombs (including diagonals).
	counts := make([][]int, h)
	for i := range counts {
		counts[i] = make([]int, w)
	}
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1} // Directions for the 8 possible neighbors (top-left, top, top-right, etc.).
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '*' { // Skip bomb cells — we don't need to count for them.
				continue
			}
			cnt := 0 // Count bombs in all 8 neighboring cells
			for k := 0; k < 8; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if InBounds(h, w, ni, nj) && grid[ni][nj] == '*' {
					cnt++
				}
			}
			counts[i][j] = cnt
		}
	}
	return counts
}

func FloodReveal(h, w, x, y int, revealed [][]bool, counts [][]int) { // FloodReveal reveals a cell and, if it's empty (0 bombs around), recursively reveals all connected empty cells and their neighbors (like Minesweeper flood fill).
	if !InBounds(h, w, x, y) || revealed[x][y] {                      // If the cell is outside the board or already revealed, stop.
		return
	}
	queue := [][2]int{{x, y}} // Use a queue to store cells to reveal (BFS flood fill).
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for len(queue) > 0 {
		cur := queue[0] // Take the first cell from the queue.
		queue = queue[1:]
		cx, cy := cur[0], cur[1]
		if revealed[cx][cy] {
			continue
		}
		revealed[cx][cy] = true  // Reveal the current cell.
		if counts[cx][cy] == 0 { // If this cell has no bombs nearby, add all neighbors to the queue.
			for k := 0; k < 8; k++ {
				ni, nj := cx+dx[k], cy+dy[k]
				if InBounds(h, w, ni, nj) && !revealed[ni][nj] {
					queue = append(queue, [2]int{ni, nj})
				}
			}
		}
	}
}

// helper to center text in a fixed-width cell
func CenterText(s string, width int) string {
	if len(s) >= width {
		return s
	}
	left := (width - len(s)) / 2
	right := width - len(s) - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

// DrawBoard renders the Minesweeper board on the screen
func DrawBoard(h, w int, revealed [][]bool, grid [][]rune, counts [][]int) {
	cellW := 7

	// Column numbers
	fmt.Printf("   ")
	for j := 0; j < w; j++ {
		fmt.Printf(" ")
		numStr := strconv.Itoa(j + 1) // simpler than manual digit extraction
		fmt.Printf("%s", CenterText(numStr, cellW))
	}
	fmt.Printf("\n")

	// Top line
	fmt.Printf("    ")
	for i := 0; i < w*8-1; i++ {
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
				fmt.Printf("%s", CenterText("", cellW))
			} else {
				fmt.Printf("%s", CenterText("XXXXXXX", cellW))
			}
		}
		fmt.Printf("|\n")

		// content line
		if i+1 < 10 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d ", i+1)
		for j := 0; j < w; j++ {
			fmt.Printf("|")
			if revealed[i][j] {
				if grid[i][j] == '*' {
					fmt.Printf("   *   ")
				} else if counts[i][j] == 0 {
					fmt.Printf("       ")
				} else {
					numStr := strconv.Itoa(counts[i][j])
					fmt.Printf("%s", CenterText(numStr, cellW))
				}
			} else {
				fmt.Printf("XXXXXXX")
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
