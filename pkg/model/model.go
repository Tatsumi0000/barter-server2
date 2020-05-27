package model

import "time"

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

// ポイントをもらう側のユーザリストを管理する
type CommunityUser struct {
	ID   int
	Name string
}

// ポイントをアップデートする
type PointUpdate struct {
	SendUserID    int
	ReceiveUserID int
	Point         int
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

type PointHistory struct {
	ID            int `gorm:"primary_key"`
	SendUserID    int
	ReceiveUserID int
	SendPoint     int
	Date          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type UserInfo struct {
	JWT   string // JWT
	ID    int    // ID
	Name  string // ユーザ名
	Point int    // 所持ポイント
}

// tb_point_historys	point_id	int	PRIMARY_KEY	AUTO_INCREMENT
//  id int PRIMARY_KEY AUTO_INCREMENT ポイントを送った番号
// 	send_user_id	int			送ったユーザのID
// 	receive_user_id	int			受取ったユーザのID
// 	send_point	int			送ったポイント
// 	date	DATETIME			日付
