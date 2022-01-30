package huaweistore

import (
	"net/http"
	"time"
)

// Client implements IAPClient
type Client struct {
	OrderSiteURL        string
	SubscriptionSiteURL string
	TokenURL            string
	httpCli             *http.Client
}

// New creates a client object
func New() *Client {
	return &Client{
		OrderSiteURL:        OrderSiteGermany,
		SubscriptionSiteURL: SubscriptionSiteGermany,
		TokenURL:            ProductTokenURL,
		httpCli: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// NewWithClient creates a client with a custom http client.
func NewWithClient(hc *http.Client) *Client {
	return &Client{
		OrderSiteURL:        OrderSiteGermany,
		SubscriptionSiteURL: SubscriptionSiteGermany,
		TokenURL:            ProductTokenURL,
		httpCli:             hc,
	}
}
