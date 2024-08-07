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
var __Harp_pedals__dummysDeclaration__ models.Harp_pedals
var __Harp_pedals_time__dummyDeclaration time.Duration

var mutexHarp_pedals sync.Mutex

// An Harp_pedalsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHarp_pedals updateHarp_pedals deleteHarp_pedals
type Harp_pedalsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Harp_pedalsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHarp_pedals updateHarp_pedals
type Harp_pedalsInput struct {
	// The Harp_pedals to submit or modify
	// in: body
	Harp_pedals *orm.Harp_pedalsAPI
}

// GetHarp_pedalss
//
// swagger:route GET /harp_pedalss harp_pedalss getHarp_pedalss
//
// # Get all harp_pedalss
//
// Responses:
// default: genericError
//
//	200: harp_pedalsDBResponse
func (controller *Controller) GetHarp_pedalss(c *gin.Context) {

	// source slice
	var harp_pedalsDBs []orm.Harp_pedalsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarp_pedalss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarp_pedals.GetDB()

	query := db.Find(&harp_pedalsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	harp_pedalsAPIs := make([]orm.Harp_pedalsAPI, 0)

	// for each harp_pedals, update fields from the database nullable fields
	for idx := range harp_pedalsDBs {
		harp_pedalsDB := &harp_pedalsDBs[idx]
		_ = harp_pedalsDB
		var harp_pedalsAPI orm.Harp_pedalsAPI

		// insertion point for updating fields
		harp_pedalsAPI.ID = harp_pedalsDB.ID
		harp_pedalsDB.CopyBasicFieldsToHarp_pedals_WOP(&harp_pedalsAPI.Harp_pedals_WOP)
		harp_pedalsAPI.Harp_pedalsPointersEncoding = harp_pedalsDB.Harp_pedalsPointersEncoding
		harp_pedalsAPIs = append(harp_pedalsAPIs, harp_pedalsAPI)
	}

	c.JSON(http.StatusOK, harp_pedalsAPIs)
}

// PostHarp_pedals
//
// swagger:route POST /harp_pedalss harp_pedalss postHarp_pedals
//
// Creates a harp_pedals
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHarp_pedals(c *gin.Context) {

	mutexHarp_pedals.Lock()
	defer mutexHarp_pedals.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHarp_pedalss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarp_pedals.GetDB()

	// Validate input
	var input orm.Harp_pedalsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create harp_pedals
	harp_pedalsDB := orm.Harp_pedalsDB{}
	harp_pedalsDB.Harp_pedalsPointersEncoding = input.Harp_pedalsPointersEncoding
	harp_pedalsDB.CopyBasicFieldsFromHarp_pedals_WOP(&input.Harp_pedals_WOP)

	query := db.Create(&harp_pedalsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHarp_pedals.CheckoutPhaseOneInstance(&harp_pedalsDB)
	harp_pedals := backRepo.BackRepoHarp_pedals.Map_Harp_pedalsDBID_Harp_pedalsPtr[harp_pedalsDB.ID]

	if harp_pedals != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), harp_pedals)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, harp_pedalsDB)
}

// GetHarp_pedals
//
// swagger:route GET /harp_pedalss/{ID} harp_pedalss getHarp_pedals
//
// Gets the details for a harp_pedals.
//
// Responses:
// default: genericError
//
//	200: harp_pedalsDBResponse
func (controller *Controller) GetHarp_pedals(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarp_pedals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarp_pedals.GetDB()

	// Get harp_pedalsDB in DB
	var harp_pedalsDB orm.Harp_pedalsDB
	if err := db.First(&harp_pedalsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var harp_pedalsAPI orm.Harp_pedalsAPI
	harp_pedalsAPI.ID = harp_pedalsDB.ID
	harp_pedalsAPI.Harp_pedalsPointersEncoding = harp_pedalsDB.Harp_pedalsPointersEncoding
	harp_pedalsDB.CopyBasicFieldsToHarp_pedals_WOP(&harp_pedalsAPI.Harp_pedals_WOP)

	c.JSON(http.StatusOK, harp_pedalsAPI)
}

// UpdateHarp_pedals
//
// swagger:route PATCH /harp_pedalss/{ID} harp_pedalss updateHarp_pedals
//
// # Update a harp_pedals
//
// Responses:
// default: genericError
//
//	200: harp_pedalsDBResponse
func (controller *Controller) UpdateHarp_pedals(c *gin.Context) {

	mutexHarp_pedals.Lock()
	defer mutexHarp_pedals.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHarp_pedals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarp_pedals.GetDB()

	// Validate input
	var input orm.Harp_pedalsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var harp_pedalsDB orm.Harp_pedalsDB

	// fetch the harp_pedals
	query := db.First(&harp_pedalsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	harp_pedalsDB.CopyBasicFieldsFromHarp_pedals_WOP(&input.Harp_pedals_WOP)
	harp_pedalsDB.Harp_pedalsPointersEncoding = input.Harp_pedalsPointersEncoding

	query = db.Model(&harp_pedalsDB).Updates(harp_pedalsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	harp_pedalsNew := new(models.Harp_pedals)
	harp_pedalsDB.CopyBasicFieldsToHarp_pedals(harp_pedalsNew)

	// redeem pointers
	harp_pedalsDB.DecodePointers(backRepo, harp_pedalsNew)

	// get stage instance from DB instance, and call callback function
	harp_pedalsOld := backRepo.BackRepoHarp_pedals.Map_Harp_pedalsDBID_Harp_pedalsPtr[harp_pedalsDB.ID]
	if harp_pedalsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), harp_pedalsOld, harp_pedalsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the harp_pedalsDB
	c.JSON(http.StatusOK, harp_pedalsDB)
}

// DeleteHarp_pedals
//
// swagger:route DELETE /harp_pedalss/{ID} harp_pedalss deleteHarp_pedals
//
// # Delete a harp_pedals
//
// default: genericError
//
//	200: harp_pedalsDBResponse
func (controller *Controller) DeleteHarp_pedals(c *gin.Context) {

	mutexHarp_pedals.Lock()
	defer mutexHarp_pedals.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHarp_pedals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarp_pedals.GetDB()

	// Get model if exist
	var harp_pedalsDB orm.Harp_pedalsDB
	if err := db.First(&harp_pedalsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&harp_pedalsDB)

	// get an instance (not staged) from DB instance, and call callback function
	harp_pedalsDeleted := new(models.Harp_pedals)
	harp_pedalsDB.CopyBasicFieldsToHarp_pedals(harp_pedalsDeleted)

	// get stage instance from DB instance, and call callback function
	harp_pedalsStaged := backRepo.BackRepoHarp_pedals.Map_Harp_pedalsDBID_Harp_pedalsPtr[harp_pedalsDB.ID]
	if harp_pedalsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), harp_pedalsStaged, harp_pedalsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
