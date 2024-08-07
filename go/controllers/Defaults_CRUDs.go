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
var __Defaults__dummysDeclaration__ models.Defaults
var __Defaults_time__dummyDeclaration time.Duration

var mutexDefaults sync.Mutex

// An DefaultsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDefaults updateDefaults deleteDefaults
type DefaultsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DefaultsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDefaults updateDefaults
type DefaultsInput struct {
	// The Defaults to submit or modify
	// in: body
	Defaults *orm.DefaultsAPI
}

// GetDefaultss
//
// swagger:route GET /defaultss defaultss getDefaultss
//
// # Get all defaultss
//
// Responses:
// default: genericError
//
//	200: defaultsDBResponse
func (controller *Controller) GetDefaultss(c *gin.Context) {

	// source slice
	var defaultsDBs []orm.DefaultsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDefaultss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDefaults.GetDB()

	query := db.Find(&defaultsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	defaultsAPIs := make([]orm.DefaultsAPI, 0)

	// for each defaults, update fields from the database nullable fields
	for idx := range defaultsDBs {
		defaultsDB := &defaultsDBs[idx]
		_ = defaultsDB
		var defaultsAPI orm.DefaultsAPI

		// insertion point for updating fields
		defaultsAPI.ID = defaultsDB.ID
		defaultsDB.CopyBasicFieldsToDefaults_WOP(&defaultsAPI.Defaults_WOP)
		defaultsAPI.DefaultsPointersEncoding = defaultsDB.DefaultsPointersEncoding
		defaultsAPIs = append(defaultsAPIs, defaultsAPI)
	}

	c.JSON(http.StatusOK, defaultsAPIs)
}

// PostDefaults
//
// swagger:route POST /defaultss defaultss postDefaults
//
// Creates a defaults
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDefaults(c *gin.Context) {

	mutexDefaults.Lock()
	defer mutexDefaults.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDefaultss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDefaults.GetDB()

	// Validate input
	var input orm.DefaultsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create defaults
	defaultsDB := orm.DefaultsDB{}
	defaultsDB.DefaultsPointersEncoding = input.DefaultsPointersEncoding
	defaultsDB.CopyBasicFieldsFromDefaults_WOP(&input.Defaults_WOP)

	query := db.Create(&defaultsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDefaults.CheckoutPhaseOneInstance(&defaultsDB)
	defaults := backRepo.BackRepoDefaults.Map_DefaultsDBID_DefaultsPtr[defaultsDB.ID]

	if defaults != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), defaults)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, defaultsDB)
}

// GetDefaults
//
// swagger:route GET /defaultss/{ID} defaultss getDefaults
//
// Gets the details for a defaults.
//
// Responses:
// default: genericError
//
//	200: defaultsDBResponse
func (controller *Controller) GetDefaults(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDefaults", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDefaults.GetDB()

	// Get defaultsDB in DB
	var defaultsDB orm.DefaultsDB
	if err := db.First(&defaultsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var defaultsAPI orm.DefaultsAPI
	defaultsAPI.ID = defaultsDB.ID
	defaultsAPI.DefaultsPointersEncoding = defaultsDB.DefaultsPointersEncoding
	defaultsDB.CopyBasicFieldsToDefaults_WOP(&defaultsAPI.Defaults_WOP)

	c.JSON(http.StatusOK, defaultsAPI)
}

// UpdateDefaults
//
// swagger:route PATCH /defaultss/{ID} defaultss updateDefaults
//
// # Update a defaults
//
// Responses:
// default: genericError
//
//	200: defaultsDBResponse
func (controller *Controller) UpdateDefaults(c *gin.Context) {

	mutexDefaults.Lock()
	defer mutexDefaults.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDefaults", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDefaults.GetDB()

	// Validate input
	var input orm.DefaultsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var defaultsDB orm.DefaultsDB

	// fetch the defaults
	query := db.First(&defaultsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	defaultsDB.CopyBasicFieldsFromDefaults_WOP(&input.Defaults_WOP)
	defaultsDB.DefaultsPointersEncoding = input.DefaultsPointersEncoding

	query = db.Model(&defaultsDB).Updates(defaultsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	defaultsNew := new(models.Defaults)
	defaultsDB.CopyBasicFieldsToDefaults(defaultsNew)

	// redeem pointers
	defaultsDB.DecodePointers(backRepo, defaultsNew)

	// get stage instance from DB instance, and call callback function
	defaultsOld := backRepo.BackRepoDefaults.Map_DefaultsDBID_DefaultsPtr[defaultsDB.ID]
	if defaultsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), defaultsOld, defaultsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the defaultsDB
	c.JSON(http.StatusOK, defaultsDB)
}

// DeleteDefaults
//
// swagger:route DELETE /defaultss/{ID} defaultss deleteDefaults
//
// # Delete a defaults
//
// default: genericError
//
//	200: defaultsDBResponse
func (controller *Controller) DeleteDefaults(c *gin.Context) {

	mutexDefaults.Lock()
	defer mutexDefaults.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDefaults", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDefaults.GetDB()

	// Get model if exist
	var defaultsDB orm.DefaultsDB
	if err := db.First(&defaultsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&defaultsDB)

	// get an instance (not staged) from DB instance, and call callback function
	defaultsDeleted := new(models.Defaults)
	defaultsDB.CopyBasicFieldsToDefaults(defaultsDeleted)

	// get stage instance from DB instance, and call callback function
	defaultsStaged := backRepo.BackRepoDefaults.Map_DefaultsDBID_DefaultsPtr[defaultsDB.ID]
	if defaultsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), defaultsStaged, defaultsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
