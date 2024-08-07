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
var __Supports__dummysDeclaration__ models.Supports
var __Supports_time__dummyDeclaration time.Duration

var mutexSupports sync.Mutex

// An SupportsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSupports updateSupports deleteSupports
type SupportsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SupportsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSupports updateSupports
type SupportsInput struct {
	// The Supports to submit or modify
	// in: body
	Supports *orm.SupportsAPI
}

// GetSupportss
//
// swagger:route GET /supportss supportss getSupportss
//
// # Get all supportss
//
// Responses:
// default: genericError
//
//	200: supportsDBResponse
func (controller *Controller) GetSupportss(c *gin.Context) {

	// source slice
	var supportsDBs []orm.SupportsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSupportss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSupports.GetDB()

	query := db.Find(&supportsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	supportsAPIs := make([]orm.SupportsAPI, 0)

	// for each supports, update fields from the database nullable fields
	for idx := range supportsDBs {
		supportsDB := &supportsDBs[idx]
		_ = supportsDB
		var supportsAPI orm.SupportsAPI

		// insertion point for updating fields
		supportsAPI.ID = supportsDB.ID
		supportsDB.CopyBasicFieldsToSupports_WOP(&supportsAPI.Supports_WOP)
		supportsAPI.SupportsPointersEncoding = supportsDB.SupportsPointersEncoding
		supportsAPIs = append(supportsAPIs, supportsAPI)
	}

	c.JSON(http.StatusOK, supportsAPIs)
}

// PostSupports
//
// swagger:route POST /supportss supportss postSupports
//
// Creates a supports
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSupports(c *gin.Context) {

	mutexSupports.Lock()
	defer mutexSupports.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSupportss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSupports.GetDB()

	// Validate input
	var input orm.SupportsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create supports
	supportsDB := orm.SupportsDB{}
	supportsDB.SupportsPointersEncoding = input.SupportsPointersEncoding
	supportsDB.CopyBasicFieldsFromSupports_WOP(&input.Supports_WOP)

	query := db.Create(&supportsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSupports.CheckoutPhaseOneInstance(&supportsDB)
	supports := backRepo.BackRepoSupports.Map_SupportsDBID_SupportsPtr[supportsDB.ID]

	if supports != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), supports)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, supportsDB)
}

// GetSupports
//
// swagger:route GET /supportss/{ID} supportss getSupports
//
// Gets the details for a supports.
//
// Responses:
// default: genericError
//
//	200: supportsDBResponse
func (controller *Controller) GetSupports(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSupports", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSupports.GetDB()

	// Get supportsDB in DB
	var supportsDB orm.SupportsDB
	if err := db.First(&supportsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var supportsAPI orm.SupportsAPI
	supportsAPI.ID = supportsDB.ID
	supportsAPI.SupportsPointersEncoding = supportsDB.SupportsPointersEncoding
	supportsDB.CopyBasicFieldsToSupports_WOP(&supportsAPI.Supports_WOP)

	c.JSON(http.StatusOK, supportsAPI)
}

// UpdateSupports
//
// swagger:route PATCH /supportss/{ID} supportss updateSupports
//
// # Update a supports
//
// Responses:
// default: genericError
//
//	200: supportsDBResponse
func (controller *Controller) UpdateSupports(c *gin.Context) {

	mutexSupports.Lock()
	defer mutexSupports.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSupports", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSupports.GetDB()

	// Validate input
	var input orm.SupportsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var supportsDB orm.SupportsDB

	// fetch the supports
	query := db.First(&supportsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	supportsDB.CopyBasicFieldsFromSupports_WOP(&input.Supports_WOP)
	supportsDB.SupportsPointersEncoding = input.SupportsPointersEncoding

	query = db.Model(&supportsDB).Updates(supportsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	supportsNew := new(models.Supports)
	supportsDB.CopyBasicFieldsToSupports(supportsNew)

	// redeem pointers
	supportsDB.DecodePointers(backRepo, supportsNew)

	// get stage instance from DB instance, and call callback function
	supportsOld := backRepo.BackRepoSupports.Map_SupportsDBID_SupportsPtr[supportsDB.ID]
	if supportsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), supportsOld, supportsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the supportsDB
	c.JSON(http.StatusOK, supportsDB)
}

// DeleteSupports
//
// swagger:route DELETE /supportss/{ID} supportss deleteSupports
//
// # Delete a supports
//
// default: genericError
//
//	200: supportsDBResponse
func (controller *Controller) DeleteSupports(c *gin.Context) {

	mutexSupports.Lock()
	defer mutexSupports.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSupports", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSupports.GetDB()

	// Get model if exist
	var supportsDB orm.SupportsDB
	if err := db.First(&supportsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&supportsDB)

	// get an instance (not staged) from DB instance, and call callback function
	supportsDeleted := new(models.Supports)
	supportsDB.CopyBasicFieldsToSupports(supportsDeleted)

	// get stage instance from DB instance, and call callback function
	supportsStaged := backRepo.BackRepoSupports.Map_SupportsDBID_SupportsPtr[supportsDB.ID]
	if supportsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), supportsStaged, supportsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
