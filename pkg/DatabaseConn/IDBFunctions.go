package DatabaseConn

import (
	"search-engine/pkg/Models"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBFunctions interface {
	UploadWebpage(webpage *Models.Webpage) (*mongo.InsertOneResult, error)
	Search(keys []string) ([]Models.Webpage, error)
	AllPagesInCollection() ([]Models.Webpage, error)
}
