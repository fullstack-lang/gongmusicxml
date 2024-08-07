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
var __Part_list__dummysDeclaration__ models.Part_list
var __Part_list_time__dummyDeclaration time.Duration

var mutexPart_list sync.Mutex

// An Part_listID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPart_list updatePart_list deletePart_list
type Part_listID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Part_listInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPart_list updatePart_list
type Part_listInput struct {
	// The Part_list to submit or modify
	// in: body
	Part_list *orm.Part_listAPI
}

// GetPart_lists
//
// swagger:route GET /part_lists part_lists getPart_lists
//
// # Get all part_lists
//
// Responses:
// default: genericError
//
//	200: part_listDBResponse
func (controller *Controller) GetPart_lists(c *gin.Context) {

	// source slice
	var part_listDBs []orm.Part_listDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_lists", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_list.GetDB()

	query := db.Find(&part_listDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	part_listAPIs := make([]orm.Part_listAPI, 0)

	// for each part_list, update fields from the database nullable fields
	for idx := range part_listDBs {
		part_listDB := &part_listDBs[idx]
		_ = part_listDB
		var part_listAPI orm.Part_listAPI

		// insertion point for updating fields
		part_listAPI.ID = part_listDB.ID
		part_listDB.CopyBasicFieldsToPart_list_WOP(&part_listAPI.Part_list_WOP)
		part_listAPI.Part_listPointersEncoding = part_listDB.Part_listPointersEncoding
		part_listAPIs = append(part_listAPIs, part_listAPI)
	}

	c.JSON(http.StatusOK, part_listAPIs)
}

// PostPart_list
//
// swagger:route POST /part_lists part_lists postPart_list
//
// Creates a part_list
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPart_list(c *gin.Context) {

	mutexPart_list.Lock()
	defer mutexPart_list.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPart_lists", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_list.GetDB()

	// Validate input
	var input orm.Part_listAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create part_list
	part_listDB := orm.Part_listDB{}
	part_listDB.Part_listPointersEncoding = input.Part_listPointersEncoding
	part_listDB.CopyBasicFieldsFromPart_list_WOP(&input.Part_list_WOP)

	query := db.Create(&part_listDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPart_list.CheckoutPhaseOneInstance(&part_listDB)
	part_list := backRepo.BackRepoPart_list.Map_Part_listDBID_Part_listPtr[part_listDB.ID]

	if part_list != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), part_list)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, part_listDB)
}

// GetPart_list
//
// swagger:route GET /part_lists/{ID} part_lists getPart_list
//
// Gets the details for a part_list.
//
// Responses:
// default: genericError
//
//	200: part_listDBResponse
func (controller *Controller) GetPart_list(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_list", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_list.GetDB()

	// Get part_listDB in DB
	var part_listDB orm.Part_listDB
	if err := db.First(&part_listDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var part_listAPI orm.Part_listAPI
	part_listAPI.ID = part_listDB.ID
	part_listAPI.Part_listPointersEncoding = part_listDB.Part_listPointersEncoding
	part_listDB.CopyBasicFieldsToPart_list_WOP(&part_listAPI.Part_list_WOP)

	c.JSON(http.StatusOK, part_listAPI)
}

// UpdatePart_list
//
// swagger:route PATCH /part_lists/{ID} part_lists updatePart_list
//
// # Update a part_list
//
// Responses:
// default: genericError
//
//	200: part_listDBResponse
func (controller *Controller) UpdatePart_list(c *gin.Context) {

	mutexPart_list.Lock()
	defer mutexPart_list.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePart_list", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_list.GetDB()

	// Validate input
	var input orm.Part_listAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var part_listDB orm.Part_listDB

	// fetch the part_list
	query := db.First(&part_listDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	part_listDB.CopyBasicFieldsFromPart_list_WOP(&input.Part_list_WOP)
	part_listDB.Part_listPointersEncoding = input.Part_listPointersEncoding

	query = db.Model(&part_listDB).Updates(part_listDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	part_listNew := new(models.Part_list)
	part_listDB.CopyBasicFieldsToPart_list(part_listNew)

	// redeem pointers
	part_listDB.DecodePointers(backRepo, part_listNew)

	// get stage instance from DB instance, and call callback function
	part_listOld := backRepo.BackRepoPart_list.Map_Part_listDBID_Part_listPtr[part_listDB.ID]
	if part_listOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), part_listOld, part_listNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the part_listDB
	c.JSON(http.StatusOK, part_listDB)
}

// DeletePart_list
//
// swagger:route DELETE /part_lists/{ID} part_lists deletePart_list
//
// # Delete a part_list
//
// default: genericError
//
//	200: part_listDBResponse
func (controller *Controller) DeletePart_list(c *gin.Context) {

	mutexPart_list.Lock()
	defer mutexPart_list.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePart_list", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_list.GetDB()

	// Get model if exist
	var part_listDB orm.Part_listDB
	if err := db.First(&part_listDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&part_listDB)

	// get an instance (not staged) from DB instance, and call callback function
	part_listDeleted := new(models.Part_list)
	part_listDB.CopyBasicFieldsToPart_list(part_listDeleted)

	// get stage instance from DB instance, and call callback function
	part_listStaged := backRepo.BackRepoPart_list.Map_Part_listDBID_Part_listPtr[part_listDB.ID]
	if part_listStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), part_listStaged, part_listDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
