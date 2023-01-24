package DatabaseConn

import (
	"errors"
	mocks "search-engine/mocks/pkg/DatabaseConn"
	"search-engine/pkg/Models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestUploadWebpageResultCheck(t *testing.T) {

	// mockColl := mocks.NewCollectionHelper(t)
	mockColl := mocks.NewCollectionHelper(t)

	currId := primitive.NewObjectID()
	testpage := Models.Webpage{Id: currId, Title: "p1", Keywords: []string{"anvesh", "gali"}}

	mockColl.On("InsertOne", mock.Anything, &testpage).Return(&mongo.InsertOneResult{}, nil)

	mockApiColl := NewDBFunctions(mockColl)
	result, err := mockApiColl.UploadWebpage(&testpage)

	assert.Nil(t, err)
	assert.Equal(t, &mongo.InsertOneResult{}, result)

}

func TestUploadWebpageErrorCheck(t *testing.T) {

	mockColl := mocks.NewCollectionHelper(t)

	currId := primitive.NewObjectID()
	testpage := Models.Webpage{Id: currId, Title: "p1", Keywords: []string{"anvesh", "gali"}}

	mockColl.On("InsertOne", mock.Anything, &testpage).Return(nil, errors.New("error-while-inserting"))

	mockApiColl := NewDBFunctions(mockColl)
	_, err := mockApiColl.UploadWebpage(&testpage)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "error-while-inserting")
}

// here main &{ 0xc00013a000 <nil> 1 0xc0001621c0 0xc00069e750 <nil>}
// here main &{ <nil> <nil> 0 <nil> <nil> <nil>}

func TestAllCollectionResultCheck(t *testing.T) {

	mockColl := mocks.NewCollectionHelper(t)
	currID := primitive.NewObjectID()
	docs := []interface{}{bson.M{"_id": currID, "title": "anvesh", "keywords": []string{"a", "b"}}}
	dummy, _ := mongo.NewCursorFromDocuments(docs, nil, nil)
	mockCursor1 := dummy
	// mockCursor2 := *dummy

	expectedResult := []Models.Webpage{
		Models.Webpage{
			Id:       currID,
			Title:    "anvesh",
			Keywords: []string{"a", "b"},
		},
	}

	mockColl.On("Find", mock.Anything, bson.M{}).Return(mockCursor1, nil)

	mockApi := NewDBFunctions(mockColl)
	result, err := mockApi.AllPagesInCollection()

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestAllCollectionErrorCheck(t *testing.T) {

	mockColl := mocks.NewCollectionHelper(t)

	mockColl.On("Find", mock.Anything, bson.M{}).Return(nil, errors.New("error while fetching records"))

	mockApi := NewDBFunctions(mockColl)
	_, err := mockApi.AllPagesInCollection()

	assert.NotNil(t, err)
	assert.EqualError(t, err, "error while fetching records")

}

func TestSearchingResultCheck(t *testing.T) {

	keys := []string{"bmw", "microsoft", "paris"}
	mockColl := mocks.NewCollectionHelper(t)
	currID := primitive.NewObjectID()
	docs := []interface{}{
		bson.M{"_id": currID, "title": "twitter", "keywords": []string{"tesla", "ford", "BMW"}},
	}
	dummy, _ := mongo.NewCursorFromDocuments(docs, nil, nil)
	mockCursor1 := dummy

	expectedResult := []Models.Webpage{
		Models.Webpage{
			Id:       currID,
			Title:    "twitter",
			Keywords: []string{"tesla", "ford", "BMW"},
		},
	}

	mockColl.On("Find", mock.Anything, mock.Anything).Return(mockCursor1, nil)

	mockApi := NewDBFunctions(mockColl)
	result, err := mockApi.Search(keys)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestSearchingErrorCheck(t *testing.T) {

	mockColl := mocks.NewCollectionHelper(t)
	keys := []string{"tesla", "ford"}

	mockColl.On("Find", mock.Anything, mock.Anything).Return(nil, errors.New("error while fetching records"))

	mockApi := NewDBFunctions(mockColl)
	_, err := mockApi.Search(keys)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "error while fetching records")

}
