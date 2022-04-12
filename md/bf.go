package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/russross/blackfriday/v2"
)

type TermRender struct {
	params []int
	buf    bytes.Buffer
}

func New() *TermRender {
	return &TermRender{}
}

func (s *TermRender) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	switch node.Type {
	case blackfriday.Text:
		w.Write([]byte(string(node.Literal)))
	case blackfriday.Softbreak:
		w.Write([]byte("-----------------------------\n"))
	case blackfriday.Hardbreak:
		w.Write([]byte("=============================\n"))
	case blackfriday.Emph:
		if entering {
			w.Write([]byte("E<"))
		} else {
			w.Write([]byte(">E"))
		}
	case blackfriday.Strong:
		if entering {
			w.Write([]byte("S<"))
		} else {
			w.Write([]byte(">S"))
		}
	case blackfriday.Del:
		if entering {
			w.Write([]byte("D<"))
		} else {
			w.Write([]byte(">D"))
		}
	case blackfriday.Link:
		if entering {
			w.Write([]byte("LN<"))
		} else {
			w.Write([]byte("> -> " + string(node.LinkData.Destination) + "\n"))
		}
	case blackfriday.Image:
		if entering {
			w.Write([]byte("LN<"))
		} else {
			w.Write([]byte("> -> " + string(node.LinkData.Destination) + "\n"))
		}
	case blackfriday.Code:
		w.Write([]byte("#<" + string(node.Literal) + ">\n"))
	case blackfriday.Document:
		if entering {
			w.Write([]byte("=DOC START"))
		} else {
			w.Write([]byte("=DOC END"))
		}
	case blackfriday.Paragraph:
		if entering {
			w.Write([]byte("P<"))
		} else {
			w.Write([]byte(">P\n"))
		}
	case blackfriday.BlockQuote:
		if entering {
			w.Write([]byte("BQ<"))
		} else {
			w.Write([]byte(">BQ\n"))
		}
	case blackfriday.HTMLBlock:
		w.Write([]byte("HB<" + string(node.Literal) + ">HB\n"))
	case blackfriday.HTMLSpan:
		w.Write([]byte("HS<" + string(node.Literal) + ">HS\n"))
	case blackfriday.Heading:
		if entering {
			w.Write([]byte("H" + fmt.Sprintf("%d", node.Level) + " " + string(node.HeadingID)))
		} else {
			w.Write([]byte("\n"))
		}

	case blackfriday.HorizontalRule:
		w.Write([]byte("HR\n"))
	case blackfriday.List:
		ordered := node.ListFlags & blackfriday.ListTypeOrdered
		if entering {
			if ordered == 0 {
				w.Write([]byte("ul<\n"))
			} else {
				w.Write([]byte("ol<\n"))
			}
		} else {
			w.Write([]byte(">\n"))
		}
	case blackfriday.Item:
		if entering {
			w.Write([]byte("-"))
		} else {
			w.Write([]byte("-\n"))
		}
	case blackfriday.CodeBlock:
		w.Write([]byte(string(node.Info) + "<\n" + string(node.Literal) + ">\n"))
	case blackfriday.Table:
		if entering {
			w.Write([]byte("T<\n"))
		} else {
			w.Write([]byte(">\n"))
		}
	case blackfriday.TableCell:
		if entering {
			w.Write([]byte("TC<\n"))
		} else {
			w.Write([]byte(">\n"))
		}
	case blackfriday.TableHead:
		if entering {
			w.Write([]byte("TH<\n"))
		} else {
			w.Write([]byte(">\n"))
		}
	case blackfriday.TableBody:
		if entering {
			w.Write([]byte("TB<\n"))
		} else {
			w.Write([]byte(">\n"))
		}
	case blackfriday.TableRow:
		if entering {
			w.Write([]byte("TR<\n"))
		} else {
			w.Write([]byte(">\n"))
		}
	default:
	}
	return blackfriday.GoToNext
}
func (s *TermRender) RenderHeader(w io.Writer, ast *blackfriday.Node) {}
func (s *TermRender) RenderFooter(w io.Writer, ast *blackfriday.Node) {}

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
		s.buf.WriteRune(';')
	}
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
		s._add_params(TERM_RESET_UNDERLIN)
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
