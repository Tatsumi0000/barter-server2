package main

import (
	"barter-server2/pkg/auth"
	"barter-server2/pkg/handler"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// CORSの設定
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:5000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})
	r := mux.NewRouter()
	r.Handle("/public", handler.Public)
	r.Handle("/private", auth.JwtMiddleware.Handler(handler.Private))
	// r.Handle("/auth", auth.GetTokenHandler)
	// ログインのやつ
	// r.HandleFunc("/login", auth.GetTokenHandler).Methods("POST")
  r.HandleFunc("/login", auth.GetTokenHandler).Methods("POST")
  
	// 認証したあとに見るやつ
	// r.HandleFunc("/pointhistory", auth.JwtMiddleware.Handler(handler.PointHistory)).Methods("POST")

	// デバッグ用のやつ
	r.Handle("/public2", auth.Public2).Methods("POST")

	// 8080番ポートでリッスンするよ〜。
	if err := http.ListenAndServe(":1234", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}
