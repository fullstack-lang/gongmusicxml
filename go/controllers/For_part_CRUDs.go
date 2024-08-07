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
var __For_part__dummysDeclaration__ models.For_part
var __For_part_time__dummyDeclaration time.Duration

var mutexFor_part sync.Mutex

// An For_partID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFor_part updateFor_part deleteFor_part
type For_partID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// For_partInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFor_part updateFor_part
type For_partInput struct {
	// The For_part to submit or modify
	// in: body
	For_part *orm.For_partAPI
}

// GetFor_parts
//
// swagger:route GET /for_parts for_parts getFor_parts
//
// # Get all for_parts
//
// Responses:
// default: genericError
//
//	200: for_partDBResponse
func (controller *Controller) GetFor_parts(c *gin.Context) {

	// source slice
	var for_partDBs []orm.For_partDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFor_parts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFor_part.GetDB()

	query := db.Find(&for_partDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	for_partAPIs := make([]orm.For_partAPI, 0)

	// for each for_part, update fields from the database nullable fields
	for idx := range for_partDBs {
		for_partDB := &for_partDBs[idx]
		_ = for_partDB
		var for_partAPI orm.For_partAPI

		// insertion point for updating fields
		for_partAPI.ID = for_partDB.ID
		for_partDB.CopyBasicFieldsToFor_part_WOP(&for_partAPI.For_part_WOP)
		for_partAPI.For_partPointersEncoding = for_partDB.For_partPointersEncoding
		for_partAPIs = append(for_partAPIs, for_partAPI)
	}

	c.JSON(http.StatusOK, for_partAPIs)
}

// PostFor_part
//
// swagger:route POST /for_parts for_parts postFor_part
//
// Creates a for_part
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFor_part(c *gin.Context) {

	mutexFor_part.Lock()
	defer mutexFor_part.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFor_parts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFor_part.GetDB()

	// Validate input
	var input orm.For_partAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create for_part
	for_partDB := orm.For_partDB{}
	for_partDB.For_partPointersEncoding = input.For_partPointersEncoding
	for_partDB.CopyBasicFieldsFromFor_part_WOP(&input.For_part_WOP)

	query := db.Create(&for_partDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFor_part.CheckoutPhaseOneInstance(&for_partDB)
	for_part := backRepo.BackRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID]

	if for_part != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), for_part)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, for_partDB)
}

// GetFor_part
//
// swagger:route GET /for_parts/{ID} for_parts getFor_part
//
// Gets the details for a for_part.
//
// Responses:
// default: genericError
//
//	200: for_partDBResponse
func (controller *Controller) GetFor_part(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFor_part", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFor_part.GetDB()

	// Get for_partDB in DB
	var for_partDB orm.For_partDB
	if err := db.First(&for_partDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var for_partAPI orm.For_partAPI
	for_partAPI.ID = for_partDB.ID
	for_partAPI.For_partPointersEncoding = for_partDB.For_partPointersEncoding
	for_partDB.CopyBasicFieldsToFor_part_WOP(&for_partAPI.For_part_WOP)

	c.JSON(http.StatusOK, for_partAPI)
}

// UpdateFor_part
//
// swagger:route PATCH /for_parts/{ID} for_parts updateFor_part
//
// # Update a for_part
//
// Responses:
// default: genericError
//
//	200: for_partDBResponse
func (controller *Controller) UpdateFor_part(c *gin.Context) {

	mutexFor_part.Lock()
	defer mutexFor_part.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFor_part", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFor_part.GetDB()

	// Validate input
	var input orm.For_partAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var for_partDB orm.For_partDB

	// fetch the for_part
	query := db.First(&for_partDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	for_partDB.CopyBasicFieldsFromFor_part_WOP(&input.For_part_WOP)
	for_partDB.For_partPointersEncoding = input.For_partPointersEncoding

	query = db.Model(&for_partDB).Updates(for_partDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	for_partNew := new(models.For_part)
	for_partDB.CopyBasicFieldsToFor_part(for_partNew)

	// redeem pointers
	for_partDB.DecodePointers(backRepo, for_partNew)

	// get stage instance from DB instance, and call callback function
	for_partOld := backRepo.BackRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID]
	if for_partOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), for_partOld, for_partNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the for_partDB
	c.JSON(http.StatusOK, for_partDB)
}

// DeleteFor_part
//
// swagger:route DELETE /for_parts/{ID} for_parts deleteFor_part
//
// # Delete a for_part
//
// default: genericError
//
//	200: for_partDBResponse
func (controller *Controller) DeleteFor_part(c *gin.Context) {

	mutexFor_part.Lock()
	defer mutexFor_part.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFor_part", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFor_part.GetDB()

	// Get model if exist
	var for_partDB orm.For_partDB
	if err := db.First(&for_partDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&for_partDB)

	// get an instance (not staged) from DB instance, and call callback function
	for_partDeleted := new(models.For_part)
	for_partDB.CopyBasicFieldsToFor_part(for_partDeleted)

	// get stage instance from DB instance, and call callback function
	for_partStaged := backRepo.BackRepoFor_part.Map_For_partDBID_For_partPtr[for_partDB.ID]
	if for_partStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), for_partStaged, for_partDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
