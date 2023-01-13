package nlp

type Unigram uint16

type Bigram uint32

type Trigram struct {
	a, b, c Unigram
}

type Quadrigram uint64

type Fivegram struct {
	a, b, c, d, e Unigram
}
