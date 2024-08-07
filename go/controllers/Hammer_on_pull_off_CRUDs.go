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
var __Hammer_on_pull_off__dummysDeclaration__ models.Hammer_on_pull_off
var __Hammer_on_pull_off_time__dummyDeclaration time.Duration

var mutexHammer_on_pull_off sync.Mutex

// An Hammer_on_pull_offID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHammer_on_pull_off updateHammer_on_pull_off deleteHammer_on_pull_off
type Hammer_on_pull_offID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Hammer_on_pull_offInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHammer_on_pull_off updateHammer_on_pull_off
type Hammer_on_pull_offInput struct {
	// The Hammer_on_pull_off to submit or modify
	// in: body
	Hammer_on_pull_off *orm.Hammer_on_pull_offAPI
}

// GetHammer_on_pull_offs
//
// swagger:route GET /hammer_on_pull_offs hammer_on_pull_offs getHammer_on_pull_offs
//
// # Get all hammer_on_pull_offs
//
// Responses:
// default: genericError
//
//	200: hammer_on_pull_offDBResponse
func (controller *Controller) GetHammer_on_pull_offs(c *gin.Context) {

	// source slice
	var hammer_on_pull_offDBs []orm.Hammer_on_pull_offDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHammer_on_pull_offs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHammer_on_pull_off.GetDB()

	query := db.Find(&hammer_on_pull_offDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	hammer_on_pull_offAPIs := make([]orm.Hammer_on_pull_offAPI, 0)

	// for each hammer_on_pull_off, update fields from the database nullable fields
	for idx := range hammer_on_pull_offDBs {
		hammer_on_pull_offDB := &hammer_on_pull_offDBs[idx]
		_ = hammer_on_pull_offDB
		var hammer_on_pull_offAPI orm.Hammer_on_pull_offAPI

		// insertion point for updating fields
		hammer_on_pull_offAPI.ID = hammer_on_pull_offDB.ID
		hammer_on_pull_offDB.CopyBasicFieldsToHammer_on_pull_off_WOP(&hammer_on_pull_offAPI.Hammer_on_pull_off_WOP)
		hammer_on_pull_offAPI.Hammer_on_pull_offPointersEncoding = hammer_on_pull_offDB.Hammer_on_pull_offPointersEncoding
		hammer_on_pull_offAPIs = append(hammer_on_pull_offAPIs, hammer_on_pull_offAPI)
	}

	c.JSON(http.StatusOK, hammer_on_pull_offAPIs)
}

// PostHammer_on_pull_off
//
// swagger:route POST /hammer_on_pull_offs hammer_on_pull_offs postHammer_on_pull_off
//
// Creates a hammer_on_pull_off
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHammer_on_pull_off(c *gin.Context) {

	mutexHammer_on_pull_off.Lock()
	defer mutexHammer_on_pull_off.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHammer_on_pull_offs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHammer_on_pull_off.GetDB()

	// Validate input
	var input orm.Hammer_on_pull_offAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create hammer_on_pull_off
	hammer_on_pull_offDB := orm.Hammer_on_pull_offDB{}
	hammer_on_pull_offDB.Hammer_on_pull_offPointersEncoding = input.Hammer_on_pull_offPointersEncoding
	hammer_on_pull_offDB.CopyBasicFieldsFromHammer_on_pull_off_WOP(&input.Hammer_on_pull_off_WOP)

	query := db.Create(&hammer_on_pull_offDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHammer_on_pull_off.CheckoutPhaseOneInstance(&hammer_on_pull_offDB)
	hammer_on_pull_off := backRepo.BackRepoHammer_on_pull_off.Map_Hammer_on_pull_offDBID_Hammer_on_pull_offPtr[hammer_on_pull_offDB.ID]

	if hammer_on_pull_off != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), hammer_on_pull_off)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, hammer_on_pull_offDB)
}

