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
var __Empty_trill_sound__dummysDeclaration__ models.Empty_trill_sound
var __Empty_trill_sound_time__dummyDeclaration time.Duration

var mutexEmpty_trill_sound sync.Mutex

// An Empty_trill_soundID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_trill_sound updateEmpty_trill_sound deleteEmpty_trill_sound
type Empty_trill_soundID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_trill_soundInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_trill_sound updateEmpty_trill_sound
type Empty_trill_soundInput struct {
	// The Empty_trill_sound to submit or modify
	// in: body
	Empty_trill_sound *orm.Empty_trill_soundAPI
}

// GetEmpty_trill_sounds
//
// swagger:route GET /empty_trill_sounds empty_trill_sounds getEmpty_trill_sounds
//
// # Get all empty_trill_sounds
//
// Responses:
// default: genericError
//
//	200: empty_trill_soundDBResponse
func (controller *Controller) GetEmpty_trill_sounds(c *gin.Context) {

	// source slice
	var empty_trill_soundDBs []orm.Empty_trill_soundDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_trill_sounds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_trill_sound.GetDB()

	query := db.Find(&empty_trill_soundDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_trill_soundAPIs := make([]orm.Empty_trill_soundAPI, 0)

	// for each empty_trill_sound, update fields from the database nullable fields
	for idx := range empty_trill_soundDBs {
		empty_trill_soundDB := &empty_trill_soundDBs[idx]
		_ = empty_trill_soundDB
		var empty_trill_soundAPI orm.Empty_trill_soundAPI

		// insertion point for updating fields
		empty_trill_soundAPI.ID = empty_trill_soundDB.ID
		empty_trill_soundDB.CopyBasicFieldsToEmpty_trill_sound_WOP(&empty_trill_soundAPI.Empty_trill_sound_WOP)
		empty_trill_soundAPI.Empty_trill_soundPointersEncoding = empty_trill_soundDB.Empty_trill_soundPointersEncoding
		empty_trill_soundAPIs = append(empty_trill_soundAPIs, empty_trill_soundAPI)
	}

	c.JSON(http.StatusOK, empty_trill_soundAPIs)
}

// PostEmpty_trill_sound
//
// swagger:route POST /empty_trill_sounds empty_trill_sounds postEmpty_trill_sound
//
// Creates a empty_trill_sound
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_trill_sound(c *gin.Context) {

	mutexEmpty_trill_sound.Lock()
	defer mutexEmpty_trill_sound.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_trill_sounds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_trill_sound.GetDB()

	// Validate input
	var input orm.Empty_trill_soundAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_trill_sound
	empty_trill_soundDB := orm.Empty_trill_soundDB{}
	empty_trill_soundDB.Empty_trill_soundPointersEncoding = input.Empty_trill_soundPointersEncoding
	empty_trill_soundDB.CopyBasicFieldsFromEmpty_trill_sound_WOP(&input.Empty_trill_sound_WOP)

	query := db.Create(&empty_trill_soundDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_trill_sound.CheckoutPhaseOneInstance(&empty_trill_soundDB)
	empty_trill_sound := backRepo.BackRepoEmpty_trill_sound.Map_Empty_trill_soundDBID_Empty_trill_soundPtr[empty_trill_soundDB.ID]

	if empty_trill_sound != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_trill_sound)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_trill_soundDB)
}

// GetEmpty_trill_sound
//
// swagger:route GET /empty_trill_sounds/{ID} empty_trill_sounds getEmpty_trill_sound
//
// Gets the details for a empty_trill_sound.
//
// Responses:
// default: genericError
//
//	200: empty_trill_soundDBResponse
func (controller *Controller) GetEmpty_trill_sound(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_trill_sound", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_trill_sound.GetDB()

	// Get empty_trill_soundDB in DB
	var empty_trill_soundDB orm.Empty_trill_soundDB
	if err := db.First(&empty_trill_soundDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_trill_soundAPI orm.Empty_trill_soundAPI
	empty_trill_soundAPI.ID = empty_trill_soundDB.ID
	empty_trill_soundAPI.Empty_trill_soundPointersEncoding = empty_trill_soundDB.Empty_trill_soundPointersEncoding
	empty_trill_soundDB.CopyBasicFieldsToEmpty_trill_sound_WOP(&empty_trill_soundAPI.Empty_trill_sound_WOP)

	c.JSON(http.StatusOK, empty_trill_soundAPI)
}

// UpdateEmpty_trill_sound
//
// swagger:route PATCH /empty_trill_sounds/{ID} empty_trill_sounds updateEmpty_trill_sound
//
// # Update a empty_trill_sound
//
// Responses:
// default: genericError
//
//	200: empty_trill_soundDBResponse
func (controller *Controller) UpdateEmpty_trill_sound(c *gin.Context) {

	mutexEmpty_trill_sound.Lock()
	defer mutexEmpty_trill_sound.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_trill_sound", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_trill_sound.GetDB()

	// Validate input
	var input orm.Empty_trill_soundAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_trill_soundDB orm.Empty_trill_soundDB

	// fetch the empty_trill_sound
	query := db.First(&empty_trill_soundDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_trill_soundDB.CopyBasicFieldsFromEmpty_trill_sound_WOP(&input.Empty_trill_sound_WOP)
	empty_trill_soundDB.Empty_trill_soundPointersEncoding = input.Empty_trill_soundPointersEncoding

	query = db.Model(&empty_trill_soundDB).Updates(empty_trill_soundDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_trill_soundNew := new(models.Empty_trill_sound)
	empty_trill_soundDB.CopyBasicFieldsToEmpty_trill_sound(empty_trill_soundNew)

	// redeem pointers
	empty_trill_soundDB.DecodePointers(backRepo, empty_trill_soundNew)

	// get stage instance from DB instance, and call callback function
	empty_trill_soundOld := backRepo.BackRepoEmpty_trill_sound.Map_Empty_trill_soundDBID_Empty_trill_soundPtr[empty_trill_soundDB.ID]
	if empty_trill_soundOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_trill_soundOld, empty_trill_soundNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_trill_soundDB
	c.JSON(http.StatusOK, empty_trill_soundDB)
}

// DeleteEmpty_trill_sound
//
// swagger:route DELETE /empty_trill_sounds/{ID} empty_trill_sounds deleteEmpty_trill_sound
//
// # Delete a empty_trill_sound
//
// default: genericError
//
//	200: empty_trill_soundDBResponse
func (controller *Controller) DeleteEmpty_trill_sound(c *gin.Context) {

	mutexEmpty_trill_sound.Lock()
	defer mutexEmpty_trill_sound.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_trill_sound", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_trill_sound.GetDB()

	// Get model if exist
	var empty_trill_soundDB orm.Empty_trill_soundDB
	if err := db.First(&empty_trill_soundDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_trill_soundDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_trill_soundDeleted := new(models.Empty_trill_sound)
	empty_trill_soundDB.CopyBasicFieldsToEmpty_trill_sound(empty_trill_soundDeleted)

	// get stage instance from DB instance, and call callback function
	empty_trill_soundStaged := backRepo.BackRepoEmpty_trill_sound.Map_Empty_trill_soundDBID_Empty_trill_soundPtr[empty_trill_soundDB.ID]
	if empty_trill_soundStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_trill_soundStaged, empty_trill_soundDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
