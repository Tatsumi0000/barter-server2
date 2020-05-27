package handler

import (
	"barter-server2/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

var Public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
		Tag:   "Vue.js",
		URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	}
	json.NewEncoder(w).Encode(post)
})

var Private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &Post{
		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
		Tag:   "Go",
		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	}
	json.NewEncoder(w).Encode(post)
})

// ユーザ情報を返すや〜つ
func UserInfo(w http.ResponseWriter, r *http.Request) {

	var pointUpdate model.PointUpdate
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&pointUpdate)
	fmt.Println("http.Rquestの中身: ", r)
	fmt.Println("POSTの中身:", pointUpdate)
	// ポイントの更新をする。
}

// ポイントをアップデートするや〜つ
func UpdatePoint(w http.ResponseWriter, r *http.Request) {

	var pointUpdate model.PointUpdate
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&pointUpdate)
	fmt.Println("http.Rquestの中身: ", r)
	fmt.Println("POSTの中身:", pointUpdate)
	// ポイントの更新をする。
}

// ポイントの履歴を返すや〜つ
func PointHistory(w http.ResponseWriter, r *http.Request) {

	var pointHistory model.PointHistory
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&pointHistory)
	fmt.Println("http.Rquestの中身: ", r)
	fmt.Println("POSTの中身:", pointHistory)
	// ポイントの履歴を取得する
  
	// ポイントの履歴を返す。

}

// デバッグ用のやつ
var Public2 = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var userAuth model.UserAuth
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&userAuth)
	fmt.Println("http.Rquestの中身: ", r)
	fmt.Println("POSTの中身:", userAuth)
	fmt.Println("Bodyの中身: ", r.Body)

	// post := &post{
	// 	Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
	// 	Tag:   "Vue.js",
	// 	URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	// }
	json.NewEncoder(w).Encode(userAuth)
})
