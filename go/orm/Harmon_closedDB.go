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
var dummy_Harmon_closed_sql sql.NullBool
var dummy_Harmon_closed_time time.Duration
var dummy_Harmon_closed_sort sort.Float64Slice

// Harmon_closedAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model harmon_closedAPI
type Harmon_closedAPI struct {
	gorm.Model

	models.Harmon_closed_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Harmon_closedPointersEncoding Harmon_closedPointersEncoding
}

// Harmon_closedPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Harmon_closedPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Harmon_closedDB describes a harmon_closed in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model harmon_closedDB
type Harmon_closedDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field harmon_closedDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Harmon_closedPointersEncoding
}

// Harmon_closedDBs arrays harmon_closedDBs
// swagger:response harmon_closedDBsResponse
type Harmon_closedDBs []Harmon_closedDB

// Harmon_closedDBResponse provides response
// swagger:response harmon_closedDBResponse
type Harmon_closedDBResponse struct {
	Harmon_closedDB
}

// Harmon_closedWOP is a Harmon_closed without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Harmon_closedWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Harmon_closed_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoHarmon_closedStruct struct {
	// stores Harmon_closedDB according to their gorm ID
	Map_Harmon_closedDBID_Harmon_closedDB map[uint]*Harmon_closedDB

	// stores Harmon_closedDB ID according to Harmon_closed address
	Map_Harmon_closedPtr_Harmon_closedDBID map[*models.Harmon_closed]uint

	// stores Harmon_closed according to their gorm ID
	Map_Harmon_closedDBID_Harmon_closedPtr map[uint]*models.Harmon_closed

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoHarmon_closed.stage
	return
}

func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) GetDB() *gorm.DB {
	return backRepoHarmon_closed.db
}

// GetHarmon_closedDBFromHarmon_closedPtr is a handy function to access the back repo instance from the stage instance
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) GetHarmon_closedDBFromHarmon_closedPtr(harmon_closed *models.Harmon_closed) (harmon_closedDB *Harmon_closedDB) {
	id := backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed]
	harmon_closedDB = backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[id]
	return
}

// BackRepoHarmon_closed.CommitPhaseOne commits all staged instances of Harmon_closed to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for harmon_closed := range stage.Harmon_closeds {
		backRepoHarmon_closed.CommitPhaseOneInstance(harmon_closed)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, harmon_closed := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr {
		if _, ok := stage.Harmon_closeds[harmon_closed]; !ok {
			backRepoHarmon_closed.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoHarmon_closed.CommitDeleteInstance commits deletion of Harmon_closed to the BackRepo
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CommitDeleteInstance(id uint) (Error error) {

	harmon_closed := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[id]

	// harmon_closed is not staged anymore, remove harmon_closedDB
	harmon_closedDB := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[id]
	query := backRepoHarmon_closed.db.Unscoped().Delete(&harmon_closedDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID, harmon_closed)
	delete(backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr, id)
	delete(backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB, id)

	return
}

// BackRepoHarmon_closed.CommitPhaseOneInstance commits harmon_closed staged instances of Harmon_closed to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CommitPhaseOneInstance(harmon_closed *models.Harmon_closed) (Error error) {

	// check if the harmon_closed is not commited yet
	if _, ok := backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed]; ok {
		return
	}

	// initiate harmon_closed
	var harmon_closedDB Harmon_closedDB
	harmon_closedDB.CopyBasicFieldsFromHarmon_closed(harmon_closed)

	query := backRepoHarmon_closed.db.Create(&harmon_closedDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed] = harmon_closedDB.ID
	backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID] = harmon_closed
	backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[harmon_closedDB.ID] = &harmon_closedDB

	return
}

// BackRepoHarmon_closed.CommitPhaseTwo commits all staged instances of Harmon_closed to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, harmon_closed := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr {
		backRepoHarmon_closed.CommitPhaseTwoInstance(backRepo, idx, harmon_closed)
	}

	return
}

// BackRepoHarmon_closed.CommitPhaseTwoInstance commits {{structname }} of models.Harmon_closed to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, harmon_closed *models.Harmon_closed) (Error error) {

	// fetch matching harmon_closedDB
	if harmon_closedDB, ok := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[idx]; ok {

		harmon_closedDB.CopyBasicFieldsFromHarmon_closed(harmon_closed)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoHarmon_closed.db.Save(&harmon_closedDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Harmon_closed intance %s", harmon_closed.Name))
		return err
	}

	return
}

