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
var __Part_clef__dummysDeclaration__ models.Part_clef
var __Part_clef_time__dummyDeclaration time.Duration

var mutexPart_clef sync.Mutex

// An Part_clefID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPart_clef updatePart_clef deletePart_clef
type Part_clefID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Part_clefInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPart_clef updatePart_clef
type Part_clefInput struct {
	// The Part_clef to submit or modify
	// in: body
	Part_clef *orm.Part_clefAPI
}

// GetPart_clefs
//
// swagger:route GET /part_clefs part_clefs getPart_clefs
//
// # Get all part_clefs
//
// Responses:
// default: genericError
//
//	200: part_clefDBResponse
func (controller *Controller) GetPart_clefs(c *gin.Context) {

	// source slice
	var part_clefDBs []orm.Part_clefDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_clefs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_clef.GetDB()

	query := db.Find(&part_clefDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	part_clefAPIs := make([]orm.Part_clefAPI, 0)

	// for each part_clef, update fields from the database nullable fields
	for idx := range part_clefDBs {
		part_clefDB := &part_clefDBs[idx]
		_ = part_clefDB
		var part_clefAPI orm.Part_clefAPI

		// insertion point for updating fields
		part_clefAPI.ID = part_clefDB.ID
		part_clefDB.CopyBasicFieldsToPart_clef_WOP(&part_clefAPI.Part_clef_WOP)
		part_clefAPI.Part_clefPointersEncoding = part_clefDB.Part_clefPointersEncoding
		part_clefAPIs = append(part_clefAPIs, part_clefAPI)
	}

	c.JSON(http.StatusOK, part_clefAPIs)
}

// PostPart_clef
//
// swagger:route POST /part_clefs part_clefs postPart_clef
//
// Creates a part_clef
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPart_clef(c *gin.Context) {

	mutexPart_clef.Lock()
	defer mutexPart_clef.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPart_clefs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_clef.GetDB()

	// Validate input
	var input orm.Part_clefAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create part_clef
	part_clefDB := orm.Part_clefDB{}
	part_clefDB.Part_clefPointersEncoding = input.Part_clefPointersEncoding
	part_clefDB.CopyBasicFieldsFromPart_clef_WOP(&input.Part_clef_WOP)

	query := db.Create(&part_clefDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPart_clef.CheckoutPhaseOneInstance(&part_clefDB)
	part_clef := backRepo.BackRepoPart_clef.Map_Part_clefDBID_Part_clefPtr[part_clefDB.ID]

	if part_clef != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), part_clef)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, part_clefDB)
}

// GetPart_clef
//
// swagger:route GET /part_clefs/{ID} part_clefs getPart_clef
//
// Gets the details for a part_clef.
//
// Responses:
// default: genericError
//
//	200: part_clefDBResponse
func (controller *Controller) GetPart_clef(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_clef", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_clef.GetDB()

	// Get part_clefDB in DB
	var part_clefDB orm.Part_clefDB
	if err := db.First(&part_clefDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var part_clefAPI orm.Part_clefAPI
	part_clefAPI.ID = part_clefDB.ID
	part_clefAPI.Part_clefPointersEncoding = part_clefDB.Part_clefPointersEncoding
	part_clefDB.CopyBasicFieldsToPart_clef_WOP(&part_clefAPI.Part_clef_WOP)

	c.JSON(http.StatusOK, part_clefAPI)
}

// UpdatePart_clef
//
// swagger:route PATCH /part_clefs/{ID} part_clefs updatePart_clef
//
// # Update a part_clef
//
// Responses:
// default: genericError
//
//	200: part_clefDBResponse
func (controller *Controller) UpdatePart_clef(c *gin.Context) {

	mutexPart_clef.Lock()
	defer mutexPart_clef.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePart_clef", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_clef.GetDB()

	// Validate input
	var input orm.Part_clefAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var part_clefDB orm.Part_clefDB

	// fetch the part_clef
	query := db.First(&part_clefDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	part_clefDB.CopyBasicFieldsFromPart_clef_WOP(&input.Part_clef_WOP)
	part_clefDB.Part_clefPointersEncoding = input.Part_clefPointersEncoding

	query = db.Model(&part_clefDB).Updates(part_clefDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	part_clefNew := new(models.Part_clef)
	part_clefDB.CopyBasicFieldsToPart_clef(part_clefNew)

	// redeem pointers
	part_clefDB.DecodePointers(backRepo, part_clefNew)

	// get stage instance from DB instance, and call callback function
	part_clefOld := backRepo.BackRepoPart_clef.Map_Part_clefDBID_Part_clefPtr[part_clefDB.ID]
	if part_clefOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), part_clefOld, part_clefNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the part_clefDB
	c.JSON(http.StatusOK, part_clefDB)
}

// DeletePart_clef
//
// swagger:route DELETE /part_clefs/{ID} part_clefs deletePart_clef
//
// # Delete a part_clef
//
// default: genericError
//
//	200: part_clefDBResponse
func (controller *Controller) DeletePart_clef(c *gin.Context) {

	mutexPart_clef.Lock()
	defer mutexPart_clef.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePart_clef", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_clef.GetDB()

	// Get model if exist
	var part_clefDB orm.Part_clefDB
	if err := db.First(&part_clefDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&part_clefDB)

	// get an instance (not staged) from DB instance, and call callback function
	part_clefDeleted := new(models.Part_clef)
	part_clefDB.CopyBasicFieldsToPart_clef(part_clefDeleted)

	// get stage instance from DB instance, and call callback function
	part_clefStaged := backRepo.BackRepoPart_clef.Map_Part_clefDBID_Part_clefPtr[part_clefDB.ID]
	if part_clefStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), part_clefStaged, part_clefDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
