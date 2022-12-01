package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64 `json:"id" gorm:"primaryKey;auto_increment"`
	FirstName string `json:"firstName" gorm:"not null;default:null"`
	LastName  string `json:"lastName" gorm:"not null;default:null"`
	Email     string `json:"email" gorm:"not null;default:null"`
	Age       uint   `json:"age" gorm:"not null;default:0"`
}

func GetUsers() ([]User, error) {
	// get all users

	var users []User

	tx := db.Find(&users)

	if tx.Error != nil {
		return users, tx.Error
	}

	return users, nil
}

func GetUser(id uint64) (User, error) {
	// get user by id

	var user User

	tx := db.Where("id = ?", id).First(&user)

	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func CreateUser(user User) error {
	// create user

	tx := db.Create(&user)

	return tx.Error
}

func UpdateUser(id uint64, user User) error {
	// update user

	tx := db.Where("id = ?", id).Updates(user)

	return tx.Error
}

func DeleteUser(id uint64) error {
	// delete user

	tx := db.Unscoped().Delete(&User{}, id)

	return tx.Error
}
