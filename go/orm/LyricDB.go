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
var dummy_Lyric_sql sql.NullBool
var dummy_Lyric_time time.Duration
var dummy_Lyric_sort sort.Float64Slice

// LyricAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model lyricAPI
type LyricAPI struct {
	gorm.Model

	models.Lyric_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	LyricPointersEncoding LyricPointersEncoding
}

// LyricPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LyricPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field End_line is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	End_lineID sql.NullInt64

	// field End_paragraph is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	End_paragraphID sql.NullInt64

	// field Extend is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	ExtendID sql.NullInt64

	// field Laughing is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	LaughingID sql.NullInt64

	// field Humming is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	HummingID sql.NullInt64
}

// LyricDB describes a lyric in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model lyricDB
type LyricDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field lyricDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	LyricPointersEncoding
}

// LyricDBs arrays lyricDBs
// swagger:response lyricDBsResponse
type LyricDBs []LyricDB

// LyricDBResponse provides response
// swagger:response lyricDBResponse
type LyricDBResponse struct {
	LyricDB
}

// LyricWOP is a Lyric without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LyricWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Lyric_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoLyricStruct struct {
	// stores LyricDB according to their gorm ID
	Map_LyricDBID_LyricDB map[uint]*LyricDB

	// stores LyricDB ID according to Lyric address
	Map_LyricPtr_LyricDBID map[*models.Lyric]uint

	// stores Lyric according to their gorm ID
	Map_LyricDBID_LyricPtr map[uint]*models.Lyric

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoLyric *BackRepoLyricStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoLyric.stage
	return
}

func (backRepoLyric *BackRepoLyricStruct) GetDB() *gorm.DB {
	return backRepoLyric.db
}

// GetLyricDBFromLyricPtr is a handy function to access the back repo instance from the stage instance
func (backRepoLyric *BackRepoLyricStruct) GetLyricDBFromLyricPtr(lyric *models.Lyric) (lyricDB *LyricDB) {
	id := backRepoLyric.Map_LyricPtr_LyricDBID[lyric]
	lyricDB = backRepoLyric.Map_LyricDBID_LyricDB[id]
	return
}

// BackRepoLyric.CommitPhaseOne commits all staged instances of Lyric to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLyric *BackRepoLyricStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for lyric := range stage.Lyrics {
		backRepoLyric.CommitPhaseOneInstance(lyric)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, lyric := range backRepoLyric.Map_LyricDBID_LyricPtr {
		if _, ok := stage.Lyrics[lyric]; !ok {
			backRepoLyric.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLyric.CommitDeleteInstance commits deletion of Lyric to the BackRepo
func (backRepoLyric *BackRepoLyricStruct) CommitDeleteInstance(id uint) (Error error) {

	lyric := backRepoLyric.Map_LyricDBID_LyricPtr[id]

	// lyric is not staged anymore, remove lyricDB
	lyricDB := backRepoLyric.Map_LyricDBID_LyricDB[id]
	query := backRepoLyric.db.Unscoped().Delete(&lyricDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoLyric.Map_LyricPtr_LyricDBID, lyric)
	delete(backRepoLyric.Map_LyricDBID_LyricPtr, id)
	delete(backRepoLyric.Map_LyricDBID_LyricDB, id)

	return
}

// BackRepoLyric.CommitPhaseOneInstance commits lyric staged instances of Lyric to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLyric *BackRepoLyricStruct) CommitPhaseOneInstance(lyric *models.Lyric) (Error error) {

	// check if the lyric is not commited yet
	if _, ok := backRepoLyric.Map_LyricPtr_LyricDBID[lyric]; ok {
		return
	}

	// initiate lyric
	var lyricDB LyricDB
	lyricDB.CopyBasicFieldsFromLyric(lyric)

	query := backRepoLyric.db.Create(&lyricDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoLyric.Map_LyricPtr_LyricDBID[lyric] = lyricDB.ID
	backRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID] = lyric
	backRepoLyric.Map_LyricDBID_LyricDB[lyricDB.ID] = &lyricDB

	return
}

// BackRepoLyric.CommitPhaseTwo commits all staged instances of Lyric to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLyric *BackRepoLyricStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, lyric := range backRepoLyric.Map_LyricDBID_LyricPtr {
		backRepoLyric.CommitPhaseTwoInstance(backRepo, idx, lyric)
	}

	return
}

