package models

import (
	jwt "github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Claim es la estrutura usada para procesar el JWT */

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"` /*`bson: "_id" json:"_id,omitempty"`*/
	jwt.StandardClaims
}
