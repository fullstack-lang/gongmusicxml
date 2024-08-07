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
var __Other_notation__dummysDeclaration__ models.Other_notation
var __Other_notation_time__dummyDeclaration time.Duration

var mutexOther_notation sync.Mutex

// An Other_notationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOther_notation updateOther_notation deleteOther_notation
type Other_notationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Other_notationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOther_notation updateOther_notation
type Other_notationInput struct {
	// The Other_notation to submit or modify
	// in: body
	Other_notation *orm.Other_notationAPI
}

// GetOther_notations
//
// swagger:route GET /other_notations other_notations getOther_notations
//
// # Get all other_notations
//
// Responses:
// default: genericError
//
//	200: other_notationDBResponse
func (controller *Controller) GetOther_notations(c *gin.Context) {

	// source slice
	var other_notationDBs []orm.Other_notationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_notations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_notation.GetDB()

	query := db.Find(&other_notationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	other_notationAPIs := make([]orm.Other_notationAPI, 0)

	// for each other_notation, update fields from the database nullable fields
	for idx := range other_notationDBs {
		other_notationDB := &other_notationDBs[idx]
		_ = other_notationDB
		var other_notationAPI orm.Other_notationAPI

		// insertion point for updating fields
		other_notationAPI.ID = other_notationDB.ID
		other_notationDB.CopyBasicFieldsToOther_notation_WOP(&other_notationAPI.Other_notation_WOP)
		other_notationAPI.Other_notationPointersEncoding = other_notationDB.Other_notationPointersEncoding
		other_notationAPIs = append(other_notationAPIs, other_notationAPI)
	}

	c.JSON(http.StatusOK, other_notationAPIs)
}

// PostOther_notation
//
// swagger:route POST /other_notations other_notations postOther_notation
//
// Creates a other_notation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOther_notation(c *gin.Context) {

	mutexOther_notation.Lock()
	defer mutexOther_notation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOther_notations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_notation.GetDB()

	// Validate input
	var input orm.Other_notationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create other_notation
	other_notationDB := orm.Other_notationDB{}
	other_notationDB.Other_notationPointersEncoding = input.Other_notationPointersEncoding
	other_notationDB.CopyBasicFieldsFromOther_notation_WOP(&input.Other_notation_WOP)

	query := db.Create(&other_notationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOther_notation.CheckoutPhaseOneInstance(&other_notationDB)
	other_notation := backRepo.BackRepoOther_notation.Map_Other_notationDBID_Other_notationPtr[other_notationDB.ID]

	if other_notation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), other_notation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, other_notationDB)
}

// GetOther_notation
//
// swagger:route GET /other_notations/{ID} other_notations getOther_notation
//
// Gets the details for a other_notation.
//
// Responses:
// default: genericError
//
//	200: other_notationDBResponse
func (controller *Controller) GetOther_notation(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_notation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_notation.GetDB()

	// Get other_notationDB in DB
	var other_notationDB orm.Other_notationDB
	if err := db.First(&other_notationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var other_notationAPI orm.Other_notationAPI
	other_notationAPI.ID = other_notationDB.ID
	other_notationAPI.Other_notationPointersEncoding = other_notationDB.Other_notationPointersEncoding
	other_notationDB.CopyBasicFieldsToOther_notation_WOP(&other_notationAPI.Other_notation_WOP)

	c.JSON(http.StatusOK, other_notationAPI)
}

// UpdateOther_notation
//
// swagger:route PATCH /other_notations/{ID} other_notations updateOther_notation
//
// # Update a other_notation
//
// Responses:
// default: genericError
//
//	200: other_notationDBResponse
func (controller *Controller) UpdateOther_notation(c *gin.Context) {

	mutexOther_notation.Lock()
	defer mutexOther_notation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOther_notation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_notation.GetDB()

	// Validate input
	var input orm.Other_notationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var other_notationDB orm.Other_notationDB

	// fetch the other_notation
	query := db.First(&other_notationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	other_notationDB.CopyBasicFieldsFromOther_notation_WOP(&input.Other_notation_WOP)
	other_notationDB.Other_notationPointersEncoding = input.Other_notationPointersEncoding

	query = db.Model(&other_notationDB).Updates(other_notationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	other_notationNew := new(models.Other_notation)
	other_notationDB.CopyBasicFieldsToOther_notation(other_notationNew)

	// redeem pointers
	other_notationDB.DecodePointers(backRepo, other_notationNew)

	// get stage instance from DB instance, and call callback function
	other_notationOld := backRepo.BackRepoOther_notation.Map_Other_notationDBID_Other_notationPtr[other_notationDB.ID]
	if other_notationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), other_notationOld, other_notationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the other_notationDB
	c.JSON(http.StatusOK, other_notationDB)
}

// DeleteOther_notation
//
// swagger:route DELETE /other_notations/{ID} other_notations deleteOther_notation
//
// # Delete a other_notation
//
// default: genericError
//
//	200: other_notationDBResponse
func (controller *Controller) DeleteOther_notation(c *gin.Context) {

	mutexOther_notation.Lock()
	defer mutexOther_notation.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOther_notation", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_notation.GetDB()

	// Get model if exist
	var other_notationDB orm.Other_notationDB
	if err := db.First(&other_notationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&other_notationDB)

	// get an instance (not staged) from DB instance, and call callback function
	other_notationDeleted := new(models.Other_notation)
	other_notationDB.CopyBasicFieldsToOther_notation(other_notationDeleted)

	// get stage instance from DB instance, and call callback function
	other_notationStaged := backRepo.BackRepoOther_notation.Map_Other_notationDBID_Other_notationPtr[other_notationDB.ID]
	if other_notationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), other_notationStaged, other_notationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
