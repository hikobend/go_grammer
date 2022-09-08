package controllers

import (
	"net/http"

	"example.com/m/grammer14/config"
)

func StarMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
