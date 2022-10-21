package services

import (
	"errors"
	"github.com/yogasab/go-unit-test-playground/entities"
	"github.com/yogasab/go-unit-test-playground/repositories"
)

type categoryService struct {
	CategoryRepository repositories.CategoryRepository
}

type CategoryService interface {
	GetCategory(ID string) (*entities.Category, error)
}

func NewCategoryService(CategoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{
		CategoryRepository: CategoryRepository,
	}
}

func (s *categoryService) GetCategory(ID string) (*entities.Category, error) {
	category := s.CategoryRepository.FindByID(ID)
	if category == nil {
		return category, errors.New("category is not found")
	}
	return category, nil
}
