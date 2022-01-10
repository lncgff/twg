package main

import (
	"net/http"

	"github.com/lncgff/twg/app"
)

func main() {
	http.ListenAndServe(":3000", &app.Server{})
}
