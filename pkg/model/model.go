// package main
package model

// import (
// 	"barter-server2/pkg/db"
// )

// Gormは最初大文字じゃないとカラムを作ってくれない。
// ユーザのデータを持つ構造体
type User struct {
	ID       int `gorm:"primary_key"`
	Password string
	Name     string
	Point    int `gorm:"default:'100'"`
}

type UserAuth struct {
	ID       int
	Password string
}

// func main() {
// 	db := db.GetDBConn()
// 	// ユーザのテーブルを作成
// 	db.AutoMigrate(&User{})
// }

// user_id	int	PRIMARY_KEY	AUTO_INCREMENT	ユーザID
// user_pass	VARCHAR(32)			ユーザPASS
// user_name	VARCHAR(64)			ユーザ名
// point	int			ポイント
