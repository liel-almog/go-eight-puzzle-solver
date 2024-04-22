package algorithm

type Tiles = [][]int
type TilesArray = []Tiles

type Solver interface {
	Solve(targetBoard []int) (TilesArray, error)
}
