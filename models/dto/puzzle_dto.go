package dto

type BoardDimensionsDTO struct {
	RowCount    int `json:"rowCount" validate:"required,min=2,max=8"`
	ColumnCount int `json:"columnCount" validate:"required,min=2,max=8"`
}

type TilesDTO struct {
	Tiles [][]int `json:"tiles" validate:"required,min=2,dive,required,min=2"`
}
