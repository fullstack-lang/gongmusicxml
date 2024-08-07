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
var __Metronome_tuplet__dummysDeclaration__ models.Metronome_tuplet
var __Metronome_tuplet_time__dummyDeclaration time.Duration

var mutexMetronome_tuplet sync.Mutex

// An Metronome_tupletID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMetronome_tuplet updateMetronome_tuplet deleteMetronome_tuplet
type Metronome_tupletID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Metronome_tupletInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMetronome_tuplet updateMetronome_tuplet
type Metronome_tupletInput struct {
	// The Metronome_tuplet to submit or modify
	// in: body
	Metronome_tuplet *orm.Metronome_tupletAPI
}

// GetMetronome_tuplets
//
// swagger:route GET /metronome_tuplets metronome_tuplets getMetronome_tuplets
//
// # Get all metronome_tuplets
//
// Responses:
// default: genericError
//
//	200: metronome_tupletDBResponse
func (controller *Controller) GetMetronome_tuplets(c *gin.Context) {

	// source slice
	var metronome_tupletDBs []orm.Metronome_tupletDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_tuplets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tuplet.GetDB()

	query := db.Find(&metronome_tupletDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metronome_tupletAPIs := make([]orm.Metronome_tupletAPI, 0)

	// for each metronome_tuplet, update fields from the database nullable fields
	for idx := range metronome_tupletDBs {
		metronome_tupletDB := &metronome_tupletDBs[idx]
		_ = metronome_tupletDB
		var metronome_tupletAPI orm.Metronome_tupletAPI

		// insertion point for updating fields
		metronome_tupletAPI.ID = metronome_tupletDB.ID
		metronome_tupletDB.CopyBasicFieldsToMetronome_tuplet_WOP(&metronome_tupletAPI.Metronome_tuplet_WOP)
		metronome_tupletAPI.Metronome_tupletPointersEncoding = metronome_tupletDB.Metronome_tupletPointersEncoding
		metronome_tupletAPIs = append(metronome_tupletAPIs, metronome_tupletAPI)
	}

	c.JSON(http.StatusOK, metronome_tupletAPIs)
}

// PostMetronome_tuplet
//
// swagger:route POST /metronome_tuplets metronome_tuplets postMetronome_tuplet
//
// Creates a metronome_tuplet
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMetronome_tuplet(c *gin.Context) {

	mutexMetronome_tuplet.Lock()
	defer mutexMetronome_tuplet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetronome_tuplets", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tuplet.GetDB()

	// Validate input
	var input orm.Metronome_tupletAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create metronome_tuplet
	metronome_tupletDB := orm.Metronome_tupletDB{}
	metronome_tupletDB.Metronome_tupletPointersEncoding = input.Metronome_tupletPointersEncoding
	metronome_tupletDB.CopyBasicFieldsFromMetronome_tuplet_WOP(&input.Metronome_tuplet_WOP)

	query := db.Create(&metronome_tupletDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMetronome_tuplet.CheckoutPhaseOneInstance(&metronome_tupletDB)
	metronome_tuplet := backRepo.BackRepoMetronome_tuplet.Map_Metronome_tupletDBID_Metronome_tupletPtr[metronome_tupletDB.ID]

	if metronome_tuplet != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), metronome_tuplet)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metronome_tupletDB)
}

// GetMetronome_tuplet
//
// swagger:route GET /metronome_tuplets/{ID} metronome_tuplets getMetronome_tuplet
//
// Gets the details for a metronome_tuplet.
//
// Responses:
// default: genericError
//
//	200: metronome_tupletDBResponse
func (controller *Controller) GetMetronome_tuplet(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_tuplet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tuplet.GetDB()

	// Get metronome_tupletDB in DB
	var metronome_tupletDB orm.Metronome_tupletDB
	if err := db.First(&metronome_tupletDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metronome_tupletAPI orm.Metronome_tupletAPI
	metronome_tupletAPI.ID = metronome_tupletDB.ID
	metronome_tupletAPI.Metronome_tupletPointersEncoding = metronome_tupletDB.Metronome_tupletPointersEncoding
	metronome_tupletDB.CopyBasicFieldsToMetronome_tuplet_WOP(&metronome_tupletAPI.Metronome_tuplet_WOP)

	c.JSON(http.StatusOK, metronome_tupletAPI)
}

// UpdateMetronome_tuplet
//
// swagger:route PATCH /metronome_tuplets/{ID} metronome_tuplets updateMetronome_tuplet
//
// # Update a metronome_tuplet
//
// Responses:
// default: genericError
//
//	200: metronome_tupletDBResponse
func (controller *Controller) UpdateMetronome_tuplet(c *gin.Context) {

	mutexMetronome_tuplet.Lock()
	defer mutexMetronome_tuplet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMetronome_tuplet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tuplet.GetDB()

	// Validate input
	var input orm.Metronome_tupletAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metronome_tupletDB orm.Metronome_tupletDB

	// fetch the metronome_tuplet
	query := db.First(&metronome_tupletDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metronome_tupletDB.CopyBasicFieldsFromMetronome_tuplet_WOP(&input.Metronome_tuplet_WOP)
	metronome_tupletDB.Metronome_tupletPointersEncoding = input.Metronome_tupletPointersEncoding

	query = db.Model(&metronome_tupletDB).Updates(metronome_tupletDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metronome_tupletNew := new(models.Metronome_tuplet)
	metronome_tupletDB.CopyBasicFieldsToMetronome_tuplet(metronome_tupletNew)

	// redeem pointers
	metronome_tupletDB.DecodePointers(backRepo, metronome_tupletNew)

	// get stage instance from DB instance, and call callback function
	metronome_tupletOld := backRepo.BackRepoMetronome_tuplet.Map_Metronome_tupletDBID_Metronome_tupletPtr[metronome_tupletDB.ID]
	if metronome_tupletOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metronome_tupletOld, metronome_tupletNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metronome_tupletDB
	c.JSON(http.StatusOK, metronome_tupletDB)
}

// DeleteMetronome_tuplet
//
// swagger:route DELETE /metronome_tuplets/{ID} metronome_tuplets deleteMetronome_tuplet
//
// # Delete a metronome_tuplet
//
// default: genericError
//
//	200: metronome_tupletDBResponse
func (controller *Controller) DeleteMetronome_tuplet(c *gin.Context) {

	mutexMetronome_tuplet.Lock()
	defer mutexMetronome_tuplet.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMetronome_tuplet", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tuplet.GetDB()

	// Get model if exist
	var metronome_tupletDB orm.Metronome_tupletDB
	if err := db.First(&metronome_tupletDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&metronome_tupletDB)

	// get an instance (not staged) from DB instance, and call callback function
	metronome_tupletDeleted := new(models.Metronome_tuplet)
	metronome_tupletDB.CopyBasicFieldsToMetronome_tuplet(metronome_tupletDeleted)

	// get stage instance from DB instance, and call callback function
	metronome_tupletStaged := backRepo.BackRepoMetronome_tuplet.Map_Metronome_tupletDBID_Metronome_tupletPtr[metronome_tupletDB.ID]
	if metronome_tupletStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metronome_tupletStaged, metronome_tupletDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
