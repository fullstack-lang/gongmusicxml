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
var __Technical__dummysDeclaration__ models.Technical
var __Technical_time__dummyDeclaration time.Duration

var mutexTechnical sync.Mutex

// An TechnicalID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTechnical updateTechnical deleteTechnical
type TechnicalID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TechnicalInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTechnical updateTechnical
type TechnicalInput struct {
	// The Technical to submit or modify
	// in: body
	Technical *orm.TechnicalAPI
}

// GetTechnicals
//
// swagger:route GET /technicals technicals getTechnicals
//
// # Get all technicals
//
// Responses:
// default: genericError
//
//	200: technicalDBResponse
func (controller *Controller) GetTechnicals(c *gin.Context) {

	// source slice
	var technicalDBs []orm.TechnicalDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTechnicals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTechnical.GetDB()

	query := db.Find(&technicalDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	technicalAPIs := make([]orm.TechnicalAPI, 0)

	// for each technical, update fields from the database nullable fields
	for idx := range technicalDBs {
		technicalDB := &technicalDBs[idx]
		_ = technicalDB
		var technicalAPI orm.TechnicalAPI

		// insertion point for updating fields
		technicalAPI.ID = technicalDB.ID
		technicalDB.CopyBasicFieldsToTechnical_WOP(&technicalAPI.Technical_WOP)
		technicalAPI.TechnicalPointersEncoding = technicalDB.TechnicalPointersEncoding
		technicalAPIs = append(technicalAPIs, technicalAPI)
	}

	c.JSON(http.StatusOK, technicalAPIs)
}

// PostTechnical
//
// swagger:route POST /technicals technicals postTechnical
//
// Creates a technical
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTechnical(c *gin.Context) {

	mutexTechnical.Lock()
	defer mutexTechnical.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTechnicals", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTechnical.GetDB()

	// Validate input
	var input orm.TechnicalAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create technical
	technicalDB := orm.TechnicalDB{}
	technicalDB.TechnicalPointersEncoding = input.TechnicalPointersEncoding
	technicalDB.CopyBasicFieldsFromTechnical_WOP(&input.Technical_WOP)

	query := db.Create(&technicalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTechnical.CheckoutPhaseOneInstance(&technicalDB)
	technical := backRepo.BackRepoTechnical.Map_TechnicalDBID_TechnicalPtr[technicalDB.ID]

	if technical != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), technical)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, technicalDB)
}

// GetTechnical
//
// swagger:route GET /technicals/{ID} technicals getTechnical
//
// Gets the details for a technical.
//
// Responses:
// default: genericError
//
//	200: technicalDBResponse
func (controller *Controller) GetTechnical(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTechnical", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTechnical.GetDB()

	// Get technicalDB in DB
	var technicalDB orm.TechnicalDB
	if err := db.First(&technicalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var technicalAPI orm.TechnicalAPI
	technicalAPI.ID = technicalDB.ID
	technicalAPI.TechnicalPointersEncoding = technicalDB.TechnicalPointersEncoding
	technicalDB.CopyBasicFieldsToTechnical_WOP(&technicalAPI.Technical_WOP)

	c.JSON(http.StatusOK, technicalAPI)
}

// UpdateTechnical
//
// swagger:route PATCH /technicals/{ID} technicals updateTechnical
//
// # Update a technical
//
// Responses:
// default: genericError
//
//	200: technicalDBResponse
func (controller *Controller) UpdateTechnical(c *gin.Context) {

	mutexTechnical.Lock()
	defer mutexTechnical.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTechnical", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTechnical.GetDB()

	// Validate input
	var input orm.TechnicalAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var technicalDB orm.TechnicalDB

	// fetch the technical
	query := db.First(&technicalDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	technicalDB.CopyBasicFieldsFromTechnical_WOP(&input.Technical_WOP)
	technicalDB.TechnicalPointersEncoding = input.TechnicalPointersEncoding

	query = db.Model(&technicalDB).Updates(technicalDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	technicalNew := new(models.Technical)
	technicalDB.CopyBasicFieldsToTechnical(technicalNew)

	// redeem pointers
	technicalDB.DecodePointers(backRepo, technicalNew)

	// get stage instance from DB instance, and call callback function
	technicalOld := backRepo.BackRepoTechnical.Map_TechnicalDBID_TechnicalPtr[technicalDB.ID]
	if technicalOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), technicalOld, technicalNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the technicalDB
	c.JSON(http.StatusOK, technicalDB)
}

// DeleteTechnical
//
// swagger:route DELETE /technicals/{ID} technicals deleteTechnical
//
// # Delete a technical
//
// default: genericError
//
//	200: technicalDBResponse
func (controller *Controller) DeleteTechnical(c *gin.Context) {

	mutexTechnical.Lock()
	defer mutexTechnical.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTechnical", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTechnical.GetDB()

	// Get model if exist
	var technicalDB orm.TechnicalDB
	if err := db.First(&technicalDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&technicalDB)

	// get an instance (not staged) from DB instance, and call callback function
	technicalDeleted := new(models.Technical)
	technicalDB.CopyBasicFieldsToTechnical(technicalDeleted)

	// get stage instance from DB instance, and call callback function
	technicalStaged := backRepo.BackRepoTechnical.Map_TechnicalDBID_TechnicalPtr[technicalDB.ID]
	if technicalStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), technicalStaged, technicalDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
