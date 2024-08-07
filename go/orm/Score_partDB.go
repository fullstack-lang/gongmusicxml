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
var dummy_Score_part_sql sql.NullBool
var dummy_Score_part_time time.Duration
var dummy_Score_part_sort sort.Float64Slice

// Score_partAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model score_partAPI
type Score_partAPI struct {
	gorm.Model

	models.Score_part_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	Score_partPointersEncoding Score_partPointersEncoding
}

// Score_partPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type Score_partPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Identification is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	IdentificationID sql.NullInt64

	// field Part_link is a slice of pointers to another Struct (optional or 0..1)
	Part_link IntSlice `gorm:"type:TEXT"`

	// field Part_name_display is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	Part_name_displayID sql.NullInt64

	// field Part_abbreviation_display is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	Part_abbreviation_displayID sql.NullInt64

	// field Score_instrument is a slice of pointers to another Struct (optional or 0..1)
	Score_instrument IntSlice `gorm:"type:TEXT"`

	// field Player is a slice of pointers to another Struct (optional or 0..1)
	Player IntSlice `gorm:"type:TEXT"`
}

// Score_partDB describes a score_part in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model score_partDB
type Score_partDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field score_partDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	Score_partPointersEncoding
}

// Score_partDBs arrays score_partDBs
// swagger:response score_partDBsResponse
type Score_partDBs []Score_partDB

// Score_partDBResponse provides response
// swagger:response score_partDBResponse
type Score_partDBResponse struct {
	Score_partDB
}

// Score_partWOP is a Score_part without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type Score_partWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Score_part_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoScore_partStruct struct {
	// stores Score_partDB according to their gorm ID
	Map_Score_partDBID_Score_partDB map[uint]*Score_partDB

	// stores Score_partDB ID according to Score_part address
	Map_Score_partPtr_Score_partDBID map[*models.Score_part]uint

	// stores Score_part according to their gorm ID
	Map_Score_partDBID_Score_partPtr map[uint]*models.Score_part

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoScore_part *BackRepoScore_partStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoScore_part.stage
	return
}

func (backRepoScore_part *BackRepoScore_partStruct) GetDB() *gorm.DB {
	return backRepoScore_part.db
}

// GetScore_partDBFromScore_partPtr is a handy function to access the back repo instance from the stage instance
func (backRepoScore_part *BackRepoScore_partStruct) GetScore_partDBFromScore_partPtr(score_part *models.Score_part) (score_partDB *Score_partDB) {
	id := backRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part]
	score_partDB = backRepoScore_part.Map_Score_partDBID_Score_partDB[id]
	return
}

// BackRepoScore_part.CommitPhaseOne commits all staged instances of Score_part to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoScore_part *BackRepoScore_partStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for score_part := range stage.Score_parts {
		backRepoScore_part.CommitPhaseOneInstance(score_part)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, score_part := range backRepoScore_part.Map_Score_partDBID_Score_partPtr {
		if _, ok := stage.Score_parts[score_part]; !ok {
			backRepoScore_part.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoScore_part.CommitDeleteInstance commits deletion of Score_part to the BackRepo
func (backRepoScore_part *BackRepoScore_partStruct) CommitDeleteInstance(id uint) (Error error) {

	score_part := backRepoScore_part.Map_Score_partDBID_Score_partPtr[id]

	// score_part is not staged anymore, remove score_partDB
	score_partDB := backRepoScore_part.Map_Score_partDBID_Score_partDB[id]
	query := backRepoScore_part.db.Unscoped().Delete(&score_partDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoScore_part.Map_Score_partPtr_Score_partDBID, score_part)
	delete(backRepoScore_part.Map_Score_partDBID_Score_partPtr, id)
	delete(backRepoScore_part.Map_Score_partDBID_Score_partDB, id)

	return
}

// BackRepoScore_part.CommitPhaseOneInstance commits score_part staged instances of Score_part to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoScore_part *BackRepoScore_partStruct) CommitPhaseOneInstance(score_part *models.Score_part) (Error error) {

	// check if the score_part is not commited yet
	if _, ok := backRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part]; ok {
		return
	}

	// initiate score_part
	var score_partDB Score_partDB
	score_partDB.CopyBasicFieldsFromScore_part(score_part)

	query := backRepoScore_part.db.Create(&score_partDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part] = score_partDB.ID
	backRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID] = score_part
	backRepoScore_part.Map_Score_partDBID_Score_partDB[score_partDB.ID] = &score_partDB

	return
}

