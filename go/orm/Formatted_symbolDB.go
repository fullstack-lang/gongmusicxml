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
var dummy_Formatted_symbol_sql sql.NullBool
var dummy_Formatted_symbol_time time.Duration
var dummy_Formatted_symbol_sort sort.Float64Slice

// Formatted_symbolAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model formatted_symbolAPI
type Formatted_symbolAPI struct {
	gorm.Model

	models.Formatted_symbol_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Formatted_symbolPointersEncoding Formatted_symbolPointersEncoding
}

// Formatted_symbolPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Formatted_symbolPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Formatted_symbolDB describes a formatted_symbol in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model formatted_symbolDB
type Formatted_symbolDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field formatted_symbolDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Formatted_symbolPointersEncoding
}

// Formatted_symbolDBs arrays formatted_symbolDBs
// swagger:response formatted_symbolDBsResponse
type Formatted_symbolDBs []Formatted_symbolDB

// Formatted_symbolDBResponse provides response
// swagger:response formatted_symbolDBResponse
type Formatted_symbolDBResponse struct {
	Formatted_symbolDB
}

// Formatted_symbolWOP is a Formatted_symbol without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Formatted_symbolWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Formatted_symbol_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoFormatted_symbolStruct struct {
	// stores Formatted_symbolDB according to their gorm ID
	Map_Formatted_symbolDBID_Formatted_symbolDB map[uint]*Formatted_symbolDB

	// stores Formatted_symbolDB ID according to Formatted_symbol address
	Map_Formatted_symbolPtr_Formatted_symbolDBID map[*models.Formatted_symbol]uint

	// stores Formatted_symbol according to their gorm ID
	Map_Formatted_symbolDBID_Formatted_symbolPtr map[uint]*models.Formatted_symbol

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoFormatted_symbol.stage
	return
}

func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) GetDB() *gorm.DB {
	return backRepoFormatted_symbol.db
}

// GetFormatted_symbolDBFromFormatted_symbolPtr is a handy function to access the back repo instance from the stage instance
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) GetFormatted_symbolDBFromFormatted_symbolPtr(formatted_symbol *models.Formatted_symbol) (formatted_symbolDB *Formatted_symbolDB) {
	id := backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol]
	formatted_symbolDB = backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[id]
	return
}

// BackRepoFormatted_symbol.CommitPhaseOne commits all staged instances of Formatted_symbol to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for formatted_symbol := range stage.Formatted_symbols {
		backRepoFormatted_symbol.CommitPhaseOneInstance(formatted_symbol)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, formatted_symbol := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr {
		if _, ok := stage.Formatted_symbols[formatted_symbol]; !ok {
			backRepoFormatted_symbol.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoFormatted_symbol.CommitDeleteInstance commits deletion of Formatted_symbol to the BackRepo
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CommitDeleteInstance(id uint) (Error error) {

	formatted_symbol := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[id]

	// formatted_symbol is not staged anymore, remove formatted_symbolDB
	formatted_symbolDB := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[id]
	query := backRepoFormatted_symbol.db.Unscoped().Delete(&formatted_symbolDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID, formatted_symbol)
	delete(backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr, id)
	delete(backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB, id)

	return
}

// BackRepoFormatted_symbol.CommitPhaseOneInstance commits formatted_symbol staged instances of Formatted_symbol to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CommitPhaseOneInstance(formatted_symbol *models.Formatted_symbol) (Error error) {

	// check if the formatted_symbol is not commited yet
	if _, ok := backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol]; ok {
		return
	}

	// initiate formatted_symbol
	var formatted_symbolDB Formatted_symbolDB
	formatted_symbolDB.CopyBasicFieldsFromFormatted_symbol(formatted_symbol)

	query := backRepoFormatted_symbol.db.Create(&formatted_symbolDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol] = formatted_symbolDB.ID
	backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID] = formatted_symbol
	backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[formatted_symbolDB.ID] = &formatted_symbolDB

	return
}

