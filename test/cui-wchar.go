package main

import (
	"fmt"
	"log"

	"github.com/CaoYnag/gocui"
)

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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "中文测试English"
		n1, _ := fmt.Fprintln(v, "ascii asdnioqu172841958109(*^&%%$$%^)") // 37
		n2, _ := fmt.Fprintln(v, "这是一段中文！！")                              // 8
		n3, _ := fmt.Fprintln(v, "やばい!")                                  // 4
		n4, _ := fmt.Fprintln(v, "▲№§→←↑◎◇↓↓℃℃￣＼＆☆")                      // 16
		n5, _ := fmt.Fprintln(v, "end")                                   // 4
		fmt.Fprintf(v, "%d %d %d %d %d\n", n1, n2, n3, n4, n5)
	}
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
