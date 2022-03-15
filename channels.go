package sendbird

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/fwaygo/sendbird-go/api"
)

func (c *client) ChannelsCreate(ctx context.Context, request api.ChannelCreateRequest) (*api.ChannelCreateResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	log.Printf("%s", body)

	req, err := http.NewRequest(http.MethodPost, c.url()+"/group_channels", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.ChannelCreateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) ChannelsList(ctx context.Context, request api.ChannelListRequest) (*api.ChannelListResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	log.Printf("%s", body)

	requestPtr := &request
	parameters := EncodeParameters(requestPtr)

	req, err := http.NewRequest(http.MethodGet, c.url()+"/group_channels"+parameters, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.ChannelListResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
