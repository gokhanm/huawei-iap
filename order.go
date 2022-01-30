package huaweistore

import (
	"context"
	"encoding/json"
)

const (
	// OrderSiteChina site Information for China
	OrderSiteChina = "https://orders-drcn.iap.hicloud.com"
	// OrderSiteGermany site Information for China
	OrderSiteGermany = "https://orders-dre.iap.hicloud.com"
	// OrderSiteSingapore site Information for China
	OrderSiteSingapore = "https://orders-dra.iap.hicloud.com"
	// OrderSiteRussia site Information for China
	OrderSiteRussia = "https://orders-drru.iap.hicloud.com"
	// OrderSiteAppTouch site Information for China
	OrderSiteAppTouch = "https://orders-at-dre.iap.dbankcloud.com"
)

// VerifyToken purchase Token Verification for the Order Service
func (c *Client) VerifyToken(ctx context.Context, reqBody IAPRequest) (*IAPResponse, error) {
	url := c.OrderSiteURL + "/applications/purchases/tokens/verify"

	body, err := c.sendRequest(ctx, url, reqBody)
	if err != nil {
		return nil, err
	}

	var result IAPResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
