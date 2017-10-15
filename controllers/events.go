package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nickhsine/test_backend/models"
	"github.com/nickhsine/test_backend/storage"
	"net/http"
	"strconv"
)

// NewEventController ...
func NewEventController(s storage.EventStorage) Controller {
	return &EventController{s}
}

// EventController ...
type EventController struct {
	Storage storage.EventStorage
}

// Close is the method of Controller interface
func (ec *EventController) Close() error {
	err := ec.Storage.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetEvents ...
func (ec *EventController) GetEvents(c *gin.Context) {
	var err error
	var events []models.Event
	var total uint

	_limit := c.Query("limit")
	_offset := c.Query("offset")
	limit, _ := strconv.Atoi(_limit)
	offset, _ := strconv.Atoi(_offset)

	if limit == 0 {
		limit = 100
	}

	events, total, err = ec.Storage.GetEvents(limit, offset)
	if err != nil {
		appErr := err.(models.AppError)
		c.JSON(appErr.StatusCode, gin.H{"status": "error", "message": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": struct {
		Total  uint           `json:"total"`
		Limit  int            `json:"limit"`
		Offset int            `json:"offset"`
		Events []models.Event `json:"events"`
	}{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Events: events,
	}})

}

// ViewEvent ...
func (ec *EventController) ViewEvent(c *gin.Context) {
	type postBody struct {
		EventID uint `json:"event_id" binding:"required"`
	}
	var body postBody
	var e models.Event
	var err error

	if err = c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	e, err = ec.Storage.GetEvent(body.EventID)
	if err != nil {
		appErr := err.(models.AppError)
		c.JSON(appErr.StatusCode, gin.H{"status": "error", "message": appErr.Error()})
		return
	}

	e.IsViewed = true
	err = ec.Storage.UpdateEvent(&e)
	if err != nil {
		appErr := err.(models.AppError)
		c.JSON(appErr.StatusCode, gin.H{"status": "error", "message": appErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": e})
}

// CreateEvent ...
func (ec *EventController) CreateEvent(c *gin.Context) {
	var e models.Event
	var err error

	if err = c.Bind(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	err = ec.Storage.CreateEvent(&e)
	if err != nil {
		appErr := err.(models.AppError)
		c.JSON(appErr.StatusCode, gin.H{"status": "error", "message": appErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": e})
}

// SetRoute is the method of Controller interface
func (ec *EventController) SetRoute(group *gin.RouterGroup) *gin.RouterGroup {
	// endpoints for events
	group.GET("/new-alarm-events/", ec.GetEvents)
	group.POST("/event-viewed/event-id/", ec.ViewEvent)
	group.POST("/new-alarm-events/", ec.CreateEvent)

	return group
}
