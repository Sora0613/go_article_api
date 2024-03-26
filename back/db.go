package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func gormConnect() *gorm.DB {
	USER := "go_user"
	PASS := "go_password"
	PROTOCOL := "tcp(go-db:3306)"
	DBNAME := "go_task_db"
	DSN := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	fmt.Println("--------------------------------- ")
	fmt.Println("Database connected: ", &db)
	fmt.Println("--------------------------------- ")
	return db
}

func CloseDB(db *gorm.DB) {
	cdb, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	defer cdb.Close()
}
