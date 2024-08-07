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
var __Fingering__dummysDeclaration__ models.Fingering
var __Fingering_time__dummyDeclaration time.Duration

var mutexFingering sync.Mutex

// An FingeringID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFingering updateFingering deleteFingering
type FingeringID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FingeringInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFingering updateFingering
type FingeringInput struct {
	// The Fingering to submit or modify
	// in: body
	Fingering *orm.FingeringAPI
}

// GetFingerings
//
// swagger:route GET /fingerings fingerings getFingerings
//
// # Get all fingerings
//
// Responses:
// default: genericError
//
//	200: fingeringDBResponse
func (controller *Controller) GetFingerings(c *gin.Context) {

	// source slice
	var fingeringDBs []orm.FingeringDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFingerings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFingering.GetDB()

	query := db.Find(&fingeringDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fingeringAPIs := make([]orm.FingeringAPI, 0)

	// for each fingering, update fields from the database nullable fields
	for idx := range fingeringDBs {
		fingeringDB := &fingeringDBs[idx]
		_ = fingeringDB
		var fingeringAPI orm.FingeringAPI

		// insertion point for updating fields
		fingeringAPI.ID = fingeringDB.ID
		fingeringDB.CopyBasicFieldsToFingering_WOP(&fingeringAPI.Fingering_WOP)
		fingeringAPI.FingeringPointersEncoding = fingeringDB.FingeringPointersEncoding
		fingeringAPIs = append(fingeringAPIs, fingeringAPI)
	}

	c.JSON(http.StatusOK, fingeringAPIs)
}

// PostFingering
//
// swagger:route POST /fingerings fingerings postFingering
//
// Creates a fingering
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFingering(c *gin.Context) {

	mutexFingering.Lock()
	defer mutexFingering.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFingerings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFingering.GetDB()

	// Validate input
	var input orm.FingeringAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create fingering
	fingeringDB := orm.FingeringDB{}
	fingeringDB.FingeringPointersEncoding = input.FingeringPointersEncoding
	fingeringDB.CopyBasicFieldsFromFingering_WOP(&input.Fingering_WOP)

	query := db.Create(&fingeringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFingering.CheckoutPhaseOneInstance(&fingeringDB)
	fingering := backRepo.BackRepoFingering.Map_FingeringDBID_FingeringPtr[fingeringDB.ID]

	if fingering != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), fingering)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fingeringDB)
}

// GetFingering
//
// swagger:route GET /fingerings/{ID} fingerings getFingering
//
// Gets the details for a fingering.
//
// Responses:
// default: genericError
//
//	200: fingeringDBResponse
func (controller *Controller) GetFingering(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFingering", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFingering.GetDB()

	// Get fingeringDB in DB
	var fingeringDB orm.FingeringDB
	if err := db.First(&fingeringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fingeringAPI orm.FingeringAPI
	fingeringAPI.ID = fingeringDB.ID
	fingeringAPI.FingeringPointersEncoding = fingeringDB.FingeringPointersEncoding
	fingeringDB.CopyBasicFieldsToFingering_WOP(&fingeringAPI.Fingering_WOP)

	c.JSON(http.StatusOK, fingeringAPI)
}

// UpdateFingering
//
// swagger:route PATCH /fingerings/{ID} fingerings updateFingering
//
// # Update a fingering
//
// Responses:
// default: genericError
//
//	200: fingeringDBResponse
func (controller *Controller) UpdateFingering(c *gin.Context) {

	mutexFingering.Lock()
	defer mutexFingering.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFingering", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFingering.GetDB()

	// Validate input
	var input orm.FingeringAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fingeringDB orm.FingeringDB

	// fetch the fingering
	query := db.First(&fingeringDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fingeringDB.CopyBasicFieldsFromFingering_WOP(&input.Fingering_WOP)
	fingeringDB.FingeringPointersEncoding = input.FingeringPointersEncoding

	query = db.Model(&fingeringDB).Updates(fingeringDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fingeringNew := new(models.Fingering)
	fingeringDB.CopyBasicFieldsToFingering(fingeringNew)

	// redeem pointers
	fingeringDB.DecodePointers(backRepo, fingeringNew)

	// get stage instance from DB instance, and call callback function
	fingeringOld := backRepo.BackRepoFingering.Map_FingeringDBID_FingeringPtr[fingeringDB.ID]
	if fingeringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fingeringOld, fingeringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fingeringDB
	c.JSON(http.StatusOK, fingeringDB)
}

// DeleteFingering
//
// swagger:route DELETE /fingerings/{ID} fingerings deleteFingering
//
// # Delete a fingering
//
// default: genericError
//
//	200: fingeringDBResponse
func (controller *Controller) DeleteFingering(c *gin.Context) {

	mutexFingering.Lock()
	defer mutexFingering.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFingering", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFingering.GetDB()

	// Get model if exist
	var fingeringDB orm.FingeringDB
	if err := db.First(&fingeringDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&fingeringDB)

	// get an instance (not staged) from DB instance, and call callback function
	fingeringDeleted := new(models.Fingering)
	fingeringDB.CopyBasicFieldsToFingering(fingeringDeleted)

	// get stage instance from DB instance, and call callback function
	fingeringStaged := backRepo.BackRepoFingering.Map_FingeringDBID_FingeringPtr[fingeringDB.ID]
	if fingeringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fingeringStaged, fingeringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
