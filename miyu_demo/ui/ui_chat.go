package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

type _chat_ui struct {
	_g    *gocui.Gui
	_ctx  *gocui.View // msg area
	_inp  *gocui.View // input area
	_stat *gocui.View // user info
	_info *gocui.View

	_data struct {
		name string
	}
}

func GetChat(name string) *_chat_ui {
	rslt := &_chat_ui{}
	rslt._data.name = name
	return rslt
}

func (s *_chat_ui) Init() error {
	var err error
	s._g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	s._g.Cursor = true

	s._g.SetManagerFunc(s.layout)
	return nil
}

func (s *_chat_ui) Run() error {
	if err := s._g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (s *_chat_ui) Release() {
	s._g.Close()
}

func (s *_chat_ui) layout(g *gocui.Gui) error {
	if e := s.keybindings(s._g); e != nil {
		return e
	}
	maxX, maxY := g.Size()

	var err error
	s._stat, err = g.SetView(STAT, -1, -1, maxX, STAT_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(s._stat, s._data.name) // user stat here
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
	s._inp, err = g.SetView(INP, -1, maxY-WIDGET_HGT-INP_HGT, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._inp.Editable = true
		s._inp.Autoscroll = true
		g.SetCurrentView(INP)
	}
	s._ctx, err = g.SetView(CTX, -1, STAT_HGT, maxX, maxY-WIDGET_HGT-INP_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._ctx.Editable = false
		s._ctx.Autoscroll = true
		s.show_msg(s._data.name, "在 嘛 ?", time.Now())
		s.show_msg("spes", "不 在 !", time.Now())
		s.show_msg("spes", "里 四 居 !", time.Now())
	}
	return nil
}

func (s *_chat_ui) keybindings(g *gocui.Gui) error {
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	// input
	if err := set_key_binding(g, INP, gocui.KeyEnter, gocui.ModNone, s.snd_msg); err != nil {
		return err
	}
	// key conflict
	if err := set_key_binding(g, INP, gocui.KeyCtrlSpace, gocui.ModNone, s.new_line); err != nil {
		return err
	}
	return nil
}
func (s *_chat_ui) show_msg(name, msg string, ts time.Time) {
	if name == "spes" {
		// self
		fmt.Fprintf(s._ctx, "\033[32;1m%s %s\033[0m\n", name, ts.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Fprintf(s._ctx, "\033[34;1m%s %s\033[0m\n", name, ts.Format("2006-01-02 15:04:05"))
	}
	fmt.Fprintf(s._ctx, "%s\n\n", msg)
}

func (s *_chat_ui) show_info(msg string) {
	s._info.Clear()
	fmt.Fprintf(s._info, "\033[31;1m%s\033[0m", msg)
}

func (s *_chat_ui) snd_msg(g *gocui.Gui, v *gocui.View) error {
	m := s._inp.Buffer()
	m = m[:len(m)-1]
	s.show_msg("spes", m, time.Now())
	s._inp.Clear()
	x, y := s._inp.Origin()
	s._inp.SetCursor(x, y)
	return nil
}

func (s *_chat_ui) new_line(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintf(s._inp, "\n")
	return nil
}
