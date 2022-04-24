package msg

import "time"

const (
	STATE_ONLINE  = 0
	STATE_OFFLINE = 1
)

type UserInfo struct {
	Avatar string
	Name   string
	Nick   string
	Sex    string
	State  int // online offline
	Birth  time.Time
}

type GroupInfo struct {
	ICON   string
	Name   string
	Create time.Time
	Member []string
}

type ChatMsg struct {
	From    string    // user
	Device  string    // device
	To      string    // user or a chat room, or a dev in future?
	Content string    // encrypted data
	Ts      time.Time // timestamp
	Ref     string    // reference or reply
}
