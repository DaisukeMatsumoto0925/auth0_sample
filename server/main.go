package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/DaisukeMatsumoto0925/auth0_sample/handlers/v1"
	"github.com/DaisukeMatsumoto0925/auth0_sample/handlers/v1/users/me"
	"github.com/DaisukeMatsumoto0925/auth0_sample/middlewares/auth0"
	"github.com/rs/cors"
)

const (
	port     = 8000
	domain   = ""
	clientID = ""
)

func main() {
	jwks, err := auth0.FetchJWKS(domain)
	if err != nil {
		log.Fatal(err)
	}

	jwtMiddleware, err := auth0.NewMiddleware(domain, clientID, jwks)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	// /v1へのリクエストが来た場合のハンドラを追加
	mux.HandleFunc("/v1", v1.HandleIndex)
	mux.Handle("/v1/users/me", auth0.UseJWT(http.HandlerFunc(me.HandleIndex)))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	wrappedMux := auth0.WithJWTMiddleware(jwtMiddleware)(mux)
	wrappedMux = c.Handler(wrappedMux)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Listeng on %s", addr)
	// localhost:8000 でサーバーを立ち上げる
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
