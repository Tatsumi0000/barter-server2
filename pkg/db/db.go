package db

// package main

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
	PASS := "barterdb8"
	DBNAME := "db_barter"
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME
	return gorm.Open(DBMS, CONNECT)
}

// Connect DB
func GetDBConfig() (string, string) {
	DBMS := "mysql"
	USER := "root"
	PASS := "barterdb8"
	DBNAME := "db_barter"
	CONNECT := USER + ":" + PASS + "@" + "/" + DBNAME
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
func GetCommunityUser(db *gorm.DB) []model.CommunityUser {
	var communityUsers []model.CommunityUser
	// DB内からIDと名前を全て引っこ抜く。
	db.Table("users").Select("id, name").Find(&communityUsers)
	fmt.Println(communityUsers)
	return communityUsers
}

// ポイントを更新するやーつ
func UpdatePoint(db *gorm.DB, pointUpdate model.PointUpdate) {
	// 受け取った方のやつ UPDATE users SET point = point + 100 WHERE id = 1;
	db.Table("users").Where("id = ?", pointUpdate.ReceiveUserID).Update("point", gorm.Expr("point + ?", pointUpdate.Point))
	// 送った方のやつ UPDATE users SET point = point + 100 WHERE id = 1;
	db.Table("users").Where("id = ?", pointUpdate.SendUserID).Update("point", gorm.Expr("point - ?", pointUpdate.Point))
}

// func main() {
// 	// var communityUser []model.CommunityUser
// 	// var pointUpdate = model.PointUpdate{ReceiveUserID: 1, SendUserID: 2, Point: 200}
// 	db := GetDBConn()
// 	// UpdatePoint(db, pointUpdate)
// 	// GORMは、自動でFind()のやつからテーブルを参照するがTableで指定すればそっちを参照する
// 	// db.Table("users").Select("id, name").Find(&communityUser)
// 	// fmt.Println(communityUser)
// 	// テーブル作成
// 	db.AutoMigrate(&model.PointHistory{})
// }
