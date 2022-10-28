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

// // Create
// db.Create(&Product{Code: "D42", Price: 100})

// // Read
// var product Product
// db.First(&product, 1) // find product with integer primary key
// db.First(&product, "code = ?", "D42") // find product with code D42

// // Update - update product's price to 200
// db.Model(&product).Update("Price", 200)
// // Update - update multiple fields
// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

// // Delete - delete product
// db.Delete(&product, 1)
