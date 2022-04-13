package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type TermRender struct {
	params []int
	buf    bytes.Buffer
}

func New() *TermRender {
	return &TermRender{}
}

func (s *TermRender) flush_to(w io.Writer) {
	s.reset()
	s.buf.WriteTo(w)
}

// output datas, and clear all stats
func (s *TermRender) flush() []byte {
	s.reset() // recover status
	d := s.buf.Bytes()
	s.buf.Reset()
	return d
}

func (s *TermRender) write(str string) *TermRender {
	if len(s.params) > 0 {
		s._write_params()
	}
	s.buf.WriteString(str)
	return s
}

func (s *TermRender) _write_params() {
	s.buf.WriteString("\x1b[")
	for i := range s.params {
		s.buf.WriteString(strconv.Itoa(s.params[i]))
		fmt.Printf("write param: %d\n", s.params[i])
		s.buf.WriteRune(';')
	}
	s.params = nil
	s.buf.Truncate(s.buf.Len() - 1)
	s.buf.WriteRune('m')
}

func (s *TermRender) _add_params(v ...int) *TermRender {
	s.params = append(s.params, v...)
	return s
}

func (s *TermRender) reset() *TermRender {
	// ignore params not use
	if len(s.params) > 0 {
		s.params = nil
	}
	s._add_params(TERM_RESET)
	return s
}
func (s *TermRender) fg(c int) *TermRender {
	s._add_params(c)
	return s
}
func (s *TermRender) bg(c int) *TermRender {
	s._add_params(c)
	return s
}
func (s *TermRender) fg256(c int) *TermRender {
	s._add_params(38, 5, c)
	return s
}
func (s *TermRender) bg256(c int) *TermRender {
	s._add_params(48, 5, c)
	return s
}
func (s *TermRender) italic(v bool) *TermRender {
	if v {
		s._add_params(TERM_ITALIC)
	} else {
		s._add_params(TERM_RESET_ITALIC)
	}
	return s
}
func (s *TermRender) blink(v bool) *TermRender {
	if v {
		s._add_params(TERM_BLINK)
	} else {
		s._add_params(TERM_RESET_BLINK)
	}
	return s
}
func (s *TermRender) reverse(v bool) *TermRender {
	if v {
		s._add_params(TERM_REVERSE)
	} else {
		s._add_params(TERM_RESET_REVERSE)
	}
	return s
}
func (s *TermRender) dim(v bool) *TermRender {
	if v {
		s._add_params(TERM_DIM)
	} else {
		s._add_params(TERM_RESET_DIM)
	}
	return s
}
func (s *TermRender) hidden(v bool) *TermRender {
	if v {
		s._add_params(TERM_HIDDEN)
	} else {
		s._add_params(TERM_RESET_HIDDEN)
	}
	return s
}
func (s *TermRender) bold(v bool) *TermRender {
	if v {
		s._add_params(TERM_BOLD)
	} else {
		s._add_params(TERM_RESET_BOLD)
	}
	return s
}
func (s *TermRender) underline(v bool) *TermRender {
	if v {
		s._add_params(TERM_UNDERLINE)
	} else {
		s._add_params(TERM_RESET_UNDERLINE)
	}
	return s
}
func (s *TermRender) strikethrough(v bool) *TermRender {
	if v {
		s._add_params(TERM_STRIKETHROUGH)
	} else {
		s._add_params(TERM_RESET_STRIKETHROUGH)
	}
	return s
}
