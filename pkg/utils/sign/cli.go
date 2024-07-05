package sign

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// Sign signs a message using the provided private key.
func Sign(message string, privateKey *ecdsa.PrivateKey) (string, error) {

	// Hash the message
	hash := sha256.Sum256([]byte(message))

	// Sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	// Encode the signature
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	// Encode the signature to hex
	signatureHex := hex.EncodeToString(signature)
	return signatureHex, nil
}
