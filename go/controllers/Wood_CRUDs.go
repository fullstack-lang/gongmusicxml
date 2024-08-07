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
var __Wood__dummysDeclaration__ models.Wood
var __Wood_time__dummyDeclaration time.Duration

var mutexWood sync.Mutex

// An WoodID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getWood updateWood deleteWood
type WoodID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// WoodInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postWood updateWood
type WoodInput struct {
	// The Wood to submit or modify
	// in: body
	Wood *orm.WoodAPI
}

// GetWoods
//
// swagger:route GET /woods woods getWoods
//
// # Get all woods
//
// Responses:
// default: genericError
//
//	200: woodDBResponse
func (controller *Controller) GetWoods(c *gin.Context) {

	// source slice
	var woodDBs []orm.WoodDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWoods", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWood.GetDB()

	query := db.Find(&woodDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	woodAPIs := make([]orm.WoodAPI, 0)

	// for each wood, update fields from the database nullable fields
	for idx := range woodDBs {
		woodDB := &woodDBs[idx]
		_ = woodDB
		var woodAPI orm.WoodAPI

		// insertion point for updating fields
		woodAPI.ID = woodDB.ID
		woodDB.CopyBasicFieldsToWood_WOP(&woodAPI.Wood_WOP)
		woodAPI.WoodPointersEncoding = woodDB.WoodPointersEncoding
		woodAPIs = append(woodAPIs, woodAPI)
	}

	c.JSON(http.StatusOK, woodAPIs)
}

// PostWood
//
// swagger:route POST /woods woods postWood
//
// Creates a wood
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostWood(c *gin.Context) {

	mutexWood.Lock()
	defer mutexWood.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostWoods", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWood.GetDB()

	// Validate input
	var input orm.WoodAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create wood
	woodDB := orm.WoodDB{}
	woodDB.WoodPointersEncoding = input.WoodPointersEncoding
	woodDB.CopyBasicFieldsFromWood_WOP(&input.Wood_WOP)

	query := db.Create(&woodDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoWood.CheckoutPhaseOneInstance(&woodDB)
	wood := backRepo.BackRepoWood.Map_WoodDBID_WoodPtr[woodDB.ID]

	if wood != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), wood)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, woodDB)
}

// GetWood
//
// swagger:route GET /woods/{ID} woods getWood
//
// Gets the details for a wood.
//
// Responses:
// default: genericError
//
//	200: woodDBResponse
func (controller *Controller) GetWood(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetWood", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWood.GetDB()

	// Get woodDB in DB
	var woodDB orm.WoodDB
	if err := db.First(&woodDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var woodAPI orm.WoodAPI
	woodAPI.ID = woodDB.ID
	woodAPI.WoodPointersEncoding = woodDB.WoodPointersEncoding
	woodDB.CopyBasicFieldsToWood_WOP(&woodAPI.Wood_WOP)

	c.JSON(http.StatusOK, woodAPI)
}

// UpdateWood
//
// swagger:route PATCH /woods/{ID} woods updateWood
//
// # Update a wood
//
// Responses:
// default: genericError
//
//	200: woodDBResponse
func (controller *Controller) UpdateWood(c *gin.Context) {

	mutexWood.Lock()
	defer mutexWood.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateWood", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWood.GetDB()

	// Validate input
	var input orm.WoodAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var woodDB orm.WoodDB

	// fetch the wood
	query := db.First(&woodDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	woodDB.CopyBasicFieldsFromWood_WOP(&input.Wood_WOP)
	woodDB.WoodPointersEncoding = input.WoodPointersEncoding

	query = db.Model(&woodDB).Updates(woodDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	woodNew := new(models.Wood)
	woodDB.CopyBasicFieldsToWood(woodNew)

	// redeem pointers
	woodDB.DecodePointers(backRepo, woodNew)

	// get stage instance from DB instance, and call callback function
	woodOld := backRepo.BackRepoWood.Map_WoodDBID_WoodPtr[woodDB.ID]
	if woodOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), woodOld, woodNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the woodDB
	c.JSON(http.StatusOK, woodDB)
}

// DeleteWood
//
// swagger:route DELETE /woods/{ID} woods deleteWood
//
// # Delete a wood
//
// default: genericError
//
//	200: woodDBResponse
func (controller *Controller) DeleteWood(c *gin.Context) {

	mutexWood.Lock()
	defer mutexWood.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteWood", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoWood.GetDB()

	// Get model if exist
	var woodDB orm.WoodDB
	if err := db.First(&woodDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&woodDB)

	// get an instance (not staged) from DB instance, and call callback function
	woodDeleted := new(models.Wood)
	woodDB.CopyBasicFieldsToWood(woodDeleted)

	// get stage instance from DB instance, and call callback function
	woodStaged := backRepo.BackRepoWood.Map_WoodDBID_WoodPtr[woodDB.ID]
	if woodStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), woodStaged, woodDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
