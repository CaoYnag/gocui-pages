package ui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

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
func _new_btn(g *gocui.Gui, label string, x, y, w, h int, cb func(*gocui.Gui, *gocui.View) error) (v *gocui.View, err error) {
	vn := "btn_" + label
	v, err = g.SetView(vn, x, y, x+w, y+h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
		v.Editable = false
		v.Wrap = false
		v.Frame = true
		// TODO align center
		ns := (w - 1 - len(label)) / 2
		if ns < 0 {
			ns = 0
		}
		fmt.Fprintf(v, "%s%s", strings.Repeat(" ", ns), label)
		err = set_key_binding(g, vn, gocui.KeyEnter, gocui.ModNone, cb)
		if err != nil {
			return
		}
	}
	return
}
func set_cur_top_view(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func set_key_binding(g *gocui.Gui, v string, k interface{}, m gocui.Modifier, h func(*gocui.Gui, *gocui.View) error) error {
	// delete exists keybinding first
	g.DeleteKeybinding(v, k, m)
	return g.SetKeybinding(v, k, m, h)
}

func get_value(v *gocui.View) string {
	val := strings.Fields(v.Buffer())
	if len(val) > 0 {
		return val[0]
	}
	return ""
}
