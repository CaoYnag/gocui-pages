package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type _reg_ui struct {
	_g     *gocui.Gui
	_v     *gocui.View // form view
	_help  *gocui.View // help
	_ghelp *gocui.View // global help
	_info  *gocui.View // info view
	_act   int         // active view
	_views []*gocui.View

	_vname     *gocui.View
	_vnick     *gocui.View
	_vpsw      *gocui.View
	_vrpsw     *gocui.View // repeat psw
	_vsex      *gocui.View // sex btn
	_vdev      *gocui.View // dev id
	_btn_reg   *gocui.View
	_btn_login *gocui.View

	_user struct {
		name, nick, psw, sex string
	}
	_dev struct {
		id, pub string
	}
}

const (
	MAX_REG_HELP_WID = 40
)

func GetReg(name string) *_reg_ui {
	rslt := &_reg_ui{}
	rslt._user.name = name
	rslt._user.sex = "female"
	return rslt
}

func (s *_reg_ui) Init() error {
	var err error
	s._g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	s._g.Highlight = true
	s._g.SelFgColor = gocui.ColorRed
	s._views = make([]*gocui.View, 8)
	s._act = 0

	s._g.SetManagerFunc(s.layout)
	return nil
}

func (s *_reg_ui) Run() error {
	if err := s._g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (s *_reg_ui) Release() {
	s._g.Close()
}

func (s *_reg_ui) layout(g *gocui.Gui) error {
	if e := s.keybindings(s._g); e != nil {
		return e
	}
	maxX, maxY := g.Size()
	if maxX > 100 {
		maxX = 100
	}
	REG_HELP_WID := maxX / 3

	var err error
	{
		var help *gocui.View
		help, err = g.SetView("help_view", maxX-REG_HELP_WID, 0, maxX-1, maxY-WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			help.Title = "Help"
			help.Editable = false
		}
		GLOBAL_HELP_HGT := maxY / 3
		s._ghelp, err = g.SetView("global_help", maxX-REG_HELP_WID, 0, maxX-1, GLOBAL_HELP_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._ghelp.Editable = false
			s._ghelp.Wrap = true
			fmt.Fprintln(s._ghelp, "TAB:    next field")
			fmt.Fprintln(s._ghelp, "Ctrl+P: prev field")
		}
		s._help, err = g.SetView(HELP, maxX-REG_HELP_WID, GLOBAL_HELP_HGT, maxX-1, maxY-WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._help.Editable = false
			s._help.Wrap = true
			s.show_help()
		}

	}
	s._v, err = g.SetView(REG, 0, 0, maxX-REG_HELP_WID-1, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._v.Title = "Register"
		s._v.Editable = false
		s._v.Wrap = true
	}
	s._info, err = g.SetView(INFO, 0, maxY-WIDGET_HGT, maxX, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._info.Editable = true
		s._info.Wrap = true
		s._info.Frame = false
		s.show_info("success...")
	}
	X := (maxX-REG_HELP_WID-ENTRY_WID-LABEL_WID)/2 - 1
	Y := 1
	{
		if err = _new_label(g, "name", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
			return err
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
			fmt.Fprint(s._vname, s._user.name)
		}
		Y += WIDGET_HGT
	}
	{
		if err = _new_label(g, "nickname", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
			return err
		}
		s._vnick, err = g.SetView(NICK, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._vnick.Editable = true
			s._vnick.Wrap = false
			s._views[1] = s._vnick
		}
		Y += WIDGET_HGT
	}
	{
		if err = _new_label(g, "password", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
			return err
		}
		s._vpsw, err = g.SetView(PSW, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._vpsw.Editable = true
			s._vpsw.Wrap = false
			s._vpsw.Mask = '*'
			s._views[2] = s._vpsw
		}
		Y += WIDGET_HGT
	}
	{
		if err = _new_label(g, "repeat", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
			return err
		}
		s._vrpsw, err = g.SetView(RPSW, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._vrpsw.Editable = true
			s._vrpsw.Wrap = false
			s._vrpsw.Mask = '*'
			s._views[3] = s._vrpsw
		}
		Y += WIDGET_HGT
	}
	{
		if err = _new_label(g, "sex", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
			return err
		}
		s._vsex, err = g.SetView(SEX, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._vsex.Editable = false
			s._vsex.Wrap = false
			s._vsex.FgColor = gocui.ColorCyan
			s.change_sex(nil, nil)
			s._views[4] = s._vsex
		}
		Y += WIDGET_HGT
	}
	{
		if err = _new_label(g, "device", X, Y, LABEL_WID, WIDGET_HGT); err != nil {
			return err
		}
		s._vdev, err = g.SetView(DEV, X+LABEL_WID, Y, X+LABEL_WID+ENTRY_WID, Y+WIDGET_HGT)
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
			s._vdev.Editable = true
			s._vdev.Wrap = false
			s._views[5] = s._vdev
		}
		Y += WIDGET_HGT
	}
	// login register button
	BTN_WID := 10
	X = (maxX - REG_HELP_WID - BTN_WID*2 - 2) / 4
	Y++
	if s._btn_login, err = _new_btn(g, "login", X, Y, BTN_WID, WIDGET_HGT, s.to_login); err != nil {
		return nil
	}
	s._views[6] = s._btn_login
	X = X*3 + BTN_WID
	if s._btn_reg, err = _new_btn(g, "register", X, Y, BTN_WID, WIDGET_HGT, s.submit); err != nil {
		return nil
	}
	s._views[7] = s._btn_reg
	return nil
}