// BackRepoLyric.CommitPhaseTwoInstance commits {{structname }} of models.Lyric to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLyric *BackRepoLyricStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, lyric *models.Lyric) (Error error) {

	// fetch matching lyricDB
	if lyricDB, ok := backRepoLyric.Map_LyricDBID_LyricDB[idx]; ok {

		lyricDB.CopyBasicFieldsFromLyric(lyric)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value lyric.End_line translates to updating the lyric.End_lineID
		lyricDB.End_lineID.Valid = true // allow for a 0 value (nil association)
		if lyric.End_line != nil {
			if End_lineId, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[lyric.End_line]; ok {
				lyricDB.End_lineID.Int64 = int64(End_lineId)
				lyricDB.End_lineID.Valid = true
			}
		} else {
			lyricDB.End_lineID.Int64 = 0
			lyricDB.End_lineID.Valid = true
		}

		// commit pointer value lyric.End_paragraph translates to updating the lyric.End_paragraphID
		lyricDB.End_paragraphID.Valid = true // allow for a 0 value (nil association)
		if lyric.End_paragraph != nil {
			if End_paragraphId, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[lyric.End_paragraph]; ok {
				lyricDB.End_paragraphID.Int64 = int64(End_paragraphId)
				lyricDB.End_paragraphID.Valid = true
			}
		} else {
			lyricDB.End_paragraphID.Int64 = 0
			lyricDB.End_paragraphID.Valid = true
		}

		// commit pointer value lyric.Extend translates to updating the lyric.ExtendID
		lyricDB.ExtendID.Valid = true // allow for a 0 value (nil association)
		if lyric.Extend != nil {
			if ExtendId, ok := backRepo.BackRepoExtend.Map_ExtendPtr_ExtendDBID[lyric.Extend]; ok {
				lyricDB.ExtendID.Int64 = int64(ExtendId)
				lyricDB.ExtendID.Valid = true
			}
		} else {
			lyricDB.ExtendID.Int64 = 0
			lyricDB.ExtendID.Valid = true
		}

		// commit pointer value lyric.Laughing translates to updating the lyric.LaughingID
		lyricDB.LaughingID.Valid = true // allow for a 0 value (nil association)
		if lyric.Laughing != nil {
			if LaughingId, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[lyric.Laughing]; ok {
				lyricDB.LaughingID.Int64 = int64(LaughingId)
				lyricDB.LaughingID.Valid = true
			}
		} else {
			lyricDB.LaughingID.Int64 = 0
			lyricDB.LaughingID.Valid = true
		}

		// commit pointer value lyric.Humming translates to updating the lyric.HummingID
		lyricDB.HummingID.Valid = true // allow for a 0 value (nil association)
		if lyric.Humming != nil {
			if HummingId, ok := backRepo.BackRepoEmpty.Map_EmptyPtr_EmptyDBID[lyric.Humming]; ok {
				lyricDB.HummingID.Int64 = int64(HummingId)
				lyricDB.HummingID.Valid = true
			}
		} else {
			lyricDB.HummingID.Int64 = 0
			lyricDB.HummingID.Valid = true
		}

		query := backRepoLyric.db.Save(&lyricDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Lyric intance %s", lyric.Name))
		return err
	}

	return
}

