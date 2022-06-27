package models

import (
	"html"
	"sispendik-api/utils/token"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	// Sekolah
	Users struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		RoleID    uint      `json:"role_id"`
		Nama      string    `json:"nama"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Users{}

	err = db.Model(Users{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Users) SaveUser(db *gorm.DB) (*Users, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Users{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}
