package ui

import "github.com/jroimartin/gocui"

type CUI interface {
	Init(g *gocui.Gui) error
	Layout(g *gocui.Gui) error
	Keybindings(g *gocui.Gui) error // call this in Layout
	Release()
}
