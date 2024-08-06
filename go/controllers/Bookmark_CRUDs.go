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
var __Bookmark__dummysDeclaration__ models.Bookmark
var __Bookmark_time__dummyDeclaration time.Duration

var mutexBookmark sync.Mutex

// An BookmarkID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBookmark updateBookmark deleteBookmark
type BookmarkID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// BookmarkInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBookmark updateBookmark
type BookmarkInput struct {
	// The Bookmark to submit or modify
	// in: body
	Bookmark *orm.BookmarkAPI
}

// GetBookmarks
//
// swagger:route GET /bookmarks bookmarks getBookmarks
//
// # Get all bookmarks
//
// Responses:
// default: genericError
//
//	200: bookmarkDBResponse
func (controller *Controller) GetBookmarks(c *gin.Context) {

	// source slice
	var bookmarkDBs []orm.BookmarkDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookmarks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookmark.GetDB()

	query := db.Find(&bookmarkDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	bookmarkAPIs := make([]orm.BookmarkAPI, 0)

	// for each bookmark, update fields from the database nullable fields
	for idx := range bookmarkDBs {
		bookmarkDB := &bookmarkDBs[idx]
		_ = bookmarkDB
		var bookmarkAPI orm.BookmarkAPI

		// insertion point for updating fields
		bookmarkAPI.ID = bookmarkDB.ID
		bookmarkDB.CopyBasicFieldsToBookmark_WOP(&bookmarkAPI.Bookmark_WOP)
		bookmarkAPI.BookmarkPointersEncoding = bookmarkDB.BookmarkPointersEncoding
		bookmarkAPIs = append(bookmarkAPIs, bookmarkAPI)
	}

	c.JSON(http.StatusOK, bookmarkAPIs)
}

// PostBookmark
//
// swagger:route POST /bookmarks bookmarks postBookmark
//
// Creates a bookmark
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBookmark(c *gin.Context) {

	mutexBookmark.Lock()
	defer mutexBookmark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBookmarks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookmark.GetDB()

	// Validate input
	var input orm.BookmarkAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create bookmark
	bookmarkDB := orm.BookmarkDB{}
	bookmarkDB.BookmarkPointersEncoding = input.BookmarkPointersEncoding
	bookmarkDB.CopyBasicFieldsFromBookmark_WOP(&input.Bookmark_WOP)

	query := db.Create(&bookmarkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBookmark.CheckoutPhaseOneInstance(&bookmarkDB)
	bookmark := backRepo.BackRepoBookmark.Map_BookmarkDBID_BookmarkPtr[bookmarkDB.ID]

	if bookmark != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), bookmark)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, bookmarkDB)
}

// GetBookmark
//
// swagger:route GET /bookmarks/{ID} bookmarks getBookmark
//
// Gets the details for a bookmark.
//
// Responses:
// default: genericError
//
//	200: bookmarkDBResponse
func (controller *Controller) GetBookmark(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBookmark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookmark.GetDB()

	// Get bookmarkDB in DB
	var bookmarkDB orm.BookmarkDB
	if err := db.First(&bookmarkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var bookmarkAPI orm.BookmarkAPI
	bookmarkAPI.ID = bookmarkDB.ID
	bookmarkAPI.BookmarkPointersEncoding = bookmarkDB.BookmarkPointersEncoding
	bookmarkDB.CopyBasicFieldsToBookmark_WOP(&bookmarkAPI.Bookmark_WOP)

	c.JSON(http.StatusOK, bookmarkAPI)
}

// UpdateBookmark
//
// swagger:route PATCH /bookmarks/{ID} bookmarks updateBookmark
//
// # Update a bookmark
//
// Responses:
// default: genericError
//
//	200: bookmarkDBResponse
func (controller *Controller) UpdateBookmark(c *gin.Context) {

	mutexBookmark.Lock()
	defer mutexBookmark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBookmark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookmark.GetDB()

	// Validate input
	var input orm.BookmarkAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var bookmarkDB orm.BookmarkDB

	// fetch the bookmark
	query := db.First(&bookmarkDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	bookmarkDB.CopyBasicFieldsFromBookmark_WOP(&input.Bookmark_WOP)
	bookmarkDB.BookmarkPointersEncoding = input.BookmarkPointersEncoding

	query = db.Model(&bookmarkDB).Updates(bookmarkDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	bookmarkNew := new(models.Bookmark)
	bookmarkDB.CopyBasicFieldsToBookmark(bookmarkNew)

	// redeem pointers
	bookmarkDB.DecodePointers(backRepo, bookmarkNew)

	// get stage instance from DB instance, and call callback function
	bookmarkOld := backRepo.BackRepoBookmark.Map_BookmarkDBID_BookmarkPtr[bookmarkDB.ID]
	if bookmarkOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), bookmarkOld, bookmarkNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the bookmarkDB
	c.JSON(http.StatusOK, bookmarkDB)
}

// DeleteBookmark
//
// swagger:route DELETE /bookmarks/{ID} bookmarks deleteBookmark
//
// # Delete a bookmark
//
// default: genericError
//
//	200: bookmarkDBResponse
func (controller *Controller) DeleteBookmark(c *gin.Context) {

	mutexBookmark.Lock()
	defer mutexBookmark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBookmark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBookmark.GetDB()

	// Get model if exist
	var bookmarkDB orm.BookmarkDB
	if err := db.First(&bookmarkDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&bookmarkDB)

	// get an instance (not staged) from DB instance, and call callback function
	bookmarkDeleted := new(models.Bookmark)
	bookmarkDB.CopyBasicFieldsToBookmark(bookmarkDeleted)

	// get stage instance from DB instance, and call callback function
	bookmarkStaged := backRepo.BackRepoBookmark.Map_BookmarkDBID_BookmarkPtr[bookmarkDB.ID]
	if bookmarkStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), bookmarkStaged, bookmarkDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
