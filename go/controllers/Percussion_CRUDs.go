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
var __Percussion__dummysDeclaration__ models.Percussion
var __Percussion_time__dummyDeclaration time.Duration

var mutexPercussion sync.Mutex

// An PercussionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPercussion updatePercussion deletePercussion
type PercussionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PercussionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPercussion updatePercussion
type PercussionInput struct {
	// The Percussion to submit or modify
	// in: body
	Percussion *orm.PercussionAPI
}

// GetPercussions
//
// swagger:route GET /percussions percussions getPercussions
//
// # Get all percussions
//
// Responses:
// default: genericError
//
//	200: percussionDBResponse
func (controller *Controller) GetPercussions(c *gin.Context) {

	// source slice
	var percussionDBs []orm.PercussionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPercussions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPercussion.GetDB()

	query := db.Find(&percussionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	percussionAPIs := make([]orm.PercussionAPI, 0)

	// for each percussion, update fields from the database nullable fields
	for idx := range percussionDBs {
		percussionDB := &percussionDBs[idx]
		_ = percussionDB
		var percussionAPI orm.PercussionAPI

		// insertion point for updating fields
		percussionAPI.ID = percussionDB.ID
		percussionDB.CopyBasicFieldsToPercussion_WOP(&percussionAPI.Percussion_WOP)
		percussionAPI.PercussionPointersEncoding = percussionDB.PercussionPointersEncoding
		percussionAPIs = append(percussionAPIs, percussionAPI)
	}

	c.JSON(http.StatusOK, percussionAPIs)
}

// PostPercussion
//
// swagger:route POST /percussions percussions postPercussion
//
// Creates a percussion
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPercussion(c *gin.Context) {

	mutexPercussion.Lock()
	defer mutexPercussion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPercussions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPercussion.GetDB()

	// Validate input
	var input orm.PercussionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create percussion
	percussionDB := orm.PercussionDB{}
	percussionDB.PercussionPointersEncoding = input.PercussionPointersEncoding
	percussionDB.CopyBasicFieldsFromPercussion_WOP(&input.Percussion_WOP)

	query := db.Create(&percussionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPercussion.CheckoutPhaseOneInstance(&percussionDB)
	percussion := backRepo.BackRepoPercussion.Map_PercussionDBID_PercussionPtr[percussionDB.ID]

	if percussion != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), percussion)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, percussionDB)
}

// GetPercussion
//
// swagger:route GET /percussions/{ID} percussions getPercussion
//
// Gets the details for a percussion.
//
// Responses:
// default: genericError
//
//	200: percussionDBResponse
func (controller *Controller) GetPercussion(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPercussion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPercussion.GetDB()

	// Get percussionDB in DB
	var percussionDB orm.PercussionDB
	if err := db.First(&percussionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var percussionAPI orm.PercussionAPI
	percussionAPI.ID = percussionDB.ID
	percussionAPI.PercussionPointersEncoding = percussionDB.PercussionPointersEncoding
	percussionDB.CopyBasicFieldsToPercussion_WOP(&percussionAPI.Percussion_WOP)

	c.JSON(http.StatusOK, percussionAPI)
}

// UpdatePercussion
//
// swagger:route PATCH /percussions/{ID} percussions updatePercussion
//
// # Update a percussion
//
// Responses:
// default: genericError
//
//	200: percussionDBResponse
func (controller *Controller) UpdatePercussion(c *gin.Context) {

	mutexPercussion.Lock()
	defer mutexPercussion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePercussion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPercussion.GetDB()

	// Validate input
	var input orm.PercussionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var percussionDB orm.PercussionDB

	// fetch the percussion
	query := db.First(&percussionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	percussionDB.CopyBasicFieldsFromPercussion_WOP(&input.Percussion_WOP)
	percussionDB.PercussionPointersEncoding = input.PercussionPointersEncoding

	query = db.Model(&percussionDB).Updates(percussionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	percussionNew := new(models.Percussion)
	percussionDB.CopyBasicFieldsToPercussion(percussionNew)

	// redeem pointers
	percussionDB.DecodePointers(backRepo, percussionNew)

	// get stage instance from DB instance, and call callback function
	percussionOld := backRepo.BackRepoPercussion.Map_PercussionDBID_PercussionPtr[percussionDB.ID]
	if percussionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), percussionOld, percussionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the percussionDB
	c.JSON(http.StatusOK, percussionDB)
}

// DeletePercussion
//
// swagger:route DELETE /percussions/{ID} percussions deletePercussion
//
// # Delete a percussion
//
// default: genericError
//
//	200: percussionDBResponse
func (controller *Controller) DeletePercussion(c *gin.Context) {

	mutexPercussion.Lock()
	defer mutexPercussion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePercussion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPercussion.GetDB()

	// Get model if exist
	var percussionDB orm.PercussionDB
	if err := db.First(&percussionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&percussionDB)

	// get an instance (not staged) from DB instance, and call callback function
	percussionDeleted := new(models.Percussion)
	percussionDB.CopyBasicFieldsToPercussion(percussionDeleted)

	// get stage instance from DB instance, and call callback function
	percussionStaged := backRepo.BackRepoPercussion.Map_PercussionDBID_PercussionPtr[percussionDB.ID]
	if percussionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), percussionStaged, percussionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
