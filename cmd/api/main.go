package main

import (
	"log"

	"github.com/rubenbupe/go-auth-server/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
