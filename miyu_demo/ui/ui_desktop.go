/*
Desktop

2 ways to switch panel
1. Alt=(1, 2, 3, 4)
2. use TAB to focus navbar, then use arrow and enter select panel
*/
package ui

import (
	"fmt"
	"gocui-demo/miyu_demo/mocks/msg"

	"github.com/jroimartin/gocui"
)

type _desktop_ui struct {
	_g        *gocui.Gui
	_nav      *gocui.View
	_ssn      *gocui.View
	_usrs     *gocui.View
	_grps     *gocui.View
	_settings *gocui.View
	_cur      *gocui.View
	_info     *gocui.View

	VIEWS []string
	LEN   int
	NUMS  map[string]int

	_users    map[string]*msg.UserInfo  // users
	_groups   map[string]*msg.GroupInfo // groups
	_sessions map[string]*msg.ChatMsg   // user or group name
}

const (
	MIN_SIDE_VIEW = 12
)

func GetDesktop() *_desktop_ui {
	rslt := &_desktop_ui{}
	rslt.VIEWS = []string{SSN, USRS, GRPS, SETTINGS}
	rslt.LEN = len(rslt.VIEWS)
	rslt.NUMS = make(map[string]int)
	return rslt
}

func (s *_desktop_ui) Init() error {
	var err error
	s._g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}

	s._g.SetManagerFunc(s.layout)
	if e := s.keybindings(s._g); e != nil {
		return e
	}
	s._users = make(map[string]*msg.UserInfo)
	s._groups = make(map[string]*msg.GroupInfo)
	s._sessions = make(map[string]*msg.ChatMsg)
	s.load_cache()
	return nil
}

func (s *_desktop_ui) Run() error {
	if err := s._g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (s *_desktop_ui) Release() {
	s._g.Close()
}

func (s *_desktop_ui) layout(g *gocui.Gui) error {
	var err error
	maxX, maxY := g.Size()
	if maxX > 120 {
		maxX = 120
	}
	SIDE_WID := maxX / 5
	if SIDE_WID < MIN_SIDE_VIEW {
		SIDE_WID = MIN_SIDE_VIEW
	}
	s._nav, err = g.SetView(NAV, -1, -1, SIDE_WID, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._nav.Highlight = true
		s._nav.SelBgColor = gocui.ColorGreen
		fmt.Fprintln(s._nav, "Session")
		fmt.Fprintln(s._nav, "Users")
		fmt.Fprintln(s._nav, "Chat Rooms")
		fmt.Fprint(s._nav, "Settings")
	}

	s._usrs, err = g.SetView(USRS, SIDE_WID, -1, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._usrs.Highlight = true
		s._usrs.SelBgColor = gocui.ColorGreen
		s._usrs.Editable = false
		s.refresh_users()
	}
	s._grps, err = g.SetView(GRPS, SIDE_WID, -1, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._grps.Highlight = true
		s._grps.SelBgColor = gocui.ColorGreen
		s._grps.Editable = false
		s._grps.Autoscroll = true
		s.refresh_groups()
	}
	s._settings, err = g.SetView(SETTINGS, SIDE_WID, -1, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._settings.Editable = false
		fmt.Fprintf(s._settings, "%s", SETTINGS)
	}
	s._ssn, err = g.SetView(SSN, SIDE_WID, -1, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._ssn.Highlight = true
		s._ssn.SelBgColor = gocui.ColorGreen
		s._ssn.Editable = false
		s.refresh_snss()
		if e := s.focus_ssn(g, nil); e != nil {
			return err
		}
	}

	s._info, err = g.SetView(INFO, -1, maxY-WIDGET_HGT, maxX, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._info.Editable = false
		s._info.Autoscroll = true
		fmt.Fprintf(s._info, "nothing...")
	}

	return nil
}

func (s *_desktop_ui) keybindings(g *gocui.Gui) error {
	// global
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, '1', gocui.ModAlt, s.focus_ssn); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, '2', gocui.ModAlt, s.focus_usrs); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, '3', gocui.ModAlt, s.focus_grps); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, '4', gocui.ModAlt, s.focus_settings); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, gocui.KeyTab, gocui.ModNone, s.focus_nav); err != nil {
		return err
	}

	// nav
	if err := set_key_binding(g, NAV, gocui.KeyArrowDown, gocui.ModNone, s.nav_cursor_down); err != nil {
		return err
	}
	if err := set_key_binding(g, NAV, gocui.KeyArrowUp, gocui.ModNone, s.nav_cursor_up); err != nil {
		return err
	}
	if err := set_key_binding(g, NAV, gocui.KeyEnter, gocui.ModNone, s.focus_panel); err != nil {
		return err
	}

	// ssns
	if err := set_key_binding(g, SSN, gocui.KeyArrowDown, gocui.ModNone, s.cursor_down); err != nil {
		return err
	}
	if err := set_key_binding(g, SSN, gocui.KeyArrowUp, gocui.ModNone, s.cursor_up); err != nil {
		return err
	}
	if err := set_key_binding(g, SSN, gocui.KeyEnter, gocui.ModNone, s.enter_ssn); err != nil {
		return err
	}

	// usrs
	if err := set_key_binding(g, USRS, gocui.KeyArrowDown, gocui.ModNone, s.cursor_down); err != nil {
		return err
	}
	if err := set_key_binding(g, USRS, gocui.KeyArrowUp, gocui.ModNone, s.cursor_up); err != nil {
		return err
	}
	if err := set_key_binding(g, USRS, gocui.KeyEnter, gocui.ModNone, s.chat_with_usr); err != nil {
		return err
	}

	// rooms
	if err := set_key_binding(g, GRPS, gocui.KeyArrowDown, gocui.ModNone, s.cursor_down); err != nil {
		return err
	}
	if err := set_key_binding(g, GRPS, gocui.KeyArrowUp, gocui.ModNone, s.cursor_up); err != nil {
		return err
	}
	if err := set_key_binding(g, GRPS, gocui.KeyEnter, gocui.ModNone, s.enter_room); err != nil {
		return err
	}

	// mocks
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlU, gocui.ModNone, s.rand_user); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlM, gocui.ModNone, s.rand_msg); err != nil {
		return err
	}

	return nil
}

