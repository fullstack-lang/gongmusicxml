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
var __Listen__dummysDeclaration__ models.Listen
var __Listen_time__dummyDeclaration time.Duration

var mutexListen sync.Mutex

// An ListenID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getListen updateListen deleteListen
type ListenID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ListenInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postListen updateListen
type ListenInput struct {
	// The Listen to submit or modify
	// in: body
	Listen *orm.ListenAPI
}

// GetListens
//
// swagger:route GET /listens listens getListens
//
// # Get all listens
//
// Responses:
// default: genericError
//
//	200: listenDBResponse
func (controller *Controller) GetListens(c *gin.Context) {

	// source slice
	var listenDBs []orm.ListenDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetListens", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListen.GetDB()

	query := db.Find(&listenDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	listenAPIs := make([]orm.ListenAPI, 0)

	// for each listen, update fields from the database nullable fields
	for idx := range listenDBs {
		listenDB := &listenDBs[idx]
		_ = listenDB
		var listenAPI orm.ListenAPI

		// insertion point for updating fields
		listenAPI.ID = listenDB.ID
		listenDB.CopyBasicFieldsToListen_WOP(&listenAPI.Listen_WOP)
		listenAPI.ListenPointersEncoding = listenDB.ListenPointersEncoding
		listenAPIs = append(listenAPIs, listenAPI)
	}

	c.JSON(http.StatusOK, listenAPIs)
}

// PostListen
//
// swagger:route POST /listens listens postListen
//
// Creates a listen
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostListen(c *gin.Context) {

	mutexListen.Lock()
	defer mutexListen.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostListens", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListen.GetDB()

	// Validate input
	var input orm.ListenAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create listen
	listenDB := orm.ListenDB{}
	listenDB.ListenPointersEncoding = input.ListenPointersEncoding
	listenDB.CopyBasicFieldsFromListen_WOP(&input.Listen_WOP)

	query := db.Create(&listenDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoListen.CheckoutPhaseOneInstance(&listenDB)
	listen := backRepo.BackRepoListen.Map_ListenDBID_ListenPtr[listenDB.ID]

	if listen != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), listen)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, listenDB)
}

// GetListen
//
// swagger:route GET /listens/{ID} listens getListen
//
// Gets the details for a listen.
//
// Responses:
// default: genericError
//
//	200: listenDBResponse
func (controller *Controller) GetListen(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetListen", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListen.GetDB()

	// Get listenDB in DB
	var listenDB orm.ListenDB
	if err := db.First(&listenDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var listenAPI orm.ListenAPI
	listenAPI.ID = listenDB.ID
	listenAPI.ListenPointersEncoding = listenDB.ListenPointersEncoding
	listenDB.CopyBasicFieldsToListen_WOP(&listenAPI.Listen_WOP)

	c.JSON(http.StatusOK, listenAPI)
}

// UpdateListen
//
// swagger:route PATCH /listens/{ID} listens updateListen
//
// # Update a listen
//
// Responses:
// default: genericError
//
//	200: listenDBResponse
func (controller *Controller) UpdateListen(c *gin.Context) {

	mutexListen.Lock()
	defer mutexListen.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateListen", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListen.GetDB()

	// Validate input
	var input orm.ListenAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var listenDB orm.ListenDB

	// fetch the listen
	query := db.First(&listenDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	listenDB.CopyBasicFieldsFromListen_WOP(&input.Listen_WOP)
	listenDB.ListenPointersEncoding = input.ListenPointersEncoding

	query = db.Model(&listenDB).Updates(listenDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	listenNew := new(models.Listen)
	listenDB.CopyBasicFieldsToListen(listenNew)

	// redeem pointers
	listenDB.DecodePointers(backRepo, listenNew)

	// get stage instance from DB instance, and call callback function
	listenOld := backRepo.BackRepoListen.Map_ListenDBID_ListenPtr[listenDB.ID]
	if listenOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), listenOld, listenNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the listenDB
	c.JSON(http.StatusOK, listenDB)
}

// DeleteListen
//
// swagger:route DELETE /listens/{ID} listens deleteListen
//
// # Delete a listen
//
// default: genericError
//
//	200: listenDBResponse
func (controller *Controller) DeleteListen(c *gin.Context) {

	mutexListen.Lock()
	defer mutexListen.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteListen", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoListen.GetDB()

	// Get model if exist
	var listenDB orm.ListenDB
	if err := db.First(&listenDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&listenDB)

	// get an instance (not staged) from DB instance, and call callback function
	listenDeleted := new(models.Listen)
	listenDB.CopyBasicFieldsToListen(listenDeleted)

	// get stage instance from DB instance, and call callback function
	listenStaged := backRepo.BackRepoListen.Map_ListenDBID_ListenPtr[listenDB.ID]
	if listenStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), listenStaged, listenDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
