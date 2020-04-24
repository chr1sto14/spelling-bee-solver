package main

import (
	"fmt"
	"strings"

	"github.com/chr1sto14/enchant"
)

const (
	MinWordLength int = 4
	MaxWordLength int = 10
)

var (
	letters     []rune           = []rune{'m', 'o', 'n', 't', 'y', 'r', 'p'}
	center      rune             = 'm'
	suffixcache map[int][]string = make(map[int][]string)
)

func subs(size int) (vals []string) {
	if size == 0 {
		return
	}
	var str strings.Builder
	for _, letter := range letters {
		subsize := size - 1
		suffixes, ok := suffixcache[subsize]
		if !ok {
			suffixes = subs(subsize)
			suffixcache[subsize] = suffixes
		}
		nSuffix := len(suffixes)
		if nSuffix == 0 {
			str.Reset()
			str.WriteRune(letter)
			vals = append(vals, str.String())
			continue
		}
		nVal := len(vals)
		vals = append(vals, make([]string, nSuffix)...)
		for i, suffix := range suffixes {
			str.Reset()
			str.WriteRune(letter)
			str.WriteString(suffix)
			vals[nVal+i] = str.String()
		}
	}
	return
}

func main() {
	// create a new enchant instance
	dict, err := enchant.NewEnchant()
	if err != nil {
		panic("Enchant error: " + err.Error())
	}
	// defer freeing memory to the end of this program
	defer dict.Free()
	// check whether a certain dictionary exists on the system
	if !dict.DictExists("en_US") {
		panic("need en_US")
	}
	dict.LoadDict("en_US")

	var answers []string
	for size := MinWordLength; size < MaxWordLength; size++ {
		for _, word := range subs(size) {
			if !strings.ContainsRune(word, center) {
				continue
			}
			if dict.Check(word) {
				answers = append(answers, word)
			}
		}
	}
	fmt.Println(answers)
}
