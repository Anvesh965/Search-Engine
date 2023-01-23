package Routes

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"search-engine/pkg/Models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	mocks "search-engine/mocks/pkg/DatabaseConn"
	. "search-engine/pkg/Controllers"
)

func TestGetAllWebPages(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	//rdb := &DatabaseConn.RealDBFunction{}

	//webpage := Models.Webpage{Id: primitive.NilObjectID, Title: "page-1", Keywords: []string{"wrd1", "wrd2"}}
	// Create mock for DB function
	rdbMock := mocks.NewDBFunctions(t)
	rdbMock.On("AllPagesInCollection").Return([]Models.Webpage{})
	router.GET("/allpages", func(c *gin.Context) {
		GetAllWebPages(c, rdbMock)
	})

	// Replace real DB function with mock

	// Create request to endpoint
	//jsonInput, _ := json.Marshal(webpage)
	//input := `{"title":"page-1","keywords":["wrd1","wrd2"]}`
	req, err := http.NewRequest("GET", "/allpages", nil)
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusOK, resp.Code)
	// Check if mock function was called
	rdbMock.AssertExpectations(t)
}
func TestCreateWebPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	rdbMock := mocks.NewDBFunctions(t)
	rdbMock.On("UploadWebpage", &Models.Webpage{}).Return(Models.Webpage{})
	router.POST("/savepage", func(c *gin.Context) {
		CreateWebPage(c, rdbMock)
	})

	// Replace real DB function with mock

	// Create request to endpoint
	//jsonInput, _ := json.Marshal(webpage)
	input := `{}`
	req, err := http.NewRequest("POST", "/savepage", bytes.NewBuffer([]byte(input)))
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusPartialContent, resp.Code)
	// Check if mock function was called
	rdbMock.AssertExpectations(t)
}
