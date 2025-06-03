package main

import (
	"log"
	"web-authentication/src"
)

func main() {
	log.SetFlags(0)
	if err := src.StartTCPServer(); err != nil {
		log.Fatal(err)
	}
	log.Print("Program has been successfully executed")
}
