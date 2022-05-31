package model

import (
	coordTransform "github.com/qichengzx/coordtransform"
	"minepin/com/constvar"
	"minepin/com/db"
	"minepin/com/log"
	"minepin/com/utils"
)

type Pin struct {
	BaseModel
	Location string
	Lat      string
	Lng      string
	CRS      string `gorm:"default:'BD09'"`
	Note     string
	UserId   uint64
	GroupId  uint64
	Group    PinGroup
}

type PinBind struct {
	Location string
	Lat      string
	Lng      string
	Note     string
	CRS      string
	Group    PinGroup
}

type Pins struct {
	Group       PinGroup
	Pins        []Pin
	BaiduAK     string
	TianDiTuKey string
}

func (u *User) CreatePin(pb PinBind) (pin Pin, err error) {
	pin = Pin{
		Location: pb.Location,
		Lat:      pb.Lat,
		Lng:      pb.Lng,
		Note:     pb.Note,
		CRS:      pb.CRS,
		UserId:   u.Id,
		GroupId:  pb.Group.Id,
	}
	//err = db.DB.Create(&pin).Error
	err = db.DB.Model(&u).Association("Pins").Append(&pin)
	return
}

func (u *User) PinList() (pins []Pin, err error) {
	err = db.DB.Model(&u).Order("createdAt desc").Association("Pins").Find(&pins)
	//for i, p := range pins {
	//	pins[i].Group, _ = u.GetGroupByID(p.GroupId)
	//}
	TransformPins(&pins)
	if err != nil {
		return nil, err
	}
	return
}

func (u *User) GetPinByUUID(uid string) (pin Pin, err error) {
	err = db.DB.Model(&u).
		Where("uuid = ?", uid).Association("Pins").Find(&pin)
	return
}

func (u *User) TransfromWithBD09() {
	var pins []Pin
	err := db.DB.Model(&u).
		Where("crs = ?", "BD09").Association("Pins").Find(&pins)
	if err != nil {
		return
	}
	if len(pins) != 0 {
		for _, pin := range pins {
			pin.TransformBD09()
			err := pin.UpdatePin()
			if err != nil {
				log.ErrorF("Trans BD09 with pin [UUID=%v] failed - %v",
					pin.UUID, err.Error())
			}
		}
	}
}

func (p *Pin) User() (user User) {
	db.DB.First(&user, p.UserId)
	return
}

func (p *Pin) Groups() (groups []PinGroup) {
	user := p.User()
	groups, _ = user.GroupList()
	return
}

//func (p *Pin) GetGroup() (group PinGroup) {
//	groups, _ = db.DB.Model(&p).Association("Group").Find(&group)
//	return
//}

func (p *Pin) UpdatePin() (err error) {
	// 保证仅更新非零字段
	err = db.DB.Where("uuid = ?", p.UUID).Updates(Pin{
		Location: p.Location,
		Lat:      p.Lat,
		Lng:      p.Lng,
		Note:     p.Note,
		GroupId:  p.Group.Id,
		CRS:      p.CRS,
		//Group:    p.Group,
	}).Error
	return
}

func (p *Pin) Delete() (err error) {
	return db.DB.Delete(&p).Error
}

func (p *Pin) TransformBD09() {
	if p.CRS == constvar.CRSBd09 {
		lng, lat := coordTransform.BD09toWGS84(utils.StrToFloat64(p.Lng), utils.StrToFloat64(p.Lat))
		p.Lng = utils.Float64ToStr(lng)
		p.Lat = utils.Float64ToStr(lat)
		p.CRS = constvar.CRSWgs84
	}
}

func TransformPins(pins *[]Pin) {
	for k, _ := range *pins {
		(*pins)[k].TransformBD09()
	}
}

//func GetPinByUUID(pid string) (p Pin, err error) {
//	err = db.DB.Where("uuid = ?", pid).First(&p).Error
//	return
//}
