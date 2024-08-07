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
var __Staff_tuning__dummysDeclaration__ models.Staff_tuning
var __Staff_tuning_time__dummyDeclaration time.Duration

var mutexStaff_tuning sync.Mutex

// An Staff_tuningID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaff_tuning updateStaff_tuning deleteStaff_tuning
type Staff_tuningID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Staff_tuningInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaff_tuning updateStaff_tuning
type Staff_tuningInput struct {
	// The Staff_tuning to submit or modify
	// in: body
	Staff_tuning *orm.Staff_tuningAPI
}

// GetStaff_tunings
//
// swagger:route GET /staff_tunings staff_tunings getStaff_tunings
//
// # Get all staff_tunings
//
// Responses:
// default: genericError
//
//	200: staff_tuningDBResponse
func (controller *Controller) GetStaff_tunings(c *gin.Context) {

	// source slice
	var staff_tuningDBs []orm.Staff_tuningDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_tunings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_tuning.GetDB()

	query := db.Find(&staff_tuningDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staff_tuningAPIs := make([]orm.Staff_tuningAPI, 0)

	// for each staff_tuning, update fields from the database nullable fields
	for idx := range staff_tuningDBs {
		staff_tuningDB := &staff_tuningDBs[idx]
		_ = staff_tuningDB
		var staff_tuningAPI orm.Staff_tuningAPI

		// insertion point for updating fields
		staff_tuningAPI.ID = staff_tuningDB.ID
		staff_tuningDB.CopyBasicFieldsToStaff_tuning_WOP(&staff_tuningAPI.Staff_tuning_WOP)
		staff_tuningAPI.Staff_tuningPointersEncoding = staff_tuningDB.Staff_tuningPointersEncoding
		staff_tuningAPIs = append(staff_tuningAPIs, staff_tuningAPI)
	}

	c.JSON(http.StatusOK, staff_tuningAPIs)
}

// PostStaff_tuning
//
// swagger:route POST /staff_tunings staff_tunings postStaff_tuning
//
// Creates a staff_tuning
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaff_tuning(c *gin.Context) {

	mutexStaff_tuning.Lock()
	defer mutexStaff_tuning.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaff_tunings", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_tuning.GetDB()

	// Validate input
	var input orm.Staff_tuningAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staff_tuning
	staff_tuningDB := orm.Staff_tuningDB{}
	staff_tuningDB.Staff_tuningPointersEncoding = input.Staff_tuningPointersEncoding
	staff_tuningDB.CopyBasicFieldsFromStaff_tuning_WOP(&input.Staff_tuning_WOP)

	query := db.Create(&staff_tuningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaff_tuning.CheckoutPhaseOneInstance(&staff_tuningDB)
	staff_tuning := backRepo.BackRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID]

	if staff_tuning != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staff_tuning)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staff_tuningDB)
}

// GetStaff_tuning
//
// swagger:route GET /staff_tunings/{ID} staff_tunings getStaff_tuning
//
// Gets the details for a staff_tuning.
//
// Responses:
// default: genericError
//
//	200: staff_tuningDBResponse
func (controller *Controller) GetStaff_tuning(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_tuning", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_tuning.GetDB()

	// Get staff_tuningDB in DB
	var staff_tuningDB orm.Staff_tuningDB
	if err := db.First(&staff_tuningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staff_tuningAPI orm.Staff_tuningAPI
	staff_tuningAPI.ID = staff_tuningDB.ID
	staff_tuningAPI.Staff_tuningPointersEncoding = staff_tuningDB.Staff_tuningPointersEncoding
	staff_tuningDB.CopyBasicFieldsToStaff_tuning_WOP(&staff_tuningAPI.Staff_tuning_WOP)

	c.JSON(http.StatusOK, staff_tuningAPI)
}

// UpdateStaff_tuning
//
// swagger:route PATCH /staff_tunings/{ID} staff_tunings updateStaff_tuning
//
// # Update a staff_tuning
//
// Responses:
// default: genericError
//
//	200: staff_tuningDBResponse
func (controller *Controller) UpdateStaff_tuning(c *gin.Context) {

	mutexStaff_tuning.Lock()
	defer mutexStaff_tuning.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaff_tuning", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_tuning.GetDB()

	// Validate input
	var input orm.Staff_tuningAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staff_tuningDB orm.Staff_tuningDB

	// fetch the staff_tuning
	query := db.First(&staff_tuningDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staff_tuningDB.CopyBasicFieldsFromStaff_tuning_WOP(&input.Staff_tuning_WOP)
	staff_tuningDB.Staff_tuningPointersEncoding = input.Staff_tuningPointersEncoding

	query = db.Model(&staff_tuningDB).Updates(staff_tuningDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staff_tuningNew := new(models.Staff_tuning)
	staff_tuningDB.CopyBasicFieldsToStaff_tuning(staff_tuningNew)

	// redeem pointers
	staff_tuningDB.DecodePointers(backRepo, staff_tuningNew)

	// get stage instance from DB instance, and call callback function
	staff_tuningOld := backRepo.BackRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID]
	if staff_tuningOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staff_tuningOld, staff_tuningNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staff_tuningDB
	c.JSON(http.StatusOK, staff_tuningDB)
}

// DeleteStaff_tuning
//
// swagger:route DELETE /staff_tunings/{ID} staff_tunings deleteStaff_tuning
//
// # Delete a staff_tuning
//
// default: genericError
//
//	200: staff_tuningDBResponse
func (controller *Controller) DeleteStaff_tuning(c *gin.Context) {

	mutexStaff_tuning.Lock()
	defer mutexStaff_tuning.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaff_tuning", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_tuning.GetDB()

	// Get model if exist
	var staff_tuningDB orm.Staff_tuningDB
	if err := db.First(&staff_tuningDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&staff_tuningDB)

	// get an instance (not staged) from DB instance, and call callback function
	staff_tuningDeleted := new(models.Staff_tuning)
	staff_tuningDB.CopyBasicFieldsToStaff_tuning(staff_tuningDeleted)

	// get stage instance from DB instance, and call callback function
	staff_tuningStaged := backRepo.BackRepoStaff_tuning.Map_Staff_tuningDBID_Staff_tuningPtr[staff_tuningDB.ID]
	if staff_tuningStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staff_tuningStaged, staff_tuningDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
