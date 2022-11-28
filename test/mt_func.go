package main

import (
	"fmt"
	"time"
	"unsafe"
)

type S struct{}

func (s *S) foo() {
	go s.foo2()
	fmt.Printf("s in foo: %v\n", unsafe.Pointer(s))
}
func (s *S) foo2() {
	fmt.Printf("s in foo2: %v\n", unsafe.Pointer(s))
}
func main() {
	s := S{}
	s.foo()
	<-time.After(time.Millisecond)
}
