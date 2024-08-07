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
var __Heel_toe__dummysDeclaration__ models.Heel_toe
var __Heel_toe_time__dummyDeclaration time.Duration

var mutexHeel_toe sync.Mutex

// An Heel_toeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHeel_toe updateHeel_toe deleteHeel_toe
type Heel_toeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Heel_toeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHeel_toe updateHeel_toe
type Heel_toeInput struct {
	// The Heel_toe to submit or modify
	// in: body
	Heel_toe *orm.Heel_toeAPI
}

// GetHeel_toes
//
// swagger:route GET /heel_toes heel_toes getHeel_toes
//
// # Get all heel_toes
//
// Responses:
// default: genericError
//
//	200: heel_toeDBResponse
func (controller *Controller) GetHeel_toes(c *gin.Context) {

	// source slice
	var heel_toeDBs []orm.Heel_toeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHeel_toes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHeel_toe.GetDB()

	query := db.Find(&heel_toeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	heel_toeAPIs := make([]orm.Heel_toeAPI, 0)

	// for each heel_toe, update fields from the database nullable fields
	for idx := range heel_toeDBs {
		heel_toeDB := &heel_toeDBs[idx]
		_ = heel_toeDB
		var heel_toeAPI orm.Heel_toeAPI

		// insertion point for updating fields
		heel_toeAPI.ID = heel_toeDB.ID
		heel_toeDB.CopyBasicFieldsToHeel_toe_WOP(&heel_toeAPI.Heel_toe_WOP)
		heel_toeAPI.Heel_toePointersEncoding = heel_toeDB.Heel_toePointersEncoding
		heel_toeAPIs = append(heel_toeAPIs, heel_toeAPI)
	}

	c.JSON(http.StatusOK, heel_toeAPIs)
}

// PostHeel_toe
//
// swagger:route POST /heel_toes heel_toes postHeel_toe
//
// Creates a heel_toe
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHeel_toe(c *gin.Context) {

	mutexHeel_toe.Lock()
	defer mutexHeel_toe.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHeel_toes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHeel_toe.GetDB()

	// Validate input
	var input orm.Heel_toeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create heel_toe
	heel_toeDB := orm.Heel_toeDB{}
	heel_toeDB.Heel_toePointersEncoding = input.Heel_toePointersEncoding
	heel_toeDB.CopyBasicFieldsFromHeel_toe_WOP(&input.Heel_toe_WOP)

	query := db.Create(&heel_toeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHeel_toe.CheckoutPhaseOneInstance(&heel_toeDB)
	heel_toe := backRepo.BackRepoHeel_toe.Map_Heel_toeDBID_Heel_toePtr[heel_toeDB.ID]

	if heel_toe != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), heel_toe)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, heel_toeDB)
}

// GetHeel_toe
//
// swagger:route GET /heel_toes/{ID} heel_toes getHeel_toe
//
// Gets the details for a heel_toe.
//
// Responses:
// default: genericError
//
//	200: heel_toeDBResponse
func (controller *Controller) GetHeel_toe(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHeel_toe", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHeel_toe.GetDB()

	// Get heel_toeDB in DB
	var heel_toeDB orm.Heel_toeDB
	if err := db.First(&heel_toeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var heel_toeAPI orm.Heel_toeAPI
	heel_toeAPI.ID = heel_toeDB.ID
	heel_toeAPI.Heel_toePointersEncoding = heel_toeDB.Heel_toePointersEncoding
	heel_toeDB.CopyBasicFieldsToHeel_toe_WOP(&heel_toeAPI.Heel_toe_WOP)

	c.JSON(http.StatusOK, heel_toeAPI)
}

// UpdateHeel_toe
//
// swagger:route PATCH /heel_toes/{ID} heel_toes updateHeel_toe
//
// # Update a heel_toe
//
// Responses:
// default: genericError
//
//	200: heel_toeDBResponse
func (controller *Controller) UpdateHeel_toe(c *gin.Context) {

	mutexHeel_toe.Lock()
	defer mutexHeel_toe.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHeel_toe", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHeel_toe.GetDB()

	// Validate input
	var input orm.Heel_toeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var heel_toeDB orm.Heel_toeDB

	// fetch the heel_toe
	query := db.First(&heel_toeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	heel_toeDB.CopyBasicFieldsFromHeel_toe_WOP(&input.Heel_toe_WOP)
	heel_toeDB.Heel_toePointersEncoding = input.Heel_toePointersEncoding

	query = db.Model(&heel_toeDB).Updates(heel_toeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	heel_toeNew := new(models.Heel_toe)
	heel_toeDB.CopyBasicFieldsToHeel_toe(heel_toeNew)

	// redeem pointers
	heel_toeDB.DecodePointers(backRepo, heel_toeNew)

	// get stage instance from DB instance, and call callback function
	heel_toeOld := backRepo.BackRepoHeel_toe.Map_Heel_toeDBID_Heel_toePtr[heel_toeDB.ID]
	if heel_toeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), heel_toeOld, heel_toeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the heel_toeDB
	c.JSON(http.StatusOK, heel_toeDB)
}

// DeleteHeel_toe
//
// swagger:route DELETE /heel_toes/{ID} heel_toes deleteHeel_toe
//
// # Delete a heel_toe
//
// default: genericError
//
//	200: heel_toeDBResponse
func (controller *Controller) DeleteHeel_toe(c *gin.Context) {

	mutexHeel_toe.Lock()
	defer mutexHeel_toe.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHeel_toe", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHeel_toe.GetDB()

	// Get model if exist
	var heel_toeDB orm.Heel_toeDB
	if err := db.First(&heel_toeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&heel_toeDB)

	// get an instance (not staged) from DB instance, and call callback function
	heel_toeDeleted := new(models.Heel_toe)
	heel_toeDB.CopyBasicFieldsToHeel_toe(heel_toeDeleted)

	// get stage instance from DB instance, and call callback function
	heel_toeStaged := backRepo.BackRepoHeel_toe.Map_Heel_toeDBID_Heel_toePtr[heel_toeDB.ID]
	if heel_toeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), heel_toeStaged, heel_toeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
