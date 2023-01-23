package Controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"search-engine/pkg/Models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

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
