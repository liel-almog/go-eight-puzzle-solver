package algorithm

import (
	"github.com/lielalmog/go-be-eight-puzzle-solver/board"
	"github.com/lielalmog/go-be-eight-puzzle-solver/stack"
)

type DfsSolver struct {
	InitialBoard *board.Board
}

func NewDfsSolver(b *board.Board) Solver {
	return &DfsSolver{
		InitialBoard: b,
	}
}

func (dSolver *DfsSolver) Solve(targetBoard []int) (board.TilesArray, error) {
	s := stack.NewSliceStack[board.Board]()
	s.Push(*dSolver.InitialBoard)

	var visited = make(map[string]bool)

	// The key is a string representation of the board
	// The value is a string representation of the board be
	var parent = make(map[string]string)

	initialBoardKey := arrayToString(dSolver.InitialBoard.GetTiles())
	parent[initialBoardKey] = ""

	var iterations int

	for !s.IsEmpty() {
		iterations++
		b, err := s.Pop()
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
				path := reconstructPath(*dSolver.InitialBoard, currBoardKey, parent)
				return path, nil
			}

			if _, found := visited[currBoardKey]; !found {
				s.Push(currBoard)
				parent[currBoardKey] = key
			}
		}
	}

	return nil, ErrNoSolution
}
