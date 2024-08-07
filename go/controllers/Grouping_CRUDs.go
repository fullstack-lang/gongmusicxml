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
var __Grouping__dummysDeclaration__ models.Grouping
var __Grouping_time__dummyDeclaration time.Duration

var mutexGrouping sync.Mutex

// An GroupingID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGrouping updateGrouping deleteGrouping
type GroupingID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GroupingInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGrouping updateGrouping
type GroupingInput struct {
	// The Grouping to submit or modify
	// in: body
	Grouping *orm.GroupingAPI
}

// GetGroupings
//
// swagger:route GET /groupings groupings getGroupings
//
// # Get all groupings
//
// Responses:
// default: genericError
//
//	200: groupingDBResponse
func (controller *Controller) GetGroupings(c *gin.Context) {

	// source slice
	var groupingDBs []orm.GroupingDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroupings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrouping.GetDB()

	query := db.Find(&groupingDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	groupingAPIs := make([]orm.GroupingAPI, 0)

	// for each grouping, update fields from the database nullable fields
	for idx := range groupingDBs {
		groupingDB := &groupingDBs[idx]
		_ = groupingDB
		var groupingAPI orm.GroupingAPI

		// insertion point for updating fields
		groupingAPI.ID = groupingDB.ID
		groupingDB.CopyBasicFieldsToGrouping_WOP(&groupingAPI.Grouping_WOP)
		groupingAPI.GroupingPointersEncoding = groupingDB.GroupingPointersEncoding
		groupingAPIs = append(groupingAPIs, groupingAPI)
	}

	c.JSON(http.StatusOK, groupingAPIs)
}

// PostGrouping
//
// swagger:route POST /groupings groupings postGrouping
//
// Creates a grouping
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGrouping(c *gin.Context) {

	mutexGrouping.Lock()
	defer mutexGrouping.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroupings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrouping.GetDB()

	// Validate input
	var input orm.GroupingAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create grouping
	groupingDB := orm.GroupingDB{}
	groupingDB.GroupingPointersEncoding = input.GroupingPointersEncoding
	groupingDB.CopyBasicFieldsFromGrouping_WOP(&input.Grouping_WOP)

	query := db.Create(&groupingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGrouping.CheckoutPhaseOneInstance(&groupingDB)
	grouping := backRepo.BackRepoGrouping.Map_GroupingDBID_GroupingPtr[groupingDB.ID]

	if grouping != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), grouping)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, groupingDB)
}

// GetGrouping
//
// swagger:route GET /groupings/{ID} groupings getGrouping
//
// Gets the details for a grouping.
//
// Responses:
// default: genericError
//
//	200: groupingDBResponse
func (controller *Controller) GetGrouping(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGrouping", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrouping.GetDB()

	// Get groupingDB in DB
	var groupingDB orm.GroupingDB
	if err := db.First(&groupingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var groupingAPI orm.GroupingAPI
	groupingAPI.ID = groupingDB.ID
	groupingAPI.GroupingPointersEncoding = groupingDB.GroupingPointersEncoding
	groupingDB.CopyBasicFieldsToGrouping_WOP(&groupingAPI.Grouping_WOP)

	c.JSON(http.StatusOK, groupingAPI)
}

// UpdateGrouping
//
// swagger:route PATCH /groupings/{ID} groupings updateGrouping
//
// # Update a grouping
//
// Responses:
// default: genericError
//
//	200: groupingDBResponse
func (controller *Controller) UpdateGrouping(c *gin.Context) {

	mutexGrouping.Lock()
	defer mutexGrouping.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGrouping", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrouping.GetDB()

	// Validate input
	var input orm.GroupingAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var groupingDB orm.GroupingDB

	// fetch the grouping
	query := db.First(&groupingDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	groupingDB.CopyBasicFieldsFromGrouping_WOP(&input.Grouping_WOP)
	groupingDB.GroupingPointersEncoding = input.GroupingPointersEncoding

	query = db.Model(&groupingDB).Updates(groupingDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	groupingNew := new(models.Grouping)
	groupingDB.CopyBasicFieldsToGrouping(groupingNew)

	// redeem pointers
	groupingDB.DecodePointers(backRepo, groupingNew)

	// get stage instance from DB instance, and call callback function
	groupingOld := backRepo.BackRepoGrouping.Map_GroupingDBID_GroupingPtr[groupingDB.ID]
	if groupingOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), groupingOld, groupingNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the groupingDB
	c.JSON(http.StatusOK, groupingDB)
}

// DeleteGrouping
//
// swagger:route DELETE /groupings/{ID} groupings deleteGrouping
//
// # Delete a grouping
//
// default: genericError
//
//	200: groupingDBResponse
func (controller *Controller) DeleteGrouping(c *gin.Context) {

	mutexGrouping.Lock()
	defer mutexGrouping.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGrouping", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGrouping.GetDB()

	// Get model if exist
	var groupingDB orm.GroupingDB
	if err := db.First(&groupingDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&groupingDB)

	// get an instance (not staged) from DB instance, and call callback function
	groupingDeleted := new(models.Grouping)
	groupingDB.CopyBasicFieldsToGrouping(groupingDeleted)

	// get stage instance from DB instance, and call callback function
	groupingStaged := backRepo.BackRepoGrouping.Map_GroupingDBID_GroupingPtr[groupingDB.ID]
	if groupingStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), groupingStaged, groupingDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
