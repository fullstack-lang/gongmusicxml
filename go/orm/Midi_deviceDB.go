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
var dummy_Midi_device_sql sql.NullBool
var dummy_Midi_device_time time.Duration
var dummy_Midi_device_sort sort.Float64Slice

// Midi_deviceAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model midi_deviceAPI
type Midi_deviceAPI struct {
	gorm.Model

	models.Midi_device_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Midi_devicePointersEncoding Midi_devicePointersEncoding
}

// Midi_devicePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Midi_devicePointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// Midi_deviceDB describes a midi_device in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model midi_deviceDB
type Midi_deviceDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field midi_deviceDB.Name
	Name_Data sql.NullString

	// Declation for basic field midi_deviceDB.Value
	Value_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Midi_devicePointersEncoding
}

// Midi_deviceDBs arrays midi_deviceDBs
// swagger:response midi_deviceDBsResponse
type Midi_deviceDBs []Midi_deviceDB

// Midi_deviceDBResponse provides response
// swagger:response midi_deviceDBResponse
type Midi_deviceDBResponse struct {
	Midi_deviceDB
}

// Midi_deviceWOP is a Midi_device without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Midi_deviceWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Midi_device_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoMidi_deviceStruct struct {
	// stores Midi_deviceDB according to their gorm ID
	Map_Midi_deviceDBID_Midi_deviceDB map[uint]*Midi_deviceDB

	// stores Midi_deviceDB ID according to Midi_device address
	Map_Midi_devicePtr_Midi_deviceDBID map[*models.Midi_device]uint

	// stores Midi_device according to their gorm ID
	Map_Midi_deviceDBID_Midi_devicePtr map[uint]*models.Midi_device

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoMidi_device *BackRepoMidi_deviceStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoMidi_device.stage
	return
}

func (backRepoMidi_device *BackRepoMidi_deviceStruct) GetDB() *gorm.DB {
	return backRepoMidi_device.db
}

// GetMidi_deviceDBFromMidi_devicePtr is a handy function to access the back repo instance from the stage instance
func (backRepoMidi_device *BackRepoMidi_deviceStruct) GetMidi_deviceDBFromMidi_devicePtr(midi_device *models.Midi_device) (midi_deviceDB *Midi_deviceDB) {
	id := backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device]
	midi_deviceDB = backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[id]
	return
}

// BackRepoMidi_device.CommitPhaseOne commits all staged instances of Midi_device to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for midi_device := range stage.Midi_devices {
		backRepoMidi_device.CommitPhaseOneInstance(midi_device)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, midi_device := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr {
		if _, ok := stage.Midi_devices[midi_device]; !ok {
			backRepoMidi_device.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMidi_device.CommitDeleteInstance commits deletion of Midi_device to the BackRepo
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CommitDeleteInstance(id uint) (Error error) {

	midi_device := backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[id]

	// midi_device is not staged anymore, remove midi_deviceDB
	midi_deviceDB := backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[id]
	query := backRepoMidi_device.db.Unscoped().Delete(&midi_deviceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID, midi_device)
	delete(backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr, id)
	delete(backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB, id)

	return
}

// BackRepoMidi_device.CommitPhaseOneInstance commits midi_device staged instances of Midi_device to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CommitPhaseOneInstance(midi_device *models.Midi_device) (Error error) {

	// check if the midi_device is not commited yet
	if _, ok := backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device]; ok {
		return
	}

	// initiate midi_device
	var midi_deviceDB Midi_deviceDB
	midi_deviceDB.CopyBasicFieldsFromMidi_device(midi_device)

	query := backRepoMidi_device.db.Create(&midi_deviceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device] = midi_deviceDB.ID
	backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID] = midi_device
	backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[midi_deviceDB.ID] = &midi_deviceDB

	return
}

// BackRepoMidi_device.CommitPhaseTwo commits all staged instances of Midi_device to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, midi_device := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr {
		backRepoMidi_device.CommitPhaseTwoInstance(backRepo, idx, midi_device)
	}

	return
}

// BackRepoMidi_device.CommitPhaseTwoInstance commits {{structname }} of models.Midi_device to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, midi_device *models.Midi_device) (Error error) {

	// fetch matching midi_deviceDB
	if midi_deviceDB, ok := backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[idx]; ok {

		midi_deviceDB.CopyBasicFieldsFromMidi_device(midi_device)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoMidi_device.db.Save(&midi_deviceDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Midi_device intance %s", midi_device.Name))
		return err
	}

	return
}

