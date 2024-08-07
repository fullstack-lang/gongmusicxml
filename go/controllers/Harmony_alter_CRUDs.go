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
var __Harmony_alter__dummysDeclaration__ models.Harmony_alter
var __Harmony_alter_time__dummyDeclaration time.Duration

var mutexHarmony_alter sync.Mutex

// An Harmony_alterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHarmony_alter updateHarmony_alter deleteHarmony_alter
type Harmony_alterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Harmony_alterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHarmony_alter updateHarmony_alter
type Harmony_alterInput struct {
	// The Harmony_alter to submit or modify
	// in: body
	Harmony_alter *orm.Harmony_alterAPI
}

// GetHarmony_alters
//
// swagger:route GET /harmony_alters harmony_alters getHarmony_alters
//
// # Get all harmony_alters
//
// Responses:
// default: genericError
//
//	200: harmony_alterDBResponse
func (controller *Controller) GetHarmony_alters(c *gin.Context) {

	// source slice
	var harmony_alterDBs []orm.Harmony_alterDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmony_alters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony_alter.GetDB()

	query := db.Find(&harmony_alterDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	harmony_alterAPIs := make([]orm.Harmony_alterAPI, 0)

	// for each harmony_alter, update fields from the database nullable fields
	for idx := range harmony_alterDBs {
		harmony_alterDB := &harmony_alterDBs[idx]
		_ = harmony_alterDB
		var harmony_alterAPI orm.Harmony_alterAPI

		// insertion point for updating fields
		harmony_alterAPI.ID = harmony_alterDB.ID
		harmony_alterDB.CopyBasicFieldsToHarmony_alter_WOP(&harmony_alterAPI.Harmony_alter_WOP)
		harmony_alterAPI.Harmony_alterPointersEncoding = harmony_alterDB.Harmony_alterPointersEncoding
		harmony_alterAPIs = append(harmony_alterAPIs, harmony_alterAPI)
	}

	c.JSON(http.StatusOK, harmony_alterAPIs)
}

// PostHarmony_alter
//
// swagger:route POST /harmony_alters harmony_alters postHarmony_alter
//
// Creates a harmony_alter
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHarmony_alter(c *gin.Context) {

	mutexHarmony_alter.Lock()
	defer mutexHarmony_alter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHarmony_alters", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony_alter.GetDB()

	// Validate input
	var input orm.Harmony_alterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create harmony_alter
	harmony_alterDB := orm.Harmony_alterDB{}
	harmony_alterDB.Harmony_alterPointersEncoding = input.Harmony_alterPointersEncoding
	harmony_alterDB.CopyBasicFieldsFromHarmony_alter_WOP(&input.Harmony_alter_WOP)

	query := db.Create(&harmony_alterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHarmony_alter.CheckoutPhaseOneInstance(&harmony_alterDB)
	harmony_alter := backRepo.BackRepoHarmony_alter.Map_Harmony_alterDBID_Harmony_alterPtr[harmony_alterDB.ID]

	if harmony_alter != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), harmony_alter)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, harmony_alterDB)
}

// GetHarmony_alter
//
// swagger:route GET /harmony_alters/{ID} harmony_alters getHarmony_alter
//
// Gets the details for a harmony_alter.
//
// Responses:
// default: genericError
//
//	200: harmony_alterDBResponse
func (controller *Controller) GetHarmony_alter(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmony_alter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony_alter.GetDB()

	// Get harmony_alterDB in DB
	var harmony_alterDB orm.Harmony_alterDB
	if err := db.First(&harmony_alterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var harmony_alterAPI orm.Harmony_alterAPI
	harmony_alterAPI.ID = harmony_alterDB.ID
	harmony_alterAPI.Harmony_alterPointersEncoding = harmony_alterDB.Harmony_alterPointersEncoding
	harmony_alterDB.CopyBasicFieldsToHarmony_alter_WOP(&harmony_alterAPI.Harmony_alter_WOP)

	c.JSON(http.StatusOK, harmony_alterAPI)
}

// UpdateHarmony_alter
//
// swagger:route PATCH /harmony_alters/{ID} harmony_alters updateHarmony_alter
//
// # Update a harmony_alter
//
// Responses:
// default: genericError
//
//	200: harmony_alterDBResponse
func (controller *Controller) UpdateHarmony_alter(c *gin.Context) {

	mutexHarmony_alter.Lock()
	defer mutexHarmony_alter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHarmony_alter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony_alter.GetDB()

	// Validate input
	var input orm.Harmony_alterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var harmony_alterDB orm.Harmony_alterDB

	// fetch the harmony_alter
	query := db.First(&harmony_alterDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	harmony_alterDB.CopyBasicFieldsFromHarmony_alter_WOP(&input.Harmony_alter_WOP)
	harmony_alterDB.Harmony_alterPointersEncoding = input.Harmony_alterPointersEncoding

	query = db.Model(&harmony_alterDB).Updates(harmony_alterDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	harmony_alterNew := new(models.Harmony_alter)
	harmony_alterDB.CopyBasicFieldsToHarmony_alter(harmony_alterNew)

	// redeem pointers
	harmony_alterDB.DecodePointers(backRepo, harmony_alterNew)

	// get stage instance from DB instance, and call callback function
	harmony_alterOld := backRepo.BackRepoHarmony_alter.Map_Harmony_alterDBID_Harmony_alterPtr[harmony_alterDB.ID]
	if harmony_alterOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), harmony_alterOld, harmony_alterNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the harmony_alterDB
	c.JSON(http.StatusOK, harmony_alterDB)
}

// DeleteHarmony_alter
//
// swagger:route DELETE /harmony_alters/{ID} harmony_alters deleteHarmony_alter
//
// # Delete a harmony_alter
//
// default: genericError
//
//	200: harmony_alterDBResponse
func (controller *Controller) DeleteHarmony_alter(c *gin.Context) {

	mutexHarmony_alter.Lock()
	defer mutexHarmony_alter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHarmony_alter", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony_alter.GetDB()

	// Get model if exist
	var harmony_alterDB orm.Harmony_alterDB
	if err := db.First(&harmony_alterDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&harmony_alterDB)

	// get an instance (not staged) from DB instance, and call callback function
	harmony_alterDeleted := new(models.Harmony_alter)
	harmony_alterDB.CopyBasicFieldsToHarmony_alter(harmony_alterDeleted)

	// get stage instance from DB instance, and call callback function
	harmony_alterStaged := backRepo.BackRepoHarmony_alter.Map_Harmony_alterDBID_Harmony_alterPtr[harmony_alterDB.ID]
	if harmony_alterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), harmony_alterStaged, harmony_alterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