// BackRepoScore_part.CommitPhaseTwo commits all staged instances of Score_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoScore_part *BackRepoScore_partStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, score_part := range backRepoScore_part.Map_Score_partDBID_Score_partPtr {
		backRepoScore_part.CommitPhaseTwoInstance(backRepo, idx, score_part)
	}

	return
}

// BackRepoScore_part.CommitPhaseTwoInstance commits {{structname }} of models.Score_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoScore_part *BackRepoScore_partStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, score_part *models.Score_part) (Error error) {

	// fetch matching score_partDB
	if score_partDB, ok := backRepoScore_part.Map_Score_partDBID_Score_partDB[idx]; ok {

		score_partDB.CopyBasicFieldsFromScore_part(score_part)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value score_part.Identification translates to updating the score_part.IdentificationID
		score_partDB.IdentificationID.Valid = true // allow for a 0 value (nil association)
		if score_part.Identification != nil {
			if IdentificationId, ok := backRepo.BackRepoIdentification.Map_IdentificationPtr_IdentificationDBID[score_part.Identification]; ok {
				score_partDB.IdentificationID.Int64 = int64(IdentificationId)
				score_partDB.IdentificationID.Valid = true
			}
		} else {
			score_partDB.IdentificationID.Int64 = 0
			score_partDB.IdentificationID.Valid = true
		}

		// 1. reset
		score_partDB.Score_partPointersEncoding.Part_link = make([]int, 0)
		// 2. encode
		for _, part_linkAssocEnd := range score_part.Part_link {
			part_linkAssocEnd_DB :=
				backRepo.BackRepoPart_link.GetPart_linkDBFromPart_linkPtr(part_linkAssocEnd)
			
			// the stage might be inconsistant, meaning that the part_linkAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if part_linkAssocEnd_DB == nil {
				continue
			}
			
			score_partDB.Score_partPointersEncoding.Part_link =
				append(score_partDB.Score_partPointersEncoding.Part_link, int(part_linkAssocEnd_DB.ID))
		}

		// commit pointer value score_part.Part_name_display translates to updating the score_part.Part_name_displayID
		score_partDB.Part_name_displayID.Valid = true // allow for a 0 value (nil association)
		if score_part.Part_name_display != nil {
			if Part_name_displayId, ok := backRepo.BackRepoName_display.Map_Name_displayPtr_Name_displayDBID[score_part.Part_name_display]; ok {
				score_partDB.Part_name_displayID.Int64 = int64(Part_name_displayId)
				score_partDB.Part_name_displayID.Valid = true
			}
		} else {
			score_partDB.Part_name_displayID.Int64 = 0
			score_partDB.Part_name_displayID.Valid = true
		}

		// commit pointer value score_part.Part_abbreviation_display translates to updating the score_part.Part_abbreviation_displayID
		score_partDB.Part_abbreviation_displayID.Valid = true // allow for a 0 value (nil association)
		if score_part.Part_abbreviation_display != nil {
			if Part_abbreviation_displayId, ok := backRepo.BackRepoName_display.Map_Name_displayPtr_Name_displayDBID[score_part.Part_abbreviation_display]; ok {
				score_partDB.Part_abbreviation_displayID.Int64 = int64(Part_abbreviation_displayId)
				score_partDB.Part_abbreviation_displayID.Valid = true
			}
		} else {
			score_partDB.Part_abbreviation_displayID.Int64 = 0
			score_partDB.Part_abbreviation_displayID.Valid = true
		}

		// 1. reset
		score_partDB.Score_partPointersEncoding.Score_instrument = make([]int, 0)
		// 2. encode
		for _, score_instrumentAssocEnd := range score_part.Score_instrument {
			score_instrumentAssocEnd_DB :=
				backRepo.BackRepoScore_instrument.GetScore_instrumentDBFromScore_instrumentPtr(score_instrumentAssocEnd)
			
			// the stage might be inconsistant, meaning that the score_instrumentAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if score_instrumentAssocEnd_DB == nil {
				continue
			}
			
			score_partDB.Score_partPointersEncoding.Score_instrument =
				append(score_partDB.Score_partPointersEncoding.Score_instrument, int(score_instrumentAssocEnd_DB.ID))
		}

		// 1. reset
		score_partDB.Score_partPointersEncoding.Player = make([]int, 0)
		// 2. encode
		for _, playerAssocEnd := range score_part.Player {
			playerAssocEnd_DB :=
				backRepo.BackRepoPlayer.GetPlayerDBFromPlayerPtr(playerAssocEnd)
			
			// the stage might be inconsistant, meaning that the playerAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if playerAssocEnd_DB == nil {
				continue
			}
			
			score_partDB.Score_partPointersEncoding.Player =
				append(score_partDB.Score_partPointersEncoding.Player, int(playerAssocEnd_DB.ID))
		}

		query := backRepoScore_part.db.Save(&score_partDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Score_part intance %s", score_part.Name))
		return err
	}

	return
}

