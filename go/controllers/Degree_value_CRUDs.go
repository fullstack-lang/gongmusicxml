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
var __Degree_value__dummysDeclaration__ models.Degree_value
var __Degree_value_time__dummyDeclaration time.Duration

var mutexDegree_value sync.Mutex

// An Degree_valueID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDegree_value updateDegree_value deleteDegree_value
type Degree_valueID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Degree_valueInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDegree_value updateDegree_value
type Degree_valueInput struct {
	// The Degree_value to submit or modify
	// in: body
	Degree_value *orm.Degree_valueAPI
}

// GetDegree_values
//
// swagger:route GET /degree_values degree_values getDegree_values
//
// # Get all degree_values
//
// Responses:
// default: genericError
//
//	200: degree_valueDBResponse
func (controller *Controller) GetDegree_values(c *gin.Context) {

	// source slice
	var degree_valueDBs []orm.Degree_valueDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree_values", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_value.GetDB()

	query := db.Find(&degree_valueDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	degree_valueAPIs := make([]orm.Degree_valueAPI, 0)

	// for each degree_value, update fields from the database nullable fields
	for idx := range degree_valueDBs {
		degree_valueDB := &degree_valueDBs[idx]
		_ = degree_valueDB
		var degree_valueAPI orm.Degree_valueAPI

		// insertion point for updating fields
		degree_valueAPI.ID = degree_valueDB.ID
		degree_valueDB.CopyBasicFieldsToDegree_value_WOP(&degree_valueAPI.Degree_value_WOP)
		degree_valueAPI.Degree_valuePointersEncoding = degree_valueDB.Degree_valuePointersEncoding
		degree_valueAPIs = append(degree_valueAPIs, degree_valueAPI)
	}

	c.JSON(http.StatusOK, degree_valueAPIs)
}

// PostDegree_value
//
// swagger:route POST /degree_values degree_values postDegree_value
//
// Creates a degree_value
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDegree_value(c *gin.Context) {

	mutexDegree_value.Lock()
	defer mutexDegree_value.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDegree_values", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_value.GetDB()

	// Validate input
	var input orm.Degree_valueAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create degree_value
	degree_valueDB := orm.Degree_valueDB{}
	degree_valueDB.Degree_valuePointersEncoding = input.Degree_valuePointersEncoding
	degree_valueDB.CopyBasicFieldsFromDegree_value_WOP(&input.Degree_value_WOP)

	query := db.Create(&degree_valueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDegree_value.CheckoutPhaseOneInstance(&degree_valueDB)
	degree_value := backRepo.BackRepoDegree_value.Map_Degree_valueDBID_Degree_valuePtr[degree_valueDB.ID]

	if degree_value != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), degree_value)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, degree_valueDB)
}

// GetDegree_value
//
// swagger:route GET /degree_values/{ID} degree_values getDegree_value
//
// Gets the details for a degree_value.
//
// Responses:
// default: genericError
//
//	200: degree_valueDBResponse
func (controller *Controller) GetDegree_value(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree_value", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_value.GetDB()

	// Get degree_valueDB in DB
	var degree_valueDB orm.Degree_valueDB
	if err := db.First(&degree_valueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var degree_valueAPI orm.Degree_valueAPI
	degree_valueAPI.ID = degree_valueDB.ID
	degree_valueAPI.Degree_valuePointersEncoding = degree_valueDB.Degree_valuePointersEncoding
	degree_valueDB.CopyBasicFieldsToDegree_value_WOP(&degree_valueAPI.Degree_value_WOP)

	c.JSON(http.StatusOK, degree_valueAPI)
}

// UpdateDegree_value
//
// swagger:route PATCH /degree_values/{ID} degree_values updateDegree_value
//
// # Update a degree_value
//
// Responses:
// default: genericError
//
//	200: degree_valueDBResponse
func (controller *Controller) UpdateDegree_value(c *gin.Context) {

	mutexDegree_value.Lock()
	defer mutexDegree_value.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDegree_value", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_value.GetDB()

	// Validate input
	var input orm.Degree_valueAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var degree_valueDB orm.Degree_valueDB

	// fetch the degree_value
	query := db.First(&degree_valueDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	degree_valueDB.CopyBasicFieldsFromDegree_value_WOP(&input.Degree_value_WOP)
	degree_valueDB.Degree_valuePointersEncoding = input.Degree_valuePointersEncoding

	query = db.Model(&degree_valueDB).Updates(degree_valueDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	degree_valueNew := new(models.Degree_value)
	degree_valueDB.CopyBasicFieldsToDegree_value(degree_valueNew)

	// redeem pointers
	degree_valueDB.DecodePointers(backRepo, degree_valueNew)

	// get stage instance from DB instance, and call callback function
	degree_valueOld := backRepo.BackRepoDegree_value.Map_Degree_valueDBID_Degree_valuePtr[degree_valueDB.ID]
	if degree_valueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), degree_valueOld, degree_valueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the degree_valueDB
	c.JSON(http.StatusOK, degree_valueDB)
}

// DeleteDegree_value
//
// swagger:route DELETE /degree_values/{ID} degree_values deleteDegree_value
//
// # Delete a degree_value
//
// default: genericError
//
//	200: degree_valueDBResponse
func (controller *Controller) DeleteDegree_value(c *gin.Context) {

	mutexDegree_value.Lock()
	defer mutexDegree_value.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDegree_value", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_value.GetDB()

	// Get model if exist
	var degree_valueDB orm.Degree_valueDB
	if err := db.First(&degree_valueDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&degree_valueDB)

	// get an instance (not staged) from DB instance, and call callback function
	degree_valueDeleted := new(models.Degree_value)
	degree_valueDB.CopyBasicFieldsToDegree_value(degree_valueDeleted)

	// get stage instance from DB instance, and call callback function
	degree_valueStaged := backRepo.BackRepoDegree_value.Map_Degree_valueDBID_Degree_valuePtr[degree_valueDB.ID]
	if degree_valueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), degree_valueStaged, degree_valueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
