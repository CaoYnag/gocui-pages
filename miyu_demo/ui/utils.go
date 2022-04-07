package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func change_ui(ui CUI) (err error) {
	if _cur != nil {
		_cur.Release()
	}
	err = ui.Init(_g)
	if err != nil {
		return
	}
	err = ui.Keybindings(_g)
	if err != nil {
		return
	}
	_g.SetManagerFunc(ui.Layout)
	_cur = ui
	return nil
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
func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