// BackRepoScore_part.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoScore_part *BackRepoScore_partStruct) CheckoutPhaseOne() (Error error) {

	score_partDBArray := make([]Score_partDB, 0)
	query := backRepoScore_part.db.Find(&score_partDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	score_partInstancesToBeRemovedFromTheStage := make(map[*models.Score_part]any)
	for key, value := range backRepoScore_part.stage.Score_parts {
		score_partInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, score_partDB := range score_partDBArray {
		backRepoScore_part.CheckoutPhaseOneInstance(&score_partDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		score_part, ok := backRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID]
		if ok {
			delete(score_partInstancesToBeRemovedFromTheStage, score_part)
		}
	}

	// remove from stage and back repo's 3 maps all score_parts that are not in the checkout
	for score_part := range score_partInstancesToBeRemovedFromTheStage {
		score_part.Unstage(backRepoScore_part.GetStage())

		// remove instance from the back repo 3 maps
		score_partID := backRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part]
		delete(backRepoScore_part.Map_Score_partPtr_Score_partDBID, score_part)
		delete(backRepoScore_part.Map_Score_partDBID_Score_partDB, score_partID)
		delete(backRepoScore_part.Map_Score_partDBID_Score_partPtr, score_partID)
	}

	return
}

// CheckoutPhaseOneInstance takes a score_partDB that has been found in the DB, updates the backRepo and stages the
// models version of the score_partDB
func (backRepoScore_part *BackRepoScore_partStruct) CheckoutPhaseOneInstance(score_partDB *Score_partDB) (Error error) {

	score_part, ok := backRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID]
	if !ok {
		score_part = new(models.Score_part)

		backRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID] = score_part
		backRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part] = score_partDB.ID

		// append model store with the new element
		score_part.Name = score_partDB.Name_Data.String
		score_part.Stage(backRepoScore_part.GetStage())
	}
	score_partDB.CopyBasicFieldsToScore_part(score_part)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	score_part.Stage(backRepoScore_part.GetStage())

	// preserve pointer to score_partDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_Score_partDBID_Score_partDB)[score_partDB hold variable pointers
	score_partDB_Data := *score_partDB
	preservedPtrToScore_part := &score_partDB_Data
	backRepoScore_part.Map_Score_partDBID_Score_partDB[score_partDB.ID] = preservedPtrToScore_part

	return
}

// BackRepoScore_part.CheckoutPhaseTwo Checkouts all staged instances of Score_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoScore_part *BackRepoScore_partStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, score_partDB := range backRepoScore_part.Map_Score_partDBID_Score_partDB {
		backRepoScore_part.CheckoutPhaseTwoInstance(backRepo, score_partDB)
	}
	return
}

// BackRepoScore_part.CheckoutPhaseTwoInstance Checkouts staged instances of Score_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoScore_part *BackRepoScore_partStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, score_partDB *Score_partDB) (Error error) {

	score_part := backRepoScore_part.Map_Score_partDBID_Score_partPtr[score_partDB.ID]

	score_partDB.DecodePointers(backRepo, score_part)

	return
}

