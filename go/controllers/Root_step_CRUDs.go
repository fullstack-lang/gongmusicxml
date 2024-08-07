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
var __Root_step__dummysDeclaration__ models.Root_step
var __Root_step_time__dummyDeclaration time.Duration

var mutexRoot_step sync.Mutex

// An Root_stepID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRoot_step updateRoot_step deleteRoot_step
type Root_stepID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Root_stepInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRoot_step updateRoot_step
type Root_stepInput struct {
	// The Root_step to submit or modify
	// in: body
	Root_step *orm.Root_stepAPI
}

// GetRoot_steps
//
// swagger:route GET /root_steps root_steps getRoot_steps
//
// # Get all root_steps
//
// Responses:
// default: genericError
//
//	200: root_stepDBResponse
func (controller *Controller) GetRoot_steps(c *gin.Context) {

	// source slice
	var root_stepDBs []orm.Root_stepDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRoot_steps", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot_step.GetDB()

	query := db.Find(&root_stepDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	root_stepAPIs := make([]orm.Root_stepAPI, 0)

	// for each root_step, update fields from the database nullable fields
	for idx := range root_stepDBs {
		root_stepDB := &root_stepDBs[idx]
		_ = root_stepDB
		var root_stepAPI orm.Root_stepAPI

		// insertion point for updating fields
		root_stepAPI.ID = root_stepDB.ID
		root_stepDB.CopyBasicFieldsToRoot_step_WOP(&root_stepAPI.Root_step_WOP)
		root_stepAPI.Root_stepPointersEncoding = root_stepDB.Root_stepPointersEncoding
		root_stepAPIs = append(root_stepAPIs, root_stepAPI)
	}

	c.JSON(http.StatusOK, root_stepAPIs)
}

// PostRoot_step
//
// swagger:route POST /root_steps root_steps postRoot_step
//
// Creates a root_step
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRoot_step(c *gin.Context) {

	mutexRoot_step.Lock()
	defer mutexRoot_step.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRoot_steps", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot_step.GetDB()

	// Validate input
	var input orm.Root_stepAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create root_step
	root_stepDB := orm.Root_stepDB{}
	root_stepDB.Root_stepPointersEncoding = input.Root_stepPointersEncoding
	root_stepDB.CopyBasicFieldsFromRoot_step_WOP(&input.Root_step_WOP)

	query := db.Create(&root_stepDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRoot_step.CheckoutPhaseOneInstance(&root_stepDB)
	root_step := backRepo.BackRepoRoot_step.Map_Root_stepDBID_Root_stepPtr[root_stepDB.ID]

	if root_step != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), root_step)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, root_stepDB)
}

// GetRoot_step
//
// swagger:route GET /root_steps/{ID} root_steps getRoot_step
//
// Gets the details for a root_step.
//
// Responses:
// default: genericError
//
//	200: root_stepDBResponse
func (controller *Controller) GetRoot_step(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRoot_step", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot_step.GetDB()

	// Get root_stepDB in DB
	var root_stepDB orm.Root_stepDB
	if err := db.First(&root_stepDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var root_stepAPI orm.Root_stepAPI
	root_stepAPI.ID = root_stepDB.ID
	root_stepAPI.Root_stepPointersEncoding = root_stepDB.Root_stepPointersEncoding
	root_stepDB.CopyBasicFieldsToRoot_step_WOP(&root_stepAPI.Root_step_WOP)

	c.JSON(http.StatusOK, root_stepAPI)
}

// UpdateRoot_step
//
// swagger:route PATCH /root_steps/{ID} root_steps updateRoot_step
//
// # Update a root_step
//
// Responses:
// default: genericError
//
//	200: root_stepDBResponse
func (controller *Controller) UpdateRoot_step(c *gin.Context) {

	mutexRoot_step.Lock()
	defer mutexRoot_step.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRoot_step", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot_step.GetDB()

	// Validate input
	var input orm.Root_stepAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var root_stepDB orm.Root_stepDB

	// fetch the root_step
	query := db.First(&root_stepDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	root_stepDB.CopyBasicFieldsFromRoot_step_WOP(&input.Root_step_WOP)
	root_stepDB.Root_stepPointersEncoding = input.Root_stepPointersEncoding

	query = db.Model(&root_stepDB).Updates(root_stepDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	root_stepNew := new(models.Root_step)
	root_stepDB.CopyBasicFieldsToRoot_step(root_stepNew)

	// redeem pointers
	root_stepDB.DecodePointers(backRepo, root_stepNew)

	// get stage instance from DB instance, and call callback function
	root_stepOld := backRepo.BackRepoRoot_step.Map_Root_stepDBID_Root_stepPtr[root_stepDB.ID]
	if root_stepOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), root_stepOld, root_stepNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the root_stepDB
	c.JSON(http.StatusOK, root_stepDB)
}

// DeleteRoot_step
//
// swagger:route DELETE /root_steps/{ID} root_steps deleteRoot_step
//
// # Delete a root_step
//
// default: genericError
//
//	200: root_stepDBResponse
func (controller *Controller) DeleteRoot_step(c *gin.Context) {

	mutexRoot_step.Lock()
	defer mutexRoot_step.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRoot_step", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoRoot_step.GetDB()

	// Get model if exist
	var root_stepDB orm.Root_stepDB
	if err := db.First(&root_stepDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&root_stepDB)

	// get an instance (not staged) from DB instance, and call callback function
	root_stepDeleted := new(models.Root_step)
	root_stepDB.CopyBasicFieldsToRoot_step(root_stepDeleted)

	// get stage instance from DB instance, and call callback function
	root_stepStaged := backRepo.BackRepoRoot_step.Map_Root_stepDBID_Root_stepPtr[root_stepDB.ID]
	if root_stepStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), root_stepStaged, root_stepDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
