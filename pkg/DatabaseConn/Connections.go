package DatabaseConn

import (
	"context"
	"log"
	"search-engine/pkg/Models"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	. "search-engine/cmd/config"
)

var collPtr *mongo.Collection

type RealDBFunction struct {
}

func (rdb *RealDBFunction) Start() {

	// Connecting to MongoDB database

	URI := Config.Database.Protocol + "://" + Config.Database.Host + ":" + strconv.Itoa(Config.Database.Port)
	log.Println("URI:", URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	myOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, myOptions)

	if err != nil {
		log.Fatal(err)
	}

	// checking whether connection was successful or not.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Connected to MongoDB")

	// creating database and collections

	collPtr = client.Database(Config.Database.DBName).Collection(Config.Database.Collection)
}

func (rdb *RealDBFunction) UploadWebpage(webpage *Models.Webpage) {

	webpage.Id = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// adding document to webpages colletion
	result, err := collPtr.InsertOne(ctx, webpage)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Record inserted with id:", result.InsertedID)

}

func (rdb *RealDBFunction) Search(keys []string) []Models.Webpage {

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

func (rdb *RealDBFunction) AllPagesInCollection() []Models.Webpage {
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
