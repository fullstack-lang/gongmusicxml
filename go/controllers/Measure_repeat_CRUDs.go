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
var __Measure_repeat__dummysDeclaration__ models.Measure_repeat
var __Measure_repeat_time__dummyDeclaration time.Duration

var mutexMeasure_repeat sync.Mutex

// An Measure_repeatID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeasure_repeat updateMeasure_repeat deleteMeasure_repeat
type Measure_repeatID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Measure_repeatInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeasure_repeat updateMeasure_repeat
type Measure_repeatInput struct {
	// The Measure_repeat to submit or modify
	// in: body
	Measure_repeat *orm.Measure_repeatAPI
}

// GetMeasure_repeats
//
// swagger:route GET /measure_repeats measure_repeats getMeasure_repeats
//
// # Get all measure_repeats
//
// Responses:
// default: genericError
//
//	200: measure_repeatDBResponse
func (controller *Controller) GetMeasure_repeats(c *gin.Context) {

	// source slice
	var measure_repeatDBs []orm.Measure_repeatDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_repeats", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_repeat.GetDB()

	query := db.Find(&measure_repeatDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	measure_repeatAPIs := make([]orm.Measure_repeatAPI, 0)

	// for each measure_repeat, update fields from the database nullable fields
	for idx := range measure_repeatDBs {
		measure_repeatDB := &measure_repeatDBs[idx]
		_ = measure_repeatDB
		var measure_repeatAPI orm.Measure_repeatAPI

		// insertion point for updating fields
		measure_repeatAPI.ID = measure_repeatDB.ID
		measure_repeatDB.CopyBasicFieldsToMeasure_repeat_WOP(&measure_repeatAPI.Measure_repeat_WOP)
		measure_repeatAPI.Measure_repeatPointersEncoding = measure_repeatDB.Measure_repeatPointersEncoding
		measure_repeatAPIs = append(measure_repeatAPIs, measure_repeatAPI)
	}

	c.JSON(http.StatusOK, measure_repeatAPIs)
}

// PostMeasure_repeat
//
// swagger:route POST /measure_repeats measure_repeats postMeasure_repeat
//
// Creates a measure_repeat
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeasure_repeat(c *gin.Context) {

	mutexMeasure_repeat.Lock()
	defer mutexMeasure_repeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeasure_repeats", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_repeat.GetDB()

	// Validate input
	var input orm.Measure_repeatAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create measure_repeat
	measure_repeatDB := orm.Measure_repeatDB{}
	measure_repeatDB.Measure_repeatPointersEncoding = input.Measure_repeatPointersEncoding
	measure_repeatDB.CopyBasicFieldsFromMeasure_repeat_WOP(&input.Measure_repeat_WOP)

	query := db.Create(&measure_repeatDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeasure_repeat.CheckoutPhaseOneInstance(&measure_repeatDB)
	measure_repeat := backRepo.BackRepoMeasure_repeat.Map_Measure_repeatDBID_Measure_repeatPtr[measure_repeatDB.ID]

	if measure_repeat != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), measure_repeat)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, measure_repeatDB)
}

// GetMeasure_repeat
//
// swagger:route GET /measure_repeats/{ID} measure_repeats getMeasure_repeat
//
// Gets the details for a measure_repeat.
//
// Responses:
// default: genericError
//
//	200: measure_repeatDBResponse
func (controller *Controller) GetMeasure_repeat(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_repeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_repeat.GetDB()

	// Get measure_repeatDB in DB
	var measure_repeatDB orm.Measure_repeatDB
	if err := db.First(&measure_repeatDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var measure_repeatAPI orm.Measure_repeatAPI
	measure_repeatAPI.ID = measure_repeatDB.ID
	measure_repeatAPI.Measure_repeatPointersEncoding = measure_repeatDB.Measure_repeatPointersEncoding
	measure_repeatDB.CopyBasicFieldsToMeasure_repeat_WOP(&measure_repeatAPI.Measure_repeat_WOP)

	c.JSON(http.StatusOK, measure_repeatAPI)
}

// UpdateMeasure_repeat
//
// swagger:route PATCH /measure_repeats/{ID} measure_repeats updateMeasure_repeat
//
// # Update a measure_repeat
//
// Responses:
// default: genericError
//
//	200: measure_repeatDBResponse
func (controller *Controller) UpdateMeasure_repeat(c *gin.Context) {

	mutexMeasure_repeat.Lock()
	defer mutexMeasure_repeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMeasure_repeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_repeat.GetDB()

	// Validate input
	var input orm.Measure_repeatAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var measure_repeatDB orm.Measure_repeatDB

	// fetch the measure_repeat
	query := db.First(&measure_repeatDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	measure_repeatDB.CopyBasicFieldsFromMeasure_repeat_WOP(&input.Measure_repeat_WOP)
	measure_repeatDB.Measure_repeatPointersEncoding = input.Measure_repeatPointersEncoding

	query = db.Model(&measure_repeatDB).Updates(measure_repeatDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	measure_repeatNew := new(models.Measure_repeat)
	measure_repeatDB.CopyBasicFieldsToMeasure_repeat(measure_repeatNew)

	// redeem pointers
	measure_repeatDB.DecodePointers(backRepo, measure_repeatNew)

	// get stage instance from DB instance, and call callback function
	measure_repeatOld := backRepo.BackRepoMeasure_repeat.Map_Measure_repeatDBID_Measure_repeatPtr[measure_repeatDB.ID]
	if measure_repeatOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), measure_repeatOld, measure_repeatNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the measure_repeatDB
	c.JSON(http.StatusOK, measure_repeatDB)
}

// DeleteMeasure_repeat
//
// swagger:route DELETE /measure_repeats/{ID} measure_repeats deleteMeasure_repeat
//
// # Delete a measure_repeat
//
// default: genericError
//
//	200: measure_repeatDBResponse
func (controller *Controller) DeleteMeasure_repeat(c *gin.Context) {

	mutexMeasure_repeat.Lock()
	defer mutexMeasure_repeat.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeasure_repeat", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_repeat.GetDB()

	// Get model if exist
	var measure_repeatDB orm.Measure_repeatDB
	if err := db.First(&measure_repeatDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&measure_repeatDB)

	// get an instance (not staged) from DB instance, and call callback function
	measure_repeatDeleted := new(models.Measure_repeat)
	measure_repeatDB.CopyBasicFieldsToMeasure_repeat(measure_repeatDeleted)

	// get stage instance from DB instance, and call callback function
	measure_repeatStaged := backRepo.BackRepoMeasure_repeat.Map_Measure_repeatDBID_Measure_repeatPtr[measure_repeatDB.ID]
	if measure_repeatStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), measure_repeatStaged, measure_repeatDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
