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
var __Articulations__dummysDeclaration__ models.Articulations
var __Articulations_time__dummyDeclaration time.Duration

var mutexArticulations sync.Mutex

// An ArticulationsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getArticulations updateArticulations deleteArticulations
type ArticulationsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ArticulationsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postArticulations updateArticulations
type ArticulationsInput struct {
	// The Articulations to submit or modify
	// in: body
	Articulations *orm.ArticulationsAPI
}

// GetArticulationss
//
// swagger:route GET /articulationss articulationss getArticulationss
//
// # Get all articulationss
//
// Responses:
// default: genericError
//
//	200: articulationsDBResponse
func (controller *Controller) GetArticulationss(c *gin.Context) {

	// source slice
	var articulationsDBs []orm.ArticulationsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArticulationss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArticulations.GetDB()

	query := db.Find(&articulationsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	articulationsAPIs := make([]orm.ArticulationsAPI, 0)

	// for each articulations, update fields from the database nullable fields
	for idx := range articulationsDBs {
		articulationsDB := &articulationsDBs[idx]
		_ = articulationsDB
		var articulationsAPI orm.ArticulationsAPI

		// insertion point for updating fields
		articulationsAPI.ID = articulationsDB.ID
		articulationsDB.CopyBasicFieldsToArticulations_WOP(&articulationsAPI.Articulations_WOP)
		articulationsAPI.ArticulationsPointersEncoding = articulationsDB.ArticulationsPointersEncoding
		articulationsAPIs = append(articulationsAPIs, articulationsAPI)
	}

	c.JSON(http.StatusOK, articulationsAPIs)
}

// PostArticulations
//
// swagger:route POST /articulationss articulationss postArticulations
//
// Creates a articulations
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostArticulations(c *gin.Context) {

	mutexArticulations.Lock()
	defer mutexArticulations.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostArticulationss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArticulations.GetDB()

	// Validate input
	var input orm.ArticulationsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create articulations
	articulationsDB := orm.ArticulationsDB{}
	articulationsDB.ArticulationsPointersEncoding = input.ArticulationsPointersEncoding
	articulationsDB.CopyBasicFieldsFromArticulations_WOP(&input.Articulations_WOP)

	query := db.Create(&articulationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoArticulations.CheckoutPhaseOneInstance(&articulationsDB)
	articulations := backRepo.BackRepoArticulations.Map_ArticulationsDBID_ArticulationsPtr[articulationsDB.ID]

	if articulations != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), articulations)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, articulationsDB)
}

// GetArticulations
//
// swagger:route GET /articulationss/{ID} articulationss getArticulations
//
// Gets the details for a articulations.
//
// Responses:
// default: genericError
//
//	200: articulationsDBResponse
func (controller *Controller) GetArticulations(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetArticulations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArticulations.GetDB()

	// Get articulationsDB in DB
	var articulationsDB orm.ArticulationsDB
	if err := db.First(&articulationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var articulationsAPI orm.ArticulationsAPI
	articulationsAPI.ID = articulationsDB.ID
	articulationsAPI.ArticulationsPointersEncoding = articulationsDB.ArticulationsPointersEncoding
	articulationsDB.CopyBasicFieldsToArticulations_WOP(&articulationsAPI.Articulations_WOP)

	c.JSON(http.StatusOK, articulationsAPI)
}

// UpdateArticulations
//
// swagger:route PATCH /articulationss/{ID} articulationss updateArticulations
//
// # Update a articulations
//
// Responses:
// default: genericError
//
//	200: articulationsDBResponse
func (controller *Controller) UpdateArticulations(c *gin.Context) {

	mutexArticulations.Lock()
	defer mutexArticulations.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateArticulations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArticulations.GetDB()

	// Validate input
	var input orm.ArticulationsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var articulationsDB orm.ArticulationsDB

	// fetch the articulations
	query := db.First(&articulationsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	articulationsDB.CopyBasicFieldsFromArticulations_WOP(&input.Articulations_WOP)
	articulationsDB.ArticulationsPointersEncoding = input.ArticulationsPointersEncoding

	query = db.Model(&articulationsDB).Updates(articulationsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	articulationsNew := new(models.Articulations)
	articulationsDB.CopyBasicFieldsToArticulations(articulationsNew)

	// redeem pointers
	articulationsDB.DecodePointers(backRepo, articulationsNew)

	// get stage instance from DB instance, and call callback function
	articulationsOld := backRepo.BackRepoArticulations.Map_ArticulationsDBID_ArticulationsPtr[articulationsDB.ID]
	if articulationsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), articulationsOld, articulationsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the articulationsDB
	c.JSON(http.StatusOK, articulationsDB)
}

// DeleteArticulations
//
// swagger:route DELETE /articulationss/{ID} articulationss deleteArticulations
//
// # Delete a articulations
//
// default: genericError
//
//	200: articulationsDBResponse
func (controller *Controller) DeleteArticulations(c *gin.Context) {

	mutexArticulations.Lock()
	defer mutexArticulations.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteArticulations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoArticulations.GetDB()

	// Get model if exist
	var articulationsDB orm.ArticulationsDB
	if err := db.First(&articulationsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&articulationsDB)

	// get an instance (not staged) from DB instance, and call callback function
	articulationsDeleted := new(models.Articulations)
	articulationsDB.CopyBasicFieldsToArticulations(articulationsDeleted)

	// get stage instance from DB instance, and call callback function
	articulationsStaged := backRepo.BackRepoArticulations.Map_ArticulationsDBID_ArticulationsPtr[articulationsDB.ID]
	if articulationsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), articulationsStaged, articulationsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
