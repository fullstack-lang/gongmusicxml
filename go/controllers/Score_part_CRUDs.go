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
var __Score_part__dummysDeclaration__ models.Score_part
var __Score_part_time__dummyDeclaration time.Duration

var mutexScore_part sync.Mutex

// An Score_partID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getScore_part updateScore_part deleteScore_part
type Score_partID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Score_partInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postScore_part updateScore_part
type Score_partInput struct {
	// The Score_part to submit or modify
	// in: body
	Score_part *orm.Score_partAPI
}

// GetScore_parts
//
// swagger:route GET /score_parts score_parts getScore_parts
//
// # Get all score_parts
//
// Responses:
// default: genericError
//
//	200: score_partDBResponse
func (controller *Controller) GetScore_parts(c *gin.Context) {

	// source slice
	var score_partDBs []orm.Score_partDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_parts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_part.GetDB()

	query := db.Find(&score_partDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	score_partAPIs := make([]orm.Score_partAPI, 0)

	// for each score_part, update fields from the database nullable fields
	for idx := range score_partDBs {
		score_partDB := &score_partDBs[idx]
		_ = score_partDB
		var score_partAPI orm.Score_partAPI

		// insertion point for updating fields
		score_partAPI.ID = score_partDB.ID
		score_partDB.CopyBasicFieldsToScore_part_WOP(&score_partAPI.Score_part_WOP)
		score_partAPI.Score_partPointersEncoding = score_partDB.Score_partPointersEncoding
		score_partAPIs = append(score_partAPIs, score_partAPI)
	}

	c.JSON(http.StatusOK, score_partAPIs)
}

// PostScore_part
//
// swagger:route POST /score_parts score_parts postScore_part
//
// Creates a score_part
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostScore_part(c *gin.Context) {

	mutexScore_part.Lock()
	defer mutexScore_part.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostScore_parts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_part.GetDB()

	// Validate input
	var input orm.Score_partAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create score_part
	score_partDB := orm.Score_partDB{}
	score_partDB.Score_partPointersEncoding = input.Score_partPointersEncoding
	score_partDB.CopyBasicFieldsFromScore_part_WOP(&input.Score_part_WOP)

	query := db.Create(&score_partDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoScore_part.CheckoutPhaseOneInstance(&score_partDB)
	score_part := backRepo.BackRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID]

	if score_part != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), score_part)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, score_partDB)
}

// GetScore_part
//
// swagger:route GET /score_parts/{ID} score_parts getScore_part
//
// Gets the details for a score_part.
//
// Responses:
// default: genericError
//
//	200: score_partDBResponse
func (controller *Controller) GetScore_part(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetScore_part", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_part.GetDB()

	// Get score_partDB in DB
	var score_partDB orm.Score_partDB
	if err := db.First(&score_partDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var score_partAPI orm.Score_partAPI
	score_partAPI.ID = score_partDB.ID
	score_partAPI.Score_partPointersEncoding = score_partDB.Score_partPointersEncoding
	score_partDB.CopyBasicFieldsToScore_part_WOP(&score_partAPI.Score_part_WOP)

	c.JSON(http.StatusOK, score_partAPI)
}

// UpdateScore_part
//
// swagger:route PATCH /score_parts/{ID} score_parts updateScore_part
//
// # Update a score_part
//
// Responses:
// default: genericError
//
//	200: score_partDBResponse
func (controller *Controller) UpdateScore_part(c *gin.Context) {

	mutexScore_part.Lock()
	defer mutexScore_part.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateScore_part", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_part.GetDB()

	// Validate input
	var input orm.Score_partAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var score_partDB orm.Score_partDB

	// fetch the score_part
	query := db.First(&score_partDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	score_partDB.CopyBasicFieldsFromScore_part_WOP(&input.Score_part_WOP)
	score_partDB.Score_partPointersEncoding = input.Score_partPointersEncoding

	query = db.Model(&score_partDB).Updates(score_partDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	score_partNew := new(models.Score_part)
	score_partDB.CopyBasicFieldsToScore_part(score_partNew)

	// redeem pointers
	score_partDB.DecodePointers(backRepo, score_partNew)

	// get stage instance from DB instance, and call callback function
	score_partOld := backRepo.BackRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID]
	if score_partOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), score_partOld, score_partNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the score_partDB
	c.JSON(http.StatusOK, score_partDB)
}

// DeleteScore_part
//
// swagger:route DELETE /score_parts/{ID} score_parts deleteScore_part
//
// # Delete a score_part
//
// default: genericError
//
//	200: score_partDBResponse
func (controller *Controller) DeleteScore_part(c *gin.Context) {

	mutexScore_part.Lock()
	defer mutexScore_part.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteScore_part", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoScore_part.GetDB()

	// Get model if exist
	var score_partDB orm.Score_partDB
	if err := db.First(&score_partDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&score_partDB)

	// get an instance (not staged) from DB instance, and call callback function
	score_partDeleted := new(models.Score_part)
	score_partDB.CopyBasicFieldsToScore_part(score_partDeleted)

	// get stage instance from DB instance, and call callback function
	score_partStaged := backRepo.BackRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID]
	if score_partStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), score_partStaged, score_partDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
