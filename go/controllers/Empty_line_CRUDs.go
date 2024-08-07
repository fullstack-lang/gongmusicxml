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
var __Empty_line__dummysDeclaration__ models.Empty_line
var __Empty_line_time__dummyDeclaration time.Duration

var mutexEmpty_line sync.Mutex

// An Empty_lineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_line updateEmpty_line deleteEmpty_line
type Empty_lineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_lineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_line updateEmpty_line
type Empty_lineInput struct {
	// The Empty_line to submit or modify
	// in: body
	Empty_line *orm.Empty_lineAPI
}

// GetEmpty_lines
//
// swagger:route GET /empty_lines empty_lines getEmpty_lines
//
// # Get all empty_lines
//
// Responses:
// default: genericError
//
//	200: empty_lineDBResponse
func (controller *Controller) GetEmpty_lines(c *gin.Context) {

	// source slice
	var empty_lineDBs []orm.Empty_lineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_lines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_line.GetDB()

	query := db.Find(&empty_lineDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_lineAPIs := make([]orm.Empty_lineAPI, 0)

	// for each empty_line, update fields from the database nullable fields
	for idx := range empty_lineDBs {
		empty_lineDB := &empty_lineDBs[idx]
		_ = empty_lineDB
		var empty_lineAPI orm.Empty_lineAPI

		// insertion point for updating fields
		empty_lineAPI.ID = empty_lineDB.ID
		empty_lineDB.CopyBasicFieldsToEmpty_line_WOP(&empty_lineAPI.Empty_line_WOP)
		empty_lineAPI.Empty_linePointersEncoding = empty_lineDB.Empty_linePointersEncoding
		empty_lineAPIs = append(empty_lineAPIs, empty_lineAPI)
	}

	c.JSON(http.StatusOK, empty_lineAPIs)
}

// PostEmpty_line
//
// swagger:route POST /empty_lines empty_lines postEmpty_line
//
// Creates a empty_line
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_line(c *gin.Context) {

	mutexEmpty_line.Lock()
	defer mutexEmpty_line.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_lines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_line.GetDB()

	// Validate input
	var input orm.Empty_lineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_line
	empty_lineDB := orm.Empty_lineDB{}
	empty_lineDB.Empty_linePointersEncoding = input.Empty_linePointersEncoding
	empty_lineDB.CopyBasicFieldsFromEmpty_line_WOP(&input.Empty_line_WOP)

	query := db.Create(&empty_lineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_line.CheckoutPhaseOneInstance(&empty_lineDB)
	empty_line := backRepo.BackRepoEmpty_line.Map_Empty_lineDBID_Empty_linePtr[empty_lineDB.ID]

	if empty_line != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_line)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_lineDB)
}

// GetEmpty_line
//
// swagger:route GET /empty_lines/{ID} empty_lines getEmpty_line
//
// Gets the details for a empty_line.
//
// Responses:
// default: genericError
//
//	200: empty_lineDBResponse
func (controller *Controller) GetEmpty_line(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_line", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_line.GetDB()

	// Get empty_lineDB in DB
	var empty_lineDB orm.Empty_lineDB
	if err := db.First(&empty_lineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_lineAPI orm.Empty_lineAPI
	empty_lineAPI.ID = empty_lineDB.ID
	empty_lineAPI.Empty_linePointersEncoding = empty_lineDB.Empty_linePointersEncoding
	empty_lineDB.CopyBasicFieldsToEmpty_line_WOP(&empty_lineAPI.Empty_line_WOP)

	c.JSON(http.StatusOK, empty_lineAPI)
}

// UpdateEmpty_line
//
// swagger:route PATCH /empty_lines/{ID} empty_lines updateEmpty_line
//
// # Update a empty_line
//
// Responses:
// default: genericError
//
//	200: empty_lineDBResponse
func (controller *Controller) UpdateEmpty_line(c *gin.Context) {

	mutexEmpty_line.Lock()
	defer mutexEmpty_line.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_line", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_line.GetDB()

	// Validate input
	var input orm.Empty_lineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_lineDB orm.Empty_lineDB

	// fetch the empty_line
	query := db.First(&empty_lineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_lineDB.CopyBasicFieldsFromEmpty_line_WOP(&input.Empty_line_WOP)
	empty_lineDB.Empty_linePointersEncoding = input.Empty_linePointersEncoding

	query = db.Model(&empty_lineDB).Updates(empty_lineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_lineNew := new(models.Empty_line)
	empty_lineDB.CopyBasicFieldsToEmpty_line(empty_lineNew)

	// redeem pointers
	empty_lineDB.DecodePointers(backRepo, empty_lineNew)

	// get stage instance from DB instance, and call callback function
	empty_lineOld := backRepo.BackRepoEmpty_line.Map_Empty_lineDBID_Empty_linePtr[empty_lineDB.ID]
	if empty_lineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_lineOld, empty_lineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_lineDB
	c.JSON(http.StatusOK, empty_lineDB)
}

// DeleteEmpty_line
//
// swagger:route DELETE /empty_lines/{ID} empty_lines deleteEmpty_line
//
// # Delete a empty_line
//
// default: genericError
//
//	200: empty_lineDBResponse
func (controller *Controller) DeleteEmpty_line(c *gin.Context) {

	mutexEmpty_line.Lock()
	defer mutexEmpty_line.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_line", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_line.GetDB()

	// Get model if exist
	var empty_lineDB orm.Empty_lineDB
	if err := db.First(&empty_lineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_lineDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_lineDeleted := new(models.Empty_line)
	empty_lineDB.CopyBasicFieldsToEmpty_line(empty_lineDeleted)

	// get stage instance from DB instance, and call callback function
	empty_lineStaged := backRepo.BackRepoEmpty_line.Map_Empty_lineDBID_Empty_linePtr[empty_lineDB.ID]
	if empty_lineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_lineStaged, empty_lineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
