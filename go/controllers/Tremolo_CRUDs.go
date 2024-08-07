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
var __Tremolo__dummysDeclaration__ models.Tremolo
var __Tremolo_time__dummyDeclaration time.Duration

var mutexTremolo sync.Mutex

// An TremoloID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTremolo updateTremolo deleteTremolo
type TremoloID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TremoloInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTremolo updateTremolo
type TremoloInput struct {
	// The Tremolo to submit or modify
	// in: body
	Tremolo *orm.TremoloAPI
}

// GetTremolos
//
// swagger:route GET /tremolos tremolos getTremolos
//
// # Get all tremolos
//
// Responses:
// default: genericError
//
//	200: tremoloDBResponse
func (controller *Controller) GetTremolos(c *gin.Context) {

	// source slice
	var tremoloDBs []orm.TremoloDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTremolos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTremolo.GetDB()

	query := db.Find(&tremoloDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tremoloAPIs := make([]orm.TremoloAPI, 0)

	// for each tremolo, update fields from the database nullable fields
	for idx := range tremoloDBs {
		tremoloDB := &tremoloDBs[idx]
		_ = tremoloDB
		var tremoloAPI orm.TremoloAPI

		// insertion point for updating fields
		tremoloAPI.ID = tremoloDB.ID
		tremoloDB.CopyBasicFieldsToTremolo_WOP(&tremoloAPI.Tremolo_WOP)
		tremoloAPI.TremoloPointersEncoding = tremoloDB.TremoloPointersEncoding
		tremoloAPIs = append(tremoloAPIs, tremoloAPI)
	}

	c.JSON(http.StatusOK, tremoloAPIs)
}

// PostTremolo
//
// swagger:route POST /tremolos tremolos postTremolo
//
// Creates a tremolo
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTremolo(c *gin.Context) {

	mutexTremolo.Lock()
	defer mutexTremolo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTremolos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTremolo.GetDB()

	// Validate input
	var input orm.TremoloAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tremolo
	tremoloDB := orm.TremoloDB{}
	tremoloDB.TremoloPointersEncoding = input.TremoloPointersEncoding
	tremoloDB.CopyBasicFieldsFromTremolo_WOP(&input.Tremolo_WOP)

	query := db.Create(&tremoloDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTremolo.CheckoutPhaseOneInstance(&tremoloDB)
	tremolo := backRepo.BackRepoTremolo.Map_TremoloDBID_TremoloPtr[tremoloDB.ID]

	if tremolo != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tremolo)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tremoloDB)
}

// GetTremolo
//
// swagger:route GET /tremolos/{ID} tremolos getTremolo
//
// Gets the details for a tremolo.
//
// Responses:
// default: genericError
//
//	200: tremoloDBResponse
func (controller *Controller) GetTremolo(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTremolo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTremolo.GetDB()

	// Get tremoloDB in DB
	var tremoloDB orm.TremoloDB
	if err := db.First(&tremoloDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tremoloAPI orm.TremoloAPI
	tremoloAPI.ID = tremoloDB.ID
	tremoloAPI.TremoloPointersEncoding = tremoloDB.TremoloPointersEncoding
	tremoloDB.CopyBasicFieldsToTremolo_WOP(&tremoloAPI.Tremolo_WOP)

	c.JSON(http.StatusOK, tremoloAPI)
}

// UpdateTremolo
//
// swagger:route PATCH /tremolos/{ID} tremolos updateTremolo
//
// # Update a tremolo
//
// Responses:
// default: genericError
//
//	200: tremoloDBResponse
func (controller *Controller) UpdateTremolo(c *gin.Context) {

	mutexTremolo.Lock()
	defer mutexTremolo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTremolo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTremolo.GetDB()

	// Validate input
	var input orm.TremoloAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tremoloDB orm.TremoloDB

	// fetch the tremolo
	query := db.First(&tremoloDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tremoloDB.CopyBasicFieldsFromTremolo_WOP(&input.Tremolo_WOP)
	tremoloDB.TremoloPointersEncoding = input.TremoloPointersEncoding

	query = db.Model(&tremoloDB).Updates(tremoloDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tremoloNew := new(models.Tremolo)
	tremoloDB.CopyBasicFieldsToTremolo(tremoloNew)

	// redeem pointers
	tremoloDB.DecodePointers(backRepo, tremoloNew)

	// get stage instance from DB instance, and call callback function
	tremoloOld := backRepo.BackRepoTremolo.Map_TremoloDBID_TremoloPtr[tremoloDB.ID]
	if tremoloOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tremoloOld, tremoloNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tremoloDB
	c.JSON(http.StatusOK, tremoloDB)
}

// DeleteTremolo
//
// swagger:route DELETE /tremolos/{ID} tremolos deleteTremolo
//
// # Delete a tremolo
//
// default: genericError
//
//	200: tremoloDBResponse
func (controller *Controller) DeleteTremolo(c *gin.Context) {

	mutexTremolo.Lock()
	defer mutexTremolo.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTremolo", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTremolo.GetDB()

	// Get model if exist
	var tremoloDB orm.TremoloDB
	if err := db.First(&tremoloDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tremoloDB)

	// get an instance (not staged) from DB instance, and call callback function
	tremoloDeleted := new(models.Tremolo)
	tremoloDB.CopyBasicFieldsToTremolo(tremoloDeleted)

	// get stage instance from DB instance, and call callback function
	tremoloStaged := backRepo.BackRepoTremolo.Map_TremoloDBID_TremoloPtr[tremoloDB.ID]
	if tremoloStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tremoloStaged, tremoloDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
