package middlew

import (
	"net/http"

	"github.com/jisdisai/RedSocialT/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
