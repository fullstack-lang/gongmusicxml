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
var __Ornaments__dummysDeclaration__ models.Ornaments
var __Ornaments_time__dummyDeclaration time.Duration

var mutexOrnaments sync.Mutex

// An OrnamentsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOrnaments updateOrnaments deleteOrnaments
type OrnamentsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// OrnamentsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOrnaments updateOrnaments
type OrnamentsInput struct {
	// The Ornaments to submit or modify
	// in: body
	Ornaments *orm.OrnamentsAPI
}

// GetOrnamentss
//
// swagger:route GET /ornamentss ornamentss getOrnamentss
//
// # Get all ornamentss
//
// Responses:
// default: genericError
//
//	200: ornamentsDBResponse
func (controller *Controller) GetOrnamentss(c *gin.Context) {

	// source slice
	var ornamentsDBs []orm.OrnamentsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOrnamentss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOrnaments.GetDB()

	query := db.Find(&ornamentsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	ornamentsAPIs := make([]orm.OrnamentsAPI, 0)

	// for each ornaments, update fields from the database nullable fields
	for idx := range ornamentsDBs {
		ornamentsDB := &ornamentsDBs[idx]
		_ = ornamentsDB
		var ornamentsAPI orm.OrnamentsAPI

		// insertion point for updating fields
		ornamentsAPI.ID = ornamentsDB.ID
		ornamentsDB.CopyBasicFieldsToOrnaments_WOP(&ornamentsAPI.Ornaments_WOP)
		ornamentsAPI.OrnamentsPointersEncoding = ornamentsDB.OrnamentsPointersEncoding
		ornamentsAPIs = append(ornamentsAPIs, ornamentsAPI)
	}

	c.JSON(http.StatusOK, ornamentsAPIs)
}

// PostOrnaments
//
// swagger:route POST /ornamentss ornamentss postOrnaments
//
// Creates a ornaments
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOrnaments(c *gin.Context) {

	mutexOrnaments.Lock()
	defer mutexOrnaments.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOrnamentss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOrnaments.GetDB()

	// Validate input
	var input orm.OrnamentsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create ornaments
	ornamentsDB := orm.OrnamentsDB{}
	ornamentsDB.OrnamentsPointersEncoding = input.OrnamentsPointersEncoding
	ornamentsDB.CopyBasicFieldsFromOrnaments_WOP(&input.Ornaments_WOP)

	query := db.Create(&ornamentsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOrnaments.CheckoutPhaseOneInstance(&ornamentsDB)
	ornaments := backRepo.BackRepoOrnaments.Map_OrnamentsDBID_OrnamentsPtr[ornamentsDB.ID]

	if ornaments != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), ornaments)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, ornamentsDB)
}

// GetOrnaments
//
// swagger:route GET /ornamentss/{ID} ornamentss getOrnaments
//
// Gets the details for a ornaments.
//
// Responses:
// default: genericError
//
//	200: ornamentsDBResponse
func (controller *Controller) GetOrnaments(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOrnaments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOrnaments.GetDB()

	// Get ornamentsDB in DB
	var ornamentsDB orm.OrnamentsDB
	if err := db.First(&ornamentsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var ornamentsAPI orm.OrnamentsAPI
	ornamentsAPI.ID = ornamentsDB.ID
	ornamentsAPI.OrnamentsPointersEncoding = ornamentsDB.OrnamentsPointersEncoding
	ornamentsDB.CopyBasicFieldsToOrnaments_WOP(&ornamentsAPI.Ornaments_WOP)

	c.JSON(http.StatusOK, ornamentsAPI)
}

// UpdateOrnaments
//
// swagger:route PATCH /ornamentss/{ID} ornamentss updateOrnaments
//
// # Update a ornaments
//
// Responses:
// default: genericError
//
//	200: ornamentsDBResponse
func (controller *Controller) UpdateOrnaments(c *gin.Context) {

	mutexOrnaments.Lock()
	defer mutexOrnaments.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOrnaments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOrnaments.GetDB()

	// Validate input
	var input orm.OrnamentsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var ornamentsDB orm.OrnamentsDB

	// fetch the ornaments
	query := db.First(&ornamentsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	ornamentsDB.CopyBasicFieldsFromOrnaments_WOP(&input.Ornaments_WOP)
	ornamentsDB.OrnamentsPointersEncoding = input.OrnamentsPointersEncoding

	query = db.Model(&ornamentsDB).Updates(ornamentsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	ornamentsNew := new(models.Ornaments)
	ornamentsDB.CopyBasicFieldsToOrnaments(ornamentsNew)

	// redeem pointers
	ornamentsDB.DecodePointers(backRepo, ornamentsNew)

	// get stage instance from DB instance, and call callback function
	ornamentsOld := backRepo.BackRepoOrnaments.Map_OrnamentsDBID_OrnamentsPtr[ornamentsDB.ID]
	if ornamentsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), ornamentsOld, ornamentsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the ornamentsDB
	c.JSON(http.StatusOK, ornamentsDB)
}

// DeleteOrnaments
//
// swagger:route DELETE /ornamentss/{ID} ornamentss deleteOrnaments
//
// # Delete a ornaments
//
// default: genericError
//
//	200: ornamentsDBResponse
func (controller *Controller) DeleteOrnaments(c *gin.Context) {

	mutexOrnaments.Lock()
	defer mutexOrnaments.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOrnaments", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOrnaments.GetDB()

	// Get model if exist
	var ornamentsDB orm.OrnamentsDB
	if err := db.First(&ornamentsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&ornamentsDB)

	// get an instance (not staged) from DB instance, and call callback function
	ornamentsDeleted := new(models.Ornaments)
	ornamentsDB.CopyBasicFieldsToOrnaments(ornamentsDeleted)

	// get stage instance from DB instance, and call callback function
	ornamentsStaged := backRepo.BackRepoOrnaments.Map_OrnamentsDBID_OrnamentsPtr[ornamentsDB.ID]
	if ornamentsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), ornamentsStaged, ornamentsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
