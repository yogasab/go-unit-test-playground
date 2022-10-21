package repositories

import "github.com/yogasab/go-unit-test-playground/entities"

type CategoryRepository interface {
	FindByID(ID string) *entities.Category
}
