package algorithm

import (
	"math/rand"
)

func createRandomArray(size int) []int {
	arr := make([]int, size)

	arr[0] = int(BoardBlankValue)

	for i := 1; i < len(arr); i++ {
		arr[i] = i
	}

	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })

	return arr
}

func convertTo1D(arr [][]int) []int {
	rowCount := len(arr)
	columnCount := len(arr[0])
	a := make([]int, rowCount*columnCount)

	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			a[i*columnCount+j] = arr[i][j]
		}
	}

	return a
}

func isOdd(n int) bool {
	return n%2 == 1
}

func isEven(n int) bool {
	return n%2 == 0
}

func GenerateTargetBoard(rowCount, columnCount int) []int {
	arr := make([]int, rowCount*columnCount)
	arr[0] = BoardBlankValue

	for i := 1; i < rowCount*columnCount; i++ {
		arr[i] = i
	}

	return arr
}

func getBlankTileIndex(b *Board) int {
	return b.blankTilePosition.row*b.columnCount + b.blankTilePosition.column
}

func cloneTiles(tiles Tiles) (Tiles, error) {
	rowCount := len(tiles)

	if rowCount < 1 {
		return nil, ErrEmptyTiles
	}
	columnCount := len(tiles[0])

	newTiles := make(Tiles, rowCount)

	return newTiles, nil
}
