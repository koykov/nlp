package tokenizer

import (
	"github.com/koykov/byteseq"
	"github.com/koykov/nlp"
)

// group: SpaceTokenizer

type SpaceTokenizer[T byteseq.Byteseq] struct {
	nlp.StringTokenizer[T]
}

func NewSpaceTokenizer[T byteseq.Byteseq]() SpaceTokenizer[T] {
	return SpaceTokenizer[T]{
		StringTokenizer: nlp.StringTokenizer[T]{
			Separator:  " ",
			BlankLines: nlp.TokenizerBlankLinesKeep,
		},
	}
}

func (t SpaceTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t SpaceTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	t.Separator, t.BlankLines = " ", nlp.TokenizerBlankLinesKeep
	return t.StringTokenizer.AppendTokenize(dst, x)
}

// endgroup: SpaceTokenizer

// group: TabTokenizer

type TabTokenizer[T byteseq.Byteseq] struct {
	nlp.StringTokenizer[T]
}

func NewTabTokenizer[T byteseq.Byteseq]() TabTokenizer[T] {
	return TabTokenizer[T]{
		StringTokenizer: nlp.StringTokenizer[T]{
			Separator:  "\t",
			BlankLines: nlp.TokenizerBlankLinesKeep,
		},
	}
}

func (t TabTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t TabTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	t.Separator, t.BlankLines = "\t", nlp.TokenizerBlankLinesKeep
	return t.StringTokenizer.AppendTokenize(dst, x)
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
	nlp.StringTokenizer[T]
	BlankLines nlp.TokenizerBlankLines
}

func NewLineTokenizer[T byteseq.Byteseq](blankLines nlp.TokenizerBlankLines) LineTokenizer[T] {
	return LineTokenizer[T]{
		StringTokenizer: nlp.StringTokenizer[T]{
			Separator:  "\n",
			BlankLines: blankLines,
		},
	}
}

func (t LineTokenizer[T]) Tokenize(x T) nlp.Tokens {
	return t.AppendTokenize(nil, x)
}

func (t LineTokenizer[T]) AppendTokenize(dst nlp.Tokens, x T) nlp.Tokens {
	t.Separator, t.BlankLines = "\n", nlp.TokenizerBlankLinesKeep
	return t.StringTokenizer.AppendTokenize(dst, x)
}

// endgroup: LineTokenizer
