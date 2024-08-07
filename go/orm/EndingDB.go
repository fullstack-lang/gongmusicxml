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
var dummy_Ending_sql sql.NullBool
var dummy_Ending_time time.Duration
var dummy_Ending_sort sort.Float64Slice

// EndingAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model endingAPI
type EndingAPI struct {
	gorm.Model

	models.Ending_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	EndingPointersEncoding EndingPointersEncoding
}

// EndingPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type EndingPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// EndingDB describes a ending in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model endingDB
type EndingDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field endingDB.Name
	Name_Data sql.NullString

	// Declation for basic field endingDB.Value
	Value_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	EndingPointersEncoding
}

// EndingDBs arrays endingDBs
// swagger:response endingDBsResponse
type EndingDBs []EndingDB

// EndingDBResponse provides response
// swagger:response endingDBResponse
type EndingDBResponse struct {
	EndingDB
}

// EndingWOP is a Ending without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type EndingWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Ending_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoEndingStruct struct {
	// stores EndingDB according to their gorm ID
	Map_EndingDBID_EndingDB map[uint]*EndingDB

	// stores EndingDB ID according to Ending address
	Map_EndingPtr_EndingDBID map[*models.Ending]uint

	// stores Ending according to their gorm ID
	Map_EndingDBID_EndingPtr map[uint]*models.Ending

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoEnding *BackRepoEndingStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoEnding.stage
	return
}

func (backRepoEnding *BackRepoEndingStruct) GetDB() *gorm.DB {
	return backRepoEnding.db
}

// GetEndingDBFromEndingPtr is a handy function to access the back repo instance from the stage instance
func (backRepoEnding *BackRepoEndingStruct) GetEndingDBFromEndingPtr(ending *models.Ending) (endingDB *EndingDB) {
	id := backRepoEnding.Map_EndingPtr_EndingDBID[ending]
	endingDB = backRepoEnding.Map_EndingDBID_EndingDB[id]
	return
}

// BackRepoEnding.CommitPhaseOne commits all staged instances of Ending to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEnding *BackRepoEndingStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for ending := range stage.Endings {
		backRepoEnding.CommitPhaseOneInstance(ending)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, ending := range backRepoEnding.Map_EndingDBID_EndingPtr {
		if _, ok := stage.Endings[ending]; !ok {
			backRepoEnding.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoEnding.CommitDeleteInstance commits deletion of Ending to the BackRepo
func (backRepoEnding *BackRepoEndingStruct) CommitDeleteInstance(id uint) (Error error) {

	ending := backRepoEnding.Map_EndingDBID_EndingPtr[id]

	// ending is not staged anymore, remove endingDB
	endingDB := backRepoEnding.Map_EndingDBID_EndingDB[id]
	query := backRepoEnding.db.Unscoped().Delete(&endingDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoEnding.Map_EndingPtr_EndingDBID, ending)
	delete(backRepoEnding.Map_EndingDBID_EndingPtr, id)
	delete(backRepoEnding.Map_EndingDBID_EndingDB, id)

	return
}

// BackRepoEnding.CommitPhaseOneInstance commits ending staged instances of Ending to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEnding *BackRepoEndingStruct) CommitPhaseOneInstance(ending *models.Ending) (Error error) {

	// check if the ending is not commited yet
	if _, ok := backRepoEnding.Map_EndingPtr_EndingDBID[ending]; ok {
		return
	}

	// initiate ending
	var endingDB EndingDB
	endingDB.CopyBasicFieldsFromEnding(ending)

	query := backRepoEnding.db.Create(&endingDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoEnding.Map_EndingPtr_EndingDBID[ending] = endingDB.ID
	backRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID] = ending
	backRepoEnding.Map_EndingDBID_EndingDB[endingDB.ID] = &endingDB

	return
}

// BackRepoEnding.CommitPhaseTwo commits all staged instances of Ending to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEnding *BackRepoEndingStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, ending := range backRepoEnding.Map_EndingDBID_EndingPtr {
		backRepoEnding.CommitPhaseTwoInstance(backRepo, idx, ending)
	}

	return
}

// BackRepoEnding.CommitPhaseTwoInstance commits {{structname }} of models.Ending to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEnding *BackRepoEndingStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, ending *models.Ending) (Error error) {

	// fetch matching endingDB
	if endingDB, ok := backRepoEnding.Map_EndingDBID_EndingDB[idx]; ok {

		endingDB.CopyBasicFieldsFromEnding(ending)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoEnding.db.Save(&endingDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Ending intance %s", ending.Name))
		return err
	}

	return
}

