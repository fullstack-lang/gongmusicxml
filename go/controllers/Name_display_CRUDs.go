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
var __Name_display__dummysDeclaration__ models.Name_display
var __Name_display_time__dummyDeclaration time.Duration

var mutexName_display sync.Mutex

// An Name_displayID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getName_display updateName_display deleteName_display
type Name_displayID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Name_displayInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postName_display updateName_display
type Name_displayInput struct {
	// The Name_display to submit or modify
	// in: body
	Name_display *orm.Name_displayAPI
}

// GetName_displays
//
// swagger:route GET /name_displays name_displays getName_displays
//
// # Get all name_displays
//
// Responses:
// default: genericError
//
//	200: name_displayDBResponse
func (controller *Controller) GetName_displays(c *gin.Context) {

	// source slice
	var name_displayDBs []orm.Name_displayDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetName_displays", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoName_display.GetDB()

	query := db.Find(&name_displayDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	name_displayAPIs := make([]orm.Name_displayAPI, 0)

	// for each name_display, update fields from the database nullable fields
	for idx := range name_displayDBs {
		name_displayDB := &name_displayDBs[idx]
		_ = name_displayDB
		var name_displayAPI orm.Name_displayAPI

		// insertion point for updating fields
		name_displayAPI.ID = name_displayDB.ID
		name_displayDB.CopyBasicFieldsToName_display_WOP(&name_displayAPI.Name_display_WOP)
		name_displayAPI.Name_displayPointersEncoding = name_displayDB.Name_displayPointersEncoding
		name_displayAPIs = append(name_displayAPIs, name_displayAPI)
	}

	c.JSON(http.StatusOK, name_displayAPIs)
}

// PostName_display
//
// swagger:route POST /name_displays name_displays postName_display
//
// Creates a name_display
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostName_display(c *gin.Context) {

	mutexName_display.Lock()
	defer mutexName_display.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostName_displays", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoName_display.GetDB()

	// Validate input
	var input orm.Name_displayAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create name_display
	name_displayDB := orm.Name_displayDB{}
	name_displayDB.Name_displayPointersEncoding = input.Name_displayPointersEncoding
	name_displayDB.CopyBasicFieldsFromName_display_WOP(&input.Name_display_WOP)

	query := db.Create(&name_displayDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoName_display.CheckoutPhaseOneInstance(&name_displayDB)
	name_display := backRepo.BackRepoName_display.Map_Name_displayDBID_Name_displayPtr[name_displayDB.ID]

	if name_display != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), name_display)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, name_displayDB)
}

// GetName_display
//
// swagger:route GET /name_displays/{ID} name_displays getName_display
//
// Gets the details for a name_display.
//
// Responses:
// default: genericError
//
//	200: name_displayDBResponse
func (controller *Controller) GetName_display(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetName_display", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoName_display.GetDB()

	// Get name_displayDB in DB
	var name_displayDB orm.Name_displayDB
	if err := db.First(&name_displayDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var name_displayAPI orm.Name_displayAPI
	name_displayAPI.ID = name_displayDB.ID
	name_displayAPI.Name_displayPointersEncoding = name_displayDB.Name_displayPointersEncoding
	name_displayDB.CopyBasicFieldsToName_display_WOP(&name_displayAPI.Name_display_WOP)

	c.JSON(http.StatusOK, name_displayAPI)
}

// UpdateName_display
//
// swagger:route PATCH /name_displays/{ID} name_displays updateName_display
//
// # Update a name_display
//
// Responses:
// default: genericError
//
//	200: name_displayDBResponse
func (controller *Controller) UpdateName_display(c *gin.Context) {

	mutexName_display.Lock()
	defer mutexName_display.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateName_display", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoName_display.GetDB()

	// Validate input
	var input orm.Name_displayAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var name_displayDB orm.Name_displayDB

	// fetch the name_display
	query := db.First(&name_displayDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	name_displayDB.CopyBasicFieldsFromName_display_WOP(&input.Name_display_WOP)
	name_displayDB.Name_displayPointersEncoding = input.Name_displayPointersEncoding

	query = db.Model(&name_displayDB).Updates(name_displayDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	name_displayNew := new(models.Name_display)
	name_displayDB.CopyBasicFieldsToName_display(name_displayNew)

	// redeem pointers
	name_displayDB.DecodePointers(backRepo, name_displayNew)

	// get stage instance from DB instance, and call callback function
	name_displayOld := backRepo.BackRepoName_display.Map_Name_displayDBID_Name_displayPtr[name_displayDB.ID]
	if name_displayOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), name_displayOld, name_displayNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the name_displayDB
	c.JSON(http.StatusOK, name_displayDB)
}

// DeleteName_display
//
// swagger:route DELETE /name_displays/{ID} name_displays deleteName_display
//
// # Delete a name_display
//
// default: genericError
//
//	200: name_displayDBResponse
func (controller *Controller) DeleteName_display(c *gin.Context) {

	mutexName_display.Lock()
	defer mutexName_display.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteName_display", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoName_display.GetDB()

	// Get model if exist
	var name_displayDB orm.Name_displayDB
	if err := db.First(&name_displayDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&name_displayDB)

	// get an instance (not staged) from DB instance, and call callback function
	name_displayDeleted := new(models.Name_display)
	name_displayDB.CopyBasicFieldsToName_display(name_displayDeleted)

	// get stage instance from DB instance, and call callback function
	name_displayStaged := backRepo.BackRepoName_display.Map_Name_displayDBID_Name_displayPtr[name_displayDB.ID]
	if name_displayStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), name_displayStaged, name_displayDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
