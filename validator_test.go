package huaweistore

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	expected := &Client{
		OrderSiteURL:        OrderSiteGermany,
		SubscriptionSiteURL: SubscriptionSiteGermany,
		TokenURL:            ProductTokenURL,
		httpCli: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	actual := New()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestNewWithClient(t *testing.T) {
	expected := &Client{
		OrderSiteURL:        OrderSiteGermany,
		SubscriptionSiteURL: SubscriptionSiteGermany,
		TokenURL:            ProductTokenURL,
		httpCli: &http.Client{
			Timeout: 20 * time.Second,
		},
	}

	actual := NewWithClient(&http.Client{
		Timeout: 20 * time.Second,
	})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestVerifySignature(t *testing.T) {
	var testCase = []struct {
		Name      string
		Content   string
		PublicKey string
		Signature string
		Error     error
	}{
		{
			Name:      "verification error",
			Content:   "messagetobesigned",
			PublicKey: `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5dXHzCplyOMtHV2Wsv6aJYJbiWUEdI6MfKND44DkHFbyz90Jfe3ntJOhWCFiT0Zbm/pUFSUocnExbebOnT/Z5DyCDz8d6Ek1o2Z8KOeSWf00e756v7D3R6GpLfOw03yKcxZWjYgKdb5nDaC/qa4wiz+TaJy625DzTjyZ7hvq5owIDAQAB`,
			Signature: `ORpOhMMyKmbgBx3GaLtOZkGStkT/dSkrPTyEagDkLSVOAf+NpsDkiVOEHQ7/8Xsbspx37IzNSjL00YInCEIaniSzGU97szv8u1pvoLAPpdRBR4djZbNDdeOKUJ8Zi5LodtU+/V8RdjidU7xhCdOMINmaaX1rYjTfq7RBHDI/zCo=`,
			Error:     errors.New("crypto/rsa: verification error"),
		},
		{
			Name:      "no error",
			Content:   "hello world",
			PublicKey: `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC5dXHzCplyOMtHV2Wsv6aJYJbiWUEdI6MfKND44DkHFbyz90Jfe3ntJOhWCFiT0Zbm/pUFSUocnExbebOnT/Z5DyCDz8d6Ek1o2Z8KOeSWf00e756v7D3R6GpLfOw03yKcxZWjYgKdb5nDaC/qa4wiz+TaJy625DzTjyZ7hvq5owIDAQAB`,
			Signature: `a9Wr8F3fQlcy3lq5jlEwLpAV/Gyj1aadE+1tUAHmP7NK0f+/ZlueJOD8wfFBlBGNjnSKsZbsnBEHvuaaMr42UtFn/zCmYavX2yyZAZLmW0merrQO+2BlHKz7ipZjo+TEkWevY0o8mta3So79caK/O1YdKwOFh2RMiUIPrHZm9L0=`,
			Error:     nil,
		},
	}

	for _, test := range testCase {
		t.Run(fmt.Sprintf("test name %s", test.Name), func(t *testing.T) {
			if err := VerifySignature(test.Content, test.Signature, test.PublicKey); err != nil {
				if err == test.Error {
					t.Errorf("Expected error for input")
				}
			}
		})
	}
}
