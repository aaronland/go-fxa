package main

import (
	"flag"
	"github.com/aaronland/go-fxa"
	"log"
)

func main() {

	flag.Parse()

	cl, err := fxa.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	email := "e@example.com"
	pswd := "example"

	err = cl.Login(email, pswd)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(cl)
}
