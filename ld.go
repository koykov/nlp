package nlp

type LanguageDetector interface {
	Detect(ctx *Ctx, text []byte, def Language) Language
	DetectString(ctx *Ctx, text string, def Language) Language
}
