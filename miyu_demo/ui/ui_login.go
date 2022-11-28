package ui

import (
	"fmt"

	"github.com/CaoYnag/gocui"
)

type _login_ui struct {
	_g         *gocui.Gui
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

func GetLogin(name, psw string) *_login_ui {
	rslt := &_login_ui{}
	rslt._data.name = name
	rslt._data.psw = psw
	return rslt
}

func (s *_login_ui) Init() error {
	var err error
	s._g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	s._g.Highlight = true
	s._g.SelFgColor = gocui.ColorRed
	s._views = make([]*gocui.View, 4)
	s._act = 0

	s._g.SetManagerFunc(s.layout)
	if e := s.keybindings(s._g); e != nil {
		return e
	}
	return nil
}

func (s *_login_ui) Run() error {
	if err := s._g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (s *_login_ui) Release() {
	s._g.Close()
}

func (s *_login_ui) layout(g *gocui.Gui) error {
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
	s._vpsw, err = g.SetView(PSW, X+LABEL_WID, Y+WIDGET_HGT, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT*2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._vpsw.Editable = true
		s._vpsw.Wrap = false
		s._vpsw.Mask = '#'
		s._views[1] = s._vpsw
		fmt.Fprint(s._vpsw, s._data.psw)
	}
	s._info, err = g.SetView(INFO, X, Y+WIDGET_HGT*2, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT*3)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._info.Editable = true
		s._info.Wrap = true
		s._info.Frame = false
	}
	s._vname, err = g.SetView(NAME, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._vname.Editable = true
		s._vname.Wrap = false
		if _, err := g.SetCurrentView(NAME); err != nil {
			return err
		}
		s._views[0] = s._vname
		fmt.Fprint(s._vname, s._data.name)
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

func (s *_login_ui) keybindings(g *gocui.Gui) error {
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, gocui.KeyTab, gocui.ModNone, s.next_widget); err != nil {
		return err
	}
	if err := set_key_binding(g, PSW, gocui.KeyEnter, gocui.ModNone, s.do_login); err != nil {
		return err
	}
	return nil
}

func (s *_login_ui) next_widget(g *gocui.Gui, v *gocui.View) error {
	s._act++
	if s._act >= len(s._views) {
		s._act = 0
	}
	set_cur_top_view(g, s._views[s._act].Name())
	return nil
}
func (s *_login_ui) do_login(g *gocui.Gui, v *gocui.View) error {
	s._data.name = get_value(s._vname)
	s._data.psw = get_value(s._vpsw)
	// TODO do login, and handle result
	s.show_info(fmt.Sprintf("login: %s - %s", s._data.name, s._data.psw))
	if s._data.name == "spes" && s._data.psw == "123" {
		_jump_to(GetSync())
		return gocui.ErrQuit
	}
	return nil
}
func (s *_login_ui) to_reg(g *gocui.Gui, v *gocui.View) error {
	_jump_to(GetReg(get_value(s._vname)))
	return gocui.ErrQuit
}
func (s *_login_ui) show_info(msg string) {
	s._info.Clear()
	fmt.Fprintf(s._info, "\033[31;1m%s\033[0m", msg)
}
