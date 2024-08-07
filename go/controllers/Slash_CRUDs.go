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
var __Slash__dummysDeclaration__ models.Slash
var __Slash_time__dummyDeclaration time.Duration

var mutexSlash sync.Mutex

// An SlashID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSlash updateSlash deleteSlash
type SlashID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SlashInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSlash updateSlash
type SlashInput struct {
	// The Slash to submit or modify
	// in: body
	Slash *orm.SlashAPI
}

// GetSlashs
//
// swagger:route GET /slashs slashs getSlashs
//
// # Get all slashs
//
// Responses:
// default: genericError
//
//	200: slashDBResponse
func (controller *Controller) GetSlashs(c *gin.Context) {

	// source slice
	var slashDBs []orm.SlashDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlashs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlash.GetDB()

	query := db.Find(&slashDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	slashAPIs := make([]orm.SlashAPI, 0)

	// for each slash, update fields from the database nullable fields
	for idx := range slashDBs {
		slashDB := &slashDBs[idx]
		_ = slashDB
		var slashAPI orm.SlashAPI

		// insertion point for updating fields
		slashAPI.ID = slashDB.ID
		slashDB.CopyBasicFieldsToSlash_WOP(&slashAPI.Slash_WOP)
		slashAPI.SlashPointersEncoding = slashDB.SlashPointersEncoding
		slashAPIs = append(slashAPIs, slashAPI)
	}

	c.JSON(http.StatusOK, slashAPIs)
}

// PostSlash
//
// swagger:route POST /slashs slashs postSlash
//
// Creates a slash
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSlash(c *gin.Context) {

	mutexSlash.Lock()
	defer mutexSlash.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSlashs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlash.GetDB()

	// Validate input
	var input orm.SlashAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create slash
	slashDB := orm.SlashDB{}
	slashDB.SlashPointersEncoding = input.SlashPointersEncoding
	slashDB.CopyBasicFieldsFromSlash_WOP(&input.Slash_WOP)

	query := db.Create(&slashDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSlash.CheckoutPhaseOneInstance(&slashDB)
	slash := backRepo.BackRepoSlash.Map_SlashDBID_SlashPtr[slashDB.ID]

	if slash != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), slash)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, slashDB)
}

// GetSlash
//
// swagger:route GET /slashs/{ID} slashs getSlash
//
// Gets the details for a slash.
//
// Responses:
// default: genericError
//
//	200: slashDBResponse
func (controller *Controller) GetSlash(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlash", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlash.GetDB()

	// Get slashDB in DB
	var slashDB orm.SlashDB
	if err := db.First(&slashDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var slashAPI orm.SlashAPI
	slashAPI.ID = slashDB.ID
	slashAPI.SlashPointersEncoding = slashDB.SlashPointersEncoding
	slashDB.CopyBasicFieldsToSlash_WOP(&slashAPI.Slash_WOP)

	c.JSON(http.StatusOK, slashAPI)
}

// UpdateSlash
//
// swagger:route PATCH /slashs/{ID} slashs updateSlash
//
// # Update a slash
//
// Responses:
// default: genericError
//
//	200: slashDBResponse
func (controller *Controller) UpdateSlash(c *gin.Context) {

	mutexSlash.Lock()
	defer mutexSlash.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSlash", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlash.GetDB()

	// Validate input
	var input orm.SlashAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var slashDB orm.SlashDB

	// fetch the slash
	query := db.First(&slashDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	slashDB.CopyBasicFieldsFromSlash_WOP(&input.Slash_WOP)
	slashDB.SlashPointersEncoding = input.SlashPointersEncoding

	query = db.Model(&slashDB).Updates(slashDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	slashNew := new(models.Slash)
	slashDB.CopyBasicFieldsToSlash(slashNew)

	// redeem pointers
	slashDB.DecodePointers(backRepo, slashNew)

	// get stage instance from DB instance, and call callback function
	slashOld := backRepo.BackRepoSlash.Map_SlashDBID_SlashPtr[slashDB.ID]
	if slashOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), slashOld, slashNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the slashDB
	c.JSON(http.StatusOK, slashDB)
}

// DeleteSlash
//
// swagger:route DELETE /slashs/{ID} slashs deleteSlash
//
// # Delete a slash
//
// default: genericError
//
//	200: slashDBResponse
func (controller *Controller) DeleteSlash(c *gin.Context) {

	mutexSlash.Lock()
	defer mutexSlash.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSlash", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlash.GetDB()

	// Get model if exist
	var slashDB orm.SlashDB
	if err := db.First(&slashDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&slashDB)

	// get an instance (not staged) from DB instance, and call callback function
	slashDeleted := new(models.Slash)
	slashDB.CopyBasicFieldsToSlash(slashDeleted)

	// get stage instance from DB instance, and call callback function
	slashStaged := backRepo.BackRepoSlash.Map_SlashDBID_SlashPtr[slashDB.ID]
	if slashStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), slashStaged, slashDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
