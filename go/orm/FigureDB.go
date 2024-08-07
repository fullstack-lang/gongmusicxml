// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Figure_sql sql.NullBool
var dummy_Figure_time time.Duration
var dummy_Figure_sort sort.Float64Slice

// FigureAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model figureAPI
type FigureAPI struct {
	gorm.Model

	models.Figure_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	FigurePointersEncoding FigurePointersEncoding
}

// FigurePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type FigurePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Extend is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ExtendID sql.NullInt64
}

// FigureDB describes a figure in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model figureDB
type FigureDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field figureDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	FigurePointersEncoding
}

// FigureDBs arrays figureDBs
// swagger:response figureDBsResponse
type FigureDBs []FigureDB

// FigureDBResponse provides response
// swagger:response figureDBResponse
type FigureDBResponse struct {
	FigureDB
}

// FigureWOP is a Figure without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type FigureWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Figure_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoFigureStruct struct {
	// stores FigureDB according to their gorm ID
	Map_FigureDBID_FigureDB map[uint]*FigureDB

	// stores FigureDB ID according to Figure address
	Map_FigurePtr_FigureDBID map[*models.Figure]uint

	// stores Figure according to their gorm ID
	Map_FigureDBID_FigurePtr map[uint]*models.Figure

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFigure *BackRepoFigureStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFigure.stage
	return
}

func (backRepoFigure *BackRepoFigureStruct) GetDB() *gorm.DB {
	return backRepoFigure.db
}

// GetFigureDBFromFigurePtr is a handy function to access the back repo instance from the stage instance
func (backRepoFigure *BackRepoFigureStruct) GetFigureDBFromFigurePtr(figure *models.Figure) (figureDB *FigureDB) {
	id := backRepoFigure.Map_FigurePtr_FigureDBID[figure]
	figureDB = backRepoFigure.Map_FigureDBID_FigureDB[id]
	return
}

// BackRepoFigure.CommitPhaseOne commits all staged instances of Figure to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFigure *BackRepoFigureStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for figure := range stage.Figures {
		backRepoFigure.CommitPhaseOneInstance(figure)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, figure := range backRepoFigure.Map_FigureDBID_FigurePtr {
		if _, ok := stage.Figures[figure]; !ok {
			backRepoFigure.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFigure.CommitDeleteInstance commits deletion of Figure to the BackRepo
func (backRepoFigure *BackRepoFigureStruct) CommitDeleteInstance(id uint) (Error error) {

	figure := backRepoFigure.Map_FigureDBID_FigurePtr[id]

	// figure is not staged anymore, remove figureDB
	figureDB := backRepoFigure.Map_FigureDBID_FigureDB[id]
	query := backRepoFigure.db.Unscoped().Delete(&figureDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFigure.Map_FigurePtr_FigureDBID, figure)
	delete(backRepoFigure.Map_FigureDBID_FigurePtr, id)
	delete(backRepoFigure.Map_FigureDBID_FigureDB, id)

	return
}

// BackRepoFigure.CommitPhaseOneInstance commits figure staged instances of Figure to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFigure *BackRepoFigureStruct) CommitPhaseOneInstance(figure *models.Figure) (Error error) {

	// check if the figure is not commited yet
	if _, ok := backRepoFigure.Map_FigurePtr_FigureDBID[figure]; ok {
		return
	}

	// initiate figure
	var figureDB FigureDB
	figureDB.CopyBasicFieldsFromFigure(figure)

	query := backRepoFigure.db.Create(&figureDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFigure.Map_FigurePtr_FigureDBID[figure] = figureDB.ID
	backRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID] = figure
	backRepoFigure.Map_FigureDBID_FigureDB[figureDB.ID] = &figureDB

	return
}

// BackRepoFigure.CommitPhaseTwo commits all staged instances of Figure to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFigure *BackRepoFigureStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, figure := range backRepoFigure.Map_FigureDBID_FigurePtr {
		backRepoFigure.CommitPhaseTwoInstance(backRepo, idx, figure)
	}

	return
}

// BackRepoFigure.CommitPhaseTwoInstance commits {{structname }} of models.Figure to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFigure *BackRepoFigureStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, figure *models.Figure) (Error error) {

	// fetch matching figureDB
	if figureDB, ok := backRepoFigure.Map_FigureDBID_FigureDB[idx]; ok {

		figureDB.CopyBasicFieldsFromFigure(figure)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value figure.Extend translates to updating the figure.ExtendID
		figureDB.ExtendID.Valid = true // allow for a 0 value (nil association)
		if figure.Extend != nil {
			if ExtendId, ok := backRepo.BackRepoExtend.Map_ExtendPtr_ExtendDBID[figure.Extend]; ok {
				figureDB.ExtendID.Int64 = int64(ExtendId)
				figureDB.ExtendID.Valid = true
			}
		} else {
			figureDB.ExtendID.Int64 = 0
			figureDB.ExtendID.Valid = true
		}

		query := backRepoFigure.db.Save(&figureDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Figure intance %s", figure.Name))
		return err
	}

	return
}

