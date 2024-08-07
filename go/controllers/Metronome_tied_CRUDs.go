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
var __Metronome_tied__dummysDeclaration__ models.Metronome_tied
var __Metronome_tied_time__dummyDeclaration time.Duration

var mutexMetronome_tied sync.Mutex

// An Metronome_tiedID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMetronome_tied updateMetronome_tied deleteMetronome_tied
type Metronome_tiedID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Metronome_tiedInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMetronome_tied updateMetronome_tied
type Metronome_tiedInput struct {
	// The Metronome_tied to submit or modify
	// in: body
	Metronome_tied *orm.Metronome_tiedAPI
}

// GetMetronome_tieds
//
// swagger:route GET /metronome_tieds metronome_tieds getMetronome_tieds
//
// # Get all metronome_tieds
//
// Responses:
// default: genericError
//
//	200: metronome_tiedDBResponse
func (controller *Controller) GetMetronome_tieds(c *gin.Context) {

	// source slice
	var metronome_tiedDBs []orm.Metronome_tiedDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_tieds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tied.GetDB()

	query := db.Find(&metronome_tiedDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metronome_tiedAPIs := make([]orm.Metronome_tiedAPI, 0)

	// for each metronome_tied, update fields from the database nullable fields
	for idx := range metronome_tiedDBs {
		metronome_tiedDB := &metronome_tiedDBs[idx]
		_ = metronome_tiedDB
		var metronome_tiedAPI orm.Metronome_tiedAPI

		// insertion point for updating fields
		metronome_tiedAPI.ID = metronome_tiedDB.ID
		metronome_tiedDB.CopyBasicFieldsToMetronome_tied_WOP(&metronome_tiedAPI.Metronome_tied_WOP)
		metronome_tiedAPI.Metronome_tiedPointersEncoding = metronome_tiedDB.Metronome_tiedPointersEncoding
		metronome_tiedAPIs = append(metronome_tiedAPIs, metronome_tiedAPI)
	}

	c.JSON(http.StatusOK, metronome_tiedAPIs)
}

// PostMetronome_tied
//
// swagger:route POST /metronome_tieds metronome_tieds postMetronome_tied
//
// Creates a metronome_tied
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMetronome_tied(c *gin.Context) {

	mutexMetronome_tied.Lock()
	defer mutexMetronome_tied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetronome_tieds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tied.GetDB()

	// Validate input
	var input orm.Metronome_tiedAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create metronome_tied
	metronome_tiedDB := orm.Metronome_tiedDB{}
	metronome_tiedDB.Metronome_tiedPointersEncoding = input.Metronome_tiedPointersEncoding
	metronome_tiedDB.CopyBasicFieldsFromMetronome_tied_WOP(&input.Metronome_tied_WOP)

	query := db.Create(&metronome_tiedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMetronome_tied.CheckoutPhaseOneInstance(&metronome_tiedDB)
	metronome_tied := backRepo.BackRepoMetronome_tied.Map_Metronome_tiedDBID_Metronome_tiedPtr[metronome_tiedDB.ID]

	if metronome_tied != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), metronome_tied)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metronome_tiedDB)
}

// GetMetronome_tied
//
// swagger:route GET /metronome_tieds/{ID} metronome_tieds getMetronome_tied
//
// Gets the details for a metronome_tied.
//
// Responses:
// default: genericError
//
//	200: metronome_tiedDBResponse
func (controller *Controller) GetMetronome_tied(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome_tied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tied.GetDB()

	// Get metronome_tiedDB in DB
	var metronome_tiedDB orm.Metronome_tiedDB
	if err := db.First(&metronome_tiedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metronome_tiedAPI orm.Metronome_tiedAPI
	metronome_tiedAPI.ID = metronome_tiedDB.ID
	metronome_tiedAPI.Metronome_tiedPointersEncoding = metronome_tiedDB.Metronome_tiedPointersEncoding
	metronome_tiedDB.CopyBasicFieldsToMetronome_tied_WOP(&metronome_tiedAPI.Metronome_tied_WOP)

	c.JSON(http.StatusOK, metronome_tiedAPI)
}

// UpdateMetronome_tied
//
// swagger:route PATCH /metronome_tieds/{ID} metronome_tieds updateMetronome_tied
//
// # Update a metronome_tied
//
// Responses:
// default: genericError
//
//	200: metronome_tiedDBResponse
func (controller *Controller) UpdateMetronome_tied(c *gin.Context) {

	mutexMetronome_tied.Lock()
	defer mutexMetronome_tied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMetronome_tied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tied.GetDB()

	// Validate input
	var input orm.Metronome_tiedAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metronome_tiedDB orm.Metronome_tiedDB

	// fetch the metronome_tied
	query := db.First(&metronome_tiedDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metronome_tiedDB.CopyBasicFieldsFromMetronome_tied_WOP(&input.Metronome_tied_WOP)
	metronome_tiedDB.Metronome_tiedPointersEncoding = input.Metronome_tiedPointersEncoding

	query = db.Model(&metronome_tiedDB).Updates(metronome_tiedDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metronome_tiedNew := new(models.Metronome_tied)
	metronome_tiedDB.CopyBasicFieldsToMetronome_tied(metronome_tiedNew)

	// redeem pointers
	metronome_tiedDB.DecodePointers(backRepo, metronome_tiedNew)

	// get stage instance from DB instance, and call callback function
	metronome_tiedOld := backRepo.BackRepoMetronome_tied.Map_Metronome_tiedDBID_Metronome_tiedPtr[metronome_tiedDB.ID]
	if metronome_tiedOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metronome_tiedOld, metronome_tiedNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metronome_tiedDB
	c.JSON(http.StatusOK, metronome_tiedDB)
}

// DeleteMetronome_tied
//
// swagger:route DELETE /metronome_tieds/{ID} metronome_tieds deleteMetronome_tied
//
// # Delete a metronome_tied
//
// default: genericError
//
//	200: metronome_tiedDBResponse
func (controller *Controller) DeleteMetronome_tied(c *gin.Context) {

	mutexMetronome_tied.Lock()
	defer mutexMetronome_tied.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMetronome_tied", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome_tied.GetDB()

	// Get model if exist
	var metronome_tiedDB orm.Metronome_tiedDB
	if err := db.First(&metronome_tiedDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&metronome_tiedDB)

	// get an instance (not staged) from DB instance, and call callback function
	metronome_tiedDeleted := new(models.Metronome_tied)
	metronome_tiedDB.CopyBasicFieldsToMetronome_tied(metronome_tiedDeleted)

	// get stage instance from DB instance, and call callback function
	metronome_tiedStaged := backRepo.BackRepoMetronome_tied.Map_Metronome_tiedDBID_Metronome_tiedPtr[metronome_tiedDB.ID]
	if metronome_tiedStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metronome_tiedStaged, metronome_tiedDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
