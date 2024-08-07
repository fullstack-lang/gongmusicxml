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
var __Numeral__dummysDeclaration__ models.Numeral
var __Numeral_time__dummyDeclaration time.Duration

var mutexNumeral sync.Mutex

// An NumeralID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getNumeral updateNumeral deleteNumeral
type NumeralID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// NumeralInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postNumeral updateNumeral
type NumeralInput struct {
	// The Numeral to submit or modify
	// in: body
	Numeral *orm.NumeralAPI
}

// GetNumerals
//
// swagger:route GET /numerals numerals getNumerals
//
// # Get all numerals
//
// Responses:
// default: genericError
//
//	200: numeralDBResponse
func (controller *Controller) GetNumerals(c *gin.Context) {

	// source slice
	var numeralDBs []orm.NumeralDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNumerals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral.GetDB()

	query := db.Find(&numeralDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	numeralAPIs := make([]orm.NumeralAPI, 0)

	// for each numeral, update fields from the database nullable fields
	for idx := range numeralDBs {
		numeralDB := &numeralDBs[idx]
		_ = numeralDB
		var numeralAPI orm.NumeralAPI

		// insertion point for updating fields
		numeralAPI.ID = numeralDB.ID
		numeralDB.CopyBasicFieldsToNumeral_WOP(&numeralAPI.Numeral_WOP)
		numeralAPI.NumeralPointersEncoding = numeralDB.NumeralPointersEncoding
		numeralAPIs = append(numeralAPIs, numeralAPI)
	}

	c.JSON(http.StatusOK, numeralAPIs)
}

// PostNumeral
//
// swagger:route POST /numerals numerals postNumeral
//
// Creates a numeral
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostNumeral(c *gin.Context) {

	mutexNumeral.Lock()
	defer mutexNumeral.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostNumerals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral.GetDB()

	// Validate input
	var input orm.NumeralAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create numeral
	numeralDB := orm.NumeralDB{}
	numeralDB.NumeralPointersEncoding = input.NumeralPointersEncoding
	numeralDB.CopyBasicFieldsFromNumeral_WOP(&input.Numeral_WOP)

	query := db.Create(&numeralDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoNumeral.CheckoutPhaseOneInstance(&numeralDB)
	numeral := backRepo.BackRepoNumeral.Map_NumeralDBID_NumeralPtr[numeralDB.ID]

	if numeral != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), numeral)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, numeralDB)
}

// GetNumeral
//
// swagger:route GET /numerals/{ID} numerals getNumeral
//
// Gets the details for a numeral.
//
// Responses:
// default: genericError
//
//	200: numeralDBResponse
func (controller *Controller) GetNumeral(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetNumeral", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral.GetDB()

	// Get numeralDB in DB
	var numeralDB orm.NumeralDB
	if err := db.First(&numeralDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var numeralAPI orm.NumeralAPI
	numeralAPI.ID = numeralDB.ID
	numeralAPI.NumeralPointersEncoding = numeralDB.NumeralPointersEncoding
	numeralDB.CopyBasicFieldsToNumeral_WOP(&numeralAPI.Numeral_WOP)

	c.JSON(http.StatusOK, numeralAPI)
}

// UpdateNumeral
//
// swagger:route PATCH /numerals/{ID} numerals updateNumeral
//
// # Update a numeral
//
// Responses:
// default: genericError
//
//	200: numeralDBResponse
func (controller *Controller) UpdateNumeral(c *gin.Context) {

	mutexNumeral.Lock()
	defer mutexNumeral.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateNumeral", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral.GetDB()

	// Validate input
	var input orm.NumeralAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var numeralDB orm.NumeralDB

	// fetch the numeral
	query := db.First(&numeralDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	numeralDB.CopyBasicFieldsFromNumeral_WOP(&input.Numeral_WOP)
	numeralDB.NumeralPointersEncoding = input.NumeralPointersEncoding

	query = db.Model(&numeralDB).Updates(numeralDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	numeralNew := new(models.Numeral)
	numeralDB.CopyBasicFieldsToNumeral(numeralNew)

	// redeem pointers
	numeralDB.DecodePointers(backRepo, numeralNew)

	// get stage instance from DB instance, and call callback function
	numeralOld := backRepo.BackRepoNumeral.Map_NumeralDBID_NumeralPtr[numeralDB.ID]
	if numeralOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), numeralOld, numeralNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the numeralDB
	c.JSON(http.StatusOK, numeralDB)
}

// DeleteNumeral
//
// swagger:route DELETE /numerals/{ID} numerals deleteNumeral
//
// # Delete a numeral
//
// default: genericError
//
//	200: numeralDBResponse
func (controller *Controller) DeleteNumeral(c *gin.Context) {

	mutexNumeral.Lock()
	defer mutexNumeral.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteNumeral", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoNumeral.GetDB()

	// Get model if exist
	var numeralDB orm.NumeralDB
	if err := db.First(&numeralDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&numeralDB)

	// get an instance (not staged) from DB instance, and call callback function
	numeralDeleted := new(models.Numeral)
	numeralDB.CopyBasicFieldsToNumeral(numeralDeleted)

	// get stage instance from DB instance, and call callback function
	numeralStaged := backRepo.BackRepoNumeral.Map_NumeralDBID_NumeralPtr[numeralDB.ID]
	if numeralStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), numeralStaged, numeralDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