// BackRepoLyric.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoLyric *BackRepoLyricStruct) CheckoutPhaseOne() (Error error) {

	lyricDBArray := make([]LyricDB, 0)
	query := backRepoLyric.db.Find(&lyricDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	lyricInstancesToBeRemovedFromTheStage := make(map[*models.Lyric]any)
	for key, value := range backRepoLyric.stage.Lyrics {
		lyricInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, lyricDB := range lyricDBArray {
		backRepoLyric.CheckoutPhaseOneInstance(&lyricDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		lyric, ok := backRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID]
		if ok {
			delete(lyricInstancesToBeRemovedFromTheStage, lyric)
		}
	}

	// remove from stage and back repo's 3 maps all lyrics that are not in the checkout
	for lyric := range lyricInstancesToBeRemovedFromTheStage {
		lyric.Unstage(backRepoLyric.GetStage())

		// remove instance from the back repo 3 maps
		lyricID := backRepoLyric.Map_LyricPtr_LyricDBID[lyric]
		delete(backRepoLyric.Map_LyricPtr_LyricDBID, lyric)
		delete(backRepoLyric.Map_LyricDBID_LyricDB, lyricID)
		delete(backRepoLyric.Map_LyricDBID_LyricPtr, lyricID)
	}

	return
}

// CheckoutPhaseOneInstance takes a lyricDB that has been found in the DB, updates the backRepo and stages the
// models version of the lyricDB
func (backRepoLyric *BackRepoLyricStruct) CheckoutPhaseOneInstance(lyricDB *LyricDB) (Error error) {

	lyric, ok := backRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID]
	if !ok {
		lyric = new(models.Lyric)

		backRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID] = lyric
		backRepoLyric.Map_LyricPtr_LyricDBID[lyric] = lyricDB.ID

		// append model store with the new element
		lyric.Name = lyricDB.Name_Data.String
		lyric.Stage(backRepoLyric.GetStage())
	}
	lyricDB.CopyBasicFieldsToLyric(lyric)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	lyric.Stage(backRepoLyric.GetStage())

	// preserve pointer to lyricDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LyricDBID_LyricDB)[lyricDB hold variable pointers
	lyricDB_Data := *lyricDB
	preservedPtrToLyric := &lyricDB_Data
	backRepoLyric.Map_LyricDBID_LyricDB[lyricDB.ID] = preservedPtrToLyric

	return
}

// BackRepoLyric.CheckoutPhaseTwo Checkouts all staged instances of Lyric to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLyric *BackRepoLyricStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, lyricDB := range backRepoLyric.Map_LyricDBID_LyricDB {
		backRepoLyric.CheckoutPhaseTwoInstance(backRepo, lyricDB)
	}
	return
}

// BackRepoLyric.CheckoutPhaseTwoInstance Checkouts staged instances of Lyric to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLyric *BackRepoLyricStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, lyricDB *LyricDB) (Error error) {

	lyric := backRepoLyric.Map_LyricDBID_LyricPtr[lyricDB.ID]

	lyricDB.DecodePointers(backRepo, lyric)

	return
}

func (lyricDB *LyricDB) DecodePointers(backRepo *BackRepoStruct, lyric *models.Lyric) {

	// insertion point for checkout of pointer encoding
	// End_line field
	lyric.End_line = nil
	if lyricDB.End_lineID.Int64 != 0 {
		lyric.End_line = backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[uint(lyricDB.End_lineID.Int64)]
	}
	// End_paragraph field
	lyric.End_paragraph = nil
	if lyricDB.End_paragraphID.Int64 != 0 {
		lyric.End_paragraph = backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[uint(lyricDB.End_paragraphID.Int64)]
	}
	// Extend field
	lyric.Extend = nil
	if lyricDB.ExtendID.Int64 != 0 {
		lyric.Extend = backRepo.BackRepoExtend.Map_ExtendDBID_ExtendPtr[uint(lyricDB.ExtendID.Int64)]
	}
	// Laughing field
	lyric.Laughing = nil
	if lyricDB.LaughingID.Int64 != 0 {
		lyric.Laughing = backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[uint(lyricDB.LaughingID.Int64)]
	}
	// Humming field
	lyric.Humming = nil
	if lyricDB.HummingID.Int64 != 0 {
		lyric.Humming = backRepo.BackRepoEmpty.Map_EmptyDBID_EmptyPtr[uint(lyricDB.HummingID.Int64)]
	}
	return
}

