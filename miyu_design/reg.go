/*
Register
*/
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

	HELP_WID = 40
)

var (
	_v     *gocui.View // form view
	_vname *gocui.View
	_vnick *gocui.View
	_vpsw  *gocui.View
	_vsex  *gocui.View     // sex btn
	_vrpsw *gocui.View     // repeat psw
	_vdev  *gocui.View     // dev id
	_act   int         = 0 // active view
	_views             = []*gocui.View{
		_vname, _vnick, _vpsw, _vsex, _vrpsw, _vdev,
	}

	_user struct {
		name, nick, psw string
	}
	_dev struct {
		id, pub string
	}
	_help *gocui.View // help
	_info *gocui.View // info view
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	g.Close()
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
	_v, err = g.SetView("Register", 0, 0, maxX-HELP_WID-1, maxY-1)
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
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func change_view(g *gocui.Gui, v *gocui.View) error {
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

func register(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	var err error
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		log.Panicln(err)
	}
	err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, change_view)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}
