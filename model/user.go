package model

import (
	"minepin/com/constvar"
	"minepin/com/db"
	"minepin/com/utils"
	"net/http"
)

type User struct {
	BaseModel
	Role     constvar.UserType `json:"role" gorm:"column:role;default:1"`
	Name     string            `json:"name" gorm:"column:name;not null;default:-" validate:"min=1,max=128"`
	Email    string            `json:"email" gorm:"column:email;unique" validate:"max=64"`
	Password string            `json:"password" gorm:"column:password;not null" validate:"min=5,max=128"`
}

type Session struct {
	BaseModel
	Email  string `json:"email" gorm:"column:email" validate:"max=64"`
	UserId uint64 `gorm:"column:uid;comment:UserID" `
}

func (u *User) Create() error {
	u.Password = utils.Encrypt(u.Password)
	return db.DB.Create(&u).Error
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{
		Email:  u.Email,
		UserId: u.Id,
	}
	err = db.DB.Create(&session).Error
	return
}

func (u *User) Session() (session Session, err error) {
	err = db.DB.First(&session, u.Id).Error
	return
}

func (s *Session) User() (user User, err error) {
	err = db.DB.First(&user, s.UserId).Error
	return
}

func UserByEmail(email string) (u User, err error) {
	d := db.DB.Where("email = ?", email).First(&u)
	err = d.Error
	return
}

func Check(sid string) (s Session, err error) {
	err = db.DB.Where("uuid = ?", sid).First(&s).Error
	return
}

func CheckSession(request *http.Request) (session Session, err error) {
	cookie, err := request.Cookie("_cookie")
	session, err = Check(cookie.Value)
	return
}

func DeleteSession(uuid string) {
	db.DB.Where("uuid = ?", uuid).Delete(&Session{})
}
