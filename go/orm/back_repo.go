// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongmusicxml/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAccidental BackRepoAccidentalStruct

	BackRepoAccidental_mark BackRepoAccidental_markStruct

	BackRepoAccidental_text BackRepoAccidental_textStruct

	BackRepoAccord BackRepoAccordStruct

	BackRepoAccordion_registration BackRepoAccordion_registrationStruct

	BackRepoAnyType BackRepoAnyTypeStruct

	BackRepoAppearance BackRepoAppearanceStruct

	BackRepoArpeggiate BackRepoArpeggiateStruct

	BackRepoArrow BackRepoArrowStruct

	BackRepoArticulations BackRepoArticulationsStruct

	BackRepoAssess BackRepoAssessStruct

	BackRepoAttributes BackRepoAttributesStruct

	BackRepoBackup BackRepoBackupStruct

	BackRepoBar_style_color BackRepoBar_style_colorStruct

	BackRepoBarline BackRepoBarlineStruct

	BackRepoBarre BackRepoBarreStruct

	BackRepoBass BackRepoBassStruct

	BackRepoBass_step BackRepoBass_stepStruct

	BackRepoBeam BackRepoBeamStruct

	BackRepoBeat_repeat BackRepoBeat_repeatStruct

	BackRepoBeat_unit_tied BackRepoBeat_unit_tiedStruct

	BackRepoBeater BackRepoBeaterStruct

	BackRepoBend BackRepoBendStruct

	BackRepoBookmark BackRepoBookmarkStruct

	BackRepoBracket BackRepoBracketStruct

	BackRepoBreath_mark BackRepoBreath_markStruct

	BackRepoCaesura BackRepoCaesuraStruct

	BackRepoCancel BackRepoCancelStruct

	BackRepoClef BackRepoClefStruct

	BackRepoCoda BackRepoCodaStruct

	BackRepoCredit BackRepoCreditStruct

	BackRepoDashes BackRepoDashesStruct

	BackRepoDefaults BackRepoDefaultsStruct

	BackRepoDegree BackRepoDegreeStruct

	BackRepoDegree_alter BackRepoDegree_alterStruct

	BackRepoDegree_type BackRepoDegree_typeStruct

	BackRepoDegree_value BackRepoDegree_valueStruct

	BackRepoDirection BackRepoDirectionStruct

	BackRepoDirection_type BackRepoDirection_typeStruct

	BackRepoDistance BackRepoDistanceStruct

	BackRepoDouble BackRepoDoubleStruct

	BackRepoDynamics BackRepoDynamicsStruct

	BackRepoEffect BackRepoEffectStruct

	BackRepoElision BackRepoElisionStruct

	BackRepoEmpty BackRepoEmptyStruct

	BackRepoEmpty_font BackRepoEmpty_fontStruct

	BackRepoEmpty_line BackRepoEmpty_lineStruct

	BackRepoEmpty_placement BackRepoEmpty_placementStruct

	BackRepoEmpty_placement_smufl BackRepoEmpty_placement_smuflStruct

	BackRepoEmpty_print_object_style_align BackRepoEmpty_print_object_style_alignStruct

	BackRepoEmpty_print_style BackRepoEmpty_print_styleStruct

	BackRepoEmpty_print_style_align BackRepoEmpty_print_style_alignStruct

	BackRepoEmpty_print_style_align_id BackRepoEmpty_print_style_align_idStruct

	BackRepoEmpty_trill_sound BackRepoEmpty_trill_soundStruct

	BackRepoEncoding BackRepoEncodingStruct

	BackRepoEnding BackRepoEndingStruct

	BackRepoExtend BackRepoExtendStruct

	BackRepoFeature BackRepoFeatureStruct

	BackRepoFermata BackRepoFermataStruct

	BackRepoFigure BackRepoFigureStruct

	BackRepoFigured_bass BackRepoFigured_bassStruct

	BackRepoFingering BackRepoFingeringStruct

	BackRepoFirst_fret BackRepoFirst_fretStruct

	BackRepoFoo BackRepoFooStruct

	BackRepoFor_part BackRepoFor_partStruct

	BackRepoFormatted_symbol BackRepoFormatted_symbolStruct

	BackRepoFormatted_symbol_id BackRepoFormatted_symbol_idStruct

	BackRepoForward BackRepoForwardStruct

	BackRepoFrame BackRepoFrameStruct

	BackRepoFrame_note BackRepoFrame_noteStruct

	BackRepoFret BackRepoFretStruct

	BackRepoGlass BackRepoGlassStruct

	BackRepoGlissando BackRepoGlissandoStruct

	BackRepoGlyph BackRepoGlyphStruct

	BackRepoGrace BackRepoGraceStruct

	BackRepoGroup_barline BackRepoGroup_barlineStruct

	BackRepoGroup_symbol BackRepoGroup_symbolStruct

	BackRepoGrouping BackRepoGroupingStruct

	BackRepoHammer_on_pull_off BackRepoHammer_on_pull_offStruct

	BackRepoHandbell BackRepoHandbellStruct

	BackRepoHarmon_closed BackRepoHarmon_closedStruct

	BackRepoHarmon_mute BackRepoHarmon_muteStruct

	BackRepoHarmonic BackRepoHarmonicStruct

	BackRepoHarmony BackRepoHarmonyStruct

	BackRepoHarmony_alter BackRepoHarmony_alterStruct

	BackRepoHarp_pedals BackRepoHarp_pedalsStruct

	BackRepoHeel_toe BackRepoHeel_toeStruct

	BackRepoHole BackRepoHoleStruct

	BackRepoHole_closed BackRepoHole_closedStruct

	BackRepoHorizontal_turn BackRepoHorizontal_turnStruct

	BackRepoIdentification BackRepoIdentificationStruct

	BackRepoImage BackRepoImageStruct

	BackRepoInstrument BackRepoInstrumentStruct

	BackRepoInstrument_change BackRepoInstrument_changeStruct

	BackRepoInstrument_link BackRepoInstrument_linkStruct

	BackRepoInterchangeable BackRepoInterchangeableStruct

	BackRepoInversion BackRepoInversionStruct

	BackRepoKey BackRepoKeyStruct

	BackRepoKey_accidental BackRepoKey_accidentalStruct

	BackRepoKey_octave BackRepoKey_octaveStruct

	BackRepoKind BackRepoKindStruct

	BackRepoLevel BackRepoLevelStruct

	BackRepoLine_detail BackRepoLine_detailStruct

	BackRepoLine_width BackRepoLine_widthStruct

	BackRepoLink BackRepoLinkStruct

	BackRepoListen BackRepoListenStruct

	BackRepoListening BackRepoListeningStruct

	BackRepoLyric BackRepoLyricStruct

	BackRepoLyric_font BackRepoLyric_fontStruct

	BackRepoLyric_language BackRepoLyric_languageStruct

	BackRepoMeasure_layout BackRepoMeasure_layoutStruct

	BackRepoMeasure_numbering BackRepoMeasure_numberingStruct

	BackRepoMeasure_repeat BackRepoMeasure_repeatStruct

	BackRepoMeasure_style BackRepoMeasure_styleStruct

	BackRepoMembrane BackRepoMembraneStruct

	BackRepoMetal BackRepoMetalStruct

	BackRepoMetronome BackRepoMetronomeStruct

	BackRepoMetronome_beam BackRepoMetronome_beamStruct

	BackRepoMetronome_note BackRepoMetronome_noteStruct

	BackRepoMetronome_tied BackRepoMetronome_tiedStruct

	BackRepoMetronome_tuplet BackRepoMetronome_tupletStruct

	BackRepoMidi_device BackRepoMidi_deviceStruct

	BackRepoMidi_instrument BackRepoMidi_instrumentStruct

	BackRepoMiscellaneous BackRepoMiscellaneousStruct

	BackRepoMiscellaneous_field BackRepoMiscellaneous_fieldStruct

	BackRepoMordent BackRepoMordentStruct

	BackRepoMultiple_rest BackRepoMultiple_restStruct

	BackRepoName_display BackRepoName_displayStruct

	BackRepoNon_arpeggiate BackRepoNon_arpeggiateStruct

	BackRepoNotations BackRepoNotationsStruct

	BackRepoNote BackRepoNoteStruct

	BackRepoNote_size BackRepoNote_sizeStruct

	BackRepoNote_type BackRepoNote_typeStruct

	BackRepoNotehead BackRepoNoteheadStruct

	BackRepoNotehead_text BackRepoNotehead_textStruct

	BackRepoNumeral BackRepoNumeralStruct

	BackRepoNumeral_key BackRepoNumeral_keyStruct

	BackRepoNumeral_root BackRepoNumeral_rootStruct

	BackRepoOctave_shift BackRepoOctave_shiftStruct

	BackRepoOffset BackRepoOffsetStruct

	BackRepoOpus BackRepoOpusStruct

	BackRepoOrnaments BackRepoOrnamentsStruct

	BackRepoOther_appearance BackRepoOther_appearanceStruct

	BackRepoOther_listening BackRepoOther_listeningStruct

	BackRepoOther_notation BackRepoOther_notationStruct

	BackRepoOther_play BackRepoOther_playStruct

	BackRepoPage_layout BackRepoPage_layoutStruct

	BackRepoPage_margins BackRepoPage_marginsStruct

	BackRepoPart_clef BackRepoPart_clefStruct

	BackRepoPart_group BackRepoPart_groupStruct

	BackRepoPart_link BackRepoPart_linkStruct

	BackRepoPart_list BackRepoPart_listStruct

	BackRepoPart_symbol BackRepoPart_symbolStruct

	BackRepoPart_transpose BackRepoPart_transposeStruct

	BackRepoPedal BackRepoPedalStruct

	BackRepoPedal_tuning BackRepoPedal_tuningStruct

	BackRepoPercussion BackRepoPercussionStruct

	BackRepoPitch BackRepoPitchStruct

	BackRepoPitched BackRepoPitchedStruct

	BackRepoPlay BackRepoPlayStruct

	BackRepoPlayer BackRepoPlayerStruct

	BackRepoPrincipal_voice BackRepoPrincipal_voiceStruct

	BackRepoPrint BackRepoPrintStruct

	BackRepoRelease BackRepoReleaseStruct

	BackRepoRepeat BackRepoRepeatStruct

	BackRepoRest BackRepoRestStruct

	BackRepoRoot BackRepoRootStruct

	BackRepoRoot_step BackRepoRoot_stepStruct

	BackRepoScaling BackRepoScalingStruct

	BackRepoScordatura BackRepoScordaturaStruct

	BackRepoScore_instrument BackRepoScore_instrumentStruct

	BackRepoScore_part BackRepoScore_partStruct

	BackRepoScore_partwise BackRepoScore_partwiseStruct

	BackRepoScore_timewise BackRepoScore_timewiseStruct

	BackRepoSegno BackRepoSegnoStruct

	BackRepoSlash BackRepoSlashStruct

	BackRepoSlide BackRepoSlideStruct

	BackRepoSlur BackRepoSlurStruct

	BackRepoSound BackRepoSoundStruct

	BackRepoStaff_details BackRepoStaff_detailsStruct

	BackRepoStaff_divide BackRepoStaff_divideStruct

	BackRepoStaff_layout BackRepoStaff_layoutStruct

	BackRepoStaff_size BackRepoStaff_sizeStruct

	BackRepoStaff_tuning BackRepoStaff_tuningStruct

	BackRepoStem BackRepoStemStruct

	BackRepoStick BackRepoStickStruct

	BackRepoString_mute BackRepoString_muteStruct

	BackRepoStrong_accent BackRepoStrong_accentStruct

	BackRepoSupports BackRepoSupportsStruct

	BackRepoSwing BackRepoSwingStruct

	BackRepoSync BackRepoSyncStruct

	BackRepoSystem_dividers BackRepoSystem_dividersStruct

	BackRepoSystem_layout BackRepoSystem_layoutStruct

	BackRepoSystem_margins BackRepoSystem_marginsStruct

	BackRepoTap BackRepoTapStruct

	BackRepoTechnical BackRepoTechnicalStruct

	BackRepoText_element_data BackRepoText_element_dataStruct

	BackRepoTie BackRepoTieStruct

	BackRepoTied BackRepoTiedStruct

	BackRepoTime BackRepoTimeStruct

	BackRepoTime_modification BackRepoTime_modificationStruct

	BackRepoTimpani BackRepoTimpaniStruct

	BackRepoTranspose BackRepoTransposeStruct

	BackRepoTremolo BackRepoTremoloStruct

	BackRepoTuplet BackRepoTupletStruct

	BackRepoTuplet_dot BackRepoTuplet_dotStruct

	BackRepoTuplet_number BackRepoTuplet_numberStruct

	BackRepoTuplet_portion BackRepoTuplet_portionStruct

	BackRepoTuplet_type BackRepoTuplet_typeStruct

	BackRepoTyped_text BackRepoTyped_textStruct

	BackRepoUnpitched BackRepoUnpitchedStruct

	BackRepoVirtual_instrument BackRepoVirtual_instrumentStruct

	BackRepoWait BackRepoWaitStruct

	BackRepoWavy_line BackRepoWavy_lineStruct

	BackRepoWedge BackRepoWedgeStruct

	BackRepoWood BackRepoWoodStruct

	BackRepoWork BackRepoWorkStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&AccidentalDB{},
		&Accidental_markDB{},
		&Accidental_textDB{},
		&AccordDB{},
		&Accordion_registrationDB{},
		&AnyTypeDB{},
		&AppearanceDB{},
		&ArpeggiateDB{},
		&ArrowDB{},
		&ArticulationsDB{},
		&AssessDB{},
		&AttributesDB{},
		&BackupDB{},
		&Bar_style_colorDB{},
		&BarlineDB{},
		&BarreDB{},
		&BassDB{},
		&Bass_stepDB{},
		&BeamDB{},
		&Beat_repeatDB{},
		&Beat_unit_tiedDB{},
		&BeaterDB{},
		&BendDB{},
		&BookmarkDB{},
		&BracketDB{},
		&Breath_markDB{},
		&CaesuraDB{},
		&CancelDB{},
		&ClefDB{},
		&CodaDB{},
		&CreditDB{},
		&DashesDB{},
		&DefaultsDB{},
		&DegreeDB{},
		&Degree_alterDB{},
		&Degree_typeDB{},
		&Degree_valueDB{},
		&DirectionDB{},
		&Direction_typeDB{},
		&DistanceDB{},
		&DoubleDB{},
		&DynamicsDB{},
		&EffectDB{},
		&ElisionDB{},
		&EmptyDB{},
		&Empty_fontDB{},
		&Empty_lineDB{},
		&Empty_placementDB{},
		&Empty_placement_smuflDB{},
		&Empty_print_object_style_alignDB{},
		&Empty_print_styleDB{},
		&Empty_print_style_alignDB{},
		&Empty_print_style_align_idDB{},
		&Empty_trill_soundDB{},
		&EncodingDB{},
		&EndingDB{},
		&ExtendDB{},
		&FeatureDB{},
		&FermataDB{},
		&FigureDB{},
		&Figured_bassDB{},
		&FingeringDB{},
		&First_fretDB{},
		&FooDB{},
		&For_partDB{},
		&Formatted_symbolDB{},
		&Formatted_symbol_idDB{},
		&ForwardDB{},
		&FrameDB{},
		&Frame_noteDB{},
		&FretDB{},
		&GlassDB{},
		&GlissandoDB{},
		&GlyphDB{},
		&GraceDB{},
		&Group_barlineDB{},
		&Group_symbolDB{},
		&GroupingDB{},
		&Hammer_on_pull_offDB{},
		&HandbellDB{},
		&Harmon_closedDB{},
		&Harmon_muteDB{},
		&HarmonicDB{},
		&HarmonyDB{},
		&Harmony_alterDB{},
		&Harp_pedalsDB{},
		&Heel_toeDB{},
		&HoleDB{},
		&Hole_closedDB{},
		&Horizontal_turnDB{},
		&IdentificationDB{},
		&ImageDB{},
		&InstrumentDB{},
		&Instrument_changeDB{},
		&Instrument_linkDB{},
		&InterchangeableDB{},
		&InversionDB{},
		&KeyDB{},
		&Key_accidentalDB{},
		&Key_octaveDB{},
		&KindDB{},
		&LevelDB{},
		&Line_detailDB{},
		&Line_widthDB{},
		&LinkDB{},
		&ListenDB{},
		&ListeningDB{},
		&LyricDB{},
		&Lyric_fontDB{},
		&Lyric_languageDB{},
		&Measure_layoutDB{},
		&Measure_numberingDB{},
		&Measure_repeatDB{},
		&Measure_styleDB{},
		&MembraneDB{},
		&MetalDB{},
		&MetronomeDB{},
		&Metronome_beamDB{},
		&Metronome_noteDB{},
		&Metronome_tiedDB{},
		&Metronome_tupletDB{},
		&Midi_deviceDB{},
		&Midi_instrumentDB{},
		&MiscellaneousDB{},
		&Miscellaneous_fieldDB{},
		&MordentDB{},
		&Multiple_restDB{},
		&Name_displayDB{},
		&Non_arpeggiateDB{},
		&NotationsDB{},
		&NoteDB{},
		&Note_sizeDB{},
		&Note_typeDB{},
		&NoteheadDB{},
		&Notehead_textDB{},
		&NumeralDB{},
		&Numeral_keyDB{},
		&Numeral_rootDB{},
		&Octave_shiftDB{},
		&OffsetDB{},
		&OpusDB{},
		&OrnamentsDB{},
		&Other_appearanceDB{},
		&Other_listeningDB{},
		&Other_notationDB{},
		&Other_playDB{},
		&Page_layoutDB{},
		&Page_marginsDB{},
		&Part_clefDB{},
		&Part_groupDB{},
		&Part_linkDB{},
		&Part_listDB{},
		&Part_symbolDB{},
		&Part_transposeDB{},
		&PedalDB{},
		&Pedal_tuningDB{},
		&PercussionDB{},
		&PitchDB{},
		&PitchedDB{},
		&PlayDB{},
		&PlayerDB{},
		&Principal_voiceDB{},
		&PrintDB{},
		&ReleaseDB{},
		&RepeatDB{},
		&RestDB{},
		&RootDB{},
		&Root_stepDB{},
		&ScalingDB{},
		&ScordaturaDB{},
		&Score_instrumentDB{},
		&Score_partDB{},
		&Score_partwiseDB{},
		&Score_timewiseDB{},
		&SegnoDB{},
		&SlashDB{},
		&SlideDB{},
		&SlurDB{},
		&SoundDB{},
		&Staff_detailsDB{},
		&Staff_divideDB{},
		&Staff_layoutDB{},
		&Staff_sizeDB{},
		&Staff_tuningDB{},
		&StemDB{},
		&StickDB{},
		&String_muteDB{},
		&Strong_accentDB{},
		&SupportsDB{},
		&SwingDB{},
		&SyncDB{},
		&System_dividersDB{},
		&System_layoutDB{},
		&System_marginsDB{},
		&TapDB{},
		&TechnicalDB{},
		&Text_element_dataDB{},
		&TieDB{},
		&TiedDB{},
		&TimeDB{},
		&Time_modificationDB{},
		&TimpaniDB{},
		&TransposeDB{},
		&TremoloDB{},
		&TupletDB{},
		&Tuplet_dotDB{},
		&Tuplet_numberDB{},
		&Tuplet_portionDB{},
		&Tuplet_typeDB{},
		&Typed_textDB{},
		&UnpitchedDB{},
		&Virtual_instrumentDB{},
		&WaitDB{},
		&Wavy_lineDB{},
		&WedgeDB{},
		&WoodDB{},
		&WorkDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAccidental = BackRepoAccidentalStruct{
		Map_AccidentalDBID_AccidentalPtr: make(map[uint]*models.Accidental, 0),
		Map_AccidentalDBID_AccidentalDB:  make(map[uint]*AccidentalDB, 0),
		Map_AccidentalPtr_AccidentalDBID: make(map[*models.Accidental]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAccidental_mark = BackRepoAccidental_markStruct{
		Map_Accidental_markDBID_Accidental_markPtr: make(map[uint]*models.Accidental_mark, 0),
		Map_Accidental_markDBID_Accidental_markDB:  make(map[uint]*Accidental_markDB, 0),
		Map_Accidental_markPtr_Accidental_markDBID: make(map[*models.Accidental_mark]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAccidental_text = BackRepoAccidental_textStruct{
		Map_Accidental_textDBID_Accidental_textPtr: make(map[uint]*models.Accidental_text, 0),
		Map_Accidental_textDBID_Accidental_textDB:  make(map[uint]*Accidental_textDB, 0),
		Map_Accidental_textPtr_Accidental_textDBID: make(map[*models.Accidental_text]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAccord = BackRepoAccordStruct{
		Map_AccordDBID_AccordPtr: make(map[uint]*models.Accord, 0),
		Map_AccordDBID_AccordDB:  make(map[uint]*AccordDB, 0),
		Map_AccordPtr_AccordDBID: make(map[*models.Accord]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAccordion_registration = BackRepoAccordion_registrationStruct{
		Map_Accordion_registrationDBID_Accordion_registrationPtr: make(map[uint]*models.Accordion_registration, 0),
		Map_Accordion_registrationDBID_Accordion_registrationDB:  make(map[uint]*Accordion_registrationDB, 0),
		Map_Accordion_registrationPtr_Accordion_registrationDBID: make(map[*models.Accordion_registration]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAnyType = BackRepoAnyTypeStruct{
		Map_AnyTypeDBID_AnyTypePtr: make(map[uint]*models.AnyType, 0),
		Map_AnyTypeDBID_AnyTypeDB:  make(map[uint]*AnyTypeDB, 0),
		Map_AnyTypePtr_AnyTypeDBID: make(map[*models.AnyType]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAppearance = BackRepoAppearanceStruct{
		Map_AppearanceDBID_AppearancePtr: make(map[uint]*models.Appearance, 0),
		Map_AppearanceDBID_AppearanceDB:  make(map[uint]*AppearanceDB, 0),
		Map_AppearancePtr_AppearanceDBID: make(map[*models.Appearance]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoArpeggiate = BackRepoArpeggiateStruct{
		Map_ArpeggiateDBID_ArpeggiatePtr: make(map[uint]*models.Arpeggiate, 0),
		Map_ArpeggiateDBID_ArpeggiateDB:  make(map[uint]*ArpeggiateDB, 0),
		Map_ArpeggiatePtr_ArpeggiateDBID: make(map[*models.Arpeggiate]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoArrow = BackRepoArrowStruct{
		Map_ArrowDBID_ArrowPtr: make(map[uint]*models.Arrow, 0),
		Map_ArrowDBID_ArrowDB:  make(map[uint]*ArrowDB, 0),
		Map_ArrowPtr_ArrowDBID: make(map[*models.Arrow]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoArticulations = BackRepoArticulationsStruct{
		Map_ArticulationsDBID_ArticulationsPtr: make(map[uint]*models.Articulations, 0),
		Map_ArticulationsDBID_ArticulationsDB:  make(map[uint]*ArticulationsDB, 0),
		Map_ArticulationsPtr_ArticulationsDBID: make(map[*models.Articulations]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAssess = BackRepoAssessStruct{
		Map_AssessDBID_AssessPtr: make(map[uint]*models.Assess, 0),
		Map_AssessDBID_AssessDB:  make(map[uint]*AssessDB, 0),
		Map_AssessPtr_AssessDBID: make(map[*models.Assess]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAttributes = BackRepoAttributesStruct{
		Map_AttributesDBID_AttributesPtr: make(map[uint]*models.Attributes, 0),
		Map_AttributesDBID_AttributesDB:  make(map[uint]*AttributesDB, 0),
		Map_AttributesPtr_AttributesDBID: make(map[*models.Attributes]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBackup = BackRepoBackupStruct{
		Map_BackupDBID_BackupPtr: make(map[uint]*models.Backup, 0),
		Map_BackupDBID_BackupDB:  make(map[uint]*BackupDB, 0),
		Map_BackupPtr_BackupDBID: make(map[*models.Backup]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBar_style_color = BackRepoBar_style_colorStruct{
		Map_Bar_style_colorDBID_Bar_style_colorPtr: make(map[uint]*models.Bar_style_color, 0),
		Map_Bar_style_colorDBID_Bar_style_colorDB:  make(map[uint]*Bar_style_colorDB, 0),
		Map_Bar_style_colorPtr_Bar_style_colorDBID: make(map[*models.Bar_style_color]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBarline = BackRepoBarlineStruct{
		Map_BarlineDBID_BarlinePtr: make(map[uint]*models.Barline, 0),
		Map_BarlineDBID_BarlineDB:  make(map[uint]*BarlineDB, 0),
		Map_BarlinePtr_BarlineDBID: make(map[*models.Barline]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBarre = BackRepoBarreStruct{
		Map_BarreDBID_BarrePtr: make(map[uint]*models.Barre, 0),
		Map_BarreDBID_BarreDB:  make(map[uint]*BarreDB, 0),
		Map_BarrePtr_BarreDBID: make(map[*models.Barre]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBass = BackRepoBassStruct{
		Map_BassDBID_BassPtr: make(map[uint]*models.Bass, 0),
		Map_BassDBID_BassDB:  make(map[uint]*BassDB, 0),
		Map_BassPtr_BassDBID: make(map[*models.Bass]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBass_step = BackRepoBass_stepStruct{
		Map_Bass_stepDBID_Bass_stepPtr: make(map[uint]*models.Bass_step, 0),
		Map_Bass_stepDBID_Bass_stepDB:  make(map[uint]*Bass_stepDB, 0),
		Map_Bass_stepPtr_Bass_stepDBID: make(map[*models.Bass_step]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBeam = BackRepoBeamStruct{
		Map_BeamDBID_BeamPtr: make(map[uint]*models.Beam, 0),
		Map_BeamDBID_BeamDB:  make(map[uint]*BeamDB, 0),
		Map_BeamPtr_BeamDBID: make(map[*models.Beam]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBeat_repeat = BackRepoBeat_repeatStruct{
		Map_Beat_repeatDBID_Beat_repeatPtr: make(map[uint]*models.Beat_repeat, 0),
		Map_Beat_repeatDBID_Beat_repeatDB:  make(map[uint]*Beat_repeatDB, 0),
		Map_Beat_repeatPtr_Beat_repeatDBID: make(map[*models.Beat_repeat]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBeat_unit_tied = BackRepoBeat_unit_tiedStruct{
		Map_Beat_unit_tiedDBID_Beat_unit_tiedPtr: make(map[uint]*models.Beat_unit_tied, 0),
		Map_Beat_unit_tiedDBID_Beat_unit_tiedDB:  make(map[uint]*Beat_unit_tiedDB, 0),
		Map_Beat_unit_tiedPtr_Beat_unit_tiedDBID: make(map[*models.Beat_unit_tied]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBeater = BackRepoBeaterStruct{
		Map_BeaterDBID_BeaterPtr: make(map[uint]*models.Beater, 0),
		Map_BeaterDBID_BeaterDB:  make(map[uint]*BeaterDB, 0),
		Map_BeaterPtr_BeaterDBID: make(map[*models.Beater]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBend = BackRepoBendStruct{
		Map_BendDBID_BendPtr: make(map[uint]*models.Bend, 0),
		Map_BendDBID_BendDB:  make(map[uint]*BendDB, 0),
		Map_BendPtr_BendDBID: make(map[*models.Bend]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBookmark = BackRepoBookmarkStruct{
		Map_BookmarkDBID_BookmarkPtr: make(map[uint]*models.Bookmark, 0),
		Map_BookmarkDBID_BookmarkDB:  make(map[uint]*BookmarkDB, 0),
		Map_BookmarkPtr_BookmarkDBID: make(map[*models.Bookmark]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBracket = BackRepoBracketStruct{
		Map_BracketDBID_BracketPtr: make(map[uint]*models.Bracket, 0),
		Map_BracketDBID_BracketDB:  make(map[uint]*BracketDB, 0),
		Map_BracketPtr_BracketDBID: make(map[*models.Bracket]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoBreath_mark = BackRepoBreath_markStruct{
		Map_Breath_markDBID_Breath_markPtr: make(map[uint]*models.Breath_mark, 0),
		Map_Breath_markDBID_Breath_markDB:  make(map[uint]*Breath_markDB, 0),
		Map_Breath_markPtr_Breath_markDBID: make(map[*models.Breath_mark]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCaesura = BackRepoCaesuraStruct{
		Map_CaesuraDBID_CaesuraPtr: make(map[uint]*models.Caesura, 0),
		Map_CaesuraDBID_CaesuraDB:  make(map[uint]*CaesuraDB, 0),
		Map_CaesuraPtr_CaesuraDBID: make(map[*models.Caesura]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCancel = BackRepoCancelStruct{
		Map_CancelDBID_CancelPtr: make(map[uint]*models.Cancel, 0),
		Map_CancelDBID_CancelDB:  make(map[uint]*CancelDB, 0),
		Map_CancelPtr_CancelDBID: make(map[*models.Cancel]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoClef = BackRepoClefStruct{
		Map_ClefDBID_ClefPtr: make(map[uint]*models.Clef, 0),
		Map_ClefDBID_ClefDB:  make(map[uint]*ClefDB, 0),
		Map_ClefPtr_ClefDBID: make(map[*models.Clef]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCoda = BackRepoCodaStruct{
		Map_CodaDBID_CodaPtr: make(map[uint]*models.Coda, 0),
		Map_CodaDBID_CodaDB:  make(map[uint]*CodaDB, 0),
		Map_CodaPtr_CodaDBID: make(map[*models.Coda]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoCredit = BackRepoCreditStruct{
		Map_CreditDBID_CreditPtr: make(map[uint]*models.Credit, 0),
		Map_CreditDBID_CreditDB:  make(map[uint]*CreditDB, 0),
		Map_CreditPtr_CreditDBID: make(map[*models.Credit]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDashes = BackRepoDashesStruct{
		Map_DashesDBID_DashesPtr: make(map[uint]*models.Dashes, 0),
		Map_DashesDBID_DashesDB:  make(map[uint]*DashesDB, 0),
		Map_DashesPtr_DashesDBID: make(map[*models.Dashes]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDefaults = BackRepoDefaultsStruct{
		Map_DefaultsDBID_DefaultsPtr: make(map[uint]*models.Defaults, 0),
		Map_DefaultsDBID_DefaultsDB:  make(map[uint]*DefaultsDB, 0),
		Map_DefaultsPtr_DefaultsDBID: make(map[*models.Defaults]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDegree = BackRepoDegreeStruct{
		Map_DegreeDBID_DegreePtr: make(map[uint]*models.Degree, 0),
		Map_DegreeDBID_DegreeDB:  make(map[uint]*DegreeDB, 0),
		Map_DegreePtr_DegreeDBID: make(map[*models.Degree]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDegree_alter = BackRepoDegree_alterStruct{
		Map_Degree_alterDBID_Degree_alterPtr: make(map[uint]*models.Degree_alter, 0),
		Map_Degree_alterDBID_Degree_alterDB:  make(map[uint]*Degree_alterDB, 0),
		Map_Degree_alterPtr_Degree_alterDBID: make(map[*models.Degree_alter]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDegree_type = BackRepoDegree_typeStruct{
		Map_Degree_typeDBID_Degree_typePtr: make(map[uint]*models.Degree_type, 0),
		Map_Degree_typeDBID_Degree_typeDB:  make(map[uint]*Degree_typeDB, 0),
		Map_Degree_typePtr_Degree_typeDBID: make(map[*models.Degree_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDegree_value = BackRepoDegree_valueStruct{
		Map_Degree_valueDBID_Degree_valuePtr: make(map[uint]*models.Degree_value, 0),
		Map_Degree_valueDBID_Degree_valueDB:  make(map[uint]*Degree_valueDB, 0),
		Map_Degree_valuePtr_Degree_valueDBID: make(map[*models.Degree_value]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDirection = BackRepoDirectionStruct{
		Map_DirectionDBID_DirectionPtr: make(map[uint]*models.Direction, 0),
		Map_DirectionDBID_DirectionDB:  make(map[uint]*DirectionDB, 0),
		Map_DirectionPtr_DirectionDBID: make(map[*models.Direction]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDirection_type = BackRepoDirection_typeStruct{
		Map_Direction_typeDBID_Direction_typePtr: make(map[uint]*models.Direction_type, 0),
		Map_Direction_typeDBID_Direction_typeDB:  make(map[uint]*Direction_typeDB, 0),
		Map_Direction_typePtr_Direction_typeDBID: make(map[*models.Direction_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDistance = BackRepoDistanceStruct{
		Map_DistanceDBID_DistancePtr: make(map[uint]*models.Distance, 0),
		Map_DistanceDBID_DistanceDB:  make(map[uint]*DistanceDB, 0),
		Map_DistancePtr_DistanceDBID: make(map[*models.Distance]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDouble = BackRepoDoubleStruct{
		Map_DoubleDBID_DoublePtr: make(map[uint]*models.Double, 0),
		Map_DoubleDBID_DoubleDB:  make(map[uint]*DoubleDB, 0),
		Map_DoublePtr_DoubleDBID: make(map[*models.Double]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDynamics = BackRepoDynamicsStruct{
		Map_DynamicsDBID_DynamicsPtr: make(map[uint]*models.Dynamics, 0),
		Map_DynamicsDBID_DynamicsDB:  make(map[uint]*DynamicsDB, 0),
		Map_DynamicsPtr_DynamicsDBID: make(map[*models.Dynamics]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEffect = BackRepoEffectStruct{
		Map_EffectDBID_EffectPtr: make(map[uint]*models.Effect, 0),
		Map_EffectDBID_EffectDB:  make(map[uint]*EffectDB, 0),
		Map_EffectPtr_EffectDBID: make(map[*models.Effect]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoElision = BackRepoElisionStruct{
		Map_ElisionDBID_ElisionPtr: make(map[uint]*models.Elision, 0),
		Map_ElisionDBID_ElisionDB:  make(map[uint]*ElisionDB, 0),
		Map_ElisionPtr_ElisionDBID: make(map[*models.Elision]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty = BackRepoEmptyStruct{
		Map_EmptyDBID_EmptyPtr: make(map[uint]*models.Empty, 0),
		Map_EmptyDBID_EmptyDB:  make(map[uint]*EmptyDB, 0),
		Map_EmptyPtr_EmptyDBID: make(map[*models.Empty]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_font = BackRepoEmpty_fontStruct{
		Map_Empty_fontDBID_Empty_fontPtr: make(map[uint]*models.Empty_font, 0),
		Map_Empty_fontDBID_Empty_fontDB:  make(map[uint]*Empty_fontDB, 0),
		Map_Empty_fontPtr_Empty_fontDBID: make(map[*models.Empty_font]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_line = BackRepoEmpty_lineStruct{
		Map_Empty_lineDBID_Empty_linePtr: make(map[uint]*models.Empty_line, 0),
		Map_Empty_lineDBID_Empty_lineDB:  make(map[uint]*Empty_lineDB, 0),
		Map_Empty_linePtr_Empty_lineDBID: make(map[*models.Empty_line]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_placement = BackRepoEmpty_placementStruct{
		Map_Empty_placementDBID_Empty_placementPtr: make(map[uint]*models.Empty_placement, 0),
		Map_Empty_placementDBID_Empty_placementDB:  make(map[uint]*Empty_placementDB, 0),
		Map_Empty_placementPtr_Empty_placementDBID: make(map[*models.Empty_placement]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_placement_smufl = BackRepoEmpty_placement_smuflStruct{
		Map_Empty_placement_smuflDBID_Empty_placement_smuflPtr: make(map[uint]*models.Empty_placement_smufl, 0),
		Map_Empty_placement_smuflDBID_Empty_placement_smuflDB:  make(map[uint]*Empty_placement_smuflDB, 0),
		Map_Empty_placement_smuflPtr_Empty_placement_smuflDBID: make(map[*models.Empty_placement_smufl]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_print_object_style_align = BackRepoEmpty_print_object_style_alignStruct{
		Map_Empty_print_object_style_alignDBID_Empty_print_object_style_alignPtr: make(map[uint]*models.Empty_print_object_style_align, 0),
		Map_Empty_print_object_style_alignDBID_Empty_print_object_style_alignDB:  make(map[uint]*Empty_print_object_style_alignDB, 0),
		Map_Empty_print_object_style_alignPtr_Empty_print_object_style_alignDBID: make(map[*models.Empty_print_object_style_align]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_print_style = BackRepoEmpty_print_styleStruct{
		Map_Empty_print_styleDBID_Empty_print_stylePtr: make(map[uint]*models.Empty_print_style, 0),
		Map_Empty_print_styleDBID_Empty_print_styleDB:  make(map[uint]*Empty_print_styleDB, 0),
		Map_Empty_print_stylePtr_Empty_print_styleDBID: make(map[*models.Empty_print_style]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_print_style_align = BackRepoEmpty_print_style_alignStruct{
		Map_Empty_print_style_alignDBID_Empty_print_style_alignPtr: make(map[uint]*models.Empty_print_style_align, 0),
		Map_Empty_print_style_alignDBID_Empty_print_style_alignDB:  make(map[uint]*Empty_print_style_alignDB, 0),
		Map_Empty_print_style_alignPtr_Empty_print_style_alignDBID: make(map[*models.Empty_print_style_align]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_print_style_align_id = BackRepoEmpty_print_style_align_idStruct{
		Map_Empty_print_style_align_idDBID_Empty_print_style_align_idPtr: make(map[uint]*models.Empty_print_style_align_id, 0),
		Map_Empty_print_style_align_idDBID_Empty_print_style_align_idDB:  make(map[uint]*Empty_print_style_align_idDB, 0),
		Map_Empty_print_style_align_idPtr_Empty_print_style_align_idDBID: make(map[*models.Empty_print_style_align_id]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEmpty_trill_sound = BackRepoEmpty_trill_soundStruct{
		Map_Empty_trill_soundDBID_Empty_trill_soundPtr: make(map[uint]*models.Empty_trill_sound, 0),
		Map_Empty_trill_soundDBID_Empty_trill_soundDB:  make(map[uint]*Empty_trill_soundDB, 0),
		Map_Empty_trill_soundPtr_Empty_trill_soundDBID: make(map[*models.Empty_trill_sound]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEncoding = BackRepoEncodingStruct{
		Map_EncodingDBID_EncodingPtr: make(map[uint]*models.Encoding, 0),
		Map_EncodingDBID_EncodingDB:  make(map[uint]*EncodingDB, 0),
		Map_EncodingPtr_EncodingDBID: make(map[*models.Encoding]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEnding = BackRepoEndingStruct{
		Map_EndingDBID_EndingPtr: make(map[uint]*models.Ending, 0),
		Map_EndingDBID_EndingDB:  make(map[uint]*EndingDB, 0),
		Map_EndingPtr_EndingDBID: make(map[*models.Ending]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoExtend = BackRepoExtendStruct{
		Map_ExtendDBID_ExtendPtr: make(map[uint]*models.Extend, 0),
		Map_ExtendDBID_ExtendDB:  make(map[uint]*ExtendDB, 0),
		Map_ExtendPtr_ExtendDBID: make(map[*models.Extend]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFeature = BackRepoFeatureStruct{
		Map_FeatureDBID_FeaturePtr: make(map[uint]*models.Feature, 0),
		Map_FeatureDBID_FeatureDB:  make(map[uint]*FeatureDB, 0),
		Map_FeaturePtr_FeatureDBID: make(map[*models.Feature]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFermata = BackRepoFermataStruct{
		Map_FermataDBID_FermataPtr: make(map[uint]*models.Fermata, 0),
		Map_FermataDBID_FermataDB:  make(map[uint]*FermataDB, 0),
		Map_FermataPtr_FermataDBID: make(map[*models.Fermata]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFigure = BackRepoFigureStruct{
		Map_FigureDBID_FigurePtr: make(map[uint]*models.Figure, 0),
		Map_FigureDBID_FigureDB:  make(map[uint]*FigureDB, 0),
		Map_FigurePtr_FigureDBID: make(map[*models.Figure]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFigured_bass = BackRepoFigured_bassStruct{
		Map_Figured_bassDBID_Figured_bassPtr: make(map[uint]*models.Figured_bass, 0),
		Map_Figured_bassDBID_Figured_bassDB:  make(map[uint]*Figured_bassDB, 0),
		Map_Figured_bassPtr_Figured_bassDBID: make(map[*models.Figured_bass]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFingering = BackRepoFingeringStruct{
		Map_FingeringDBID_FingeringPtr: make(map[uint]*models.Fingering, 0),
		Map_FingeringDBID_FingeringDB:  make(map[uint]*FingeringDB, 0),
		Map_FingeringPtr_FingeringDBID: make(map[*models.Fingering]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFirst_fret = BackRepoFirst_fretStruct{
		Map_First_fretDBID_First_fretPtr: make(map[uint]*models.First_fret, 0),
		Map_First_fretDBID_First_fretDB:  make(map[uint]*First_fretDB, 0),
		Map_First_fretPtr_First_fretDBID: make(map[*models.First_fret]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFoo = BackRepoFooStruct{
		Map_FooDBID_FooPtr: make(map[uint]*models.Foo, 0),
		Map_FooDBID_FooDB:  make(map[uint]*FooDB, 0),
		Map_FooPtr_FooDBID: make(map[*models.Foo]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFor_part = BackRepoFor_partStruct{
		Map_For_partDBID_For_partPtr: make(map[uint]*models.For_part, 0),
		Map_For_partDBID_For_partDB:  make(map[uint]*For_partDB, 0),
		Map_For_partPtr_For_partDBID: make(map[*models.For_part]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormatted_symbol = BackRepoFormatted_symbolStruct{
		Map_Formatted_symbolDBID_Formatted_symbolPtr: make(map[uint]*models.Formatted_symbol, 0),
		Map_Formatted_symbolDBID_Formatted_symbolDB:  make(map[uint]*Formatted_symbolDB, 0),
		Map_Formatted_symbolPtr_Formatted_symbolDBID: make(map[*models.Formatted_symbol]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFormatted_symbol_id = BackRepoFormatted_symbol_idStruct{
		Map_Formatted_symbol_idDBID_Formatted_symbol_idPtr: make(map[uint]*models.Formatted_symbol_id, 0),
		Map_Formatted_symbol_idDBID_Formatted_symbol_idDB:  make(map[uint]*Formatted_symbol_idDB, 0),
		Map_Formatted_symbol_idPtr_Formatted_symbol_idDBID: make(map[*models.Formatted_symbol_id]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoForward = BackRepoForwardStruct{
		Map_ForwardDBID_ForwardPtr: make(map[uint]*models.Forward, 0),
		Map_ForwardDBID_ForwardDB:  make(map[uint]*ForwardDB, 0),
		Map_ForwardPtr_ForwardDBID: make(map[*models.Forward]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFrame = BackRepoFrameStruct{
		Map_FrameDBID_FramePtr: make(map[uint]*models.Frame, 0),
		Map_FrameDBID_FrameDB:  make(map[uint]*FrameDB, 0),
		Map_FramePtr_FrameDBID: make(map[*models.Frame]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFrame_note = BackRepoFrame_noteStruct{
		Map_Frame_noteDBID_Frame_notePtr: make(map[uint]*models.Frame_note, 0),
		Map_Frame_noteDBID_Frame_noteDB:  make(map[uint]*Frame_noteDB, 0),
		Map_Frame_notePtr_Frame_noteDBID: make(map[*models.Frame_note]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoFret = BackRepoFretStruct{
		Map_FretDBID_FretPtr: make(map[uint]*models.Fret, 0),
		Map_FretDBID_FretDB:  make(map[uint]*FretDB, 0),
		Map_FretPtr_FretDBID: make(map[*models.Fret]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGlass = BackRepoGlassStruct{
		Map_GlassDBID_GlassPtr: make(map[uint]*models.Glass, 0),
		Map_GlassDBID_GlassDB:  make(map[uint]*GlassDB, 0),
		Map_GlassPtr_GlassDBID: make(map[*models.Glass]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGlissando = BackRepoGlissandoStruct{
		Map_GlissandoDBID_GlissandoPtr: make(map[uint]*models.Glissando, 0),
		Map_GlissandoDBID_GlissandoDB:  make(map[uint]*GlissandoDB, 0),
		Map_GlissandoPtr_GlissandoDBID: make(map[*models.Glissando]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGlyph = BackRepoGlyphStruct{
		Map_GlyphDBID_GlyphPtr: make(map[uint]*models.Glyph, 0),
		Map_GlyphDBID_GlyphDB:  make(map[uint]*GlyphDB, 0),
		Map_GlyphPtr_GlyphDBID: make(map[*models.Glyph]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGrace = BackRepoGraceStruct{
		Map_GraceDBID_GracePtr: make(map[uint]*models.Grace, 0),
		Map_GraceDBID_GraceDB:  make(map[uint]*GraceDB, 0),
		Map_GracePtr_GraceDBID: make(map[*models.Grace]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGroup_barline = BackRepoGroup_barlineStruct{
		Map_Group_barlineDBID_Group_barlinePtr: make(map[uint]*models.Group_barline, 0),
		Map_Group_barlineDBID_Group_barlineDB:  make(map[uint]*Group_barlineDB, 0),
		Map_Group_barlinePtr_Group_barlineDBID: make(map[*models.Group_barline]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGroup_symbol = BackRepoGroup_symbolStruct{
		Map_Group_symbolDBID_Group_symbolPtr: make(map[uint]*models.Group_symbol, 0),
		Map_Group_symbolDBID_Group_symbolDB:  make(map[uint]*Group_symbolDB, 0),
		Map_Group_symbolPtr_Group_symbolDBID: make(map[*models.Group_symbol]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGrouping = BackRepoGroupingStruct{
		Map_GroupingDBID_GroupingPtr: make(map[uint]*models.Grouping, 0),
		Map_GroupingDBID_GroupingDB:  make(map[uint]*GroupingDB, 0),
		Map_GroupingPtr_GroupingDBID: make(map[*models.Grouping]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHammer_on_pull_off = BackRepoHammer_on_pull_offStruct{
		Map_Hammer_on_pull_offDBID_Hammer_on_pull_offPtr: make(map[uint]*models.Hammer_on_pull_off, 0),
		Map_Hammer_on_pull_offDBID_Hammer_on_pull_offDB:  make(map[uint]*Hammer_on_pull_offDB, 0),
		Map_Hammer_on_pull_offPtr_Hammer_on_pull_offDBID: make(map[*models.Hammer_on_pull_off]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHandbell = BackRepoHandbellStruct{
		Map_HandbellDBID_HandbellPtr: make(map[uint]*models.Handbell, 0),
		Map_HandbellDBID_HandbellDB:  make(map[uint]*HandbellDB, 0),
		Map_HandbellPtr_HandbellDBID: make(map[*models.Handbell]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHarmon_closed = BackRepoHarmon_closedStruct{
		Map_Harmon_closedDBID_Harmon_closedPtr: make(map[uint]*models.Harmon_closed, 0),
		Map_Harmon_closedDBID_Harmon_closedDB:  make(map[uint]*Harmon_closedDB, 0),
		Map_Harmon_closedPtr_Harmon_closedDBID: make(map[*models.Harmon_closed]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHarmon_mute = BackRepoHarmon_muteStruct{
		Map_Harmon_muteDBID_Harmon_mutePtr: make(map[uint]*models.Harmon_mute, 0),
		Map_Harmon_muteDBID_Harmon_muteDB:  make(map[uint]*Harmon_muteDB, 0),
		Map_Harmon_mutePtr_Harmon_muteDBID: make(map[*models.Harmon_mute]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHarmonic = BackRepoHarmonicStruct{
		Map_HarmonicDBID_HarmonicPtr: make(map[uint]*models.Harmonic, 0),
		Map_HarmonicDBID_HarmonicDB:  make(map[uint]*HarmonicDB, 0),
		Map_HarmonicPtr_HarmonicDBID: make(map[*models.Harmonic]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHarmony = BackRepoHarmonyStruct{
		Map_HarmonyDBID_HarmonyPtr: make(map[uint]*models.Harmony, 0),
		Map_HarmonyDBID_HarmonyDB:  make(map[uint]*HarmonyDB, 0),
		Map_HarmonyPtr_HarmonyDBID: make(map[*models.Harmony]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHarmony_alter = BackRepoHarmony_alterStruct{
		Map_Harmony_alterDBID_Harmony_alterPtr: make(map[uint]*models.Harmony_alter, 0),
		Map_Harmony_alterDBID_Harmony_alterDB:  make(map[uint]*Harmony_alterDB, 0),
		Map_Harmony_alterPtr_Harmony_alterDBID: make(map[*models.Harmony_alter]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHarp_pedals = BackRepoHarp_pedalsStruct{
		Map_Harp_pedalsDBID_Harp_pedalsPtr: make(map[uint]*models.Harp_pedals, 0),
		Map_Harp_pedalsDBID_Harp_pedalsDB:  make(map[uint]*Harp_pedalsDB, 0),
		Map_Harp_pedalsPtr_Harp_pedalsDBID: make(map[*models.Harp_pedals]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHeel_toe = BackRepoHeel_toeStruct{
		Map_Heel_toeDBID_Heel_toePtr: make(map[uint]*models.Heel_toe, 0),
		Map_Heel_toeDBID_Heel_toeDB:  make(map[uint]*Heel_toeDB, 0),
		Map_Heel_toePtr_Heel_toeDBID: make(map[*models.Heel_toe]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHole = BackRepoHoleStruct{
		Map_HoleDBID_HolePtr: make(map[uint]*models.Hole, 0),
		Map_HoleDBID_HoleDB:  make(map[uint]*HoleDB, 0),
		Map_HolePtr_HoleDBID: make(map[*models.Hole]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHole_closed = BackRepoHole_closedStruct{
		Map_Hole_closedDBID_Hole_closedPtr: make(map[uint]*models.Hole_closed, 0),
		Map_Hole_closedDBID_Hole_closedDB:  make(map[uint]*Hole_closedDB, 0),
		Map_Hole_closedPtr_Hole_closedDBID: make(map[*models.Hole_closed]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoHorizontal_turn = BackRepoHorizontal_turnStruct{
		Map_Horizontal_turnDBID_Horizontal_turnPtr: make(map[uint]*models.Horizontal_turn, 0),
		Map_Horizontal_turnDBID_Horizontal_turnDB:  make(map[uint]*Horizontal_turnDB, 0),
		Map_Horizontal_turnPtr_Horizontal_turnDBID: make(map[*models.Horizontal_turn]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoIdentification = BackRepoIdentificationStruct{
		Map_IdentificationDBID_IdentificationPtr: make(map[uint]*models.Identification, 0),
		Map_IdentificationDBID_IdentificationDB:  make(map[uint]*IdentificationDB, 0),
		Map_IdentificationPtr_IdentificationDBID: make(map[*models.Identification]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoImage = BackRepoImageStruct{
		Map_ImageDBID_ImagePtr: make(map[uint]*models.Image, 0),
		Map_ImageDBID_ImageDB:  make(map[uint]*ImageDB, 0),
		Map_ImagePtr_ImageDBID: make(map[*models.Image]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoInstrument = BackRepoInstrumentStruct{
		Map_InstrumentDBID_InstrumentPtr: make(map[uint]*models.Instrument, 0),
		Map_InstrumentDBID_InstrumentDB:  make(map[uint]*InstrumentDB, 0),
		Map_InstrumentPtr_InstrumentDBID: make(map[*models.Instrument]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoInstrument_change = BackRepoInstrument_changeStruct{
		Map_Instrument_changeDBID_Instrument_changePtr: make(map[uint]*models.Instrument_change, 0),
		Map_Instrument_changeDBID_Instrument_changeDB:  make(map[uint]*Instrument_changeDB, 0),
		Map_Instrument_changePtr_Instrument_changeDBID: make(map[*models.Instrument_change]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoInstrument_link = BackRepoInstrument_linkStruct{
		Map_Instrument_linkDBID_Instrument_linkPtr: make(map[uint]*models.Instrument_link, 0),
		Map_Instrument_linkDBID_Instrument_linkDB:  make(map[uint]*Instrument_linkDB, 0),
		Map_Instrument_linkPtr_Instrument_linkDBID: make(map[*models.Instrument_link]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoInterchangeable = BackRepoInterchangeableStruct{
		Map_InterchangeableDBID_InterchangeablePtr: make(map[uint]*models.Interchangeable, 0),
		Map_InterchangeableDBID_InterchangeableDB:  make(map[uint]*InterchangeableDB, 0),
		Map_InterchangeablePtr_InterchangeableDBID: make(map[*models.Interchangeable]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoInversion = BackRepoInversionStruct{
		Map_InversionDBID_InversionPtr: make(map[uint]*models.Inversion, 0),
		Map_InversionDBID_InversionDB:  make(map[uint]*InversionDB, 0),
		Map_InversionPtr_InversionDBID: make(map[*models.Inversion]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKey = BackRepoKeyStruct{
		Map_KeyDBID_KeyPtr: make(map[uint]*models.Key, 0),
		Map_KeyDBID_KeyDB:  make(map[uint]*KeyDB, 0),
		Map_KeyPtr_KeyDBID: make(map[*models.Key]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKey_accidental = BackRepoKey_accidentalStruct{
		Map_Key_accidentalDBID_Key_accidentalPtr: make(map[uint]*models.Key_accidental, 0),
		Map_Key_accidentalDBID_Key_accidentalDB:  make(map[uint]*Key_accidentalDB, 0),
		Map_Key_accidentalPtr_Key_accidentalDBID: make(map[*models.Key_accidental]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKey_octave = BackRepoKey_octaveStruct{
		Map_Key_octaveDBID_Key_octavePtr: make(map[uint]*models.Key_octave, 0),
		Map_Key_octaveDBID_Key_octaveDB:  make(map[uint]*Key_octaveDB, 0),
		Map_Key_octavePtr_Key_octaveDBID: make(map[*models.Key_octave]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoKind = BackRepoKindStruct{
		Map_KindDBID_KindPtr: make(map[uint]*models.Kind, 0),
		Map_KindDBID_KindDB:  make(map[uint]*KindDB, 0),
		Map_KindPtr_KindDBID: make(map[*models.Kind]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLevel = BackRepoLevelStruct{
		Map_LevelDBID_LevelPtr: make(map[uint]*models.Level, 0),
		Map_LevelDBID_LevelDB:  make(map[uint]*LevelDB, 0),
		Map_LevelPtr_LevelDBID: make(map[*models.Level]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLine_detail = BackRepoLine_detailStruct{
		Map_Line_detailDBID_Line_detailPtr: make(map[uint]*models.Line_detail, 0),
		Map_Line_detailDBID_Line_detailDB:  make(map[uint]*Line_detailDB, 0),
		Map_Line_detailPtr_Line_detailDBID: make(map[*models.Line_detail]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLine_width = BackRepoLine_widthStruct{
		Map_Line_widthDBID_Line_widthPtr: make(map[uint]*models.Line_width, 0),
		Map_Line_widthDBID_Line_widthDB:  make(map[uint]*Line_widthDB, 0),
		Map_Line_widthPtr_Line_widthDBID: make(map[*models.Line_width]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLink = BackRepoLinkStruct{
		Map_LinkDBID_LinkPtr: make(map[uint]*models.Link, 0),
		Map_LinkDBID_LinkDB:  make(map[uint]*LinkDB, 0),
		Map_LinkPtr_LinkDBID: make(map[*models.Link]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoListen = BackRepoListenStruct{
		Map_ListenDBID_ListenPtr: make(map[uint]*models.Listen, 0),
		Map_ListenDBID_ListenDB:  make(map[uint]*ListenDB, 0),
		Map_ListenPtr_ListenDBID: make(map[*models.Listen]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoListening = BackRepoListeningStruct{
		Map_ListeningDBID_ListeningPtr: make(map[uint]*models.Listening, 0),
		Map_ListeningDBID_ListeningDB:  make(map[uint]*ListeningDB, 0),
		Map_ListeningPtr_ListeningDBID: make(map[*models.Listening]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLyric = BackRepoLyricStruct{
		Map_LyricDBID_LyricPtr: make(map[uint]*models.Lyric, 0),
		Map_LyricDBID_LyricDB:  make(map[uint]*LyricDB, 0),
		Map_LyricPtr_LyricDBID: make(map[*models.Lyric]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLyric_font = BackRepoLyric_fontStruct{
		Map_Lyric_fontDBID_Lyric_fontPtr: make(map[uint]*models.Lyric_font, 0),
		Map_Lyric_fontDBID_Lyric_fontDB:  make(map[uint]*Lyric_fontDB, 0),
		Map_Lyric_fontPtr_Lyric_fontDBID: make(map[*models.Lyric_font]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLyric_language = BackRepoLyric_languageStruct{
		Map_Lyric_languageDBID_Lyric_languagePtr: make(map[uint]*models.Lyric_language, 0),
		Map_Lyric_languageDBID_Lyric_languageDB:  make(map[uint]*Lyric_languageDB, 0),
		Map_Lyric_languagePtr_Lyric_languageDBID: make(map[*models.Lyric_language]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeasure_layout = BackRepoMeasure_layoutStruct{
		Map_Measure_layoutDBID_Measure_layoutPtr: make(map[uint]*models.Measure_layout, 0),
		Map_Measure_layoutDBID_Measure_layoutDB:  make(map[uint]*Measure_layoutDB, 0),
		Map_Measure_layoutPtr_Measure_layoutDBID: make(map[*models.Measure_layout]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeasure_numbering = BackRepoMeasure_numberingStruct{
		Map_Measure_numberingDBID_Measure_numberingPtr: make(map[uint]*models.Measure_numbering, 0),
		Map_Measure_numberingDBID_Measure_numberingDB:  make(map[uint]*Measure_numberingDB, 0),
		Map_Measure_numberingPtr_Measure_numberingDBID: make(map[*models.Measure_numbering]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeasure_repeat = BackRepoMeasure_repeatStruct{
		Map_Measure_repeatDBID_Measure_repeatPtr: make(map[uint]*models.Measure_repeat, 0),
		Map_Measure_repeatDBID_Measure_repeatDB:  make(map[uint]*Measure_repeatDB, 0),
		Map_Measure_repeatPtr_Measure_repeatDBID: make(map[*models.Measure_repeat]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMeasure_style = BackRepoMeasure_styleStruct{
		Map_Measure_styleDBID_Measure_stylePtr: make(map[uint]*models.Measure_style, 0),
		Map_Measure_styleDBID_Measure_styleDB:  make(map[uint]*Measure_styleDB, 0),
		Map_Measure_stylePtr_Measure_styleDBID: make(map[*models.Measure_style]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMembrane = BackRepoMembraneStruct{
		Map_MembraneDBID_MembranePtr: make(map[uint]*models.Membrane, 0),
		Map_MembraneDBID_MembraneDB:  make(map[uint]*MembraneDB, 0),
		Map_MembranePtr_MembraneDBID: make(map[*models.Membrane]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetal = BackRepoMetalStruct{
		Map_MetalDBID_MetalPtr: make(map[uint]*models.Metal, 0),
		Map_MetalDBID_MetalDB:  make(map[uint]*MetalDB, 0),
		Map_MetalPtr_MetalDBID: make(map[*models.Metal]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetronome = BackRepoMetronomeStruct{
		Map_MetronomeDBID_MetronomePtr: make(map[uint]*models.Metronome, 0),
		Map_MetronomeDBID_MetronomeDB:  make(map[uint]*MetronomeDB, 0),
		Map_MetronomePtr_MetronomeDBID: make(map[*models.Metronome]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetronome_beam = BackRepoMetronome_beamStruct{
		Map_Metronome_beamDBID_Metronome_beamPtr: make(map[uint]*models.Metronome_beam, 0),
		Map_Metronome_beamDBID_Metronome_beamDB:  make(map[uint]*Metronome_beamDB, 0),
		Map_Metronome_beamPtr_Metronome_beamDBID: make(map[*models.Metronome_beam]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetronome_note = BackRepoMetronome_noteStruct{
		Map_Metronome_noteDBID_Metronome_notePtr: make(map[uint]*models.Metronome_note, 0),
		Map_Metronome_noteDBID_Metronome_noteDB:  make(map[uint]*Metronome_noteDB, 0),
		Map_Metronome_notePtr_Metronome_noteDBID: make(map[*models.Metronome_note]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetronome_tied = BackRepoMetronome_tiedStruct{
		Map_Metronome_tiedDBID_Metronome_tiedPtr: make(map[uint]*models.Metronome_tied, 0),
		Map_Metronome_tiedDBID_Metronome_tiedDB:  make(map[uint]*Metronome_tiedDB, 0),
		Map_Metronome_tiedPtr_Metronome_tiedDBID: make(map[*models.Metronome_tied]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMetronome_tuplet = BackRepoMetronome_tupletStruct{
		Map_Metronome_tupletDBID_Metronome_tupletPtr: make(map[uint]*models.Metronome_tuplet, 0),
		Map_Metronome_tupletDBID_Metronome_tupletDB:  make(map[uint]*Metronome_tupletDB, 0),
		Map_Metronome_tupletPtr_Metronome_tupletDBID: make(map[*models.Metronome_tuplet]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMidi_device = BackRepoMidi_deviceStruct{
		Map_Midi_deviceDBID_Midi_devicePtr: make(map[uint]*models.Midi_device, 0),
		Map_Midi_deviceDBID_Midi_deviceDB:  make(map[uint]*Midi_deviceDB, 0),
		Map_Midi_devicePtr_Midi_deviceDBID: make(map[*models.Midi_device]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMidi_instrument = BackRepoMidi_instrumentStruct{
		Map_Midi_instrumentDBID_Midi_instrumentPtr: make(map[uint]*models.Midi_instrument, 0),
		Map_Midi_instrumentDBID_Midi_instrumentDB:  make(map[uint]*Midi_instrumentDB, 0),
		Map_Midi_instrumentPtr_Midi_instrumentDBID: make(map[*models.Midi_instrument]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMiscellaneous = BackRepoMiscellaneousStruct{
		Map_MiscellaneousDBID_MiscellaneousPtr: make(map[uint]*models.Miscellaneous, 0),
		Map_MiscellaneousDBID_MiscellaneousDB:  make(map[uint]*MiscellaneousDB, 0),
		Map_MiscellaneousPtr_MiscellaneousDBID: make(map[*models.Miscellaneous]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMiscellaneous_field = BackRepoMiscellaneous_fieldStruct{
		Map_Miscellaneous_fieldDBID_Miscellaneous_fieldPtr: make(map[uint]*models.Miscellaneous_field, 0),
		Map_Miscellaneous_fieldDBID_Miscellaneous_fieldDB:  make(map[uint]*Miscellaneous_fieldDB, 0),
		Map_Miscellaneous_fieldPtr_Miscellaneous_fieldDBID: make(map[*models.Miscellaneous_field]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMordent = BackRepoMordentStruct{
		Map_MordentDBID_MordentPtr: make(map[uint]*models.Mordent, 0),
		Map_MordentDBID_MordentDB:  make(map[uint]*MordentDB, 0),
		Map_MordentPtr_MordentDBID: make(map[*models.Mordent]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMultiple_rest = BackRepoMultiple_restStruct{
		Map_Multiple_restDBID_Multiple_restPtr: make(map[uint]*models.Multiple_rest, 0),
		Map_Multiple_restDBID_Multiple_restDB:  make(map[uint]*Multiple_restDB, 0),
		Map_Multiple_restPtr_Multiple_restDBID: make(map[*models.Multiple_rest]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoName_display = BackRepoName_displayStruct{
		Map_Name_displayDBID_Name_displayPtr: make(map[uint]*models.Name_display, 0),
		Map_Name_displayDBID_Name_displayDB:  make(map[uint]*Name_displayDB, 0),
		Map_Name_displayPtr_Name_displayDBID: make(map[*models.Name_display]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNon_arpeggiate = BackRepoNon_arpeggiateStruct{
		Map_Non_arpeggiateDBID_Non_arpeggiatePtr: make(map[uint]*models.Non_arpeggiate, 0),
		Map_Non_arpeggiateDBID_Non_arpeggiateDB:  make(map[uint]*Non_arpeggiateDB, 0),
		Map_Non_arpeggiatePtr_Non_arpeggiateDBID: make(map[*models.Non_arpeggiate]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNotations = BackRepoNotationsStruct{
		Map_NotationsDBID_NotationsPtr: make(map[uint]*models.Notations, 0),
		Map_NotationsDBID_NotationsDB:  make(map[uint]*NotationsDB, 0),
		Map_NotationsPtr_NotationsDBID: make(map[*models.Notations]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNote = BackRepoNoteStruct{
		Map_NoteDBID_NotePtr: make(map[uint]*models.Note, 0),
		Map_NoteDBID_NoteDB:  make(map[uint]*NoteDB, 0),
		Map_NotePtr_NoteDBID: make(map[*models.Note]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNote_size = BackRepoNote_sizeStruct{
		Map_Note_sizeDBID_Note_sizePtr: make(map[uint]*models.Note_size, 0),
		Map_Note_sizeDBID_Note_sizeDB:  make(map[uint]*Note_sizeDB, 0),
		Map_Note_sizePtr_Note_sizeDBID: make(map[*models.Note_size]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNote_type = BackRepoNote_typeStruct{
		Map_Note_typeDBID_Note_typePtr: make(map[uint]*models.Note_type, 0),
		Map_Note_typeDBID_Note_typeDB:  make(map[uint]*Note_typeDB, 0),
		Map_Note_typePtr_Note_typeDBID: make(map[*models.Note_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNotehead = BackRepoNoteheadStruct{
		Map_NoteheadDBID_NoteheadPtr: make(map[uint]*models.Notehead, 0),
		Map_NoteheadDBID_NoteheadDB:  make(map[uint]*NoteheadDB, 0),
		Map_NoteheadPtr_NoteheadDBID: make(map[*models.Notehead]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNotehead_text = BackRepoNotehead_textStruct{
		Map_Notehead_textDBID_Notehead_textPtr: make(map[uint]*models.Notehead_text, 0),
		Map_Notehead_textDBID_Notehead_textDB:  make(map[uint]*Notehead_textDB, 0),
		Map_Notehead_textPtr_Notehead_textDBID: make(map[*models.Notehead_text]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNumeral = BackRepoNumeralStruct{
		Map_NumeralDBID_NumeralPtr: make(map[uint]*models.Numeral, 0),
		Map_NumeralDBID_NumeralDB:  make(map[uint]*NumeralDB, 0),
		Map_NumeralPtr_NumeralDBID: make(map[*models.Numeral]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNumeral_key = BackRepoNumeral_keyStruct{
		Map_Numeral_keyDBID_Numeral_keyPtr: make(map[uint]*models.Numeral_key, 0),
		Map_Numeral_keyDBID_Numeral_keyDB:  make(map[uint]*Numeral_keyDB, 0),
		Map_Numeral_keyPtr_Numeral_keyDBID: make(map[*models.Numeral_key]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoNumeral_root = BackRepoNumeral_rootStruct{
		Map_Numeral_rootDBID_Numeral_rootPtr: make(map[uint]*models.Numeral_root, 0),
		Map_Numeral_rootDBID_Numeral_rootDB:  make(map[uint]*Numeral_rootDB, 0),
		Map_Numeral_rootPtr_Numeral_rootDBID: make(map[*models.Numeral_root]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOctave_shift = BackRepoOctave_shiftStruct{
		Map_Octave_shiftDBID_Octave_shiftPtr: make(map[uint]*models.Octave_shift, 0),
		Map_Octave_shiftDBID_Octave_shiftDB:  make(map[uint]*Octave_shiftDB, 0),
		Map_Octave_shiftPtr_Octave_shiftDBID: make(map[*models.Octave_shift]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOffset = BackRepoOffsetStruct{
		Map_OffsetDBID_OffsetPtr: make(map[uint]*models.Offset, 0),
		Map_OffsetDBID_OffsetDB:  make(map[uint]*OffsetDB, 0),
		Map_OffsetPtr_OffsetDBID: make(map[*models.Offset]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOpus = BackRepoOpusStruct{
		Map_OpusDBID_OpusPtr: make(map[uint]*models.Opus, 0),
		Map_OpusDBID_OpusDB:  make(map[uint]*OpusDB, 0),
		Map_OpusPtr_OpusDBID: make(map[*models.Opus]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOrnaments = BackRepoOrnamentsStruct{
		Map_OrnamentsDBID_OrnamentsPtr: make(map[uint]*models.Ornaments, 0),
		Map_OrnamentsDBID_OrnamentsDB:  make(map[uint]*OrnamentsDB, 0),
		Map_OrnamentsPtr_OrnamentsDBID: make(map[*models.Ornaments]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOther_appearance = BackRepoOther_appearanceStruct{
		Map_Other_appearanceDBID_Other_appearancePtr: make(map[uint]*models.Other_appearance, 0),
		Map_Other_appearanceDBID_Other_appearanceDB:  make(map[uint]*Other_appearanceDB, 0),
		Map_Other_appearancePtr_Other_appearanceDBID: make(map[*models.Other_appearance]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOther_listening = BackRepoOther_listeningStruct{
		Map_Other_listeningDBID_Other_listeningPtr: make(map[uint]*models.Other_listening, 0),
		Map_Other_listeningDBID_Other_listeningDB:  make(map[uint]*Other_listeningDB, 0),
		Map_Other_listeningPtr_Other_listeningDBID: make(map[*models.Other_listening]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOther_notation = BackRepoOther_notationStruct{
		Map_Other_notationDBID_Other_notationPtr: make(map[uint]*models.Other_notation, 0),
		Map_Other_notationDBID_Other_notationDB:  make(map[uint]*Other_notationDB, 0),
		Map_Other_notationPtr_Other_notationDBID: make(map[*models.Other_notation]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoOther_play = BackRepoOther_playStruct{
		Map_Other_playDBID_Other_playPtr: make(map[uint]*models.Other_play, 0),
		Map_Other_playDBID_Other_playDB:  make(map[uint]*Other_playDB, 0),
		Map_Other_playPtr_Other_playDBID: make(map[*models.Other_play]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPage_layout = BackRepoPage_layoutStruct{
		Map_Page_layoutDBID_Page_layoutPtr: make(map[uint]*models.Page_layout, 0),
		Map_Page_layoutDBID_Page_layoutDB:  make(map[uint]*Page_layoutDB, 0),
		Map_Page_layoutPtr_Page_layoutDBID: make(map[*models.Page_layout]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPage_margins = BackRepoPage_marginsStruct{
		Map_Page_marginsDBID_Page_marginsPtr: make(map[uint]*models.Page_margins, 0),
		Map_Page_marginsDBID_Page_marginsDB:  make(map[uint]*Page_marginsDB, 0),
		Map_Page_marginsPtr_Page_marginsDBID: make(map[*models.Page_margins]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPart_clef = BackRepoPart_clefStruct{
		Map_Part_clefDBID_Part_clefPtr: make(map[uint]*models.Part_clef, 0),
		Map_Part_clefDBID_Part_clefDB:  make(map[uint]*Part_clefDB, 0),
		Map_Part_clefPtr_Part_clefDBID: make(map[*models.Part_clef]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPart_group = BackRepoPart_groupStruct{
		Map_Part_groupDBID_Part_groupPtr: make(map[uint]*models.Part_group, 0),
		Map_Part_groupDBID_Part_groupDB:  make(map[uint]*Part_groupDB, 0),
		Map_Part_groupPtr_Part_groupDBID: make(map[*models.Part_group]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPart_link = BackRepoPart_linkStruct{
		Map_Part_linkDBID_Part_linkPtr: make(map[uint]*models.Part_link, 0),
		Map_Part_linkDBID_Part_linkDB:  make(map[uint]*Part_linkDB, 0),
		Map_Part_linkPtr_Part_linkDBID: make(map[*models.Part_link]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPart_list = BackRepoPart_listStruct{
		Map_Part_listDBID_Part_listPtr: make(map[uint]*models.Part_list, 0),
		Map_Part_listDBID_Part_listDB:  make(map[uint]*Part_listDB, 0),
		Map_Part_listPtr_Part_listDBID: make(map[*models.Part_list]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPart_symbol = BackRepoPart_symbolStruct{
		Map_Part_symbolDBID_Part_symbolPtr: make(map[uint]*models.Part_symbol, 0),
		Map_Part_symbolDBID_Part_symbolDB:  make(map[uint]*Part_symbolDB, 0),
		Map_Part_symbolPtr_Part_symbolDBID: make(map[*models.Part_symbol]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPart_transpose = BackRepoPart_transposeStruct{
		Map_Part_transposeDBID_Part_transposePtr: make(map[uint]*models.Part_transpose, 0),
		Map_Part_transposeDBID_Part_transposeDB:  make(map[uint]*Part_transposeDB, 0),
		Map_Part_transposePtr_Part_transposeDBID: make(map[*models.Part_transpose]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPedal = BackRepoPedalStruct{
		Map_PedalDBID_PedalPtr: make(map[uint]*models.Pedal, 0),
		Map_PedalDBID_PedalDB:  make(map[uint]*PedalDB, 0),
		Map_PedalPtr_PedalDBID: make(map[*models.Pedal]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPedal_tuning = BackRepoPedal_tuningStruct{
		Map_Pedal_tuningDBID_Pedal_tuningPtr: make(map[uint]*models.Pedal_tuning, 0),
		Map_Pedal_tuningDBID_Pedal_tuningDB:  make(map[uint]*Pedal_tuningDB, 0),
		Map_Pedal_tuningPtr_Pedal_tuningDBID: make(map[*models.Pedal_tuning]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPercussion = BackRepoPercussionStruct{
		Map_PercussionDBID_PercussionPtr: make(map[uint]*models.Percussion, 0),
		Map_PercussionDBID_PercussionDB:  make(map[uint]*PercussionDB, 0),
		Map_PercussionPtr_PercussionDBID: make(map[*models.Percussion]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPitch = BackRepoPitchStruct{
		Map_PitchDBID_PitchPtr: make(map[uint]*models.Pitch, 0),
		Map_PitchDBID_PitchDB:  make(map[uint]*PitchDB, 0),
		Map_PitchPtr_PitchDBID: make(map[*models.Pitch]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPitched = BackRepoPitchedStruct{
		Map_PitchedDBID_PitchedPtr: make(map[uint]*models.Pitched, 0),
		Map_PitchedDBID_PitchedDB:  make(map[uint]*PitchedDB, 0),
		Map_PitchedPtr_PitchedDBID: make(map[*models.Pitched]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPlay = BackRepoPlayStruct{
		Map_PlayDBID_PlayPtr: make(map[uint]*models.Play, 0),
		Map_PlayDBID_PlayDB:  make(map[uint]*PlayDB, 0),
		Map_PlayPtr_PlayDBID: make(map[*models.Play]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPlayer = BackRepoPlayerStruct{
		Map_PlayerDBID_PlayerPtr: make(map[uint]*models.Player, 0),
		Map_PlayerDBID_PlayerDB:  make(map[uint]*PlayerDB, 0),
		Map_PlayerPtr_PlayerDBID: make(map[*models.Player]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPrincipal_voice = BackRepoPrincipal_voiceStruct{
		Map_Principal_voiceDBID_Principal_voicePtr: make(map[uint]*models.Principal_voice, 0),
		Map_Principal_voiceDBID_Principal_voiceDB:  make(map[uint]*Principal_voiceDB, 0),
		Map_Principal_voicePtr_Principal_voiceDBID: make(map[*models.Principal_voice]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPrint = BackRepoPrintStruct{
		Map_PrintDBID_PrintPtr: make(map[uint]*models.Print, 0),
		Map_PrintDBID_PrintDB:  make(map[uint]*PrintDB, 0),
		Map_PrintPtr_PrintDBID: make(map[*models.Print]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRelease = BackRepoReleaseStruct{
		Map_ReleaseDBID_ReleasePtr: make(map[uint]*models.Release, 0),
		Map_ReleaseDBID_ReleaseDB:  make(map[uint]*ReleaseDB, 0),
		Map_ReleasePtr_ReleaseDBID: make(map[*models.Release]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRepeat = BackRepoRepeatStruct{
		Map_RepeatDBID_RepeatPtr: make(map[uint]*models.Repeat, 0),
		Map_RepeatDBID_RepeatDB:  make(map[uint]*RepeatDB, 0),
		Map_RepeatPtr_RepeatDBID: make(map[*models.Repeat]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRest = BackRepoRestStruct{
		Map_RestDBID_RestPtr: make(map[uint]*models.Rest, 0),
		Map_RestDBID_RestDB:  make(map[uint]*RestDB, 0),
		Map_RestPtr_RestDBID: make(map[*models.Rest]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRoot = BackRepoRootStruct{
		Map_RootDBID_RootPtr: make(map[uint]*models.Root, 0),
		Map_RootDBID_RootDB:  make(map[uint]*RootDB, 0),
		Map_RootPtr_RootDBID: make(map[*models.Root]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRoot_step = BackRepoRoot_stepStruct{
		Map_Root_stepDBID_Root_stepPtr: make(map[uint]*models.Root_step, 0),
		Map_Root_stepDBID_Root_stepDB:  make(map[uint]*Root_stepDB, 0),
		Map_Root_stepPtr_Root_stepDBID: make(map[*models.Root_step]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScaling = BackRepoScalingStruct{
		Map_ScalingDBID_ScalingPtr: make(map[uint]*models.Scaling, 0),
		Map_ScalingDBID_ScalingDB:  make(map[uint]*ScalingDB, 0),
		Map_ScalingPtr_ScalingDBID: make(map[*models.Scaling]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScordatura = BackRepoScordaturaStruct{
		Map_ScordaturaDBID_ScordaturaPtr: make(map[uint]*models.Scordatura, 0),
		Map_ScordaturaDBID_ScordaturaDB:  make(map[uint]*ScordaturaDB, 0),
		Map_ScordaturaPtr_ScordaturaDBID: make(map[*models.Scordatura]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScore_instrument = BackRepoScore_instrumentStruct{
		Map_Score_instrumentDBID_Score_instrumentPtr: make(map[uint]*models.Score_instrument, 0),
		Map_Score_instrumentDBID_Score_instrumentDB:  make(map[uint]*Score_instrumentDB, 0),
		Map_Score_instrumentPtr_Score_instrumentDBID: make(map[*models.Score_instrument]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScore_part = BackRepoScore_partStruct{
		Map_Score_partDBID_Score_partPtr: make(map[uint]*models.Score_part, 0),
		Map_Score_partDBID_Score_partDB:  make(map[uint]*Score_partDB, 0),
		Map_Score_partPtr_Score_partDBID: make(map[*models.Score_part]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScore_partwise = BackRepoScore_partwiseStruct{
		Map_Score_partwiseDBID_Score_partwisePtr: make(map[uint]*models.Score_partwise, 0),
		Map_Score_partwiseDBID_Score_partwiseDB:  make(map[uint]*Score_partwiseDB, 0),
		Map_Score_partwisePtr_Score_partwiseDBID: make(map[*models.Score_partwise]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoScore_timewise = BackRepoScore_timewiseStruct{
		Map_Score_timewiseDBID_Score_timewisePtr: make(map[uint]*models.Score_timewise, 0),
		Map_Score_timewiseDBID_Score_timewiseDB:  make(map[uint]*Score_timewiseDB, 0),
		Map_Score_timewisePtr_Score_timewiseDBID: make(map[*models.Score_timewise]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSegno = BackRepoSegnoStruct{
		Map_SegnoDBID_SegnoPtr: make(map[uint]*models.Segno, 0),
		Map_SegnoDBID_SegnoDB:  make(map[uint]*SegnoDB, 0),
		Map_SegnoPtr_SegnoDBID: make(map[*models.Segno]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSlash = BackRepoSlashStruct{
		Map_SlashDBID_SlashPtr: make(map[uint]*models.Slash, 0),
		Map_SlashDBID_SlashDB:  make(map[uint]*SlashDB, 0),
		Map_SlashPtr_SlashDBID: make(map[*models.Slash]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSlide = BackRepoSlideStruct{
		Map_SlideDBID_SlidePtr: make(map[uint]*models.Slide, 0),
		Map_SlideDBID_SlideDB:  make(map[uint]*SlideDB, 0),
		Map_SlidePtr_SlideDBID: make(map[*models.Slide]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSlur = BackRepoSlurStruct{
		Map_SlurDBID_SlurPtr: make(map[uint]*models.Slur, 0),
		Map_SlurDBID_SlurDB:  make(map[uint]*SlurDB, 0),
		Map_SlurPtr_SlurDBID: make(map[*models.Slur]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSound = BackRepoSoundStruct{
		Map_SoundDBID_SoundPtr: make(map[uint]*models.Sound, 0),
		Map_SoundDBID_SoundDB:  make(map[uint]*SoundDB, 0),
		Map_SoundPtr_SoundDBID: make(map[*models.Sound]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStaff_details = BackRepoStaff_detailsStruct{
		Map_Staff_detailsDBID_Staff_detailsPtr: make(map[uint]*models.Staff_details, 0),
		Map_Staff_detailsDBID_Staff_detailsDB:  make(map[uint]*Staff_detailsDB, 0),
		Map_Staff_detailsPtr_Staff_detailsDBID: make(map[*models.Staff_details]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStaff_divide = BackRepoStaff_divideStruct{
		Map_Staff_divideDBID_Staff_dividePtr: make(map[uint]*models.Staff_divide, 0),
		Map_Staff_divideDBID_Staff_divideDB:  make(map[uint]*Staff_divideDB, 0),
		Map_Staff_dividePtr_Staff_divideDBID: make(map[*models.Staff_divide]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStaff_layout = BackRepoStaff_layoutStruct{
		Map_Staff_layoutDBID_Staff_layoutPtr: make(map[uint]*models.Staff_layout, 0),
		Map_Staff_layoutDBID_Staff_layoutDB:  make(map[uint]*Staff_layoutDB, 0),
		Map_Staff_layoutPtr_Staff_layoutDBID: make(map[*models.Staff_layout]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStaff_size = BackRepoStaff_sizeStruct{
		Map_Staff_sizeDBID_Staff_sizePtr: make(map[uint]*models.Staff_size, 0),
		Map_Staff_sizeDBID_Staff_sizeDB:  make(map[uint]*Staff_sizeDB, 0),
		Map_Staff_sizePtr_Staff_sizeDBID: make(map[*models.Staff_size]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStaff_tuning = BackRepoStaff_tuningStruct{
		Map_Staff_tuningDBID_Staff_tuningPtr: make(map[uint]*models.Staff_tuning, 0),
		Map_Staff_tuningDBID_Staff_tuningDB:  make(map[uint]*Staff_tuningDB, 0),
		Map_Staff_tuningPtr_Staff_tuningDBID: make(map[*models.Staff_tuning]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStem = BackRepoStemStruct{
		Map_StemDBID_StemPtr: make(map[uint]*models.Stem, 0),
		Map_StemDBID_StemDB:  make(map[uint]*StemDB, 0),
		Map_StemPtr_StemDBID: make(map[*models.Stem]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStick = BackRepoStickStruct{
		Map_StickDBID_StickPtr: make(map[uint]*models.Stick, 0),
		Map_StickDBID_StickDB:  make(map[uint]*StickDB, 0),
		Map_StickPtr_StickDBID: make(map[*models.Stick]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoString_mute = BackRepoString_muteStruct{
		Map_String_muteDBID_String_mutePtr: make(map[uint]*models.String_mute, 0),
		Map_String_muteDBID_String_muteDB:  make(map[uint]*String_muteDB, 0),
		Map_String_mutePtr_String_muteDBID: make(map[*models.String_mute]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoStrong_accent = BackRepoStrong_accentStruct{
		Map_Strong_accentDBID_Strong_accentPtr: make(map[uint]*models.Strong_accent, 0),
		Map_Strong_accentDBID_Strong_accentDB:  make(map[uint]*Strong_accentDB, 0),
		Map_Strong_accentPtr_Strong_accentDBID: make(map[*models.Strong_accent]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSupports = BackRepoSupportsStruct{
		Map_SupportsDBID_SupportsPtr: make(map[uint]*models.Supports, 0),
		Map_SupportsDBID_SupportsDB:  make(map[uint]*SupportsDB, 0),
		Map_SupportsPtr_SupportsDBID: make(map[*models.Supports]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSwing = BackRepoSwingStruct{
		Map_SwingDBID_SwingPtr: make(map[uint]*models.Swing, 0),
		Map_SwingDBID_SwingDB:  make(map[uint]*SwingDB, 0),
		Map_SwingPtr_SwingDBID: make(map[*models.Swing]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSync = BackRepoSyncStruct{
		Map_SyncDBID_SyncPtr: make(map[uint]*models.Sync, 0),
		Map_SyncDBID_SyncDB:  make(map[uint]*SyncDB, 0),
		Map_SyncPtr_SyncDBID: make(map[*models.Sync]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSystem_dividers = BackRepoSystem_dividersStruct{
		Map_System_dividersDBID_System_dividersPtr: make(map[uint]*models.System_dividers, 0),
		Map_System_dividersDBID_System_dividersDB:  make(map[uint]*System_dividersDB, 0),
		Map_System_dividersPtr_System_dividersDBID: make(map[*models.System_dividers]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSystem_layout = BackRepoSystem_layoutStruct{
		Map_System_layoutDBID_System_layoutPtr: make(map[uint]*models.System_layout, 0),
		Map_System_layoutDBID_System_layoutDB:  make(map[uint]*System_layoutDB, 0),
		Map_System_layoutPtr_System_layoutDBID: make(map[*models.System_layout]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSystem_margins = BackRepoSystem_marginsStruct{
		Map_System_marginsDBID_System_marginsPtr: make(map[uint]*models.System_margins, 0),
		Map_System_marginsDBID_System_marginsDB:  make(map[uint]*System_marginsDB, 0),
		Map_System_marginsPtr_System_marginsDBID: make(map[*models.System_margins]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTap = BackRepoTapStruct{
		Map_TapDBID_TapPtr: make(map[uint]*models.Tap, 0),
		Map_TapDBID_TapDB:  make(map[uint]*TapDB, 0),
		Map_TapPtr_TapDBID: make(map[*models.Tap]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTechnical = BackRepoTechnicalStruct{
		Map_TechnicalDBID_TechnicalPtr: make(map[uint]*models.Technical, 0),
		Map_TechnicalDBID_TechnicalDB:  make(map[uint]*TechnicalDB, 0),
		Map_TechnicalPtr_TechnicalDBID: make(map[*models.Technical]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoText_element_data = BackRepoText_element_dataStruct{
		Map_Text_element_dataDBID_Text_element_dataPtr: make(map[uint]*models.Text_element_data, 0),
		Map_Text_element_dataDBID_Text_element_dataDB:  make(map[uint]*Text_element_dataDB, 0),
		Map_Text_element_dataPtr_Text_element_dataDBID: make(map[*models.Text_element_data]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTie = BackRepoTieStruct{
		Map_TieDBID_TiePtr: make(map[uint]*models.Tie, 0),
		Map_TieDBID_TieDB:  make(map[uint]*TieDB, 0),
		Map_TiePtr_TieDBID: make(map[*models.Tie]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTied = BackRepoTiedStruct{
		Map_TiedDBID_TiedPtr: make(map[uint]*models.Tied, 0),
		Map_TiedDBID_TiedDB:  make(map[uint]*TiedDB, 0),
		Map_TiedPtr_TiedDBID: make(map[*models.Tied]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTime = BackRepoTimeStruct{
		Map_TimeDBID_TimePtr: make(map[uint]*models.Time, 0),
		Map_TimeDBID_TimeDB:  make(map[uint]*TimeDB, 0),
		Map_TimePtr_TimeDBID: make(map[*models.Time]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTime_modification = BackRepoTime_modificationStruct{
		Map_Time_modificationDBID_Time_modificationPtr: make(map[uint]*models.Time_modification, 0),
		Map_Time_modificationDBID_Time_modificationDB:  make(map[uint]*Time_modificationDB, 0),
		Map_Time_modificationPtr_Time_modificationDBID: make(map[*models.Time_modification]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTimpani = BackRepoTimpaniStruct{
		Map_TimpaniDBID_TimpaniPtr: make(map[uint]*models.Timpani, 0),
		Map_TimpaniDBID_TimpaniDB:  make(map[uint]*TimpaniDB, 0),
		Map_TimpaniPtr_TimpaniDBID: make(map[*models.Timpani]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTranspose = BackRepoTransposeStruct{
		Map_TransposeDBID_TransposePtr: make(map[uint]*models.Transpose, 0),
		Map_TransposeDBID_TransposeDB:  make(map[uint]*TransposeDB, 0),
		Map_TransposePtr_TransposeDBID: make(map[*models.Transpose]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTremolo = BackRepoTremoloStruct{
		Map_TremoloDBID_TremoloPtr: make(map[uint]*models.Tremolo, 0),
		Map_TremoloDBID_TremoloDB:  make(map[uint]*TremoloDB, 0),
		Map_TremoloPtr_TremoloDBID: make(map[*models.Tremolo]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTuplet = BackRepoTupletStruct{
		Map_TupletDBID_TupletPtr: make(map[uint]*models.Tuplet, 0),
		Map_TupletDBID_TupletDB:  make(map[uint]*TupletDB, 0),
		Map_TupletPtr_TupletDBID: make(map[*models.Tuplet]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTuplet_dot = BackRepoTuplet_dotStruct{
		Map_Tuplet_dotDBID_Tuplet_dotPtr: make(map[uint]*models.Tuplet_dot, 0),
		Map_Tuplet_dotDBID_Tuplet_dotDB:  make(map[uint]*Tuplet_dotDB, 0),
		Map_Tuplet_dotPtr_Tuplet_dotDBID: make(map[*models.Tuplet_dot]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTuplet_number = BackRepoTuplet_numberStruct{
		Map_Tuplet_numberDBID_Tuplet_numberPtr: make(map[uint]*models.Tuplet_number, 0),
		Map_Tuplet_numberDBID_Tuplet_numberDB:  make(map[uint]*Tuplet_numberDB, 0),
		Map_Tuplet_numberPtr_Tuplet_numberDBID: make(map[*models.Tuplet_number]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTuplet_portion = BackRepoTuplet_portionStruct{
		Map_Tuplet_portionDBID_Tuplet_portionPtr: make(map[uint]*models.Tuplet_portion, 0),
		Map_Tuplet_portionDBID_Tuplet_portionDB:  make(map[uint]*Tuplet_portionDB, 0),
		Map_Tuplet_portionPtr_Tuplet_portionDBID: make(map[*models.Tuplet_portion]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTuplet_type = BackRepoTuplet_typeStruct{
		Map_Tuplet_typeDBID_Tuplet_typePtr: make(map[uint]*models.Tuplet_type, 0),
		Map_Tuplet_typeDBID_Tuplet_typeDB:  make(map[uint]*Tuplet_typeDB, 0),
		Map_Tuplet_typePtr_Tuplet_typeDBID: make(map[*models.Tuplet_type]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoTyped_text = BackRepoTyped_textStruct{
		Map_Typed_textDBID_Typed_textPtr: make(map[uint]*models.Typed_text, 0),
		Map_Typed_textDBID_Typed_textDB:  make(map[uint]*Typed_textDB, 0),
		Map_Typed_textPtr_Typed_textDBID: make(map[*models.Typed_text]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUnpitched = BackRepoUnpitchedStruct{
		Map_UnpitchedDBID_UnpitchedPtr: make(map[uint]*models.Unpitched, 0),
		Map_UnpitchedDBID_UnpitchedDB:  make(map[uint]*UnpitchedDB, 0),
		Map_UnpitchedPtr_UnpitchedDBID: make(map[*models.Unpitched]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoVirtual_instrument = BackRepoVirtual_instrumentStruct{
		Map_Virtual_instrumentDBID_Virtual_instrumentPtr: make(map[uint]*models.Virtual_instrument, 0),
		Map_Virtual_instrumentDBID_Virtual_instrumentDB:  make(map[uint]*Virtual_instrumentDB, 0),
		Map_Virtual_instrumentPtr_Virtual_instrumentDBID: make(map[*models.Virtual_instrument]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWait = BackRepoWaitStruct{
		Map_WaitDBID_WaitPtr: make(map[uint]*models.Wait, 0),
		Map_WaitDBID_WaitDB:  make(map[uint]*WaitDB, 0),
		Map_WaitPtr_WaitDBID: make(map[*models.Wait]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWavy_line = BackRepoWavy_lineStruct{
		Map_Wavy_lineDBID_Wavy_linePtr: make(map[uint]*models.Wavy_line, 0),
		Map_Wavy_lineDBID_Wavy_lineDB:  make(map[uint]*Wavy_lineDB, 0),
		Map_Wavy_linePtr_Wavy_lineDBID: make(map[*models.Wavy_line]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWedge = BackRepoWedgeStruct{
		Map_WedgeDBID_WedgePtr: make(map[uint]*models.Wedge, 0),
		Map_WedgeDBID_WedgeDB:  make(map[uint]*WedgeDB, 0),
		Map_WedgePtr_WedgeDBID: make(map[*models.Wedge]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWood = BackRepoWoodStruct{
		Map_WoodDBID_WoodPtr: make(map[uint]*models.Wood, 0),
		Map_WoodDBID_WoodDB:  make(map[uint]*WoodDB, 0),
		Map_WoodPtr_WoodDBID: make(map[*models.Wood]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWork = BackRepoWorkStruct{
		Map_WorkDBID_WorkPtr: make(map[uint]*models.Work, 0),
		Map_WorkDBID_WorkDB:  make(map[uint]*WorkDB, 0),
		Map_WorkPtr_WorkDBID: make(map[*models.Work]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAccidental.CommitPhaseOne(stage)
	backRepo.BackRepoAccidental_mark.CommitPhaseOne(stage)
	backRepo.BackRepoAccidental_text.CommitPhaseOne(stage)
	backRepo.BackRepoAccord.CommitPhaseOne(stage)
	backRepo.BackRepoAccordion_registration.CommitPhaseOne(stage)
	backRepo.BackRepoAnyType.CommitPhaseOne(stage)
	backRepo.BackRepoAppearance.CommitPhaseOne(stage)
	backRepo.BackRepoArpeggiate.CommitPhaseOne(stage)
	backRepo.BackRepoArrow.CommitPhaseOne(stage)
	backRepo.BackRepoArticulations.CommitPhaseOne(stage)
	backRepo.BackRepoAssess.CommitPhaseOne(stage)
	backRepo.BackRepoAttributes.CommitPhaseOne(stage)
	backRepo.BackRepoBackup.CommitPhaseOne(stage)
	backRepo.BackRepoBar_style_color.CommitPhaseOne(stage)
	backRepo.BackRepoBarline.CommitPhaseOne(stage)
	backRepo.BackRepoBarre.CommitPhaseOne(stage)
	backRepo.BackRepoBass.CommitPhaseOne(stage)
	backRepo.BackRepoBass_step.CommitPhaseOne(stage)
	backRepo.BackRepoBeam.CommitPhaseOne(stage)
	backRepo.BackRepoBeat_repeat.CommitPhaseOne(stage)
	backRepo.BackRepoBeat_unit_tied.CommitPhaseOne(stage)
	backRepo.BackRepoBeater.CommitPhaseOne(stage)
	backRepo.BackRepoBend.CommitPhaseOne(stage)
	backRepo.BackRepoBookmark.CommitPhaseOne(stage)
	backRepo.BackRepoBracket.CommitPhaseOne(stage)
	backRepo.BackRepoBreath_mark.CommitPhaseOne(stage)
	backRepo.BackRepoCaesura.CommitPhaseOne(stage)
	backRepo.BackRepoCancel.CommitPhaseOne(stage)
	backRepo.BackRepoClef.CommitPhaseOne(stage)
	backRepo.BackRepoCoda.CommitPhaseOne(stage)
	backRepo.BackRepoCredit.CommitPhaseOne(stage)
	backRepo.BackRepoDashes.CommitPhaseOne(stage)
	backRepo.BackRepoDefaults.CommitPhaseOne(stage)
	backRepo.BackRepoDegree.CommitPhaseOne(stage)
	backRepo.BackRepoDegree_alter.CommitPhaseOne(stage)
	backRepo.BackRepoDegree_type.CommitPhaseOne(stage)
	backRepo.BackRepoDegree_value.CommitPhaseOne(stage)
	backRepo.BackRepoDirection.CommitPhaseOne(stage)
	backRepo.BackRepoDirection_type.CommitPhaseOne(stage)
	backRepo.BackRepoDistance.CommitPhaseOne(stage)
	backRepo.BackRepoDouble.CommitPhaseOne(stage)
	backRepo.BackRepoDynamics.CommitPhaseOne(stage)
	backRepo.BackRepoEffect.CommitPhaseOne(stage)
	backRepo.BackRepoElision.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_font.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_line.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_placement.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_placement_smufl.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_print_object_style_align.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_print_style.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_print_style_align.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_print_style_align_id.CommitPhaseOne(stage)
	backRepo.BackRepoEmpty_trill_sound.CommitPhaseOne(stage)
	backRepo.BackRepoEncoding.CommitPhaseOne(stage)
	backRepo.BackRepoEnding.CommitPhaseOne(stage)
	backRepo.BackRepoExtend.CommitPhaseOne(stage)
	backRepo.BackRepoFeature.CommitPhaseOne(stage)
	backRepo.BackRepoFermata.CommitPhaseOne(stage)
	backRepo.BackRepoFigure.CommitPhaseOne(stage)
	backRepo.BackRepoFigured_bass.CommitPhaseOne(stage)
	backRepo.BackRepoFingering.CommitPhaseOne(stage)
	backRepo.BackRepoFirst_fret.CommitPhaseOne(stage)
	backRepo.BackRepoFoo.CommitPhaseOne(stage)
	backRepo.BackRepoFor_part.CommitPhaseOne(stage)
	backRepo.BackRepoFormatted_symbol.CommitPhaseOne(stage)
	backRepo.BackRepoFormatted_symbol_id.CommitPhaseOne(stage)
	backRepo.BackRepoForward.CommitPhaseOne(stage)
	backRepo.BackRepoFrame.CommitPhaseOne(stage)
	backRepo.BackRepoFrame_note.CommitPhaseOne(stage)
	backRepo.BackRepoFret.CommitPhaseOne(stage)
	backRepo.BackRepoGlass.CommitPhaseOne(stage)
	backRepo.BackRepoGlissando.CommitPhaseOne(stage)
	backRepo.BackRepoGlyph.CommitPhaseOne(stage)
	backRepo.BackRepoGrace.CommitPhaseOne(stage)
	backRepo.BackRepoGroup_barline.CommitPhaseOne(stage)
	backRepo.BackRepoGroup_symbol.CommitPhaseOne(stage)
	backRepo.BackRepoGrouping.CommitPhaseOne(stage)
	backRepo.BackRepoHammer_on_pull_off.CommitPhaseOne(stage)
	backRepo.BackRepoHandbell.CommitPhaseOne(stage)
	backRepo.BackRepoHarmon_closed.CommitPhaseOne(stage)
	backRepo.BackRepoHarmon_mute.CommitPhaseOne(stage)
	backRepo.BackRepoHarmonic.CommitPhaseOne(stage)
	backRepo.BackRepoHarmony.CommitPhaseOne(stage)
	backRepo.BackRepoHarmony_alter.CommitPhaseOne(stage)
	backRepo.BackRepoHarp_pedals.CommitPhaseOne(stage)
	backRepo.BackRepoHeel_toe.CommitPhaseOne(stage)
	backRepo.BackRepoHole.CommitPhaseOne(stage)
	backRepo.BackRepoHole_closed.CommitPhaseOne(stage)
	backRepo.BackRepoHorizontal_turn.CommitPhaseOne(stage)
	backRepo.BackRepoIdentification.CommitPhaseOne(stage)
	backRepo.BackRepoImage.CommitPhaseOne(stage)
	backRepo.BackRepoInstrument.CommitPhaseOne(stage)
	backRepo.BackRepoInstrument_change.CommitPhaseOne(stage)
	backRepo.BackRepoInstrument_link.CommitPhaseOne(stage)
	backRepo.BackRepoInterchangeable.CommitPhaseOne(stage)
	backRepo.BackRepoInversion.CommitPhaseOne(stage)
	backRepo.BackRepoKey.CommitPhaseOne(stage)
	backRepo.BackRepoKey_accidental.CommitPhaseOne(stage)
	backRepo.BackRepoKey_octave.CommitPhaseOne(stage)
	backRepo.BackRepoKind.CommitPhaseOne(stage)
	backRepo.BackRepoLevel.CommitPhaseOne(stage)
	backRepo.BackRepoLine_detail.CommitPhaseOne(stage)
	backRepo.BackRepoLine_width.CommitPhaseOne(stage)
	backRepo.BackRepoLink.CommitPhaseOne(stage)
	backRepo.BackRepoListen.CommitPhaseOne(stage)
	backRepo.BackRepoListening.CommitPhaseOne(stage)
	backRepo.BackRepoLyric.CommitPhaseOne(stage)
	backRepo.BackRepoLyric_font.CommitPhaseOne(stage)
	backRepo.BackRepoLyric_language.CommitPhaseOne(stage)
	backRepo.BackRepoMeasure_layout.CommitPhaseOne(stage)
	backRepo.BackRepoMeasure_numbering.CommitPhaseOne(stage)
	backRepo.BackRepoMeasure_repeat.CommitPhaseOne(stage)
	backRepo.BackRepoMeasure_style.CommitPhaseOne(stage)
	backRepo.BackRepoMembrane.CommitPhaseOne(stage)
	backRepo.BackRepoMetal.CommitPhaseOne(stage)
	backRepo.BackRepoMetronome.CommitPhaseOne(stage)
	backRepo.BackRepoMetronome_beam.CommitPhaseOne(stage)
	backRepo.BackRepoMetronome_note.CommitPhaseOne(stage)
	backRepo.BackRepoMetronome_tied.CommitPhaseOne(stage)
	backRepo.BackRepoMetronome_tuplet.CommitPhaseOne(stage)
	backRepo.BackRepoMidi_device.CommitPhaseOne(stage)
	backRepo.BackRepoMidi_instrument.CommitPhaseOne(stage)
	backRepo.BackRepoMiscellaneous.CommitPhaseOne(stage)
	backRepo.BackRepoMiscellaneous_field.CommitPhaseOne(stage)
	backRepo.BackRepoMordent.CommitPhaseOne(stage)
	backRepo.BackRepoMultiple_rest.CommitPhaseOne(stage)
	backRepo.BackRepoName_display.CommitPhaseOne(stage)
	backRepo.BackRepoNon_arpeggiate.CommitPhaseOne(stage)
	backRepo.BackRepoNotations.CommitPhaseOne(stage)
	backRepo.BackRepoNote.CommitPhaseOne(stage)
	backRepo.BackRepoNote_size.CommitPhaseOne(stage)
	backRepo.BackRepoNote_type.CommitPhaseOne(stage)
	backRepo.BackRepoNotehead.CommitPhaseOne(stage)
	backRepo.BackRepoNotehead_text.CommitPhaseOne(stage)
	backRepo.BackRepoNumeral.CommitPhaseOne(stage)
	backRepo.BackRepoNumeral_key.CommitPhaseOne(stage)
	backRepo.BackRepoNumeral_root.CommitPhaseOne(stage)
	backRepo.BackRepoOctave_shift.CommitPhaseOne(stage)
	backRepo.BackRepoOffset.CommitPhaseOne(stage)
	backRepo.BackRepoOpus.CommitPhaseOne(stage)
	backRepo.BackRepoOrnaments.CommitPhaseOne(stage)
	backRepo.BackRepoOther_appearance.CommitPhaseOne(stage)
	backRepo.BackRepoOther_listening.CommitPhaseOne(stage)
	backRepo.BackRepoOther_notation.CommitPhaseOne(stage)
	backRepo.BackRepoOther_play.CommitPhaseOne(stage)
	backRepo.BackRepoPage_layout.CommitPhaseOne(stage)
	backRepo.BackRepoPage_margins.CommitPhaseOne(stage)
	backRepo.BackRepoPart_clef.CommitPhaseOne(stage)
	backRepo.BackRepoPart_group.CommitPhaseOne(stage)
	backRepo.BackRepoPart_link.CommitPhaseOne(stage)
	backRepo.BackRepoPart_list.CommitPhaseOne(stage)
	backRepo.BackRepoPart_symbol.CommitPhaseOne(stage)
	backRepo.BackRepoPart_transpose.CommitPhaseOne(stage)
	backRepo.BackRepoPedal.CommitPhaseOne(stage)
	backRepo.BackRepoPedal_tuning.CommitPhaseOne(stage)
	backRepo.BackRepoPercussion.CommitPhaseOne(stage)
	backRepo.BackRepoPitch.CommitPhaseOne(stage)
	backRepo.BackRepoPitched.CommitPhaseOne(stage)
	backRepo.BackRepoPlay.CommitPhaseOne(stage)
	backRepo.BackRepoPlayer.CommitPhaseOne(stage)
	backRepo.BackRepoPrincipal_voice.CommitPhaseOne(stage)
	backRepo.BackRepoPrint.CommitPhaseOne(stage)
	backRepo.BackRepoRelease.CommitPhaseOne(stage)
	backRepo.BackRepoRepeat.CommitPhaseOne(stage)
	backRepo.BackRepoRest.CommitPhaseOne(stage)
	backRepo.BackRepoRoot.CommitPhaseOne(stage)
	backRepo.BackRepoRoot_step.CommitPhaseOne(stage)
	backRepo.BackRepoScaling.CommitPhaseOne(stage)
	backRepo.BackRepoScordatura.CommitPhaseOne(stage)
	backRepo.BackRepoScore_instrument.CommitPhaseOne(stage)
	backRepo.BackRepoScore_part.CommitPhaseOne(stage)
	backRepo.BackRepoScore_partwise.CommitPhaseOne(stage)
	backRepo.BackRepoScore_timewise.CommitPhaseOne(stage)
	backRepo.BackRepoSegno.CommitPhaseOne(stage)
	backRepo.BackRepoSlash.CommitPhaseOne(stage)
	backRepo.BackRepoSlide.CommitPhaseOne(stage)
	backRepo.BackRepoSlur.CommitPhaseOne(stage)
	backRepo.BackRepoSound.CommitPhaseOne(stage)
	backRepo.BackRepoStaff_details.CommitPhaseOne(stage)
	backRepo.BackRepoStaff_divide.CommitPhaseOne(stage)
	backRepo.BackRepoStaff_layout.CommitPhaseOne(stage)
	backRepo.BackRepoStaff_size.CommitPhaseOne(stage)
	backRepo.BackRepoStaff_tuning.CommitPhaseOne(stage)
	backRepo.BackRepoStem.CommitPhaseOne(stage)
	backRepo.BackRepoStick.CommitPhaseOne(stage)
	backRepo.BackRepoString_mute.CommitPhaseOne(stage)
	backRepo.BackRepoStrong_accent.CommitPhaseOne(stage)
	backRepo.BackRepoSupports.CommitPhaseOne(stage)
	backRepo.BackRepoSwing.CommitPhaseOne(stage)
	backRepo.BackRepoSync.CommitPhaseOne(stage)
	backRepo.BackRepoSystem_dividers.CommitPhaseOne(stage)
	backRepo.BackRepoSystem_layout.CommitPhaseOne(stage)
	backRepo.BackRepoSystem_margins.CommitPhaseOne(stage)
	backRepo.BackRepoTap.CommitPhaseOne(stage)
	backRepo.BackRepoTechnical.CommitPhaseOne(stage)
	backRepo.BackRepoText_element_data.CommitPhaseOne(stage)
	backRepo.BackRepoTie.CommitPhaseOne(stage)
	backRepo.BackRepoTied.CommitPhaseOne(stage)
	backRepo.BackRepoTime.CommitPhaseOne(stage)
	backRepo.BackRepoTime_modification.CommitPhaseOne(stage)
	backRepo.BackRepoTimpani.CommitPhaseOne(stage)
	backRepo.BackRepoTranspose.CommitPhaseOne(stage)
	backRepo.BackRepoTremolo.CommitPhaseOne(stage)
	backRepo.BackRepoTuplet.CommitPhaseOne(stage)
	backRepo.BackRepoTuplet_dot.CommitPhaseOne(stage)
	backRepo.BackRepoTuplet_number.CommitPhaseOne(stage)
	backRepo.BackRepoTuplet_portion.CommitPhaseOne(stage)
	backRepo.BackRepoTuplet_type.CommitPhaseOne(stage)
	backRepo.BackRepoTyped_text.CommitPhaseOne(stage)
	backRepo.BackRepoUnpitched.CommitPhaseOne(stage)
	backRepo.BackRepoVirtual_instrument.CommitPhaseOne(stage)
	backRepo.BackRepoWait.CommitPhaseOne(stage)
	backRepo.BackRepoWavy_line.CommitPhaseOne(stage)
	backRepo.BackRepoWedge.CommitPhaseOne(stage)
	backRepo.BackRepoWood.CommitPhaseOne(stage)
	backRepo.BackRepoWork.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAccidental.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAccidental_mark.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAccidental_text.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAccord.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAccordion_registration.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAnyType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAppearance.CommitPhaseTwo(backRepo)
	backRepo.BackRepoArpeggiate.CommitPhaseTwo(backRepo)
	backRepo.BackRepoArrow.CommitPhaseTwo(backRepo)
	backRepo.BackRepoArticulations.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAssess.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAttributes.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBackup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBar_style_color.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBarline.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBarre.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBass.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBass_step.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBeam.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBeat_repeat.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBeat_unit_tied.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBeater.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBend.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBookmark.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBracket.CommitPhaseTwo(backRepo)
	backRepo.BackRepoBreath_mark.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCaesura.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCancel.CommitPhaseTwo(backRepo)
	backRepo.BackRepoClef.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCoda.CommitPhaseTwo(backRepo)
	backRepo.BackRepoCredit.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDashes.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDefaults.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDegree.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDegree_alter.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDegree_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDegree_value.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDirection.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDirection_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDistance.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDouble.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDynamics.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEffect.CommitPhaseTwo(backRepo)
	backRepo.BackRepoElision.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_font.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_line.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_placement.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_placement_smufl.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_object_style_align.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_style.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_style_align.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_style_align_id.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_trill_sound.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEncoding.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEnding.CommitPhaseTwo(backRepo)
	backRepo.BackRepoExtend.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFeature.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFermata.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFigure.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFigured_bass.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFingering.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFirst_fret.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFoo.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFor_part.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormatted_symbol.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFormatted_symbol_id.CommitPhaseTwo(backRepo)
	backRepo.BackRepoForward.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFrame.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFrame_note.CommitPhaseTwo(backRepo)
	backRepo.BackRepoFret.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGlass.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGlissando.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGlyph.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGrace.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGroup_barline.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGroup_symbol.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGrouping.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHammer_on_pull_off.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHandbell.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHarmon_closed.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHarmon_mute.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHarmonic.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHarmony.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHarmony_alter.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHarp_pedals.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHeel_toe.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHole.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHole_closed.CommitPhaseTwo(backRepo)
	backRepo.BackRepoHorizontal_turn.CommitPhaseTwo(backRepo)
	backRepo.BackRepoIdentification.CommitPhaseTwo(backRepo)
	backRepo.BackRepoImage.CommitPhaseTwo(backRepo)
	backRepo.BackRepoInstrument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoInstrument_change.CommitPhaseTwo(backRepo)
	backRepo.BackRepoInstrument_link.CommitPhaseTwo(backRepo)
	backRepo.BackRepoInterchangeable.CommitPhaseTwo(backRepo)
	backRepo.BackRepoInversion.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKey.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKey_accidental.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKey_octave.CommitPhaseTwo(backRepo)
	backRepo.BackRepoKind.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLevel.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLine_detail.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLine_width.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLink.CommitPhaseTwo(backRepo)
	backRepo.BackRepoListen.CommitPhaseTwo(backRepo)
	backRepo.BackRepoListening.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLyric.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLyric_font.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLyric_language.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_layout.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_numbering.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_repeat.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_style.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMembrane.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetal.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetronome.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_beam.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_note.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_tied.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_tuplet.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMidi_device.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMidi_instrument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMiscellaneous.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMiscellaneous_field.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMordent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMultiple_rest.CommitPhaseTwo(backRepo)
	backRepo.BackRepoName_display.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNon_arpeggiate.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNotations.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNote.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNote_size.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNote_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNotehead.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNotehead_text.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNumeral.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNumeral_key.CommitPhaseTwo(backRepo)
	backRepo.BackRepoNumeral_root.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOctave_shift.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOffset.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOpus.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOrnaments.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOther_appearance.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOther_listening.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOther_notation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoOther_play.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPage_layout.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPage_margins.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPart_clef.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPart_group.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPart_link.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPart_list.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPart_symbol.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPart_transpose.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPedal.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPedal_tuning.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPercussion.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPitch.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPitched.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPlay.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPlayer.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPrincipal_voice.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPrint.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRelease.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRepeat.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRest.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRoot.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRoot_step.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScaling.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScordatura.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScore_instrument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScore_part.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScore_partwise.CommitPhaseTwo(backRepo)
	backRepo.BackRepoScore_timewise.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSegno.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSlash.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSlide.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSlur.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSound.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStaff_details.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStaff_divide.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStaff_layout.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStaff_size.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStaff_tuning.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStem.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStick.CommitPhaseTwo(backRepo)
	backRepo.BackRepoString_mute.CommitPhaseTwo(backRepo)
	backRepo.BackRepoStrong_accent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSupports.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSwing.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSync.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSystem_dividers.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSystem_layout.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSystem_margins.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTap.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTechnical.CommitPhaseTwo(backRepo)
	backRepo.BackRepoText_element_data.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTie.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTied.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTime.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTime_modification.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTimpani.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTranspose.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTremolo.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTuplet.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_dot.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_number.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_portion.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_type.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTyped_text.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUnpitched.CommitPhaseTwo(backRepo)
	backRepo.BackRepoVirtual_instrument.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWait.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWavy_line.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWedge.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWood.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWork.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAccidental.CheckoutPhaseOne()
	backRepo.BackRepoAccidental_mark.CheckoutPhaseOne()
	backRepo.BackRepoAccidental_text.CheckoutPhaseOne()
	backRepo.BackRepoAccord.CheckoutPhaseOne()
	backRepo.BackRepoAccordion_registration.CheckoutPhaseOne()
	backRepo.BackRepoAnyType.CheckoutPhaseOne()
	backRepo.BackRepoAppearance.CheckoutPhaseOne()
	backRepo.BackRepoArpeggiate.CheckoutPhaseOne()
	backRepo.BackRepoArrow.CheckoutPhaseOne()
	backRepo.BackRepoArticulations.CheckoutPhaseOne()
	backRepo.BackRepoAssess.CheckoutPhaseOne()
	backRepo.BackRepoAttributes.CheckoutPhaseOne()
	backRepo.BackRepoBackup.CheckoutPhaseOne()
	backRepo.BackRepoBar_style_color.CheckoutPhaseOne()
	backRepo.BackRepoBarline.CheckoutPhaseOne()
	backRepo.BackRepoBarre.CheckoutPhaseOne()
	backRepo.BackRepoBass.CheckoutPhaseOne()
	backRepo.BackRepoBass_step.CheckoutPhaseOne()
	backRepo.BackRepoBeam.CheckoutPhaseOne()
	backRepo.BackRepoBeat_repeat.CheckoutPhaseOne()
	backRepo.BackRepoBeat_unit_tied.CheckoutPhaseOne()
	backRepo.BackRepoBeater.CheckoutPhaseOne()
	backRepo.BackRepoBend.CheckoutPhaseOne()
	backRepo.BackRepoBookmark.CheckoutPhaseOne()
	backRepo.BackRepoBracket.CheckoutPhaseOne()
	backRepo.BackRepoBreath_mark.CheckoutPhaseOne()
	backRepo.BackRepoCaesura.CheckoutPhaseOne()
	backRepo.BackRepoCancel.CheckoutPhaseOne()
	backRepo.BackRepoClef.CheckoutPhaseOne()
	backRepo.BackRepoCoda.CheckoutPhaseOne()
	backRepo.BackRepoCredit.CheckoutPhaseOne()
	backRepo.BackRepoDashes.CheckoutPhaseOne()
	backRepo.BackRepoDefaults.CheckoutPhaseOne()
	backRepo.BackRepoDegree.CheckoutPhaseOne()
	backRepo.BackRepoDegree_alter.CheckoutPhaseOne()
	backRepo.BackRepoDegree_type.CheckoutPhaseOne()
	backRepo.BackRepoDegree_value.CheckoutPhaseOne()
	backRepo.BackRepoDirection.CheckoutPhaseOne()
	backRepo.BackRepoDirection_type.CheckoutPhaseOne()
	backRepo.BackRepoDistance.CheckoutPhaseOne()
	backRepo.BackRepoDouble.CheckoutPhaseOne()
	backRepo.BackRepoDynamics.CheckoutPhaseOne()
	backRepo.BackRepoEffect.CheckoutPhaseOne()
	backRepo.BackRepoElision.CheckoutPhaseOne()
	backRepo.BackRepoEmpty.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_font.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_line.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_placement.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_placement_smufl.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_print_object_style_align.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_print_style.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_print_style_align.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_print_style_align_id.CheckoutPhaseOne()
	backRepo.BackRepoEmpty_trill_sound.CheckoutPhaseOne()
	backRepo.BackRepoEncoding.CheckoutPhaseOne()
	backRepo.BackRepoEnding.CheckoutPhaseOne()
	backRepo.BackRepoExtend.CheckoutPhaseOne()
	backRepo.BackRepoFeature.CheckoutPhaseOne()
	backRepo.BackRepoFermata.CheckoutPhaseOne()
	backRepo.BackRepoFigure.CheckoutPhaseOne()
	backRepo.BackRepoFigured_bass.CheckoutPhaseOne()
	backRepo.BackRepoFingering.CheckoutPhaseOne()
	backRepo.BackRepoFirst_fret.CheckoutPhaseOne()
	backRepo.BackRepoFoo.CheckoutPhaseOne()
	backRepo.BackRepoFor_part.CheckoutPhaseOne()
	backRepo.BackRepoFormatted_symbol.CheckoutPhaseOne()
	backRepo.BackRepoFormatted_symbol_id.CheckoutPhaseOne()
	backRepo.BackRepoForward.CheckoutPhaseOne()
	backRepo.BackRepoFrame.CheckoutPhaseOne()
	backRepo.BackRepoFrame_note.CheckoutPhaseOne()
	backRepo.BackRepoFret.CheckoutPhaseOne()
	backRepo.BackRepoGlass.CheckoutPhaseOne()
	backRepo.BackRepoGlissando.CheckoutPhaseOne()
	backRepo.BackRepoGlyph.CheckoutPhaseOne()
	backRepo.BackRepoGrace.CheckoutPhaseOne()
	backRepo.BackRepoGroup_barline.CheckoutPhaseOne()
	backRepo.BackRepoGroup_symbol.CheckoutPhaseOne()
	backRepo.BackRepoGrouping.CheckoutPhaseOne()
	backRepo.BackRepoHammer_on_pull_off.CheckoutPhaseOne()
	backRepo.BackRepoHandbell.CheckoutPhaseOne()
	backRepo.BackRepoHarmon_closed.CheckoutPhaseOne()
	backRepo.BackRepoHarmon_mute.CheckoutPhaseOne()
	backRepo.BackRepoHarmonic.CheckoutPhaseOne()
	backRepo.BackRepoHarmony.CheckoutPhaseOne()
	backRepo.BackRepoHarmony_alter.CheckoutPhaseOne()
	backRepo.BackRepoHarp_pedals.CheckoutPhaseOne()
	backRepo.BackRepoHeel_toe.CheckoutPhaseOne()
	backRepo.BackRepoHole.CheckoutPhaseOne()
	backRepo.BackRepoHole_closed.CheckoutPhaseOne()
	backRepo.BackRepoHorizontal_turn.CheckoutPhaseOne()
	backRepo.BackRepoIdentification.CheckoutPhaseOne()
	backRepo.BackRepoImage.CheckoutPhaseOne()
	backRepo.BackRepoInstrument.CheckoutPhaseOne()
	backRepo.BackRepoInstrument_change.CheckoutPhaseOne()
	backRepo.BackRepoInstrument_link.CheckoutPhaseOne()
	backRepo.BackRepoInterchangeable.CheckoutPhaseOne()
	backRepo.BackRepoInversion.CheckoutPhaseOne()
	backRepo.BackRepoKey.CheckoutPhaseOne()
	backRepo.BackRepoKey_accidental.CheckoutPhaseOne()
	backRepo.BackRepoKey_octave.CheckoutPhaseOne()
	backRepo.BackRepoKind.CheckoutPhaseOne()
	backRepo.BackRepoLevel.CheckoutPhaseOne()
	backRepo.BackRepoLine_detail.CheckoutPhaseOne()
	backRepo.BackRepoLine_width.CheckoutPhaseOne()
	backRepo.BackRepoLink.CheckoutPhaseOne()
	backRepo.BackRepoListen.CheckoutPhaseOne()
	backRepo.BackRepoListening.CheckoutPhaseOne()
	backRepo.BackRepoLyric.CheckoutPhaseOne()
	backRepo.BackRepoLyric_font.CheckoutPhaseOne()
	backRepo.BackRepoLyric_language.CheckoutPhaseOne()
	backRepo.BackRepoMeasure_layout.CheckoutPhaseOne()
	backRepo.BackRepoMeasure_numbering.CheckoutPhaseOne()
	backRepo.BackRepoMeasure_repeat.CheckoutPhaseOne()
	backRepo.BackRepoMeasure_style.CheckoutPhaseOne()
	backRepo.BackRepoMembrane.CheckoutPhaseOne()
	backRepo.BackRepoMetal.CheckoutPhaseOne()
	backRepo.BackRepoMetronome.CheckoutPhaseOne()
	backRepo.BackRepoMetronome_beam.CheckoutPhaseOne()
	backRepo.BackRepoMetronome_note.CheckoutPhaseOne()
	backRepo.BackRepoMetronome_tied.CheckoutPhaseOne()
	backRepo.BackRepoMetronome_tuplet.CheckoutPhaseOne()
	backRepo.BackRepoMidi_device.CheckoutPhaseOne()
	backRepo.BackRepoMidi_instrument.CheckoutPhaseOne()
	backRepo.BackRepoMiscellaneous.CheckoutPhaseOne()
	backRepo.BackRepoMiscellaneous_field.CheckoutPhaseOne()
	backRepo.BackRepoMordent.CheckoutPhaseOne()
	backRepo.BackRepoMultiple_rest.CheckoutPhaseOne()
	backRepo.BackRepoName_display.CheckoutPhaseOne()
	backRepo.BackRepoNon_arpeggiate.CheckoutPhaseOne()
	backRepo.BackRepoNotations.CheckoutPhaseOne()
	backRepo.BackRepoNote.CheckoutPhaseOne()
	backRepo.BackRepoNote_size.CheckoutPhaseOne()
	backRepo.BackRepoNote_type.CheckoutPhaseOne()
	backRepo.BackRepoNotehead.CheckoutPhaseOne()
	backRepo.BackRepoNotehead_text.CheckoutPhaseOne()
	backRepo.BackRepoNumeral.CheckoutPhaseOne()
	backRepo.BackRepoNumeral_key.CheckoutPhaseOne()
	backRepo.BackRepoNumeral_root.CheckoutPhaseOne()
	backRepo.BackRepoOctave_shift.CheckoutPhaseOne()
	backRepo.BackRepoOffset.CheckoutPhaseOne()
	backRepo.BackRepoOpus.CheckoutPhaseOne()
	backRepo.BackRepoOrnaments.CheckoutPhaseOne()
	backRepo.BackRepoOther_appearance.CheckoutPhaseOne()
	backRepo.BackRepoOther_listening.CheckoutPhaseOne()
	backRepo.BackRepoOther_notation.CheckoutPhaseOne()
	backRepo.BackRepoOther_play.CheckoutPhaseOne()
	backRepo.BackRepoPage_layout.CheckoutPhaseOne()
	backRepo.BackRepoPage_margins.CheckoutPhaseOne()
	backRepo.BackRepoPart_clef.CheckoutPhaseOne()
	backRepo.BackRepoPart_group.CheckoutPhaseOne()
	backRepo.BackRepoPart_link.CheckoutPhaseOne()
	backRepo.BackRepoPart_list.CheckoutPhaseOne()
	backRepo.BackRepoPart_symbol.CheckoutPhaseOne()
	backRepo.BackRepoPart_transpose.CheckoutPhaseOne()
	backRepo.BackRepoPedal.CheckoutPhaseOne()
	backRepo.BackRepoPedal_tuning.CheckoutPhaseOne()
	backRepo.BackRepoPercussion.CheckoutPhaseOne()
	backRepo.BackRepoPitch.CheckoutPhaseOne()
	backRepo.BackRepoPitched.CheckoutPhaseOne()
	backRepo.BackRepoPlay.CheckoutPhaseOne()
	backRepo.BackRepoPlayer.CheckoutPhaseOne()
	backRepo.BackRepoPrincipal_voice.CheckoutPhaseOne()
	backRepo.BackRepoPrint.CheckoutPhaseOne()
	backRepo.BackRepoRelease.CheckoutPhaseOne()
	backRepo.BackRepoRepeat.CheckoutPhaseOne()
	backRepo.BackRepoRest.CheckoutPhaseOne()
	backRepo.BackRepoRoot.CheckoutPhaseOne()
	backRepo.BackRepoRoot_step.CheckoutPhaseOne()
	backRepo.BackRepoScaling.CheckoutPhaseOne()
	backRepo.BackRepoScordatura.CheckoutPhaseOne()
	backRepo.BackRepoScore_instrument.CheckoutPhaseOne()
	backRepo.BackRepoScore_part.CheckoutPhaseOne()
	backRepo.BackRepoScore_partwise.CheckoutPhaseOne()
	backRepo.BackRepoScore_timewise.CheckoutPhaseOne()
	backRepo.BackRepoSegno.CheckoutPhaseOne()
	backRepo.BackRepoSlash.CheckoutPhaseOne()
	backRepo.BackRepoSlide.CheckoutPhaseOne()
	backRepo.BackRepoSlur.CheckoutPhaseOne()
	backRepo.BackRepoSound.CheckoutPhaseOne()
	backRepo.BackRepoStaff_details.CheckoutPhaseOne()
	backRepo.BackRepoStaff_divide.CheckoutPhaseOne()
	backRepo.BackRepoStaff_layout.CheckoutPhaseOne()
	backRepo.BackRepoStaff_size.CheckoutPhaseOne()
	backRepo.BackRepoStaff_tuning.CheckoutPhaseOne()
	backRepo.BackRepoStem.CheckoutPhaseOne()
	backRepo.BackRepoStick.CheckoutPhaseOne()
	backRepo.BackRepoString_mute.CheckoutPhaseOne()
	backRepo.BackRepoStrong_accent.CheckoutPhaseOne()
	backRepo.BackRepoSupports.CheckoutPhaseOne()
	backRepo.BackRepoSwing.CheckoutPhaseOne()
	backRepo.BackRepoSync.CheckoutPhaseOne()
	backRepo.BackRepoSystem_dividers.CheckoutPhaseOne()
	backRepo.BackRepoSystem_layout.CheckoutPhaseOne()
	backRepo.BackRepoSystem_margins.CheckoutPhaseOne()
	backRepo.BackRepoTap.CheckoutPhaseOne()
	backRepo.BackRepoTechnical.CheckoutPhaseOne()
	backRepo.BackRepoText_element_data.CheckoutPhaseOne()
	backRepo.BackRepoTie.CheckoutPhaseOne()
	backRepo.BackRepoTied.CheckoutPhaseOne()
	backRepo.BackRepoTime.CheckoutPhaseOne()
	backRepo.BackRepoTime_modification.CheckoutPhaseOne()
	backRepo.BackRepoTimpani.CheckoutPhaseOne()
	backRepo.BackRepoTranspose.CheckoutPhaseOne()
	backRepo.BackRepoTremolo.CheckoutPhaseOne()
	backRepo.BackRepoTuplet.CheckoutPhaseOne()
	backRepo.BackRepoTuplet_dot.CheckoutPhaseOne()
	backRepo.BackRepoTuplet_number.CheckoutPhaseOne()
	backRepo.BackRepoTuplet_portion.CheckoutPhaseOne()
	backRepo.BackRepoTuplet_type.CheckoutPhaseOne()
	backRepo.BackRepoTyped_text.CheckoutPhaseOne()
	backRepo.BackRepoUnpitched.CheckoutPhaseOne()
	backRepo.BackRepoVirtual_instrument.CheckoutPhaseOne()
	backRepo.BackRepoWait.CheckoutPhaseOne()
	backRepo.BackRepoWavy_line.CheckoutPhaseOne()
	backRepo.BackRepoWedge.CheckoutPhaseOne()
	backRepo.BackRepoWood.CheckoutPhaseOne()
	backRepo.BackRepoWork.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAccidental.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAccidental_mark.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAccidental_text.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAccord.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAccordion_registration.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAnyType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAppearance.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoArpeggiate.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoArrow.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoArticulations.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAssess.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAttributes.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBackup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBar_style_color.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBarline.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBarre.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBass.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBass_step.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBeam.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBeat_repeat.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBeat_unit_tied.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBeater.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBend.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBookmark.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBracket.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoBreath_mark.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCaesura.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCancel.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoClef.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCoda.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoCredit.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDashes.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDefaults.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDegree.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDegree_alter.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDegree_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDegree_value.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDirection.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDirection_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDistance.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDouble.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDynamics.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEffect.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoElision.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_font.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_line.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_placement.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_placement_smufl.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_object_style_align.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_style.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_style_align.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_print_style_align_id.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEmpty_trill_sound.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEncoding.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEnding.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoExtend.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFeature.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFermata.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFigure.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFigured_bass.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFingering.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFirst_fret.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFoo.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFor_part.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormatted_symbol.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFormatted_symbol_id.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoForward.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFrame.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFrame_note.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoFret.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGlass.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGlissando.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGlyph.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGrace.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGroup_barline.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGroup_symbol.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGrouping.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHammer_on_pull_off.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHandbell.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHarmon_closed.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHarmon_mute.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHarmonic.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHarmony.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHarmony_alter.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHarp_pedals.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHeel_toe.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHole.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHole_closed.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoHorizontal_turn.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoIdentification.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoImage.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoInstrument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoInstrument_change.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoInstrument_link.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoInterchangeable.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoInversion.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKey.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKey_accidental.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKey_octave.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoKind.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLevel.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLine_detail.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLine_width.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLink.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoListen.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoListening.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLyric.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLyric_font.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLyric_language.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_layout.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_numbering.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_repeat.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMeasure_style.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMembrane.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetal.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetronome.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_beam.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_note.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_tied.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMetronome_tuplet.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMidi_device.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMidi_instrument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMiscellaneous.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMiscellaneous_field.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMordent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMultiple_rest.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoName_display.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNon_arpeggiate.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNotations.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNote.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNote_size.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNote_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNotehead.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNotehead_text.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNumeral.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNumeral_key.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoNumeral_root.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOctave_shift.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOffset.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOpus.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOrnaments.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOther_appearance.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOther_listening.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOther_notation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoOther_play.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPage_layout.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPage_margins.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPart_clef.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPart_group.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPart_link.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPart_list.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPart_symbol.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPart_transpose.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPedal.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPedal_tuning.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPercussion.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPitch.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPitched.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPlay.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPlayer.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPrincipal_voice.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPrint.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRelease.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRepeat.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRest.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRoot.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRoot_step.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScaling.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScordatura.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScore_instrument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScore_part.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScore_partwise.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoScore_timewise.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSegno.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSlash.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSlide.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSlur.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSound.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStaff_details.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStaff_divide.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStaff_layout.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStaff_size.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStaff_tuning.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStem.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStick.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoString_mute.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoStrong_accent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSupports.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSwing.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSync.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSystem_dividers.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSystem_layout.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSystem_margins.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTap.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTechnical.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoText_element_data.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTie.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTied.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTime.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTime_modification.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTimpani.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTranspose.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTremolo.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTuplet.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_dot.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_number.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_portion.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTuplet_type.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTyped_text.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUnpitched.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoVirtual_instrument.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWait.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWavy_line.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWedge.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWood.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWork.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAccidental.Backup(dirPath)
	backRepo.BackRepoAccidental_mark.Backup(dirPath)
	backRepo.BackRepoAccidental_text.Backup(dirPath)
	backRepo.BackRepoAccord.Backup(dirPath)
	backRepo.BackRepoAccordion_registration.Backup(dirPath)
	backRepo.BackRepoAnyType.Backup(dirPath)
	backRepo.BackRepoAppearance.Backup(dirPath)
	backRepo.BackRepoArpeggiate.Backup(dirPath)
	backRepo.BackRepoArrow.Backup(dirPath)
	backRepo.BackRepoArticulations.Backup(dirPath)
	backRepo.BackRepoAssess.Backup(dirPath)
	backRepo.BackRepoAttributes.Backup(dirPath)
	backRepo.BackRepoBackup.Backup(dirPath)
	backRepo.BackRepoBar_style_color.Backup(dirPath)
	backRepo.BackRepoBarline.Backup(dirPath)
	backRepo.BackRepoBarre.Backup(dirPath)
	backRepo.BackRepoBass.Backup(dirPath)
	backRepo.BackRepoBass_step.Backup(dirPath)
	backRepo.BackRepoBeam.Backup(dirPath)
	backRepo.BackRepoBeat_repeat.Backup(dirPath)
	backRepo.BackRepoBeat_unit_tied.Backup(dirPath)
	backRepo.BackRepoBeater.Backup(dirPath)
	backRepo.BackRepoBend.Backup(dirPath)
	backRepo.BackRepoBookmark.Backup(dirPath)
	backRepo.BackRepoBracket.Backup(dirPath)
	backRepo.BackRepoBreath_mark.Backup(dirPath)
	backRepo.BackRepoCaesura.Backup(dirPath)
	backRepo.BackRepoCancel.Backup(dirPath)
	backRepo.BackRepoClef.Backup(dirPath)
	backRepo.BackRepoCoda.Backup(dirPath)
	backRepo.BackRepoCredit.Backup(dirPath)
	backRepo.BackRepoDashes.Backup(dirPath)
	backRepo.BackRepoDefaults.Backup(dirPath)
	backRepo.BackRepoDegree.Backup(dirPath)
	backRepo.BackRepoDegree_alter.Backup(dirPath)
	backRepo.BackRepoDegree_type.Backup(dirPath)
	backRepo.BackRepoDegree_value.Backup(dirPath)
	backRepo.BackRepoDirection.Backup(dirPath)
	backRepo.BackRepoDirection_type.Backup(dirPath)
	backRepo.BackRepoDistance.Backup(dirPath)
	backRepo.BackRepoDouble.Backup(dirPath)
	backRepo.BackRepoDynamics.Backup(dirPath)
	backRepo.BackRepoEffect.Backup(dirPath)
	backRepo.BackRepoElision.Backup(dirPath)
	backRepo.BackRepoEmpty.Backup(dirPath)
	backRepo.BackRepoEmpty_font.Backup(dirPath)
	backRepo.BackRepoEmpty_line.Backup(dirPath)
	backRepo.BackRepoEmpty_placement.Backup(dirPath)
	backRepo.BackRepoEmpty_placement_smufl.Backup(dirPath)
	backRepo.BackRepoEmpty_print_object_style_align.Backup(dirPath)
	backRepo.BackRepoEmpty_print_style.Backup(dirPath)
	backRepo.BackRepoEmpty_print_style_align.Backup(dirPath)
	backRepo.BackRepoEmpty_print_style_align_id.Backup(dirPath)
	backRepo.BackRepoEmpty_trill_sound.Backup(dirPath)
	backRepo.BackRepoEncoding.Backup(dirPath)
	backRepo.BackRepoEnding.Backup(dirPath)
	backRepo.BackRepoExtend.Backup(dirPath)
	backRepo.BackRepoFeature.Backup(dirPath)
	backRepo.BackRepoFermata.Backup(dirPath)
	backRepo.BackRepoFigure.Backup(dirPath)
	backRepo.BackRepoFigured_bass.Backup(dirPath)
	backRepo.BackRepoFingering.Backup(dirPath)
	backRepo.BackRepoFirst_fret.Backup(dirPath)
	backRepo.BackRepoFoo.Backup(dirPath)
	backRepo.BackRepoFor_part.Backup(dirPath)
	backRepo.BackRepoFormatted_symbol.Backup(dirPath)
	backRepo.BackRepoFormatted_symbol_id.Backup(dirPath)
	backRepo.BackRepoForward.Backup(dirPath)
	backRepo.BackRepoFrame.Backup(dirPath)
	backRepo.BackRepoFrame_note.Backup(dirPath)
	backRepo.BackRepoFret.Backup(dirPath)
	backRepo.BackRepoGlass.Backup(dirPath)
	backRepo.BackRepoGlissando.Backup(dirPath)
	backRepo.BackRepoGlyph.Backup(dirPath)
	backRepo.BackRepoGrace.Backup(dirPath)
	backRepo.BackRepoGroup_barline.Backup(dirPath)
	backRepo.BackRepoGroup_symbol.Backup(dirPath)
	backRepo.BackRepoGrouping.Backup(dirPath)
	backRepo.BackRepoHammer_on_pull_off.Backup(dirPath)
	backRepo.BackRepoHandbell.Backup(dirPath)
	backRepo.BackRepoHarmon_closed.Backup(dirPath)
	backRepo.BackRepoHarmon_mute.Backup(dirPath)
	backRepo.BackRepoHarmonic.Backup(dirPath)
	backRepo.BackRepoHarmony.Backup(dirPath)
	backRepo.BackRepoHarmony_alter.Backup(dirPath)
	backRepo.BackRepoHarp_pedals.Backup(dirPath)
	backRepo.BackRepoHeel_toe.Backup(dirPath)
	backRepo.BackRepoHole.Backup(dirPath)
	backRepo.BackRepoHole_closed.Backup(dirPath)
	backRepo.BackRepoHorizontal_turn.Backup(dirPath)
	backRepo.BackRepoIdentification.Backup(dirPath)
	backRepo.BackRepoImage.Backup(dirPath)
	backRepo.BackRepoInstrument.Backup(dirPath)
	backRepo.BackRepoInstrument_change.Backup(dirPath)
	backRepo.BackRepoInstrument_link.Backup(dirPath)
	backRepo.BackRepoInterchangeable.Backup(dirPath)
	backRepo.BackRepoInversion.Backup(dirPath)
	backRepo.BackRepoKey.Backup(dirPath)
	backRepo.BackRepoKey_accidental.Backup(dirPath)
	backRepo.BackRepoKey_octave.Backup(dirPath)
	backRepo.BackRepoKind.Backup(dirPath)
	backRepo.BackRepoLevel.Backup(dirPath)
	backRepo.BackRepoLine_detail.Backup(dirPath)
	backRepo.BackRepoLine_width.Backup(dirPath)
	backRepo.BackRepoLink.Backup(dirPath)
	backRepo.BackRepoListen.Backup(dirPath)
	backRepo.BackRepoListening.Backup(dirPath)
	backRepo.BackRepoLyric.Backup(dirPath)
	backRepo.BackRepoLyric_font.Backup(dirPath)
	backRepo.BackRepoLyric_language.Backup(dirPath)
	backRepo.BackRepoMeasure_layout.Backup(dirPath)
	backRepo.BackRepoMeasure_numbering.Backup(dirPath)
	backRepo.BackRepoMeasure_repeat.Backup(dirPath)
	backRepo.BackRepoMeasure_style.Backup(dirPath)
	backRepo.BackRepoMembrane.Backup(dirPath)
	backRepo.BackRepoMetal.Backup(dirPath)
	backRepo.BackRepoMetronome.Backup(dirPath)
	backRepo.BackRepoMetronome_beam.Backup(dirPath)
	backRepo.BackRepoMetronome_note.Backup(dirPath)
	backRepo.BackRepoMetronome_tied.Backup(dirPath)
	backRepo.BackRepoMetronome_tuplet.Backup(dirPath)
	backRepo.BackRepoMidi_device.Backup(dirPath)
	backRepo.BackRepoMidi_instrument.Backup(dirPath)
	backRepo.BackRepoMiscellaneous.Backup(dirPath)
	backRepo.BackRepoMiscellaneous_field.Backup(dirPath)
	backRepo.BackRepoMordent.Backup(dirPath)
	backRepo.BackRepoMultiple_rest.Backup(dirPath)
	backRepo.BackRepoName_display.Backup(dirPath)
	backRepo.BackRepoNon_arpeggiate.Backup(dirPath)
	backRepo.BackRepoNotations.Backup(dirPath)
	backRepo.BackRepoNote.Backup(dirPath)
	backRepo.BackRepoNote_size.Backup(dirPath)
	backRepo.BackRepoNote_type.Backup(dirPath)
	backRepo.BackRepoNotehead.Backup(dirPath)
	backRepo.BackRepoNotehead_text.Backup(dirPath)
	backRepo.BackRepoNumeral.Backup(dirPath)
	backRepo.BackRepoNumeral_key.Backup(dirPath)
	backRepo.BackRepoNumeral_root.Backup(dirPath)
	backRepo.BackRepoOctave_shift.Backup(dirPath)
	backRepo.BackRepoOffset.Backup(dirPath)
	backRepo.BackRepoOpus.Backup(dirPath)
	backRepo.BackRepoOrnaments.Backup(dirPath)
	backRepo.BackRepoOther_appearance.Backup(dirPath)
	backRepo.BackRepoOther_listening.Backup(dirPath)
	backRepo.BackRepoOther_notation.Backup(dirPath)
	backRepo.BackRepoOther_play.Backup(dirPath)
	backRepo.BackRepoPage_layout.Backup(dirPath)
	backRepo.BackRepoPage_margins.Backup(dirPath)
	backRepo.BackRepoPart_clef.Backup(dirPath)
	backRepo.BackRepoPart_group.Backup(dirPath)
	backRepo.BackRepoPart_link.Backup(dirPath)
	backRepo.BackRepoPart_list.Backup(dirPath)
	backRepo.BackRepoPart_symbol.Backup(dirPath)
	backRepo.BackRepoPart_transpose.Backup(dirPath)
	backRepo.BackRepoPedal.Backup(dirPath)
	backRepo.BackRepoPedal_tuning.Backup(dirPath)
	backRepo.BackRepoPercussion.Backup(dirPath)
	backRepo.BackRepoPitch.Backup(dirPath)
	backRepo.BackRepoPitched.Backup(dirPath)
	backRepo.BackRepoPlay.Backup(dirPath)
	backRepo.BackRepoPlayer.Backup(dirPath)
	backRepo.BackRepoPrincipal_voice.Backup(dirPath)
	backRepo.BackRepoPrint.Backup(dirPath)
	backRepo.BackRepoRelease.Backup(dirPath)
	backRepo.BackRepoRepeat.Backup(dirPath)
	backRepo.BackRepoRest.Backup(dirPath)
	backRepo.BackRepoRoot.Backup(dirPath)
	backRepo.BackRepoRoot_step.Backup(dirPath)
	backRepo.BackRepoScaling.Backup(dirPath)
	backRepo.BackRepoScordatura.Backup(dirPath)
	backRepo.BackRepoScore_instrument.Backup(dirPath)
	backRepo.BackRepoScore_part.Backup(dirPath)
	backRepo.BackRepoScore_partwise.Backup(dirPath)
	backRepo.BackRepoScore_timewise.Backup(dirPath)
	backRepo.BackRepoSegno.Backup(dirPath)
	backRepo.BackRepoSlash.Backup(dirPath)
	backRepo.BackRepoSlide.Backup(dirPath)
	backRepo.BackRepoSlur.Backup(dirPath)
	backRepo.BackRepoSound.Backup(dirPath)
	backRepo.BackRepoStaff_details.Backup(dirPath)
	backRepo.BackRepoStaff_divide.Backup(dirPath)
	backRepo.BackRepoStaff_layout.Backup(dirPath)
	backRepo.BackRepoStaff_size.Backup(dirPath)
	backRepo.BackRepoStaff_tuning.Backup(dirPath)
	backRepo.BackRepoStem.Backup(dirPath)
	backRepo.BackRepoStick.Backup(dirPath)
	backRepo.BackRepoString_mute.Backup(dirPath)
	backRepo.BackRepoStrong_accent.Backup(dirPath)
	backRepo.BackRepoSupports.Backup(dirPath)
	backRepo.BackRepoSwing.Backup(dirPath)
	backRepo.BackRepoSync.Backup(dirPath)
	backRepo.BackRepoSystem_dividers.Backup(dirPath)
	backRepo.BackRepoSystem_layout.Backup(dirPath)
	backRepo.BackRepoSystem_margins.Backup(dirPath)
	backRepo.BackRepoTap.Backup(dirPath)
	backRepo.BackRepoTechnical.Backup(dirPath)
	backRepo.BackRepoText_element_data.Backup(dirPath)
	backRepo.BackRepoTie.Backup(dirPath)
	backRepo.BackRepoTied.Backup(dirPath)
	backRepo.BackRepoTime.Backup(dirPath)
	backRepo.BackRepoTime_modification.Backup(dirPath)
	backRepo.BackRepoTimpani.Backup(dirPath)
	backRepo.BackRepoTranspose.Backup(dirPath)
	backRepo.BackRepoTremolo.Backup(dirPath)
	backRepo.BackRepoTuplet.Backup(dirPath)
	backRepo.BackRepoTuplet_dot.Backup(dirPath)
	backRepo.BackRepoTuplet_number.Backup(dirPath)
	backRepo.BackRepoTuplet_portion.Backup(dirPath)
	backRepo.BackRepoTuplet_type.Backup(dirPath)
	backRepo.BackRepoTyped_text.Backup(dirPath)
	backRepo.BackRepoUnpitched.Backup(dirPath)
	backRepo.BackRepoVirtual_instrument.Backup(dirPath)
	backRepo.BackRepoWait.Backup(dirPath)
	backRepo.BackRepoWavy_line.Backup(dirPath)
	backRepo.BackRepoWedge.Backup(dirPath)
	backRepo.BackRepoWood.Backup(dirPath)
	backRepo.BackRepoWork.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAccidental.BackupXL(file)
	backRepo.BackRepoAccidental_mark.BackupXL(file)
	backRepo.BackRepoAccidental_text.BackupXL(file)
	backRepo.BackRepoAccord.BackupXL(file)
	backRepo.BackRepoAccordion_registration.BackupXL(file)
	backRepo.BackRepoAnyType.BackupXL(file)
	backRepo.BackRepoAppearance.BackupXL(file)
	backRepo.BackRepoArpeggiate.BackupXL(file)
	backRepo.BackRepoArrow.BackupXL(file)
	backRepo.BackRepoArticulations.BackupXL(file)
	backRepo.BackRepoAssess.BackupXL(file)
	backRepo.BackRepoAttributes.BackupXL(file)
	backRepo.BackRepoBackup.BackupXL(file)
	backRepo.BackRepoBar_style_color.BackupXL(file)
	backRepo.BackRepoBarline.BackupXL(file)
	backRepo.BackRepoBarre.BackupXL(file)
	backRepo.BackRepoBass.BackupXL(file)
	backRepo.BackRepoBass_step.BackupXL(file)
	backRepo.BackRepoBeam.BackupXL(file)
	backRepo.BackRepoBeat_repeat.BackupXL(file)
	backRepo.BackRepoBeat_unit_tied.BackupXL(file)
	backRepo.BackRepoBeater.BackupXL(file)
	backRepo.BackRepoBend.BackupXL(file)
	backRepo.BackRepoBookmark.BackupXL(file)
	backRepo.BackRepoBracket.BackupXL(file)
	backRepo.BackRepoBreath_mark.BackupXL(file)
	backRepo.BackRepoCaesura.BackupXL(file)
	backRepo.BackRepoCancel.BackupXL(file)
	backRepo.BackRepoClef.BackupXL(file)
	backRepo.BackRepoCoda.BackupXL(file)
	backRepo.BackRepoCredit.BackupXL(file)
	backRepo.BackRepoDashes.BackupXL(file)
	backRepo.BackRepoDefaults.BackupXL(file)
	backRepo.BackRepoDegree.BackupXL(file)
	backRepo.BackRepoDegree_alter.BackupXL(file)
	backRepo.BackRepoDegree_type.BackupXL(file)
	backRepo.BackRepoDegree_value.BackupXL(file)
	backRepo.BackRepoDirection.BackupXL(file)
	backRepo.BackRepoDirection_type.BackupXL(file)
	backRepo.BackRepoDistance.BackupXL(file)
	backRepo.BackRepoDouble.BackupXL(file)
	backRepo.BackRepoDynamics.BackupXL(file)
	backRepo.BackRepoEffect.BackupXL(file)
	backRepo.BackRepoElision.BackupXL(file)
	backRepo.BackRepoEmpty.BackupXL(file)
	backRepo.BackRepoEmpty_font.BackupXL(file)
	backRepo.BackRepoEmpty_line.BackupXL(file)
	backRepo.BackRepoEmpty_placement.BackupXL(file)
	backRepo.BackRepoEmpty_placement_smufl.BackupXL(file)
	backRepo.BackRepoEmpty_print_object_style_align.BackupXL(file)
	backRepo.BackRepoEmpty_print_style.BackupXL(file)
	backRepo.BackRepoEmpty_print_style_align.BackupXL(file)
	backRepo.BackRepoEmpty_print_style_align_id.BackupXL(file)
	backRepo.BackRepoEmpty_trill_sound.BackupXL(file)
	backRepo.BackRepoEncoding.BackupXL(file)
	backRepo.BackRepoEnding.BackupXL(file)
	backRepo.BackRepoExtend.BackupXL(file)
	backRepo.BackRepoFeature.BackupXL(file)
	backRepo.BackRepoFermata.BackupXL(file)
	backRepo.BackRepoFigure.BackupXL(file)
	backRepo.BackRepoFigured_bass.BackupXL(file)
	backRepo.BackRepoFingering.BackupXL(file)
	backRepo.BackRepoFirst_fret.BackupXL(file)
	backRepo.BackRepoFoo.BackupXL(file)
	backRepo.BackRepoFor_part.BackupXL(file)
	backRepo.BackRepoFormatted_symbol.BackupXL(file)
	backRepo.BackRepoFormatted_symbol_id.BackupXL(file)
	backRepo.BackRepoForward.BackupXL(file)
	backRepo.BackRepoFrame.BackupXL(file)
	backRepo.BackRepoFrame_note.BackupXL(file)
	backRepo.BackRepoFret.BackupXL(file)
	backRepo.BackRepoGlass.BackupXL(file)
	backRepo.BackRepoGlissando.BackupXL(file)
	backRepo.BackRepoGlyph.BackupXL(file)
	backRepo.BackRepoGrace.BackupXL(file)
	backRepo.BackRepoGroup_barline.BackupXL(file)
	backRepo.BackRepoGroup_symbol.BackupXL(file)
	backRepo.BackRepoGrouping.BackupXL(file)
	backRepo.BackRepoHammer_on_pull_off.BackupXL(file)
	backRepo.BackRepoHandbell.BackupXL(file)
	backRepo.BackRepoHarmon_closed.BackupXL(file)
	backRepo.BackRepoHarmon_mute.BackupXL(file)
	backRepo.BackRepoHarmonic.BackupXL(file)
	backRepo.BackRepoHarmony.BackupXL(file)
	backRepo.BackRepoHarmony_alter.BackupXL(file)
	backRepo.BackRepoHarp_pedals.BackupXL(file)
	backRepo.BackRepoHeel_toe.BackupXL(file)
	backRepo.BackRepoHole.BackupXL(file)
	backRepo.BackRepoHole_closed.BackupXL(file)
	backRepo.BackRepoHorizontal_turn.BackupXL(file)
	backRepo.BackRepoIdentification.BackupXL(file)
	backRepo.BackRepoImage.BackupXL(file)
	backRepo.BackRepoInstrument.BackupXL(file)
	backRepo.BackRepoInstrument_change.BackupXL(file)
	backRepo.BackRepoInstrument_link.BackupXL(file)
	backRepo.BackRepoInterchangeable.BackupXL(file)
	backRepo.BackRepoInversion.BackupXL(file)
	backRepo.BackRepoKey.BackupXL(file)
	backRepo.BackRepoKey_accidental.BackupXL(file)
	backRepo.BackRepoKey_octave.BackupXL(file)
	backRepo.BackRepoKind.BackupXL(file)
	backRepo.BackRepoLevel.BackupXL(file)
	backRepo.BackRepoLine_detail.BackupXL(file)
	backRepo.BackRepoLine_width.BackupXL(file)
	backRepo.BackRepoLink.BackupXL(file)
	backRepo.BackRepoListen.BackupXL(file)
	backRepo.BackRepoListening.BackupXL(file)
	backRepo.BackRepoLyric.BackupXL(file)
	backRepo.BackRepoLyric_font.BackupXL(file)
	backRepo.BackRepoLyric_language.BackupXL(file)
	backRepo.BackRepoMeasure_layout.BackupXL(file)
	backRepo.BackRepoMeasure_numbering.BackupXL(file)
	backRepo.BackRepoMeasure_repeat.BackupXL(file)
	backRepo.BackRepoMeasure_style.BackupXL(file)
	backRepo.BackRepoMembrane.BackupXL(file)
	backRepo.BackRepoMetal.BackupXL(file)
	backRepo.BackRepoMetronome.BackupXL(file)
	backRepo.BackRepoMetronome_beam.BackupXL(file)
	backRepo.BackRepoMetronome_note.BackupXL(file)
	backRepo.BackRepoMetronome_tied.BackupXL(file)
	backRepo.BackRepoMetronome_tuplet.BackupXL(file)
	backRepo.BackRepoMidi_device.BackupXL(file)
	backRepo.BackRepoMidi_instrument.BackupXL(file)
	backRepo.BackRepoMiscellaneous.BackupXL(file)
	backRepo.BackRepoMiscellaneous_field.BackupXL(file)
	backRepo.BackRepoMordent.BackupXL(file)
	backRepo.BackRepoMultiple_rest.BackupXL(file)
	backRepo.BackRepoName_display.BackupXL(file)
	backRepo.BackRepoNon_arpeggiate.BackupXL(file)
	backRepo.BackRepoNotations.BackupXL(file)
	backRepo.BackRepoNote.BackupXL(file)
	backRepo.BackRepoNote_size.BackupXL(file)
	backRepo.BackRepoNote_type.BackupXL(file)
	backRepo.BackRepoNotehead.BackupXL(file)
	backRepo.BackRepoNotehead_text.BackupXL(file)
	backRepo.BackRepoNumeral.BackupXL(file)
	backRepo.BackRepoNumeral_key.BackupXL(file)
	backRepo.BackRepoNumeral_root.BackupXL(file)
	backRepo.BackRepoOctave_shift.BackupXL(file)
	backRepo.BackRepoOffset.BackupXL(file)
	backRepo.BackRepoOpus.BackupXL(file)
	backRepo.BackRepoOrnaments.BackupXL(file)
	backRepo.BackRepoOther_appearance.BackupXL(file)
	backRepo.BackRepoOther_listening.BackupXL(file)
	backRepo.BackRepoOther_notation.BackupXL(file)
	backRepo.BackRepoOther_play.BackupXL(file)
	backRepo.BackRepoPage_layout.BackupXL(file)
	backRepo.BackRepoPage_margins.BackupXL(file)
	backRepo.BackRepoPart_clef.BackupXL(file)
	backRepo.BackRepoPart_group.BackupXL(file)
	backRepo.BackRepoPart_link.BackupXL(file)
	backRepo.BackRepoPart_list.BackupXL(file)
	backRepo.BackRepoPart_symbol.BackupXL(file)
	backRepo.BackRepoPart_transpose.BackupXL(file)
	backRepo.BackRepoPedal.BackupXL(file)
	backRepo.BackRepoPedal_tuning.BackupXL(file)
	backRepo.BackRepoPercussion.BackupXL(file)
	backRepo.BackRepoPitch.BackupXL(file)
	backRepo.BackRepoPitched.BackupXL(file)
	backRepo.BackRepoPlay.BackupXL(file)
	backRepo.BackRepoPlayer.BackupXL(file)
	backRepo.BackRepoPrincipal_voice.BackupXL(file)
	backRepo.BackRepoPrint.BackupXL(file)
	backRepo.BackRepoRelease.BackupXL(file)
	backRepo.BackRepoRepeat.BackupXL(file)
	backRepo.BackRepoRest.BackupXL(file)
	backRepo.BackRepoRoot.BackupXL(file)
	backRepo.BackRepoRoot_step.BackupXL(file)
	backRepo.BackRepoScaling.BackupXL(file)
	backRepo.BackRepoScordatura.BackupXL(file)
	backRepo.BackRepoScore_instrument.BackupXL(file)
	backRepo.BackRepoScore_part.BackupXL(file)
	backRepo.BackRepoScore_partwise.BackupXL(file)
	backRepo.BackRepoScore_timewise.BackupXL(file)
	backRepo.BackRepoSegno.BackupXL(file)
	backRepo.BackRepoSlash.BackupXL(file)
	backRepo.BackRepoSlide.BackupXL(file)
	backRepo.BackRepoSlur.BackupXL(file)
	backRepo.BackRepoSound.BackupXL(file)
	backRepo.BackRepoStaff_details.BackupXL(file)
	backRepo.BackRepoStaff_divide.BackupXL(file)
	backRepo.BackRepoStaff_layout.BackupXL(file)
	backRepo.BackRepoStaff_size.BackupXL(file)
	backRepo.BackRepoStaff_tuning.BackupXL(file)
	backRepo.BackRepoStem.BackupXL(file)
	backRepo.BackRepoStick.BackupXL(file)
	backRepo.BackRepoString_mute.BackupXL(file)
	backRepo.BackRepoStrong_accent.BackupXL(file)
	backRepo.BackRepoSupports.BackupXL(file)
	backRepo.BackRepoSwing.BackupXL(file)
	backRepo.BackRepoSync.BackupXL(file)
	backRepo.BackRepoSystem_dividers.BackupXL(file)
	backRepo.BackRepoSystem_layout.BackupXL(file)
	backRepo.BackRepoSystem_margins.BackupXL(file)
	backRepo.BackRepoTap.BackupXL(file)
	backRepo.BackRepoTechnical.BackupXL(file)
	backRepo.BackRepoText_element_data.BackupXL(file)
	backRepo.BackRepoTie.BackupXL(file)
	backRepo.BackRepoTied.BackupXL(file)
	backRepo.BackRepoTime.BackupXL(file)
	backRepo.BackRepoTime_modification.BackupXL(file)
	backRepo.BackRepoTimpani.BackupXL(file)
	backRepo.BackRepoTranspose.BackupXL(file)
	backRepo.BackRepoTremolo.BackupXL(file)
	backRepo.BackRepoTuplet.BackupXL(file)
	backRepo.BackRepoTuplet_dot.BackupXL(file)
	backRepo.BackRepoTuplet_number.BackupXL(file)
	backRepo.BackRepoTuplet_portion.BackupXL(file)
	backRepo.BackRepoTuplet_type.BackupXL(file)
	backRepo.BackRepoTyped_text.BackupXL(file)
	backRepo.BackRepoUnpitched.BackupXL(file)
	backRepo.BackRepoVirtual_instrument.BackupXL(file)
	backRepo.BackRepoWait.BackupXL(file)
	backRepo.BackRepoWavy_line.BackupXL(file)
	backRepo.BackRepoWedge.BackupXL(file)
	backRepo.BackRepoWood.BackupXL(file)
	backRepo.BackRepoWork.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAccidental.RestorePhaseOne(dirPath)
	backRepo.BackRepoAccidental_mark.RestorePhaseOne(dirPath)
	backRepo.BackRepoAccidental_text.RestorePhaseOne(dirPath)
	backRepo.BackRepoAccord.RestorePhaseOne(dirPath)
	backRepo.BackRepoAccordion_registration.RestorePhaseOne(dirPath)
	backRepo.BackRepoAnyType.RestorePhaseOne(dirPath)
	backRepo.BackRepoAppearance.RestorePhaseOne(dirPath)
	backRepo.BackRepoArpeggiate.RestorePhaseOne(dirPath)
	backRepo.BackRepoArrow.RestorePhaseOne(dirPath)
	backRepo.BackRepoArticulations.RestorePhaseOne(dirPath)
	backRepo.BackRepoAssess.RestorePhaseOne(dirPath)
	backRepo.BackRepoAttributes.RestorePhaseOne(dirPath)
	backRepo.BackRepoBackup.RestorePhaseOne(dirPath)
	backRepo.BackRepoBar_style_color.RestorePhaseOne(dirPath)
	backRepo.BackRepoBarline.RestorePhaseOne(dirPath)
	backRepo.BackRepoBarre.RestorePhaseOne(dirPath)
	backRepo.BackRepoBass.RestorePhaseOne(dirPath)
	backRepo.BackRepoBass_step.RestorePhaseOne(dirPath)
	backRepo.BackRepoBeam.RestorePhaseOne(dirPath)
	backRepo.BackRepoBeat_repeat.RestorePhaseOne(dirPath)
	backRepo.BackRepoBeat_unit_tied.RestorePhaseOne(dirPath)
	backRepo.BackRepoBeater.RestorePhaseOne(dirPath)
	backRepo.BackRepoBend.RestorePhaseOne(dirPath)
	backRepo.BackRepoBookmark.RestorePhaseOne(dirPath)
	backRepo.BackRepoBracket.RestorePhaseOne(dirPath)
	backRepo.BackRepoBreath_mark.RestorePhaseOne(dirPath)
	backRepo.BackRepoCaesura.RestorePhaseOne(dirPath)
	backRepo.BackRepoCancel.RestorePhaseOne(dirPath)
	backRepo.BackRepoClef.RestorePhaseOne(dirPath)
	backRepo.BackRepoCoda.RestorePhaseOne(dirPath)
	backRepo.BackRepoCredit.RestorePhaseOne(dirPath)
	backRepo.BackRepoDashes.RestorePhaseOne(dirPath)
	backRepo.BackRepoDefaults.RestorePhaseOne(dirPath)
	backRepo.BackRepoDegree.RestorePhaseOne(dirPath)
	backRepo.BackRepoDegree_alter.RestorePhaseOne(dirPath)
	backRepo.BackRepoDegree_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoDegree_value.RestorePhaseOne(dirPath)
	backRepo.BackRepoDirection.RestorePhaseOne(dirPath)
	backRepo.BackRepoDirection_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoDistance.RestorePhaseOne(dirPath)
	backRepo.BackRepoDouble.RestorePhaseOne(dirPath)
	backRepo.BackRepoDynamics.RestorePhaseOne(dirPath)
	backRepo.BackRepoEffect.RestorePhaseOne(dirPath)
	backRepo.BackRepoElision.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_font.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_line.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_placement.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_placement_smufl.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_print_object_style_align.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_print_style.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_print_style_align.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_print_style_align_id.RestorePhaseOne(dirPath)
	backRepo.BackRepoEmpty_trill_sound.RestorePhaseOne(dirPath)
	backRepo.BackRepoEncoding.RestorePhaseOne(dirPath)
	backRepo.BackRepoEnding.RestorePhaseOne(dirPath)
	backRepo.BackRepoExtend.RestorePhaseOne(dirPath)
	backRepo.BackRepoFeature.RestorePhaseOne(dirPath)
	backRepo.BackRepoFermata.RestorePhaseOne(dirPath)
	backRepo.BackRepoFigure.RestorePhaseOne(dirPath)
	backRepo.BackRepoFigured_bass.RestorePhaseOne(dirPath)
	backRepo.BackRepoFingering.RestorePhaseOne(dirPath)
	backRepo.BackRepoFirst_fret.RestorePhaseOne(dirPath)
	backRepo.BackRepoFoo.RestorePhaseOne(dirPath)
	backRepo.BackRepoFor_part.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormatted_symbol.RestorePhaseOne(dirPath)
	backRepo.BackRepoFormatted_symbol_id.RestorePhaseOne(dirPath)
	backRepo.BackRepoForward.RestorePhaseOne(dirPath)
	backRepo.BackRepoFrame.RestorePhaseOne(dirPath)
	backRepo.BackRepoFrame_note.RestorePhaseOne(dirPath)
	backRepo.BackRepoFret.RestorePhaseOne(dirPath)
	backRepo.BackRepoGlass.RestorePhaseOne(dirPath)
	backRepo.BackRepoGlissando.RestorePhaseOne(dirPath)
	backRepo.BackRepoGlyph.RestorePhaseOne(dirPath)
	backRepo.BackRepoGrace.RestorePhaseOne(dirPath)
	backRepo.BackRepoGroup_barline.RestorePhaseOne(dirPath)
	backRepo.BackRepoGroup_symbol.RestorePhaseOne(dirPath)
	backRepo.BackRepoGrouping.RestorePhaseOne(dirPath)
	backRepo.BackRepoHammer_on_pull_off.RestorePhaseOne(dirPath)
	backRepo.BackRepoHandbell.RestorePhaseOne(dirPath)
	backRepo.BackRepoHarmon_closed.RestorePhaseOne(dirPath)
	backRepo.BackRepoHarmon_mute.RestorePhaseOne(dirPath)
	backRepo.BackRepoHarmonic.RestorePhaseOne(dirPath)
	backRepo.BackRepoHarmony.RestorePhaseOne(dirPath)
	backRepo.BackRepoHarmony_alter.RestorePhaseOne(dirPath)
	backRepo.BackRepoHarp_pedals.RestorePhaseOne(dirPath)
	backRepo.BackRepoHeel_toe.RestorePhaseOne(dirPath)
	backRepo.BackRepoHole.RestorePhaseOne(dirPath)
	backRepo.BackRepoHole_closed.RestorePhaseOne(dirPath)
	backRepo.BackRepoHorizontal_turn.RestorePhaseOne(dirPath)
	backRepo.BackRepoIdentification.RestorePhaseOne(dirPath)
	backRepo.BackRepoImage.RestorePhaseOne(dirPath)
	backRepo.BackRepoInstrument.RestorePhaseOne(dirPath)
	backRepo.BackRepoInstrument_change.RestorePhaseOne(dirPath)
	backRepo.BackRepoInstrument_link.RestorePhaseOne(dirPath)
	backRepo.BackRepoInterchangeable.RestorePhaseOne(dirPath)
	backRepo.BackRepoInversion.RestorePhaseOne(dirPath)
	backRepo.BackRepoKey.RestorePhaseOne(dirPath)
	backRepo.BackRepoKey_accidental.RestorePhaseOne(dirPath)
	backRepo.BackRepoKey_octave.RestorePhaseOne(dirPath)
	backRepo.BackRepoKind.RestorePhaseOne(dirPath)
	backRepo.BackRepoLevel.RestorePhaseOne(dirPath)
	backRepo.BackRepoLine_detail.RestorePhaseOne(dirPath)
	backRepo.BackRepoLine_width.RestorePhaseOne(dirPath)
	backRepo.BackRepoLink.RestorePhaseOne(dirPath)
	backRepo.BackRepoListen.RestorePhaseOne(dirPath)
	backRepo.BackRepoListening.RestorePhaseOne(dirPath)
	backRepo.BackRepoLyric.RestorePhaseOne(dirPath)
	backRepo.BackRepoLyric_font.RestorePhaseOne(dirPath)
	backRepo.BackRepoLyric_language.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeasure_layout.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeasure_numbering.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeasure_repeat.RestorePhaseOne(dirPath)
	backRepo.BackRepoMeasure_style.RestorePhaseOne(dirPath)
	backRepo.BackRepoMembrane.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetal.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetronome.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetronome_beam.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetronome_note.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetronome_tied.RestorePhaseOne(dirPath)
	backRepo.BackRepoMetronome_tuplet.RestorePhaseOne(dirPath)
	backRepo.BackRepoMidi_device.RestorePhaseOne(dirPath)
	backRepo.BackRepoMidi_instrument.RestorePhaseOne(dirPath)
	backRepo.BackRepoMiscellaneous.RestorePhaseOne(dirPath)
	backRepo.BackRepoMiscellaneous_field.RestorePhaseOne(dirPath)
	backRepo.BackRepoMordent.RestorePhaseOne(dirPath)
	backRepo.BackRepoMultiple_rest.RestorePhaseOne(dirPath)
	backRepo.BackRepoName_display.RestorePhaseOne(dirPath)
	backRepo.BackRepoNon_arpeggiate.RestorePhaseOne(dirPath)
	backRepo.BackRepoNotations.RestorePhaseOne(dirPath)
	backRepo.BackRepoNote.RestorePhaseOne(dirPath)
	backRepo.BackRepoNote_size.RestorePhaseOne(dirPath)
	backRepo.BackRepoNote_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoNotehead.RestorePhaseOne(dirPath)
	backRepo.BackRepoNotehead_text.RestorePhaseOne(dirPath)
	backRepo.BackRepoNumeral.RestorePhaseOne(dirPath)
	backRepo.BackRepoNumeral_key.RestorePhaseOne(dirPath)
	backRepo.BackRepoNumeral_root.RestorePhaseOne(dirPath)
	backRepo.BackRepoOctave_shift.RestorePhaseOne(dirPath)
	backRepo.BackRepoOffset.RestorePhaseOne(dirPath)
	backRepo.BackRepoOpus.RestorePhaseOne(dirPath)
	backRepo.BackRepoOrnaments.RestorePhaseOne(dirPath)
	backRepo.BackRepoOther_appearance.RestorePhaseOne(dirPath)
	backRepo.BackRepoOther_listening.RestorePhaseOne(dirPath)
	backRepo.BackRepoOther_notation.RestorePhaseOne(dirPath)
	backRepo.BackRepoOther_play.RestorePhaseOne(dirPath)
	backRepo.BackRepoPage_layout.RestorePhaseOne(dirPath)
	backRepo.BackRepoPage_margins.RestorePhaseOne(dirPath)
	backRepo.BackRepoPart_clef.RestorePhaseOne(dirPath)
	backRepo.BackRepoPart_group.RestorePhaseOne(dirPath)
	backRepo.BackRepoPart_link.RestorePhaseOne(dirPath)
	backRepo.BackRepoPart_list.RestorePhaseOne(dirPath)
	backRepo.BackRepoPart_symbol.RestorePhaseOne(dirPath)
	backRepo.BackRepoPart_transpose.RestorePhaseOne(dirPath)
	backRepo.BackRepoPedal.RestorePhaseOne(dirPath)
	backRepo.BackRepoPedal_tuning.RestorePhaseOne(dirPath)
	backRepo.BackRepoPercussion.RestorePhaseOne(dirPath)
	backRepo.BackRepoPitch.RestorePhaseOne(dirPath)
	backRepo.BackRepoPitched.RestorePhaseOne(dirPath)
	backRepo.BackRepoPlay.RestorePhaseOne(dirPath)
	backRepo.BackRepoPlayer.RestorePhaseOne(dirPath)
	backRepo.BackRepoPrincipal_voice.RestorePhaseOne(dirPath)
	backRepo.BackRepoPrint.RestorePhaseOne(dirPath)
	backRepo.BackRepoRelease.RestorePhaseOne(dirPath)
	backRepo.BackRepoRepeat.RestorePhaseOne(dirPath)
	backRepo.BackRepoRest.RestorePhaseOne(dirPath)
	backRepo.BackRepoRoot.RestorePhaseOne(dirPath)
	backRepo.BackRepoRoot_step.RestorePhaseOne(dirPath)
	backRepo.BackRepoScaling.RestorePhaseOne(dirPath)
	backRepo.BackRepoScordatura.RestorePhaseOne(dirPath)
	backRepo.BackRepoScore_instrument.RestorePhaseOne(dirPath)
	backRepo.BackRepoScore_part.RestorePhaseOne(dirPath)
	backRepo.BackRepoScore_partwise.RestorePhaseOne(dirPath)
	backRepo.BackRepoScore_timewise.RestorePhaseOne(dirPath)
	backRepo.BackRepoSegno.RestorePhaseOne(dirPath)
	backRepo.BackRepoSlash.RestorePhaseOne(dirPath)
	backRepo.BackRepoSlide.RestorePhaseOne(dirPath)
	backRepo.BackRepoSlur.RestorePhaseOne(dirPath)
	backRepo.BackRepoSound.RestorePhaseOne(dirPath)
	backRepo.BackRepoStaff_details.RestorePhaseOne(dirPath)
	backRepo.BackRepoStaff_divide.RestorePhaseOne(dirPath)
	backRepo.BackRepoStaff_layout.RestorePhaseOne(dirPath)
	backRepo.BackRepoStaff_size.RestorePhaseOne(dirPath)
	backRepo.BackRepoStaff_tuning.RestorePhaseOne(dirPath)
	backRepo.BackRepoStem.RestorePhaseOne(dirPath)
	backRepo.BackRepoStick.RestorePhaseOne(dirPath)
	backRepo.BackRepoString_mute.RestorePhaseOne(dirPath)
	backRepo.BackRepoStrong_accent.RestorePhaseOne(dirPath)
	backRepo.BackRepoSupports.RestorePhaseOne(dirPath)
	backRepo.BackRepoSwing.RestorePhaseOne(dirPath)
	backRepo.BackRepoSync.RestorePhaseOne(dirPath)
	backRepo.BackRepoSystem_dividers.RestorePhaseOne(dirPath)
	backRepo.BackRepoSystem_layout.RestorePhaseOne(dirPath)
	backRepo.BackRepoSystem_margins.RestorePhaseOne(dirPath)
	backRepo.BackRepoTap.RestorePhaseOne(dirPath)
	backRepo.BackRepoTechnical.RestorePhaseOne(dirPath)
	backRepo.BackRepoText_element_data.RestorePhaseOne(dirPath)
	backRepo.BackRepoTie.RestorePhaseOne(dirPath)
	backRepo.BackRepoTied.RestorePhaseOne(dirPath)
	backRepo.BackRepoTime.RestorePhaseOne(dirPath)
	backRepo.BackRepoTime_modification.RestorePhaseOne(dirPath)
	backRepo.BackRepoTimpani.RestorePhaseOne(dirPath)
	backRepo.BackRepoTranspose.RestorePhaseOne(dirPath)
	backRepo.BackRepoTremolo.RestorePhaseOne(dirPath)
	backRepo.BackRepoTuplet.RestorePhaseOne(dirPath)
	backRepo.BackRepoTuplet_dot.RestorePhaseOne(dirPath)
	backRepo.BackRepoTuplet_number.RestorePhaseOne(dirPath)
	backRepo.BackRepoTuplet_portion.RestorePhaseOne(dirPath)
	backRepo.BackRepoTuplet_type.RestorePhaseOne(dirPath)
	backRepo.BackRepoTyped_text.RestorePhaseOne(dirPath)
	backRepo.BackRepoUnpitched.RestorePhaseOne(dirPath)
	backRepo.BackRepoVirtual_instrument.RestorePhaseOne(dirPath)
	backRepo.BackRepoWait.RestorePhaseOne(dirPath)
	backRepo.BackRepoWavy_line.RestorePhaseOne(dirPath)
	backRepo.BackRepoWedge.RestorePhaseOne(dirPath)
	backRepo.BackRepoWood.RestorePhaseOne(dirPath)
	backRepo.BackRepoWork.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAccidental.RestorePhaseTwo()
	backRepo.BackRepoAccidental_mark.RestorePhaseTwo()
	backRepo.BackRepoAccidental_text.RestorePhaseTwo()
	backRepo.BackRepoAccord.RestorePhaseTwo()
	backRepo.BackRepoAccordion_registration.RestorePhaseTwo()
	backRepo.BackRepoAnyType.RestorePhaseTwo()
	backRepo.BackRepoAppearance.RestorePhaseTwo()
	backRepo.BackRepoArpeggiate.RestorePhaseTwo()
	backRepo.BackRepoArrow.RestorePhaseTwo()
	backRepo.BackRepoArticulations.RestorePhaseTwo()
	backRepo.BackRepoAssess.RestorePhaseTwo()
	backRepo.BackRepoAttributes.RestorePhaseTwo()
	backRepo.BackRepoBackup.RestorePhaseTwo()
	backRepo.BackRepoBar_style_color.RestorePhaseTwo()
	backRepo.BackRepoBarline.RestorePhaseTwo()
	backRepo.BackRepoBarre.RestorePhaseTwo()
	backRepo.BackRepoBass.RestorePhaseTwo()
	backRepo.BackRepoBass_step.RestorePhaseTwo()
	backRepo.BackRepoBeam.RestorePhaseTwo()
	backRepo.BackRepoBeat_repeat.RestorePhaseTwo()
	backRepo.BackRepoBeat_unit_tied.RestorePhaseTwo()
	backRepo.BackRepoBeater.RestorePhaseTwo()
	backRepo.BackRepoBend.RestorePhaseTwo()
	backRepo.BackRepoBookmark.RestorePhaseTwo()
	backRepo.BackRepoBracket.RestorePhaseTwo()
	backRepo.BackRepoBreath_mark.RestorePhaseTwo()
	backRepo.BackRepoCaesura.RestorePhaseTwo()
	backRepo.BackRepoCancel.RestorePhaseTwo()
	backRepo.BackRepoClef.RestorePhaseTwo()
	backRepo.BackRepoCoda.RestorePhaseTwo()
	backRepo.BackRepoCredit.RestorePhaseTwo()
	backRepo.BackRepoDashes.RestorePhaseTwo()
	backRepo.BackRepoDefaults.RestorePhaseTwo()
	backRepo.BackRepoDegree.RestorePhaseTwo()
	backRepo.BackRepoDegree_alter.RestorePhaseTwo()
	backRepo.BackRepoDegree_type.RestorePhaseTwo()
	backRepo.BackRepoDegree_value.RestorePhaseTwo()
	backRepo.BackRepoDirection.RestorePhaseTwo()
	backRepo.BackRepoDirection_type.RestorePhaseTwo()
	backRepo.BackRepoDistance.RestorePhaseTwo()
	backRepo.BackRepoDouble.RestorePhaseTwo()
	backRepo.BackRepoDynamics.RestorePhaseTwo()
	backRepo.BackRepoEffect.RestorePhaseTwo()
	backRepo.BackRepoElision.RestorePhaseTwo()
	backRepo.BackRepoEmpty.RestorePhaseTwo()
	backRepo.BackRepoEmpty_font.RestorePhaseTwo()
	backRepo.BackRepoEmpty_line.RestorePhaseTwo()
	backRepo.BackRepoEmpty_placement.RestorePhaseTwo()
	backRepo.BackRepoEmpty_placement_smufl.RestorePhaseTwo()
	backRepo.BackRepoEmpty_print_object_style_align.RestorePhaseTwo()
	backRepo.BackRepoEmpty_print_style.RestorePhaseTwo()
	backRepo.BackRepoEmpty_print_style_align.RestorePhaseTwo()
	backRepo.BackRepoEmpty_print_style_align_id.RestorePhaseTwo()
	backRepo.BackRepoEmpty_trill_sound.RestorePhaseTwo()
	backRepo.BackRepoEncoding.RestorePhaseTwo()
	backRepo.BackRepoEnding.RestorePhaseTwo()
	backRepo.BackRepoExtend.RestorePhaseTwo()
	backRepo.BackRepoFeature.RestorePhaseTwo()
	backRepo.BackRepoFermata.RestorePhaseTwo()
	backRepo.BackRepoFigure.RestorePhaseTwo()
	backRepo.BackRepoFigured_bass.RestorePhaseTwo()
	backRepo.BackRepoFingering.RestorePhaseTwo()
	backRepo.BackRepoFirst_fret.RestorePhaseTwo()
	backRepo.BackRepoFoo.RestorePhaseTwo()
	backRepo.BackRepoFor_part.RestorePhaseTwo()
	backRepo.BackRepoFormatted_symbol.RestorePhaseTwo()
	backRepo.BackRepoFormatted_symbol_id.RestorePhaseTwo()
	backRepo.BackRepoForward.RestorePhaseTwo()
	backRepo.BackRepoFrame.RestorePhaseTwo()
	backRepo.BackRepoFrame_note.RestorePhaseTwo()
	backRepo.BackRepoFret.RestorePhaseTwo()
	backRepo.BackRepoGlass.RestorePhaseTwo()
	backRepo.BackRepoGlissando.RestorePhaseTwo()
	backRepo.BackRepoGlyph.RestorePhaseTwo()
	backRepo.BackRepoGrace.RestorePhaseTwo()
	backRepo.BackRepoGroup_barline.RestorePhaseTwo()
	backRepo.BackRepoGroup_symbol.RestorePhaseTwo()
	backRepo.BackRepoGrouping.RestorePhaseTwo()
	backRepo.BackRepoHammer_on_pull_off.RestorePhaseTwo()
	backRepo.BackRepoHandbell.RestorePhaseTwo()
	backRepo.BackRepoHarmon_closed.RestorePhaseTwo()
	backRepo.BackRepoHarmon_mute.RestorePhaseTwo()
	backRepo.BackRepoHarmonic.RestorePhaseTwo()
	backRepo.BackRepoHarmony.RestorePhaseTwo()
	backRepo.BackRepoHarmony_alter.RestorePhaseTwo()
	backRepo.BackRepoHarp_pedals.RestorePhaseTwo()
	backRepo.BackRepoHeel_toe.RestorePhaseTwo()
	backRepo.BackRepoHole.RestorePhaseTwo()
	backRepo.BackRepoHole_closed.RestorePhaseTwo()
	backRepo.BackRepoHorizontal_turn.RestorePhaseTwo()
	backRepo.BackRepoIdentification.RestorePhaseTwo()
	backRepo.BackRepoImage.RestorePhaseTwo()
	backRepo.BackRepoInstrument.RestorePhaseTwo()
	backRepo.BackRepoInstrument_change.RestorePhaseTwo()
	backRepo.BackRepoInstrument_link.RestorePhaseTwo()
	backRepo.BackRepoInterchangeable.RestorePhaseTwo()
	backRepo.BackRepoInversion.RestorePhaseTwo()
	backRepo.BackRepoKey.RestorePhaseTwo()
	backRepo.BackRepoKey_accidental.RestorePhaseTwo()
	backRepo.BackRepoKey_octave.RestorePhaseTwo()
	backRepo.BackRepoKind.RestorePhaseTwo()
	backRepo.BackRepoLevel.RestorePhaseTwo()
	backRepo.BackRepoLine_detail.RestorePhaseTwo()
	backRepo.BackRepoLine_width.RestorePhaseTwo()
	backRepo.BackRepoLink.RestorePhaseTwo()
	backRepo.BackRepoListen.RestorePhaseTwo()
	backRepo.BackRepoListening.RestorePhaseTwo()
	backRepo.BackRepoLyric.RestorePhaseTwo()
	backRepo.BackRepoLyric_font.RestorePhaseTwo()
	backRepo.BackRepoLyric_language.RestorePhaseTwo()
	backRepo.BackRepoMeasure_layout.RestorePhaseTwo()
	backRepo.BackRepoMeasure_numbering.RestorePhaseTwo()
	backRepo.BackRepoMeasure_repeat.RestorePhaseTwo()
	backRepo.BackRepoMeasure_style.RestorePhaseTwo()
	backRepo.BackRepoMembrane.RestorePhaseTwo()
	backRepo.BackRepoMetal.RestorePhaseTwo()
	backRepo.BackRepoMetronome.RestorePhaseTwo()
	backRepo.BackRepoMetronome_beam.RestorePhaseTwo()
	backRepo.BackRepoMetronome_note.RestorePhaseTwo()
	backRepo.BackRepoMetronome_tied.RestorePhaseTwo()
	backRepo.BackRepoMetronome_tuplet.RestorePhaseTwo()
	backRepo.BackRepoMidi_device.RestorePhaseTwo()
	backRepo.BackRepoMidi_instrument.RestorePhaseTwo()
	backRepo.BackRepoMiscellaneous.RestorePhaseTwo()
	backRepo.BackRepoMiscellaneous_field.RestorePhaseTwo()
	backRepo.BackRepoMordent.RestorePhaseTwo()
	backRepo.BackRepoMultiple_rest.RestorePhaseTwo()
	backRepo.BackRepoName_display.RestorePhaseTwo()
	backRepo.BackRepoNon_arpeggiate.RestorePhaseTwo()
	backRepo.BackRepoNotations.RestorePhaseTwo()
	backRepo.BackRepoNote.RestorePhaseTwo()
	backRepo.BackRepoNote_size.RestorePhaseTwo()
	backRepo.BackRepoNote_type.RestorePhaseTwo()
	backRepo.BackRepoNotehead.RestorePhaseTwo()
	backRepo.BackRepoNotehead_text.RestorePhaseTwo()
	backRepo.BackRepoNumeral.RestorePhaseTwo()
	backRepo.BackRepoNumeral_key.RestorePhaseTwo()
	backRepo.BackRepoNumeral_root.RestorePhaseTwo()
	backRepo.BackRepoOctave_shift.RestorePhaseTwo()
	backRepo.BackRepoOffset.RestorePhaseTwo()
	backRepo.BackRepoOpus.RestorePhaseTwo()
	backRepo.BackRepoOrnaments.RestorePhaseTwo()
	backRepo.BackRepoOther_appearance.RestorePhaseTwo()
	backRepo.BackRepoOther_listening.RestorePhaseTwo()
	backRepo.BackRepoOther_notation.RestorePhaseTwo()
	backRepo.BackRepoOther_play.RestorePhaseTwo()
	backRepo.BackRepoPage_layout.RestorePhaseTwo()
	backRepo.BackRepoPage_margins.RestorePhaseTwo()
	backRepo.BackRepoPart_clef.RestorePhaseTwo()
	backRepo.BackRepoPart_group.RestorePhaseTwo()
	backRepo.BackRepoPart_link.RestorePhaseTwo()
	backRepo.BackRepoPart_list.RestorePhaseTwo()
	backRepo.BackRepoPart_symbol.RestorePhaseTwo()
	backRepo.BackRepoPart_transpose.RestorePhaseTwo()
	backRepo.BackRepoPedal.RestorePhaseTwo()
	backRepo.BackRepoPedal_tuning.RestorePhaseTwo()
	backRepo.BackRepoPercussion.RestorePhaseTwo()
	backRepo.BackRepoPitch.RestorePhaseTwo()
	backRepo.BackRepoPitched.RestorePhaseTwo()
	backRepo.BackRepoPlay.RestorePhaseTwo()
	backRepo.BackRepoPlayer.RestorePhaseTwo()
	backRepo.BackRepoPrincipal_voice.RestorePhaseTwo()
	backRepo.BackRepoPrint.RestorePhaseTwo()
	backRepo.BackRepoRelease.RestorePhaseTwo()
	backRepo.BackRepoRepeat.RestorePhaseTwo()
	backRepo.BackRepoRest.RestorePhaseTwo()
	backRepo.BackRepoRoot.RestorePhaseTwo()
	backRepo.BackRepoRoot_step.RestorePhaseTwo()
	backRepo.BackRepoScaling.RestorePhaseTwo()
	backRepo.BackRepoScordatura.RestorePhaseTwo()
	backRepo.BackRepoScore_instrument.RestorePhaseTwo()
	backRepo.BackRepoScore_part.RestorePhaseTwo()
	backRepo.BackRepoScore_partwise.RestorePhaseTwo()
	backRepo.BackRepoScore_timewise.RestorePhaseTwo()
	backRepo.BackRepoSegno.RestorePhaseTwo()
	backRepo.BackRepoSlash.RestorePhaseTwo()
	backRepo.BackRepoSlide.RestorePhaseTwo()
	backRepo.BackRepoSlur.RestorePhaseTwo()
	backRepo.BackRepoSound.RestorePhaseTwo()
	backRepo.BackRepoStaff_details.RestorePhaseTwo()
	backRepo.BackRepoStaff_divide.RestorePhaseTwo()
	backRepo.BackRepoStaff_layout.RestorePhaseTwo()
	backRepo.BackRepoStaff_size.RestorePhaseTwo()
	backRepo.BackRepoStaff_tuning.RestorePhaseTwo()
	backRepo.BackRepoStem.RestorePhaseTwo()
	backRepo.BackRepoStick.RestorePhaseTwo()
	backRepo.BackRepoString_mute.RestorePhaseTwo()
	backRepo.BackRepoStrong_accent.RestorePhaseTwo()
	backRepo.BackRepoSupports.RestorePhaseTwo()
	backRepo.BackRepoSwing.RestorePhaseTwo()
	backRepo.BackRepoSync.RestorePhaseTwo()
	backRepo.BackRepoSystem_dividers.RestorePhaseTwo()
	backRepo.BackRepoSystem_layout.RestorePhaseTwo()
	backRepo.BackRepoSystem_margins.RestorePhaseTwo()
	backRepo.BackRepoTap.RestorePhaseTwo()
	backRepo.BackRepoTechnical.RestorePhaseTwo()
	backRepo.BackRepoText_element_data.RestorePhaseTwo()
	backRepo.BackRepoTie.RestorePhaseTwo()
	backRepo.BackRepoTied.RestorePhaseTwo()
	backRepo.BackRepoTime.RestorePhaseTwo()
	backRepo.BackRepoTime_modification.RestorePhaseTwo()
	backRepo.BackRepoTimpani.RestorePhaseTwo()
	backRepo.BackRepoTranspose.RestorePhaseTwo()
	backRepo.BackRepoTremolo.RestorePhaseTwo()
	backRepo.BackRepoTuplet.RestorePhaseTwo()
	backRepo.BackRepoTuplet_dot.RestorePhaseTwo()
	backRepo.BackRepoTuplet_number.RestorePhaseTwo()
	backRepo.BackRepoTuplet_portion.RestorePhaseTwo()
	backRepo.BackRepoTuplet_type.RestorePhaseTwo()
	backRepo.BackRepoTyped_text.RestorePhaseTwo()
	backRepo.BackRepoUnpitched.RestorePhaseTwo()
	backRepo.BackRepoVirtual_instrument.RestorePhaseTwo()
	backRepo.BackRepoWait.RestorePhaseTwo()
	backRepo.BackRepoWavy_line.RestorePhaseTwo()
	backRepo.BackRepoWedge.RestorePhaseTwo()
	backRepo.BackRepoWood.RestorePhaseTwo()
	backRepo.BackRepoWork.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAccidental.RestoreXLPhaseOne(file)
	backRepo.BackRepoAccidental_mark.RestoreXLPhaseOne(file)
	backRepo.BackRepoAccidental_text.RestoreXLPhaseOne(file)
	backRepo.BackRepoAccord.RestoreXLPhaseOne(file)
	backRepo.BackRepoAccordion_registration.RestoreXLPhaseOne(file)
	backRepo.BackRepoAnyType.RestoreXLPhaseOne(file)
	backRepo.BackRepoAppearance.RestoreXLPhaseOne(file)
	backRepo.BackRepoArpeggiate.RestoreXLPhaseOne(file)
	backRepo.BackRepoArrow.RestoreXLPhaseOne(file)
	backRepo.BackRepoArticulations.RestoreXLPhaseOne(file)
	backRepo.BackRepoAssess.RestoreXLPhaseOne(file)
	backRepo.BackRepoAttributes.RestoreXLPhaseOne(file)
	backRepo.BackRepoBackup.RestoreXLPhaseOne(file)
	backRepo.BackRepoBar_style_color.RestoreXLPhaseOne(file)
	backRepo.BackRepoBarline.RestoreXLPhaseOne(file)
	backRepo.BackRepoBarre.RestoreXLPhaseOne(file)
	backRepo.BackRepoBass.RestoreXLPhaseOne(file)
	backRepo.BackRepoBass_step.RestoreXLPhaseOne(file)
	backRepo.BackRepoBeam.RestoreXLPhaseOne(file)
	backRepo.BackRepoBeat_repeat.RestoreXLPhaseOne(file)
	backRepo.BackRepoBeat_unit_tied.RestoreXLPhaseOne(file)
	backRepo.BackRepoBeater.RestoreXLPhaseOne(file)
	backRepo.BackRepoBend.RestoreXLPhaseOne(file)
	backRepo.BackRepoBookmark.RestoreXLPhaseOne(file)
	backRepo.BackRepoBracket.RestoreXLPhaseOne(file)
	backRepo.BackRepoBreath_mark.RestoreXLPhaseOne(file)
	backRepo.BackRepoCaesura.RestoreXLPhaseOne(file)
	backRepo.BackRepoCancel.RestoreXLPhaseOne(file)
	backRepo.BackRepoClef.RestoreXLPhaseOne(file)
	backRepo.BackRepoCoda.RestoreXLPhaseOne(file)
	backRepo.BackRepoCredit.RestoreXLPhaseOne(file)
	backRepo.BackRepoDashes.RestoreXLPhaseOne(file)
	backRepo.BackRepoDefaults.RestoreXLPhaseOne(file)
	backRepo.BackRepoDegree.RestoreXLPhaseOne(file)
	backRepo.BackRepoDegree_alter.RestoreXLPhaseOne(file)
	backRepo.BackRepoDegree_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoDegree_value.RestoreXLPhaseOne(file)
	backRepo.BackRepoDirection.RestoreXLPhaseOne(file)
	backRepo.BackRepoDirection_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoDistance.RestoreXLPhaseOne(file)
	backRepo.BackRepoDouble.RestoreXLPhaseOne(file)
	backRepo.BackRepoDynamics.RestoreXLPhaseOne(file)
	backRepo.BackRepoEffect.RestoreXLPhaseOne(file)
	backRepo.BackRepoElision.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_font.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_line.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_placement.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_placement_smufl.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_print_object_style_align.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_print_style.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_print_style_align.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_print_style_align_id.RestoreXLPhaseOne(file)
	backRepo.BackRepoEmpty_trill_sound.RestoreXLPhaseOne(file)
	backRepo.BackRepoEncoding.RestoreXLPhaseOne(file)
	backRepo.BackRepoEnding.RestoreXLPhaseOne(file)
	backRepo.BackRepoExtend.RestoreXLPhaseOne(file)
	backRepo.BackRepoFeature.RestoreXLPhaseOne(file)
	backRepo.BackRepoFermata.RestoreXLPhaseOne(file)
	backRepo.BackRepoFigure.RestoreXLPhaseOne(file)
	backRepo.BackRepoFigured_bass.RestoreXLPhaseOne(file)
	backRepo.BackRepoFingering.RestoreXLPhaseOne(file)
	backRepo.BackRepoFirst_fret.RestoreXLPhaseOne(file)
	backRepo.BackRepoFoo.RestoreXLPhaseOne(file)
	backRepo.BackRepoFor_part.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormatted_symbol.RestoreXLPhaseOne(file)
	backRepo.BackRepoFormatted_symbol_id.RestoreXLPhaseOne(file)
	backRepo.BackRepoForward.RestoreXLPhaseOne(file)
	backRepo.BackRepoFrame.RestoreXLPhaseOne(file)
	backRepo.BackRepoFrame_note.RestoreXLPhaseOne(file)
	backRepo.BackRepoFret.RestoreXLPhaseOne(file)
	backRepo.BackRepoGlass.RestoreXLPhaseOne(file)
	backRepo.BackRepoGlissando.RestoreXLPhaseOne(file)
	backRepo.BackRepoGlyph.RestoreXLPhaseOne(file)
	backRepo.BackRepoGrace.RestoreXLPhaseOne(file)
	backRepo.BackRepoGroup_barline.RestoreXLPhaseOne(file)
	backRepo.BackRepoGroup_symbol.RestoreXLPhaseOne(file)
	backRepo.BackRepoGrouping.RestoreXLPhaseOne(file)
	backRepo.BackRepoHammer_on_pull_off.RestoreXLPhaseOne(file)
	backRepo.BackRepoHandbell.RestoreXLPhaseOne(file)
	backRepo.BackRepoHarmon_closed.RestoreXLPhaseOne(file)
	backRepo.BackRepoHarmon_mute.RestoreXLPhaseOne(file)
	backRepo.BackRepoHarmonic.RestoreXLPhaseOne(file)
	backRepo.BackRepoHarmony.RestoreXLPhaseOne(file)
	backRepo.BackRepoHarmony_alter.RestoreXLPhaseOne(file)
	backRepo.BackRepoHarp_pedals.RestoreXLPhaseOne(file)
	backRepo.BackRepoHeel_toe.RestoreXLPhaseOne(file)
	backRepo.BackRepoHole.RestoreXLPhaseOne(file)
	backRepo.BackRepoHole_closed.RestoreXLPhaseOne(file)
	backRepo.BackRepoHorizontal_turn.RestoreXLPhaseOne(file)
	backRepo.BackRepoIdentification.RestoreXLPhaseOne(file)
	backRepo.BackRepoImage.RestoreXLPhaseOne(file)
	backRepo.BackRepoInstrument.RestoreXLPhaseOne(file)
	backRepo.BackRepoInstrument_change.RestoreXLPhaseOne(file)
	backRepo.BackRepoInstrument_link.RestoreXLPhaseOne(file)
	backRepo.BackRepoInterchangeable.RestoreXLPhaseOne(file)
	backRepo.BackRepoInversion.RestoreXLPhaseOne(file)
	backRepo.BackRepoKey.RestoreXLPhaseOne(file)
	backRepo.BackRepoKey_accidental.RestoreXLPhaseOne(file)
	backRepo.BackRepoKey_octave.RestoreXLPhaseOne(file)
	backRepo.BackRepoKind.RestoreXLPhaseOne(file)
	backRepo.BackRepoLevel.RestoreXLPhaseOne(file)
	backRepo.BackRepoLine_detail.RestoreXLPhaseOne(file)
	backRepo.BackRepoLine_width.RestoreXLPhaseOne(file)
	backRepo.BackRepoLink.RestoreXLPhaseOne(file)
	backRepo.BackRepoListen.RestoreXLPhaseOne(file)
	backRepo.BackRepoListening.RestoreXLPhaseOne(file)
	backRepo.BackRepoLyric.RestoreXLPhaseOne(file)
	backRepo.BackRepoLyric_font.RestoreXLPhaseOne(file)
	backRepo.BackRepoLyric_language.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeasure_layout.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeasure_numbering.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeasure_repeat.RestoreXLPhaseOne(file)
	backRepo.BackRepoMeasure_style.RestoreXLPhaseOne(file)
	backRepo.BackRepoMembrane.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetal.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetronome.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetronome_beam.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetronome_note.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetronome_tied.RestoreXLPhaseOne(file)
	backRepo.BackRepoMetronome_tuplet.RestoreXLPhaseOne(file)
	backRepo.BackRepoMidi_device.RestoreXLPhaseOne(file)
	backRepo.BackRepoMidi_instrument.RestoreXLPhaseOne(file)
	backRepo.BackRepoMiscellaneous.RestoreXLPhaseOne(file)
	backRepo.BackRepoMiscellaneous_field.RestoreXLPhaseOne(file)
	backRepo.BackRepoMordent.RestoreXLPhaseOne(file)
	backRepo.BackRepoMultiple_rest.RestoreXLPhaseOne(file)
	backRepo.BackRepoName_display.RestoreXLPhaseOne(file)
	backRepo.BackRepoNon_arpeggiate.RestoreXLPhaseOne(file)
	backRepo.BackRepoNotations.RestoreXLPhaseOne(file)
	backRepo.BackRepoNote.RestoreXLPhaseOne(file)
	backRepo.BackRepoNote_size.RestoreXLPhaseOne(file)
	backRepo.BackRepoNote_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoNotehead.RestoreXLPhaseOne(file)
	backRepo.BackRepoNotehead_text.RestoreXLPhaseOne(file)
	backRepo.BackRepoNumeral.RestoreXLPhaseOne(file)
	backRepo.BackRepoNumeral_key.RestoreXLPhaseOne(file)
	backRepo.BackRepoNumeral_root.RestoreXLPhaseOne(file)
	backRepo.BackRepoOctave_shift.RestoreXLPhaseOne(file)
	backRepo.BackRepoOffset.RestoreXLPhaseOne(file)
	backRepo.BackRepoOpus.RestoreXLPhaseOne(file)
	backRepo.BackRepoOrnaments.RestoreXLPhaseOne(file)
	backRepo.BackRepoOther_appearance.RestoreXLPhaseOne(file)
	backRepo.BackRepoOther_listening.RestoreXLPhaseOne(file)
	backRepo.BackRepoOther_notation.RestoreXLPhaseOne(file)
	backRepo.BackRepoOther_play.RestoreXLPhaseOne(file)
	backRepo.BackRepoPage_layout.RestoreXLPhaseOne(file)
	backRepo.BackRepoPage_margins.RestoreXLPhaseOne(file)
	backRepo.BackRepoPart_clef.RestoreXLPhaseOne(file)
	backRepo.BackRepoPart_group.RestoreXLPhaseOne(file)
	backRepo.BackRepoPart_link.RestoreXLPhaseOne(file)
	backRepo.BackRepoPart_list.RestoreXLPhaseOne(file)
	backRepo.BackRepoPart_symbol.RestoreXLPhaseOne(file)
	backRepo.BackRepoPart_transpose.RestoreXLPhaseOne(file)
	backRepo.BackRepoPedal.RestoreXLPhaseOne(file)
	backRepo.BackRepoPedal_tuning.RestoreXLPhaseOne(file)
	backRepo.BackRepoPercussion.RestoreXLPhaseOne(file)
	backRepo.BackRepoPitch.RestoreXLPhaseOne(file)
	backRepo.BackRepoPitched.RestoreXLPhaseOne(file)
	backRepo.BackRepoPlay.RestoreXLPhaseOne(file)
	backRepo.BackRepoPlayer.RestoreXLPhaseOne(file)
	backRepo.BackRepoPrincipal_voice.RestoreXLPhaseOne(file)
	backRepo.BackRepoPrint.RestoreXLPhaseOne(file)
	backRepo.BackRepoRelease.RestoreXLPhaseOne(file)
	backRepo.BackRepoRepeat.RestoreXLPhaseOne(file)
	backRepo.BackRepoRest.RestoreXLPhaseOne(file)
	backRepo.BackRepoRoot.RestoreXLPhaseOne(file)
	backRepo.BackRepoRoot_step.RestoreXLPhaseOne(file)
	backRepo.BackRepoScaling.RestoreXLPhaseOne(file)
	backRepo.BackRepoScordatura.RestoreXLPhaseOne(file)
	backRepo.BackRepoScore_instrument.RestoreXLPhaseOne(file)
	backRepo.BackRepoScore_part.RestoreXLPhaseOne(file)
	backRepo.BackRepoScore_partwise.RestoreXLPhaseOne(file)
	backRepo.BackRepoScore_timewise.RestoreXLPhaseOne(file)
	backRepo.BackRepoSegno.RestoreXLPhaseOne(file)
	backRepo.BackRepoSlash.RestoreXLPhaseOne(file)
	backRepo.BackRepoSlide.RestoreXLPhaseOne(file)
	backRepo.BackRepoSlur.RestoreXLPhaseOne(file)
	backRepo.BackRepoSound.RestoreXLPhaseOne(file)
	backRepo.BackRepoStaff_details.RestoreXLPhaseOne(file)
	backRepo.BackRepoStaff_divide.RestoreXLPhaseOne(file)
	backRepo.BackRepoStaff_layout.RestoreXLPhaseOne(file)
	backRepo.BackRepoStaff_size.RestoreXLPhaseOne(file)
	backRepo.BackRepoStaff_tuning.RestoreXLPhaseOne(file)
	backRepo.BackRepoStem.RestoreXLPhaseOne(file)
	backRepo.BackRepoStick.RestoreXLPhaseOne(file)
	backRepo.BackRepoString_mute.RestoreXLPhaseOne(file)
	backRepo.BackRepoStrong_accent.RestoreXLPhaseOne(file)
	backRepo.BackRepoSupports.RestoreXLPhaseOne(file)
	backRepo.BackRepoSwing.RestoreXLPhaseOne(file)
	backRepo.BackRepoSync.RestoreXLPhaseOne(file)
	backRepo.BackRepoSystem_dividers.RestoreXLPhaseOne(file)
	backRepo.BackRepoSystem_layout.RestoreXLPhaseOne(file)
	backRepo.BackRepoSystem_margins.RestoreXLPhaseOne(file)
	backRepo.BackRepoTap.RestoreXLPhaseOne(file)
	backRepo.BackRepoTechnical.RestoreXLPhaseOne(file)
	backRepo.BackRepoText_element_data.RestoreXLPhaseOne(file)
	backRepo.BackRepoTie.RestoreXLPhaseOne(file)
	backRepo.BackRepoTied.RestoreXLPhaseOne(file)
	backRepo.BackRepoTime.RestoreXLPhaseOne(file)
	backRepo.BackRepoTime_modification.RestoreXLPhaseOne(file)
	backRepo.BackRepoTimpani.RestoreXLPhaseOne(file)
	backRepo.BackRepoTranspose.RestoreXLPhaseOne(file)
	backRepo.BackRepoTremolo.RestoreXLPhaseOne(file)
	backRepo.BackRepoTuplet.RestoreXLPhaseOne(file)
	backRepo.BackRepoTuplet_dot.RestoreXLPhaseOne(file)
	backRepo.BackRepoTuplet_number.RestoreXLPhaseOne(file)
	backRepo.BackRepoTuplet_portion.RestoreXLPhaseOne(file)
	backRepo.BackRepoTuplet_type.RestoreXLPhaseOne(file)
	backRepo.BackRepoTyped_text.RestoreXLPhaseOne(file)
	backRepo.BackRepoUnpitched.RestoreXLPhaseOne(file)
	backRepo.BackRepoVirtual_instrument.RestoreXLPhaseOne(file)
	backRepo.BackRepoWait.RestoreXLPhaseOne(file)
	backRepo.BackRepoWavy_line.RestoreXLPhaseOne(file)
	backRepo.BackRepoWedge.RestoreXLPhaseOne(file)
	backRepo.BackRepoWood.RestoreXLPhaseOne(file)
	backRepo.BackRepoWork.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
