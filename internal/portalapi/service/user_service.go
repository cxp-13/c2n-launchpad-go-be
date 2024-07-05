package service

import (
	"crypto/ecdsa"
	"github.com/cxp-13/c2n-launchpad-go-be/pkg/utils/sign"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// UserService defines the interface for the signing service.
type UserService struct {
	privateKey *ecdsa.PrivateKey // 嵌入私钥
}

// NewUserService creates a new UserService with the given private key.
func NewUserService(privateKey string) *UserService {
	// 将十六进制私钥转换为 *ecdsa.PrivateKey
	privateKeyBytes, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("HexToECDSA fail: %s", err.Error())
	}
	return &UserService{
		privateKey: privateKeyBytes,
	}
}

// Sign signs a message using the provided private key.
func (s *UserService) Sign(message string) (string, error) {
	// 使用私钥对消息进行签名
	return sign.Sign(message, s.privateKey)
}
