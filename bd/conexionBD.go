package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MpngoCN es el Objeto ala conexion a la BD*/
var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://salgadoisai:<password>@cluster0.sjyak.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD es la funcion que permite connetar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la BD")
	return client
}

/* ChequeoConnection es el ping a la base de datos */
func ChequeoConnection() int {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
