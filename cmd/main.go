package main

import (
	"fmt"
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"
)

// temp DB
func main() {
	fmt.Println("test")
	DatabaseConn.Start()
	StartServer()
}
