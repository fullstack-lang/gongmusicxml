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
var __Direction__dummysDeclaration__ models.Direction
var __Direction_time__dummyDeclaration time.Duration

var mutexDirection sync.Mutex

// An DirectionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDirection updateDirection deleteDirection
type DirectionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DirectionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDirection updateDirection
type DirectionInput struct {
	// The Direction to submit or modify
	// in: body
	Direction *orm.DirectionAPI
}

// GetDirections
//
// swagger:route GET /directions directions getDirections
//
// # Get all directions
//
// Responses:
// default: genericError
//
//	200: directionDBResponse
func (controller *Controller) GetDirections(c *gin.Context) {

	// source slice
	var directionDBs []orm.DirectionDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDirections", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection.GetDB()

	query := db.Find(&directionDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	directionAPIs := make([]orm.DirectionAPI, 0)

	// for each direction, update fields from the database nullable fields
	for idx := range directionDBs {
		directionDB := &directionDBs[idx]
		_ = directionDB
		var directionAPI orm.DirectionAPI

		// insertion point for updating fields
		directionAPI.ID = directionDB.ID
		directionDB.CopyBasicFieldsToDirection_WOP(&directionAPI.Direction_WOP)
		directionAPI.DirectionPointersEncoding = directionDB.DirectionPointersEncoding
		directionAPIs = append(directionAPIs, directionAPI)
	}

	c.JSON(http.StatusOK, directionAPIs)
}

// PostDirection
//
// swagger:route POST /directions directions postDirection
//
// Creates a direction
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDirection(c *gin.Context) {

	mutexDirection.Lock()
	defer mutexDirection.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDirections", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection.GetDB()

	// Validate input
	var input orm.DirectionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create direction
	directionDB := orm.DirectionDB{}
	directionDB.DirectionPointersEncoding = input.DirectionPointersEncoding
	directionDB.CopyBasicFieldsFromDirection_WOP(&input.Direction_WOP)

	query := db.Create(&directionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDirection.CheckoutPhaseOneInstance(&directionDB)
	direction := backRepo.BackRepoDirection.Map_DirectionDBID_DirectionPtr[directionDB.ID]

	if direction != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), direction)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, directionDB)
}

// GetDirection
//
// swagger:route GET /directions/{ID} directions getDirection
//
// Gets the details for a direction.
//
// Responses:
// default: genericError
//
//	200: directionDBResponse
func (controller *Controller) GetDirection(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDirection", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection.GetDB()

	// Get directionDB in DB
	var directionDB orm.DirectionDB
	if err := db.First(&directionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var directionAPI orm.DirectionAPI
	directionAPI.ID = directionDB.ID
	directionAPI.DirectionPointersEncoding = directionDB.DirectionPointersEncoding
	directionDB.CopyBasicFieldsToDirection_WOP(&directionAPI.Direction_WOP)

	c.JSON(http.StatusOK, directionAPI)
}

// UpdateDirection
//
// swagger:route PATCH /directions/{ID} directions updateDirection
//
// # Update a direction
//
// Responses:
// default: genericError
//
//	200: directionDBResponse
func (controller *Controller) UpdateDirection(c *gin.Context) {

	mutexDirection.Lock()
	defer mutexDirection.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDirection", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection.GetDB()

	// Validate input
	var input orm.DirectionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var directionDB orm.DirectionDB

	// fetch the direction
	query := db.First(&directionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	directionDB.CopyBasicFieldsFromDirection_WOP(&input.Direction_WOP)
	directionDB.DirectionPointersEncoding = input.DirectionPointersEncoding

	query = db.Model(&directionDB).Updates(directionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	directionNew := new(models.Direction)
	directionDB.CopyBasicFieldsToDirection(directionNew)

	// redeem pointers
	directionDB.DecodePointers(backRepo, directionNew)

	// get stage instance from DB instance, and call callback function
	directionOld := backRepo.BackRepoDirection.Map_DirectionDBID_DirectionPtr[directionDB.ID]
	if directionOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), directionOld, directionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the directionDB
	c.JSON(http.StatusOK, directionDB)
}

// DeleteDirection
//
// swagger:route DELETE /directions/{ID} directions deleteDirection
//
// # Delete a direction
//
// default: genericError
//
//	200: directionDBResponse
func (controller *Controller) DeleteDirection(c *gin.Context) {

	mutexDirection.Lock()
	defer mutexDirection.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDirection", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection.GetDB()

	// Get model if exist
	var directionDB orm.DirectionDB
	if err := db.First(&directionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&directionDB)

	// get an instance (not staged) from DB instance, and call callback function
	directionDeleted := new(models.Direction)
	directionDB.CopyBasicFieldsToDirection(directionDeleted)

	// get stage instance from DB instance, and call callback function
	directionStaged := backRepo.BackRepoDirection.Map_DirectionDBID_DirectionPtr[directionDB.ID]
	if directionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), directionStaged, directionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
