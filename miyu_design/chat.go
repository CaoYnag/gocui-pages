/*
MainUI
*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

var (
	_ctx  *gocui.View
	_inp  *gocui.View // input area
	_stat *gocui.View // user info
	_info *gocui.View // info bar
)

const ()

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
