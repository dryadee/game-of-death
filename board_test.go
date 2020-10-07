package god

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBoard(t *testing.T) {
	assert := assert.New(t)

	dimension := 10

	b := EmptyBoard(dimension)

	assert.Equal(b.dimension, dimension)
	assert.Equal(len(b.board), dimension)
	for _, row := range b.board {
		assert.Equal(len(row), dimension)
	}
	assert.True(b.gen == 0)
	assert.NotNil(b)
	for _, row := range b.board {
		for _, col := range row {
			assert.True(col == 0)
		}
	}
}

func TestSet(t *testing.T) {
	assert := assert.New(t)

	b := EmptyBoard(10)

	for x, row := range b.board {
		for y := range row {
			assert.True(b.At(x, y) == 0)
		}
	}
	b.Set(5, 5, 1)
	assert.False(b.At(5, 5) == 0)
	assert.True(b.At(5, 5) == 1)
}

func TestAliveNeighbors(t *testing.T) {
	assert := assert.New(t)

	dim := 10
	b := EmptyBoard(dim)

	x := 5
	y := 5

	// Empty Board should have no alive neighbors at any cell
	aliveB, aliveG := b.AliveNeighbors(x, y)
	assert.True(aliveB == 0)
	assert.Equal(aliveB, aliveG)

	// case horizontal / vertical
	/*
		0 2 0
		1 1 1
		0 2 0
	*/
	b.Set(x, y, 1)
	b.Set(x-1, y, 1)
	b.Set(x+1, y, 1)
	b.Set(x, y-1, 2)
	b.Set(x, y+1, 2)

	aliveB, aliveG = b.AliveNeighbors(x, y)
	assert.True(aliveB == 2)
	assert.Equal(aliveB, aliveG)

	// kill color we check neighbors for
	// should have no impact on result

	aliveB2, aliveG2 := b.AliveNeighbors(x, y)
	assert.Equal(aliveB, aliveB2)
	assert.Equal(aliveG, aliveG2)

	// add diagonal
	/*
		1 2 2
		1 1 1
		1 2 2
	*/
	b.Set(x-1, y-1, 1)
	b.Set(x-1, y+1, 1)
	b.Set(x+1, y-1, 2)
	b.Set(x+1, y+1, 2)

	aliveB, aliveG = b.AliveNeighbors(x, y)
	assert.True(aliveB == 4)
	assert.Equal(aliveB, aliveG)
}
