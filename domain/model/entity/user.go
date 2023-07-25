package entity

import (
	"golang-gin-gorm-template/helpers"

	"gorm.io/gorm"
)

type User struct {
	ID             string `gorm:"primary_key;size:36;not_null" json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email" binding:"email"`
	NoTelp         string `json:"no_telp"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	MembershipTier string `json:"membership"`
	PhotoURL       string `json:"photo_url"`
	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
