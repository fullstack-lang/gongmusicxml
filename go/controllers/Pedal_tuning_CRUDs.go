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
var __Pedal_tuning__dummysDeclaration__ models.Pedal_tuning
var __Pedal_tuning_time__dummyDeclaration time.Duration

var mutexPedal_tuning sync.Mutex

// An Pedal_tuningID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPedal_tuning updatePedal_tuning deletePedal_tuning
type Pedal_tuningID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Pedal_tuningInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPedal_tuning updatePedal_tuning
type Pedal_tuningInput struct {
	// The Pedal_tuning to submit or modify
	// in: body
	Pedal_tuning *orm.Pedal_tuningAPI
}

// GetPedal_tunings
//
// swagger:route GET /pedal_tunings pedal_tunings getPedal_tunings
//
// # Get all pedal_tunings
//
// Responses:
// default: genericError
//
//	200: pedal_tuningDBResponse
func (controller *Controller) GetPedal_tunings(c *gin.Context) {

	// source slice
	var pedal_tuningDBs []orm.Pedal_tuningDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPedal_tunings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal_tuning.GetDB()

	query := db.Find(&pedal_tuningDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pedal_tuningAPIs := make([]orm.Pedal_tuningAPI, 0)

	// for each pedal_tuning, update fields from the database nullable fields
	for idx := range pedal_tuningDBs {
		pedal_tuningDB := &pedal_tuningDBs[idx]
		_ = pedal_tuningDB
		var pedal_tuningAPI orm.Pedal_tuningAPI

		// insertion point for updating fields
		pedal_tuningAPI.ID = pedal_tuningDB.ID
		pedal_tuningDB.CopyBasicFieldsToPedal_tuning_WOP(&pedal_tuningAPI.Pedal_tuning_WOP)
		pedal_tuningAPI.Pedal_tuningPointersEncoding = pedal_tuningDB.Pedal_tuningPointersEncoding
		pedal_tuningAPIs = append(pedal_tuningAPIs, pedal_tuningAPI)
	}

	c.JSON(http.StatusOK, pedal_tuningAPIs)
}

// PostPedal_tuning
//
// swagger:route POST /pedal_tunings pedal_tunings postPedal_tuning
//
// Creates a pedal_tuning
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPedal_tuning(c *gin.Context) {

	mutexPedal_tuning.Lock()
	defer mutexPedal_tuning.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPedal_tunings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal_tuning.GetDB()

	// Validate input
	var input orm.Pedal_tuningAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pedal_tuning
	pedal_tuningDB := orm.Pedal_tuningDB{}
	pedal_tuningDB.Pedal_tuningPointersEncoding = input.Pedal_tuningPointersEncoding
	pedal_tuningDB.CopyBasicFieldsFromPedal_tuning_WOP(&input.Pedal_tuning_WOP)

	query := db.Create(&pedal_tuningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPedal_tuning.CheckoutPhaseOneInstance(&pedal_tuningDB)
	pedal_tuning := backRepo.BackRepoPedal_tuning.Map_Pedal_tuningDBID_Pedal_tuningPtr[pedal_tuningDB.ID]

	if pedal_tuning != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pedal_tuning)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pedal_tuningDB)
}

// GetPedal_tuning
//
// swagger:route GET /pedal_tunings/{ID} pedal_tunings getPedal_tuning
//
// Gets the details for a pedal_tuning.
//
// Responses:
// default: genericError
//
//	200: pedal_tuningDBResponse
func (controller *Controller) GetPedal_tuning(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPedal_tuning", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal_tuning.GetDB()

	// Get pedal_tuningDB in DB
	var pedal_tuningDB orm.Pedal_tuningDB
	if err := db.First(&pedal_tuningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pedal_tuningAPI orm.Pedal_tuningAPI
	pedal_tuningAPI.ID = pedal_tuningDB.ID
	pedal_tuningAPI.Pedal_tuningPointersEncoding = pedal_tuningDB.Pedal_tuningPointersEncoding
	pedal_tuningDB.CopyBasicFieldsToPedal_tuning_WOP(&pedal_tuningAPI.Pedal_tuning_WOP)

	c.JSON(http.StatusOK, pedal_tuningAPI)
}

// UpdatePedal_tuning
//
// swagger:route PATCH /pedal_tunings/{ID} pedal_tunings updatePedal_tuning
//
// # Update a pedal_tuning
//
// Responses:
// default: genericError
//
//	200: pedal_tuningDBResponse
func (controller *Controller) UpdatePedal_tuning(c *gin.Context) {

	mutexPedal_tuning.Lock()
	defer mutexPedal_tuning.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePedal_tuning", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal_tuning.GetDB()

	// Validate input
	var input orm.Pedal_tuningAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pedal_tuningDB orm.Pedal_tuningDB

	// fetch the pedal_tuning
	query := db.First(&pedal_tuningDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pedal_tuningDB.CopyBasicFieldsFromPedal_tuning_WOP(&input.Pedal_tuning_WOP)
	pedal_tuningDB.Pedal_tuningPointersEncoding = input.Pedal_tuningPointersEncoding

	query = db.Model(&pedal_tuningDB).Updates(pedal_tuningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pedal_tuningNew := new(models.Pedal_tuning)
	pedal_tuningDB.CopyBasicFieldsToPedal_tuning(pedal_tuningNew)

	// redeem pointers
	pedal_tuningDB.DecodePointers(backRepo, pedal_tuningNew)

	// get stage instance from DB instance, and call callback function
	pedal_tuningOld := backRepo.BackRepoPedal_tuning.Map_Pedal_tuningDBID_Pedal_tuningPtr[pedal_tuningDB.ID]
	if pedal_tuningOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pedal_tuningOld, pedal_tuningNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pedal_tuningDB
	c.JSON(http.StatusOK, pedal_tuningDB)
}

// DeletePedal_tuning
//
// swagger:route DELETE /pedal_tunings/{ID} pedal_tunings deletePedal_tuning
//
// # Delete a pedal_tuning
//
// default: genericError
//
//	200: pedal_tuningDBResponse
func (controller *Controller) DeletePedal_tuning(c *gin.Context) {

	mutexPedal_tuning.Lock()
	defer mutexPedal_tuning.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePedal_tuning", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal_tuning.GetDB()

	// Get model if exist
	var pedal_tuningDB orm.Pedal_tuningDB
	if err := db.First(&pedal_tuningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&pedal_tuningDB)

	// get an instance (not staged) from DB instance, and call callback function
	pedal_tuningDeleted := new(models.Pedal_tuning)
	pedal_tuningDB.CopyBasicFieldsToPedal_tuning(pedal_tuningDeleted)

	// get stage instance from DB instance, and call callback function
	pedal_tuningStaged := backRepo.BackRepoPedal_tuning.Map_Pedal_tuningDBID_Pedal_tuningPtr[pedal_tuningDB.ID]
	if pedal_tuningStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pedal_tuningStaged, pedal_tuningDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
