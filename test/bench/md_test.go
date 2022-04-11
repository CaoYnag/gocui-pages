package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/russross/blackfriday/v2"
	"github.com/yuin/goldmark"
)

var (
	SOURCE []byte
)

func setup() error {
	var e error
	SOURCE, e = os.ReadFile("/home/spes/Documents/notes/backend/Spring&Boot.md")
	return e
}

func gold_mark() {
	var buf bytes.Buffer
	goldmark.Convert(SOURCE, &buf)
}

func gomarkdown() {
	markdown.ToHTML(SOURCE, nil, nil)
}

func blackfridaymarkdown() {
	blackfriday.Run(SOURCE)
}

func BenchmarkGoldMark(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gold_mark()
	}
}

func BenchmarkGomarkdown(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gomarkdown()
	}
}

func BenchmarkBlackfriday(b *testing.B) {
	for n := 0; n < b.N; n++ {
		blackfridaymarkdown()
	}
}

func TestMain(m *testing.M) {
	if e := setup(); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
