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
var __Dynamics__dummysDeclaration__ models.Dynamics
var __Dynamics_time__dummyDeclaration time.Duration

var mutexDynamics sync.Mutex

// An DynamicsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDynamics updateDynamics deleteDynamics
type DynamicsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DynamicsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDynamics updateDynamics
type DynamicsInput struct {
	// The Dynamics to submit or modify
	// in: body
	Dynamics *orm.DynamicsAPI
}

// GetDynamicss
//
// swagger:route GET /dynamicss dynamicss getDynamicss
//
// # Get all dynamicss
//
// Responses:
// default: genericError
//
//	200: dynamicsDBResponse
func (controller *Controller) GetDynamicss(c *gin.Context) {

	// source slice
	var dynamicsDBs []orm.DynamicsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDynamicss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDynamics.GetDB()

	query := db.Find(&dynamicsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	dynamicsAPIs := make([]orm.DynamicsAPI, 0)

	// for each dynamics, update fields from the database nullable fields
	for idx := range dynamicsDBs {
		dynamicsDB := &dynamicsDBs[idx]
		_ = dynamicsDB
		var dynamicsAPI orm.DynamicsAPI

		// insertion point for updating fields
		dynamicsAPI.ID = dynamicsDB.ID
		dynamicsDB.CopyBasicFieldsToDynamics_WOP(&dynamicsAPI.Dynamics_WOP)
		dynamicsAPI.DynamicsPointersEncoding = dynamicsDB.DynamicsPointersEncoding
		dynamicsAPIs = append(dynamicsAPIs, dynamicsAPI)
	}

	c.JSON(http.StatusOK, dynamicsAPIs)
}

// PostDynamics
//
// swagger:route POST /dynamicss dynamicss postDynamics
//
// Creates a dynamics
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDynamics(c *gin.Context) {

	mutexDynamics.Lock()
	defer mutexDynamics.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDynamicss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDynamics.GetDB()

	// Validate input
	var input orm.DynamicsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create dynamics
	dynamicsDB := orm.DynamicsDB{}
	dynamicsDB.DynamicsPointersEncoding = input.DynamicsPointersEncoding
	dynamicsDB.CopyBasicFieldsFromDynamics_WOP(&input.Dynamics_WOP)

	query := db.Create(&dynamicsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDynamics.CheckoutPhaseOneInstance(&dynamicsDB)
	dynamics := backRepo.BackRepoDynamics.Map_DynamicsDBID_DynamicsPtr[dynamicsDB.ID]

	if dynamics != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), dynamics)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, dynamicsDB)
}

// GetDynamics
//
// swagger:route GET /dynamicss/{ID} dynamicss getDynamics
//
// Gets the details for a dynamics.
//
// Responses:
// default: genericError
//
//	200: dynamicsDBResponse
func (controller *Controller) GetDynamics(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDynamics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDynamics.GetDB()

	// Get dynamicsDB in DB
	var dynamicsDB orm.DynamicsDB
	if err := db.First(&dynamicsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var dynamicsAPI orm.DynamicsAPI
	dynamicsAPI.ID = dynamicsDB.ID
	dynamicsAPI.DynamicsPointersEncoding = dynamicsDB.DynamicsPointersEncoding
	dynamicsDB.CopyBasicFieldsToDynamics_WOP(&dynamicsAPI.Dynamics_WOP)

	c.JSON(http.StatusOK, dynamicsAPI)
}

// UpdateDynamics
//
// swagger:route PATCH /dynamicss/{ID} dynamicss updateDynamics
//
// # Update a dynamics
//
// Responses:
// default: genericError
//
//	200: dynamicsDBResponse
func (controller *Controller) UpdateDynamics(c *gin.Context) {

	mutexDynamics.Lock()
	defer mutexDynamics.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDynamics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDynamics.GetDB()

	// Validate input
	var input orm.DynamicsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var dynamicsDB orm.DynamicsDB

	// fetch the dynamics
	query := db.First(&dynamicsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	dynamicsDB.CopyBasicFieldsFromDynamics_WOP(&input.Dynamics_WOP)
	dynamicsDB.DynamicsPointersEncoding = input.DynamicsPointersEncoding

	query = db.Model(&dynamicsDB).Updates(dynamicsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	dynamicsNew := new(models.Dynamics)
	dynamicsDB.CopyBasicFieldsToDynamics(dynamicsNew)

	// redeem pointers
	dynamicsDB.DecodePointers(backRepo, dynamicsNew)

	// get stage instance from DB instance, and call callback function
	dynamicsOld := backRepo.BackRepoDynamics.Map_DynamicsDBID_DynamicsPtr[dynamicsDB.ID]
	if dynamicsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), dynamicsOld, dynamicsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the dynamicsDB
	c.JSON(http.StatusOK, dynamicsDB)
}

// DeleteDynamics
//
// swagger:route DELETE /dynamicss/{ID} dynamicss deleteDynamics
//
// # Delete a dynamics
//
// default: genericError
//
//	200: dynamicsDBResponse
func (controller *Controller) DeleteDynamics(c *gin.Context) {

	mutexDynamics.Lock()
	defer mutexDynamics.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDynamics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDynamics.GetDB()

	// Get model if exist
	var dynamicsDB orm.DynamicsDB
	if err := db.First(&dynamicsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&dynamicsDB)

	// get an instance (not staged) from DB instance, and call callback function
	dynamicsDeleted := new(models.Dynamics)
	dynamicsDB.CopyBasicFieldsToDynamics(dynamicsDeleted)

	// get stage instance from DB instance, and call callback function
	dynamicsStaged := backRepo.BackRepoDynamics.Map_DynamicsDBID_DynamicsPtr[dynamicsDB.ID]
	if dynamicsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), dynamicsStaged, dynamicsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
