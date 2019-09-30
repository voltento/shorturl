package main

import (
	"shorturl/internal/app/urlshorter"
	"shorturl/pkg/log"
)

func main() {
	app, err := &urlshorter.NewApp()
	if err == nil {
		err = app.Start()
	}
	if err != nil {
		log.Fatalf("Got fatal error on start applications. Error: %v"
		err.Error())
	}

}
