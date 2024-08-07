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
var __Listening__dummysDeclaration__ models.Listening
var __Listening_time__dummyDeclaration time.Duration

var mutexListening sync.Mutex

// An ListeningID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getListening updateListening deleteListening
type ListeningID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ListeningInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postListening updateListening
type ListeningInput struct {
	// The Listening to submit or modify
	// in: body
	Listening *orm.ListeningAPI
}

// GetListenings
//
// swagger:route GET /listenings listenings getListenings
//
// # Get all listenings
//
// Responses:
// default: genericError
//
//	200: listeningDBResponse
func (controller *Controller) GetListenings(c *gin.Context) {

	// source slice
	var listeningDBs []orm.ListeningDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetListenings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListening.GetDB()

	query := db.Find(&listeningDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	listeningAPIs := make([]orm.ListeningAPI, 0)

	// for each listening, update fields from the database nullable fields
	for idx := range listeningDBs {
		listeningDB := &listeningDBs[idx]
		_ = listeningDB
		var listeningAPI orm.ListeningAPI

		// insertion point for updating fields
		listeningAPI.ID = listeningDB.ID
		listeningDB.CopyBasicFieldsToListening_WOP(&listeningAPI.Listening_WOP)
		listeningAPI.ListeningPointersEncoding = listeningDB.ListeningPointersEncoding
		listeningAPIs = append(listeningAPIs, listeningAPI)
	}

	c.JSON(http.StatusOK, listeningAPIs)
}

// PostListening
//
// swagger:route POST /listenings listenings postListening
//
// Creates a listening
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostListening(c *gin.Context) {

	mutexListening.Lock()
	defer mutexListening.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostListenings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListening.GetDB()

	// Validate input
	var input orm.ListeningAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create listening
	listeningDB := orm.ListeningDB{}
	listeningDB.ListeningPointersEncoding = input.ListeningPointersEncoding
	listeningDB.CopyBasicFieldsFromListening_WOP(&input.Listening_WOP)

	query := db.Create(&listeningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoListening.CheckoutPhaseOneInstance(&listeningDB)
	listening := backRepo.BackRepoListening.Map_ListeningDBID_ListeningPtr[listeningDB.ID]

	if listening != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), listening)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, listeningDB)
}

// GetListening
//
// swagger:route GET /listenings/{ID} listenings getListening
//
// Gets the details for a listening.
//
// Responses:
// default: genericError
//
//	200: listeningDBResponse
func (controller *Controller) GetListening(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetListening", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListening.GetDB()

	// Get listeningDB in DB
	var listeningDB orm.ListeningDB
	if err := db.First(&listeningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var listeningAPI orm.ListeningAPI
	listeningAPI.ID = listeningDB.ID
	listeningAPI.ListeningPointersEncoding = listeningDB.ListeningPointersEncoding
	listeningDB.CopyBasicFieldsToListening_WOP(&listeningAPI.Listening_WOP)

	c.JSON(http.StatusOK, listeningAPI)
}

// UpdateListening
//
// swagger:route PATCH /listenings/{ID} listenings updateListening
//
// # Update a listening
//
// Responses:
// default: genericError
//
//	200: listeningDBResponse
func (controller *Controller) UpdateListening(c *gin.Context) {

	mutexListening.Lock()
	defer mutexListening.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateListening", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListening.GetDB()

	// Validate input
	var input orm.ListeningAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var listeningDB orm.ListeningDB

	// fetch the listening
	query := db.First(&listeningDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	listeningDB.CopyBasicFieldsFromListening_WOP(&input.Listening_WOP)
	listeningDB.ListeningPointersEncoding = input.ListeningPointersEncoding

	query = db.Model(&listeningDB).Updates(listeningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	listeningNew := new(models.Listening)
	listeningDB.CopyBasicFieldsToListening(listeningNew)

	// redeem pointers
	listeningDB.DecodePointers(backRepo, listeningNew)

	// get stage instance from DB instance, and call callback function
	listeningOld := backRepo.BackRepoListening.Map_ListeningDBID_ListeningPtr[listeningDB.ID]
	if listeningOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), listeningOld, listeningNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the listeningDB
	c.JSON(http.StatusOK, listeningDB)
}

// DeleteListening
//
// swagger:route DELETE /listenings/{ID} listenings deleteListening
//
// # Delete a listening
//
// default: genericError
//
//	200: listeningDBResponse
func (controller *Controller) DeleteListening(c *gin.Context) {

	mutexListening.Lock()
	defer mutexListening.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteListening", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListening.GetDB()

	// Get model if exist
	var listeningDB orm.ListeningDB
	if err := db.First(&listeningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&listeningDB)

	// get an instance (not staged) from DB instance, and call callback function
	listeningDeleted := new(models.Listening)
	listeningDB.CopyBasicFieldsToListening(listeningDeleted)

	// get stage instance from DB instance, and call callback function
	listeningStaged := backRepo.BackRepoListening.Map_ListeningDBID_ListeningPtr[listeningDB.ID]
	if listeningStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), listeningStaged, listeningDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
