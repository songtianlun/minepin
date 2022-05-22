package main

import (
	"fmt"
	"minepin/com/cfg"
	"minepin/com/cli"
	"minepin/com/db"
	"minepin/com/log"
	"minepin/com/v"
	"minepin/com/web"
	"minepin/handle"
	"minepin/model"
)

func initHandle() {

	// static file
	web.RegisterFile("/static/", cfg.GetString("Static"), true)

	// index
	web.RegisterHandle("/", handle.Index)

	// error
	web.RegisterHandle("/err", handle.Err)

	// defined in route_auth.go
	web.RegisterHandle("/login", handle.Login)
	web.RegisterHandle("/signup", handle.Signup)
	web.RegisterHandle("/signup_account", handle.SignupAccount)
	web.RegisterHandle("/authenticate", handle.Authenticate)

	web.RegisterHandle("/logout", handle.Logout)

	// defined minepin
	web.RegisterHandle("/minepin", handle.MinePinIndex, handle.Auth)
	web.RegisterHandle("/pin/new", handle.NewPin, handle.Auth)
	web.RegisterHandle("/pin/create", handle.CreatePin, handle.Auth)
	web.RegisterHandle("/pin/edit", handle.EditPin, handle.Auth)
	web.RegisterHandle("/pin/update", handle.UpdatePin, handle.Auth)
	web.RegisterHandle("/pin/delete", handle.DeletePin, handle.Auth)

	// defined group
	web.RegisterHandle("/group", handle.PinGroupIndex, handle.Auth)
	web.RegisterHandle("/group/new", handle.NewGroup, handle.Auth)
	web.RegisterHandle("/group/create", handle.CreateGroup, handle.Auth)
	web.RegisterHandle("/group/edit", handle.EditGroup, handle.Auth)
	web.RegisterHandle("/group/update", handle.UpdateGroup, handle.Auth)
	web.RegisterHandle("/group/delete", handle.DeleteGroup, handle.Auth)
	web.RegisterHandle("/group/show", handle.ShowGroup, handle.Auth)
}

func initCfg() {
	// 首先完成配置项的注册
	cfg.RegisterCfg("Port", 6008, "int64")
	//cfg.RegisterCfg("Address", "0.0.0.0:6008", "string")
	cfg.RegisterCfg("ReadTimeout", 10, "int64")
	cfg.RegisterCfg("WriteTimeout", 600, "int64")
	cfg.RegisterCfg("Static", "public", "string")
	cfg.RegisterCfg("SessionTimeoutHour", 6, "int64")
	cfg.RegisterCfg("BaiduAK", "<YOUR_BAIDU_AK>", "string")
	// log
	cfg.RegisterCfg("log.level", "info", "string")
	cfg.RegisterCfg("log.file_name", "log/minegin.log", "string")
	cfg.RegisterCfg("log.max_size_mb", 1, "int")
	cfg.RegisterCfg("log.max_file_num", 64, "int")
	cfg.RegisterCfg("log.max_file_day", 7, "int")
	cfg.RegisterCfg("log.compress", false, "bool")
	cfg.RegisterCfg("log.stdout", true, "bool")
	cfg.RegisterCfg("log.only_stdout", false, "bool")
	// db
	cfg.RegisterCfg("db.type", "sqlite", "string")
	cfg.RegisterCfg("db.addr", "./minepin.db", "string")
	cfg.RegisterCfg("db.name", "minepin", "string")
	cfg.RegisterCfg("db.username", "minepin", "string")
	cfg.RegisterCfg("db.password", "minepin", "string")

	// 之后再进行初始化
	err := cfg.Init("")
	if err != nil {
		panic(fmt.Sprintf("init cfg failed: %s", err))
	}
}

func initLog() {
	log.InitLogger(
		cfg.GetString("log.file_name"),
		cfg.GetString("log.level"),
		cfg.GetInt("log.max_size_mb"),
		cfg.GetInt("log.max_file_num"),
		cfg.GetInt("log.max_file_day"),
		cfg.GetBool("log.compress"))
}

func initDB() {
	db.InitDB(&db.CfgDb{Typ: "sqlite", Addr: cfg.GetString("db.addr")})

	db.MigrateModel(model.User{})
	db.MigrateModel(model.Session{})
	db.MigrateModel(model.Pin{})
	db.MigrateModel(model.PinGroup{})
}

func runCLI() (isCli bool) {
	cli.RegisterCLI("version", "V", "show version info.", func() {
		fmt.Println(v.GetVersionStr())
	})
	return cli.CheckCLI()
}
