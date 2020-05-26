package db
// package main

import (
	"barter-server2/pkg/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


var user model.User

// Connect to DB
func Dbconnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	DBNAME := "db_barter"
	CONNECT := USER + ":" + "@" + "/" + DBNAME
	return gorm.Open(DBMS, CONNECT)
}

// Connect DB
func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	DBNAME := "db_barter"
	CONNECT := USER + ":" + "@" + "/" + DBNAME
	return DBMS, CONNECT
}

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	if err != nil {
	   panic(err)
	} else {
	   fmt.Println("DB接続成功")
	}
 
	db.LogMode(true)
	return db
 }

 // ユーザを検索する
 func FindUser(db *gorm.DB, id int, pass string) {
	 db.Where(&model.User{ID: id, Password: pass}).Find(&user)
	 fmt.Println(user)
 }
 

// func main() {
// 	db := GetDBConn()
// 	FindUser(db, 1, "0000")
// 	// テーブル作成
// 	// db.AutoMigrate(&model.User{})
// }
