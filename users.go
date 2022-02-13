package sendbird

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/fwaygo/sendbird-go/api"
)

func (c *client) UsersCreate(ctx context.Context, request api.UsersCreateRequest) (*api.UsersCreateResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url()+"/users", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.UsersCreateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) UsersGet(ctx context.Context, request api.UsersGetRequest) (*api.UsersGetResponse, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, c.url()+"/users", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var response api.UsersCreateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *client) UsersUpdate(ctx context.Context, request api.UsersUpdateRequest) (*api.UsersUpdateResponse, error) {
	return nil, nil
}

func (c *client) UsersDelete(ctx context.Context, request api.UsersDeleteRequest) error {
	return nil
}
