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
var __Part_group__dummysDeclaration__ models.Part_group
var __Part_group_time__dummyDeclaration time.Duration

var mutexPart_group sync.Mutex

// An Part_groupID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPart_group updatePart_group deletePart_group
type Part_groupID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Part_groupInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPart_group updatePart_group
type Part_groupInput struct {
	// The Part_group to submit or modify
	// in: body
	Part_group *orm.Part_groupAPI
}

// GetPart_groups
//
// swagger:route GET /part_groups part_groups getPart_groups
//
// # Get all part_groups
//
// Responses:
// default: genericError
//
//	200: part_groupDBResponse
func (controller *Controller) GetPart_groups(c *gin.Context) {

	// source slice
	var part_groupDBs []orm.Part_groupDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_groups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_group.GetDB()

	query := db.Find(&part_groupDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	part_groupAPIs := make([]orm.Part_groupAPI, 0)

	// for each part_group, update fields from the database nullable fields
	for idx := range part_groupDBs {
		part_groupDB := &part_groupDBs[idx]
		_ = part_groupDB
		var part_groupAPI orm.Part_groupAPI

		// insertion point for updating fields
		part_groupAPI.ID = part_groupDB.ID
		part_groupDB.CopyBasicFieldsToPart_group_WOP(&part_groupAPI.Part_group_WOP)
		part_groupAPI.Part_groupPointersEncoding = part_groupDB.Part_groupPointersEncoding
		part_groupAPIs = append(part_groupAPIs, part_groupAPI)
	}

	c.JSON(http.StatusOK, part_groupAPIs)
}

// PostPart_group
//
// swagger:route POST /part_groups part_groups postPart_group
//
// Creates a part_group
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPart_group(c *gin.Context) {

	mutexPart_group.Lock()
	defer mutexPart_group.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPart_groups", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_group.GetDB()

	// Validate input
	var input orm.Part_groupAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create part_group
	part_groupDB := orm.Part_groupDB{}
	part_groupDB.Part_groupPointersEncoding = input.Part_groupPointersEncoding
	part_groupDB.CopyBasicFieldsFromPart_group_WOP(&input.Part_group_WOP)

	query := db.Create(&part_groupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPart_group.CheckoutPhaseOneInstance(&part_groupDB)
	part_group := backRepo.BackRepoPart_group.Map_Part_groupDBID_Part_groupPtr[part_groupDB.ID]

	if part_group != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), part_group)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, part_groupDB)
}

// GetPart_group
//
// swagger:route GET /part_groups/{ID} part_groups getPart_group
//
// Gets the details for a part_group.
//
// Responses:
// default: genericError
//
//	200: part_groupDBResponse
func (controller *Controller) GetPart_group(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_group", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_group.GetDB()

	// Get part_groupDB in DB
	var part_groupDB orm.Part_groupDB
	if err := db.First(&part_groupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var part_groupAPI orm.Part_groupAPI
	part_groupAPI.ID = part_groupDB.ID
	part_groupAPI.Part_groupPointersEncoding = part_groupDB.Part_groupPointersEncoding
	part_groupDB.CopyBasicFieldsToPart_group_WOP(&part_groupAPI.Part_group_WOP)

	c.JSON(http.StatusOK, part_groupAPI)
}

// UpdatePart_group
//
// swagger:route PATCH /part_groups/{ID} part_groups updatePart_group
//
// # Update a part_group
//
// Responses:
// default: genericError
//
//	200: part_groupDBResponse
func (controller *Controller) UpdatePart_group(c *gin.Context) {

	mutexPart_group.Lock()
	defer mutexPart_group.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePart_group", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_group.GetDB()

	// Validate input
	var input orm.Part_groupAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var part_groupDB orm.Part_groupDB

	// fetch the part_group
	query := db.First(&part_groupDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	part_groupDB.CopyBasicFieldsFromPart_group_WOP(&input.Part_group_WOP)
	part_groupDB.Part_groupPointersEncoding = input.Part_groupPointersEncoding

	query = db.Model(&part_groupDB).Updates(part_groupDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	part_groupNew := new(models.Part_group)
	part_groupDB.CopyBasicFieldsToPart_group(part_groupNew)

	// redeem pointers
	part_groupDB.DecodePointers(backRepo, part_groupNew)

	// get stage instance from DB instance, and call callback function
	part_groupOld := backRepo.BackRepoPart_group.Map_Part_groupDBID_Part_groupPtr[part_groupDB.ID]
	if part_groupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), part_groupOld, part_groupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the part_groupDB
	c.JSON(http.StatusOK, part_groupDB)
}

// DeletePart_group
//
// swagger:route DELETE /part_groups/{ID} part_groups deletePart_group
//
// # Delete a part_group
//
// default: genericError
//
//	200: part_groupDBResponse
func (controller *Controller) DeletePart_group(c *gin.Context) {

	mutexPart_group.Lock()
	defer mutexPart_group.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePart_group", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_group.GetDB()

	// Get model if exist
	var part_groupDB orm.Part_groupDB
	if err := db.First(&part_groupDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&part_groupDB)

	// get an instance (not staged) from DB instance, and call callback function
	part_groupDeleted := new(models.Part_group)
	part_groupDB.CopyBasicFieldsToPart_group(part_groupDeleted)

	// get stage instance from DB instance, and call callback function
	part_groupStaged := backRepo.BackRepoPart_group.Map_Part_groupDBID_Part_groupPtr[part_groupDB.ID]
	if part_groupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), part_groupStaged, part_groupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
