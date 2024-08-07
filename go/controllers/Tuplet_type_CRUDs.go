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
var __Tuplet_type__dummysDeclaration__ models.Tuplet_type
var __Tuplet_type_time__dummyDeclaration time.Duration

var mutexTuplet_type sync.Mutex

// An Tuplet_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTuplet_type updateTuplet_type deleteTuplet_type
type Tuplet_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Tuplet_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTuplet_type updateTuplet_type
type Tuplet_typeInput struct {
	// The Tuplet_type to submit or modify
	// in: body
	Tuplet_type *orm.Tuplet_typeAPI
}

// GetTuplet_types
//
// swagger:route GET /tuplet_types tuplet_types getTuplet_types
//
// # Get all tuplet_types
//
// Responses:
// default: genericError
//
//	200: tuplet_typeDBResponse
func (controller *Controller) GetTuplet_types(c *gin.Context) {

	// source slice
	var tuplet_typeDBs []orm.Tuplet_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_type.GetDB()

	query := db.Find(&tuplet_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tuplet_typeAPIs := make([]orm.Tuplet_typeAPI, 0)

	// for each tuplet_type, update fields from the database nullable fields
	for idx := range tuplet_typeDBs {
		tuplet_typeDB := &tuplet_typeDBs[idx]
		_ = tuplet_typeDB
		var tuplet_typeAPI orm.Tuplet_typeAPI

		// insertion point for updating fields
		tuplet_typeAPI.ID = tuplet_typeDB.ID
		tuplet_typeDB.CopyBasicFieldsToTuplet_type_WOP(&tuplet_typeAPI.Tuplet_type_WOP)
		tuplet_typeAPI.Tuplet_typePointersEncoding = tuplet_typeDB.Tuplet_typePointersEncoding
		tuplet_typeAPIs = append(tuplet_typeAPIs, tuplet_typeAPI)
	}

	c.JSON(http.StatusOK, tuplet_typeAPIs)
}

// PostTuplet_type
//
// swagger:route POST /tuplet_types tuplet_types postTuplet_type
//
// Creates a tuplet_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTuplet_type(c *gin.Context) {

	mutexTuplet_type.Lock()
	defer mutexTuplet_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTuplet_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_type.GetDB()

	// Validate input
	var input orm.Tuplet_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tuplet_type
	tuplet_typeDB := orm.Tuplet_typeDB{}
	tuplet_typeDB.Tuplet_typePointersEncoding = input.Tuplet_typePointersEncoding
	tuplet_typeDB.CopyBasicFieldsFromTuplet_type_WOP(&input.Tuplet_type_WOP)

	query := db.Create(&tuplet_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTuplet_type.CheckoutPhaseOneInstance(&tuplet_typeDB)
	tuplet_type := backRepo.BackRepoTuplet_type.Map_Tuplet_typeDBID_Tuplet_typePtr[tuplet_typeDB.ID]

	if tuplet_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tuplet_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tuplet_typeDB)
}

// GetTuplet_type
//
// swagger:route GET /tuplet_types/{ID} tuplet_types getTuplet_type
//
// Gets the details for a tuplet_type.
//
// Responses:
// default: genericError
//
//	200: tuplet_typeDBResponse
func (controller *Controller) GetTuplet_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_type.GetDB()

	// Get tuplet_typeDB in DB
	var tuplet_typeDB orm.Tuplet_typeDB
	if err := db.First(&tuplet_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tuplet_typeAPI orm.Tuplet_typeAPI
	tuplet_typeAPI.ID = tuplet_typeDB.ID
	tuplet_typeAPI.Tuplet_typePointersEncoding = tuplet_typeDB.Tuplet_typePointersEncoding
	tuplet_typeDB.CopyBasicFieldsToTuplet_type_WOP(&tuplet_typeAPI.Tuplet_type_WOP)

	c.JSON(http.StatusOK, tuplet_typeAPI)
}

// UpdateTuplet_type
//
// swagger:route PATCH /tuplet_types/{ID} tuplet_types updateTuplet_type
//
// # Update a tuplet_type
//
// Responses:
// default: genericError
//
//	200: tuplet_typeDBResponse
func (controller *Controller) UpdateTuplet_type(c *gin.Context) {

	mutexTuplet_type.Lock()
	defer mutexTuplet_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTuplet_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_type.GetDB()

	// Validate input
	var input orm.Tuplet_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tuplet_typeDB orm.Tuplet_typeDB

	// fetch the tuplet_type
	query := db.First(&tuplet_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tuplet_typeDB.CopyBasicFieldsFromTuplet_type_WOP(&input.Tuplet_type_WOP)
	tuplet_typeDB.Tuplet_typePointersEncoding = input.Tuplet_typePointersEncoding

	query = db.Model(&tuplet_typeDB).Updates(tuplet_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_typeNew := new(models.Tuplet_type)
	tuplet_typeDB.CopyBasicFieldsToTuplet_type(tuplet_typeNew)

	// redeem pointers
	tuplet_typeDB.DecodePointers(backRepo, tuplet_typeNew)

	// get stage instance from DB instance, and call callback function
	tuplet_typeOld := backRepo.BackRepoTuplet_type.Map_Tuplet_typeDBID_Tuplet_typePtr[tuplet_typeDB.ID]
	if tuplet_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tuplet_typeOld, tuplet_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tuplet_typeDB
	c.JSON(http.StatusOK, tuplet_typeDB)
}

// DeleteTuplet_type
//
// swagger:route DELETE /tuplet_types/{ID} tuplet_types deleteTuplet_type
//
// # Delete a tuplet_type
//
// default: genericError
//
//	200: tuplet_typeDBResponse
func (controller *Controller) DeleteTuplet_type(c *gin.Context) {

	mutexTuplet_type.Lock()
	defer mutexTuplet_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTuplet_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_type.GetDB()

	// Get model if exist
	var tuplet_typeDB orm.Tuplet_typeDB
	if err := db.First(&tuplet_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tuplet_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_typeDeleted := new(models.Tuplet_type)
	tuplet_typeDB.CopyBasicFieldsToTuplet_type(tuplet_typeDeleted)

	// get stage instance from DB instance, and call callback function
	tuplet_typeStaged := backRepo.BackRepoTuplet_type.Map_Tuplet_typeDBID_Tuplet_typePtr[tuplet_typeDB.ID]
	if tuplet_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tuplet_typeStaged, tuplet_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
