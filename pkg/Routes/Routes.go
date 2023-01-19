package Routes

import (
	. "search-engine/cmd/config"
	. "search-engine/pkg/Controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	return r
}
func HandleRoutes(router *gin.Engine) {

	router.GET("/", StatusCheck)
	HandleVersion1(router)

}
func HandleVersion1(router *gin.Engine) {
	var api1 = router.Group("/v1")
	api1.GET("/", ServerHome)
	api1.POST("/savepage", CreateWebPage)
	api1.GET("/querypages", QueryHandle)
	api1.GET("/allpages", GetAllWebPages)
}

func StartServer() {
	router := GetRouter()
	HandleRoutes(router)

	url := ":" + strconv.Itoa(Config.Server.Port)
	router.Run(url)
}
