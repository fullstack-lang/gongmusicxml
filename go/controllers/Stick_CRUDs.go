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
var __Stick__dummysDeclaration__ models.Stick
var __Stick_time__dummyDeclaration time.Duration

var mutexStick sync.Mutex

// An StickID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStick updateStick deleteStick
type StickID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StickInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStick updateStick
type StickInput struct {
	// The Stick to submit or modify
	// in: body
	Stick *orm.StickAPI
}

// GetSticks
//
// swagger:route GET /sticks sticks getSticks
//
// # Get all sticks
//
// Responses:
// default: genericError
//
//	200: stickDBResponse
func (controller *Controller) GetSticks(c *gin.Context) {

	// source slice
	var stickDBs []orm.StickDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSticks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStick.GetDB()

	query := db.Find(&stickDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	stickAPIs := make([]orm.StickAPI, 0)

	// for each stick, update fields from the database nullable fields
	for idx := range stickDBs {
		stickDB := &stickDBs[idx]
		_ = stickDB
		var stickAPI orm.StickAPI

		// insertion point for updating fields
		stickAPI.ID = stickDB.ID
		stickDB.CopyBasicFieldsToStick_WOP(&stickAPI.Stick_WOP)
		stickAPI.StickPointersEncoding = stickDB.StickPointersEncoding
		stickAPIs = append(stickAPIs, stickAPI)
	}

	c.JSON(http.StatusOK, stickAPIs)
}

// PostStick
//
// swagger:route POST /sticks sticks postStick
//
// Creates a stick
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStick(c *gin.Context) {

	mutexStick.Lock()
	defer mutexStick.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSticks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStick.GetDB()

	// Validate input
	var input orm.StickAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create stick
	stickDB := orm.StickDB{}
	stickDB.StickPointersEncoding = input.StickPointersEncoding
	stickDB.CopyBasicFieldsFromStick_WOP(&input.Stick_WOP)

	query := db.Create(&stickDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStick.CheckoutPhaseOneInstance(&stickDB)
	stick := backRepo.BackRepoStick.Map_StickDBID_StickPtr[stickDB.ID]

	if stick != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), stick)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, stickDB)
}

// GetStick
//
// swagger:route GET /sticks/{ID} sticks getStick
//
// Gets the details for a stick.
//
// Responses:
// default: genericError
//
//	200: stickDBResponse
func (controller *Controller) GetStick(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStick", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStick.GetDB()

	// Get stickDB in DB
	var stickDB orm.StickDB
	if err := db.First(&stickDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var stickAPI orm.StickAPI
	stickAPI.ID = stickDB.ID
	stickAPI.StickPointersEncoding = stickDB.StickPointersEncoding
	stickDB.CopyBasicFieldsToStick_WOP(&stickAPI.Stick_WOP)

	c.JSON(http.StatusOK, stickAPI)
}

// UpdateStick
//
// swagger:route PATCH /sticks/{ID} sticks updateStick
//
// # Update a stick
//
// Responses:
// default: genericError
//
//	200: stickDBResponse
func (controller *Controller) UpdateStick(c *gin.Context) {

	mutexStick.Lock()
	defer mutexStick.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStick", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStick.GetDB()

	// Validate input
	var input orm.StickAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var stickDB orm.StickDB

	// fetch the stick
	query := db.First(&stickDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	stickDB.CopyBasicFieldsFromStick_WOP(&input.Stick_WOP)
	stickDB.StickPointersEncoding = input.StickPointersEncoding

	query = db.Model(&stickDB).Updates(stickDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	stickNew := new(models.Stick)
	stickDB.CopyBasicFieldsToStick(stickNew)

	// redeem pointers
	stickDB.DecodePointers(backRepo, stickNew)

	// get stage instance from DB instance, and call callback function
	stickOld := backRepo.BackRepoStick.Map_StickDBID_StickPtr[stickDB.ID]
	if stickOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), stickOld, stickNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the stickDB
	c.JSON(http.StatusOK, stickDB)
}

// DeleteStick
//
// swagger:route DELETE /sticks/{ID} sticks deleteStick
//
// # Delete a stick
//
// default: genericError
//
//	200: stickDBResponse
func (controller *Controller) DeleteStick(c *gin.Context) {

	mutexStick.Lock()
	defer mutexStick.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStick", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStick.GetDB()

	// Get model if exist
	var stickDB orm.StickDB
	if err := db.First(&stickDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&stickDB)

	// get an instance (not staged) from DB instance, and call callback function
	stickDeleted := new(models.Stick)
	stickDB.CopyBasicFieldsToStick(stickDeleted)

	// get stage instance from DB instance, and call callback function
	stickStaged := backRepo.BackRepoStick.Map_StickDBID_StickPtr[stickDB.ID]
	if stickStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), stickStaged, stickDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
