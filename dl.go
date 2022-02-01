package nlp

import "github.com/koykov/fastconv"

type LanguageScore struct {
	Language Language
	Score    float32
}

type LanguageProba []LanguageScore

func (s LanguageProba) Len() int {
	return len(s)
}

func (s LanguageProba) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s *LanguageProba) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func DetectLanguage(ctx *Ctx, text []byte) (Language, error) {
	// ...
	return 0, nil
}

func DetectLanguageString(ctx *Ctx, text string) (Language, error) {
	return DetectLanguage(ctx, fastconv.S2B(text))
}

func DetectLanguageProba(ctx *Ctx, text []byte) (LanguageProba, error) {
	// ...
	return nil, nil
}

func DetectLanguageStringProba(ctx *Ctx, text string) (LanguageProba, error) {
	return DetectLanguageProba(ctx, fastconv.S2B(text))
}
