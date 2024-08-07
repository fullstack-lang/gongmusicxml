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
var __Hole__dummysDeclaration__ models.Hole
var __Hole_time__dummyDeclaration time.Duration

var mutexHole sync.Mutex

// An HoleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHole updateHole deleteHole
type HoleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HoleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHole updateHole
type HoleInput struct {
	// The Hole to submit or modify
	// in: body
	Hole *orm.HoleAPI
}

// GetHoles
//
// swagger:route GET /holes holes getHoles
//
// # Get all holes
//
// Responses:
// default: genericError
//
//	200: holeDBResponse
func (controller *Controller) GetHoles(c *gin.Context) {

	// source slice
	var holeDBs []orm.HoleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHoles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole.GetDB()

	query := db.Find(&holeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	holeAPIs := make([]orm.HoleAPI, 0)

	// for each hole, update fields from the database nullable fields
	for idx := range holeDBs {
		holeDB := &holeDBs[idx]
		_ = holeDB
		var holeAPI orm.HoleAPI

		// insertion point for updating fields
		holeAPI.ID = holeDB.ID
		holeDB.CopyBasicFieldsToHole_WOP(&holeAPI.Hole_WOP)
		holeAPI.HolePointersEncoding = holeDB.HolePointersEncoding
		holeAPIs = append(holeAPIs, holeAPI)
	}

	c.JSON(http.StatusOK, holeAPIs)
}

// PostHole
//
// swagger:route POST /holes holes postHole
//
// Creates a hole
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHole(c *gin.Context) {

	mutexHole.Lock()
	defer mutexHole.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHoles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole.GetDB()

	// Validate input
	var input orm.HoleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create hole
	holeDB := orm.HoleDB{}
	holeDB.HolePointersEncoding = input.HolePointersEncoding
	holeDB.CopyBasicFieldsFromHole_WOP(&input.Hole_WOP)

	query := db.Create(&holeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHole.CheckoutPhaseOneInstance(&holeDB)
	hole := backRepo.BackRepoHole.Map_HoleDBID_HolePtr[holeDB.ID]

	if hole != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), hole)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, holeDB)
}

// GetHole
//
// swagger:route GET /holes/{ID} holes getHole
//
// Gets the details for a hole.
//
// Responses:
// default: genericError
//
//	200: holeDBResponse
func (controller *Controller) GetHole(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHole", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole.GetDB()

	// Get holeDB in DB
	var holeDB orm.HoleDB
	if err := db.First(&holeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var holeAPI orm.HoleAPI
	holeAPI.ID = holeDB.ID
	holeAPI.HolePointersEncoding = holeDB.HolePointersEncoding
	holeDB.CopyBasicFieldsToHole_WOP(&holeAPI.Hole_WOP)

	c.JSON(http.StatusOK, holeAPI)
}

// UpdateHole
//
// swagger:route PATCH /holes/{ID} holes updateHole
//
// # Update a hole
//
// Responses:
// default: genericError
//
//	200: holeDBResponse
func (controller *Controller) UpdateHole(c *gin.Context) {

	mutexHole.Lock()
	defer mutexHole.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHole", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole.GetDB()

	// Validate input
	var input orm.HoleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var holeDB orm.HoleDB

	// fetch the hole
	query := db.First(&holeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	holeDB.CopyBasicFieldsFromHole_WOP(&input.Hole_WOP)
	holeDB.HolePointersEncoding = input.HolePointersEncoding

	query = db.Model(&holeDB).Updates(holeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	holeNew := new(models.Hole)
	holeDB.CopyBasicFieldsToHole(holeNew)

	// redeem pointers
	holeDB.DecodePointers(backRepo, holeNew)

	// get stage instance from DB instance, and call callback function
	holeOld := backRepo.BackRepoHole.Map_HoleDBID_HolePtr[holeDB.ID]
	if holeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), holeOld, holeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the holeDB
	c.JSON(http.StatusOK, holeDB)
}

// DeleteHole
//
// swagger:route DELETE /holes/{ID} holes deleteHole
//
// # Delete a hole
//
// default: genericError
//
//	200: holeDBResponse
func (controller *Controller) DeleteHole(c *gin.Context) {

	mutexHole.Lock()
	defer mutexHole.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHole", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHole.GetDB()

	// Get model if exist
	var holeDB orm.HoleDB
	if err := db.First(&holeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&holeDB)

	// get an instance (not staged) from DB instance, and call callback function
	holeDeleted := new(models.Hole)
	holeDB.CopyBasicFieldsToHole(holeDeleted)

	// get stage instance from DB instance, and call callback function
	holeStaged := backRepo.BackRepoHole.Map_HoleDBID_HolePtr[holeDB.ID]
	if holeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), holeStaged, holeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
