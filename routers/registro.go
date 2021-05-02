package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jisdisai/RedSocialT/bd"
	"github.com/jisdisai/RedSocialT/models"
)

/* Registro es la funcion para crear en la BD el Registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los Datos Recibidos "+err.Error(), 400)
		return

	}

	/* Validacion del correo */

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido ", 400)
		return
	}

	/* validacion de la contraseña*/

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos de 6 Caracteres ", 400)
		return
	}

	/* -validacion contra los Datos detecta que el usuario o email ya existe en la BD */
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya Existe un usuario  con ese email  ", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio  un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado insertar el usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
