package nlp

import "github.com/koykov/byteseq"

type Cleaner[T byteseq.Byteseq] interface {
	Clean(x T) []rune
	AppendClean(dst []rune, x T) []rune
}

type Modifier[T byteseq.Byteseq] interface {
	Modify(x T) T
	AppendModify(dst []rune, x T) []rune
}

type Tokenizer[T byteseq.Byteseq] interface {
	Tokenize(x T) Tokens
	AppendTokenize(dst Tokens, x T) Tokens
}

type ScriptDetector[T byteseq.Byteseq] interface {
	Detect(ctx *Ctx[T]) (Script, error)
	DetectProba(ctx *Ctx[T]) (ScriptProba, error)
}
