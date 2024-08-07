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
var __Accordion_registration__dummysDeclaration__ models.Accordion_registration
var __Accordion_registration_time__dummyDeclaration time.Duration

var mutexAccordion_registration sync.Mutex

// An Accordion_registrationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAccordion_registration updateAccordion_registration deleteAccordion_registration
type Accordion_registrationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Accordion_registrationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAccordion_registration updateAccordion_registration
type Accordion_registrationInput struct {
	// The Accordion_registration to submit or modify
	// in: body
	Accordion_registration *orm.Accordion_registrationAPI
}

// GetAccordion_registrations
//
// swagger:route GET /accordion_registrations accordion_registrations getAccordion_registrations
//
// # Get all accordion_registrations
//
// Responses:
// default: genericError
//
//	200: accordion_registrationDBResponse
func (controller *Controller) GetAccordion_registrations(c *gin.Context) {

	// source slice
	var accordion_registrationDBs []orm.Accordion_registrationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccordion_registrations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccordion_registration.GetDB()

	query := db.Find(&accordion_registrationDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	accordion_registrationAPIs := make([]orm.Accordion_registrationAPI, 0)

	// for each accordion_registration, update fields from the database nullable fields
	for idx := range accordion_registrationDBs {
		accordion_registrationDB := &accordion_registrationDBs[idx]
		_ = accordion_registrationDB
		var accordion_registrationAPI orm.Accordion_registrationAPI

		// insertion point for updating fields
		accordion_registrationAPI.ID = accordion_registrationDB.ID
		accordion_registrationDB.CopyBasicFieldsToAccordion_registration_WOP(&accordion_registrationAPI.Accordion_registration_WOP)
		accordion_registrationAPI.Accordion_registrationPointersEncoding = accordion_registrationDB.Accordion_registrationPointersEncoding
		accordion_registrationAPIs = append(accordion_registrationAPIs, accordion_registrationAPI)
	}

	c.JSON(http.StatusOK, accordion_registrationAPIs)
}

// PostAccordion_registration
//
// swagger:route POST /accordion_registrations accordion_registrations postAccordion_registration
//
// Creates a accordion_registration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAccordion_registration(c *gin.Context) {

	mutexAccordion_registration.Lock()
	defer mutexAccordion_registration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAccordion_registrations", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccordion_registration.GetDB()

	// Validate input
	var input orm.Accordion_registrationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create accordion_registration
	accordion_registrationDB := orm.Accordion_registrationDB{}
	accordion_registrationDB.Accordion_registrationPointersEncoding = input.Accordion_registrationPointersEncoding
	accordion_registrationDB.CopyBasicFieldsFromAccordion_registration_WOP(&input.Accordion_registration_WOP)

	query := db.Create(&accordion_registrationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAccordion_registration.CheckoutPhaseOneInstance(&accordion_registrationDB)
	accordion_registration := backRepo.BackRepoAccordion_registration.Map_Accordion_registrationDBID_Accordion_registrationPtr[accordion_registrationDB.ID]

	if accordion_registration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), accordion_registration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, accordion_registrationDB)
}

// GetAccordion_registration
//
// swagger:route GET /accordion_registrations/{ID} accordion_registrations getAccordion_registration
//
// Gets the details for a accordion_registration.
//
// Responses:
// default: genericError
//
//	200: accordion_registrationDBResponse
func (controller *Controller) GetAccordion_registration(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAccordion_registration", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccordion_registration.GetDB()

	// Get accordion_registrationDB in DB
	var accordion_registrationDB orm.Accordion_registrationDB
	if err := db.First(&accordion_registrationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var accordion_registrationAPI orm.Accordion_registrationAPI
	accordion_registrationAPI.ID = accordion_registrationDB.ID
	accordion_registrationAPI.Accordion_registrationPointersEncoding = accordion_registrationDB.Accordion_registrationPointersEncoding
	accordion_registrationDB.CopyBasicFieldsToAccordion_registration_WOP(&accordion_registrationAPI.Accordion_registration_WOP)

	c.JSON(http.StatusOK, accordion_registrationAPI)
}

// UpdateAccordion_registration
//
// swagger:route PATCH /accordion_registrations/{ID} accordion_registrations updateAccordion_registration
//
// # Update a accordion_registration
//
// Responses:
// default: genericError
//
//	200: accordion_registrationDBResponse
func (controller *Controller) UpdateAccordion_registration(c *gin.Context) {

	mutexAccordion_registration.Lock()
	defer mutexAccordion_registration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAccordion_registration", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccordion_registration.GetDB()

	// Validate input
	var input orm.Accordion_registrationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var accordion_registrationDB orm.Accordion_registrationDB

	// fetch the accordion_registration
	query := db.First(&accordion_registrationDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	accordion_registrationDB.CopyBasicFieldsFromAccordion_registration_WOP(&input.Accordion_registration_WOP)
	accordion_registrationDB.Accordion_registrationPointersEncoding = input.Accordion_registrationPointersEncoding

	query = db.Model(&accordion_registrationDB).Updates(accordion_registrationDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	accordion_registrationNew := new(models.Accordion_registration)
	accordion_registrationDB.CopyBasicFieldsToAccordion_registration(accordion_registrationNew)

	// redeem pointers
	accordion_registrationDB.DecodePointers(backRepo, accordion_registrationNew)

	// get stage instance from DB instance, and call callback function
	accordion_registrationOld := backRepo.BackRepoAccordion_registration.Map_Accordion_registrationDBID_Accordion_registrationPtr[accordion_registrationDB.ID]
	if accordion_registrationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), accordion_registrationOld, accordion_registrationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the accordion_registrationDB
	c.JSON(http.StatusOK, accordion_registrationDB)
}

// DeleteAccordion_registration
//
// swagger:route DELETE /accordion_registrations/{ID} accordion_registrations deleteAccordion_registration
//
// # Delete a accordion_registration
//
// default: genericError
//
//	200: accordion_registrationDBResponse
func (controller *Controller) DeleteAccordion_registration(c *gin.Context) {

	mutexAccordion_registration.Lock()
	defer mutexAccordion_registration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAccordion_registration", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAccordion_registration.GetDB()

	// Get model if exist
	var accordion_registrationDB orm.Accordion_registrationDB
	if err := db.First(&accordion_registrationDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&accordion_registrationDB)

	// get an instance (not staged) from DB instance, and call callback function
	accordion_registrationDeleted := new(models.Accordion_registration)
	accordion_registrationDB.CopyBasicFieldsToAccordion_registration(accordion_registrationDeleted)

	// get stage instance from DB instance, and call callback function
	accordion_registrationStaged := backRepo.BackRepoAccordion_registration.Map_Accordion_registrationDBID_Accordion_registrationPtr[accordion_registrationDB.ID]
	if accordion_registrationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), accordion_registrationStaged, accordion_registrationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
