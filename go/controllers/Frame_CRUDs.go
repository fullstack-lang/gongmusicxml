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
var __Frame__dummysDeclaration__ models.Frame
var __Frame_time__dummyDeclaration time.Duration

var mutexFrame sync.Mutex

// An FrameID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getFrame updateFrame deleteFrame
type FrameID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FrameInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postFrame updateFrame
type FrameInput struct {
	// The Frame to submit or modify
	// in: body
	Frame *orm.FrameAPI
}

// GetFrames
//
// swagger:route GET /frames frames getFrames
//
// # Get all frames
//
// Responses:
// default: genericError
//
//	200: frameDBResponse
func (controller *Controller) GetFrames(c *gin.Context) {

	// source slice
	var frameDBs []orm.FrameDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrames", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame.GetDB()

	query := db.Find(&frameDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	frameAPIs := make([]orm.FrameAPI, 0)

	// for each frame, update fields from the database nullable fields
	for idx := range frameDBs {
		frameDB := &frameDBs[idx]
		_ = frameDB
		var frameAPI orm.FrameAPI

		// insertion point for updating fields
		frameAPI.ID = frameDB.ID
		frameDB.CopyBasicFieldsToFrame_WOP(&frameAPI.Frame_WOP)
		frameAPI.FramePointersEncoding = frameDB.FramePointersEncoding
		frameAPIs = append(frameAPIs, frameAPI)
	}

	c.JSON(http.StatusOK, frameAPIs)
}

// PostFrame
//
// swagger:route POST /frames frames postFrame
//
// Creates a frame
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostFrame(c *gin.Context) {

	mutexFrame.Lock()
	defer mutexFrame.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFrames", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame.GetDB()

	// Validate input
	var input orm.FrameAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create frame
	frameDB := orm.FrameDB{}
	frameDB.FramePointersEncoding = input.FramePointersEncoding
	frameDB.CopyBasicFieldsFromFrame_WOP(&input.Frame_WOP)

	query := db.Create(&frameDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoFrame.CheckoutPhaseOneInstance(&frameDB)
	frame := backRepo.BackRepoFrame.Map_FrameDBID_FramePtr[frameDB.ID]

	if frame != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), frame)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, frameDB)
}

// GetFrame
//
// swagger:route GET /frames/{ID} frames getFrame
//
// Gets the details for a frame.
//
// Responses:
// default: genericError
//
//	200: frameDBResponse
func (controller *Controller) GetFrame(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFrame", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame.GetDB()

	// Get frameDB in DB
	var frameDB orm.FrameDB
	if err := db.First(&frameDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var frameAPI orm.FrameAPI
	frameAPI.ID = frameDB.ID
	frameAPI.FramePointersEncoding = frameDB.FramePointersEncoding
	frameDB.CopyBasicFieldsToFrame_WOP(&frameAPI.Frame_WOP)

	c.JSON(http.StatusOK, frameAPI)
}

// UpdateFrame
//
// swagger:route PATCH /frames/{ID} frames updateFrame
//
// # Update a frame
//
// Responses:
// default: genericError
//
//	200: frameDBResponse
func (controller *Controller) UpdateFrame(c *gin.Context) {

	mutexFrame.Lock()
	defer mutexFrame.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateFrame", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame.GetDB()

	// Validate input
	var input orm.FrameAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var frameDB orm.FrameDB

	// fetch the frame
	query := db.First(&frameDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	frameDB.CopyBasicFieldsFromFrame_WOP(&input.Frame_WOP)
	frameDB.FramePointersEncoding = input.FramePointersEncoding

	query = db.Model(&frameDB).Updates(frameDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	frameNew := new(models.Frame)
	frameDB.CopyBasicFieldsToFrame(frameNew)

	// redeem pointers
	frameDB.DecodePointers(backRepo, frameNew)

	// get stage instance from DB instance, and call callback function
	frameOld := backRepo.BackRepoFrame.Map_FrameDBID_FramePtr[frameDB.ID]
	if frameOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), frameOld, frameNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the frameDB
	c.JSON(http.StatusOK, frameDB)
}

// DeleteFrame
//
// swagger:route DELETE /frames/{ID} frames deleteFrame
//
// # Delete a frame
//
// default: genericError
//
//	200: frameDBResponse
func (controller *Controller) DeleteFrame(c *gin.Context) {

	mutexFrame.Lock()
	defer mutexFrame.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteFrame", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoFrame.GetDB()

	// Get model if exist
	var frameDB orm.FrameDB
	if err := db.First(&frameDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&frameDB)

	// get an instance (not staged) from DB instance, and call callback function
	frameDeleted := new(models.Frame)
	frameDB.CopyBasicFieldsToFrame(frameDeleted)

	// get stage instance from DB instance, and call callback function
	frameStaged := backRepo.BackRepoFrame.Map_FrameDBID_FramePtr[frameDB.ID]
	if frameStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), frameStaged, frameDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
