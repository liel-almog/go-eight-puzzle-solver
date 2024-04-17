package algorithm

import (
	"github.com/jinzhu/copier"
)

type BfsSolver struct {
	InitialBaord *Board
}

func (bSolver *BfsSolver) Solve() ([]Board, error) {
	bSolver.neighbours(bSolver.InitialBaord)

	return nil, nil
}

func (bSolver *BfsSolver) neighbours(b *Board) []Board {
	var adj []Board = make([]Board, 0)

	directions := GetDirections()
	for i := 0; i < len(directions); i++ {
		d := directions[i]

		var newBoard Board
		copier.Copy(&newBoard, b)

		if err := newBoard.move(d); err != nil {
			adj = append(adj, newBoard)
		}
	}

	return adj
}