// BackRepoHarmon_closed.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CheckoutPhaseOne() (Error error) {

	harmon_closedDBArray := make([]Harmon_closedDB, 0)
	query := backRepoHarmon_closed.db.Find(&harmon_closedDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	harmon_closedInstancesToBeRemovedFromTheStage := make(map[*models.Harmon_closed]any)
	for key, value := range backRepoHarmon_closed.stage.Harmon_closeds {
		harmon_closedInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, harmon_closedDB := range harmon_closedDBArray {
		backRepoHarmon_closed.CheckoutPhaseOneInstance(&harmon_closedDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		harmon_closed, ok := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID]
		if ok {
			delete(harmon_closedInstancesToBeRemovedFromTheStage, harmon_closed)
		}
	}

	// remove from stage and back repo's 3 maps all harmon_closeds that are not in the checkout
	for harmon_closed := range harmon_closedInstancesToBeRemovedFromTheStage {
		harmon_closed.Unstage(backRepoHarmon_closed.GetStage())

		// remove instance from the back repo 3 maps
		harmon_closedID := backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed]
		delete(backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID, harmon_closed)
		delete(backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB, harmon_closedID)
		delete(backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr, harmon_closedID)
	}

	return
}

// CheckoutPhaseOneInstance takes a harmon_closedDB that has been found in the DB, updates the backRepo and stages the
// models version of the harmon_closedDB
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CheckoutPhaseOneInstance(harmon_closedDB *Harmon_closedDB) (Error error) {

	harmon_closed, ok := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID]
	if !ok {
		harmon_closed = new(models.Harmon_closed)

		backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID] = harmon_closed
		backRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed] = harmon_closedDB.ID

		// append model store with the new element
		harmon_closed.Name = harmon_closedDB.Name_Data.String
		harmon_closed.Stage(backRepoHarmon_closed.GetStage())
	}
	harmon_closedDB.CopyBasicFieldsToHarmon_closed(harmon_closed)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	harmon_closed.Stage(backRepoHarmon_closed.GetStage())

	// preserve pointer to harmon_closedDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Harmon_closedDBID_Harmon_closedDB)[harmon_closedDB hold variable pointers
	harmon_closedDB_Data := *harmon_closedDB
	preservedPtrToHarmon_closed := &harmon_closedDB_Data
	backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[harmon_closedDB.ID] = preservedPtrToHarmon_closed

	return
}

// BackRepoHarmon_closed.CheckoutPhaseTwo Checkouts all staged instances of Harmon_closed to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, harmon_closedDB := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB {
		backRepoHarmon_closed.CheckoutPhaseTwoInstance(backRepo, harmon_closedDB)
	}
	return
}

// BackRepoHarmon_closed.CheckoutPhaseTwoInstance Checkouts staged instances of Harmon_closed to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, harmon_closedDB *Harmon_closedDB) (Error error) {

	harmon_closed := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr[harmon_closedDB.ID]

	harmon_closedDB.DecodePointers(backRepo, harmon_closed)

	return
}

