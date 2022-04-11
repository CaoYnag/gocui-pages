package ui

import (
	"fmt"
	"time"

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
	s._g.Highlight = true
	s._g.SelFgColor = gocui.ColorRed

	s._g.SetManagerFunc(s.layout)
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
	if e := s.keybindings(s._g); e != nil {
		return e
	}
	var err error
	maxX, maxY := g.Size()
	if maxX > 120 {
		maxX = 120
	}
	SIDE_WID := maxX / 5
	if SIDE_WID < MIN_SIDE_VIEW {
		SIDE_WID = MIN_SIDE_VIEW
	}
	s._nav, err = g.SetView(NAV, -1, -1, SIDE_WID, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._nav.Highlight = true
		s._nav.SelBgColor = gocui.ColorGreen
		s._nav.SelFgColor = gocui.ColorBlack
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
		s._usrs.SelBgColor = gocui.ColorGreen
		s._usrs.SelFgColor = gocui.ColorBlack
		s._usrs.Editable = false
		s.add_usr("spes2", true)
		s.add_usr("spes3", false)
		s.add_usr("ehh1", true)
	}
	s._grps, err = g.SetView(GRPS, SIDE_WID, -1, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._grps.Editable = false
		s.add_rooms("Just do IT")
		s.add_rooms("老 司 机 带 带 窝 ! ")
		s.add_rooms("游 戏 开 黑 搞 起 !")
		s.add_rooms("三 月 读 书 会 ")
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
		s._ssn.Editable = false
		s._ssn.Autoscroll = false
		if e := s.focus_ssn(g, nil); e != nil {
			return err
		}
		s.add_ssn("ehh", "hello", time.Now(), true)
		s.add_ssn("ehh2", "hello", time.Now(), false)
		s.add_ssn("Just do IT", "yes!", time.Now(), false)
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
	if err := set_key_binding(g, "", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := set_key_binding(g, "", '1', gocui.ModAlt, s.focus_ssn); err != nil {
		return err
	}
	if err := set_key_binding(g, "", '2', gocui.ModAlt, s.focus_usrs); err != nil {
		return err
	}
	if err := set_key_binding(g, "", '3', gocui.ModAlt, s.focus_grps); err != nil {
		return err
	}
	if err := set_key_binding(g, "", '4', gocui.ModAlt, s.focus_settings); err != nil {
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

	return nil
}

func (s *_desktop_ui) show_info(msg string) {
	s._info.Clear()
	fmt.Fprintf(s._info, "\033[31;1m%s\033[0m", msg)
}

func (s *_desktop_ui) add_ssn(from, msg string, ts time.Time, unread bool) {
	s.NUMS[SSN]++
	if unread {
		fmt.Fprintf(s._ssn, "\033[31;1m%s:%s\033[0m\n", from, msg)
	} else {
		fmt.Fprintf(s._ssn, "%s:%s\n", from, msg)
	}
}

func (s *_desktop_ui) add_usr(name string, online bool) {
	s.NUMS[USRS]++
	if online {
		fmt.Fprintf(s._usrs, "\033[32;1m%s\033[0m\n", name)
	} else {
		fmt.Fprintf(s._usrs, "\033[31;1m%s\033[0m\n", name)
	}
}

func (s *_desktop_ui) add_rooms(name string) {
	s.NUMS[GRPS]++
	fmt.Fprintf(s._grps, "%s\n", name)
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

func (s *_desktop_ui) chat_with_usr(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	fmt.Fprintf(s._info, "\nchat to %s", v.BufferLines()[cy-oy])
	return nil
}
func (s *_desktop_ui) enter_ssn(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	fmt.Fprintf(s._info, "\nenter ssn: %s", v.BufferLines()[cy-oy])
	return nil
}
func (s *_desktop_ui) enter_room(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	fmt.Fprintf(s._info, "\nenter room: %s", v.BufferLines()[cy-oy])
	return nil
}
func (s *_desktop_ui) focus_ssn(g *gocui.Gui, v *gocui.View) error {
	set_cur_top_view(g, SSN)
	return nil
}
func (s *_desktop_ui) focus_usrs(g *gocui.Gui, v *gocui.View) error {
	set_cur_top_view(g, USRS)
	return nil
}
func (s *_desktop_ui) focus_grps(g *gocui.Gui, v *gocui.View) error {
	set_cur_top_view(g, GRPS)
	return nil
}
func (s *_desktop_ui) focus_settings(g *gocui.Gui, v *gocui.View) error {
	set_cur_top_view(g, SETTINGS)
	return nil
}
