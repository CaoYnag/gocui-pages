package ui

import (
	"github.com/jroimartin/gocui"
)

var (
	_g   *gocui.Gui
	_cur CUI
)

func Init() error {
	var err error
	_g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	/*err = ui.Keybindings(_g)
	if err != nil {
		return
	}*/
	change_ui(GetLogin())
	return nil
}

func Run() error {
	defer _g.Close()
	if err := _g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}
