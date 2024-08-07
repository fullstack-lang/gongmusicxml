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
var __Bass_step__dummysDeclaration__ models.Bass_step
var __Bass_step_time__dummyDeclaration time.Duration

var mutexBass_step sync.Mutex

// An Bass_stepID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBass_step updateBass_step deleteBass_step
type Bass_stepID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Bass_stepInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBass_step updateBass_step
type Bass_stepInput struct {
	// The Bass_step to submit or modify
	// in: body
	Bass_step *orm.Bass_stepAPI
}

// GetBass_steps
//
// swagger:route GET /bass_steps bass_steps getBass_steps
//
// # Get all bass_steps
//
// Responses:
// default: genericError
//
//	200: bass_stepDBResponse
func (controller *Controller) GetBass_steps(c *gin.Context) {

	// source slice
	var bass_stepDBs []orm.Bass_stepDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBass_steps", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass_step.GetDB()

	query := db.Find(&bass_stepDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bass_stepAPIs := make([]orm.Bass_stepAPI, 0)

	// for each bass_step, update fields from the database nullable fields
	for idx := range bass_stepDBs {
		bass_stepDB := &bass_stepDBs[idx]
		_ = bass_stepDB
		var bass_stepAPI orm.Bass_stepAPI

		// insertion point for updating fields
		bass_stepAPI.ID = bass_stepDB.ID
		bass_stepDB.CopyBasicFieldsToBass_step_WOP(&bass_stepAPI.Bass_step_WOP)
		bass_stepAPI.Bass_stepPointersEncoding = bass_stepDB.Bass_stepPointersEncoding
		bass_stepAPIs = append(bass_stepAPIs, bass_stepAPI)
	}

	c.JSON(http.StatusOK, bass_stepAPIs)
}

// PostBass_step
//
// swagger:route POST /bass_steps bass_steps postBass_step
//
// Creates a bass_step
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBass_step(c *gin.Context) {

	mutexBass_step.Lock()
	defer mutexBass_step.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBass_steps", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass_step.GetDB()

	// Validate input
	var input orm.Bass_stepAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bass_step
	bass_stepDB := orm.Bass_stepDB{}
	bass_stepDB.Bass_stepPointersEncoding = input.Bass_stepPointersEncoding
	bass_stepDB.CopyBasicFieldsFromBass_step_WOP(&input.Bass_step_WOP)

	query := db.Create(&bass_stepDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBass_step.CheckoutPhaseOneInstance(&bass_stepDB)
	bass_step := backRepo.BackRepoBass_step.Map_Bass_stepDBID_Bass_stepPtr[bass_stepDB.ID]

	if bass_step != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bass_step)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bass_stepDB)
}

// GetBass_step
//
// swagger:route GET /bass_steps/{ID} bass_steps getBass_step
//
// Gets the details for a bass_step.
//
// Responses:
// default: genericError
//
//	200: bass_stepDBResponse
func (controller *Controller) GetBass_step(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBass_step", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass_step.GetDB()

	// Get bass_stepDB in DB
	var bass_stepDB orm.Bass_stepDB
	if err := db.First(&bass_stepDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bass_stepAPI orm.Bass_stepAPI
	bass_stepAPI.ID = bass_stepDB.ID
	bass_stepAPI.Bass_stepPointersEncoding = bass_stepDB.Bass_stepPointersEncoding
	bass_stepDB.CopyBasicFieldsToBass_step_WOP(&bass_stepAPI.Bass_step_WOP)

	c.JSON(http.StatusOK, bass_stepAPI)
}

// UpdateBass_step
//
// swagger:route PATCH /bass_steps/{ID} bass_steps updateBass_step
//
// # Update a bass_step
//
// Responses:
// default: genericError
//
//	200: bass_stepDBResponse
func (controller *Controller) UpdateBass_step(c *gin.Context) {

	mutexBass_step.Lock()
	defer mutexBass_step.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBass_step", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass_step.GetDB()

	// Validate input
	var input orm.Bass_stepAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bass_stepDB orm.Bass_stepDB

	// fetch the bass_step
	query := db.First(&bass_stepDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bass_stepDB.CopyBasicFieldsFromBass_step_WOP(&input.Bass_step_WOP)
	bass_stepDB.Bass_stepPointersEncoding = input.Bass_stepPointersEncoding

	query = db.Model(&bass_stepDB).Updates(bass_stepDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bass_stepNew := new(models.Bass_step)
	bass_stepDB.CopyBasicFieldsToBass_step(bass_stepNew)

	// redeem pointers
	bass_stepDB.DecodePointers(backRepo, bass_stepNew)

	// get stage instance from DB instance, and call callback function
	bass_stepOld := backRepo.BackRepoBass_step.Map_Bass_stepDBID_Bass_stepPtr[bass_stepDB.ID]
	if bass_stepOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bass_stepOld, bass_stepNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bass_stepDB
	c.JSON(http.StatusOK, bass_stepDB)
}

// DeleteBass_step
//
// swagger:route DELETE /bass_steps/{ID} bass_steps deleteBass_step
//
// # Delete a bass_step
//
// default: genericError
//
//	200: bass_stepDBResponse
func (controller *Controller) DeleteBass_step(c *gin.Context) {

	mutexBass_step.Lock()
	defer mutexBass_step.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBass_step", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass_step.GetDB()

	// Get model if exist
	var bass_stepDB orm.Bass_stepDB
	if err := db.First(&bass_stepDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bass_stepDB)

	// get an instance (not staged) from DB instance, and call callback function
	bass_stepDeleted := new(models.Bass_step)
	bass_stepDB.CopyBasicFieldsToBass_step(bass_stepDeleted)

	// get stage instance from DB instance, and call callback function
	bass_stepStaged := backRepo.BackRepoBass_step.Map_Bass_stepDBID_Bass_stepPtr[bass_stepDB.ID]
	if bass_stepStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bass_stepStaged, bass_stepDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
