package nlp

type LanguageDetector interface {
	Detect(text, def string) Language
	DetectName(text, def string) string
}
