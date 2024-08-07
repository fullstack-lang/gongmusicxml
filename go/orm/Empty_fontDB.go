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
var dummy_Empty_font_sql sql.NullBool
var dummy_Empty_font_time time.Duration
var dummy_Empty_font_sort sort.Float64Slice

// Empty_fontAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model empty_fontAPI
type Empty_fontAPI struct {
	gorm.Model

	models.Empty_font_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Empty_fontPointersEncoding Empty_fontPointersEncoding
}

// Empty_fontPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Empty_fontPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Empty_fontDB describes a empty_font in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model empty_fontDB
type Empty_fontDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field empty_fontDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Empty_fontPointersEncoding
}

// Empty_fontDBs arrays empty_fontDBs
// swagger:response empty_fontDBsResponse
type Empty_fontDBs []Empty_fontDB

// Empty_fontDBResponse provides response
// swagger:response empty_fontDBResponse
type Empty_fontDBResponse struct {
	Empty_fontDB
}

// Empty_fontWOP is a Empty_font without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Empty_fontWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Empty_font_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoEmpty_fontStruct struct {
	// stores Empty_fontDB according to their gorm ID
	Map_Empty_fontDBID_Empty_fontDB map[uint]*Empty_fontDB

	// stores Empty_fontDB ID according to Empty_font address
	Map_Empty_fontPtr_Empty_fontDBID map[*models.Empty_font]uint

	// stores Empty_font according to their gorm ID
	Map_Empty_fontDBID_Empty_fontPtr map[uint]*models.Empty_font

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoEmpty_font *BackRepoEmpty_fontStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoEmpty_font.stage
	return
}

func (backRepoEmpty_font *BackRepoEmpty_fontStruct) GetDB() *gorm.DB {
	return backRepoEmpty_font.db
}

// GetEmpty_fontDBFromEmpty_fontPtr is a handy function to access the back repo instance from the stage instance
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) GetEmpty_fontDBFromEmpty_fontPtr(empty_font *models.Empty_font) (empty_fontDB *Empty_fontDB) {
	id := backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font]
	empty_fontDB = backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[id]
	return
}

// BackRepoEmpty_font.CommitPhaseOne commits all staged instances of Empty_font to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for empty_font := range stage.Empty_fonts {
		backRepoEmpty_font.CommitPhaseOneInstance(empty_font)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, empty_font := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr {
		if _, ok := stage.Empty_fonts[empty_font]; !ok {
			backRepoEmpty_font.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoEmpty_font.CommitDeleteInstance commits deletion of Empty_font to the BackRepo
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CommitDeleteInstance(id uint) (Error error) {

	empty_font := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[id]

	// empty_font is not staged anymore, remove empty_fontDB
	empty_fontDB := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[id]
	query := backRepoEmpty_font.db.Unscoped().Delete(&empty_fontDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID, empty_font)
	delete(backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr, id)
	delete(backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB, id)

	return
}

// BackRepoEmpty_font.CommitPhaseOneInstance commits empty_font staged instances of Empty_font to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CommitPhaseOneInstance(empty_font *models.Empty_font) (Error error) {

	// check if the empty_font is not commited yet
	if _, ok := backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font]; ok {
		return
	}

	// initiate empty_font
	var empty_fontDB Empty_fontDB
	empty_fontDB.CopyBasicFieldsFromEmpty_font(empty_font)

	query := backRepoEmpty_font.db.Create(&empty_fontDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font] = empty_fontDB.ID
	backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID] = empty_font
	backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[empty_fontDB.ID] = &empty_fontDB

	return
}

// BackRepoEmpty_font.CommitPhaseTwo commits all staged instances of Empty_font to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, empty_font := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr {
		backRepoEmpty_font.CommitPhaseTwoInstance(backRepo, idx, empty_font)
	}

	return
}

// BackRepoEmpty_font.CommitPhaseTwoInstance commits {{structname }} of models.Empty_font to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, empty_font *models.Empty_font) (Error error) {

	// fetch matching empty_fontDB
	if empty_fontDB, ok := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[idx]; ok {

		empty_fontDB.CopyBasicFieldsFromEmpty_font(empty_font)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoEmpty_font.db.Save(&empty_fontDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Empty_font intance %s", empty_font.Name))
		return err
	}

	return
}

