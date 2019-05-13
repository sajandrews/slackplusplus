package plusplus

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sajandrews/slackplusplus/redis"

	"github.com/nlopes/slack"
)

type Slacker interface {
	GetInfo() *slack.Info
	GetUserInfo(string) (*slack.User, error)
	SendMessage(*slack.OutgoingMessage)
	NewOutgoingMessage(string, string, ...slack.RTMsgOption) *slack.OutgoingMessage
	OpenIMChannel(string) (bool, bool, string, error)
}

func Process(rtm Slacker, ev *slack.MessageEvent) error {
	var err error

	botID := fmt.Sprintf("<@%s> ", rtm.GetInfo().User.ID)

	if strings.Contains(ev.Msg.Text, botID) {
		err = processBotMessage(rtm, ev, botID)
	} else {
		//Lets see if we need to plus anyone!
		err = processPossiblePlusMessage(rtm, ev)
	}

	return err
}

func processPossiblePlusMessage(rtm Slacker, ev *slack.MessageEvent) error {
	re := regexp.MustCompile(`<@([A-Z0-9]*)>\s*\+\+`)
	matches := re.FindSubmatch([]byte(ev.Msg.Text))

	if len(matches) == 2 {
		rdb := redis.GetClient()

		userID := string(matches[1])

		numPlusses, err := rdb.Incr(redis.GetKey(userID)).Result()

		if err != nil {
			return err
		}

		user, err := rtm.GetUserInfo(userID)

		if err != nil {
			return err
		}

		message := fmt.Sprintf("%s now has %d plusses!", user.Name, numPlusses)

		rtm.SendMessage(rtm.NewOutgoingMessage(message, ev.Channel))

		return nil
	}

	re = regexp.MustCompile(`<@([A-Z0-9]*)>\s*--`)
	matches = re.FindSubmatch([]byte(ev.Msg.Text))

	if len(matches) == 2 {
		rdb := redis.GetClient()

		userID := string(matches[1])

		numPlusses, err := rdb.Decr(redis.GetKey(userID)).Result()

		if err != nil {
			return err
		}

		user, err := rtm.GetUserInfo(userID)

		if err != nil {
			return err
		}

		message := fmt.Sprintf("%s now has %d plusses!", user.Name, numPlusses)

		rtm.SendMessage(rtm.NewOutgoingMessage(message, ev.Channel))
	}

	return nil
}

func processBotMessage(rtm Slacker, ev *slack.MessageEvent, botID string) error {
	message := strings.Replace(ev.Msg.Text, botID, "", -1)

	if message == "help" {
		err := postHelp(rtm, ev)

		if err != nil {
			return err
		}
	}

	//Lets see if they are asking for stats?
	re := regexp.MustCompile(`<@([A-Z0-9]*)>`)
	matches := re.FindSubmatch([]byte(message))

	if len(matches) == 2 {
		rdb := redis.GetClient()

		userID := string(matches[1])

		numPlusses, err := rdb.Get(redis.GetKey(userID)).Result()

		if err != nil {
			if err.Error() == "redis: nil" {
				numPlusses = "0"
			} else {
				return err
			}
		}

		user, err := rtm.GetUserInfo(userID)

		if err != nil {
			return err
		}

		message := fmt.Sprintf("%s has %s plusses", user.Name, numPlusses)

		rtm.SendMessage(rtm.NewOutgoingMessage(message, ev.Channel))
	}

	return nil

}

func postHelp(rtm Slacker, ev *slack.MessageEvent) error {
	_, _, cID, err := rtm.OpenIMChannel(ev.User)

	if err != nil {
		return err
	}

	rtm.SendMessage(rtm.NewOutgoingMessage(getHelpMessage(), cID))

	return nil
}

func getHelpMessage() string {
	return `
Hey!
Plus plus is super simple:
@user++ to give kudos
@user-- to remove kudos
@plusplus @username to get stats
@plusplus help for this message`
}
