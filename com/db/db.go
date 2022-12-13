package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"minepin/com/log"
)

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
		panic("failed to connect database - " + err.Error())
	}

	return db
}

func openMySqlDB(path string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect with MySQL - " + err.Error())
	}
	return db
}

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
	case "mysql":
		gdb = openMySqlDB(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Username, c.Passwd, c.Addr, c.Name))
	default:
		panic("unknown database type: " + c.Typ)
	}
	log.Infof("connected to %s with %s",
		c.Addr, c.Typ)

	DB = gdb
	// DB = &Database{
	//	DB: gdb,
	// }

	// err := DB.Callback().Query().Register("check_acl", CheckACL)
	// if err != nil {
	//	log.Error("Error of Register Callback - " + err.Error())
	// }
}
