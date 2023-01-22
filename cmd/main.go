package main

import (
	. "search-engine/cmd/config"
	"search-engine/pkg/DatabaseConn"
	. "search-engine/pkg/Routes"

	_ "search-engine/docs"
)

// @title Search-Engine-API
// @version 2.0
// @description Search-Engine-Rest_API. You can visit the GitHub repository at https://github.com/Anvesh965/Search-Engine
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:4000
// @BasePath /
// @query.collection.format multi
func main() {
	LoadConfig()
	rdb := &DatabaseConn.RealDBFunction{}
	rdb.Start()
	StartServer(rdb)
}
