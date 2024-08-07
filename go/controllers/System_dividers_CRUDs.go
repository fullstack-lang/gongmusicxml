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
var __System_dividers__dummysDeclaration__ models.System_dividers
var __System_dividers_time__dummyDeclaration time.Duration

var mutexSystem_dividers sync.Mutex

// An System_dividersID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSystem_dividers updateSystem_dividers deleteSystem_dividers
type System_dividersID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// System_dividersInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSystem_dividers updateSystem_dividers
type System_dividersInput struct {
	// The System_dividers to submit or modify
	// in: body
	System_dividers *orm.System_dividersAPI
}

// GetSystem_dividerss
//
// swagger:route GET /system_dividerss system_dividerss getSystem_dividerss
//
// # Get all system_dividerss
//
// Responses:
// default: genericError
//
//	200: system_dividersDBResponse
func (controller *Controller) GetSystem_dividerss(c *gin.Context) {

	// source slice
	var system_dividersDBs []orm.System_dividersDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSystem_dividerss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_dividers.GetDB()

	query := db.Find(&system_dividersDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	system_dividersAPIs := make([]orm.System_dividersAPI, 0)

	// for each system_dividers, update fields from the database nullable fields
	for idx := range system_dividersDBs {
		system_dividersDB := &system_dividersDBs[idx]
		_ = system_dividersDB
		var system_dividersAPI orm.System_dividersAPI

		// insertion point for updating fields
		system_dividersAPI.ID = system_dividersDB.ID
		system_dividersDB.CopyBasicFieldsToSystem_dividers_WOP(&system_dividersAPI.System_dividers_WOP)
		system_dividersAPI.System_dividersPointersEncoding = system_dividersDB.System_dividersPointersEncoding
		system_dividersAPIs = append(system_dividersAPIs, system_dividersAPI)
	}

	c.JSON(http.StatusOK, system_dividersAPIs)
}

// PostSystem_dividers
//
// swagger:route POST /system_dividerss system_dividerss postSystem_dividers
//
// Creates a system_dividers
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSystem_dividers(c *gin.Context) {

	mutexSystem_dividers.Lock()
	defer mutexSystem_dividers.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSystem_dividerss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_dividers.GetDB()

	// Validate input
	var input orm.System_dividersAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create system_dividers
	system_dividersDB := orm.System_dividersDB{}
	system_dividersDB.System_dividersPointersEncoding = input.System_dividersPointersEncoding
	system_dividersDB.CopyBasicFieldsFromSystem_dividers_WOP(&input.System_dividers_WOP)

	query := db.Create(&system_dividersDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSystem_dividers.CheckoutPhaseOneInstance(&system_dividersDB)
	system_dividers := backRepo.BackRepoSystem_dividers.Map_System_dividersDBID_System_dividersPtr[system_dividersDB.ID]

	if system_dividers != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), system_dividers)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, system_dividersDB)
}

// GetSystem_dividers
//
// swagger:route GET /system_dividerss/{ID} system_dividerss getSystem_dividers
//
// Gets the details for a system_dividers.
//
// Responses:
// default: genericError
//
//	200: system_dividersDBResponse
func (controller *Controller) GetSystem_dividers(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSystem_dividers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_dividers.GetDB()

	// Get system_dividersDB in DB
	var system_dividersDB orm.System_dividersDB
	if err := db.First(&system_dividersDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var system_dividersAPI orm.System_dividersAPI
	system_dividersAPI.ID = system_dividersDB.ID
	system_dividersAPI.System_dividersPointersEncoding = system_dividersDB.System_dividersPointersEncoding
	system_dividersDB.CopyBasicFieldsToSystem_dividers_WOP(&system_dividersAPI.System_dividers_WOP)

	c.JSON(http.StatusOK, system_dividersAPI)
}

// UpdateSystem_dividers
//
// swagger:route PATCH /system_dividerss/{ID} system_dividerss updateSystem_dividers
//
// # Update a system_dividers
//
// Responses:
// default: genericError
//
//	200: system_dividersDBResponse
func (controller *Controller) UpdateSystem_dividers(c *gin.Context) {

	mutexSystem_dividers.Lock()
	defer mutexSystem_dividers.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSystem_dividers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_dividers.GetDB()

	// Validate input
	var input orm.System_dividersAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var system_dividersDB orm.System_dividersDB

	// fetch the system_dividers
	query := db.First(&system_dividersDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	system_dividersDB.CopyBasicFieldsFromSystem_dividers_WOP(&input.System_dividers_WOP)
	system_dividersDB.System_dividersPointersEncoding = input.System_dividersPointersEncoding

	query = db.Model(&system_dividersDB).Updates(system_dividersDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	system_dividersNew := new(models.System_dividers)
	system_dividersDB.CopyBasicFieldsToSystem_dividers(system_dividersNew)

	// redeem pointers
	system_dividersDB.DecodePointers(backRepo, system_dividersNew)

	// get stage instance from DB instance, and call callback function
	system_dividersOld := backRepo.BackRepoSystem_dividers.Map_System_dividersDBID_System_dividersPtr[system_dividersDB.ID]
	if system_dividersOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), system_dividersOld, system_dividersNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the system_dividersDB
	c.JSON(http.StatusOK, system_dividersDB)
}

// DeleteSystem_dividers
//
// swagger:route DELETE /system_dividerss/{ID} system_dividerss deleteSystem_dividers
//
// # Delete a system_dividers
//
// default: genericError
//
//	200: system_dividersDBResponse
func (controller *Controller) DeleteSystem_dividers(c *gin.Context) {

	mutexSystem_dividers.Lock()
	defer mutexSystem_dividers.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSystem_dividers", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_dividers.GetDB()

	// Get model if exist
	var system_dividersDB orm.System_dividersDB
	if err := db.First(&system_dividersDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&system_dividersDB)

	// get an instance (not staged) from DB instance, and call callback function
	system_dividersDeleted := new(models.System_dividers)
	system_dividersDB.CopyBasicFieldsToSystem_dividers(system_dividersDeleted)

	// get stage instance from DB instance, and call callback function
	system_dividersStaged := backRepo.BackRepoSystem_dividers.Map_System_dividersDBID_System_dividersPtr[system_dividersDB.ID]
	if system_dividersStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), system_dividersStaged, system_dividersDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
