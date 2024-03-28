package Database

import (
	"fmt"
	"go_api/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormConnect() *gorm.DB {
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
	db.AutoMigrate(
		&Models.Article{},
		&Models.Company{},
		&Models.OBVisits{},
	)

	fmt.Println("--------------------------------- ")
	fmt.Println("Database connected: ", &db)
	fmt.Println("--------------------------------- ")
	return db
}
