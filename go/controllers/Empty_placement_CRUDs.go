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
var __Empty_placement__dummysDeclaration__ models.Empty_placement
var __Empty_placement_time__dummyDeclaration time.Duration

var mutexEmpty_placement sync.Mutex

// An Empty_placementID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_placement updateEmpty_placement deleteEmpty_placement
type Empty_placementID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_placementInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_placement updateEmpty_placement
type Empty_placementInput struct {
	// The Empty_placement to submit or modify
	// in: body
	Empty_placement *orm.Empty_placementAPI
}

// GetEmpty_placements
//
// swagger:route GET /empty_placements empty_placements getEmpty_placements
//
// # Get all empty_placements
//
// Responses:
// default: genericError
//
//	200: empty_placementDBResponse
func (controller *Controller) GetEmpty_placements(c *gin.Context) {

	// source slice
	var empty_placementDBs []orm.Empty_placementDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_placements", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement.GetDB()

	query := db.Find(&empty_placementDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_placementAPIs := make([]orm.Empty_placementAPI, 0)

	// for each empty_placement, update fields from the database nullable fields
	for idx := range empty_placementDBs {
		empty_placementDB := &empty_placementDBs[idx]
		_ = empty_placementDB
		var empty_placementAPI orm.Empty_placementAPI

		// insertion point for updating fields
		empty_placementAPI.ID = empty_placementDB.ID
		empty_placementDB.CopyBasicFieldsToEmpty_placement_WOP(&empty_placementAPI.Empty_placement_WOP)
		empty_placementAPI.Empty_placementPointersEncoding = empty_placementDB.Empty_placementPointersEncoding
		empty_placementAPIs = append(empty_placementAPIs, empty_placementAPI)
	}

	c.JSON(http.StatusOK, empty_placementAPIs)
}

// PostEmpty_placement
//
// swagger:route POST /empty_placements empty_placements postEmpty_placement
//
// Creates a empty_placement
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_placement(c *gin.Context) {

	mutexEmpty_placement.Lock()
	defer mutexEmpty_placement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_placements", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement.GetDB()

	// Validate input
	var input orm.Empty_placementAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_placement
	empty_placementDB := orm.Empty_placementDB{}
	empty_placementDB.Empty_placementPointersEncoding = input.Empty_placementPointersEncoding
	empty_placementDB.CopyBasicFieldsFromEmpty_placement_WOP(&input.Empty_placement_WOP)

	query := db.Create(&empty_placementDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_placement.CheckoutPhaseOneInstance(&empty_placementDB)
	empty_placement := backRepo.BackRepoEmpty_placement.Map_Empty_placementDBID_Empty_placementPtr[empty_placementDB.ID]

	if empty_placement != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_placement)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_placementDB)
}

// GetEmpty_placement
//
// swagger:route GET /empty_placements/{ID} empty_placements getEmpty_placement
//
// Gets the details for a empty_placement.
//
// Responses:
// default: genericError
//
//	200: empty_placementDBResponse
func (controller *Controller) GetEmpty_placement(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_placement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement.GetDB()

	// Get empty_placementDB in DB
	var empty_placementDB orm.Empty_placementDB
	if err := db.First(&empty_placementDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_placementAPI orm.Empty_placementAPI
	empty_placementAPI.ID = empty_placementDB.ID
	empty_placementAPI.Empty_placementPointersEncoding = empty_placementDB.Empty_placementPointersEncoding
	empty_placementDB.CopyBasicFieldsToEmpty_placement_WOP(&empty_placementAPI.Empty_placement_WOP)

	c.JSON(http.StatusOK, empty_placementAPI)
}

// UpdateEmpty_placement
//
// swagger:route PATCH /empty_placements/{ID} empty_placements updateEmpty_placement
//
// # Update a empty_placement
//
// Responses:
// default: genericError
//
//	200: empty_placementDBResponse
func (controller *Controller) UpdateEmpty_placement(c *gin.Context) {

	mutexEmpty_placement.Lock()
	defer mutexEmpty_placement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_placement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement.GetDB()

	// Validate input
	var input orm.Empty_placementAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_placementDB orm.Empty_placementDB

	// fetch the empty_placement
	query := db.First(&empty_placementDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_placementDB.CopyBasicFieldsFromEmpty_placement_WOP(&input.Empty_placement_WOP)
	empty_placementDB.Empty_placementPointersEncoding = input.Empty_placementPointersEncoding

	query = db.Model(&empty_placementDB).Updates(empty_placementDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_placementNew := new(models.Empty_placement)
	empty_placementDB.CopyBasicFieldsToEmpty_placement(empty_placementNew)

	// redeem pointers
	empty_placementDB.DecodePointers(backRepo, empty_placementNew)

	// get stage instance from DB instance, and call callback function
	empty_placementOld := backRepo.BackRepoEmpty_placement.Map_Empty_placementDBID_Empty_placementPtr[empty_placementDB.ID]
	if empty_placementOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_placementOld, empty_placementNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_placementDB
	c.JSON(http.StatusOK, empty_placementDB)
}

// DeleteEmpty_placement
//
// swagger:route DELETE /empty_placements/{ID} empty_placements deleteEmpty_placement
//
// # Delete a empty_placement
//
// default: genericError
//
//	200: empty_placementDBResponse
func (controller *Controller) DeleteEmpty_placement(c *gin.Context) {

	mutexEmpty_placement.Lock()
	defer mutexEmpty_placement.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_placement", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement.GetDB()

	// Get model if exist
	var empty_placementDB orm.Empty_placementDB
	if err := db.First(&empty_placementDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_placementDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_placementDeleted := new(models.Empty_placement)
	empty_placementDB.CopyBasicFieldsToEmpty_placement(empty_placementDeleted)

	// get stage instance from DB instance, and call callback function
	empty_placementStaged := backRepo.BackRepoEmpty_placement.Map_Empty_placementDBID_Empty_placementPtr[empty_placementDB.ID]
	if empty_placementStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_placementStaged, empty_placementDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
