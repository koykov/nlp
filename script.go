package nlp

import "unicode"

type Script uint

func ScriptsSupported() []Script {
	return __sl
}

// Evaluate checks if given rune r is written on script s.
// Use precompiled SRE (script rune evaluator) to speed up evaluation.
// See performance tests https://github.com/koykov/versus/blob/master/nlp_script/evaluate_test.go
func (s Script) Evaluate(r rune) bool {
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

func (s Script) Languages() []Language {
	if int(s) >= len(__sl_idx) {
		return nil
	}
	lo, hi := __sl_idx[s].Decode()
	if lo > hi || hi >= uint16(len(__sl_buf)) {
		return nil
	}
	return __sl_buf[lo:hi]
}
