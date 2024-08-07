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
var __Virtual_instrument__dummysDeclaration__ models.Virtual_instrument
var __Virtual_instrument_time__dummyDeclaration time.Duration

var mutexVirtual_instrument sync.Mutex

// An Virtual_instrumentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVirtual_instrument updateVirtual_instrument deleteVirtual_instrument
type Virtual_instrumentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Virtual_instrumentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVirtual_instrument updateVirtual_instrument
type Virtual_instrumentInput struct {
	// The Virtual_instrument to submit or modify
	// in: body
	Virtual_instrument *orm.Virtual_instrumentAPI
}

// GetVirtual_instruments
//
// swagger:route GET /virtual_instruments virtual_instruments getVirtual_instruments
//
// # Get all virtual_instruments
//
// Responses:
// default: genericError
//
//	200: virtual_instrumentDBResponse
func (controller *Controller) GetVirtual_instruments(c *gin.Context) {

	// source slice
	var virtual_instrumentDBs []orm.Virtual_instrumentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVirtual_instruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVirtual_instrument.GetDB()

	query := db.Find(&virtual_instrumentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	virtual_instrumentAPIs := make([]orm.Virtual_instrumentAPI, 0)

	// for each virtual_instrument, update fields from the database nullable fields
	for idx := range virtual_instrumentDBs {
		virtual_instrumentDB := &virtual_instrumentDBs[idx]
		_ = virtual_instrumentDB
		var virtual_instrumentAPI orm.Virtual_instrumentAPI

		// insertion point for updating fields
		virtual_instrumentAPI.ID = virtual_instrumentDB.ID
		virtual_instrumentDB.CopyBasicFieldsToVirtual_instrument_WOP(&virtual_instrumentAPI.Virtual_instrument_WOP)
		virtual_instrumentAPI.Virtual_instrumentPointersEncoding = virtual_instrumentDB.Virtual_instrumentPointersEncoding
		virtual_instrumentAPIs = append(virtual_instrumentAPIs, virtual_instrumentAPI)
	}

	c.JSON(http.StatusOK, virtual_instrumentAPIs)
}

// PostVirtual_instrument
//
// swagger:route POST /virtual_instruments virtual_instruments postVirtual_instrument
//
// Creates a virtual_instrument
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostVirtual_instrument(c *gin.Context) {

	mutexVirtual_instrument.Lock()
	defer mutexVirtual_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostVirtual_instruments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVirtual_instrument.GetDB()

	// Validate input
	var input orm.Virtual_instrumentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create virtual_instrument
	virtual_instrumentDB := orm.Virtual_instrumentDB{}
	virtual_instrumentDB.Virtual_instrumentPointersEncoding = input.Virtual_instrumentPointersEncoding
	virtual_instrumentDB.CopyBasicFieldsFromVirtual_instrument_WOP(&input.Virtual_instrument_WOP)

	query := db.Create(&virtual_instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoVirtual_instrument.CheckoutPhaseOneInstance(&virtual_instrumentDB)
	virtual_instrument := backRepo.BackRepoVirtual_instrument.Map_Virtual_instrumentDBID_Virtual_instrumentPtr[virtual_instrumentDB.ID]

	if virtual_instrument != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), virtual_instrument)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, virtual_instrumentDB)
}

// GetVirtual_instrument
//
// swagger:route GET /virtual_instruments/{ID} virtual_instruments getVirtual_instrument
//
// Gets the details for a virtual_instrument.
//
// Responses:
// default: genericError
//
//	200: virtual_instrumentDBResponse
func (controller *Controller) GetVirtual_instrument(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetVirtual_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVirtual_instrument.GetDB()

	// Get virtual_instrumentDB in DB
	var virtual_instrumentDB orm.Virtual_instrumentDB
	if err := db.First(&virtual_instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var virtual_instrumentAPI orm.Virtual_instrumentAPI
	virtual_instrumentAPI.ID = virtual_instrumentDB.ID
	virtual_instrumentAPI.Virtual_instrumentPointersEncoding = virtual_instrumentDB.Virtual_instrumentPointersEncoding
	virtual_instrumentDB.CopyBasicFieldsToVirtual_instrument_WOP(&virtual_instrumentAPI.Virtual_instrument_WOP)

	c.JSON(http.StatusOK, virtual_instrumentAPI)
}

// UpdateVirtual_instrument
//
// swagger:route PATCH /virtual_instruments/{ID} virtual_instruments updateVirtual_instrument
//
// # Update a virtual_instrument
//
// Responses:
// default: genericError
//
//	200: virtual_instrumentDBResponse
func (controller *Controller) UpdateVirtual_instrument(c *gin.Context) {

	mutexVirtual_instrument.Lock()
	defer mutexVirtual_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateVirtual_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVirtual_instrument.GetDB()

	// Validate input
	var input orm.Virtual_instrumentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var virtual_instrumentDB orm.Virtual_instrumentDB

	// fetch the virtual_instrument
	query := db.First(&virtual_instrumentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	virtual_instrumentDB.CopyBasicFieldsFromVirtual_instrument_WOP(&input.Virtual_instrument_WOP)
	virtual_instrumentDB.Virtual_instrumentPointersEncoding = input.Virtual_instrumentPointersEncoding

	query = db.Model(&virtual_instrumentDB).Updates(virtual_instrumentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	virtual_instrumentNew := new(models.Virtual_instrument)
	virtual_instrumentDB.CopyBasicFieldsToVirtual_instrument(virtual_instrumentNew)

	// redeem pointers
	virtual_instrumentDB.DecodePointers(backRepo, virtual_instrumentNew)

	// get stage instance from DB instance, and call callback function
	virtual_instrumentOld := backRepo.BackRepoVirtual_instrument.Map_Virtual_instrumentDBID_Virtual_instrumentPtr[virtual_instrumentDB.ID]
	if virtual_instrumentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), virtual_instrumentOld, virtual_instrumentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the virtual_instrumentDB
	c.JSON(http.StatusOK, virtual_instrumentDB)
}

// DeleteVirtual_instrument
//
// swagger:route DELETE /virtual_instruments/{ID} virtual_instruments deleteVirtual_instrument
//
// # Delete a virtual_instrument
//
// default: genericError
//
//	200: virtual_instrumentDBResponse
func (controller *Controller) DeleteVirtual_instrument(c *gin.Context) {

	mutexVirtual_instrument.Lock()
	defer mutexVirtual_instrument.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteVirtual_instrument", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoVirtual_instrument.GetDB()

	// Get model if exist
	var virtual_instrumentDB orm.Virtual_instrumentDB
	if err := db.First(&virtual_instrumentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&virtual_instrumentDB)

	// get an instance (not staged) from DB instance, and call callback function
	virtual_instrumentDeleted := new(models.Virtual_instrument)
	virtual_instrumentDB.CopyBasicFieldsToVirtual_instrument(virtual_instrumentDeleted)

	// get stage instance from DB instance, and call callback function
	virtual_instrumentStaged := backRepo.BackRepoVirtual_instrument.Map_Virtual_instrumentDBID_Virtual_instrumentPtr[virtual_instrumentDB.ID]
	if virtual_instrumentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), virtual_instrumentStaged, virtual_instrumentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
