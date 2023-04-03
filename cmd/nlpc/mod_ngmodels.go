package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/koykov/fastconv"
	"github.com/koykov/nlp"
	"github.com/koykov/nlp/cleaner"
	"github.com/koykov/nlp/tokenizer"
)

type ngmodelsModule struct{}

func (m ngmodelsModule) Validate(input, target string) error {
	if len(input) == 0 {
		return fmt.Errorf("param -input is required")
	}
	if isDirWR(target) {
		return fmt.Errorf("target '%s' isn't writable", target)
	}
	return nil
}

func (m ngmodelsModule) Compile(w moduleWriter, input, target string) (err error) {
	_, _ = w, target
	var resp *http.Response
	if resp, err = http.Get(input); err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()

	scanner := bufio.NewScanner(resp.Body)

	ctx := nlp.NewCtx[string]()
	model := nlp.NGModel[string]{}
	for scanner.Scan() {
		line := scanner.Bytes()
		ctx.Reset().
			SetText(fastconv.B2S(line)).
			WithCleaner(cleaner.Macros[string]{Left: "{", Right: "}"}).
			WithCleaner(cleaner.Macros[string]{Left: "[", Right: "]"}).
			WithCleaner(cleaner.Space[string]{}).
			WithCleaner(nlp.UnicodeCleaner[string]{Mask: nlp.DefaultCleanMask}).
			Clean().
			WithTokenizer(tokenizer.SpaceTokenizer[string]{}).
			Tokenize().
			GetTokens().
			Each(func(i int, t nlp.Token) {
				println(t.String())
				model.Parse(t.String())
			})
	}

	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = model.Write(f)

	return
}
