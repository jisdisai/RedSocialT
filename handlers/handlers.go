package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo el puerto,  el Handler y pongo a escuchar el sevidor */
func Manejadores() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll()._Handler(router)
	log.Fatal(http.ListenAndServer(":"+PORT, handler))

}
