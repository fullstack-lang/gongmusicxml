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
var __Wavy_line__dummysDeclaration__ models.Wavy_line
var __Wavy_line_time__dummyDeclaration time.Duration

var mutexWavy_line sync.Mutex

// An Wavy_lineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWavy_line updateWavy_line deleteWavy_line
type Wavy_lineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Wavy_lineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWavy_line updateWavy_line
type Wavy_lineInput struct {
	// The Wavy_line to submit or modify
	// in: body
	Wavy_line *orm.Wavy_lineAPI
}

// GetWavy_lines
//
// swagger:route GET /wavy_lines wavy_lines getWavy_lines
//
// # Get all wavy_lines
//
// Responses:
// default: genericError
//
//	200: wavy_lineDBResponse
func (controller *Controller) GetWavy_lines(c *gin.Context) {

	// source slice
	var wavy_lineDBs []orm.Wavy_lineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWavy_lines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWavy_line.GetDB()

	query := db.Find(&wavy_lineDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	wavy_lineAPIs := make([]orm.Wavy_lineAPI, 0)

	// for each wavy_line, update fields from the database nullable fields
	for idx := range wavy_lineDBs {
		wavy_lineDB := &wavy_lineDBs[idx]
		_ = wavy_lineDB
		var wavy_lineAPI orm.Wavy_lineAPI

		// insertion point for updating fields
		wavy_lineAPI.ID = wavy_lineDB.ID
		wavy_lineDB.CopyBasicFieldsToWavy_line_WOP(&wavy_lineAPI.Wavy_line_WOP)
		wavy_lineAPI.Wavy_linePointersEncoding = wavy_lineDB.Wavy_linePointersEncoding
		wavy_lineAPIs = append(wavy_lineAPIs, wavy_lineAPI)
	}

	c.JSON(http.StatusOK, wavy_lineAPIs)
}

// PostWavy_line
//
// swagger:route POST /wavy_lines wavy_lines postWavy_line
//
// Creates a wavy_line
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWavy_line(c *gin.Context) {

	mutexWavy_line.Lock()
	defer mutexWavy_line.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWavy_lines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWavy_line.GetDB()

	// Validate input
	var input orm.Wavy_lineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create wavy_line
	wavy_lineDB := orm.Wavy_lineDB{}
	wavy_lineDB.Wavy_linePointersEncoding = input.Wavy_linePointersEncoding
	wavy_lineDB.CopyBasicFieldsFromWavy_line_WOP(&input.Wavy_line_WOP)

	query := db.Create(&wavy_lineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWavy_line.CheckoutPhaseOneInstance(&wavy_lineDB)
	wavy_line := backRepo.BackRepoWavy_line.Map_Wavy_lineDBID_Wavy_linePtr[wavy_lineDB.ID]

	if wavy_line != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), wavy_line)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, wavy_lineDB)
}

// GetWavy_line
//
// swagger:route GET /wavy_lines/{ID} wavy_lines getWavy_line
//
// Gets the details for a wavy_line.
//
// Responses:
// default: genericError
//
//	200: wavy_lineDBResponse
func (controller *Controller) GetWavy_line(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWavy_line", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWavy_line.GetDB()

	// Get wavy_lineDB in DB
	var wavy_lineDB orm.Wavy_lineDB
	if err := db.First(&wavy_lineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var wavy_lineAPI orm.Wavy_lineAPI
	wavy_lineAPI.ID = wavy_lineDB.ID
	wavy_lineAPI.Wavy_linePointersEncoding = wavy_lineDB.Wavy_linePointersEncoding
	wavy_lineDB.CopyBasicFieldsToWavy_line_WOP(&wavy_lineAPI.Wavy_line_WOP)

	c.JSON(http.StatusOK, wavy_lineAPI)
}

// UpdateWavy_line
//
// swagger:route PATCH /wavy_lines/{ID} wavy_lines updateWavy_line
//
// # Update a wavy_line
//
// Responses:
// default: genericError
//
//	200: wavy_lineDBResponse
func (controller *Controller) UpdateWavy_line(c *gin.Context) {

	mutexWavy_line.Lock()
	defer mutexWavy_line.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWavy_line", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWavy_line.GetDB()

	// Validate input
	var input orm.Wavy_lineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var wavy_lineDB orm.Wavy_lineDB

	// fetch the wavy_line
	query := db.First(&wavy_lineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	wavy_lineDB.CopyBasicFieldsFromWavy_line_WOP(&input.Wavy_line_WOP)
	wavy_lineDB.Wavy_linePointersEncoding = input.Wavy_linePointersEncoding

	query = db.Model(&wavy_lineDB).Updates(wavy_lineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	wavy_lineNew := new(models.Wavy_line)
	wavy_lineDB.CopyBasicFieldsToWavy_line(wavy_lineNew)

	// redeem pointers
	wavy_lineDB.DecodePointers(backRepo, wavy_lineNew)

	// get stage instance from DB instance, and call callback function
	wavy_lineOld := backRepo.BackRepoWavy_line.Map_Wavy_lineDBID_Wavy_linePtr[wavy_lineDB.ID]
	if wavy_lineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), wavy_lineOld, wavy_lineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the wavy_lineDB
	c.JSON(http.StatusOK, wavy_lineDB)
}

// DeleteWavy_line
//
// swagger:route DELETE /wavy_lines/{ID} wavy_lines deleteWavy_line
//
// # Delete a wavy_line
//
// default: genericError
//
//	200: wavy_lineDBResponse
func (controller *Controller) DeleteWavy_line(c *gin.Context) {

	mutexWavy_line.Lock()
	defer mutexWavy_line.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWavy_line", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWavy_line.GetDB()

	// Get model if exist
	var wavy_lineDB orm.Wavy_lineDB
	if err := db.First(&wavy_lineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&wavy_lineDB)

	// get an instance (not staged) from DB instance, and call callback function
	wavy_lineDeleted := new(models.Wavy_line)
	wavy_lineDB.CopyBasicFieldsToWavy_line(wavy_lineDeleted)

	// get stage instance from DB instance, and call callback function
	wavy_lineStaged := backRepo.BackRepoWavy_line.Map_Wavy_lineDBID_Wavy_linePtr[wavy_lineDB.ID]
	if wavy_lineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), wavy_lineStaged, wavy_lineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
