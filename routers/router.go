package routers

import (
	// log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nickhsine/test_backend/controllers"
	"github.com/nickhsine/test_backend/utils"
)

// SetupRouter ...
func SetupRouter(cf *controllers.ControllerFactory) *gin.Engine {
	engine := gin.Default()

	config := cors.DefaultConfig()

	if utils.Cfg.Environment != "development" {
		config.AllowOrigins = utils.Cfg.CorsSettings.AllowOrigins
	} else {
		config.AllowAllOrigins = true
	}

	engine.Use(cors.New(config))

	routerGroup := engine.Group("/v1")
	{
		menuitems := new(controllers.MenuItemsController)
		routerGroup.GET("/ping", menuitems.Retrieve)
	}

	routerGroup = cf.SetRoute(routerGroup)

	return engine
}
