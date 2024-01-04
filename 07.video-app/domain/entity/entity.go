package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Person struct {
	Id        uint   `json:"id" gorm:"primary_key;auto_increment"`
	FirstName string `json:"firstName" binding:"required" gorm:"column:first_name"`
	LastName  string `json:"lastName" binding:"required" gorm:"column:last_name"`
	Age       int    `json:"age" binding:"gte=1,lte=100" gorm:"column:age"`
	Email     string `json:"email" binding:"required,email" gorm:"column:email"`
}

func (Person) TableName() string {
	return "Person"
}

type Video struct {
	Id          uint      `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title" binding:"min=2,max=20" gorm:"column:title;type:varchar(100)"`
	Description string    `json:"description" binding:"max=40" gorm:"column:description;type:varchar(100)"`
	Url         string    `json:"url" binding:"required,url" gorm:"column:url;type:varchar(100)"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonId"`
	PersonId    uint      `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type User struct {
	Email    string `json:"email" binding:"required" gorm:"primary_key;type:varchar(100)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(256)"`
}

type CustomJwtClaims struct {
	jwt.RegisteredClaims
	Authorities string `json:"authorities"`
	Scope       string `json:"scopes"`
}
