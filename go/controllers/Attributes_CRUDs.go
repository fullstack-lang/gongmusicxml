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
var __Attributes__dummysDeclaration__ models.Attributes
var __Attributes_time__dummyDeclaration time.Duration

var mutexAttributes sync.Mutex

// An AttributesID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAttributes updateAttributes deleteAttributes
type AttributesID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AttributesInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAttributes updateAttributes
type AttributesInput struct {
	// The Attributes to submit or modify
	// in: body
	Attributes *orm.AttributesAPI
}

// GetAttributess
//
// swagger:route GET /attributess attributess getAttributess
//
// # Get all attributess
//
// Responses:
// default: genericError
//
//	200: attributesDBResponse
func (controller *Controller) GetAttributess(c *gin.Context) {

	// source slice
	var attributesDBs []orm.AttributesDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributes.GetDB()

	query := db.Find(&attributesDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attributesAPIs := make([]orm.AttributesAPI, 0)

	// for each attributes, update fields from the database nullable fields
	for idx := range attributesDBs {
		attributesDB := &attributesDBs[idx]
		_ = attributesDB
		var attributesAPI orm.AttributesAPI

		// insertion point for updating fields
		attributesAPI.ID = attributesDB.ID
		attributesDB.CopyBasicFieldsToAttributes_WOP(&attributesAPI.Attributes_WOP)
		attributesAPI.AttributesPointersEncoding = attributesDB.AttributesPointersEncoding
		attributesAPIs = append(attributesAPIs, attributesAPI)
	}

	c.JSON(http.StatusOK, attributesAPIs)
}

// PostAttributes
//
// swagger:route POST /attributess attributess postAttributes
//
// Creates a attributes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAttributes(c *gin.Context) {

	mutexAttributes.Lock()
	defer mutexAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAttributess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributes.GetDB()

	// Validate input
	var input orm.AttributesAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attributes
	attributesDB := orm.AttributesDB{}
	attributesDB.AttributesPointersEncoding = input.AttributesPointersEncoding
	attributesDB.CopyBasicFieldsFromAttributes_WOP(&input.Attributes_WOP)

	query := db.Create(&attributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAttributes.CheckoutPhaseOneInstance(&attributesDB)
	attributes := backRepo.BackRepoAttributes.Map_AttributesDBID_AttributesPtr[attributesDB.ID]

	if attributes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attributes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attributesDB)
}

// GetAttributes
//
// swagger:route GET /attributess/{ID} attributess getAttributes
//
// Gets the details for a attributes.
//
// Responses:
// default: genericError
//
//	200: attributesDBResponse
func (controller *Controller) GetAttributes(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributes.GetDB()

	// Get attributesDB in DB
	var attributesDB orm.AttributesDB
	if err := db.First(&attributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attributesAPI orm.AttributesAPI
	attributesAPI.ID = attributesDB.ID
	attributesAPI.AttributesPointersEncoding = attributesDB.AttributesPointersEncoding
	attributesDB.CopyBasicFieldsToAttributes_WOP(&attributesAPI.Attributes_WOP)

	c.JSON(http.StatusOK, attributesAPI)
}

// UpdateAttributes
//
// swagger:route PATCH /attributess/{ID} attributess updateAttributes
//
// # Update a attributes
//
// Responses:
// default: genericError
//
//	200: attributesDBResponse
func (controller *Controller) UpdateAttributes(c *gin.Context) {

	mutexAttributes.Lock()
	defer mutexAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributes.GetDB()

	// Validate input
	var input orm.AttributesAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attributesDB orm.AttributesDB

	// fetch the attributes
	query := db.First(&attributesDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attributesDB.CopyBasicFieldsFromAttributes_WOP(&input.Attributes_WOP)
	attributesDB.AttributesPointersEncoding = input.AttributesPointersEncoding

	query = db.Model(&attributesDB).Updates(attributesDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attributesNew := new(models.Attributes)
	attributesDB.CopyBasicFieldsToAttributes(attributesNew)

	// redeem pointers
	attributesDB.DecodePointers(backRepo, attributesNew)

	// get stage instance from DB instance, and call callback function
	attributesOld := backRepo.BackRepoAttributes.Map_AttributesDBID_AttributesPtr[attributesDB.ID]
	if attributesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attributesOld, attributesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attributesDB
	c.JSON(http.StatusOK, attributesDB)
}

// DeleteAttributes
//
// swagger:route DELETE /attributess/{ID} attributess deleteAttributes
//
// # Delete a attributes
//
// default: genericError
//
//	200: attributesDBResponse
func (controller *Controller) DeleteAttributes(c *gin.Context) {

	mutexAttributes.Lock()
	defer mutexAttributes.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAttributes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAttributes.GetDB()

	// Get model if exist
	var attributesDB orm.AttributesDB
	if err := db.First(&attributesDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attributesDB)

	// get an instance (not staged) from DB instance, and call callback function
	attributesDeleted := new(models.Attributes)
	attributesDB.CopyBasicFieldsToAttributes(attributesDeleted)

	// get stage instance from DB instance, and call callback function
	attributesStaged := backRepo.BackRepoAttributes.Map_AttributesDBID_AttributesPtr[attributesDB.ID]
	if attributesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attributesStaged, attributesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
