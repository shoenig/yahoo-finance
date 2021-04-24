package yahoofinance

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

const (
	yf = "https://query1.finance.yahoo.com"

	// do we need this?
	otherURL = "https://finance.yahoo.com/quote"
)

type Client interface {
	Lookup(ticker string) (*Chart, error)
}

type client struct {
	httpClient *http.Client
}

type Options struct {
	HTTPClient *http.Client
}

func New(opts *Options) Client {
	if opts == nil {
		opts = new(Options)
	}

	httpClient := opts.HTTPClient
	if httpClient == nil {
		httpClient = cleanhttp.DefaultPooledClient()
	}

	return &client{
		httpClient: httpClient,
	}
}

func (c *client) Lookup(ticker string) (*Chart, error) {
	yURL := fmt.Sprintf("%s/v8/finance/chart/%s", yf, ticker)

	request, reqErr := http.NewRequest(http.MethodGet, yURL, nil)
	if reqErr != nil {
		return nil, reqErr
	}

	response, reqErr := c.httpClient.Do(request)
	if reqErr != nil {
		return nil, reqErr
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("received status code %d", response.StatusCode)
	}

	var chart Chart
	if decErr := json.NewDecoder(response.Body).Decode(&chart); decErr != nil {
		return nil, decErr
	}

	return &chart, nil
}
