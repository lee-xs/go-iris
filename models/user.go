package models

import (
	"go-iris/database"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model

	Username string `gorm:"not null VARCHAR(64)"`
	Password string `gorm:"not null VARCHAR(64)"`
	Name 	 string
	Email 	 string
}

type UserJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name 	 string `json:"name"`
	Email 	 string `json:"email"`
}

var db = database.DB

func CreateUser(data *UserJson) (user *User,err error) {

	user = new(User)

	user.Username = data.Username
	user.Password = data.Password
	user.Name = data.Name
	user.Email = data.Email

	if u := GetUserByUsername(user.Username); u.ID != 0 {
		log.Println("CreateUser 用户名已存在")
		return user, errors.New("CreateUser 用户名已存在")
	}

	if err = db.Create(&user).Error; err != nil {
		log.Println("CreateUser 创建数据时出现错误")
	}
	return
}

func GetUserById(id uint) *User{
	user := new(User)
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println("GetUserById 查询数据库出现错误")
	}
	return user
}

func GetAllUsers() (users *[]User) {
	users = new([]User)
	if err := db.Find(&users).Error; err != nil {
		log.Println("GetAll 查询数据库出现错误")
	}
	return users
}

func GetUserByUsername(username string) *User {
	user := new(User)
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("GetUserByUsername -> [%s] 不存在", username)
	}
	return user
}

func DelUserById(id uint) {
	user := new(User)
	user.ID = id
	if err := db.Delete(&user).Error; err != nil {
		log.Println("DelUserById 删除数据失败", err)
	}
	fmt.Println(user)
}
