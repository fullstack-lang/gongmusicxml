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
var __Level__dummysDeclaration__ models.Level
var __Level_time__dummyDeclaration time.Duration

var mutexLevel sync.Mutex

// An LevelID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getLevel updateLevel deleteLevel
type LevelID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// LevelInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postLevel updateLevel
type LevelInput struct {
	// The Level to submit or modify
	// in: body
	Level *orm.LevelAPI
}

// GetLevels
//
// swagger:route GET /levels levels getLevels
//
// # Get all levels
//
// Responses:
// default: genericError
//
//	200: levelDBResponse
func (controller *Controller) GetLevels(c *gin.Context) {

	// source slice
	var levelDBs []orm.LevelDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLevels", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLevel.GetDB()

	query := db.Find(&levelDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	levelAPIs := make([]orm.LevelAPI, 0)

	// for each level, update fields from the database nullable fields
	for idx := range levelDBs {
		levelDB := &levelDBs[idx]
		_ = levelDB
		var levelAPI orm.LevelAPI

		// insertion point for updating fields
		levelAPI.ID = levelDB.ID
		levelDB.CopyBasicFieldsToLevel_WOP(&levelAPI.Level_WOP)
		levelAPI.LevelPointersEncoding = levelDB.LevelPointersEncoding
		levelAPIs = append(levelAPIs, levelAPI)
	}

	c.JSON(http.StatusOK, levelAPIs)
}

// PostLevel
//
// swagger:route POST /levels levels postLevel
//
// Creates a level
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostLevel(c *gin.Context) {

	mutexLevel.Lock()
	defer mutexLevel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostLevels", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLevel.GetDB()

	// Validate input
	var input orm.LevelAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create level
	levelDB := orm.LevelDB{}
	levelDB.LevelPointersEncoding = input.LevelPointersEncoding
	levelDB.CopyBasicFieldsFromLevel_WOP(&input.Level_WOP)

	query := db.Create(&levelDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoLevel.CheckoutPhaseOneInstance(&levelDB)
	level := backRepo.BackRepoLevel.Map_LevelDBID_LevelPtr[levelDB.ID]

	if level != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), level)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, levelDB)
}

// GetLevel
//
// swagger:route GET /levels/{ID} levels getLevel
//
// Gets the details for a level.
//
// Responses:
// default: genericError
//
//	200: levelDBResponse
func (controller *Controller) GetLevel(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLevel", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLevel.GetDB()

	// Get levelDB in DB
	var levelDB orm.LevelDB
	if err := db.First(&levelDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var levelAPI orm.LevelAPI
	levelAPI.ID = levelDB.ID
	levelAPI.LevelPointersEncoding = levelDB.LevelPointersEncoding
	levelDB.CopyBasicFieldsToLevel_WOP(&levelAPI.Level_WOP)

	c.JSON(http.StatusOK, levelAPI)
}

// UpdateLevel
//
// swagger:route PATCH /levels/{ID} levels updateLevel
//
// # Update a level
//
// Responses:
// default: genericError
//
//	200: levelDBResponse
func (controller *Controller) UpdateLevel(c *gin.Context) {

	mutexLevel.Lock()
	defer mutexLevel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateLevel", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLevel.GetDB()

	// Validate input
	var input orm.LevelAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var levelDB orm.LevelDB

	// fetch the level
	query := db.First(&levelDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	levelDB.CopyBasicFieldsFromLevel_WOP(&input.Level_WOP)
	levelDB.LevelPointersEncoding = input.LevelPointersEncoding

	query = db.Model(&levelDB).Updates(levelDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	levelNew := new(models.Level)
	levelDB.CopyBasicFieldsToLevel(levelNew)

	// redeem pointers
	levelDB.DecodePointers(backRepo, levelNew)

	// get stage instance from DB instance, and call callback function
	levelOld := backRepo.BackRepoLevel.Map_LevelDBID_LevelPtr[levelDB.ID]
	if levelOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), levelOld, levelNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the levelDB
	c.JSON(http.StatusOK, levelDB)
}

// DeleteLevel
//
// swagger:route DELETE /levels/{ID} levels deleteLevel
//
// # Delete a level
//
// default: genericError
//
//	200: levelDBResponse
func (controller *Controller) DeleteLevel(c *gin.Context) {

	mutexLevel.Lock()
	defer mutexLevel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteLevel", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoLevel.GetDB()

	// Get model if exist
	var levelDB orm.LevelDB
	if err := db.First(&levelDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&levelDB)

	// get an instance (not staged) from DB instance, and call callback function
	levelDeleted := new(models.Level)
	levelDB.CopyBasicFieldsToLevel(levelDeleted)

	// get stage instance from DB instance, and call callback function
	levelStaged := backRepo.BackRepoLevel.Map_LevelDBID_LevelPtr[levelDB.ID]
	if levelStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), levelStaged, levelDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
