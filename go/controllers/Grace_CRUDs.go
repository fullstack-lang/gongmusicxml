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
var __Grace__dummysDeclaration__ models.Grace
var __Grace_time__dummyDeclaration time.Duration

var mutexGrace sync.Mutex

// An GraceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGrace updateGrace deleteGrace
type GraceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GraceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGrace updateGrace
type GraceInput struct {
	// The Grace to submit or modify
	// in: body
	Grace *orm.GraceAPI
}

// GetGraces
//
// swagger:route GET /graces graces getGraces
//
// # Get all graces
//
// Responses:
// default: genericError
//
//	200: graceDBResponse
func (controller *Controller) GetGraces(c *gin.Context) {

	// source slice
	var graceDBs []orm.GraceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGraces", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrace.GetDB()

	query := db.Find(&graceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	graceAPIs := make([]orm.GraceAPI, 0)

	// for each grace, update fields from the database nullable fields
	for idx := range graceDBs {
		graceDB := &graceDBs[idx]
		_ = graceDB
		var graceAPI orm.GraceAPI

		// insertion point for updating fields
		graceAPI.ID = graceDB.ID
		graceDB.CopyBasicFieldsToGrace_WOP(&graceAPI.Grace_WOP)
		graceAPI.GracePointersEncoding = graceDB.GracePointersEncoding
		graceAPIs = append(graceAPIs, graceAPI)
	}

	c.JSON(http.StatusOK, graceAPIs)
}

// PostGrace
//
// swagger:route POST /graces graces postGrace
//
// Creates a grace
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGrace(c *gin.Context) {

	mutexGrace.Lock()
	defer mutexGrace.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGraces", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrace.GetDB()

	// Validate input
	var input orm.GraceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create grace
	graceDB := orm.GraceDB{}
	graceDB.GracePointersEncoding = input.GracePointersEncoding
	graceDB.CopyBasicFieldsFromGrace_WOP(&input.Grace_WOP)

	query := db.Create(&graceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGrace.CheckoutPhaseOneInstance(&graceDB)
	grace := backRepo.BackRepoGrace.Map_GraceDBID_GracePtr[graceDB.ID]

	if grace != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), grace)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, graceDB)
}

// GetGrace
//
// swagger:route GET /graces/{ID} graces getGrace
//
// Gets the details for a grace.
//
// Responses:
// default: genericError
//
//	200: graceDBResponse
func (controller *Controller) GetGrace(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGrace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrace.GetDB()

	// Get graceDB in DB
	var graceDB orm.GraceDB
	if err := db.First(&graceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var graceAPI orm.GraceAPI
	graceAPI.ID = graceDB.ID
	graceAPI.GracePointersEncoding = graceDB.GracePointersEncoding
	graceDB.CopyBasicFieldsToGrace_WOP(&graceAPI.Grace_WOP)

	c.JSON(http.StatusOK, graceAPI)
}

// UpdateGrace
//
// swagger:route PATCH /graces/{ID} graces updateGrace
//
// # Update a grace
//
// Responses:
// default: genericError
//
//	200: graceDBResponse
func (controller *Controller) UpdateGrace(c *gin.Context) {

	mutexGrace.Lock()
	defer mutexGrace.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGrace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrace.GetDB()

	// Validate input
	var input orm.GraceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var graceDB orm.GraceDB

	// fetch the grace
	query := db.First(&graceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	graceDB.CopyBasicFieldsFromGrace_WOP(&input.Grace_WOP)
	graceDB.GracePointersEncoding = input.GracePointersEncoding

	query = db.Model(&graceDB).Updates(graceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	graceNew := new(models.Grace)
	graceDB.CopyBasicFieldsToGrace(graceNew)

	// redeem pointers
	graceDB.DecodePointers(backRepo, graceNew)

	// get stage instance from DB instance, and call callback function
	graceOld := backRepo.BackRepoGrace.Map_GraceDBID_GracePtr[graceDB.ID]
	if graceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), graceOld, graceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the graceDB
	c.JSON(http.StatusOK, graceDB)
}

// DeleteGrace
//
// swagger:route DELETE /graces/{ID} graces deleteGrace
//
// # Delete a grace
//
// default: genericError
//
//	200: graceDBResponse
func (controller *Controller) DeleteGrace(c *gin.Context) {

	mutexGrace.Lock()
	defer mutexGrace.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGrace", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrace.GetDB()

	// Get model if exist
	var graceDB orm.GraceDB
	if err := db.First(&graceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&graceDB)

	// get an instance (not staged) from DB instance, and call callback function
	graceDeleted := new(models.Grace)
	graceDB.CopyBasicFieldsToGrace(graceDeleted)

	// get stage instance from DB instance, and call callback function
	graceStaged := backRepo.BackRepoGrace.Map_GraceDBID_GracePtr[graceDB.ID]
	if graceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), graceStaged, graceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
