package nlp

import "unicode"

type Script uint

func (s Script) Evaluate(p []byte) bool {
	// todo implement
	_ = p
	return false
}

func (s Script) EvaluateString(p string) bool {
	// todo implement
	_ = p
	return false
}

func (s Script) EvaluateRunes(p []rune) bool {
	// todo implement
	_ = p
	return false
}

func (s Script) EvaluateByte(b byte) bool {
	// todo implement
	_ = b
	return false
}

// EvaluateRune checks if given rune r is written on script s.
// Use precompiled SRE (script rune evaluator) to speed up evaluation.
// See performance tests https://github.com/koykov/versus/blob/master/nlp_script/evaluate_test.go
func (s Script) EvaluateRune(r rune) bool {
	if int(s) >= len(__sre_buf) {
		return false
	}
	return __sre_buf[s].Evaluate(r)
}

func (s Script) Is(r rune) bool {
	if int(s) >= len(__sre_buf) {
		return false
	}
	return unicode.Is(__sre_buf[s].t, r)
}
