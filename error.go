package nlp

import "errors"

var (
	ErrEmptyInput    = errors.New("input text is empty")
	ErrUnknownScript = errors.New("unknown script")
)
