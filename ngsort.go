package nlp

type unisort []Unigram

func toUnisort(m map[Unigram]struct{}) unisort {
	b := make(unisort, 0, len(m))
	for n := range m {
		b = append(b, n)
	}
	return b
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

func toBisort(m map[Bigram]struct{}) bisort {
	b := make(bisort, 0, len(m))
	for n := range m {
		b = append(b, n)
	}
	return b
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

func toTrisort(m map[Trigram]struct{}) trisort {
	b := make(trisort, 0, len(m))
	for n := range m {
		b = append(b, n)
	}
	return b
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

func toQuadsort(m map[Quadrigram]struct{}) quadsort {
	b := make(quadsort, 0, len(m))
	for n := range m {
		b = append(b, n)
	}
	return b
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

func toFivesort(m map[Fivegram]struct{}) fivesort {
	b := make(fivesort, 0, len(m))
	for n := range m {
		b = append(b, n)
	}
	return b
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
