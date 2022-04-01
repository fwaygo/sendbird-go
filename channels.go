package sendbird

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fwaygo/sendbird-go/api"
)

func (c *client) ChannelsCreate(ctx context.Context, request api.ChannelCreateRequest) (*api.ChannelResponse, error) {
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

	var response api.ChannelResponse
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

func (c *client) ChannelsView(ctx context.Context, request api.ChannelGetRequest) (*api.ChannelResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, c.url()+"/group_channels/"+request.ChannelUrl+strconv.FormatBool(*request.ShowMember), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.ChannelGetResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) ChannelsAddMember(ctx context.Context, request api.AddMemberRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, c.url()+"/group_channels/"+request.ChannelUrl+"/join", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	_, err = c.do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) ChannelsHide(ctx context.Context, request api.ChannelHideRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPut,
		fmt.Sprintf(
			"%s/group_channels/%s/hide",
			c.url(),
			request.ChannelUrl,
		),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	_, err = c.do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) ChannelsUnhide(ctx context.Context, request api.ChannelUnhideRequest) error {
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf(
			"%s/group_channels/%s/hide?user_id=%s",
			c.url(),
			request.ChannelUrl,
			request.UserID,
		),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	_, err = c.do(req)
	if err != nil {
		return err
	}

	return nil
}
