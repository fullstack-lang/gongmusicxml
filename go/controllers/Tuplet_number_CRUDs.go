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
var __Tuplet_number__dummysDeclaration__ models.Tuplet_number
var __Tuplet_number_time__dummyDeclaration time.Duration

var mutexTuplet_number sync.Mutex

// An Tuplet_numberID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTuplet_number updateTuplet_number deleteTuplet_number
type Tuplet_numberID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Tuplet_numberInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTuplet_number updateTuplet_number
type Tuplet_numberInput struct {
	// The Tuplet_number to submit or modify
	// in: body
	Tuplet_number *orm.Tuplet_numberAPI
}

// GetTuplet_numbers
//
// swagger:route GET /tuplet_numbers tuplet_numbers getTuplet_numbers
//
// # Get all tuplet_numbers
//
// Responses:
// default: genericError
//
//	200: tuplet_numberDBResponse
func (controller *Controller) GetTuplet_numbers(c *gin.Context) {

	// source slice
	var tuplet_numberDBs []orm.Tuplet_numberDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_numbers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_number.GetDB()

	query := db.Find(&tuplet_numberDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tuplet_numberAPIs := make([]orm.Tuplet_numberAPI, 0)

	// for each tuplet_number, update fields from the database nullable fields
	for idx := range tuplet_numberDBs {
		tuplet_numberDB := &tuplet_numberDBs[idx]
		_ = tuplet_numberDB
		var tuplet_numberAPI orm.Tuplet_numberAPI

		// insertion point for updating fields
		tuplet_numberAPI.ID = tuplet_numberDB.ID
		tuplet_numberDB.CopyBasicFieldsToTuplet_number_WOP(&tuplet_numberAPI.Tuplet_number_WOP)
		tuplet_numberAPI.Tuplet_numberPointersEncoding = tuplet_numberDB.Tuplet_numberPointersEncoding
		tuplet_numberAPIs = append(tuplet_numberAPIs, tuplet_numberAPI)
	}

	c.JSON(http.StatusOK, tuplet_numberAPIs)
}

// PostTuplet_number
//
// swagger:route POST /tuplet_numbers tuplet_numbers postTuplet_number
//
// Creates a tuplet_number
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTuplet_number(c *gin.Context) {

	mutexTuplet_number.Lock()
	defer mutexTuplet_number.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTuplet_numbers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_number.GetDB()

	// Validate input
	var input orm.Tuplet_numberAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tuplet_number
	tuplet_numberDB := orm.Tuplet_numberDB{}
	tuplet_numberDB.Tuplet_numberPointersEncoding = input.Tuplet_numberPointersEncoding
	tuplet_numberDB.CopyBasicFieldsFromTuplet_number_WOP(&input.Tuplet_number_WOP)

	query := db.Create(&tuplet_numberDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTuplet_number.CheckoutPhaseOneInstance(&tuplet_numberDB)
	tuplet_number := backRepo.BackRepoTuplet_number.Map_Tuplet_numberDBID_Tuplet_numberPtr[tuplet_numberDB.ID]

	if tuplet_number != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tuplet_number)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tuplet_numberDB)
}

// GetTuplet_number
//
// swagger:route GET /tuplet_numbers/{ID} tuplet_numbers getTuplet_number
//
// Gets the details for a tuplet_number.
//
// Responses:
// default: genericError
//
//	200: tuplet_numberDBResponse
func (controller *Controller) GetTuplet_number(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTuplet_number", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_number.GetDB()

	// Get tuplet_numberDB in DB
	var tuplet_numberDB orm.Tuplet_numberDB
	if err := db.First(&tuplet_numberDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tuplet_numberAPI orm.Tuplet_numberAPI
	tuplet_numberAPI.ID = tuplet_numberDB.ID
	tuplet_numberAPI.Tuplet_numberPointersEncoding = tuplet_numberDB.Tuplet_numberPointersEncoding
	tuplet_numberDB.CopyBasicFieldsToTuplet_number_WOP(&tuplet_numberAPI.Tuplet_number_WOP)

	c.JSON(http.StatusOK, tuplet_numberAPI)
}

// UpdateTuplet_number
//
// swagger:route PATCH /tuplet_numbers/{ID} tuplet_numbers updateTuplet_number
//
// # Update a tuplet_number
//
// Responses:
// default: genericError
//
//	200: tuplet_numberDBResponse
func (controller *Controller) UpdateTuplet_number(c *gin.Context) {

	mutexTuplet_number.Lock()
	defer mutexTuplet_number.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTuplet_number", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_number.GetDB()

	// Validate input
	var input orm.Tuplet_numberAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tuplet_numberDB orm.Tuplet_numberDB

	// fetch the tuplet_number
	query := db.First(&tuplet_numberDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tuplet_numberDB.CopyBasicFieldsFromTuplet_number_WOP(&input.Tuplet_number_WOP)
	tuplet_numberDB.Tuplet_numberPointersEncoding = input.Tuplet_numberPointersEncoding

	query = db.Model(&tuplet_numberDB).Updates(tuplet_numberDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_numberNew := new(models.Tuplet_number)
	tuplet_numberDB.CopyBasicFieldsToTuplet_number(tuplet_numberNew)

	// redeem pointers
	tuplet_numberDB.DecodePointers(backRepo, tuplet_numberNew)

	// get stage instance from DB instance, and call callback function
	tuplet_numberOld := backRepo.BackRepoTuplet_number.Map_Tuplet_numberDBID_Tuplet_numberPtr[tuplet_numberDB.ID]
	if tuplet_numberOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tuplet_numberOld, tuplet_numberNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tuplet_numberDB
	c.JSON(http.StatusOK, tuplet_numberDB)
}

// DeleteTuplet_number
//
// swagger:route DELETE /tuplet_numbers/{ID} tuplet_numbers deleteTuplet_number
//
// # Delete a tuplet_number
//
// default: genericError
//
//	200: tuplet_numberDBResponse
func (controller *Controller) DeleteTuplet_number(c *gin.Context) {

	mutexTuplet_number.Lock()
	defer mutexTuplet_number.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTuplet_number", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTuplet_number.GetDB()

	// Get model if exist
	var tuplet_numberDB orm.Tuplet_numberDB
	if err := db.First(&tuplet_numberDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tuplet_numberDB)

	// get an instance (not staged) from DB instance, and call callback function
	tuplet_numberDeleted := new(models.Tuplet_number)
	tuplet_numberDB.CopyBasicFieldsToTuplet_number(tuplet_numberDeleted)

	// get stage instance from DB instance, and call callback function
	tuplet_numberStaged := backRepo.BackRepoTuplet_number.Map_Tuplet_numberDBID_Tuplet_numberPtr[tuplet_numberDB.ID]
	if tuplet_numberStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tuplet_numberStaged, tuplet_numberDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
