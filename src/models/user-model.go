package models

// import (
// 	"github.com/go-playground/validator/v10"
// )

//? ID string `json:"id" gorm:"primaryKey"`
//! Business rules with gorm tags ^

type User struct {
	ID       string `json:"_id" bson:"_id" `
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=6"`
}

// func (u *User) Validate() {
// 	v := validator.New()

// 	// Validate the user struct
// 	err := v.Struct(u)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
