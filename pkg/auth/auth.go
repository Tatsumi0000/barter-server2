package auth

import (
	"barter-server2/pkg/db"
	// "barter-server2/pkg/handler"
	"barter-server2/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
)

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// GetTokenHandler get token
func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
  var userAuth model.UserAuth
  decoder := schema.NewDecoder()
  err := r.ParseForm()
  err = decoder.Decode(&userAuth, r.PostForm)
  if err != nil {
    fmt.Println("ERROR")
  }
  fmt.Println("POSTの中身は: ", userAuth.ID, userAuth.Password)
	fmt.Println("http.Rquestの中身: ", r.Body)
	fmt.Println("POSTの中身:", userAuth)
	fmt.Println("ID: ", userAuth.ID, "\nPASS: ", userAuth.Password)
	// DB接続
	dbGorm := db.GetDBConn()
	// ユーザIDとパスワードを検索し、データの中身を代入
	user := db.FindUser(dbGorm, userAuth.ID, userAuth.Password)
	// IDとPASSがどっちともDBにないとゼロ値（0）が入る
	if user.ID == 0 {
		// fmt.Println("IDもしくはPASSが間違っています。")
		w.Write([]byte("IDもしくはパスワードが間違っています。"))
		return
	}

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID // 個人の識別子を埋め込む
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

  userInfo := &model.UserInfo{
    JWT: tokenString,
    ID: user.ID,
    Name: user.Name,
    Point: user.Point,
  }



  	// post := &post{
	// 	Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
	// 	Tag:   "Vue.js",
	// 	URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
  // }
	// JWTを返却
	json.NewEncoder(w).Encode(userInfo)
}

var GetTokenHandler2 = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var userAuth model.UserAuth
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&userAuth)
	fmt.Println("http.Rquestの中身: ", r)
	fmt.Println("POSTの中身:", userAuth)
	fmt.Println("ID: ", userAuth.ID, "\nPASS: ", userAuth.Password)
	// DB接続
	dbGorm := db.GetDBConn()
	// ユーザIDとパスワードを検索し、データの中身を代入
	user := db.FindUser(dbGorm, userAuth.ID, userAuth.Password)
	// IDとPASSがどっちともDBにないとゼロ値（0）が入る
	if user.ID == 0 {
		fmt.Println("IDもしくはPASSが間違っています。")
		w.Write([]byte("IDもしくはパスワードが間違っています。"))
		return
	}

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID // 個人の識別子を埋め込む
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	w.Write([]byte(tokenString))
	// w.Write([]byte("AAAAAA"))
})

var Public2 = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var userAuth model.UserAuth
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&userAuth)
	fmt.Println("POSTの中身:", userAuth)
	// post := &handler.Post{
	// 	Title: "VueCLIからVue.js入門①【VueCLIで出てくるファイルを概要図で理解】",
	// 	Tag:   "Vue.js",
	// 	URL:   "https://qiita.com/po3rin/items/3968f825f3c86f9c4e21",
	// }
	w.Write([]byte("post"))
	// json.NewEncoder(w).Encode(post)
})
