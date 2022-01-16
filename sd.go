package nlp

import "unicode"

type ScriptDetector interface {
	Detect(ctx *Ctx, text []byte) *unicode.RangeTable
	DetectString(ctx *Ctx, text string) *unicode.RangeTable

	DetectProba(ctx *Ctx, text []byte) []*unicode.RangeTable
	DetectProbaString(ctx *Ctx, text string) []*unicode.RangeTable
}
