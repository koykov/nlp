package nlp

type unigram uint16

type bigram uint32

type trigram struct {
	a, b, c unigram
}

type quadrigram uint64

type fivegram struct {
	a, b, c, d, e unigram
}
