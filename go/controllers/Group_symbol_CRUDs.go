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
var __Group_symbol__dummysDeclaration__ models.Group_symbol
var __Group_symbol_time__dummyDeclaration time.Duration

var mutexGroup_symbol sync.Mutex

// An Group_symbolID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGroup_symbol updateGroup_symbol deleteGroup_symbol
type Group_symbolID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Group_symbolInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGroup_symbol updateGroup_symbol
type Group_symbolInput struct {
	// The Group_symbol to submit or modify
	// in: body
	Group_symbol *orm.Group_symbolAPI
}

// GetGroup_symbols
//
// swagger:route GET /group_symbols group_symbols getGroup_symbols
//
// # Get all group_symbols
//
// Responses:
// default: genericError
//
//	200: group_symbolDBResponse
func (controller *Controller) GetGroup_symbols(c *gin.Context) {

	// source slice
	var group_symbolDBs []orm.Group_symbolDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroup_symbols", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_symbol.GetDB()

	query := db.Find(&group_symbolDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	group_symbolAPIs := make([]orm.Group_symbolAPI, 0)

	// for each group_symbol, update fields from the database nullable fields
	for idx := range group_symbolDBs {
		group_symbolDB := &group_symbolDBs[idx]
		_ = group_symbolDB
		var group_symbolAPI orm.Group_symbolAPI

		// insertion point for updating fields
		group_symbolAPI.ID = group_symbolDB.ID
		group_symbolDB.CopyBasicFieldsToGroup_symbol_WOP(&group_symbolAPI.Group_symbol_WOP)
		group_symbolAPI.Group_symbolPointersEncoding = group_symbolDB.Group_symbolPointersEncoding
		group_symbolAPIs = append(group_symbolAPIs, group_symbolAPI)
	}

	c.JSON(http.StatusOK, group_symbolAPIs)
}

// PostGroup_symbol
//
// swagger:route POST /group_symbols group_symbols postGroup_symbol
//
// Creates a group_symbol
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGroup_symbol(c *gin.Context) {

	mutexGroup_symbol.Lock()
	defer mutexGroup_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroup_symbols", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_symbol.GetDB()

	// Validate input
	var input orm.Group_symbolAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create group_symbol
	group_symbolDB := orm.Group_symbolDB{}
	group_symbolDB.Group_symbolPointersEncoding = input.Group_symbolPointersEncoding
	group_symbolDB.CopyBasicFieldsFromGroup_symbol_WOP(&input.Group_symbol_WOP)

	query := db.Create(&group_symbolDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGroup_symbol.CheckoutPhaseOneInstance(&group_symbolDB)
	group_symbol := backRepo.BackRepoGroup_symbol.Map_Group_symbolDBID_Group_symbolPtr[group_symbolDB.ID]

	if group_symbol != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), group_symbol)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, group_symbolDB)
}

// GetGroup_symbol
//
// swagger:route GET /group_symbols/{ID} group_symbols getGroup_symbol
//
// Gets the details for a group_symbol.
//
// Responses:
// default: genericError
//
//	200: group_symbolDBResponse
func (controller *Controller) GetGroup_symbol(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroup_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_symbol.GetDB()

	// Get group_symbolDB in DB
	var group_symbolDB orm.Group_symbolDB
	if err := db.First(&group_symbolDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var group_symbolAPI orm.Group_symbolAPI
	group_symbolAPI.ID = group_symbolDB.ID
	group_symbolAPI.Group_symbolPointersEncoding = group_symbolDB.Group_symbolPointersEncoding
	group_symbolDB.CopyBasicFieldsToGroup_symbol_WOP(&group_symbolAPI.Group_symbol_WOP)

	c.JSON(http.StatusOK, group_symbolAPI)
}

// UpdateGroup_symbol
//
// swagger:route PATCH /group_symbols/{ID} group_symbols updateGroup_symbol
//
// # Update a group_symbol
//
// Responses:
// default: genericError
//
//	200: group_symbolDBResponse
func (controller *Controller) UpdateGroup_symbol(c *gin.Context) {

	mutexGroup_symbol.Lock()
	defer mutexGroup_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGroup_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_symbol.GetDB()

	// Validate input
	var input orm.Group_symbolAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var group_symbolDB orm.Group_symbolDB

	// fetch the group_symbol
	query := db.First(&group_symbolDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	group_symbolDB.CopyBasicFieldsFromGroup_symbol_WOP(&input.Group_symbol_WOP)
	group_symbolDB.Group_symbolPointersEncoding = input.Group_symbolPointersEncoding

	query = db.Model(&group_symbolDB).Updates(group_symbolDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	group_symbolNew := new(models.Group_symbol)
	group_symbolDB.CopyBasicFieldsToGroup_symbol(group_symbolNew)

	// redeem pointers
	group_symbolDB.DecodePointers(backRepo, group_symbolNew)

	// get stage instance from DB instance, and call callback function
	group_symbolOld := backRepo.BackRepoGroup_symbol.Map_Group_symbolDBID_Group_symbolPtr[group_symbolDB.ID]
	if group_symbolOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), group_symbolOld, group_symbolNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the group_symbolDB
	c.JSON(http.StatusOK, group_symbolDB)
}

// DeleteGroup_symbol
//
// swagger:route DELETE /group_symbols/{ID} group_symbols deleteGroup_symbol
//
// # Delete a group_symbol
//
// default: genericError
//
//	200: group_symbolDBResponse
func (controller *Controller) DeleteGroup_symbol(c *gin.Context) {

	mutexGroup_symbol.Lock()
	defer mutexGroup_symbol.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGroup_symbol", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_symbol.GetDB()

	// Get model if exist
	var group_symbolDB orm.Group_symbolDB
	if err := db.First(&group_symbolDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&group_symbolDB)

	// get an instance (not staged) from DB instance, and call callback function
	group_symbolDeleted := new(models.Group_symbol)
	group_symbolDB.CopyBasicFieldsToGroup_symbol(group_symbolDeleted)

	// get stage instance from DB instance, and call callback function
	group_symbolStaged := backRepo.BackRepoGroup_symbol.Map_Group_symbolDBID_Group_symbolPtr[group_symbolDB.ID]
	if group_symbolStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), group_symbolStaged, group_symbolDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
