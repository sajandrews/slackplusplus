package plusplus_test

import (
	"fmt"
	"testing"

	re "github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"

	"github.com/sajandrews/slackplusplus/mocks"
	"github.com/sajandrews/slackplusplus/plusplus"
	"github.com/sajandrews/slackplusplus/redis"
)

func getHelpMessage() string {
	return `
Hey!
Plus plus is super simple:
@user++ to give kudos
@user-- to remove kudos
@plusplus @username to get stats
@plusplus help for this message`
}

func TestGetHelp(t *testing.T) {
	assert := assert.New(t)

	testIMID := "TESTIMID"
	testID := 12
	testBotID := "TESTBOTID"

	slackerCTRL := gomock.NewController(t)
	mockSlacker := mocks.NewMockSlacker(slackerCTRL)

	mockOutgoingMessage := &slack.OutgoingMessage{
		ID:      testID,
		Type:    "message",
		Channel: testIMID,
		Text:    getHelpMessage(),
	}

	mockInfo := &slack.Info{
		User: &slack.UserDetails{ID: testBotID},
	}

	gomock.InOrder(
		mockSlacker.EXPECT().GetInfo().Return(mockInfo).Times(1),
		mockSlacker.EXPECT().OpenIMChannel(testBotID).Return(true, true, testIMID, nil).Times(1),
		mockSlacker.EXPECT().NewOutgoingMessage(getHelpMessage(), testIMID).Return(mockOutgoingMessage).Times(1),
		mockSlacker.EXPECT().SendMessage(mockOutgoingMessage).Return().Times(1),
	)

	mockEV := &slack.MessageEvent{
		Msg: slack.Msg{
			Type: "message",
			Text: fmt.Sprintf("%s help", testBotID),
		},
	}

	err := plusplus.Process(mockSlacker, mockEV)

	assert.Nil(err)
}

func TestGetStats(t *testing.T) {
	assert := assert.New(t)

	testChannelID := "TESTCHANNELID"
	testID := 12
	testBotID := "TESTBOTID"
	testUserID := "<@USER>"
	testUserName := "TESTUSERNAME"
	testNumPlusses := "2"
	testMessage := fmt.Sprintf("%s %s", testBotID, testUserID)

	testOutGoingMessage := fmt.Sprintf("%s has %s plusses", testUserName, testNumPlusses)

	slackerCTRL := gomock.NewController(t)
	mockSlacker := mocks.NewMockSlacker(slackerCTRL)

	mockOutgoingMessage := &slack.OutgoingMessage{
		ID:      testID,
		Type:    "message",
		Channel: testChannelID,
		Text:    testOutGoingMessage,
	}

	mockInfo := &slack.Info{
		User: &slack.UserDetails{ID: testBotID},
	}

	mockUser := &slack.User{
		Name: testUserName,
	}

	gomock.InOrder(
		mockSlacker.EXPECT().GetInfo().Return(mockInfo).Times(1),
		mockSlacker.EXPECT().GetUserInfo("USER").Return(mockUser, nil).Times(1),
		mockSlacker.EXPECT().NewOutgoingMessage(testOutGoingMessage, testChannelID).Return(mockOutgoingMessage).Times(1),
		mockSlacker.EXPECT().SendMessage(mockOutgoingMessage).Return().Times(1),
	)

	redisCTRL := gomock.NewController(t)
	mockRediser := mocks.NewMockRediser(redisCTRL)

	mockGetResult := re.NewStringCmd("get", testNumPlusses)

	gomock.InOrder(
		mockRediser.EXPECT().Get(redis.GetKey("USER")).Return(mockGetResult).Times(1),
	)

	redis.TestingRDB = mockRediser

	mockEV := &slack.MessageEvent{
		Msg: slack.Msg{
			Type:    "message",
			Text:    testMessage,
			Channel: testChannelID,
		},
	}

	err := plusplus.Process(mockSlacker, mockEV)

	assert.Nil(err)

	redis.TestingRDB = nil
}

