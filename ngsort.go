package nlp

import "sort"

type ngsort interface {
	sort.Interface
	Each(func(int, appenderTo))
}

type appenderTo interface {
	AppendTo([]byte) []byte
}

// group: unisort

type unisort []Unigram

func appendUnisort(dst unisort, m map[Unigram]struct{}) unisort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (q unisort) Len() int {
	return len(q)
}

func (q unisort) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q *unisort) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q unisort) Each(fn func(int, appenderTo)) {
	for i := 0; i < len(q); i++ {
		fn(i, q[i])
	}
}

// endgroup: unisort

// group: bisort

type bisort []Bigram

func appendBisort(dst bisort, m map[Bigram]struct{}) bisort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (b bisort) Len() int {
	return len(b)
}

func (b bisort) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *bisort) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b bisort) Each(fn func(int, appenderTo)) {
	for i := 0; i < len(b); i++ {
		fn(i, b[i])
	}
}

// group: bisort

// endgroup: trisort

type trisort []Trigram

func appendTrisort(dst trisort, m map[Trigram]struct{}) trisort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (t trisort) Len() int {
	return len(t)
}

func (t trisort) Less(i, j int) bool {
	return t[i].a+t[i].b+t[i].c < t[j].a+t[j].b+t[j].c
}

func (t *trisort) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t trisort) Each(fn func(int, appenderTo)) {
	for i := 0; i < len(t); i++ {
		fn(i, t[i])
	}
}

// group: trisort

// endgroup: quadsort

type quadsort []Quadrigram

func appendQuadsort(dst quadsort, m map[Quadrigram]struct{}) quadsort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (q quadsort) Len() int {
	return len(q)
}

func (q quadsort) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q *quadsort) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q quadsort) Each(fn func(int, appenderTo)) {
	for i := 0; i < len(q); i++ {
		fn(i, q[i])
	}
}

// group: quadsort

// endgroup: fivesort

type fivesort []Fivegram

func appendFivesort(dst fivesort, m map[Fivegram]struct{}) fivesort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (f fivesort) Len() int {
	return len(f)
}

func (f fivesort) Less(i, j int) bool {
	return f[i].a+f[i].b+f[i].c+f[i].d+f[i].e < f[j].a+f[j].b+f[j].c+f[j].d+f[j].e
}

func (f *fivesort) Swap(i, j int) {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f fivesort) Each(fn func(int, appenderTo)) {
	for i := 0; i < len(f); i++ {
		fn(i, f[i])
	}
}

// endgroup: fivesort
