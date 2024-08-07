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
var __Mordent__dummysDeclaration__ models.Mordent
var __Mordent_time__dummyDeclaration time.Duration

var mutexMordent sync.Mutex

// An MordentID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMordent updateMordent deleteMordent
type MordentID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MordentInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMordent updateMordent
type MordentInput struct {
	// The Mordent to submit or modify
	// in: body
	Mordent *orm.MordentAPI
}

// GetMordents
//
// swagger:route GET /mordents mordents getMordents
//
// # Get all mordents
//
// Responses:
// default: genericError
//
//	200: mordentDBResponse
func (controller *Controller) GetMordents(c *gin.Context) {

	// source slice
	var mordentDBs []orm.MordentDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMordents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMordent.GetDB()

	query := db.Find(&mordentDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	mordentAPIs := make([]orm.MordentAPI, 0)

	// for each mordent, update fields from the database nullable fields
	for idx := range mordentDBs {
		mordentDB := &mordentDBs[idx]
		_ = mordentDB
		var mordentAPI orm.MordentAPI

		// insertion point for updating fields
		mordentAPI.ID = mordentDB.ID
		mordentDB.CopyBasicFieldsToMordent_WOP(&mordentAPI.Mordent_WOP)
		mordentAPI.MordentPointersEncoding = mordentDB.MordentPointersEncoding
		mordentAPIs = append(mordentAPIs, mordentAPI)
	}

	c.JSON(http.StatusOK, mordentAPIs)
}

// PostMordent
//
// swagger:route POST /mordents mordents postMordent
//
// Creates a mordent
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMordent(c *gin.Context) {

	mutexMordent.Lock()
	defer mutexMordent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMordents", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMordent.GetDB()

	// Validate input
	var input orm.MordentAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create mordent
	mordentDB := orm.MordentDB{}
	mordentDB.MordentPointersEncoding = input.MordentPointersEncoding
	mordentDB.CopyBasicFieldsFromMordent_WOP(&input.Mordent_WOP)

	query := db.Create(&mordentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMordent.CheckoutPhaseOneInstance(&mordentDB)
	mordent := backRepo.BackRepoMordent.Map_MordentDBID_MordentPtr[mordentDB.ID]

	if mordent != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), mordent)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, mordentDB)
}

// GetMordent
//
// swagger:route GET /mordents/{ID} mordents getMordent
//
// Gets the details for a mordent.
//
// Responses:
// default: genericError
//
//	200: mordentDBResponse
func (controller *Controller) GetMordent(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMordent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMordent.GetDB()

	// Get mordentDB in DB
	var mordentDB orm.MordentDB
	if err := db.First(&mordentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var mordentAPI orm.MordentAPI
	mordentAPI.ID = mordentDB.ID
	mordentAPI.MordentPointersEncoding = mordentDB.MordentPointersEncoding
	mordentDB.CopyBasicFieldsToMordent_WOP(&mordentAPI.Mordent_WOP)

	c.JSON(http.StatusOK, mordentAPI)
}

// UpdateMordent
//
// swagger:route PATCH /mordents/{ID} mordents updateMordent
//
// # Update a mordent
//
// Responses:
// default: genericError
//
//	200: mordentDBResponse
func (controller *Controller) UpdateMordent(c *gin.Context) {

	mutexMordent.Lock()
	defer mutexMordent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMordent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMordent.GetDB()

	// Validate input
	var input orm.MordentAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var mordentDB orm.MordentDB

	// fetch the mordent
	query := db.First(&mordentDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	mordentDB.CopyBasicFieldsFromMordent_WOP(&input.Mordent_WOP)
	mordentDB.MordentPointersEncoding = input.MordentPointersEncoding

	query = db.Model(&mordentDB).Updates(mordentDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	mordentNew := new(models.Mordent)
	mordentDB.CopyBasicFieldsToMordent(mordentNew)

	// redeem pointers
	mordentDB.DecodePointers(backRepo, mordentNew)

	// get stage instance from DB instance, and call callback function
	mordentOld := backRepo.BackRepoMordent.Map_MordentDBID_MordentPtr[mordentDB.ID]
	if mordentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), mordentOld, mordentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the mordentDB
	c.JSON(http.StatusOK, mordentDB)
}

// DeleteMordent
//
// swagger:route DELETE /mordents/{ID} mordents deleteMordent
//
// # Delete a mordent
//
// default: genericError
//
//	200: mordentDBResponse
func (controller *Controller) DeleteMordent(c *gin.Context) {

	mutexMordent.Lock()
	defer mutexMordent.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMordent", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMordent.GetDB()

	// Get model if exist
	var mordentDB orm.MordentDB
	if err := db.First(&mordentDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&mordentDB)

	// get an instance (not staged) from DB instance, and call callback function
	mordentDeleted := new(models.Mordent)
	mordentDB.CopyBasicFieldsToMordent(mordentDeleted)

	// get stage instance from DB instance, and call callback function
	mordentStaged := backRepo.BackRepoMordent.Map_MordentDBID_MordentPtr[mordentDB.ID]
	if mordentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), mordentStaged, mordentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
