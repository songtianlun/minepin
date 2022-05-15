package model

import "minepin/com/db"

type Post struct {
	BaseModel
	Body     string
	UserId   uint64
	ThreadId uint64
}

func (u *User) CreatePost(thread Thread, body string) (post Post, err error) {
	post = Post{
		Body:     body,
		UserId:   u.Id,
		ThreadId: thread.Id,
	}
	err = db.DB.Create(&post).Error
	return
}

func (post *Post) User() (user User) {
	db.DB.First(&user, post.UserId)
	return
}
