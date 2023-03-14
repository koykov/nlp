package modifier

import "github.com/koykov/byteseq"

type stage[T byteseq.Byteseq] struct {
	key string
	src T
	exp string
}
