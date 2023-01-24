package Routes

import (
	. "search-engine/cmd/config"
	. "search-engine/pkg/Controllers"
	. "search-engine/pkg/DatabaseConn"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	return r
}
func HandleRoutes(router *gin.Engine, rdb DBFunctions) {

	router.GET("/", StatusCheck)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	HandleVersion1(router, rdb)

}
func HandleVersion1(router *gin.Engine, rdb DBFunctions) {
	var api1 = router.Group("/v1")
	api1.GET("/", ServerHome)
	api1.POST("/savepage", func(c *gin.Context) {
		CreateWebPage(c, rdb)
	})
	api1.POST("/querypages", func(c *gin.Context) {
		QueryHandle(c, rdb)
	})
	api1.GET("/allpages", func(c *gin.Context) {
		GetAllWebPages(c, rdb)
	})
}

func StartServer(rdb DBFunctions) {
	router := GetRouter()
	HandleRoutes(router, rdb)

	url := ":" + strconv.Itoa(Config.Server.Port)
	router.Run(url)
}
