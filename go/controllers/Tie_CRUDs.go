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
var __Tie__dummysDeclaration__ models.Tie
var __Tie_time__dummyDeclaration time.Duration

var mutexTie sync.Mutex

// An TieID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTie updateTie deleteTie
type TieID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TieInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTie updateTie
type TieInput struct {
	// The Tie to submit or modify
	// in: body
	Tie *orm.TieAPI
}

// GetTies
//
// swagger:route GET /ties ties getTies
//
// # Get all ties
//
// Responses:
// default: genericError
//
//	200: tieDBResponse
func (controller *Controller) GetTies(c *gin.Context) {

	// source slice
	var tieDBs []orm.TieDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTies", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTie.GetDB()

	query := db.Find(&tieDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	tieAPIs := make([]orm.TieAPI, 0)

	// for each tie, update fields from the database nullable fields
	for idx := range tieDBs {
		tieDB := &tieDBs[idx]
		_ = tieDB
		var tieAPI orm.TieAPI

		// insertion point for updating fields
		tieAPI.ID = tieDB.ID
		tieDB.CopyBasicFieldsToTie_WOP(&tieAPI.Tie_WOP)
		tieAPI.TiePointersEncoding = tieDB.TiePointersEncoding
		tieAPIs = append(tieAPIs, tieAPI)
	}

	c.JSON(http.StatusOK, tieAPIs)
}

// PostTie
//
// swagger:route POST /ties ties postTie
//
// Creates a tie
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTie(c *gin.Context) {

	mutexTie.Lock()
	defer mutexTie.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTies", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTie.GetDB()

	// Validate input
	var input orm.TieAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create tie
	tieDB := orm.TieDB{}
	tieDB.TiePointersEncoding = input.TiePointersEncoding
	tieDB.CopyBasicFieldsFromTie_WOP(&input.Tie_WOP)

	query := db.Create(&tieDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTie.CheckoutPhaseOneInstance(&tieDB)
	tie := backRepo.BackRepoTie.Map_TieDBID_TiePtr[tieDB.ID]

	if tie != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), tie)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, tieDB)
}

// GetTie
//
// swagger:route GET /ties/{ID} ties getTie
//
// Gets the details for a tie.
//
// Responses:
// default: genericError
//
//	200: tieDBResponse
func (controller *Controller) GetTie(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTie.GetDB()

	// Get tieDB in DB
	var tieDB orm.TieDB
	if err := db.First(&tieDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var tieAPI orm.TieAPI
	tieAPI.ID = tieDB.ID
	tieAPI.TiePointersEncoding = tieDB.TiePointersEncoding
	tieDB.CopyBasicFieldsToTie_WOP(&tieAPI.Tie_WOP)

	c.JSON(http.StatusOK, tieAPI)
}

// UpdateTie
//
// swagger:route PATCH /ties/{ID} ties updateTie
//
// # Update a tie
//
// Responses:
// default: genericError
//
//	200: tieDBResponse
func (controller *Controller) UpdateTie(c *gin.Context) {

	mutexTie.Lock()
	defer mutexTie.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTie.GetDB()

	// Validate input
	var input orm.TieAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var tieDB orm.TieDB

	// fetch the tie
	query := db.First(&tieDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	tieDB.CopyBasicFieldsFromTie_WOP(&input.Tie_WOP)
	tieDB.TiePointersEncoding = input.TiePointersEncoding

	query = db.Model(&tieDB).Updates(tieDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	tieNew := new(models.Tie)
	tieDB.CopyBasicFieldsToTie(tieNew)

	// redeem pointers
	tieDB.DecodePointers(backRepo, tieNew)

	// get stage instance from DB instance, and call callback function
	tieOld := backRepo.BackRepoTie.Map_TieDBID_TiePtr[tieDB.ID]
	if tieOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), tieOld, tieNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the tieDB
	c.JSON(http.StatusOK, tieDB)
}

// DeleteTie
//
// swagger:route DELETE /ties/{ID} ties deleteTie
//
// # Delete a tie
//
// default: genericError
//
//	200: tieDBResponse
func (controller *Controller) DeleteTie(c *gin.Context) {

	mutexTie.Lock()
	defer mutexTie.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTie", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTie.GetDB()

	// Get model if exist
	var tieDB orm.TieDB
	if err := db.First(&tieDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&tieDB)

	// get an instance (not staged) from DB instance, and call callback function
	tieDeleted := new(models.Tie)
	tieDB.CopyBasicFieldsToTie(tieDeleted)

	// get stage instance from DB instance, and call callback function
	tieStaged := backRepo.BackRepoTie.Map_TieDBID_TiePtr[tieDB.ID]
	if tieStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), tieStaged, tieDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
