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
var __Measure_numbering__dummysDeclaration__ models.Measure_numbering
var __Measure_numbering_time__dummyDeclaration time.Duration

var mutexMeasure_numbering sync.Mutex

// An Measure_numberingID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeasure_numbering updateMeasure_numbering deleteMeasure_numbering
type Measure_numberingID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Measure_numberingInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeasure_numbering updateMeasure_numbering
type Measure_numberingInput struct {
	// The Measure_numbering to submit or modify
	// in: body
	Measure_numbering *orm.Measure_numberingAPI
}

// GetMeasure_numberings
//
// swagger:route GET /measure_numberings measure_numberings getMeasure_numberings
//
// # Get all measure_numberings
//
// Responses:
// default: genericError
//
//	200: measure_numberingDBResponse
func (controller *Controller) GetMeasure_numberings(c *gin.Context) {

	// source slice
	var measure_numberingDBs []orm.Measure_numberingDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_numberings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_numbering.GetDB()

	query := db.Find(&measure_numberingDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	measure_numberingAPIs := make([]orm.Measure_numberingAPI, 0)

	// for each measure_numbering, update fields from the database nullable fields
	for idx := range measure_numberingDBs {
		measure_numberingDB := &measure_numberingDBs[idx]
		_ = measure_numberingDB
		var measure_numberingAPI orm.Measure_numberingAPI

		// insertion point for updating fields
		measure_numberingAPI.ID = measure_numberingDB.ID
		measure_numberingDB.CopyBasicFieldsToMeasure_numbering_WOP(&measure_numberingAPI.Measure_numbering_WOP)
		measure_numberingAPI.Measure_numberingPointersEncoding = measure_numberingDB.Measure_numberingPointersEncoding
		measure_numberingAPIs = append(measure_numberingAPIs, measure_numberingAPI)
	}

	c.JSON(http.StatusOK, measure_numberingAPIs)
}

// PostMeasure_numbering
//
// swagger:route POST /measure_numberings measure_numberings postMeasure_numbering
//
// Creates a measure_numbering
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeasure_numbering(c *gin.Context) {

	mutexMeasure_numbering.Lock()
	defer mutexMeasure_numbering.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeasure_numberings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_numbering.GetDB()

	// Validate input
	var input orm.Measure_numberingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create measure_numbering
	measure_numberingDB := orm.Measure_numberingDB{}
	measure_numberingDB.Measure_numberingPointersEncoding = input.Measure_numberingPointersEncoding
	measure_numberingDB.CopyBasicFieldsFromMeasure_numbering_WOP(&input.Measure_numbering_WOP)

	query := db.Create(&measure_numberingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeasure_numbering.CheckoutPhaseOneInstance(&measure_numberingDB)
	measure_numbering := backRepo.BackRepoMeasure_numbering.Map_Measure_numberingDBID_Measure_numberingPtr[measure_numberingDB.ID]

	if measure_numbering != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), measure_numbering)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, measure_numberingDB)
}

// GetMeasure_numbering
//
// swagger:route GET /measure_numberings/{ID} measure_numberings getMeasure_numbering
//
// Gets the details for a measure_numbering.
//
// Responses:
// default: genericError
//
//	200: measure_numberingDBResponse
func (controller *Controller) GetMeasure_numbering(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_numbering", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_numbering.GetDB()

	// Get measure_numberingDB in DB
	var measure_numberingDB orm.Measure_numberingDB
	if err := db.First(&measure_numberingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var measure_numberingAPI orm.Measure_numberingAPI
	measure_numberingAPI.ID = measure_numberingDB.ID
	measure_numberingAPI.Measure_numberingPointersEncoding = measure_numberingDB.Measure_numberingPointersEncoding
	measure_numberingDB.CopyBasicFieldsToMeasure_numbering_WOP(&measure_numberingAPI.Measure_numbering_WOP)

	c.JSON(http.StatusOK, measure_numberingAPI)
}

// UpdateMeasure_numbering
//
// swagger:route PATCH /measure_numberings/{ID} measure_numberings updateMeasure_numbering
//
// # Update a measure_numbering
//
// Responses:
// default: genericError
//
//	200: measure_numberingDBResponse
func (controller *Controller) UpdateMeasure_numbering(c *gin.Context) {

	mutexMeasure_numbering.Lock()
	defer mutexMeasure_numbering.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMeasure_numbering", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_numbering.GetDB()

	// Validate input
	var input orm.Measure_numberingAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var measure_numberingDB orm.Measure_numberingDB

	// fetch the measure_numbering
	query := db.First(&measure_numberingDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	measure_numberingDB.CopyBasicFieldsFromMeasure_numbering_WOP(&input.Measure_numbering_WOP)
	measure_numberingDB.Measure_numberingPointersEncoding = input.Measure_numberingPointersEncoding

	query = db.Model(&measure_numberingDB).Updates(measure_numberingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	measure_numberingNew := new(models.Measure_numbering)
	measure_numberingDB.CopyBasicFieldsToMeasure_numbering(measure_numberingNew)

	// redeem pointers
	measure_numberingDB.DecodePointers(backRepo, measure_numberingNew)

	// get stage instance from DB instance, and call callback function
	measure_numberingOld := backRepo.BackRepoMeasure_numbering.Map_Measure_numberingDBID_Measure_numberingPtr[measure_numberingDB.ID]
	if measure_numberingOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), measure_numberingOld, measure_numberingNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the measure_numberingDB
	c.JSON(http.StatusOK, measure_numberingDB)
}

// DeleteMeasure_numbering
//
// swagger:route DELETE /measure_numberings/{ID} measure_numberings deleteMeasure_numbering
//
// # Delete a measure_numbering
//
// default: genericError
//
//	200: measure_numberingDBResponse
func (controller *Controller) DeleteMeasure_numbering(c *gin.Context) {

	mutexMeasure_numbering.Lock()
	defer mutexMeasure_numbering.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeasure_numbering", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_numbering.GetDB()

	// Get model if exist
	var measure_numberingDB orm.Measure_numberingDB
	if err := db.First(&measure_numberingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&measure_numberingDB)

	// get an instance (not staged) from DB instance, and call callback function
	measure_numberingDeleted := new(models.Measure_numbering)
	measure_numberingDB.CopyBasicFieldsToMeasure_numbering(measure_numberingDeleted)

	// get stage instance from DB instance, and call callback function
	measure_numberingStaged := backRepo.BackRepoMeasure_numbering.Map_Measure_numberingDBID_Measure_numberingPtr[measure_numberingDB.ID]
	if measure_numberingStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), measure_numberingStaged, measure_numberingDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
