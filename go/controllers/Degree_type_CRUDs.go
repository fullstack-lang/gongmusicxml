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
var __Degree_type__dummysDeclaration__ models.Degree_type
var __Degree_type_time__dummyDeclaration time.Duration

var mutexDegree_type sync.Mutex

// An Degree_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDegree_type updateDegree_type deleteDegree_type
type Degree_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Degree_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDegree_type updateDegree_type
type Degree_typeInput struct {
	// The Degree_type to submit or modify
	// in: body
	Degree_type *orm.Degree_typeAPI
}

// GetDegree_types
//
// swagger:route GET /degree_types degree_types getDegree_types
//
// # Get all degree_types
//
// Responses:
// default: genericError
//
//	200: degree_typeDBResponse
func (controller *Controller) GetDegree_types(c *gin.Context) {

	// source slice
	var degree_typeDBs []orm.Degree_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_type.GetDB()

	query := db.Find(&degree_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	degree_typeAPIs := make([]orm.Degree_typeAPI, 0)

	// for each degree_type, update fields from the database nullable fields
	for idx := range degree_typeDBs {
		degree_typeDB := &degree_typeDBs[idx]
		_ = degree_typeDB
		var degree_typeAPI orm.Degree_typeAPI

		// insertion point for updating fields
		degree_typeAPI.ID = degree_typeDB.ID
		degree_typeDB.CopyBasicFieldsToDegree_type_WOP(&degree_typeAPI.Degree_type_WOP)
		degree_typeAPI.Degree_typePointersEncoding = degree_typeDB.Degree_typePointersEncoding
		degree_typeAPIs = append(degree_typeAPIs, degree_typeAPI)
	}

	c.JSON(http.StatusOK, degree_typeAPIs)
}

// PostDegree_type
//
// swagger:route POST /degree_types degree_types postDegree_type
//
// Creates a degree_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDegree_type(c *gin.Context) {

	mutexDegree_type.Lock()
	defer mutexDegree_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDegree_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_type.GetDB()

	// Validate input
	var input orm.Degree_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create degree_type
	degree_typeDB := orm.Degree_typeDB{}
	degree_typeDB.Degree_typePointersEncoding = input.Degree_typePointersEncoding
	degree_typeDB.CopyBasicFieldsFromDegree_type_WOP(&input.Degree_type_WOP)

	query := db.Create(&degree_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDegree_type.CheckoutPhaseOneInstance(&degree_typeDB)
	degree_type := backRepo.BackRepoDegree_type.Map_Degree_typeDBID_Degree_typePtr[degree_typeDB.ID]

	if degree_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), degree_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, degree_typeDB)
}

// GetDegree_type
//
// swagger:route GET /degree_types/{ID} degree_types getDegree_type
//
// Gets the details for a degree_type.
//
// Responses:
// default: genericError
//
//	200: degree_typeDBResponse
func (controller *Controller) GetDegree_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_type.GetDB()

	// Get degree_typeDB in DB
	var degree_typeDB orm.Degree_typeDB
	if err := db.First(&degree_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var degree_typeAPI orm.Degree_typeAPI
	degree_typeAPI.ID = degree_typeDB.ID
	degree_typeAPI.Degree_typePointersEncoding = degree_typeDB.Degree_typePointersEncoding
	degree_typeDB.CopyBasicFieldsToDegree_type_WOP(&degree_typeAPI.Degree_type_WOP)

	c.JSON(http.StatusOK, degree_typeAPI)
}

// UpdateDegree_type
//
// swagger:route PATCH /degree_types/{ID} degree_types updateDegree_type
//
// # Update a degree_type
//
// Responses:
// default: genericError
//
//	200: degree_typeDBResponse
func (controller *Controller) UpdateDegree_type(c *gin.Context) {

	mutexDegree_type.Lock()
	defer mutexDegree_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDegree_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_type.GetDB()

	// Validate input
	var input orm.Degree_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var degree_typeDB orm.Degree_typeDB

	// fetch the degree_type
	query := db.First(&degree_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	degree_typeDB.CopyBasicFieldsFromDegree_type_WOP(&input.Degree_type_WOP)
	degree_typeDB.Degree_typePointersEncoding = input.Degree_typePointersEncoding

	query = db.Model(&degree_typeDB).Updates(degree_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	degree_typeNew := new(models.Degree_type)
	degree_typeDB.CopyBasicFieldsToDegree_type(degree_typeNew)

	// redeem pointers
	degree_typeDB.DecodePointers(backRepo, degree_typeNew)

	// get stage instance from DB instance, and call callback function
	degree_typeOld := backRepo.BackRepoDegree_type.Map_Degree_typeDBID_Degree_typePtr[degree_typeDB.ID]
	if degree_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), degree_typeOld, degree_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the degree_typeDB
	c.JSON(http.StatusOK, degree_typeDB)
}

// DeleteDegree_type
//
// swagger:route DELETE /degree_types/{ID} degree_types deleteDegree_type
//
// # Delete a degree_type
//
// default: genericError
//
//	200: degree_typeDBResponse
func (controller *Controller) DeleteDegree_type(c *gin.Context) {

	mutexDegree_type.Lock()
	defer mutexDegree_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDegree_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree_type.GetDB()

	// Get model if exist
	var degree_typeDB orm.Degree_typeDB
	if err := db.First(&degree_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&degree_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	degree_typeDeleted := new(models.Degree_type)
	degree_typeDB.CopyBasicFieldsToDegree_type(degree_typeDeleted)

	// get stage instance from DB instance, and call callback function
	degree_typeStaged := backRepo.BackRepoDegree_type.Map_Degree_typeDBID_Degree_typePtr[degree_typeDB.ID]
	if degree_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), degree_typeStaged, degree_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
