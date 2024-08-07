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
var dummy_Numeral_key_sql sql.NullBool
var dummy_Numeral_key_time time.Duration
var dummy_Numeral_key_sort sort.Float64Slice

// Numeral_keyAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model numeral_keyAPI
type Numeral_keyAPI struct {
	gorm.Model

	models.Numeral_key_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Numeral_keyPointersEncoding Numeral_keyPointersEncoding
}

// Numeral_keyPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Numeral_keyPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Numeral_keyDB describes a numeral_key in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model numeral_keyDB
type Numeral_keyDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field numeral_keyDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Numeral_keyPointersEncoding
}

// Numeral_keyDBs arrays numeral_keyDBs
// swagger:response numeral_keyDBsResponse
type Numeral_keyDBs []Numeral_keyDB

// Numeral_keyDBResponse provides response
// swagger:response numeral_keyDBResponse
type Numeral_keyDBResponse struct {
	Numeral_keyDB
}

// Numeral_keyWOP is a Numeral_key without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Numeral_keyWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Numeral_key_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoNumeral_keyStruct struct {
	// stores Numeral_keyDB according to their gorm ID
	Map_Numeral_keyDBID_Numeral_keyDB map[uint]*Numeral_keyDB

	// stores Numeral_keyDB ID according to Numeral_key address
	Map_Numeral_keyPtr_Numeral_keyDBID map[*models.Numeral_key]uint

	// stores Numeral_key according to their gorm ID
	Map_Numeral_keyDBID_Numeral_keyPtr map[uint]*models.Numeral_key

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoNumeral_key *BackRepoNumeral_keyStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoNumeral_key.stage
	return
}

func (backRepoNumeral_key *BackRepoNumeral_keyStruct) GetDB() *gorm.DB {
	return backRepoNumeral_key.db
}

// GetNumeral_keyDBFromNumeral_keyPtr is a handy function to access the back repo instance from the stage instance
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) GetNumeral_keyDBFromNumeral_keyPtr(numeral_key *models.Numeral_key) (numeral_keyDB *Numeral_keyDB) {
	id := backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key]
	numeral_keyDB = backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[id]
	return
}

// BackRepoNumeral_key.CommitPhaseOne commits all staged instances of Numeral_key to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for numeral_key := range stage.Numeral_keys {
		backRepoNumeral_key.CommitPhaseOneInstance(numeral_key)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, numeral_key := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr {
		if _, ok := stage.Numeral_keys[numeral_key]; !ok {
			backRepoNumeral_key.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoNumeral_key.CommitDeleteInstance commits deletion of Numeral_key to the BackRepo
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CommitDeleteInstance(id uint) (Error error) {

	numeral_key := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[id]

	// numeral_key is not staged anymore, remove numeral_keyDB
	numeral_keyDB := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[id]
	query := backRepoNumeral_key.db.Unscoped().Delete(&numeral_keyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID, numeral_key)
	delete(backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr, id)
	delete(backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB, id)

	return
}

// BackRepoNumeral_key.CommitPhaseOneInstance commits numeral_key staged instances of Numeral_key to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CommitPhaseOneInstance(numeral_key *models.Numeral_key) (Error error) {

	// check if the numeral_key is not commited yet
	if _, ok := backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key]; ok {
		return
	}

	// initiate numeral_key
	var numeral_keyDB Numeral_keyDB
	numeral_keyDB.CopyBasicFieldsFromNumeral_key(numeral_key)

	query := backRepoNumeral_key.db.Create(&numeral_keyDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key] = numeral_keyDB.ID
	backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID] = numeral_key
	backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[numeral_keyDB.ID] = &numeral_keyDB

	return
}

// BackRepoNumeral_key.CommitPhaseTwo commits all staged instances of Numeral_key to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, numeral_key := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr {
		backRepoNumeral_key.CommitPhaseTwoInstance(backRepo, idx, numeral_key)
	}

	return
}

// BackRepoNumeral_key.CommitPhaseTwoInstance commits {{structname }} of models.Numeral_key to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, numeral_key *models.Numeral_key) (Error error) {

	// fetch matching numeral_keyDB
	if numeral_keyDB, ok := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[idx]; ok {

		numeral_keyDB.CopyBasicFieldsFromNumeral_key(numeral_key)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoNumeral_key.db.Save(&numeral_keyDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Numeral_key intance %s", numeral_key.Name))
		return err
	}

	return
}

