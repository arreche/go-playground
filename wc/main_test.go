package main

import (
	"testing"
)

func Test_CountWords(t *testing.T) {
	want := 3
	got := countWords(" Hello  my\nfrield")
	if got != want {
		t.Errorf("Got: %d, Want: %d", got, want)
	}
}

func Benchmark_CountWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		countWords("Hello world")
	}
}
