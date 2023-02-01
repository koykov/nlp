package nlp

import "errors"

var (
	ErrEmptyInput = errors.New("input text is empty")
	ErrBadVersion = errors.New("incompatible version")
)
