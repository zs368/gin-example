package main

import (
	"net/http"

	"github.com/zs368/gin-example/internal/pkg/routing"
)

func main() {
	r := routing.NewRouter()
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	s.ListenAndServe()
}
