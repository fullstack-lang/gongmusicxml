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
var __Assess__dummysDeclaration__ models.Assess
var __Assess_time__dummyDeclaration time.Duration

var mutexAssess sync.Mutex

// An AssessID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getAssess updateAssess deleteAssess
type AssessID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// AssessInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAssess updateAssess
type AssessInput struct {
	// The Assess to submit or modify
	// in: body
	Assess *orm.AssessAPI
}

// GetAssesss
//
// swagger:route GET /assesss assesss getAssesss
//
// # Get all assesss
//
// Responses:
// default: genericError
//
//	200: assessDBResponse
func (controller *Controller) GetAssesss(c *gin.Context) {

	// source slice
	var assessDBs []orm.AssessDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAssesss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAssess.GetDB()

	query := db.Find(&assessDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	assessAPIs := make([]orm.AssessAPI, 0)

	// for each assess, update fields from the database nullable fields
	for idx := range assessDBs {
		assessDB := &assessDBs[idx]
		_ = assessDB
		var assessAPI orm.AssessAPI

		// insertion point for updating fields
		assessAPI.ID = assessDB.ID
		assessDB.CopyBasicFieldsToAssess_WOP(&assessAPI.Assess_WOP)
		assessAPI.AssessPointersEncoding = assessDB.AssessPointersEncoding
		assessAPIs = append(assessAPIs, assessAPI)
	}

	c.JSON(http.StatusOK, assessAPIs)
}

// PostAssess
//
// swagger:route POST /assesss assesss postAssess
//
// Creates a assess
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostAssess(c *gin.Context) {

	mutexAssess.Lock()
	defer mutexAssess.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostAssesss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAssess.GetDB()

	// Validate input
	var input orm.AssessAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create assess
	assessDB := orm.AssessDB{}
	assessDB.AssessPointersEncoding = input.AssessPointersEncoding
	assessDB.CopyBasicFieldsFromAssess_WOP(&input.Assess_WOP)

	query := db.Create(&assessDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoAssess.CheckoutPhaseOneInstance(&assessDB)
	assess := backRepo.BackRepoAssess.Map_AssessDBID_AssessPtr[assessDB.ID]

	if assess != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), assess)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, assessDB)
}

// GetAssess
//
// swagger:route GET /assesss/{ID} assesss getAssess
//
// Gets the details for a assess.
//
// Responses:
// default: genericError
//
//	200: assessDBResponse
func (controller *Controller) GetAssess(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetAssess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAssess.GetDB()

	// Get assessDB in DB
	var assessDB orm.AssessDB
	if err := db.First(&assessDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var assessAPI orm.AssessAPI
	assessAPI.ID = assessDB.ID
	assessAPI.AssessPointersEncoding = assessDB.AssessPointersEncoding
	assessDB.CopyBasicFieldsToAssess_WOP(&assessAPI.Assess_WOP)

	c.JSON(http.StatusOK, assessAPI)
}

// UpdateAssess
//
// swagger:route PATCH /assesss/{ID} assesss updateAssess
//
// # Update a assess
//
// Responses:
// default: genericError
//
//	200: assessDBResponse
func (controller *Controller) UpdateAssess(c *gin.Context) {

	mutexAssess.Lock()
	defer mutexAssess.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateAssess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAssess.GetDB()

	// Validate input
	var input orm.AssessAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var assessDB orm.AssessDB

	// fetch the assess
	query := db.First(&assessDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	assessDB.CopyBasicFieldsFromAssess_WOP(&input.Assess_WOP)
	assessDB.AssessPointersEncoding = input.AssessPointersEncoding

	query = db.Model(&assessDB).Updates(assessDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	assessNew := new(models.Assess)
	assessDB.CopyBasicFieldsToAssess(assessNew)

	// redeem pointers
	assessDB.DecodePointers(backRepo, assessNew)

	// get stage instance from DB instance, and call callback function
	assessOld := backRepo.BackRepoAssess.Map_AssessDBID_AssessPtr[assessDB.ID]
	if assessOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), assessOld, assessNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the assessDB
	c.JSON(http.StatusOK, assessDB)
}

// DeleteAssess
//
// swagger:route DELETE /assesss/{ID} assesss deleteAssess
//
// # Delete a assess
//
// default: genericError
//
//	200: assessDBResponse
func (controller *Controller) DeleteAssess(c *gin.Context) {

	mutexAssess.Lock()
	defer mutexAssess.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteAssess", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoAssess.GetDB()

	// Get model if exist
	var assessDB orm.AssessDB
	if err := db.First(&assessDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&assessDB)

	// get an instance (not staged) from DB instance, and call callback function
	assessDeleted := new(models.Assess)
	assessDB.CopyBasicFieldsToAssess(assessDeleted)

	// get stage instance from DB instance, and call callback function
	assessStaged := backRepo.BackRepoAssess.Map_AssessDBID_AssessPtr[assessDB.ID]
	if assessStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), assessStaged, assessDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
