package handlers

import (
	"fmt"
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/service"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"net/http"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		service: userService,
	}
}

func (e *UserHandler) Registered(c *gin.Context) {
	userAddress := c.PostForm("userAddress")
	contractAddress := c.PostForm("contractAddress")

	if userAddress == "" || contractAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	cleanUserAddr := common.HexToAddress(userAddress).String()
	cleanContractAddr := common.HexToAddress(contractAddress).String()

	message := strings.ToLower(cleanUserAddr + cleanContractAddr)

	signature, err := e.service.Sign(message)
	if err != nil {
		log.Printf("Error signing registration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"signature": signature})
}

func (e *UserHandler) Participation(c *gin.Context) {
	userAddress := c.PostForm("userAddress")
	amount := c.PostForm("amount")
	contractAddress := c.PostForm("contractAddress")

	if userAddress == "" || amount == "" || contractAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	cleanUserAddress := common.HexToAddress(userAddress).String()
	cleanContractAddress := common.HexToAddress(contractAddress).String()
	hexString := common.Bytes2Hex([]byte(fmt.Sprintf("%s%s%s", cleanUserAddress, amountInt.String(), cleanContractAddress)))

	signature, err := e.service.Sign(hexString)
	if err != nil {
		log.Printf("Error signing participation: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	c.JSON(http.StatusOK, gin.H{"signature": signature})
}
