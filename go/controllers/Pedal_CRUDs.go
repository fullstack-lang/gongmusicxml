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
var __Pedal__dummysDeclaration__ models.Pedal
var __Pedal_time__dummyDeclaration time.Duration

var mutexPedal sync.Mutex

// An PedalID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPedal updatePedal deletePedal
type PedalID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PedalInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPedal updatePedal
type PedalInput struct {
	// The Pedal to submit or modify
	// in: body
	Pedal *orm.PedalAPI
}

// GetPedals
//
// swagger:route GET /pedals pedals getPedals
//
// # Get all pedals
//
// Responses:
// default: genericError
//
//	200: pedalDBResponse
func (controller *Controller) GetPedals(c *gin.Context) {

	// source slice
	var pedalDBs []orm.PedalDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPedals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal.GetDB()

	query := db.Find(&pedalDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	pedalAPIs := make([]orm.PedalAPI, 0)

	// for each pedal, update fields from the database nullable fields
	for idx := range pedalDBs {
		pedalDB := &pedalDBs[idx]
		_ = pedalDB
		var pedalAPI orm.PedalAPI

		// insertion point for updating fields
		pedalAPI.ID = pedalDB.ID
		pedalDB.CopyBasicFieldsToPedal_WOP(&pedalAPI.Pedal_WOP)
		pedalAPI.PedalPointersEncoding = pedalDB.PedalPointersEncoding
		pedalAPIs = append(pedalAPIs, pedalAPI)
	}

	c.JSON(http.StatusOK, pedalAPIs)
}

// PostPedal
//
// swagger:route POST /pedals pedals postPedal
//
// Creates a pedal
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPedal(c *gin.Context) {

	mutexPedal.Lock()
	defer mutexPedal.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPedals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal.GetDB()

	// Validate input
	var input orm.PedalAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create pedal
	pedalDB := orm.PedalDB{}
	pedalDB.PedalPointersEncoding = input.PedalPointersEncoding
	pedalDB.CopyBasicFieldsFromPedal_WOP(&input.Pedal_WOP)

	query := db.Create(&pedalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPedal.CheckoutPhaseOneInstance(&pedalDB)
	pedal := backRepo.BackRepoPedal.Map_PedalDBID_PedalPtr[pedalDB.ID]

	if pedal != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), pedal)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, pedalDB)
}

// GetPedal
//
// swagger:route GET /pedals/{ID} pedals getPedal
//
// Gets the details for a pedal.
//
// Responses:
// default: genericError
//
//	200: pedalDBResponse
func (controller *Controller) GetPedal(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPedal", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal.GetDB()

	// Get pedalDB in DB
	var pedalDB orm.PedalDB
	if err := db.First(&pedalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var pedalAPI orm.PedalAPI
	pedalAPI.ID = pedalDB.ID
	pedalAPI.PedalPointersEncoding = pedalDB.PedalPointersEncoding
	pedalDB.CopyBasicFieldsToPedal_WOP(&pedalAPI.Pedal_WOP)

	c.JSON(http.StatusOK, pedalAPI)
}

// UpdatePedal
//
// swagger:route PATCH /pedals/{ID} pedals updatePedal
//
// # Update a pedal
//
// Responses:
// default: genericError
//
//	200: pedalDBResponse
func (controller *Controller) UpdatePedal(c *gin.Context) {

	mutexPedal.Lock()
	defer mutexPedal.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePedal", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal.GetDB()

	// Validate input
	var input orm.PedalAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var pedalDB orm.PedalDB

	// fetch the pedal
	query := db.First(&pedalDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	pedalDB.CopyBasicFieldsFromPedal_WOP(&input.Pedal_WOP)
	pedalDB.PedalPointersEncoding = input.PedalPointersEncoding

	query = db.Model(&pedalDB).Updates(pedalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	pedalNew := new(models.Pedal)
	pedalDB.CopyBasicFieldsToPedal(pedalNew)

	// redeem pointers
	pedalDB.DecodePointers(backRepo, pedalNew)

	// get stage instance from DB instance, and call callback function
	pedalOld := backRepo.BackRepoPedal.Map_PedalDBID_PedalPtr[pedalDB.ID]
	if pedalOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), pedalOld, pedalNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the pedalDB
	c.JSON(http.StatusOK, pedalDB)
}

// DeletePedal
//
// swagger:route DELETE /pedals/{ID} pedals deletePedal
//
// # Delete a pedal
//
// default: genericError
//
//	200: pedalDBResponse
func (controller *Controller) DeletePedal(c *gin.Context) {

	mutexPedal.Lock()
	defer mutexPedal.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePedal", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPedal.GetDB()

	// Get model if exist
	var pedalDB orm.PedalDB
	if err := db.First(&pedalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&pedalDB)

	// get an instance (not staged) from DB instance, and call callback function
	pedalDeleted := new(models.Pedal)
	pedalDB.CopyBasicFieldsToPedal(pedalDeleted)

	// get stage instance from DB instance, and call callback function
	pedalStaged := backRepo.BackRepoPedal.Map_PedalDBID_PedalPtr[pedalDB.ID]
	if pedalStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), pedalStaged, pedalDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
