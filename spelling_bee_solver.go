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
	letters []rune = []rune{'m', 'o', 'n', 't', 'y', 'r', 'p'}
	center  rune   = 'm'
)

func subs(size int) (vals []string) {
	if size == 0 {
		return
	}
	var str strings.Builder
	for _, letter := range letters {
		suffixes := subs(size - 1)
		if len(suffixes) == 0 {
			str.Reset()
			str.WriteRune(letter)
			vals = append(vals, str.String())
		}
		for _, suffix := range suffixes {
			str.Reset()
			str.WriteRune(letter)
			str.WriteString(suffix)
			vals = append(vals, str.String())
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
