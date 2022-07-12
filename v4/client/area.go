package client

import (
	"fmt"
	"github.com/bpalermo/football-data-go/v4/api"
	"net/http"
)

const (
	areasEndpoint = "/v4/areas"
)

func (c *Client) GetAreas() (a *api.Areas, err error) {
	a = &api.Areas{}
	err = c.do(http.MethodGet, areasEndpoint, nil, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Client) GetArea(id string) (a *api.Area, err error) {
	a = &api.Area{}
	err = c.do(http.MethodGet, fmt.Sprintf("%s/%s", areasEndpoint, id), nil, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
