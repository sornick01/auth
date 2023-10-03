package main

import (
	"log"

	"github.com/sornick01/auth/internal/app"
)

func main() {
	a := app.New()
	err := a.RunApp("9933")
	if err != nil {
		log.Fatal(err)
	}
}
