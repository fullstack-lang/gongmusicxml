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
var __Sync__dummysDeclaration__ models.Sync
var __Sync_time__dummyDeclaration time.Duration

var mutexSync sync.Mutex

// An SyncID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSync updateSync deleteSync
type SyncID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SyncInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSync updateSync
type SyncInput struct {
	// The Sync to submit or modify
	// in: body
	Sync *orm.SyncAPI
}

// GetSyncs
//
// swagger:route GET /syncs syncs getSyncs
//
// # Get all syncs
//
// Responses:
// default: genericError
//
//	200: syncDBResponse
func (controller *Controller) GetSyncs(c *gin.Context) {

	// source slice
	var syncDBs []orm.SyncDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSyncs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSync.GetDB()

	query := db.Find(&syncDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	syncAPIs := make([]orm.SyncAPI, 0)

	// for each sync, update fields from the database nullable fields
	for idx := range syncDBs {
		syncDB := &syncDBs[idx]
		_ = syncDB
		var syncAPI orm.SyncAPI

		// insertion point for updating fields
		syncAPI.ID = syncDB.ID
		syncDB.CopyBasicFieldsToSync_WOP(&syncAPI.Sync_WOP)
		syncAPI.SyncPointersEncoding = syncDB.SyncPointersEncoding
		syncAPIs = append(syncAPIs, syncAPI)
	}

	c.JSON(http.StatusOK, syncAPIs)
}

// PostSync
//
// swagger:route POST /syncs syncs postSync
//
// Creates a sync
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSync(c *gin.Context) {

	mutexSync.Lock()
	defer mutexSync.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSyncs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSync.GetDB()

	// Validate input
	var input orm.SyncAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create sync
	syncDB := orm.SyncDB{}
	syncDB.SyncPointersEncoding = input.SyncPointersEncoding
	syncDB.CopyBasicFieldsFromSync_WOP(&input.Sync_WOP)

	query := db.Create(&syncDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSync.CheckoutPhaseOneInstance(&syncDB)
	sync := backRepo.BackRepoSync.Map_SyncDBID_SyncPtr[syncDB.ID]

	if sync != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), sync)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, syncDB)
}

// GetSync
//
// swagger:route GET /syncs/{ID} syncs getSync
//
// Gets the details for a sync.
//
// Responses:
// default: genericError
//
//	200: syncDBResponse
func (controller *Controller) GetSync(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSync", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSync.GetDB()

	// Get syncDB in DB
	var syncDB orm.SyncDB
	if err := db.First(&syncDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var syncAPI orm.SyncAPI
	syncAPI.ID = syncDB.ID
	syncAPI.SyncPointersEncoding = syncDB.SyncPointersEncoding
	syncDB.CopyBasicFieldsToSync_WOP(&syncAPI.Sync_WOP)

	c.JSON(http.StatusOK, syncAPI)
}

// UpdateSync
//
// swagger:route PATCH /syncs/{ID} syncs updateSync
//
// # Update a sync
//
// Responses:
// default: genericError
//
//	200: syncDBResponse
func (controller *Controller) UpdateSync(c *gin.Context) {

	mutexSync.Lock()
	defer mutexSync.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSync", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSync.GetDB()

	// Validate input
	var input orm.SyncAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var syncDB orm.SyncDB

	// fetch the sync
	query := db.First(&syncDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	syncDB.CopyBasicFieldsFromSync_WOP(&input.Sync_WOP)
	syncDB.SyncPointersEncoding = input.SyncPointersEncoding

	query = db.Model(&syncDB).Updates(syncDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	syncNew := new(models.Sync)
	syncDB.CopyBasicFieldsToSync(syncNew)

	// redeem pointers
	syncDB.DecodePointers(backRepo, syncNew)

	// get stage instance from DB instance, and call callback function
	syncOld := backRepo.BackRepoSync.Map_SyncDBID_SyncPtr[syncDB.ID]
	if syncOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), syncOld, syncNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the syncDB
	c.JSON(http.StatusOK, syncDB)
}

// DeleteSync
//
// swagger:route DELETE /syncs/{ID} syncs deleteSync
//
// # Delete a sync
//
// default: genericError
//
//	200: syncDBResponse
func (controller *Controller) DeleteSync(c *gin.Context) {

	mutexSync.Lock()
	defer mutexSync.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSync", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSync.GetDB()

	// Get model if exist
	var syncDB orm.SyncDB
	if err := db.First(&syncDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&syncDB)

	// get an instance (not staged) from DB instance, and call callback function
	syncDeleted := new(models.Sync)
	syncDB.CopyBasicFieldsToSync(syncDeleted)

	// get stage instance from DB instance, and call callback function
	syncStaged := backRepo.BackRepoSync.Map_SyncDBID_SyncPtr[syncDB.ID]
	if syncStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), syncStaged, syncDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
