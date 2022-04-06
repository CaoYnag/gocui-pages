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
	LABEL_WID  = 10
	ALL_HEIGHT = 2

	HELP_WID = 40

	REG  = "reg"
	NAME = "name"
	NICK = "nick"
	PSW  = "psw"
	RPSW = "rpsw"
	SEX  = "sex"
	DEV  = "dev"
	INFO = "info"
	HELP = "help"
)

var (
	_v     *gocui.View // form view
	_vname *gocui.View
	_vnick *gocui.View
	_vpsw  *gocui.View
	_vrpsw *gocui.View     // repeat psw
	_vsex  *gocui.View     // sex btn
	_vdev  *gocui.View     // dev id
	_act   int         = 0 // active view
	_views             = make([]*gocui.View, 6)

	_user = struct {
		name, nick, psw, sex string
	}{
		sex: "female",
	}
	_dev struct {
		id, pub string
	}
	_help  *gocui.View // help
	_ghelp *gocui.View // global help
	_info  *gocui.View // info view
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

	var err error
	{
		var help *gocui.View
		help, err = g.SetView("help_view", maxX-HELP_WID, 0, maxX-1, maxY-ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			help.Title = "Help"
			help.Editable = false
		}
		GLOBAL_HELP_HGT := 10
		_ghelp, err = g.SetView("global_help", maxX-HELP_WID, 0, maxX-1, GLOBAL_HELP_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_ghelp.Editable = false
			fmt.Fprintln(_ghelp, "TAB:    switch input field")
			fmt.Fprintln(_ghelp, "Ctrl+s: submit")
		}
		_help, err = g.SetView(HELP, maxX-HELP_WID, GLOBAL_HELP_HGT, maxX-1, maxY-ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_help.Editable = false
			show_help()
		}

	}
	_v, err = g.SetView(REG, 0, 0, maxX-HELP_WID-1, maxY-ALL_HEIGHT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_v.Title = "Register"
		_v.Editable = false
		_v.Wrap = true
	}
	_info, err = g.SetView(INFO, 0, maxY-ALL_HEIGHT, maxX, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_info.Editable = true
		_info.Wrap = true
		_info.Frame = false
		show_info("success...")
	}
	X := (maxX-HELP_WID-ENTRY_WID-LABEL_WID)/2 - 1
	Y := 1
	{
		if err = _new_label(g, "name", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
			return err
		}
		_vname, err = g.SetView(NAME, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_vname.Editable = true
			_vname.Wrap = false
			if _, err := g.SetCurrentView(NAME); err != nil {
				return err
			}
			_views[0] = _vname
		}
		Y += ALL_HEIGHT
	}
	{
		if err = _new_label(g, "nickname", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
			return err
		}
		_vnick, err = g.SetView(NICK, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_vnick.Editable = true
			_vnick.Wrap = false
			_views[1] = _vnick
		}
		Y += ALL_HEIGHT
	}
	{
		if err = _new_label(g, "password", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
			return err
		}
		_vpsw, err = g.SetView(PSW, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_vpsw.Editable = true
			_vpsw.Wrap = false
			_vpsw.Mask = '*'
			_views[2] = _vpsw
		}
		Y += ALL_HEIGHT
	}
	{
		if err = _new_label(g, "repeat", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
			return err
		}
		_vrpsw, err = g.SetView(RPSW, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_vrpsw.Editable = true
			_vrpsw.Wrap = false
			_vrpsw.Mask = '*'
			_views[3] = _vrpsw
		}
		Y += ALL_HEIGHT
	}
	{
		if err = _new_label(g, "sex", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
			return err
		}
		_vsex, err = g.SetView(SEX, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_vsex.Editable = false
			_vsex.Wrap = false
			_vsex.FgColor = gocui.ColorCyan
			change_sex(nil, nil)
			_views[4] = _vsex
		}
		Y += ALL_HEIGHT
	}
	{
		if err = _new_label(g, "device", X, Y, LABEL_WID, ALL_HEIGHT); err != nil {
			return err
		}
		_vdev, err = g.SetView(DEV, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+ALL_HEIGHT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			_vdev.Editable = true
			_vdev.Wrap = false
			_views[5] = _vdev
		}
		Y += ALL_HEIGHT
	}

	return nil
}

func change_view(g *gocui.Gui, v *gocui.View) error {
	_act++
	if _act == len(_views) {
		_act = 0
	}
	g.SetCurrentView(_views[_act].Name())
	show_help()
	return nil
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func change_sex(g *gocui.Gui, v *gocui.View) error {
	_change_sex()
	return nil
}
func submit(g *gocui.Gui, v *gocui.View) error {
	show_info("register success...")
	return nil

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
	err = g.SetKeybinding("", gocui.KeyCtrlS, gocui.ModNone, submit)
	if err != nil {
		log.Panicln(err)
	}
	err = g.SetKeybinding(SEX, gocui.KeyEnter, gocui.ModNone, change_sex)
	if err != nil {
		log.Panicln(err)
	}
	return nil
}

func show_info(msg string) {
	_info.Clear()
	fmt.Fprintf(_info, msg)
}

func _change_sex() {
	if _user.sex == "female" {
		_user.sex = "male"
	} else {
		_user.sex = "female"
	}
	_vsex.Clear()
	fmt.Fprint(_vsex, _user.sex)
}

func show_help() {
	switch _act {
	case 0:
		_show_help("input user name, also your accout\nlength between 8 and 16\ndo not contain special character")
	case 1:
		_show_help("nickname\nusually shows to other")
	case 2:
		_show_help("input password\nlength between 8 and 16\nmust contains number and character at least")
	case 3:
		_show_help("repeat your password")
	case 4:
		_show_help("press ENTER to change your sex")
	case 5:
		_show_help("use a name to identify this device")
	}
}
func _show_help(msg string) {
	_help.Clear()
	fmt.Fprint(_help, msg)
}
