package users

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RoleType string

const (
	RoleAdmin RoleType = "admin"
	RoleUser  RoleType = "user"
)

type User struct {
	ID             uint     `gorm:"column:id;primaryKey"`
	EpicID         string   `gorm:"column:epic_id;unique"`
	Email          string   `gorm:"column:email;unique"`
	Password       string   `gorm:"column:password"        json:"-"`
	Balance        uint     `gorm:"column:balance"`
	IsBlocked      bool     `gorm:"column:is_blocked"`
	Role           RoleType `gorm:"column:role"`
	ProfilePicture string   `gorm:"column:profile_picture"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ClearPassword() {
	u.Password = ""
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	err = u.HashPassword()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	return nil
}

type Token struct {
	ID           uint   `gorm:"column:id;primaryKey"`
	RefreshToken string `gorm:"column:refresh_token"`
	UserID       uint
	User         User
}

type UserWithTokens struct {
	User         User
	AccessToken  string
	RefreshToken string
}

type GetUsersQuery struct {
	ID       uint   `json:"id"        form:"id"`
	EpicID   string `json:"epic_id"   form:"epic_id"`
	Email    string `json:"email"     form:"email"`
	Sort     string `json:"sort"      form:"sort"`
	Page     int    `json:"page"      form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type RegisterRequest struct {
	Email    string `json:"email"    binding:"required"`
	EpicID   string `json:"epic_id"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}
