package main

import (
	"barter-server2/pkg/auth"
	"barter-server2/pkg/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/public", handler.Public)
	r.Handle("/private", auth.JwtMiddleware.Handler(handler.Private))
    // r.Handle("/auth", auth.GetTokenHandler)
    // ログインのやつ
    r.Handle("/login", auth.GetTokenHandler).Methods("POST")
    // デバッグ用のやつ
    r.Handle("/public2", handler.Public2).Methods("POST")

	// 8080番ポートでリッスン
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}