// BackRepoFormatted_symbol.CommitPhaseTwo commits all staged instances of Formatted_symbol to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, formatted_symbol := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr {
		backRepoFormatted_symbol.CommitPhaseTwoInstance(backRepo, idx, formatted_symbol)
	}

	return
}

// BackRepoFormatted_symbol.CommitPhaseTwoInstance commits {{structname }} of models.Formatted_symbol to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, formatted_symbol *models.Formatted_symbol) (Error error) {

	// fetch matching formatted_symbolDB
	if formatted_symbolDB, ok := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[idx]; ok {

		formatted_symbolDB.CopyBasicFieldsFromFormatted_symbol(formatted_symbol)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoFormatted_symbol.db.Save(&formatted_symbolDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Formatted_symbol intance %s", formatted_symbol.Name))
		return err
	}

	return
}

// BackRepoFormatted_symbol.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CheckoutPhaseOne() (Error error) {

	formatted_symbolDBArray := make([]Formatted_symbolDB, 0)
	query := backRepoFormatted_symbol.db.Find(&formatted_symbolDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	formatted_symbolInstancesToBeRemovedFromTheStage := make(map[*models.Formatted_symbol]any)
	for key, value := range backRepoFormatted_symbol.stage.Formatted_symbols {
		formatted_symbolInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, formatted_symbolDB := range formatted_symbolDBArray {
		backRepoFormatted_symbol.CheckoutPhaseOneInstance(&formatted_symbolDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		formatted_symbol, ok := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID]
		if ok {
			delete(formatted_symbolInstancesToBeRemovedFromTheStage, formatted_symbol)
		}
	}

	// remove from stage and back repo's 3 maps all formatted_symbols that are not in the checkout
	for formatted_symbol := range formatted_symbolInstancesToBeRemovedFromTheStage {
		formatted_symbol.Unstage(backRepoFormatted_symbol.GetStage())

		// remove instance from the back repo 3 maps
		formatted_symbolID := backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol]
		delete(backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID, formatted_symbol)
		delete(backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB, formatted_symbolID)
		delete(backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr, formatted_symbolID)
	}

	return
}

// CheckoutPhaseOneInstance takes a formatted_symbolDB that has been found in the DB, updates the backRepo and stages the
// models version of the formatted_symbolDB
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CheckoutPhaseOneInstance(formatted_symbolDB *Formatted_symbolDB) (Error error) {

	formatted_symbol, ok := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID]
	if !ok {
		formatted_symbol = new(models.Formatted_symbol)

		backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID] = formatted_symbol
		backRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol] = formatted_symbolDB.ID

		// append model store with the new element
		formatted_symbol.Name = formatted_symbolDB.Name_Data.String
		formatted_symbol.Stage(backRepoFormatted_symbol.GetStage())
	}
	formatted_symbolDB.CopyBasicFieldsToFormatted_symbol(formatted_symbol)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	formatted_symbol.Stage(backRepoFormatted_symbol.GetStage())

	// preserve pointer to formatted_symbolDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Formatted_symbolDBID_Formatted_symbolDB)[formatted_symbolDB hold variable pointers
	formatted_symbolDB_Data := *formatted_symbolDB
	preservedPtrToFormatted_symbol := &formatted_symbolDB_Data
	backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[formatted_symbolDB.ID] = preservedPtrToFormatted_symbol

	return
}

// BackRepoFormatted_symbol.CheckoutPhaseTwo Checkouts all staged instances of Formatted_symbol to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, formatted_symbolDB := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB {
		backRepoFormatted_symbol.CheckoutPhaseTwoInstance(backRepo, formatted_symbolDB)
	}
	return
}

