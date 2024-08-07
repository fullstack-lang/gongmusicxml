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
var __Appearance__dummysDeclaration__ models.Appearance
var __Appearance_time__dummyDeclaration time.Duration

var mutexAppearance sync.Mutex

// An AppearanceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAppearance updateAppearance deleteAppearance
type AppearanceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AppearanceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAppearance updateAppearance
type AppearanceInput struct {
	// The Appearance to submit or modify
	// in: body
	Appearance *orm.AppearanceAPI
}

// GetAppearances
//
// swagger:route GET /appearances appearances getAppearances
//
// # Get all appearances
//
// Responses:
// default: genericError
//
//	200: appearanceDBResponse
func (controller *Controller) GetAppearances(c *gin.Context) {

	// source slice
	var appearanceDBs []orm.AppearanceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAppearances", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAppearance.GetDB()

	query := db.Find(&appearanceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	appearanceAPIs := make([]orm.AppearanceAPI, 0)

	// for each appearance, update fields from the database nullable fields
	for idx := range appearanceDBs {
		appearanceDB := &appearanceDBs[idx]
		_ = appearanceDB
		var appearanceAPI orm.AppearanceAPI

		// insertion point for updating fields
		appearanceAPI.ID = appearanceDB.ID
		appearanceDB.CopyBasicFieldsToAppearance_WOP(&appearanceAPI.Appearance_WOP)
		appearanceAPI.AppearancePointersEncoding = appearanceDB.AppearancePointersEncoding
		appearanceAPIs = append(appearanceAPIs, appearanceAPI)
	}

	c.JSON(http.StatusOK, appearanceAPIs)
}

// PostAppearance
//
// swagger:route POST /appearances appearances postAppearance
//
// Creates a appearance
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAppearance(c *gin.Context) {

	mutexAppearance.Lock()
	defer mutexAppearance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAppearances", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAppearance.GetDB()

	// Validate input
	var input orm.AppearanceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create appearance
	appearanceDB := orm.AppearanceDB{}
	appearanceDB.AppearancePointersEncoding = input.AppearancePointersEncoding
	appearanceDB.CopyBasicFieldsFromAppearance_WOP(&input.Appearance_WOP)

	query := db.Create(&appearanceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAppearance.CheckoutPhaseOneInstance(&appearanceDB)
	appearance := backRepo.BackRepoAppearance.Map_AppearanceDBID_AppearancePtr[appearanceDB.ID]

	if appearance != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), appearance)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, appearanceDB)
}

// GetAppearance
//
// swagger:route GET /appearances/{ID} appearances getAppearance
//
// Gets the details for a appearance.
//
// Responses:
// default: genericError
//
//	200: appearanceDBResponse
func (controller *Controller) GetAppearance(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAppearance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAppearance.GetDB()

	// Get appearanceDB in DB
	var appearanceDB orm.AppearanceDB
	if err := db.First(&appearanceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var appearanceAPI orm.AppearanceAPI
	appearanceAPI.ID = appearanceDB.ID
	appearanceAPI.AppearancePointersEncoding = appearanceDB.AppearancePointersEncoding
	appearanceDB.CopyBasicFieldsToAppearance_WOP(&appearanceAPI.Appearance_WOP)

	c.JSON(http.StatusOK, appearanceAPI)
}

// UpdateAppearance
//
// swagger:route PATCH /appearances/{ID} appearances updateAppearance
//
// # Update a appearance
//
// Responses:
// default: genericError
//
//	200: appearanceDBResponse
func (controller *Controller) UpdateAppearance(c *gin.Context) {

	mutexAppearance.Lock()
	defer mutexAppearance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAppearance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAppearance.GetDB()

	// Validate input
	var input orm.AppearanceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var appearanceDB orm.AppearanceDB

	// fetch the appearance
	query := db.First(&appearanceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	appearanceDB.CopyBasicFieldsFromAppearance_WOP(&input.Appearance_WOP)
	appearanceDB.AppearancePointersEncoding = input.AppearancePointersEncoding

	query = db.Model(&appearanceDB).Updates(appearanceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	appearanceNew := new(models.Appearance)
	appearanceDB.CopyBasicFieldsToAppearance(appearanceNew)

	// redeem pointers
	appearanceDB.DecodePointers(backRepo, appearanceNew)

	// get stage instance from DB instance, and call callback function
	appearanceOld := backRepo.BackRepoAppearance.Map_AppearanceDBID_AppearancePtr[appearanceDB.ID]
	if appearanceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), appearanceOld, appearanceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the appearanceDB
	c.JSON(http.StatusOK, appearanceDB)
}

// DeleteAppearance
//
// swagger:route DELETE /appearances/{ID} appearances deleteAppearance
//
// # Delete a appearance
//
// default: genericError
//
//	200: appearanceDBResponse
func (controller *Controller) DeleteAppearance(c *gin.Context) {

	mutexAppearance.Lock()
	defer mutexAppearance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAppearance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAppearance.GetDB()

	// Get model if exist
	var appearanceDB orm.AppearanceDB
	if err := db.First(&appearanceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&appearanceDB)

	// get an instance (not staged) from DB instance, and call callback function
	appearanceDeleted := new(models.Appearance)
	appearanceDB.CopyBasicFieldsToAppearance(appearanceDeleted)

	// get stage instance from DB instance, and call callback function
	appearanceStaged := backRepo.BackRepoAppearance.Map_AppearanceDBID_AppearancePtr[appearanceDB.ID]
	if appearanceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), appearanceStaged, appearanceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
