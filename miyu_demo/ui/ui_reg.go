package ui

import (
	"github.com/jroimartin/gocui"
)

type _reg_ui struct {
	_v     *gocui.View
	_vname *gocui.View
	_vpsw  *gocui.View
	_info  *gocui.View
	_act   int // active view

	_data struct {
		name, psw string
	}
}

func GetReg() *_reg_ui {
	return &_reg_ui{}
}

func (s *_reg_ui) Init(g *gocui.Gui) error {
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	return nil
}

func (s *_reg_ui) Layout(g *gocui.Gui) error {
	if e := s.Keybindings(g); e != nil {
		return e
	}
	maxX, maxY := g.Size()
	if maxX > 80 {
		maxX = 80
	}
	if maxY > 36 {
		maxY = 36
	}

	var err error
	s._v, err = g.SetView("reg", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._v.Title = "Register"
		s._v.Editable = false
		s._v.Wrap = true
	}
	X := (maxX - ENTRY_WID - LABEL_WID) / 2
	Y := maxY/2 - 3
	if err = _new_label(g, "name", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
		return err
	}
	if err = _new_label(g, "psw", X, Y+WIDGET_HGT, LABEL_WID, WIDGET_HGT); err != nil {
		return err
	}
	s._vpsw, err = g.SetView("password", X+LABEL_WID, Y+WIDGET_HGT, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT*2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._vpsw.Editable = true
		s._vpsw.Wrap = false
		s._vpsw.Mask = '#'
	}
	s._info, err = g.SetView("info", X, Y+WIDGET_HGT*2, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT*3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._info.Editable = true
		s._info.Wrap = true
		s._info.Frame = false
	}
	s._vname, err = g.SetView("username", X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._vname.Editable = true
		s._vname.Wrap = false
		if _, err := g.SetCurrentView("username"); err != nil {
			return err
		}
	}
	return nil
}

func (s *_reg_ui) Keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, s.next_view); err != nil {
		return err
	}
	if err := g.SetKeybinding("password", gocui.KeyEnter, gocui.ModNone, s.do_reg); err != nil {
		return err
	}
	return nil
}

func (s *_reg_ui) Release() {}

func (s *_reg_ui) next_view(g *gocui.Gui, v *gocui.View) error {
	if s._act == 0 {
		// name
		if _, err := set_cur_top_view(g, "password"); err != nil {
			return err
		}
		s._act = 1
	} else {
		// psw
		if _, err := set_cur_top_view(g, "username"); err != nil {
			return err
		}
		s._act = 0
	}
	return nil
}
func (s *_reg_ui) do_reg(g *gocui.Gui, v *gocui.View) error {
	s._data.name = s._vname.Buffer()
	s._data.psw = s._vpsw.Buffer()
	// TODO handle success and failed
	return nil
}
