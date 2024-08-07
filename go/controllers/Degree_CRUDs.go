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
var __Degree__dummysDeclaration__ models.Degree
var __Degree_time__dummyDeclaration time.Duration

var mutexDegree sync.Mutex

// An DegreeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDegree updateDegree deleteDegree
type DegreeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DegreeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDegree updateDegree
type DegreeInput struct {
	// The Degree to submit or modify
	// in: body
	Degree *orm.DegreeAPI
}

// GetDegrees
//
// swagger:route GET /degrees degrees getDegrees
//
// # Get all degrees
//
// Responses:
// default: genericError
//
//	200: degreeDBResponse
func (controller *Controller) GetDegrees(c *gin.Context) {

	// source slice
	var degreeDBs []orm.DegreeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegrees", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree.GetDB()

	query := db.Find(&degreeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	degreeAPIs := make([]orm.DegreeAPI, 0)

	// for each degree, update fields from the database nullable fields
	for idx := range degreeDBs {
		degreeDB := &degreeDBs[idx]
		_ = degreeDB
		var degreeAPI orm.DegreeAPI

		// insertion point for updating fields
		degreeAPI.ID = degreeDB.ID
		degreeDB.CopyBasicFieldsToDegree_WOP(&degreeAPI.Degree_WOP)
		degreeAPI.DegreePointersEncoding = degreeDB.DegreePointersEncoding
		degreeAPIs = append(degreeAPIs, degreeAPI)
	}

	c.JSON(http.StatusOK, degreeAPIs)
}

// PostDegree
//
// swagger:route POST /degrees degrees postDegree
//
// Creates a degree
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDegree(c *gin.Context) {

	mutexDegree.Lock()
	defer mutexDegree.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDegrees", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree.GetDB()

	// Validate input
	var input orm.DegreeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create degree
	degreeDB := orm.DegreeDB{}
	degreeDB.DegreePointersEncoding = input.DegreePointersEncoding
	degreeDB.CopyBasicFieldsFromDegree_WOP(&input.Degree_WOP)

	query := db.Create(&degreeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDegree.CheckoutPhaseOneInstance(&degreeDB)
	degree := backRepo.BackRepoDegree.Map_DegreeDBID_DegreePtr[degreeDB.ID]

	if degree != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), degree)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, degreeDB)
}

// GetDegree
//
// swagger:route GET /degrees/{ID} degrees getDegree
//
// Gets the details for a degree.
//
// Responses:
// default: genericError
//
//	200: degreeDBResponse
func (controller *Controller) GetDegree(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDegree", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree.GetDB()

	// Get degreeDB in DB
	var degreeDB orm.DegreeDB
	if err := db.First(&degreeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var degreeAPI orm.DegreeAPI
	degreeAPI.ID = degreeDB.ID
	degreeAPI.DegreePointersEncoding = degreeDB.DegreePointersEncoding
	degreeDB.CopyBasicFieldsToDegree_WOP(&degreeAPI.Degree_WOP)

	c.JSON(http.StatusOK, degreeAPI)
}

// UpdateDegree
//
// swagger:route PATCH /degrees/{ID} degrees updateDegree
//
// # Update a degree
//
// Responses:
// default: genericError
//
//	200: degreeDBResponse
func (controller *Controller) UpdateDegree(c *gin.Context) {

	mutexDegree.Lock()
	defer mutexDegree.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDegree", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree.GetDB()

	// Validate input
	var input orm.DegreeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var degreeDB orm.DegreeDB

	// fetch the degree
	query := db.First(&degreeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	degreeDB.CopyBasicFieldsFromDegree_WOP(&input.Degree_WOP)
	degreeDB.DegreePointersEncoding = input.DegreePointersEncoding

	query = db.Model(&degreeDB).Updates(degreeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	degreeNew := new(models.Degree)
	degreeDB.CopyBasicFieldsToDegree(degreeNew)

	// redeem pointers
	degreeDB.DecodePointers(backRepo, degreeNew)

	// get stage instance from DB instance, and call callback function
	degreeOld := backRepo.BackRepoDegree.Map_DegreeDBID_DegreePtr[degreeDB.ID]
	if degreeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), degreeOld, degreeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the degreeDB
	c.JSON(http.StatusOK, degreeDB)
}

// DeleteDegree
//
// swagger:route DELETE /degrees/{ID} degrees deleteDegree
//
// # Delete a degree
//
// default: genericError
//
//	200: degreeDBResponse
func (controller *Controller) DeleteDegree(c *gin.Context) {

	mutexDegree.Lock()
	defer mutexDegree.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDegree", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDegree.GetDB()

	// Get model if exist
	var degreeDB orm.DegreeDB
	if err := db.First(&degreeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&degreeDB)

	// get an instance (not staged) from DB instance, and call callback function
	degreeDeleted := new(models.Degree)
	degreeDB.CopyBasicFieldsToDegree(degreeDeleted)

	// get stage instance from DB instance, and call callback function
	degreeStaged := backRepo.BackRepoDegree.Map_DegreeDBID_DegreePtr[degreeDB.ID]
	if degreeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), degreeStaged, degreeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