func TestAddPlus(t *testing.T) {
	assert := assert.New(t)

	testChannelID := "TESTCHANNELID"
	testID := 12
	testBotID := "TESTBOTID"
	testUserID := "<@USER>"
	testUserName := "TESTUSERNAME"
	testNumPlusses := 2
	testMessage := fmt.Sprintf("%s ++", testUserID)

	//TODO https://github.com/go-redis/redis/blob/3f9b4d7004b1270253e54665d3b233afb06e9784/command.go#L393 set IntCmd val???
	testOutGoingMessage := fmt.Sprintf("%s now has %d plusses!", testUserName, 0)

	slackerCTRL := gomock.NewController(t)
	mockSlacker := mocks.NewMockSlacker(slackerCTRL)

	mockOutgoingMessage := &slack.OutgoingMessage{
		ID:      testID,
		Type:    "message",
		Channel: testChannelID,
		Text:    testOutGoingMessage,
	}

	mockInfo := &slack.Info{
		User: &slack.UserDetails{ID: testBotID},
	}

	mockUser := &slack.User{
		Name: testUserName,
	}

	gomock.InOrder(
		mockSlacker.EXPECT().GetInfo().Return(mockInfo).Times(1),
		mockSlacker.EXPECT().GetUserInfo("USER").Return(mockUser, nil).Times(1),
		mockSlacker.EXPECT().NewOutgoingMessage(testOutGoingMessage, testChannelID).Return(mockOutgoingMessage).Times(1),
		mockSlacker.EXPECT().SendMessage(mockOutgoingMessage).Return().Times(1),
	)

	redisCTRL := gomock.NewController(t)
	mockRediser := mocks.NewMockRediser(redisCTRL)

	mockGetResult := re.NewIntCmd("incr", testNumPlusses)

	gomock.InOrder(
		mockRediser.EXPECT().Incr(redis.GetKey("USER")).Return(mockGetResult).Times(1),
	)

	redis.TestingRDB = mockRediser

	mockEV := &slack.MessageEvent{
		Msg: slack.Msg{
			Type:    "message",
			Text:    testMessage,
			Channel: testChannelID,
		},
	}

	err := plusplus.Process(mockSlacker, mockEV)

	assert.Nil(err)

	redis.TestingRDB = nil
}

func TestRemovePlus(t *testing.T) {
	assert := assert.New(t)

	testChannelID := "TESTCHANNELID"
	testID := 12
	testBotID := "TESTBOTID"
	testUserID := "<@USER>"
	testUserName := "TESTUSERNAME"
	testNumPlusses := 2
	testMessage := fmt.Sprintf("%s --", testUserID)

	//TODO https://github.com/go-redis/redis/blob/3f9b4d7004b1270253e54665d3b233afb06e9784/command.go#L393 set IntCmd val???
	testOutGoingMessage := fmt.Sprintf("%s now has %d plusses!", testUserName, 0)

	slackerCTRL := gomock.NewController(t)
	mockSlacker := mocks.NewMockSlacker(slackerCTRL)

	mockOutgoingMessage := &slack.OutgoingMessage{
		ID:      testID,
		Type:    "message",
		Channel: testChannelID,
		Text:    testOutGoingMessage,
	}

	mockInfo := &slack.Info{
		User: &slack.UserDetails{ID: testBotID},
	}

	mockUser := &slack.User{
		Name: testUserName,
	}

	gomock.InOrder(
		mockSlacker.EXPECT().GetInfo().Return(mockInfo).Times(1),
		mockSlacker.EXPECT().GetUserInfo("USER").Return(mockUser, nil).Times(1),
		mockSlacker.EXPECT().NewOutgoingMessage(testOutGoingMessage, testChannelID).Return(mockOutgoingMessage).Times(1),
		mockSlacker.EXPECT().SendMessage(mockOutgoingMessage).Return().Times(1),
	)

	redisCTRL := gomock.NewController(t)
	mockRediser := mocks.NewMockRediser(redisCTRL)

	mockGetResult := re.NewIntCmd("incr", testNumPlusses)

	gomock.InOrder(
		mockRediser.EXPECT().Decr(redis.GetKey("USER")).Return(mockGetResult).Times(1),
	)

	redis.TestingRDB = mockRediser

	mockEV := &slack.MessageEvent{
		Msg: slack.Msg{
			Type:    "message",
			Text:    testMessage,
			Channel: testChannelID,
		},
	}

	err := plusplus.Process(mockSlacker, mockEV)

	assert.Nil(err)

	redis.TestingRDB = nil
}
