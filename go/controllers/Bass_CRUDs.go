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
var __Bass__dummysDeclaration__ models.Bass
var __Bass_time__dummyDeclaration time.Duration

var mutexBass sync.Mutex

// An BassID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBass updateBass deleteBass
type BassID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BassInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBass updateBass
type BassInput struct {
	// The Bass to submit or modify
	// in: body
	Bass *orm.BassAPI
}

// GetBasss
//
// swagger:route GET /basss basss getBasss
//
// # Get all basss
//
// Responses:
// default: genericError
//
//	200: bassDBResponse
func (controller *Controller) GetBasss(c *gin.Context) {

	// source slice
	var bassDBs []orm.BassDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBasss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass.GetDB()

	query := db.Find(&bassDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bassAPIs := make([]orm.BassAPI, 0)

	// for each bass, update fields from the database nullable fields
	for idx := range bassDBs {
		bassDB := &bassDBs[idx]
		_ = bassDB
		var bassAPI orm.BassAPI

		// insertion point for updating fields
		bassAPI.ID = bassDB.ID
		bassDB.CopyBasicFieldsToBass_WOP(&bassAPI.Bass_WOP)
		bassAPI.BassPointersEncoding = bassDB.BassPointersEncoding
		bassAPIs = append(bassAPIs, bassAPI)
	}

	c.JSON(http.StatusOK, bassAPIs)
}

// PostBass
//
// swagger:route POST /basss basss postBass
//
// Creates a bass
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBass(c *gin.Context) {

	mutexBass.Lock()
	defer mutexBass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBasss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass.GetDB()

	// Validate input
	var input orm.BassAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bass
	bassDB := orm.BassDB{}
	bassDB.BassPointersEncoding = input.BassPointersEncoding
	bassDB.CopyBasicFieldsFromBass_WOP(&input.Bass_WOP)

	query := db.Create(&bassDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBass.CheckoutPhaseOneInstance(&bassDB)
	bass := backRepo.BackRepoBass.Map_BassDBID_BassPtr[bassDB.ID]

	if bass != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bass)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bassDB)
}

// GetBass
//
// swagger:route GET /basss/{ID} basss getBass
//
// Gets the details for a bass.
//
// Responses:
// default: genericError
//
//	200: bassDBResponse
func (controller *Controller) GetBass(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass.GetDB()

	// Get bassDB in DB
	var bassDB orm.BassDB
	if err := db.First(&bassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bassAPI orm.BassAPI
	bassAPI.ID = bassDB.ID
	bassAPI.BassPointersEncoding = bassDB.BassPointersEncoding
	bassDB.CopyBasicFieldsToBass_WOP(&bassAPI.Bass_WOP)

	c.JSON(http.StatusOK, bassAPI)
}

// UpdateBass
//
// swagger:route PATCH /basss/{ID} basss updateBass
//
// # Update a bass
//
// Responses:
// default: genericError
//
//	200: bassDBResponse
func (controller *Controller) UpdateBass(c *gin.Context) {

	mutexBass.Lock()
	defer mutexBass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass.GetDB()

	// Validate input
	var input orm.BassAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bassDB orm.BassDB

	// fetch the bass
	query := db.First(&bassDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bassDB.CopyBasicFieldsFromBass_WOP(&input.Bass_WOP)
	bassDB.BassPointersEncoding = input.BassPointersEncoding

	query = db.Model(&bassDB).Updates(bassDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bassNew := new(models.Bass)
	bassDB.CopyBasicFieldsToBass(bassNew)

	// redeem pointers
	bassDB.DecodePointers(backRepo, bassNew)

	// get stage instance from DB instance, and call callback function
	bassOld := backRepo.BackRepoBass.Map_BassDBID_BassPtr[bassDB.ID]
	if bassOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bassOld, bassNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bassDB
	c.JSON(http.StatusOK, bassDB)
}

// DeleteBass
//
// swagger:route DELETE /basss/{ID} basss deleteBass
//
// # Delete a bass
//
// default: genericError
//
//	200: bassDBResponse
func (controller *Controller) DeleteBass(c *gin.Context) {

	mutexBass.Lock()
	defer mutexBass.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBass", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBass.GetDB()

	// Get model if exist
	var bassDB orm.BassDB
	if err := db.First(&bassDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bassDB)

	// get an instance (not staged) from DB instance, and call callback function
	bassDeleted := new(models.Bass)
	bassDB.CopyBasicFieldsToBass(bassDeleted)

	// get stage instance from DB instance, and call callback function
	bassStaged := backRepo.BackRepoBass.Map_BassDBID_BassPtr[bassDB.ID]
	if bassStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bassStaged, bassDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
