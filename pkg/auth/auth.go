package auth

import (
	"barter-server2/pkg/db"
	"barter-server2/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var userAuth model.UserAuth
	// ここでPOSTをいい感じにする
	json.NewDecoder(r.Body).Decode(&userAuth)
	fmt.Println("http.Rquestの中身: ", r)
	fmt.Println("POSTの中身:", userAuth)
	fmt.Println("ID: ", userAuth.ID, "\nPASS: ", userAuth.Password)
	fmt.Println("Bodyの中身: ", r.Body)
	// DB接続
	dbGorm := db.GetDBConn()
	// ユーザIDとパスワードを検索
	db.FindUser(dbGorm, 1, "0000")

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
