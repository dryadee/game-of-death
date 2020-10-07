package god

import (
	"math/rand"
	"time"
)

const (
	NUM_CELLS int = 10
	NONE Color = 0x0
	BLUE Color = 0x1
	GREEN Color = 0x2
)

type Color byte

type Config struct {
	// Defaults could also be provided like this
	// Birth int `default:"3"`
	// See https://github.com/creasty/defaults
	Birth int
	Isolation int
	Overpopulation int
}

type Board struct {
	board [][]Color
	dimension int
	gen uint64
}


var (
	board *Board
	config *Config
)

// Config with default values
func DefaultConfig() *Config {
	return &Config{Birth: 3, Isolation: 2, Overpopulation: 3}
}

// Constructs an empty board
func EmptyBoard(dim int) *Board {
	board := make([][]Color, dim)
	for i:=0; i < dim; i++ {
		board[i] = make([]Color, dim)
	}
	return &Board{board: board, dimension: dim, gen: 0}
}

// Randomizes the value of each cell in the board
func (b *Board) RandomizeBoard() {
	rand.Seed(time.Now().UnixNano())

	n := b.dimension

	for x:=0; x < n; x++ {
		for y:=0; y < n; y++ {
			if rand.Intn((n * n) / NUM_CELLS) == 1 {
				if rand.Intn(2) == 0 {
					b.Set(x, y, BLUE) // Blue
				} else {
					b.Set(x, y, GREEN) // Green
				}
			}
		}
	}
}

// Set value of Cell at (x, y)
func (b *Board) Set(x int, y int, p Color) Color {
	replaced := b.At(x, y)
	b.board[x][y] = p
	return replaced
}

func (b *Board) Invert(x int, y int) {
	c := b.At(x, y)
	if c == NONE {
		return
	}
	b.Set(x, y, 3 - c)
}

// Returns value of Cell at (x, y)
func (b *Board) At(x int, y int) Color {
	return b.board[x][y]
}

// Returns number of blue and green neighbors for Cell at (x, y)
func (b *Board) AliveNeighbors(x int, y int) (int, int) {
	aliveBlue := 0
	aliveGreen := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if !OutOfBounds(x+i, y+j) {
				if b.At(x+i, y+j) == BLUE {
					aliveBlue++
				} else if b.At(x+i, y+j) == GREEN {
					aliveGreen++
				}
			}
		}
	}

	return aliveBlue, aliveGreen
}

// Advance Generation by applying modified GoL rules
func (b *Board) Step() {

	n := b.dimension
	next := EmptyBoard(n)
	next.gen = b.gen + 1

	for x:=0; x < n; x++ {
		for y:=0; y < n; y++ {
			bc, gc := b.AliveNeighbors(x, y)
			total := bc+gc
			if value := b.At(x, y); value == NONE && total == config.Birth {
				if bc > gc {
					next.Set(x, y, BLUE)
				} else {
					next.Set(x, y, GREEN)
				}
			} else if value == BLUE || value == GREEN{
				switch {
					case total > config.Isolation && total < config.Overpopulation:
						if bc > gc {
							next.Set(x, y, BLUE)
						} else {
							next.Set(x, y, GREEN)
						}
					default:
						next.Set(x, y, NONE)
				}
			}

		}
	}
	board = next
}

// Helper for AliveNeighbors
func OutOfBounds(x int, y int) bool {
	return (x < 0 || y < 0)
}






