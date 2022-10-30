package models

type User struct {
	ID       string `json:"_id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

//? ID string `json:"id" gorm:"primaryKey"`
//! Business rules with gorm tags ^
