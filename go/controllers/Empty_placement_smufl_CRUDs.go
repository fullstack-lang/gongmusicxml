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
var __Empty_placement_smufl__dummysDeclaration__ models.Empty_placement_smufl
var __Empty_placement_smufl_time__dummyDeclaration time.Duration

var mutexEmpty_placement_smufl sync.Mutex

// An Empty_placement_smuflID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_placement_smufl updateEmpty_placement_smufl deleteEmpty_placement_smufl
type Empty_placement_smuflID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_placement_smuflInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_placement_smufl updateEmpty_placement_smufl
type Empty_placement_smuflInput struct {
	// The Empty_placement_smufl to submit or modify
	// in: body
	Empty_placement_smufl *orm.Empty_placement_smuflAPI
}

// GetEmpty_placement_smufls
//
// swagger:route GET /empty_placement_smufls empty_placement_smufls getEmpty_placement_smufls
//
// # Get all empty_placement_smufls
//
// Responses:
// default: genericError
//
//	200: empty_placement_smuflDBResponse
func (controller *Controller) GetEmpty_placement_smufls(c *gin.Context) {

	// source slice
	var empty_placement_smuflDBs []orm.Empty_placement_smuflDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_placement_smufls", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement_smufl.GetDB()

	query := db.Find(&empty_placement_smuflDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_placement_smuflAPIs := make([]orm.Empty_placement_smuflAPI, 0)

	// for each empty_placement_smufl, update fields from the database nullable fields
	for idx := range empty_placement_smuflDBs {
		empty_placement_smuflDB := &empty_placement_smuflDBs[idx]
		_ = empty_placement_smuflDB
		var empty_placement_smuflAPI orm.Empty_placement_smuflAPI

		// insertion point for updating fields
		empty_placement_smuflAPI.ID = empty_placement_smuflDB.ID
		empty_placement_smuflDB.CopyBasicFieldsToEmpty_placement_smufl_WOP(&empty_placement_smuflAPI.Empty_placement_smufl_WOP)
		empty_placement_smuflAPI.Empty_placement_smuflPointersEncoding = empty_placement_smuflDB.Empty_placement_smuflPointersEncoding
		empty_placement_smuflAPIs = append(empty_placement_smuflAPIs, empty_placement_smuflAPI)
	}

	c.JSON(http.StatusOK, empty_placement_smuflAPIs)
}

// PostEmpty_placement_smufl
//
// swagger:route POST /empty_placement_smufls empty_placement_smufls postEmpty_placement_smufl
//
// Creates a empty_placement_smufl
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_placement_smufl(c *gin.Context) {

	mutexEmpty_placement_smufl.Lock()
	defer mutexEmpty_placement_smufl.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_placement_smufls", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement_smufl.GetDB()

	// Validate input
	var input orm.Empty_placement_smuflAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_placement_smufl
	empty_placement_smuflDB := orm.Empty_placement_smuflDB{}
	empty_placement_smuflDB.Empty_placement_smuflPointersEncoding = input.Empty_placement_smuflPointersEncoding
	empty_placement_smuflDB.CopyBasicFieldsFromEmpty_placement_smufl_WOP(&input.Empty_placement_smufl_WOP)

	query := db.Create(&empty_placement_smuflDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_placement_smufl.CheckoutPhaseOneInstance(&empty_placement_smuflDB)
	empty_placement_smufl := backRepo.BackRepoEmpty_placement_smufl.Map_Empty_placement_smuflDBID_Empty_placement_smuflPtr[empty_placement_smuflDB.ID]

	if empty_placement_smufl != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_placement_smufl)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_placement_smuflDB)
}

