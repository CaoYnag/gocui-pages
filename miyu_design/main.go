/*
MainUI
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CaoYnag/gocui"
)

var (
	_nav      *gocui.View
	_ssn      *gocui.View
	_usrs     *gocui.View
	_grps     *gocui.View
	_settings *gocui.View
	_views    []*gocui.View
	_cur      *gocui.View
	_info     *gocui.View

	VIEWS = []string{SSN, USRS, GRPS, SETTINGS}
	LEN   = len(VIEWS)

	NUMS = make(map[string]int)
)

const (
	NAV      = "nav"
	SSN      = "ssn"
	USRS     = "usrs"
	GRPS     = "grps"
	SETTINGS = "settings"
	INFO     = "info"

	SIDE_WID = 12

	INFO_HGT = 2
)

func layout(g *gocui.Gui) error {
	var err error
	maxX, maxY := g.Size()
	_nav, err = g.SetView(NAV, -1, -1, SIDE_WID, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_nav.Highlight = true
		_nav.SelBgColor = gocui.ColorGreen
		_nav.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(_nav, "Session")
		fmt.Fprintln(_nav, "Users")
		fmt.Fprintln(_nav, "Chat Rooms")
		fmt.Fprint(_nav, "Settings")
	}

	_usrs, err = g.SetView(USRS, SIDE_WID, -1, maxX, maxY-INFO_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_nav.SelBgColor = gocui.ColorGreen
		_nav.SelFgColor = gocui.ColorBlack
		add_usr("spes2", true)
		add_usr("spes3", false)
		add_usr("ehh1", true)
	}
	_grps, err = g.SetView(GRPS, SIDE_WID, -1, maxX, maxY-INFO_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		add_rooms("Just do IT")
		add_rooms("老 司 机 带 带 窝 ! ")
		add_rooms("游 戏 开 黑 搞 起 !")
		add_rooms("三 月 读 书 会 ")
	}
	_settings, err = g.SetView(SETTINGS, SIDE_WID, -1, maxX, maxY-INFO_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintf(_settings, "%s", SETTINGS)
	}
	_ssn, err = g.SetView(SSN, SIDE_WID, -1, maxX, maxY-INFO_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_ssn.Editable = true
		_ssn.Autoscroll = true
		if e := focus_view(g, SSN); e != nil {
			return err
		}
		add_ssn("ehh", "hello", time.Now(), true)
		add_ssn("ehh2", "hello", time.Now(), false)
		add_ssn("Just do IT", "yes!", time.Now(), false)
	}

	_info, err = g.SetView(INFO, -1, maxY-INFO_HGT, maxX, maxY)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		_info.Editable = false
		_info.Autoscroll = true
		fmt.Fprintf(_info, "nothing...")
	}

	_views = []*gocui.View{_ssn, _usrs, _grps, _settings}
	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func add_ssn(from, msg string, ts time.Time, unread bool) {
	NUMS[SSN]++
	if unread {
		fmt.Fprintf(_ssn, "\033[31;1m%s:%s\033[0m\n", from, msg)
	} else {
		fmt.Fprintf(_ssn, "%s:%s\n", from, msg)
	}
}

func add_usr(name string, online bool) {
	NUMS[USRS]++
	if online {
		fmt.Fprintf(_usrs, "\033[32;1m%s\033[0m\n", name)
	} else {
		fmt.Fprintf(_usrs, "\033[31;1m%s\033[0m\n", name)
	}
}

func add_rooms(name string) {
	NUMS[GRPS]++
	fmt.Fprintf(_grps, "%s\n", name)
}

func nav_cursor_down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		_, cy := v.Cursor()
		ty := cy + 1
		if ty-oy >= LEN {
			ty = oy + LEN - 1
		}
		if err := v.SetCursor(ox, ty); err != nil {
			return err
		}
		to_top(g, VIEWS[ty])
	}
	return nil
}

func nav_cursor_up(g *gocui.Gui, v *gocui.View) error {
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
		to_top(g, VIEWS[ty])
	}
	return nil
}

func cursor_down(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		_, cy := v.Cursor()
		ty := cy + 1
		if ty-oy >= NUMS[v.Name()] {
			ty = oy + NUMS[v.Name()] - 1
		}
		if err := v.SetCursor(ox, ty); err != nil {
			return err
		}
		fmt.Fprintf(_info, "\nsel[%s]: %s", v.Name(), v.BufferLines()[ty])
	}
	return nil
}

func cursor_up(g *gocui.Gui, v *gocui.View) error {
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
		fmt.Fprintf(_info, "\nsel[%s]: %s", v.Name(), v.BufferLines()[ty])
	}
	return nil
}

func chat_with_usr(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	fmt.Fprintf(_info, "\nchat to %s", v.BufferLines()[cy-oy])
	return nil
}
func enter_ssn(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	fmt.Fprintf(_info, "\nenter ssn: %s", v.BufferLines()[cy-oy])
	return nil
}
func enter_room(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	fmt.Fprintf(_info, "\nenter room: %s", v.BufferLines()[cy-oy])
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func keybindings(g *gocui.Gui) error {
	// global
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, focus_nav); err != nil {
		return err
	}

	// nav side
	if err := g.SetKeybinding(NAV, 'n', gocui.ModNone, nav_cursor_down); err != nil {
		return err
	}
	if err := g.SetKeybinding(NAV, 'p', gocui.ModNone, nav_cursor_up); err != nil {
		return err
	}
	if err := g.SetKeybinding(NAV, 'e', gocui.ModNone, change_view); err != nil {
		return err
	}

	// ssns
	if err := g.SetKeybinding(SSN, 'n', gocui.ModNone, cursor_down); err != nil {
		return err
	}
	if err := g.SetKeybinding(SSN, 'p', gocui.ModNone, cursor_up); err != nil {
		return err
	}
	if err := g.SetKeybinding(SSN, 'e', gocui.ModNone, enter_ssn); err != nil {
		return err
	}

	// usrs
	if err := g.SetKeybinding(USRS, 'n', gocui.ModNone, cursor_down); err != nil {
		return err
	}
	if err := g.SetKeybinding(USRS, 'p', gocui.ModNone, cursor_up); err != nil {
		return err
	}
	if err := g.SetKeybinding(USRS, 'e', gocui.ModNone, chat_with_usr); err != nil {
		return err
	}

	// rooms
	if err := g.SetKeybinding(GRPS, 'n', gocui.ModNone, cursor_down); err != nil {
		return err
	}
	if err := g.SetKeybinding(GRPS, 'p', gocui.ModNone, cursor_up); err != nil {
		return err
	}
	if err := g.SetKeybinding(GRPS, 'e', gocui.ModNone, enter_room); err != nil {
		return err
	}

	return nil
}

func focus_view(g *gocui.Gui, name string) error {
	var e error
	// recover _cur
	if _cur != nil {
		_cur.BgColor = gocui.ColorWhite
	}
	v, e := g.SetCurrentView(name)
	if e != nil {
		return e
	}
	_cur = v
	// update active view
	_cur.BgColor = gocui.ColorCyan
	return nil
}

func to_top(g *gocui.Gui, name string) error {
	_, e := g.SetViewOnTop(name)
	return e
}

func change_view(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if cy >= LEN {
		return fmt.Errorf("bad nav item")
	}
	focus_view(g, VIEWS[cy])
	return nil
}

func focus_nav(g *gocui.Gui, v *gocui.View) error {
	return focus_view(g, NAV)
}
