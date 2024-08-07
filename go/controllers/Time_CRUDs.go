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
var __Time__dummysDeclaration__ models.Time
var __Time_time__dummyDeclaration time.Duration

var mutexTime sync.Mutex

// An TimeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTime updateTime deleteTime
type TimeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TimeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTime updateTime
type TimeInput struct {
	// The Time to submit or modify
	// in: body
	Time *orm.TimeAPI
}

// GetTimes
//
// swagger:route GET /times times getTimes
//
// # Get all times
//
// Responses:
// default: genericError
//
//	200: timeDBResponse
func (controller *Controller) GetTimes(c *gin.Context) {

	// source slice
	var timeDBs []orm.TimeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime.GetDB()

	query := db.Find(&timeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	timeAPIs := make([]orm.TimeAPI, 0)

	// for each time, update fields from the database nullable fields
	for idx := range timeDBs {
		timeDB := &timeDBs[idx]
		_ = timeDB
		var timeAPI orm.TimeAPI

		// insertion point for updating fields
		timeAPI.ID = timeDB.ID
		timeDB.CopyBasicFieldsToTime_WOP(&timeAPI.Time_WOP)
		timeAPI.TimePointersEncoding = timeDB.TimePointersEncoding
		timeAPIs = append(timeAPIs, timeAPI)
	}

	c.JSON(http.StatusOK, timeAPIs)
}

// PostTime
//
// swagger:route POST /times times postTime
//
// Creates a time
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTime(c *gin.Context) {

	mutexTime.Lock()
	defer mutexTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTimes", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime.GetDB()

	// Validate input
	var input orm.TimeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create time
	timeDB := orm.TimeDB{}
	timeDB.TimePointersEncoding = input.TimePointersEncoding
	timeDB.CopyBasicFieldsFromTime_WOP(&input.Time_WOP)

	query := db.Create(&timeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTime.CheckoutPhaseOneInstance(&timeDB)
	time := backRepo.BackRepoTime.Map_TimeDBID_TimePtr[timeDB.ID]

	if time != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), time)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, timeDB)
}

// GetTime
//
// swagger:route GET /times/{ID} times getTime
//
// Gets the details for a time.
//
// Responses:
// default: genericError
//
//	200: timeDBResponse
func (controller *Controller) GetTime(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime.GetDB()

	// Get timeDB in DB
	var timeDB orm.TimeDB
	if err := db.First(&timeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var timeAPI orm.TimeAPI
	timeAPI.ID = timeDB.ID
	timeAPI.TimePointersEncoding = timeDB.TimePointersEncoding
	timeDB.CopyBasicFieldsToTime_WOP(&timeAPI.Time_WOP)

	c.JSON(http.StatusOK, timeAPI)
}

// UpdateTime
//
// swagger:route PATCH /times/{ID} times updateTime
//
// # Update a time
//
// Responses:
// default: genericError
//
//	200: timeDBResponse
func (controller *Controller) UpdateTime(c *gin.Context) {

	mutexTime.Lock()
	defer mutexTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime.GetDB()

	// Validate input
	var input orm.TimeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var timeDB orm.TimeDB

	// fetch the time
	query := db.First(&timeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	timeDB.CopyBasicFieldsFromTime_WOP(&input.Time_WOP)
	timeDB.TimePointersEncoding = input.TimePointersEncoding

	query = db.Model(&timeDB).Updates(timeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	timeNew := new(models.Time)
	timeDB.CopyBasicFieldsToTime(timeNew)

	// redeem pointers
	timeDB.DecodePointers(backRepo, timeNew)

	// get stage instance from DB instance, and call callback function
	timeOld := backRepo.BackRepoTime.Map_TimeDBID_TimePtr[timeDB.ID]
	if timeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), timeOld, timeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the timeDB
	c.JSON(http.StatusOK, timeDB)
}

// DeleteTime
//
// swagger:route DELETE /times/{ID} times deleteTime
//
// # Delete a time
//
// default: genericError
//
//	200: timeDBResponse
func (controller *Controller) DeleteTime(c *gin.Context) {

	mutexTime.Lock()
	defer mutexTime.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTime", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTime.GetDB()

	// Get model if exist
	var timeDB orm.TimeDB
	if err := db.First(&timeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&timeDB)

	// get an instance (not staged) from DB instance, and call callback function
	timeDeleted := new(models.Time)
	timeDB.CopyBasicFieldsToTime(timeDeleted)

	// get stage instance from DB instance, and call callback function
	timeStaged := backRepo.BackRepoTime.Map_TimeDBID_TimePtr[timeDB.ID]
	if timeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), timeStaged, timeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
