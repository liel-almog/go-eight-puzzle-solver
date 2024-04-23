package algorithm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lielalmog/go-be-eight-puzzle-solver/board"
)

func arrayToString(a [][]int) string {
	var builder strings.Builder
	for i, row := range a {
		if i > 0 {
			builder.WriteString("|") // Use '|' as a row delimiter
		}

		for j, val := range row {
			if j > 0 {
				builder.WriteString(",") // Use ',' as a column delimiter
			}

			fmt.Fprintf(&builder, "%d", val)
		}
	}

	return builder.String()
}

func stringToArray(rowCount, columnCount int, s string) [][]int {
	var tiles [][]int = make([][]int, rowCount)

	for i := 0; i < rowCount; i++ {
		tiles[i] = make([]int, columnCount)
	}

	rows := strings.Split(s, "|")
	if len(rows) != rowCount {
		panic("can not transform string to array")
	}

	for i := 0; i < rowCount; i++ {
		row := strings.Split(rows[i], ",")
		if len(row) != columnCount {
			panic("can not transform string to array")
		}

		for j := 0; j < columnCount; j++ {
			n, err := strconv.Atoi(row[j])
			if err != nil {
				panic("can not transform string to array")
			}

			tiles[i][j] = n
		}
	}

	return tiles
}

func reconstructPath(initialBoard board.Board, state string, stateMap map[string]string) board.TilesArray {
	var sPath []string
	for step := state; step != ""; step = stateMap[step] {
		sPath = append(sPath, step)
	}

	var path = make([][][]int, len(sPath))

	for i := 0; i < len(sPath); i++ {
		rowCount := initialBoard.GetRowCount()
		columnCount := initialBoard.GetColumnCount()
		t := stringToArray(rowCount, columnCount, sPath[i])

		path[i] = t
	}

	// Perform an In-Place Array Reversal
	// We swap the first with the last, the second with the second last, etc...
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
