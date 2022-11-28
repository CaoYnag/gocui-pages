package ui

import (
	"fmt"
	"gocui-demo/miyu_demo/mocks/data"
	"gocui-demo/miyu_demo/mocks/msg"
	"math/rand"
	"sort"
	"time"

	"github.com/CaoYnag/gocui"
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

// user stat change
func (s *_desktop_ui) on_user(u *msg.UserInfo) {
	s._users[u.Name] = u
	s.refresh_users()
}

// TODO finish groups in future
func (s *_desktop_ui) on_group() {}

// rcv chat msg TODO user msg only, finish group msg in future
func (s *_desktop_ui) on_msg(m *msg.ChatMsg) {
	s._sessions[m.From] = m
	s.refresh_snss()
}

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
		Content: "fuck u nvidia!",
		Ts:      time.Now(),
	}
	s._sessions["test2"] = &msg.ChatMsg{
		From:    "test2",
		To:      "spes",
		Content: "终 有 一 天,hhh",
		Ts:      time.Now(),
	}
}

func (s *_desktop_ui) refresh_users() {
	// TODO do not sort, just show all user by default.
	var usrs []*msg.UserInfo
	for key := range s._users {
		usrs = append(usrs, s._users[key])
	}

	s._usrs.Clear()
	s.NUMS[USRS] = len(usrs)
	for idx := range usrs {
		u := usrs[idx]
		if u.State == msg.STATE_ONLINE {
			fmt.Fprintf(s._usrs, "\033[32;1m%s\033[0m(%s)\n", u.Nick, u.Name)
		} else {
			fmt.Fprintf(s._usrs, "\033[31;1m%s\033[0m(%s)\n", u.Nick, u.Name)
		}
	}
}
func (s *_desktop_ui) refresh_groups() {
	s.NUMS[GRPS] = 0
}
func (s *_desktop_ui) refresh_snss() {
	// sort by ts
	var ssns []*msg.ChatMsg = nil
	for key := range s._sessions {
		ssns = append(ssns, s._sessions[key])
	}
	sort.SliceStable(ssns, func(i, j int) bool {
		return ssns[i].Ts.After(ssns[j].Ts)
	})

	s._ssn.Clear()
	for idx := range ssns {
		ssn := ssns[idx]
		fmt.Fprintf(s._ssn, "\x1b[31;1m%s:%s\x1b[0m\n", ssn.From, ssn.Content)
	}
	s.NUMS[SSN] = len(ssns)
}

/**************************************/
/**       some mocks functions       **/
/**************************************/

func (s *_desktop_ui) rand_user(g *gocui.Gui, v *gocui.View) error {
	s.on_user(data.RandUser())
	return nil
}

func (s *_desktop_ui) rand_msg(g *gocui.Gui, v *gocui.View) error {
	var usrs []*msg.UserInfo
	for key := range s._users {
		usrs = append(usrs, s._users[key])
	}
	d := data.RandMsg(usrs[rand.Intn(len(usrs))].Name)
	d.To = "spes"
	s.on_msg(d)
	return nil
}
