package nlp

import "unicode"

type ScriptDetector interface {
	Detect(ctx *Ctx, text []byte) *unicode.RangeTable
	DetectString(ctx *Ctx, text string) *unicode.RangeTable

	DetectScore(ctx *Ctx, text []byte) ScriptDetectProba
	DetectProbaString(ctx *Ctx, text string) ScriptDetectProba
}

type ScriptDetectScore struct {
	Script *unicode.RangeTable
	Score  float32
}

type ScriptDetectProba []ScriptDetectScore

func (s ScriptDetectProba) Len() int {
	return len(s)
}

func (s ScriptDetectProba) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s *ScriptDetectProba) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
