package main

import (
	"log"
	"web-authentication/src"
)

func main() {
	log.SetFlags(0)
	src.HTTPServer()
	log.Print("Program has been successfully executed")
}
