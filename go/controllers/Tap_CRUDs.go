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
var __Tap__dummysDeclaration__ models.Tap
var __Tap_time__dummyDeclaration time.Duration

var mutexTap sync.Mutex

// An TapID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTap updateTap deleteTap
type TapID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TapInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTap updateTap
type TapInput struct {
	// The Tap to submit or modify
	// in: body
	Tap *orm.TapAPI
}

// GetTaps
//
// swagger:route GET /taps taps getTaps
//
// # Get all taps
//
// Responses:
// default: genericError
//
//	200: tapDBResponse
func (controller *Controller) GetTaps(c *gin.Context) {

	// source slice
	var tapDBs []orm.TapDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTaps", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTap.GetDB()

	query := db.Find(&tapDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tapAPIs := make([]orm.TapAPI, 0)

	// for each tap, update fields from the database nullable fields
	for idx := range tapDBs {
		tapDB := &tapDBs[idx]
		_ = tapDB
		var tapAPI orm.TapAPI

		// insertion point for updating fields
		tapAPI.ID = tapDB.ID
		tapDB.CopyBasicFieldsToTap_WOP(&tapAPI.Tap_WOP)
		tapAPI.TapPointersEncoding = tapDB.TapPointersEncoding
		tapAPIs = append(tapAPIs, tapAPI)
	}

	c.JSON(http.StatusOK, tapAPIs)
}

// PostTap
//
// swagger:route POST /taps taps postTap
//
// Creates a tap
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTap(c *gin.Context) {

	mutexTap.Lock()
	defer mutexTap.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTaps", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTap.GetDB()

	// Validate input
	var input orm.TapAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tap
	tapDB := orm.TapDB{}
	tapDB.TapPointersEncoding = input.TapPointersEncoding
	tapDB.CopyBasicFieldsFromTap_WOP(&input.Tap_WOP)

	query := db.Create(&tapDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTap.CheckoutPhaseOneInstance(&tapDB)
	tap := backRepo.BackRepoTap.Map_TapDBID_TapPtr[tapDB.ID]

	if tap != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tap)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tapDB)
}

// GetTap
//
// swagger:route GET /taps/{ID} taps getTap
//
// Gets the details for a tap.
//
// Responses:
// default: genericError
//
//	200: tapDBResponse
func (controller *Controller) GetTap(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTap", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTap.GetDB()

	// Get tapDB in DB
	var tapDB orm.TapDB
	if err := db.First(&tapDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tapAPI orm.TapAPI
	tapAPI.ID = tapDB.ID
	tapAPI.TapPointersEncoding = tapDB.TapPointersEncoding
	tapDB.CopyBasicFieldsToTap_WOP(&tapAPI.Tap_WOP)

	c.JSON(http.StatusOK, tapAPI)
}

// UpdateTap
//
// swagger:route PATCH /taps/{ID} taps updateTap
//
// # Update a tap
//
// Responses:
// default: genericError
//
//	200: tapDBResponse
func (controller *Controller) UpdateTap(c *gin.Context) {

	mutexTap.Lock()
	defer mutexTap.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTap", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTap.GetDB()

	// Validate input
	var input orm.TapAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tapDB orm.TapDB

	// fetch the tap
	query := db.First(&tapDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tapDB.CopyBasicFieldsFromTap_WOP(&input.Tap_WOP)
	tapDB.TapPointersEncoding = input.TapPointersEncoding

	query = db.Model(&tapDB).Updates(tapDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tapNew := new(models.Tap)
	tapDB.CopyBasicFieldsToTap(tapNew)

	// redeem pointers
	tapDB.DecodePointers(backRepo, tapNew)

	// get stage instance from DB instance, and call callback function
	tapOld := backRepo.BackRepoTap.Map_TapDBID_TapPtr[tapDB.ID]
	if tapOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tapOld, tapNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tapDB
	c.JSON(http.StatusOK, tapDB)
}

// DeleteTap
//
// swagger:route DELETE /taps/{ID} taps deleteTap
//
// # Delete a tap
//
// default: genericError
//
//	200: tapDBResponse
func (controller *Controller) DeleteTap(c *gin.Context) {

	mutexTap.Lock()
	defer mutexTap.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTap", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTap.GetDB()

	// Get model if exist
	var tapDB orm.TapDB
	if err := db.First(&tapDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tapDB)

	// get an instance (not staged) from DB instance, and call callback function
	tapDeleted := new(models.Tap)
	tapDB.CopyBasicFieldsToTap(tapDeleted)

	// get stage instance from DB instance, and call callback function
	tapStaged := backRepo.BackRepoTap.Map_TapDBID_TapPtr[tapDB.ID]
	if tapStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tapStaged, tapDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
