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
var __Text_element_data__dummysDeclaration__ models.Text_element_data
var __Text_element_data_time__dummyDeclaration time.Duration

var mutexText_element_data sync.Mutex

// An Text_element_dataID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getText_element_data updateText_element_data deleteText_element_data
type Text_element_dataID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Text_element_dataInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postText_element_data updateText_element_data
type Text_element_dataInput struct {
	// The Text_element_data to submit or modify
	// in: body
	Text_element_data *orm.Text_element_dataAPI
}

// GetText_element_datas
//
// swagger:route GET /text_element_datas text_element_datas getText_element_datas
//
// # Get all text_element_datas
//
// Responses:
// default: genericError
//
//	200: text_element_dataDBResponse
func (controller *Controller) GetText_element_datas(c *gin.Context) {

	// source slice
	var text_element_dataDBs []orm.Text_element_dataDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetText_element_datas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText_element_data.GetDB()

	query := db.Find(&text_element_dataDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	text_element_dataAPIs := make([]orm.Text_element_dataAPI, 0)

	// for each text_element_data, update fields from the database nullable fields
	for idx := range text_element_dataDBs {
		text_element_dataDB := &text_element_dataDBs[idx]
		_ = text_element_dataDB
		var text_element_dataAPI orm.Text_element_dataAPI

		// insertion point for updating fields
		text_element_dataAPI.ID = text_element_dataDB.ID
		text_element_dataDB.CopyBasicFieldsToText_element_data_WOP(&text_element_dataAPI.Text_element_data_WOP)
		text_element_dataAPI.Text_element_dataPointersEncoding = text_element_dataDB.Text_element_dataPointersEncoding
		text_element_dataAPIs = append(text_element_dataAPIs, text_element_dataAPI)
	}

	c.JSON(http.StatusOK, text_element_dataAPIs)
}

// PostText_element_data
//
// swagger:route POST /text_element_datas text_element_datas postText_element_data
//
// Creates a text_element_data
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostText_element_data(c *gin.Context) {

	mutexText_element_data.Lock()
	defer mutexText_element_data.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostText_element_datas", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText_element_data.GetDB()

	// Validate input
	var input orm.Text_element_dataAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create text_element_data
	text_element_dataDB := orm.Text_element_dataDB{}
	text_element_dataDB.Text_element_dataPointersEncoding = input.Text_element_dataPointersEncoding
	text_element_dataDB.CopyBasicFieldsFromText_element_data_WOP(&input.Text_element_data_WOP)

	query := db.Create(&text_element_dataDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoText_element_data.CheckoutPhaseOneInstance(&text_element_dataDB)
	text_element_data := backRepo.BackRepoText_element_data.Map_Text_element_dataDBID_Text_element_dataPtr[text_element_dataDB.ID]

	if text_element_data != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), text_element_data)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, text_element_dataDB)
}

// GetText_element_data
//
// swagger:route GET /text_element_datas/{ID} text_element_datas getText_element_data
//
// Gets the details for a text_element_data.
//
// Responses:
// default: genericError
//
//	200: text_element_dataDBResponse
func (controller *Controller) GetText_element_data(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetText_element_data", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText_element_data.GetDB()

	// Get text_element_dataDB in DB
	var text_element_dataDB orm.Text_element_dataDB
	if err := db.First(&text_element_dataDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var text_element_dataAPI orm.Text_element_dataAPI
	text_element_dataAPI.ID = text_element_dataDB.ID
	text_element_dataAPI.Text_element_dataPointersEncoding = text_element_dataDB.Text_element_dataPointersEncoding
	text_element_dataDB.CopyBasicFieldsToText_element_data_WOP(&text_element_dataAPI.Text_element_data_WOP)

	c.JSON(http.StatusOK, text_element_dataAPI)
}

// UpdateText_element_data
//
// swagger:route PATCH /text_element_datas/{ID} text_element_datas updateText_element_data
//
// # Update a text_element_data
//
// Responses:
// default: genericError
//
//	200: text_element_dataDBResponse
func (controller *Controller) UpdateText_element_data(c *gin.Context) {

	mutexText_element_data.Lock()
	defer mutexText_element_data.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateText_element_data", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText_element_data.GetDB()

	// Validate input
	var input orm.Text_element_dataAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var text_element_dataDB orm.Text_element_dataDB

	// fetch the text_element_data
	query := db.First(&text_element_dataDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	text_element_dataDB.CopyBasicFieldsFromText_element_data_WOP(&input.Text_element_data_WOP)
	text_element_dataDB.Text_element_dataPointersEncoding = input.Text_element_dataPointersEncoding

	query = db.Model(&text_element_dataDB).Updates(text_element_dataDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	text_element_dataNew := new(models.Text_element_data)
	text_element_dataDB.CopyBasicFieldsToText_element_data(text_element_dataNew)

	// redeem pointers
	text_element_dataDB.DecodePointers(backRepo, text_element_dataNew)

	// get stage instance from DB instance, and call callback function
	text_element_dataOld := backRepo.BackRepoText_element_data.Map_Text_element_dataDBID_Text_element_dataPtr[text_element_dataDB.ID]
	if text_element_dataOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), text_element_dataOld, text_element_dataNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the text_element_dataDB
	c.JSON(http.StatusOK, text_element_dataDB)
}

// DeleteText_element_data
//
// swagger:route DELETE /text_element_datas/{ID} text_element_datas deleteText_element_data
//
// # Delete a text_element_data
//
// default: genericError
//
//	200: text_element_dataDBResponse
func (controller *Controller) DeleteText_element_data(c *gin.Context) {

	mutexText_element_data.Lock()
	defer mutexText_element_data.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteText_element_data", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoText_element_data.GetDB()

	// Get model if exist
	var text_element_dataDB orm.Text_element_dataDB
	if err := db.First(&text_element_dataDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&text_element_dataDB)

	// get an instance (not staged) from DB instance, and call callback function
	text_element_dataDeleted := new(models.Text_element_data)
	text_element_dataDB.CopyBasicFieldsToText_element_data(text_element_dataDeleted)

	// get stage instance from DB instance, and call callback function
	text_element_dataStaged := backRepo.BackRepoText_element_data.Map_Text_element_dataDBID_Text_element_dataPtr[text_element_dataDB.ID]
	if text_element_dataStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), text_element_dataStaged, text_element_dataDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
