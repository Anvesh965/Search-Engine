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

type CollectionHelper interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}

// type API struct {
// 	coll CollectionHelper
// 	ctx  context.Context
// }

type RealDBFunction struct {
	collPtr CollectionHelper
}

func NewDBFunctions(ch CollectionHelper) DBFunctions {
	return &RealDBFunction{collPtr: ch}
}

func Start() CollectionHelper {

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

	myCollPtr := client.Database(Config.Database.DBName).Collection(Config.Database.Collection)

	return myCollPtr
}

func (rdb *RealDBFunction) UploadWebpage(webpage *Models.Webpage) (*mongo.InsertOneResult, error) {

	webpage.Id = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// adding document to webpages colletion
	result, err := rdb.collPtr.InsertOne(ctx, webpage)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Record inserted with id:", result.InsertedID)

	return result, err

}

func (rdb *RealDBFunction) Search(keys []string) ([]Models.Webpage, error) {

	var orOptions []bson.M

	for _, key := range keys {
		each := bson.M{"keywords": bson.M{"$regex": "\\b" + key + "\\b", "$options": "i"}}
		orOptions = append(orOptions, each)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// filter := bson.M{"keywords": bson.M{"$in": keys}}
	filter := bson.M{"$or": orOptions}
	cursor, err := rdb.collPtr.Find(ctx, filter)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	matchedRecords := []Models.Webpage{}

	for cursor.Next(ctx) {

		var webpage Models.Webpage
		if err := cursor.Decode(&webpage); err != nil {
			log.Println(err)
			return matchedRecords, err
		}

		matchedRecords = append(matchedRecords, webpage)
	}

	return matchedRecords, err

}

func (rdb *RealDBFunction) AllPagesInCollection() ([]Models.Webpage, error) {
	ctx, cance := context.WithTimeout(context.Background(), 5*time.Second)
	defer cance()

	filter := bson.M{}

	cursor, err := rdb.collPtr.Find(ctx, filter)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	allpages := []Models.Webpage{}

	for cursor.Next(ctx) {
		var webpage Models.Webpage
		if err := cursor.Decode(&webpage); err != nil {
			log.Println(err)
			return allpages, err
		}
		allpages = append(allpages, webpage)
	}
	return allpages, err

}