// CommitLyric allows commit of a single lyric (if already staged)
func (backRepo *BackRepoStruct) CommitLyric(lyric *models.Lyric) {
	backRepo.BackRepoLyric.CommitPhaseOneInstance(lyric)
	if id, ok := backRepo.BackRepoLyric.Map_LyricPtr_LyricDBID[lyric]; ok {
		backRepo.BackRepoLyric.CommitPhaseTwoInstance(backRepo, id, lyric)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitLyric allows checkout of a single lyric (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLyric(lyric *models.Lyric) {
	// check if the lyric is staged
	if _, ok := backRepo.BackRepoLyric.Map_LyricPtr_LyricDBID[lyric]; ok {

		if id, ok := backRepo.BackRepoLyric.Map_LyricPtr_LyricDBID[lyric]; ok {
			var lyricDB LyricDB
			lyricDB.ID = id

			if err := backRepo.BackRepoLyric.db.First(&lyricDB, id).Error; err != nil {
				log.Fatalln("CheckoutLyric : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLyric.CheckoutPhaseOneInstance(&lyricDB)
			backRepo.BackRepoLyric.CheckoutPhaseTwoInstance(backRepo, &lyricDB)
		}
	}
}

// CopyBasicFieldsFromLyric
func (lyricDB *LyricDB) CopyBasicFieldsFromLyric(lyric *models.Lyric) {
	// insertion point for fields commit

	lyricDB.Name_Data.String = lyric.Name
	lyricDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromLyric_WOP
func (lyricDB *LyricDB) CopyBasicFieldsFromLyric_WOP(lyric *models.Lyric_WOP) {
	// insertion point for fields commit

	lyricDB.Name_Data.String = lyric.Name
	lyricDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromLyricWOP
func (lyricDB *LyricDB) CopyBasicFieldsFromLyricWOP(lyric *LyricWOP) {
	// insertion point for fields commit

	lyricDB.Name_Data.String = lyric.Name
	lyricDB.Name_Data.Valid = true
}

// CopyBasicFieldsToLyric
func (lyricDB *LyricDB) CopyBasicFieldsToLyric(lyric *models.Lyric) {
	// insertion point for checkout of basic fields (back repo to stage)
	lyric.Name = lyricDB.Name_Data.String
}

// CopyBasicFieldsToLyric_WOP
func (lyricDB *LyricDB) CopyBasicFieldsToLyric_WOP(lyric *models.Lyric_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	lyric.Name = lyricDB.Name_Data.String
}

// CopyBasicFieldsToLyricWOP
func (lyricDB *LyricDB) CopyBasicFieldsToLyricWOP(lyric *LyricWOP) {
	lyric.ID = int(lyricDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	lyric.Name = lyricDB.Name_Data.String
}

// Backup generates a json file from a slice of all LyricDB instances in the backrepo
func (backRepoLyric *BackRepoLyricStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LyricDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LyricDB, 0)
	for _, lyricDB := range backRepoLyric.Map_LyricDBID_LyricDB {
		forBackup = append(forBackup, lyricDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Lyric ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Lyric file", err.Error())
	}
}

// Backup generates a json file from a slice of all LyricDB instances in the backrepo
func (backRepoLyric *BackRepoLyricStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LyricDB, 0)
	for _, lyricDB := range backRepoLyric.Map_LyricDBID_LyricDB {
		forBackup = append(forBackup, lyricDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Lyric")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Lyric_Fields, -1)
	for _, lyricDB := range forBackup {

		var lyricWOP LyricWOP
		lyricDB.CopyBasicFieldsToLyricWOP(&lyricWOP)

		row := sh.AddRow()
		row.WriteStruct(&lyricWOP, -1)
	}
}

// RestoreXL from the "Lyric" sheet all LyricDB instances
func (backRepoLyric *BackRepoLyricStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLyricid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Lyric"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLyric.rowVisitorLyric)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoLyric *BackRepoLyricStruct) rowVisitorLyric(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var lyricWOP LyricWOP
		row.ReadStruct(&lyricWOP)

		// add the unmarshalled struct to the stage
		lyricDB := new(LyricDB)
		lyricDB.CopyBasicFieldsFromLyricWOP(&lyricWOP)

		lyricDB_ID_atBackupTime := lyricDB.ID
		lyricDB.ID = 0
		query := backRepoLyric.db.Create(lyricDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoLyric.Map_LyricDBID_LyricDB[lyricDB.ID] = lyricDB
		BackRepoLyricid_atBckpTime_newID[lyricDB_ID_atBackupTime] = lyricDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LyricDB.json" in dirPath that stores an array
// of LyricDB and stores it in the database
// the map BackRepoLyricid_atBckpTime_newID is updated accordingly
func (backRepoLyric *BackRepoLyricStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLyricid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LyricDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Lyric file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LyricDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LyricDBID_LyricDB
	for _, lyricDB := range forRestore {

		lyricDB_ID_atBackupTime := lyricDB.ID
		lyricDB.ID = 0
		query := backRepoLyric.db.Create(lyricDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoLyric.Map_LyricDBID_LyricDB[lyricDB.ID] = lyricDB
		BackRepoLyricid_atBckpTime_newID[lyricDB_ID_atBackupTime] = lyricDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Lyric file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Lyric>id_atBckpTime_newID
// to compute new index
func (backRepoLyric *BackRepoLyricStruct) RestorePhaseTwo() {

	for _, lyricDB := range backRepoLyric.Map_LyricDBID_LyricDB {

		// next line of code is to avert unused variable compilation error
		_ = lyricDB

		// insertion point for reindexing pointers encoding
		// reindexing End_line field
		if lyricDB.End_lineID.Int64 != 0 {
			lyricDB.End_lineID.Int64 = int64(BackRepoEmptyid_atBckpTime_newID[uint(lyricDB.End_lineID.Int64)])
			lyricDB.End_lineID.Valid = true
		}

		// reindexing End_paragraph field
		if lyricDB.End_paragraphID.Int64 != 0 {
			lyricDB.End_paragraphID.Int64 = int64(BackRepoEmptyid_atBckpTime_newID[uint(lyricDB.End_paragraphID.Int64)])
			lyricDB.End_paragraphID.Valid = true
		}

		// reindexing Extend field
		if lyricDB.ExtendID.Int64 != 0 {
			lyricDB.ExtendID.Int64 = int64(BackRepoExtendid_atBckpTime_newID[uint(lyricDB.ExtendID.Int64)])
			lyricDB.ExtendID.Valid = true
		}

		// reindexing Laughing field
		if lyricDB.LaughingID.Int64 != 0 {
			lyricDB.LaughingID.Int64 = int64(BackRepoEmptyid_atBckpTime_newID[uint(lyricDB.LaughingID.Int64)])
			lyricDB.LaughingID.Valid = true
		}

		// reindexing Humming field
		if lyricDB.HummingID.Int64 != 0 {
			lyricDB.HummingID.Int64 = int64(BackRepoEmptyid_atBckpTime_newID[uint(lyricDB.HummingID.Int64)])
			lyricDB.HummingID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoLyric.db.Model(lyricDB).Updates(*lyricDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoLyric.ResetReversePointers commits all staged instances of Lyric to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLyric *BackRepoLyricStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, lyric := range backRepoLyric.Map_LyricDBID_LyricPtr {
		backRepoLyric.ResetReversePointersInstance(backRepo, idx, lyric)
	}

	return
}

func (backRepoLyric *BackRepoLyricStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, lyric *models.Lyric) (Error error) {

	// fetch matching lyricDB
	if lyricDB, ok := backRepoLyric.Map_LyricDBID_LyricDB[idx]; ok {
		_ = lyricDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLyricid_atBckpTime_newID map[uint]uint
