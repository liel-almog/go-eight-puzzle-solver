package board

type MoveDirection struct {
	rowChange    int
	columnChange int
}

func GetDirections() []MoveDirection {
	directionLeft := MoveDirection{
		rowChange:    0,
		columnChange: -1,
	}

	directionRight := MoveDirection{
		rowChange:    0,
		columnChange: 1,
	}

	directionUp := MoveDirection{
		rowChange:    -1,
		columnChange: 0,
	}

	directionDown := MoveDirection{
		rowChange:    1,
		columnChange: 0,
	}

	return []MoveDirection{directionLeft, directionRight, directionDown, directionUp}
}
