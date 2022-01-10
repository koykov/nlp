package nlp

import "github.com/koykov/fastconv"

type Language int

func (l Language) String() string {
	lo, hi := __lt_lst[l].name.decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}

func (l Language) Native() string {
	lo, hi := __lt_lst[l].native.decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}

func (l Language) Iso6391() string {
	lo, hi := __lt_lst[l].iso1.decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}

func (l Language) Iso6393() string {
	lo, hi := __lt_lst[l].iso3.decode()
	return fastconv.B2S(__lt_buf[lo:hi])
}
