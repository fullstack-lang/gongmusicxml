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
var __Direction_type__dummysDeclaration__ models.Direction_type
var __Direction_type_time__dummyDeclaration time.Duration

var mutexDirection_type sync.Mutex

// An Direction_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDirection_type updateDirection_type deleteDirection_type
type Direction_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Direction_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDirection_type updateDirection_type
type Direction_typeInput struct {
	// The Direction_type to submit or modify
	// in: body
	Direction_type *orm.Direction_typeAPI
}

// GetDirection_types
//
// swagger:route GET /direction_types direction_types getDirection_types
//
// # Get all direction_types
//
// Responses:
// default: genericError
//
//	200: direction_typeDBResponse
func (controller *Controller) GetDirection_types(c *gin.Context) {

	// source slice
	var direction_typeDBs []orm.Direction_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDirection_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection_type.GetDB()

	query := db.Find(&direction_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	direction_typeAPIs := make([]orm.Direction_typeAPI, 0)

	// for each direction_type, update fields from the database nullable fields
	for idx := range direction_typeDBs {
		direction_typeDB := &direction_typeDBs[idx]
		_ = direction_typeDB
		var direction_typeAPI orm.Direction_typeAPI

		// insertion point for updating fields
		direction_typeAPI.ID = direction_typeDB.ID
		direction_typeDB.CopyBasicFieldsToDirection_type_WOP(&direction_typeAPI.Direction_type_WOP)
		direction_typeAPI.Direction_typePointersEncoding = direction_typeDB.Direction_typePointersEncoding
		direction_typeAPIs = append(direction_typeAPIs, direction_typeAPI)
	}

	c.JSON(http.StatusOK, direction_typeAPIs)
}

// PostDirection_type
//
// swagger:route POST /direction_types direction_types postDirection_type
//
// Creates a direction_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDirection_type(c *gin.Context) {

	mutexDirection_type.Lock()
	defer mutexDirection_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDirection_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection_type.GetDB()

	// Validate input
	var input orm.Direction_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create direction_type
	direction_typeDB := orm.Direction_typeDB{}
	direction_typeDB.Direction_typePointersEncoding = input.Direction_typePointersEncoding
	direction_typeDB.CopyBasicFieldsFromDirection_type_WOP(&input.Direction_type_WOP)

	query := db.Create(&direction_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDirection_type.CheckoutPhaseOneInstance(&direction_typeDB)
	direction_type := backRepo.BackRepoDirection_type.Map_Direction_typeDBID_Direction_typePtr[direction_typeDB.ID]

	if direction_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), direction_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, direction_typeDB)
}

// GetDirection_type
//
// swagger:route GET /direction_types/{ID} direction_types getDirection_type
//
// Gets the details for a direction_type.
//
// Responses:
// default: genericError
//
//	200: direction_typeDBResponse
func (controller *Controller) GetDirection_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDirection_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection_type.GetDB()

	// Get direction_typeDB in DB
	var direction_typeDB orm.Direction_typeDB
	if err := db.First(&direction_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var direction_typeAPI orm.Direction_typeAPI
	direction_typeAPI.ID = direction_typeDB.ID
	direction_typeAPI.Direction_typePointersEncoding = direction_typeDB.Direction_typePointersEncoding
	direction_typeDB.CopyBasicFieldsToDirection_type_WOP(&direction_typeAPI.Direction_type_WOP)

	c.JSON(http.StatusOK, direction_typeAPI)
}

// UpdateDirection_type
//
// swagger:route PATCH /direction_types/{ID} direction_types updateDirection_type
//
// # Update a direction_type
//
// Responses:
// default: genericError
//
//	200: direction_typeDBResponse
func (controller *Controller) UpdateDirection_type(c *gin.Context) {

	mutexDirection_type.Lock()
	defer mutexDirection_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDirection_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection_type.GetDB()

	// Validate input
	var input orm.Direction_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var direction_typeDB orm.Direction_typeDB

	// fetch the direction_type
	query := db.First(&direction_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	direction_typeDB.CopyBasicFieldsFromDirection_type_WOP(&input.Direction_type_WOP)
	direction_typeDB.Direction_typePointersEncoding = input.Direction_typePointersEncoding

	query = db.Model(&direction_typeDB).Updates(direction_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	direction_typeNew := new(models.Direction_type)
	direction_typeDB.CopyBasicFieldsToDirection_type(direction_typeNew)

	// redeem pointers
	direction_typeDB.DecodePointers(backRepo, direction_typeNew)

	// get stage instance from DB instance, and call callback function
	direction_typeOld := backRepo.BackRepoDirection_type.Map_Direction_typeDBID_Direction_typePtr[direction_typeDB.ID]
	if direction_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), direction_typeOld, direction_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the direction_typeDB
	c.JSON(http.StatusOK, direction_typeDB)
}

// DeleteDirection_type
//
// swagger:route DELETE /direction_types/{ID} direction_types deleteDirection_type
//
// # Delete a direction_type
//
// default: genericError
//
//	200: direction_typeDBResponse
func (controller *Controller) DeleteDirection_type(c *gin.Context) {

	mutexDirection_type.Lock()
	defer mutexDirection_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDirection_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDirection_type.GetDB()

	// Get model if exist
	var direction_typeDB orm.Direction_typeDB
	if err := db.First(&direction_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&direction_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	direction_typeDeleted := new(models.Direction_type)
	direction_typeDB.CopyBasicFieldsToDirection_type(direction_typeDeleted)

	// get stage instance from DB instance, and call callback function
	direction_typeStaged := backRepo.BackRepoDirection_type.Map_Direction_typeDBID_Direction_typePtr[direction_typeDB.ID]
	if direction_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), direction_typeStaged, direction_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
