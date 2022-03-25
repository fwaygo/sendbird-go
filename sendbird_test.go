package sendbird

import (
	"context"
	"fmt"
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

// func TestNewClient(t *testing.T) {
// 	config := getConfig(t)
// 	_, err := NewClient(ClientConfig{
// 		ApplicationID: config.SendBirdAppID,
// 		APIToken:      config.SendBirdAPIToken,
// 		Version:       config.SendBirdVersion,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestCreateChannel(t *testing.T) {
// 	config := getConfig(t)
// 	client, err := NewClient(ClientConfig{
// 		ApplicationID: config.SendBirdAppID,
// 		APIToken:      config.SendBirdAPIToken,
// 		Version:       config.SendBirdVersion,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ctx := context.Background()
// 	channelName := "iusrjFGvKc1HMHNMI9QHy0hGE3reAuq1:iusrMG8ADbdBCVOkr1mIpNnnMKDdwTz1"
// 	users := []string{"iusrjFGvKc1HMHNMI9QHy0hGE3reAuq1", "iusrMG8ADbdBCVOkr1mIpNnnMKDdwTz1"}

// 	channel, err := client.ChannelsCreate(ctx, api.ChannelCreateRequest{
// 		UserIDs:    users,
// 		Name:       &channelName,
// 		IsDistinct: true,
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	log.Printf("Created Channel %s", channel.ChannelUrl)
// }

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

func TestReactions(t *testing.T) {
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

	reaction, err := client.AddReaction(ctx, api.AddReactionRequest{
		UserID:   string(user1),
		Reaction: "celebrate",
	}, api.ChannelParams{
		ChannelType: api.GROUP_CHANNEL,
		ChannelUrl:  channel.ChannelUrl,
		MessageID:   &message.MessageID,
	})
	if err != nil {
		t.Fatalf("could not add reaction -- %s", err)
	}

	log.Printf("Reaction: %v", reaction)
}

func TestListChannels(t *testing.T) {
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
	user3 := format.NewUserID()

	users1 := []string{user1.String(), user2.String()}
	users2 := []string{user1.String(), user3.String()}
	users3 := []string{user1.String(), user2.String(), user3.String()}

	usersArr := [][]string{
		users1,
		users2,
		users3,
	}
	channelName := user1.String() + ":" + user2.String()

	for _, j := range users3 {
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

	for _, users := range usersArr {
		_, err := client.ChannelsCreate(ctx, api.ChannelCreateRequest{
			UserIDs: users,
			Name:    &channelName,
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, users := range usersArr {

		channelList, err := client.ChannelsList(ctx, api.ChannelListRequest{
			ShowEmpty:           true,
			ShowMember:          false,
			ShowDeliveryReceipt: false,
			ShowReadReceipt:     false,
			ShowMetadata:        false,
			ShowFrozen:          false,
			MembersExactlyIn:    users,
		})

		for _, channel := range channelList.Channels {
			fmt.Printf("Channel id: %s\n", channel.ChannelUrl)
		}

		if err != nil {
			t.Fatal(err)
		}
		if len(channelList.Channels) != 1 {
			t.Fatal("Query incorrect: request responded with multiple channels")
		}
	}
	log.Println("Query Pass")
}

func TestEncoder(t *testing.T) {
	user1 := format.NewUserID()
	user2 := format.NewUserID()
	user3 := format.NewUserID()
	users := []string{user1.String(), user2.String(), user3.String()}

	EncodeParameters(&api.ChannelListRequest{
		ShowEmpty:           true,
		ShowMember:          false,
		ShowDeliveryReceipt: false,
		ShowReadReceipt:     false,
		ShowMetadata:        false,
		ShowFrozen:          false,
		MembersExactlyIn:    users,
	})
}

// func TestAddMember(t *testing.T) {
// 	config := getConfig(t)
// 	client, err := NewClient(ClientConfig{
// 		ApplicationID: config.SendBirdAppID,
// 		APIToken:      config.SendBirdAPIToken,
// 		Version:       config.SendBirdVersion,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ctx := context.Background()

// 	user1 := format.NewUserID()
// 	user2 := format.NewUserID()
// 	user3 := format.NewUserID()
// 	users := []string{user1.String(), user2.String(), user3.String()}

// 	for _, user := range users {
// 		_, err := client.UsersCreate(ctx, api.UsersCreateRequest{
// 			UserID:     user,
// 			Nickname:   user,
// 			ProfileURL: "",
// 		})
// 		if err != nil {
// 			t.Fatalf("error creating user: %s", err)
// 		}
// 	}

// 	memberToAdd := format.NewUserID()

// 	_, err = client.UsersCreate(ctx, api.UsersCreateRequest{
// 		UserID: memberToAdd.String(),
// 	})

// 	if err != nil {
// 		t.Fatalf("Error creating member to add: %s", err)
// 	}

// 	channelName := "addingMembers:" + memberToAdd.String()
// 	isPublicPtr := true

// 	chResponse, err := client.ChannelsCreate(ctx, api.ChannelCreateRequest{
// 		UserIDs:  users,
// 		Name:     &channelName,
// 		IsPublic: &isPublicPtr,
// 	})

// 	if err != nil {
// 		t.Fatal("failed to create channel")
// 	}

// 	_, err = client.AddMemberToGroupChannel(ctx, api.AddMemberRequest{
// 		ChannelUrl: chResponse.ChannelUrl,
// 		UserId:     memberToAdd.String(),
// 	})

// 	if err != nil {
// 		t.Fatalf("failed to add member: %s", err)
// 	}
// }
