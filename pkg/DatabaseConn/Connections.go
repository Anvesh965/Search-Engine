package DatabaseConn

import (
	"context"
	"fmt"
	"log"
	"search-engine/pkg/Models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const URI = "mongodb+srv://anvesh9652:12345@cluster0.n763atd.mongodb.net/test"
const URI = "mongodb://localhost:27017"
const collectionName = "webpages"
const databaseName = "Search-Engine"

var collPtr *mongo.Collection

func Start() {

	// Connecting to MongoDB database
	myOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), myOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Connected to MongoDB")

	// creating database and collections
	collPtr = client.Database(databaseName).Collection(collectionName)

}

func UploadWebpage(webpage Models.Webpage) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// adding document to webpages colletion
	result, err := collPtr.InsertOne(ctx, webpage)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Record inserted with id: ", result.InsertedID)

}

func Search(keys []string) []primitive.M {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"keywords": bson.M{"$in": keys}}

	cursor, err := collPtr.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	var matchedRecords []primitive.M
	for cursor.Next(ctx) {

		var webpage bson.M
		if err := cursor.Decode(&webpage); err != nil {
			log.Fatal(err)
		}

		matchedRecords = append(matchedRecords, webpage)
	}
	return matchedRecords

}
