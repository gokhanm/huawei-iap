package huaweistore

import (
	"context"
	"encoding/json"
	"fmt"
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
func (c *Client) VerifyToken(ctx context.Context, reqBody IAPRequest) (*InappPurchaseData, error) {
	url := c.OrderSiteURL + "/applications/purchases/tokens/verify"

	body, err := c.sendRequest(ctx, url, reqBody)
	if err != nil {
		return nil, err
	}

	var orderResponse IAPOrderResponse
	if err := json.Unmarshal(body, &orderResponse); err != nil {
		return nil, err
	}

	if orderResponse.ResponseCode != IAPResponseResultOK {
		//TODO parse api result codes
		// https://developer.huawei.com/consumer/en/doc/development/HMS-References/iap-api-specification-related-v4#API-ErrorCode

		return nil, fmt.Errorf("response code: %s, response message: %s", orderResponse.ResponseCode, orderResponse.ResponseMessage)
	}

	var purchaseData InappPurchaseData
	if err := json.Unmarshal([]byte(orderResponse.PurchaseTokenData), &purchaseData); err != nil {
		return nil, err
	}

	return &purchaseData, nil
}
