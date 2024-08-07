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
var __Scaling__dummysDeclaration__ models.Scaling
var __Scaling_time__dummyDeclaration time.Duration

var mutexScaling sync.Mutex

// An ScalingID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScaling updateScaling deleteScaling
type ScalingID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ScalingInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScaling updateScaling
type ScalingInput struct {
	// The Scaling to submit or modify
	// in: body
	Scaling *orm.ScalingAPI
}

// GetScalings
//
// swagger:route GET /scalings scalings getScalings
//
// # Get all scalings
//
// Responses:
// default: genericError
//
//	200: scalingDBResponse
func (controller *Controller) GetScalings(c *gin.Context) {

	// source slice
	var scalingDBs []orm.ScalingDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScalings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScaling.GetDB()

	query := db.Find(&scalingDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	scalingAPIs := make([]orm.ScalingAPI, 0)

	// for each scaling, update fields from the database nullable fields
	for idx := range scalingDBs {
		scalingDB := &scalingDBs[idx]
		_ = scalingDB
		var scalingAPI orm.ScalingAPI

		// insertion point for updating fields
		scalingAPI.ID = scalingDB.ID
		scalingDB.CopyBasicFieldsToScaling_WOP(&scalingAPI.Scaling_WOP)
		scalingAPI.ScalingPointersEncoding = scalingDB.ScalingPointersEncoding
		scalingAPIs = append(scalingAPIs, scalingAPI)
	}

	c.JSON(http.StatusOK, scalingAPIs)
}

// PostScaling
//
// swagger:route POST /scalings scalings postScaling
//
// Creates a scaling
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScaling(c *gin.Context) {

	mutexScaling.Lock()
	defer mutexScaling.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScalings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScaling.GetDB()

	// Validate input
	var input orm.ScalingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create scaling
	scalingDB := orm.ScalingDB{}
	scalingDB.ScalingPointersEncoding = input.ScalingPointersEncoding
	scalingDB.CopyBasicFieldsFromScaling_WOP(&input.Scaling_WOP)

	query := db.Create(&scalingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScaling.CheckoutPhaseOneInstance(&scalingDB)
	scaling := backRepo.BackRepoScaling.Map_ScalingDBID_ScalingPtr[scalingDB.ID]

	if scaling != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), scaling)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, scalingDB)
}

// GetScaling
//
// swagger:route GET /scalings/{ID} scalings getScaling
//
// Gets the details for a scaling.
//
// Responses:
// default: genericError
//
//	200: scalingDBResponse
func (controller *Controller) GetScaling(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScaling", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScaling.GetDB()

	// Get scalingDB in DB
	var scalingDB orm.ScalingDB
	if err := db.First(&scalingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var scalingAPI orm.ScalingAPI
	scalingAPI.ID = scalingDB.ID
	scalingAPI.ScalingPointersEncoding = scalingDB.ScalingPointersEncoding
	scalingDB.CopyBasicFieldsToScaling_WOP(&scalingAPI.Scaling_WOP)

	c.JSON(http.StatusOK, scalingAPI)
}

// UpdateScaling
//
// swagger:route PATCH /scalings/{ID} scalings updateScaling
//
// # Update a scaling
//
// Responses:
// default: genericError
//
//	200: scalingDBResponse
func (controller *Controller) UpdateScaling(c *gin.Context) {

	mutexScaling.Lock()
	defer mutexScaling.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScaling", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScaling.GetDB()

	// Validate input
	var input orm.ScalingAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var scalingDB orm.ScalingDB

	// fetch the scaling
	query := db.First(&scalingDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	scalingDB.CopyBasicFieldsFromScaling_WOP(&input.Scaling_WOP)
	scalingDB.ScalingPointersEncoding = input.ScalingPointersEncoding

	query = db.Model(&scalingDB).Updates(scalingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	scalingNew := new(models.Scaling)
	scalingDB.CopyBasicFieldsToScaling(scalingNew)

	// redeem pointers
	scalingDB.DecodePointers(backRepo, scalingNew)

	// get stage instance from DB instance, and call callback function
	scalingOld := backRepo.BackRepoScaling.Map_ScalingDBID_ScalingPtr[scalingDB.ID]
	if scalingOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), scalingOld, scalingNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the scalingDB
	c.JSON(http.StatusOK, scalingDB)
}

// DeleteScaling
//
// swagger:route DELETE /scalings/{ID} scalings deleteScaling
//
// # Delete a scaling
//
// default: genericError
//
//	200: scalingDBResponse
func (controller *Controller) DeleteScaling(c *gin.Context) {

	mutexScaling.Lock()
	defer mutexScaling.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScaling", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScaling.GetDB()

	// Get model if exist
	var scalingDB orm.ScalingDB
	if err := db.First(&scalingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&scalingDB)

	// get an instance (not staged) from DB instance, and call callback function
	scalingDeleted := new(models.Scaling)
	scalingDB.CopyBasicFieldsToScaling(scalingDeleted)

	// get stage instance from DB instance, and call callback function
	scalingStaged := backRepo.BackRepoScaling.Map_ScalingDBID_ScalingPtr[scalingDB.ID]
	if scalingStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), scalingStaged, scalingDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