// BackRepoEmpty_font.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CheckoutPhaseOne() (Error error) {

	empty_fontDBArray := make([]Empty_fontDB, 0)
	query := backRepoEmpty_font.db.Find(&empty_fontDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	empty_fontInstancesToBeRemovedFromTheStage := make(map[*models.Empty_font]any)
	for key, value := range backRepoEmpty_font.stage.Empty_fonts {
		empty_fontInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, empty_fontDB := range empty_fontDBArray {
		backRepoEmpty_font.CheckoutPhaseOneInstance(&empty_fontDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		empty_font, ok := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID]
		if ok {
			delete(empty_fontInstancesToBeRemovedFromTheStage, empty_font)
		}
	}

	// remove from stage and back repo's 3 maps all empty_fonts that are not in the checkout
	for empty_font := range empty_fontInstancesToBeRemovedFromTheStage {
		empty_font.Unstage(backRepoEmpty_font.GetStage())

		// remove instance from the back repo 3 maps
		empty_fontID := backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font]
		delete(backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID, empty_font)
		delete(backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB, empty_fontID)
		delete(backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr, empty_fontID)
	}

	return
}

// CheckoutPhaseOneInstance takes a empty_fontDB that has been found in the DB, updates the backRepo and stages the
// models version of the empty_fontDB
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CheckoutPhaseOneInstance(empty_fontDB *Empty_fontDB) (Error error) {

	empty_font, ok := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID]
	if !ok {
		empty_font = new(models.Empty_font)

		backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID] = empty_font
		backRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font] = empty_fontDB.ID

		// append model store with the new element
		empty_font.Name = empty_fontDB.Name_Data.String
		empty_font.Stage(backRepoEmpty_font.GetStage())
	}
	empty_fontDB.CopyBasicFieldsToEmpty_font(empty_font)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	empty_font.Stage(backRepoEmpty_font.GetStage())

	// preserve pointer to empty_fontDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Empty_fontDBID_Empty_fontDB)[empty_fontDB hold variable pointers
	empty_fontDB_Data := *empty_fontDB
	preservedPtrToEmpty_font := &empty_fontDB_Data
	backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[empty_fontDB.ID] = preservedPtrToEmpty_font

	return
}

// BackRepoEmpty_font.CheckoutPhaseTwo Checkouts all staged instances of Empty_font to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, empty_fontDB := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB {
		backRepoEmpty_font.CheckoutPhaseTwoInstance(backRepo, empty_fontDB)
	}
	return
}

// BackRepoEmpty_font.CheckoutPhaseTwoInstance Checkouts staged instances of Empty_font to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, empty_fontDB *Empty_fontDB) (Error error) {

	empty_font := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr[empty_fontDB.ID]

	empty_fontDB.DecodePointers(backRepo, empty_font)

	return
}

