package main

import (
	"fmt"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

// isWordChar reports whether r should be kept as part of a word
// (letters and apostrophes, so "don't" stays one word).
func isWordChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '\''
}

// TopWords returns the top n most frequent words in text.
func TopWords(text string, n int) []wordCount {
	counts := make(map[string]int)

	// Split on anything that isn't a letter/apostrophe.
	fields := strings.FieldsFunc(text, func(r rune) bool {
		return !isWordChar(r)
	})

	for _, w := range fields {
		w = strings.ToLower(strings.Trim(w, "'")) // drop stray leading/trailing quotes
		if w == "" {
			continue
		}
		counts[w]++
	}

	list := make([]wordCount, 0, len(counts))
	for w, c := range counts {
		list = append(list, wordCount{w, c})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].count != list[j].count {
			return list[i].count > list[j].count
		}
		return list[i].word < list[j].word // tie-break alphabetically
	})

	if n > len(list) {
		n = len(list)
	}
	return list[:n]
}

func main() {
	paragraph := `The quick brown fox jumps over the lazy dog. The dog barks,
	but the fox is already gone. The lazy dog just watches the fox run
	quickly into the woods, and the quick fox disappears.`

	top := TopWords(paragraph, 5)

	fmt.Println("Top 5 words:")
	for i, wc := range top {
		fmt.Printf("%d. %-10s %d\n", i+1, wc.word, wc.count)
	}
}