// GetHammer_on_pull_off
//
// swagger:route GET /hammer_on_pull_offs/{ID} hammer_on_pull_offs getHammer_on_pull_off
//
// Gets the details for a hammer_on_pull_off.
//
// Responses:
// default: genericError
//
//	200: hammer_on_pull_offDBResponse
func (controller *Controller) GetHammer_on_pull_off(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHammer_on_pull_off", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHammer_on_pull_off.GetDB()

	// Get hammer_on_pull_offDB in DB
	var hammer_on_pull_offDB orm.Hammer_on_pull_offDB
	if err := db.First(&hammer_on_pull_offDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var hammer_on_pull_offAPI orm.Hammer_on_pull_offAPI
	hammer_on_pull_offAPI.ID = hammer_on_pull_offDB.ID
	hammer_on_pull_offAPI.Hammer_on_pull_offPointersEncoding = hammer_on_pull_offDB.Hammer_on_pull_offPointersEncoding
	hammer_on_pull_offDB.CopyBasicFieldsToHammer_on_pull_off_WOP(&hammer_on_pull_offAPI.Hammer_on_pull_off_WOP)

	c.JSON(http.StatusOK, hammer_on_pull_offAPI)
}

// UpdateHammer_on_pull_off
//
// swagger:route PATCH /hammer_on_pull_offs/{ID} hammer_on_pull_offs updateHammer_on_pull_off
//
// # Update a hammer_on_pull_off
//
// Responses:
// default: genericError
//
//	200: hammer_on_pull_offDBResponse
func (controller *Controller) UpdateHammer_on_pull_off(c *gin.Context) {

	mutexHammer_on_pull_off.Lock()
	defer mutexHammer_on_pull_off.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHammer_on_pull_off", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHammer_on_pull_off.GetDB()

	// Validate input
	var input orm.Hammer_on_pull_offAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var hammer_on_pull_offDB orm.Hammer_on_pull_offDB

	// fetch the hammer_on_pull_off
	query := db.First(&hammer_on_pull_offDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	hammer_on_pull_offDB.CopyBasicFieldsFromHammer_on_pull_off_WOP(&input.Hammer_on_pull_off_WOP)
	hammer_on_pull_offDB.Hammer_on_pull_offPointersEncoding = input.Hammer_on_pull_offPointersEncoding

	query = db.Model(&hammer_on_pull_offDB).Updates(hammer_on_pull_offDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	hammer_on_pull_offNew := new(models.Hammer_on_pull_off)
	hammer_on_pull_offDB.CopyBasicFieldsToHammer_on_pull_off(hammer_on_pull_offNew)

	// redeem pointers
	hammer_on_pull_offDB.DecodePointers(backRepo, hammer_on_pull_offNew)

	// get stage instance from DB instance, and call callback function
	hammer_on_pull_offOld := backRepo.BackRepoHammer_on_pull_off.Map_Hammer_on_pull_offDBID_Hammer_on_pull_offPtr[hammer_on_pull_offDB.ID]
	if hammer_on_pull_offOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), hammer_on_pull_offOld, hammer_on_pull_offNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the hammer_on_pull_offDB
	c.JSON(http.StatusOK, hammer_on_pull_offDB)
}

// DeleteHammer_on_pull_off
//
// swagger:route DELETE /hammer_on_pull_offs/{ID} hammer_on_pull_offs deleteHammer_on_pull_off
//
// # Delete a hammer_on_pull_off
//
// default: genericError
//
//	200: hammer_on_pull_offDBResponse
func (controller *Controller) DeleteHammer_on_pull_off(c *gin.Context) {

	mutexHammer_on_pull_off.Lock()
	defer mutexHammer_on_pull_off.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHammer_on_pull_off", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHammer_on_pull_off.GetDB()

	// Get model if exist
	var hammer_on_pull_offDB orm.Hammer_on_pull_offDB
	if err := db.First(&hammer_on_pull_offDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&hammer_on_pull_offDB)

	// get an instance (not staged) from DB instance, and call callback function
	hammer_on_pull_offDeleted := new(models.Hammer_on_pull_off)
	hammer_on_pull_offDB.CopyBasicFieldsToHammer_on_pull_off(hammer_on_pull_offDeleted)

	// get stage instance from DB instance, and call callback function
	hammer_on_pull_offStaged := backRepo.BackRepoHammer_on_pull_off.Map_Hammer_on_pull_offDBID_Hammer_on_pull_offPtr[hammer_on_pull_offDB.ID]
	if hammer_on_pull_offStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), hammer_on_pull_offStaged, hammer_on_pull_offDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
