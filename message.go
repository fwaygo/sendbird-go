package sendbird

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/fwaygo/sendbird-go/api"
)

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
