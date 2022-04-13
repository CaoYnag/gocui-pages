/*
Markdown in cui
*/
package main

import (
	"fmt"
	"log"

	"github.com/CaoYnag/gocui"
	"github.com/russross/blackfriday/v2"
)

func main() {
	//style_test()
	md_render()
}

func style_test() {
	t := New()
	t.write("normal ")
	t.bold(true).write("bold").bold(false).write(" ")
	t.italic(true).write("italic").italic(false).write(" ")
	t.dim(true).write("dim").dim(false).write(" ")
	t.strikethrough(true).write("strikethrough").strikethrough(false).write(" ")
	t.underline(true).write("underline").underline(false).write(" ")
	t.blink(true).write("blink").blink(false).write(" ")
	t.reverse(true).write("reverse").reverse(false).write(" ")
	t.hidden(true).write("hidden").hidden(false).write(" ")
	fmt.Println(string(t.flush()))
}

func md_render() {
	opt := blackfriday.WithRenderer(New())
	buf2 := blackfriday.Run([]byte(SOURCE), opt)
	fmt.Println(string(buf2))
}

func ui() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
