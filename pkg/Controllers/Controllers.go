package Controllers

import (
	"log"
	"net/http"
	. "search-engine/pkg/DatabaseConn"
	"search-engine/pkg/Models"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ranks struct {
	PageName string `json:"title"`
	Value    int    `json:"rank"`
}
type Message struct {
	Msg     string `json:"message"`
	Version string `json:"version",omitempty`
}

func StatusCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Server Running")
}

// @Summary get version data
// @ID get-version-details
// @Produce json
// @Success 200 {object} Message
// @Router /v1/ [get]
func ServerHome(c *gin.Context) {
	msg := Message{"Search-Engine-Rest-Api", "v1"}
	c.IndentedJSON(http.StatusOK, msg)
}

// @Summary get all pages in the webpages
// @ID get-all-webpages
// @Produce json
// @Success 200 {object} Models.Webpage
// @Router /v1/allpages [get]
func GetAllWebPages(c *gin.Context) {
	allPages := AllPagesInCollection()
	c.IndentedJSON(http.StatusOK, allPages)
}

// @Summary add a new webpage to the webpages list
// @ID create-web-page
// @Accept	json
// @Produce json
// @Param Page body Models.Page true "The input webpage details"
// @Success 201 {object} Models.Webpage
// @Failure 206 {object} Message
// @Failure 406 {object} Message
// @Router /v1/savepage [post]
func CreateWebPage(c *gin.Context) {
	var webpage Models.Webpage

	var msg Message
	msg.Msg = "Enter a valid Content"
	err := c.BindJSON(&webpage)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, msg)
		return
	}
	if webpage.Check() {
		c.IndentedJSON(http.StatusPartialContent, msg)
		return
	}
	webpage.ModifyKeysLength()
	webpage.Id = primitive.NewObjectID()
	UploadWebpage(&webpage)
	c.IndentedJSON(http.StatusCreated, webpage)
}

// @Summary get page ranks for keywords
// @ID get-page-ranks
// @Accept json
// @Produce json
// @Param data body Models.Keys true "The input Keyword list"
// @Success 200 {object} Ranks
// @Failure 404 {object} Message
// @Router /v1/querypages [post]
func QueryHandle(c *gin.Context) {
	var webpage Models.Webpage
	var msg Message
	msg.Msg = "Enter a valid Content"
	log.Println(c.Params)
	if err := c.BindJSON(&webpage); err != nil {
		c.IndentedJSON(http.StatusNoContent, msg)
		return
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
