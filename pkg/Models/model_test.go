package Models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Check(t *testing.T) {

	testcases := []struct {
		webpage  Webpage
		expected bool
	}{
		{webpage: Webpage{Title: "", Keywords: []string{}}, expected: true},
		{webpage: Webpage{Title: "", Keywords: []string{"wrd"}}, expected: true},
		{webpage: Webpage{Title: "Page", Keywords: []string{}}, expected: true},
		{webpage: Webpage{Title: "    ", Keywords: []string{"one"}}, expected: true},
		{webpage: Webpage{Title: "Title", Keywords: []string{"wrd1", "wrd2"}}, expected: false},
	}
	for _, e := range testcases {
		actual := e.webpage.Check()
		assert.Equal(t, e.expected, actual)
	}
}
func TestModifyKeyLength(t *testing.T) {
	testcases := []struct {
		webpage  Webpage
		expected Webpage
	}{
		{webpage: Webpage{Title: "page", Keywords: []string{}}, expected: Webpage{Title: "page", Keywords: []string{}}},
		{webpage: Webpage{Title: "page", Keywords: []string{"wrd1", "wrd2"}}, expected: Webpage{Title: "page", Keywords: []string{"wrd1", "wrd2"}}},
		{webpage: Webpage{Title: "", Keywords: []string{}}, expected: Webpage{Title: "", Keywords: []string{}}},
		{webpage: Webpage{Title: "", Keywords: []string{"1", "2", "3", "4", "5"}}, expected: Webpage{Title: "", Keywords: []string{"1", "2", "3", "4", "5"}}},
		{webpage: Webpage{Title: "page", Keywords: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}}, expected: Webpage{Title: "page", Keywords: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}}},
	}
	for _, e := range testcases {
		e.webpage.ModifyKeysLength()
		assert.Equal(t, e.expected, e.webpage)
	}
}
