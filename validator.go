package huaweistore

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
)

// VerifySignature verifying the signature in the returned result.
// Using SHA256withRSA RSA Signature Algorithm
func VerifySignature(jsonContent string, signature string, publicKey string) error {
	pKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return err
	}

	pub, err := x509.ParsePKIXPublicKey(pKey)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256([]byte(jsonContent))
	sgn, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hashed[:], sgn)
}
