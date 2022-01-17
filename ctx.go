package nlp

type Ctx struct {
	sd ScriptDetector
	ld LanguageDetector

	BufSDS ScriptDetectScores
	BufLDS LanguageDetectScores
}

func NewCtx() *Ctx {
	ctx := Ctx{}
	return &ctx
}

func (c *Ctx) Reset() {
	c.BufSDS = c.BufSDS[:0]
	c.BufLDS = c.BufLDS[:0]
}
