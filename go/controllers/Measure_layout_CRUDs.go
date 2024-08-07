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
var __Measure_layout__dummysDeclaration__ models.Measure_layout
var __Measure_layout_time__dummyDeclaration time.Duration

var mutexMeasure_layout sync.Mutex

// An Measure_layoutID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMeasure_layout updateMeasure_layout deleteMeasure_layout
type Measure_layoutID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Measure_layoutInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMeasure_layout updateMeasure_layout
type Measure_layoutInput struct {
	// The Measure_layout to submit or modify
	// in: body
	Measure_layout *orm.Measure_layoutAPI
}

// GetMeasure_layouts
//
// swagger:route GET /measure_layouts measure_layouts getMeasure_layouts
//
// # Get all measure_layouts
//
// Responses:
// default: genericError
//
//	200: measure_layoutDBResponse
func (controller *Controller) GetMeasure_layouts(c *gin.Context) {

	// source slice
	var measure_layoutDBs []orm.Measure_layoutDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_layout.GetDB()

	query := db.Find(&measure_layoutDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	measure_layoutAPIs := make([]orm.Measure_layoutAPI, 0)

	// for each measure_layout, update fields from the database nullable fields
	for idx := range measure_layoutDBs {
		measure_layoutDB := &measure_layoutDBs[idx]
		_ = measure_layoutDB
		var measure_layoutAPI orm.Measure_layoutAPI

		// insertion point for updating fields
		measure_layoutAPI.ID = measure_layoutDB.ID
		measure_layoutDB.CopyBasicFieldsToMeasure_layout_WOP(&measure_layoutAPI.Measure_layout_WOP)
		measure_layoutAPI.Measure_layoutPointersEncoding = measure_layoutDB.Measure_layoutPointersEncoding
		measure_layoutAPIs = append(measure_layoutAPIs, measure_layoutAPI)
	}

	c.JSON(http.StatusOK, measure_layoutAPIs)
}

// PostMeasure_layout
//
// swagger:route POST /measure_layouts measure_layouts postMeasure_layout
//
// Creates a measure_layout
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMeasure_layout(c *gin.Context) {

	mutexMeasure_layout.Lock()
	defer mutexMeasure_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMeasure_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_layout.GetDB()

	// Validate input
	var input orm.Measure_layoutAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create measure_layout
	measure_layoutDB := orm.Measure_layoutDB{}
	measure_layoutDB.Measure_layoutPointersEncoding = input.Measure_layoutPointersEncoding
	measure_layoutDB.CopyBasicFieldsFromMeasure_layout_WOP(&input.Measure_layout_WOP)

	query := db.Create(&measure_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMeasure_layout.CheckoutPhaseOneInstance(&measure_layoutDB)
	measure_layout := backRepo.BackRepoMeasure_layout.Map_Measure_layoutDBID_Measure_layoutPtr[measure_layoutDB.ID]

	if measure_layout != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), measure_layout)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, measure_layoutDB)
}

// GetMeasure_layout
//
// swagger:route GET /measure_layouts/{ID} measure_layouts getMeasure_layout
//
// Gets the details for a measure_layout.
//
// Responses:
// default: genericError
//
//	200: measure_layoutDBResponse
func (controller *Controller) GetMeasure_layout(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMeasure_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_layout.GetDB()

	// Get measure_layoutDB in DB
	var measure_layoutDB orm.Measure_layoutDB
	if err := db.First(&measure_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var measure_layoutAPI orm.Measure_layoutAPI
	measure_layoutAPI.ID = measure_layoutDB.ID
	measure_layoutAPI.Measure_layoutPointersEncoding = measure_layoutDB.Measure_layoutPointersEncoding
	measure_layoutDB.CopyBasicFieldsToMeasure_layout_WOP(&measure_layoutAPI.Measure_layout_WOP)

	c.JSON(http.StatusOK, measure_layoutAPI)
}

// UpdateMeasure_layout
//
// swagger:route PATCH /measure_layouts/{ID} measure_layouts updateMeasure_layout
//
// # Update a measure_layout
//
// Responses:
// default: genericError
//
//	200: measure_layoutDBResponse
func (controller *Controller) UpdateMeasure_layout(c *gin.Context) {

	mutexMeasure_layout.Lock()
	defer mutexMeasure_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMeasure_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_layout.GetDB()

	// Validate input
	var input orm.Measure_layoutAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var measure_layoutDB orm.Measure_layoutDB

	// fetch the measure_layout
	query := db.First(&measure_layoutDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	measure_layoutDB.CopyBasicFieldsFromMeasure_layout_WOP(&input.Measure_layout_WOP)
	measure_layoutDB.Measure_layoutPointersEncoding = input.Measure_layoutPointersEncoding

	query = db.Model(&measure_layoutDB).Updates(measure_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	measure_layoutNew := new(models.Measure_layout)
	measure_layoutDB.CopyBasicFieldsToMeasure_layout(measure_layoutNew)

	// redeem pointers
	measure_layoutDB.DecodePointers(backRepo, measure_layoutNew)

	// get stage instance from DB instance, and call callback function
	measure_layoutOld := backRepo.BackRepoMeasure_layout.Map_Measure_layoutDBID_Measure_layoutPtr[measure_layoutDB.ID]
	if measure_layoutOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), measure_layoutOld, measure_layoutNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the measure_layoutDB
	c.JSON(http.StatusOK, measure_layoutDB)
}

// DeleteMeasure_layout
//
// swagger:route DELETE /measure_layouts/{ID} measure_layouts deleteMeasure_layout
//
// # Delete a measure_layout
//
// default: genericError
//
//	200: measure_layoutDBResponse
func (controller *Controller) DeleteMeasure_layout(c *gin.Context) {

	mutexMeasure_layout.Lock()
	defer mutexMeasure_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMeasure_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMeasure_layout.GetDB()

	// Get model if exist
	var measure_layoutDB orm.Measure_layoutDB
	if err := db.First(&measure_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&measure_layoutDB)

	// get an instance (not staged) from DB instance, and call callback function
	measure_layoutDeleted := new(models.Measure_layout)
	measure_layoutDB.CopyBasicFieldsToMeasure_layout(measure_layoutDeleted)

	// get stage instance from DB instance, and call callback function
	measure_layoutStaged := backRepo.BackRepoMeasure_layout.Map_Measure_layoutDBID_Measure_layoutPtr[measure_layoutDB.ID]
	if measure_layoutStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), measure_layoutStaged, measure_layoutDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
