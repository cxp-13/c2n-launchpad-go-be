package service

import (
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/models"
	"github.com/cxp-13/c2n-launchpad-go-be/internal/portalapi/repositories"
)

type ProjectService struct {
	repo *repositories.ProjectRepository
}

func NewProjectService(repo *repositories.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) GetByID(productID uint) (*models.Project, error) {
	return s.repo.GetById(productID)
}

func (s *ProjectService) List() ([]models.Project, error) {
	return s.repo.List()
}

func (s *ProjectService) UpdateByProject(project *models.Project) error {
	return s.repo.Update(project)
}
