package main

import (
	"fmt"
	. "search-engine/cmd/config"
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"
)

// temp DB
func main() {
	fmt.Println("test")
	LoadConfig()
	DatabaseConn.Start()
	StartServer()
}
