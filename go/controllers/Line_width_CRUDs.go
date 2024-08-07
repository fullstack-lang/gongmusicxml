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
var __Line_width__dummysDeclaration__ models.Line_width
var __Line_width_time__dummyDeclaration time.Duration

var mutexLine_width sync.Mutex

// An Line_widthID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLine_width updateLine_width deleteLine_width
type Line_widthID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Line_widthInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLine_width updateLine_width
type Line_widthInput struct {
	// The Line_width to submit or modify
	// in: body
	Line_width *orm.Line_widthAPI
}

// GetLine_widths
//
// swagger:route GET /line_widths line_widths getLine_widths
//
// # Get all line_widths
//
// Responses:
// default: genericError
//
//	200: line_widthDBResponse
func (controller *Controller) GetLine_widths(c *gin.Context) {

	// source slice
	var line_widthDBs []orm.Line_widthDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLine_widths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_width.GetDB()

	query := db.Find(&line_widthDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	line_widthAPIs := make([]orm.Line_widthAPI, 0)

	// for each line_width, update fields from the database nullable fields
	for idx := range line_widthDBs {
		line_widthDB := &line_widthDBs[idx]
		_ = line_widthDB
		var line_widthAPI orm.Line_widthAPI

		// insertion point for updating fields
		line_widthAPI.ID = line_widthDB.ID
		line_widthDB.CopyBasicFieldsToLine_width_WOP(&line_widthAPI.Line_width_WOP)
		line_widthAPI.Line_widthPointersEncoding = line_widthDB.Line_widthPointersEncoding
		line_widthAPIs = append(line_widthAPIs, line_widthAPI)
	}

	c.JSON(http.StatusOK, line_widthAPIs)
}

// PostLine_width
//
// swagger:route POST /line_widths line_widths postLine_width
//
// Creates a line_width
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLine_width(c *gin.Context) {

	mutexLine_width.Lock()
	defer mutexLine_width.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLine_widths", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_width.GetDB()

	// Validate input
	var input orm.Line_widthAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create line_width
	line_widthDB := orm.Line_widthDB{}
	line_widthDB.Line_widthPointersEncoding = input.Line_widthPointersEncoding
	line_widthDB.CopyBasicFieldsFromLine_width_WOP(&input.Line_width_WOP)

	query := db.Create(&line_widthDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLine_width.CheckoutPhaseOneInstance(&line_widthDB)
	line_width := backRepo.BackRepoLine_width.Map_Line_widthDBID_Line_widthPtr[line_widthDB.ID]

	if line_width != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), line_width)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, line_widthDB)
}

// GetLine_width
//
// swagger:route GET /line_widths/{ID} line_widths getLine_width
//
// Gets the details for a line_width.
//
// Responses:
// default: genericError
//
//	200: line_widthDBResponse
func (controller *Controller) GetLine_width(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLine_width", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_width.GetDB()

	// Get line_widthDB in DB
	var line_widthDB orm.Line_widthDB
	if err := db.First(&line_widthDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var line_widthAPI orm.Line_widthAPI
	line_widthAPI.ID = line_widthDB.ID
	line_widthAPI.Line_widthPointersEncoding = line_widthDB.Line_widthPointersEncoding
	line_widthDB.CopyBasicFieldsToLine_width_WOP(&line_widthAPI.Line_width_WOP)

	c.JSON(http.StatusOK, line_widthAPI)
}

// UpdateLine_width
//
// swagger:route PATCH /line_widths/{ID} line_widths updateLine_width
//
// # Update a line_width
//
// Responses:
// default: genericError
//
//	200: line_widthDBResponse
func (controller *Controller) UpdateLine_width(c *gin.Context) {

	mutexLine_width.Lock()
	defer mutexLine_width.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLine_width", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_width.GetDB()

	// Validate input
	var input orm.Line_widthAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var line_widthDB orm.Line_widthDB

	// fetch the line_width
	query := db.First(&line_widthDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	line_widthDB.CopyBasicFieldsFromLine_width_WOP(&input.Line_width_WOP)
	line_widthDB.Line_widthPointersEncoding = input.Line_widthPointersEncoding

	query = db.Model(&line_widthDB).Updates(line_widthDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	line_widthNew := new(models.Line_width)
	line_widthDB.CopyBasicFieldsToLine_width(line_widthNew)

	// redeem pointers
	line_widthDB.DecodePointers(backRepo, line_widthNew)

	// get stage instance from DB instance, and call callback function
	line_widthOld := backRepo.BackRepoLine_width.Map_Line_widthDBID_Line_widthPtr[line_widthDB.ID]
	if line_widthOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), line_widthOld, line_widthNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the line_widthDB
	c.JSON(http.StatusOK, line_widthDB)
}

// DeleteLine_width
//
// swagger:route DELETE /line_widths/{ID} line_widths deleteLine_width
//
// # Delete a line_width
//
// default: genericError
//
//	200: line_widthDBResponse
func (controller *Controller) DeleteLine_width(c *gin.Context) {

	mutexLine_width.Lock()
	defer mutexLine_width.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLine_width", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_width.GetDB()

	// Get model if exist
	var line_widthDB orm.Line_widthDB
	if err := db.First(&line_widthDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&line_widthDB)

	// get an instance (not staged) from DB instance, and call callback function
	line_widthDeleted := new(models.Line_width)
	line_widthDB.CopyBasicFieldsToLine_width(line_widthDeleted)

	// get stage instance from DB instance, and call callback function
	line_widthStaged := backRepo.BackRepoLine_width.Map_Line_widthDBID_Line_widthPtr[line_widthDB.ID]
	if line_widthStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), line_widthStaged, line_widthDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
