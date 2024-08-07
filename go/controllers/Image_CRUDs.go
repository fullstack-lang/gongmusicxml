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
var __Image__dummysDeclaration__ models.Image
var __Image_time__dummyDeclaration time.Duration

var mutexImage sync.Mutex

// An ImageID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getImage updateImage deleteImage
type ImageID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ImageInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postImage updateImage
type ImageInput struct {
	// The Image to submit or modify
	// in: body
	Image *orm.ImageAPI
}

// GetImages
//
// swagger:route GET /images images getImages
//
// # Get all images
//
// Responses:
// default: genericError
//
//	200: imageDBResponse
func (controller *Controller) GetImages(c *gin.Context) {

	// source slice
	var imageDBs []orm.ImageDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetImages", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoImage.GetDB()

	query := db.Find(&imageDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	imageAPIs := make([]orm.ImageAPI, 0)

	// for each image, update fields from the database nullable fields
	for idx := range imageDBs {
		imageDB := &imageDBs[idx]
		_ = imageDB
		var imageAPI orm.ImageAPI

		// insertion point for updating fields
		imageAPI.ID = imageDB.ID
		imageDB.CopyBasicFieldsToImage_WOP(&imageAPI.Image_WOP)
		imageAPI.ImagePointersEncoding = imageDB.ImagePointersEncoding
		imageAPIs = append(imageAPIs, imageAPI)
	}

	c.JSON(http.StatusOK, imageAPIs)
}

// PostImage
//
// swagger:route POST /images images postImage
//
// Creates a image
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostImage(c *gin.Context) {

	mutexImage.Lock()
	defer mutexImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostImages", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoImage.GetDB()

	// Validate input
	var input orm.ImageAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create image
	imageDB := orm.ImageDB{}
	imageDB.ImagePointersEncoding = input.ImagePointersEncoding
	imageDB.CopyBasicFieldsFromImage_WOP(&input.Image_WOP)

	query := db.Create(&imageDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoImage.CheckoutPhaseOneInstance(&imageDB)
	image := backRepo.BackRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID]

	if image != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), image)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, imageDB)
}

// GetImage
//
// swagger:route GET /images/{ID} images getImage
//
// Gets the details for a image.
//
// Responses:
// default: genericError
//
//	200: imageDBResponse
func (controller *Controller) GetImage(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetImage", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoImage.GetDB()

	// Get imageDB in DB
	var imageDB orm.ImageDB
	if err := db.First(&imageDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var imageAPI orm.ImageAPI
	imageAPI.ID = imageDB.ID
	imageAPI.ImagePointersEncoding = imageDB.ImagePointersEncoding
	imageDB.CopyBasicFieldsToImage_WOP(&imageAPI.Image_WOP)

	c.JSON(http.StatusOK, imageAPI)
}

// UpdateImage
//
// swagger:route PATCH /images/{ID} images updateImage
//
// # Update a image
//
// Responses:
// default: genericError
//
//	200: imageDBResponse
func (controller *Controller) UpdateImage(c *gin.Context) {

	mutexImage.Lock()
	defer mutexImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateImage", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoImage.GetDB()

	// Validate input
	var input orm.ImageAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var imageDB orm.ImageDB

	// fetch the image
	query := db.First(&imageDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	imageDB.CopyBasicFieldsFromImage_WOP(&input.Image_WOP)
	imageDB.ImagePointersEncoding = input.ImagePointersEncoding

	query = db.Model(&imageDB).Updates(imageDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	imageNew := new(models.Image)
	imageDB.CopyBasicFieldsToImage(imageNew)

	// redeem pointers
	imageDB.DecodePointers(backRepo, imageNew)

	// get stage instance from DB instance, and call callback function
	imageOld := backRepo.BackRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID]
	if imageOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), imageOld, imageNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the imageDB
	c.JSON(http.StatusOK, imageDB)
}

// DeleteImage
//
// swagger:route DELETE /images/{ID} images deleteImage
//
// # Delete a image
//
// default: genericError
//
//	200: imageDBResponse
func (controller *Controller) DeleteImage(c *gin.Context) {

	mutexImage.Lock()
	defer mutexImage.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteImage", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongmusicxml/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoImage.GetDB()

	// Get model if exist
	var imageDB orm.ImageDB
	if err := db.First(&imageDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&imageDB)

	// get an instance (not staged) from DB instance, and call callback function
	imageDeleted := new(models.Image)
	imageDB.CopyBasicFieldsToImage(imageDeleted)

	// get stage instance from DB instance, and call callback function
	imageStaged := backRepo.BackRepoImage.Map_ImageDBID_ImagePtr[imageDB.ID]
	if imageStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), imageStaged, imageDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
