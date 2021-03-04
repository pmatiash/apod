package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pmatiash/apod/db"
	"time"
)

type ImageOfTheDay struct {
	Date        string `json:"date"`
	Copyright   string `json:"copyright"`
	Explanation string `json:"explanation"`
	Hdurl       string `json:"hdurl"`
	MediaType   string `json:"media_type"`
	Title       string `json:"title"`
	Url         string `json:"url"`
}

type Stats struct {
	SingleImportStartDate string `json:"singleImportStartDate"`
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	fmt.Println("Start...")

	dbClient := db.InitClient()
	defer dbClient.Destruct()

	var stats Stats

	t := time.Now()
	today := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())

	// set start of single day import to avoid crossing with bulk import
	err := dbClient.FindOne("stats", bson.M{}, &stats)
	if err != nil {
		stats.SingleImportStartDate = today
		dbClient.Insert("stats", stats)
	}
	// @TODO: we need to update single record in stats if already exists

	var item ImageOfTheDay
	json.Unmarshal(getPost(), &item)
	dbClient.Insert("posts", item)

	fmt.Println("DONE!")
}
