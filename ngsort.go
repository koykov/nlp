package nlp

type unisort []Unigram

func appendUnisort(dst unisort, m map[Unigram]struct{}) unisort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (b unisort) Len() int {
	return len(b)
}

func (b unisort) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *unisort) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

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

type trisort []Trigram

func appendTrisort(dst trisort, m map[Trigram]struct{}) trisort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (b trisort) Len() int {
	return len(b)
}

func (b trisort) Less(i, j int) bool {
	return b[i].a+b[i].b+b[i].c < b[j].a+b[j].b+b[j].c
}

func (b *trisort) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

type quadsort []Quadrigram

func appendQuadsort(dst quadsort, m map[Quadrigram]struct{}) quadsort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (b quadsort) Len() int {
	return len(b)
}

func (b quadsort) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b *quadsort) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

type fivesort []Fivegram

func appendFivesort(dst fivesort, m map[Fivegram]struct{}) fivesort {
	for n := range m {
		dst = append(dst, n)
	}
	return dst
}

func (b fivesort) Len() int {
	return len(b)
}

func (b fivesort) Less(i, j int) bool {
	return b[i].a+b[i].b+b[i].c+b[i].d+b[i].e < b[j].a+b[j].b+b[j].c+b[j].d+b[j].e
}

func (b *fivesort) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}
