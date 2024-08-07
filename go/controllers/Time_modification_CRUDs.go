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
var __Time_modification__dummysDeclaration__ models.Time_modification
var __Time_modification_time__dummyDeclaration time.Duration

var mutexTime_modification sync.Mutex

// An Time_modificationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTime_modification updateTime_modification deleteTime_modification
type Time_modificationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Time_modificationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTime_modification updateTime_modification
type Time_modificationInput struct {
	// The Time_modification to submit or modify
	// in: body
	Time_modification *orm.Time_modificationAPI
}

// GetTime_modifications
//
// swagger:route GET /time_modifications time_modifications getTime_modifications
//
// # Get all time_modifications
//
// Responses:
// default: genericError
//
//	200: time_modificationDBResponse
func (controller *Controller) GetTime_modifications(c *gin.Context) {

	// source slice
	var time_modificationDBs []orm.Time_modificationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTime_modifications", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime_modification.GetDB()

	query := db.Find(&time_modificationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	time_modificationAPIs := make([]orm.Time_modificationAPI, 0)

	// for each time_modification, update fields from the database nullable fields
	for idx := range time_modificationDBs {
		time_modificationDB := &time_modificationDBs[idx]
		_ = time_modificationDB
		var time_modificationAPI orm.Time_modificationAPI

		// insertion point for updating fields
		time_modificationAPI.ID = time_modificationDB.ID
		time_modificationDB.CopyBasicFieldsToTime_modification_WOP(&time_modificationAPI.Time_modification_WOP)
		time_modificationAPI.Time_modificationPointersEncoding = time_modificationDB.Time_modificationPointersEncoding
		time_modificationAPIs = append(time_modificationAPIs, time_modificationAPI)
	}

	c.JSON(http.StatusOK, time_modificationAPIs)
}

// PostTime_modification
//
// swagger:route POST /time_modifications time_modifications postTime_modification
//
// Creates a time_modification
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTime_modification(c *gin.Context) {

	mutexTime_modification.Lock()
	defer mutexTime_modification.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTime_modifications", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime_modification.GetDB()

	// Validate input
	var input orm.Time_modificationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create time_modification
	time_modificationDB := orm.Time_modificationDB{}
	time_modificationDB.Time_modificationPointersEncoding = input.Time_modificationPointersEncoding
	time_modificationDB.CopyBasicFieldsFromTime_modification_WOP(&input.Time_modification_WOP)

	query := db.Create(&time_modificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTime_modification.CheckoutPhaseOneInstance(&time_modificationDB)
	time_modification := backRepo.BackRepoTime_modification.Map_Time_modificationDBID_Time_modificationPtr[time_modificationDB.ID]

	if time_modification != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), time_modification)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, time_modificationDB)
}

// GetTime_modification
//
// swagger:route GET /time_modifications/{ID} time_modifications getTime_modification
//
// Gets the details for a time_modification.
//
// Responses:
// default: genericError
//
//	200: time_modificationDBResponse
func (controller *Controller) GetTime_modification(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTime_modification", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime_modification.GetDB()

	// Get time_modificationDB in DB
	var time_modificationDB orm.Time_modificationDB
	if err := db.First(&time_modificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var time_modificationAPI orm.Time_modificationAPI
	time_modificationAPI.ID = time_modificationDB.ID
	time_modificationAPI.Time_modificationPointersEncoding = time_modificationDB.Time_modificationPointersEncoding
	time_modificationDB.CopyBasicFieldsToTime_modification_WOP(&time_modificationAPI.Time_modification_WOP)

	c.JSON(http.StatusOK, time_modificationAPI)
}

// UpdateTime_modification
//
// swagger:route PATCH /time_modifications/{ID} time_modifications updateTime_modification
//
// # Update a time_modification
//
// Responses:
// default: genericError
//
//	200: time_modificationDBResponse
func (controller *Controller) UpdateTime_modification(c *gin.Context) {

	mutexTime_modification.Lock()
	defer mutexTime_modification.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTime_modification", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime_modification.GetDB()

	// Validate input
	var input orm.Time_modificationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var time_modificationDB orm.Time_modificationDB

	// fetch the time_modification
	query := db.First(&time_modificationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	time_modificationDB.CopyBasicFieldsFromTime_modification_WOP(&input.Time_modification_WOP)
	time_modificationDB.Time_modificationPointersEncoding = input.Time_modificationPointersEncoding

	query = db.Model(&time_modificationDB).Updates(time_modificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	time_modificationNew := new(models.Time_modification)
	time_modificationDB.CopyBasicFieldsToTime_modification(time_modificationNew)

	// redeem pointers
	time_modificationDB.DecodePointers(backRepo, time_modificationNew)

	// get stage instance from DB instance, and call callback function
	time_modificationOld := backRepo.BackRepoTime_modification.Map_Time_modificationDBID_Time_modificationPtr[time_modificationDB.ID]
	if time_modificationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), time_modificationOld, time_modificationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the time_modificationDB
	c.JSON(http.StatusOK, time_modificationDB)
}

// DeleteTime_modification
//
// swagger:route DELETE /time_modifications/{ID} time_modifications deleteTime_modification
//
// # Delete a time_modification
//
// default: genericError
//
//	200: time_modificationDBResponse
func (controller *Controller) DeleteTime_modification(c *gin.Context) {

	mutexTime_modification.Lock()
	defer mutexTime_modification.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTime_modification", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime_modification.GetDB()

	// Get model if exist
	var time_modificationDB orm.Time_modificationDB
	if err := db.First(&time_modificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&time_modificationDB)

	// get an instance (not staged) from DB instance, and call callback function
	time_modificationDeleted := new(models.Time_modification)
	time_modificationDB.CopyBasicFieldsToTime_modification(time_modificationDeleted)

	// get stage instance from DB instance, and call callback function
	time_modificationStaged := backRepo.BackRepoTime_modification.Map_Time_modificationDBID_Time_modificationPtr[time_modificationDB.ID]
	if time_modificationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), time_modificationStaged, time_modificationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
