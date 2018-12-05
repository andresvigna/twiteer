package domain

import "time"

type Tweet struct {
	User string
	Text string
	Date *time.Time
}
∫
func NewTweet(user string, text string) *Tweet {
	var tim time.Time = time.Now()
	return &Tweet{user, text, &tim}
}
