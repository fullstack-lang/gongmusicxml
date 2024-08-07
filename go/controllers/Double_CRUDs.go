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
var __Double__dummysDeclaration__ models.Double
var __Double_time__dummyDeclaration time.Duration

var mutexDouble sync.Mutex

// An DoubleID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDouble updateDouble deleteDouble
type DoubleID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DoubleInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDouble updateDouble
type DoubleInput struct {
	// The Double to submit or modify
	// in: body
	Double *orm.DoubleAPI
}

// GetDoubles
//
// swagger:route GET /doubles doubles getDoubles
//
// # Get all doubles
//
// Responses:
// default: genericError
//
//	200: doubleDBResponse
func (controller *Controller) GetDoubles(c *gin.Context) {

	// source slice
	var doubleDBs []orm.DoubleDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDoubles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDouble.GetDB()

	query := db.Find(&doubleDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	doubleAPIs := make([]orm.DoubleAPI, 0)

	// for each double, update fields from the database nullable fields
	for idx := range doubleDBs {
		doubleDB := &doubleDBs[idx]
		_ = doubleDB
		var doubleAPI orm.DoubleAPI

		// insertion point for updating fields
		doubleAPI.ID = doubleDB.ID
		doubleDB.CopyBasicFieldsToDouble_WOP(&doubleAPI.Double_WOP)
		doubleAPI.DoublePointersEncoding = doubleDB.DoublePointersEncoding
		doubleAPIs = append(doubleAPIs, doubleAPI)
	}

	c.JSON(http.StatusOK, doubleAPIs)
}

// PostDouble
//
// swagger:route POST /doubles doubles postDouble
//
// Creates a double
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDouble(c *gin.Context) {

	mutexDouble.Lock()
	defer mutexDouble.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDoubles", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDouble.GetDB()

	// Validate input
	var input orm.DoubleAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create double
	doubleDB := orm.DoubleDB{}
	doubleDB.DoublePointersEncoding = input.DoublePointersEncoding
	doubleDB.CopyBasicFieldsFromDouble_WOP(&input.Double_WOP)

	query := db.Create(&doubleDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDouble.CheckoutPhaseOneInstance(&doubleDB)
	double := backRepo.BackRepoDouble.Map_DoubleDBID_DoublePtr[doubleDB.ID]

	if double != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), double)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, doubleDB)
}

// GetDouble
//
// swagger:route GET /doubles/{ID} doubles getDouble
//
// Gets the details for a double.
//
// Responses:
// default: genericError
//
//	200: doubleDBResponse
func (controller *Controller) GetDouble(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDouble", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDouble.GetDB()

	// Get doubleDB in DB
	var doubleDB orm.DoubleDB
	if err := db.First(&doubleDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var doubleAPI orm.DoubleAPI
	doubleAPI.ID = doubleDB.ID
	doubleAPI.DoublePointersEncoding = doubleDB.DoublePointersEncoding
	doubleDB.CopyBasicFieldsToDouble_WOP(&doubleAPI.Double_WOP)

	c.JSON(http.StatusOK, doubleAPI)
}

// UpdateDouble
//
// swagger:route PATCH /doubles/{ID} doubles updateDouble
//
// # Update a double
//
// Responses:
// default: genericError
//
//	200: doubleDBResponse
func (controller *Controller) UpdateDouble(c *gin.Context) {

	mutexDouble.Lock()
	defer mutexDouble.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDouble", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDouble.GetDB()

	// Validate input
	var input orm.DoubleAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var doubleDB orm.DoubleDB

	// fetch the double
	query := db.First(&doubleDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	doubleDB.CopyBasicFieldsFromDouble_WOP(&input.Double_WOP)
	doubleDB.DoublePointersEncoding = input.DoublePointersEncoding

	query = db.Model(&doubleDB).Updates(doubleDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	doubleNew := new(models.Double)
	doubleDB.CopyBasicFieldsToDouble(doubleNew)

	// redeem pointers
	doubleDB.DecodePointers(backRepo, doubleNew)

	// get stage instance from DB instance, and call callback function
	doubleOld := backRepo.BackRepoDouble.Map_DoubleDBID_DoublePtr[doubleDB.ID]
	if doubleOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), doubleOld, doubleNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the doubleDB
	c.JSON(http.StatusOK, doubleDB)
}

// DeleteDouble
//
// swagger:route DELETE /doubles/{ID} doubles deleteDouble
//
// # Delete a double
//
// default: genericError
//
//	200: doubleDBResponse
func (controller *Controller) DeleteDouble(c *gin.Context) {

	mutexDouble.Lock()
	defer mutexDouble.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDouble", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDouble.GetDB()

	// Get model if exist
	var doubleDB orm.DoubleDB
	if err := db.First(&doubleDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&doubleDB)

	// get an instance (not staged) from DB instance, and call callback function
	doubleDeleted := new(models.Double)
	doubleDB.CopyBasicFieldsToDouble(doubleDeleted)

	// get stage instance from DB instance, and call callback function
	doubleStaged := backRepo.BackRepoDouble.Map_DoubleDBID_DoublePtr[doubleDB.ID]
	if doubleStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), doubleStaged, doubleDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
