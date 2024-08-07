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
var __System_margins__dummysDeclaration__ models.System_margins
var __System_margins_time__dummyDeclaration time.Duration

var mutexSystem_margins sync.Mutex

// An System_marginsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSystem_margins updateSystem_margins deleteSystem_margins
type System_marginsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// System_marginsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSystem_margins updateSystem_margins
type System_marginsInput struct {
	// The System_margins to submit or modify
	// in: body
	System_margins *orm.System_marginsAPI
}

// GetSystem_marginss
//
// swagger:route GET /system_marginss system_marginss getSystem_marginss
//
// # Get all system_marginss
//
// Responses:
// default: genericError
//
//	200: system_marginsDBResponse
func (controller *Controller) GetSystem_marginss(c *gin.Context) {

	// source slice
	var system_marginsDBs []orm.System_marginsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSystem_marginss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_margins.GetDB()

	query := db.Find(&system_marginsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	system_marginsAPIs := make([]orm.System_marginsAPI, 0)

	// for each system_margins, update fields from the database nullable fields
	for idx := range system_marginsDBs {
		system_marginsDB := &system_marginsDBs[idx]
		_ = system_marginsDB
		var system_marginsAPI orm.System_marginsAPI

		// insertion point for updating fields
		system_marginsAPI.ID = system_marginsDB.ID
		system_marginsDB.CopyBasicFieldsToSystem_margins_WOP(&system_marginsAPI.System_margins_WOP)
		system_marginsAPI.System_marginsPointersEncoding = system_marginsDB.System_marginsPointersEncoding
		system_marginsAPIs = append(system_marginsAPIs, system_marginsAPI)
	}

	c.JSON(http.StatusOK, system_marginsAPIs)
}

// PostSystem_margins
//
// swagger:route POST /system_marginss system_marginss postSystem_margins
//
// Creates a system_margins
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSystem_margins(c *gin.Context) {

	mutexSystem_margins.Lock()
	defer mutexSystem_margins.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSystem_marginss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_margins.GetDB()

	// Validate input
	var input orm.System_marginsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create system_margins
	system_marginsDB := orm.System_marginsDB{}
	system_marginsDB.System_marginsPointersEncoding = input.System_marginsPointersEncoding
	system_marginsDB.CopyBasicFieldsFromSystem_margins_WOP(&input.System_margins_WOP)

	query := db.Create(&system_marginsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSystem_margins.CheckoutPhaseOneInstance(&system_marginsDB)
	system_margins := backRepo.BackRepoSystem_margins.Map_System_marginsDBID_System_marginsPtr[system_marginsDB.ID]

	if system_margins != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), system_margins)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, system_marginsDB)
}

// GetSystem_margins
//
// swagger:route GET /system_marginss/{ID} system_marginss getSystem_margins
//
// Gets the details for a system_margins.
//
// Responses:
// default: genericError
//
//	200: system_marginsDBResponse
func (controller *Controller) GetSystem_margins(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSystem_margins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_margins.GetDB()

	// Get system_marginsDB in DB
	var system_marginsDB orm.System_marginsDB
	if err := db.First(&system_marginsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var system_marginsAPI orm.System_marginsAPI
	system_marginsAPI.ID = system_marginsDB.ID
	system_marginsAPI.System_marginsPointersEncoding = system_marginsDB.System_marginsPointersEncoding
	system_marginsDB.CopyBasicFieldsToSystem_margins_WOP(&system_marginsAPI.System_margins_WOP)

	c.JSON(http.StatusOK, system_marginsAPI)
}

// UpdateSystem_margins
//
// swagger:route PATCH /system_marginss/{ID} system_marginss updateSystem_margins
//
// # Update a system_margins
//
// Responses:
// default: genericError
//
//	200: system_marginsDBResponse
func (controller *Controller) UpdateSystem_margins(c *gin.Context) {

	mutexSystem_margins.Lock()
	defer mutexSystem_margins.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSystem_margins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_margins.GetDB()

	// Validate input
	var input orm.System_marginsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var system_marginsDB orm.System_marginsDB

	// fetch the system_margins
	query := db.First(&system_marginsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	system_marginsDB.CopyBasicFieldsFromSystem_margins_WOP(&input.System_margins_WOP)
	system_marginsDB.System_marginsPointersEncoding = input.System_marginsPointersEncoding

	query = db.Model(&system_marginsDB).Updates(system_marginsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	system_marginsNew := new(models.System_margins)
	system_marginsDB.CopyBasicFieldsToSystem_margins(system_marginsNew)

	// redeem pointers
	system_marginsDB.DecodePointers(backRepo, system_marginsNew)

	// get stage instance from DB instance, and call callback function
	system_marginsOld := backRepo.BackRepoSystem_margins.Map_System_marginsDBID_System_marginsPtr[system_marginsDB.ID]
	if system_marginsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), system_marginsOld, system_marginsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the system_marginsDB
	c.JSON(http.StatusOK, system_marginsDB)
}

// DeleteSystem_margins
//
// swagger:route DELETE /system_marginss/{ID} system_marginss deleteSystem_margins
//
// # Delete a system_margins
//
// default: genericError
//
//	200: system_marginsDBResponse
func (controller *Controller) DeleteSystem_margins(c *gin.Context) {

	mutexSystem_margins.Lock()
	defer mutexSystem_margins.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSystem_margins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_margins.GetDB()

	// Get model if exist
	var system_marginsDB orm.System_marginsDB
	if err := db.First(&system_marginsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&system_marginsDB)

	// get an instance (not staged) from DB instance, and call callback function
	system_marginsDeleted := new(models.System_margins)
	system_marginsDB.CopyBasicFieldsToSystem_margins(system_marginsDeleted)

	// get stage instance from DB instance, and call callback function
	system_marginsStaged := backRepo.BackRepoSystem_margins.Map_System_marginsDBID_System_marginsPtr[system_marginsDB.ID]
	if system_marginsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), system_marginsStaged, system_marginsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
