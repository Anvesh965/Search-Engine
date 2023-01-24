package Controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"search-engine/pkg/Models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mocks "search-engine/mocks/pkg/DatabaseConn"
)

func TestGetAllWebPages(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	rdbMock := mocks.NewDBFunctions(t)
	rdbMock.On("AllPagesInCollection").Return([]Models.Webpage{})
	router.GET("/allpages", func(c *gin.Context) {
		GetAllWebPages(c, rdbMock)
	})

	req, err := http.NewRequest("GET", "/allpages", nil)
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusOK, resp.Code)
	// Check if mock function was called
	//rdbMock.AssertExpectations(t)

}
func TestCreateWebPage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	rdbMock := mocks.NewDBFunctions(t)

	//To Test 206 status code when body is there but required data is not mentioned
	//rdbMock.On("UploadWebpage").Return([]Models.Webpage{})
	router.POST("/savepage", func(c *gin.Context) {
		CreateWebPage(c, rdbMock)
	})
	input := `{}`
	req, err := http.NewRequest("POST", "/savepage", bytes.NewBuffer([]byte(input)))
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusPartialContent, resp.Code)

	//test for body nil sttaus code 400
	input = `{"Title":"123","Keys":["123","234"]}`
	req, err = http.NewRequest("POST", "/savepage", nil)
	if err != nil {
		log.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//check when BindJson error occurs statuscode 400
	input = `"user":"name","password":"123","number":123,"mail":"email"`
	req, err = http.NewRequest("POST", "/savepage", bytes.NewBuffer([]byte(input)))
	if err != nil {
		log.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	rdbMock.On("UploadWebpage", mock.Anything).Return()
	input = `"title":"page","keywords":["wrd1"]`
	webpage := Models.Webpage{Title: "page", Keywords: []string{"wrd1"}}

	jsonInput, _ := json.Marshal(webpage)
	req, err = http.NewRequest("POST", "/savepage", bytes.NewBuffer(jsonInput))
	if err != nil {
		log.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.NotEmpty(t, resp.Body)
}

func TestQueryHandle(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	rdbMock := mocks.NewDBFunctions(t)
	router.POST("/querypages", func(c *gin.Context) {
		QueryHandle(c, rdbMock)
	})

	req, err := http.NewRequest("POST", "/querypages", nil)
	if err != nil {
		log.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.NotNil(t, resp.Body)

	rdbMock.On("Search", mock.Anything).Return([]Models.Webpage{
		{Id: primitive.NewObjectID(), Title: "page-1", Keywords: []string{"Ford", "Review", "Car"}},
		{Id: primitive.NewObjectID(), Title: "page-2", Keywords: []string{"BMW", "Gin", "", "GO", "Car"}},
		{Id: primitive.NewObjectID(), Title: "page-3", Keywords: []string{"Car", "Toyota", "Mock"}},
		{Id: primitive.NewObjectID(), Title: "page-4", Keywords: []string{"KIA", "Car"}},
		{Id: primitive.NewObjectID(), Title: "page-5", Keywords: []string{"", "Review", "", "Car"}},
	})
	keys := Models.Keys{Keywords: []string{"Car"}}
	inputjson, _ := json.Marshal(keys)
	req, err = http.NewRequest("POST", "/querypages", bytes.NewBuffer([]byte(inputjson)))
	if err != nil {
		log.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.NotNil(t, resp.Body)

}

func TestStatusCheck(t *testing.T) {
	mockResponse := "\"Server Running\""
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/", StatusCheck)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestHomepageHandler(t *testing.T) {

	mockResponse := []byte(`{
    "message": "Search-Engine-Rest-Api",
    "version": "v1"
}`)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/v1/", ServerHome)
	req, _ := http.NewRequest("GET", "/v1/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGenerateRanks(t *testing.T) {
	rdbMock := mocks.NewDBFunctions(t)
	rdbMock.On("Search", mock.Anything).Return([]Models.Webpage{
		{Id: primitive.NewObjectID(), Title: "page-1", Keywords: []string{"Ford", "Review", "Car"}},
		{Id: primitive.NewObjectID(), Title: "page-2", Keywords: []string{"BMW", "Gin", "", "GO", "Car"}},
		{Id: primitive.NewObjectID(), Title: "page-3", Keywords: []string{"Car", "Toyota", "Mock"}},
		{Id: primitive.NewObjectID(), Title: "page-4", Keywords: []string{"KIA", "Car"}},
		{Id: primitive.NewObjectID(), Title: "page-5", Keywords: []string{"", "Review", "", "Car"}},
	})
	params := []string{"Car"}
	expected := []Ranks{
		{PageName: "page-3", Value: 100},
		{PageName: "page-4", Value: 90},
		{PageName: "page-1", Value: 80},
		{PageName: "page-5", Value: 70},
		{PageName: "page-2", Value: 60},
	}
	actual := GeneratePageRanks(params, rdbMock)
	assert.Equal(t, expected, actual)
}
func TestGetScore(t *testing.T) {
	testcases := []struct {
		keywords []string
		params   []string
		expected int
	}{
		{[]string{"one", "two", "three"}, []string{"one", "three"}, 172},
		{[]string{"1", "2", "3", "4"}, []string{"4", "3"}, 142},
		{[]string{}, []string{"4", "3"}, 0},
		{[]string{"1", "2", "3", "4"}, []string{}, 0},
		{[]string{"1", "5", "6"}, []string{"4", "3"}, 0},
	}
	for _, e := range testcases {
		actual := GetScore(e.keywords, e.params)
		assert.Equal(t, e.expected, actual)
	}
}
func TestMin(t *testing.T) {
	testcases := []struct {
		a        int
		b        int
		expected int
	}{
		{10, 10, 10},
		{-10, -12., -12},
		{10, 11, 10},
	}
	for _, e := range testcases {
		actual := min(e.a, e.b)
		assert.Equal(t, e.expected, actual)
	}
}
