package nlp

type LanguageDetectScore struct {
	Language Language
	Score    float32
}

type LanguageDetectScores []LanguageDetectScore

func (s LanguageDetectScores) Len() int {
	return len(s)
}

func (s LanguageDetectScores) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func (s *LanguageDetectScores) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
