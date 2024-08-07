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
var __Dashes__dummysDeclaration__ models.Dashes
var __Dashes_time__dummyDeclaration time.Duration

var mutexDashes sync.Mutex

// An DashesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDashes updateDashes deleteDashes
type DashesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DashesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDashes updateDashes
type DashesInput struct {
	// The Dashes to submit or modify
	// in: body
	Dashes *orm.DashesAPI
}

// GetDashess
//
// swagger:route GET /dashess dashess getDashess
//
// # Get all dashess
//
// Responses:
// default: genericError
//
//	200: dashesDBResponse
func (controller *Controller) GetDashess(c *gin.Context) {

	// source slice
	var dashesDBs []orm.DashesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDashess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDashes.GetDB()

	query := db.Find(&dashesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dashesAPIs := make([]orm.DashesAPI, 0)

	// for each dashes, update fields from the database nullable fields
	for idx := range dashesDBs {
		dashesDB := &dashesDBs[idx]
		_ = dashesDB
		var dashesAPI orm.DashesAPI

		// insertion point for updating fields
		dashesAPI.ID = dashesDB.ID
		dashesDB.CopyBasicFieldsToDashes_WOP(&dashesAPI.Dashes_WOP)
		dashesAPI.DashesPointersEncoding = dashesDB.DashesPointersEncoding
		dashesAPIs = append(dashesAPIs, dashesAPI)
	}

	c.JSON(http.StatusOK, dashesAPIs)
}

// PostDashes
//
// swagger:route POST /dashess dashess postDashes
//
// Creates a dashes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDashes(c *gin.Context) {

	mutexDashes.Lock()
	defer mutexDashes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDashess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDashes.GetDB()

	// Validate input
	var input orm.DashesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dashes
	dashesDB := orm.DashesDB{}
	dashesDB.DashesPointersEncoding = input.DashesPointersEncoding
	dashesDB.CopyBasicFieldsFromDashes_WOP(&input.Dashes_WOP)

	query := db.Create(&dashesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDashes.CheckoutPhaseOneInstance(&dashesDB)
	dashes := backRepo.BackRepoDashes.Map_DashesDBID_DashesPtr[dashesDB.ID]

	if dashes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), dashes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, dashesDB)
}

// GetDashes
//
// swagger:route GET /dashess/{ID} dashess getDashes
//
// Gets the details for a dashes.
//
// Responses:
// default: genericError
//
//	200: dashesDBResponse
func (controller *Controller) GetDashes(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDashes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDashes.GetDB()

	// Get dashesDB in DB
	var dashesDB orm.DashesDB
	if err := db.First(&dashesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var dashesAPI orm.DashesAPI
	dashesAPI.ID = dashesDB.ID
	dashesAPI.DashesPointersEncoding = dashesDB.DashesPointersEncoding
	dashesDB.CopyBasicFieldsToDashes_WOP(&dashesAPI.Dashes_WOP)

	c.JSON(http.StatusOK, dashesAPI)
}

// UpdateDashes
//
// swagger:route PATCH /dashess/{ID} dashess updateDashes
//
// # Update a dashes
//
// Responses:
// default: genericError
//
//	200: dashesDBResponse
func (controller *Controller) UpdateDashes(c *gin.Context) {

	mutexDashes.Lock()
	defer mutexDashes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDashes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDashes.GetDB()

	// Validate input
	var input orm.DashesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dashesDB orm.DashesDB

	// fetch the dashes
	query := db.First(&dashesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dashesDB.CopyBasicFieldsFromDashes_WOP(&input.Dashes_WOP)
	dashesDB.DashesPointersEncoding = input.DashesPointersEncoding

	query = db.Model(&dashesDB).Updates(dashesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dashesNew := new(models.Dashes)
	dashesDB.CopyBasicFieldsToDashes(dashesNew)

	// redeem pointers
	dashesDB.DecodePointers(backRepo, dashesNew)

	// get stage instance from DB instance, and call callback function
	dashesOld := backRepo.BackRepoDashes.Map_DashesDBID_DashesPtr[dashesDB.ID]
	if dashesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), dashesOld, dashesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dashesDB
	c.JSON(http.StatusOK, dashesDB)
}

// DeleteDashes
//
// swagger:route DELETE /dashess/{ID} dashess deleteDashes
//
// # Delete a dashes
//
// default: genericError
//
//	200: dashesDBResponse
func (controller *Controller) DeleteDashes(c *gin.Context) {

	mutexDashes.Lock()
	defer mutexDashes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDashes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDashes.GetDB()

	// Get model if exist
	var dashesDB orm.DashesDB
	if err := db.First(&dashesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&dashesDB)

	// get an instance (not staged) from DB instance, and call callback function
	dashesDeleted := new(models.Dashes)
	dashesDB.CopyBasicFieldsToDashes(dashesDeleted)

	// get stage instance from DB instance, and call callback function
	dashesStaged := backRepo.BackRepoDashes.Map_DashesDBID_DashesPtr[dashesDB.ID]
	if dashesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), dashesStaged, dashesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
