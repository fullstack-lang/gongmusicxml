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
var __Horizontal_turn__dummysDeclaration__ models.Horizontal_turn
var __Horizontal_turn_time__dummyDeclaration time.Duration

var mutexHorizontal_turn sync.Mutex

// An Horizontal_turnID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHorizontal_turn updateHorizontal_turn deleteHorizontal_turn
type Horizontal_turnID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Horizontal_turnInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHorizontal_turn updateHorizontal_turn
type Horizontal_turnInput struct {
	// The Horizontal_turn to submit or modify
	// in: body
	Horizontal_turn *orm.Horizontal_turnAPI
}

// GetHorizontal_turns
//
// swagger:route GET /horizontal_turns horizontal_turns getHorizontal_turns
//
// # Get all horizontal_turns
//
// Responses:
// default: genericError
//
//	200: horizontal_turnDBResponse
func (controller *Controller) GetHorizontal_turns(c *gin.Context) {

	// source slice
	var horizontal_turnDBs []orm.Horizontal_turnDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHorizontal_turns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontal_turn.GetDB()

	query := db.Find(&horizontal_turnDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	horizontal_turnAPIs := make([]orm.Horizontal_turnAPI, 0)

	// for each horizontal_turn, update fields from the database nullable fields
	for idx := range horizontal_turnDBs {
		horizontal_turnDB := &horizontal_turnDBs[idx]
		_ = horizontal_turnDB
		var horizontal_turnAPI orm.Horizontal_turnAPI

		// insertion point for updating fields
		horizontal_turnAPI.ID = horizontal_turnDB.ID
		horizontal_turnDB.CopyBasicFieldsToHorizontal_turn_WOP(&horizontal_turnAPI.Horizontal_turn_WOP)
		horizontal_turnAPI.Horizontal_turnPointersEncoding = horizontal_turnDB.Horizontal_turnPointersEncoding
		horizontal_turnAPIs = append(horizontal_turnAPIs, horizontal_turnAPI)
	}

	c.JSON(http.StatusOK, horizontal_turnAPIs)
}

// PostHorizontal_turn
//
// swagger:route POST /horizontal_turns horizontal_turns postHorizontal_turn
//
// Creates a horizontal_turn
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHorizontal_turn(c *gin.Context) {

	mutexHorizontal_turn.Lock()
	defer mutexHorizontal_turn.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHorizontal_turns", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontal_turn.GetDB()

	// Validate input
	var input orm.Horizontal_turnAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create horizontal_turn
	horizontal_turnDB := orm.Horizontal_turnDB{}
	horizontal_turnDB.Horizontal_turnPointersEncoding = input.Horizontal_turnPointersEncoding
	horizontal_turnDB.CopyBasicFieldsFromHorizontal_turn_WOP(&input.Horizontal_turn_WOP)

	query := db.Create(&horizontal_turnDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHorizontal_turn.CheckoutPhaseOneInstance(&horizontal_turnDB)
	horizontal_turn := backRepo.BackRepoHorizontal_turn.Map_Horizontal_turnDBID_Horizontal_turnPtr[horizontal_turnDB.ID]

	if horizontal_turn != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), horizontal_turn)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, horizontal_turnDB)
}

// GetHorizontal_turn
//
// swagger:route GET /horizontal_turns/{ID} horizontal_turns getHorizontal_turn
//
// Gets the details for a horizontal_turn.
//
// Responses:
// default: genericError
//
//	200: horizontal_turnDBResponse
func (controller *Controller) GetHorizontal_turn(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHorizontal_turn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontal_turn.GetDB()

	// Get horizontal_turnDB in DB
	var horizontal_turnDB orm.Horizontal_turnDB
	if err := db.First(&horizontal_turnDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var horizontal_turnAPI orm.Horizontal_turnAPI
	horizontal_turnAPI.ID = horizontal_turnDB.ID
	horizontal_turnAPI.Horizontal_turnPointersEncoding = horizontal_turnDB.Horizontal_turnPointersEncoding
	horizontal_turnDB.CopyBasicFieldsToHorizontal_turn_WOP(&horizontal_turnAPI.Horizontal_turn_WOP)

	c.JSON(http.StatusOK, horizontal_turnAPI)
}

// UpdateHorizontal_turn
//
// swagger:route PATCH /horizontal_turns/{ID} horizontal_turns updateHorizontal_turn
//
// # Update a horizontal_turn
//
// Responses:
// default: genericError
//
//	200: horizontal_turnDBResponse
func (controller *Controller) UpdateHorizontal_turn(c *gin.Context) {

	mutexHorizontal_turn.Lock()
	defer mutexHorizontal_turn.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHorizontal_turn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontal_turn.GetDB()

	// Validate input
	var input orm.Horizontal_turnAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var horizontal_turnDB orm.Horizontal_turnDB

	// fetch the horizontal_turn
	query := db.First(&horizontal_turnDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	horizontal_turnDB.CopyBasicFieldsFromHorizontal_turn_WOP(&input.Horizontal_turn_WOP)
	horizontal_turnDB.Horizontal_turnPointersEncoding = input.Horizontal_turnPointersEncoding

	query = db.Model(&horizontal_turnDB).Updates(horizontal_turnDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	horizontal_turnNew := new(models.Horizontal_turn)
	horizontal_turnDB.CopyBasicFieldsToHorizontal_turn(horizontal_turnNew)

	// redeem pointers
	horizontal_turnDB.DecodePointers(backRepo, horizontal_turnNew)

	// get stage instance from DB instance, and call callback function
	horizontal_turnOld := backRepo.BackRepoHorizontal_turn.Map_Horizontal_turnDBID_Horizontal_turnPtr[horizontal_turnDB.ID]
	if horizontal_turnOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), horizontal_turnOld, horizontal_turnNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the horizontal_turnDB
	c.JSON(http.StatusOK, horizontal_turnDB)
}

// DeleteHorizontal_turn
//
// swagger:route DELETE /horizontal_turns/{ID} horizontal_turns deleteHorizontal_turn
//
// # Delete a horizontal_turn
//
// default: genericError
//
//	200: horizontal_turnDBResponse
func (controller *Controller) DeleteHorizontal_turn(c *gin.Context) {

	mutexHorizontal_turn.Lock()
	defer mutexHorizontal_turn.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHorizontal_turn", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHorizontal_turn.GetDB()

	// Get model if exist
	var horizontal_turnDB orm.Horizontal_turnDB
	if err := db.First(&horizontal_turnDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&horizontal_turnDB)

	// get an instance (not staged) from DB instance, and call callback function
	horizontal_turnDeleted := new(models.Horizontal_turn)
	horizontal_turnDB.CopyBasicFieldsToHorizontal_turn(horizontal_turnDeleted)

	// get stage instance from DB instance, and call callback function
	horizontal_turnStaged := backRepo.BackRepoHorizontal_turn.Map_Horizontal_turnDBID_Horizontal_turnPtr[horizontal_turnDB.ID]
	if horizontal_turnStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), horizontal_turnStaged, horizontal_turnDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
