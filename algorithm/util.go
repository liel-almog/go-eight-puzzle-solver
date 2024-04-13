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

func convertTo2D(arr []int, rowCount, columnCount int) [][]int {
	a := make([][]int, rowCount)

	for i := range a {
		a[i] = make([]int, columnCount)
	}

	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			a[i][j] = arr[i*columnCount+j]
		}
	}

	return a
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
