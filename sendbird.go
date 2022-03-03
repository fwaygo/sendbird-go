package sendbird

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fwaygo/sendbird-go/api"
)

type Client interface {
	UsersCreate(context.Context, api.UsersCreateRequest) (*api.UsersCreateResponse, error)
	UsersGet(context.Context, api.UsersGetRequest) (*api.UsersGetResponse, error)
	UsersUpdate(context.Context, api.UsersUpdateRequest) (*api.UsersUpdateResponse, error)
	UsersDelete(context.Context, api.UsersDeleteRequest) error

	ChannelsCreate(context.Context, api.ChannelCreateRequest) (*api.ChannelCreateResponse, error)

	SendUserMessage(context.Context, api.SendUserMessageRequest, api.ChannelParams) (*api.UserMessage, error)
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

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}
	defer response.Body.Close()

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
