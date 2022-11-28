package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"testing"
)

var (
	d1k   []byte
	d1m   []byte
	d10m  []byte
	d100m []byte
	d1g   []byte
)

func setupsha() error {
	d1k = make([]byte, 1<<10)
	d1m = make([]byte, 1<<20)
	d10m = make([]byte, 10<<20)
	d100m = make([]byte, 100<<20)
	d1g = make([]byte, 1<<30)
	return nil
}

func sha(d []byte) {
	h := sha256.New()
	h.Write(d)
	h.Sum(nil)
}

func BenchmarkSha1K(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sha(d1k)
	}
}
func BenchmarkSha1M(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sha(d1m)
	}
}
func BenchmarkSha10M(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sha(d10m)
	}
}
func BenchmarkSha100M(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sha(d100m)
	}
}
func BenchmarkSha1G(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sha(d1g)
	}
}
func TestMain(m *testing.M) {
	if e := setupsha(); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
