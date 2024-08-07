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
var __Extend__dummysDeclaration__ models.Extend
var __Extend_time__dummyDeclaration time.Duration

var mutexExtend sync.Mutex

// An ExtendID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getExtend updateExtend deleteExtend
type ExtendID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ExtendInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postExtend updateExtend
type ExtendInput struct {
	// The Extend to submit or modify
	// in: body
	Extend *orm.ExtendAPI
}

// GetExtends
//
// swagger:route GET /extends extends getExtends
//
// # Get all extends
//
// Responses:
// default: genericError
//
//	200: extendDBResponse
func (controller *Controller) GetExtends(c *gin.Context) {

	// source slice
	var extendDBs []orm.ExtendDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtends", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtend.GetDB()

	query := db.Find(&extendDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	extendAPIs := make([]orm.ExtendAPI, 0)

	// for each extend, update fields from the database nullable fields
	for idx := range extendDBs {
		extendDB := &extendDBs[idx]
		_ = extendDB
		var extendAPI orm.ExtendAPI

		// insertion point for updating fields
		extendAPI.ID = extendDB.ID
		extendDB.CopyBasicFieldsToExtend_WOP(&extendAPI.Extend_WOP)
		extendAPI.ExtendPointersEncoding = extendDB.ExtendPointersEncoding
		extendAPIs = append(extendAPIs, extendAPI)
	}

	c.JSON(http.StatusOK, extendAPIs)
}

// PostExtend
//
// swagger:route POST /extends extends postExtend
//
// Creates a extend
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostExtend(c *gin.Context) {

	mutexExtend.Lock()
	defer mutexExtend.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostExtends", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtend.GetDB()

	// Validate input
	var input orm.ExtendAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create extend
	extendDB := orm.ExtendDB{}
	extendDB.ExtendPointersEncoding = input.ExtendPointersEncoding
	extendDB.CopyBasicFieldsFromExtend_WOP(&input.Extend_WOP)

	query := db.Create(&extendDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoExtend.CheckoutPhaseOneInstance(&extendDB)
	extend := backRepo.BackRepoExtend.Map_ExtendDBID_ExtendPtr[extendDB.ID]

	if extend != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), extend)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, extendDB)
}

// GetExtend
//
// swagger:route GET /extends/{ID} extends getExtend
//
// Gets the details for a extend.
//
// Responses:
// default: genericError
//
//	200: extendDBResponse
func (controller *Controller) GetExtend(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetExtend", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtend.GetDB()

	// Get extendDB in DB
	var extendDB orm.ExtendDB
	if err := db.First(&extendDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var extendAPI orm.ExtendAPI
	extendAPI.ID = extendDB.ID
	extendAPI.ExtendPointersEncoding = extendDB.ExtendPointersEncoding
	extendDB.CopyBasicFieldsToExtend_WOP(&extendAPI.Extend_WOP)

	c.JSON(http.StatusOK, extendAPI)
}

// UpdateExtend
//
// swagger:route PATCH /extends/{ID} extends updateExtend
//
// # Update a extend
//
// Responses:
// default: genericError
//
//	200: extendDBResponse
func (controller *Controller) UpdateExtend(c *gin.Context) {

	mutexExtend.Lock()
	defer mutexExtend.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateExtend", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtend.GetDB()

	// Validate input
	var input orm.ExtendAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var extendDB orm.ExtendDB

	// fetch the extend
	query := db.First(&extendDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	extendDB.CopyBasicFieldsFromExtend_WOP(&input.Extend_WOP)
	extendDB.ExtendPointersEncoding = input.ExtendPointersEncoding

	query = db.Model(&extendDB).Updates(extendDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	extendNew := new(models.Extend)
	extendDB.CopyBasicFieldsToExtend(extendNew)

	// redeem pointers
	extendDB.DecodePointers(backRepo, extendNew)

	// get stage instance from DB instance, and call callback function
	extendOld := backRepo.BackRepoExtend.Map_ExtendDBID_ExtendPtr[extendDB.ID]
	if extendOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), extendOld, extendNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the extendDB
	c.JSON(http.StatusOK, extendDB)
}

// DeleteExtend
//
// swagger:route DELETE /extends/{ID} extends deleteExtend
//
// # Delete a extend
//
// default: genericError
//
//	200: extendDBResponse
func (controller *Controller) DeleteExtend(c *gin.Context) {

	mutexExtend.Lock()
	defer mutexExtend.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteExtend", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoExtend.GetDB()

	// Get model if exist
	var extendDB orm.ExtendDB
	if err := db.First(&extendDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&extendDB)

	// get an instance (not staged) from DB instance, and call callback function
	extendDeleted := new(models.Extend)
	extendDB.CopyBasicFieldsToExtend(extendDeleted)

	// get stage instance from DB instance, and call callback function
	extendStaged := backRepo.BackRepoExtend.Map_ExtendDBID_ExtendPtr[extendDB.ID]
	if extendStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), extendStaged, extendDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
