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
var __Rest__dummysDeclaration__ models.Rest
var __Rest_time__dummyDeclaration time.Duration

var mutexRest sync.Mutex

// An RestID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRest updateRest deleteRest
type RestID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RestInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRest updateRest
type RestInput struct {
	// The Rest to submit or modify
	// in: body
	Rest *orm.RestAPI
}

// GetRests
//
// swagger:route GET /rests rests getRests
//
// # Get all rests
//
// Responses:
// default: genericError
//
//	200: restDBResponse
func (controller *Controller) GetRests(c *gin.Context) {

	// source slice
	var restDBs []orm.RestDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRests", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRest.GetDB()

	query := db.Find(&restDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	restAPIs := make([]orm.RestAPI, 0)

	// for each rest, update fields from the database nullable fields
	for idx := range restDBs {
		restDB := &restDBs[idx]
		_ = restDB
		var restAPI orm.RestAPI

		// insertion point for updating fields
		restAPI.ID = restDB.ID
		restDB.CopyBasicFieldsToRest_WOP(&restAPI.Rest_WOP)
		restAPI.RestPointersEncoding = restDB.RestPointersEncoding
		restAPIs = append(restAPIs, restAPI)
	}

	c.JSON(http.StatusOK, restAPIs)
}

// PostRest
//
// swagger:route POST /rests rests postRest
//
// Creates a rest
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRest(c *gin.Context) {

	mutexRest.Lock()
	defer mutexRest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRests", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRest.GetDB()

	// Validate input
	var input orm.RestAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rest
	restDB := orm.RestDB{}
	restDB.RestPointersEncoding = input.RestPointersEncoding
	restDB.CopyBasicFieldsFromRest_WOP(&input.Rest_WOP)

	query := db.Create(&restDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRest.CheckoutPhaseOneInstance(&restDB)
	rest := backRepo.BackRepoRest.Map_RestDBID_RestPtr[restDB.ID]

	if rest != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rest)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, restDB)
}

// GetRest
//
// swagger:route GET /rests/{ID} rests getRest
//
// Gets the details for a rest.
//
// Responses:
// default: genericError
//
//	200: restDBResponse
func (controller *Controller) GetRest(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRest.GetDB()

	// Get restDB in DB
	var restDB orm.RestDB
	if err := db.First(&restDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var restAPI orm.RestAPI
	restAPI.ID = restDB.ID
	restAPI.RestPointersEncoding = restDB.RestPointersEncoding
	restDB.CopyBasicFieldsToRest_WOP(&restAPI.Rest_WOP)

	c.JSON(http.StatusOK, restAPI)
}

// UpdateRest
//
// swagger:route PATCH /rests/{ID} rests updateRest
//
// # Update a rest
//
// Responses:
// default: genericError
//
//	200: restDBResponse
func (controller *Controller) UpdateRest(c *gin.Context) {

	mutexRest.Lock()
	defer mutexRest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRest.GetDB()

	// Validate input
	var input orm.RestAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var restDB orm.RestDB

	// fetch the rest
	query := db.First(&restDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	restDB.CopyBasicFieldsFromRest_WOP(&input.Rest_WOP)
	restDB.RestPointersEncoding = input.RestPointersEncoding

	query = db.Model(&restDB).Updates(restDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	restNew := new(models.Rest)
	restDB.CopyBasicFieldsToRest(restNew)

	// redeem pointers
	restDB.DecodePointers(backRepo, restNew)

	// get stage instance from DB instance, and call callback function
	restOld := backRepo.BackRepoRest.Map_RestDBID_RestPtr[restDB.ID]
	if restOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), restOld, restNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the restDB
	c.JSON(http.StatusOK, restDB)
}

// DeleteRest
//
// swagger:route DELETE /rests/{ID} rests deleteRest
//
// # Delete a rest
//
// default: genericError
//
//	200: restDBResponse
func (controller *Controller) DeleteRest(c *gin.Context) {

	mutexRest.Lock()
	defer mutexRest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRest.GetDB()

	// Get model if exist
	var restDB orm.RestDB
	if err := db.First(&restDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&restDB)

	// get an instance (not staged) from DB instance, and call callback function
	restDeleted := new(models.Rest)
	restDB.CopyBasicFieldsToRest(restDeleted)

	// get stage instance from DB instance, and call callback function
	restStaged := backRepo.BackRepoRest.Map_RestDBID_RestPtr[restDB.ID]
	if restStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), restStaged, restDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
