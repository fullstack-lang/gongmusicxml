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
var __Harmon_closed__dummysDeclaration__ models.Harmon_closed
var __Harmon_closed_time__dummyDeclaration time.Duration

var mutexHarmon_closed sync.Mutex

// An Harmon_closedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHarmon_closed updateHarmon_closed deleteHarmon_closed
type Harmon_closedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Harmon_closedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHarmon_closed updateHarmon_closed
type Harmon_closedInput struct {
	// The Harmon_closed to submit or modify
	// in: body
	Harmon_closed *orm.Harmon_closedAPI
}

// GetHarmon_closeds
//
// swagger:route GET /harmon_closeds harmon_closeds getHarmon_closeds
//
// # Get all harmon_closeds
//
// Responses:
// default: genericError
//
//	200: harmon_closedDBResponse
func (controller *Controller) GetHarmon_closeds(c *gin.Context) {

	// source slice
	var harmon_closedDBs []orm.Harmon_closedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmon_closeds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_closed.GetDB()

	query := db.Find(&harmon_closedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	harmon_closedAPIs := make([]orm.Harmon_closedAPI, 0)

	// for each harmon_closed, update fields from the database nullable fields
	for idx := range harmon_closedDBs {
		harmon_closedDB := &harmon_closedDBs[idx]
		_ = harmon_closedDB
		var harmon_closedAPI orm.Harmon_closedAPI

		// insertion point for updating fields
		harmon_closedAPI.ID = harmon_closedDB.ID
		harmon_closedDB.CopyBasicFieldsToHarmon_closed_WOP(&harmon_closedAPI.Harmon_closed_WOP)
		harmon_closedAPI.Harmon_closedPointersEncoding = harmon_closedDB.Harmon_closedPointersEncoding
		harmon_closedAPIs = append(harmon_closedAPIs, harmon_closedAPI)
	}

	c.JSON(http.StatusOK, harmon_closedAPIs)
}

// PostHarmon_closed
//
// swagger:route POST /harmon_closeds harmon_closeds postHarmon_closed
//
// Creates a harmon_closed
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHarmon_closed(c *gin.Context) {

	mutexHarmon_closed.Lock()
	defer mutexHarmon_closed.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHarmon_closeds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_closed.GetDB()

	// Validate input
	var input orm.Harmon_closedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create harmon_closed
	harmon_closedDB := orm.Harmon_closedDB{}
	harmon_closedDB.Harmon_closedPointersEncoding = input.Harmon_closedPointersEncoding
	harmon_closedDB.CopyBasicFieldsFromHarmon_closed_WOP(&input.Harmon_closed_WOP)

	query := db.Create(&harmon_closedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHarmon_closed.CheckoutPhaseOneInstance(&harmon_closedDB)
	harmon_closed := backRepo.BackRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID]

	if harmon_closed != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), harmon_closed)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, harmon_closedDB)
}

// GetHarmon_closed
//
// swagger:route GET /harmon_closeds/{ID} harmon_closeds getHarmon_closed
//
// Gets the details for a harmon_closed.
//
// Responses:
// default: genericError
//
//	200: harmon_closedDBResponse
func (controller *Controller) GetHarmon_closed(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmon_closed", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_closed.GetDB()

	// Get harmon_closedDB in DB
	var harmon_closedDB orm.Harmon_closedDB
	if err := db.First(&harmon_closedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var harmon_closedAPI orm.Harmon_closedAPI
	harmon_closedAPI.ID = harmon_closedDB.ID
	harmon_closedAPI.Harmon_closedPointersEncoding = harmon_closedDB.Harmon_closedPointersEncoding
	harmon_closedDB.CopyBasicFieldsToHarmon_closed_WOP(&harmon_closedAPI.Harmon_closed_WOP)

	c.JSON(http.StatusOK, harmon_closedAPI)
}

// UpdateHarmon_closed
//
// swagger:route PATCH /harmon_closeds/{ID} harmon_closeds updateHarmon_closed
//
// # Update a harmon_closed
//
// Responses:
// default: genericError
//
//	200: harmon_closedDBResponse
func (controller *Controller) UpdateHarmon_closed(c *gin.Context) {

	mutexHarmon_closed.Lock()
	defer mutexHarmon_closed.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHarmon_closed", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_closed.GetDB()

	// Validate input
	var input orm.Harmon_closedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var harmon_closedDB orm.Harmon_closedDB

	// fetch the harmon_closed
	query := db.First(&harmon_closedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	harmon_closedDB.CopyBasicFieldsFromHarmon_closed_WOP(&input.Harmon_closed_WOP)
	harmon_closedDB.Harmon_closedPointersEncoding = input.Harmon_closedPointersEncoding

	query = db.Model(&harmon_closedDB).Updates(harmon_closedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	harmon_closedNew := new(models.Harmon_closed)
	harmon_closedDB.CopyBasicFieldsToHarmon_closed(harmon_closedNew)

	// redeem pointers
	harmon_closedDB.DecodePointers(backRepo, harmon_closedNew)

	// get stage instance from DB instance, and call callback function
	harmon_closedOld := backRepo.BackRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID]
	if harmon_closedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), harmon_closedOld, harmon_closedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the harmon_closedDB
	c.JSON(http.StatusOK, harmon_closedDB)
}

// DeleteHarmon_closed
//
// swagger:route DELETE /harmon_closeds/{ID} harmon_closeds deleteHarmon_closed
//
// # Delete a harmon_closed
//
// default: genericError
//
//	200: harmon_closedDBResponse
func (controller *Controller) DeleteHarmon_closed(c *gin.Context) {

	mutexHarmon_closed.Lock()
	defer mutexHarmon_closed.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHarmon_closed", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_closed.GetDB()

	// Get model if exist
	var harmon_closedDB orm.Harmon_closedDB
	if err := db.First(&harmon_closedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&harmon_closedDB)

	// get an instance (not staged) from DB instance, and call callback function
	harmon_closedDeleted := new(models.Harmon_closed)
	harmon_closedDB.CopyBasicFieldsToHarmon_closed(harmon_closedDeleted)

	// get stage instance from DB instance, and call callback function
	harmon_closedStaged := backRepo.BackRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID]
	if harmon_closedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), harmon_closedStaged, harmon_closedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
