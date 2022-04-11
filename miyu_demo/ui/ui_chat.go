/*
Chat UI
TODO:
- fix chinese display
*/

package ui

import (
	"fmt"
	"time"

	"github.com/d4l3k/go-highlight"
	"github.com/jroimartin/gocui"
)

type _chat_ui struct {
	_g      *gocui.Gui
	_ctx    *gocui.View // msg area
	_inp_nm *gocui.View // input area
	_inp    int
	_vinp   *gocui.View
	_inp_md *gocui.View // markdown input area
	_stat   *gocui.View // user info
	_info   *gocui.View

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
	s._inp = INPUT_NORMAL // default input method could be conf in settings

	s._g.SetManagerFunc(s.layout)
	if e := s.keybindings(s._g); e != nil {
		return e
	}
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
	s._inp_nm, err = g.SetView(INP_NM, -1, maxY-WIDGET_HGT-INP_HGT, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._inp_nm.Editable = true
		s._inp_nm.Autoscroll = true
		g.SetCurrentView(INP_NM)
	}
	s._inp_md, err = g.SetView(INP_MD, -1, maxY-WIDGET_HGT-INP_HGT, maxX, maxY-WIDGET_HGT)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._inp_md.Editable = true
		s._inp_md.Autoscroll = true
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
	// Global
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlC, gocui.ModNone, s.to_desktop); err != nil {
		return err
	}
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlS, gocui.ModNone, s.switch_inp); err != nil {
		return err
	}

	// input normal
	if err := set_key_binding(g, INP_NM, gocui.KeyEnter, gocui.ModNone, s.snd_msg); err != nil {
		return err
	}
	// TODO key conflict
	if err := set_key_binding(g, INP_NM, gocui.KeyCtrlSpace, gocui.ModNone, s.new_line); err != nil {
		return err
	}

	// input markdown
	if err := set_key_binding(g, INP_MD, gocui.KeyCtrlSpace, gocui.ModNone, s.snd_msg); err != nil {
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

func highlight_code(lang, code string) string {
	s, _ := highlight.Term(lang, []byte(code))
	return string(s)
}
func (s *_chat_ui) snd_msg(g *gocui.Gui, v *gocui.View) error {
	m := s._vinp.Buffer()
	m = m[:len(m)-1]
	if s._inp == INPUT_MD {
		m = highlight_code("c++", m)
	}
	s.show_msg("spes", m, time.Now())
	s._vinp.Clear()
	x, y := s._vinp.Origin()
	s._vinp.SetCursor(x, y)
	return nil
}

func (s *_chat_ui) new_line(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintf(s._inp_nm, "\n")
	return nil
}

func (s *_chat_ui) to_desktop(g *gocui.Gui, v *gocui.View) error {
	_jump_to(GetDesktop())
	return gocui.ErrQuit
}

func (s *_chat_ui) switch_inp(g *gocui.Gui, v *gocui.View) error {
	if s._inp == INPUT_NORMAL {
		s._vinp = s._inp_md
		s.show_info("switch to markdown input, use Ctrl+Space to send msg.")
	} else {
		s._vinp = s._inp_nm
		s.show_info("switch to normal input, use Ctrl+Space for newline, enter to send msg.")
	}
	s._inp = 1 - s._inp
	set_cur_top_view(g, s._vinp.Name())
	return nil
}
