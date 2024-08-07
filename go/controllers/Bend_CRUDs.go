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
var __Bend__dummysDeclaration__ models.Bend
var __Bend_time__dummyDeclaration time.Duration

var mutexBend sync.Mutex

// An BendID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBend updateBend deleteBend
type BendID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BendInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBend updateBend
type BendInput struct {
	// The Bend to submit or modify
	// in: body
	Bend *orm.BendAPI
}

// GetBends
//
// swagger:route GET /bends bends getBends
//
// # Get all bends
//
// Responses:
// default: genericError
//
//	200: bendDBResponse
func (controller *Controller) GetBends(c *gin.Context) {

	// source slice
	var bendDBs []orm.BendDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBends", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBend.GetDB()

	query := db.Find(&bendDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bendAPIs := make([]orm.BendAPI, 0)

	// for each bend, update fields from the database nullable fields
	for idx := range bendDBs {
		bendDB := &bendDBs[idx]
		_ = bendDB
		var bendAPI orm.BendAPI

		// insertion point for updating fields
		bendAPI.ID = bendDB.ID
		bendDB.CopyBasicFieldsToBend_WOP(&bendAPI.Bend_WOP)
		bendAPI.BendPointersEncoding = bendDB.BendPointersEncoding
		bendAPIs = append(bendAPIs, bendAPI)
	}

	c.JSON(http.StatusOK, bendAPIs)
}

// PostBend
//
// swagger:route POST /bends bends postBend
//
// Creates a bend
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBend(c *gin.Context) {

	mutexBend.Lock()
	defer mutexBend.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBends", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBend.GetDB()

	// Validate input
	var input orm.BendAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bend
	bendDB := orm.BendDB{}
	bendDB.BendPointersEncoding = input.BendPointersEncoding
	bendDB.CopyBasicFieldsFromBend_WOP(&input.Bend_WOP)

	query := db.Create(&bendDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBend.CheckoutPhaseOneInstance(&bendDB)
	bend := backRepo.BackRepoBend.Map_BendDBID_BendPtr[bendDB.ID]

	if bend != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bend)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bendDB)
}

// GetBend
//
// swagger:route GET /bends/{ID} bends getBend
//
// Gets the details for a bend.
//
// Responses:
// default: genericError
//
//	200: bendDBResponse
func (controller *Controller) GetBend(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBend", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBend.GetDB()

	// Get bendDB in DB
	var bendDB orm.BendDB
	if err := db.First(&bendDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bendAPI orm.BendAPI
	bendAPI.ID = bendDB.ID
	bendAPI.BendPointersEncoding = bendDB.BendPointersEncoding
	bendDB.CopyBasicFieldsToBend_WOP(&bendAPI.Bend_WOP)

	c.JSON(http.StatusOK, bendAPI)
}

// UpdateBend
//
// swagger:route PATCH /bends/{ID} bends updateBend
//
// # Update a bend
//
// Responses:
// default: genericError
//
//	200: bendDBResponse
func (controller *Controller) UpdateBend(c *gin.Context) {

	mutexBend.Lock()
	defer mutexBend.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBend", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBend.GetDB()

	// Validate input
	var input orm.BendAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bendDB orm.BendDB

	// fetch the bend
	query := db.First(&bendDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bendDB.CopyBasicFieldsFromBend_WOP(&input.Bend_WOP)
	bendDB.BendPointersEncoding = input.BendPointersEncoding

	query = db.Model(&bendDB).Updates(bendDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bendNew := new(models.Bend)
	bendDB.CopyBasicFieldsToBend(bendNew)

	// redeem pointers
	bendDB.DecodePointers(backRepo, bendNew)

	// get stage instance from DB instance, and call callback function
	bendOld := backRepo.BackRepoBend.Map_BendDBID_BendPtr[bendDB.ID]
	if bendOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bendOld, bendNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bendDB
	c.JSON(http.StatusOK, bendDB)
}

// DeleteBend
//
// swagger:route DELETE /bends/{ID} bends deleteBend
//
// # Delete a bend
//
// default: genericError
//
//	200: bendDBResponse
func (controller *Controller) DeleteBend(c *gin.Context) {

	mutexBend.Lock()
	defer mutexBend.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBend", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBend.GetDB()

	// Get model if exist
	var bendDB orm.BendDB
	if err := db.First(&bendDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bendDB)

	// get an instance (not staged) from DB instance, and call callback function
	bendDeleted := new(models.Bend)
	bendDB.CopyBasicFieldsToBend(bendDeleted)

	// get stage instance from DB instance, and call callback function
	bendStaged := backRepo.BackRepoBend.Map_BendDBID_BendPtr[bendDB.ID]
	if bendStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bendStaged, bendDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