func (score_partDB *Score_partDB) DecodePointers(backRepo *BackRepoStruct, score_part *models.Score_part) {

	// insertion point for checkout of pointer encoding
	// Identification field
	score_part.Identification = nil
	if score_partDB.IdentificationID.Int64 != 0 {
		score_part.Identification = backRepo.BackRepoIdentification.Map_IdentificationDBID_IdentificationPtr[uint(score_partDB.IdentificationID.Int64)]
	}
	// This loop redeem score_part.Part_link in the stage from the encode in the back repo
	// It parses all Part_linkDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	score_part.Part_link = score_part.Part_link[:0]
	for _, _Part_linkid := range score_partDB.Score_partPointersEncoding.Part_link {
		score_part.Part_link = append(score_part.Part_link, backRepo.BackRepoPart_link.Map_Part_linkDBID_Part_linkPtr[uint(_Part_linkid)])
	}

	// Part_name_display field
	score_part.Part_name_display = nil
	if score_partDB.Part_name_displayID.Int64 != 0 {
		score_part.Part_name_display = backRepo.BackRepoName_display.Map_Name_displayDBID_Name_displayPtr[uint(score_partDB.Part_name_displayID.Int64)]
	}
	// Part_abbreviation_display field
	score_part.Part_abbreviation_display = nil
	if score_partDB.Part_abbreviation_displayID.Int64 != 0 {
		score_part.Part_abbreviation_display = backRepo.BackRepoName_display.Map_Name_displayDBID_Name_displayPtr[uint(score_partDB.Part_abbreviation_displayID.Int64)]
	}
	// This loop redeem score_part.Score_instrument in the stage from the encode in the back repo
	// It parses all Score_instrumentDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	score_part.Score_instrument = score_part.Score_instrument[:0]
	for _, _Score_instrumentid := range score_partDB.Score_partPointersEncoding.Score_instrument {
		score_part.Score_instrument = append(score_part.Score_instrument, backRepo.BackRepoScore_instrument.Map_Score_instrumentDBID_Score_instrumentPtr[uint(_Score_instrumentid)])
	}

	// This loop redeem score_part.Player in the stage from the encode in the back repo
	// It parses all PlayerDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	score_part.Player = score_part.Player[:0]
	for _, _Playerid := range score_partDB.Score_partPointersEncoding.Player {
		score_part.Player = append(score_part.Player, backRepo.BackRepoPlayer.Map_PlayerDBID_PlayerPtr[uint(_Playerid)])
	}

	return
}

