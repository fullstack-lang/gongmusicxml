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
var __Accord__dummysDeclaration__ models.Accord
var __Accord_time__dummyDeclaration time.Duration

var mutexAccord sync.Mutex

// An AccordID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAccord updateAccord deleteAccord
type AccordID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AccordInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAccord updateAccord
type AccordInput struct {
	// The Accord to submit or modify
	// in: body
	Accord *orm.AccordAPI
}

// GetAccords
//
// swagger:route GET /accords accords getAccords
//
// # Get all accords
//
// Responses:
// default: genericError
//
//	200: accordDBResponse
func (controller *Controller) GetAccords(c *gin.Context) {

	// source slice
	var accordDBs []orm.AccordDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccords", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccord.GetDB()

	query := db.Find(&accordDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	accordAPIs := make([]orm.AccordAPI, 0)

	// for each accord, update fields from the database nullable fields
	for idx := range accordDBs {
		accordDB := &accordDBs[idx]
		_ = accordDB
		var accordAPI orm.AccordAPI

		// insertion point for updating fields
		accordAPI.ID = accordDB.ID
		accordDB.CopyBasicFieldsToAccord_WOP(&accordAPI.Accord_WOP)
		accordAPI.AccordPointersEncoding = accordDB.AccordPointersEncoding
		accordAPIs = append(accordAPIs, accordAPI)
	}

	c.JSON(http.StatusOK, accordAPIs)
}

// PostAccord
//
// swagger:route POST /accords accords postAccord
//
// Creates a accord
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAccord(c *gin.Context) {

	mutexAccord.Lock()
	defer mutexAccord.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAccords", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccord.GetDB()

	// Validate input
	var input orm.AccordAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create accord
	accordDB := orm.AccordDB{}
	accordDB.AccordPointersEncoding = input.AccordPointersEncoding
	accordDB.CopyBasicFieldsFromAccord_WOP(&input.Accord_WOP)

	query := db.Create(&accordDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAccord.CheckoutPhaseOneInstance(&accordDB)
	accord := backRepo.BackRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID]

	if accord != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), accord)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, accordDB)
}

// GetAccord
//
// swagger:route GET /accords/{ID} accords getAccord
//
// Gets the details for a accord.
//
// Responses:
// default: genericError
//
//	200: accordDBResponse
func (controller *Controller) GetAccord(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccord", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccord.GetDB()

	// Get accordDB in DB
	var accordDB orm.AccordDB
	if err := db.First(&accordDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var accordAPI orm.AccordAPI
	accordAPI.ID = accordDB.ID
	accordAPI.AccordPointersEncoding = accordDB.AccordPointersEncoding
	accordDB.CopyBasicFieldsToAccord_WOP(&accordAPI.Accord_WOP)

	c.JSON(http.StatusOK, accordAPI)
}

// UpdateAccord
//
// swagger:route PATCH /accords/{ID} accords updateAccord
//
// # Update a accord
//
// Responses:
// default: genericError
//
//	200: accordDBResponse
func (controller *Controller) UpdateAccord(c *gin.Context) {

	mutexAccord.Lock()
	defer mutexAccord.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAccord", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccord.GetDB()

	// Validate input
	var input orm.AccordAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var accordDB orm.AccordDB

	// fetch the accord
	query := db.First(&accordDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	accordDB.CopyBasicFieldsFromAccord_WOP(&input.Accord_WOP)
	accordDB.AccordPointersEncoding = input.AccordPointersEncoding

	query = db.Model(&accordDB).Updates(accordDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	accordNew := new(models.Accord)
	accordDB.CopyBasicFieldsToAccord(accordNew)

	// redeem pointers
	accordDB.DecodePointers(backRepo, accordNew)

	// get stage instance from DB instance, and call callback function
	accordOld := backRepo.BackRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID]
	if accordOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), accordOld, accordNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the accordDB
	c.JSON(http.StatusOK, accordDB)
}

// DeleteAccord
//
// swagger:route DELETE /accords/{ID} accords deleteAccord
//
// # Delete a accord
//
// default: genericError
//
//	200: accordDBResponse
func (controller *Controller) DeleteAccord(c *gin.Context) {

	mutexAccord.Lock()
	defer mutexAccord.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAccord", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccord.GetDB()

	// Get model if exist
	var accordDB orm.AccordDB
	if err := db.First(&accordDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&accordDB)

	// get an instance (not staged) from DB instance, and call callback function
	accordDeleted := new(models.Accord)
	accordDB.CopyBasicFieldsToAccord(accordDeleted)

	// get stage instance from DB instance, and call callback function
	accordStaged := backRepo.BackRepoAccord.Map_AccordDBID_AccordPtr[accordDB.ID]
	if accordStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), accordStaged, accordDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
