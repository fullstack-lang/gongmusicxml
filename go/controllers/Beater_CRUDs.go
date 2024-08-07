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
var __Beater__dummysDeclaration__ models.Beater
var __Beater_time__dummyDeclaration time.Duration

var mutexBeater sync.Mutex

// An BeaterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBeater updateBeater deleteBeater
type BeaterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BeaterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBeater updateBeater
type BeaterInput struct {
	// The Beater to submit or modify
	// in: body
	Beater *orm.BeaterAPI
}

// GetBeaters
//
// swagger:route GET /beaters beaters getBeaters
//
// # Get all beaters
//
// Responses:
// default: genericError
//
//	200: beaterDBResponse
func (controller *Controller) GetBeaters(c *gin.Context) {

	// source slice
	var beaterDBs []orm.BeaterDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeaters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeater.GetDB()

	query := db.Find(&beaterDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	beaterAPIs := make([]orm.BeaterAPI, 0)

	// for each beater, update fields from the database nullable fields
	for idx := range beaterDBs {
		beaterDB := &beaterDBs[idx]
		_ = beaterDB
		var beaterAPI orm.BeaterAPI

		// insertion point for updating fields
		beaterAPI.ID = beaterDB.ID
		beaterDB.CopyBasicFieldsToBeater_WOP(&beaterAPI.Beater_WOP)
		beaterAPI.BeaterPointersEncoding = beaterDB.BeaterPointersEncoding
		beaterAPIs = append(beaterAPIs, beaterAPI)
	}

	c.JSON(http.StatusOK, beaterAPIs)
}

// PostBeater
//
// swagger:route POST /beaters beaters postBeater
//
// Creates a beater
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBeater(c *gin.Context) {

	mutexBeater.Lock()
	defer mutexBeater.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBeaters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeater.GetDB()

	// Validate input
	var input orm.BeaterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create beater
	beaterDB := orm.BeaterDB{}
	beaterDB.BeaterPointersEncoding = input.BeaterPointersEncoding
	beaterDB.CopyBasicFieldsFromBeater_WOP(&input.Beater_WOP)

	query := db.Create(&beaterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBeater.CheckoutPhaseOneInstance(&beaterDB)
	beater := backRepo.BackRepoBeater.Map_BeaterDBID_BeaterPtr[beaterDB.ID]

	if beater != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), beater)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, beaterDB)
}

// GetBeater
//
// swagger:route GET /beaters/{ID} beaters getBeater
//
// Gets the details for a beater.
//
// Responses:
// default: genericError
//
//	200: beaterDBResponse
func (controller *Controller) GetBeater(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBeater", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeater.GetDB()

	// Get beaterDB in DB
	var beaterDB orm.BeaterDB
	if err := db.First(&beaterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var beaterAPI orm.BeaterAPI
	beaterAPI.ID = beaterDB.ID
	beaterAPI.BeaterPointersEncoding = beaterDB.BeaterPointersEncoding
	beaterDB.CopyBasicFieldsToBeater_WOP(&beaterAPI.Beater_WOP)

	c.JSON(http.StatusOK, beaterAPI)
}

// UpdateBeater
//
// swagger:route PATCH /beaters/{ID} beaters updateBeater
//
// # Update a beater
//
// Responses:
// default: genericError
//
//	200: beaterDBResponse
func (controller *Controller) UpdateBeater(c *gin.Context) {

	mutexBeater.Lock()
	defer mutexBeater.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBeater", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeater.GetDB()

	// Validate input
	var input orm.BeaterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var beaterDB orm.BeaterDB

	// fetch the beater
	query := db.First(&beaterDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	beaterDB.CopyBasicFieldsFromBeater_WOP(&input.Beater_WOP)
	beaterDB.BeaterPointersEncoding = input.BeaterPointersEncoding

	query = db.Model(&beaterDB).Updates(beaterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	beaterNew := new(models.Beater)
	beaterDB.CopyBasicFieldsToBeater(beaterNew)

	// redeem pointers
	beaterDB.DecodePointers(backRepo, beaterNew)

	// get stage instance from DB instance, and call callback function
	beaterOld := backRepo.BackRepoBeater.Map_BeaterDBID_BeaterPtr[beaterDB.ID]
	if beaterOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), beaterOld, beaterNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the beaterDB
	c.JSON(http.StatusOK, beaterDB)
}

// DeleteBeater
//
// swagger:route DELETE /beaters/{ID} beaters deleteBeater
//
// # Delete a beater
//
// default: genericError
//
//	200: beaterDBResponse
func (controller *Controller) DeleteBeater(c *gin.Context) {

	mutexBeater.Lock()
	defer mutexBeater.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBeater", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBeater.GetDB()

	// Get model if exist
	var beaterDB orm.BeaterDB
	if err := db.First(&beaterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&beaterDB)

	// get an instance (not staged) from DB instance, and call callback function
	beaterDeleted := new(models.Beater)
	beaterDB.CopyBasicFieldsToBeater(beaterDeleted)

	// get stage instance from DB instance, and call callback function
	beaterStaged := backRepo.BackRepoBeater.Map_BeaterDBID_BeaterPtr[beaterDB.ID]
	if beaterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), beaterStaged, beaterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
