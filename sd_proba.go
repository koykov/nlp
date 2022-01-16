package nlp

import "unicode"

type ScriptDetectProba struct {
	Script *unicode.RangeTable
	Proba  float32
}
