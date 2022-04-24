package ui

import (
	"fmt"
	"gocui-demo/miyu_demo/mocks/msg"
	"time"

	"github.com/jroimartin/gocui"
)

func (s *_desktop_ui) chat_with_usr(g *gocui.Gui, v *gocui.View) error {
	_, oy := v.Origin()
	_, cy := v.Cursor()
	_jump_to(GetChat(v.BufferLines()[cy-oy]))
	return gocui.ErrQuit
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

// user stat change
func (s *_desktop_ui) on_user(u *msg.UserInfo) {
	u, found := s._users[u.Name]
	if !found {
		// new register

	}
}

// TODO finish groups in future
func (s *_desktop_ui) on_group() {}

// rcv chat msg TODO user msg only, finish group msg in future
func (s *_desktop_ui) on_msg(m *msg.ChatMsg) {}

// load user/group/session cache in init
func (s *_desktop_ui) load_cache() {
	// NOTICE just create fake data in this demo.

	// users
	s._users["spes"] = &msg.UserInfo{
		Name:  "spes",
		Nick:  "Spes",
		Sex:   "male",
		State: msg.STATE_ONLINE,
	}
	s._users["test1"] = &msg.UserInfo{
		Name:  "test1",
		Nick:  "Test1",
		Sex:   "male",
		State: msg.STATE_ONLINE,
	}
	s._users["test2"] = &msg.UserInfo{
		Name:  "test2",
		Nick:  "Test2",
		Sex:   "male",
		State: msg.STATE_ONLINE,
	}

	// groups
	// sessions
	s._sessions["test1"] = &msg.ChatMsg{
		From:    "test1",
		To:      "spes",
		Content: "hello",
		Ts:      time.Now(),
	}
	s._sessions["test2"] = &msg.ChatMsg{
		From:    "test2",
		To:      "spes",
		Content: "testtestpingpingping",
		Ts:      time.Now(),
	}
}

func (s *_desktop_ui) refresh_users() {

}
func (s *_desktop_ui) refresh_groups() {

}
func (s *_desktop_ui) refresh_snss() {

}
