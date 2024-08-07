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
var __Tuplet__dummysDeclaration__ models.Tuplet
var __Tuplet_time__dummyDeclaration time.Duration

var mutexTuplet sync.Mutex

// An TupletID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTuplet updateTuplet deleteTuplet
type TupletID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TupletInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTuplet updateTuplet
type TupletInput struct {
	// The Tuplet to submit or modify
	// in: body
	Tuplet *orm.TupletAPI
}

// GetTuplets
//
// swagger:route GET /tuplets tuplets getTuplets
//
// # Get all tuplets
//
// Responses:
// default: genericError
//
//	200: tupletDBResponse
func (controller *Controller) GetTuplets(c *gin.Context) {

	// source slice
	var tupletDBs []orm.TupletDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet.GetDB()

	query := db.Find(&tupletDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tupletAPIs := make([]orm.TupletAPI, 0)

	// for each tuplet, update fields from the database nullable fields
	for idx := range tupletDBs {
		tupletDB := &tupletDBs[idx]
		_ = tupletDB
		var tupletAPI orm.TupletAPI

		// insertion point for updating fields
		tupletAPI.ID = tupletDB.ID
		tupletDB.CopyBasicFieldsToTuplet_WOP(&tupletAPI.Tuplet_WOP)
		tupletAPI.TupletPointersEncoding = tupletDB.TupletPointersEncoding
		tupletAPIs = append(tupletAPIs, tupletAPI)
	}

	c.JSON(http.StatusOK, tupletAPIs)
}

// PostTuplet
//
// swagger:route POST /tuplets tuplets postTuplet
//
// Creates a tuplet
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTuplet(c *gin.Context) {

	mutexTuplet.Lock()
	defer mutexTuplet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTuplets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet.GetDB()

	// Validate input
	var input orm.TupletAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tuplet
	tupletDB := orm.TupletDB{}
	tupletDB.TupletPointersEncoding = input.TupletPointersEncoding
	tupletDB.CopyBasicFieldsFromTuplet_WOP(&input.Tuplet_WOP)

	query := db.Create(&tupletDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTuplet.CheckoutPhaseOneInstance(&tupletDB)
	tuplet := backRepo.BackRepoTuplet.Map_TupletDBID_TupletPtr[tupletDB.ID]

	if tuplet != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tuplet)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tupletDB)
}

// GetTuplet
//
// swagger:route GET /tuplets/{ID} tuplets getTuplet
//
// Gets the details for a tuplet.
//
// Responses:
// default: genericError
//
//	200: tupletDBResponse
func (controller *Controller) GetTuplet(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet.GetDB()

	// Get tupletDB in DB
	var tupletDB orm.TupletDB
	if err := db.First(&tupletDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tupletAPI orm.TupletAPI
	tupletAPI.ID = tupletDB.ID
	tupletAPI.TupletPointersEncoding = tupletDB.TupletPointersEncoding
	tupletDB.CopyBasicFieldsToTuplet_WOP(&tupletAPI.Tuplet_WOP)

	c.JSON(http.StatusOK, tupletAPI)
}

// UpdateTuplet
//
// swagger:route PATCH /tuplets/{ID} tuplets updateTuplet
//
// # Update a tuplet
//
// Responses:
// default: genericError
//
//	200: tupletDBResponse
func (controller *Controller) UpdateTuplet(c *gin.Context) {

	mutexTuplet.Lock()
	defer mutexTuplet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTuplet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet.GetDB()

	// Validate input
	var input orm.TupletAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tupletDB orm.TupletDB

	// fetch the tuplet
	query := db.First(&tupletDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tupletDB.CopyBasicFieldsFromTuplet_WOP(&input.Tuplet_WOP)
	tupletDB.TupletPointersEncoding = input.TupletPointersEncoding

	query = db.Model(&tupletDB).Updates(tupletDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tupletNew := new(models.Tuplet)
	tupletDB.CopyBasicFieldsToTuplet(tupletNew)

	// redeem pointers
	tupletDB.DecodePointers(backRepo, tupletNew)

	// get stage instance from DB instance, and call callback function
	tupletOld := backRepo.BackRepoTuplet.Map_TupletDBID_TupletPtr[tupletDB.ID]
	if tupletOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tupletOld, tupletNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tupletDB
	c.JSON(http.StatusOK, tupletDB)
}

// DeleteTuplet
//
// swagger:route DELETE /tuplets/{ID} tuplets deleteTuplet
//
// # Delete a tuplet
//
// default: genericError
//
//	200: tupletDBResponse
func (controller *Controller) DeleteTuplet(c *gin.Context) {

	mutexTuplet.Lock()
	defer mutexTuplet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTuplet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet.GetDB()

	// Get model if exist
	var tupletDB orm.TupletDB
	if err := db.First(&tupletDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tupletDB)

	// get an instance (not staged) from DB instance, and call callback function
	tupletDeleted := new(models.Tuplet)
	tupletDB.CopyBasicFieldsToTuplet(tupletDeleted)

	// get stage instance from DB instance, and call callback function
	tupletStaged := backRepo.BackRepoTuplet.Map_TupletDBID_TupletPtr[tupletDB.ID]
	if tupletStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tupletStaged, tupletDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
