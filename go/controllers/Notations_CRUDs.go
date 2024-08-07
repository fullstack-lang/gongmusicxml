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
var __Notations__dummysDeclaration__ models.Notations
var __Notations_time__dummyDeclaration time.Duration

var mutexNotations sync.Mutex

// An NotationsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNotations updateNotations deleteNotations
type NotationsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NotationsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNotations updateNotations
type NotationsInput struct {
	// The Notations to submit or modify
	// in: body
	Notations *orm.NotationsAPI
}

// GetNotationss
//
// swagger:route GET /notationss notationss getNotationss
//
// # Get all notationss
//
// Responses:
// default: genericError
//
//	200: notationsDBResponse
func (controller *Controller) GetNotationss(c *gin.Context) {

	// source slice
	var notationsDBs []orm.NotationsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotationss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotations.GetDB()

	query := db.Find(&notationsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	notationsAPIs := make([]orm.NotationsAPI, 0)

	// for each notations, update fields from the database nullable fields
	for idx := range notationsDBs {
		notationsDB := &notationsDBs[idx]
		_ = notationsDB
		var notationsAPI orm.NotationsAPI

		// insertion point for updating fields
		notationsAPI.ID = notationsDB.ID
		notationsDB.CopyBasicFieldsToNotations_WOP(&notationsAPI.Notations_WOP)
		notationsAPI.NotationsPointersEncoding = notationsDB.NotationsPointersEncoding
		notationsAPIs = append(notationsAPIs, notationsAPI)
	}

	c.JSON(http.StatusOK, notationsAPIs)
}

// PostNotations
//
// swagger:route POST /notationss notationss postNotations
//
// Creates a notations
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNotations(c *gin.Context) {

	mutexNotations.Lock()
	defer mutexNotations.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNotationss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotations.GetDB()

	// Validate input
	var input orm.NotationsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create notations
	notationsDB := orm.NotationsDB{}
	notationsDB.NotationsPointersEncoding = input.NotationsPointersEncoding
	notationsDB.CopyBasicFieldsFromNotations_WOP(&input.Notations_WOP)

	query := db.Create(&notationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNotations.CheckoutPhaseOneInstance(&notationsDB)
	notations := backRepo.BackRepoNotations.Map_NotationsDBID_NotationsPtr[notationsDB.ID]

	if notations != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), notations)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, notationsDB)
}

// GetNotations
//
// swagger:route GET /notationss/{ID} notationss getNotations
//
// Gets the details for a notations.
//
// Responses:
// default: genericError
//
//	200: notationsDBResponse
func (controller *Controller) GetNotations(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNotations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotations.GetDB()

	// Get notationsDB in DB
	var notationsDB orm.NotationsDB
	if err := db.First(&notationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var notationsAPI orm.NotationsAPI
	notationsAPI.ID = notationsDB.ID
	notationsAPI.NotationsPointersEncoding = notationsDB.NotationsPointersEncoding
	notationsDB.CopyBasicFieldsToNotations_WOP(&notationsAPI.Notations_WOP)

	c.JSON(http.StatusOK, notationsAPI)
}

// UpdateNotations
//
// swagger:route PATCH /notationss/{ID} notationss updateNotations
//
// # Update a notations
//
// Responses:
// default: genericError
//
//	200: notationsDBResponse
func (controller *Controller) UpdateNotations(c *gin.Context) {

	mutexNotations.Lock()
	defer mutexNotations.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNotations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotations.GetDB()

	// Validate input
	var input orm.NotationsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var notationsDB orm.NotationsDB

	// fetch the notations
	query := db.First(&notationsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	notationsDB.CopyBasicFieldsFromNotations_WOP(&input.Notations_WOP)
	notationsDB.NotationsPointersEncoding = input.NotationsPointersEncoding

	query = db.Model(&notationsDB).Updates(notationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	notationsNew := new(models.Notations)
	notationsDB.CopyBasicFieldsToNotations(notationsNew)

	// redeem pointers
	notationsDB.DecodePointers(backRepo, notationsNew)

	// get stage instance from DB instance, and call callback function
	notationsOld := backRepo.BackRepoNotations.Map_NotationsDBID_NotationsPtr[notationsDB.ID]
	if notationsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), notationsOld, notationsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the notationsDB
	c.JSON(http.StatusOK, notationsDB)
}

// DeleteNotations
//
// swagger:route DELETE /notationss/{ID} notationss deleteNotations
//
// # Delete a notations
//
// default: genericError
//
//	200: notationsDBResponse
func (controller *Controller) DeleteNotations(c *gin.Context) {

	mutexNotations.Lock()
	defer mutexNotations.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNotations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNotations.GetDB()

	// Get model if exist
	var notationsDB orm.NotationsDB
	if err := db.First(&notationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&notationsDB)

	// get an instance (not staged) from DB instance, and call callback function
	notationsDeleted := new(models.Notations)
	notationsDB.CopyBasicFieldsToNotations(notationsDeleted)

	// get stage instance from DB instance, and call callback function
	notationsStaged := backRepo.BackRepoNotations.Map_NotationsDBID_NotationsPtr[notationsDB.ID]
	if notationsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), notationsStaged, notationsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
