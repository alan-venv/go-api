package repositories

import (
	"example/go-api/src/models"
)

type IUserRepository interface {
	ReadAll() ([]models.User, error)
	Read(string) (models.User, error)
	Create(models.User) error
	Delete(string) error
}
