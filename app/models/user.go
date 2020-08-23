package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	gorm.Model
	Name			string `gorm:"size:255"`
	Email          	string `gorm:"type:varchar(100);unique_index"`
	HashedPassword 	[]byte
	Active         	bool
}

func (user *User) SetNewPassword(passwordString string) {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordString), bcrypt.DefaultCost)
	user.HashedPassword = bcryptPassword
}

func (user *User) IsValidPassword(passwordString string) bool {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(passwordString))
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}
