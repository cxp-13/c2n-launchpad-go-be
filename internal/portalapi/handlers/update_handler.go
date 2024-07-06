package handlers

import (
	"net/http"

	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/models"
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/service"
	"github.com/gin-gonic/gin"
)

type UpdateHandler struct {
	projectService *service.ProjectService
}

func NewUpdateHandler(projectService *service.ProjectService) *UpdateHandler {
	return &UpdateHandler{
		projectService: projectService,
	}
}

func (h *UpdateHandler) Update(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.projectService.UpdateByProject(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}
