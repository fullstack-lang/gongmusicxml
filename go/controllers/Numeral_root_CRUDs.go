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
var __Numeral_root__dummysDeclaration__ models.Numeral_root
var __Numeral_root_time__dummyDeclaration time.Duration

var mutexNumeral_root sync.Mutex

// An Numeral_rootID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNumeral_root updateNumeral_root deleteNumeral_root
type Numeral_rootID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Numeral_rootInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNumeral_root updateNumeral_root
type Numeral_rootInput struct {
	// The Numeral_root to submit or modify
	// in: body
	Numeral_root *orm.Numeral_rootAPI
}

// GetNumeral_roots
//
// swagger:route GET /numeral_roots numeral_roots getNumeral_roots
//
// # Get all numeral_roots
//
// Responses:
// default: genericError
//
//	200: numeral_rootDBResponse
func (controller *Controller) GetNumeral_roots(c *gin.Context) {

	// source slice
	var numeral_rootDBs []orm.Numeral_rootDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNumeral_roots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_root.GetDB()

	query := db.Find(&numeral_rootDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	numeral_rootAPIs := make([]orm.Numeral_rootAPI, 0)

	// for each numeral_root, update fields from the database nullable fields
	for idx := range numeral_rootDBs {
		numeral_rootDB := &numeral_rootDBs[idx]
		_ = numeral_rootDB
		var numeral_rootAPI orm.Numeral_rootAPI

		// insertion point for updating fields
		numeral_rootAPI.ID = numeral_rootDB.ID
		numeral_rootDB.CopyBasicFieldsToNumeral_root_WOP(&numeral_rootAPI.Numeral_root_WOP)
		numeral_rootAPI.Numeral_rootPointersEncoding = numeral_rootDB.Numeral_rootPointersEncoding
		numeral_rootAPIs = append(numeral_rootAPIs, numeral_rootAPI)
	}

	c.JSON(http.StatusOK, numeral_rootAPIs)
}

// PostNumeral_root
//
// swagger:route POST /numeral_roots numeral_roots postNumeral_root
//
// Creates a numeral_root
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNumeral_root(c *gin.Context) {

	mutexNumeral_root.Lock()
	defer mutexNumeral_root.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNumeral_roots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_root.GetDB()

	// Validate input
	var input orm.Numeral_rootAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create numeral_root
	numeral_rootDB := orm.Numeral_rootDB{}
	numeral_rootDB.Numeral_rootPointersEncoding = input.Numeral_rootPointersEncoding
	numeral_rootDB.CopyBasicFieldsFromNumeral_root_WOP(&input.Numeral_root_WOP)

	query := db.Create(&numeral_rootDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNumeral_root.CheckoutPhaseOneInstance(&numeral_rootDB)
	numeral_root := backRepo.BackRepoNumeral_root.Map_Numeral_rootDBID_Numeral_rootPtr[numeral_rootDB.ID]

	if numeral_root != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), numeral_root)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, numeral_rootDB)
}

// GetNumeral_root
//
// swagger:route GET /numeral_roots/{ID} numeral_roots getNumeral_root
//
// Gets the details for a numeral_root.
//
// Responses:
// default: genericError
//
//	200: numeral_rootDBResponse
func (controller *Controller) GetNumeral_root(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNumeral_root", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_root.GetDB()

	// Get numeral_rootDB in DB
	var numeral_rootDB orm.Numeral_rootDB
	if err := db.First(&numeral_rootDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var numeral_rootAPI orm.Numeral_rootAPI
	numeral_rootAPI.ID = numeral_rootDB.ID
	numeral_rootAPI.Numeral_rootPointersEncoding = numeral_rootDB.Numeral_rootPointersEncoding
	numeral_rootDB.CopyBasicFieldsToNumeral_root_WOP(&numeral_rootAPI.Numeral_root_WOP)

	c.JSON(http.StatusOK, numeral_rootAPI)
}

// UpdateNumeral_root
//
// swagger:route PATCH /numeral_roots/{ID} numeral_roots updateNumeral_root
//
// # Update a numeral_root
//
// Responses:
// default: genericError
//
//	200: numeral_rootDBResponse
func (controller *Controller) UpdateNumeral_root(c *gin.Context) {

	mutexNumeral_root.Lock()
	defer mutexNumeral_root.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNumeral_root", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_root.GetDB()

	// Validate input
	var input orm.Numeral_rootAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var numeral_rootDB orm.Numeral_rootDB

	// fetch the numeral_root
	query := db.First(&numeral_rootDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	numeral_rootDB.CopyBasicFieldsFromNumeral_root_WOP(&input.Numeral_root_WOP)
	numeral_rootDB.Numeral_rootPointersEncoding = input.Numeral_rootPointersEncoding

	query = db.Model(&numeral_rootDB).Updates(numeral_rootDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	numeral_rootNew := new(models.Numeral_root)
	numeral_rootDB.CopyBasicFieldsToNumeral_root(numeral_rootNew)

	// redeem pointers
	numeral_rootDB.DecodePointers(backRepo, numeral_rootNew)

	// get stage instance from DB instance, and call callback function
	numeral_rootOld := backRepo.BackRepoNumeral_root.Map_Numeral_rootDBID_Numeral_rootPtr[numeral_rootDB.ID]
	if numeral_rootOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), numeral_rootOld, numeral_rootNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the numeral_rootDB
	c.JSON(http.StatusOK, numeral_rootDB)
}

// DeleteNumeral_root
//
// swagger:route DELETE /numeral_roots/{ID} numeral_roots deleteNumeral_root
//
// # Delete a numeral_root
//
// default: genericError
//
//	200: numeral_rootDBResponse
func (controller *Controller) DeleteNumeral_root(c *gin.Context) {

	mutexNumeral_root.Lock()
	defer mutexNumeral_root.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNumeral_root", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral_root.GetDB()

	// Get model if exist
	var numeral_rootDB orm.Numeral_rootDB
	if err := db.First(&numeral_rootDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&numeral_rootDB)

	// get an instance (not staged) from DB instance, and call callback function
	numeral_rootDeleted := new(models.Numeral_root)
	numeral_rootDB.CopyBasicFieldsToNumeral_root(numeral_rootDeleted)

	// get stage instance from DB instance, and call callback function
	numeral_rootStaged := backRepo.BackRepoNumeral_root.Map_Numeral_rootDBID_Numeral_rootPtr[numeral_rootDB.ID]
	if numeral_rootStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), numeral_rootStaged, numeral_rootDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
