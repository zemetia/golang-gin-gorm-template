package config

import (
	"fmt"
	entity "golang-gin-gorm-template/domain/model/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func gormDatabaseConnect(dbName string) (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func createDatabase() (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")

	db, err := gormDatabaseConnect("information_schema")
	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE `%s`", dbName)
	db.Exec(createDatabaseCommand)
	db.Commit()

	db, err = gormDatabaseConnect(dbName)

	return db, err
}

func SetupDatabaseConnection() *gorm.DB {
	dbName := os.Getenv("DB_NAME")

	db, err := gormDatabaseConnect(dbName)
	if err != nil {
		fmt.Println("Database not Found! Trying to automatically create a new database instead!")
		db, err = createDatabase()
		if err != nil {
			fmt.Println("Can't create a new database!")
			fmt.Println(err)
			panic(err)
		}
	}

	if err := db.AutoMigrate(
		entity.User{},
	); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	dbSQL.Close()
}
