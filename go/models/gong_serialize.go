// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *StageStruct, filename string) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelize[Accidental](stage, f)
		SerializeExcelize[Accidental_mark](stage, f)
		SerializeExcelize[Accidental_text](stage, f)
		SerializeExcelize[Accord](stage, f)
		SerializeExcelize[Accordion_registration](stage, f)
		SerializeExcelize[AnyType](stage, f)
		SerializeExcelize[Appearance](stage, f)
		SerializeExcelize[Arpeggiate](stage, f)
		SerializeExcelize[Arrow](stage, f)
		SerializeExcelize[Articulations](stage, f)
		SerializeExcelize[Assess](stage, f)
		SerializeExcelize[Attributes](stage, f)
		SerializeExcelize[Backup](stage, f)
		SerializeExcelize[Bar_style_color](stage, f)
		SerializeExcelize[Barline](stage, f)
		SerializeExcelize[Barre](stage, f)
		SerializeExcelize[Bass](stage, f)
		SerializeExcelize[Bass_step](stage, f)
		SerializeExcelize[Beam](stage, f)
		SerializeExcelize[Beat_repeat](stage, f)
		SerializeExcelize[Beat_unit_tied](stage, f)
		SerializeExcelize[Beater](stage, f)
		SerializeExcelize[Bend](stage, f)
		SerializeExcelize[Bookmark](stage, f)
		SerializeExcelize[Bracket](stage, f)
		SerializeExcelize[Breath_mark](stage, f)
		SerializeExcelize[Caesura](stage, f)
		SerializeExcelize[Cancel](stage, f)
		SerializeExcelize[Clef](stage, f)
		SerializeExcelize[Coda](stage, f)
		SerializeExcelize[Credit](stage, f)
		SerializeExcelize[Dashes](stage, f)
		SerializeExcelize[Defaults](stage, f)
		SerializeExcelize[Degree](stage, f)
		SerializeExcelize[Degree_alter](stage, f)
		SerializeExcelize[Degree_type](stage, f)
		SerializeExcelize[Degree_value](stage, f)
		SerializeExcelize[Direction](stage, f)
		SerializeExcelize[Direction_type](stage, f)
		SerializeExcelize[Distance](stage, f)
		SerializeExcelize[Double](stage, f)
		SerializeExcelize[Dynamics](stage, f)
		SerializeExcelize[Effect](stage, f)
		SerializeExcelize[Elision](stage, f)
		SerializeExcelize[Empty](stage, f)
		SerializeExcelize[Empty_font](stage, f)
		SerializeExcelize[Empty_line](stage, f)
		SerializeExcelize[Empty_placement](stage, f)
		SerializeExcelize[Empty_placement_smufl](stage, f)
		SerializeExcelize[Empty_print_object_style_align](stage, f)
		SerializeExcelize[Empty_print_style](stage, f)
		SerializeExcelize[Empty_print_style_align](stage, f)
		SerializeExcelize[Empty_print_style_align_id](stage, f)
		SerializeExcelize[Empty_trill_sound](stage, f)
		SerializeExcelize[Encoding](stage, f)
		SerializeExcelize[Ending](stage, f)
		SerializeExcelize[Extend](stage, f)
		SerializeExcelize[Feature](stage, f)
		SerializeExcelize[Fermata](stage, f)
		SerializeExcelize[Figure](stage, f)
		SerializeExcelize[Figured_bass](stage, f)
		SerializeExcelize[Fingering](stage, f)
		SerializeExcelize[First_fret](stage, f)
		SerializeExcelize[Foo](stage, f)
		SerializeExcelize[For_part](stage, f)
		SerializeExcelize[Formatted_symbol](stage, f)
		SerializeExcelize[Formatted_symbol_id](stage, f)
		SerializeExcelize[Forward](stage, f)
		SerializeExcelize[Frame](stage, f)
		SerializeExcelize[Frame_note](stage, f)
		SerializeExcelize[Fret](stage, f)
		SerializeExcelize[Glass](stage, f)
		SerializeExcelize[Glissando](stage, f)
		SerializeExcelize[Glyph](stage, f)
		SerializeExcelize[Grace](stage, f)
		SerializeExcelize[Group_barline](stage, f)
		SerializeExcelize[Group_symbol](stage, f)
		SerializeExcelize[Grouping](stage, f)
		SerializeExcelize[Hammer_on_pull_off](stage, f)
		SerializeExcelize[Handbell](stage, f)
		SerializeExcelize[Harmon_closed](stage, f)
		SerializeExcelize[Harmon_mute](stage, f)
		SerializeExcelize[Harmonic](stage, f)
		SerializeExcelize[Harmony](stage, f)
		SerializeExcelize[Harmony_alter](stage, f)
		SerializeExcelize[Harp_pedals](stage, f)
		SerializeExcelize[Heel_toe](stage, f)
		SerializeExcelize[Hole](stage, f)
		SerializeExcelize[Hole_closed](stage, f)
		SerializeExcelize[Horizontal_turn](stage, f)
		SerializeExcelize[Identification](stage, f)
		SerializeExcelize[Image](stage, f)
		SerializeExcelize[Instrument](stage, f)
		SerializeExcelize[Instrument_change](stage, f)
		SerializeExcelize[Instrument_link](stage, f)
		SerializeExcelize[Interchangeable](stage, f)
		SerializeExcelize[Inversion](stage, f)
		SerializeExcelize[Key](stage, f)
		SerializeExcelize[Key_accidental](stage, f)
		SerializeExcelize[Key_octave](stage, f)
		SerializeExcelize[Kind](stage, f)
		SerializeExcelize[Level](stage, f)
		SerializeExcelize[Line_detail](stage, f)
		SerializeExcelize[Line_width](stage, f)
		SerializeExcelize[Link](stage, f)
		SerializeExcelize[Listen](stage, f)
		SerializeExcelize[Listening](stage, f)
		SerializeExcelize[Lyric](stage, f)
		SerializeExcelize[Lyric_font](stage, f)
		SerializeExcelize[Lyric_language](stage, f)
		SerializeExcelize[Measure_layout](stage, f)
		SerializeExcelize[Measure_numbering](stage, f)
		SerializeExcelize[Measure_repeat](stage, f)
		SerializeExcelize[Measure_style](stage, f)
		SerializeExcelize[Membrane](stage, f)
		SerializeExcelize[Metal](stage, f)
		SerializeExcelize[Metronome](stage, f)
		SerializeExcelize[Metronome_beam](stage, f)
		SerializeExcelize[Metronome_note](stage, f)
		SerializeExcelize[Metronome_tied](stage, f)
		SerializeExcelize[Metronome_tuplet](stage, f)
		SerializeExcelize[Midi_device](stage, f)
		SerializeExcelize[Midi_instrument](stage, f)
		SerializeExcelize[Miscellaneous](stage, f)
		SerializeExcelize[Miscellaneous_field](stage, f)
		SerializeExcelize[Mordent](stage, f)
		SerializeExcelize[Multiple_rest](stage, f)
		SerializeExcelize[Name_display](stage, f)
		SerializeExcelize[Non_arpeggiate](stage, f)
		SerializeExcelize[Notations](stage, f)
		SerializeExcelize[Note](stage, f)
		SerializeExcelize[Note_size](stage, f)
		SerializeExcelize[Note_type](stage, f)
		SerializeExcelize[Notehead](stage, f)
		SerializeExcelize[Notehead_text](stage, f)
		SerializeExcelize[Numeral](stage, f)
		SerializeExcelize[Numeral_key](stage, f)
		SerializeExcelize[Numeral_root](stage, f)
		SerializeExcelize[Octave_shift](stage, f)
		SerializeExcelize[Offset](stage, f)
		SerializeExcelize[Opus](stage, f)
		SerializeExcelize[Ornaments](stage, f)
		SerializeExcelize[Other_appearance](stage, f)
		SerializeExcelize[Other_listening](stage, f)
		SerializeExcelize[Other_notation](stage, f)
		SerializeExcelize[Other_play](stage, f)
		SerializeExcelize[Page_layout](stage, f)
		SerializeExcelize[Page_margins](stage, f)
		SerializeExcelize[Part_clef](stage, f)
		SerializeExcelize[Part_group](stage, f)
		SerializeExcelize[Part_link](stage, f)
		SerializeExcelize[Part_list](stage, f)
		SerializeExcelize[Part_symbol](stage, f)
		SerializeExcelize[Part_transpose](stage, f)
		SerializeExcelize[Pedal](stage, f)
		SerializeExcelize[Pedal_tuning](stage, f)
		SerializeExcelize[Percussion](stage, f)
		SerializeExcelize[Pitch](stage, f)
		SerializeExcelize[Pitched](stage, f)
		SerializeExcelize[Play](stage, f)
		SerializeExcelize[Player](stage, f)
		SerializeExcelize[Principal_voice](stage, f)
		SerializeExcelize[Print](stage, f)
		SerializeExcelize[Release](stage, f)
		SerializeExcelize[Repeat](stage, f)
		SerializeExcelize[Rest](stage, f)
		SerializeExcelize[Root](stage, f)
		SerializeExcelize[Root_step](stage, f)
		SerializeExcelize[Scaling](stage, f)
		SerializeExcelize[Scordatura](stage, f)
		SerializeExcelize[Score_instrument](stage, f)
		SerializeExcelize[Score_part](stage, f)
		SerializeExcelize[Score_partwise](stage, f)
		SerializeExcelize[Score_timewise](stage, f)
		SerializeExcelize[Segno](stage, f)
		SerializeExcelize[Slash](stage, f)
		SerializeExcelize[Slide](stage, f)
		SerializeExcelize[Slur](stage, f)
		SerializeExcelize[Sound](stage, f)
		SerializeExcelize[Staff_details](stage, f)
		SerializeExcelize[Staff_divide](stage, f)
		SerializeExcelize[Staff_layout](stage, f)
		SerializeExcelize[Staff_size](stage, f)
		SerializeExcelize[Staff_tuning](stage, f)
		SerializeExcelize[Stem](stage, f)
		SerializeExcelize[Stick](stage, f)
		SerializeExcelize[String_mute](stage, f)
		SerializeExcelize[Strong_accent](stage, f)
		SerializeExcelize[Supports](stage, f)
		SerializeExcelize[Swing](stage, f)
		SerializeExcelize[Sync](stage, f)
		SerializeExcelize[System_dividers](stage, f)
		SerializeExcelize[System_layout](stage, f)
		SerializeExcelize[System_margins](stage, f)
		SerializeExcelize[Tap](stage, f)
		SerializeExcelize[Technical](stage, f)
		SerializeExcelize[Text_element_data](stage, f)
		SerializeExcelize[Tie](stage, f)
		SerializeExcelize[Tied](stage, f)
		SerializeExcelize[Time](stage, f)
		SerializeExcelize[Time_modification](stage, f)
		SerializeExcelize[Timpani](stage, f)
		SerializeExcelize[Transpose](stage, f)
		SerializeExcelize[Tremolo](stage, f)
		SerializeExcelize[Tuplet](stage, f)
		SerializeExcelize[Tuplet_dot](stage, f)
		SerializeExcelize[Tuplet_number](stage, f)
		SerializeExcelize[Tuplet_portion](stage, f)
		SerializeExcelize[Tuplet_type](stage, f)
		SerializeExcelize[Typed_text](stage, f)
		SerializeExcelize[Unpitched](stage, f)
		SerializeExcelize[Virtual_instrument](stage, f)
		SerializeExcelize[Wait](stage, f)
		SerializeExcelize[Wavy_line](stage, f)
		SerializeExcelize[Wedge](stage, f)
		SerializeExcelize[Wood](stage, f)
		SerializeExcelize[Work](stage, f)
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
		if err := f.SaveAs(filename); err != nil {
			fmt.Println("cannot write xl file : ", err)
		}
	}

}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

func Serialize[Type Gongstruct](stage *StageStruct, tab Tabulator) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	tab.AddSheet(sheetName)

	headerRowIndex := tab.AddRow(sheetName)
	for colIndex, fieldName := range GetFields[Type]() {
		tab.AddCell(sheetName, headerRowIndex, colIndex, fieldName)
		// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}

	set := *GetGongstructInstancesSet[Type](stage)
	for instance := range set {
		line := tab.AddRow(sheetName)
		for index, fieldName := range GetFields[Type]() {
			tab.AddCell(sheetName, line, index, GetFieldStringValue(
				any(*instance).(Type), fieldName))
			// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
			// 	any(*instance).(Type), fieldName))
		}
	}
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

func SerializeExcelize[Type Gongstruct](stage *StageStruct, f *excelize.File) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSet[Type](stage)
	line := 1

	for index, fieldName := range GetFields[Type]() {
		f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for instance := range set {
		line = line + 1
		for index, fieldName := range GetFields[Type]() {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
				any(*instance).(Type), fieldName))
		}
	}

	// Autofit all columns according to their text content
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Panicln("SerializeExcelize")
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			log.Panicln("SerializeExcelize")
		}
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
