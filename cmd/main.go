package main

import (
	. "search-engine/cmd/config"
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"
)

// temp DB
func main() {
	LoadConfig()
	DatabaseConn.Start()
	StartServer()
}
