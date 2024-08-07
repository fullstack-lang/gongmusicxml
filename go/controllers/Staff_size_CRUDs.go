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
var __Staff_size__dummysDeclaration__ models.Staff_size
var __Staff_size_time__dummyDeclaration time.Duration

var mutexStaff_size sync.Mutex

// An Staff_sizeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaff_size updateStaff_size deleteStaff_size
type Staff_sizeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Staff_sizeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaff_size updateStaff_size
type Staff_sizeInput struct {
	// The Staff_size to submit or modify
	// in: body
	Staff_size *orm.Staff_sizeAPI
}

// GetStaff_sizes
//
// swagger:route GET /staff_sizes staff_sizes getStaff_sizes
//
// # Get all staff_sizes
//
// Responses:
// default: genericError
//
//	200: staff_sizeDBResponse
func (controller *Controller) GetStaff_sizes(c *gin.Context) {

	// source slice
	var staff_sizeDBs []orm.Staff_sizeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_sizes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_size.GetDB()

	query := db.Find(&staff_sizeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staff_sizeAPIs := make([]orm.Staff_sizeAPI, 0)

	// for each staff_size, update fields from the database nullable fields
	for idx := range staff_sizeDBs {
		staff_sizeDB := &staff_sizeDBs[idx]
		_ = staff_sizeDB
		var staff_sizeAPI orm.Staff_sizeAPI

		// insertion point for updating fields
		staff_sizeAPI.ID = staff_sizeDB.ID
		staff_sizeDB.CopyBasicFieldsToStaff_size_WOP(&staff_sizeAPI.Staff_size_WOP)
		staff_sizeAPI.Staff_sizePointersEncoding = staff_sizeDB.Staff_sizePointersEncoding
		staff_sizeAPIs = append(staff_sizeAPIs, staff_sizeAPI)
	}

	c.JSON(http.StatusOK, staff_sizeAPIs)
}

// PostStaff_size
//
// swagger:route POST /staff_sizes staff_sizes postStaff_size
//
// Creates a staff_size
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaff_size(c *gin.Context) {

	mutexStaff_size.Lock()
	defer mutexStaff_size.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaff_sizes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_size.GetDB()

	// Validate input
	var input orm.Staff_sizeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staff_size
	staff_sizeDB := orm.Staff_sizeDB{}
	staff_sizeDB.Staff_sizePointersEncoding = input.Staff_sizePointersEncoding
	staff_sizeDB.CopyBasicFieldsFromStaff_size_WOP(&input.Staff_size_WOP)

	query := db.Create(&staff_sizeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaff_size.CheckoutPhaseOneInstance(&staff_sizeDB)
	staff_size := backRepo.BackRepoStaff_size.Map_Staff_sizeDBID_Staff_sizePtr[staff_sizeDB.ID]

	if staff_size != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staff_size)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staff_sizeDB)
}

// GetStaff_size
//
// swagger:route GET /staff_sizes/{ID} staff_sizes getStaff_size
//
// Gets the details for a staff_size.
//
// Responses:
// default: genericError
//
//	200: staff_sizeDBResponse
func (controller *Controller) GetStaff_size(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_size", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_size.GetDB()

	// Get staff_sizeDB in DB
	var staff_sizeDB orm.Staff_sizeDB
	if err := db.First(&staff_sizeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staff_sizeAPI orm.Staff_sizeAPI
	staff_sizeAPI.ID = staff_sizeDB.ID
	staff_sizeAPI.Staff_sizePointersEncoding = staff_sizeDB.Staff_sizePointersEncoding
	staff_sizeDB.CopyBasicFieldsToStaff_size_WOP(&staff_sizeAPI.Staff_size_WOP)

	c.JSON(http.StatusOK, staff_sizeAPI)
}

// UpdateStaff_size
//
// swagger:route PATCH /staff_sizes/{ID} staff_sizes updateStaff_size
//
// # Update a staff_size
//
// Responses:
// default: genericError
//
//	200: staff_sizeDBResponse
func (controller *Controller) UpdateStaff_size(c *gin.Context) {

	mutexStaff_size.Lock()
	defer mutexStaff_size.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaff_size", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_size.GetDB()

	// Validate input
	var input orm.Staff_sizeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staff_sizeDB orm.Staff_sizeDB

	// fetch the staff_size
	query := db.First(&staff_sizeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staff_sizeDB.CopyBasicFieldsFromStaff_size_WOP(&input.Staff_size_WOP)
	staff_sizeDB.Staff_sizePointersEncoding = input.Staff_sizePointersEncoding

	query = db.Model(&staff_sizeDB).Updates(staff_sizeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staff_sizeNew := new(models.Staff_size)
	staff_sizeDB.CopyBasicFieldsToStaff_size(staff_sizeNew)

	// redeem pointers
	staff_sizeDB.DecodePointers(backRepo, staff_sizeNew)

	// get stage instance from DB instance, and call callback function
	staff_sizeOld := backRepo.BackRepoStaff_size.Map_Staff_sizeDBID_Staff_sizePtr[staff_sizeDB.ID]
	if staff_sizeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staff_sizeOld, staff_sizeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staff_sizeDB
	c.JSON(http.StatusOK, staff_sizeDB)
}

// DeleteStaff_size
//
// swagger:route DELETE /staff_sizes/{ID} staff_sizes deleteStaff_size
//
// # Delete a staff_size
//
// default: genericError
//
//	200: staff_sizeDBResponse
func (controller *Controller) DeleteStaff_size(c *gin.Context) {

	mutexStaff_size.Lock()
	defer mutexStaff_size.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaff_size", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_size.GetDB()

	// Get model if exist
	var staff_sizeDB orm.Staff_sizeDB
	if err := db.First(&staff_sizeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&staff_sizeDB)

	// get an instance (not staged) from DB instance, and call callback function
	staff_sizeDeleted := new(models.Staff_size)
	staff_sizeDB.CopyBasicFieldsToStaff_size(staff_sizeDeleted)

	// get stage instance from DB instance, and call callback function
	staff_sizeStaged := backRepo.BackRepoStaff_size.Map_Staff_sizeDBID_Staff_sizePtr[staff_sizeDB.ID]
	if staff_sizeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staff_sizeStaged, staff_sizeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
