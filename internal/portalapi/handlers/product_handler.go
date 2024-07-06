package handlers

import (
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service *service.ProjectService
}

func NewProductHandler(service *service.ProjectService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) BaseInfo(c *gin.Context) {
	productIDStr := c.Query("productId")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	product, err := h.service.GetByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductHandler) List(c *gin.Context) {
	projects, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}
