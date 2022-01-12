package nlp

type LanguageDetector interface {
	Detect(text []byte, def Language) Language
	DetectString(text string, def Language) Language
}
