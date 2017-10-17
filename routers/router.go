package routers

import (
	// "net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nickhsine/test_backend/controllers"
	//log "github.com/Sirupsen/logrus"
)

// SetupRouter ...
func SetupRouter(ec *controllers.EventController) *gin.Engine {
	engine := gin.Default()

	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AddAllowHeaders("Access-Control-Allow-Headers", "X-Requested-With", "Origin")
	config.AddAllowMethods("OPTIONS")

	engine.Use(cors.New(config))

	routerGroup := engine.Group("/v1")
	{
		menuitems := new(controllers.MenuItemsController)
		routerGroup.GET("/ping", menuitems.Retrieve)
	}

	routerGroup = ec.SetRoute(routerGroup)

	return engine
}
