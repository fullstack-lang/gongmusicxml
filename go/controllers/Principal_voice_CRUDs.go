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
var __Principal_voice__dummysDeclaration__ models.Principal_voice
var __Principal_voice_time__dummyDeclaration time.Duration

var mutexPrincipal_voice sync.Mutex

// An Principal_voiceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPrincipal_voice updatePrincipal_voice deletePrincipal_voice
type Principal_voiceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Principal_voiceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPrincipal_voice updatePrincipal_voice
type Principal_voiceInput struct {
	// The Principal_voice to submit or modify
	// in: body
	Principal_voice *orm.Principal_voiceAPI
}

// GetPrincipal_voices
//
// swagger:route GET /principal_voices principal_voices getPrincipal_voices
//
// # Get all principal_voices
//
// Responses:
// default: genericError
//
//	200: principal_voiceDBResponse
func (controller *Controller) GetPrincipal_voices(c *gin.Context) {

	// source slice
	var principal_voiceDBs []orm.Principal_voiceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPrincipal_voices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrincipal_voice.GetDB()

	query := db.Find(&principal_voiceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	principal_voiceAPIs := make([]orm.Principal_voiceAPI, 0)

	// for each principal_voice, update fields from the database nullable fields
	for idx := range principal_voiceDBs {
		principal_voiceDB := &principal_voiceDBs[idx]
		_ = principal_voiceDB
		var principal_voiceAPI orm.Principal_voiceAPI

		// insertion point for updating fields
		principal_voiceAPI.ID = principal_voiceDB.ID
		principal_voiceDB.CopyBasicFieldsToPrincipal_voice_WOP(&principal_voiceAPI.Principal_voice_WOP)
		principal_voiceAPI.Principal_voicePointersEncoding = principal_voiceDB.Principal_voicePointersEncoding
		principal_voiceAPIs = append(principal_voiceAPIs, principal_voiceAPI)
	}

	c.JSON(http.StatusOK, principal_voiceAPIs)
}

// PostPrincipal_voice
//
// swagger:route POST /principal_voices principal_voices postPrincipal_voice
//
// Creates a principal_voice
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPrincipal_voice(c *gin.Context) {

	mutexPrincipal_voice.Lock()
	defer mutexPrincipal_voice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPrincipal_voices", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrincipal_voice.GetDB()

	// Validate input
	var input orm.Principal_voiceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create principal_voice
	principal_voiceDB := orm.Principal_voiceDB{}
	principal_voiceDB.Principal_voicePointersEncoding = input.Principal_voicePointersEncoding
	principal_voiceDB.CopyBasicFieldsFromPrincipal_voice_WOP(&input.Principal_voice_WOP)

	query := db.Create(&principal_voiceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPrincipal_voice.CheckoutPhaseOneInstance(&principal_voiceDB)
	principal_voice := backRepo.BackRepoPrincipal_voice.Map_Principal_voiceDBID_Principal_voicePtr[principal_voiceDB.ID]

	if principal_voice != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), principal_voice)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, principal_voiceDB)
}

// GetPrincipal_voice
//
// swagger:route GET /principal_voices/{ID} principal_voices getPrincipal_voice
//
// Gets the details for a principal_voice.
//
// Responses:
// default: genericError
//
//	200: principal_voiceDBResponse
func (controller *Controller) GetPrincipal_voice(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPrincipal_voice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrincipal_voice.GetDB()

	// Get principal_voiceDB in DB
	var principal_voiceDB orm.Principal_voiceDB
	if err := db.First(&principal_voiceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var principal_voiceAPI orm.Principal_voiceAPI
	principal_voiceAPI.ID = principal_voiceDB.ID
	principal_voiceAPI.Principal_voicePointersEncoding = principal_voiceDB.Principal_voicePointersEncoding
	principal_voiceDB.CopyBasicFieldsToPrincipal_voice_WOP(&principal_voiceAPI.Principal_voice_WOP)

	c.JSON(http.StatusOK, principal_voiceAPI)
}

// UpdatePrincipal_voice
//
// swagger:route PATCH /principal_voices/{ID} principal_voices updatePrincipal_voice
//
// # Update a principal_voice
//
// Responses:
// default: genericError
//
//	200: principal_voiceDBResponse
func (controller *Controller) UpdatePrincipal_voice(c *gin.Context) {

	mutexPrincipal_voice.Lock()
	defer mutexPrincipal_voice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePrincipal_voice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrincipal_voice.GetDB()

	// Validate input
	var input orm.Principal_voiceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var principal_voiceDB orm.Principal_voiceDB

	// fetch the principal_voice
	query := db.First(&principal_voiceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	principal_voiceDB.CopyBasicFieldsFromPrincipal_voice_WOP(&input.Principal_voice_WOP)
	principal_voiceDB.Principal_voicePointersEncoding = input.Principal_voicePointersEncoding

	query = db.Model(&principal_voiceDB).Updates(principal_voiceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	principal_voiceNew := new(models.Principal_voice)
	principal_voiceDB.CopyBasicFieldsToPrincipal_voice(principal_voiceNew)

	// redeem pointers
	principal_voiceDB.DecodePointers(backRepo, principal_voiceNew)

	// get stage instance from DB instance, and call callback function
	principal_voiceOld := backRepo.BackRepoPrincipal_voice.Map_Principal_voiceDBID_Principal_voicePtr[principal_voiceDB.ID]
	if principal_voiceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), principal_voiceOld, principal_voiceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the principal_voiceDB
	c.JSON(http.StatusOK, principal_voiceDB)
}

// DeletePrincipal_voice
//
// swagger:route DELETE /principal_voices/{ID} principal_voices deletePrincipal_voice
//
// # Delete a principal_voice
//
// default: genericError
//
//	200: principal_voiceDBResponse
func (controller *Controller) DeletePrincipal_voice(c *gin.Context) {

	mutexPrincipal_voice.Lock()
	defer mutexPrincipal_voice.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePrincipal_voice", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPrincipal_voice.GetDB()

	// Get model if exist
	var principal_voiceDB orm.Principal_voiceDB
	if err := db.First(&principal_voiceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&principal_voiceDB)

	// get an instance (not staged) from DB instance, and call callback function
	principal_voiceDeleted := new(models.Principal_voice)
	principal_voiceDB.CopyBasicFieldsToPrincipal_voice(principal_voiceDeleted)

	// get stage instance from DB instance, and call callback function
	principal_voiceStaged := backRepo.BackRepoPrincipal_voice.Map_Principal_voiceDBID_Principal_voicePtr[principal_voiceDB.ID]
	if principal_voiceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), principal_voiceStaged, principal_voiceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
