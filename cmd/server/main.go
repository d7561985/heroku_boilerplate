package main

import (
	"github.com/d7561985/heroku_boilerplate/pkg/config"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":"+config.V.Port, http.HandlerFunc(func(r http.ResponseWriter, _ *http.Request) {
		_, _ = r.Write([]byte{'O', 'K'})
	})); err != nil {
		panic(err)
	}
}
