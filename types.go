package nlp

type Script uint

type ScriptScore struct {
	Script Script
	Score  float32
}

type ScriptProba []ScriptScore

func (s ScriptProba) Len() int {
	return len(s)
}

func (s ScriptProba) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s *ScriptProba) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
