package models

type User struct {
	ID       string `gorm:"primary_key" json:"id" bson:"id"`
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=6"`
}