func (empty_fontDB *Empty_fontDB) DecodePointers(backRepo *BackRepoStruct, empty_font *models.Empty_font) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitEmpty_font allows commit of a single empty_font (if already staged)
func (backRepo *BackRepoStruct) CommitEmpty_font(empty_font *models.Empty_font) {
	backRepo.BackRepoEmpty_font.CommitPhaseOneInstance(empty_font)
	if id, ok := backRepo.BackRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font]; ok {
		backRepo.BackRepoEmpty_font.CommitPhaseTwoInstance(backRepo, id, empty_font)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitEmpty_font allows checkout of a single empty_font (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutEmpty_font(empty_font *models.Empty_font) {
	// check if the empty_font is staged
	if _, ok := backRepo.BackRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font]; ok {

		if id, ok := backRepo.BackRepoEmpty_font.Map_Empty_fontPtr_Empty_fontDBID[empty_font]; ok {
			var empty_fontDB Empty_fontDB
			empty_fontDB.ID = id

			if err := backRepo.BackRepoEmpty_font.db.First(&empty_fontDB, id).Error; err != nil {
				log.Fatalln("CheckoutEmpty_font : Problem with getting object with id:", id)
			}
			backRepo.BackRepoEmpty_font.CheckoutPhaseOneInstance(&empty_fontDB)
			backRepo.BackRepoEmpty_font.CheckoutPhaseTwoInstance(backRepo, &empty_fontDB)
		}
	}
}

// CopyBasicFieldsFromEmpty_font
func (empty_fontDB *Empty_fontDB) CopyBasicFieldsFromEmpty_font(empty_font *models.Empty_font) {
	// insertion point for fields commit

	empty_fontDB.Name_Data.String = empty_font.Name
	empty_fontDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromEmpty_font_WOP
func (empty_fontDB *Empty_fontDB) CopyBasicFieldsFromEmpty_font_WOP(empty_font *models.Empty_font_WOP) {
	// insertion point for fields commit

	empty_fontDB.Name_Data.String = empty_font.Name
	empty_fontDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromEmpty_fontWOP
func (empty_fontDB *Empty_fontDB) CopyBasicFieldsFromEmpty_fontWOP(empty_font *Empty_fontWOP) {
	// insertion point for fields commit

	empty_fontDB.Name_Data.String = empty_font.Name
	empty_fontDB.Name_Data.Valid = true
}

// CopyBasicFieldsToEmpty_font
func (empty_fontDB *Empty_fontDB) CopyBasicFieldsToEmpty_font(empty_font *models.Empty_font) {
	// insertion point for checkout of basic fields (back repo to stage)
	empty_font.Name = empty_fontDB.Name_Data.String
}

// CopyBasicFieldsToEmpty_font_WOP
func (empty_fontDB *Empty_fontDB) CopyBasicFieldsToEmpty_font_WOP(empty_font *models.Empty_font_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	empty_font.Name = empty_fontDB.Name_Data.String
}

// CopyBasicFieldsToEmpty_fontWOP
func (empty_fontDB *Empty_fontDB) CopyBasicFieldsToEmpty_fontWOP(empty_font *Empty_fontWOP) {
	empty_font.ID = int(empty_fontDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	empty_font.Name = empty_fontDB.Name_Data.String
}

// Backup generates a json file from a slice of all Empty_fontDB instances in the backrepo
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Empty_fontDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Empty_fontDB, 0)
	for _, empty_fontDB := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB {
		forBackup = append(forBackup, empty_fontDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Empty_font ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Empty_font file", err.Error())
	}
}

// Backup generates a json file from a slice of all Empty_fontDB instances in the backrepo
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Empty_fontDB, 0)
	for _, empty_fontDB := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB {
		forBackup = append(forBackup, empty_fontDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Empty_font")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Empty_font_Fields, -1)
	for _, empty_fontDB := range forBackup {

		var empty_fontWOP Empty_fontWOP
		empty_fontDB.CopyBasicFieldsToEmpty_fontWOP(&empty_fontWOP)

		row := sh.AddRow()
		row.WriteStruct(&empty_fontWOP, -1)
	}
}

// RestoreXL from the "Empty_font" sheet all Empty_fontDB instances
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoEmpty_fontid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Empty_font"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoEmpty_font.rowVisitorEmpty_font)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoEmpty_font *BackRepoEmpty_fontStruct) rowVisitorEmpty_font(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var empty_fontWOP Empty_fontWOP
		row.ReadStruct(&empty_fontWOP)

		// add the unmarshalled struct to the stage
		empty_fontDB := new(Empty_fontDB)
		empty_fontDB.CopyBasicFieldsFromEmpty_fontWOP(&empty_fontWOP)

		empty_fontDB_ID_atBackupTime := empty_fontDB.ID
		empty_fontDB.ID = 0
		query := backRepoEmpty_font.db.Create(empty_fontDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[empty_fontDB.ID] = empty_fontDB
		BackRepoEmpty_fontid_atBckpTime_newID[empty_fontDB_ID_atBackupTime] = empty_fontDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Empty_fontDB.json" in dirPath that stores an array
// of Empty_fontDB and stores it in the database
// the map BackRepoEmpty_fontid_atBckpTime_newID is updated accordingly
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoEmpty_fontid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Empty_fontDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Empty_font file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Empty_fontDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Empty_fontDBID_Empty_fontDB
	for _, empty_fontDB := range forRestore {

		empty_fontDB_ID_atBackupTime := empty_fontDB.ID
		empty_fontDB.ID = 0
		query := backRepoEmpty_font.db.Create(empty_fontDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[empty_fontDB.ID] = empty_fontDB
		BackRepoEmpty_fontid_atBckpTime_newID[empty_fontDB_ID_atBackupTime] = empty_fontDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Empty_font file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Empty_font>id_atBckpTime_newID
// to compute new index
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) RestorePhaseTwo() {

	for _, empty_fontDB := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB {

		// next line of code is to avert unused variable compilation error
		_ = empty_fontDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoEmpty_font.db.Model(empty_fontDB).Updates(*empty_fontDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoEmpty_font.ResetReversePointers commits all staged instances of Empty_font to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoEmpty_font *BackRepoEmpty_fontStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, empty_font := range backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontPtr {
		backRepoEmpty_font.ResetReversePointersInstance(backRepo, idx, empty_font)
	}

	return
}

func (backRepoEmpty_font *BackRepoEmpty_fontStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, empty_font *models.Empty_font) (Error error) {

	// fetch matching empty_fontDB
	if empty_fontDB, ok := backRepoEmpty_font.Map_Empty_fontDBID_Empty_fontDB[idx]; ok {
		_ = empty_fontDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoEmpty_fontid_atBckpTime_newID map[uint]uint
