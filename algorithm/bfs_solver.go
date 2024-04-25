package algorithm

import (
	"github.com/lielalmog/go-be-eight-puzzle-solver/board"
	"github.com/lielalmog/go-be-eight-puzzle-solver/queue"
)

type BfsSolver struct {
	InitialBoard *board.Board
}

func NewBfsSolver(b *board.Board) Solver {
	return &BfsSolver{
		InitialBoard: b,
	}
}

func (bSolver *BfsSolver) Solve(targetBoard []int) ([][][]int, error) {
	q := queue.NewSliceQueue[board.Board]()
	q.Enqueue(*bSolver.InitialBoard)

	var visited = make(map[string]bool)

	// The key is a string representation of the board
	// The value is a string representation of the board be
	var parent = make(map[string]string)

	initialBoardKey := arrayToString(bSolver.InitialBoard.GetTiles())
	parent[initialBoardKey] = ""

	var iterations int

	for !q.IsEmpty() {
		iterations++
		b, err := q.Dequeue()
		if err != nil {
			return nil, err
		}

		// Set this board as visited
		key := arrayToString(b.GetTiles())
		visited[key] = true

		n, err := b.Neighbours()
		if err != nil {
			return nil, err
		}

		for _, currBoard := range n {
			currBoardKey := arrayToString(currBoard.GetTiles())

			if currBoard.IsSolved(targetBoard) {
				parent[currBoardKey] = key
				path := reconstructPath(*bSolver.InitialBoard, currBoardKey, parent)
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
