package main

import (
	. "search-engine/cmd/config"
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"

	_ "search-engine/docs"
)

// @title Search-Engine-API
// @version 2.0
// @description Search-Engine-Rest-API. You can visit the GitHub repository at https://github.com/Anvesh965/Search-Engine
// @host localhost:4000
// @BasePath /
// @query.collection.format multi
func main() {
	LoadConfig()
	// rdb := &DatabaseConn.RealDBFunction{}
	ch := DatabaseConn.Start()
	rdb := DatabaseConn.NewDBFunctions(ch)
	StartServer(rdb)
}
