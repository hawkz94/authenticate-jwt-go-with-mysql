package database

import (
	"fmt"
	"log"
	"os"
	"tecmentor-api/database/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func StartDB() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	password := os.Getenv("DATABASE_PASSWORD")
	user := os.Getenv("DATABASE_USER")
	database := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")

	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Mysql Database")
		log.Fatal("Error: ", err)
	}
	migrations.RunMigrate(db)
	return db
}

// func CloseConn() error {
// 	config, err := db.DB()
// 	if err != nil {
// 		fmt.Println("Caiiiiuuuuuuuuuuu")
// 		return err
// 	}

// 	err = config.Close()
// 	if err != nil {
// 		fmt.Println("Caiiiiuuuuuuuuuuu 2322222")
// 		return err
// 	}

// 	return nil
// }

func GetDatabase() *gorm.DB {
	Db = StartDB()
	return Db
}
