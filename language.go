package nlp

type Language int

const (
	Unknown Language = iota
	English
	Russian
)

func (l Language) String() string {
	return ""
}
