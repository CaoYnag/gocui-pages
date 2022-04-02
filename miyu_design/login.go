package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

const (
	ENTRY_WID  = 20
	LABEL_WID  = 6
	ALL_HEIGHT = 2
)

var (
	_v          *gocui.View
	_vname      *gocui.View
	_vpsw       *gocui.View
	_info       *gocui.View
	_act        int // active view
	_name, _psw string
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("password", gocui.KeyEnter, gocui.ModNone, doLogin); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	g.Close()
	fmt.Printf("name: %s, psw: %s\n", _name, _psw)
}

func _new_label(g *gocui.Gui, label string, x, y, w, h int) error {
	_v, err := g.SetView("label_"+label, x, y, x+w, y+h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_v.Editable = false
		_v.Wrap = false
		_v.Frame = false
		fmt.Fprintf(_v, label)
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if maxX > 80 {
		maxX = 80
	}
	if maxY > 36 {
		maxY = 36
	}

	var err error
	_v, err = g.SetView("login", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_v.Title = "Login"
		_v.Editable = false
		_v.Wrap = true
	}
	X := (maxX - ENTRY_WID - LABEL_WID) / 2
	Y := maxY/2 - 3
	if err = _new_label(g, "name", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
		return err
	}
	if err = _new_label(g, "psw", X, Y+ALL_HEIGHT, LABEL_WID, ALL_HEIGHT); err != nil {
		return err
	}
	_vpsw, err = g.SetView("password", X+LABEL_WID, Y+ALL_HEIGHT, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT*2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_vpsw.Editable = true
		_vpsw.Wrap = false
		_vpsw.Mask = '#'
	}
	_info, err = g.SetView("info", X, Y+ALL_HEIGHT*2, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT*3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_info.Editable = true
		_info.Wrap = true
		_info.Frame = false
	}
	_vname, err = g.SetView("username", X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_vname.Editable = true
		_vname.Wrap = false
		if _, err := g.SetCurrentView("username"); err != nil {
			return err
		}
	}
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	if _act == 0 {
		// name
		if _, err := setCurrentViewOnTop(g, "password"); err != nil {
			return err
		}
		_act = 1
	} else {
		// psw
		if _, err := setCurrentViewOnTop(g, "username"); err != nil {
			return err
		}
		_act = 0
	}
	return nil
}

func doLogin(g *gocui.Gui, v *gocui.View) error {
	_name = _vname.Buffer()
	_psw = _vpsw.Buffer()
	fmt.Fprintf(_info, "\033[32;1mloging...\033[0m \033[31;1mfailed\033[0m")
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
