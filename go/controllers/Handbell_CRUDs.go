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
var __Handbell__dummysDeclaration__ models.Handbell
var __Handbell_time__dummyDeclaration time.Duration

var mutexHandbell sync.Mutex

// An HandbellID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHandbell updateHandbell deleteHandbell
type HandbellID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HandbellInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHandbell updateHandbell
type HandbellInput struct {
	// The Handbell to submit or modify
	// in: body
	Handbell *orm.HandbellAPI
}

// GetHandbells
//
// swagger:route GET /handbells handbells getHandbells
//
// # Get all handbells
//
// Responses:
// default: genericError
//
//	200: handbellDBResponse
func (controller *Controller) GetHandbells(c *gin.Context) {

	// source slice
	var handbellDBs []orm.HandbellDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHandbells", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHandbell.GetDB()

	query := db.Find(&handbellDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	handbellAPIs := make([]orm.HandbellAPI, 0)

	// for each handbell, update fields from the database nullable fields
	for idx := range handbellDBs {
		handbellDB := &handbellDBs[idx]
		_ = handbellDB
		var handbellAPI orm.HandbellAPI

		// insertion point for updating fields
		handbellAPI.ID = handbellDB.ID
		handbellDB.CopyBasicFieldsToHandbell_WOP(&handbellAPI.Handbell_WOP)
		handbellAPI.HandbellPointersEncoding = handbellDB.HandbellPointersEncoding
		handbellAPIs = append(handbellAPIs, handbellAPI)
	}

	c.JSON(http.StatusOK, handbellAPIs)
}

// PostHandbell
//
// swagger:route POST /handbells handbells postHandbell
//
// Creates a handbell
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHandbell(c *gin.Context) {

	mutexHandbell.Lock()
	defer mutexHandbell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHandbells", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHandbell.GetDB()

	// Validate input
	var input orm.HandbellAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create handbell
	handbellDB := orm.HandbellDB{}
	handbellDB.HandbellPointersEncoding = input.HandbellPointersEncoding
	handbellDB.CopyBasicFieldsFromHandbell_WOP(&input.Handbell_WOP)

	query := db.Create(&handbellDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHandbell.CheckoutPhaseOneInstance(&handbellDB)
	handbell := backRepo.BackRepoHandbell.Map_HandbellDBID_HandbellPtr[handbellDB.ID]

	if handbell != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), handbell)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, handbellDB)
}

// GetHandbell
//
// swagger:route GET /handbells/{ID} handbells getHandbell
//
// Gets the details for a handbell.
//
// Responses:
// default: genericError
//
//	200: handbellDBResponse
func (controller *Controller) GetHandbell(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHandbell", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHandbell.GetDB()

	// Get handbellDB in DB
	var handbellDB orm.HandbellDB
	if err := db.First(&handbellDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var handbellAPI orm.HandbellAPI
	handbellAPI.ID = handbellDB.ID
	handbellAPI.HandbellPointersEncoding = handbellDB.HandbellPointersEncoding
	handbellDB.CopyBasicFieldsToHandbell_WOP(&handbellAPI.Handbell_WOP)

	c.JSON(http.StatusOK, handbellAPI)
}

// UpdateHandbell
//
// swagger:route PATCH /handbells/{ID} handbells updateHandbell
//
// # Update a handbell
//
// Responses:
// default: genericError
//
//	200: handbellDBResponse
func (controller *Controller) UpdateHandbell(c *gin.Context) {

	mutexHandbell.Lock()
	defer mutexHandbell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHandbell", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHandbell.GetDB()

	// Validate input
	var input orm.HandbellAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var handbellDB orm.HandbellDB

	// fetch the handbell
	query := db.First(&handbellDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	handbellDB.CopyBasicFieldsFromHandbell_WOP(&input.Handbell_WOP)
	handbellDB.HandbellPointersEncoding = input.HandbellPointersEncoding

	query = db.Model(&handbellDB).Updates(handbellDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	handbellNew := new(models.Handbell)
	handbellDB.CopyBasicFieldsToHandbell(handbellNew)

	// redeem pointers
	handbellDB.DecodePointers(backRepo, handbellNew)

	// get stage instance from DB instance, and call callback function
	handbellOld := backRepo.BackRepoHandbell.Map_HandbellDBID_HandbellPtr[handbellDB.ID]
	if handbellOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), handbellOld, handbellNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the handbellDB
	c.JSON(http.StatusOK, handbellDB)
}

// DeleteHandbell
//
// swagger:route DELETE /handbells/{ID} handbells deleteHandbell
//
// # Delete a handbell
//
// default: genericError
//
//	200: handbellDBResponse
func (controller *Controller) DeleteHandbell(c *gin.Context) {

	mutexHandbell.Lock()
	defer mutexHandbell.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHandbell", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHandbell.GetDB()

	// Get model if exist
	var handbellDB orm.HandbellDB
	if err := db.First(&handbellDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&handbellDB)

	// get an instance (not staged) from DB instance, and call callback function
	handbellDeleted := new(models.Handbell)
	handbellDB.CopyBasicFieldsToHandbell(handbellDeleted)

	// get stage instance from DB instance, and call callback function
	handbellStaged := backRepo.BackRepoHandbell.Map_HandbellDBID_HandbellPtr[handbellDB.ID]
	if handbellStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), handbellStaged, handbellDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
