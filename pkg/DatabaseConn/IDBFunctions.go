package DatabaseConn

import "search-engine/pkg/Models"

type DBFunctions interface {
	Start()
	UploadWebpage(webpage *Models.Webpage)
	Search(keys []string) []Models.Webpage
	AllPagesInCollection() []Models.Webpage
}
