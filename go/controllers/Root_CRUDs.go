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
var __Root__dummysDeclaration__ models.Root
var __Root_time__dummyDeclaration time.Duration

var mutexRoot sync.Mutex

// An RootID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRoot updateRoot deleteRoot
type RootID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RootInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRoot updateRoot
type RootInput struct {
	// The Root to submit or modify
	// in: body
	Root *orm.RootAPI
}

// GetRoots
//
// swagger:route GET /roots roots getRoots
//
// # Get all roots
//
// Responses:
// default: genericError
//
//	200: rootDBResponse
func (controller *Controller) GetRoots(c *gin.Context) {

	// source slice
	var rootDBs []orm.RootDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRoots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot.GetDB()

	query := db.Find(&rootDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rootAPIs := make([]orm.RootAPI, 0)

	// for each root, update fields from the database nullable fields
	for idx := range rootDBs {
		rootDB := &rootDBs[idx]
		_ = rootDB
		var rootAPI orm.RootAPI

		// insertion point for updating fields
		rootAPI.ID = rootDB.ID
		rootDB.CopyBasicFieldsToRoot_WOP(&rootAPI.Root_WOP)
		rootAPI.RootPointersEncoding = rootDB.RootPointersEncoding
		rootAPIs = append(rootAPIs, rootAPI)
	}

	c.JSON(http.StatusOK, rootAPIs)
}

// PostRoot
//
// swagger:route POST /roots roots postRoot
//
// Creates a root
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRoot(c *gin.Context) {

	mutexRoot.Lock()
	defer mutexRoot.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRoots", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot.GetDB()

	// Validate input
	var input orm.RootAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create root
	rootDB := orm.RootDB{}
	rootDB.RootPointersEncoding = input.RootPointersEncoding
	rootDB.CopyBasicFieldsFromRoot_WOP(&input.Root_WOP)

	query := db.Create(&rootDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRoot.CheckoutPhaseOneInstance(&rootDB)
	root := backRepo.BackRepoRoot.Map_RootDBID_RootPtr[rootDB.ID]

	if root != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), root)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rootDB)
}

// GetRoot
//
// swagger:route GET /roots/{ID} roots getRoot
//
// Gets the details for a root.
//
// Responses:
// default: genericError
//
//	200: rootDBResponse
func (controller *Controller) GetRoot(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRoot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot.GetDB()

	// Get rootDB in DB
	var rootDB orm.RootDB
	if err := db.First(&rootDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rootAPI orm.RootAPI
	rootAPI.ID = rootDB.ID
	rootAPI.RootPointersEncoding = rootDB.RootPointersEncoding
	rootDB.CopyBasicFieldsToRoot_WOP(&rootAPI.Root_WOP)

	c.JSON(http.StatusOK, rootAPI)
}

// UpdateRoot
//
// swagger:route PATCH /roots/{ID} roots updateRoot
//
// # Update a root
//
// Responses:
// default: genericError
//
//	200: rootDBResponse
func (controller *Controller) UpdateRoot(c *gin.Context) {

	mutexRoot.Lock()
	defer mutexRoot.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRoot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot.GetDB()

	// Validate input
	var input orm.RootAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rootDB orm.RootDB

	// fetch the root
	query := db.First(&rootDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rootDB.CopyBasicFieldsFromRoot_WOP(&input.Root_WOP)
	rootDB.RootPointersEncoding = input.RootPointersEncoding

	query = db.Model(&rootDB).Updates(rootDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rootNew := new(models.Root)
	rootDB.CopyBasicFieldsToRoot(rootNew)

	// redeem pointers
	rootDB.DecodePointers(backRepo, rootNew)

	// get stage instance from DB instance, and call callback function
	rootOld := backRepo.BackRepoRoot.Map_RootDBID_RootPtr[rootDB.ID]
	if rootOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rootOld, rootNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rootDB
	c.JSON(http.StatusOK, rootDB)
}

// DeleteRoot
//
// swagger:route DELETE /roots/{ID} roots deleteRoot
//
// # Delete a root
//
// default: genericError
//
//	200: rootDBResponse
func (controller *Controller) DeleteRoot(c *gin.Context) {

	mutexRoot.Lock()
	defer mutexRoot.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRoot", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot.GetDB()

	// Get model if exist
	var rootDB orm.RootDB
	if err := db.First(&rootDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&rootDB)

	// get an instance (not staged) from DB instance, and call callback function
	rootDeleted := new(models.Root)
	rootDB.CopyBasicFieldsToRoot(rootDeleted)

	// get stage instance from DB instance, and call callback function
	rootStaged := backRepo.BackRepoRoot.Map_RootDBID_RootPtr[rootDB.ID]
	if rootStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rootStaged, rootDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
