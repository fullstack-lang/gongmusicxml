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
var __System_layout__dummysDeclaration__ models.System_layout
var __System_layout_time__dummyDeclaration time.Duration

var mutexSystem_layout sync.Mutex

// An System_layoutID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSystem_layout updateSystem_layout deleteSystem_layout
type System_layoutID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// System_layoutInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSystem_layout updateSystem_layout
type System_layoutInput struct {
	// The System_layout to submit or modify
	// in: body
	System_layout *orm.System_layoutAPI
}

// GetSystem_layouts
//
// swagger:route GET /system_layouts system_layouts getSystem_layouts
//
// # Get all system_layouts
//
// Responses:
// default: genericError
//
//	200: system_layoutDBResponse
func (controller *Controller) GetSystem_layouts(c *gin.Context) {

	// source slice
	var system_layoutDBs []orm.System_layoutDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSystem_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_layout.GetDB()

	query := db.Find(&system_layoutDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	system_layoutAPIs := make([]orm.System_layoutAPI, 0)

	// for each system_layout, update fields from the database nullable fields
	for idx := range system_layoutDBs {
		system_layoutDB := &system_layoutDBs[idx]
		_ = system_layoutDB
		var system_layoutAPI orm.System_layoutAPI

		// insertion point for updating fields
		system_layoutAPI.ID = system_layoutDB.ID
		system_layoutDB.CopyBasicFieldsToSystem_layout_WOP(&system_layoutAPI.System_layout_WOP)
		system_layoutAPI.System_layoutPointersEncoding = system_layoutDB.System_layoutPointersEncoding
		system_layoutAPIs = append(system_layoutAPIs, system_layoutAPI)
	}

	c.JSON(http.StatusOK, system_layoutAPIs)
}

// PostSystem_layout
//
// swagger:route POST /system_layouts system_layouts postSystem_layout
//
// Creates a system_layout
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSystem_layout(c *gin.Context) {

	mutexSystem_layout.Lock()
	defer mutexSystem_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSystem_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_layout.GetDB()

	// Validate input
	var input orm.System_layoutAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create system_layout
	system_layoutDB := orm.System_layoutDB{}
	system_layoutDB.System_layoutPointersEncoding = input.System_layoutPointersEncoding
	system_layoutDB.CopyBasicFieldsFromSystem_layout_WOP(&input.System_layout_WOP)

	query := db.Create(&system_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSystem_layout.CheckoutPhaseOneInstance(&system_layoutDB)
	system_layout := backRepo.BackRepoSystem_layout.Map_System_layoutDBID_System_layoutPtr[system_layoutDB.ID]

	if system_layout != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), system_layout)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, system_layoutDB)
}

// GetSystem_layout
//
// swagger:route GET /system_layouts/{ID} system_layouts getSystem_layout
//
// Gets the details for a system_layout.
//
// Responses:
// default: genericError
//
//	200: system_layoutDBResponse
func (controller *Controller) GetSystem_layout(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSystem_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_layout.GetDB()

	// Get system_layoutDB in DB
	var system_layoutDB orm.System_layoutDB
	if err := db.First(&system_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var system_layoutAPI orm.System_layoutAPI
	system_layoutAPI.ID = system_layoutDB.ID
	system_layoutAPI.System_layoutPointersEncoding = system_layoutDB.System_layoutPointersEncoding
	system_layoutDB.CopyBasicFieldsToSystem_layout_WOP(&system_layoutAPI.System_layout_WOP)

	c.JSON(http.StatusOK, system_layoutAPI)
}

// UpdateSystem_layout
//
// swagger:route PATCH /system_layouts/{ID} system_layouts updateSystem_layout
//
// # Update a system_layout
//
// Responses:
// default: genericError
//
//	200: system_layoutDBResponse
func (controller *Controller) UpdateSystem_layout(c *gin.Context) {

	mutexSystem_layout.Lock()
	defer mutexSystem_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSystem_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_layout.GetDB()

	// Validate input
	var input orm.System_layoutAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var system_layoutDB orm.System_layoutDB

	// fetch the system_layout
	query := db.First(&system_layoutDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	system_layoutDB.CopyBasicFieldsFromSystem_layout_WOP(&input.System_layout_WOP)
	system_layoutDB.System_layoutPointersEncoding = input.System_layoutPointersEncoding

	query = db.Model(&system_layoutDB).Updates(system_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	system_layoutNew := new(models.System_layout)
	system_layoutDB.CopyBasicFieldsToSystem_layout(system_layoutNew)

	// redeem pointers
	system_layoutDB.DecodePointers(backRepo, system_layoutNew)

	// get stage instance from DB instance, and call callback function
	system_layoutOld := backRepo.BackRepoSystem_layout.Map_System_layoutDBID_System_layoutPtr[system_layoutDB.ID]
	if system_layoutOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), system_layoutOld, system_layoutNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the system_layoutDB
	c.JSON(http.StatusOK, system_layoutDB)
}

// DeleteSystem_layout
//
// swagger:route DELETE /system_layouts/{ID} system_layouts deleteSystem_layout
//
// # Delete a system_layout
//
// default: genericError
//
//	200: system_layoutDBResponse
func (controller *Controller) DeleteSystem_layout(c *gin.Context) {

	mutexSystem_layout.Lock()
	defer mutexSystem_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSystem_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSystem_layout.GetDB()

	// Get model if exist
	var system_layoutDB orm.System_layoutDB
	if err := db.First(&system_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&system_layoutDB)

	// get an instance (not staged) from DB instance, and call callback function
	system_layoutDeleted := new(models.System_layout)
	system_layoutDB.CopyBasicFieldsToSystem_layout(system_layoutDeleted)

	// get stage instance from DB instance, and call callback function
	system_layoutStaged := backRepo.BackRepoSystem_layout.Map_System_layoutDBID_System_layoutPtr[system_layoutDB.ID]
	if system_layoutStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), system_layoutStaged, system_layoutDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
