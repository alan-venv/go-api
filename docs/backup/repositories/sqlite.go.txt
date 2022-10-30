package repository

import (
	"example/go-api/src/database"
	"example/go-api/src/models"
)

func ReadUsers() ([]models.User, error) {
	db := database.Get()
	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func ReadUser(id int) (models.User, error) {
	db := database.Get()
	var user models.User

	err := db.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func CreateUser(user models.User) error {
	db := database.Get()

	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(user models.User) error {
	db := database.Get()

	err := db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	db := database.Get()

	err := db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
