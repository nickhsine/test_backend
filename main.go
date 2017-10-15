package main

import (
	"go/build"
	"net/http"
	"path/filepath"
	"time"

	"github.com/nickhsine/test_backend/controllers"
	"github.com/nickhsine/test_backend/routers"
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

	cf, err := controllers.NewControllerFactory()

	if err != nil {
		panic(err)
	}

	defer cf.Close()

	// set up the router
	router := routers.SetupRouter(cf)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	s.ListenAndServe()

}
