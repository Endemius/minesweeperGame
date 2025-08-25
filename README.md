# Minesweeper

A terminal Minesweeper game written in Go using only the standard library. You draw the minefield yourself, then open cells by typing their coordinates until you clear the board or hit a bomb.

## Features

- Custom boards from 3×3 up to 20×20, with bombs placed wherever you want
- Numbers showing how many bombs touch each cell, diagonals included
- Automatic opening of empty areas (flood fill), like in classic Minesweeper
- Input checks for board size, row length, allowed characters and bomb count
- All bombs revealed on game over
- End-of-game statistics: field size, number of bombs and number of moves

## How to run

You need [Go](https://go.dev/dl/) installed.

```bash
git clone https://github.com/Endemius/minesweeperGame.git
cd minesweeperGame/codes/minesweeper
go run main.go
```

## How to play

1. **Enter the board size** as `height width`. Both must be between 3 and 20.
2. **Draw the map** as `height` rows of exactly `width` characters each:
   - `.` is a safe cell
   - `*` is a bomb

   The map needs at least 2 bombs and at least one safe cell.
3. **Open cells** by typing `row column`. Both start at 1, and `1 1` is the top-left cell.

If you open a cell with no bombs around it, the game also opens every connected empty cell and its neighbours. You **win** when every safe cell is open and **lose** if you open a bomb.

### Board symbols

| Symbol | Meaning |
|---|---|
| `XXXXXXX` | Hidden cell |
| `1`–`8` | Number of bombs touching this cell |
| *(blank)* | Open cell with no bombs around it |
| `*` | Bomb, shown when the game ends |

## Example game

```
Enter height and width:
3 3
Enter the content of the map(. or *):
*..
...
..*
       1       2       3
    _______________________
   |XXXXXXX|XXXXXXX|XXXXXXX|
 1 |XXXXXXX|XXXXXXX|XXXXXXX|
   |XXXXXXX|XXXXXXX|XXXXXXX|
   |XXXXXXX|XXXXXXX|XXXXXXX|
 2 |XXXXXXX|XXXXXXX|XXXXXXX|
   |XXXXXXX|XXXXXXX|XXXXXXX|
   |XXXXXXX|XXXXXXX|XXXXXXX|
 3 |XXXXXXX|XXXXXXX|XXXXXXX|
   |XXXXXXX|XXXXXXX|XXXXXXX|
Enter coordinates:
1 3
       1       2       3
    _______________________
   |XXXXXXX|       |       |
 1 |XXXXXXX|   1   |       |
   |XXXXXXX|_______|_______|
   |XXXXXXX|       |       |
 2 |XXXXXXX|   2   |   1   |
   |XXXXXXX|_______|_______|
   |XXXXXXX|XXXXXXX|XXXXXXX|
 3 |XXXXXXX|XXXXXXX|XXXXXXX|
   |XXXXXXX|XXXXXXX|XXXXXXX|
Enter coordinates:
3 1
       1       2       3
    _______________________
   |XXXXXXX|       |       |
 1 |XXXXXXX|   1   |       |
   |XXXXXXX|_______|_______|
   |       |       |       |
 2 |   1   |   2   |   1   |
   |_______|_______|_______|
   |       |       |XXXXXXX|
 3 |       |   1   |XXXXXXX|
   |_______|_______|XXXXXXX|
You Win!
Your statistics:
- Field size: 3x3
- Number of bombs: 2
- Number of moves: 2
```

## Input checks

The game stops with an error if the map is invalid:

| Rule | Error message |
|---|---|
| Height and width must be numbers | `not a digit` |
| Height and width must be at least 3 | `the map is too small` |
| Height and width must be at most 20 | `the map is too big` |
| Each row must be exactly `width` characters | `rows don't match the height or width` |
| Only `.` and `*` are allowed | `only '.' or '*' allowed` |
| At least 2 bombs | `not enough bombs` |
| At least one safe cell | `too many bombs` |

During play, coordinates that are off the board or already open aren't counted as a move. The game asks for new coordinates instead.

## How it works

- **`ComputeCounts`** counts the bombs in the 8 neighbours of every cell once, at the start of the game.
- **`FloodReveal`** opens cells with a breadth-first search (BFS). It uses a queue, so large empty areas open without deep recursion. When it reaches a cell with a count of 0, it adds all of that cell's neighbours to the queue.
- **`DrawBoard`** prints the board with 7-character-wide cells and row and column numbers, so it's easy to read the coordinates.
- **Win check:** after each move, the game compares the number of open safe cells with the total number of safe cells.

## Project structure

```
codes/
└── minesweeper/
    ├── main.go                  # The full game: input, checks, game loop, drawing
    ├── board/
    │   └── board.go             # DrawBoard as its own package
    └── gamelogic/
        └── gamelogic.go         # InBounds, ComputeCounts, FloodReveal as their own package
```

`main.go` has its own copy of every function, so it runs as a single file. The `board` and `gamelogic` folders hold the same logic split into separate packages.
