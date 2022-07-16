package client

import (
	"encoding/json"
	"fmt"
	api2 "github.com/bpalermo/football-data-go/v4/api"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultBaseURL   = "https://api.football-data.org"
	defaultAuthToken = ""
	defaultTimeout   = 30 * time.Second
)

type ApiClient interface {
	do(method, endpoint string, params map[string]string, ret interface{}) (err error)
	GetAreas() (a *api2.Areas, err error)
	GetArea(id string) (a []*api2.Area, err error)
	GetCompetitions() (a *api2.Competitions, err error)
	GetCompetition() (a *api2.Competition, err error)
}

// Option is a functional option for configuring the API client
type Option func(*Client) error

// Client holds information necessary to make a request to your API
type Client struct {
	baseURL    string
	httpClient *http.Client
	authToken  string
}

// BaseURL allows overriding of API client baseURL for testing
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// AuthToken sets the authentication token
func AuthToken(authToken string) Option {
	return func(c *Client) error {
		c.authToken = authToken
		return nil
	}
}

// Timeout http client timeout
func Timeout(timeout time.Duration) Option {
	return func(c *Client) error {
		c.httpClient.Timeout = timeout
		return nil
	}
}

// New creates a new API client
func New(opts ...Option) (*Client, error) {
	client := &Client{
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		authToken: defaultAuthToken,
	}

	if err := client.parseOptions(opts...); err != nil {
		return nil, err
	}

	return client, nil
}

// parseOptions parses the supplied options functions and returns a configured
// *Client instance
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each function option and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) do(method, endpoint string, params map[string]string, ret interface{}) (err error) {
	baseURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	if c.authToken != "" {
		req.Header.Add("X-Auth-Token", c.authToken)
	}

	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer handleReadError(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, ret); err != nil {
		return err
	}

	return nil
}

func handleReadError(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		return
	}
}
