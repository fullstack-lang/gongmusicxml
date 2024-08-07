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
var __Octave_shift__dummysDeclaration__ models.Octave_shift
var __Octave_shift_time__dummyDeclaration time.Duration

var mutexOctave_shift sync.Mutex

// An Octave_shiftID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOctave_shift updateOctave_shift deleteOctave_shift
type Octave_shiftID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Octave_shiftInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOctave_shift updateOctave_shift
type Octave_shiftInput struct {
	// The Octave_shift to submit or modify
	// in: body
	Octave_shift *orm.Octave_shiftAPI
}

// GetOctave_shifts
//
// swagger:route GET /octave_shifts octave_shifts getOctave_shifts
//
// # Get all octave_shifts
//
// Responses:
// default: genericError
//
//	200: octave_shiftDBResponse
func (controller *Controller) GetOctave_shifts(c *gin.Context) {

	// source slice
	var octave_shiftDBs []orm.Octave_shiftDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOctave_shifts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOctave_shift.GetDB()

	query := db.Find(&octave_shiftDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	octave_shiftAPIs := make([]orm.Octave_shiftAPI, 0)

	// for each octave_shift, update fields from the database nullable fields
	for idx := range octave_shiftDBs {
		octave_shiftDB := &octave_shiftDBs[idx]
		_ = octave_shiftDB
		var octave_shiftAPI orm.Octave_shiftAPI

		// insertion point for updating fields
		octave_shiftAPI.ID = octave_shiftDB.ID
		octave_shiftDB.CopyBasicFieldsToOctave_shift_WOP(&octave_shiftAPI.Octave_shift_WOP)
		octave_shiftAPI.Octave_shiftPointersEncoding = octave_shiftDB.Octave_shiftPointersEncoding
		octave_shiftAPIs = append(octave_shiftAPIs, octave_shiftAPI)
	}

	c.JSON(http.StatusOK, octave_shiftAPIs)
}

// PostOctave_shift
//
// swagger:route POST /octave_shifts octave_shifts postOctave_shift
//
// Creates a octave_shift
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOctave_shift(c *gin.Context) {

	mutexOctave_shift.Lock()
	defer mutexOctave_shift.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOctave_shifts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOctave_shift.GetDB()

	// Validate input
	var input orm.Octave_shiftAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create octave_shift
	octave_shiftDB := orm.Octave_shiftDB{}
	octave_shiftDB.Octave_shiftPointersEncoding = input.Octave_shiftPointersEncoding
	octave_shiftDB.CopyBasicFieldsFromOctave_shift_WOP(&input.Octave_shift_WOP)

	query := db.Create(&octave_shiftDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOctave_shift.CheckoutPhaseOneInstance(&octave_shiftDB)
	octave_shift := backRepo.BackRepoOctave_shift.Map_Octave_shiftDBID_Octave_shiftPtr[octave_shiftDB.ID]

	if octave_shift != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), octave_shift)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, octave_shiftDB)
}

// GetOctave_shift
//
// swagger:route GET /octave_shifts/{ID} octave_shifts getOctave_shift
//
// Gets the details for a octave_shift.
//
// Responses:
// default: genericError
//
//	200: octave_shiftDBResponse
func (controller *Controller) GetOctave_shift(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOctave_shift", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOctave_shift.GetDB()

	// Get octave_shiftDB in DB
	var octave_shiftDB orm.Octave_shiftDB
	if err := db.First(&octave_shiftDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var octave_shiftAPI orm.Octave_shiftAPI
	octave_shiftAPI.ID = octave_shiftDB.ID
	octave_shiftAPI.Octave_shiftPointersEncoding = octave_shiftDB.Octave_shiftPointersEncoding
	octave_shiftDB.CopyBasicFieldsToOctave_shift_WOP(&octave_shiftAPI.Octave_shift_WOP)

	c.JSON(http.StatusOK, octave_shiftAPI)
}

// UpdateOctave_shift
//
// swagger:route PATCH /octave_shifts/{ID} octave_shifts updateOctave_shift
//
// # Update a octave_shift
//
// Responses:
// default: genericError
//
//	200: octave_shiftDBResponse
func (controller *Controller) UpdateOctave_shift(c *gin.Context) {

	mutexOctave_shift.Lock()
	defer mutexOctave_shift.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOctave_shift", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOctave_shift.GetDB()

	// Validate input
	var input orm.Octave_shiftAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var octave_shiftDB orm.Octave_shiftDB

	// fetch the octave_shift
	query := db.First(&octave_shiftDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	octave_shiftDB.CopyBasicFieldsFromOctave_shift_WOP(&input.Octave_shift_WOP)
	octave_shiftDB.Octave_shiftPointersEncoding = input.Octave_shiftPointersEncoding

	query = db.Model(&octave_shiftDB).Updates(octave_shiftDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	octave_shiftNew := new(models.Octave_shift)
	octave_shiftDB.CopyBasicFieldsToOctave_shift(octave_shiftNew)

	// redeem pointers
	octave_shiftDB.DecodePointers(backRepo, octave_shiftNew)

	// get stage instance from DB instance, and call callback function
	octave_shiftOld := backRepo.BackRepoOctave_shift.Map_Octave_shiftDBID_Octave_shiftPtr[octave_shiftDB.ID]
	if octave_shiftOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), octave_shiftOld, octave_shiftNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the octave_shiftDB
	c.JSON(http.StatusOK, octave_shiftDB)
}

// DeleteOctave_shift
//
// swagger:route DELETE /octave_shifts/{ID} octave_shifts deleteOctave_shift
//
// # Delete a octave_shift
//
// default: genericError
//
//	200: octave_shiftDBResponse
func (controller *Controller) DeleteOctave_shift(c *gin.Context) {

	mutexOctave_shift.Lock()
	defer mutexOctave_shift.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOctave_shift", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOctave_shift.GetDB()

	// Get model if exist
	var octave_shiftDB orm.Octave_shiftDB
	if err := db.First(&octave_shiftDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&octave_shiftDB)

	// get an instance (not staged) from DB instance, and call callback function
	octave_shiftDeleted := new(models.Octave_shift)
	octave_shiftDB.CopyBasicFieldsToOctave_shift(octave_shiftDeleted)

	// get stage instance from DB instance, and call callback function
	octave_shiftStaged := backRepo.BackRepoOctave_shift.Map_Octave_shiftDBID_Octave_shiftPtr[octave_shiftDB.ID]
	if octave_shiftStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), octave_shiftStaged, octave_shiftDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
