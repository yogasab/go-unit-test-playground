package repositories

import (
	"github.com/stretchr/testify/mock"
	"github.com/yogasab/go-unit-test-playground/entities"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindByID(ID string) *entities.Category {
	args := repository.Mock.Called(ID)
	if args.Get(0) == nil {
		return nil
	}
	category := args.Get(0).(entities.Category)
	return &category
}
