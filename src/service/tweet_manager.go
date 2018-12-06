package service

import (
	"fmt"

	"github.com/twiteer/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	Tweet = tweet
	if Tweet.User == "" {
		return fmt.Errorf("user is required")
	}

	if Tweet.Text == "" {
		return fmt.Errorf("text is required")
	}

	if len(Tweet.Text) > 140 {
		return fmt.Errorf("Exceeded 140 characters")
	}
	return nil
}

func GetTweet() *domain.Tweet {
	return Tweet
}
