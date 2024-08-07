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
var __Scordatura__dummysDeclaration__ models.Scordatura
var __Scordatura_time__dummyDeclaration time.Duration

var mutexScordatura sync.Mutex

// An ScordaturaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScordatura updateScordatura deleteScordatura
type ScordaturaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ScordaturaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScordatura updateScordatura
type ScordaturaInput struct {
	// The Scordatura to submit or modify
	// in: body
	Scordatura *orm.ScordaturaAPI
}

// GetScordaturas
//
// swagger:route GET /scordaturas scordaturas getScordaturas
//
// # Get all scordaturas
//
// Responses:
// default: genericError
//
//	200: scordaturaDBResponse
func (controller *Controller) GetScordaturas(c *gin.Context) {

	// source slice
	var scordaturaDBs []orm.ScordaturaDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScordaturas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScordatura.GetDB()

	query := db.Find(&scordaturaDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	scordaturaAPIs := make([]orm.ScordaturaAPI, 0)

	// for each scordatura, update fields from the database nullable fields
	for idx := range scordaturaDBs {
		scordaturaDB := &scordaturaDBs[idx]
		_ = scordaturaDB
		var scordaturaAPI orm.ScordaturaAPI

		// insertion point for updating fields
		scordaturaAPI.ID = scordaturaDB.ID
		scordaturaDB.CopyBasicFieldsToScordatura_WOP(&scordaturaAPI.Scordatura_WOP)
		scordaturaAPI.ScordaturaPointersEncoding = scordaturaDB.ScordaturaPointersEncoding
		scordaturaAPIs = append(scordaturaAPIs, scordaturaAPI)
	}

	c.JSON(http.StatusOK, scordaturaAPIs)
}

// PostScordatura
//
// swagger:route POST /scordaturas scordaturas postScordatura
//
// Creates a scordatura
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScordatura(c *gin.Context) {

	mutexScordatura.Lock()
	defer mutexScordatura.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScordaturas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScordatura.GetDB()

	// Validate input
	var input orm.ScordaturaAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create scordatura
	scordaturaDB := orm.ScordaturaDB{}
	scordaturaDB.ScordaturaPointersEncoding = input.ScordaturaPointersEncoding
	scordaturaDB.CopyBasicFieldsFromScordatura_WOP(&input.Scordatura_WOP)

	query := db.Create(&scordaturaDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScordatura.CheckoutPhaseOneInstance(&scordaturaDB)
	scordatura := backRepo.BackRepoScordatura.Map_ScordaturaDBID_ScordaturaPtr[scordaturaDB.ID]

	if scordatura != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), scordatura)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, scordaturaDB)
}

// GetScordatura
//
// swagger:route GET /scordaturas/{ID} scordaturas getScordatura
//
// Gets the details for a scordatura.
//
// Responses:
// default: genericError
//
//	200: scordaturaDBResponse
func (controller *Controller) GetScordatura(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScordatura", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScordatura.GetDB()

	// Get scordaturaDB in DB
	var scordaturaDB orm.ScordaturaDB
	if err := db.First(&scordaturaDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var scordaturaAPI orm.ScordaturaAPI
	scordaturaAPI.ID = scordaturaDB.ID
	scordaturaAPI.ScordaturaPointersEncoding = scordaturaDB.ScordaturaPointersEncoding
	scordaturaDB.CopyBasicFieldsToScordatura_WOP(&scordaturaAPI.Scordatura_WOP)

	c.JSON(http.StatusOK, scordaturaAPI)
}

// UpdateScordatura
//
// swagger:route PATCH /scordaturas/{ID} scordaturas updateScordatura
//
// # Update a scordatura
//
// Responses:
// default: genericError
//
//	200: scordaturaDBResponse
func (controller *Controller) UpdateScordatura(c *gin.Context) {

	mutexScordatura.Lock()
	defer mutexScordatura.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScordatura", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScordatura.GetDB()

	// Validate input
	var input orm.ScordaturaAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var scordaturaDB orm.ScordaturaDB

	// fetch the scordatura
	query := db.First(&scordaturaDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	scordaturaDB.CopyBasicFieldsFromScordatura_WOP(&input.Scordatura_WOP)
	scordaturaDB.ScordaturaPointersEncoding = input.ScordaturaPointersEncoding

	query = db.Model(&scordaturaDB).Updates(scordaturaDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	scordaturaNew := new(models.Scordatura)
	scordaturaDB.CopyBasicFieldsToScordatura(scordaturaNew)

	// redeem pointers
	scordaturaDB.DecodePointers(backRepo, scordaturaNew)

	// get stage instance from DB instance, and call callback function
	scordaturaOld := backRepo.BackRepoScordatura.Map_ScordaturaDBID_ScordaturaPtr[scordaturaDB.ID]
	if scordaturaOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), scordaturaOld, scordaturaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the scordaturaDB
	c.JSON(http.StatusOK, scordaturaDB)
}

// DeleteScordatura
//
// swagger:route DELETE /scordaturas/{ID} scordaturas deleteScordatura
//
// # Delete a scordatura
//
// default: genericError
//
//	200: scordaturaDBResponse
func (controller *Controller) DeleteScordatura(c *gin.Context) {

	mutexScordatura.Lock()
	defer mutexScordatura.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScordatura", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScordatura.GetDB()

	// Get model if exist
	var scordaturaDB orm.ScordaturaDB
	if err := db.First(&scordaturaDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&scordaturaDB)

	// get an instance (not staged) from DB instance, and call callback function
	scordaturaDeleted := new(models.Scordatura)
	scordaturaDB.CopyBasicFieldsToScordatura(scordaturaDeleted)

	// get stage instance from DB instance, and call callback function
	scordaturaStaged := backRepo.BackRepoScordatura.Map_ScordaturaDBID_ScordaturaPtr[scordaturaDB.ID]
	if scordaturaStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), scordaturaStaged, scordaturaDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
