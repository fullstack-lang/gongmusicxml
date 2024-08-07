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
var __Other_play__dummysDeclaration__ models.Other_play
var __Other_play_time__dummyDeclaration time.Duration

var mutexOther_play sync.Mutex

// An Other_playID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getOther_play updateOther_play deleteOther_play
type Other_playID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Other_playInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postOther_play updateOther_play
type Other_playInput struct {
	// The Other_play to submit or modify
	// in: body
	Other_play *orm.Other_playAPI
}

// GetOther_plays
//
// swagger:route GET /other_plays other_plays getOther_plays
//
// # Get all other_plays
//
// Responses:
// default: genericError
//
//	200: other_playDBResponse
func (controller *Controller) GetOther_plays(c *gin.Context) {

	// source slice
	var other_playDBs []orm.Other_playDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_plays", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_play.GetDB()

	query := db.Find(&other_playDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	other_playAPIs := make([]orm.Other_playAPI, 0)

	// for each other_play, update fields from the database nullable fields
	for idx := range other_playDBs {
		other_playDB := &other_playDBs[idx]
		_ = other_playDB
		var other_playAPI orm.Other_playAPI

		// insertion point for updating fields
		other_playAPI.ID = other_playDB.ID
		other_playDB.CopyBasicFieldsToOther_play_WOP(&other_playAPI.Other_play_WOP)
		other_playAPI.Other_playPointersEncoding = other_playDB.Other_playPointersEncoding
		other_playAPIs = append(other_playAPIs, other_playAPI)
	}

	c.JSON(http.StatusOK, other_playAPIs)
}

// PostOther_play
//
// swagger:route POST /other_plays other_plays postOther_play
//
// Creates a other_play
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostOther_play(c *gin.Context) {

	mutexOther_play.Lock()
	defer mutexOther_play.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostOther_plays", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_play.GetDB()

	// Validate input
	var input orm.Other_playAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create other_play
	other_playDB := orm.Other_playDB{}
	other_playDB.Other_playPointersEncoding = input.Other_playPointersEncoding
	other_playDB.CopyBasicFieldsFromOther_play_WOP(&input.Other_play_WOP)

	query := db.Create(&other_playDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoOther_play.CheckoutPhaseOneInstance(&other_playDB)
	other_play := backRepo.BackRepoOther_play.Map_Other_playDBID_Other_playPtr[other_playDB.ID]

	if other_play != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), other_play)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, other_playDB)
}

// GetOther_play
//
// swagger:route GET /other_plays/{ID} other_plays getOther_play
//
// Gets the details for a other_play.
//
// Responses:
// default: genericError
//
//	200: other_playDBResponse
func (controller *Controller) GetOther_play(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetOther_play", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_play.GetDB()

	// Get other_playDB in DB
	var other_playDB orm.Other_playDB
	if err := db.First(&other_playDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var other_playAPI orm.Other_playAPI
	other_playAPI.ID = other_playDB.ID
	other_playAPI.Other_playPointersEncoding = other_playDB.Other_playPointersEncoding
	other_playDB.CopyBasicFieldsToOther_play_WOP(&other_playAPI.Other_play_WOP)

	c.JSON(http.StatusOK, other_playAPI)
}

// UpdateOther_play
//
// swagger:route PATCH /other_plays/{ID} other_plays updateOther_play
//
// # Update a other_play
//
// Responses:
// default: genericError
//
//	200: other_playDBResponse
func (controller *Controller) UpdateOther_play(c *gin.Context) {

	mutexOther_play.Lock()
	defer mutexOther_play.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateOther_play", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_play.GetDB()

	// Validate input
	var input orm.Other_playAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var other_playDB orm.Other_playDB

	// fetch the other_play
	query := db.First(&other_playDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	other_playDB.CopyBasicFieldsFromOther_play_WOP(&input.Other_play_WOP)
	other_playDB.Other_playPointersEncoding = input.Other_playPointersEncoding

	query = db.Model(&other_playDB).Updates(other_playDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	other_playNew := new(models.Other_play)
	other_playDB.CopyBasicFieldsToOther_play(other_playNew)

	// redeem pointers
	other_playDB.DecodePointers(backRepo, other_playNew)

	// get stage instance from DB instance, and call callback function
	other_playOld := backRepo.BackRepoOther_play.Map_Other_playDBID_Other_playPtr[other_playDB.ID]
	if other_playOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), other_playOld, other_playNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the other_playDB
	c.JSON(http.StatusOK, other_playDB)
}

// DeleteOther_play
//
// swagger:route DELETE /other_plays/{ID} other_plays deleteOther_play
//
// # Delete a other_play
//
// default: genericError
//
//	200: other_playDBResponse
func (controller *Controller) DeleteOther_play(c *gin.Context) {

	mutexOther_play.Lock()
	defer mutexOther_play.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteOther_play", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoOther_play.GetDB()

	// Get model if exist
	var other_playDB orm.Other_playDB
	if err := db.First(&other_playDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&other_playDB)

	// get an instance (not staged) from DB instance, and call callback function
	other_playDeleted := new(models.Other_play)
	other_playDB.CopyBasicFieldsToOther_play(other_playDeleted)

	// get stage instance from DB instance, and call callback function
	other_playStaged := backRepo.BackRepoOther_play.Map_Other_playDBID_Other_playPtr[other_playDB.ID]
	if other_playStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), other_playStaged, other_playDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
