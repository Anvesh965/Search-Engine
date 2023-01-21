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

/*func TestPostMethod(t *testing.T) {
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
}*/
