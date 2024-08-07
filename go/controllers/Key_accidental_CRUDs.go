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
var __Key_accidental__dummysDeclaration__ models.Key_accidental
var __Key_accidental_time__dummyDeclaration time.Duration

var mutexKey_accidental sync.Mutex

// An Key_accidentalID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKey_accidental updateKey_accidental deleteKey_accidental
type Key_accidentalID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Key_accidentalInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKey_accidental updateKey_accidental
type Key_accidentalInput struct {
	// The Key_accidental to submit or modify
	// in: body
	Key_accidental *orm.Key_accidentalAPI
}

// GetKey_accidentals
//
// swagger:route GET /key_accidentals key_accidentals getKey_accidentals
//
// # Get all key_accidentals
//
// Responses:
// default: genericError
//
//	200: key_accidentalDBResponse
func (controller *Controller) GetKey_accidentals(c *gin.Context) {

	// source slice
	var key_accidentalDBs []orm.Key_accidentalDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKey_accidentals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_accidental.GetDB()

	query := db.Find(&key_accidentalDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	key_accidentalAPIs := make([]orm.Key_accidentalAPI, 0)

	// for each key_accidental, update fields from the database nullable fields
	for idx := range key_accidentalDBs {
		key_accidentalDB := &key_accidentalDBs[idx]
		_ = key_accidentalDB
		var key_accidentalAPI orm.Key_accidentalAPI

		// insertion point for updating fields
		key_accidentalAPI.ID = key_accidentalDB.ID
		key_accidentalDB.CopyBasicFieldsToKey_accidental_WOP(&key_accidentalAPI.Key_accidental_WOP)
		key_accidentalAPI.Key_accidentalPointersEncoding = key_accidentalDB.Key_accidentalPointersEncoding
		key_accidentalAPIs = append(key_accidentalAPIs, key_accidentalAPI)
	}

	c.JSON(http.StatusOK, key_accidentalAPIs)
}

// PostKey_accidental
//
// swagger:route POST /key_accidentals key_accidentals postKey_accidental
//
// Creates a key_accidental
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKey_accidental(c *gin.Context) {

	mutexKey_accidental.Lock()
	defer mutexKey_accidental.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKey_accidentals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_accidental.GetDB()

	// Validate input
	var input orm.Key_accidentalAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create key_accidental
	key_accidentalDB := orm.Key_accidentalDB{}
	key_accidentalDB.Key_accidentalPointersEncoding = input.Key_accidentalPointersEncoding
	key_accidentalDB.CopyBasicFieldsFromKey_accidental_WOP(&input.Key_accidental_WOP)

	query := db.Create(&key_accidentalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKey_accidental.CheckoutPhaseOneInstance(&key_accidentalDB)
	key_accidental := backRepo.BackRepoKey_accidental.Map_Key_accidentalDBID_Key_accidentalPtr[key_accidentalDB.ID]

	if key_accidental != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), key_accidental)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, key_accidentalDB)
}

// GetKey_accidental
//
// swagger:route GET /key_accidentals/{ID} key_accidentals getKey_accidental
//
// Gets the details for a key_accidental.
//
// Responses:
// default: genericError
//
//	200: key_accidentalDBResponse
func (controller *Controller) GetKey_accidental(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKey_accidental", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_accidental.GetDB()

	// Get key_accidentalDB in DB
	var key_accidentalDB orm.Key_accidentalDB
	if err := db.First(&key_accidentalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var key_accidentalAPI orm.Key_accidentalAPI
	key_accidentalAPI.ID = key_accidentalDB.ID
	key_accidentalAPI.Key_accidentalPointersEncoding = key_accidentalDB.Key_accidentalPointersEncoding
	key_accidentalDB.CopyBasicFieldsToKey_accidental_WOP(&key_accidentalAPI.Key_accidental_WOP)

	c.JSON(http.StatusOK, key_accidentalAPI)
}

// UpdateKey_accidental
//
// swagger:route PATCH /key_accidentals/{ID} key_accidentals updateKey_accidental
//
// # Update a key_accidental
//
// Responses:
// default: genericError
//
//	200: key_accidentalDBResponse
func (controller *Controller) UpdateKey_accidental(c *gin.Context) {

	mutexKey_accidental.Lock()
	defer mutexKey_accidental.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKey_accidental", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_accidental.GetDB()

	// Validate input
	var input orm.Key_accidentalAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var key_accidentalDB orm.Key_accidentalDB

	// fetch the key_accidental
	query := db.First(&key_accidentalDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	key_accidentalDB.CopyBasicFieldsFromKey_accidental_WOP(&input.Key_accidental_WOP)
	key_accidentalDB.Key_accidentalPointersEncoding = input.Key_accidentalPointersEncoding

	query = db.Model(&key_accidentalDB).Updates(key_accidentalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	key_accidentalNew := new(models.Key_accidental)
	key_accidentalDB.CopyBasicFieldsToKey_accidental(key_accidentalNew)

	// redeem pointers
	key_accidentalDB.DecodePointers(backRepo, key_accidentalNew)

	// get stage instance from DB instance, and call callback function
	key_accidentalOld := backRepo.BackRepoKey_accidental.Map_Key_accidentalDBID_Key_accidentalPtr[key_accidentalDB.ID]
	if key_accidentalOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), key_accidentalOld, key_accidentalNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the key_accidentalDB
	c.JSON(http.StatusOK, key_accidentalDB)
}

// DeleteKey_accidental
//
// swagger:route DELETE /key_accidentals/{ID} key_accidentals deleteKey_accidental
//
// # Delete a key_accidental
//
// default: genericError
//
//	200: key_accidentalDBResponse
func (controller *Controller) DeleteKey_accidental(c *gin.Context) {

	mutexKey_accidental.Lock()
	defer mutexKey_accidental.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKey_accidental", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_accidental.GetDB()

	// Get model if exist
	var key_accidentalDB orm.Key_accidentalDB
	if err := db.First(&key_accidentalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&key_accidentalDB)

	// get an instance (not staged) from DB instance, and call callback function
	key_accidentalDeleted := new(models.Key_accidental)
	key_accidentalDB.CopyBasicFieldsToKey_accidental(key_accidentalDeleted)

	// get stage instance from DB instance, and call callback function
	key_accidentalStaged := backRepo.BackRepoKey_accidental.Map_Key_accidentalDBID_Key_accidentalPtr[key_accidentalDB.ID]
	if key_accidentalStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), key_accidentalStaged, key_accidentalDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
