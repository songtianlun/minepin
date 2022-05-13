package model

import "minepin/com/db"

type Thread struct {
	BaseModel
	Topic  string
	UserId uint64
}

func (t *Thread) NumReplies() (count int64) {
	//db.DB.Where("thread_id = ?", t.Id).Find(&Post{})
	db.DB.Model(&Post{}).Where("thread_id = ?", t.Id).Count(&count)
	//	count = d.RowsAffected
	return
}

func (u *User) CreateThread(topic string) (thread Thread, err error) {
	thread = Thread{
		Topic:  topic,
		UserId: u.Id,
	}
	err = db.DB.Create(&thread).Error
	return
}

func Threads() (threads []Thread, err error) {
	err = db.DB.Find(&threads).Error
	return
}

func (t *Thread) CreatedAtDate() string {
	return t.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (t *Thread) Posts() (posts []Post, err error) {
	err = db.DB.Where("thread_id = ?", t.Id).Find(&posts).Error
	return
}

func (t *Thread) User() (user User) {
	db.DB.First(&user, t.UserId)
	return
}

func ThreadByUUID(tid string) (thread Thread, err error) {
	err = db.DB.Where("uuid = ?", tid).First(&thread).Error
	return
}
