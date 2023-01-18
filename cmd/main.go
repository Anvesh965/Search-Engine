package main

import (
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"
)

// temp DB
func main() {
	DatabaseConn.Start()
	StartServer()
}
