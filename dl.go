package nlp

import (
	"sort"

	"github.com/koykov/fastconv"
)

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
	return DetectLanguageString(ctx, fastconv.B2S(text))
}

func DetectLanguageString(ctx *Ctx, text string) (Language, error) {
	if err := dlProba(ctx, text); err != nil {
		return 0, err
	}
	var (
		mx float32
		mi int
	)
	_ = ctx.BufLP[len(ctx.BufLP)-1]
	for i := 0; i < len(ctx.BufLP); i++ {
		if score := ctx.BufLP[i].Score; score > mx {
			mx, mi = score, i
		}
	}
	return ctx.BufLP[mi].Language, nil
}

func DetectLanguageProba(ctx *Ctx, text []byte) (LanguageProba, error) {
	return DetectLanguageStringProba(ctx, fastconv.B2S(text))
}

func DetectLanguageStringProba(ctx *Ctx, text string) (LanguageProba, error) {
	if err := dlProba(ctx, text); err != nil {
		return nil, err
	}
	sort.Sort(&ctx.BufLP)
	return ctx.BufLP, nil
}

func dlProba(ctx *Ctx, text string) error {
	_, _ = ctx, text
	// todo implement me
	return nil
}
