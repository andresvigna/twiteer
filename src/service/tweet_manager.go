package service

var Tweet string

func PublishTweet(t string) {
	Tweet = t
}

func GetTweet() string {
	return Tweet
}
