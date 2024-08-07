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
var __Metronome__dummysDeclaration__ models.Metronome
var __Metronome_time__dummyDeclaration time.Duration

var mutexMetronome sync.Mutex

// An MetronomeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMetronome updateMetronome deleteMetronome
type MetronomeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MetronomeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMetronome updateMetronome
type MetronomeInput struct {
	// The Metronome to submit or modify
	// in: body
	Metronome *orm.MetronomeAPI
}

// GetMetronomes
//
// swagger:route GET /metronomes metronomes getMetronomes
//
// # Get all metronomes
//
// Responses:
// default: genericError
//
//	200: metronomeDBResponse
func (controller *Controller) GetMetronomes(c *gin.Context) {

	// source slice
	var metronomeDBs []orm.MetronomeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronomes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome.GetDB()

	query := db.Find(&metronomeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metronomeAPIs := make([]orm.MetronomeAPI, 0)

	// for each metronome, update fields from the database nullable fields
	for idx := range metronomeDBs {
		metronomeDB := &metronomeDBs[idx]
		_ = metronomeDB
		var metronomeAPI orm.MetronomeAPI

		// insertion point for updating fields
		metronomeAPI.ID = metronomeDB.ID
		metronomeDB.CopyBasicFieldsToMetronome_WOP(&metronomeAPI.Metronome_WOP)
		metronomeAPI.MetronomePointersEncoding = metronomeDB.MetronomePointersEncoding
		metronomeAPIs = append(metronomeAPIs, metronomeAPI)
	}

	c.JSON(http.StatusOK, metronomeAPIs)
}

// PostMetronome
//
// swagger:route POST /metronomes metronomes postMetronome
//
// Creates a metronome
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMetronome(c *gin.Context) {

	mutexMetronome.Lock()
	defer mutexMetronome.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetronomes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome.GetDB()

	// Validate input
	var input orm.MetronomeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create metronome
	metronomeDB := orm.MetronomeDB{}
	metronomeDB.MetronomePointersEncoding = input.MetronomePointersEncoding
	metronomeDB.CopyBasicFieldsFromMetronome_WOP(&input.Metronome_WOP)

	query := db.Create(&metronomeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMetronome.CheckoutPhaseOneInstance(&metronomeDB)
	metronome := backRepo.BackRepoMetronome.Map_MetronomeDBID_MetronomePtr[metronomeDB.ID]

	if metronome != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), metronome)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metronomeDB)
}

// GetMetronome
//
// swagger:route GET /metronomes/{ID} metronomes getMetronome
//
// Gets the details for a metronome.
//
// Responses:
// default: genericError
//
//	200: metronomeDBResponse
func (controller *Controller) GetMetronome(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetronome", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome.GetDB()

	// Get metronomeDB in DB
	var metronomeDB orm.MetronomeDB
	if err := db.First(&metronomeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metronomeAPI orm.MetronomeAPI
	metronomeAPI.ID = metronomeDB.ID
	metronomeAPI.MetronomePointersEncoding = metronomeDB.MetronomePointersEncoding
	metronomeDB.CopyBasicFieldsToMetronome_WOP(&metronomeAPI.Metronome_WOP)

	c.JSON(http.StatusOK, metronomeAPI)
}

// UpdateMetronome
//
// swagger:route PATCH /metronomes/{ID} metronomes updateMetronome
//
// # Update a metronome
//
// Responses:
// default: genericError
//
//	200: metronomeDBResponse
func (controller *Controller) UpdateMetronome(c *gin.Context) {

	mutexMetronome.Lock()
	defer mutexMetronome.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMetronome", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome.GetDB()

	// Validate input
	var input orm.MetronomeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metronomeDB orm.MetronomeDB

	// fetch the metronome
	query := db.First(&metronomeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metronomeDB.CopyBasicFieldsFromMetronome_WOP(&input.Metronome_WOP)
	metronomeDB.MetronomePointersEncoding = input.MetronomePointersEncoding

	query = db.Model(&metronomeDB).Updates(metronomeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metronomeNew := new(models.Metronome)
	metronomeDB.CopyBasicFieldsToMetronome(metronomeNew)

	// redeem pointers
	metronomeDB.DecodePointers(backRepo, metronomeNew)

	// get stage instance from DB instance, and call callback function
	metronomeOld := backRepo.BackRepoMetronome.Map_MetronomeDBID_MetronomePtr[metronomeDB.ID]
	if metronomeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metronomeOld, metronomeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metronomeDB
	c.JSON(http.StatusOK, metronomeDB)
}

// DeleteMetronome
//
// swagger:route DELETE /metronomes/{ID} metronomes deleteMetronome
//
// # Delete a metronome
//
// default: genericError
//
//	200: metronomeDBResponse
func (controller *Controller) DeleteMetronome(c *gin.Context) {

	mutexMetronome.Lock()
	defer mutexMetronome.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMetronome", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMetronome.GetDB()

	// Get model if exist
	var metronomeDB orm.MetronomeDB
	if err := db.First(&metronomeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&metronomeDB)

	// get an instance (not staged) from DB instance, and call callback function
	metronomeDeleted := new(models.Metronome)
	metronomeDB.CopyBasicFieldsToMetronome(metronomeDeleted)

	// get stage instance from DB instance, and call callback function
	metronomeStaged := backRepo.BackRepoMetronome.Map_MetronomeDBID_MetronomePtr[metronomeDB.ID]
	if metronomeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metronomeStaged, metronomeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
