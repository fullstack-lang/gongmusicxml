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
var __Pitched__dummysDeclaration__ models.Pitched
var __Pitched_time__dummyDeclaration time.Duration

var mutexPitched sync.Mutex

// An PitchedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPitched updatePitched deletePitched
type PitchedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PitchedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPitched updatePitched
type PitchedInput struct {
	// The Pitched to submit or modify
	// in: body
	Pitched *orm.PitchedAPI
}

// GetPitcheds
//
// swagger:route GET /pitcheds pitcheds getPitcheds
//
// # Get all pitcheds
//
// Responses:
// default: genericError
//
//	200: pitchedDBResponse
func (controller *Controller) GetPitcheds(c *gin.Context) {

	// source slice
	var pitchedDBs []orm.PitchedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPitcheds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitched.GetDB()

	query := db.Find(&pitchedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pitchedAPIs := make([]orm.PitchedAPI, 0)

	// for each pitched, update fields from the database nullable fields
	for idx := range pitchedDBs {
		pitchedDB := &pitchedDBs[idx]
		_ = pitchedDB
		var pitchedAPI orm.PitchedAPI

		// insertion point for updating fields
		pitchedAPI.ID = pitchedDB.ID
		pitchedDB.CopyBasicFieldsToPitched_WOP(&pitchedAPI.Pitched_WOP)
		pitchedAPI.PitchedPointersEncoding = pitchedDB.PitchedPointersEncoding
		pitchedAPIs = append(pitchedAPIs, pitchedAPI)
	}

	c.JSON(http.StatusOK, pitchedAPIs)
}

// PostPitched
//
// swagger:route POST /pitcheds pitcheds postPitched
//
// Creates a pitched
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPitched(c *gin.Context) {

	mutexPitched.Lock()
	defer mutexPitched.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPitcheds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitched.GetDB()

	// Validate input
	var input orm.PitchedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pitched
	pitchedDB := orm.PitchedDB{}
	pitchedDB.PitchedPointersEncoding = input.PitchedPointersEncoding
	pitchedDB.CopyBasicFieldsFromPitched_WOP(&input.Pitched_WOP)

	query := db.Create(&pitchedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPitched.CheckoutPhaseOneInstance(&pitchedDB)
	pitched := backRepo.BackRepoPitched.Map_PitchedDBID_PitchedPtr[pitchedDB.ID]

	if pitched != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pitched)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pitchedDB)
}

// GetPitched
//
// swagger:route GET /pitcheds/{ID} pitcheds getPitched
//
// Gets the details for a pitched.
//
// Responses:
// default: genericError
//
//	200: pitchedDBResponse
func (controller *Controller) GetPitched(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPitched", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitched.GetDB()

	// Get pitchedDB in DB
	var pitchedDB orm.PitchedDB
	if err := db.First(&pitchedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pitchedAPI orm.PitchedAPI
	pitchedAPI.ID = pitchedDB.ID
	pitchedAPI.PitchedPointersEncoding = pitchedDB.PitchedPointersEncoding
	pitchedDB.CopyBasicFieldsToPitched_WOP(&pitchedAPI.Pitched_WOP)

	c.JSON(http.StatusOK, pitchedAPI)
}

// UpdatePitched
//
// swagger:route PATCH /pitcheds/{ID} pitcheds updatePitched
//
// # Update a pitched
//
// Responses:
// default: genericError
//
//	200: pitchedDBResponse
func (controller *Controller) UpdatePitched(c *gin.Context) {

	mutexPitched.Lock()
	defer mutexPitched.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePitched", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitched.GetDB()

	// Validate input
	var input orm.PitchedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pitchedDB orm.PitchedDB

	// fetch the pitched
	query := db.First(&pitchedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pitchedDB.CopyBasicFieldsFromPitched_WOP(&input.Pitched_WOP)
	pitchedDB.PitchedPointersEncoding = input.PitchedPointersEncoding

	query = db.Model(&pitchedDB).Updates(pitchedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pitchedNew := new(models.Pitched)
	pitchedDB.CopyBasicFieldsToPitched(pitchedNew)

	// redeem pointers
	pitchedDB.DecodePointers(backRepo, pitchedNew)

	// get stage instance from DB instance, and call callback function
	pitchedOld := backRepo.BackRepoPitched.Map_PitchedDBID_PitchedPtr[pitchedDB.ID]
	if pitchedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pitchedOld, pitchedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pitchedDB
	c.JSON(http.StatusOK, pitchedDB)
}

// DeletePitched
//
// swagger:route DELETE /pitcheds/{ID} pitcheds deletePitched
//
// # Delete a pitched
//
// default: genericError
//
//	200: pitchedDBResponse
func (controller *Controller) DeletePitched(c *gin.Context) {

	mutexPitched.Lock()
	defer mutexPitched.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePitched", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitched.GetDB()

	// Get model if exist
	var pitchedDB orm.PitchedDB
	if err := db.First(&pitchedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&pitchedDB)

	// get an instance (not staged) from DB instance, and call callback function
	pitchedDeleted := new(models.Pitched)
	pitchedDB.CopyBasicFieldsToPitched(pitchedDeleted)

	// get stage instance from DB instance, and call callback function
	pitchedStaged := backRepo.BackRepoPitched.Map_PitchedDBID_PitchedPtr[pitchedDB.ID]
	if pitchedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pitchedStaged, pitchedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