// BackRepoEnding.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoEnding *BackRepoEndingStruct) CheckoutPhaseOne() (Error error) {

	endingDBArray := make([]EndingDB, 0)
	query := backRepoEnding.db.Find(&endingDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	endingInstancesToBeRemovedFromTheStage := make(map[*models.Ending]any)
	for key, value := range backRepoEnding.stage.Endings {
		endingInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, endingDB := range endingDBArray {
		backRepoEnding.CheckoutPhaseOneInstance(&endingDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		ending, ok := backRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID]
		if ok {
			delete(endingInstancesToBeRemovedFromTheStage, ending)
		}
	}

	// remove from stage and back repo's 3 maps all endings that are not in the checkout
	for ending := range endingInstancesToBeRemovedFromTheStage {
		ending.Unstage(backRepoEnding.GetStage())

		// remove instance from the back repo 3 maps
		endingID := backRepoEnding.Map_EndingPtr_EndingDBID[ending]
		delete(backRepoEnding.Map_EndingPtr_EndingDBID, ending)
		delete(backRepoEnding.Map_EndingDBID_EndingDB, endingID)
		delete(backRepoEnding.Map_EndingDBID_EndingPtr, endingID)
	}

	return
}

// CheckoutPhaseOneInstance takes a endingDB that has been found in the DB, updates the backRepo and stages the
// models version of the endingDB
func (backRepoEnding *BackRepoEndingStruct) CheckoutPhaseOneInstance(endingDB *EndingDB) (Error error) {

	ending, ok := backRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID]
	if !ok {
		ending = new(models.Ending)

		backRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID] = ending
		backRepoEnding.Map_EndingPtr_EndingDBID[ending] = endingDB.ID

		// append model store with the new element
		ending.Name = endingDB.Name_Data.String
		ending.Stage(backRepoEnding.GetStage())
	}
	endingDB.CopyBasicFieldsToEnding(ending)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	ending.Stage(backRepoEnding.GetStage())

	// preserve pointer to endingDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_EndingDBID_EndingDB)[endingDB hold variable pointers
	endingDB_Data := *endingDB
	preservedPtrToEnding := &endingDB_Data
	backRepoEnding.Map_EndingDBID_EndingDB[endingDB.ID] = preservedPtrToEnding

	return
}

// BackRepoEnding.CheckoutPhaseTwo Checkouts all staged instances of Ending to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEnding *BackRepoEndingStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, endingDB := range backRepoEnding.Map_EndingDBID_EndingDB {
		backRepoEnding.CheckoutPhaseTwoInstance(backRepo, endingDB)
	}
	return
}

// BackRepoEnding.CheckoutPhaseTwoInstance Checkouts staged instances of Ending to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEnding *BackRepoEndingStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, endingDB *EndingDB) (Error error) {

	ending := backRepoEnding.Map_EndingDBID_EndingPtr[endingDB.ID]

	endingDB.DecodePointers(backRepo, ending)

	return
}

