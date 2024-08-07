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
var __Effect__dummysDeclaration__ models.Effect
var __Effect_time__dummyDeclaration time.Duration

var mutexEffect sync.Mutex

// An EffectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEffect updateEffect deleteEffect
type EffectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EffectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEffect updateEffect
type EffectInput struct {
	// The Effect to submit or modify
	// in: body
	Effect *orm.EffectAPI
}

// GetEffects
//
// swagger:route GET /effects effects getEffects
//
// # Get all effects
//
// Responses:
// default: genericError
//
//	200: effectDBResponse
func (controller *Controller) GetEffects(c *gin.Context) {

	// source slice
	var effectDBs []orm.EffectDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEffects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEffect.GetDB()

	query := db.Find(&effectDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	effectAPIs := make([]orm.EffectAPI, 0)

	// for each effect, update fields from the database nullable fields
	for idx := range effectDBs {
		effectDB := &effectDBs[idx]
		_ = effectDB
		var effectAPI orm.EffectAPI

		// insertion point for updating fields
		effectAPI.ID = effectDB.ID
		effectDB.CopyBasicFieldsToEffect_WOP(&effectAPI.Effect_WOP)
		effectAPI.EffectPointersEncoding = effectDB.EffectPointersEncoding
		effectAPIs = append(effectAPIs, effectAPI)
	}

	c.JSON(http.StatusOK, effectAPIs)
}

// PostEffect
//
// swagger:route POST /effects effects postEffect
//
// Creates a effect
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEffect(c *gin.Context) {

	mutexEffect.Lock()
	defer mutexEffect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEffects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEffect.GetDB()

	// Validate input
	var input orm.EffectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create effect
	effectDB := orm.EffectDB{}
	effectDB.EffectPointersEncoding = input.EffectPointersEncoding
	effectDB.CopyBasicFieldsFromEffect_WOP(&input.Effect_WOP)

	query := db.Create(&effectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEffect.CheckoutPhaseOneInstance(&effectDB)
	effect := backRepo.BackRepoEffect.Map_EffectDBID_EffectPtr[effectDB.ID]

	if effect != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), effect)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, effectDB)
}

// GetEffect
//
// swagger:route GET /effects/{ID} effects getEffect
//
// Gets the details for a effect.
//
// Responses:
// default: genericError
//
//	200: effectDBResponse
func (controller *Controller) GetEffect(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEffect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEffect.GetDB()

	// Get effectDB in DB
	var effectDB orm.EffectDB
	if err := db.First(&effectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var effectAPI orm.EffectAPI
	effectAPI.ID = effectDB.ID
	effectAPI.EffectPointersEncoding = effectDB.EffectPointersEncoding
	effectDB.CopyBasicFieldsToEffect_WOP(&effectAPI.Effect_WOP)

	c.JSON(http.StatusOK, effectAPI)
}

// UpdateEffect
//
// swagger:route PATCH /effects/{ID} effects updateEffect
//
// # Update a effect
//
// Responses:
// default: genericError
//
//	200: effectDBResponse
func (controller *Controller) UpdateEffect(c *gin.Context) {

	mutexEffect.Lock()
	defer mutexEffect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEffect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEffect.GetDB()

	// Validate input
	var input orm.EffectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var effectDB orm.EffectDB

	// fetch the effect
	query := db.First(&effectDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	effectDB.CopyBasicFieldsFromEffect_WOP(&input.Effect_WOP)
	effectDB.EffectPointersEncoding = input.EffectPointersEncoding

	query = db.Model(&effectDB).Updates(effectDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	effectNew := new(models.Effect)
	effectDB.CopyBasicFieldsToEffect(effectNew)

	// redeem pointers
	effectDB.DecodePointers(backRepo, effectNew)

	// get stage instance from DB instance, and call callback function
	effectOld := backRepo.BackRepoEffect.Map_EffectDBID_EffectPtr[effectDB.ID]
	if effectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), effectOld, effectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the effectDB
	c.JSON(http.StatusOK, effectDB)
}

// DeleteEffect
//
// swagger:route DELETE /effects/{ID} effects deleteEffect
//
// # Delete a effect
//
// default: genericError
//
//	200: effectDBResponse
func (controller *Controller) DeleteEffect(c *gin.Context) {

	mutexEffect.Lock()
	defer mutexEffect.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEffect", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoEffect.GetDB()

	// Get model if exist
	var effectDB orm.EffectDB
	if err := db.First(&effectDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&effectDB)

	// get an instance (not staged) from DB instance, and call callback function
	effectDeleted := new(models.Effect)
	effectDB.CopyBasicFieldsToEffect(effectDeleted)

	// get stage instance from DB instance, and call callback function
	effectStaged := backRepo.BackRepoEffect.Map_EffectDBID_EffectPtr[effectDB.ID]
	if effectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), effectStaged, effectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
