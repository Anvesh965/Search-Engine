package testing

import (
	"encoding/json"
	"log"
	"net/http"
	"search-engine/pkg/DatabaseConn"
	"search-engine/pkg/Models"

	"github.com/gorilla/mux"
)

func Initialize() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", HandleUpload).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	var webpage Models.Webpage

	json.NewDecoder(r.Body).Decode(&webpage)

	DatabaseConn.UploadWebpage(&webpage)

	json.NewEncoder(w).Encode(webpage)
}
