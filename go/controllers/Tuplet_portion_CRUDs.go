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
var __Tuplet_portion__dummysDeclaration__ models.Tuplet_portion
var __Tuplet_portion_time__dummyDeclaration time.Duration

var mutexTuplet_portion sync.Mutex

// An Tuplet_portionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTuplet_portion updateTuplet_portion deleteTuplet_portion
type Tuplet_portionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Tuplet_portionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTuplet_portion updateTuplet_portion
type Tuplet_portionInput struct {
	// The Tuplet_portion to submit or modify
	// in: body
	Tuplet_portion *orm.Tuplet_portionAPI
}

// GetTuplet_portions
//
// swagger:route GET /tuplet_portions tuplet_portions getTuplet_portions
//
// # Get all tuplet_portions
//
// Responses:
// default: genericError
//
//	200: tuplet_portionDBResponse
func (controller *Controller) GetTuplet_portions(c *gin.Context) {

	// source slice
	var tuplet_portionDBs []orm.Tuplet_portionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_portions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_portion.GetDB()

	query := db.Find(&tuplet_portionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tuplet_portionAPIs := make([]orm.Tuplet_portionAPI, 0)

	// for each tuplet_portion, update fields from the database nullable fields
	for idx := range tuplet_portionDBs {
		tuplet_portionDB := &tuplet_portionDBs[idx]
		_ = tuplet_portionDB
		var tuplet_portionAPI orm.Tuplet_portionAPI

		// insertion point for updating fields
		tuplet_portionAPI.ID = tuplet_portionDB.ID
		tuplet_portionDB.CopyBasicFieldsToTuplet_portion_WOP(&tuplet_portionAPI.Tuplet_portion_WOP)
		tuplet_portionAPI.Tuplet_portionPointersEncoding = tuplet_portionDB.Tuplet_portionPointersEncoding
		tuplet_portionAPIs = append(tuplet_portionAPIs, tuplet_portionAPI)
	}

	c.JSON(http.StatusOK, tuplet_portionAPIs)
}

// PostTuplet_portion
//
// swagger:route POST /tuplet_portions tuplet_portions postTuplet_portion
//
// Creates a tuplet_portion
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTuplet_portion(c *gin.Context) {

	mutexTuplet_portion.Lock()
	defer mutexTuplet_portion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTuplet_portions", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_portion.GetDB()

	// Validate input
	var input orm.Tuplet_portionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tuplet_portion
	tuplet_portionDB := orm.Tuplet_portionDB{}
	tuplet_portionDB.Tuplet_portionPointersEncoding = input.Tuplet_portionPointersEncoding
	tuplet_portionDB.CopyBasicFieldsFromTuplet_portion_WOP(&input.Tuplet_portion_WOP)

	query := db.Create(&tuplet_portionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTuplet_portion.CheckoutPhaseOneInstance(&tuplet_portionDB)
	tuplet_portion := backRepo.BackRepoTuplet_portion.Map_Tuplet_portionDBID_Tuplet_portionPtr[tuplet_portionDB.ID]

	if tuplet_portion != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tuplet_portion)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tuplet_portionDB)
}

// GetTuplet_portion
//
// swagger:route GET /tuplet_portions/{ID} tuplet_portions getTuplet_portion
//
// Gets the details for a tuplet_portion.
//
// Responses:
// default: genericError
//
//	200: tuplet_portionDBResponse
func (controller *Controller) GetTuplet_portion(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_portion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_portion.GetDB()

	// Get tuplet_portionDB in DB
	var tuplet_portionDB orm.Tuplet_portionDB
	if err := db.First(&tuplet_portionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tuplet_portionAPI orm.Tuplet_portionAPI
	tuplet_portionAPI.ID = tuplet_portionDB.ID
	tuplet_portionAPI.Tuplet_portionPointersEncoding = tuplet_portionDB.Tuplet_portionPointersEncoding
	tuplet_portionDB.CopyBasicFieldsToTuplet_portion_WOP(&tuplet_portionAPI.Tuplet_portion_WOP)

	c.JSON(http.StatusOK, tuplet_portionAPI)
}

// UpdateTuplet_portion
//
// swagger:route PATCH /tuplet_portions/{ID} tuplet_portions updateTuplet_portion
//
// # Update a tuplet_portion
//
// Responses:
// default: genericError
//
//	200: tuplet_portionDBResponse
func (controller *Controller) UpdateTuplet_portion(c *gin.Context) {

	mutexTuplet_portion.Lock()
	defer mutexTuplet_portion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTuplet_portion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_portion.GetDB()

	// Validate input
	var input orm.Tuplet_portionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tuplet_portionDB orm.Tuplet_portionDB

	// fetch the tuplet_portion
	query := db.First(&tuplet_portionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tuplet_portionDB.CopyBasicFieldsFromTuplet_portion_WOP(&input.Tuplet_portion_WOP)
	tuplet_portionDB.Tuplet_portionPointersEncoding = input.Tuplet_portionPointersEncoding

	query = db.Model(&tuplet_portionDB).Updates(tuplet_portionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_portionNew := new(models.Tuplet_portion)
	tuplet_portionDB.CopyBasicFieldsToTuplet_portion(tuplet_portionNew)

	// redeem pointers
	tuplet_portionDB.DecodePointers(backRepo, tuplet_portionNew)

	// get stage instance from DB instance, and call callback function
	tuplet_portionOld := backRepo.BackRepoTuplet_portion.Map_Tuplet_portionDBID_Tuplet_portionPtr[tuplet_portionDB.ID]
	if tuplet_portionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tuplet_portionOld, tuplet_portionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tuplet_portionDB
	c.JSON(http.StatusOK, tuplet_portionDB)
}

// DeleteTuplet_portion
//
// swagger:route DELETE /tuplet_portions/{ID} tuplet_portions deleteTuplet_portion
//
// # Delete a tuplet_portion
//
// default: genericError
//
//	200: tuplet_portionDBResponse
func (controller *Controller) DeleteTuplet_portion(c *gin.Context) {

	mutexTuplet_portion.Lock()
	defer mutexTuplet_portion.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTuplet_portion", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_portion.GetDB()

	// Get model if exist
	var tuplet_portionDB orm.Tuplet_portionDB
	if err := db.First(&tuplet_portionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tuplet_portionDB)

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_portionDeleted := new(models.Tuplet_portion)
	tuplet_portionDB.CopyBasicFieldsToTuplet_portion(tuplet_portionDeleted)

	// get stage instance from DB instance, and call callback function
	tuplet_portionStaged := backRepo.BackRepoTuplet_portion.Map_Tuplet_portionDBID_Tuplet_portionPtr[tuplet_portionDB.ID]
	if tuplet_portionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tuplet_portionStaged, tuplet_portionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
