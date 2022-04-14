package main

import (
	"log"
	"os"
)

var (
	SOURCE []byte
)

func init() {
	var e error
	SOURCE, e = os.ReadFile("test.md")
	if e != nil {
		log.Panic(e)
	}
}
