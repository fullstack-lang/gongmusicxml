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
var __Staff_divide__dummysDeclaration__ models.Staff_divide
var __Staff_divide_time__dummyDeclaration time.Duration

var mutexStaff_divide sync.Mutex

// An Staff_divideID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaff_divide updateStaff_divide deleteStaff_divide
type Staff_divideID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Staff_divideInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaff_divide updateStaff_divide
type Staff_divideInput struct {
	// The Staff_divide to submit or modify
	// in: body
	Staff_divide *orm.Staff_divideAPI
}

// GetStaff_divides
//
// swagger:route GET /staff_divides staff_divides getStaff_divides
//
// # Get all staff_divides
//
// Responses:
// default: genericError
//
//	200: staff_divideDBResponse
func (controller *Controller) GetStaff_divides(c *gin.Context) {

	// source slice
	var staff_divideDBs []orm.Staff_divideDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_divides", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_divide.GetDB()

	query := db.Find(&staff_divideDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staff_divideAPIs := make([]orm.Staff_divideAPI, 0)

	// for each staff_divide, update fields from the database nullable fields
	for idx := range staff_divideDBs {
		staff_divideDB := &staff_divideDBs[idx]
		_ = staff_divideDB
		var staff_divideAPI orm.Staff_divideAPI

		// insertion point for updating fields
		staff_divideAPI.ID = staff_divideDB.ID
		staff_divideDB.CopyBasicFieldsToStaff_divide_WOP(&staff_divideAPI.Staff_divide_WOP)
		staff_divideAPI.Staff_dividePointersEncoding = staff_divideDB.Staff_dividePointersEncoding
		staff_divideAPIs = append(staff_divideAPIs, staff_divideAPI)
	}

	c.JSON(http.StatusOK, staff_divideAPIs)
}

// PostStaff_divide
//
// swagger:route POST /staff_divides staff_divides postStaff_divide
//
// Creates a staff_divide
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaff_divide(c *gin.Context) {

	mutexStaff_divide.Lock()
	defer mutexStaff_divide.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaff_divides", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_divide.GetDB()

	// Validate input
	var input orm.Staff_divideAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staff_divide
	staff_divideDB := orm.Staff_divideDB{}
	staff_divideDB.Staff_dividePointersEncoding = input.Staff_dividePointersEncoding
	staff_divideDB.CopyBasicFieldsFromStaff_divide_WOP(&input.Staff_divide_WOP)

	query := db.Create(&staff_divideDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaff_divide.CheckoutPhaseOneInstance(&staff_divideDB)
	staff_divide := backRepo.BackRepoStaff_divide.Map_Staff_divideDBID_Staff_dividePtr[staff_divideDB.ID]

	if staff_divide != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staff_divide)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staff_divideDB)
}

// GetStaff_divide
//
// swagger:route GET /staff_divides/{ID} staff_divides getStaff_divide
//
// Gets the details for a staff_divide.
//
// Responses:
// default: genericError
//
//	200: staff_divideDBResponse
func (controller *Controller) GetStaff_divide(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaff_divide", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_divide.GetDB()

	// Get staff_divideDB in DB
	var staff_divideDB orm.Staff_divideDB
	if err := db.First(&staff_divideDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staff_divideAPI orm.Staff_divideAPI
	staff_divideAPI.ID = staff_divideDB.ID
	staff_divideAPI.Staff_dividePointersEncoding = staff_divideDB.Staff_dividePointersEncoding
	staff_divideDB.CopyBasicFieldsToStaff_divide_WOP(&staff_divideAPI.Staff_divide_WOP)

	c.JSON(http.StatusOK, staff_divideAPI)
}

// UpdateStaff_divide
//
// swagger:route PATCH /staff_divides/{ID} staff_divides updateStaff_divide
//
// # Update a staff_divide
//
// Responses:
// default: genericError
//
//	200: staff_divideDBResponse
func (controller *Controller) UpdateStaff_divide(c *gin.Context) {

	mutexStaff_divide.Lock()
	defer mutexStaff_divide.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaff_divide", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_divide.GetDB()

	// Validate input
	var input orm.Staff_divideAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staff_divideDB orm.Staff_divideDB

	// fetch the staff_divide
	query := db.First(&staff_divideDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staff_divideDB.CopyBasicFieldsFromStaff_divide_WOP(&input.Staff_divide_WOP)
	staff_divideDB.Staff_dividePointersEncoding = input.Staff_dividePointersEncoding

	query = db.Model(&staff_divideDB).Updates(staff_divideDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staff_divideNew := new(models.Staff_divide)
	staff_divideDB.CopyBasicFieldsToStaff_divide(staff_divideNew)

	// redeem pointers
	staff_divideDB.DecodePointers(backRepo, staff_divideNew)

	// get stage instance from DB instance, and call callback function
	staff_divideOld := backRepo.BackRepoStaff_divide.Map_Staff_divideDBID_Staff_dividePtr[staff_divideDB.ID]
	if staff_divideOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staff_divideOld, staff_divideNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staff_divideDB
	c.JSON(http.StatusOK, staff_divideDB)
}

// DeleteStaff_divide
//
// swagger:route DELETE /staff_divides/{ID} staff_divides deleteStaff_divide
//
// # Delete a staff_divide
//
// default: genericError
//
//	200: staff_divideDBResponse
func (controller *Controller) DeleteStaff_divide(c *gin.Context) {

	mutexStaff_divide.Lock()
	defer mutexStaff_divide.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaff_divide", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoStaff_divide.GetDB()

	// Get model if exist
	var staff_divideDB orm.Staff_divideDB
	if err := db.First(&staff_divideDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&staff_divideDB)

	// get an instance (not staged) from DB instance, and call callback function
	staff_divideDeleted := new(models.Staff_divide)
	staff_divideDB.CopyBasicFieldsToStaff_divide(staff_divideDeleted)

	// get stage instance from DB instance, and call callback function
	staff_divideStaged := backRepo.BackRepoStaff_divide.Map_Staff_divideDBID_Staff_dividePtr[staff_divideDB.ID]
	if staff_divideStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staff_divideStaged, staff_divideDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
