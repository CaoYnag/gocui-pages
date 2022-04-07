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
	change_ui(GetLogin())
	return nil
}

func Run() error {
	if err := _g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}
