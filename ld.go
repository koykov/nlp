package nlp

type LanguageDetector interface {
	Detect(ctx *Ctx, text []byte) Language
	DetectString(ctx *Ctx, text string) Language

	DetectScore(ctx *Ctx, text []byte) LanguageDetectProba
	DetectProbaString(ctx *Ctx, text string) LanguageDetectProba
}

type LanguageDetectScore struct {
	Language Language
	Score    float32
}

type LanguageDetectProba []LanguageDetectScore

func (s LanguageDetectProba) Len() int {
	return len(s)
}

func (s LanguageDetectProba) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s *LanguageDetectProba) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
