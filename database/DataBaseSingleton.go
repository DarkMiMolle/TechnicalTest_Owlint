package database

import (
	"context"
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
	return client.Database("technical_test_owlint")
}()

func DataBase() *mongo.Database {
	return db
}

// TODO: filter function maker or something more intuitive/helpful to use
