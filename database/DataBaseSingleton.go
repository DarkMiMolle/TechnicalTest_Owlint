package database

import (
	"context"
	"fmt"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// db should be initialized at the beginning and only once. It represents the database
var db *mongo.Database = func() *mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	util.PanicErr(err)

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	util.PanicErr(err)

	// get collection as ref
	db := client.Database("technical_test_owlint")
	fmt.Println("database created")
	return db
}()

// DataBase actually return the used Collection since we use only one collection
func DataBase() *mongo.Collection {
	return db.Collection("Comment")
}

// TODO: filter function maker or something more intuitive/helpful to use
