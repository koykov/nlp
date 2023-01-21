package nlp

import "github.com/koykov/fastconv"

type Byteseq interface {
	~string | ~[]byte
}

func q2s[T Byteseq](x T) string {
	ix := any(x)
	switch ix.(type) {
	case string:
		return ix.(string)
	case []byte:
		return fastconv.B2S(ix.([]byte))
	default:
		return ""
	}
}

func q2b[T Byteseq](x T) []byte {
	ix := any(x)
	switch ix.(type) {
	case string:
		return fastconv.S2B(ix.(string))
	case []byte:
		return ix.([]byte)
	default:
		return nil
	}
}
