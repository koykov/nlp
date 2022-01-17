package nlp

import "unicode"

type ScriptDetector interface {
	Detect(ctx *Ctx, text []byte) *unicode.RangeTable
	DetectString(ctx *Ctx, text string) *unicode.RangeTable

	DetectScore(ctx *Ctx, text []byte) ScriptDetectScores
	DetectScoreString(ctx *Ctx, text string) ScriptDetectScores
}
