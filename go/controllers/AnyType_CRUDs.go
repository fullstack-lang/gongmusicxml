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
var __AnyType__dummysDeclaration__ models.AnyType
var __AnyType_time__dummyDeclaration time.Duration

var mutexAnyType sync.Mutex

// An AnyTypeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAnyType updateAnyType deleteAnyType
type AnyTypeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AnyTypeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAnyType updateAnyType
type AnyTypeInput struct {
	// The AnyType to submit or modify
	// in: body
	AnyType *orm.AnyTypeAPI
}

// GetAnyTypes
//
// swagger:route GET /anytypes anytypes getAnyTypes
//
// # Get all anytypes
//
// Responses:
// default: genericError
//
//	200: anytypeDBResponse
func (controller *Controller) GetAnyTypes(c *gin.Context) {

	// source slice
	var anytypeDBs []orm.AnyTypeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnyTypes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnyType.GetDB()

	query := db.Find(&anytypeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	anytypeAPIs := make([]orm.AnyTypeAPI, 0)

	// for each anytype, update fields from the database nullable fields
	for idx := range anytypeDBs {
		anytypeDB := &anytypeDBs[idx]
		_ = anytypeDB
		var anytypeAPI orm.AnyTypeAPI

		// insertion point for updating fields
		anytypeAPI.ID = anytypeDB.ID
		anytypeDB.CopyBasicFieldsToAnyType_WOP(&anytypeAPI.AnyType_WOP)
		anytypeAPI.AnyTypePointersEncoding = anytypeDB.AnyTypePointersEncoding
		anytypeAPIs = append(anytypeAPIs, anytypeAPI)
	}

	c.JSON(http.StatusOK, anytypeAPIs)
}

// PostAnyType
//
// swagger:route POST /anytypes anytypes postAnyType
//
// Creates a anytype
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAnyType(c *gin.Context) {

	mutexAnyType.Lock()
	defer mutexAnyType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAnyTypes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnyType.GetDB()

	// Validate input
	var input orm.AnyTypeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create anytype
	anytypeDB := orm.AnyTypeDB{}
	anytypeDB.AnyTypePointersEncoding = input.AnyTypePointersEncoding
	anytypeDB.CopyBasicFieldsFromAnyType_WOP(&input.AnyType_WOP)

	query := db.Create(&anytypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAnyType.CheckoutPhaseOneInstance(&anytypeDB)
	anytype := backRepo.BackRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID]

	if anytype != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), anytype)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, anytypeDB)
}

// GetAnyType
//
// swagger:route GET /anytypes/{ID} anytypes getAnyType
//
// Gets the details for a anytype.
//
// Responses:
// default: genericError
//
//	200: anytypeDBResponse
func (controller *Controller) GetAnyType(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAnyType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnyType.GetDB()

	// Get anytypeDB in DB
	var anytypeDB orm.AnyTypeDB
	if err := db.First(&anytypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var anytypeAPI orm.AnyTypeAPI
	anytypeAPI.ID = anytypeDB.ID
	anytypeAPI.AnyTypePointersEncoding = anytypeDB.AnyTypePointersEncoding
	anytypeDB.CopyBasicFieldsToAnyType_WOP(&anytypeAPI.AnyType_WOP)

	c.JSON(http.StatusOK, anytypeAPI)
}

// UpdateAnyType
//
// swagger:route PATCH /anytypes/{ID} anytypes updateAnyType
//
// # Update a anytype
//
// Responses:
// default: genericError
//
//	200: anytypeDBResponse
func (controller *Controller) UpdateAnyType(c *gin.Context) {

	mutexAnyType.Lock()
	defer mutexAnyType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAnyType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnyType.GetDB()

	// Validate input
	var input orm.AnyTypeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var anytypeDB orm.AnyTypeDB

	// fetch the anytype
	query := db.First(&anytypeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	anytypeDB.CopyBasicFieldsFromAnyType_WOP(&input.AnyType_WOP)
	anytypeDB.AnyTypePointersEncoding = input.AnyTypePointersEncoding

	query = db.Model(&anytypeDB).Updates(anytypeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	anytypeNew := new(models.AnyType)
	anytypeDB.CopyBasicFieldsToAnyType(anytypeNew)

	// redeem pointers
	anytypeDB.DecodePointers(backRepo, anytypeNew)

	// get stage instance from DB instance, and call callback function
	anytypeOld := backRepo.BackRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID]
	if anytypeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), anytypeOld, anytypeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the anytypeDB
	c.JSON(http.StatusOK, anytypeDB)
}

// DeleteAnyType
//
// swagger:route DELETE /anytypes/{ID} anytypes deleteAnyType
//
// # Delete a anytype
//
// default: genericError
//
//	200: anytypeDBResponse
func (controller *Controller) DeleteAnyType(c *gin.Context) {

	mutexAnyType.Lock()
	defer mutexAnyType.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAnyType", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAnyType.GetDB()

	// Get model if exist
	var anytypeDB orm.AnyTypeDB
	if err := db.First(&anytypeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&anytypeDB)

	// get an instance (not staged) from DB instance, and call callback function
	anytypeDeleted := new(models.AnyType)
	anytypeDB.CopyBasicFieldsToAnyType(anytypeDeleted)

	// get stage instance from DB instance, and call callback function
	anytypeStaged := backRepo.BackRepoAnyType.Map_AnyTypeDBID_AnyTypePtr[anytypeDB.ID]
	if anytypeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), anytypeStaged, anytypeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