// BackRepoMidi_device.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CheckoutPhaseOne() (Error error) {

	midi_deviceDBArray := make([]Midi_deviceDB, 0)
	query := backRepoMidi_device.db.Find(&midi_deviceDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	midi_deviceInstancesToBeRemovedFromTheStage := make(map[*models.Midi_device]any)
	for key, value := range backRepoMidi_device.stage.Midi_devices {
		midi_deviceInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, midi_deviceDB := range midi_deviceDBArray {
		backRepoMidi_device.CheckoutPhaseOneInstance(&midi_deviceDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		midi_device, ok := backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID]
		if ok {
			delete(midi_deviceInstancesToBeRemovedFromTheStage, midi_device)
		}
	}

	// remove from stage and back repo's 3 maps all midi_devices that are not in the checkout
	for midi_device := range midi_deviceInstancesToBeRemovedFromTheStage {
		midi_device.Unstage(backRepoMidi_device.GetStage())

		// remove instance from the back repo 3 maps
		midi_deviceID := backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device]
		delete(backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID, midi_device)
		delete(backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB, midi_deviceID)
		delete(backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr, midi_deviceID)
	}

	return
}

// CheckoutPhaseOneInstance takes a midi_deviceDB that has been found in the DB, updates the backRepo and stages the
// models version of the midi_deviceDB
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CheckoutPhaseOneInstance(midi_deviceDB *Midi_deviceDB) (Error error) {

	midi_device, ok := backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID]
	if !ok {
		midi_device = new(models.Midi_device)

		backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID] = midi_device
		backRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device] = midi_deviceDB.ID

		// append model store with the new element
		midi_device.Name = midi_deviceDB.Name_Data.String
		midi_device.Stage(backRepoMidi_device.GetStage())
	}
	midi_deviceDB.CopyBasicFieldsToMidi_device(midi_device)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	midi_device.Stage(backRepoMidi_device.GetStage())

	// preserve pointer to midi_deviceDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Midi_deviceDBID_Midi_deviceDB)[midi_deviceDB hold variable pointers
	midi_deviceDB_Data := *midi_deviceDB
	preservedPtrToMidi_device := &midi_deviceDB_Data
	backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[midi_deviceDB.ID] = preservedPtrToMidi_device

	return
}

// BackRepoMidi_device.CheckoutPhaseTwo Checkouts all staged instances of Midi_device to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, midi_deviceDB := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB {
		backRepoMidi_device.CheckoutPhaseTwoInstance(backRepo, midi_deviceDB)
	}
	return
}

// BackRepoMidi_device.CheckoutPhaseTwoInstance Checkouts staged instances of Midi_device to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_device *BackRepoMidi_deviceStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, midi_deviceDB *Midi_deviceDB) (Error error) {

	midi_device := backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr[midi_deviceDB.ID]

	midi_deviceDB.DecodePointers(backRepo, midi_device)

	return
}

