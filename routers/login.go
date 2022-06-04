package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jisdisai/RedSocialT/bd"
	"github.com/jisdisai/RedSocialT/jwt"
	"github.com/jisdisai/RedSocialT/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El Usuario es Requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == !existe {
		http.Error(w, "Usuario y/o Contraseña invalidas", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un Error al Intentar Generar el TOken Correspondiente "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/* Grabar una Cookies*/
	/* campo fecha Expoirar */
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
