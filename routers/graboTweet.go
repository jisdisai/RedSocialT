package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jisdisai/RedSocialT/bd"
	"github.com/jisdisai/RedSocialT/models"
)

/* GraboTweet permite grabar el tweet en la BD */
func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	Registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(Registro)
	if err != nil {
		http.Error(w, "Ocurrio un Error al Intentar insertar el registro intente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado innsertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
