package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nickhsine/test_backend/constants"
	"github.com/nickhsine/test_backend/storage"
	"github.com/nickhsine/test_backend/utils"

	log "github.com/Sirupsen/logrus"
)

// Controller ...
type Controller interface {
	SetRoute(*gin.RouterGroup) *gin.RouterGroup
	Close() error
}

// ControllerFactory ...
type ControllerFactory struct {
	Controllers map[string]Controller
}

// GetController ...
func (cf *ControllerFactory) GetController(cn string) Controller {
	return cf.Controllers[cn]
}

// GetControllers returns an array of controllers
func (cf *ControllerFactory) GetControllers() []Controller {
	var cons []Controller

	for _, con := range cf.Controllers {
		cons = append(cons, con)
	}
	return cons
}

// SetController ...
func (cf *ControllerFactory) SetController(cn string, c Controller) {
	cf.Controllers[cn] = c
}

// Close this func releases the resource appropriately
func (cf *ControllerFactory) Close() error {
	var err error
	for _, controller := range cf.GetControllers() {
		err = controller.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// SetRoute set route by calling the correspoding controllers.
func (cf *ControllerFactory) SetRoute(group *gin.RouterGroup) *gin.RouterGroup {
	for _, v := range cf.Controllers {
		group = v.SetRoute(group)
	}
	return group
}

// NewControllerFactory ...
func NewControllerFactory() (*ControllerFactory, error) {
	// set up database connection
	log.Info("Connecting to MySQL cloud")
	db, err := utils.InitDB(10, 5)
	if err != nil {
		return nil, err
	}

	// set up data storage
	gs := storage.NewGormStorage(db)

	// init controllers
	ec := NewEventController(gs)

	cf := &ControllerFactory{
		Controllers: make(map[string]Controller),
	}

	cf.SetController(constants.EventController, ec)

	return cf, nil
}
