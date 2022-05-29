package main

import (
	"log"
	"github.com/Zeldoso17/API-GO-REDSOCIAL/bd"
	"github.com/Zeldoso17/API-GO-REDSOCIAL/handlers"
)

func main() {
	if !bd.ConnectionStatus() {
		log.Fatal("No Database Connection")
		return
	}
	handlers.Managers()
}