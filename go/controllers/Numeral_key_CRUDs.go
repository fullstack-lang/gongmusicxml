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
var __Numeral_key__dummysDeclaration__ models.Numeral_key
var __Numeral_key_time__dummyDeclaration time.Duration

var mutexNumeral_key sync.Mutex

// An Numeral_keyID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNumeral_key updateNumeral_key deleteNumeral_key
type Numeral_keyID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Numeral_keyInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNumeral_key updateNumeral_key
type Numeral_keyInput struct {
	// The Numeral_key to submit or modify
	// in: body
	Numeral_key *orm.Numeral_keyAPI
}

// GetNumeral_keys
//
// swagger:route GET /numeral_keys numeral_keys getNumeral_keys
//
// # Get all numeral_keys
//
// Responses:
// default: genericError
//
//	200: numeral_keyDBResponse
func (controller *Controller) GetNumeral_keys(c *gin.Context) {

	// source slice
	var numeral_keyDBs []orm.Numeral_keyDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNumeral_keys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_key.GetDB()

	query := db.Find(&numeral_keyDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	numeral_keyAPIs := make([]orm.Numeral_keyAPI, 0)

	// for each numeral_key, update fields from the database nullable fields
	for idx := range numeral_keyDBs {
		numeral_keyDB := &numeral_keyDBs[idx]
		_ = numeral_keyDB
		var numeral_keyAPI orm.Numeral_keyAPI

		// insertion point for updating fields
		numeral_keyAPI.ID = numeral_keyDB.ID
		numeral_keyDB.CopyBasicFieldsToNumeral_key_WOP(&numeral_keyAPI.Numeral_key_WOP)
		numeral_keyAPI.Numeral_keyPointersEncoding = numeral_keyDB.Numeral_keyPointersEncoding
		numeral_keyAPIs = append(numeral_keyAPIs, numeral_keyAPI)
	}

	c.JSON(http.StatusOK, numeral_keyAPIs)
}

// PostNumeral_key
//
// swagger:route POST /numeral_keys numeral_keys postNumeral_key
//
// Creates a numeral_key
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNumeral_key(c *gin.Context) {

	mutexNumeral_key.Lock()
	defer mutexNumeral_key.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNumeral_keys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_key.GetDB()

	// Validate input
	var input orm.Numeral_keyAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create numeral_key
	numeral_keyDB := orm.Numeral_keyDB{}
	numeral_keyDB.Numeral_keyPointersEncoding = input.Numeral_keyPointersEncoding
	numeral_keyDB.CopyBasicFieldsFromNumeral_key_WOP(&input.Numeral_key_WOP)

	query := db.Create(&numeral_keyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNumeral_key.CheckoutPhaseOneInstance(&numeral_keyDB)
	numeral_key := backRepo.BackRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID]

	if numeral_key != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), numeral_key)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, numeral_keyDB)
}

// GetNumeral_key
//
// swagger:route GET /numeral_keys/{ID} numeral_keys getNumeral_key
//
// Gets the details for a numeral_key.
//
// Responses:
// default: genericError
//
//	200: numeral_keyDBResponse
func (controller *Controller) GetNumeral_key(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNumeral_key", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_key.GetDB()

	// Get numeral_keyDB in DB
	var numeral_keyDB orm.Numeral_keyDB
	if err := db.First(&numeral_keyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var numeral_keyAPI orm.Numeral_keyAPI
	numeral_keyAPI.ID = numeral_keyDB.ID
	numeral_keyAPI.Numeral_keyPointersEncoding = numeral_keyDB.Numeral_keyPointersEncoding
	numeral_keyDB.CopyBasicFieldsToNumeral_key_WOP(&numeral_keyAPI.Numeral_key_WOP)

	c.JSON(http.StatusOK, numeral_keyAPI)
}

// UpdateNumeral_key
//
// swagger:route PATCH /numeral_keys/{ID} numeral_keys updateNumeral_key
//
// # Update a numeral_key
//
// Responses:
// default: genericError
//
//	200: numeral_keyDBResponse
func (controller *Controller) UpdateNumeral_key(c *gin.Context) {

	mutexNumeral_key.Lock()
	defer mutexNumeral_key.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNumeral_key", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_key.GetDB()

	// Validate input
	var input orm.Numeral_keyAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var numeral_keyDB orm.Numeral_keyDB

	// fetch the numeral_key
	query := db.First(&numeral_keyDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	numeral_keyDB.CopyBasicFieldsFromNumeral_key_WOP(&input.Numeral_key_WOP)
	numeral_keyDB.Numeral_keyPointersEncoding = input.Numeral_keyPointersEncoding

	query = db.Model(&numeral_keyDB).Updates(numeral_keyDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	numeral_keyNew := new(models.Numeral_key)
	numeral_keyDB.CopyBasicFieldsToNumeral_key(numeral_keyNew)

	// redeem pointers
	numeral_keyDB.DecodePointers(backRepo, numeral_keyNew)

	// get stage instance from DB instance, and call callback function
	numeral_keyOld := backRepo.BackRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID]
	if numeral_keyOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), numeral_keyOld, numeral_keyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the numeral_keyDB
	c.JSON(http.StatusOK, numeral_keyDB)
}

// DeleteNumeral_key
//
// swagger:route DELETE /numeral_keys/{ID} numeral_keys deleteNumeral_key
//
// # Delete a numeral_key
//
// default: genericError
//
//	200: numeral_keyDBResponse
func (controller *Controller) DeleteNumeral_key(c *gin.Context) {

	mutexNumeral_key.Lock()
	defer mutexNumeral_key.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNumeral_key", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_key.GetDB()

	// Get model if exist
	var numeral_keyDB orm.Numeral_keyDB
	if err := db.First(&numeral_keyDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&numeral_keyDB)

	// get an instance (not staged) from DB instance, and call callback function
	numeral_keyDeleted := new(models.Numeral_key)
	numeral_keyDB.CopyBasicFieldsToNumeral_key(numeral_keyDeleted)

	// get stage instance from DB instance, and call callback function
	numeral_keyStaged := backRepo.BackRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID]
	if numeral_keyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), numeral_keyStaged, numeral_keyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
