package main

import "testing"

func TestLongestCompleteString(t *testing.T) {
	cases := []struct {
		words []string
		want  string
	}{
		{[]string{"n", "ni", "nin", "ninj", "ninja", "ninga"}, "ninja"},
		{[]string{"ab", "bc"}, ""},             // single letters missing
		{[]string{"a", "b", "ab", "ba"}, "ab"}, // tie ab/ba -> lexicographic
		{[]string{"z"}, "z"},
		{nil, ""},
		{[]string{"cat", "ca", "c", "dog", "do", "d", "dogs"}, "dogs"},
	}
	for _, c := range cases {
		if got := LongestCompleteString(c.words); got != c.want {
			t.Errorf("LongestCompleteString(%v) = %q, want %q", c.words, got, c.want)
		}
	}
}
