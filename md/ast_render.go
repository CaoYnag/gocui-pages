package main

import (
	"fmt"
	"io"

	"github.com/russross/blackfriday/v2"
)

type ASTRender struct{}

func NewASTRender() *ASTRender {
	return &ASTRender{}
}

func (s *ASTRender) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
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
			w.Write([]byte("IM<"))
		} else {
			w.Write([]byte("> -> " + string(node.LinkData.Destination) + "\n"))
		}
	case blackfriday.Code:
		w.Write([]byte("#<" + string(node.Literal) + ">\n"))
	case blackfriday.Document:
	case blackfriday.Paragraph:
	case blackfriday.BlockQuote:
	case blackfriday.HTMLBlock:
	case blackfriday.HTMLSpan:
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
	case blackfriday.CodeBlock:
		w.Write([]byte(string(node.Info) + "<\n" + string(node.Literal) + ">\n"))
	case blackfriday.Table:
		if entering {
			w.Write([]byte("T<\n"))
		} else {
			w.Write([]byte(">\n"))
		}
	case blackfriday.TableCell:
	case blackfriday.TableHead:
	case blackfriday.TableBody:
	case blackfriday.TableRow:
	default:
	}
	return blackfriday.GoToNext
}
func (s *ASTRender) RenderHeader(w io.Writer, ast *blackfriday.Node) {}
func (s *ASTRender) RenderFooter(w io.Writer, ast *blackfriday.Node) {}
