package storage

import (
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/nickhsine/test_backend/models"

	log "github.com/Sirupsen/logrus"
)

// ErrRecordNotFound record not found error, happens when haven't find any matched data when looking up with a struct
var ErrRecordNotFound = gorm.ErrRecordNotFound

// ErrDuplicateEntry record is already existed in MySQL database
var ErrDuplicateEntry uint16 = 1062

// ErrMgoDuplicateEntry record is already existed in MongoDB
var ErrMgoDuplicateEntry = 11000

// NewStorageError return AppError with detailed information.
// This method is mainly used to deal with MySQL error.
func (g *GormStorage) NewStorageError(err error, where string, message string) error {
	errStruct, ok := err.(*mysql.MySQLError)

	if err != nil && err.Error() == ErrRecordNotFound.Error() {
		return models.NewAppError(where, "Record not found", fmt.Sprintf("%v : %v", message, err.Error()), http.StatusNotFound)
	} else if ok && errStruct.Number == ErrDuplicateEntry {
		return models.NewAppError(where, "Record is already existed", fmt.Sprintf("%v : %v", message, err.Error()), http.StatusConflict)
	} else if err != nil {
		log.Error(err.Error())
		return models.NewAppError(where, "Internal server error", fmt.Sprintf("%v : %v", message, err.Error()), http.StatusInternalServerError)
	}
	return nil
}
