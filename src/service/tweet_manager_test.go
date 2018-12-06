package service_test

import (
	"testing"

	"github.com/twiteer/src/domain"

	"github.com/twiteer/src/service"
)

func TestPublishTweetIsSaved(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expeted tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	//Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required, recieved: " + err.Error())
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required, recieved: " + err.Error())
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	//Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "bbbbbbbababababababababababbabababababababababbababbbbbbbababababababababababbabababababababababbababbbbbbbababababababababababbabababababababababbaba"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	err = service.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "Exceeded 140 characters" {
		t.Error("Expected error is Exceeded 140 characters, recieved: " + err.Error())
	}
}
