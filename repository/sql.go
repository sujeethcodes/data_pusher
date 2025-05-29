package repository

import (
	"fmt"
	"os"

	"github.com/matryer/resync"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var onceMysql resync.Once
var mysqlConn *gorm.DB

type MysqlCon struct {
	Connection *gorm.DB
}

func SingletonMysqlCon() *MysqlCon {
	onceMysql.Do(func() {
		fmt.Println("Enter DBs Connection")
		userName := os.Getenv("USER_NAME")
		password := os.Getenv("PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, dbHost, dbName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("could not create db:", err)
			return
		}
		mysqlConn = db
	})
	return &MysqlCon{Connection: mysqlConn}
}
