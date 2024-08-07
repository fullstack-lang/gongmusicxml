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
var __Formatted_symbol__dummysDeclaration__ models.Formatted_symbol
var __Formatted_symbol_time__dummyDeclaration time.Duration

var mutexFormatted_symbol sync.Mutex

// An Formatted_symbolID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFormatted_symbol updateFormatted_symbol deleteFormatted_symbol
type Formatted_symbolID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Formatted_symbolInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFormatted_symbol updateFormatted_symbol
type Formatted_symbolInput struct {
	// The Formatted_symbol to submit or modify
	// in: body
	Formatted_symbol *orm.Formatted_symbolAPI
}

// GetFormatted_symbols
//
// swagger:route GET /formatted_symbols formatted_symbols getFormatted_symbols
//
// # Get all formatted_symbols
//
// Responses:
// default: genericError
//
//	200: formatted_symbolDBResponse
func (controller *Controller) GetFormatted_symbols(c *gin.Context) {

	// source slice
	var formatted_symbolDBs []orm.Formatted_symbolDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormatted_symbols", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol.GetDB()

	query := db.Find(&formatted_symbolDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	formatted_symbolAPIs := make([]orm.Formatted_symbolAPI, 0)

	// for each formatted_symbol, update fields from the database nullable fields
	for idx := range formatted_symbolDBs {
		formatted_symbolDB := &formatted_symbolDBs[idx]
		_ = formatted_symbolDB
		var formatted_symbolAPI orm.Formatted_symbolAPI

		// insertion point for updating fields
		formatted_symbolAPI.ID = formatted_symbolDB.ID
		formatted_symbolDB.CopyBasicFieldsToFormatted_symbol_WOP(&formatted_symbolAPI.Formatted_symbol_WOP)
		formatted_symbolAPI.Formatted_symbolPointersEncoding = formatted_symbolDB.Formatted_symbolPointersEncoding
		formatted_symbolAPIs = append(formatted_symbolAPIs, formatted_symbolAPI)
	}

	c.JSON(http.StatusOK, formatted_symbolAPIs)
}

// PostFormatted_symbol
//
// swagger:route POST /formatted_symbols formatted_symbols postFormatted_symbol
//
// Creates a formatted_symbol
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFormatted_symbol(c *gin.Context) {

	mutexFormatted_symbol.Lock()
	defer mutexFormatted_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFormatted_symbols", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol.GetDB()

	// Validate input
	var input orm.Formatted_symbolAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create formatted_symbol
	formatted_symbolDB := orm.Formatted_symbolDB{}
	formatted_symbolDB.Formatted_symbolPointersEncoding = input.Formatted_symbolPointersEncoding
	formatted_symbolDB.CopyBasicFieldsFromFormatted_symbol_WOP(&input.Formatted_symbol_WOP)

	query := db.Create(&formatted_symbolDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFormatted_symbol.CheckoutPhaseOneInstance(&formatted_symbolDB)
	formatted_symbol := backRepo.BackRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID]

	if formatted_symbol != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), formatted_symbol)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, formatted_symbolDB)
}

// GetFormatted_symbol
//
// swagger:route GET /formatted_symbols/{ID} formatted_symbols getFormatted_symbol
//
// Gets the details for a formatted_symbol.
//
// Responses:
// default: genericError
//
//	200: formatted_symbolDBResponse
func (controller *Controller) GetFormatted_symbol(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFormatted_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol.GetDB()

	// Get formatted_symbolDB in DB
	var formatted_symbolDB orm.Formatted_symbolDB
	if err := db.First(&formatted_symbolDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var formatted_symbolAPI orm.Formatted_symbolAPI
	formatted_symbolAPI.ID = formatted_symbolDB.ID
	formatted_symbolAPI.Formatted_symbolPointersEncoding = formatted_symbolDB.Formatted_symbolPointersEncoding
	formatted_symbolDB.CopyBasicFieldsToFormatted_symbol_WOP(&formatted_symbolAPI.Formatted_symbol_WOP)

	c.JSON(http.StatusOK, formatted_symbolAPI)
}

// UpdateFormatted_symbol
//
// swagger:route PATCH /formatted_symbols/{ID} formatted_symbols updateFormatted_symbol
//
// # Update a formatted_symbol
//
// Responses:
// default: genericError
//
//	200: formatted_symbolDBResponse
func (controller *Controller) UpdateFormatted_symbol(c *gin.Context) {

	mutexFormatted_symbol.Lock()
	defer mutexFormatted_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFormatted_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol.GetDB()

	// Validate input
	var input orm.Formatted_symbolAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var formatted_symbolDB orm.Formatted_symbolDB

	// fetch the formatted_symbol
	query := db.First(&formatted_symbolDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	formatted_symbolDB.CopyBasicFieldsFromFormatted_symbol_WOP(&input.Formatted_symbol_WOP)
	formatted_symbolDB.Formatted_symbolPointersEncoding = input.Formatted_symbolPointersEncoding

	query = db.Model(&formatted_symbolDB).Updates(formatted_symbolDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	formatted_symbolNew := new(models.Formatted_symbol)
	formatted_symbolDB.CopyBasicFieldsToFormatted_symbol(formatted_symbolNew)

	// redeem pointers
	formatted_symbolDB.DecodePointers(backRepo, formatted_symbolNew)

	// get stage instance from DB instance, and call callback function
	formatted_symbolOld := backRepo.BackRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID]
	if formatted_symbolOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), formatted_symbolOld, formatted_symbolNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the formatted_symbolDB
	c.JSON(http.StatusOK, formatted_symbolDB)
}

// DeleteFormatted_symbol
//
// swagger:route DELETE /formatted_symbols/{ID} formatted_symbols deleteFormatted_symbol
//
// # Delete a formatted_symbol
//
// default: genericError
//
//	200: formatted_symbolDBResponse
func (controller *Controller) DeleteFormatted_symbol(c *gin.Context) {

	mutexFormatted_symbol.Lock()
	defer mutexFormatted_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFormatted_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFormatted_symbol.GetDB()

	// Get model if exist
	var formatted_symbolDB orm.Formatted_symbolDB
	if err := db.First(&formatted_symbolDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&formatted_symbolDB)

	// get an instance (not staged) from DB instance, and call callback function
	formatted_symbolDeleted := new(models.Formatted_symbol)
	formatted_symbolDB.CopyBasicFieldsToFormatted_symbol(formatted_symbolDeleted)

	// get stage instance from DB instance, and call callback function
	formatted_symbolStaged := backRepo.BackRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID]
	if formatted_symbolStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), formatted_symbolStaged, formatted_symbolDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
