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
var __Score_partwise__dummysDeclaration__ models.Score_partwise
var __Score_partwise_time__dummyDeclaration time.Duration

var mutexScore_partwise sync.Mutex

// An Score_partwiseID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScore_partwise updateScore_partwise deleteScore_partwise
type Score_partwiseID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Score_partwiseInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScore_partwise updateScore_partwise
type Score_partwiseInput struct {
	// The Score_partwise to submit or modify
	// in: body
	Score_partwise *orm.Score_partwiseAPI
}

// GetScore_partwises
//
// swagger:route GET /score_partwises score_partwises getScore_partwises
//
// # Get all score_partwises
//
// Responses:
// default: genericError
//
//	200: score_partwiseDBResponse
func (controller *Controller) GetScore_partwises(c *gin.Context) {

	// source slice
	var score_partwiseDBs []orm.Score_partwiseDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_partwises", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_partwise.GetDB()

	query := db.Find(&score_partwiseDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	score_partwiseAPIs := make([]orm.Score_partwiseAPI, 0)

	// for each score_partwise, update fields from the database nullable fields
	for idx := range score_partwiseDBs {
		score_partwiseDB := &score_partwiseDBs[idx]
		_ = score_partwiseDB
		var score_partwiseAPI orm.Score_partwiseAPI

		// insertion point for updating fields
		score_partwiseAPI.ID = score_partwiseDB.ID
		score_partwiseDB.CopyBasicFieldsToScore_partwise_WOP(&score_partwiseAPI.Score_partwise_WOP)
		score_partwiseAPI.Score_partwisePointersEncoding = score_partwiseDB.Score_partwisePointersEncoding
		score_partwiseAPIs = append(score_partwiseAPIs, score_partwiseAPI)
	}

	c.JSON(http.StatusOK, score_partwiseAPIs)
}

// PostScore_partwise
//
// swagger:route POST /score_partwises score_partwises postScore_partwise
//
// Creates a score_partwise
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScore_partwise(c *gin.Context) {

	mutexScore_partwise.Lock()
	defer mutexScore_partwise.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScore_partwises", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_partwise.GetDB()

	// Validate input
	var input orm.Score_partwiseAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create score_partwise
	score_partwiseDB := orm.Score_partwiseDB{}
	score_partwiseDB.Score_partwisePointersEncoding = input.Score_partwisePointersEncoding
	score_partwiseDB.CopyBasicFieldsFromScore_partwise_WOP(&input.Score_partwise_WOP)

	query := db.Create(&score_partwiseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScore_partwise.CheckoutPhaseOneInstance(&score_partwiseDB)
	score_partwise := backRepo.BackRepoScore_partwise.Map_Score_partwiseDBID_Score_partwisePtr[score_partwiseDB.ID]

	if score_partwise != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), score_partwise)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, score_partwiseDB)
}

// GetScore_partwise
//
// swagger:route GET /score_partwises/{ID} score_partwises getScore_partwise
//
// Gets the details for a score_partwise.
//
// Responses:
// default: genericError
//
//	200: score_partwiseDBResponse
func (controller *Controller) GetScore_partwise(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_partwise", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_partwise.GetDB()

	// Get score_partwiseDB in DB
	var score_partwiseDB orm.Score_partwiseDB
	if err := db.First(&score_partwiseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var score_partwiseAPI orm.Score_partwiseAPI
	score_partwiseAPI.ID = score_partwiseDB.ID
	score_partwiseAPI.Score_partwisePointersEncoding = score_partwiseDB.Score_partwisePointersEncoding
	score_partwiseDB.CopyBasicFieldsToScore_partwise_WOP(&score_partwiseAPI.Score_partwise_WOP)

	c.JSON(http.StatusOK, score_partwiseAPI)
}

// UpdateScore_partwise
//
// swagger:route PATCH /score_partwises/{ID} score_partwises updateScore_partwise
//
// # Update a score_partwise
//
// Responses:
// default: genericError
//
//	200: score_partwiseDBResponse
func (controller *Controller) UpdateScore_partwise(c *gin.Context) {

	mutexScore_partwise.Lock()
	defer mutexScore_partwise.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScore_partwise", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_partwise.GetDB()

	// Validate input
	var input orm.Score_partwiseAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var score_partwiseDB orm.Score_partwiseDB

	// fetch the score_partwise
	query := db.First(&score_partwiseDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	score_partwiseDB.CopyBasicFieldsFromScore_partwise_WOP(&input.Score_partwise_WOP)
	score_partwiseDB.Score_partwisePointersEncoding = input.Score_partwisePointersEncoding

	query = db.Model(&score_partwiseDB).Updates(score_partwiseDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	score_partwiseNew := new(models.Score_partwise)
	score_partwiseDB.CopyBasicFieldsToScore_partwise(score_partwiseNew)

	// redeem pointers
	score_partwiseDB.DecodePointers(backRepo, score_partwiseNew)

	// get stage instance from DB instance, and call callback function
	score_partwiseOld := backRepo.BackRepoScore_partwise.Map_Score_partwiseDBID_Score_partwisePtr[score_partwiseDB.ID]
	if score_partwiseOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), score_partwiseOld, score_partwiseNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the score_partwiseDB
	c.JSON(http.StatusOK, score_partwiseDB)
}

// DeleteScore_partwise
//
// swagger:route DELETE /score_partwises/{ID} score_partwises deleteScore_partwise
//
// # Delete a score_partwise
//
// default: genericError
//
//	200: score_partwiseDBResponse
func (controller *Controller) DeleteScore_partwise(c *gin.Context) {

	mutexScore_partwise.Lock()
	defer mutexScore_partwise.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScore_partwise", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_partwise.GetDB()

	// Get model if exist
	var score_partwiseDB orm.Score_partwiseDB
	if err := db.First(&score_partwiseDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&score_partwiseDB)

	// get an instance (not staged) from DB instance, and call callback function
	score_partwiseDeleted := new(models.Score_partwise)
	score_partwiseDB.CopyBasicFieldsToScore_partwise(score_partwiseDeleted)

	// get stage instance from DB instance, and call callback function
	score_partwiseStaged := backRepo.BackRepoScore_partwise.Map_Score_partwiseDBID_Score_partwisePtr[score_partwiseDB.ID]
	if score_partwiseStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), score_partwiseStaged, score_partwiseDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
