// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongmusicxml/go/models"
	"github.com/fullstack-lang/gongmusicxml/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Backup__dummysDeclaration__ models.Backup
var __Backup_time__dummyDeclaration time.Duration

var mutexBackup sync.Mutex

// An BackupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBackup updateBackup deleteBackup
type BackupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BackupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBackup updateBackup
type BackupInput struct {
	// The Backup to submit or modify
	// in: body
	Backup *orm.BackupAPI
}

// GetBackups
//
// swagger:route GET /backups backups getBackups
//
// # Get all backups
//
// Responses:
// default: genericError
//
//	200: backupDBResponse
func (controller *Controller) GetBackups(c *gin.Context) {

	// source slice
	var backupDBs []orm.BackupDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBackups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBackup.GetDB()

	query := db.Find(&backupDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	backupAPIs := make([]orm.BackupAPI, 0)

	// for each backup, update fields from the database nullable fields
	for idx := range backupDBs {
		backupDB := &backupDBs[idx]
		_ = backupDB
		var backupAPI orm.BackupAPI

		// insertion point for updating fields
		backupAPI.ID = backupDB.ID
		backupDB.CopyBasicFieldsToBackup_WOP(&backupAPI.Backup_WOP)
		backupAPI.BackupPointersEncoding = backupDB.BackupPointersEncoding
		backupAPIs = append(backupAPIs, backupAPI)
	}

	c.JSON(http.StatusOK, backupAPIs)
}

// PostBackup
//
// swagger:route POST /backups backups postBackup
//
// Creates a backup
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBackup(c *gin.Context) {

	mutexBackup.Lock()
	defer mutexBackup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBackups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBackup.GetDB()

	// Validate input
	var input orm.BackupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create backup
	backupDB := orm.BackupDB{}
	backupDB.BackupPointersEncoding = input.BackupPointersEncoding
	backupDB.CopyBasicFieldsFromBackup_WOP(&input.Backup_WOP)

	query := db.Create(&backupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBackup.CheckoutPhaseOneInstance(&backupDB)
	backup := backRepo.BackRepoBackup.Map_BackupDBID_BackupPtr[backupDB.ID]

	if backup != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), backup)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, backupDB)
}

// GetBackup
//
// swagger:route GET /backups/{ID} backups getBackup
//
// Gets the details for a backup.
//
// Responses:
// default: genericError
//
//	200: backupDBResponse
func (controller *Controller) GetBackup(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBackup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBackup.GetDB()

	// Get backupDB in DB
	var backupDB orm.BackupDB
	if err := db.First(&backupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var backupAPI orm.BackupAPI
	backupAPI.ID = backupDB.ID
	backupAPI.BackupPointersEncoding = backupDB.BackupPointersEncoding
	backupDB.CopyBasicFieldsToBackup_WOP(&backupAPI.Backup_WOP)

	c.JSON(http.StatusOK, backupAPI)
}

// UpdateBackup
//
// swagger:route PATCH /backups/{ID} backups updateBackup
//
// # Update a backup
//
// Responses:
// default: genericError
//
//	200: backupDBResponse
func (controller *Controller) UpdateBackup(c *gin.Context) {

	mutexBackup.Lock()
	defer mutexBackup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBackup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBackup.GetDB()

	// Validate input
	var input orm.BackupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var backupDB orm.BackupDB

	// fetch the backup
	query := db.First(&backupDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	backupDB.CopyBasicFieldsFromBackup_WOP(&input.Backup_WOP)
	backupDB.BackupPointersEncoding = input.BackupPointersEncoding

	query = db.Model(&backupDB).Updates(backupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backupNew := new(models.Backup)
	backupDB.CopyBasicFieldsToBackup(backupNew)

	// redeem pointers
	backupDB.DecodePointers(backRepo, backupNew)

	// get stage instance from DB instance, and call callback function
	backupOld := backRepo.BackRepoBackup.Map_BackupDBID_BackupPtr[backupDB.ID]
	if backupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), backupOld, backupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the backupDB
	c.JSON(http.StatusOK, backupDB)
}

// DeleteBackup
//
// swagger:route DELETE /backups/{ID} backups deleteBackup
//
// # Delete a backup
//
// default: genericError
//
//	200: backupDBResponse
func (controller *Controller) DeleteBackup(c *gin.Context) {

	mutexBackup.Lock()
	defer mutexBackup.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBackup", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBackup.GetDB()

	// Get model if exist
	var backupDB orm.BackupDB
	if err := db.First(&backupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&backupDB)

	// get an instance (not staged) from DB instance, and call callback function
	backupDeleted := new(models.Backup)
	backupDB.CopyBasicFieldsToBackup(backupDeleted)

	// get stage instance from DB instance, and call callback function
	backupStaged := backRepo.BackRepoBackup.Map_BackupDBID_BackupPtr[backupDB.ID]
	if backupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), backupStaged, backupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
