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
var __Empty__dummysDeclaration__ models.Empty
var __Empty_time__dummyDeclaration time.Duration

var mutexEmpty sync.Mutex

// An EmptyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty updateEmpty deleteEmpty
type EmptyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EmptyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty updateEmpty
type EmptyInput struct {
	// The Empty to submit or modify
	// in: body
	Empty *orm.EmptyAPI
}

// GetEmptys
//
// swagger:route GET /emptys emptys getEmptys
//
// # Get all emptys
//
// Responses:
// default: genericError
//
//	200: emptyDBResponse
func (controller *Controller) GetEmptys(c *gin.Context) {

	// source slice
	var emptyDBs []orm.EmptyDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmptys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty.GetDB()

	query := db.Find(&emptyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	emptyAPIs := make([]orm.EmptyAPI, 0)

	// for each empty, update fields from the database nullable fields
	for idx := range emptyDBs {
		emptyDB := &emptyDBs[idx]
		_ = emptyDB
		var emptyAPI orm.EmptyAPI

		// insertion point for updating fields
		emptyAPI.ID = emptyDB.ID
		emptyDB.CopyBasicFieldsToEmpty_WOP(&emptyAPI.Empty_WOP)
		emptyAPI.EmptyPointersEncoding = emptyDB.EmptyPointersEncoding
		emptyAPIs = append(emptyAPIs, emptyAPI)
	}

	c.JSON(http.StatusOK, emptyAPIs)
}

// PostEmpty
//
// swagger:route POST /emptys emptys postEmpty
//
// Creates a empty
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty(c *gin.Context) {

	mutexEmpty.Lock()
	defer mutexEmpty.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmptys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty.GetDB()

	// Validate input
	var input orm.EmptyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty
	emptyDB := orm.EmptyDB{}
	emptyDB.EmptyPointersEncoding = input.EmptyPointersEncoding
	emptyDB.CopyBasicFieldsFromEmpty_WOP(&input.Empty_WOP)

	query := db.Create(&emptyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty.CheckoutPhaseOneInstance(&emptyDB)
	empty := backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID]

	if empty != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, emptyDB)
}

// GetEmpty
//
// swagger:route GET /emptys/{ID} emptys getEmpty
//
// Gets the details for a empty.
//
// Responses:
// default: genericError
//
//	200: emptyDBResponse
func (controller *Controller) GetEmpty(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty.GetDB()

	// Get emptyDB in DB
	var emptyDB orm.EmptyDB
	if err := db.First(&emptyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var emptyAPI orm.EmptyAPI
	emptyAPI.ID = emptyDB.ID
	emptyAPI.EmptyPointersEncoding = emptyDB.EmptyPointersEncoding
	emptyDB.CopyBasicFieldsToEmpty_WOP(&emptyAPI.Empty_WOP)

	c.JSON(http.StatusOK, emptyAPI)
}

// UpdateEmpty
//
// swagger:route PATCH /emptys/{ID} emptys updateEmpty
//
// # Update a empty
//
// Responses:
// default: genericError
//
//	200: emptyDBResponse
func (controller *Controller) UpdateEmpty(c *gin.Context) {

	mutexEmpty.Lock()
	defer mutexEmpty.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty.GetDB()

	// Validate input
	var input orm.EmptyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var emptyDB orm.EmptyDB

	// fetch the empty
	query := db.First(&emptyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	emptyDB.CopyBasicFieldsFromEmpty_WOP(&input.Empty_WOP)
	emptyDB.EmptyPointersEncoding = input.EmptyPointersEncoding

	query = db.Model(&emptyDB).Updates(emptyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	emptyNew := new(models.Empty)
	emptyDB.CopyBasicFieldsToEmpty(emptyNew)

	// redeem pointers
	emptyDB.DecodePointers(backRepo, emptyNew)

	// get stage instance from DB instance, and call callback function
	emptyOld := backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID]
	if emptyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), emptyOld, emptyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the emptyDB
	c.JSON(http.StatusOK, emptyDB)
}

// DeleteEmpty
//
// swagger:route DELETE /emptys/{ID} emptys deleteEmpty
//
// # Delete a empty
//
// default: genericError
//
//	200: emptyDBResponse
func (controller *Controller) DeleteEmpty(c *gin.Context) {

	mutexEmpty.Lock()
	defer mutexEmpty.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty.GetDB()

	// Get model if exist
	var emptyDB orm.EmptyDB
	if err := db.First(&emptyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&emptyDB)

	// get an instance (not staged) from DB instance, and call callback function
	emptyDeleted := new(models.Empty)
	emptyDB.CopyBasicFieldsToEmpty(emptyDeleted)

	// get stage instance from DB instance, and call callback function
	emptyStaged := backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[emptyDB.ID]
	if emptyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), emptyStaged, emptyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
