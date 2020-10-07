package god

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyBoard(t *testing.T, dimension int) {
	assert := assert.New(t)

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

	assert.True(b.board[5][5] == 0)
	b.Set(5, 5, 1)
	assert.False(b.board[5][5] == 0)
	assert.True(b.board[5][5] == 1)
}
