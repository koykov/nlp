package nlp

type LanguageDetector interface {
	Detect(ctx *Ctx, text []byte) Language
	DetectString(ctx *Ctx, text string) Language

	DetectProba(ctx *Ctx, text []byte) []LanguageDetectProba
	DetectProbaString(ctx *Ctx, text string) []LanguageDetectProba
}
