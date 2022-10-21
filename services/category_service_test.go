package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yogasab/go-unit-test-playground/entities"
	"github.com/yogasab/go-unit-test-playground/repositories"
	"testing"
)

var categoryRepository = &repositories.CategoryRepositoryMock{Mock: mock.Mock{}}
var newCategoryService = NewCategoryService(categoryRepository)

func TestCategoryService_GetNotFound(t *testing.T) {
	//	program mock
	categoryRepository.Mock.On("FindByID", "1").Return(nil)

	category, err := newCategoryService.GetCategory("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	newCategory := entities.Category{
		ID:   2,
		Name: "Smartphone",
	}
	categoryRepository.Mock.On("FindByID", "2").Return(newCategory)

	category, err := newCategoryService.GetCategory("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, newCategory.ID, category.ID)
	assert.Equal(t, newCategory.Name, category.Name)
}
