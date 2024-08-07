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
var __Pitch__dummysDeclaration__ models.Pitch
var __Pitch_time__dummyDeclaration time.Duration

var mutexPitch sync.Mutex

// An PitchID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPitch updatePitch deletePitch
type PitchID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PitchInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPitch updatePitch
type PitchInput struct {
	// The Pitch to submit or modify
	// in: body
	Pitch *orm.PitchAPI
}

// GetPitchs
//
// swagger:route GET /pitchs pitchs getPitchs
//
// # Get all pitchs
//
// Responses:
// default: genericError
//
//	200: pitchDBResponse
func (controller *Controller) GetPitchs(c *gin.Context) {

	// source slice
	var pitchDBs []orm.PitchDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPitchs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitch.GetDB()

	query := db.Find(&pitchDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pitchAPIs := make([]orm.PitchAPI, 0)

	// for each pitch, update fields from the database nullable fields
	for idx := range pitchDBs {
		pitchDB := &pitchDBs[idx]
		_ = pitchDB
		var pitchAPI orm.PitchAPI

		// insertion point for updating fields
		pitchAPI.ID = pitchDB.ID
		pitchDB.CopyBasicFieldsToPitch_WOP(&pitchAPI.Pitch_WOP)
		pitchAPI.PitchPointersEncoding = pitchDB.PitchPointersEncoding
		pitchAPIs = append(pitchAPIs, pitchAPI)
	}

	c.JSON(http.StatusOK, pitchAPIs)
}

// PostPitch
//
// swagger:route POST /pitchs pitchs postPitch
//
// Creates a pitch
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPitch(c *gin.Context) {

	mutexPitch.Lock()
	defer mutexPitch.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPitchs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitch.GetDB()

	// Validate input
	var input orm.PitchAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pitch
	pitchDB := orm.PitchDB{}
	pitchDB.PitchPointersEncoding = input.PitchPointersEncoding
	pitchDB.CopyBasicFieldsFromPitch_WOP(&input.Pitch_WOP)

	query := db.Create(&pitchDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPitch.CheckoutPhaseOneInstance(&pitchDB)
	pitch := backRepo.BackRepoPitch.Map_PitchDBID_PitchPtr[pitchDB.ID]

	if pitch != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pitch)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pitchDB)
}

// GetPitch
//
// swagger:route GET /pitchs/{ID} pitchs getPitch
//
// Gets the details for a pitch.
//
// Responses:
// default: genericError
//
//	200: pitchDBResponse
func (controller *Controller) GetPitch(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPitch", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitch.GetDB()

	// Get pitchDB in DB
	var pitchDB orm.PitchDB
	if err := db.First(&pitchDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pitchAPI orm.PitchAPI
	pitchAPI.ID = pitchDB.ID
	pitchAPI.PitchPointersEncoding = pitchDB.PitchPointersEncoding
	pitchDB.CopyBasicFieldsToPitch_WOP(&pitchAPI.Pitch_WOP)

	c.JSON(http.StatusOK, pitchAPI)
}

// UpdatePitch
//
// swagger:route PATCH /pitchs/{ID} pitchs updatePitch
//
// # Update a pitch
//
// Responses:
// default: genericError
//
//	200: pitchDBResponse
func (controller *Controller) UpdatePitch(c *gin.Context) {

	mutexPitch.Lock()
	defer mutexPitch.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePitch", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitch.GetDB()

	// Validate input
	var input orm.PitchAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pitchDB orm.PitchDB

	// fetch the pitch
	query := db.First(&pitchDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pitchDB.CopyBasicFieldsFromPitch_WOP(&input.Pitch_WOP)
	pitchDB.PitchPointersEncoding = input.PitchPointersEncoding

	query = db.Model(&pitchDB).Updates(pitchDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pitchNew := new(models.Pitch)
	pitchDB.CopyBasicFieldsToPitch(pitchNew)

	// redeem pointers
	pitchDB.DecodePointers(backRepo, pitchNew)

	// get stage instance from DB instance, and call callback function
	pitchOld := backRepo.BackRepoPitch.Map_PitchDBID_PitchPtr[pitchDB.ID]
	if pitchOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pitchOld, pitchNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pitchDB
	c.JSON(http.StatusOK, pitchDB)
}

// DeletePitch
//
// swagger:route DELETE /pitchs/{ID} pitchs deletePitch
//
// # Delete a pitch
//
// default: genericError
//
//	200: pitchDBResponse
func (controller *Controller) DeletePitch(c *gin.Context) {

	mutexPitch.Lock()
	defer mutexPitch.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePitch", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPitch.GetDB()

	// Get model if exist
	var pitchDB orm.PitchDB
	if err := db.First(&pitchDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&pitchDB)

	// get an instance (not staged) from DB instance, and call callback function
	pitchDeleted := new(models.Pitch)
	pitchDB.CopyBasicFieldsToPitch(pitchDeleted)

	// get stage instance from DB instance, and call callback function
	pitchStaged := backRepo.BackRepoPitch.Map_PitchDBID_PitchPtr[pitchDB.ID]
	if pitchStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pitchStaged, pitchDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