// BackRepoNumeral_key.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CheckoutPhaseOne() (Error error) {

	numeral_keyDBArray := make([]Numeral_keyDB, 0)
	query := backRepoNumeral_key.db.Find(&numeral_keyDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	numeral_keyInstancesToBeRemovedFromTheStage := make(map[*models.Numeral_key]any)
	for key, value := range backRepoNumeral_key.stage.Numeral_keys {
		numeral_keyInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, numeral_keyDB := range numeral_keyDBArray {
		backRepoNumeral_key.CheckoutPhaseOneInstance(&numeral_keyDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		numeral_key, ok := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID]
		if ok {
			delete(numeral_keyInstancesToBeRemovedFromTheStage, numeral_key)
		}
	}

	// remove from stage and back repo's 3 maps all numeral_keys that are not in the checkout
	for numeral_key := range numeral_keyInstancesToBeRemovedFromTheStage {
		numeral_key.Unstage(backRepoNumeral_key.GetStage())

		// remove instance from the back repo 3 maps
		numeral_keyID := backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key]
		delete(backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID, numeral_key)
		delete(backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB, numeral_keyID)
		delete(backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr, numeral_keyID)
	}

	return
}

// CheckoutPhaseOneInstance takes a numeral_keyDB that has been found in the DB, updates the backRepo and stages the
// models version of the numeral_keyDB
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CheckoutPhaseOneInstance(numeral_keyDB *Numeral_keyDB) (Error error) {

	numeral_key, ok := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID]
	if !ok {
		numeral_key = new(models.Numeral_key)

		backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID] = numeral_key
		backRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key] = numeral_keyDB.ID

		// append model store with the new element
		numeral_key.Name = numeral_keyDB.Name_Data.String
		numeral_key.Stage(backRepoNumeral_key.GetStage())
	}
	numeral_keyDB.CopyBasicFieldsToNumeral_key(numeral_key)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	numeral_key.Stage(backRepoNumeral_key.GetStage())

	// preserve pointer to numeral_keyDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Numeral_keyDBID_Numeral_keyDB)[numeral_keyDB hold variable pointers
	numeral_keyDB_Data := *numeral_keyDB
	preservedPtrToNumeral_key := &numeral_keyDB_Data
	backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[numeral_keyDB.ID] = preservedPtrToNumeral_key

	return
}

// BackRepoNumeral_key.CheckoutPhaseTwo Checkouts all staged instances of Numeral_key to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, numeral_keyDB := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB {
		backRepoNumeral_key.CheckoutPhaseTwoInstance(backRepo, numeral_keyDB)
	}
	return
}

// BackRepoNumeral_key.CheckoutPhaseTwoInstance Checkouts staged instances of Numeral_key to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, numeral_keyDB *Numeral_keyDB) (Error error) {

	numeral_key := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr[numeral_keyDB.ID]

	numeral_keyDB.DecodePointers(backRepo, numeral_key)

	return
}

