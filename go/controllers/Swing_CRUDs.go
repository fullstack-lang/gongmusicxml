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
var __Swing__dummysDeclaration__ models.Swing
var __Swing_time__dummyDeclaration time.Duration

var mutexSwing sync.Mutex

// An SwingID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSwing updateSwing deleteSwing
type SwingID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SwingInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSwing updateSwing
type SwingInput struct {
	// The Swing to submit or modify
	// in: body
	Swing *orm.SwingAPI
}

// GetSwings
//
// swagger:route GET /swings swings getSwings
//
// # Get all swings
//
// Responses:
// default: genericError
//
//	200: swingDBResponse
func (controller *Controller) GetSwings(c *gin.Context) {

	// source slice
	var swingDBs []orm.SwingDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSwings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSwing.GetDB()

	query := db.Find(&swingDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	swingAPIs := make([]orm.SwingAPI, 0)

	// for each swing, update fields from the database nullable fields
	for idx := range swingDBs {
		swingDB := &swingDBs[idx]
		_ = swingDB
		var swingAPI orm.SwingAPI

		// insertion point for updating fields
		swingAPI.ID = swingDB.ID
		swingDB.CopyBasicFieldsToSwing_WOP(&swingAPI.Swing_WOP)
		swingAPI.SwingPointersEncoding = swingDB.SwingPointersEncoding
		swingAPIs = append(swingAPIs, swingAPI)
	}

	c.JSON(http.StatusOK, swingAPIs)
}

// PostSwing
//
// swagger:route POST /swings swings postSwing
//
// Creates a swing
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSwing(c *gin.Context) {

	mutexSwing.Lock()
	defer mutexSwing.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSwings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSwing.GetDB()

	// Validate input
	var input orm.SwingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create swing
	swingDB := orm.SwingDB{}
	swingDB.SwingPointersEncoding = input.SwingPointersEncoding
	swingDB.CopyBasicFieldsFromSwing_WOP(&input.Swing_WOP)

	query := db.Create(&swingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSwing.CheckoutPhaseOneInstance(&swingDB)
	swing := backRepo.BackRepoSwing.Map_SwingDBID_SwingPtr[swingDB.ID]

	if swing != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), swing)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, swingDB)
}

// GetSwing
//
// swagger:route GET /swings/{ID} swings getSwing
//
// Gets the details for a swing.
//
// Responses:
// default: genericError
//
//	200: swingDBResponse
func (controller *Controller) GetSwing(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSwing", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSwing.GetDB()

	// Get swingDB in DB
	var swingDB orm.SwingDB
	if err := db.First(&swingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var swingAPI orm.SwingAPI
	swingAPI.ID = swingDB.ID
	swingAPI.SwingPointersEncoding = swingDB.SwingPointersEncoding
	swingDB.CopyBasicFieldsToSwing_WOP(&swingAPI.Swing_WOP)

	c.JSON(http.StatusOK, swingAPI)
}

// UpdateSwing
//
// swagger:route PATCH /swings/{ID} swings updateSwing
//
// # Update a swing
//
// Responses:
// default: genericError
//
//	200: swingDBResponse
func (controller *Controller) UpdateSwing(c *gin.Context) {

	mutexSwing.Lock()
	defer mutexSwing.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSwing", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSwing.GetDB()

	// Validate input
	var input orm.SwingAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var swingDB orm.SwingDB

	// fetch the swing
	query := db.First(&swingDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	swingDB.CopyBasicFieldsFromSwing_WOP(&input.Swing_WOP)
	swingDB.SwingPointersEncoding = input.SwingPointersEncoding

	query = db.Model(&swingDB).Updates(swingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	swingNew := new(models.Swing)
	swingDB.CopyBasicFieldsToSwing(swingNew)

	// redeem pointers
	swingDB.DecodePointers(backRepo, swingNew)

	// get stage instance from DB instance, and call callback function
	swingOld := backRepo.BackRepoSwing.Map_SwingDBID_SwingPtr[swingDB.ID]
	if swingOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), swingOld, swingNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the swingDB
	c.JSON(http.StatusOK, swingDB)
}

// DeleteSwing
//
// swagger:route DELETE /swings/{ID} swings deleteSwing
//
// # Delete a swing
//
// default: genericError
//
//	200: swingDBResponse
func (controller *Controller) DeleteSwing(c *gin.Context) {

	mutexSwing.Lock()
	defer mutexSwing.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSwing", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSwing.GetDB()

	// Get model if exist
	var swingDB orm.SwingDB
	if err := db.First(&swingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&swingDB)

	// get an instance (not staged) from DB instance, and call callback function
	swingDeleted := new(models.Swing)
	swingDB.CopyBasicFieldsToSwing(swingDeleted)

	// get stage instance from DB instance, and call callback function
	swingStaged := backRepo.BackRepoSwing.Map_SwingDBID_SwingPtr[swingDB.ID]
	if swingStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), swingStaged, swingDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
