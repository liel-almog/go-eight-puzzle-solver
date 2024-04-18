package algorithm

import "fmt"

type BfsSolver struct {
	InitialBaord *Board
}

func (bSolver *BfsSolver) Solve() ([]Board, error) {
	v := bSolver.neighbours(bSolver.InitialBaord)
	fmt.Println(len(v))
	return nil, nil
}

func (bSolver *BfsSolver) neighbours(b *Board) []Board {
	var adj []Board = make([]Board, 0)
	fmt.Println(b.tiles)

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
