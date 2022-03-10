package sendbird

import (
	"context"
	"log"
	"testing"

	"github.com/fwaygo/fwaygo-kit/format"
	"github.com/fwaygo/sendbird-go/api"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SendBirdAppID    string `envconfig:"SENDBIRD_APPLICATION_ID" default:"appid"`
	SendBirdAPIToken string `envconfig:"SENDBIRD_API_TOKEN" default:"shfjdksahfjkdshajfkdsah"`
	SendBirdVersion  string `envconfig:"SENDBIRD_VERSION" default:"v3"`
}

func getConfig(t *testing.T) *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		t.Fatal(err)
	}

	return &cfg
}

func TestNewClient(t *testing.T) {
	config := getConfig(t)
	_, err := NewClient(ClientConfig{
		ApplicationID: config.SendBirdAppID,
		APIToken:      config.SendBirdAPIToken,
		Version:       config.SendBirdVersion,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateChannel(t *testing.T) {
	config := getConfig(t)
	client, err := NewClient(ClientConfig{
		ApplicationID: config.SendBirdAppID,
		APIToken:      config.SendBirdAPIToken,
		Version:       config.SendBirdVersion,
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	channelName := "iusrjFGvKc1HMHNMI9QHy0hGE3reAuq1:iusrMG8ADbdBCVOkr1mIpNnnMKDdwTz1"
	users := []string{"iusrjFGvKc1HMHNMI9QHy0hGE3reAuq1", "iusrMG8ADbdBCVOkr1mIpNnnMKDdwTz1"}

	channel, err := client.ChannelsCreate(ctx, api.ChannelCreateRequest{
		UserIDs:    users,
		Name:       &channelName,
		IsDistinct: true,
	})

	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Created Channel %s", channel.ChannelUrl)
}

func TestSendMessage(t *testing.T) {
	config := getConfig(t)
	client, err := NewClient(ClientConfig{
		ApplicationID: config.SendBirdAppID,
		APIToken:      config.SendBirdAPIToken,
		Version:       config.SendBirdVersion,
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	user1 := format.NewUserID()
	user2 := format.NewUserID()
	users := []string{string(user1), string(user2)}
	channelName := string(user1) + ":" + string(user2)

	for _, j := range users {
		user, err := client.UsersCreate(ctx, api.UsersCreateRequest{
			UserID:     j,
			Nickname:   j,
			ProfileURL: "",
		})
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("New User %s: %s", user.UserID, user.CreatedAt)
	}

	channel, err := client.ChannelsCreate(ctx, api.ChannelCreateRequest{
		UserIDs: users,
		Name:    &channelName,
	})
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Created Channel %s", channel.ChannelUrl)
	request := api.SendBaseMessageRequest{
		MessageType: "MESG",
	}

	message, err := client.SendUserMessage(ctx, api.SendUserMessageRequest{
		SendBaseMessageRequest: request,
		UserID:                 string(user1),
		Message:                "first test message",
	}, api.ChannelParams{
		ChannelType: "group_channels",
		ChannelUrl:  channel.ChannelUrl,
	})
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Message ID: %d", message.MessageID)
}
