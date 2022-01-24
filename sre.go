package nlp

import "unicode"

const (
	// see https://github.com/golang/go/blob/master/src/unicode/letter.go#L84
	sreLinearMax = 18
)

// SRE ...
// Nested functions approach performance https://github.com/koykov/lab/tree/master/call_perf
type SRE struct {
	Evaluate func(r rune) bool
}

func sreEvalBinary16(ranges []unicode.Range16, r uint16) bool {
	lo := 0
	hi := len(ranges)
	for lo < hi {
		m := lo + (hi-lo)/2
		rn := &ranges[m]
		if rn.Lo <= r && r <= rn.Hi {
			return rn.Stride == 1 || (r-rn.Lo)%rn.Stride == 0
		}
		if r < rn.Lo {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return false
}

func sreEvalBinary32(ranges []unicode.Range32, r uint32) bool {
	lo := 0
	hi := len(ranges)
	for lo < hi {
		m := lo + (hi-lo)/2
		rn := ranges[m]
		if rn.Lo <= r && r <= rn.Hi {
			return rn.Stride == 1 || (r-rn.Lo)%rn.Stride == 0
		}
		if r < rn.Lo {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return false
}
