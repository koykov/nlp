package tokenize

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

// group: SpaceTokenizer

type SpaceTokenizer[T byteseq.Byteseq] struct {
	o bool
	t nlp.StringTokenizer[T]
}

func NewSpaceTokenizer[T byteseq.Byteseq]() SpaceTokenizer[T] {
	return SpaceTokenizer[T]{
		t: nlp.StringTokenizer[T]{
			Separator:  " ",
			BlankLines: nlp.TokenizerBlankLinesKeep,
		},
		o: true,
	}
}

func (t SpaceTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.t.AppendTokenize(nil, x)
}

func (t SpaceTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.t.Separator = " "
		t.t.BlankLines = nlp.TokenizerBlankLinesKeep
		t.o = true
	}
	return t.t.AppendTokenize(dst, x)
}

// endgroup: SpaceTokenizer

// group: TabTokenizer

type TabTokenizer[T byteseq.Byteseq] struct {
	o bool
	t nlp.StringTokenizer[T]
}

func NewTabTokenizer[T byteseq.Byteseq]() TabTokenizer[T] {
	return TabTokenizer[T]{
		t: nlp.StringTokenizer[T]{
			Separator:  "\t",
			BlankLines: nlp.TokenizerBlankLinesKeep,
		},
		o: true,
	}
}

func (t TabTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.t.AppendTokenize(nil, x)
}

func (t TabTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.t.Separator = "\t"
		t.t.BlankLines = nlp.TokenizerBlankLinesKeep
		t.o = true
	}
	return t.t.AppendTokenize(dst, x)
}

// endgroup: TabTokenizer

// group: CharTokenizer

type CharTokenizer[T byteseq.Byteseq] struct{}

func NewCharTokenizer[T byteseq.Byteseq]() CharTokenizer[T] {
	return CharTokenizer[T]{}
}

func (t CharTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t CharTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	for i := 0; i < len(x); i++ {
		dst = append(dst, nlp.ParseToken(x, i, i+1))
	}
	return dst
}

// endgroup: CharTokenizer

// group: LineTokenizer

type LineTokenizer[T byteseq.Byteseq] struct {
	BlankLines nlp.TokenizerBlankLines

	o bool
	t nlp.StringTokenizer[T]
}

func NewLineTokenizer[T byteseq.Byteseq](blankLines nlp.TokenizerBlankLines) LineTokenizer[T] {
	return LineTokenizer[T]{
		t: nlp.StringTokenizer[T]{
			Separator:  "\n",
			BlankLines: blankLines,
		},
		o: true,
	}
}

func (t LineTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.t.AppendTokenize(nil, x)
}

func (t LineTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	if !t.o {
		t.t.Separator = "\n"
		t.t.BlankLines = t.BlankLines
		t.o = true
	}
	return t.t.AppendTokenize(dst, x)
}

// endgroup: LineTokenizer
