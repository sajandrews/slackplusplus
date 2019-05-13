package main

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/sajandrews/slackplusplus/plusplus"
)

func main() {
	api := slack.New("xoxb-455086501812-611131728262-gbQK2uNh6VgISrCfU65TNOzk")

	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			err := plusplus.Process(rtm, ev)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
