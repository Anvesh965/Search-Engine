package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	. "search-engine/pkg/Models"
	"strings"
)

var WebPages []Webpage //assuming as DB replace with actual db

type Ranks struct {
	PageName string `json:"title"`
	Value    int    `json:"rank"`
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	log.Println("ServerHome ::Called")
	w.Write([]byte("<h1>Welcome to Ranking websites Rest-Api</h1>"))
}
func GetAllWebPages(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllCourses ::Called")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(WebPages)
	//handle in v1
}
func CreateWebPage(w http.ResponseWriter, r *http.Request) {
	//handle in v1
	log.Println("CreateWebPage :: Called")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter some data")
		return
	}
	var webpage Webpage

	_ = json.NewDecoder(r.Body).Decode(&webpage)

	//do some checking of data valid or not
	WebPages = append(WebPages, webpage)
	json.NewEncoder(w).Encode(webpage)

}
func QueryHandle(w http.ResponseWriter, r *http.Request) {
	//handle in v1
	log.Println("QueryHandle :: Called")
	w.Header().Set("content-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter some data")
		return
	}
	var webpage Webpage
	_ = json.NewDecoder(r.Body).Decode(&webpage)
	PageRanks := GeneratePageRanks(webpage.Keywords)
	json.NewEncoder(w).Encode(PageRanks)

}
func GeneratePageRanks(params []string) []Ranks {

	log.Println("GeneratePageRanks")

	var PageRank []Ranks
	for i := 0; i < len(WebPages); i++ {
		score := GetScore(WebPages[i].Keywords, params)
		PageRank = append(PageRank, Ranks{WebPages[i].Title, score})
	}
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
