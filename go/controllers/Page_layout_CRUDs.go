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
var __Page_layout__dummysDeclaration__ models.Page_layout
var __Page_layout_time__dummyDeclaration time.Duration

var mutexPage_layout sync.Mutex

// An Page_layoutID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPage_layout updatePage_layout deletePage_layout
type Page_layoutID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Page_layoutInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPage_layout updatePage_layout
type Page_layoutInput struct {
	// The Page_layout to submit or modify
	// in: body
	Page_layout *orm.Page_layoutAPI
}

// GetPage_layouts
//
// swagger:route GET /page_layouts page_layouts getPage_layouts
//
// # Get all page_layouts
//
// Responses:
// default: genericError
//
//	200: page_layoutDBResponse
func (controller *Controller) GetPage_layouts(c *gin.Context) {

	// source slice
	var page_layoutDBs []orm.Page_layoutDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPage_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_layout.GetDB()

	query := db.Find(&page_layoutDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	page_layoutAPIs := make([]orm.Page_layoutAPI, 0)

	// for each page_layout, update fields from the database nullable fields
	for idx := range page_layoutDBs {
		page_layoutDB := &page_layoutDBs[idx]
		_ = page_layoutDB
		var page_layoutAPI orm.Page_layoutAPI

		// insertion point for updating fields
		page_layoutAPI.ID = page_layoutDB.ID
		page_layoutDB.CopyBasicFieldsToPage_layout_WOP(&page_layoutAPI.Page_layout_WOP)
		page_layoutAPI.Page_layoutPointersEncoding = page_layoutDB.Page_layoutPointersEncoding
		page_layoutAPIs = append(page_layoutAPIs, page_layoutAPI)
	}

	c.JSON(http.StatusOK, page_layoutAPIs)
}

// PostPage_layout
//
// swagger:route POST /page_layouts page_layouts postPage_layout
//
// Creates a page_layout
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPage_layout(c *gin.Context) {

	mutexPage_layout.Lock()
	defer mutexPage_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPage_layouts", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_layout.GetDB()

	// Validate input
	var input orm.Page_layoutAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create page_layout
	page_layoutDB := orm.Page_layoutDB{}
	page_layoutDB.Page_layoutPointersEncoding = input.Page_layoutPointersEncoding
	page_layoutDB.CopyBasicFieldsFromPage_layout_WOP(&input.Page_layout_WOP)

	query := db.Create(&page_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPage_layout.CheckoutPhaseOneInstance(&page_layoutDB)
	page_layout := backRepo.BackRepoPage_layout.Map_Page_layoutDBID_Page_layoutPtr[page_layoutDB.ID]

	if page_layout != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), page_layout)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, page_layoutDB)
}

// GetPage_layout
//
// swagger:route GET /page_layouts/{ID} page_layouts getPage_layout
//
// Gets the details for a page_layout.
//
// Responses:
// default: genericError
//
//	200: page_layoutDBResponse
func (controller *Controller) GetPage_layout(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPage_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_layout.GetDB()

	// Get page_layoutDB in DB
	var page_layoutDB orm.Page_layoutDB
	if err := db.First(&page_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var page_layoutAPI orm.Page_layoutAPI
	page_layoutAPI.ID = page_layoutDB.ID
	page_layoutAPI.Page_layoutPointersEncoding = page_layoutDB.Page_layoutPointersEncoding
	page_layoutDB.CopyBasicFieldsToPage_layout_WOP(&page_layoutAPI.Page_layout_WOP)

	c.JSON(http.StatusOK, page_layoutAPI)
}

// UpdatePage_layout
//
// swagger:route PATCH /page_layouts/{ID} page_layouts updatePage_layout
//
// # Update a page_layout
//
// Responses:
// default: genericError
//
//	200: page_layoutDBResponse
func (controller *Controller) UpdatePage_layout(c *gin.Context) {

	mutexPage_layout.Lock()
	defer mutexPage_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePage_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_layout.GetDB()

	// Validate input
	var input orm.Page_layoutAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var page_layoutDB orm.Page_layoutDB

	// fetch the page_layout
	query := db.First(&page_layoutDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	page_layoutDB.CopyBasicFieldsFromPage_layout_WOP(&input.Page_layout_WOP)
	page_layoutDB.Page_layoutPointersEncoding = input.Page_layoutPointersEncoding

	query = db.Model(&page_layoutDB).Updates(page_layoutDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	page_layoutNew := new(models.Page_layout)
	page_layoutDB.CopyBasicFieldsToPage_layout(page_layoutNew)

	// redeem pointers
	page_layoutDB.DecodePointers(backRepo, page_layoutNew)

	// get stage instance from DB instance, and call callback function
	page_layoutOld := backRepo.BackRepoPage_layout.Map_Page_layoutDBID_Page_layoutPtr[page_layoutDB.ID]
	if page_layoutOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), page_layoutOld, page_layoutNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the page_layoutDB
	c.JSON(http.StatusOK, page_layoutDB)
}

// DeletePage_layout
//
// swagger:route DELETE /page_layouts/{ID} page_layouts deletePage_layout
//
// # Delete a page_layout
//
// default: genericError
//
//	200: page_layoutDBResponse
func (controller *Controller) DeletePage_layout(c *gin.Context) {

	mutexPage_layout.Lock()
	defer mutexPage_layout.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePage_layout", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPage_layout.GetDB()

	// Get model if exist
	var page_layoutDB orm.Page_layoutDB
	if err := db.First(&page_layoutDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&page_layoutDB)

	// get an instance (not staged) from DB instance, and call callback function
	page_layoutDeleted := new(models.Page_layout)
	page_layoutDB.CopyBasicFieldsToPage_layout(page_layoutDeleted)

	// get stage instance from DB instance, and call callback function
	page_layoutStaged := backRepo.BackRepoPage_layout.Map_Page_layoutDBID_Page_layoutPtr[page_layoutDB.ID]
	if page_layoutStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), page_layoutStaged, page_layoutDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
