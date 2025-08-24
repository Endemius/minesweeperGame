package gamelogic

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
