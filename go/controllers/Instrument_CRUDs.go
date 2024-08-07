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
var __Instrument__dummysDeclaration__ models.Instrument
var __Instrument_time__dummyDeclaration time.Duration

var mutexInstrument sync.Mutex

// An InstrumentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getInstrument updateInstrument deleteInstrument
type InstrumentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// InstrumentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postInstrument updateInstrument
type InstrumentInput struct {
	// The Instrument to submit or modify
	// in: body
	Instrument *orm.InstrumentAPI
}

// GetInstruments
//
// swagger:route GET /instruments instruments getInstruments
//
// # Get all instruments
//
// Responses:
// default: genericError
//
//	200: instrumentDBResponse
func (controller *Controller) GetInstruments(c *gin.Context) {

	// source slice
	var instrumentDBs []orm.InstrumentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInstruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument.GetDB()

	query := db.Find(&instrumentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	instrumentAPIs := make([]orm.InstrumentAPI, 0)

	// for each instrument, update fields from the database nullable fields
	for idx := range instrumentDBs {
		instrumentDB := &instrumentDBs[idx]
		_ = instrumentDB
		var instrumentAPI orm.InstrumentAPI

		// insertion point for updating fields
		instrumentAPI.ID = instrumentDB.ID
		instrumentDB.CopyBasicFieldsToInstrument_WOP(&instrumentAPI.Instrument_WOP)
		instrumentAPI.InstrumentPointersEncoding = instrumentDB.InstrumentPointersEncoding
		instrumentAPIs = append(instrumentAPIs, instrumentAPI)
	}

	c.JSON(http.StatusOK, instrumentAPIs)
}

// PostInstrument
//
// swagger:route POST /instruments instruments postInstrument
//
// Creates a instrument
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostInstrument(c *gin.Context) {

	mutexInstrument.Lock()
	defer mutexInstrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostInstruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument.GetDB()

	// Validate input
	var input orm.InstrumentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create instrument
	instrumentDB := orm.InstrumentDB{}
	instrumentDB.InstrumentPointersEncoding = input.InstrumentPointersEncoding
	instrumentDB.CopyBasicFieldsFromInstrument_WOP(&input.Instrument_WOP)

	query := db.Create(&instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoInstrument.CheckoutPhaseOneInstance(&instrumentDB)
	instrument := backRepo.BackRepoInstrument.Map_InstrumentDBID_InstrumentPtr[instrumentDB.ID]

	if instrument != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), instrument)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, instrumentDB)
}

// GetInstrument
//
// swagger:route GET /instruments/{ID} instruments getInstrument
//
// Gets the details for a instrument.
//
// Responses:
// default: genericError
//
//	200: instrumentDBResponse
func (controller *Controller) GetInstrument(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInstrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument.GetDB()

	// Get instrumentDB in DB
	var instrumentDB orm.InstrumentDB
	if err := db.First(&instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var instrumentAPI orm.InstrumentAPI
	instrumentAPI.ID = instrumentDB.ID
	instrumentAPI.InstrumentPointersEncoding = instrumentDB.InstrumentPointersEncoding
	instrumentDB.CopyBasicFieldsToInstrument_WOP(&instrumentAPI.Instrument_WOP)

	c.JSON(http.StatusOK, instrumentAPI)
}

// UpdateInstrument
//
// swagger:route PATCH /instruments/{ID} instruments updateInstrument
//
// # Update a instrument
//
// Responses:
// default: genericError
//
//	200: instrumentDBResponse
func (controller *Controller) UpdateInstrument(c *gin.Context) {

	mutexInstrument.Lock()
	defer mutexInstrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateInstrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument.GetDB()

	// Validate input
	var input orm.InstrumentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var instrumentDB orm.InstrumentDB

	// fetch the instrument
	query := db.First(&instrumentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	instrumentDB.CopyBasicFieldsFromInstrument_WOP(&input.Instrument_WOP)
	instrumentDB.InstrumentPointersEncoding = input.InstrumentPointersEncoding

	query = db.Model(&instrumentDB).Updates(instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	instrumentNew := new(models.Instrument)
	instrumentDB.CopyBasicFieldsToInstrument(instrumentNew)

	// redeem pointers
	instrumentDB.DecodePointers(backRepo, instrumentNew)

	// get stage instance from DB instance, and call callback function
	instrumentOld := backRepo.BackRepoInstrument.Map_InstrumentDBID_InstrumentPtr[instrumentDB.ID]
	if instrumentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), instrumentOld, instrumentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the instrumentDB
	c.JSON(http.StatusOK, instrumentDB)
}

// DeleteInstrument
//
// swagger:route DELETE /instruments/{ID} instruments deleteInstrument
//
// # Delete a instrument
//
// default: genericError
//
//	200: instrumentDBResponse
func (controller *Controller) DeleteInstrument(c *gin.Context) {

	mutexInstrument.Lock()
	defer mutexInstrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteInstrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInstrument.GetDB()

	// Get model if exist
	var instrumentDB orm.InstrumentDB
	if err := db.First(&instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&instrumentDB)

	// get an instance (not staged) from DB instance, and call callback function
	instrumentDeleted := new(models.Instrument)
	instrumentDB.CopyBasicFieldsToInstrument(instrumentDeleted)

	// get stage instance from DB instance, and call callback function
	instrumentStaged := backRepo.BackRepoInstrument.Map_InstrumentDBID_InstrumentPtr[instrumentDB.ID]
	if instrumentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), instrumentStaged, instrumentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
