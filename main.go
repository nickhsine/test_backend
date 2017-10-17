package main

import (
	"go/build"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nickhsine/test_backend/controllers"
	"github.com/nickhsine/test_backend/routers"
	"github.com/nickhsine/test_backend/storage"
	"github.com/nickhsine/test_backend/utils"

	log "github.com/Sirupsen/logrus"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	p, _ := build.Default.Import("github.com/nickhsine/test_backend", "", build.FindOnly)

	fname := filepath.Join(p.Dir, "configs/config.json")

	// Load config file
	err := utils.LoadConfig(fname)
	if err != nil {
		log.Fatal("main.load_config.fatal_error: ", err.Error())
	}

	// set up database connection
	log.Info("Connecting to MySQL cloud")
	db, err := utils.InitDB(10, 5)
	if err != nil {
		panic(err)
	}

	// set up data storage
	gs := storage.NewGormStorage(db)

	// init controllers
	ec := controllers.NewEventController(gs)

	if err != nil {
		panic(err)
	}

	defer ec.Close()

	hub := newHub()
	go hub.run()

	// set up the router
	router := routers.SetupRouter(ec)
	routerGroup := router.Group("/v1")
	routerGroup.Any("/ws", func(c *gin.Context) {
		w := c.Writer
		r := c.Request

		serveWs(hub, w, r)
	})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
