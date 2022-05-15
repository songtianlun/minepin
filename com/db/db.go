package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"minepin/com/log"
)

//var MapDB = make(map[string]interface{})

type CfgDb struct {
	Typ      string
	Addr     string
	Username string
	Passwd   string
	Name     string
}

var DB *gorm.DB

func openSqliteDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func MigrateModel(model interface{}) {
	err := DB.AutoMigrate(&model)
	if err != nil {
		panic("fail to auto migrate db: " + err.Error())
	}
}

//func CheckACL(db *gorm.DB) {
//	log.Error("Check ACL!")
//}

func InitDB(c *CfgDb) {
	var gdb *gorm.DB
	switch c.Typ {
	case "sqlite":
		gdb = openSqliteDB(c.Addr)
	default:
		panic("unknown database type: " + c.Typ)
	}
	log.InfoF("connected to %s with %s",
		c.Addr, c.Typ)

	DB = gdb
	//DB = &Database{
	//	DB: gdb,
	//}

	//err := DB.Callback().Query().Register("check_acl", CheckACL)
	//if err != nil {
	//	log.Error("Error of Register Callback - " + err.Error())
	//}
}