func (midi_deviceDB *Midi_deviceDB) DecodePointers(backRepo *BackRepoStruct, midi_device *models.Midi_device) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitMidi_device allows commit of a single midi_device (if already staged)
func (backRepo *BackRepoStruct) CommitMidi_device(midi_device *models.Midi_device) {
	backRepo.BackRepoMidi_device.CommitPhaseOneInstance(midi_device)
	if id, ok := backRepo.BackRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device]; ok {
		backRepo.BackRepoMidi_device.CommitPhaseTwoInstance(backRepo, id, midi_device)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitMidi_device allows checkout of a single midi_device (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMidi_device(midi_device *models.Midi_device) {
	// check if the midi_device is staged
	if _, ok := backRepo.BackRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device]; ok {

		if id, ok := backRepo.BackRepoMidi_device.Map_Midi_devicePtr_Midi_deviceDBID[midi_device]; ok {
			var midi_deviceDB Midi_deviceDB
			midi_deviceDB.ID = id

			if err := backRepo.BackRepoMidi_device.db.First(&midi_deviceDB, id).Error; err != nil {
				log.Fatalln("CheckoutMidi_device : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMidi_device.CheckoutPhaseOneInstance(&midi_deviceDB)
			backRepo.BackRepoMidi_device.CheckoutPhaseTwoInstance(backRepo, &midi_deviceDB)
		}
	}
}

// CopyBasicFieldsFromMidi_device
func (midi_deviceDB *Midi_deviceDB) CopyBasicFieldsFromMidi_device(midi_device *models.Midi_device) {
	// insertion point for fields commit

	midi_deviceDB.Name_Data.String = midi_device.Name
	midi_deviceDB.Name_Data.Valid = true

	midi_deviceDB.Value_Data.String = midi_device.Value
	midi_deviceDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromMidi_device_WOP
func (midi_deviceDB *Midi_deviceDB) CopyBasicFieldsFromMidi_device_WOP(midi_device *models.Midi_device_WOP) {
	// insertion point for fields commit

	midi_deviceDB.Name_Data.String = midi_device.Name
	midi_deviceDB.Name_Data.Valid = true

	midi_deviceDB.Value_Data.String = midi_device.Value
	midi_deviceDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromMidi_deviceWOP
func (midi_deviceDB *Midi_deviceDB) CopyBasicFieldsFromMidi_deviceWOP(midi_device *Midi_deviceWOP) {
	// insertion point for fields commit

	midi_deviceDB.Name_Data.String = midi_device.Name
	midi_deviceDB.Name_Data.Valid = true

	midi_deviceDB.Value_Data.String = midi_device.Value
	midi_deviceDB.Value_Data.Valid = true
}

// CopyBasicFieldsToMidi_device
func (midi_deviceDB *Midi_deviceDB) CopyBasicFieldsToMidi_device(midi_device *models.Midi_device) {
	// insertion point for checkout of basic fields (back repo to stage)
	midi_device.Name = midi_deviceDB.Name_Data.String
	midi_device.Value = midi_deviceDB.Value_Data.String
}

// CopyBasicFieldsToMidi_device_WOP
func (midi_deviceDB *Midi_deviceDB) CopyBasicFieldsToMidi_device_WOP(midi_device *models.Midi_device_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	midi_device.Name = midi_deviceDB.Name_Data.String
	midi_device.Value = midi_deviceDB.Value_Data.String
}

// CopyBasicFieldsToMidi_deviceWOP
func (midi_deviceDB *Midi_deviceDB) CopyBasicFieldsToMidi_deviceWOP(midi_device *Midi_deviceWOP) {
	midi_device.ID = int(midi_deviceDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	midi_device.Name = midi_deviceDB.Name_Data.String
	midi_device.Value = midi_deviceDB.Value_Data.String
}

// Backup generates a json file from a slice of all Midi_deviceDB instances in the backrepo
func (backRepoMidi_device *BackRepoMidi_deviceStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Midi_deviceDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Midi_deviceDB, 0)
	for _, midi_deviceDB := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB {
		forBackup = append(forBackup, midi_deviceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Midi_device ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Midi_device file", err.Error())
	}
}

// Backup generates a json file from a slice of all Midi_deviceDB instances in the backrepo
func (backRepoMidi_device *BackRepoMidi_deviceStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Midi_deviceDB, 0)
	for _, midi_deviceDB := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB {
		forBackup = append(forBackup, midi_deviceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Midi_device")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Midi_device_Fields, -1)
	for _, midi_deviceDB := range forBackup {

		var midi_deviceWOP Midi_deviceWOP
		midi_deviceDB.CopyBasicFieldsToMidi_deviceWOP(&midi_deviceWOP)

		row := sh.AddRow()
		row.WriteStruct(&midi_deviceWOP, -1)
	}
}

// RestoreXL from the "Midi_device" sheet all Midi_deviceDB instances
func (backRepoMidi_device *BackRepoMidi_deviceStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoMidi_deviceid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Midi_device"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoMidi_device.rowVisitorMidi_device)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoMidi_device *BackRepoMidi_deviceStruct) rowVisitorMidi_device(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var midi_deviceWOP Midi_deviceWOP
		row.ReadStruct(&midi_deviceWOP)

		// add the unmarshalled struct to the stage
		midi_deviceDB := new(Midi_deviceDB)
		midi_deviceDB.CopyBasicFieldsFromMidi_deviceWOP(&midi_deviceWOP)

		midi_deviceDB_ID_atBackupTime := midi_deviceDB.ID
		midi_deviceDB.ID = 0
		query := backRepoMidi_device.db.Create(midi_deviceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[midi_deviceDB.ID] = midi_deviceDB
		BackRepoMidi_deviceid_atBckpTime_newID[midi_deviceDB_ID_atBackupTime] = midi_deviceDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Midi_deviceDB.json" in dirPath that stores an array
// of Midi_deviceDB and stores it in the database
// the map BackRepoMidi_deviceid_atBckpTime_newID is updated accordingly
func (backRepoMidi_device *BackRepoMidi_deviceStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMidi_deviceid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Midi_deviceDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Midi_device file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Midi_deviceDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Midi_deviceDBID_Midi_deviceDB
	for _, midi_deviceDB := range forRestore {

		midi_deviceDB_ID_atBackupTime := midi_deviceDB.ID
		midi_deviceDB.ID = 0
		query := backRepoMidi_device.db.Create(midi_deviceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[midi_deviceDB.ID] = midi_deviceDB
		BackRepoMidi_deviceid_atBckpTime_newID[midi_deviceDB_ID_atBackupTime] = midi_deviceDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Midi_device file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Midi_device>id_atBckpTime_newID
// to compute new index
func (backRepoMidi_device *BackRepoMidi_deviceStruct) RestorePhaseTwo() {

	for _, midi_deviceDB := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB {

		// next line of code is to avert unused variable compilation error
		_ = midi_deviceDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoMidi_device.db.Model(midi_deviceDB).Updates(*midi_deviceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoMidi_device.ResetReversePointers commits all staged instances of Midi_device to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMidi_device *BackRepoMidi_deviceStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, midi_device := range backRepoMidi_device.Map_Midi_deviceDBID_Midi_devicePtr {
		backRepoMidi_device.ResetReversePointersInstance(backRepo, idx, midi_device)
	}

	return
}

func (backRepoMidi_device *BackRepoMidi_deviceStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, midi_device *models.Midi_device) (Error error) {

	// fetch matching midi_deviceDB
	if midi_deviceDB, ok := backRepoMidi_device.Map_Midi_deviceDBID_Midi_deviceDB[idx]; ok {
		_ = midi_deviceDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMidi_deviceid_atBckpTime_newID map[uint]uint
