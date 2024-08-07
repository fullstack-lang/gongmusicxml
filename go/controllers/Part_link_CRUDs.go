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
var __Part_link__dummysDeclaration__ models.Part_link
var __Part_link_time__dummyDeclaration time.Duration

var mutexPart_link sync.Mutex

// An Part_linkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPart_link updatePart_link deletePart_link
type Part_linkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Part_linkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPart_link updatePart_link
type Part_linkInput struct {
	// The Part_link to submit or modify
	// in: body
	Part_link *orm.Part_linkAPI
}

// GetPart_links
//
// swagger:route GET /part_links part_links getPart_links
//
// # Get all part_links
//
// Responses:
// default: genericError
//
//	200: part_linkDBResponse
func (controller *Controller) GetPart_links(c *gin.Context) {

	// source slice
	var part_linkDBs []orm.Part_linkDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_links", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_link.GetDB()

	query := db.Find(&part_linkDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	part_linkAPIs := make([]orm.Part_linkAPI, 0)

	// for each part_link, update fields from the database nullable fields
	for idx := range part_linkDBs {
		part_linkDB := &part_linkDBs[idx]
		_ = part_linkDB
		var part_linkAPI orm.Part_linkAPI

		// insertion point for updating fields
		part_linkAPI.ID = part_linkDB.ID
		part_linkDB.CopyBasicFieldsToPart_link_WOP(&part_linkAPI.Part_link_WOP)
		part_linkAPI.Part_linkPointersEncoding = part_linkDB.Part_linkPointersEncoding
		part_linkAPIs = append(part_linkAPIs, part_linkAPI)
	}

	c.JSON(http.StatusOK, part_linkAPIs)
}

// PostPart_link
//
// swagger:route POST /part_links part_links postPart_link
//
// Creates a part_link
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPart_link(c *gin.Context) {

	mutexPart_link.Lock()
	defer mutexPart_link.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPart_links", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_link.GetDB()

	// Validate input
	var input orm.Part_linkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create part_link
	part_linkDB := orm.Part_linkDB{}
	part_linkDB.Part_linkPointersEncoding = input.Part_linkPointersEncoding
	part_linkDB.CopyBasicFieldsFromPart_link_WOP(&input.Part_link_WOP)

	query := db.Create(&part_linkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPart_link.CheckoutPhaseOneInstance(&part_linkDB)
	part_link := backRepo.BackRepoPart_link.Map_Part_linkDBID_Part_linkPtr[part_linkDB.ID]

	if part_link != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), part_link)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, part_linkDB)
}

// GetPart_link
//
// swagger:route GET /part_links/{ID} part_links getPart_link
//
// Gets the details for a part_link.
//
// Responses:
// default: genericError
//
//	200: part_linkDBResponse
func (controller *Controller) GetPart_link(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPart_link", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_link.GetDB()

	// Get part_linkDB in DB
	var part_linkDB orm.Part_linkDB
	if err := db.First(&part_linkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var part_linkAPI orm.Part_linkAPI
	part_linkAPI.ID = part_linkDB.ID
	part_linkAPI.Part_linkPointersEncoding = part_linkDB.Part_linkPointersEncoding
	part_linkDB.CopyBasicFieldsToPart_link_WOP(&part_linkAPI.Part_link_WOP)

	c.JSON(http.StatusOK, part_linkAPI)
}

// UpdatePart_link
//
// swagger:route PATCH /part_links/{ID} part_links updatePart_link
//
// # Update a part_link
//
// Responses:
// default: genericError
//
//	200: part_linkDBResponse
func (controller *Controller) UpdatePart_link(c *gin.Context) {

	mutexPart_link.Lock()
	defer mutexPart_link.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePart_link", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_link.GetDB()

	// Validate input
	var input orm.Part_linkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var part_linkDB orm.Part_linkDB

	// fetch the part_link
	query := db.First(&part_linkDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	part_linkDB.CopyBasicFieldsFromPart_link_WOP(&input.Part_link_WOP)
	part_linkDB.Part_linkPointersEncoding = input.Part_linkPointersEncoding

	query = db.Model(&part_linkDB).Updates(part_linkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	part_linkNew := new(models.Part_link)
	part_linkDB.CopyBasicFieldsToPart_link(part_linkNew)

	// redeem pointers
	part_linkDB.DecodePointers(backRepo, part_linkNew)

	// get stage instance from DB instance, and call callback function
	part_linkOld := backRepo.BackRepoPart_link.Map_Part_linkDBID_Part_linkPtr[part_linkDB.ID]
	if part_linkOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), part_linkOld, part_linkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the part_linkDB
	c.JSON(http.StatusOK, part_linkDB)
}

// DeletePart_link
//
// swagger:route DELETE /part_links/{ID} part_links deletePart_link
//
// # Delete a part_link
//
// default: genericError
//
//	200: part_linkDBResponse
func (controller *Controller) DeletePart_link(c *gin.Context) {

	mutexPart_link.Lock()
	defer mutexPart_link.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePart_link", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPart_link.GetDB()

	// Get model if exist
	var part_linkDB orm.Part_linkDB
	if err := db.First(&part_linkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&part_linkDB)

	// get an instance (not staged) from DB instance, and call callback function
	part_linkDeleted := new(models.Part_link)
	part_linkDB.CopyBasicFieldsToPart_link(part_linkDeleted)

	// get stage instance from DB instance, and call callback function
	part_linkStaged := backRepo.BackRepoPart_link.Map_Part_linkDBID_Part_linkPtr[part_linkDB.ID]
	if part_linkStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), part_linkStaged, part_linkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
