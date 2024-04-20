package algorithm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lielalmog/go-be-eight-puzzle-solver/queue"
)

type BfsSolver struct {
	InitialBoard *Board
}

func NewBfsSolver(b *Board) Solver {
	return &BfsSolver{
		InitialBoard: b,
	}
}

func (bSolver *BfsSolver) Solve(targetBoard []int) ([][][]int, error) {
	q := queue.NewSliceQueue[Board]()
	q.Enqueue(*bSolver.InitialBoard)

	var visited = make(map[string]bool)

	// The key is a string representation of the board
	// The value is a string representation of the board be
	var parent = make(map[string]string)

	initialBoardKey := bSolver.arrayToString(bSolver.InitialBoard.tiles)
	parent[initialBoardKey] = ""

	var iterations int

	for !q.IsEmpty() {
		iterations++
		b, err := q.Dequeue()

		// Set this board as visited
		key := bSolver.arrayToString(b.tiles)
		visited[key] = true

		if err != nil {
			return nil, err
		}

		n := bSolver.neighbours(&b)

		for _, currBoard := range n {
			currBoardKey := bSolver.arrayToString(currBoard.tiles)

			if currBoard.IsSolved(targetBoard) {
				parent[currBoardKey] = key
				path := bSolver.reconstructPath(currBoardKey, parent)
				return path, nil
			}

			if _, found := visited[currBoardKey]; !found {
				q.Enqueue(currBoard)
				parent[currBoardKey] = key
			}
		}
	}

	return nil, ErrNoSolution
}

func (bSolver *BfsSolver) neighbours(b *Board) []Board {
	var adj []Board = make([]Board, 0, 4)

	directions := GetDirections()
	for i := 0; i < len(directions); i++ {
		d := directions[i]
		var newBoard Board = *NewBoardFromBoard(b)

		if err := newBoard.move(d); err == nil {
			adj = append(adj, newBoard)
		}
	}

	return adj
}

func (bSolver *BfsSolver) reconstructPath(state string, stateMap map[string]string) [][][]int {
	var sPath []string
	for step := state; step != ""; step = stateMap[step] {
		sPath = append(sPath, step)
	}

	var path = make([][][]int, len(sPath))

	for i := 0; i < len(sPath); i++ {
		t := bSolver.stringToArray(sPath[i])

		path[i] = t
	}

	// Perform an In-Place Array Reversal
	// We swap the first with the last, the second with the second last, etc...
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (bSolver *BfsSolver) arrayToString(a [][]int) string {
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

func (bSolver *BfsSolver) stringToArray(s string) [][]int {
	rowCount := bSolver.InitialBoard.rowCount
	columnCount := bSolver.InitialBoard.columnCount

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
