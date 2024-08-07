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
var __Line_detail__dummysDeclaration__ models.Line_detail
var __Line_detail_time__dummyDeclaration time.Duration

var mutexLine_detail sync.Mutex

// An Line_detailID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLine_detail updateLine_detail deleteLine_detail
type Line_detailID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Line_detailInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLine_detail updateLine_detail
type Line_detailInput struct {
	// The Line_detail to submit or modify
	// in: body
	Line_detail *orm.Line_detailAPI
}

// GetLine_details
//
// swagger:route GET /line_details line_details getLine_details
//
// # Get all line_details
//
// Responses:
// default: genericError
//
//	200: line_detailDBResponse
func (controller *Controller) GetLine_details(c *gin.Context) {

	// source slice
	var line_detailDBs []orm.Line_detailDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLine_details", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_detail.GetDB()

	query := db.Find(&line_detailDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	line_detailAPIs := make([]orm.Line_detailAPI, 0)

	// for each line_detail, update fields from the database nullable fields
	for idx := range line_detailDBs {
		line_detailDB := &line_detailDBs[idx]
		_ = line_detailDB
		var line_detailAPI orm.Line_detailAPI

		// insertion point for updating fields
		line_detailAPI.ID = line_detailDB.ID
		line_detailDB.CopyBasicFieldsToLine_detail_WOP(&line_detailAPI.Line_detail_WOP)
		line_detailAPI.Line_detailPointersEncoding = line_detailDB.Line_detailPointersEncoding
		line_detailAPIs = append(line_detailAPIs, line_detailAPI)
	}

	c.JSON(http.StatusOK, line_detailAPIs)
}

// PostLine_detail
//
// swagger:route POST /line_details line_details postLine_detail
//
// Creates a line_detail
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLine_detail(c *gin.Context) {

	mutexLine_detail.Lock()
	defer mutexLine_detail.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLine_details", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_detail.GetDB()

	// Validate input
	var input orm.Line_detailAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create line_detail
	line_detailDB := orm.Line_detailDB{}
	line_detailDB.Line_detailPointersEncoding = input.Line_detailPointersEncoding
	line_detailDB.CopyBasicFieldsFromLine_detail_WOP(&input.Line_detail_WOP)

	query := db.Create(&line_detailDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLine_detail.CheckoutPhaseOneInstance(&line_detailDB)
	line_detail := backRepo.BackRepoLine_detail.Map_Line_detailDBID_Line_detailPtr[line_detailDB.ID]

	if line_detail != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), line_detail)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, line_detailDB)
}

// GetLine_detail
//
// swagger:route GET /line_details/{ID} line_details getLine_detail
//
// Gets the details for a line_detail.
//
// Responses:
// default: genericError
//
//	200: line_detailDBResponse
func (controller *Controller) GetLine_detail(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLine_detail", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_detail.GetDB()

	// Get line_detailDB in DB
	var line_detailDB orm.Line_detailDB
	if err := db.First(&line_detailDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var line_detailAPI orm.Line_detailAPI
	line_detailAPI.ID = line_detailDB.ID
	line_detailAPI.Line_detailPointersEncoding = line_detailDB.Line_detailPointersEncoding
	line_detailDB.CopyBasicFieldsToLine_detail_WOP(&line_detailAPI.Line_detail_WOP)

	c.JSON(http.StatusOK, line_detailAPI)
}

// UpdateLine_detail
//
// swagger:route PATCH /line_details/{ID} line_details updateLine_detail
//
// # Update a line_detail
//
// Responses:
// default: genericError
//
//	200: line_detailDBResponse
func (controller *Controller) UpdateLine_detail(c *gin.Context) {

	mutexLine_detail.Lock()
	defer mutexLine_detail.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLine_detail", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_detail.GetDB()

	// Validate input
	var input orm.Line_detailAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var line_detailDB orm.Line_detailDB

	// fetch the line_detail
	query := db.First(&line_detailDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	line_detailDB.CopyBasicFieldsFromLine_detail_WOP(&input.Line_detail_WOP)
	line_detailDB.Line_detailPointersEncoding = input.Line_detailPointersEncoding

	query = db.Model(&line_detailDB).Updates(line_detailDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	line_detailNew := new(models.Line_detail)
	line_detailDB.CopyBasicFieldsToLine_detail(line_detailNew)

	// redeem pointers
	line_detailDB.DecodePointers(backRepo, line_detailNew)

	// get stage instance from DB instance, and call callback function
	line_detailOld := backRepo.BackRepoLine_detail.Map_Line_detailDBID_Line_detailPtr[line_detailDB.ID]
	if line_detailOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), line_detailOld, line_detailNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the line_detailDB
	c.JSON(http.StatusOK, line_detailDB)
}

// DeleteLine_detail
//
// swagger:route DELETE /line_details/{ID} line_details deleteLine_detail
//
// # Delete a line_detail
//
// default: genericError
//
//	200: line_detailDBResponse
func (controller *Controller) DeleteLine_detail(c *gin.Context) {

	mutexLine_detail.Lock()
	defer mutexLine_detail.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLine_detail", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLine_detail.GetDB()

	// Get model if exist
	var line_detailDB orm.Line_detailDB
	if err := db.First(&line_detailDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&line_detailDB)

	// get an instance (not staged) from DB instance, and call callback function
	line_detailDeleted := new(models.Line_detail)
	line_detailDB.CopyBasicFieldsToLine_detail(line_detailDeleted)

	// get stage instance from DB instance, and call callback function
	line_detailStaged := backRepo.BackRepoLine_detail.Map_Line_detailDBID_Line_detailPtr[line_detailDB.ID]
	if line_detailStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), line_detailStaged, line_detailDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
