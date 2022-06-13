package model

import (
	"fmt"
	"minepin/com/constvar"
	"minepin/com/db"
)

type PinGroup struct {
	BaseModel
	Name   string `json:"name"`
	Note   string `json:"note"`
	Type   string `json:"type" gorm:"default:'cluster'"`
	UserId uint64 `json:"user_id"`
}

type GroupBind struct {
	Name   string `json:"name"`
	UserId uint64 `json:"user_id"`
	Note   string `json:"note"`
	Type   string `json:"type"`
}

func (u *User) CreateGroup(gb GroupBind) (g PinGroup, err error) {
	g = PinGroup{
		Name:   gb.Name,
		Note:   gb.Note,
		Type:   gb.Type,
		UserId: u.Id,
	}
	//err = db.DB.Create(&pin).Error
	err = db.DB.Model(&u).Association("Groups").Append(&g)
	return
}

func (u *User) GroupList() (groups []PinGroup, err error) {
	err = db.DB.Model(&u).Order("createdAt desc").Association("Groups").Find(&groups)
	return
}

func (u *User) GetGroupByUUID(uid string) (group PinGroup, err error) {
	err = db.DB.Model(&u).
		Where("uuid = ?", uid).Association("Groups").Find(&group)
	return
}

func (u *User) GetGroupByID(id uint64) (group PinGroup, err error) {
	err = db.DB.Model(&u).Association("Groups").Find(&group, id)
	return
}

func (u *User) ShowPinsByGroupID(gid uint64) (pins []Pin, err error) {
	err = db.DB.Model(&u).Where("group_id = ?", gid).
		Association("Pins").Find(&pins)
	return
}

func (u *User) GetGroupCount() int64 {
	return db.DB.Model(&u).Association("Groups").Count()
}

func (g *PinGroup) User() (user User) {
	db.DB.First(&user, g.UserId)
	return
}

func (g *PinGroup) UpdateGroup() (err error) {
	oldGroup := PinGroup{}
	err = db.DB.Model(&PinGroup{}).Where("uuid = ?", g.UUID).First(&oldGroup).Error
	if err != nil {
		return
	}
	if oldGroup.Name == constvar.DefaultGroupName {
		err = fmt.Errorf("default group can not be edited")
		return
	}
	// 保证仅更新非零字段
	err = db.DB.Where("uuid = ?", g.UUID).Updates(PinGroup{
		Name: g.Name,
		Note: g.Note,
		Type: g.Type,
	}).Error
	return
}

func (g *PinGroup) Delete() (err error) {
	if g.Name == constvar.DefaultGroupName {
		err = fmt.Errorf("default group can not be deleted")
		return
	}
	return db.DB.Delete(&g).Error
}
