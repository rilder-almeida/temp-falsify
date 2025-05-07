package main

import (
	"github.com/arquivei/go-app"
)

var (
	version = "v0.0.0-dev"
	config  struct {
		app.Config
		HTTP struct {
			Port string `default:"8000"`
		}
	}
)
