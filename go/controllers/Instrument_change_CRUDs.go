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
var __Instrument_change__dummysDeclaration__ models.Instrument_change
var __Instrument_change_time__dummyDeclaration time.Duration

var mutexInstrument_change sync.Mutex

// An Instrument_changeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getInstrument_change updateInstrument_change deleteInstrument_change
type Instrument_changeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Instrument_changeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postInstrument_change updateInstrument_change
type Instrument_changeInput struct {
	// The Instrument_change to submit or modify
	// in: body
	Instrument_change *orm.Instrument_changeAPI
}

// GetInstrument_changes
//
// swagger:route GET /instrument_changes instrument_changes getInstrument_changes
//
// # Get all instrument_changes
//
// Responses:
// default: genericError
//
//	200: instrument_changeDBResponse
func (controller *Controller) GetInstrument_changes(c *gin.Context) {

	// source slice
	var instrument_changeDBs []orm.Instrument_changeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInstrument_changes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_change.GetDB()

	query := db.Find(&instrument_changeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	instrument_changeAPIs := make([]orm.Instrument_changeAPI, 0)

	// for each instrument_change, update fields from the database nullable fields
	for idx := range instrument_changeDBs {
		instrument_changeDB := &instrument_changeDBs[idx]
		_ = instrument_changeDB
		var instrument_changeAPI orm.Instrument_changeAPI

		// insertion point for updating fields
		instrument_changeAPI.ID = instrument_changeDB.ID
		instrument_changeDB.CopyBasicFieldsToInstrument_change_WOP(&instrument_changeAPI.Instrument_change_WOP)
		instrument_changeAPI.Instrument_changePointersEncoding = instrument_changeDB.Instrument_changePointersEncoding
		instrument_changeAPIs = append(instrument_changeAPIs, instrument_changeAPI)
	}

	c.JSON(http.StatusOK, instrument_changeAPIs)
}

// PostInstrument_change
//
// swagger:route POST /instrument_changes instrument_changes postInstrument_change
//
// Creates a instrument_change
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostInstrument_change(c *gin.Context) {

	mutexInstrument_change.Lock()
	defer mutexInstrument_change.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostInstrument_changes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_change.GetDB()

	// Validate input
	var input orm.Instrument_changeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create instrument_change
	instrument_changeDB := orm.Instrument_changeDB{}
	instrument_changeDB.Instrument_changePointersEncoding = input.Instrument_changePointersEncoding
	instrument_changeDB.CopyBasicFieldsFromInstrument_change_WOP(&input.Instrument_change_WOP)

	query := db.Create(&instrument_changeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoInstrument_change.CheckoutPhaseOneInstance(&instrument_changeDB)
	instrument_change := backRepo.BackRepoInstrument_change.Map_Instrument_changeDBID_Instrument_changePtr[instrument_changeDB.ID]

	if instrument_change != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), instrument_change)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, instrument_changeDB)
}

// GetInstrument_change
//
// swagger:route GET /instrument_changes/{ID} instrument_changes getInstrument_change
//
// Gets the details for a instrument_change.
//
// Responses:
// default: genericError
//
//	200: instrument_changeDBResponse
func (controller *Controller) GetInstrument_change(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInstrument_change", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_change.GetDB()

	// Get instrument_changeDB in DB
	var instrument_changeDB orm.Instrument_changeDB
	if err := db.First(&instrument_changeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var instrument_changeAPI orm.Instrument_changeAPI
	instrument_changeAPI.ID = instrument_changeDB.ID
	instrument_changeAPI.Instrument_changePointersEncoding = instrument_changeDB.Instrument_changePointersEncoding
	instrument_changeDB.CopyBasicFieldsToInstrument_change_WOP(&instrument_changeAPI.Instrument_change_WOP)

	c.JSON(http.StatusOK, instrument_changeAPI)
}

// UpdateInstrument_change
//
// swagger:route PATCH /instrument_changes/{ID} instrument_changes updateInstrument_change
//
// # Update a instrument_change
//
// Responses:
// default: genericError
//
//	200: instrument_changeDBResponse
func (controller *Controller) UpdateInstrument_change(c *gin.Context) {

	mutexInstrument_change.Lock()
	defer mutexInstrument_change.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateInstrument_change", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_change.GetDB()

	// Validate input
	var input orm.Instrument_changeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var instrument_changeDB orm.Instrument_changeDB

	// fetch the instrument_change
	query := db.First(&instrument_changeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	instrument_changeDB.CopyBasicFieldsFromInstrument_change_WOP(&input.Instrument_change_WOP)
	instrument_changeDB.Instrument_changePointersEncoding = input.Instrument_changePointersEncoding

	query = db.Model(&instrument_changeDB).Updates(instrument_changeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	instrument_changeNew := new(models.Instrument_change)
	instrument_changeDB.CopyBasicFieldsToInstrument_change(instrument_changeNew)

	// redeem pointers
	instrument_changeDB.DecodePointers(backRepo, instrument_changeNew)

	// get stage instance from DB instance, and call callback function
	instrument_changeOld := backRepo.BackRepoInstrument_change.Map_Instrument_changeDBID_Instrument_changePtr[instrument_changeDB.ID]
	if instrument_changeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), instrument_changeOld, instrument_changeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the instrument_changeDB
	c.JSON(http.StatusOK, instrument_changeDB)
}

// DeleteInstrument_change
//
// swagger:route DELETE /instrument_changes/{ID} instrument_changes deleteInstrument_change
//
// # Delete a instrument_change
//
// default: genericError
//
//	200: instrument_changeDBResponse
func (controller *Controller) DeleteInstrument_change(c *gin.Context) {

	mutexInstrument_change.Lock()
	defer mutexInstrument_change.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteInstrument_change", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument_change.GetDB()

	// Get model if exist
	var instrument_changeDB orm.Instrument_changeDB
	if err := db.First(&instrument_changeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&instrument_changeDB)

	// get an instance (not staged) from DB instance, and call callback function
	instrument_changeDeleted := new(models.Instrument_change)
	instrument_changeDB.CopyBasicFieldsToInstrument_change(instrument_changeDeleted)

	// get stage instance from DB instance, and call callback function
	instrument_changeStaged := backRepo.BackRepoInstrument_change.Map_Instrument_changeDBID_Instrument_changePtr[instrument_changeDB.ID]
	if instrument_changeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), instrument_changeStaged, instrument_changeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
