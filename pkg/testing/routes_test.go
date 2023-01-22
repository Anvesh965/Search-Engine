package testing

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	. "search-engine/pkg/Controllers"
	. "search-engine/pkg/Routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func GetTestGinContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}
func TestStatusCheck(t *testing.T) {
	mockResponse := "\"Server Running\""
	gin.SetMode(gin.TestMode)
	r := GetRouter()
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
	r := GetRouter()
	r.GET("/v1/", ServerHome)
	req, _ := http.NewRequest("GET", "/v1/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(mockResponse), string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

/*
	func TestPostMethod(t *testing.T) {
		// Initialize a new Gin router
		router := gin.New()
		// Define the endpoint for the POST method
		router.POST("/endpoint", func(c *gin.Context) {
			// Retrieve the request body
			var jsonObject map[string]string
			c.Bind(&jsonObject)
			// Perform logic and set response
			c.JSON(200, gin.H{
				"data":    jsonObject,
				"message": "Success",
			})
		})
		// Create a request to send to the endpoint
		json := `{"key": "value"}`
		req, _ := http.NewRequest("POST", "/endpoint", bytes.NewBuffer([]byte(json)))
		req.Header.Set("Content-Type", "application/json")
		// Create a response recorder to inspect the response
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		// Check the status code and response body
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		expected := `{"message":"Success","data":{"key":"value"}}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	}

	func TestCreateWebPage(t *testing.T) {
		gin.SetMode(gin.TestMode)
		r := GetRouter()
		r.POST("/v1/savepage", CreateWebPage)
		//companyId := xid.New().String()
		webpage := Webpage{
			Title:    "demo-title",
			Keywords: []string{"wrd1", "wrd2"},
		}
		jsonValue, _ := json.Marshal(webpage)
		req, _ := http.NewRequest("POST", "/v1/savepage", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		log.Printf("returned %v\n", w.Body.String())
		assert.Equal(t, http.StatusCreated, w.Code)
	}

func router() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api_1 := router.Group("/v1")
	api_1.GET("/", ServerHome)
	api_1.GET("/allpages", GetAllWebPages)
	api_1.POST("/querypages", QueryHandle)
	api_1.POST("/savepage", CreateWebPage)
	return router
}
func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
func Test_ServerHome(t *testing.T) {
	expectedMessage := Controllers.Message{
		Msg:     "Search-Engine-Rest-Api",
		Version: "v1",
	}
	var responseMsg Controllers.Message
	writer := makeRequest("GET", "/v1/", nil)
	_ = json.NewDecoder(writer.Body).Decode(&responseMsg)

	assert.Equal(t, http.StatusOK, writer.Code)
	assert.NotEmpty(t, writer.Body)
	assert.Equal(t, expectedMessage, responseMsg)
}

func TestGetAllPages(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := GetRouter()
	r.GET("/v1/allpages", GetAllWebPages)
	req, _ := http.NewRequest("GET", "/v1/allpages", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
*/
