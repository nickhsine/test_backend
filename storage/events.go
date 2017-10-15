package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/nickhsine/test_backend/models"
	//log "github.com/Sirupsen/logrus"
)

// EventStorage defines the methods we need to implement,
// in order to fulfill the functionalities a system needs.
type EventStorage interface {
	/** Close DB Connection **/
	Close() error

	GetEvent(uint) (models.Event, error)
	GetEvents(int, int) ([]models.Event, uint, error)
	UpdateEvent(*models.Event) error
	CreateEvent(*models.Event) error
}

// NewGormStorage initializes the storage connected to MySQL database by gorm library
func NewGormStorage(db *gorm.DB) *GormStorage {
	return &GormStorage{db}
}

// GormStorage implements MembershipStorage interface
type GormStorage struct {
	db *gorm.DB
}

// Close quits the DB connection gracefully
func (gs *GormStorage) Close() error {
	err := gs.db.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetEvent ...
func (gs *GormStorage) GetEvent(eventID uint) (models.Event, error) {
	var event models.Event
	err := gs.db.Where("id = ?", eventID).Find(&event).Error
	if err != nil {
		return event, gs.NewStorageError(err, "EventStorage.GetEvents", "Querying DB record occurs error")
	}

	return event, nil
}

// GetEvents ...
func (gs *GormStorage) GetEvents(limit, offset int) ([]models.Event, uint, error) {
	var events []models.Event
	var total uint

	err := gs.db.Order("starting_timestamp desc").Limit(limit).Offset(offset).Find(&events).Error
	if err != nil {
		return events, 0, gs.NewStorageError(err, "EventStorage.GetEvents", "Querying DB record occurs error")
	}

	err = gs.db.Table("events").Count(&total).Error
	if err != nil {
		return events, 0, gs.NewStorageError(err, "EventStorage.GetEvents", "Getting total records in events table occurs error")
	}

	return events, total, nil
}

// CreateEvent ...
func (gs *GormStorage) CreateEvent(e *models.Event) error {
	err := gs.db.Create(e).Error
	if err != nil {
		return gs.NewStorageError(err, "EventStorage.CreateEvent", "Inserting DB record occurs error")
	}

	return err
}

// UpdateEvent ...
func (gs *GormStorage) UpdateEvent(e *models.Event) error {
	err := gs.db.Save(e).Error
	if err != nil {
		return gs.NewStorageError(err, "EventStorage.UpdateEvent", "Updating DB record  occurs error")
	}

	return err
}
