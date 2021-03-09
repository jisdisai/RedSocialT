package main

import (
	"log"

	"github.com/jisdisai/RedSocialT/bd"
	"github.com/jisdisai/RedSocialT/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD ")
		return
	}
	handlers.Manejadores()
}
