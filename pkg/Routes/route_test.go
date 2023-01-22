package Routes

/*
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"search-engine/pkg/DatabaseConn"
	"search-engine/pkg/Models"
	"testing"

	mocks "search-engine/mocks/pkg/DatabaseConn"
	. "search-engine/pkg/Controllers"


	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert" // import mock package
	"github.com/stretchr/testify/mock"
)

func TestCreateWebPage(t *testing.T) {
	// create new gin router
	router := gin.Default()

	dbMock := new(mocks.DBFunctions)
	dbMock.On("UploadWebpage", &Models.Webpage{}).Return(mock.Anything)
	// define endpoint
	router.POST("/v1/savepage", func(c *gin.Context) {
		CreateWebPage(c, &DatabaseConn.RealDBFunction{})
	})

	// create mock for DB function

	// replace real DB function with mock
	router.Use(func(c *gin.Context) {
		c.Set("rdb", &dbMock)
	})

	// create request to endpoint
	input := `{"title":"test-page","keywords":["test","content"]}`
	req, _ := http.NewRequest("POST", "/v1/savepage", bytes.NewBufferString(input))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// check response
	assert.Equal(t, http.StatusCreated, resp.Code)

	var webpage Models.Webpage
	err := json.Unmarshal(resp.Body.Bytes(), &webpage)
	if err != nil {
		t.Errorf("Error while Unmarshalling response")
	}
	// check if mock function was called
	dbMock.AssertExpectations(t)
}
*/
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"search-engine/pkg/DatabaseConn"
	"search-engine/pkg/Models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	mocks "search-engine/mocks/pkg/DatabaseConn"
	. "search-engine/pkg/Controllers"
)

func TestCreateWebPage(t *testing.T) {
	router := gin.Default()
	rdb := &DatabaseConn.RealDBFunction{}
	router.POST("/savepage", func(c *gin.Context) {
		CreateWebPage(c, rdb)
	})

	// Create mock for DB function
	rdbMock := new(mocks.DBFunctions)
	rdbMock.On("UploadWebpage", &Models.Webpage{}).Return()

	// Replace real DB function with mock
	router.Use(func(c *gin.Context) {
		c.Set("rdb", rdbMock)
	})

	// Create request to endpoint
	input := `{"title":"Page-1","keywords":["wrd1","wrd2"]}`
	req, _ := http.NewRequest("POST", "/savepage", bytes.NewBufferString(input))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusCreated, resp.Code)

	// Check if mock function was called
	rdbMock.AssertExpectations(t)
}
