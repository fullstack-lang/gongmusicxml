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
var __Staff_details__dummysDeclaration__ models.Staff_details
var __Staff_details_time__dummyDeclaration time.Duration

var mutexStaff_details sync.Mutex

// An Staff_detailsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaff_details updateStaff_details deleteStaff_details
type Staff_detailsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Staff_detailsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaff_details updateStaff_details
type Staff_detailsInput struct {
	// The Staff_details to submit or modify
	// in: body
	Staff_details *orm.Staff_detailsAPI
}

// GetStaff_detailss
//
// swagger:route GET /staff_detailss staff_detailss getStaff_detailss
//
// # Get all staff_detailss
//
// Responses:
// default: genericError
//
//	200: staff_detailsDBResponse
func (controller *Controller) GetStaff_detailss(c *gin.Context) {

	// source slice
	var staff_detailsDBs []orm.Staff_detailsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_detailss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_details.GetDB()

	query := db.Find(&staff_detailsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staff_detailsAPIs := make([]orm.Staff_detailsAPI, 0)

	// for each staff_details, update fields from the database nullable fields
	for idx := range staff_detailsDBs {
		staff_detailsDB := &staff_detailsDBs[idx]
		_ = staff_detailsDB
		var staff_detailsAPI orm.Staff_detailsAPI

		// insertion point for updating fields
		staff_detailsAPI.ID = staff_detailsDB.ID
		staff_detailsDB.CopyBasicFieldsToStaff_details_WOP(&staff_detailsAPI.Staff_details_WOP)
		staff_detailsAPI.Staff_detailsPointersEncoding = staff_detailsDB.Staff_detailsPointersEncoding
		staff_detailsAPIs = append(staff_detailsAPIs, staff_detailsAPI)
	}

	c.JSON(http.StatusOK, staff_detailsAPIs)
}

// PostStaff_details
//
// swagger:route POST /staff_detailss staff_detailss postStaff_details
//
// Creates a staff_details
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaff_details(c *gin.Context) {

	mutexStaff_details.Lock()
	defer mutexStaff_details.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaff_detailss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_details.GetDB()

	// Validate input
	var input orm.Staff_detailsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staff_details
	staff_detailsDB := orm.Staff_detailsDB{}
	staff_detailsDB.Staff_detailsPointersEncoding = input.Staff_detailsPointersEncoding
	staff_detailsDB.CopyBasicFieldsFromStaff_details_WOP(&input.Staff_details_WOP)

	query := db.Create(&staff_detailsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaff_details.CheckoutPhaseOneInstance(&staff_detailsDB)
	staff_details := backRepo.BackRepoStaff_details.Map_Staff_detailsDBID_Staff_detailsPtr[staff_detailsDB.ID]

	if staff_details != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staff_details)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staff_detailsDB)
}

// GetStaff_details
//
// swagger:route GET /staff_detailss/{ID} staff_detailss getStaff_details
//
// Gets the details for a staff_details.
//
// Responses:
// default: genericError
//
//	200: staff_detailsDBResponse
func (controller *Controller) GetStaff_details(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_details", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_details.GetDB()

	// Get staff_detailsDB in DB
	var staff_detailsDB orm.Staff_detailsDB
	if err := db.First(&staff_detailsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staff_detailsAPI orm.Staff_detailsAPI
	staff_detailsAPI.ID = staff_detailsDB.ID
	staff_detailsAPI.Staff_detailsPointersEncoding = staff_detailsDB.Staff_detailsPointersEncoding
	staff_detailsDB.CopyBasicFieldsToStaff_details_WOP(&staff_detailsAPI.Staff_details_WOP)

	c.JSON(http.StatusOK, staff_detailsAPI)
}

// UpdateStaff_details
//
// swagger:route PATCH /staff_detailss/{ID} staff_detailss updateStaff_details
//
// # Update a staff_details
//
// Responses:
// default: genericError
//
//	200: staff_detailsDBResponse
func (controller *Controller) UpdateStaff_details(c *gin.Context) {

	mutexStaff_details.Lock()
	defer mutexStaff_details.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaff_details", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_details.GetDB()

	// Validate input
	var input orm.Staff_detailsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staff_detailsDB orm.Staff_detailsDB

	// fetch the staff_details
	query := db.First(&staff_detailsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staff_detailsDB.CopyBasicFieldsFromStaff_details_WOP(&input.Staff_details_WOP)
	staff_detailsDB.Staff_detailsPointersEncoding = input.Staff_detailsPointersEncoding

	query = db.Model(&staff_detailsDB).Updates(staff_detailsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staff_detailsNew := new(models.Staff_details)
	staff_detailsDB.CopyBasicFieldsToStaff_details(staff_detailsNew)

	// redeem pointers
	staff_detailsDB.DecodePointers(backRepo, staff_detailsNew)

	// get stage instance from DB instance, and call callback function
	staff_detailsOld := backRepo.BackRepoStaff_details.Map_Staff_detailsDBID_Staff_detailsPtr[staff_detailsDB.ID]
	if staff_detailsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staff_detailsOld, staff_detailsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staff_detailsDB
	c.JSON(http.StatusOK, staff_detailsDB)
}

// DeleteStaff_details
//
// swagger:route DELETE /staff_detailss/{ID} staff_detailss deleteStaff_details
//
// # Delete a staff_details
//
// default: genericError
//
//	200: staff_detailsDBResponse
func (controller *Controller) DeleteStaff_details(c *gin.Context) {

	mutexStaff_details.Lock()
	defer mutexStaff_details.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaff_details", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_details.GetDB()

	// Get model if exist
	var staff_detailsDB orm.Staff_detailsDB
	if err := db.First(&staff_detailsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&staff_detailsDB)

	// get an instance (not staged) from DB instance, and call callback function
	staff_detailsDeleted := new(models.Staff_details)
	staff_detailsDB.CopyBasicFieldsToStaff_details(staff_detailsDeleted)

	// get stage instance from DB instance, and call callback function
	staff_detailsStaged := backRepo.BackRepoStaff_details.Map_Staff_detailsDBID_Staff_detailsPtr[staff_detailsDB.ID]
	if staff_detailsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staff_detailsStaged, staff_detailsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
