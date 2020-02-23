package main

import (
	"sort"
	"strings"
)

type Pair struct {
	key   string
	value int
}

type PairList []Pair

func (p PairList) Less(i, j int) bool {
	return p[i].value < p[j].value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Len() int {
	return len(p)
}

func DefCounter(str string) PairList {
	words := strings.Fields(str)

	dictionary := map[string]int{}
	for _, v := range words {
		dictionary[v] += 1
	}

	pl := make(PairList, 0, len(dictionary))

	keys := make([]string, 0, len(dictionary))
	for k, _ := range dictionary {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		pl = append(pl, Pair{k, dictionary[k]})
	}

	sort.Sort(sort.Reverse(pl))

	return pl[:10]
}
