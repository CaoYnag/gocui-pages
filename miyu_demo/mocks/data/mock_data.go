package data

import (
	"gocui-demo/miyu_demo/mocks/msg"
	"math/rand"
	"strconv"
	"time"
)

var (
	CODE int = 0 // used for generate fake data

	// come from `fortune`
	text = []string{
		"It ain't so much the things we don't know that get us in trouble.  It's the things we know that ain't so.",
		"Your lover will never wish to leave you.",
		"I kissed my first girl and smoked my first cigarette on the same day.I haven't had time for tobacco since.",
		"Going to church does not make a person religious, nor does going to school make a person educated, any more than going to a garage makes a person a car.",
		"The primary requisite for any new tax law is for it to exempt enough voters to win the next election.",
		"I'd love to go out with you, but I'm doing door-to-door collecting for staticcling.",
		"Life is like an analogy.",
	}
)

func RandUser() *msg.UserInfo {
	N := strconv.Itoa(CODE)
	CODE++
	return &msg.UserInfo{
		Name:  "gen_" + N,
		Nick:  "Test" + N,
		Sex:   "male",
		State: rand.Intn(2),
	}
}

func RandMsg(from string) *msg.ChatMsg {
	return &msg.ChatMsg{
		From:    from,
		Device:  "dev",
		Content: text[rand.Intn(len(text))],
		Ts:      time.Now(),
	}
}
