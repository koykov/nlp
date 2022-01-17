package nlp

import (
	"unicode"
)

type ScriptDetectScore struct {
	Script *unicode.RangeTable
	Score  float32
}

type ScriptDetectScores []ScriptDetectScore

func (s ScriptDetectScores) Len() int {
	return len(s)
}

func (s ScriptDetectScores) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s *ScriptDetectScores) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
