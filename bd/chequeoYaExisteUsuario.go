package bd

import (
	"context"
	"time"

	"github.com/jisdisai/RedSocialT/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya existe en la bBD */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocialt")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
