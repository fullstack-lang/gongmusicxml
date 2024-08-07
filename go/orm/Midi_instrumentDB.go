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
var dummy_Midi_instrument_sql sql.NullBool
var dummy_Midi_instrument_time time.Duration
var dummy_Midi_instrument_sort sort.Float64Slice

// Midi_instrumentAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model midi_instrumentAPI
type Midi_instrumentAPI struct {
	gorm.Model

	models.Midi_instrument_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Midi_instrumentPointersEncoding Midi_instrumentPointersEncoding
}

// Midi_instrumentPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Midi_instrumentPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Midi_instrumentDB describes a midi_instrument in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model midi_instrumentDB
type Midi_instrumentDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field midi_instrumentDB.Name
	Name_Data sql.NullString

	// Declation for basic field midi_instrumentDB.Midi_name
	Midi_name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Midi_instrumentPointersEncoding
}

// Midi_instrumentDBs arrays midi_instrumentDBs
// swagger:response midi_instrumentDBsResponse
type Midi_instrumentDBs []Midi_instrumentDB

// Midi_instrumentDBResponse provides response
// swagger:response midi_instrumentDBResponse
type Midi_instrumentDBResponse struct {
	Midi_instrumentDB
}

// Midi_instrumentWOP is a Midi_instrument without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Midi_instrumentWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Midi_name string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Midi_instrument_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Midi_name",
}

type BackRepoMidi_instrumentStruct struct {
	// stores Midi_instrumentDB according to their gorm ID
	Map_Midi_instrumentDBID_Midi_instrumentDB map[uint]*Midi_instrumentDB

	// stores Midi_instrumentDB ID according to Midi_instrument address
	Map_Midi_instrumentPtr_Midi_instrumentDBID map[*models.Midi_instrument]uint

	// stores Midi_instrument according to their gorm ID
	Map_Midi_instrumentDBID_Midi_instrumentPtr map[uint]*models.Midi_instrument

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoMidi_instrument.stage
	return
}

func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) GetDB() *gorm.DB {
	return backRepoMidi_instrument.db
}

// GetMidi_instrumentDBFromMidi_instrumentPtr is a handy function to access the back repo instance from the stage instance
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) GetMidi_instrumentDBFromMidi_instrumentPtr(midi_instrument *models.Midi_instrument) (midi_instrumentDB *Midi_instrumentDB) {
	id := backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument]
	midi_instrumentDB = backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[id]
	return
}

// BackRepoMidi_instrument.CommitPhaseOne commits all staged instances of Midi_instrument to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for midi_instrument := range stage.Midi_instruments {
		backRepoMidi_instrument.CommitPhaseOneInstance(midi_instrument)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, midi_instrument := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr {
		if _, ok := stage.Midi_instruments[midi_instrument]; !ok {
			backRepoMidi_instrument.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMidi_instrument.CommitDeleteInstance commits deletion of Midi_instrument to the BackRepo
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CommitDeleteInstance(id uint) (Error error) {

	midi_instrument := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[id]

	// midi_instrument is not staged anymore, remove midi_instrumentDB
	midi_instrumentDB := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[id]
	query := backRepoMidi_instrument.db.Unscoped().Delete(&midi_instrumentDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID, midi_instrument)
	delete(backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr, id)
	delete(backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB, id)

	return
}

// BackRepoMidi_instrument.CommitPhaseOneInstance commits midi_instrument staged instances of Midi_instrument to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CommitPhaseOneInstance(midi_instrument *models.Midi_instrument) (Error error) {

	// check if the midi_instrument is not commited yet
	if _, ok := backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument]; ok {
		return
	}

	// initiate midi_instrument
	var midi_instrumentDB Midi_instrumentDB
	midi_instrumentDB.CopyBasicFieldsFromMidi_instrument(midi_instrument)

	query := backRepoMidi_instrument.db.Create(&midi_instrumentDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument] = midi_instrumentDB.ID
	backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID] = midi_instrument
	backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[midi_instrumentDB.ID] = &midi_instrumentDB

	return
}

