package repositories

import (
	"example/go-api/src/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	Database *gorm.DB
}

func (self UserGormRepository) ReadAll() (*[]models.User, error) {
	var users []models.User

	err := self.Database.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (self UserGormRepository) Read(id string) (*models.User, error) {
	var user models.User

	err := self.Database.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (self UserGormRepository) Create(user models.User) error {
	user.ID = uuid.New().String()

	err := self.Database.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (self UserGormRepository) Update(user models.User) error {
	var saved_user models.User

	err := self.Database.First(&saved_user, "id = ?", user.ID).Error
	if err != nil {
		return err
	}

	saved_user.Name = user.Name
	saved_user.Email = user.Email
	saved_user.Password = user.Password

	err = self.Database.Save(&saved_user).Error
	if err != nil {
		return err
	}
	return nil
}

func (self UserGormRepository) Delete(id string) error {
	var user models.User

	err := self.Database.First(&user, "id = ?", id).Error
	if err != nil {
		return err
	}

	err = self.Database.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
