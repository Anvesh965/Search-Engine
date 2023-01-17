package DatabaseConn

import (
	"context"
	"fmt"
	"log"
	"search-engine/pkg/Models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const URI = "mongodb+srv://anvesh9652:12345@cluster0.n763atd.mongodb.net/test"
const URI = "mongodb://mongo-container:27017"
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

func Search(keys []string) []Models.Webpage {

	var orOptions []bson.M

	for _, key := range keys {
		each := bson.M{"keywords": bson.M{"$regex": "\\b" + key + "\\b", "$options": "i"}}
		orOptions = append(orOptions, each)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// filter := bson.M{"keywords": bson.M{"$in": keys}}
	filter := bson.M{"$or": orOptions}
	cursor, err := collPtr.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	// var matchedRecords []primitive.M
	var matchedRecords []Models.Webpage
	for cursor.Next(ctx) {

		// var webpage bson.M
		var webpage Models.Webpage
		if err := cursor.Decode(&webpage); err != nil {
			log.Fatal(err)
		}

		matchedRecords = append(matchedRecords, webpage)
	}

	return matchedRecords

}

func AllPagesInCollection() []Models.Webpage {
	ctx, cance := context.WithTimeout(context.Background(), 5*time.Second)
	defer cance()

	filter := bson.M{}

	cursor, err := collPtr.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	var allpages []Models.Webpage

	for cursor.Next(ctx) {
		var webpage Models.Webpage
		if err := cursor.Decode(&webpage); err != nil {
			log.Fatal(err)
		}
		allpages = append(allpages, webpage)
	}
	return allpages

}
