package service_test

import (
	"testing"

	"github.com/twiteer/src/service"
)

func TestPublishTweetIsSaved(t *testing.T) {
	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}
}
