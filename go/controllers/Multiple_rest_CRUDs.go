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
var __Multiple_rest__dummysDeclaration__ models.Multiple_rest
var __Multiple_rest_time__dummyDeclaration time.Duration

var mutexMultiple_rest sync.Mutex

// An Multiple_restID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMultiple_rest updateMultiple_rest deleteMultiple_rest
type Multiple_restID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Multiple_restInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMultiple_rest updateMultiple_rest
type Multiple_restInput struct {
	// The Multiple_rest to submit or modify
	// in: body
	Multiple_rest *orm.Multiple_restAPI
}

// GetMultiple_rests
//
// swagger:route GET /multiple_rests multiple_rests getMultiple_rests
//
// # Get all multiple_rests
//
// Responses:
// default: genericError
//
//	200: multiple_restDBResponse
func (controller *Controller) GetMultiple_rests(c *gin.Context) {

	// source slice
	var multiple_restDBs []orm.Multiple_restDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMultiple_rests", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMultiple_rest.GetDB()

	query := db.Find(&multiple_restDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	multiple_restAPIs := make([]orm.Multiple_restAPI, 0)

	// for each multiple_rest, update fields from the database nullable fields
	for idx := range multiple_restDBs {
		multiple_restDB := &multiple_restDBs[idx]
		_ = multiple_restDB
		var multiple_restAPI orm.Multiple_restAPI

		// insertion point for updating fields
		multiple_restAPI.ID = multiple_restDB.ID
		multiple_restDB.CopyBasicFieldsToMultiple_rest_WOP(&multiple_restAPI.Multiple_rest_WOP)
		multiple_restAPI.Multiple_restPointersEncoding = multiple_restDB.Multiple_restPointersEncoding
		multiple_restAPIs = append(multiple_restAPIs, multiple_restAPI)
	}

	c.JSON(http.StatusOK, multiple_restAPIs)
}

// PostMultiple_rest
//
// swagger:route POST /multiple_rests multiple_rests postMultiple_rest
//
// Creates a multiple_rest
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMultiple_rest(c *gin.Context) {

	mutexMultiple_rest.Lock()
	defer mutexMultiple_rest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMultiple_rests", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMultiple_rest.GetDB()

	// Validate input
	var input orm.Multiple_restAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create multiple_rest
	multiple_restDB := orm.Multiple_restDB{}
	multiple_restDB.Multiple_restPointersEncoding = input.Multiple_restPointersEncoding
	multiple_restDB.CopyBasicFieldsFromMultiple_rest_WOP(&input.Multiple_rest_WOP)

	query := db.Create(&multiple_restDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMultiple_rest.CheckoutPhaseOneInstance(&multiple_restDB)
	multiple_rest := backRepo.BackRepoMultiple_rest.Map_Multiple_restDBID_Multiple_restPtr[multiple_restDB.ID]

	if multiple_rest != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), multiple_rest)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, multiple_restDB)
}

// GetMultiple_rest
//
// swagger:route GET /multiple_rests/{ID} multiple_rests getMultiple_rest
//
// Gets the details for a multiple_rest.
//
// Responses:
// default: genericError
//
//	200: multiple_restDBResponse
func (controller *Controller) GetMultiple_rest(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMultiple_rest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMultiple_rest.GetDB()

	// Get multiple_restDB in DB
	var multiple_restDB orm.Multiple_restDB
	if err := db.First(&multiple_restDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var multiple_restAPI orm.Multiple_restAPI
	multiple_restAPI.ID = multiple_restDB.ID
	multiple_restAPI.Multiple_restPointersEncoding = multiple_restDB.Multiple_restPointersEncoding
	multiple_restDB.CopyBasicFieldsToMultiple_rest_WOP(&multiple_restAPI.Multiple_rest_WOP)

	c.JSON(http.StatusOK, multiple_restAPI)
}

// UpdateMultiple_rest
//
// swagger:route PATCH /multiple_rests/{ID} multiple_rests updateMultiple_rest
//
// # Update a multiple_rest
//
// Responses:
// default: genericError
//
//	200: multiple_restDBResponse
func (controller *Controller) UpdateMultiple_rest(c *gin.Context) {

	mutexMultiple_rest.Lock()
	defer mutexMultiple_rest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMultiple_rest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMultiple_rest.GetDB()

	// Validate input
	var input orm.Multiple_restAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var multiple_restDB orm.Multiple_restDB

	// fetch the multiple_rest
	query := db.First(&multiple_restDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	multiple_restDB.CopyBasicFieldsFromMultiple_rest_WOP(&input.Multiple_rest_WOP)
	multiple_restDB.Multiple_restPointersEncoding = input.Multiple_restPointersEncoding

	query = db.Model(&multiple_restDB).Updates(multiple_restDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	multiple_restNew := new(models.Multiple_rest)
	multiple_restDB.CopyBasicFieldsToMultiple_rest(multiple_restNew)

	// redeem pointers
	multiple_restDB.DecodePointers(backRepo, multiple_restNew)

	// get stage instance from DB instance, and call callback function
	multiple_restOld := backRepo.BackRepoMultiple_rest.Map_Multiple_restDBID_Multiple_restPtr[multiple_restDB.ID]
	if multiple_restOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), multiple_restOld, multiple_restNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the multiple_restDB
	c.JSON(http.StatusOK, multiple_restDB)
}

// DeleteMultiple_rest
//
// swagger:route DELETE /multiple_rests/{ID} multiple_rests deleteMultiple_rest
//
// # Delete a multiple_rest
//
// default: genericError
//
//	200: multiple_restDBResponse
func (controller *Controller) DeleteMultiple_rest(c *gin.Context) {

	mutexMultiple_rest.Lock()
	defer mutexMultiple_rest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMultiple_rest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMultiple_rest.GetDB()

	// Get model if exist
	var multiple_restDB orm.Multiple_restDB
	if err := db.First(&multiple_restDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&multiple_restDB)

	// get an instance (not staged) from DB instance, and call callback function
	multiple_restDeleted := new(models.Multiple_rest)
	multiple_restDB.CopyBasicFieldsToMultiple_rest(multiple_restDeleted)

	// get stage instance from DB instance, and call callback function
	multiple_restStaged := backRepo.BackRepoMultiple_rest.Map_Multiple_restDBID_Multiple_restPtr[multiple_restDB.ID]
	if multiple_restStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), multiple_restStaged, multiple_restDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
