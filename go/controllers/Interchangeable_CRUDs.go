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
var __Interchangeable__dummysDeclaration__ models.Interchangeable
var __Interchangeable_time__dummyDeclaration time.Duration

var mutexInterchangeable sync.Mutex

// An InterchangeableID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getInterchangeable updateInterchangeable deleteInterchangeable
type InterchangeableID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// InterchangeableInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postInterchangeable updateInterchangeable
type InterchangeableInput struct {
	// The Interchangeable to submit or modify
	// in: body
	Interchangeable *orm.InterchangeableAPI
}

// GetInterchangeables
//
// swagger:route GET /interchangeables interchangeables getInterchangeables
//
// # Get all interchangeables
//
// Responses:
// default: genericError
//
//	200: interchangeableDBResponse
func (controller *Controller) GetInterchangeables(c *gin.Context) {

	// source slice
	var interchangeableDBs []orm.InterchangeableDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInterchangeables", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInterchangeable.GetDB()

	query := db.Find(&interchangeableDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	interchangeableAPIs := make([]orm.InterchangeableAPI, 0)

	// for each interchangeable, update fields from the database nullable fields
	for idx := range interchangeableDBs {
		interchangeableDB := &interchangeableDBs[idx]
		_ = interchangeableDB
		var interchangeableAPI orm.InterchangeableAPI

		// insertion point for updating fields
		interchangeableAPI.ID = interchangeableDB.ID
		interchangeableDB.CopyBasicFieldsToInterchangeable_WOP(&interchangeableAPI.Interchangeable_WOP)
		interchangeableAPI.InterchangeablePointersEncoding = interchangeableDB.InterchangeablePointersEncoding
		interchangeableAPIs = append(interchangeableAPIs, interchangeableAPI)
	}

	c.JSON(http.StatusOK, interchangeableAPIs)
}

// PostInterchangeable
//
// swagger:route POST /interchangeables interchangeables postInterchangeable
//
// Creates a interchangeable
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostInterchangeable(c *gin.Context) {

	mutexInterchangeable.Lock()
	defer mutexInterchangeable.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostInterchangeables", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInterchangeable.GetDB()

	// Validate input
	var input orm.InterchangeableAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create interchangeable
	interchangeableDB := orm.InterchangeableDB{}
	interchangeableDB.InterchangeablePointersEncoding = input.InterchangeablePointersEncoding
	interchangeableDB.CopyBasicFieldsFromInterchangeable_WOP(&input.Interchangeable_WOP)

	query := db.Create(&interchangeableDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoInterchangeable.CheckoutPhaseOneInstance(&interchangeableDB)
	interchangeable := backRepo.BackRepoInterchangeable.Map_InterchangeableDBID_InterchangeablePtr[interchangeableDB.ID]

	if interchangeable != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), interchangeable)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, interchangeableDB)
}

// GetInterchangeable
//
// swagger:route GET /interchangeables/{ID} interchangeables getInterchangeable
//
// Gets the details for a interchangeable.
//
// Responses:
// default: genericError
//
//	200: interchangeableDBResponse
func (controller *Controller) GetInterchangeable(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetInterchangeable", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInterchangeable.GetDB()

	// Get interchangeableDB in DB
	var interchangeableDB orm.InterchangeableDB
	if err := db.First(&interchangeableDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var interchangeableAPI orm.InterchangeableAPI
	interchangeableAPI.ID = interchangeableDB.ID
	interchangeableAPI.InterchangeablePointersEncoding = interchangeableDB.InterchangeablePointersEncoding
	interchangeableDB.CopyBasicFieldsToInterchangeable_WOP(&interchangeableAPI.Interchangeable_WOP)

	c.JSON(http.StatusOK, interchangeableAPI)
}

// UpdateInterchangeable
//
// swagger:route PATCH /interchangeables/{ID} interchangeables updateInterchangeable
//
// # Update a interchangeable
//
// Responses:
// default: genericError
//
//	200: interchangeableDBResponse
func (controller *Controller) UpdateInterchangeable(c *gin.Context) {

	mutexInterchangeable.Lock()
	defer mutexInterchangeable.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateInterchangeable", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInterchangeable.GetDB()

	// Validate input
	var input orm.InterchangeableAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var interchangeableDB orm.InterchangeableDB

	// fetch the interchangeable
	query := db.First(&interchangeableDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	interchangeableDB.CopyBasicFieldsFromInterchangeable_WOP(&input.Interchangeable_WOP)
	interchangeableDB.InterchangeablePointersEncoding = input.InterchangeablePointersEncoding

	query = db.Model(&interchangeableDB).Updates(interchangeableDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	interchangeableNew := new(models.Interchangeable)
	interchangeableDB.CopyBasicFieldsToInterchangeable(interchangeableNew)

	// redeem pointers
	interchangeableDB.DecodePointers(backRepo, interchangeableNew)

	// get stage instance from DB instance, and call callback function
	interchangeableOld := backRepo.BackRepoInterchangeable.Map_InterchangeableDBID_InterchangeablePtr[interchangeableDB.ID]
	if interchangeableOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), interchangeableOld, interchangeableNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the interchangeableDB
	c.JSON(http.StatusOK, interchangeableDB)
}

// DeleteInterchangeable
//
// swagger:route DELETE /interchangeables/{ID} interchangeables deleteInterchangeable
//
// # Delete a interchangeable
//
// default: genericError
//
//	200: interchangeableDBResponse
func (controller *Controller) DeleteInterchangeable(c *gin.Context) {

	mutexInterchangeable.Lock()
	defer mutexInterchangeable.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteInterchangeable", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoInterchangeable.GetDB()

	// Get model if exist
	var interchangeableDB orm.InterchangeableDB
	if err := db.First(&interchangeableDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&interchangeableDB)

	// get an instance (not staged) from DB instance, and call callback function
	interchangeableDeleted := new(models.Interchangeable)
	interchangeableDB.CopyBasicFieldsToInterchangeable(interchangeableDeleted)

	// get stage instance from DB instance, and call callback function
	interchangeableStaged := backRepo.BackRepoInterchangeable.Map_InterchangeableDBID_InterchangeablePtr[interchangeableDB.ID]
	if interchangeableStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), interchangeableStaged, interchangeableDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
