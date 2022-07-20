package todoist

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//const DefaultRestUrl string = "https://api.foxpass.com/v1"

var ErrInvalidAuthn = errors.New("credentials not valid")
var ErrInvalidAuthz = errors.New("credentials not authorized to access resource")
var ErrNotFound = errors.New("requested resource not found")

type Client struct {
	httpClient *http.Client
	apiKey     string
	host       string
	ctx        context.Context
}

func NewClient(apiKey string, host string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}

	return &Client{
		httpClient: http.DefaultClient,
		apiKey:     apiKey,
		host:       host,
	}
}

func (c *Client) doRequest(ctx context.Context, method, endpoint string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.host, endpoint), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, ErrInvalidAuthn
	} else if resp.StatusCode == 403 {
		return nil, ErrInvalidAuthz
	} else if resp.StatusCode == 404 {
		return nil, ErrNotFound
	} else if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}
