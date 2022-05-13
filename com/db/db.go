package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"minepin/com/log"
)

var MapDB = make(map[string]interface{})

type CfgDb struct {
	Typ      string
	Addr     string
	Username string
	Passwd   string
	Name     string
}

//type Database struct {
//	DB *gorm.DB
//}

var DB *gorm.DB

func openSqliteDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

//func RegisterModel(name string, model interface{}) {
//	_, ok := MapDB[name]
//	if ok {
//		panic(fmt.Sprintf("DB Model %s is already registered", name))
//	}
//	MapDB[name] = model
//}

//func migrateModel() {
//	for _, v := range MapDB {
//		err := DB.AutoMigrate(v)
//		if err != nil {
//			panic("fail to auto migrate db: " + err.Error())
//		}
//	}
//}

func MigrateModel(model interface{}) {
	err := DB.AutoMigrate(&model)
	if err != nil {
		panic("fail to auto migrate db: " + err.Error())
	}
}

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
}
