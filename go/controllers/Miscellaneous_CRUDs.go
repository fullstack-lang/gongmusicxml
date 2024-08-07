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
var __Miscellaneous__dummysDeclaration__ models.Miscellaneous
var __Miscellaneous_time__dummyDeclaration time.Duration

var mutexMiscellaneous sync.Mutex

// An MiscellaneousID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMiscellaneous updateMiscellaneous deleteMiscellaneous
type MiscellaneousID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MiscellaneousInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMiscellaneous updateMiscellaneous
type MiscellaneousInput struct {
	// The Miscellaneous to submit or modify
	// in: body
	Miscellaneous *orm.MiscellaneousAPI
}

// GetMiscellaneouss
//
// swagger:route GET /miscellaneouss miscellaneouss getMiscellaneouss
//
// # Get all miscellaneouss
//
// Responses:
// default: genericError
//
//	200: miscellaneousDBResponse
func (controller *Controller) GetMiscellaneouss(c *gin.Context) {

	// source slice
	var miscellaneousDBs []orm.MiscellaneousDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMiscellaneouss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous.GetDB()

	query := db.Find(&miscellaneousDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	miscellaneousAPIs := make([]orm.MiscellaneousAPI, 0)

	// for each miscellaneous, update fields from the database nullable fields
	for idx := range miscellaneousDBs {
		miscellaneousDB := &miscellaneousDBs[idx]
		_ = miscellaneousDB
		var miscellaneousAPI orm.MiscellaneousAPI

		// insertion point for updating fields
		miscellaneousAPI.ID = miscellaneousDB.ID
		miscellaneousDB.CopyBasicFieldsToMiscellaneous_WOP(&miscellaneousAPI.Miscellaneous_WOP)
		miscellaneousAPI.MiscellaneousPointersEncoding = miscellaneousDB.MiscellaneousPointersEncoding
		miscellaneousAPIs = append(miscellaneousAPIs, miscellaneousAPI)
	}

	c.JSON(http.StatusOK, miscellaneousAPIs)
}

// PostMiscellaneous
//
// swagger:route POST /miscellaneouss miscellaneouss postMiscellaneous
//
// Creates a miscellaneous
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMiscellaneous(c *gin.Context) {

	mutexMiscellaneous.Lock()
	defer mutexMiscellaneous.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMiscellaneouss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous.GetDB()

	// Validate input
	var input orm.MiscellaneousAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create miscellaneous
	miscellaneousDB := orm.MiscellaneousDB{}
	miscellaneousDB.MiscellaneousPointersEncoding = input.MiscellaneousPointersEncoding
	miscellaneousDB.CopyBasicFieldsFromMiscellaneous_WOP(&input.Miscellaneous_WOP)

	query := db.Create(&miscellaneousDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMiscellaneous.CheckoutPhaseOneInstance(&miscellaneousDB)
	miscellaneous := backRepo.BackRepoMiscellaneous.Map_MiscellaneousDBID_MiscellaneousPtr[miscellaneousDB.ID]

	if miscellaneous != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), miscellaneous)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, miscellaneousDB)
}

// GetMiscellaneous
//
// swagger:route GET /miscellaneouss/{ID} miscellaneouss getMiscellaneous
//
// Gets the details for a miscellaneous.
//
// Responses:
// default: genericError
//
//	200: miscellaneousDBResponse
func (controller *Controller) GetMiscellaneous(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMiscellaneous", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous.GetDB()

	// Get miscellaneousDB in DB
	var miscellaneousDB orm.MiscellaneousDB
	if err := db.First(&miscellaneousDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var miscellaneousAPI orm.MiscellaneousAPI
	miscellaneousAPI.ID = miscellaneousDB.ID
	miscellaneousAPI.MiscellaneousPointersEncoding = miscellaneousDB.MiscellaneousPointersEncoding
	miscellaneousDB.CopyBasicFieldsToMiscellaneous_WOP(&miscellaneousAPI.Miscellaneous_WOP)

	c.JSON(http.StatusOK, miscellaneousAPI)
}

// UpdateMiscellaneous
//
// swagger:route PATCH /miscellaneouss/{ID} miscellaneouss updateMiscellaneous
//
// # Update a miscellaneous
//
// Responses:
// default: genericError
//
//	200: miscellaneousDBResponse
func (controller *Controller) UpdateMiscellaneous(c *gin.Context) {

	mutexMiscellaneous.Lock()
	defer mutexMiscellaneous.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMiscellaneous", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous.GetDB()

	// Validate input
	var input orm.MiscellaneousAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var miscellaneousDB orm.MiscellaneousDB

	// fetch the miscellaneous
	query := db.First(&miscellaneousDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	miscellaneousDB.CopyBasicFieldsFromMiscellaneous_WOP(&input.Miscellaneous_WOP)
	miscellaneousDB.MiscellaneousPointersEncoding = input.MiscellaneousPointersEncoding

	query = db.Model(&miscellaneousDB).Updates(miscellaneousDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	miscellaneousNew := new(models.Miscellaneous)
	miscellaneousDB.CopyBasicFieldsToMiscellaneous(miscellaneousNew)

	// redeem pointers
	miscellaneousDB.DecodePointers(backRepo, miscellaneousNew)

	// get stage instance from DB instance, and call callback function
	miscellaneousOld := backRepo.BackRepoMiscellaneous.Map_MiscellaneousDBID_MiscellaneousPtr[miscellaneousDB.ID]
	if miscellaneousOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), miscellaneousOld, miscellaneousNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the miscellaneousDB
	c.JSON(http.StatusOK, miscellaneousDB)
}

// DeleteMiscellaneous
//
// swagger:route DELETE /miscellaneouss/{ID} miscellaneouss deleteMiscellaneous
//
// # Delete a miscellaneous
//
// default: genericError
//
//	200: miscellaneousDBResponse
func (controller *Controller) DeleteMiscellaneous(c *gin.Context) {

	mutexMiscellaneous.Lock()
	defer mutexMiscellaneous.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMiscellaneous", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous.GetDB()

	// Get model if exist
	var miscellaneousDB orm.MiscellaneousDB
	if err := db.First(&miscellaneousDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&miscellaneousDB)

	// get an instance (not staged) from DB instance, and call callback function
	miscellaneousDeleted := new(models.Miscellaneous)
	miscellaneousDB.CopyBasicFieldsToMiscellaneous(miscellaneousDeleted)

	// get stage instance from DB instance, and call callback function
	miscellaneousStaged := backRepo.BackRepoMiscellaneous.Map_MiscellaneousDBID_MiscellaneousPtr[miscellaneousDB.ID]
	if miscellaneousStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), miscellaneousStaged, miscellaneousDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
