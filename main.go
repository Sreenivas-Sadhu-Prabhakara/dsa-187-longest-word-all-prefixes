package main

import "fmt"

type trie struct {
	children [26]*trie
	isWord   bool
}

func (t *trie) insert(w string) {
	n := t
	for i := 0; i < len(w); i++ {
		c := w[i] - 'a'
		if n.children[c] == nil {
			n.children[c] = &trie{}
		}
		n = n.children[c]
	}
	n.isWord = true
}

// hasAllPrefixes reports whether every prefix of w (including w) is a word.
func (t *trie) hasAllPrefixes(w string) bool {
	n := t
	for i := 0; i < len(w); i++ {
		n = n.children[w[i]-'a']
		if n == nil || !n.isWord {
			return false
		}
	}
	return true
}

// LongestCompleteString returns the longest word whose every prefix is also
// in the list (ties broken lexicographically smaller); "" when none exists.
func LongestCompleteString(words []string) string {
	root := &trie{}
	for _, w := range words {
		root.insert(w)
	}
	best := ""
	for _, w := range words {
		if !root.hasAllPrefixes(w) {
			continue
		}
		if len(w) > len(best) || (len(w) == len(best) && w < best) {
			best = w
		}
	}
	return best
}

func main() {
	fmt.Println(LongestCompleteString([]string{"n", "ni", "nin", "ninj", "ninja", "ninga"}))
}
