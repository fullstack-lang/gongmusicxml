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
var __Work__dummysDeclaration__ models.Work
var __Work_time__dummyDeclaration time.Duration

var mutexWork sync.Mutex

// An WorkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWork updateWork deleteWork
type WorkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WorkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWork updateWork
type WorkInput struct {
	// The Work to submit or modify
	// in: body
	Work *orm.WorkAPI
}

// GetWorks
//
// swagger:route GET /works works getWorks
//
// # Get all works
//
// Responses:
// default: genericError
//
//	200: workDBResponse
func (controller *Controller) GetWorks(c *gin.Context) {

	// source slice
	var workDBs []orm.WorkDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWorks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWork.GetDB()

	query := db.Find(&workDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	workAPIs := make([]orm.WorkAPI, 0)

	// for each work, update fields from the database nullable fields
	for idx := range workDBs {
		workDB := &workDBs[idx]
		_ = workDB
		var workAPI orm.WorkAPI

		// insertion point for updating fields
		workAPI.ID = workDB.ID
		workDB.CopyBasicFieldsToWork_WOP(&workAPI.Work_WOP)
		workAPI.WorkPointersEncoding = workDB.WorkPointersEncoding
		workAPIs = append(workAPIs, workAPI)
	}

	c.JSON(http.StatusOK, workAPIs)
}

// PostWork
//
// swagger:route POST /works works postWork
//
// Creates a work
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWork(c *gin.Context) {

	mutexWork.Lock()
	defer mutexWork.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWorks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWork.GetDB()

	// Validate input
	var input orm.WorkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create work
	workDB := orm.WorkDB{}
	workDB.WorkPointersEncoding = input.WorkPointersEncoding
	workDB.CopyBasicFieldsFromWork_WOP(&input.Work_WOP)

	query := db.Create(&workDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWork.CheckoutPhaseOneInstance(&workDB)
	work := backRepo.BackRepoWork.Map_WorkDBID_WorkPtr[workDB.ID]

	if work != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), work)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, workDB)
}

// GetWork
//
// swagger:route GET /works/{ID} works getWork
//
// Gets the details for a work.
//
// Responses:
// default: genericError
//
//	200: workDBResponse
func (controller *Controller) GetWork(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWork", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWork.GetDB()

	// Get workDB in DB
	var workDB orm.WorkDB
	if err := db.First(&workDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var workAPI orm.WorkAPI
	workAPI.ID = workDB.ID
	workAPI.WorkPointersEncoding = workDB.WorkPointersEncoding
	workDB.CopyBasicFieldsToWork_WOP(&workAPI.Work_WOP)

	c.JSON(http.StatusOK, workAPI)
}

// UpdateWork
//
// swagger:route PATCH /works/{ID} works updateWork
//
// # Update a work
//
// Responses:
// default: genericError
//
//	200: workDBResponse
func (controller *Controller) UpdateWork(c *gin.Context) {

	mutexWork.Lock()
	defer mutexWork.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWork", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWork.GetDB()

	// Validate input
	var input orm.WorkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var workDB orm.WorkDB

	// fetch the work
	query := db.First(&workDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	workDB.CopyBasicFieldsFromWork_WOP(&input.Work_WOP)
	workDB.WorkPointersEncoding = input.WorkPointersEncoding

	query = db.Model(&workDB).Updates(workDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	workNew := new(models.Work)
	workDB.CopyBasicFieldsToWork(workNew)

	// redeem pointers
	workDB.DecodePointers(backRepo, workNew)

	// get stage instance from DB instance, and call callback function
	workOld := backRepo.BackRepoWork.Map_WorkDBID_WorkPtr[workDB.ID]
	if workOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), workOld, workNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the workDB
	c.JSON(http.StatusOK, workDB)
}

// DeleteWork
//
// swagger:route DELETE /works/{ID} works deleteWork
//
// # Delete a work
//
// default: genericError
//
//	200: workDBResponse
func (controller *Controller) DeleteWork(c *gin.Context) {

	mutexWork.Lock()
	defer mutexWork.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWork", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWork.GetDB()

	// Get model if exist
	var workDB orm.WorkDB
	if err := db.First(&workDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&workDB)

	// get an instance (not staged) from DB instance, and call callback function
	workDeleted := new(models.Work)
	workDB.CopyBasicFieldsToWork(workDeleted)

	// get stage instance from DB instance, and call callback function
	workStaged := backRepo.BackRepoWork.Map_WorkDBID_WorkPtr[workDB.ID]
	if workStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), workStaged, workDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
