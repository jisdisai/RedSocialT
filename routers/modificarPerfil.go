package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jisdisai/RedSocialT/bd"
	"github.com/jisdisai/RedSocialT/models"
)

/* ModificarPerfil modifica el perfil de usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un Error al Intentar Modificar el Registro. Intente Nuevamente"+err.Error(), 400)
		return
	}

	if status == !status {
		http.Error(w, "No se ha Logrado modificar el usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
