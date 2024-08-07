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
var __Play__dummysDeclaration__ models.Play
var __Play_time__dummyDeclaration time.Duration

var mutexPlay sync.Mutex

// An PlayID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPlay updatePlay deletePlay
type PlayID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PlayInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPlay updatePlay
type PlayInput struct {
	// The Play to submit or modify
	// in: body
	Play *orm.PlayAPI
}

// GetPlays
//
// swagger:route GET /plays plays getPlays
//
// # Get all plays
//
// Responses:
// default: genericError
//
//	200: playDBResponse
func (controller *Controller) GetPlays(c *gin.Context) {

	// source slice
	var playDBs []orm.PlayDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPlays", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPlay.GetDB()

	query := db.Find(&playDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	playAPIs := make([]orm.PlayAPI, 0)

	// for each play, update fields from the database nullable fields
	for idx := range playDBs {
		playDB := &playDBs[idx]
		_ = playDB
		var playAPI orm.PlayAPI

		// insertion point for updating fields
		playAPI.ID = playDB.ID
		playDB.CopyBasicFieldsToPlay_WOP(&playAPI.Play_WOP)
		playAPI.PlayPointersEncoding = playDB.PlayPointersEncoding
		playAPIs = append(playAPIs, playAPI)
	}

	c.JSON(http.StatusOK, playAPIs)
}

// PostPlay
//
// swagger:route POST /plays plays postPlay
//
// Creates a play
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPlay(c *gin.Context) {

	mutexPlay.Lock()
	defer mutexPlay.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPlays", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPlay.GetDB()

	// Validate input
	var input orm.PlayAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create play
	playDB := orm.PlayDB{}
	playDB.PlayPointersEncoding = input.PlayPointersEncoding
	playDB.CopyBasicFieldsFromPlay_WOP(&input.Play_WOP)

	query := db.Create(&playDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPlay.CheckoutPhaseOneInstance(&playDB)
	play := backRepo.BackRepoPlay.Map_PlayDBID_PlayPtr[playDB.ID]

	if play != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), play)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, playDB)
}

// GetPlay
//
// swagger:route GET /plays/{ID} plays getPlay
//
// Gets the details for a play.
//
// Responses:
// default: genericError
//
//	200: playDBResponse
func (controller *Controller) GetPlay(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPlay", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPlay.GetDB()

	// Get playDB in DB
	var playDB orm.PlayDB
	if err := db.First(&playDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var playAPI orm.PlayAPI
	playAPI.ID = playDB.ID
	playAPI.PlayPointersEncoding = playDB.PlayPointersEncoding
	playDB.CopyBasicFieldsToPlay_WOP(&playAPI.Play_WOP)

	c.JSON(http.StatusOK, playAPI)
}

// UpdatePlay
//
// swagger:route PATCH /plays/{ID} plays updatePlay
//
// # Update a play
//
// Responses:
// default: genericError
//
//	200: playDBResponse
func (controller *Controller) UpdatePlay(c *gin.Context) {

	mutexPlay.Lock()
	defer mutexPlay.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePlay", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPlay.GetDB()

	// Validate input
	var input orm.PlayAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var playDB orm.PlayDB

	// fetch the play
	query := db.First(&playDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	playDB.CopyBasicFieldsFromPlay_WOP(&input.Play_WOP)
	playDB.PlayPointersEncoding = input.PlayPointersEncoding

	query = db.Model(&playDB).Updates(playDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	playNew := new(models.Play)
	playDB.CopyBasicFieldsToPlay(playNew)

	// redeem pointers
	playDB.DecodePointers(backRepo, playNew)

	// get stage instance from DB instance, and call callback function
	playOld := backRepo.BackRepoPlay.Map_PlayDBID_PlayPtr[playDB.ID]
	if playOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), playOld, playNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the playDB
	c.JSON(http.StatusOK, playDB)
}

// DeletePlay
//
// swagger:route DELETE /plays/{ID} plays deletePlay
//
// # Delete a play
//
// default: genericError
//
//	200: playDBResponse
func (controller *Controller) DeletePlay(c *gin.Context) {

	mutexPlay.Lock()
	defer mutexPlay.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePlay", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPlay.GetDB()

	// Get model if exist
	var playDB orm.PlayDB
	if err := db.First(&playDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&playDB)

	// get an instance (not staged) from DB instance, and call callback function
	playDeleted := new(models.Play)
	playDB.CopyBasicFieldsToPlay(playDeleted)

	// get stage instance from DB instance, and call callback function
	playStaged := backRepo.BackRepoPlay.Map_PlayDBID_PlayPtr[playDB.ID]
	if playStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), playStaged, playDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
