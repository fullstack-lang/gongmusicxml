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
var __Measure_style__dummysDeclaration__ models.Measure_style
var __Measure_style_time__dummyDeclaration time.Duration

var mutexMeasure_style sync.Mutex

// An Measure_styleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeasure_style updateMeasure_style deleteMeasure_style
type Measure_styleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Measure_styleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeasure_style updateMeasure_style
type Measure_styleInput struct {
	// The Measure_style to submit or modify
	// in: body
	Measure_style *orm.Measure_styleAPI
}

// GetMeasure_styles
//
// swagger:route GET /measure_styles measure_styles getMeasure_styles
//
// # Get all measure_styles
//
// Responses:
// default: genericError
//
//	200: measure_styleDBResponse
func (controller *Controller) GetMeasure_styles(c *gin.Context) {

	// source slice
	var measure_styleDBs []orm.Measure_styleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_styles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_style.GetDB()

	query := db.Find(&measure_styleDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	measure_styleAPIs := make([]orm.Measure_styleAPI, 0)

	// for each measure_style, update fields from the database nullable fields
	for idx := range measure_styleDBs {
		measure_styleDB := &measure_styleDBs[idx]
		_ = measure_styleDB
		var measure_styleAPI orm.Measure_styleAPI

		// insertion point for updating fields
		measure_styleAPI.ID = measure_styleDB.ID
		measure_styleDB.CopyBasicFieldsToMeasure_style_WOP(&measure_styleAPI.Measure_style_WOP)
		measure_styleAPI.Measure_stylePointersEncoding = measure_styleDB.Measure_stylePointersEncoding
		measure_styleAPIs = append(measure_styleAPIs, measure_styleAPI)
	}

	c.JSON(http.StatusOK, measure_styleAPIs)
}

// PostMeasure_style
//
// swagger:route POST /measure_styles measure_styles postMeasure_style
//
// Creates a measure_style
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeasure_style(c *gin.Context) {

	mutexMeasure_style.Lock()
	defer mutexMeasure_style.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeasure_styles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_style.GetDB()

	// Validate input
	var input orm.Measure_styleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create measure_style
	measure_styleDB := orm.Measure_styleDB{}
	measure_styleDB.Measure_stylePointersEncoding = input.Measure_stylePointersEncoding
	measure_styleDB.CopyBasicFieldsFromMeasure_style_WOP(&input.Measure_style_WOP)

	query := db.Create(&measure_styleDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeasure_style.CheckoutPhaseOneInstance(&measure_styleDB)
	measure_style := backRepo.BackRepoMeasure_style.Map_Measure_styleDBID_Measure_stylePtr[measure_styleDB.ID]

	if measure_style != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), measure_style)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, measure_styleDB)
}

// GetMeasure_style
//
// swagger:route GET /measure_styles/{ID} measure_styles getMeasure_style
//
// Gets the details for a measure_style.
//
// Responses:
// default: genericError
//
//	200: measure_styleDBResponse
func (controller *Controller) GetMeasure_style(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_style", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_style.GetDB()

	// Get measure_styleDB in DB
	var measure_styleDB orm.Measure_styleDB
	if err := db.First(&measure_styleDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var measure_styleAPI orm.Measure_styleAPI
	measure_styleAPI.ID = measure_styleDB.ID
	measure_styleAPI.Measure_stylePointersEncoding = measure_styleDB.Measure_stylePointersEncoding
	measure_styleDB.CopyBasicFieldsToMeasure_style_WOP(&measure_styleAPI.Measure_style_WOP)

	c.JSON(http.StatusOK, measure_styleAPI)
}

// UpdateMeasure_style
//
// swagger:route PATCH /measure_styles/{ID} measure_styles updateMeasure_style
//
// # Update a measure_style
//
// Responses:
// default: genericError
//
//	200: measure_styleDBResponse
func (controller *Controller) UpdateMeasure_style(c *gin.Context) {

	mutexMeasure_style.Lock()
	defer mutexMeasure_style.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMeasure_style", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_style.GetDB()

	// Validate input
	var input orm.Measure_styleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var measure_styleDB orm.Measure_styleDB

	// fetch the measure_style
	query := db.First(&measure_styleDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	measure_styleDB.CopyBasicFieldsFromMeasure_style_WOP(&input.Measure_style_WOP)
	measure_styleDB.Measure_stylePointersEncoding = input.Measure_stylePointersEncoding

	query = db.Model(&measure_styleDB).Updates(measure_styleDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	measure_styleNew := new(models.Measure_style)
	measure_styleDB.CopyBasicFieldsToMeasure_style(measure_styleNew)

	// redeem pointers
	measure_styleDB.DecodePointers(backRepo, measure_styleNew)

	// get stage instance from DB instance, and call callback function
	measure_styleOld := backRepo.BackRepoMeasure_style.Map_Measure_styleDBID_Measure_stylePtr[measure_styleDB.ID]
	if measure_styleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), measure_styleOld, measure_styleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the measure_styleDB
	c.JSON(http.StatusOK, measure_styleDB)
}

// DeleteMeasure_style
//
// swagger:route DELETE /measure_styles/{ID} measure_styles deleteMeasure_style
//
// # Delete a measure_style
//
// default: genericError
//
//	200: measure_styleDBResponse
func (controller *Controller) DeleteMeasure_style(c *gin.Context) {

	mutexMeasure_style.Lock()
	defer mutexMeasure_style.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeasure_style", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_style.GetDB()

	// Get model if exist
	var measure_styleDB orm.Measure_styleDB
	if err := db.First(&measure_styleDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&measure_styleDB)

	// get an instance (not staged) from DB instance, and call callback function
	measure_styleDeleted := new(models.Measure_style)
	measure_styleDB.CopyBasicFieldsToMeasure_style(measure_styleDeleted)

	// get stage instance from DB instance, and call callback function
	measure_styleStaged := backRepo.BackRepoMeasure_style.Map_Measure_styleDBID_Measure_stylePtr[measure_styleDB.ID]
	if measure_styleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), measure_styleStaged, measure_styleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
