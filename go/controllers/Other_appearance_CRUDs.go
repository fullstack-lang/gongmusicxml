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
var __Other_appearance__dummysDeclaration__ models.Other_appearance
var __Other_appearance_time__dummyDeclaration time.Duration

var mutexOther_appearance sync.Mutex

// An Other_appearanceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOther_appearance updateOther_appearance deleteOther_appearance
type Other_appearanceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Other_appearanceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOther_appearance updateOther_appearance
type Other_appearanceInput struct {
	// The Other_appearance to submit or modify
	// in: body
	Other_appearance *orm.Other_appearanceAPI
}

// GetOther_appearances
//
// swagger:route GET /other_appearances other_appearances getOther_appearances
//
// # Get all other_appearances
//
// Responses:
// default: genericError
//
//	200: other_appearanceDBResponse
func (controller *Controller) GetOther_appearances(c *gin.Context) {

	// source slice
	var other_appearanceDBs []orm.Other_appearanceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_appearances", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_appearance.GetDB()

	query := db.Find(&other_appearanceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	other_appearanceAPIs := make([]orm.Other_appearanceAPI, 0)

	// for each other_appearance, update fields from the database nullable fields
	for idx := range other_appearanceDBs {
		other_appearanceDB := &other_appearanceDBs[idx]
		_ = other_appearanceDB
		var other_appearanceAPI orm.Other_appearanceAPI

		// insertion point for updating fields
		other_appearanceAPI.ID = other_appearanceDB.ID
		other_appearanceDB.CopyBasicFieldsToOther_appearance_WOP(&other_appearanceAPI.Other_appearance_WOP)
		other_appearanceAPI.Other_appearancePointersEncoding = other_appearanceDB.Other_appearancePointersEncoding
		other_appearanceAPIs = append(other_appearanceAPIs, other_appearanceAPI)
	}

	c.JSON(http.StatusOK, other_appearanceAPIs)
}

// PostOther_appearance
//
// swagger:route POST /other_appearances other_appearances postOther_appearance
//
// Creates a other_appearance
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOther_appearance(c *gin.Context) {

	mutexOther_appearance.Lock()
	defer mutexOther_appearance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOther_appearances", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_appearance.GetDB()

	// Validate input
	var input orm.Other_appearanceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create other_appearance
	other_appearanceDB := orm.Other_appearanceDB{}
	other_appearanceDB.Other_appearancePointersEncoding = input.Other_appearancePointersEncoding
	other_appearanceDB.CopyBasicFieldsFromOther_appearance_WOP(&input.Other_appearance_WOP)

	query := db.Create(&other_appearanceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOther_appearance.CheckoutPhaseOneInstance(&other_appearanceDB)
	other_appearance := backRepo.BackRepoOther_appearance.Map_Other_appearanceDBID_Other_appearancePtr[other_appearanceDB.ID]

	if other_appearance != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), other_appearance)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, other_appearanceDB)
}

// GetOther_appearance
//
// swagger:route GET /other_appearances/{ID} other_appearances getOther_appearance
//
// Gets the details for a other_appearance.
//
// Responses:
// default: genericError
//
//	200: other_appearanceDBResponse
func (controller *Controller) GetOther_appearance(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_appearance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_appearance.GetDB()

	// Get other_appearanceDB in DB
	var other_appearanceDB orm.Other_appearanceDB
	if err := db.First(&other_appearanceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var other_appearanceAPI orm.Other_appearanceAPI
	other_appearanceAPI.ID = other_appearanceDB.ID
	other_appearanceAPI.Other_appearancePointersEncoding = other_appearanceDB.Other_appearancePointersEncoding
	other_appearanceDB.CopyBasicFieldsToOther_appearance_WOP(&other_appearanceAPI.Other_appearance_WOP)

	c.JSON(http.StatusOK, other_appearanceAPI)
}

// UpdateOther_appearance
//
// swagger:route PATCH /other_appearances/{ID} other_appearances updateOther_appearance
//
// # Update a other_appearance
//
// Responses:
// default: genericError
//
//	200: other_appearanceDBResponse
func (controller *Controller) UpdateOther_appearance(c *gin.Context) {

	mutexOther_appearance.Lock()
	defer mutexOther_appearance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOther_appearance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_appearance.GetDB()

	// Validate input
	var input orm.Other_appearanceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var other_appearanceDB orm.Other_appearanceDB

	// fetch the other_appearance
	query := db.First(&other_appearanceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	other_appearanceDB.CopyBasicFieldsFromOther_appearance_WOP(&input.Other_appearance_WOP)
	other_appearanceDB.Other_appearancePointersEncoding = input.Other_appearancePointersEncoding

	query = db.Model(&other_appearanceDB).Updates(other_appearanceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	other_appearanceNew := new(models.Other_appearance)
	other_appearanceDB.CopyBasicFieldsToOther_appearance(other_appearanceNew)

	// redeem pointers
	other_appearanceDB.DecodePointers(backRepo, other_appearanceNew)

	// get stage instance from DB instance, and call callback function
	other_appearanceOld := backRepo.BackRepoOther_appearance.Map_Other_appearanceDBID_Other_appearancePtr[other_appearanceDB.ID]
	if other_appearanceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), other_appearanceOld, other_appearanceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the other_appearanceDB
	c.JSON(http.StatusOK, other_appearanceDB)
}

// DeleteOther_appearance
//
// swagger:route DELETE /other_appearances/{ID} other_appearances deleteOther_appearance
//
// # Delete a other_appearance
//
// default: genericError
//
//	200: other_appearanceDBResponse
func (controller *Controller) DeleteOther_appearance(c *gin.Context) {

	mutexOther_appearance.Lock()
	defer mutexOther_appearance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOther_appearance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_appearance.GetDB()

	// Get model if exist
	var other_appearanceDB orm.Other_appearanceDB
	if err := db.First(&other_appearanceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&other_appearanceDB)

	// get an instance (not staged) from DB instance, and call callback function
	other_appearanceDeleted := new(models.Other_appearance)
	other_appearanceDB.CopyBasicFieldsToOther_appearance(other_appearanceDeleted)

	// get stage instance from DB instance, and call callback function
	other_appearanceStaged := backRepo.BackRepoOther_appearance.Map_Other_appearanceDBID_Other_appearancePtr[other_appearanceDB.ID]
	if other_appearanceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), other_appearanceStaged, other_appearanceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
