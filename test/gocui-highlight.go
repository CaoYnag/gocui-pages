/*
Hello Twice!
*/
package main

import (
	"fmt"
	"log"

	"github.com/d4l3k/go-highlight"
	"github.com/jroimartin/gocui"
)

var CODE = `#include <iostream>
using namespace std;

template<typename T>
class Sample{};

typedef voidp void*;

int main(int argc, char** argv)
{
	cout << "hello, world!" << endl;
	return 0;
}
`

func main() {
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

func highlight_code(lang, code string) string {
	s, _ := highlight.Term(lang, []byte(code))
	return string(s)
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, highlight_code("c++", CODE))
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
