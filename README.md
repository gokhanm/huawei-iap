# huawei-iap
![](https://img.shields.io/badge/golang-1.16+-blue.svg?style=flat)

huawei-iap verifies the purchase receipt via Huawei Store

Current API Documents:
https://developer.huawei.com/consumer/en/doc/development/HMSCore-References/api-summary-desc-0000001063744095

# Installation
```
go get github.com/gokhanm/huawei-iap
```

### Verify Subscription
```go
import (
    hms "github.com/gokhanm/huawei-iap"
)

func main() {
    client := hms.New()
	req := hms.IAPRequest{
		PurchaseToken:  token,
		SubscriptionID: subscriptionID,
		ClientID:       clientID,
		ClientSecret:   clientSecret,
	}
    
	purchaseData, err := client.VerifySubscription(context.TODO(), req)
}
```