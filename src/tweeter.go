package main

import (
	"github.com/abiosoft/ishell"
	"github.com/twiteer/src/domain"
	"github.com/twiteer/src/service"
)

func main() {
	var tweet *domain.Tweet

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Who are you: ")

			user := c.ReadLine()
			c.Print("Write your msg: ")
			text := c.ReadLine()

			tweet = domain.NewTweet(user, text)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()
			if tweet != nil {
				c.Println(*tweet)
			}

			return
		},
	})

	shell.Run()

}
