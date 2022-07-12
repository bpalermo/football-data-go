package client

import (
	"fmt"
	"github.com/bpalermo/football-data-go/v4/api"
	"net/http"
)

const (
	competitionsEndpoint = "/v4/competitions"
)

func (c *Client) GetCompetitions() (a *api.Competitions, err error) {
	a = &api.Competitions{}
	err = c.do(http.MethodGet, competitionsEndpoint, nil, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Client) GetCompetition(code string) (competition *api.Competition, err error) {
	competition = &api.Competition{}
	err = c.do(http.MethodGet, fmt.Sprintf("%s/%s", competitionsEndpoint, code), nil, competition)
	if err != nil {
		return nil, err
	}
	return competition, nil
}
