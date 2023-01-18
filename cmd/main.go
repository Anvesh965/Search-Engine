package main

import (
	. "search-engine/cmd/config"
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"
)

func main() {
	LoadConfig()
	DatabaseConn.Start()
	StartServer()
}
