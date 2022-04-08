package ui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

type _login_ui struct {
	_v         *gocui.View
	_vname     *gocui.View
	_vpsw      *gocui.View
	_btn_login *gocui.View
	_btn_reg   *gocui.View
	_info      *gocui.View
	_act       int // active view
	_views     []*gocui.View

	_data struct {
		name, psw string
	}
}

func GetLogin() *_login_ui {
	return &_login_ui{}
}

func (s *_login_ui) Init(g *gocui.Gui) error {
	g.Highlight = true
	g.SelFgColor = gocui.ColorRed
	s._views = make([]*gocui.View, 4)
	s._act = 0
	return nil
}

func (s *_login_ui) Layout(g *gocui.Gui) error {
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
	s._v, err = g.SetView("login", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._v.Title = "Login"
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
		s._views[1] = s._vpsw
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
		s._views[0] = s._vname
	}
	// login register button
	BTN_WID := 10
	X = (maxX - BTN_WID*2) / 4
	Y = Y + WIDGET_HGT*3
	if s._btn_reg, err = _new_btn(g, "register", X, Y, BTN_WID, WIDGET_HGT, s.to_reg); err != nil {
		return nil
	}
	s._views[2] = s._btn_reg
	X = X*3 + BTN_WID
	if s._btn_login, err = _new_btn(g, "login", X, Y, BTN_WID, WIDGET_HGT, s.do_login); err != nil {
		return nil
	}
	s._views[3] = s._btn_login

	return nil
}

func (s *_login_ui) Keybindings(g *gocui.Gui) error {
	if err := set_key_binding(g, "", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := set_key_binding(g, "", gocui.KeyTab, gocui.ModNone, s.next_widget); err != nil {
		return err
	}
	if err := set_key_binding(g, "password", gocui.KeyEnter, gocui.ModNone, s.do_login); err != nil {
		return err
	}
	return nil
}

func (s *_login_ui) Release() {}

func (s *_login_ui) next_widget(g *gocui.Gui, v *gocui.View) error {
	s._act++
	if s._act >= len(s._views) {
		s._act = 0
	}
	set_cur_top_view(g, s._views[s._act].Name())
	return nil
}
func (s *_login_ui) do_login(g *gocui.Gui, v *gocui.View) error {
	name := s._vname.Buffer()
	psw := s._vpsw.Buffer()
	// TODO corrupt if name or psw is empty
	s._data.name = strings.Fields(name)[0]
	s._data.psw = strings.Fields(psw)[0]
	// TODO do login, and handle result
	s.show_info(fmt.Sprintf("login: %s - %s", s._data.name, s._data.psw))
	return nil
}
func (s *_login_ui) to_reg(g *gocui.Gui, v *gocui.View) error {
	// TODO for convenience, pass username to reg ui here
	change_ui(GetReg())
	return nil
}
func (s *_login_ui) show_info(msg string) {
	s._info.Clear()
	fmt.Fprintf(s._info, "\033[31;1m%s\033[0m", msg)
}
