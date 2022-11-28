/*
Hello Twice!
*/
package main

import (
	"fmt"
	"log"

	"github.com/CaoYnag/gocui"
)

var (
	g   *gocui.Gui
	err error
)

func do_init() {
	g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	dispatch(g)
}

func dispatch(g *gocui.Gui) {
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
}

func run() {
	defer g.Close()
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func main() {
	do_init()
	run()
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
	fmt.Printf("in quit func\n")
	return gocui.ErrQuit
}
