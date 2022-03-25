package sendbird

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fwaygo/sendbird-go/api"
)

type Client interface {
	UsersCreate(context.Context, api.UsersCreateRequest) (*api.UsersCreateResponse, error)
	UsersGet(context.Context, api.UsersGetRequest) (*api.UsersGetResponse, error)
	UsersUpdate(context.Context, api.UsersUpdateRequest) (*api.UsersUpdateResponse, error)
	UsersDelete(context.Context, api.UsersDeleteRequest) error

	ChannelsCreate(context.Context, api.ChannelCreateRequest) (*api.ChannelCreateResponse, error)
	ChannelsList(context.Context, api.ChannelListRequest) (*api.ChannelListResponse, error)

	AddMemberToGroupChannel(context.Context, api.AddMemberRequest) (*api.AddMemberResponse, error)

	SendUserMessage(context.Context, api.SendUserMessageRequest, api.ChannelParams) (*api.UserMessage, error)
	SendFileMessage(context.Context, api.SendFileMessageRequest, api.ChannelParams) (*api.FileMessage, error)

	AddReaction(context.Context, api.AddReactionRequest, api.ChannelParams) (*api.ReactionUpdateResponse, error)
}

type ClientConfig struct {
	ApplicationID string `envconfig:"SENDBIRD_APPLICATION_ID"`
	APIToken      string `envconfig:"SENDBIRD_API_TOKEN"`
	Version       string `envconfig:"SENDBIRD_VERSION" default:"v3"`
}

type client struct {
	appID    string
	apiToken string
	version  string
}

func NewClient(cfg ClientConfig) (Client, error) {
	return &client{
		appID:    cfg.ApplicationID,
		apiToken: cfg.APIToken,
		version:  cfg.Version,
	}, nil
}

func (c *client) url() string {
	return fmt.Sprintf("https://api-%s.sendbird.com/%s", c.appID, c.version)
}

func (c *client) do(request *http.Request) ([]byte, error) {
	request.Header.Add("Content-Type", "application/json; charset=utf8")
	request.Header.Add("Api-Token", c.apiToken)
	log.Printf(c.apiToken)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[%d] %s", response.StatusCode, resp)
	}

	return resp, nil
}
