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
var __Harmon_mute__dummysDeclaration__ models.Harmon_mute
var __Harmon_mute_time__dummyDeclaration time.Duration

var mutexHarmon_mute sync.Mutex

// An Harmon_muteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHarmon_mute updateHarmon_mute deleteHarmon_mute
type Harmon_muteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Harmon_muteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHarmon_mute updateHarmon_mute
type Harmon_muteInput struct {
	// The Harmon_mute to submit or modify
	// in: body
	Harmon_mute *orm.Harmon_muteAPI
}

// GetHarmon_mutes
//
// swagger:route GET /harmon_mutes harmon_mutes getHarmon_mutes
//
// # Get all harmon_mutes
//
// Responses:
// default: genericError
//
//	200: harmon_muteDBResponse
func (controller *Controller) GetHarmon_mutes(c *gin.Context) {

	// source slice
	var harmon_muteDBs []orm.Harmon_muteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmon_mutes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_mute.GetDB()

	query := db.Find(&harmon_muteDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	harmon_muteAPIs := make([]orm.Harmon_muteAPI, 0)

	// for each harmon_mute, update fields from the database nullable fields
	for idx := range harmon_muteDBs {
		harmon_muteDB := &harmon_muteDBs[idx]
		_ = harmon_muteDB
		var harmon_muteAPI orm.Harmon_muteAPI

		// insertion point for updating fields
		harmon_muteAPI.ID = harmon_muteDB.ID
		harmon_muteDB.CopyBasicFieldsToHarmon_mute_WOP(&harmon_muteAPI.Harmon_mute_WOP)
		harmon_muteAPI.Harmon_mutePointersEncoding = harmon_muteDB.Harmon_mutePointersEncoding
		harmon_muteAPIs = append(harmon_muteAPIs, harmon_muteAPI)
	}

	c.JSON(http.StatusOK, harmon_muteAPIs)
}

// PostHarmon_mute
//
// swagger:route POST /harmon_mutes harmon_mutes postHarmon_mute
//
// Creates a harmon_mute
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHarmon_mute(c *gin.Context) {

	mutexHarmon_mute.Lock()
	defer mutexHarmon_mute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHarmon_mutes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_mute.GetDB()

	// Validate input
	var input orm.Harmon_muteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create harmon_mute
	harmon_muteDB := orm.Harmon_muteDB{}
	harmon_muteDB.Harmon_mutePointersEncoding = input.Harmon_mutePointersEncoding
	harmon_muteDB.CopyBasicFieldsFromHarmon_mute_WOP(&input.Harmon_mute_WOP)

	query := db.Create(&harmon_muteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHarmon_mute.CheckoutPhaseOneInstance(&harmon_muteDB)
	harmon_mute := backRepo.BackRepoHarmon_mute.Map_Harmon_muteDBID_Harmon_mutePtr[harmon_muteDB.ID]

	if harmon_mute != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), harmon_mute)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, harmon_muteDB)
}

// GetHarmon_mute
//
// swagger:route GET /harmon_mutes/{ID} harmon_mutes getHarmon_mute
//
// Gets the details for a harmon_mute.
//
// Responses:
// default: genericError
//
//	200: harmon_muteDBResponse
func (controller *Controller) GetHarmon_mute(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmon_mute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_mute.GetDB()

	// Get harmon_muteDB in DB
	var harmon_muteDB orm.Harmon_muteDB
	if err := db.First(&harmon_muteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var harmon_muteAPI orm.Harmon_muteAPI
	harmon_muteAPI.ID = harmon_muteDB.ID
	harmon_muteAPI.Harmon_mutePointersEncoding = harmon_muteDB.Harmon_mutePointersEncoding
	harmon_muteDB.CopyBasicFieldsToHarmon_mute_WOP(&harmon_muteAPI.Harmon_mute_WOP)

	c.JSON(http.StatusOK, harmon_muteAPI)
}

// UpdateHarmon_mute
//
// swagger:route PATCH /harmon_mutes/{ID} harmon_mutes updateHarmon_mute
//
// # Update a harmon_mute
//
// Responses:
// default: genericError
//
//	200: harmon_muteDBResponse
func (controller *Controller) UpdateHarmon_mute(c *gin.Context) {

	mutexHarmon_mute.Lock()
	defer mutexHarmon_mute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHarmon_mute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_mute.GetDB()

	// Validate input
	var input orm.Harmon_muteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var harmon_muteDB orm.Harmon_muteDB

	// fetch the harmon_mute
	query := db.First(&harmon_muteDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	harmon_muteDB.CopyBasicFieldsFromHarmon_mute_WOP(&input.Harmon_mute_WOP)
	harmon_muteDB.Harmon_mutePointersEncoding = input.Harmon_mutePointersEncoding

	query = db.Model(&harmon_muteDB).Updates(harmon_muteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	harmon_muteNew := new(models.Harmon_mute)
	harmon_muteDB.CopyBasicFieldsToHarmon_mute(harmon_muteNew)

	// redeem pointers
	harmon_muteDB.DecodePointers(backRepo, harmon_muteNew)

	// get stage instance from DB instance, and call callback function
	harmon_muteOld := backRepo.BackRepoHarmon_mute.Map_Harmon_muteDBID_Harmon_mutePtr[harmon_muteDB.ID]
	if harmon_muteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), harmon_muteOld, harmon_muteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the harmon_muteDB
	c.JSON(http.StatusOK, harmon_muteDB)
}

// DeleteHarmon_mute
//
// swagger:route DELETE /harmon_mutes/{ID} harmon_mutes deleteHarmon_mute
//
// # Delete a harmon_mute
//
// default: genericError
//
//	200: harmon_muteDBResponse
func (controller *Controller) DeleteHarmon_mute(c *gin.Context) {

	mutexHarmon_mute.Lock()
	defer mutexHarmon_mute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHarmon_mute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmon_mute.GetDB()

	// Get model if exist
	var harmon_muteDB orm.Harmon_muteDB
	if err := db.First(&harmon_muteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&harmon_muteDB)

	// get an instance (not staged) from DB instance, and call callback function
	harmon_muteDeleted := new(models.Harmon_mute)
	harmon_muteDB.CopyBasicFieldsToHarmon_mute(harmon_muteDeleted)

	// get stage instance from DB instance, and call callback function
	harmon_muteStaged := backRepo.BackRepoHarmon_mute.Map_Harmon_muteDBID_Harmon_mutePtr[harmon_muteDB.ID]
	if harmon_muteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), harmon_muteStaged, harmon_muteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
