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
var __String_mute__dummysDeclaration__ models.String_mute
var __String_mute_time__dummyDeclaration time.Duration

var mutexString_mute sync.Mutex

// An String_muteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getString_mute updateString_mute deleteString_mute
type String_muteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// String_muteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postString_mute updateString_mute
type String_muteInput struct {
	// The String_mute to submit or modify
	// in: body
	String_mute *orm.String_muteAPI
}

// GetString_mutes
//
// swagger:route GET /string_mutes string_mutes getString_mutes
//
// # Get all string_mutes
//
// Responses:
// default: genericError
//
//	200: string_muteDBResponse
func (controller *Controller) GetString_mutes(c *gin.Context) {

	// source slice
	var string_muteDBs []orm.String_muteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetString_mutes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoString_mute.GetDB()

	query := db.Find(&string_muteDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	string_muteAPIs := make([]orm.String_muteAPI, 0)

	// for each string_mute, update fields from the database nullable fields
	for idx := range string_muteDBs {
		string_muteDB := &string_muteDBs[idx]
		_ = string_muteDB
		var string_muteAPI orm.String_muteAPI

		// insertion point for updating fields
		string_muteAPI.ID = string_muteDB.ID
		string_muteDB.CopyBasicFieldsToString_mute_WOP(&string_muteAPI.String_mute_WOP)
		string_muteAPI.String_mutePointersEncoding = string_muteDB.String_mutePointersEncoding
		string_muteAPIs = append(string_muteAPIs, string_muteAPI)
	}

	c.JSON(http.StatusOK, string_muteAPIs)
}

// PostString_mute
//
// swagger:route POST /string_mutes string_mutes postString_mute
//
// Creates a string_mute
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostString_mute(c *gin.Context) {

	mutexString_mute.Lock()
	defer mutexString_mute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostString_mutes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoString_mute.GetDB()

	// Validate input
	var input orm.String_muteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create string_mute
	string_muteDB := orm.String_muteDB{}
	string_muteDB.String_mutePointersEncoding = input.String_mutePointersEncoding
	string_muteDB.CopyBasicFieldsFromString_mute_WOP(&input.String_mute_WOP)

	query := db.Create(&string_muteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoString_mute.CheckoutPhaseOneInstance(&string_muteDB)
	string_mute := backRepo.BackRepoString_mute.Map_String_muteDBID_String_mutePtr[string_muteDB.ID]

	if string_mute != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), string_mute)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, string_muteDB)
}

// GetString_mute
//
// swagger:route GET /string_mutes/{ID} string_mutes getString_mute
//
// Gets the details for a string_mute.
//
// Responses:
// default: genericError
//
//	200: string_muteDBResponse
func (controller *Controller) GetString_mute(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetString_mute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoString_mute.GetDB()

	// Get string_muteDB in DB
	var string_muteDB orm.String_muteDB
	if err := db.First(&string_muteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var string_muteAPI orm.String_muteAPI
	string_muteAPI.ID = string_muteDB.ID
	string_muteAPI.String_mutePointersEncoding = string_muteDB.String_mutePointersEncoding
	string_muteDB.CopyBasicFieldsToString_mute_WOP(&string_muteAPI.String_mute_WOP)

	c.JSON(http.StatusOK, string_muteAPI)
}

// UpdateString_mute
//
// swagger:route PATCH /string_mutes/{ID} string_mutes updateString_mute
//
// # Update a string_mute
//
// Responses:
// default: genericError
//
//	200: string_muteDBResponse
func (controller *Controller) UpdateString_mute(c *gin.Context) {

	mutexString_mute.Lock()
	defer mutexString_mute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateString_mute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoString_mute.GetDB()

	// Validate input
	var input orm.String_muteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var string_muteDB orm.String_muteDB

	// fetch the string_mute
	query := db.First(&string_muteDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	string_muteDB.CopyBasicFieldsFromString_mute_WOP(&input.String_mute_WOP)
	string_muteDB.String_mutePointersEncoding = input.String_mutePointersEncoding

	query = db.Model(&string_muteDB).Updates(string_muteDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	string_muteNew := new(models.String_mute)
	string_muteDB.CopyBasicFieldsToString_mute(string_muteNew)

	// redeem pointers
	string_muteDB.DecodePointers(backRepo, string_muteNew)

	// get stage instance from DB instance, and call callback function
	string_muteOld := backRepo.BackRepoString_mute.Map_String_muteDBID_String_mutePtr[string_muteDB.ID]
	if string_muteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), string_muteOld, string_muteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the string_muteDB
	c.JSON(http.StatusOK, string_muteDB)
}

// DeleteString_mute
//
// swagger:route DELETE /string_mutes/{ID} string_mutes deleteString_mute
//
// # Delete a string_mute
//
// default: genericError
//
//	200: string_muteDBResponse
func (controller *Controller) DeleteString_mute(c *gin.Context) {

	mutexString_mute.Lock()
	defer mutexString_mute.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteString_mute", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoString_mute.GetDB()

	// Get model if exist
	var string_muteDB orm.String_muteDB
	if err := db.First(&string_muteDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&string_muteDB)

	// get an instance (not staged) from DB instance, and call callback function
	string_muteDeleted := new(models.String_mute)
	string_muteDB.CopyBasicFieldsToString_mute(string_muteDeleted)

	// get stage instance from DB instance, and call callback function
	string_muteStaged := backRepo.BackRepoString_mute.Map_String_muteDBID_String_mutePtr[string_muteDB.ID]
	if string_muteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), string_muteStaged, string_muteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
