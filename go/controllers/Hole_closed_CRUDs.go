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
var __Hole_closed__dummysDeclaration__ models.Hole_closed
var __Hole_closed_time__dummyDeclaration time.Duration

var mutexHole_closed sync.Mutex

// An Hole_closedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHole_closed updateHole_closed deleteHole_closed
type Hole_closedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Hole_closedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHole_closed updateHole_closed
type Hole_closedInput struct {
	// The Hole_closed to submit or modify
	// in: body
	Hole_closed *orm.Hole_closedAPI
}

// GetHole_closeds
//
// swagger:route GET /hole_closeds hole_closeds getHole_closeds
//
// # Get all hole_closeds
//
// Responses:
// default: genericError
//
//	200: hole_closedDBResponse
func (controller *Controller) GetHole_closeds(c *gin.Context) {

	// source slice
	var hole_closedDBs []orm.Hole_closedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHole_closeds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole_closed.GetDB()

	query := db.Find(&hole_closedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	hole_closedAPIs := make([]orm.Hole_closedAPI, 0)

	// for each hole_closed, update fields from the database nullable fields
	for idx := range hole_closedDBs {
		hole_closedDB := &hole_closedDBs[idx]
		_ = hole_closedDB
		var hole_closedAPI orm.Hole_closedAPI

		// insertion point for updating fields
		hole_closedAPI.ID = hole_closedDB.ID
		hole_closedDB.CopyBasicFieldsToHole_closed_WOP(&hole_closedAPI.Hole_closed_WOP)
		hole_closedAPI.Hole_closedPointersEncoding = hole_closedDB.Hole_closedPointersEncoding
		hole_closedAPIs = append(hole_closedAPIs, hole_closedAPI)
	}

	c.JSON(http.StatusOK, hole_closedAPIs)
}

// PostHole_closed
//
// swagger:route POST /hole_closeds hole_closeds postHole_closed
//
// Creates a hole_closed
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHole_closed(c *gin.Context) {

	mutexHole_closed.Lock()
	defer mutexHole_closed.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHole_closeds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole_closed.GetDB()

	// Validate input
	var input orm.Hole_closedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create hole_closed
	hole_closedDB := orm.Hole_closedDB{}
	hole_closedDB.Hole_closedPointersEncoding = input.Hole_closedPointersEncoding
	hole_closedDB.CopyBasicFieldsFromHole_closed_WOP(&input.Hole_closed_WOP)

	query := db.Create(&hole_closedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHole_closed.CheckoutPhaseOneInstance(&hole_closedDB)
	hole_closed := backRepo.BackRepoHole_closed.Map_Hole_closedDBID_Hole_closedPtr[hole_closedDB.ID]

	if hole_closed != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), hole_closed)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, hole_closedDB)
}

// GetHole_closed
//
// swagger:route GET /hole_closeds/{ID} hole_closeds getHole_closed
//
// Gets the details for a hole_closed.
//
// Responses:
// default: genericError
//
//	200: hole_closedDBResponse
func (controller *Controller) GetHole_closed(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHole_closed", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole_closed.GetDB()

	// Get hole_closedDB in DB
	var hole_closedDB orm.Hole_closedDB
	if err := db.First(&hole_closedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var hole_closedAPI orm.Hole_closedAPI
	hole_closedAPI.ID = hole_closedDB.ID
	hole_closedAPI.Hole_closedPointersEncoding = hole_closedDB.Hole_closedPointersEncoding
	hole_closedDB.CopyBasicFieldsToHole_closed_WOP(&hole_closedAPI.Hole_closed_WOP)

	c.JSON(http.StatusOK, hole_closedAPI)
}

// UpdateHole_closed
//
// swagger:route PATCH /hole_closeds/{ID} hole_closeds updateHole_closed
//
// # Update a hole_closed
//
// Responses:
// default: genericError
//
//	200: hole_closedDBResponse
func (controller *Controller) UpdateHole_closed(c *gin.Context) {

	mutexHole_closed.Lock()
	defer mutexHole_closed.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHole_closed", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole_closed.GetDB()

	// Validate input
	var input orm.Hole_closedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var hole_closedDB orm.Hole_closedDB

	// fetch the hole_closed
	query := db.First(&hole_closedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	hole_closedDB.CopyBasicFieldsFromHole_closed_WOP(&input.Hole_closed_WOP)
	hole_closedDB.Hole_closedPointersEncoding = input.Hole_closedPointersEncoding

	query = db.Model(&hole_closedDB).Updates(hole_closedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	hole_closedNew := new(models.Hole_closed)
	hole_closedDB.CopyBasicFieldsToHole_closed(hole_closedNew)

	// redeem pointers
	hole_closedDB.DecodePointers(backRepo, hole_closedNew)

	// get stage instance from DB instance, and call callback function
	hole_closedOld := backRepo.BackRepoHole_closed.Map_Hole_closedDBID_Hole_closedPtr[hole_closedDB.ID]
	if hole_closedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), hole_closedOld, hole_closedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the hole_closedDB
	c.JSON(http.StatusOK, hole_closedDB)
}

// DeleteHole_closed
//
// swagger:route DELETE /hole_closeds/{ID} hole_closeds deleteHole_closed
//
// # Delete a hole_closed
//
// default: genericError
//
//	200: hole_closedDBResponse
func (controller *Controller) DeleteHole_closed(c *gin.Context) {

	mutexHole_closed.Lock()
	defer mutexHole_closed.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHole_closed", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole_closed.GetDB()

	// Get model if exist
	var hole_closedDB orm.Hole_closedDB
	if err := db.First(&hole_closedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&hole_closedDB)

	// get an instance (not staged) from DB instance, and call callback function
	hole_closedDeleted := new(models.Hole_closed)
	hole_closedDB.CopyBasicFieldsToHole_closed(hole_closedDeleted)

	// get stage instance from DB instance, and call callback function
	hole_closedStaged := backRepo.BackRepoHole_closed.Map_Hole_closedDBID_Hole_closedPtr[hole_closedDB.ID]
	if hole_closedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), hole_closedStaged, hole_closedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
