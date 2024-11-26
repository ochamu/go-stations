package model

type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return "not found"
}

func NewErrNotFound() *ErrNotFound {
	return &ErrNotFound{}
}