func (s *_reg_ui) keybindings(g *gocui.Gui) error {
	var err error
	err = set_key_binding(g, "", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
		return err
	}
	err = set_key_binding(g, "", gocui.KeyTab, gocui.ModNone, s.next_view)
	if err != nil {
		return err
	}
	err = set_key_binding(g, "", gocui.KeyCtrlP, gocui.ModNone, s.prev_view)
	if err != nil {
		return err
	}
	err = set_key_binding(g, "", gocui.KeyCtrlS, gocui.ModNone, s.submit)
	if err != nil {
		return err
	}
	err = set_key_binding(g, SEX, gocui.KeyEnter, gocui.ModNone, s.change_sex)
	if err != nil {
		return err
	}
	return nil
}

func (s *_reg_ui) next_view(g *gocui.Gui, v *gocui.View) error {
	s._act++
	if s._act == len(s._views) {
		s._act = 0
	}
	g.SetCurrentView(s._views[s._act].Name())
	s.show_help()
	return nil
}
func (s *_reg_ui) prev_view(g *gocui.Gui, v *gocui.View) error {
	s._act--
	if s._act < 0 {
		s._act = len(s._views) - 1
	}
	g.SetCurrentView(s._views[s._act].Name())
	s.show_help()
	return nil
}

func (s *_reg_ui) submit(g *gocui.Gui, v *gocui.View) error {
	s.show_info("register success...")
	return nil

}

func (s *_reg_ui) change_sex(g *gocui.Gui, v *gocui.View) error {
	s._change_sex()
	return nil
}

func (s *_reg_ui) show_info(msg string) {
	s._info.Clear()
	fmt.Fprintf(s._info, msg)
}

func (s *_reg_ui) _change_sex() {
	if s._user.sex == "female" {
		s._user.sex = "male"
	} else {
		s._user.sex = "female"
	}
	s._vsex.Clear()
	fmt.Fprint(s._vsex, s._user.sex)
}

func (s *_reg_ui) show_help() {
	switch s._act {
	case 0:
		s._show_help("input user name, also your accout\nlength between 8 and 16\ndo not contain special character")
	case 1:
		s._show_help("nickname\nusually shows to other")
	case 2:
		s._show_help("input password\nlength between 8 and 16\nmust contains number and character at least")
	case 3:
		s._show_help("repeat your password")
	case 4:
		s._show_help("press ENTER to change your sex")
	case 5:
		s._show_help("use a name to identify this device")
	case 6:
		s._show_help("go to login page")
	case 7:
		s._show_help("submit registration")
	}
}
func (s *_reg_ui) _show_help(msg string) {
	s._help.Clear()
	fmt.Fprint(s._help, msg)
}

func (s *_reg_ui) to_login(g *gocui.Gui, v *gocui.View) error {
	_jump_to(GetLogin(get_value(s._vname), get_value(s._vpsw)))
	return gocui.ErrQuit
}