// BackRepoFigure.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFigure *BackRepoFigureStruct) CheckoutPhaseOne() (Error error) {

	figureDBArray := make([]FigureDB, 0)
	query := backRepoFigure.db.Find(&figureDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	figureInstancesToBeRemovedFromTheStage := make(map[*models.Figure]any)
	for key, value := range backRepoFigure.stage.Figures {
		figureInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, figureDB := range figureDBArray {
		backRepoFigure.CheckoutPhaseOneInstance(&figureDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		figure, ok := backRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID]
		if ok {
			delete(figureInstancesToBeRemovedFromTheStage, figure)
		}
	}

	// remove from stage and back repo's 3 maps all figures that are not in the checkout
	for figure := range figureInstancesToBeRemovedFromTheStage {
		figure.Unstage(backRepoFigure.GetStage())

		// remove instance from the back repo 3 maps
		figureID := backRepoFigure.Map_FigurePtr_FigureDBID[figure]
		delete(backRepoFigure.Map_FigurePtr_FigureDBID, figure)
		delete(backRepoFigure.Map_FigureDBID_FigureDB, figureID)
		delete(backRepoFigure.Map_FigureDBID_FigurePtr, figureID)
	}

	return
}

// CheckoutPhaseOneInstance takes a figureDB that has been found in the DB, updates the backRepo and stages the
// models version of the figureDB
func (backRepoFigure *BackRepoFigureStruct) CheckoutPhaseOneInstance(figureDB *FigureDB) (Error error) {

	figure, ok := backRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID]
	if !ok {
		figure = new(models.Figure)

		backRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID] = figure
		backRepoFigure.Map_FigurePtr_FigureDBID[figure] = figureDB.ID

		// append model store with the new element
		figure.Name = figureDB.Name_Data.String
		figure.Stage(backRepoFigure.GetStage())
	}
	figureDB.CopyBasicFieldsToFigure(figure)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	figure.Stage(backRepoFigure.GetStage())

	// preserve pointer to figureDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_FigureDBID_FigureDB)[figureDB hold variable pointers
	figureDB_Data := *figureDB
	preservedPtrToFigure := &figureDB_Data
	backRepoFigure.Map_FigureDBID_FigureDB[figureDB.ID] = preservedPtrToFigure

	return
}

// BackRepoFigure.CheckoutPhaseTwo Checkouts all staged instances of Figure to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFigure *BackRepoFigureStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, figureDB := range backRepoFigure.Map_FigureDBID_FigureDB {
		backRepoFigure.CheckoutPhaseTwoInstance(backRepo, figureDB)
	}
	return
}

// BackRepoFigure.CheckoutPhaseTwoInstance Checkouts staged instances of Figure to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFigure *BackRepoFigureStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, figureDB *FigureDB) (Error error) {

	figure := backRepoFigure.Map_FigureDBID_FigurePtr[figureDB.ID]

	figureDB.DecodePointers(backRepo, figure)

	return
}

func (figureDB *FigureDB) DecodePointers(backRepo *BackRepoStruct, figure *models.Figure) {

	// insertion point for checkout of pointer encoding
	// Extend field
	figure.Extend = nil
	if figureDB.ExtendID.Int64 != 0 {
		figure.Extend = backRepo.BackRepoExtend.Map_ExtendDBID_ExtendPtr[uint(figureDB.ExtendID.Int64)]
	}
	return
}

