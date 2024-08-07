// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongmusicxml/go/models"
	"github.com/fullstack-lang/gongmusicxml/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Accidental:
		fillUpTable[models.Accidental](probe)
	case *models.Accidental_mark:
		fillUpTable[models.Accidental_mark](probe)
	case *models.Accidental_text:
		fillUpTable[models.Accidental_text](probe)
	case *models.Accord:
		fillUpTable[models.Accord](probe)
	case *models.Accordion_registration:
		fillUpTable[models.Accordion_registration](probe)
	case *models.AnyType:
		fillUpTable[models.AnyType](probe)
	case *models.Appearance:
		fillUpTable[models.Appearance](probe)
	case *models.Arpeggiate:
		fillUpTable[models.Arpeggiate](probe)
	case *models.Arrow:
		fillUpTable[models.Arrow](probe)
	case *models.Articulations:
		fillUpTable[models.Articulations](probe)
	case *models.Assess:
		fillUpTable[models.Assess](probe)
	case *models.Attributes:
		fillUpTable[models.Attributes](probe)
	case *models.Backup:
		fillUpTable[models.Backup](probe)
	case *models.Bar_style_color:
		fillUpTable[models.Bar_style_color](probe)
	case *models.Barline:
		fillUpTable[models.Barline](probe)
	case *models.Barre:
		fillUpTable[models.Barre](probe)
	case *models.Bass:
		fillUpTable[models.Bass](probe)
	case *models.Bass_step:
		fillUpTable[models.Bass_step](probe)
	case *models.Beam:
		fillUpTable[models.Beam](probe)
	case *models.Beat_repeat:
		fillUpTable[models.Beat_repeat](probe)
	case *models.Beat_unit_tied:
		fillUpTable[models.Beat_unit_tied](probe)
	case *models.Beater:
		fillUpTable[models.Beater](probe)
	case *models.Bend:
		fillUpTable[models.Bend](probe)
	case *models.Bookmark:
		fillUpTable[models.Bookmark](probe)
	case *models.Bracket:
		fillUpTable[models.Bracket](probe)
	case *models.Breath_mark:
		fillUpTable[models.Breath_mark](probe)
	case *models.Caesura:
		fillUpTable[models.Caesura](probe)
	case *models.Cancel:
		fillUpTable[models.Cancel](probe)
	case *models.Clef:
		fillUpTable[models.Clef](probe)
	case *models.Coda:
		fillUpTable[models.Coda](probe)
	case *models.Credit:
		fillUpTable[models.Credit](probe)
	case *models.Dashes:
		fillUpTable[models.Dashes](probe)
	case *models.Defaults:
		fillUpTable[models.Defaults](probe)
	case *models.Degree:
		fillUpTable[models.Degree](probe)
	case *models.Degree_alter:
		fillUpTable[models.Degree_alter](probe)
	case *models.Degree_type:
		fillUpTable[models.Degree_type](probe)
	case *models.Degree_value:
		fillUpTable[models.Degree_value](probe)
	case *models.Direction:
		fillUpTable[models.Direction](probe)
	case *models.Direction_type:
		fillUpTable[models.Direction_type](probe)
	case *models.Distance:
		fillUpTable[models.Distance](probe)
	case *models.Double:
		fillUpTable[models.Double](probe)
	case *models.Dynamics:
		fillUpTable[models.Dynamics](probe)
	case *models.Effect:
		fillUpTable[models.Effect](probe)
	case *models.Elision:
		fillUpTable[models.Elision](probe)
	case *models.Empty:
		fillUpTable[models.Empty](probe)
	case *models.Empty_font:
		fillUpTable[models.Empty_font](probe)
	case *models.Empty_line:
		fillUpTable[models.Empty_line](probe)
	case *models.Empty_placement:
		fillUpTable[models.Empty_placement](probe)
	case *models.Empty_placement_smufl:
		fillUpTable[models.Empty_placement_smufl](probe)
	case *models.Empty_print_object_style_align:
		fillUpTable[models.Empty_print_object_style_align](probe)
	case *models.Empty_print_style:
		fillUpTable[models.Empty_print_style](probe)
	case *models.Empty_print_style_align:
		fillUpTable[models.Empty_print_style_align](probe)
	case *models.Empty_print_style_align_id:
		fillUpTable[models.Empty_print_style_align_id](probe)
	case *models.Empty_trill_sound:
		fillUpTable[models.Empty_trill_sound](probe)
	case *models.Encoding:
		fillUpTable[models.Encoding](probe)
	case *models.Ending:
		fillUpTable[models.Ending](probe)
	case *models.Extend:
		fillUpTable[models.Extend](probe)
	case *models.Feature:
		fillUpTable[models.Feature](probe)
	case *models.Fermata:
		fillUpTable[models.Fermata](probe)
	case *models.Figure:
		fillUpTable[models.Figure](probe)
	case *models.Figured_bass:
		fillUpTable[models.Figured_bass](probe)
	case *models.Fingering:
		fillUpTable[models.Fingering](probe)
	case *models.First_fret:
		fillUpTable[models.First_fret](probe)
	case *models.Foo:
		fillUpTable[models.Foo](probe)
	case *models.For_part:
		fillUpTable[models.For_part](probe)
	case *models.Formatted_symbol:
		fillUpTable[models.Formatted_symbol](probe)
	case *models.Formatted_symbol_id:
		fillUpTable[models.Formatted_symbol_id](probe)
	case *models.Forward:
		fillUpTable[models.Forward](probe)
	case *models.Frame:
		fillUpTable[models.Frame](probe)
	case *models.Frame_note:
		fillUpTable[models.Frame_note](probe)
	case *models.Fret:
		fillUpTable[models.Fret](probe)
	case *models.Glass:
		fillUpTable[models.Glass](probe)
	case *models.Glissando:
		fillUpTable[models.Glissando](probe)
	case *models.Glyph:
		fillUpTable[models.Glyph](probe)
	case *models.Grace:
		fillUpTable[models.Grace](probe)
	case *models.Group_barline:
		fillUpTable[models.Group_barline](probe)
	case *models.Group_symbol:
		fillUpTable[models.Group_symbol](probe)
	case *models.Grouping:
		fillUpTable[models.Grouping](probe)
	case *models.Hammer_on_pull_off:
		fillUpTable[models.Hammer_on_pull_off](probe)
	case *models.Handbell:
		fillUpTable[models.Handbell](probe)
	case *models.Harmon_closed:
		fillUpTable[models.Harmon_closed](probe)
	case *models.Harmon_mute:
		fillUpTable[models.Harmon_mute](probe)
	case *models.Harmonic:
		fillUpTable[models.Harmonic](probe)
	case *models.Harmony:
		fillUpTable[models.Harmony](probe)
	case *models.Harmony_alter:
		fillUpTable[models.Harmony_alter](probe)
	case *models.Harp_pedals:
		fillUpTable[models.Harp_pedals](probe)
	case *models.Heel_toe:
		fillUpTable[models.Heel_toe](probe)
	case *models.Hole:
		fillUpTable[models.Hole](probe)
	case *models.Hole_closed:
		fillUpTable[models.Hole_closed](probe)
	case *models.Horizontal_turn:
		fillUpTable[models.Horizontal_turn](probe)
	case *models.Identification:
		fillUpTable[models.Identification](probe)
	case *models.Image:
		fillUpTable[models.Image](probe)
	case *models.Instrument:
		fillUpTable[models.Instrument](probe)
	case *models.Instrument_change:
		fillUpTable[models.Instrument_change](probe)
	case *models.Instrument_link:
		fillUpTable[models.Instrument_link](probe)
	case *models.Interchangeable:
		fillUpTable[models.Interchangeable](probe)
	case *models.Inversion:
		fillUpTable[models.Inversion](probe)
	case *models.Key:
		fillUpTable[models.Key](probe)
	case *models.Key_accidental:
		fillUpTable[models.Key_accidental](probe)
	case *models.Key_octave:
		fillUpTable[models.Key_octave](probe)
	case *models.Kind:
		fillUpTable[models.Kind](probe)
	case *models.Level:
		fillUpTable[models.Level](probe)
	case *models.Line_detail:
		fillUpTable[models.Line_detail](probe)
	case *models.Line_width:
		fillUpTable[models.Line_width](probe)
	case *models.Link:
		fillUpTable[models.Link](probe)
	case *models.Listen:
		fillUpTable[models.Listen](probe)
	case *models.Listening:
		fillUpTable[models.Listening](probe)
	case *models.Lyric:
		fillUpTable[models.Lyric](probe)
	case *models.Lyric_font:
		fillUpTable[models.Lyric_font](probe)
	case *models.Lyric_language:
		fillUpTable[models.Lyric_language](probe)
	case *models.Measure_layout:
		fillUpTable[models.Measure_layout](probe)
	case *models.Measure_numbering:
		fillUpTable[models.Measure_numbering](probe)
	case *models.Measure_repeat:
		fillUpTable[models.Measure_repeat](probe)
	case *models.Measure_style:
		fillUpTable[models.Measure_style](probe)
	case *models.Membrane:
		fillUpTable[models.Membrane](probe)
	case *models.Metal:
		fillUpTable[models.Metal](probe)
	case *models.Metronome:
		fillUpTable[models.Metronome](probe)
	case *models.Metronome_beam:
		fillUpTable[models.Metronome_beam](probe)
	case *models.Metronome_note:
		fillUpTable[models.Metronome_note](probe)
	case *models.Metronome_tied:
		fillUpTable[models.Metronome_tied](probe)
	case *models.Metronome_tuplet:
		fillUpTable[models.Metronome_tuplet](probe)
	case *models.Midi_device:
		fillUpTable[models.Midi_device](probe)
	case *models.Midi_instrument:
		fillUpTable[models.Midi_instrument](probe)
	case *models.Miscellaneous:
		fillUpTable[models.Miscellaneous](probe)
	case *models.Miscellaneous_field:
		fillUpTable[models.Miscellaneous_field](probe)
	case *models.Mordent:
		fillUpTable[models.Mordent](probe)
	case *models.Multiple_rest:
		fillUpTable[models.Multiple_rest](probe)
	case *models.Name_display:
		fillUpTable[models.Name_display](probe)
	case *models.Non_arpeggiate:
		fillUpTable[models.Non_arpeggiate](probe)
	case *models.Notations:
		fillUpTable[models.Notations](probe)
	case *models.Note:
		fillUpTable[models.Note](probe)
	case *models.Note_size:
		fillUpTable[models.Note_size](probe)
	case *models.Note_type:
		fillUpTable[models.Note_type](probe)
	case *models.Notehead:
		fillUpTable[models.Notehead](probe)
	case *models.Notehead_text:
		fillUpTable[models.Notehead_text](probe)
	case *models.Numeral:
		fillUpTable[models.Numeral](probe)
	case *models.Numeral_key:
		fillUpTable[models.Numeral_key](probe)
	case *models.Numeral_root:
		fillUpTable[models.Numeral_root](probe)
	case *models.Octave_shift:
		fillUpTable[models.Octave_shift](probe)
	case *models.Offset:
		fillUpTable[models.Offset](probe)
	case *models.Opus:
		fillUpTable[models.Opus](probe)
	case *models.Ornaments:
		fillUpTable[models.Ornaments](probe)
	case *models.Other_appearance:
		fillUpTable[models.Other_appearance](probe)
	case *models.Other_listening:
		fillUpTable[models.Other_listening](probe)
	case *models.Other_notation:
		fillUpTable[models.Other_notation](probe)
	case *models.Other_play:
		fillUpTable[models.Other_play](probe)
	case *models.Page_layout:
		fillUpTable[models.Page_layout](probe)
	case *models.Page_margins:
		fillUpTable[models.Page_margins](probe)
	case *models.Part_clef:
		fillUpTable[models.Part_clef](probe)
	case *models.Part_group:
		fillUpTable[models.Part_group](probe)
	case *models.Part_link:
		fillUpTable[models.Part_link](probe)
	case *models.Part_list:
		fillUpTable[models.Part_list](probe)
	case *models.Part_symbol:
		fillUpTable[models.Part_symbol](probe)
	case *models.Part_transpose:
		fillUpTable[models.Part_transpose](probe)
	case *models.Pedal:
		fillUpTable[models.Pedal](probe)
	case *models.Pedal_tuning:
		fillUpTable[models.Pedal_tuning](probe)
	case *models.Percussion:
		fillUpTable[models.Percussion](probe)
	case *models.Pitch:
		fillUpTable[models.Pitch](probe)
	case *models.Pitched:
		fillUpTable[models.Pitched](probe)
	case *models.Play:
		fillUpTable[models.Play](probe)
	case *models.Player:
		fillUpTable[models.Player](probe)
	case *models.Principal_voice:
		fillUpTable[models.Principal_voice](probe)
	case *models.Print:
		fillUpTable[models.Print](probe)
	case *models.Release:
		fillUpTable[models.Release](probe)
	case *models.Repeat:
		fillUpTable[models.Repeat](probe)
	case *models.Rest:
		fillUpTable[models.Rest](probe)
	case *models.Root:
		fillUpTable[models.Root](probe)
	case *models.Root_step:
		fillUpTable[models.Root_step](probe)
	case *models.Scaling:
		fillUpTable[models.Scaling](probe)
	case *models.Scordatura:
		fillUpTable[models.Scordatura](probe)
	case *models.Score_instrument:
		fillUpTable[models.Score_instrument](probe)
	case *models.Score_part:
		fillUpTable[models.Score_part](probe)
	case *models.Score_partwise:
		fillUpTable[models.Score_partwise](probe)
	case *models.Score_timewise:
		fillUpTable[models.Score_timewise](probe)
	case *models.Segno:
		fillUpTable[models.Segno](probe)
	case *models.Slash:
		fillUpTable[models.Slash](probe)
	case *models.Slide:
		fillUpTable[models.Slide](probe)
	case *models.Slur:
		fillUpTable[models.Slur](probe)
	case *models.Sound:
		fillUpTable[models.Sound](probe)
	case *models.Staff_details:
		fillUpTable[models.Staff_details](probe)
	case *models.Staff_divide:
		fillUpTable[models.Staff_divide](probe)
	case *models.Staff_layout:
		fillUpTable[models.Staff_layout](probe)
	case *models.Staff_size:
		fillUpTable[models.Staff_size](probe)
	case *models.Staff_tuning:
		fillUpTable[models.Staff_tuning](probe)
	case *models.Stem:
		fillUpTable[models.Stem](probe)
	case *models.Stick:
		fillUpTable[models.Stick](probe)
	case *models.String_mute:
		fillUpTable[models.String_mute](probe)
	case *models.Strong_accent:
		fillUpTable[models.Strong_accent](probe)
	case *models.Supports:
		fillUpTable[models.Supports](probe)
	case *models.Swing:
		fillUpTable[models.Swing](probe)
	case *models.Sync:
		fillUpTable[models.Sync](probe)
	case *models.System_dividers:
		fillUpTable[models.System_dividers](probe)
	case *models.System_layout:
		fillUpTable[models.System_layout](probe)
	case *models.System_margins:
		fillUpTable[models.System_margins](probe)
	case *models.Tap:
		fillUpTable[models.Tap](probe)
	case *models.Technical:
		fillUpTable[models.Technical](probe)
	case *models.Text_element_data:
		fillUpTable[models.Text_element_data](probe)
	case *models.Tie:
		fillUpTable[models.Tie](probe)
	case *models.Tied:
		fillUpTable[models.Tied](probe)
	case *models.Time:
		fillUpTable[models.Time](probe)
	case *models.Time_modification:
		fillUpTable[models.Time_modification](probe)
	case *models.Timpani:
		fillUpTable[models.Timpani](probe)
	case *models.Transpose:
		fillUpTable[models.Transpose](probe)
	case *models.Tremolo:
		fillUpTable[models.Tremolo](probe)
	case *models.Tuplet:
		fillUpTable[models.Tuplet](probe)
	case *models.Tuplet_dot:
		fillUpTable[models.Tuplet_dot](probe)
	case *models.Tuplet_number:
		fillUpTable[models.Tuplet_number](probe)
	case *models.Tuplet_portion:
		fillUpTable[models.Tuplet_portion](probe)
	case *models.Tuplet_type:
		fillUpTable[models.Tuplet_type](probe)
	case *models.Typed_text:
		fillUpTable[models.Typed_text](probe)
	case *models.Unpitched:
		fillUpTable[models.Unpitched](probe)
	case *models.Virtual_instrument:
		fillUpTable[models.Virtual_instrument](probe)
	case *models.Wait:
		fillUpTable[models.Wait](probe)
	case *models.Wavy_line:
		fillUpTable[models.Wavy_line](probe)
	case *models.Wedge:
		fillUpTable[models.Wedge](probe)
	case *models.Wood:
		fillUpTable[models.Wood](probe)
	case *models.Work:
		fillUpTable[models.Work](probe)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()
	probe.tableStage.Commit()

	table := new(gongtable.Table).Stage(probe.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			probe.stageOfInterest,
			probe.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
