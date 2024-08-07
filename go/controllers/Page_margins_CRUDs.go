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
var __Page_margins__dummysDeclaration__ models.Page_margins
var __Page_margins_time__dummyDeclaration time.Duration

var mutexPage_margins sync.Mutex

// An Page_marginsID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPage_margins updatePage_margins deletePage_margins
type Page_marginsID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Page_marginsInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPage_margins updatePage_margins
type Page_marginsInput struct {
	// The Page_margins to submit or modify
	// in: body
	Page_margins *orm.Page_marginsAPI
}

// GetPage_marginss
//
// swagger:route GET /page_marginss page_marginss getPage_marginss
//
// # Get all page_marginss
//
// Responses:
// default: genericError
//
//	200: page_marginsDBResponse
func (controller *Controller) GetPage_marginss(c *gin.Context) {

	// source slice
	var page_marginsDBs []orm.Page_marginsDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPage_marginss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_margins.GetDB()

	query := db.Find(&page_marginsDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	page_marginsAPIs := make([]orm.Page_marginsAPI, 0)

	// for each page_margins, update fields from the database nullable fields
	for idx := range page_marginsDBs {
		page_marginsDB := &page_marginsDBs[idx]
		_ = page_marginsDB
		var page_marginsAPI orm.Page_marginsAPI

		// insertion point for updating fields
		page_marginsAPI.ID = page_marginsDB.ID
		page_marginsDB.CopyBasicFieldsToPage_margins_WOP(&page_marginsAPI.Page_margins_WOP)
		page_marginsAPI.Page_marginsPointersEncoding = page_marginsDB.Page_marginsPointersEncoding
		page_marginsAPIs = append(page_marginsAPIs, page_marginsAPI)
	}

	c.JSON(http.StatusOK, page_marginsAPIs)
}

// PostPage_margins
//
// swagger:route POST /page_marginss page_marginss postPage_margins
//
// Creates a page_margins
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPage_margins(c *gin.Context) {

	mutexPage_margins.Lock()
	defer mutexPage_margins.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPage_marginss", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_margins.GetDB()

	// Validate input
	var input orm.Page_marginsAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create page_margins
	page_marginsDB := orm.Page_marginsDB{}
	page_marginsDB.Page_marginsPointersEncoding = input.Page_marginsPointersEncoding
	page_marginsDB.CopyBasicFieldsFromPage_margins_WOP(&input.Page_margins_WOP)

	query := db.Create(&page_marginsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPage_margins.CheckoutPhaseOneInstance(&page_marginsDB)
	page_margins := backRepo.BackRepoPage_margins.Map_Page_marginsDBID_Page_marginsPtr[page_marginsDB.ID]

	if page_margins != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), page_margins)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, page_marginsDB)
}

// GetPage_margins
//
// swagger:route GET /page_marginss/{ID} page_marginss getPage_margins
//
// Gets the details for a page_margins.
//
// Responses:
// default: genericError
//
//	200: page_marginsDBResponse
func (controller *Controller) GetPage_margins(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPage_margins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_margins.GetDB()

	// Get page_marginsDB in DB
	var page_marginsDB orm.Page_marginsDB
	if err := db.First(&page_marginsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var page_marginsAPI orm.Page_marginsAPI
	page_marginsAPI.ID = page_marginsDB.ID
	page_marginsAPI.Page_marginsPointersEncoding = page_marginsDB.Page_marginsPointersEncoding
	page_marginsDB.CopyBasicFieldsToPage_margins_WOP(&page_marginsAPI.Page_margins_WOP)

	c.JSON(http.StatusOK, page_marginsAPI)
}

// UpdatePage_margins
//
// swagger:route PATCH /page_marginss/{ID} page_marginss updatePage_margins
//
// # Update a page_margins
//
// Responses:
// default: genericError
//
//	200: page_marginsDBResponse
func (controller *Controller) UpdatePage_margins(c *gin.Context) {

	mutexPage_margins.Lock()
	defer mutexPage_margins.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePage_margins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_margins.GetDB()

	// Validate input
	var input orm.Page_marginsAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var page_marginsDB orm.Page_marginsDB

	// fetch the page_margins
	query := db.First(&page_marginsDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	page_marginsDB.CopyBasicFieldsFromPage_margins_WOP(&input.Page_margins_WOP)
	page_marginsDB.Page_marginsPointersEncoding = input.Page_marginsPointersEncoding

	query = db.Model(&page_marginsDB).Updates(page_marginsDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	page_marginsNew := new(models.Page_margins)
	page_marginsDB.CopyBasicFieldsToPage_margins(page_marginsNew)

	// redeem pointers
	page_marginsDB.DecodePointers(backRepo, page_marginsNew)

	// get stage instance from DB instance, and call callback function
	page_marginsOld := backRepo.BackRepoPage_margins.Map_Page_marginsDBID_Page_marginsPtr[page_marginsDB.ID]
	if page_marginsOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), page_marginsOld, page_marginsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the page_marginsDB
	c.JSON(http.StatusOK, page_marginsDB)
}

// DeletePage_margins
//
// swagger:route DELETE /page_marginss/{ID} page_marginss deletePage_margins
//
// # Delete a page_margins
//
// default: genericError
//
//	200: page_marginsDBResponse
func (controller *Controller) DeletePage_margins(c *gin.Context) {

	mutexPage_margins.Lock()
	defer mutexPage_margins.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePage_margins", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_margins.GetDB()

	// Get model if exist
	var page_marginsDB orm.Page_marginsDB
	if err := db.First(&page_marginsDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&page_marginsDB)

	// get an instance (not staged) from DB instance, and call callback function
	page_marginsDeleted := new(models.Page_margins)
	page_marginsDB.CopyBasicFieldsToPage_margins(page_marginsDeleted)

	// get stage instance from DB instance, and call callback function
	page_marginsStaged := backRepo.BackRepoPage_margins.Map_Page_marginsDBID_Page_marginsPtr[page_marginsDB.ID]
	if page_marginsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), page_marginsStaged, page_marginsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
