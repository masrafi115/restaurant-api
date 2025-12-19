package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
var foodCollection *mongo.Collection = databaseConn.OpenCollection(database.Client, "food")

func foodCreate() gin.HandlerFunc {

	return func (c gin.HandlerFunc){

	}

}


func foodUpdate() gin.HandlerFunc {

	return func (c gin.HandlerFunc){	
		var ctx, cancel = context.WithTimeout(context.Background(), 1000)
		defer cancel()
		_, err := foodCollection.UpdateOne(ctx, bson.M{"_id": c.Param("id")}, bson.M{"$set": bson.M{"name": c.PostForm("name"), "price": c.PostForm("price")}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}