// package db

package main

import (
	"barter-server2/pkg/model"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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
func FindUser(db *gorm.DB, id int, pass string) model.User {
  var user model.User
	//  db.Where(&model.User{ID: id, Password: pass}).Find(&user)
	db.Where("id = ? AND password = ?", id, pass).Find(&user)
	fmt.Println(user)
	return user
}

// テーブル内のユーザを全部取得
func GetCommunityUser(db *gorm.DB) model.CommunityUser {
	//  db.Where(&model.User{ID: id, Password: pass}).Find(&user)
	// users = []model.CommunityUser{}
	// DB内からIDと名前を全て引っこ抜く。
	db.Table("users").Select("id, name").Find(&communityUser)
	fmt.Println(users)
	return users
}

func main() {
	var communityUser []model.CommunityUser
	db := GetDBConn()
	// GORMは、自動でFind()のやつからテーブルを参照するがTableで指定すればそっちを参照する
	db.Table("users").Select("id, name").Find(&communityUser)
	fmt.Println(communityUser)
	// テーブル作成
	// db.AutoMigrate(&model.User{})
}
