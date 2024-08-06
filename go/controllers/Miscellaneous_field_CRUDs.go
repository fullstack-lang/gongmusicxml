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
var __Miscellaneous_field__dummysDeclaration__ models.Miscellaneous_field
var __Miscellaneous_field_time__dummyDeclaration time.Duration

var mutexMiscellaneous_field sync.Mutex

// An Miscellaneous_fieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMiscellaneous_field updateMiscellaneous_field deleteMiscellaneous_field
type Miscellaneous_fieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Miscellaneous_fieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMiscellaneous_field updateMiscellaneous_field
type Miscellaneous_fieldInput struct {
	// The Miscellaneous_field to submit or modify
	// in: body
	Miscellaneous_field *orm.Miscellaneous_fieldAPI
}

// GetMiscellaneous_fields
//
// swagger:route GET /miscellaneous_fields miscellaneous_fields getMiscellaneous_fields
//
// # Get all miscellaneous_fields
//
// Responses:
// default: genericError
//
//	200: miscellaneous_fieldDBResponse
func (controller *Controller) GetMiscellaneous_fields(c *gin.Context) {

	// source slice
	var miscellaneous_fieldDBs []orm.Miscellaneous_fieldDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMiscellaneous_fields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous_field.GetDB()

	query := db.Find(&miscellaneous_fieldDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	miscellaneous_fieldAPIs := make([]orm.Miscellaneous_fieldAPI, 0)

	// for each miscellaneous_field, update fields from the database nullable fields
	for idx := range miscellaneous_fieldDBs {
		miscellaneous_fieldDB := &miscellaneous_fieldDBs[idx]
		_ = miscellaneous_fieldDB
		var miscellaneous_fieldAPI orm.Miscellaneous_fieldAPI

		// insertion point for updating fields
		miscellaneous_fieldAPI.ID = miscellaneous_fieldDB.ID
		miscellaneous_fieldDB.CopyBasicFieldsToMiscellaneous_field_WOP(&miscellaneous_fieldAPI.Miscellaneous_field_WOP)
		miscellaneous_fieldAPI.Miscellaneous_fieldPointersEncoding = miscellaneous_fieldDB.Miscellaneous_fieldPointersEncoding
		miscellaneous_fieldAPIs = append(miscellaneous_fieldAPIs, miscellaneous_fieldAPI)
	}

	c.JSON(http.StatusOK, miscellaneous_fieldAPIs)
}

// PostMiscellaneous_field
//
// swagger:route POST /miscellaneous_fields miscellaneous_fields postMiscellaneous_field
//
// Creates a miscellaneous_field
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMiscellaneous_field(c *gin.Context) {

	mutexMiscellaneous_field.Lock()
	defer mutexMiscellaneous_field.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMiscellaneous_fields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous_field.GetDB()

	// Validate input
	var input orm.Miscellaneous_fieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create miscellaneous_field
	miscellaneous_fieldDB := orm.Miscellaneous_fieldDB{}
	miscellaneous_fieldDB.Miscellaneous_fieldPointersEncoding = input.Miscellaneous_fieldPointersEncoding
	miscellaneous_fieldDB.CopyBasicFieldsFromMiscellaneous_field_WOP(&input.Miscellaneous_field_WOP)

	query := db.Create(&miscellaneous_fieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMiscellaneous_field.CheckoutPhaseOneInstance(&miscellaneous_fieldDB)
	miscellaneous_field := backRepo.BackRepoMiscellaneous_field.Map_Miscellaneous_fieldDBID_Miscellaneous_fieldPtr[miscellaneous_fieldDB.ID]

	if miscellaneous_field != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), miscellaneous_field)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, miscellaneous_fieldDB)
}

