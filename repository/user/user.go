package user

import (
	"errors"
	"fmt"
	"todo/entity"

	"todo/utils"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

type UserRepo struct {
	Db *gorm.DB
}

// Check users islogin with payload
func (ur *UserRepo) IsLogin(email, password string) (entity.User, bool, error) {
	var u entity.User
	var pwd string

	query := "SELECT id, name, email, password FROM users WHERE email = ?"

	if err := ur.Db.Raw(query, email).Scan(&u).Error; err != nil {
		log.Warn(err)
		return entity.User{}, false, errors.New("tidak dapat select data")
	}
	fmt.Println(u)
	pwd = u.Password

	match, err := utils.CheckPasswordHash(password, pwd)
	if !match {
		log.Warn("Hash and password doesnt match")
		return u, false, err
	}

	log.Info()
	return u, true, nil
}

func (ur *UserRepo) Register(newUser entity.User) (entity.User, error) {
	if err := ur.Db.Create(&newUser).Error; err != nil {
		return entity.User{}, errors.New("tidak dapat insert data")
	}
	log.Info("insert succes!")
	return newUser, nil
}