func (endingDB *EndingDB) DecodePointers(backRepo *BackRepoStruct, ending *models.Ending) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitEnding allows commit of a single ending (if already staged)
func (backRepo *BackRepoStruct) CommitEnding(ending *models.Ending) {
	backRepo.BackRepoEnding.CommitPhaseOneInstance(ending)
	if id, ok := backRepo.BackRepoEnding.Map_EndingPtr_EndingDBID[ending]; ok {
		backRepo.BackRepoEnding.CommitPhaseTwoInstance(backRepo, id, ending)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitEnding allows checkout of a single ending (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutEnding(ending *models.Ending) {
	// check if the ending is staged
	if _, ok := backRepo.BackRepoEnding.Map_EndingPtr_EndingDBID[ending]; ok {

		if id, ok := backRepo.BackRepoEnding.Map_EndingPtr_EndingDBID[ending]; ok {
			var endingDB EndingDB
			endingDB.ID = id

			if err := backRepo.BackRepoEnding.db.First(&endingDB, id).Error; err != nil {
				log.Fatalln("CheckoutEnding : Problem with getting object with id:", id)
			}
			backRepo.BackRepoEnding.CheckoutPhaseOneInstance(&endingDB)
			backRepo.BackRepoEnding.CheckoutPhaseTwoInstance(backRepo, &endingDB)
		}
	}
}

// CopyBasicFieldsFromEnding
func (endingDB *EndingDB) CopyBasicFieldsFromEnding(ending *models.Ending) {
	// insertion point for fields commit

	endingDB.Name_Data.String = ending.Name
	endingDB.Name_Data.Valid = true

	endingDB.Value_Data.String = ending.Value
	endingDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromEnding_WOP
func (endingDB *EndingDB) CopyBasicFieldsFromEnding_WOP(ending *models.Ending_WOP) {
	// insertion point for fields commit

	endingDB.Name_Data.String = ending.Name
	endingDB.Name_Data.Valid = true

	endingDB.Value_Data.String = ending.Value
	endingDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromEndingWOP
func (endingDB *EndingDB) CopyBasicFieldsFromEndingWOP(ending *EndingWOP) {
	// insertion point for fields commit

	endingDB.Name_Data.String = ending.Name
	endingDB.Name_Data.Valid = true

	endingDB.Value_Data.String = ending.Value
	endingDB.Value_Data.Valid = true
}

// CopyBasicFieldsToEnding
func (endingDB *EndingDB) CopyBasicFieldsToEnding(ending *models.Ending) {
	// insertion point for checkout of basic fields (back repo to stage)
	ending.Name = endingDB.Name_Data.String
	ending.Value = endingDB.Value_Data.String
}

// CopyBasicFieldsToEnding_WOP
func (endingDB *EndingDB) CopyBasicFieldsToEnding_WOP(ending *models.Ending_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	ending.Name = endingDB.Name_Data.String
	ending.Value = endingDB.Value_Data.String
}

// CopyBasicFieldsToEndingWOP
func (endingDB *EndingDB) CopyBasicFieldsToEndingWOP(ending *EndingWOP) {
	ending.ID = int(endingDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	ending.Name = endingDB.Name_Data.String
	ending.Value = endingDB.Value_Data.String
}

// Backup generates a json file from a slice of all EndingDB instances in the backrepo
func (backRepoEnding *BackRepoEndingStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "EndingDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*EndingDB, 0)
	for _, endingDB := range backRepoEnding.Map_EndingDBID_EndingDB {
		forBackup = append(forBackup, endingDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Ending ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Ending file", err.Error())
	}
}

// Backup generates a json file from a slice of all EndingDB instances in the backrepo
func (backRepoEnding *BackRepoEndingStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*EndingDB, 0)
	for _, endingDB := range backRepoEnding.Map_EndingDBID_EndingDB {
		forBackup = append(forBackup, endingDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Ending")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Ending_Fields, -1)
	for _, endingDB := range forBackup {

		var endingWOP EndingWOP
		endingDB.CopyBasicFieldsToEndingWOP(&endingWOP)

		row := sh.AddRow()
		row.WriteStruct(&endingWOP, -1)
	}
}

// RestoreXL from the "Ending" sheet all EndingDB instances
func (backRepoEnding *BackRepoEndingStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoEndingid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Ending"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoEnding.rowVisitorEnding)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoEnding *BackRepoEndingStruct) rowVisitorEnding(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var endingWOP EndingWOP
		row.ReadStruct(&endingWOP)

		// add the unmarshalled struct to the stage
		endingDB := new(EndingDB)
		endingDB.CopyBasicFieldsFromEndingWOP(&endingWOP)

		endingDB_ID_atBackupTime := endingDB.ID
		endingDB.ID = 0
		query := backRepoEnding.db.Create(endingDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoEnding.Map_EndingDBID_EndingDB[endingDB.ID] = endingDB
		BackRepoEndingid_atBckpTime_newID[endingDB_ID_atBackupTime] = endingDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "EndingDB.json" in dirPath that stores an array
// of EndingDB and stores it in the database
// the map BackRepoEndingid_atBckpTime_newID is updated accordingly
func (backRepoEnding *BackRepoEndingStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoEndingid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "EndingDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Ending file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*EndingDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_EndingDBID_EndingDB
	for _, endingDB := range forRestore {

		endingDB_ID_atBackupTime := endingDB.ID
		endingDB.ID = 0
		query := backRepoEnding.db.Create(endingDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoEnding.Map_EndingDBID_EndingDB[endingDB.ID] = endingDB
		BackRepoEndingid_atBckpTime_newID[endingDB_ID_atBackupTime] = endingDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Ending file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Ending>id_atBckpTime_newID
// to compute new index
func (backRepoEnding *BackRepoEndingStruct) RestorePhaseTwo() {

	for _, endingDB := range backRepoEnding.Map_EndingDBID_EndingDB {

		// next line of code is to avert unused variable compilation error
		_ = endingDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoEnding.db.Model(endingDB).Updates(*endingDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoEnding.ResetReversePointers commits all staged instances of Ending to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEnding *BackRepoEndingStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, ending := range backRepoEnding.Map_EndingDBID_EndingPtr {
		backRepoEnding.ResetReversePointersInstance(backRepo, idx, ending)
	}

	return
}

func (backRepoEnding *BackRepoEndingStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, ending *models.Ending) (Error error) {

	// fetch matching endingDB
	if endingDB, ok := backRepoEnding.Map_EndingDBID_EndingDB[idx]; ok {
		_ = endingDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoEndingid_atBckpTime_newID map[uint]uint
