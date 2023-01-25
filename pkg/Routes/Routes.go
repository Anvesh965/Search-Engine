package Routes

import (
	. "search-engine/cmd/config"
	. "search-engine/pkg/Controllers"
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
func HandleRoutes(router *gin.Engine, pgc *PageController) {

	router.GET("/", StatusCheck)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	HandleVersion1(router, pgc)

}
func HandleVersion1(router *gin.Engine, pgc *PageController) {
	var api1 = router.Group("/v1")

	api1.GET("/", pgc.ServerHome)
	api1.POST("/savepage", pgc.CreateWebPage)
	api1.POST("/querypages", pgc.QueryHandle)
	api1.GET("/allpages", pgc.GetAllWebPages)
}

func StartServer(pgc *PageController) {
	router := GetRouter()
	HandleRoutes(router, pgc)

	url := ":" + strconv.Itoa(Config.Server.Port)
	router.Run(url)
}
