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
var __Identification__dummysDeclaration__ models.Identification
var __Identification_time__dummyDeclaration time.Duration

var mutexIdentification sync.Mutex

// An IdentificationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getIdentification updateIdentification deleteIdentification
type IdentificationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// IdentificationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postIdentification updateIdentification
type IdentificationInput struct {
	// The Identification to submit or modify
	// in: body
	Identification *orm.IdentificationAPI
}

// GetIdentifications
//
// swagger:route GET /identifications identifications getIdentifications
//
// # Get all identifications
//
// Responses:
// default: genericError
//
//	200: identificationDBResponse
func (controller *Controller) GetIdentifications(c *gin.Context) {

	// source slice
	var identificationDBs []orm.IdentificationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetIdentifications", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoIdentification.GetDB()

	query := db.Find(&identificationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	identificationAPIs := make([]orm.IdentificationAPI, 0)

	// for each identification, update fields from the database nullable fields
	for idx := range identificationDBs {
		identificationDB := &identificationDBs[idx]
		_ = identificationDB
		var identificationAPI orm.IdentificationAPI

		// insertion point for updating fields
		identificationAPI.ID = identificationDB.ID
		identificationDB.CopyBasicFieldsToIdentification_WOP(&identificationAPI.Identification_WOP)
		identificationAPI.IdentificationPointersEncoding = identificationDB.IdentificationPointersEncoding
		identificationAPIs = append(identificationAPIs, identificationAPI)
	}

	c.JSON(http.StatusOK, identificationAPIs)
}

// PostIdentification
//
// swagger:route POST /identifications identifications postIdentification
//
// Creates a identification
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostIdentification(c *gin.Context) {

	mutexIdentification.Lock()
	defer mutexIdentification.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostIdentifications", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoIdentification.GetDB()

	// Validate input
	var input orm.IdentificationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create identification
	identificationDB := orm.IdentificationDB{}
	identificationDB.IdentificationPointersEncoding = input.IdentificationPointersEncoding
	identificationDB.CopyBasicFieldsFromIdentification_WOP(&input.Identification_WOP)

	query := db.Create(&identificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoIdentification.CheckoutPhaseOneInstance(&identificationDB)
	identification := backRepo.BackRepoIdentification.Map_IdentificationDBID_IdentificationPtr[identificationDB.ID]

	if identification != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), identification)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, identificationDB)
}

// GetIdentification
//
// swagger:route GET /identifications/{ID} identifications getIdentification
//
// Gets the details for a identification.
//
// Responses:
// default: genericError
//
//	200: identificationDBResponse
func (controller *Controller) GetIdentification(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetIdentification", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoIdentification.GetDB()

	// Get identificationDB in DB
	var identificationDB orm.IdentificationDB
	if err := db.First(&identificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var identificationAPI orm.IdentificationAPI
	identificationAPI.ID = identificationDB.ID
	identificationAPI.IdentificationPointersEncoding = identificationDB.IdentificationPointersEncoding
	identificationDB.CopyBasicFieldsToIdentification_WOP(&identificationAPI.Identification_WOP)

	c.JSON(http.StatusOK, identificationAPI)
}

// UpdateIdentification
//
// swagger:route PATCH /identifications/{ID} identifications updateIdentification
//
// # Update a identification
//
// Responses:
// default: genericError
//
//	200: identificationDBResponse
func (controller *Controller) UpdateIdentification(c *gin.Context) {

	mutexIdentification.Lock()
	defer mutexIdentification.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateIdentification", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoIdentification.GetDB()

	// Validate input
	var input orm.IdentificationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var identificationDB orm.IdentificationDB

	// fetch the identification
	query := db.First(&identificationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	identificationDB.CopyBasicFieldsFromIdentification_WOP(&input.Identification_WOP)
	identificationDB.IdentificationPointersEncoding = input.IdentificationPointersEncoding

	query = db.Model(&identificationDB).Updates(identificationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	identificationNew := new(models.Identification)
	identificationDB.CopyBasicFieldsToIdentification(identificationNew)

	// redeem pointers
	identificationDB.DecodePointers(backRepo, identificationNew)

	// get stage instance from DB instance, and call callback function
	identificationOld := backRepo.BackRepoIdentification.Map_IdentificationDBID_IdentificationPtr[identificationDB.ID]
	if identificationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), identificationOld, identificationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the identificationDB
	c.JSON(http.StatusOK, identificationDB)
}

// DeleteIdentification
//
// swagger:route DELETE /identifications/{ID} identifications deleteIdentification
//
// # Delete a identification
//
// default: genericError
//
//	200: identificationDBResponse
func (controller *Controller) DeleteIdentification(c *gin.Context) {

	mutexIdentification.Lock()
	defer mutexIdentification.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteIdentification", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoIdentification.GetDB()

	// Get model if exist
	var identificationDB orm.IdentificationDB
	if err := db.First(&identificationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&identificationDB)

	// get an instance (not staged) from DB instance, and call callback function
	identificationDeleted := new(models.Identification)
	identificationDB.CopyBasicFieldsToIdentification(identificationDeleted)

	// get stage instance from DB instance, and call callback function
	identificationStaged := backRepo.BackRepoIdentification.Map_IdentificationDBID_IdentificationPtr[identificationDB.ID]
	if identificationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), identificationStaged, identificationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
