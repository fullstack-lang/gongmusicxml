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
var __Coda__dummysDeclaration__ models.Coda
var __Coda_time__dummyDeclaration time.Duration

var mutexCoda sync.Mutex

// An CodaID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCoda updateCoda deleteCoda
type CodaID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CodaInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCoda updateCoda
type CodaInput struct {
	// The Coda to submit or modify
	// in: body
	Coda *orm.CodaAPI
}

// GetCodas
//
// swagger:route GET /codas codas getCodas
//
// # Get all codas
//
// Responses:
// default: genericError
//
//	200: codaDBResponse
func (controller *Controller) GetCodas(c *gin.Context) {

	// source slice
	var codaDBs []orm.CodaDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCodas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCoda.GetDB()

	query := db.Find(&codaDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	codaAPIs := make([]orm.CodaAPI, 0)

	// for each coda, update fields from the database nullable fields
	for idx := range codaDBs {
		codaDB := &codaDBs[idx]
		_ = codaDB
		var codaAPI orm.CodaAPI

		// insertion point for updating fields
		codaAPI.ID = codaDB.ID
		codaDB.CopyBasicFieldsToCoda_WOP(&codaAPI.Coda_WOP)
		codaAPI.CodaPointersEncoding = codaDB.CodaPointersEncoding
		codaAPIs = append(codaAPIs, codaAPI)
	}

	c.JSON(http.StatusOK, codaAPIs)
}

// PostCoda
//
// swagger:route POST /codas codas postCoda
//
// Creates a coda
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCoda(c *gin.Context) {

	mutexCoda.Lock()
	defer mutexCoda.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCodas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCoda.GetDB()

	// Validate input
	var input orm.CodaAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create coda
	codaDB := orm.CodaDB{}
	codaDB.CodaPointersEncoding = input.CodaPointersEncoding
	codaDB.CopyBasicFieldsFromCoda_WOP(&input.Coda_WOP)

	query := db.Create(&codaDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCoda.CheckoutPhaseOneInstance(&codaDB)
	coda := backRepo.BackRepoCoda.Map_CodaDBID_CodaPtr[codaDB.ID]

	if coda != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), coda)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, codaDB)
}

// GetCoda
//
// swagger:route GET /codas/{ID} codas getCoda
//
// Gets the details for a coda.
//
// Responses:
// default: genericError
//
//	200: codaDBResponse
func (controller *Controller) GetCoda(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCoda", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCoda.GetDB()

	// Get codaDB in DB
	var codaDB orm.CodaDB
	if err := db.First(&codaDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var codaAPI orm.CodaAPI
	codaAPI.ID = codaDB.ID
	codaAPI.CodaPointersEncoding = codaDB.CodaPointersEncoding
	codaDB.CopyBasicFieldsToCoda_WOP(&codaAPI.Coda_WOP)

	c.JSON(http.StatusOK, codaAPI)
}

// UpdateCoda
//
// swagger:route PATCH /codas/{ID} codas updateCoda
//
// # Update a coda
//
// Responses:
// default: genericError
//
//	200: codaDBResponse
func (controller *Controller) UpdateCoda(c *gin.Context) {

	mutexCoda.Lock()
	defer mutexCoda.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCoda", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCoda.GetDB()

	// Validate input
	var input orm.CodaAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var codaDB orm.CodaDB

	// fetch the coda
	query := db.First(&codaDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	codaDB.CopyBasicFieldsFromCoda_WOP(&input.Coda_WOP)
	codaDB.CodaPointersEncoding = input.CodaPointersEncoding

	query = db.Model(&codaDB).Updates(codaDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	codaNew := new(models.Coda)
	codaDB.CopyBasicFieldsToCoda(codaNew)

	// redeem pointers
	codaDB.DecodePointers(backRepo, codaNew)

	// get stage instance from DB instance, and call callback function
	codaOld := backRepo.BackRepoCoda.Map_CodaDBID_CodaPtr[codaDB.ID]
	if codaOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), codaOld, codaNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the codaDB
	c.JSON(http.StatusOK, codaDB)
}

// DeleteCoda
//
// swagger:route DELETE /codas/{ID} codas deleteCoda
//
// # Delete a coda
//
// default: genericError
//
//	200: codaDBResponse
func (controller *Controller) DeleteCoda(c *gin.Context) {

	mutexCoda.Lock()
	defer mutexCoda.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCoda", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCoda.GetDB()

	// Get model if exist
	var codaDB orm.CodaDB
	if err := db.First(&codaDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&codaDB)

	// get an instance (not staged) from DB instance, and call callback function
	codaDeleted := new(models.Coda)
	codaDB.CopyBasicFieldsToCoda(codaDeleted)

	// get stage instance from DB instance, and call callback function
	codaStaged := backRepo.BackRepoCoda.Map_CodaDBID_CodaPtr[codaDB.ID]
	if codaStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), codaStaged, codaDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
