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
var __Membrane__dummysDeclaration__ models.Membrane
var __Membrane_time__dummyDeclaration time.Duration

var mutexMembrane sync.Mutex

// An MembraneID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMembrane updateMembrane deleteMembrane
type MembraneID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MembraneInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMembrane updateMembrane
type MembraneInput struct {
	// The Membrane to submit or modify
	// in: body
	Membrane *orm.MembraneAPI
}

// GetMembranes
//
// swagger:route GET /membranes membranes getMembranes
//
// # Get all membranes
//
// Responses:
// default: genericError
//
//	200: membraneDBResponse
func (controller *Controller) GetMembranes(c *gin.Context) {

	// source slice
	var membraneDBs []orm.MembraneDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMembranes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMembrane.GetDB()

	query := db.Find(&membraneDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	membraneAPIs := make([]orm.MembraneAPI, 0)

	// for each membrane, update fields from the database nullable fields
	for idx := range membraneDBs {
		membraneDB := &membraneDBs[idx]
		_ = membraneDB
		var membraneAPI orm.MembraneAPI

		// insertion point for updating fields
		membraneAPI.ID = membraneDB.ID
		membraneDB.CopyBasicFieldsToMembrane_WOP(&membraneAPI.Membrane_WOP)
		membraneAPI.MembranePointersEncoding = membraneDB.MembranePointersEncoding
		membraneAPIs = append(membraneAPIs, membraneAPI)
	}

	c.JSON(http.StatusOK, membraneAPIs)
}

// PostMembrane
//
// swagger:route POST /membranes membranes postMembrane
//
// Creates a membrane
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMembrane(c *gin.Context) {

	mutexMembrane.Lock()
	defer mutexMembrane.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMembranes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMembrane.GetDB()

	// Validate input
	var input orm.MembraneAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create membrane
	membraneDB := orm.MembraneDB{}
	membraneDB.MembranePointersEncoding = input.MembranePointersEncoding
	membraneDB.CopyBasicFieldsFromMembrane_WOP(&input.Membrane_WOP)

	query := db.Create(&membraneDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMembrane.CheckoutPhaseOneInstance(&membraneDB)
	membrane := backRepo.BackRepoMembrane.Map_MembraneDBID_MembranePtr[membraneDB.ID]

	if membrane != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), membrane)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, membraneDB)
}

// GetMembrane
//
// swagger:route GET /membranes/{ID} membranes getMembrane
//
// Gets the details for a membrane.
//
// Responses:
// default: genericError
//
//	200: membraneDBResponse
func (controller *Controller) GetMembrane(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMembrane", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMembrane.GetDB()

	// Get membraneDB in DB
	var membraneDB orm.MembraneDB
	if err := db.First(&membraneDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var membraneAPI orm.MembraneAPI
	membraneAPI.ID = membraneDB.ID
	membraneAPI.MembranePointersEncoding = membraneDB.MembranePointersEncoding
	membraneDB.CopyBasicFieldsToMembrane_WOP(&membraneAPI.Membrane_WOP)

	c.JSON(http.StatusOK, membraneAPI)
}

// UpdateMembrane
//
// swagger:route PATCH /membranes/{ID} membranes updateMembrane
//
// # Update a membrane
//
// Responses:
// default: genericError
//
//	200: membraneDBResponse
func (controller *Controller) UpdateMembrane(c *gin.Context) {

	mutexMembrane.Lock()
	defer mutexMembrane.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMembrane", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMembrane.GetDB()

	// Validate input
	var input orm.MembraneAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var membraneDB orm.MembraneDB

	// fetch the membrane
	query := db.First(&membraneDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	membraneDB.CopyBasicFieldsFromMembrane_WOP(&input.Membrane_WOP)
	membraneDB.MembranePointersEncoding = input.MembranePointersEncoding

	query = db.Model(&membraneDB).Updates(membraneDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	membraneNew := new(models.Membrane)
	membraneDB.CopyBasicFieldsToMembrane(membraneNew)

	// redeem pointers
	membraneDB.DecodePointers(backRepo, membraneNew)

	// get stage instance from DB instance, and call callback function
	membraneOld := backRepo.BackRepoMembrane.Map_MembraneDBID_MembranePtr[membraneDB.ID]
	if membraneOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), membraneOld, membraneNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the membraneDB
	c.JSON(http.StatusOK, membraneDB)
}

// DeleteMembrane
//
// swagger:route DELETE /membranes/{ID} membranes deleteMembrane
//
// # Delete a membrane
//
// default: genericError
//
//	200: membraneDBResponse
func (controller *Controller) DeleteMembrane(c *gin.Context) {

	mutexMembrane.Lock()
	defer mutexMembrane.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMembrane", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoMembrane.GetDB()

	// Get model if exist
	var membraneDB orm.MembraneDB
	if err := db.First(&membraneDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&membraneDB)

	// get an instance (not staged) from DB instance, and call callback function
	membraneDeleted := new(models.Membrane)
	membraneDB.CopyBasicFieldsToMembrane(membraneDeleted)

	// get stage instance from DB instance, and call callback function
	membraneStaged := backRepo.BackRepoMembrane.Map_MembraneDBID_MembranePtr[membraneDB.ID]
	if membraneStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), membraneStaged, membraneDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
