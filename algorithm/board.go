package algorithm

import (
	"github.com/lielalmog/go-be-eight-puzzle-solver/errors/apperrors"
)

type Board struct {
	rowCount          int
	columnCount       int
	tiles             Tiles
	blankTilePosition Position
}

func NewBoardFromBoard(b *Board) *Board {
	var newB Board = *b

	copyTiles := make([][]int, len(b.tiles))
	for i := range b.tiles {
		copyTiles[i] = make([]int, len(b.tiles[i]))
		copy(copyTiles[i], b.tiles[i])
	}

	newB.tiles = copyTiles
	return &newB
}

func NewBoard(rowCount, columnCount int) (*Board, error) {
	if rowCount < 1 {
		return nil, apperrors.ErrRowCountSmallerThanOne
	}

	if columnCount < 1 {
		return nil, apperrors.ErrColumnCountSmallerThanOne
	}

	// Initliaze array of arrays
	tiles := make([][]int, rowCount)

	var shufflesArr []int = createRandomArray(columnCount * rowCount)

	// Initliaze each inner array
	for i := 0; i < rowCount; i++ {
		tiles[i] = make([]int, columnCount)
	}

	var blankTilePosition Position

	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			index := i*columnCount + j
			randomValue := shufflesArr[index]

			// If this is the blank tile set this value in the struct
			if randomValue == BoardBlankValue {
				blankTilePosition = Position{
					row:    i,
					column: j,
				}
			}

			tiles[i][j] = randomValue
		}
	}

	board := &Board{
		rowCount:          rowCount,
		columnCount:       columnCount,
		tiles:             tiles,
		blankTilePosition: blankTilePosition,
	}

	return board, nil
}

func (b *Board) IsValid() bool {
	var s map[int]bool = make(map[int]bool)

	for i := 0; i < b.rowCount; i++ {
		for j := 0; j < b.columnCount; j++ {
			v := b.tiles[i][j]
			s[v] = true
		}
	}

	return len(s) == b.columnCount*b.rowCount
}

func (b *Board) inversionCount() int {
	arr := convertTo1D(b.tiles)

	// Remove the blank tile
	index := getBlankTileIndex(b)
	arr = append(arr[:index], arr[index+1:]...)

	return countInversions(arr)
}

func (b *Board) IsSolvable() bool {
	invCount := b.inversionCount()

	if isOdd(b.columnCount) {
		return isEven(invCount)
	} else {
		blankRowFromBottom := b.rowCount - b.blankTilePosition.row

		if isEven(blankRowFromBottom) {
			return isOdd(invCount)
		} else {
			return isEven(invCount)
		}
	}
}

func (b *Board) IsSolved(targetArr []int) bool {
	currArr := convertTo1D(b.tiles)

	if len(currArr) != len(targetArr) {
		return false
	}

	for i := 0; i < len(currArr); i++ {
		if currArr[i] != targetArr[i] {
			return false
		}
	}

	return true
}

func (b *Board) isOutOfBounds(pos Position) bool {
	return (pos.column < 0 || pos.column >= b.columnCount) || (pos.row < 0 || pos.row >= b.rowCount)
}

func (b *Board) move(move MoveDirection) error {
	newPos := Position{
		row:    b.blankTilePosition.row + move.rowChange,
		column: b.blankTilePosition.column + move.columnChange,
	}

	if b.isOutOfBounds(newPos) {
		return apperrors.ErrInvalidMove
	}

	v := b.tiles[newPos.row][newPos.column]

	b.tiles[newPos.row][newPos.column] = BoardBlankValue
	b.tiles[b.blankTilePosition.row][b.blankTilePosition.column] = v

	b.blankTilePosition = newPos

	return nil
}
