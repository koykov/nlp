package nlp

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestNGModel(t *testing.T) {
	t.Run("parse single", func(t *testing.T) {
		s := "foobar"
		m := NGModel[string]{}
		m.Parse(s)
		if u, b, t_, q, f := m.Stat(); u != 5 || b != 5 || t_ != 4 || q != 3 || f != 2 {
			t.FailNow()
		}
	})
	t.Run("parse multi", func(t *testing.T) {
		s := "The quick brown fox jumped over the lazy dogs"
		m := NGModel[string]{}
		m.Parse(s)
		if u, b, t_, q, f := m.Stat(); u != 26 || b != 26 || t_ != 18 || q != 10 || f != 4 {
			t.FailNow()
		}
	})
	t.Run("write", func(t *testing.T) {
		assertFile := func(a, b string) (err error) {
			var af, bf []byte
			if af, err = os.ReadFile(a); err != nil {
				return err
			}
			if bf, err = os.ReadFile(b); err != nil {
				return err
			}
			if !bytes.Equal(af, bf) {
				return fmt.Errorf("files %s and %s are not equal", a, b)
			}
			return nil
		}

		ctx := AcquireCtx[string]()
		defer ReleaseCtx[string](ctx)

		m := NGModel[string]{}
		ctx.SetText("Hello World!").
			Clean().
			Tokenize().
			GetTokens().
			Each(func(i int, t Token) {
				m.Parse(t.String())
			})
		fn := fmt.Sprintf("testdata/%d.ngm", time.Now().UnixNano())
		f, err := os.Create(fn)
		if err != nil {
			t.Fatal(err)
		}
		if _, err = m.Write(f); err != nil {
			t.Fatal(err)
		}
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
		if err = assertFile(fn, "testdata/example.ngm"); err != nil {
			t.Error(err)
		}
	})
	t.Run("read", func(t *testing.T) {
		m := NGModel[string]{Version: 0}
		f, err := os.Open("testdata/example.ngm")
		if err != nil {
			t.Fatal(err)
		}
		if _, err = m.Read(f); err != nil {
			t.Fatal(err)
		}
		if err = f.Close(); err != nil {
			t.Fatal(err)
		}
		if u, b, t_, q, f := m.Stat(); u != 7 || b != 8 || t_ != 6 || q != 4 || f != 2 {
			t.FailNow()
		}
	})
}
