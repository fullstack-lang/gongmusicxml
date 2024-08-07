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
var __Score_timewise__dummysDeclaration__ models.Score_timewise
var __Score_timewise_time__dummyDeclaration time.Duration

var mutexScore_timewise sync.Mutex

// An Score_timewiseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScore_timewise updateScore_timewise deleteScore_timewise
type Score_timewiseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Score_timewiseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScore_timewise updateScore_timewise
type Score_timewiseInput struct {
	// The Score_timewise to submit or modify
	// in: body
	Score_timewise *orm.Score_timewiseAPI
}

// GetScore_timewises
//
// swagger:route GET /score_timewises score_timewises getScore_timewises
//
// # Get all score_timewises
//
// Responses:
// default: genericError
//
//	200: score_timewiseDBResponse
func (controller *Controller) GetScore_timewises(c *gin.Context) {

	// source slice
	var score_timewiseDBs []orm.Score_timewiseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_timewises", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_timewise.GetDB()

	query := db.Find(&score_timewiseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	score_timewiseAPIs := make([]orm.Score_timewiseAPI, 0)

	// for each score_timewise, update fields from the database nullable fields
	for idx := range score_timewiseDBs {
		score_timewiseDB := &score_timewiseDBs[idx]
		_ = score_timewiseDB
		var score_timewiseAPI orm.Score_timewiseAPI

		// insertion point for updating fields
		score_timewiseAPI.ID = score_timewiseDB.ID
		score_timewiseDB.CopyBasicFieldsToScore_timewise_WOP(&score_timewiseAPI.Score_timewise_WOP)
		score_timewiseAPI.Score_timewisePointersEncoding = score_timewiseDB.Score_timewisePointersEncoding
		score_timewiseAPIs = append(score_timewiseAPIs, score_timewiseAPI)
	}

	c.JSON(http.StatusOK, score_timewiseAPIs)
}

// PostScore_timewise
//
// swagger:route POST /score_timewises score_timewises postScore_timewise
//
// Creates a score_timewise
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScore_timewise(c *gin.Context) {

	mutexScore_timewise.Lock()
	defer mutexScore_timewise.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScore_timewises", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_timewise.GetDB()

	// Validate input
	var input orm.Score_timewiseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create score_timewise
	score_timewiseDB := orm.Score_timewiseDB{}
	score_timewiseDB.Score_timewisePointersEncoding = input.Score_timewisePointersEncoding
	score_timewiseDB.CopyBasicFieldsFromScore_timewise_WOP(&input.Score_timewise_WOP)

	query := db.Create(&score_timewiseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScore_timewise.CheckoutPhaseOneInstance(&score_timewiseDB)
	score_timewise := backRepo.BackRepoScore_timewise.Map_Score_timewiseDBID_Score_timewisePtr[score_timewiseDB.ID]

	if score_timewise != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), score_timewise)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, score_timewiseDB)
}

// GetScore_timewise
//
// swagger:route GET /score_timewises/{ID} score_timewises getScore_timewise
//
// Gets the details for a score_timewise.
//
// Responses:
// default: genericError
//
//	200: score_timewiseDBResponse
func (controller *Controller) GetScore_timewise(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_timewise", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_timewise.GetDB()

	// Get score_timewiseDB in DB
	var score_timewiseDB orm.Score_timewiseDB
	if err := db.First(&score_timewiseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var score_timewiseAPI orm.Score_timewiseAPI
	score_timewiseAPI.ID = score_timewiseDB.ID
	score_timewiseAPI.Score_timewisePointersEncoding = score_timewiseDB.Score_timewisePointersEncoding
	score_timewiseDB.CopyBasicFieldsToScore_timewise_WOP(&score_timewiseAPI.Score_timewise_WOP)

	c.JSON(http.StatusOK, score_timewiseAPI)
}

// UpdateScore_timewise
//
// swagger:route PATCH /score_timewises/{ID} score_timewises updateScore_timewise
//
// # Update a score_timewise
//
// Responses:
// default: genericError
//
//	200: score_timewiseDBResponse
func (controller *Controller) UpdateScore_timewise(c *gin.Context) {

	mutexScore_timewise.Lock()
	defer mutexScore_timewise.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScore_timewise", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_timewise.GetDB()

	// Validate input
	var input orm.Score_timewiseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var score_timewiseDB orm.Score_timewiseDB

	// fetch the score_timewise
	query := db.First(&score_timewiseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	score_timewiseDB.CopyBasicFieldsFromScore_timewise_WOP(&input.Score_timewise_WOP)
	score_timewiseDB.Score_timewisePointersEncoding = input.Score_timewisePointersEncoding

	query = db.Model(&score_timewiseDB).Updates(score_timewiseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	score_timewiseNew := new(models.Score_timewise)
	score_timewiseDB.CopyBasicFieldsToScore_timewise(score_timewiseNew)

	// redeem pointers
	score_timewiseDB.DecodePointers(backRepo, score_timewiseNew)

	// get stage instance from DB instance, and call callback function
	score_timewiseOld := backRepo.BackRepoScore_timewise.Map_Score_timewiseDBID_Score_timewisePtr[score_timewiseDB.ID]
	if score_timewiseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), score_timewiseOld, score_timewiseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the score_timewiseDB
	c.JSON(http.StatusOK, score_timewiseDB)
}

// DeleteScore_timewise
//
// swagger:route DELETE /score_timewises/{ID} score_timewises deleteScore_timewise
//
// # Delete a score_timewise
//
// default: genericError
//
//	200: score_timewiseDBResponse
func (controller *Controller) DeleteScore_timewise(c *gin.Context) {

	mutexScore_timewise.Lock()
	defer mutexScore_timewise.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScore_timewise", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_timewise.GetDB()

	// Get model if exist
	var score_timewiseDB orm.Score_timewiseDB
	if err := db.First(&score_timewiseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&score_timewiseDB)

	// get an instance (not staged) from DB instance, and call callback function
	score_timewiseDeleted := new(models.Score_timewise)
	score_timewiseDB.CopyBasicFieldsToScore_timewise(score_timewiseDeleted)

	// get stage instance from DB instance, and call callback function
	score_timewiseStaged := backRepo.BackRepoScore_timewise.Map_Score_timewiseDBID_Score_timewisePtr[score_timewiseDB.ID]
	if score_timewiseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), score_timewiseStaged, score_timewiseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
