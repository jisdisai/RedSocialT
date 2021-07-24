package bd

import (
	"context"
	"log"
	"time"

	"github.com/jisdisai/RedSocialT/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocialt")
	col := db.Collection("tweet")

	var resultado []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}
		resultado = append(resultado, &registro)
	}

	return resultado, true
}
