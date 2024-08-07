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
var __Opus__dummysDeclaration__ models.Opus
var __Opus_time__dummyDeclaration time.Duration

var mutexOpus sync.Mutex

// An OpusID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOpus updateOpus deleteOpus
type OpusID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OpusInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOpus updateOpus
type OpusInput struct {
	// The Opus to submit or modify
	// in: body
	Opus *orm.OpusAPI
}

// GetOpuss
//
// swagger:route GET /opuss opuss getOpuss
//
// # Get all opuss
//
// Responses:
// default: genericError
//
//	200: opusDBResponse
func (controller *Controller) GetOpuss(c *gin.Context) {

	// source slice
	var opusDBs []orm.OpusDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOpuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpus.GetDB()

	query := db.Find(&opusDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	opusAPIs := make([]orm.OpusAPI, 0)

	// for each opus, update fields from the database nullable fields
	for idx := range opusDBs {
		opusDB := &opusDBs[idx]
		_ = opusDB
		var opusAPI orm.OpusAPI

		// insertion point for updating fields
		opusAPI.ID = opusDB.ID
		opusDB.CopyBasicFieldsToOpus_WOP(&opusAPI.Opus_WOP)
		opusAPI.OpusPointersEncoding = opusDB.OpusPointersEncoding
		opusAPIs = append(opusAPIs, opusAPI)
	}

	c.JSON(http.StatusOK, opusAPIs)
}

// PostOpus
//
// swagger:route POST /opuss opuss postOpus
//
// Creates a opus
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOpus(c *gin.Context) {

	mutexOpus.Lock()
	defer mutexOpus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOpuss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpus.GetDB()

	// Validate input
	var input orm.OpusAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create opus
	opusDB := orm.OpusDB{}
	opusDB.OpusPointersEncoding = input.OpusPointersEncoding
	opusDB.CopyBasicFieldsFromOpus_WOP(&input.Opus_WOP)

	query := db.Create(&opusDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOpus.CheckoutPhaseOneInstance(&opusDB)
	opus := backRepo.BackRepoOpus.Map_OpusDBID_OpusPtr[opusDB.ID]

	if opus != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), opus)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, opusDB)
}

// GetOpus
//
// swagger:route GET /opuss/{ID} opuss getOpus
//
// Gets the details for a opus.
//
// Responses:
// default: genericError
//
//	200: opusDBResponse
func (controller *Controller) GetOpus(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOpus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpus.GetDB()

	// Get opusDB in DB
	var opusDB orm.OpusDB
	if err := db.First(&opusDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var opusAPI orm.OpusAPI
	opusAPI.ID = opusDB.ID
	opusAPI.OpusPointersEncoding = opusDB.OpusPointersEncoding
	opusDB.CopyBasicFieldsToOpus_WOP(&opusAPI.Opus_WOP)

	c.JSON(http.StatusOK, opusAPI)
}

// UpdateOpus
//
// swagger:route PATCH /opuss/{ID} opuss updateOpus
//
// # Update a opus
//
// Responses:
// default: genericError
//
//	200: opusDBResponse
func (controller *Controller) UpdateOpus(c *gin.Context) {

	mutexOpus.Lock()
	defer mutexOpus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOpus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpus.GetDB()

	// Validate input
	var input orm.OpusAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var opusDB orm.OpusDB

	// fetch the opus
	query := db.First(&opusDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	opusDB.CopyBasicFieldsFromOpus_WOP(&input.Opus_WOP)
	opusDB.OpusPointersEncoding = input.OpusPointersEncoding

	query = db.Model(&opusDB).Updates(opusDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	opusNew := new(models.Opus)
	opusDB.CopyBasicFieldsToOpus(opusNew)

	// redeem pointers
	opusDB.DecodePointers(backRepo, opusNew)

	// get stage instance from DB instance, and call callback function
	opusOld := backRepo.BackRepoOpus.Map_OpusDBID_OpusPtr[opusDB.ID]
	if opusOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), opusOld, opusNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the opusDB
	c.JSON(http.StatusOK, opusDB)
}

// DeleteOpus
//
// swagger:route DELETE /opuss/{ID} opuss deleteOpus
//
// # Delete a opus
//
// default: genericError
//
//	200: opusDBResponse
func (controller *Controller) DeleteOpus(c *gin.Context) {

	mutexOpus.Lock()
	defer mutexOpus.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOpus", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOpus.GetDB()

	// Get model if exist
	var opusDB orm.OpusDB
	if err := db.First(&opusDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&opusDB)

	// get an instance (not staged) from DB instance, and call callback function
	opusDeleted := new(models.Opus)
	opusDB.CopyBasicFieldsToOpus(opusDeleted)

	// get stage instance from DB instance, and call callback function
	opusStaged := backRepo.BackRepoOpus.Map_OpusDBID_OpusPtr[opusDB.ID]
	if opusStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), opusStaged, opusDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
