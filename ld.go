package nlp

type LanguageDetector interface {
	Detect(ctx *Ctx, text []byte) Language
	DetectString(ctx *Ctx, text string) Language

	DetectScore(ctx *Ctx, text []byte) LanguageDetectScores
	DetectScoreString(ctx *Ctx, text string) LanguageDetectScores
}
