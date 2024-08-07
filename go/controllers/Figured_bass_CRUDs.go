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
var __Figured_bass__dummysDeclaration__ models.Figured_bass
var __Figured_bass_time__dummyDeclaration time.Duration

var mutexFigured_bass sync.Mutex

// An Figured_bassID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFigured_bass updateFigured_bass deleteFigured_bass
type Figured_bassID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Figured_bassInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFigured_bass updateFigured_bass
type Figured_bassInput struct {
	// The Figured_bass to submit or modify
	// in: body
	Figured_bass *orm.Figured_bassAPI
}

// GetFigured_basss
//
// swagger:route GET /figured_basss figured_basss getFigured_basss
//
// # Get all figured_basss
//
// Responses:
// default: genericError
//
//	200: figured_bassDBResponse
func (controller *Controller) GetFigured_basss(c *gin.Context) {

	// source slice
	var figured_bassDBs []orm.Figured_bassDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFigured_basss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigured_bass.GetDB()

	query := db.Find(&figured_bassDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	figured_bassAPIs := make([]orm.Figured_bassAPI, 0)

	// for each figured_bass, update fields from the database nullable fields
	for idx := range figured_bassDBs {
		figured_bassDB := &figured_bassDBs[idx]
		_ = figured_bassDB
		var figured_bassAPI orm.Figured_bassAPI

		// insertion point for updating fields
		figured_bassAPI.ID = figured_bassDB.ID
		figured_bassDB.CopyBasicFieldsToFigured_bass_WOP(&figured_bassAPI.Figured_bass_WOP)
		figured_bassAPI.Figured_bassPointersEncoding = figured_bassDB.Figured_bassPointersEncoding
		figured_bassAPIs = append(figured_bassAPIs, figured_bassAPI)
	}

	c.JSON(http.StatusOK, figured_bassAPIs)
}

// PostFigured_bass
//
// swagger:route POST /figured_basss figured_basss postFigured_bass
//
// Creates a figured_bass
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFigured_bass(c *gin.Context) {

	mutexFigured_bass.Lock()
	defer mutexFigured_bass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFigured_basss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigured_bass.GetDB()

	// Validate input
	var input orm.Figured_bassAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create figured_bass
	figured_bassDB := orm.Figured_bassDB{}
	figured_bassDB.Figured_bassPointersEncoding = input.Figured_bassPointersEncoding
	figured_bassDB.CopyBasicFieldsFromFigured_bass_WOP(&input.Figured_bass_WOP)

	query := db.Create(&figured_bassDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFigured_bass.CheckoutPhaseOneInstance(&figured_bassDB)
	figured_bass := backRepo.BackRepoFigured_bass.Map_Figured_bassDBID_Figured_bassPtr[figured_bassDB.ID]

	if figured_bass != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), figured_bass)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, figured_bassDB)
}

// GetFigured_bass
//
// swagger:route GET /figured_basss/{ID} figured_basss getFigured_bass
//
// Gets the details for a figured_bass.
//
// Responses:
// default: genericError
//
//	200: figured_bassDBResponse
func (controller *Controller) GetFigured_bass(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFigured_bass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigured_bass.GetDB()

	// Get figured_bassDB in DB
	var figured_bassDB orm.Figured_bassDB
	if err := db.First(&figured_bassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var figured_bassAPI orm.Figured_bassAPI
	figured_bassAPI.ID = figured_bassDB.ID
	figured_bassAPI.Figured_bassPointersEncoding = figured_bassDB.Figured_bassPointersEncoding
	figured_bassDB.CopyBasicFieldsToFigured_bass_WOP(&figured_bassAPI.Figured_bass_WOP)

	c.JSON(http.StatusOK, figured_bassAPI)
}

// UpdateFigured_bass
//
// swagger:route PATCH /figured_basss/{ID} figured_basss updateFigured_bass
//
// # Update a figured_bass
//
// Responses:
// default: genericError
//
//	200: figured_bassDBResponse
func (controller *Controller) UpdateFigured_bass(c *gin.Context) {

	mutexFigured_bass.Lock()
	defer mutexFigured_bass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFigured_bass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigured_bass.GetDB()

	// Validate input
	var input orm.Figured_bassAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var figured_bassDB orm.Figured_bassDB

	// fetch the figured_bass
	query := db.First(&figured_bassDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	figured_bassDB.CopyBasicFieldsFromFigured_bass_WOP(&input.Figured_bass_WOP)
	figured_bassDB.Figured_bassPointersEncoding = input.Figured_bassPointersEncoding

	query = db.Model(&figured_bassDB).Updates(figured_bassDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	figured_bassNew := new(models.Figured_bass)
	figured_bassDB.CopyBasicFieldsToFigured_bass(figured_bassNew)

	// redeem pointers
	figured_bassDB.DecodePointers(backRepo, figured_bassNew)

	// get stage instance from DB instance, and call callback function
	figured_bassOld := backRepo.BackRepoFigured_bass.Map_Figured_bassDBID_Figured_bassPtr[figured_bassDB.ID]
	if figured_bassOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), figured_bassOld, figured_bassNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the figured_bassDB
	c.JSON(http.StatusOK, figured_bassDB)
}

// DeleteFigured_bass
//
// swagger:route DELETE /figured_basss/{ID} figured_basss deleteFigured_bass
//
// # Delete a figured_bass
//
// default: genericError
//
//	200: figured_bassDBResponse
func (controller *Controller) DeleteFigured_bass(c *gin.Context) {

	mutexFigured_bass.Lock()
	defer mutexFigured_bass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFigured_bass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFigured_bass.GetDB()

	// Get model if exist
	var figured_bassDB orm.Figured_bassDB
	if err := db.First(&figured_bassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&figured_bassDB)

	// get an instance (not staged) from DB instance, and call callback function
	figured_bassDeleted := new(models.Figured_bass)
	figured_bassDB.CopyBasicFieldsToFigured_bass(figured_bassDeleted)

	// get stage instance from DB instance, and call callback function
	figured_bassStaged := backRepo.BackRepoFigured_bass.Map_Figured_bassDBID_Figured_bassPtr[figured_bassDB.ID]
	if figured_bassStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), figured_bassStaged, figured_bassDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
