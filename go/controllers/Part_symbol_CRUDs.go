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
var __Part_symbol__dummysDeclaration__ models.Part_symbol
var __Part_symbol_time__dummyDeclaration time.Duration

var mutexPart_symbol sync.Mutex

// An Part_symbolID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPart_symbol updatePart_symbol deletePart_symbol
type Part_symbolID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Part_symbolInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPart_symbol updatePart_symbol
type Part_symbolInput struct {
	// The Part_symbol to submit or modify
	// in: body
	Part_symbol *orm.Part_symbolAPI
}

// GetPart_symbols
//
// swagger:route GET /part_symbols part_symbols getPart_symbols
//
// # Get all part_symbols
//
// Responses:
// default: genericError
//
//	200: part_symbolDBResponse
func (controller *Controller) GetPart_symbols(c *gin.Context) {

	// source slice
	var part_symbolDBs []orm.Part_symbolDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_symbols", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_symbol.GetDB()

	query := db.Find(&part_symbolDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	part_symbolAPIs := make([]orm.Part_symbolAPI, 0)

	// for each part_symbol, update fields from the database nullable fields
	for idx := range part_symbolDBs {
		part_symbolDB := &part_symbolDBs[idx]
		_ = part_symbolDB
		var part_symbolAPI orm.Part_symbolAPI

		// insertion point for updating fields
		part_symbolAPI.ID = part_symbolDB.ID
		part_symbolDB.CopyBasicFieldsToPart_symbol_WOP(&part_symbolAPI.Part_symbol_WOP)
		part_symbolAPI.Part_symbolPointersEncoding = part_symbolDB.Part_symbolPointersEncoding
		part_symbolAPIs = append(part_symbolAPIs, part_symbolAPI)
	}

	c.JSON(http.StatusOK, part_symbolAPIs)
}

// PostPart_symbol
//
// swagger:route POST /part_symbols part_symbols postPart_symbol
//
// Creates a part_symbol
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPart_symbol(c *gin.Context) {

	mutexPart_symbol.Lock()
	defer mutexPart_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPart_symbols", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_symbol.GetDB()

	// Validate input
	var input orm.Part_symbolAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create part_symbol
	part_symbolDB := orm.Part_symbolDB{}
	part_symbolDB.Part_symbolPointersEncoding = input.Part_symbolPointersEncoding
	part_symbolDB.CopyBasicFieldsFromPart_symbol_WOP(&input.Part_symbol_WOP)

	query := db.Create(&part_symbolDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPart_symbol.CheckoutPhaseOneInstance(&part_symbolDB)
	part_symbol := backRepo.BackRepoPart_symbol.Map_Part_symbolDBID_Part_symbolPtr[part_symbolDB.ID]

	if part_symbol != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), part_symbol)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, part_symbolDB)
}

// GetPart_symbol
//
// swagger:route GET /part_symbols/{ID} part_symbols getPart_symbol
//
// Gets the details for a part_symbol.
//
// Responses:
// default: genericError
//
//	200: part_symbolDBResponse
func (controller *Controller) GetPart_symbol(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_symbol.GetDB()

	// Get part_symbolDB in DB
	var part_symbolDB orm.Part_symbolDB
	if err := db.First(&part_symbolDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var part_symbolAPI orm.Part_symbolAPI
	part_symbolAPI.ID = part_symbolDB.ID
	part_symbolAPI.Part_symbolPointersEncoding = part_symbolDB.Part_symbolPointersEncoding
	part_symbolDB.CopyBasicFieldsToPart_symbol_WOP(&part_symbolAPI.Part_symbol_WOP)

	c.JSON(http.StatusOK, part_symbolAPI)
}

// UpdatePart_symbol
//
// swagger:route PATCH /part_symbols/{ID} part_symbols updatePart_symbol
//
// # Update a part_symbol
//
// Responses:
// default: genericError
//
//	200: part_symbolDBResponse
func (controller *Controller) UpdatePart_symbol(c *gin.Context) {

	mutexPart_symbol.Lock()
	defer mutexPart_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePart_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_symbol.GetDB()

	// Validate input
	var input orm.Part_symbolAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var part_symbolDB orm.Part_symbolDB

	// fetch the part_symbol
	query := db.First(&part_symbolDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	part_symbolDB.CopyBasicFieldsFromPart_symbol_WOP(&input.Part_symbol_WOP)
	part_symbolDB.Part_symbolPointersEncoding = input.Part_symbolPointersEncoding

	query = db.Model(&part_symbolDB).Updates(part_symbolDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	part_symbolNew := new(models.Part_symbol)
	part_symbolDB.CopyBasicFieldsToPart_symbol(part_symbolNew)

	// redeem pointers
	part_symbolDB.DecodePointers(backRepo, part_symbolNew)

	// get stage instance from DB instance, and call callback function
	part_symbolOld := backRepo.BackRepoPart_symbol.Map_Part_symbolDBID_Part_symbolPtr[part_symbolDB.ID]
	if part_symbolOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), part_symbolOld, part_symbolNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the part_symbolDB
	c.JSON(http.StatusOK, part_symbolDB)
}

// DeletePart_symbol
//
// swagger:route DELETE /part_symbols/{ID} part_symbols deletePart_symbol
//
// # Delete a part_symbol
//
// default: genericError
//
//	200: part_symbolDBResponse
func (controller *Controller) DeletePart_symbol(c *gin.Context) {

	mutexPart_symbol.Lock()
	defer mutexPart_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePart_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_symbol.GetDB()

	// Get model if exist
	var part_symbolDB orm.Part_symbolDB
	if err := db.First(&part_symbolDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&part_symbolDB)

	// get an instance (not staged) from DB instance, and call callback function
	part_symbolDeleted := new(models.Part_symbol)
	part_symbolDB.CopyBasicFieldsToPart_symbol(part_symbolDeleted)

	// get stage instance from DB instance, and call callback function
	part_symbolStaged := backRepo.BackRepoPart_symbol.Map_Part_symbolDBID_Part_symbolPtr[part_symbolDB.ID]
	if part_symbolStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), part_symbolStaged, part_symbolDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
