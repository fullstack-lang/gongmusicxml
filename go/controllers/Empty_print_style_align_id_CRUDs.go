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
var __Empty_print_style_align_id__dummysDeclaration__ models.Empty_print_style_align_id
var __Empty_print_style_align_id_time__dummyDeclaration time.Duration

var mutexEmpty_print_style_align_id sync.Mutex

// An Empty_print_style_align_idID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEmpty_print_style_align_id updateEmpty_print_style_align_id deleteEmpty_print_style_align_id
type Empty_print_style_align_idID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Empty_print_style_align_idInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEmpty_print_style_align_id updateEmpty_print_style_align_id
type Empty_print_style_align_idInput struct {
	// The Empty_print_style_align_id to submit or modify
	// in: body
	Empty_print_style_align_id *orm.Empty_print_style_align_idAPI
}

// GetEmpty_print_style_align_ids
//
// swagger:route GET /empty_print_style_align_ids empty_print_style_align_ids getEmpty_print_style_align_ids
//
// # Get all empty_print_style_align_ids
//
// Responses:
// default: genericError
//
//	200: empty_print_style_align_idDBResponse
func (controller *Controller) GetEmpty_print_style_align_ids(c *gin.Context) {

	// source slice
	var empty_print_style_align_idDBs []orm.Empty_print_style_align_idDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_print_style_align_ids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align_id.GetDB()

	query := db.Find(&empty_print_style_align_idDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	empty_print_style_align_idAPIs := make([]orm.Empty_print_style_align_idAPI, 0)

	// for each empty_print_style_align_id, update fields from the database nullable fields
	for idx := range empty_print_style_align_idDBs {
		empty_print_style_align_idDB := &empty_print_style_align_idDBs[idx]
		_ = empty_print_style_align_idDB
		var empty_print_style_align_idAPI orm.Empty_print_style_align_idAPI

		// insertion point for updating fields
		empty_print_style_align_idAPI.ID = empty_print_style_align_idDB.ID
		empty_print_style_align_idDB.CopyBasicFieldsToEmpty_print_style_align_id_WOP(&empty_print_style_align_idAPI.Empty_print_style_align_id_WOP)
		empty_print_style_align_idAPI.Empty_print_style_align_idPointersEncoding = empty_print_style_align_idDB.Empty_print_style_align_idPointersEncoding
		empty_print_style_align_idAPIs = append(empty_print_style_align_idAPIs, empty_print_style_align_idAPI)
	}

	c.JSON(http.StatusOK, empty_print_style_align_idAPIs)
}

// PostEmpty_print_style_align_id
//
// swagger:route POST /empty_print_style_align_ids empty_print_style_align_ids postEmpty_print_style_align_id
//
// Creates a empty_print_style_align_id
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEmpty_print_style_align_id(c *gin.Context) {

	mutexEmpty_print_style_align_id.Lock()
	defer mutexEmpty_print_style_align_id.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEmpty_print_style_align_ids", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align_id.GetDB()

	// Validate input
	var input orm.Empty_print_style_align_idAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create empty_print_style_align_id
	empty_print_style_align_idDB := orm.Empty_print_style_align_idDB{}
	empty_print_style_align_idDB.Empty_print_style_align_idPointersEncoding = input.Empty_print_style_align_idPointersEncoding
	empty_print_style_align_idDB.CopyBasicFieldsFromEmpty_print_style_align_id_WOP(&input.Empty_print_style_align_id_WOP)

	query := db.Create(&empty_print_style_align_idDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEmpty_print_style_align_id.CheckoutPhaseOneInstance(&empty_print_style_align_idDB)
	empty_print_style_align_id := backRepo.BackRepoEmpty_print_style_align_id.Map_Empty_print_style_align_idDBID_Empty_print_style_align_idPtr[empty_print_style_align_idDB.ID]

	if empty_print_style_align_id != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), empty_print_style_align_id)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, empty_print_style_align_idDB)
}

