package services

import (
	"search-engine/pkg/Models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PageService interface {
	UploadWebpage(webpage *Models.Webpage) (*mongo.InsertOneResult, error)
	Search(keys []string) ([]Models.Webpage, error)
	AllPagesInCollection() ([]Models.Webpage, error)
}
