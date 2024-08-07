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
var __Key_octave__dummysDeclaration__ models.Key_octave
var __Key_octave_time__dummyDeclaration time.Duration

var mutexKey_octave sync.Mutex

// An Key_octaveID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getKey_octave updateKey_octave deleteKey_octave
type Key_octaveID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Key_octaveInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postKey_octave updateKey_octave
type Key_octaveInput struct {
	// The Key_octave to submit or modify
	// in: body
	Key_octave *orm.Key_octaveAPI
}

// GetKey_octaves
//
// swagger:route GET /key_octaves key_octaves getKey_octaves
//
// # Get all key_octaves
//
// Responses:
// default: genericError
//
//	200: key_octaveDBResponse
func (controller *Controller) GetKey_octaves(c *gin.Context) {

	// source slice
	var key_octaveDBs []orm.Key_octaveDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKey_octaves", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_octave.GetDB()

	query := db.Find(&key_octaveDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	key_octaveAPIs := make([]orm.Key_octaveAPI, 0)

	// for each key_octave, update fields from the database nullable fields
	for idx := range key_octaveDBs {
		key_octaveDB := &key_octaveDBs[idx]
		_ = key_octaveDB
		var key_octaveAPI orm.Key_octaveAPI

		// insertion point for updating fields
		key_octaveAPI.ID = key_octaveDB.ID
		key_octaveDB.CopyBasicFieldsToKey_octave_WOP(&key_octaveAPI.Key_octave_WOP)
		key_octaveAPI.Key_octavePointersEncoding = key_octaveDB.Key_octavePointersEncoding
		key_octaveAPIs = append(key_octaveAPIs, key_octaveAPI)
	}

	c.JSON(http.StatusOK, key_octaveAPIs)
}

// PostKey_octave
//
// swagger:route POST /key_octaves key_octaves postKey_octave
//
// Creates a key_octave
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostKey_octave(c *gin.Context) {

	mutexKey_octave.Lock()
	defer mutexKey_octave.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostKey_octaves", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_octave.GetDB()

	// Validate input
	var input orm.Key_octaveAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create key_octave
	key_octaveDB := orm.Key_octaveDB{}
	key_octaveDB.Key_octavePointersEncoding = input.Key_octavePointersEncoding
	key_octaveDB.CopyBasicFieldsFromKey_octave_WOP(&input.Key_octave_WOP)

	query := db.Create(&key_octaveDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoKey_octave.CheckoutPhaseOneInstance(&key_octaveDB)
	key_octave := backRepo.BackRepoKey_octave.Map_Key_octaveDBID_Key_octavePtr[key_octaveDB.ID]

	if key_octave != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), key_octave)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, key_octaveDB)
}

// GetKey_octave
//
// swagger:route GET /key_octaves/{ID} key_octaves getKey_octave
//
// Gets the details for a key_octave.
//
// Responses:
// default: genericError
//
//	200: key_octaveDBResponse
func (controller *Controller) GetKey_octave(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetKey_octave", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_octave.GetDB()

	// Get key_octaveDB in DB
	var key_octaveDB orm.Key_octaveDB
	if err := db.First(&key_octaveDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var key_octaveAPI orm.Key_octaveAPI
	key_octaveAPI.ID = key_octaveDB.ID
	key_octaveAPI.Key_octavePointersEncoding = key_octaveDB.Key_octavePointersEncoding
	key_octaveDB.CopyBasicFieldsToKey_octave_WOP(&key_octaveAPI.Key_octave_WOP)

	c.JSON(http.StatusOK, key_octaveAPI)
}

// UpdateKey_octave
//
// swagger:route PATCH /key_octaves/{ID} key_octaves updateKey_octave
//
// # Update a key_octave
//
// Responses:
// default: genericError
//
//	200: key_octaveDBResponse
func (controller *Controller) UpdateKey_octave(c *gin.Context) {

	mutexKey_octave.Lock()
	defer mutexKey_octave.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateKey_octave", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_octave.GetDB()

	// Validate input
	var input orm.Key_octaveAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var key_octaveDB orm.Key_octaveDB

	// fetch the key_octave
	query := db.First(&key_octaveDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	key_octaveDB.CopyBasicFieldsFromKey_octave_WOP(&input.Key_octave_WOP)
	key_octaveDB.Key_octavePointersEncoding = input.Key_octavePointersEncoding

	query = db.Model(&key_octaveDB).Updates(key_octaveDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	key_octaveNew := new(models.Key_octave)
	key_octaveDB.CopyBasicFieldsToKey_octave(key_octaveNew)

	// redeem pointers
	key_octaveDB.DecodePointers(backRepo, key_octaveNew)

	// get stage instance from DB instance, and call callback function
	key_octaveOld := backRepo.BackRepoKey_octave.Map_Key_octaveDBID_Key_octavePtr[key_octaveDB.ID]
	if key_octaveOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), key_octaveOld, key_octaveNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the key_octaveDB
	c.JSON(http.StatusOK, key_octaveDB)
}

// DeleteKey_octave
//
// swagger:route DELETE /key_octaves/{ID} key_octaves deleteKey_octave
//
// # Delete a key_octave
//
// default: genericError
//
//	200: key_octaveDBResponse
func (controller *Controller) DeleteKey_octave(c *gin.Context) {

	mutexKey_octave.Lock()
	defer mutexKey_octave.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteKey_octave", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoKey_octave.GetDB()

	// Get model if exist
	var key_octaveDB orm.Key_octaveDB
	if err := db.First(&key_octaveDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&key_octaveDB)

	// get an instance (not staged) from DB instance, and call callback function
	key_octaveDeleted := new(models.Key_octave)
	key_octaveDB.CopyBasicFieldsToKey_octave(key_octaveDeleted)

	// get stage instance from DB instance, and call callback function
	key_octaveStaged := backRepo.BackRepoKey_octave.Map_Key_octaveDBID_Key_octavePtr[key_octaveDB.ID]
	if key_octaveStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), key_octaveStaged, key_octaveDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
