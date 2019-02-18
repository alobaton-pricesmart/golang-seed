package main

import (
	"github.com/alobaton/golang-seed/apps/masters/pkg/models"
	"github.com/alobaton/golang-seed/apps/settings/pkg/config"
)

func main() {
	if err := config.ParseSettings(); err != nil {
		log.Fatal(err)
	}

	if err := models.ConnectRepo(); err != nil {
		log.Fatal(err)
	}
}