// BackRepoMidi_instrument.CommitPhaseTwo commits all staged instances of Midi_instrument to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, midi_instrument := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr {
		backRepoMidi_instrument.CommitPhaseTwoInstance(backRepo, idx, midi_instrument)
	}

	return
}

// BackRepoMidi_instrument.CommitPhaseTwoInstance commits {{structname }} of models.Midi_instrument to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, midi_instrument *models.Midi_instrument) (Error error) {

	// fetch matching midi_instrumentDB
	if midi_instrumentDB, ok := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[idx]; ok {

		midi_instrumentDB.CopyBasicFieldsFromMidi_instrument(midi_instrument)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoMidi_instrument.db.Save(&midi_instrumentDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Midi_instrument intance %s", midi_instrument.Name))
		return err
	}

	return
}

// BackRepoMidi_instrument.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CheckoutPhaseOne() (Error error) {

	midi_instrumentDBArray := make([]Midi_instrumentDB, 0)
	query := backRepoMidi_instrument.db.Find(&midi_instrumentDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	midi_instrumentInstancesToBeRemovedFromTheStage := make(map[*models.Midi_instrument]any)
	for key, value := range backRepoMidi_instrument.stage.Midi_instruments {
		midi_instrumentInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, midi_instrumentDB := range midi_instrumentDBArray {
		backRepoMidi_instrument.CheckoutPhaseOneInstance(&midi_instrumentDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		midi_instrument, ok := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID]
		if ok {
			delete(midi_instrumentInstancesToBeRemovedFromTheStage, midi_instrument)
		}
	}

	// remove from stage and back repo's 3 maps all midi_instruments that are not in the checkout
	for midi_instrument := range midi_instrumentInstancesToBeRemovedFromTheStage {
		midi_instrument.Unstage(backRepoMidi_instrument.GetStage())

		// remove instance from the back repo 3 maps
		midi_instrumentID := backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument]
		delete(backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID, midi_instrument)
		delete(backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB, midi_instrumentID)
		delete(backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr, midi_instrumentID)
	}

	return
}

// CheckoutPhaseOneInstance takes a midi_instrumentDB that has been found in the DB, updates the backRepo and stages the
// models version of the midi_instrumentDB
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CheckoutPhaseOneInstance(midi_instrumentDB *Midi_instrumentDB) (Error error) {

	midi_instrument, ok := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID]
	if !ok {
		midi_instrument = new(models.Midi_instrument)

		backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID] = midi_instrument
		backRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument] = midi_instrumentDB.ID

		// append model store with the new element
		midi_instrument.Name = midi_instrumentDB.Name_Data.String
		midi_instrument.Stage(backRepoMidi_instrument.GetStage())
	}
	midi_instrumentDB.CopyBasicFieldsToMidi_instrument(midi_instrument)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	midi_instrument.Stage(backRepoMidi_instrument.GetStage())

	// preserve pointer to midi_instrumentDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Midi_instrumentDBID_Midi_instrumentDB)[midi_instrumentDB hold variable pointers
	midi_instrumentDB_Data := *midi_instrumentDB
	preservedPtrToMidi_instrument := &midi_instrumentDB_Data
	backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[midi_instrumentDB.ID] = preservedPtrToMidi_instrument

	return
}

// BackRepoMidi_instrument.CheckoutPhaseTwo Checkouts all staged instances of Midi_instrument to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, midi_instrumentDB := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB {
		backRepoMidi_instrument.CheckoutPhaseTwoInstance(backRepo, midi_instrumentDB)
	}
	return
}

// BackRepoMidi_instrument.CheckoutPhaseTwoInstance Checkouts staged instances of Midi_instrument to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, midi_instrumentDB *Midi_instrumentDB) (Error error) {

	midi_instrument := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr[midi_instrumentDB.ID]

	midi_instrumentDB.DecodePointers(backRepo, midi_instrument)

	return
}

