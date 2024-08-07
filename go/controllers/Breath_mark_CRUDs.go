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
var __Breath_mark__dummysDeclaration__ models.Breath_mark
var __Breath_mark_time__dummyDeclaration time.Duration

var mutexBreath_mark sync.Mutex

// An Breath_markID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getBreath_mark updateBreath_mark deleteBreath_mark
type Breath_markID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Breath_markInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postBreath_mark updateBreath_mark
type Breath_markInput struct {
	// The Breath_mark to submit or modify
	// in: body
	Breath_mark *orm.Breath_markAPI
}

// GetBreath_marks
//
// swagger:route GET /breath_marks breath_marks getBreath_marks
//
// # Get all breath_marks
//
// Responses:
// default: genericError
//
//	200: breath_markDBResponse
func (controller *Controller) GetBreath_marks(c *gin.Context) {

	// source slice
	var breath_markDBs []orm.Breath_markDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBreath_marks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBreath_mark.GetDB()

	query := db.Find(&breath_markDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	breath_markAPIs := make([]orm.Breath_markAPI, 0)

	// for each breath_mark, update fields from the database nullable fields
	for idx := range breath_markDBs {
		breath_markDB := &breath_markDBs[idx]
		_ = breath_markDB
		var breath_markAPI orm.Breath_markAPI

		// insertion point for updating fields
		breath_markAPI.ID = breath_markDB.ID
		breath_markDB.CopyBasicFieldsToBreath_mark_WOP(&breath_markAPI.Breath_mark_WOP)
		breath_markAPI.Breath_markPointersEncoding = breath_markDB.Breath_markPointersEncoding
		breath_markAPIs = append(breath_markAPIs, breath_markAPI)
	}

	c.JSON(http.StatusOK, breath_markAPIs)
}

// PostBreath_mark
//
// swagger:route POST /breath_marks breath_marks postBreath_mark
//
// Creates a breath_mark
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostBreath_mark(c *gin.Context) {

	mutexBreath_mark.Lock()
	defer mutexBreath_mark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostBreath_marks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBreath_mark.GetDB()

	// Validate input
	var input orm.Breath_markAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create breath_mark
	breath_markDB := orm.Breath_markDB{}
	breath_markDB.Breath_markPointersEncoding = input.Breath_markPointersEncoding
	breath_markDB.CopyBasicFieldsFromBreath_mark_WOP(&input.Breath_mark_WOP)

	query := db.Create(&breath_markDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoBreath_mark.CheckoutPhaseOneInstance(&breath_markDB)
	breath_mark := backRepo.BackRepoBreath_mark.Map_Breath_markDBID_Breath_markPtr[breath_markDB.ID]

	if breath_mark != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), breath_mark)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, breath_markDB)
}

// GetBreath_mark
//
// swagger:route GET /breath_marks/{ID} breath_marks getBreath_mark
//
// Gets the details for a breath_mark.
//
// Responses:
// default: genericError
//
//	200: breath_markDBResponse
func (controller *Controller) GetBreath_mark(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetBreath_mark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBreath_mark.GetDB()

	// Get breath_markDB in DB
	var breath_markDB orm.Breath_markDB
	if err := db.First(&breath_markDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var breath_markAPI orm.Breath_markAPI
	breath_markAPI.ID = breath_markDB.ID
	breath_markAPI.Breath_markPointersEncoding = breath_markDB.Breath_markPointersEncoding
	breath_markDB.CopyBasicFieldsToBreath_mark_WOP(&breath_markAPI.Breath_mark_WOP)

	c.JSON(http.StatusOK, breath_markAPI)
}

// UpdateBreath_mark
//
// swagger:route PATCH /breath_marks/{ID} breath_marks updateBreath_mark
//
// # Update a breath_mark
//
// Responses:
// default: genericError
//
//	200: breath_markDBResponse
func (controller *Controller) UpdateBreath_mark(c *gin.Context) {

	mutexBreath_mark.Lock()
	defer mutexBreath_mark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateBreath_mark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBreath_mark.GetDB()

	// Validate input
	var input orm.Breath_markAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var breath_markDB orm.Breath_markDB

	// fetch the breath_mark
	query := db.First(&breath_markDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	breath_markDB.CopyBasicFieldsFromBreath_mark_WOP(&input.Breath_mark_WOP)
	breath_markDB.Breath_markPointersEncoding = input.Breath_markPointersEncoding

	query = db.Model(&breath_markDB).Updates(breath_markDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	breath_markNew := new(models.Breath_mark)
	breath_markDB.CopyBasicFieldsToBreath_mark(breath_markNew)

	// redeem pointers
	breath_markDB.DecodePointers(backRepo, breath_markNew)

	// get stage instance from DB instance, and call callback function
	breath_markOld := backRepo.BackRepoBreath_mark.Map_Breath_markDBID_Breath_markPtr[breath_markDB.ID]
	if breath_markOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), breath_markOld, breath_markNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the breath_markDB
	c.JSON(http.StatusOK, breath_markDB)
}

// DeleteBreath_mark
//
// swagger:route DELETE /breath_marks/{ID} breath_marks deleteBreath_mark
//
// # Delete a breath_mark
//
// default: genericError
//
//	200: breath_markDBResponse
func (controller *Controller) DeleteBreath_mark(c *gin.Context) {

	mutexBreath_mark.Lock()
	defer mutexBreath_mark.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteBreath_mark", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoBreath_mark.GetDB()

	// Get model if exist
	var breath_markDB orm.Breath_markDB
	if err := db.First(&breath_markDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&breath_markDB)

	// get an instance (not staged) from DB instance, and call callback function
	breath_markDeleted := new(models.Breath_mark)
	breath_markDB.CopyBasicFieldsToBreath_mark(breath_markDeleted)

	// get stage instance from DB instance, and call callback function
	breath_markStaged := backRepo.BackRepoBreath_mark.Map_Breath_markDBID_Breath_markPtr[breath_markDB.ID]
	if breath_markStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), breath_markStaged, breath_markDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
