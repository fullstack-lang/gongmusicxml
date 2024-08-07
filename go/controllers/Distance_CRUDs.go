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
var __Distance__dummysDeclaration__ models.Distance
var __Distance_time__dummyDeclaration time.Duration

var mutexDistance sync.Mutex

// An DistanceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDistance updateDistance deleteDistance
type DistanceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DistanceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDistance updateDistance
type DistanceInput struct {
	// The Distance to submit or modify
	// in: body
	Distance *orm.DistanceAPI
}

// GetDistances
//
// swagger:route GET /distances distances getDistances
//
// # Get all distances
//
// Responses:
// default: genericError
//
//	200: distanceDBResponse
func (controller *Controller) GetDistances(c *gin.Context) {

	// source slice
	var distanceDBs []orm.DistanceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDistances", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDistance.GetDB()

	query := db.Find(&distanceDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	distanceAPIs := make([]orm.DistanceAPI, 0)

	// for each distance, update fields from the database nullable fields
	for idx := range distanceDBs {
		distanceDB := &distanceDBs[idx]
		_ = distanceDB
		var distanceAPI orm.DistanceAPI

		// insertion point for updating fields
		distanceAPI.ID = distanceDB.ID
		distanceDB.CopyBasicFieldsToDistance_WOP(&distanceAPI.Distance_WOP)
		distanceAPI.DistancePointersEncoding = distanceDB.DistancePointersEncoding
		distanceAPIs = append(distanceAPIs, distanceAPI)
	}

	c.JSON(http.StatusOK, distanceAPIs)
}

// PostDistance
//
// swagger:route POST /distances distances postDistance
//
// Creates a distance
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDistance(c *gin.Context) {

	mutexDistance.Lock()
	defer mutexDistance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDistances", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDistance.GetDB()

	// Validate input
	var input orm.DistanceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create distance
	distanceDB := orm.DistanceDB{}
	distanceDB.DistancePointersEncoding = input.DistancePointersEncoding
	distanceDB.CopyBasicFieldsFromDistance_WOP(&input.Distance_WOP)

	query := db.Create(&distanceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDistance.CheckoutPhaseOneInstance(&distanceDB)
	distance := backRepo.BackRepoDistance.Map_DistanceDBID_DistancePtr[distanceDB.ID]

	if distance != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), distance)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, distanceDB)
}

// GetDistance
//
// swagger:route GET /distances/{ID} distances getDistance
//
// Gets the details for a distance.
//
// Responses:
// default: genericError
//
//	200: distanceDBResponse
func (controller *Controller) GetDistance(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDistance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDistance.GetDB()

	// Get distanceDB in DB
	var distanceDB orm.DistanceDB
	if err := db.First(&distanceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var distanceAPI orm.DistanceAPI
	distanceAPI.ID = distanceDB.ID
	distanceAPI.DistancePointersEncoding = distanceDB.DistancePointersEncoding
	distanceDB.CopyBasicFieldsToDistance_WOP(&distanceAPI.Distance_WOP)

	c.JSON(http.StatusOK, distanceAPI)
}

// UpdateDistance
//
// swagger:route PATCH /distances/{ID} distances updateDistance
//
// # Update a distance
//
// Responses:
// default: genericError
//
//	200: distanceDBResponse
func (controller *Controller) UpdateDistance(c *gin.Context) {

	mutexDistance.Lock()
	defer mutexDistance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateDistance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDistance.GetDB()

	// Validate input
	var input orm.DistanceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var distanceDB orm.DistanceDB

	// fetch the distance
	query := db.First(&distanceDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	distanceDB.CopyBasicFieldsFromDistance_WOP(&input.Distance_WOP)
	distanceDB.DistancePointersEncoding = input.DistancePointersEncoding

	query = db.Model(&distanceDB).Updates(distanceDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	distanceNew := new(models.Distance)
	distanceDB.CopyBasicFieldsToDistance(distanceNew)

	// redeem pointers
	distanceDB.DecodePointers(backRepo, distanceNew)

	// get stage instance from DB instance, and call callback function
	distanceOld := backRepo.BackRepoDistance.Map_DistanceDBID_DistancePtr[distanceDB.ID]
	if distanceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), distanceOld, distanceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the distanceDB
	c.JSON(http.StatusOK, distanceDB)
}

// DeleteDistance
//
// swagger:route DELETE /distances/{ID} distances deleteDistance
//
// # Delete a distance
//
// default: genericError
//
//	200: distanceDBResponse
func (controller *Controller) DeleteDistance(c *gin.Context) {

	mutexDistance.Lock()
	defer mutexDistance.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDistance", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoDistance.GetDB()

	// Get model if exist
	var distanceDB orm.DistanceDB
	if err := db.First(&distanceDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&distanceDB)

	// get an instance (not staged) from DB instance, and call callback function
	distanceDeleted := new(models.Distance)
	distanceDB.CopyBasicFieldsToDistance(distanceDeleted)

	// get stage instance from DB instance, and call callback function
	distanceStaged := backRepo.BackRepoDistance.Map_DistanceDBID_DistancePtr[distanceDB.ID]
	if distanceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), distanceStaged, distanceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
