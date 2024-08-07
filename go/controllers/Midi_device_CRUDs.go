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
var __Midi_device__dummysDeclaration__ models.Midi_device
var __Midi_device_time__dummyDeclaration time.Duration

var mutexMidi_device sync.Mutex

// An Midi_deviceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMidi_device updateMidi_device deleteMidi_device
type Midi_deviceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Midi_deviceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMidi_device updateMidi_device
type Midi_deviceInput struct {
	// The Midi_device to submit or modify
	// in: body
	Midi_device *orm.Midi_deviceAPI
}

// GetMidi_devices
//
// swagger:route GET /midi_devices midi_devices getMidi_devices
//
// # Get all midi_devices
//
// Responses:
// default: genericError
//
//	200: midi_deviceDBResponse
func (controller *Controller) GetMidi_devices(c *gin.Context) {

	// source slice
	var midi_deviceDBs []orm.Midi_deviceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMidi_devices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_device.GetDB()

	query := db.Find(&midi_deviceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	midi_deviceAPIs := make([]orm.Midi_deviceAPI, 0)

	// for each midi_device, update fields from the database nullable fields
	for idx := range midi_deviceDBs {
		midi_deviceDB := &midi_deviceDBs[idx]
		_ = midi_deviceDB
		var midi_deviceAPI orm.Midi_deviceAPI

		// insertion point for updating fields
		midi_deviceAPI.ID = midi_deviceDB.ID
		midi_deviceDB.CopyBasicFieldsToMidi_device_WOP(&midi_deviceAPI.Midi_device_WOP)
		midi_deviceAPI.Midi_devicePointersEncoding = midi_deviceDB.Midi_devicePointersEncoding
		midi_deviceAPIs = append(midi_deviceAPIs, midi_deviceAPI)
	}

	c.JSON(http.StatusOK, midi_deviceAPIs)
}

// PostMidi_device
//
// swagger:route POST /midi_devices midi_devices postMidi_device
//
// Creates a midi_device
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMidi_device(c *gin.Context) {

	mutexMidi_device.Lock()
	defer mutexMidi_device.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMidi_devices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_device.GetDB()

	// Validate input
	var input orm.Midi_deviceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create midi_device
	midi_deviceDB := orm.Midi_deviceDB{}
	midi_deviceDB.Midi_devicePointersEncoding = input.Midi_devicePointersEncoding
	midi_deviceDB.CopyBasicFieldsFromMidi_device_WOP(&input.Midi_device_WOP)

	query := db.Create(&midi_deviceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMidi_device.CheckoutPhaseOneInstance(&midi_deviceDB)
	midi_device := backRepo.BackRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID]

	if midi_device != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), midi_device)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, midi_deviceDB)
}

// GetMidi_device
//
// swagger:route GET /midi_devices/{ID} midi_devices getMidi_device
//
// Gets the details for a midi_device.
//
// Responses:
// default: genericError
//
//	200: midi_deviceDBResponse
func (controller *Controller) GetMidi_device(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMidi_device", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_device.GetDB()

	// Get midi_deviceDB in DB
	var midi_deviceDB orm.Midi_deviceDB
	if err := db.First(&midi_deviceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var midi_deviceAPI orm.Midi_deviceAPI
	midi_deviceAPI.ID = midi_deviceDB.ID
	midi_deviceAPI.Midi_devicePointersEncoding = midi_deviceDB.Midi_devicePointersEncoding
	midi_deviceDB.CopyBasicFieldsToMidi_device_WOP(&midi_deviceAPI.Midi_device_WOP)

	c.JSON(http.StatusOK, midi_deviceAPI)
}

// UpdateMidi_device
//
// swagger:route PATCH /midi_devices/{ID} midi_devices updateMidi_device
//
// # Update a midi_device
//
// Responses:
// default: genericError
//
//	200: midi_deviceDBResponse
func (controller *Controller) UpdateMidi_device(c *gin.Context) {

	mutexMidi_device.Lock()
	defer mutexMidi_device.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMidi_device", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_device.GetDB()

	// Validate input
	var input orm.Midi_deviceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var midi_deviceDB orm.Midi_deviceDB

	// fetch the midi_device
	query := db.First(&midi_deviceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	midi_deviceDB.CopyBasicFieldsFromMidi_device_WOP(&input.Midi_device_WOP)
	midi_deviceDB.Midi_devicePointersEncoding = input.Midi_devicePointersEncoding

	query = db.Model(&midi_deviceDB).Updates(midi_deviceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	midi_deviceNew := new(models.Midi_device)
	midi_deviceDB.CopyBasicFieldsToMidi_device(midi_deviceNew)

	// redeem pointers
	midi_deviceDB.DecodePointers(backRepo, midi_deviceNew)

	// get stage instance from DB instance, and call callback function
	midi_deviceOld := backRepo.BackRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID]
	if midi_deviceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), midi_deviceOld, midi_deviceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the midi_deviceDB
	c.JSON(http.StatusOK, midi_deviceDB)
}

// DeleteMidi_device
//
// swagger:route DELETE /midi_devices/{ID} midi_devices deleteMidi_device
//
// # Delete a midi_device
//
// default: genericError
//
//	200: midi_deviceDBResponse
func (controller *Controller) DeleteMidi_device(c *gin.Context) {

	mutexMidi_device.Lock()
	defer mutexMidi_device.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMidi_device", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMidi_device.GetDB()

	// Get model if exist
	var midi_deviceDB orm.Midi_deviceDB
	if err := db.First(&midi_deviceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&midi_deviceDB)

	// get an instance (not staged) from DB instance, and call callback function
	midi_deviceDeleted := new(models.Midi_device)
	midi_deviceDB.CopyBasicFieldsToMidi_device(midi_deviceDeleted)

	// get stage instance from DB instance, and call callback function
	midi_deviceStaged := backRepo.BackRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID]
	if midi_deviceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), midi_deviceStaged, midi_deviceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
