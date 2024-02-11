package main

import (
	"log"
	"os"
)

func msgDie(msg string) {
	log.Println("ERR:", msg)
	os.Exit(1)
}
