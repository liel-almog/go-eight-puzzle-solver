package algorithm

import (
	"math"

	"github.com/lielalmog/go-eight-puzzle-solver/board"
	"github.com/lielalmog/go-eight-puzzle-solver/heap"
)

type AStarSolver struct {
	InitialBoard *board.Board
}

type aStarNode struct {
	Board board.Board
	Cost  float64 // Cost from start node to this node
	Heur  float64 // Heuristic cost from this node to the target node
	Total float64 // Total cost (Cost + Heur)
}

func NewAStarSolver(b *board.Board) Solver {
	return &AStarSolver{
		InitialBoard: b,
	}
}

// This is one way to calculate heuristic function
// Another way that is less accurate is to calculate the number of misplaced tiles
func (aStar *AStarSolver) manhattanHeuristic(b *board.Board) int {
	distance := 0
	t := b.GetTiles()

	// Create a map to store the goal positions for quick lookup
	// This can be done in the same loop but is less readable

	goalPositions := make(map[int][2]int)
	for i := 0; i < b.GetRowCount(); i++ {
		for j := 0; j < b.GetColumnCount(); j++ {
			// Noramally this function will use the target tiles to calculate the index to the target array
			// But because we have a coversion function this is not necessary
			index := board.IndexConversion(b.GetColumnCount(), i, j)
			goalPositions[index] = [2]int{i, j}
		}
	}

	for x := 0; x < b.GetRowCount(); x++ {
		for y := 0; y < b.GetColumnCount(); y++ {
			value := t[x][y]
			if value != board.BoardBlankValue { // Skip the blank/empty space
				goalPos := goalPositions[value]
				distance += int(math.Abs(float64(x-goalPos[0])) + math.Abs(float64(y-goalPos[1])))
			}
		}
	}

	return distance
}

type PriorityQueue []*aStarNode

func (aStar *AStarSolver) Solve(targerTiles []int) (board.TilesArray, error) {
	less := func(a, b aStarNode) bool {
		return a.Total < b.Total
	}

	h := heap.NewSliceHeap[aStarNode](less)

	initialHeur := aStar.manhattanHeuristic(aStar.InitialBoard)

	err := h.Push(aStarNode{
		Board: *aStar.InitialBoard,
		Cost:  0,
		Heur:  float64(initialHeur),
		Total: float64(initialHeur),
	})
	if err != nil {
		return nil, err
	}

	var visited = make(map[string]bool)

	var parent = make(map[string]string)

	for !h.IsEmpty() {
		curr, err := h.Pop()
		if err != nil {
			return nil, err
		}

		key := arrayToString(curr.Board.GetTiles())
		visited[key] = true

		n, err := curr.Board.Neighbours()
		if err != nil {
			return nil, err
		}

		for _, currBoard := range n {
			currBoardKey := arrayToString(currBoard.GetTiles())

			if currBoard.IsSolved(targerTiles) {
				parent[currBoardKey] = key
				path := reconstructPath(*aStar.InitialBoard, currBoardKey, parent)
				return path, nil
			}

			if _, found := visited[currBoardKey]; !found {
				c := curr.Cost + 1
				heur := aStar.manhattanHeuristic(&currBoard)
				node := aStarNode{
					Board: currBoard,
					Cost:  c,
					Heur:  float64(heur),
					Total: c + float64(heur),
				}

				err := h.Push(node)
				if err != nil {
					return nil, err
				}

				parent[currBoardKey] = key
			}
		}
	}

	return nil, ErrNoSolution
}
