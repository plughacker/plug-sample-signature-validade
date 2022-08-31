package main

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
)

// Verify is a example of a function that validates the payload signature
func Verify(publicKeyRaw string, payload string, signatureTime int64, sigHex string) (bool, error) {

	sig, err := hex.DecodeString(sigHex)

	if err != nil {
		return false, err
	}

	publicKey, err := getEd25519PublicKey(publicKeyRaw)

	if err != nil {
		return false, err
	}

	body := fmt.Sprintf("%d\n%s", signatureTime, payload)
	message := []byte(body)

	isValid := ed25519.Verify(publicKey, message, sig)

	return isValid, nil
}

func getEd25519PublicKey(pubPEM string) (ed25519.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse DER encoded public key: " + err.Error())
	}

	switch pub := pub.(type) {
	case ed25519.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("unknown type of public key")
	}
}