// GetMiscellaneous_field
//
// swagger:route GET /miscellaneous_fields/{ID} miscellaneous_fields getMiscellaneous_field
//
// Gets the details for a miscellaneous_field.
//
// Responses:
// default: genericError
//
//	200: miscellaneous_fieldDBResponse
func (controller *Controller) GetMiscellaneous_field(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMiscellaneous_field", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous_field.GetDB()

	// Get miscellaneous_fieldDB in DB
	var miscellaneous_fieldDB orm.Miscellaneous_fieldDB
	if err := db.First(&miscellaneous_fieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var miscellaneous_fieldAPI orm.Miscellaneous_fieldAPI
	miscellaneous_fieldAPI.ID = miscellaneous_fieldDB.ID
	miscellaneous_fieldAPI.Miscellaneous_fieldPointersEncoding = miscellaneous_fieldDB.Miscellaneous_fieldPointersEncoding
	miscellaneous_fieldDB.CopyBasicFieldsToMiscellaneous_field_WOP(&miscellaneous_fieldAPI.Miscellaneous_field_WOP)

	c.JSON(http.StatusOK, miscellaneous_fieldAPI)
}

// UpdateMiscellaneous_field
//
// swagger:route PATCH /miscellaneous_fields/{ID} miscellaneous_fields updateMiscellaneous_field
//
// # Update a miscellaneous_field
//
// Responses:
// default: genericError
//
//	200: miscellaneous_fieldDBResponse
func (controller *Controller) UpdateMiscellaneous_field(c *gin.Context) {

	mutexMiscellaneous_field.Lock()
	defer mutexMiscellaneous_field.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMiscellaneous_field", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous_field.GetDB()

	// Validate input
	var input orm.Miscellaneous_fieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var miscellaneous_fieldDB orm.Miscellaneous_fieldDB

	// fetch the miscellaneous_field
	query := db.First(&miscellaneous_fieldDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	miscellaneous_fieldDB.CopyBasicFieldsFromMiscellaneous_field_WOP(&input.Miscellaneous_field_WOP)
	miscellaneous_fieldDB.Miscellaneous_fieldPointersEncoding = input.Miscellaneous_fieldPointersEncoding

	query = db.Model(&miscellaneous_fieldDB).Updates(miscellaneous_fieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	miscellaneous_fieldNew := new(models.Miscellaneous_field)
	miscellaneous_fieldDB.CopyBasicFieldsToMiscellaneous_field(miscellaneous_fieldNew)

	// redeem pointers
	miscellaneous_fieldDB.DecodePointers(backRepo, miscellaneous_fieldNew)

	// get stage instance from DB instance, and call callback function
	miscellaneous_fieldOld := backRepo.BackRepoMiscellaneous_field.Map_Miscellaneous_fieldDBID_Miscellaneous_fieldPtr[miscellaneous_fieldDB.ID]
	if miscellaneous_fieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), miscellaneous_fieldOld, miscellaneous_fieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the miscellaneous_fieldDB
	c.JSON(http.StatusOK, miscellaneous_fieldDB)
}

// DeleteMiscellaneous_field
//
// swagger:route DELETE /miscellaneous_fields/{ID} miscellaneous_fields deleteMiscellaneous_field
//
// # Delete a miscellaneous_field
//
// default: genericError
//
//	200: miscellaneous_fieldDBResponse
func (controller *Controller) DeleteMiscellaneous_field(c *gin.Context) {

	mutexMiscellaneous_field.Lock()
	defer mutexMiscellaneous_field.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMiscellaneous_field", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMiscellaneous_field.GetDB()

	// Get model if exist
	var miscellaneous_fieldDB orm.Miscellaneous_fieldDB
	if err := db.First(&miscellaneous_fieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&miscellaneous_fieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	miscellaneous_fieldDeleted := new(models.Miscellaneous_field)
	miscellaneous_fieldDB.CopyBasicFieldsToMiscellaneous_field(miscellaneous_fieldDeleted)

	// get stage instance from DB instance, and call callback function
	miscellaneous_fieldStaged := backRepo.BackRepoMiscellaneous_field.Map_Miscellaneous_fieldDBID_Miscellaneous_fieldPtr[miscellaneous_fieldDB.ID]
	if miscellaneous_fieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), miscellaneous_fieldStaged, miscellaneous_fieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
