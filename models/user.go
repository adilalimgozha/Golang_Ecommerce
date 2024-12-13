package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Role struct {
	RoleID   uint   `gorm:"primaryKey"`
	RoleName string `gorm:"not null"`
}

type User struct {
	UserID       uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Email        string    `gorm:"unique;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	RoleID       uint      `gorm:"not null"`
	Role         Role      `gorm:"foreignKey:RoleID"`
}

// set hash
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// check password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
