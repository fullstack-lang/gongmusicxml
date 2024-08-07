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
var __Staff_layout__dummysDeclaration__ models.Staff_layout
var __Staff_layout_time__dummyDeclaration time.Duration

var mutexStaff_layout sync.Mutex

// An Staff_layoutID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaff_layout updateStaff_layout deleteStaff_layout
type Staff_layoutID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Staff_layoutInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaff_layout updateStaff_layout
type Staff_layoutInput struct {
	// The Staff_layout to submit or modify
	// in: body
	Staff_layout *orm.Staff_layoutAPI
}

// GetStaff_layouts
//
// swagger:route GET /staff_layouts staff_layouts getStaff_layouts
//
// # Get all staff_layouts
//
// Responses:
// default: genericError
//
//	200: staff_layoutDBResponse
func (controller *Controller) GetStaff_layouts(c *gin.Context) {

	// source slice
	var staff_layoutDBs []orm.Staff_layoutDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_layout.GetDB()

	query := db.Find(&staff_layoutDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staff_layoutAPIs := make([]orm.Staff_layoutAPI, 0)

	// for each staff_layout, update fields from the database nullable fields
	for idx := range staff_layoutDBs {
		staff_layoutDB := &staff_layoutDBs[idx]
		_ = staff_layoutDB
		var staff_layoutAPI orm.Staff_layoutAPI

		// insertion point for updating fields
		staff_layoutAPI.ID = staff_layoutDB.ID
		staff_layoutDB.CopyBasicFieldsToStaff_layout_WOP(&staff_layoutAPI.Staff_layout_WOP)
		staff_layoutAPI.Staff_layoutPointersEncoding = staff_layoutDB.Staff_layoutPointersEncoding
		staff_layoutAPIs = append(staff_layoutAPIs, staff_layoutAPI)
	}

	c.JSON(http.StatusOK, staff_layoutAPIs)
}

// PostStaff_layout
//
// swagger:route POST /staff_layouts staff_layouts postStaff_layout
//
// Creates a staff_layout
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaff_layout(c *gin.Context) {

	mutexStaff_layout.Lock()
	defer mutexStaff_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaff_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_layout.GetDB()

	// Validate input
	var input orm.Staff_layoutAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staff_layout
	staff_layoutDB := orm.Staff_layoutDB{}
	staff_layoutDB.Staff_layoutPointersEncoding = input.Staff_layoutPointersEncoding
	staff_layoutDB.CopyBasicFieldsFromStaff_layout_WOP(&input.Staff_layout_WOP)

	query := db.Create(&staff_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaff_layout.CheckoutPhaseOneInstance(&staff_layoutDB)
	staff_layout := backRepo.BackRepoStaff_layout.Map_Staff_layoutDBID_Staff_layoutPtr[staff_layoutDB.ID]

	if staff_layout != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staff_layout)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staff_layoutDB)
}

// GetStaff_layout
//
// swagger:route GET /staff_layouts/{ID} staff_layouts getStaff_layout
//
// Gets the details for a staff_layout.
//
// Responses:
// default: genericError
//
//	200: staff_layoutDBResponse
func (controller *Controller) GetStaff_layout(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_layout.GetDB()

	// Get staff_layoutDB in DB
	var staff_layoutDB orm.Staff_layoutDB
	if err := db.First(&staff_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staff_layoutAPI orm.Staff_layoutAPI
	staff_layoutAPI.ID = staff_layoutDB.ID
	staff_layoutAPI.Staff_layoutPointersEncoding = staff_layoutDB.Staff_layoutPointersEncoding
	staff_layoutDB.CopyBasicFieldsToStaff_layout_WOP(&staff_layoutAPI.Staff_layout_WOP)

	c.JSON(http.StatusOK, staff_layoutAPI)
}

// UpdateStaff_layout
//
// swagger:route PATCH /staff_layouts/{ID} staff_layouts updateStaff_layout
//
// # Update a staff_layout
//
// Responses:
// default: genericError
//
//	200: staff_layoutDBResponse
func (controller *Controller) UpdateStaff_layout(c *gin.Context) {

	mutexStaff_layout.Lock()
	defer mutexStaff_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaff_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_layout.GetDB()

	// Validate input
	var input orm.Staff_layoutAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staff_layoutDB orm.Staff_layoutDB

	// fetch the staff_layout
	query := db.First(&staff_layoutDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staff_layoutDB.CopyBasicFieldsFromStaff_layout_WOP(&input.Staff_layout_WOP)
	staff_layoutDB.Staff_layoutPointersEncoding = input.Staff_layoutPointersEncoding

	query = db.Model(&staff_layoutDB).Updates(staff_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staff_layoutNew := new(models.Staff_layout)
	staff_layoutDB.CopyBasicFieldsToStaff_layout(staff_layoutNew)

	// redeem pointers
	staff_layoutDB.DecodePointers(backRepo, staff_layoutNew)

	// get stage instance from DB instance, and call callback function
	staff_layoutOld := backRepo.BackRepoStaff_layout.Map_Staff_layoutDBID_Staff_layoutPtr[staff_layoutDB.ID]
	if staff_layoutOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staff_layoutOld, staff_layoutNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staff_layoutDB
	c.JSON(http.StatusOK, staff_layoutDB)
}

// DeleteStaff_layout
//
// swagger:route DELETE /staff_layouts/{ID} staff_layouts deleteStaff_layout
//
// # Delete a staff_layout
//
// default: genericError
//
//	200: staff_layoutDBResponse
func (controller *Controller) DeleteStaff_layout(c *gin.Context) {

	mutexStaff_layout.Lock()
	defer mutexStaff_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaff_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_layout.GetDB()

	// Get model if exist
	var staff_layoutDB orm.Staff_layoutDB
	if err := db.First(&staff_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&staff_layoutDB)

	// get an instance (not staged) from DB instance, and call callback function
	staff_layoutDeleted := new(models.Staff_layout)
	staff_layoutDB.CopyBasicFieldsToStaff_layout(staff_layoutDeleted)

	// get stage instance from DB instance, and call callback function
	staff_layoutStaged := backRepo.BackRepoStaff_layout.Map_Staff_layoutDBID_Staff_layoutPtr[staff_layoutDB.ID]
	if staff_layoutStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staff_layoutStaged, staff_layoutDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
