package main

import (
	"fmt"
	"minepin/com/cfg"
	"minepin/com/db"
	"minepin/com/log"
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
	web.RegisterHandle("/logout", handle.Logout)
	web.RegisterHandle("/signup", handle.Signup)
	web.RegisterHandle("/signup_account", handle.SignupAccount)
	web.RegisterHandle("/authenticate", handle.Authenticate)

	// defined in route_thread.go
	web.RegisterHandle("/thread/new", handle.NewThread)
	web.RegisterHandle("/thread/create", handle.CreateThread)
	web.RegisterHandle("/thread/post", handle.PostThread)
	web.RegisterHandle("/thread/read", handle.ReadThread)

	// defined minepin
	web.RegisterHandle("/minepin", handle.MinePinIndex)
	web.RegisterHandle("/pin/new", handle.NewPin)
	web.RegisterHandle("/pin/create", handle.CreatePin)
	web.RegisterHandle("/pin/edit", handle.EditPin)
	web.RegisterHandle("/pin/update", handle.UpdatePin)
}

func initCfg() {
	// 首先完成配置项的注册
	cfg.RegisterCfg("Address", "0.0.0.0:6008", "string")
	cfg.RegisterCfg("ReadTimeout", 10, "int64")
	cfg.RegisterCfg("WriteTimeout", 600, "int64")
	cfg.RegisterCfg("Static", "public", "string")
	cfg.RegisterCfg("SessionTimeoutHour", 6, "int64")
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

	// 初始化结束后配置文件正常存取
	fmt.Printf("get cfg %v\n", cfg.Get("address", false))
	fmt.Printf("get cfg %v\n", cfg.Get("ReadTimeout", false))
	fmt.Printf("get cfg %v\n", cfg.Get("WriteTimeout", false))
	fmt.Printf("get cfg %v\n", cfg.Get("Static", false))
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
	db.InitDB(&db.CfgDb{Typ: "sqlite", Addr: "./minepin1.db"})

	db.MigrateModel(model.User{})
	db.MigrateModel(model.Session{})
	db.MigrateModel(model.Thread{})
	db.MigrateModel(model.Post{})
	db.MigrateModel(model.Pin{})
}
