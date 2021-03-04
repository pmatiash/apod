package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DbClient struct {
	Client  *mongo.Client
	Context context.Context
}

func InitClient() *DbClient {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &DbClient{client, ctx}
}

func (dbClient *DbClient) Destruct() {
	dbClient.Client.Disconnect(dbClient.Context)
}

func (dbClient *DbClient) Insert(collName string, data interface{}) {
	fmt.Println("Inserting data to DB...")

	insertResult, err := dbClient.GetCollection(collName).InsertOne(dbClient.Context, data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}

func (dbClient *DbClient) GetCollection(collName string) *mongo.Collection {
	return dbClient.Client.Database(os.Getenv("DB_NAME")).Collection(collName)
}

func (dbClient *DbClient) FindOne(collName string, condition interface{}, decodeTo interface{}) error {
	return dbClient.GetCollection(collName).FindOne(dbClient.Context, condition).Decode(decodeTo)
}
