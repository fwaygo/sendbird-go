package sendbird

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fwaygo/sendbird-go/api"
)

func (c *client) SearchMessage(ctx context.Context, request api.SearchMessageRequest) (*api.Message, error) {
	req, err := http.NewRequest(http.MethodGet,
		c.url()+fmt.Sprintf("/%s/%s/messages/%s",
			request.ChannelType,
			request.ChannelURL,
			strconv.FormatUint(request.MessageID, 10)),
		nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.Message
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) SendUserMessage(ctx context.Context, request api.SendUserMessageRequest, channel api.ChannelParams) (*api.UserMessage, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url()+"/"+string(channel.ChannelType)+"/"+string(channel.ChannelUrl)+"/messages", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.UserMessage
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) SendFileMessage(ctx context.Context, request api.SendFileMessageRequest, channel api.ChannelParams) (*api.FileMessage, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url()+"/"+string(channel.ChannelType)+"/"+string(channel.ChannelUrl)+"/messages", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.FileMessage
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) DeleteMessage(ctx context.Context, request api.DeleteMessageRequest) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		c.url()+fmt.Sprintf("/%s/%s/messages/%s",
			request.ChannelType,
			request.ChannelURL,
			strconv.FormatUint(request.MessageID, 10)),
		nil)
	if err != nil {
		return err
	}

	_, err = c.do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) AddReaction(ctx context.Context, request api.AddReactionRequest, channel api.ChannelParams) (*api.ReactionUpdateResponse, error) {
	if channel.MessageID == nil {
		return nil, fmt.Errorf("message id is nil")
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost,
		c.url()+
			"/"+string(channel.ChannelType)+
			"/"+string(channel.ChannelUrl)+
			"/messages/"+
			strconv.FormatUint(*channel.MessageID, 10)+
			"/reactions", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.ReactionUpdateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) RemoveReaction(ctx context.Context, request api.RemoveReactionRequest, channel api.ChannelParams) (*api.ReactionUpdateResponse, error) {
	if channel.MessageID == nil {
		return nil, fmt.Errorf("message id is nil")
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodDelete,
		c.url()+
			"/"+string(channel.ChannelType)+
			"/"+string(channel.ChannelUrl)+
			"/messages/"+
			strconv.FormatUint(*channel.MessageID, 10)+
			"/reactions"+
			"?user_id="+request.UserID+
			"&reaction="+request.Reaction,
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.ReactionUpdateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