func (numeral_keyDB *Numeral_keyDB) DecodePointers(backRepo *BackRepoStruct, numeral_key *models.Numeral_key) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitNumeral_key allows commit of a single numeral_key (if already staged)
func (backRepo *BackRepoStruct) CommitNumeral_key(numeral_key *models.Numeral_key) {
	backRepo.BackRepoNumeral_key.CommitPhaseOneInstance(numeral_key)
	if id, ok := backRepo.BackRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key]; ok {
		backRepo.BackRepoNumeral_key.CommitPhaseTwoInstance(backRepo, id, numeral_key)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitNumeral_key allows checkout of a single numeral_key (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutNumeral_key(numeral_key *models.Numeral_key) {
	// check if the numeral_key is staged
	if _, ok := backRepo.BackRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key]; ok {

		if id, ok := backRepo.BackRepoNumeral_key.Map_Numeral_keyPtr_Numeral_keyDBID[numeral_key]; ok {
			var numeral_keyDB Numeral_keyDB
			numeral_keyDB.ID = id

			if err := backRepo.BackRepoNumeral_key.db.First(&numeral_keyDB, id).Error; err != nil {
				log.Fatalln("CheckoutNumeral_key : Problem with getting object with id:", id)
			}
			backRepo.BackRepoNumeral_key.CheckoutPhaseOneInstance(&numeral_keyDB)
			backRepo.BackRepoNumeral_key.CheckoutPhaseTwoInstance(backRepo, &numeral_keyDB)
		}
	}
}

// CopyBasicFieldsFromNumeral_key
func (numeral_keyDB *Numeral_keyDB) CopyBasicFieldsFromNumeral_key(numeral_key *models.Numeral_key) {
	// insertion point for fields commit

	numeral_keyDB.Name_Data.String = numeral_key.Name
	numeral_keyDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromNumeral_key_WOP
func (numeral_keyDB *Numeral_keyDB) CopyBasicFieldsFromNumeral_key_WOP(numeral_key *models.Numeral_key_WOP) {
	// insertion point for fields commit

	numeral_keyDB.Name_Data.String = numeral_key.Name
	numeral_keyDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromNumeral_keyWOP
func (numeral_keyDB *Numeral_keyDB) CopyBasicFieldsFromNumeral_keyWOP(numeral_key *Numeral_keyWOP) {
	// insertion point for fields commit

	numeral_keyDB.Name_Data.String = numeral_key.Name
	numeral_keyDB.Name_Data.Valid = true
}

// CopyBasicFieldsToNumeral_key
func (numeral_keyDB *Numeral_keyDB) CopyBasicFieldsToNumeral_key(numeral_key *models.Numeral_key) {
	// insertion point for checkout of basic fields (back repo to stage)
	numeral_key.Name = numeral_keyDB.Name_Data.String
}

// CopyBasicFieldsToNumeral_key_WOP
func (numeral_keyDB *Numeral_keyDB) CopyBasicFieldsToNumeral_key_WOP(numeral_key *models.Numeral_key_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	numeral_key.Name = numeral_keyDB.Name_Data.String
}

// CopyBasicFieldsToNumeral_keyWOP
func (numeral_keyDB *Numeral_keyDB) CopyBasicFieldsToNumeral_keyWOP(numeral_key *Numeral_keyWOP) {
	numeral_key.ID = int(numeral_keyDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	numeral_key.Name = numeral_keyDB.Name_Data.String
}

// Backup generates a json file from a slice of all Numeral_keyDB instances in the backrepo
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Numeral_keyDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Numeral_keyDB, 0)
	for _, numeral_keyDB := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB {
		forBackup = append(forBackup, numeral_keyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Numeral_key ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Numeral_key file", err.Error())
	}
}

// Backup generates a json file from a slice of all Numeral_keyDB instances in the backrepo
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Numeral_keyDB, 0)
	for _, numeral_keyDB := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB {
		forBackup = append(forBackup, numeral_keyDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Numeral_key")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Numeral_key_Fields, -1)
	for _, numeral_keyDB := range forBackup {

		var numeral_keyWOP Numeral_keyWOP
		numeral_keyDB.CopyBasicFieldsToNumeral_keyWOP(&numeral_keyWOP)

		row := sh.AddRow()
		row.WriteStruct(&numeral_keyWOP, -1)
	}
}

// RestoreXL from the "Numeral_key" sheet all Numeral_keyDB instances
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoNumeral_keyid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Numeral_key"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoNumeral_key.rowVisitorNumeral_key)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoNumeral_key *BackRepoNumeral_keyStruct) rowVisitorNumeral_key(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var numeral_keyWOP Numeral_keyWOP
		row.ReadStruct(&numeral_keyWOP)

		// add the unmarshalled struct to the stage
		numeral_keyDB := new(Numeral_keyDB)
		numeral_keyDB.CopyBasicFieldsFromNumeral_keyWOP(&numeral_keyWOP)

		numeral_keyDB_ID_atBackupTime := numeral_keyDB.ID
		numeral_keyDB.ID = 0
		query := backRepoNumeral_key.db.Create(numeral_keyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[numeral_keyDB.ID] = numeral_keyDB
		BackRepoNumeral_keyid_atBckpTime_newID[numeral_keyDB_ID_atBackupTime] = numeral_keyDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Numeral_keyDB.json" in dirPath that stores an array
// of Numeral_keyDB and stores it in the database
// the map BackRepoNumeral_keyid_atBckpTime_newID is updated accordingly
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoNumeral_keyid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Numeral_keyDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Numeral_key file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Numeral_keyDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Numeral_keyDBID_Numeral_keyDB
	for _, numeral_keyDB := range forRestore {

		numeral_keyDB_ID_atBackupTime := numeral_keyDB.ID
		numeral_keyDB.ID = 0
		query := backRepoNumeral_key.db.Create(numeral_keyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[numeral_keyDB.ID] = numeral_keyDB
		BackRepoNumeral_keyid_atBckpTime_newID[numeral_keyDB_ID_atBackupTime] = numeral_keyDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Numeral_key file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Numeral_key>id_atBckpTime_newID
// to compute new index
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) RestorePhaseTwo() {

	for _, numeral_keyDB := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB {

		// next line of code is to avert unused variable compilation error
		_ = numeral_keyDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoNumeral_key.db.Model(numeral_keyDB).Updates(*numeral_keyDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoNumeral_key.ResetReversePointers commits all staged instances of Numeral_key to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoNumeral_key *BackRepoNumeral_keyStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, numeral_key := range backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyPtr {
		backRepoNumeral_key.ResetReversePointersInstance(backRepo, idx, numeral_key)
	}

	return
}

func (backRepoNumeral_key *BackRepoNumeral_keyStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, numeral_key *models.Numeral_key) (Error error) {

	// fetch matching numeral_keyDB
	if numeral_keyDB, ok := backRepoNumeral_key.Map_Numeral_keyDBID_Numeral_keyDB[idx]; ok {
		_ = numeral_keyDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoNumeral_keyid_atBckpTime_newID map[uint]uint
