package ui

import (
	"fmt"
	"time"

	"github.com/CaoYnag/gocui"
)

const (
	// some const string for display
	S_STAT_LEN  = 5
	S_STAT_DONE = "\033[32;1mDONE\033[0m"
	S_STAT_FAIL = "\033[31;1mFAIL\033[0m"
)

type _sync_ui struct {
	_g    *gocui.Gui
	_v    *gocui.View
	_stat int
	_sfmt string
}

func GetSync() *_sync_ui {
	rslt := &_sync_ui{
		_stat: 0,
	}
	return rslt
}

func (s *_sync_ui) Init() error {
	var err error
	s._g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}

	s._g.SetManagerFunc(s.layout)
	if e := s.keybindings(s._g); e != nil {
		return e
	}
	return nil
}

func (s *_sync_ui) Run() error {
	go sync(s)

	if err := s._g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func (s *_sync_ui) Release() {
	s._g.Close()
}

func (s *_sync_ui) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if maxX > 80 {
		maxX = 80
	}
	if maxY > 36 {
		maxY = 36
	}
	s._sfmt = fmt.Sprintf("%%-%ds", maxX-2-S_STAT_LEN) // 2: border

	var err error
	s._v, err = g.SetView("sync", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		s._v.Title = "Login"
		s._v.Editable = false
		s._v.Wrap = true
		fmt.Fprintln(s._v, s._sfmt)
	}

	return nil
}
func (s *_sync_ui) keybindings(g *gocui.Gui) error {
	if err := set_key_binding(g, GLOBAL, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func (s *_sync_ui) to_desktop(*gocui.Gui) error {
	_jump_to(GetDesktop())
	return gocui.ErrQuit
}
func (s *_sync_ui) show_info(msg string) {
	fmt.Fprintf(s._v, "%s", msg)
	s._g.Update(func(*gocui.Gui) error { return nil })
}
func (s *_sync_ui) show_msg(msg string) {
	s.show_info(fmt.Sprintf(s._sfmt, msg))
}
func (s *_sync_ui) stat_done() {
	s.show_info(S_STAT_DONE + "\n")
}
func (s *_sync_ui) stat_fail() {
	s.show_info(S_STAT_FAIL + "\n")
}

func sync(s *_sync_ui) {
	s.load_cache()
	s.sync_user()
	s.sync_sessions()
	s.sync_settings()

	// s._g.Update(s.to_desktop) // jump to desktop
}

func (s *_sync_ui) sync_sessions() {
	s.show_msg("syncing sessions")
	<-time.After(time.Millisecond * 700)
	s.stat_done()
}
func (s *_sync_ui) sync_user() {
	s.show_msg("syncing user")
	<-time.After(time.Millisecond * 700)
	s.stat_done()
}
func (s *_sync_ui) load_cache() {
	s.show_msg("loading cache")
	<-time.After(time.Millisecond * 700)
	s.stat_done()
}
func (s *_sync_ui) sync_settings() {
	s.show_msg("syncing settings")
	<-time.After(time.Millisecond * 700)
	s.stat_fail()
}
