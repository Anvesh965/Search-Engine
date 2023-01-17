package Routes

import (
	"log"
	"net/http"
	. "search-engine/pkg/Models"
	. "search-engine/pkg/Controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
func HandleRoutes(router *mux.Router) {

	HandleVersion1(router)
	//HandleVersion2(router)

}
func HandleVersion1(router *mux.Router) {
	var api1 = router.PathPrefix("/v1").Subrouter()
	api1.HandleFunc("/", ServerHome).Methods("GET")
	api1.HandleFunc("/savepage", CreateWebPage).Methods("POST")
	api1.HandleFunc("/querypages", QueryHandle).Methods("GET")
	api1.HandleFunc("/allpages", GetAllWebPages).Methods("GET")
}

func HandleVersion2(router *mux.Router) {
	//var api2 = router.PathPrefix("/v2").Subrouter()
	log.Println("HandleVersion2:: Called")

}
func StartServer() {
	router := GetRouter()
	HandleRoutes(router)
	log.Println("Listeninig on port 4000.......")

	//seeding
	WebPages = append(WebPages, Webpage{Id: primitive.NilObjectID, Title: "page-demo", Keywords: []string{"wrd1", "wrd2", "wrd3", "wrd4"}})
	log.Fatal(http.ListenAndServe(":4000", router))
}