func (harmon_closedDB *Harmon_closedDB) DecodePointers(backRepo *BackRepoStruct, harmon_closed *models.Harmon_closed) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitHarmon_closed allows commit of a single harmon_closed (if already staged)
func (backRepo *BackRepoStruct) CommitHarmon_closed(harmon_closed *models.Harmon_closed) {
	backRepo.BackRepoHarmon_closed.CommitPhaseOneInstance(harmon_closed)
	if id, ok := backRepo.BackRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed]; ok {
		backRepo.BackRepoHarmon_closed.CommitPhaseTwoInstance(backRepo, id, harmon_closed)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitHarmon_closed allows checkout of a single harmon_closed (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutHarmon_closed(harmon_closed *models.Harmon_closed) {
	// check if the harmon_closed is staged
	if _, ok := backRepo.BackRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed]; ok {

		if id, ok := backRepo.BackRepoHarmon_closed.Map_Harmon_closedPtr_Harmon_closedDBID[harmon_closed]; ok {
			var harmon_closedDB Harmon_closedDB
			harmon_closedDB.ID = id

			if err := backRepo.BackRepoHarmon_closed.db.First(&harmon_closedDB, id).Error; err != nil {
				log.Fatalln("CheckoutHarmon_closed : Problem with getting object with id:", id)
			}
			backRepo.BackRepoHarmon_closed.CheckoutPhaseOneInstance(&harmon_closedDB)
			backRepo.BackRepoHarmon_closed.CheckoutPhaseTwoInstance(backRepo, &harmon_closedDB)
		}
	}
}

// CopyBasicFieldsFromHarmon_closed
func (harmon_closedDB *Harmon_closedDB) CopyBasicFieldsFromHarmon_closed(harmon_closed *models.Harmon_closed) {
	// insertion point for fields commit

	harmon_closedDB.Name_Data.String = harmon_closed.Name
	harmon_closedDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromHarmon_closed_WOP
func (harmon_closedDB *Harmon_closedDB) CopyBasicFieldsFromHarmon_closed_WOP(harmon_closed *models.Harmon_closed_WOP) {
	// insertion point for fields commit

	harmon_closedDB.Name_Data.String = harmon_closed.Name
	harmon_closedDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromHarmon_closedWOP
func (harmon_closedDB *Harmon_closedDB) CopyBasicFieldsFromHarmon_closedWOP(harmon_closed *Harmon_closedWOP) {
	// insertion point for fields commit

	harmon_closedDB.Name_Data.String = harmon_closed.Name
	harmon_closedDB.Name_Data.Valid = true
}

// CopyBasicFieldsToHarmon_closed
func (harmon_closedDB *Harmon_closedDB) CopyBasicFieldsToHarmon_closed(harmon_closed *models.Harmon_closed) {
	// insertion point for checkout of basic fields (back repo to stage)
	harmon_closed.Name = harmon_closedDB.Name_Data.String
}

// CopyBasicFieldsToHarmon_closed_WOP
func (harmon_closedDB *Harmon_closedDB) CopyBasicFieldsToHarmon_closed_WOP(harmon_closed *models.Harmon_closed_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	harmon_closed.Name = harmon_closedDB.Name_Data.String
}

// CopyBasicFieldsToHarmon_closedWOP
func (harmon_closedDB *Harmon_closedDB) CopyBasicFieldsToHarmon_closedWOP(harmon_closed *Harmon_closedWOP) {
	harmon_closed.ID = int(harmon_closedDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	harmon_closed.Name = harmon_closedDB.Name_Data.String
}

// Backup generates a json file from a slice of all Harmon_closedDB instances in the backrepo
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Harmon_closedDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Harmon_closedDB, 0)
	for _, harmon_closedDB := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB {
		forBackup = append(forBackup, harmon_closedDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Harmon_closed ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Harmon_closed file", err.Error())
	}
}

// Backup generates a json file from a slice of all Harmon_closedDB instances in the backrepo
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Harmon_closedDB, 0)
	for _, harmon_closedDB := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB {
		forBackup = append(forBackup, harmon_closedDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Harmon_closed")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Harmon_closed_Fields, -1)
	for _, harmon_closedDB := range forBackup {

		var harmon_closedWOP Harmon_closedWOP
		harmon_closedDB.CopyBasicFieldsToHarmon_closedWOP(&harmon_closedWOP)

		row := sh.AddRow()
		row.WriteStruct(&harmon_closedWOP, -1)
	}
}

// RestoreXL from the "Harmon_closed" sheet all Harmon_closedDB instances
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoHarmon_closedid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Harmon_closed"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoHarmon_closed.rowVisitorHarmon_closed)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) rowVisitorHarmon_closed(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var harmon_closedWOP Harmon_closedWOP
		row.ReadStruct(&harmon_closedWOP)

		// add the unmarshalled struct to the stage
		harmon_closedDB := new(Harmon_closedDB)
		harmon_closedDB.CopyBasicFieldsFromHarmon_closedWOP(&harmon_closedWOP)

		harmon_closedDB_ID_atBackupTime := harmon_closedDB.ID
		harmon_closedDB.ID = 0
		query := backRepoHarmon_closed.db.Create(harmon_closedDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[harmon_closedDB.ID] = harmon_closedDB
		BackRepoHarmon_closedid_atBckpTime_newID[harmon_closedDB_ID_atBackupTime] = harmon_closedDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Harmon_closedDB.json" in dirPath that stores an array
// of Harmon_closedDB and stores it in the database
// the map BackRepoHarmon_closedid_atBckpTime_newID is updated accordingly
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoHarmon_closedid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Harmon_closedDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Harmon_closed file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Harmon_closedDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Harmon_closedDBID_Harmon_closedDB
	for _, harmon_closedDB := range forRestore {

		harmon_closedDB_ID_atBackupTime := harmon_closedDB.ID
		harmon_closedDB.ID = 0
		query := backRepoHarmon_closed.db.Create(harmon_closedDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[harmon_closedDB.ID] = harmon_closedDB
		BackRepoHarmon_closedid_atBckpTime_newID[harmon_closedDB_ID_atBackupTime] = harmon_closedDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Harmon_closed file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Harmon_closed>id_atBckpTime_newID
// to compute new index
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) RestorePhaseTwo() {

	for _, harmon_closedDB := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB {

		// next line of code is to avert unused variable compilation error
		_ = harmon_closedDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoHarmon_closed.db.Model(harmon_closedDB).Updates(*harmon_closedDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoHarmon_closed.ResetReversePointers commits all staged instances of Harmon_closed to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, harmon_closed := range backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedPtr {
		backRepoHarmon_closed.ResetReversePointersInstance(backRepo, idx, harmon_closed)
	}

	return
}

func (backRepoHarmon_closed *BackRepoHarmon_closedStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, harmon_closed *models.Harmon_closed) (Error error) {

	// fetch matching harmon_closedDB
	if harmon_closedDB, ok := backRepoHarmon_closed.Map_Harmon_closedDBID_Harmon_closedDB[idx]; ok {
		_ = harmon_closedDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoHarmon_closedid_atBckpTime_newID map[uint]uint
