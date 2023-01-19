package Controllers

import (
	"net/http"
	. "search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Models"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ranks struct {
	PageName string `json:"title"`
	Value    int    `json:"rank"`
}

func StatusCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Status": "Server Running",
	})
}
func ServerHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Search-Engine-Rest-Api",
		"Version": "v1",
	})
}
func GetAllWebPages(c *gin.Context) {
	allPages := AllPagesInCollection()
	c.IndentedJSON(http.StatusOK, allPages)
}
func CreateWebPage(c *gin.Context) {
	var webpage Webpage

	if err := c.BindJSON(&webpage); err != nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{
			"message": "Enter a valid Content",
		})
	}
	webpage.Id = primitive.NewObjectID()
	UploadWebpage(&webpage)
	c.IndentedJSON(http.StatusCreated, webpage)
}
func QueryHandle(c *gin.Context) {
	var webpage Webpage
	if err := c.BindJSON(&webpage); err != nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{
			"message": "Enter a valid Content",
		})
	}
	PageRanks := GeneratePageRanks(webpage.Keywords)
	c.IndentedJSON(http.StatusOK, PageRanks)
}
func GeneratePageRanks(params []string) []Ranks {
	WebPages := Search(params)
	var PageRank []Ranks
	for _, webpage := range WebPages {
		score := GetScore(webpage.Keywords, params)
		PageRank = append(PageRank, Ranks{webpage.Title, score})
	}
	sort.Slice(PageRank, func(i int, j int) bool {
		return PageRank[i].Value > PageRank[j].Value
	})
	size := min(5, len(WebPages))
	PageRank = PageRank[:size]

	return PageRank
}

func GetScore(Keywords, params []string) int {
	ans := 0

	for i := 0; i < len(Keywords); i++ {
		for j := 0; j < len(params); j++ {
			if strings.EqualFold(Keywords[i], params[j]) {
				ans = ans + (10-i)*(10-j)
				break
			}
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
