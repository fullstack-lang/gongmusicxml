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
var __Other_listening__dummysDeclaration__ models.Other_listening
var __Other_listening_time__dummyDeclaration time.Duration

var mutexOther_listening sync.Mutex

// An Other_listeningID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOther_listening updateOther_listening deleteOther_listening
type Other_listeningID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Other_listeningInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOther_listening updateOther_listening
type Other_listeningInput struct {
	// The Other_listening to submit or modify
	// in: body
	Other_listening *orm.Other_listeningAPI
}

// GetOther_listenings
//
// swagger:route GET /other_listenings other_listenings getOther_listenings
//
// # Get all other_listenings
//
// Responses:
// default: genericError
//
//	200: other_listeningDBResponse
func (controller *Controller) GetOther_listenings(c *gin.Context) {

	// source slice
	var other_listeningDBs []orm.Other_listeningDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_listenings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_listening.GetDB()

	query := db.Find(&other_listeningDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	other_listeningAPIs := make([]orm.Other_listeningAPI, 0)

	// for each other_listening, update fields from the database nullable fields
	for idx := range other_listeningDBs {
		other_listeningDB := &other_listeningDBs[idx]
		_ = other_listeningDB
		var other_listeningAPI orm.Other_listeningAPI

		// insertion point for updating fields
		other_listeningAPI.ID = other_listeningDB.ID
		other_listeningDB.CopyBasicFieldsToOther_listening_WOP(&other_listeningAPI.Other_listening_WOP)
		other_listeningAPI.Other_listeningPointersEncoding = other_listeningDB.Other_listeningPointersEncoding
		other_listeningAPIs = append(other_listeningAPIs, other_listeningAPI)
	}

	c.JSON(http.StatusOK, other_listeningAPIs)
}

// PostOther_listening
//
// swagger:route POST /other_listenings other_listenings postOther_listening
//
// Creates a other_listening
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOther_listening(c *gin.Context) {

	mutexOther_listening.Lock()
	defer mutexOther_listening.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOther_listenings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_listening.GetDB()

	// Validate input
	var input orm.Other_listeningAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create other_listening
	other_listeningDB := orm.Other_listeningDB{}
	other_listeningDB.Other_listeningPointersEncoding = input.Other_listeningPointersEncoding
	other_listeningDB.CopyBasicFieldsFromOther_listening_WOP(&input.Other_listening_WOP)

	query := db.Create(&other_listeningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOther_listening.CheckoutPhaseOneInstance(&other_listeningDB)
	other_listening := backRepo.BackRepoOther_listening.Map_Other_listeningDBID_Other_listeningPtr[other_listeningDB.ID]

	if other_listening != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), other_listening)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, other_listeningDB)
}

// GetOther_listening
//
// swagger:route GET /other_listenings/{ID} other_listenings getOther_listening
//
// Gets the details for a other_listening.
//
// Responses:
// default: genericError
//
//	200: other_listeningDBResponse
func (controller *Controller) GetOther_listening(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_listening", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_listening.GetDB()

	// Get other_listeningDB in DB
	var other_listeningDB orm.Other_listeningDB
	if err := db.First(&other_listeningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var other_listeningAPI orm.Other_listeningAPI
	other_listeningAPI.ID = other_listeningDB.ID
	other_listeningAPI.Other_listeningPointersEncoding = other_listeningDB.Other_listeningPointersEncoding
	other_listeningDB.CopyBasicFieldsToOther_listening_WOP(&other_listeningAPI.Other_listening_WOP)

	c.JSON(http.StatusOK, other_listeningAPI)
}

// UpdateOther_listening
//
// swagger:route PATCH /other_listenings/{ID} other_listenings updateOther_listening
//
// # Update a other_listening
//
// Responses:
// default: genericError
//
//	200: other_listeningDBResponse
func (controller *Controller) UpdateOther_listening(c *gin.Context) {

	mutexOther_listening.Lock()
	defer mutexOther_listening.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOther_listening", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_listening.GetDB()

	// Validate input
	var input orm.Other_listeningAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var other_listeningDB orm.Other_listeningDB

	// fetch the other_listening
	query := db.First(&other_listeningDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	other_listeningDB.CopyBasicFieldsFromOther_listening_WOP(&input.Other_listening_WOP)
	other_listeningDB.Other_listeningPointersEncoding = input.Other_listeningPointersEncoding

	query = db.Model(&other_listeningDB).Updates(other_listeningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	other_listeningNew := new(models.Other_listening)
	other_listeningDB.CopyBasicFieldsToOther_listening(other_listeningNew)

	// redeem pointers
	other_listeningDB.DecodePointers(backRepo, other_listeningNew)

	// get stage instance from DB instance, and call callback function
	other_listeningOld := backRepo.BackRepoOther_listening.Map_Other_listeningDBID_Other_listeningPtr[other_listeningDB.ID]
	if other_listeningOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), other_listeningOld, other_listeningNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the other_listeningDB
	c.JSON(http.StatusOK, other_listeningDB)
}

// DeleteOther_listening
//
// swagger:route DELETE /other_listenings/{ID} other_listenings deleteOther_listening
//
// # Delete a other_listening
//
// default: genericError
//
//	200: other_listeningDBResponse
func (controller *Controller) DeleteOther_listening(c *gin.Context) {

	mutexOther_listening.Lock()
	defer mutexOther_listening.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOther_listening", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_listening.GetDB()

	// Get model if exist
	var other_listeningDB orm.Other_listeningDB
	if err := db.First(&other_listeningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&other_listeningDB)

	// get an instance (not staged) from DB instance, and call callback function
	other_listeningDeleted := new(models.Other_listening)
	other_listeningDB.CopyBasicFieldsToOther_listening(other_listeningDeleted)

	// get stage instance from DB instance, and call callback function
	other_listeningStaged := backRepo.BackRepoOther_listening.Map_Other_listeningDBID_Other_listeningPtr[other_listeningDB.ID]
	if other_listeningStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), other_listeningStaged, other_listeningDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
