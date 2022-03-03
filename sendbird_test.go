package sendbird

import (
	"context"
	"log"
	"testing"

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

	var ctx = context.Background()
	channelName := "testChannelName"
	users := []string{"test1user", "test2user"}

	channel, err := client.ChannelsCreate(ctx, api.ChannelCreateRequest{
		UserIDs: users,
		Name:    &channelName,
	})
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Created Channel %s", channel.ChannelUrl)
}
