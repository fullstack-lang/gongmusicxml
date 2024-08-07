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
var __Harmonic__dummysDeclaration__ models.Harmonic
var __Harmonic_time__dummyDeclaration time.Duration

var mutexHarmonic sync.Mutex

// An HarmonicID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHarmonic updateHarmonic deleteHarmonic
type HarmonicID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HarmonicInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHarmonic updateHarmonic
type HarmonicInput struct {
	// The Harmonic to submit or modify
	// in: body
	Harmonic *orm.HarmonicAPI
}

// GetHarmonics
//
// swagger:route GET /harmonics harmonics getHarmonics
//
// # Get all harmonics
//
// Responses:
// default: genericError
//
//	200: harmonicDBResponse
func (controller *Controller) GetHarmonics(c *gin.Context) {

	// source slice
	var harmonicDBs []orm.HarmonicDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmonics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmonic.GetDB()

	query := db.Find(&harmonicDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	harmonicAPIs := make([]orm.HarmonicAPI, 0)

	// for each harmonic, update fields from the database nullable fields
	for idx := range harmonicDBs {
		harmonicDB := &harmonicDBs[idx]
		_ = harmonicDB
		var harmonicAPI orm.HarmonicAPI

		// insertion point for updating fields
		harmonicAPI.ID = harmonicDB.ID
		harmonicDB.CopyBasicFieldsToHarmonic_WOP(&harmonicAPI.Harmonic_WOP)
		harmonicAPI.HarmonicPointersEncoding = harmonicDB.HarmonicPointersEncoding
		harmonicAPIs = append(harmonicAPIs, harmonicAPI)
	}

	c.JSON(http.StatusOK, harmonicAPIs)
}

// PostHarmonic
//
// swagger:route POST /harmonics harmonics postHarmonic
//
// Creates a harmonic
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHarmonic(c *gin.Context) {

	mutexHarmonic.Lock()
	defer mutexHarmonic.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHarmonics", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmonic.GetDB()

	// Validate input
	var input orm.HarmonicAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create harmonic
	harmonicDB := orm.HarmonicDB{}
	harmonicDB.HarmonicPointersEncoding = input.HarmonicPointersEncoding
	harmonicDB.CopyBasicFieldsFromHarmonic_WOP(&input.Harmonic_WOP)

	query := db.Create(&harmonicDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHarmonic.CheckoutPhaseOneInstance(&harmonicDB)
	harmonic := backRepo.BackRepoHarmonic.Map_HarmonicDBID_HarmonicPtr[harmonicDB.ID]

	if harmonic != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), harmonic)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, harmonicDB)
}

// GetHarmonic
//
// swagger:route GET /harmonics/{ID} harmonics getHarmonic
//
// Gets the details for a harmonic.
//
// Responses:
// default: genericError
//
//	200: harmonicDBResponse
func (controller *Controller) GetHarmonic(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHarmonic", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmonic.GetDB()

	// Get harmonicDB in DB
	var harmonicDB orm.HarmonicDB
	if err := db.First(&harmonicDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var harmonicAPI orm.HarmonicAPI
	harmonicAPI.ID = harmonicDB.ID
	harmonicAPI.HarmonicPointersEncoding = harmonicDB.HarmonicPointersEncoding
	harmonicDB.CopyBasicFieldsToHarmonic_WOP(&harmonicAPI.Harmonic_WOP)

	c.JSON(http.StatusOK, harmonicAPI)
}

// UpdateHarmonic
//
// swagger:route PATCH /harmonics/{ID} harmonics updateHarmonic
//
// # Update a harmonic
//
// Responses:
// default: genericError
//
//	200: harmonicDBResponse
func (controller *Controller) UpdateHarmonic(c *gin.Context) {

	mutexHarmonic.Lock()
	defer mutexHarmonic.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHarmonic", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmonic.GetDB()

	// Validate input
	var input orm.HarmonicAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var harmonicDB orm.HarmonicDB

	// fetch the harmonic
	query := db.First(&harmonicDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	harmonicDB.CopyBasicFieldsFromHarmonic_WOP(&input.Harmonic_WOP)
	harmonicDB.HarmonicPointersEncoding = input.HarmonicPointersEncoding

	query = db.Model(&harmonicDB).Updates(harmonicDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	harmonicNew := new(models.Harmonic)
	harmonicDB.CopyBasicFieldsToHarmonic(harmonicNew)

	// redeem pointers
	harmonicDB.DecodePointers(backRepo, harmonicNew)

	// get stage instance from DB instance, and call callback function
	harmonicOld := backRepo.BackRepoHarmonic.Map_HarmonicDBID_HarmonicPtr[harmonicDB.ID]
	if harmonicOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), harmonicOld, harmonicNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the harmonicDB
	c.JSON(http.StatusOK, harmonicDB)
}

// DeleteHarmonic
//
// swagger:route DELETE /harmonics/{ID} harmonics deleteHarmonic
//
// # Delete a harmonic
//
// default: genericError
//
//	200: harmonicDBResponse
func (controller *Controller) DeleteHarmonic(c *gin.Context) {

	mutexHarmonic.Lock()
	defer mutexHarmonic.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHarmonic", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHarmonic.GetDB()

	// Get model if exist
	var harmonicDB orm.HarmonicDB
	if err := db.First(&harmonicDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&harmonicDB)

	// get an instance (not staged) from DB instance, and call callback function
	harmonicDeleted := new(models.Harmonic)
	harmonicDB.CopyBasicFieldsToHarmonic(harmonicDeleted)

	// get stage instance from DB instance, and call callback function
	harmonicStaged := backRepo.BackRepoHarmonic.Map_HarmonicDBID_HarmonicPtr[harmonicDB.ID]
	if harmonicStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), harmonicStaged, harmonicDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
