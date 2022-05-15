package model

import "minepin/com/db"

type Pin struct {
	BaseModel
	Location string
	Lat      string
	Lng      string
	Note     string
	UserId   uint64
}

type PinBind struct {
	Location string
	Lat      string
	Lng      string
	Note     string
}

func (u *User) CreatePin(pb PinBind) (pin Pin, err error) {
	pin = Pin{
		Location: pb.Location,
		Lat:      pb.Lat,
		Lng:      pb.Lng,
		Note:     pb.Note,
		UserId:   u.Id,
	}
	//err = db.DB.Create(&pin).Error
	err = db.DB.Model(&u).Association("Pins").Append(&pin)
	return
}

func (u *User) PinList() (pins []Pin, err error) {
	err = db.DB.Model(&u).Association("Pins").Find(&pins)
	if err != nil {
		return nil, err
	}
	return
}

func (p *Pin) User() (user User) {
	db.DB.First(&user, p.UserId)
	return
}

func (p *Pin) UpdatePin() (err error) {
	// 保证仅更新非零字段
	err = db.DB.Where("uuid = ?", p.UUID).Updates(Pin{
		Location: p.Location,
		Lat:      p.Lat,
		Lng:      p.Lng,
		Note:     p.Note,
	}).Error
	return
}

func GetPinByUUID(pid string) (p Pin, err error) {
	err = db.DB.Where("uuid = ?", pid).First(&p).Error
	return
}