// GetEmpty_print_style_align_id
//
// swagger:route GET /empty_print_style_align_ids/{ID} empty_print_style_align_ids getEmpty_print_style_align_id
//
// Gets the details for a empty_print_style_align_id.
//
// Responses:
// default: genericError
//
//	200: empty_print_style_align_idDBResponse
func (controller *Controller) GetEmpty_print_style_align_id(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEmpty_print_style_align_id", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align_id.GetDB()

	// Get empty_print_style_align_idDB in DB
	var empty_print_style_align_idDB orm.Empty_print_style_align_idDB
	if err := db.First(&empty_print_style_align_idDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var empty_print_style_align_idAPI orm.Empty_print_style_align_idAPI
	empty_print_style_align_idAPI.ID = empty_print_style_align_idDB.ID
	empty_print_style_align_idAPI.Empty_print_style_align_idPointersEncoding = empty_print_style_align_idDB.Empty_print_style_align_idPointersEncoding
	empty_print_style_align_idDB.CopyBasicFieldsToEmpty_print_style_align_id_WOP(&empty_print_style_align_idAPI.Empty_print_style_align_id_WOP)

	c.JSON(http.StatusOK, empty_print_style_align_idAPI)
}

// UpdateEmpty_print_style_align_id
//
// swagger:route PATCH /empty_print_style_align_ids/{ID} empty_print_style_align_ids updateEmpty_print_style_align_id
//
// # Update a empty_print_style_align_id
//
// Responses:
// default: genericError
//
//	200: empty_print_style_align_idDBResponse
func (controller *Controller) UpdateEmpty_print_style_align_id(c *gin.Context) {

	mutexEmpty_print_style_align_id.Lock()
	defer mutexEmpty_print_style_align_id.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEmpty_print_style_align_id", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align_id.GetDB()

	// Validate input
	var input orm.Empty_print_style_align_idAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var empty_print_style_align_idDB orm.Empty_print_style_align_idDB

	// fetch the empty_print_style_align_id
	query := db.First(&empty_print_style_align_idDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	empty_print_style_align_idDB.CopyBasicFieldsFromEmpty_print_style_align_id_WOP(&input.Empty_print_style_align_id_WOP)
	empty_print_style_align_idDB.Empty_print_style_align_idPointersEncoding = input.Empty_print_style_align_idPointersEncoding

	query = db.Model(&empty_print_style_align_idDB).Updates(empty_print_style_align_idDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	empty_print_style_align_idNew := new(models.Empty_print_style_align_id)
	empty_print_style_align_idDB.CopyBasicFieldsToEmpty_print_style_align_id(empty_print_style_align_idNew)

	// redeem pointers
	empty_print_style_align_idDB.DecodePointers(backRepo, empty_print_style_align_idNew)

	// get stage instance from DB instance, and call callback function
	empty_print_style_align_idOld := backRepo.BackRepoEmpty_print_style_align_id.Map_Empty_print_style_align_idDBID_Empty_print_style_align_idPtr[empty_print_style_align_idDB.ID]
	if empty_print_style_align_idOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), empty_print_style_align_idOld, empty_print_style_align_idNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the empty_print_style_align_idDB
	c.JSON(http.StatusOK, empty_print_style_align_idDB)
}

// DeleteEmpty_print_style_align_id
//
// swagger:route DELETE /empty_print_style_align_ids/{ID} empty_print_style_align_ids deleteEmpty_print_style_align_id
//
// # Delete a empty_print_style_align_id
//
// default: genericError
//
//	200: empty_print_style_align_idDBResponse
func (controller *Controller) DeleteEmpty_print_style_align_id(c *gin.Context) {

	mutexEmpty_print_style_align_id.Lock()
	defer mutexEmpty_print_style_align_id.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEmpty_print_style_align_id", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEmpty_print_style_align_id.GetDB()

	// Get model if exist
	var empty_print_style_align_idDB orm.Empty_print_style_align_idDB
	if err := db.First(&empty_print_style_align_idDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&empty_print_style_align_idDB)

	// get an instance (not staged) from DB instance, and call callback function
	empty_print_style_align_idDeleted := new(models.Empty_print_style_align_id)
	empty_print_style_align_idDB.CopyBasicFieldsToEmpty_print_style_align_id(empty_print_style_align_idDeleted)

	// get stage instance from DB instance, and call callback function
	empty_print_style_align_idStaged := backRepo.BackRepoEmpty_print_style_align_id.Map_Empty_print_style_align_idDBID_Empty_print_style_align_idPtr[empty_print_style_align_idDB.ID]
	if empty_print_style_align_idStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), empty_print_style_align_idStaged, empty_print_style_align_idDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
