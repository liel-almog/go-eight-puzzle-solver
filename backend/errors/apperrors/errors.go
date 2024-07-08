package apperrors

import "errors"

// ErrRowCountSmallerThanOne is return when tring to create a board with less than one row
var ErrRowCountSmallerThanOne = errors.New("the number or rows must be greater then equal to one")

// ErrColumnCountSmallerThanOne is return when tring to create a board with less than one column
var ErrColumnCountSmallerThanOne = errors.New("the number or columns must be greater then equal to one")

var ErrInvalidMove = errors.New("can not move to this direction")
