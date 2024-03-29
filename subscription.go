package huaweistore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// SubscriptionSiteChina site Information for China
	SubscriptionSiteChina = "https://subscr-drcn.iap.hicloud.com"
	// SubscriptionSiteGermany site Information for China
	SubscriptionSiteGermany = "https://subscr-dre.iap.hicloud.com"
	// SubscriptionSiteSingapore site Information for China
	SubscriptionSiteSingapore = "https://subscr-dra.iap.hicloud.com"
	// SubscriptionSiteRussia site Information for China
	SubscriptionSiteRussia = "https://subscr-drru.iap.hicloud.com"
	// SubscriptionSiteAppTouchURL site Information for China
	SubscriptionSiteAppTouchURL = "https://subscr-at-dre.iap.dbankcloud.com"
)

// VerifySubscription purchase Token Verification for the Subscription Service
func (c *Client) VerifySubscription(ctx context.Context, reqBody IAPRequest) (*InappPurchaseData, error) {
	url := c.SubscriptionSiteURL + "/sub/applications/v2/purchases/get"

	body, err := c.sendRequest(ctx, url, reqBody)
	if err != nil {
		return nil, err
	}

	var subsResponse IAPSubscriptionResponse
	if err := json.Unmarshal(body, &subsResponse); err != nil {
		return nil, err
	}

	if subsResponse.ResponseCode != IAPResponseResultOK {
		//TODO parse api result codes
		// https://developer.huawei.com/consumer/en/doc/development/HMS-References/iap-api-specification-related-v4#API-ErrorCode

		return nil, fmt.Errorf("response code: %s, response message: %s", subsResponse.ResponseCode, subsResponse.ResponseMessage)
	}

	var purchaseData InappPurchaseData
	if err := json.Unmarshal([]byte(subsResponse.InappPurchaseData), &purchaseData); err != nil {
		return nil, err
	}

	return &purchaseData, nil
}

// sendRequest sends receipts and gets validation result
func (c *Client) sendRequest(ctx context.Context, url string, reqBody IAPRequest) ([]byte, error) {
	authHeader, err := c.GetHeaderToken(reqBody)
	if err != nil {
		return nil, err
	}

	encodeBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(encodeBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", authHeader)
	req = req.WithContext(ctx)

	resp, err := c.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("received http status code %d from the store server", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
