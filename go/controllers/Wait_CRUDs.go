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
var __Wait__dummysDeclaration__ models.Wait
var __Wait_time__dummyDeclaration time.Duration

var mutexWait sync.Mutex

// An WaitID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWait updateWait deleteWait
type WaitID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WaitInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWait updateWait
type WaitInput struct {
	// The Wait to submit or modify
	// in: body
	Wait *orm.WaitAPI
}

// GetWaits
//
// swagger:route GET /waits waits getWaits
//
// # Get all waits
//
// Responses:
// default: genericError
//
//	200: waitDBResponse
func (controller *Controller) GetWaits(c *gin.Context) {

	// source slice
	var waitDBs []orm.WaitDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWaits", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWait.GetDB()

	query := db.Find(&waitDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	waitAPIs := make([]orm.WaitAPI, 0)

	// for each wait, update fields from the database nullable fields
	for idx := range waitDBs {
		waitDB := &waitDBs[idx]
		_ = waitDB
		var waitAPI orm.WaitAPI

		// insertion point for updating fields
		waitAPI.ID = waitDB.ID
		waitDB.CopyBasicFieldsToWait_WOP(&waitAPI.Wait_WOP)
		waitAPI.WaitPointersEncoding = waitDB.WaitPointersEncoding
		waitAPIs = append(waitAPIs, waitAPI)
	}

	c.JSON(http.StatusOK, waitAPIs)
}

// PostWait
//
// swagger:route POST /waits waits postWait
//
// Creates a wait
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWait(c *gin.Context) {

	mutexWait.Lock()
	defer mutexWait.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWaits", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWait.GetDB()

	// Validate input
	var input orm.WaitAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create wait
	waitDB := orm.WaitDB{}
	waitDB.WaitPointersEncoding = input.WaitPointersEncoding
	waitDB.CopyBasicFieldsFromWait_WOP(&input.Wait_WOP)

	query := db.Create(&waitDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWait.CheckoutPhaseOneInstance(&waitDB)
	wait := backRepo.BackRepoWait.Map_WaitDBID_WaitPtr[waitDB.ID]

	if wait != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), wait)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, waitDB)
}

// GetWait
//
// swagger:route GET /waits/{ID} waits getWait
//
// Gets the details for a wait.
//
// Responses:
// default: genericError
//
//	200: waitDBResponse
func (controller *Controller) GetWait(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWait", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWait.GetDB()

	// Get waitDB in DB
	var waitDB orm.WaitDB
	if err := db.First(&waitDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var waitAPI orm.WaitAPI
	waitAPI.ID = waitDB.ID
	waitAPI.WaitPointersEncoding = waitDB.WaitPointersEncoding
	waitDB.CopyBasicFieldsToWait_WOP(&waitAPI.Wait_WOP)

	c.JSON(http.StatusOK, waitAPI)
}

// UpdateWait
//
// swagger:route PATCH /waits/{ID} waits updateWait
//
// # Update a wait
//
// Responses:
// default: genericError
//
//	200: waitDBResponse
func (controller *Controller) UpdateWait(c *gin.Context) {

	mutexWait.Lock()
	defer mutexWait.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWait", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWait.GetDB()

	// Validate input
	var input orm.WaitAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var waitDB orm.WaitDB

	// fetch the wait
	query := db.First(&waitDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	waitDB.CopyBasicFieldsFromWait_WOP(&input.Wait_WOP)
	waitDB.WaitPointersEncoding = input.WaitPointersEncoding

	query = db.Model(&waitDB).Updates(waitDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	waitNew := new(models.Wait)
	waitDB.CopyBasicFieldsToWait(waitNew)

	// redeem pointers
	waitDB.DecodePointers(backRepo, waitNew)

	// get stage instance from DB instance, and call callback function
	waitOld := backRepo.BackRepoWait.Map_WaitDBID_WaitPtr[waitDB.ID]
	if waitOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), waitOld, waitNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the waitDB
	c.JSON(http.StatusOK, waitDB)
}

// DeleteWait
//
// swagger:route DELETE /waits/{ID} waits deleteWait
//
// # Delete a wait
//
// default: genericError
//
//	200: waitDBResponse
func (controller *Controller) DeleteWait(c *gin.Context) {

	mutexWait.Lock()
	defer mutexWait.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWait", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWait.GetDB()

	// Get model if exist
	var waitDB orm.WaitDB
	if err := db.First(&waitDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&waitDB)

	// get an instance (not staged) from DB instance, and call callback function
	waitDeleted := new(models.Wait)
	waitDB.CopyBasicFieldsToWait(waitDeleted)

	// get stage instance from DB instance, and call callback function
	waitStaged := backRepo.BackRepoWait.Map_WaitDBID_WaitPtr[waitDB.ID]
	if waitStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), waitStaged, waitDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
