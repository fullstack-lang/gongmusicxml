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
var __Group_barline__dummysDeclaration__ models.Group_barline
var __Group_barline_time__dummyDeclaration time.Duration

var mutexGroup_barline sync.Mutex

// An Group_barlineID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGroup_barline updateGroup_barline deleteGroup_barline
type Group_barlineID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Group_barlineInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGroup_barline updateGroup_barline
type Group_barlineInput struct {
	// The Group_barline to submit or modify
	// in: body
	Group_barline *orm.Group_barlineAPI
}

// GetGroup_barlines
//
// swagger:route GET /group_barlines group_barlines getGroup_barlines
//
// # Get all group_barlines
//
// Responses:
// default: genericError
//
//	200: group_barlineDBResponse
func (controller *Controller) GetGroup_barlines(c *gin.Context) {

	// source slice
	var group_barlineDBs []orm.Group_barlineDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroup_barlines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_barline.GetDB()

	query := db.Find(&group_barlineDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	group_barlineAPIs := make([]orm.Group_barlineAPI, 0)

	// for each group_barline, update fields from the database nullable fields
	for idx := range group_barlineDBs {
		group_barlineDB := &group_barlineDBs[idx]
		_ = group_barlineDB
		var group_barlineAPI orm.Group_barlineAPI

		// insertion point for updating fields
		group_barlineAPI.ID = group_barlineDB.ID
		group_barlineDB.CopyBasicFieldsToGroup_barline_WOP(&group_barlineAPI.Group_barline_WOP)
		group_barlineAPI.Group_barlinePointersEncoding = group_barlineDB.Group_barlinePointersEncoding
		group_barlineAPIs = append(group_barlineAPIs, group_barlineAPI)
	}

	c.JSON(http.StatusOK, group_barlineAPIs)
}

// PostGroup_barline
//
// swagger:route POST /group_barlines group_barlines postGroup_barline
//
// Creates a group_barline
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGroup_barline(c *gin.Context) {

	mutexGroup_barline.Lock()
	defer mutexGroup_barline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGroup_barlines", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_barline.GetDB()

	// Validate input
	var input orm.Group_barlineAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create group_barline
	group_barlineDB := orm.Group_barlineDB{}
	group_barlineDB.Group_barlinePointersEncoding = input.Group_barlinePointersEncoding
	group_barlineDB.CopyBasicFieldsFromGroup_barline_WOP(&input.Group_barline_WOP)

	query := db.Create(&group_barlineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGroup_barline.CheckoutPhaseOneInstance(&group_barlineDB)
	group_barline := backRepo.BackRepoGroup_barline.Map_Group_barlineDBID_Group_barlinePtr[group_barlineDB.ID]

	if group_barline != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), group_barline)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, group_barlineDB)
}

// GetGroup_barline
//
// swagger:route GET /group_barlines/{ID} group_barlines getGroup_barline
//
// Gets the details for a group_barline.
//
// Responses:
// default: genericError
//
//	200: group_barlineDBResponse
func (controller *Controller) GetGroup_barline(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGroup_barline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_barline.GetDB()

	// Get group_barlineDB in DB
	var group_barlineDB orm.Group_barlineDB
	if err := db.First(&group_barlineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var group_barlineAPI orm.Group_barlineAPI
	group_barlineAPI.ID = group_barlineDB.ID
	group_barlineAPI.Group_barlinePointersEncoding = group_barlineDB.Group_barlinePointersEncoding
	group_barlineDB.CopyBasicFieldsToGroup_barline_WOP(&group_barlineAPI.Group_barline_WOP)

	c.JSON(http.StatusOK, group_barlineAPI)
}

// UpdateGroup_barline
//
// swagger:route PATCH /group_barlines/{ID} group_barlines updateGroup_barline
//
// # Update a group_barline
//
// Responses:
// default: genericError
//
//	200: group_barlineDBResponse
func (controller *Controller) UpdateGroup_barline(c *gin.Context) {

	mutexGroup_barline.Lock()
	defer mutexGroup_barline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGroup_barline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_barline.GetDB()

	// Validate input
	var input orm.Group_barlineAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var group_barlineDB orm.Group_barlineDB

	// fetch the group_barline
	query := db.First(&group_barlineDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	group_barlineDB.CopyBasicFieldsFromGroup_barline_WOP(&input.Group_barline_WOP)
	group_barlineDB.Group_barlinePointersEncoding = input.Group_barlinePointersEncoding

	query = db.Model(&group_barlineDB).Updates(group_barlineDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	group_barlineNew := new(models.Group_barline)
	group_barlineDB.CopyBasicFieldsToGroup_barline(group_barlineNew)

	// redeem pointers
	group_barlineDB.DecodePointers(backRepo, group_barlineNew)

	// get stage instance from DB instance, and call callback function
	group_barlineOld := backRepo.BackRepoGroup_barline.Map_Group_barlineDBID_Group_barlinePtr[group_barlineDB.ID]
	if group_barlineOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), group_barlineOld, group_barlineNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the group_barlineDB
	c.JSON(http.StatusOK, group_barlineDB)
}

// DeleteGroup_barline
//
// swagger:route DELETE /group_barlines/{ID} group_barlines deleteGroup_barline
//
// # Delete a group_barline
//
// default: genericError
//
//	200: group_barlineDBResponse
func (controller *Controller) DeleteGroup_barline(c *gin.Context) {

	mutexGroup_barline.Lock()
	defer mutexGroup_barline.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGroup_barline", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoGroup_barline.GetDB()

	// Get model if exist
	var group_barlineDB orm.Group_barlineDB
	if err := db.First(&group_barlineDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&group_barlineDB)

	// get an instance (not staged) from DB instance, and call callback function
	group_barlineDeleted := new(models.Group_barline)
	group_barlineDB.CopyBasicFieldsToGroup_barline(group_barlineDeleted)

	// get stage instance from DB instance, and call callback function
	group_barlineStaged := backRepo.BackRepoGroup_barline.Map_Group_barlineDBID_Group_barlinePtr[group_barlineDB.ID]
	if group_barlineStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), group_barlineStaged, group_barlineDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
