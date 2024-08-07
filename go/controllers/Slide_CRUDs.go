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
var __Slide__dummysDeclaration__ models.Slide
var __Slide_time__dummyDeclaration time.Duration

var mutexSlide sync.Mutex

// An SlideID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSlide updateSlide deleteSlide
type SlideID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SlideInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSlide updateSlide
type SlideInput struct {
	// The Slide to submit or modify
	// in: body
	Slide *orm.SlideAPI
}

// GetSlides
//
// swagger:route GET /slides slides getSlides
//
// # Get all slides
//
// Responses:
// default: genericError
//
//	200: slideDBResponse
func (controller *Controller) GetSlides(c *gin.Context) {

	// source slice
	var slideDBs []orm.SlideDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlides", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlide.GetDB()

	query := db.Find(&slideDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	slideAPIs := make([]orm.SlideAPI, 0)

	// for each slide, update fields from the database nullable fields
	for idx := range slideDBs {
		slideDB := &slideDBs[idx]
		_ = slideDB
		var slideAPI orm.SlideAPI

		// insertion point for updating fields
		slideAPI.ID = slideDB.ID
		slideDB.CopyBasicFieldsToSlide_WOP(&slideAPI.Slide_WOP)
		slideAPI.SlidePointersEncoding = slideDB.SlidePointersEncoding
		slideAPIs = append(slideAPIs, slideAPI)
	}

	c.JSON(http.StatusOK, slideAPIs)
}

// PostSlide
//
// swagger:route POST /slides slides postSlide
//
// Creates a slide
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSlide(c *gin.Context) {

	mutexSlide.Lock()
	defer mutexSlide.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSlides", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlide.GetDB()

	// Validate input
	var input orm.SlideAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create slide
	slideDB := orm.SlideDB{}
	slideDB.SlidePointersEncoding = input.SlidePointersEncoding
	slideDB.CopyBasicFieldsFromSlide_WOP(&input.Slide_WOP)

	query := db.Create(&slideDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSlide.CheckoutPhaseOneInstance(&slideDB)
	slide := backRepo.BackRepoSlide.Map_SlideDBID_SlidePtr[slideDB.ID]

	if slide != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), slide)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, slideDB)
}

// GetSlide
//
// swagger:route GET /slides/{ID} slides getSlide
//
// Gets the details for a slide.
//
// Responses:
// default: genericError
//
//	200: slideDBResponse
func (controller *Controller) GetSlide(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSlide", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlide.GetDB()

	// Get slideDB in DB
	var slideDB orm.SlideDB
	if err := db.First(&slideDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var slideAPI orm.SlideAPI
	slideAPI.ID = slideDB.ID
	slideAPI.SlidePointersEncoding = slideDB.SlidePointersEncoding
	slideDB.CopyBasicFieldsToSlide_WOP(&slideAPI.Slide_WOP)

	c.JSON(http.StatusOK, slideAPI)
}

// UpdateSlide
//
// swagger:route PATCH /slides/{ID} slides updateSlide
//
// # Update a slide
//
// Responses:
// default: genericError
//
//	200: slideDBResponse
func (controller *Controller) UpdateSlide(c *gin.Context) {

	mutexSlide.Lock()
	defer mutexSlide.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSlide", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlide.GetDB()

	// Validate input
	var input orm.SlideAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var slideDB orm.SlideDB

	// fetch the slide
	query := db.First(&slideDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	slideDB.CopyBasicFieldsFromSlide_WOP(&input.Slide_WOP)
	slideDB.SlidePointersEncoding = input.SlidePointersEncoding

	query = db.Model(&slideDB).Updates(slideDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	slideNew := new(models.Slide)
	slideDB.CopyBasicFieldsToSlide(slideNew)

	// redeem pointers
	slideDB.DecodePointers(backRepo, slideNew)

	// get stage instance from DB instance, and call callback function
	slideOld := backRepo.BackRepoSlide.Map_SlideDBID_SlidePtr[slideDB.ID]
	if slideOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), slideOld, slideNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the slideDB
	c.JSON(http.StatusOK, slideDB)
}

// DeleteSlide
//
// swagger:route DELETE /slides/{ID} slides deleteSlide
//
// # Delete a slide
//
// default: genericError
//
//	200: slideDBResponse
func (controller *Controller) DeleteSlide(c *gin.Context) {

	mutexSlide.Lock()
	defer mutexSlide.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSlide", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSlide.GetDB()

	// Get model if exist
	var slideDB orm.SlideDB
	if err := db.First(&slideDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&slideDB)

	// get an instance (not staged) from DB instance, and call callback function
	slideDeleted := new(models.Slide)
	slideDB.CopyBasicFieldsToSlide(slideDeleted)

	// get stage instance from DB instance, and call callback function
	slideStaged := backRepo.BackRepoSlide.Map_SlideDBID_SlidePtr[slideDB.ID]
	if slideStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), slideStaged, slideDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
