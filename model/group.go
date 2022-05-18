package model

import "minepin/com/db"

type PinGroup struct {
	BaseModel
	Name   string `json:"name"`
	UserId uint64 `json:"user_id"`
	Note   string `json:"note"`
}

type GroupBind struct {
	Name   string `json:"name"`
	UserId uint64 `json:"user_id"`
	Note   string `json:"note"`
}

func (u *User) CreateGroup(gb GroupBind) (g PinGroup, err error) {
	g = PinGroup{
		Name:   gb.Name,
		Note:   gb.Note,
		UserId: u.Id,
	}
	//err = db.DB.Create(&pin).Error
	err = db.DB.Model(&u).Association("Groups").Append(&g)
	return
}

func (u *User) GroupList() (groups []PinGroup, err error) {
	err = db.DB.Model(&u).Order("createdAt desc").Association("Groups").Find(&groups)
	if err != nil {
		return nil, err
	}
	return
}

func (u *User) GetGroupByUUID(uid string) (group PinGroup, err error) {
	err = db.DB.Model(&u).
		Where("uuid = ?", uid).Association("Groups").Find(&group)
	return
}

func (g *PinGroup) User() (user User) {
	db.DB.First(&user, g.UserId)
	return
}

func (g *PinGroup) UpdateGroup() (err error) {
	// 保证仅更新非零字段
	err = db.DB.Where("uuid = ?", g.UUID).Updates(PinGroup{
		Name: g.Name,
		Note: g.Note,
	}).Error
	return
}

func (g *PinGroup) Delete() (err error) {
	return db.DB.Delete(&g).Error
}