func (s *_desktop_ui) show_info(msg string) {
	s._info.Clear()
	fmt.Fprintf(s._info, "\033[31;1m%s\033[0m", msg)
}

func (s *_desktop_ui) cursor_down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		_, cy := v.Cursor()
		ty := cy + 1
		if ty-oy >= s.NUMS[v.Name()] {
			ty = oy + s.NUMS[v.Name()] - 1
		}
		if err := v.SetCursor(ox, ty); err != nil {
			return err
		}
		fmt.Fprintf(s._info, "\nsel[%s]: %s", v.Name(), v.BufferLines()[ty])
	}
	return nil
}

func (s *_desktop_ui) cursor_up(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		_, cy := v.Cursor()
		ty := cy - 1
		if ty < oy {
			ty = oy
		}
		if err := v.SetCursor(ox, ty); err != nil && oy > 0 {
			return err
		}
		fmt.Fprintf(s._info, "\nsel[%s]: %s", v.Name(), v.BufferLines()[ty])
	}
	return nil
}
func (s *_desktop_ui) focus_ssn(g *gocui.Gui, v *gocui.View) error {
	return s.focus_view(g, SSN)
}
func (s *_desktop_ui) focus_usrs(g *gocui.Gui, v *gocui.View) error {
	return s.focus_view(g, USRS)
}
func (s *_desktop_ui) focus_grps(g *gocui.Gui, v *gocui.View) error {
	return s.focus_view(g, GRPS)
}
func (s *_desktop_ui) focus_settings(g *gocui.Gui, v *gocui.View) error {
	return s.focus_view(g, SETTINGS)
}
func (s *_desktop_ui) focus_nav(g *gocui.Gui, v *gocui.View) error {
	return s.focus_view(g, NAV)
}
func (s *_desktop_ui) nav_cursor_down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		_, cy := v.Cursor()
		ty := cy + 1
		if ty-oy >= s.LEN {
			ty = oy + s.LEN - 1
		}
		if err := v.SetCursor(ox, ty); err != nil {
			return err
		}
		to_top(g, s.VIEWS[ty])
	}
	return nil
}
func (s *_desktop_ui) nav_cursor_up(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		_, cy := v.Cursor()
		ty := cy - 1
		if ty < oy {
			ty = oy
		}
		if err := v.SetCursor(ox, ty); err != nil && oy > 0 {
			return err
		}
		to_top(g, s.VIEWS[ty])
	}
	return nil
}
func (s *_desktop_ui) focus_panel(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if cy >= s.LEN {
		return fmt.Errorf("bad nav item")
	}
	s.focus_view(g, s.VIEWS[cy])
	return nil
}
func (s *_desktop_ui) focus_view(g *gocui.Gui, vn string) error {
	var e error
	// recover _cur
	if s._cur != nil {
		s._cur.BgColor = gocui.ColorDefault
	}
	v, e := set_cur_top_view(g, vn)
	if e != nil {
		return e
	}
	s._cur = v
	// update active view
	s._cur.BgColor = gocui.ColorCyan
	return nil
}
