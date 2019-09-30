package main

import (
	"github.com/voltento/shorturl/internal/app/urlshorter"
	"github.com/voltento/shorturl/pkg/log"
)

func main() {
	app, err := urlshorter.NewApp()
	if err == nil {
		err = app.Start()
	}
	if err != nil {
		log.Fatalw("Got fatal error on start applications", "Error", err.Error())
	}

	log.Infof("Application started.")

}