// CommitScore_part allows commit of a single score_part (if already staged)
func (backRepo *BackRepoStruct) CommitScore_part(score_part *models.Score_part) {
	backRepo.BackRepoScore_part.CommitPhaseOneInstance(score_part)
	if id, ok := backRepo.BackRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part]; ok {
		backRepo.BackRepoScore_part.CommitPhaseTwoInstance(backRepo, id, score_part)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitScore_part allows checkout of a single score_part (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutScore_part(score_part *models.Score_part) {
	// check if the score_part is staged
	if _, ok := backRepo.BackRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part]; ok {

		if id, ok := backRepo.BackRepoScore_part.Map_Score_partPtr_Score_partDBID[score_part]; ok {
			var score_partDB Score_partDB
			score_partDB.ID = id

			if err := backRepo.BackRepoScore_part.db.First(&score_partDB, id).Error; err != nil {
				log.Fatalln("CheckoutScore_part : Problem with getting object with id:", id)
			}
			backRepo.BackRepoScore_part.CheckoutPhaseOneInstance(&score_partDB)
			backRepo.BackRepoScore_part.CheckoutPhaseTwoInstance(backRepo, &score_partDB)
		}
	}
}

// CopyBasicFieldsFromScore_part
func (score_partDB *Score_partDB) CopyBasicFieldsFromScore_part(score_part *models.Score_part) {
	// insertion point for fields commit

	score_partDB.Name_Data.String = score_part.Name
	score_partDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromScore_part_WOP
func (score_partDB *Score_partDB) CopyBasicFieldsFromScore_part_WOP(score_part *models.Score_part_WOP) {
	// insertion point for fields commit

	score_partDB.Name_Data.String = score_part.Name
	score_partDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromScore_partWOP
func (score_partDB *Score_partDB) CopyBasicFieldsFromScore_partWOP(score_part *Score_partWOP) {
	// insertion point for fields commit

	score_partDB.Name_Data.String = score_part.Name
	score_partDB.Name_Data.Valid = true
}

// CopyBasicFieldsToScore_part
func (score_partDB *Score_partDB) CopyBasicFieldsToScore_part(score_part *models.Score_part) {
	// insertion point for checkout of basic fields (back repo to stage)
	score_part.Name = score_partDB.Name_Data.String
}

// CopyBasicFieldsToScore_part_WOP
func (score_partDB *Score_partDB) CopyBasicFieldsToScore_part_WOP(score_part *models.Score_part_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	score_part.Name = score_partDB.Name_Data.String
}

// CopyBasicFieldsToScore_partWOP
func (score_partDB *Score_partDB) CopyBasicFieldsToScore_partWOP(score_part *Score_partWOP) {
	score_part.ID = int(score_partDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	score_part.Name = score_partDB.Name_Data.String
}

// Backup generates a json file from a slice of all Score_partDB instances in the backrepo
func (backRepoScore_part *BackRepoScore_partStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "Score_partDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Score_partDB, 0)
	for _, score_partDB := range backRepoScore_part.Map_Score_partDBID_Score_partDB {
		forBackup = append(forBackup, score_partDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Score_part ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Score_part file", err.Error())
	}
}

// Backup generates a json file from a slice of all Score_partDB instances in the backrepo
func (backRepoScore_part *BackRepoScore_partStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*Score_partDB, 0)
	for _, score_partDB := range backRepoScore_part.Map_Score_partDBID_Score_partDB {
		forBackup = append(forBackup, score_partDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Score_part")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Score_part_Fields, -1)
	for _, score_partDB := range forBackup {

		var score_partWOP Score_partWOP
		score_partDB.CopyBasicFieldsToScore_partWOP(&score_partWOP)

		row := sh.AddRow()
		row.WriteStruct(&score_partWOP, -1)
	}
}

// RestoreXL from the "Score_part" sheet all Score_partDB instances
func (backRepoScore_part *BackRepoScore_partStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoScore_partid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Score_part"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoScore_part.rowVisitorScore_part)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoScore_part *BackRepoScore_partStruct) rowVisitorScore_part(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var score_partWOP Score_partWOP
		row.ReadStruct(&score_partWOP)

		// add the unmarshalled struct to the stage
		score_partDB := new(Score_partDB)
		score_partDB.CopyBasicFieldsFromScore_partWOP(&score_partWOP)

		score_partDB_ID_atBackupTime := score_partDB.ID
		score_partDB.ID = 0
		query := backRepoScore_part.db.Create(score_partDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoScore_part.Map_Score_partDBID_Score_partDB[score_partDB.ID] = score_partDB
		BackRepoScore_partid_atBckpTime_newID[score_partDB_ID_atBackupTime] = score_partDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "Score_partDB.json" in dirPath that stores an array
// of Score_partDB and stores it in the database
// the map BackRepoScore_partid_atBckpTime_newID is updated accordingly
func (backRepoScore_part *BackRepoScore_partStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoScore_partid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "Score_partDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Score_part file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*Score_partDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_Score_partDBID_Score_partDB
	for _, score_partDB := range forRestore {

		score_partDB_ID_atBackupTime := score_partDB.ID
		score_partDB.ID = 0
		query := backRepoScore_part.db.Create(score_partDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoScore_part.Map_Score_partDBID_Score_partDB[score_partDB.ID] = score_partDB
		BackRepoScore_partid_atBckpTime_newID[score_partDB_ID_atBackupTime] = score_partDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Score_part file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Score_part>id_atBckpTime_newID
// to compute new index
func (backRepoScore_part *BackRepoScore_partStruct) RestorePhaseTwo() {

	for _, score_partDB := range backRepoScore_part.Map_Score_partDBID_Score_partDB {

		// next line of code is to avert unused variable compilation error
		_ = score_partDB

		// insertion point for reindexing pointers encoding
		// reindexing Identification field
		if score_partDB.IdentificationID.Int64 != 0 {
			score_partDB.IdentificationID.Int64 = int64(BackRepoIdentificationid_atBckpTime_newID[uint(score_partDB.IdentificationID.Int64)])
			score_partDB.IdentificationID.Valid = true
		}

		// reindexing Part_name_display field
		if score_partDB.Part_name_displayID.Int64 != 0 {
			score_partDB.Part_name_displayID.Int64 = int64(BackRepoName_displayid_atBckpTime_newID[uint(score_partDB.Part_name_displayID.Int64)])
			score_partDB.Part_name_displayID.Valid = true
		}

		// reindexing Part_abbreviation_display field
		if score_partDB.Part_abbreviation_displayID.Int64 != 0 {
			score_partDB.Part_abbreviation_displayID.Int64 = int64(BackRepoName_displayid_atBckpTime_newID[uint(score_partDB.Part_abbreviation_displayID.Int64)])
			score_partDB.Part_abbreviation_displayID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoScore_part.db.Model(score_partDB).Updates(*score_partDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoScore_part.ResetReversePointers commits all staged instances of Score_part to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoScore_part *BackRepoScore_partStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, score_part := range backRepoScore_part.Map_Score_partDBID_Score_partPtr {
		backRepoScore_part.ResetReversePointersInstance(backRepo, idx, score_part)
	}

	return
}

func (backRepoScore_part *BackRepoScore_partStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, score_part *models.Score_part) (Error error) {

	// fetch matching score_partDB
	if score_partDB, ok := backRepoScore_part.Map_Score_partDBID_Score_partDB[idx]; ok {
		_ = score_partDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoScore_partid_atBckpTime_newID map[uint]uint
