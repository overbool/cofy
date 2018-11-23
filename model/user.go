package model

import (
	"github.com/jinzhu/gorm"
	"github.com/overbool/cofy/core/storage/mysql"
	"github.com/overbool/cofy/pkg/auth"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Create creates a new user.
func (u *User) Create() error {
	return mysql.DB.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint) error {
	user := User{}
	user.Model.ID = id
	return mysql.DB.Delete(&user).Error
}

// Update updates an user information.
func (u *User) Update() error {
	return mysql.DB.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*User, error) {
	u := &User{}
	d := mysql.DB.Where("username = ?", username).First(&u)
	return u, d.Error
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}
