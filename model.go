package huaweistore

const (
	// https://developer.huawei.com/consumer/en/doc/development/HMS-References/iap-api-specification-related-v4#API-ErrorCode
	IAPResponseResultOK                  = "0"
	IAPResponseResultUSER_CANCELED       = "1"
	IAPResponseResultSERVICE_UNAVAILABLE = "2"
	IAPResponseResultBILLING_UNAVAILABLE = "3"
	IAPResponseResultITEM_UNAVAILABLE    = "4"
	IAPResponseResultDEVELOPER_ERROR     = "5"
	IAPResponseResultERROR               = "6"
	IAPResponseResultITEM_ALREADY_OWNED  = "7"
	IAPResponseResultITEM_NOT_OWNED      = "8"
	IAPResponseResultITEM_CONSUMED       = "9"
)

// IAPRequest input Parameters for order and subscription
type IAPRequest struct {
	PurchaseToken        string `json:"purchaseToken,omitempty"`
	ProductID            string `json:"productId,omitempty"`
	ClientID             string
	ClientSecret         string
	ApplicationPublicKey string
	SubscriptionID       string `json:"subscriptionId,omitempty"`
}

// TokenResponse is
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

// IAPSubscriptionResponse store subscription response
type IAPSubscriptionResponse struct {
	ResponseCode      string `json:"responseCode,omitempty"`
	ResponseMessage   string `json:"responseMessage,omitempty"`
	InappPurchaseData string `json:"inappPurchaseData,omitempty"`
}

// IAPSubscriptionResponse store subscription response
type IAPOrderResponse struct {
	ResponseCode      string `json:"responseCode,omitempty"`
	ResponseMessage   string `json:"responseMessage,omitempty"`
	PurchaseTokenData string `json:"purchaseTokenData,omitempty"`
	DataSignature     string `json:"dataSignature,omitempty"`
}

// https://developer.huawei.com/consumer/en/doc/development/HMS-References/iap-InAppPurchaseDetails-v4
type InappPurchaseData struct {
	SurveyReason           *int   `json:"surveyReason,omitempty"`
	SurveyDetails          string `json:"surveyDetails,omitempty"`
	ApplicationID          int    `json:"applicationId,omitempty"`
	AutoRenewing           bool   `json:"autoRenewing,omitempty"`
	OrderID                string `json:"orderId,omitempty"`
	Kind                   int    `json:"kind,omitempty"`
	PackageName            string `json:"packageName,omitempty"`
	ProductID              string `json:"productId,omitempty"`
	ProductName            string `json:"productName,omitempty"`
	PurchaseTime           int64  `json:"purchaseTime,omitempty"`
	PurchaseState          int    `json:"purchaseState,omitempty"`
	DeveloperPayload       string `json:"developerPayload,omitempty"`
	DeveloperChallenge     string `json:"developerChallenge,omitempty"`
	ConsumptionState       *int   `json:"consumptionState,omitempty"`
	PurchaseToken          string `json:"purchaseToken,omitempty"`
	PurchaseType           *int   `json:"purchaseType,omitempty"`
	Currency               string `json:"currency,omitempty"`
	Price                  int    `json:"price,omitempty"`
	Country                string `json:"country,omitempty"`
	PayType                string `json:"payType,omitempty"`
	PayOrderID             string `json:"payOrderId,omitempty"`
	AccountFlag            *int   `json:"accountFlag,omitempty"`
	LastOrderID            string `json:"lastOrderId,omitempty"`
	ProductGroup           string `json:"productGroup,omitempty"`
	OriginalPurchaseTime   int    `json:"oriPurchaseTime,omitempty"`
	SubscriptionID         string `json:"subscriptionId,omitempty"`
	OriginalSubscriptionID string `json:"oriSubscriptionId,omitempty"`
	Quantity               int    `json:"quantity,omitempty"`
	DaysLasted             int    `json:"daysLasted,omitempty"`
	NumOfPeriods           *int   `json:"numOfPeriods,omitempty"`
	NumOfDiscount          int    `json:"numOfDiscount,omitempty"`
	ExpirationDate         int64  `json:"expirationDate,omitempty"`
	ExpirationIntent       int    `json:"expirationIntent,omitempty"`
	RetryFlag              *int   `json:"retryFlag,omitempty"`
	IntroductoryFlag       *int   `json:"introductoryFlag,omitempty"`
	TrialFlag              *int   `json:"trialFlag,omitempty"`
	CancelTime             int    `json:"cancelTime,omitempty"`
	CancelReason           *int   `json:"cancelReason,omitempty"`
	AppInfo                string `json:"appInfo,omitempty"`
	NotifyClosed           *int   `json:"notifyClosed,omitempty"`
	RenewStatus            *int   `json:"renewStatus,omitempty"`
	PriceConsentStatus     *int   `json:"priceConsentStatus,omitempty"`
	RenewPrice             int    `json:"renewPrice,omitempty"`
	SubscriptionIsvalid    bool   `json:"subIsvalid,omitempty"`
	DeferFlag              int    `json:"deferFlag,omitempty"`
	CancelWay              *int   `json:"cancelWay,omitempty"`
	CancellationTime       int    `json:"cancellationTime,omitempty"`
	CancelledSubKeepDays   int    `json:"cancelledSubKeepDays,omitempty"`
	Confirmed              *int   `json:"confirmed,omitempty"`
	ResumeTime             int    `json:"resumeTime,omitempty"`
	SdkChannel             string `json:"sdkChannel,omitempty"`
}
