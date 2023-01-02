package nlp

import "errors"

var (
	ErrEmptyInput       = errors.New("input text is empty")
	ErrNoScriptDetector = errors.New("no script detector provided")
	ErrUnknownScript    = errors.New("unknown script")
)
