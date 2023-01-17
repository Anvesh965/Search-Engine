package Controllers

import (
	"encoding/json"
	"log"
	"net/http"
	. "search-engine/pkg/Models"
)

var WebPages []Webpage

type Ranks struct {
	PageName string `json:"title"`
	Value    int    `json:"rank"`
}

var PageRanks []Ranks

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
	//var params []string

	//result := GeneratePageRanks(params)

	//log.Println("Request data:", params)
	//json.NewEncoder(w).Encode(result)
	json.NewEncoder(w).Encode(r.Body)

}
func GeneratePageRanks(params []string) {

	log.Println("GeneratePageRanks")
}