func (midi_instrumentDB *Midi_instrumentDB) DecodePointers(backRepo *BackRepoStruct, midi_instrument *models.Midi_instrument) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitMidi_instrument allows commit of a single midi_instrument (if already staged)
func (backRepo *BackRepoStruct) CommitMidi_instrument(midi_instrument *models.Midi_instrument) {
	backRepo.BackRepoMidi_instrument.CommitPhaseOneInstance(midi_instrument)
	if id, ok := backRepo.BackRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument]; ok {
		backRepo.BackRepoMidi_instrument.CommitPhaseTwoInstance(backRepo, id, midi_instrument)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitMidi_instrument allows checkout of a single midi_instrument (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMidi_instrument(midi_instrument *models.Midi_instrument) {
	// check if the midi_instrument is staged
	if _, ok := backRepo.BackRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument]; ok {

		if id, ok := backRepo.BackRepoMidi_instrument.Map_Midi_instrumentPtr_Midi_instrumentDBID[midi_instrument]; ok {
			var midi_instrumentDB Midi_instrumentDB
			midi_instrumentDB.ID = id

			if err := backRepo.BackRepoMidi_instrument.db.First(&midi_instrumentDB, id).Error; err != nil {
				log.Fatalln("CheckoutMidi_instrument : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMidi_instrument.CheckoutPhaseOneInstance(&midi_instrumentDB)
			backRepo.BackRepoMidi_instrument.CheckoutPhaseTwoInstance(backRepo, &midi_instrumentDB)
		}
	}
}

// CopyBasicFieldsFromMidi_instrument
func (midi_instrumentDB *Midi_instrumentDB) CopyBasicFieldsFromMidi_instrument(midi_instrument *models.Midi_instrument) {
	// insertion point for fields commit

	midi_instrumentDB.Name_Data.String = midi_instrument.Name
	midi_instrumentDB.Name_Data.Valid = true

	midi_instrumentDB.Midi_name_Data.String = midi_instrument.Midi_name
	midi_instrumentDB.Midi_name_Data.Valid = true
}

// CopyBasicFieldsFromMidi_instrument_WOP
func (midi_instrumentDB *Midi_instrumentDB) CopyBasicFieldsFromMidi_instrument_WOP(midi_instrument *models.Midi_instrument_WOP) {
	// insertion point for fields commit

	midi_instrumentDB.Name_Data.String = midi_instrument.Name
	midi_instrumentDB.Name_Data.Valid = true

	midi_instrumentDB.Midi_name_Data.String = midi_instrument.Midi_name
	midi_instrumentDB.Midi_name_Data.Valid = true
}

// CopyBasicFieldsFromMidi_instrumentWOP
func (midi_instrumentDB *Midi_instrumentDB) CopyBasicFieldsFromMidi_instrumentWOP(midi_instrument *Midi_instrumentWOP) {
	// insertion point for fields commit

	midi_instrumentDB.Name_Data.String = midi_instrument.Name
	midi_instrumentDB.Name_Data.Valid = true

	midi_instrumentDB.Midi_name_Data.String = midi_instrument.Midi_name
	midi_instrumentDB.Midi_name_Data.Valid = true
}

// CopyBasicFieldsToMidi_instrument
func (midi_instrumentDB *Midi_instrumentDB) CopyBasicFieldsToMidi_instrument(midi_instrument *models.Midi_instrument) {
	// insertion point for checkout of basic fields (back repo to stage)
	midi_instrument.Name = midi_instrumentDB.Name_Data.String
	midi_instrument.Midi_name = midi_instrumentDB.Midi_name_Data.String
}

// CopyBasicFieldsToMidi_instrument_WOP
func (midi_instrumentDB *Midi_instrumentDB) CopyBasicFieldsToMidi_instrument_WOP(midi_instrument *models.Midi_instrument_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	midi_instrument.Name = midi_instrumentDB.Name_Data.String
	midi_instrument.Midi_name = midi_instrumentDB.Midi_name_Data.String
}

// CopyBasicFieldsToMidi_instrumentWOP
func (midi_instrumentDB *Midi_instrumentDB) CopyBasicFieldsToMidi_instrumentWOP(midi_instrument *Midi_instrumentWOP) {
	midi_instrument.ID = int(midi_instrumentDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	midi_instrument.Name = midi_instrumentDB.Name_Data.String
	midi_instrument.Midi_name = midi_instrumentDB.Midi_name_Data.String
}

// Backup generates a json file from a slice of all Midi_instrumentDB instances in the backrepo
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Midi_instrumentDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Midi_instrumentDB, 0)
	for _, midi_instrumentDB := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB {
		forBackup = append(forBackup, midi_instrumentDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Midi_instrument ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Midi_instrument file", err.Error())
	}
}

// Backup generates a json file from a slice of all Midi_instrumentDB instances in the backrepo
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Midi_instrumentDB, 0)
	for _, midi_instrumentDB := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB {
		forBackup = append(forBackup, midi_instrumentDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Midi_instrument")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Midi_instrument_Fields, -1)
	for _, midi_instrumentDB := range forBackup {

		var midi_instrumentWOP Midi_instrumentWOP
		midi_instrumentDB.CopyBasicFieldsToMidi_instrumentWOP(&midi_instrumentWOP)

		row := sh.AddRow()
		row.WriteStruct(&midi_instrumentWOP, -1)
	}
}

// RestoreXL from the "Midi_instrument" sheet all Midi_instrumentDB instances
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoMidi_instrumentid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Midi_instrument"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoMidi_instrument.rowVisitorMidi_instrument)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) rowVisitorMidi_instrument(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var midi_instrumentWOP Midi_instrumentWOP
		row.ReadStruct(&midi_instrumentWOP)

		// add the unmarshalled struct to the stage
		midi_instrumentDB := new(Midi_instrumentDB)
		midi_instrumentDB.CopyBasicFieldsFromMidi_instrumentWOP(&midi_instrumentWOP)

		midi_instrumentDB_ID_atBackupTime := midi_instrumentDB.ID
		midi_instrumentDB.ID = 0
		query := backRepoMidi_instrument.db.Create(midi_instrumentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[midi_instrumentDB.ID] = midi_instrumentDB
		BackRepoMidi_instrumentid_atBckpTime_newID[midi_instrumentDB_ID_atBackupTime] = midi_instrumentDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Midi_instrumentDB.json" in dirPath that stores an array
// of Midi_instrumentDB and stores it in the database
// the map BackRepoMidi_instrumentid_atBckpTime_newID is updated accordingly
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMidi_instrumentid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Midi_instrumentDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Midi_instrument file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Midi_instrumentDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Midi_instrumentDBID_Midi_instrumentDB
	for _, midi_instrumentDB := range forRestore {

		midi_instrumentDB_ID_atBackupTime := midi_instrumentDB.ID
		midi_instrumentDB.ID = 0
		query := backRepoMidi_instrument.db.Create(midi_instrumentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[midi_instrumentDB.ID] = midi_instrumentDB
		BackRepoMidi_instrumentid_atBckpTime_newID[midi_instrumentDB_ID_atBackupTime] = midi_instrumentDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Midi_instrument file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Midi_instrument>id_atBckpTime_newID
// to compute new index
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) RestorePhaseTwo() {

	for _, midi_instrumentDB := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB {

		// next line of code is to avert unused variable compilation error
		_ = midi_instrumentDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoMidi_instrument.db.Model(midi_instrumentDB).Updates(*midi_instrumentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoMidi_instrument.ResetReversePointers commits all staged instances of Midi_instrument to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, midi_instrument := range backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentPtr {
		backRepoMidi_instrument.ResetReversePointersInstance(backRepo, idx, midi_instrument)
	}

	return
}

func (backRepoMidi_instrument *BackRepoMidi_instrumentStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, midi_instrument *models.Midi_instrument) (Error error) {

	// fetch matching midi_instrumentDB
	if midi_instrumentDB, ok := backRepoMidi_instrument.Map_Midi_instrumentDBID_Midi_instrumentDB[idx]; ok {
		_ = midi_instrumentDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMidi_instrumentid_atBckpTime_newID map[uint]uint
