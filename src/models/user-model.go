package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null;default:null"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//? ID string `json:"id" gorm:"primaryKey"`
//! Business rules with gorm tags ^
