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
var __Midi_instrument__dummysDeclaration__ models.Midi_instrument
var __Midi_instrument_time__dummyDeclaration time.Duration

var mutexMidi_instrument sync.Mutex

// An Midi_instrumentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMidi_instrument updateMidi_instrument deleteMidi_instrument
type Midi_instrumentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Midi_instrumentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMidi_instrument updateMidi_instrument
type Midi_instrumentInput struct {
	// The Midi_instrument to submit or modify
	// in: body
	Midi_instrument *orm.Midi_instrumentAPI
}

// GetMidi_instruments
//
// swagger:route GET /midi_instruments midi_instruments getMidi_instruments
//
// # Get all midi_instruments
//
// Responses:
// default: genericError
//
//	200: midi_instrumentDBResponse
func (controller *Controller) GetMidi_instruments(c *gin.Context) {

	// source slice
	var midi_instrumentDBs []orm.Midi_instrumentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMidi_instruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_instrument.GetDB()

	query := db.Find(&midi_instrumentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	midi_instrumentAPIs := make([]orm.Midi_instrumentAPI, 0)

	// for each midi_instrument, update fields from the database nullable fields
	for idx := range midi_instrumentDBs {
		midi_instrumentDB := &midi_instrumentDBs[idx]
		_ = midi_instrumentDB
		var midi_instrumentAPI orm.Midi_instrumentAPI

		// insertion point for updating fields
		midi_instrumentAPI.ID = midi_instrumentDB.ID
		midi_instrumentDB.CopyBasicFieldsToMidi_instrument_WOP(&midi_instrumentAPI.Midi_instrument_WOP)
		midi_instrumentAPI.Midi_instrumentPointersEncoding = midi_instrumentDB.Midi_instrumentPointersEncoding
		midi_instrumentAPIs = append(midi_instrumentAPIs, midi_instrumentAPI)
	}

	c.JSON(http.StatusOK, midi_instrumentAPIs)
}

// PostMidi_instrument
//
// swagger:route POST /midi_instruments midi_instruments postMidi_instrument
//
// Creates a midi_instrument
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMidi_instrument(c *gin.Context) {

	mutexMidi_instrument.Lock()
	defer mutexMidi_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMidi_instruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_instrument.GetDB()

	// Validate input
	var input orm.Midi_instrumentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create midi_instrument
	midi_instrumentDB := orm.Midi_instrumentDB{}
	midi_instrumentDB.Midi_instrumentPointersEncoding = input.Midi_instrumentPointersEncoding
	midi_instrumentDB.CopyBasicFieldsFromMidi_instrument_WOP(&input.Midi_instrument_WOP)

	query := db.Create(&midi_instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMidi_instrument.CheckoutPhaseOneInstance(&midi_instrumentDB)
	midi_instrument := backRepo.BackRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID]

	if midi_instrument != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), midi_instrument)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, midi_instrumentDB)
}

// GetMidi_instrument
//
// swagger:route GET /midi_instruments/{ID} midi_instruments getMidi_instrument
//
// Gets the details for a midi_instrument.
//
// Responses:
// default: genericError
//
//	200: midi_instrumentDBResponse
func (controller *Controller) GetMidi_instrument(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMidi_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_instrument.GetDB()

	// Get midi_instrumentDB in DB
	var midi_instrumentDB orm.Midi_instrumentDB
	if err := db.First(&midi_instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var midi_instrumentAPI orm.Midi_instrumentAPI
	midi_instrumentAPI.ID = midi_instrumentDB.ID
	midi_instrumentAPI.Midi_instrumentPointersEncoding = midi_instrumentDB.Midi_instrumentPointersEncoding
	midi_instrumentDB.CopyBasicFieldsToMidi_instrument_WOP(&midi_instrumentAPI.Midi_instrument_WOP)

	c.JSON(http.StatusOK, midi_instrumentAPI)
}

// UpdateMidi_instrument
//
// swagger:route PATCH /midi_instruments/{ID} midi_instruments updateMidi_instrument
//
// # Update a midi_instrument
//
// Responses:
// default: genericError
//
//	200: midi_instrumentDBResponse
func (controller *Controller) UpdateMidi_instrument(c *gin.Context) {

	mutexMidi_instrument.Lock()
	defer mutexMidi_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMidi_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_instrument.GetDB()

	// Validate input
	var input orm.Midi_instrumentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var midi_instrumentDB orm.Midi_instrumentDB

	// fetch the midi_instrument
	query := db.First(&midi_instrumentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	midi_instrumentDB.CopyBasicFieldsFromMidi_instrument_WOP(&input.Midi_instrument_WOP)
	midi_instrumentDB.Midi_instrumentPointersEncoding = input.Midi_instrumentPointersEncoding

	query = db.Model(&midi_instrumentDB).Updates(midi_instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	midi_instrumentNew := new(models.Midi_instrument)
	midi_instrumentDB.CopyBasicFieldsToMidi_instrument(midi_instrumentNew)

	// redeem pointers
	midi_instrumentDB.DecodePointers(backRepo, midi_instrumentNew)

	// get stage instance from DB instance, and call callback function
	midi_instrumentOld := backRepo.BackRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID]
	if midi_instrumentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), midi_instrumentOld, midi_instrumentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the midi_instrumentDB
	c.JSON(http.StatusOK, midi_instrumentDB)
}

// DeleteMidi_instrument
//
// swagger:route DELETE /midi_instruments/{ID} midi_instruments deleteMidi_instrument
//
// # Delete a midi_instrument
//
// default: genericError
//
//	200: midi_instrumentDBResponse
func (controller *Controller) DeleteMidi_instrument(c *gin.Context) {

	mutexMidi_instrument.Lock()
	defer mutexMidi_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMidi_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_instrument.GetDB()

	// Get model if exist
	var midi_instrumentDB orm.Midi_instrumentDB
	if err := db.First(&midi_instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&midi_instrumentDB)

	// get an instance (not staged) from DB instance, and call callback function
	midi_instrumentDeleted := new(models.Midi_instrument)
	midi_instrumentDB.CopyBasicFieldsToMidi_instrument(midi_instrumentDeleted)

	// get stage instance from DB instance, and call callback function
	midi_instrumentStaged := backRepo.BackRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID]
	if midi_instrumentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), midi_instrumentStaged, midi_instrumentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
