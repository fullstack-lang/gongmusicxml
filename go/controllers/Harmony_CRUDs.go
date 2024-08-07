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
var __Harmony__dummysDeclaration__ models.Harmony
var __Harmony_time__dummyDeclaration time.Duration

var mutexHarmony sync.Mutex

// An HarmonyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHarmony updateHarmony deleteHarmony
type HarmonyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HarmonyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHarmony updateHarmony
type HarmonyInput struct {
	// The Harmony to submit or modify
	// in: body
	Harmony *orm.HarmonyAPI
}

// GetHarmonys
//
// swagger:route GET /harmonys harmonys getHarmonys
//
// # Get all harmonys
//
// Responses:
// default: genericError
//
//	200: harmonyDBResponse
func (controller *Controller) GetHarmonys(c *gin.Context) {

	// source slice
	var harmonyDBs []orm.HarmonyDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmonys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony.GetDB()

	query := db.Find(&harmonyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	harmonyAPIs := make([]orm.HarmonyAPI, 0)

	// for each harmony, update fields from the database nullable fields
	for idx := range harmonyDBs {
		harmonyDB := &harmonyDBs[idx]
		_ = harmonyDB
		var harmonyAPI orm.HarmonyAPI

		// insertion point for updating fields
		harmonyAPI.ID = harmonyDB.ID
		harmonyDB.CopyBasicFieldsToHarmony_WOP(&harmonyAPI.Harmony_WOP)
		harmonyAPI.HarmonyPointersEncoding = harmonyDB.HarmonyPointersEncoding
		harmonyAPIs = append(harmonyAPIs, harmonyAPI)
	}

	c.JSON(http.StatusOK, harmonyAPIs)
}

// PostHarmony
//
// swagger:route POST /harmonys harmonys postHarmony
//
// Creates a harmony
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHarmony(c *gin.Context) {

	mutexHarmony.Lock()
	defer mutexHarmony.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHarmonys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony.GetDB()

	// Validate input
	var input orm.HarmonyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create harmony
	harmonyDB := orm.HarmonyDB{}
	harmonyDB.HarmonyPointersEncoding = input.HarmonyPointersEncoding
	harmonyDB.CopyBasicFieldsFromHarmony_WOP(&input.Harmony_WOP)

	query := db.Create(&harmonyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHarmony.CheckoutPhaseOneInstance(&harmonyDB)
	harmony := backRepo.BackRepoHarmony.Map_HarmonyDBID_HarmonyPtr[harmonyDB.ID]

	if harmony != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), harmony)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, harmonyDB)
}

// GetHarmony
//
// swagger:route GET /harmonys/{ID} harmonys getHarmony
//
// Gets the details for a harmony.
//
// Responses:
// default: genericError
//
//	200: harmonyDBResponse
func (controller *Controller) GetHarmony(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmony", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony.GetDB()

	// Get harmonyDB in DB
	var harmonyDB orm.HarmonyDB
	if err := db.First(&harmonyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var harmonyAPI orm.HarmonyAPI
	harmonyAPI.ID = harmonyDB.ID
	harmonyAPI.HarmonyPointersEncoding = harmonyDB.HarmonyPointersEncoding
	harmonyDB.CopyBasicFieldsToHarmony_WOP(&harmonyAPI.Harmony_WOP)

	c.JSON(http.StatusOK, harmonyAPI)
}

// UpdateHarmony
//
// swagger:route PATCH /harmonys/{ID} harmonys updateHarmony
//
// # Update a harmony
//
// Responses:
// default: genericError
//
//	200: harmonyDBResponse
func (controller *Controller) UpdateHarmony(c *gin.Context) {

	mutexHarmony.Lock()
	defer mutexHarmony.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHarmony", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony.GetDB()

	// Validate input
	var input orm.HarmonyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var harmonyDB orm.HarmonyDB

	// fetch the harmony
	query := db.First(&harmonyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	harmonyDB.CopyBasicFieldsFromHarmony_WOP(&input.Harmony_WOP)
	harmonyDB.HarmonyPointersEncoding = input.HarmonyPointersEncoding

	query = db.Model(&harmonyDB).Updates(harmonyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	harmonyNew := new(models.Harmony)
	harmonyDB.CopyBasicFieldsToHarmony(harmonyNew)

	// redeem pointers
	harmonyDB.DecodePointers(backRepo, harmonyNew)

	// get stage instance from DB instance, and call callback function
	harmonyOld := backRepo.BackRepoHarmony.Map_HarmonyDBID_HarmonyPtr[harmonyDB.ID]
	if harmonyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), harmonyOld, harmonyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the harmonyDB
	c.JSON(http.StatusOK, harmonyDB)
}

// DeleteHarmony
//
// swagger:route DELETE /harmonys/{ID} harmonys deleteHarmony
//
// # Delete a harmony
//
// default: genericError
//
//	200: harmonyDBResponse
func (controller *Controller) DeleteHarmony(c *gin.Context) {

	mutexHarmony.Lock()
	defer mutexHarmony.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHarmony", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmony.GetDB()

	// Get model if exist
	var harmonyDB orm.HarmonyDB
	if err := db.First(&harmonyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&harmonyDB)

	// get an instance (not staged) from DB instance, and call callback function
	harmonyDeleted := new(models.Harmony)
	harmonyDB.CopyBasicFieldsToHarmony(harmonyDeleted)

	// get stage instance from DB instance, and call callback function
	harmonyStaged := backRepo.BackRepoHarmony.Map_HarmonyDBID_HarmonyPtr[harmonyDB.ID]
	if harmonyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), harmonyStaged, harmonyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
