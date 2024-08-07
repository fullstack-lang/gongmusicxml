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
var __Sound__dummysDeclaration__ models.Sound
var __Sound_time__dummyDeclaration time.Duration

var mutexSound sync.Mutex

// An SoundID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSound updateSound deleteSound
type SoundID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SoundInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSound updateSound
type SoundInput struct {
	// The Sound to submit or modify
	// in: body
	Sound *orm.SoundAPI
}

// GetSounds
//
// swagger:route GET /sounds sounds getSounds
//
// # Get all sounds
//
// Responses:
// default: genericError
//
//	200: soundDBResponse
func (controller *Controller) GetSounds(c *gin.Context) {

	// source slice
	var soundDBs []orm.SoundDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSounds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSound.GetDB()

	query := db.Find(&soundDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	soundAPIs := make([]orm.SoundAPI, 0)

	// for each sound, update fields from the database nullable fields
	for idx := range soundDBs {
		soundDB := &soundDBs[idx]
		_ = soundDB
		var soundAPI orm.SoundAPI

		// insertion point for updating fields
		soundAPI.ID = soundDB.ID
		soundDB.CopyBasicFieldsToSound_WOP(&soundAPI.Sound_WOP)
		soundAPI.SoundPointersEncoding = soundDB.SoundPointersEncoding
		soundAPIs = append(soundAPIs, soundAPI)
	}

	c.JSON(http.StatusOK, soundAPIs)
}

// PostSound
//
// swagger:route POST /sounds sounds postSound
//
// Creates a sound
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSound(c *gin.Context) {

	mutexSound.Lock()
	defer mutexSound.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSounds", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSound.GetDB()

	// Validate input
	var input orm.SoundAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create sound
	soundDB := orm.SoundDB{}
	soundDB.SoundPointersEncoding = input.SoundPointersEncoding
	soundDB.CopyBasicFieldsFromSound_WOP(&input.Sound_WOP)

	query := db.Create(&soundDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSound.CheckoutPhaseOneInstance(&soundDB)
	sound := backRepo.BackRepoSound.Map_SoundDBID_SoundPtr[soundDB.ID]

	if sound != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), sound)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, soundDB)
}

// GetSound
//
// swagger:route GET /sounds/{ID} sounds getSound
//
// Gets the details for a sound.
//
// Responses:
// default: genericError
//
//	200: soundDBResponse
func (controller *Controller) GetSound(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSound", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSound.GetDB()

	// Get soundDB in DB
	var soundDB orm.SoundDB
	if err := db.First(&soundDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var soundAPI orm.SoundAPI
	soundAPI.ID = soundDB.ID
	soundAPI.SoundPointersEncoding = soundDB.SoundPointersEncoding
	soundDB.CopyBasicFieldsToSound_WOP(&soundAPI.Sound_WOP)

	c.JSON(http.StatusOK, soundAPI)
}

// UpdateSound
//
// swagger:route PATCH /sounds/{ID} sounds updateSound
//
// # Update a sound
//
// Responses:
// default: genericError
//
//	200: soundDBResponse
func (controller *Controller) UpdateSound(c *gin.Context) {

	mutexSound.Lock()
	defer mutexSound.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSound", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSound.GetDB()

	// Validate input
	var input orm.SoundAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var soundDB orm.SoundDB

	// fetch the sound
	query := db.First(&soundDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	soundDB.CopyBasicFieldsFromSound_WOP(&input.Sound_WOP)
	soundDB.SoundPointersEncoding = input.SoundPointersEncoding

	query = db.Model(&soundDB).Updates(soundDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	soundNew := new(models.Sound)
	soundDB.CopyBasicFieldsToSound(soundNew)

	// redeem pointers
	soundDB.DecodePointers(backRepo, soundNew)

	// get stage instance from DB instance, and call callback function
	soundOld := backRepo.BackRepoSound.Map_SoundDBID_SoundPtr[soundDB.ID]
	if soundOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), soundOld, soundNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the soundDB
	c.JSON(http.StatusOK, soundDB)
}

// DeleteSound
//
// swagger:route DELETE /sounds/{ID} sounds deleteSound
//
// # Delete a sound
//
// default: genericError
//
//	200: soundDBResponse
func (controller *Controller) DeleteSound(c *gin.Context) {

	mutexSound.Lock()
	defer mutexSound.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSound", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSound.GetDB()

	// Get model if exist
	var soundDB orm.SoundDB
	if err := db.First(&soundDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&soundDB)

	// get an instance (not staged) from DB instance, and call callback function
	soundDeleted := new(models.Sound)
	soundDB.CopyBasicFieldsToSound(soundDeleted)

	// get stage instance from DB instance, and call callback function
	soundStaged := backRepo.BackRepoSound.Map_SoundDBID_SoundPtr[soundDB.ID]
	if soundStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), soundStaged, soundDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
