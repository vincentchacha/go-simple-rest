package models

import (
	"github.com/example/go-simple-rest/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB



type User struct {
	gorm.Model
	//Id          string `json:"id"`
	FirstName        string `gorm:""json:"firstName"`
	LastName      string `json:"lastName"`
	Email string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User , *gorm.DB){
	var getUser User
	db:=db.Where("ID = ?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User {
	var User User
	db.Where("ID = ?", ID).Delete(User)
	return User
}