// BackRepoFormatted_symbol.CheckoutPhaseTwoInstance Checkouts staged instances of Formatted_symbol to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, formatted_symbolDB *Formatted_symbolDB) (Error error) {

	formatted_symbol := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr[formatted_symbolDB.ID]

	formatted_symbolDB.DecodePointers(backRepo, formatted_symbol)

	return
}

func (formatted_symbolDB *Formatted_symbolDB) DecodePointers(backRepo *BackRepoStruct, formatted_symbol *models.Formatted_symbol) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitFormatted_symbol allows commit of a single formatted_symbol (if already staged)
func (backRepo *BackRepoStruct) CommitFormatted_symbol(formatted_symbol *models.Formatted_symbol) {
	backRepo.BackRepoFormatted_symbol.CommitPhaseOneInstance(formatted_symbol)
	if id, ok := backRepo.BackRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol]; ok {
		backRepo.BackRepoFormatted_symbol.CommitPhaseTwoInstance(backRepo, id, formatted_symbol)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitFormatted_symbol allows checkout of a single formatted_symbol (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutFormatted_symbol(formatted_symbol *models.Formatted_symbol) {
	// check if the formatted_symbol is staged
	if _, ok := backRepo.BackRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol]; ok {

		if id, ok := backRepo.BackRepoFormatted_symbol.Map_Formatted_symbolPtr_Formatted_symbolDBID[formatted_symbol]; ok {
			var formatted_symbolDB Formatted_symbolDB
			formatted_symbolDB.ID = id

			if err := backRepo.BackRepoFormatted_symbol.db.First(&formatted_symbolDB, id).Error; err != nil {
				log.Fatalln("CheckoutFormatted_symbol : Problem with getting object with id:", id)
			}
			backRepo.BackRepoFormatted_symbol.CheckoutPhaseOneInstance(&formatted_symbolDB)
			backRepo.BackRepoFormatted_symbol.CheckoutPhaseTwoInstance(backRepo, &formatted_symbolDB)
		}
	}
}

// CopyBasicFieldsFromFormatted_symbol
func (formatted_symbolDB *Formatted_symbolDB) CopyBasicFieldsFromFormatted_symbol(formatted_symbol *models.Formatted_symbol) {
	// insertion point for fields commit

	formatted_symbolDB.Name_Data.String = formatted_symbol.Name
	formatted_symbolDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFormatted_symbol_WOP
func (formatted_symbolDB *Formatted_symbolDB) CopyBasicFieldsFromFormatted_symbol_WOP(formatted_symbol *models.Formatted_symbol_WOP) {
	// insertion point for fields commit

	formatted_symbolDB.Name_Data.String = formatted_symbol.Name
	formatted_symbolDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromFormatted_symbolWOP
func (formatted_symbolDB *Formatted_symbolDB) CopyBasicFieldsFromFormatted_symbolWOP(formatted_symbol *Formatted_symbolWOP) {
	// insertion point for fields commit

	formatted_symbolDB.Name_Data.String = formatted_symbol.Name
	formatted_symbolDB.Name_Data.Valid = true
}

// CopyBasicFieldsToFormatted_symbol
func (formatted_symbolDB *Formatted_symbolDB) CopyBasicFieldsToFormatted_symbol(formatted_symbol *models.Formatted_symbol) {
	// insertion point for checkout of basic fields (back repo to stage)
	formatted_symbol.Name = formatted_symbolDB.Name_Data.String
}

// CopyBasicFieldsToFormatted_symbol_WOP
func (formatted_symbolDB *Formatted_symbolDB) CopyBasicFieldsToFormatted_symbol_WOP(formatted_symbol *models.Formatted_symbol_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	formatted_symbol.Name = formatted_symbolDB.Name_Data.String
}

// CopyBasicFieldsToFormatted_symbolWOP
func (formatted_symbolDB *Formatted_symbolDB) CopyBasicFieldsToFormatted_symbolWOP(formatted_symbol *Formatted_symbolWOP) {
	formatted_symbol.ID = int(formatted_symbolDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	formatted_symbol.Name = formatted_symbolDB.Name_Data.String
}

// Backup generates a json file from a slice of all Formatted_symbolDB instances in the backrepo
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Formatted_symbolDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Formatted_symbolDB, 0)
	for _, formatted_symbolDB := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB {
		forBackup = append(forBackup, formatted_symbolDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Formatted_symbol ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Formatted_symbol file", err.Error())
	}
}

// Backup generates a json file from a slice of all Formatted_symbolDB instances in the backrepo
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Formatted_symbolDB, 0)
	for _, formatted_symbolDB := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB {
		forBackup = append(forBackup, formatted_symbolDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Formatted_symbol")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Formatted_symbol_Fields, -1)
	for _, formatted_symbolDB := range forBackup {

		var formatted_symbolWOP Formatted_symbolWOP
		formatted_symbolDB.CopyBasicFieldsToFormatted_symbolWOP(&formatted_symbolWOP)

		row := sh.AddRow()
		row.WriteStruct(&formatted_symbolWOP, -1)
	}
}

// RestoreXL from the "Formatted_symbol" sheet all Formatted_symbolDB instances
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoFormatted_symbolid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Formatted_symbol"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoFormatted_symbol.rowVisitorFormatted_symbol)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) rowVisitorFormatted_symbol(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var formatted_symbolWOP Formatted_symbolWOP
		row.ReadStruct(&formatted_symbolWOP)

		// add the unmarshalled struct to the stage
		formatted_symbolDB := new(Formatted_symbolDB)
		formatted_symbolDB.CopyBasicFieldsFromFormatted_symbolWOP(&formatted_symbolWOP)

		formatted_symbolDB_ID_atBackupTime := formatted_symbolDB.ID
		formatted_symbolDB.ID = 0
		query := backRepoFormatted_symbol.db.Create(formatted_symbolDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[formatted_symbolDB.ID] = formatted_symbolDB
		BackRepoFormatted_symbolid_atBckpTime_newID[formatted_symbolDB_ID_atBackupTime] = formatted_symbolDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Formatted_symbolDB.json" in dirPath that stores an array
// of Formatted_symbolDB and stores it in the database
// the map BackRepoFormatted_symbolid_atBckpTime_newID is updated accordingly
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoFormatted_symbolid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Formatted_symbolDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Formatted_symbol file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Formatted_symbolDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Formatted_symbolDBID_Formatted_symbolDB
	for _, formatted_symbolDB := range forRestore {

		formatted_symbolDB_ID_atBackupTime := formatted_symbolDB.ID
		formatted_symbolDB.ID = 0
		query := backRepoFormatted_symbol.db.Create(formatted_symbolDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[formatted_symbolDB.ID] = formatted_symbolDB
		BackRepoFormatted_symbolid_atBckpTime_newID[formatted_symbolDB_ID_atBackupTime] = formatted_symbolDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Formatted_symbol file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Formatted_symbol>id_atBckpTime_newID
// to compute new index
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) RestorePhaseTwo() {

	for _, formatted_symbolDB := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB {

		// next line of code is to avert unused variable compilation error
		_ = formatted_symbolDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoFormatted_symbol.db.Model(formatted_symbolDB).Updates(*formatted_symbolDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoFormatted_symbol.ResetReversePointers commits all staged instances of Formatted_symbol to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, formatted_symbol := range backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolPtr {
		backRepoFormatted_symbol.ResetReversePointersInstance(backRepo, idx, formatted_symbol)
	}

	return
}

func (backRepoFormatted_symbol *BackRepoFormatted_symbolStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, formatted_symbol *models.Formatted_symbol) (Error error) {

	// fetch matching formatted_symbolDB
	if formatted_symbolDB, ok := backRepoFormatted_symbol.Map_Formatted_symbolDBID_Formatted_symbolDB[idx]; ok {
		_ = formatted_symbolDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoFormatted_symbolid_atBckpTime_newID map[uint]uint