// CommitFigure allows commit of a single figure (if already staged)
func (backRepo *BackRepoStruct) CommitFigure(figure *models.Figure) {
	backRepo.BackRepoFigure.CommitPhaseOneInstance(figure)
	if id, ok := backRepo.BackRepoFigure.Map_FigurePtr_FigureDBID[figure]; ok {
		backRepo.BackRepoFigure.CommitPhaseTwoInstance(backRepo, id, figure)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFigure allows checkout of a single figure (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFigure(figure *models.Figure) {
	// check if the figure is staged
	if _, ok := backRepo.BackRepoFigure.Map_FigurePtr_FigureDBID[figure]; ok {

		if id, ok := backRepo.BackRepoFigure.Map_FigurePtr_FigureDBID[figure]; ok {
			var figureDB FigureDB
			figureDB.ID = id

			if err := backRepo.BackRepoFigure.db.First(&figureDB, id).Error; err != nil {
				log.Fatalln("CheckoutFigure : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFigure.CheckoutPhaseOneInstance(&figureDB)
			backRepo.BackRepoFigure.CheckoutPhaseTwoInstance(backRepo, &figureDB)
		}
	}
}

// CopyBasicFieldsFromFigure
func (figureDB *FigureDB) CopyBasicFieldsFromFigure(figure *models.Figure) {
	// insertion point for fields commit

	figureDB.Name_Data.String = figure.Name
	figureDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFigure_WOP
func (figureDB *FigureDB) CopyBasicFieldsFromFigure_WOP(figure *models.Figure_WOP) {
	// insertion point for fields commit

	figureDB.Name_Data.String = figure.Name
	figureDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFigureWOP
func (figureDB *FigureDB) CopyBasicFieldsFromFigureWOP(figure *FigureWOP) {
	// insertion point for fields commit

	figureDB.Name_Data.String = figure.Name
	figureDB.Name_Data.Valid = true
}

// CopyBasicFieldsToFigure
func (figureDB *FigureDB) CopyBasicFieldsToFigure(figure *models.Figure) {
	// insertion point for checkout of basic fields (back repo to stage)
	figure.Name = figureDB.Name_Data.String
}

// CopyBasicFieldsToFigure_WOP
func (figureDB *FigureDB) CopyBasicFieldsToFigure_WOP(figure *models.Figure_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	figure.Name = figureDB.Name_Data.String
}

// CopyBasicFieldsToFigureWOP
func (figureDB *FigureDB) CopyBasicFieldsToFigureWOP(figure *FigureWOP) {
	figure.ID = int(figureDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	figure.Name = figureDB.Name_Data.String
}

// Backup generates a json file from a slice of all FigureDB instances in the backrepo
func (backRepoFigure *BackRepoFigureStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "FigureDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FigureDB, 0)
	for _, figureDB := range backRepoFigure.Map_FigureDBID_FigureDB {
		forBackup = append(forBackup, figureDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Figure ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Figure file", err.Error())
	}
}

// Backup generates a json file from a slice of all FigureDB instances in the backrepo
func (backRepoFigure *BackRepoFigureStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*FigureDB, 0)
	for _, figureDB := range backRepoFigure.Map_FigureDBID_FigureDB {
		forBackup = append(forBackup, figureDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Figure")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Figure_Fields, -1)
	for _, figureDB := range forBackup {

		var figureWOP FigureWOP
		figureDB.CopyBasicFieldsToFigureWOP(&figureWOP)

		row := sh.AddRow()
		row.WriteStruct(&figureWOP, -1)
	}
}

// RestoreXL from the "Figure" sheet all FigureDB instances
func (backRepoFigure *BackRepoFigureStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFigureid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Figure"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFigure.rowVisitorFigure)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFigure *BackRepoFigureStruct) rowVisitorFigure(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var figureWOP FigureWOP
		row.ReadStruct(&figureWOP)

		// add the unmarshalled struct to the stage
		figureDB := new(FigureDB)
		figureDB.CopyBasicFieldsFromFigureWOP(&figureWOP)

		figureDB_ID_atBackupTime := figureDB.ID
		figureDB.ID = 0
		query := backRepoFigure.db.Create(figureDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFigure.Map_FigureDBID_FigureDB[figureDB.ID] = figureDB
		BackRepoFigureid_atBckpTime_newID[figureDB_ID_atBackupTime] = figureDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "FigureDB.json" in dirPath that stores an array
// of FigureDB and stores it in the database
// the map BackRepoFigureid_atBckpTime_newID is updated accordingly
func (backRepoFigure *BackRepoFigureStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFigureid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "FigureDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Figure file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*FigureDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_FigureDBID_FigureDB
	for _, figureDB := range forRestore {

		figureDB_ID_atBackupTime := figureDB.ID
		figureDB.ID = 0
		query := backRepoFigure.db.Create(figureDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFigure.Map_FigureDBID_FigureDB[figureDB.ID] = figureDB
		BackRepoFigureid_atBckpTime_newID[figureDB_ID_atBackupTime] = figureDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Figure file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Figure>id_atBckpTime_newID
// to compute new index
func (backRepoFigure *BackRepoFigureStruct) RestorePhaseTwo() {

	for _, figureDB := range backRepoFigure.Map_FigureDBID_FigureDB {

		// next line of code is to avert unused variable compilation error
		_ = figureDB

		// insertion point for reindexing pointers encoding
		// reindexing Extend field
		if figureDB.ExtendID.Int64 != 0 {
			figureDB.ExtendID.Int64 = int64(BackRepoExtendid_atBckpTime_newID[uint(figureDB.ExtendID.Int64)])
			figureDB.ExtendID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoFigure.db.Model(figureDB).Updates(*figureDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFigure.ResetReversePointers commits all staged instances of Figure to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFigure *BackRepoFigureStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, figure := range backRepoFigure.Map_FigureDBID_FigurePtr {
		backRepoFigure.ResetReversePointersInstance(backRepo, idx, figure)
	}

	return
}

func (backRepoFigure *BackRepoFigureStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, figure *models.Figure) (Error error) {

	// fetch matching figureDB
	if figureDB, ok := backRepoFigure.Map_FigureDBID_FigureDB[idx]; ok {
		_ = figureDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFigureid_atBckpTime_newID map[uint]uint
