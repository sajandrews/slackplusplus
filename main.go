package main

import (
	"flag"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/sajandrews/slackplusplus/plusplus"
)

func main() {
	oauthAccessToken := flag.String("token", "", "oauth access token")

	flag.Parse()

	if *oauthAccessToken == "" {
		fmt.Println("Please pass in a token -token=sometoken")
		return
	}

	api := slack.New(*oauthAccessToken)
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
