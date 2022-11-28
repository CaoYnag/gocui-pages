/*
MainUI
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CaoYnag/gocui"
)

var (
	_ctx  *gocui.View // msg area
	_inp  *gocui.View // input area
	_stat *gocui.View // user info
	_info *gocui.View // info bar
)

const (
	STAT_HGT = 2
	INP_HGT  = 4
	INFO_HGT = 1

	STAT = "stat"
	INP  = "inp"
	INFO = "info"
	CTX  = "ctx"
)

func layout(g *gocui.Gui) error {
	var err error
	maxX, maxY := g.Size()
	_stat, err = g.SetView(STAT, -1, -1, maxX, STAT_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(_stat, "Ehh1") // user stat here
	}

	_info, err = g.SetView(INFO, -1, maxY-INFO_HGT, maxX, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_info.Editable = false
		_info.Autoscroll = true
		fmt.Fprintf(_info, "nothing...")
	}
	_inp, err = g.SetView(INP, -1, maxY-INFO_HGT-INP_HGT, maxX, maxY-INFO_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_inp.Editable = true
		_inp.Autoscroll = true
		g.SetCurrentView(INP)
	}
	_ctx, err = g.SetView(CTX, -1, STAT_HGT, maxX, maxY-INFO_HGT-INP_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_ctx.Editable = false
		_ctx.Autoscroll = true
		show_msg("ehh1", "在 嘛 ?", time.Now())
		show_msg("spes", "不 在 !", time.Now())
		show_msg("spes", "里 四 居 !", time.Now())
	}
	return nil
}

func show_msg(name, msg string, ts time.Time) {
	if name == "spes" {
		// self
		fmt.Fprintf(_ctx, "\033[32;1m%s %s\033[0m\n", name, ts.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Fprintf(_ctx, "\033[34;1m%s %s\033[0m\n", name, ts.Format("2006-01-02 15:04:05"))
	}
	fmt.Fprintf(_ctx, "%s\n\n", msg)
}

func keybindings(g *gocui.Gui) error {
	// global
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	// input
	if err := g.SetKeybinding(INP, gocui.KeyEnter, gocui.ModNone, snd_msg); err != nil {
		return err
	}
	if err := g.SetKeybinding(INP, gocui.KeyCtrlSpace, gocui.ModNone, new_line); err != nil {
		return err
	}

	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	g.Cursor = true
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func snd_msg(g *gocui.Gui, v *gocui.View) error {
	show_msg("spes", _inp.Buffer(), time.Now())
	_inp.Clear()
	x, y := _inp.Origin()
	_inp.SetCursor(x, y)
	return nil
}

func new_line(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintf(_inp, "\n")
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