// GetEmpty_placement_smufl
//
// swagger:route GET /empty_placement_smufls/{ID} empty_placement_smufls getEmpty_placement_smufl
//
// Gets the details for a empty_placement_smufl.
//
// Responses:
// default: genericError
//
//	200: empty_placement_smuflDBResponse
func (controller *Controller) GetEmpty_placement_smufl(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_placement_smufl", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement_smufl.GetDB()

	// Get empty_placement_smuflDB in DB
	var empty_placement_smuflDB orm.Empty_placement_smuflDB
	if err := db.First(&empty_placement_smuflDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_placement_smuflAPI orm.Empty_placement_smuflAPI
	empty_placement_smuflAPI.ID = empty_placement_smuflDB.ID
	empty_placement_smuflAPI.Empty_placement_smuflPointersEncoding = empty_placement_smuflDB.Empty_placement_smuflPointersEncoding
	empty_placement_smuflDB.CopyBasicFieldsToEmpty_placement_smufl_WOP(&empty_placement_smuflAPI.Empty_placement_smufl_WOP)

	c.JSON(http.StatusOK, empty_placement_smuflAPI)
}

// UpdateEmpty_placement_smufl
//
// swagger:route PATCH /empty_placement_smufls/{ID} empty_placement_smufls updateEmpty_placement_smufl
//
// # Update a empty_placement_smufl
//
// Responses:
// default: genericError
//
//	200: empty_placement_smuflDBResponse
func (controller *Controller) UpdateEmpty_placement_smufl(c *gin.Context) {

	mutexEmpty_placement_smufl.Lock()
	defer mutexEmpty_placement_smufl.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_placement_smufl", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement_smufl.GetDB()

	// Validate input
	var input orm.Empty_placement_smuflAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_placement_smuflDB orm.Empty_placement_smuflDB

	// fetch the empty_placement_smufl
	query := db.First(&empty_placement_smuflDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_placement_smuflDB.CopyBasicFieldsFromEmpty_placement_smufl_WOP(&input.Empty_placement_smufl_WOP)
	empty_placement_smuflDB.Empty_placement_smuflPointersEncoding = input.Empty_placement_smuflPointersEncoding

	query = db.Model(&empty_placement_smuflDB).Updates(empty_placement_smuflDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_placement_smuflNew := new(models.Empty_placement_smufl)
	empty_placement_smuflDB.CopyBasicFieldsToEmpty_placement_smufl(empty_placement_smuflNew)

	// redeem pointers
	empty_placement_smuflDB.DecodePointers(backRepo, empty_placement_smuflNew)

	// get stage instance from DB instance, and call callback function
	empty_placement_smuflOld := backRepo.BackRepoEmpty_placement_smufl.Map_Empty_placement_smuflDBID_Empty_placement_smuflPtr[empty_placement_smuflDB.ID]
	if empty_placement_smuflOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_placement_smuflOld, empty_placement_smuflNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_placement_smuflDB
	c.JSON(http.StatusOK, empty_placement_smuflDB)
}

// DeleteEmpty_placement_smufl
//
// swagger:route DELETE /empty_placement_smufls/{ID} empty_placement_smufls deleteEmpty_placement_smufl
//
// # Delete a empty_placement_smufl
//
// default: genericError
//
//	200: empty_placement_smuflDBResponse
func (controller *Controller) DeleteEmpty_placement_smufl(c *gin.Context) {

	mutexEmpty_placement_smufl.Lock()
	defer mutexEmpty_placement_smufl.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_placement_smufl", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_placement_smufl.GetDB()

	// Get model if exist
	var empty_placement_smuflDB orm.Empty_placement_smuflDB
	if err := db.First(&empty_placement_smuflDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_placement_smuflDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_placement_smuflDeleted := new(models.Empty_placement_smufl)
	empty_placement_smuflDB.CopyBasicFieldsToEmpty_placement_smufl(empty_placement_smuflDeleted)

	// get stage instance from DB instance, and call callback function
	empty_placement_smuflStaged := backRepo.BackRepoEmpty_placement_smufl.Map_Empty_placement_smuflDBID_Empty_placement_smuflPtr[empty_placement_smuflDB.ID]
	if empty_placement_smuflStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_placement_smuflStaged, empty_placement_smuflDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
