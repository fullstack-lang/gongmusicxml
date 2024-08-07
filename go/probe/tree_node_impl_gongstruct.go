// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Accidental" {
		fillUpTable[models.Accidental](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Accidental_mark" {
		fillUpTable[models.Accidental_mark](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Accidental_text" {
		fillUpTable[models.Accidental_text](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Accord" {
		fillUpTable[models.Accord](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Accordion_registration" {
		fillUpTable[models.Accordion_registration](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AnyType" {
		fillUpTable[models.AnyType](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Appearance" {
		fillUpTable[models.Appearance](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Arpeggiate" {
		fillUpTable[models.Arpeggiate](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Arrow" {
		fillUpTable[models.Arrow](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Articulations" {
		fillUpTable[models.Articulations](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Assess" {
		fillUpTable[models.Assess](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Attributes" {
		fillUpTable[models.Attributes](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Backup" {
		fillUpTable[models.Backup](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bar_style_color" {
		fillUpTable[models.Bar_style_color](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Barline" {
		fillUpTable[models.Barline](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Barre" {
		fillUpTable[models.Barre](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bass" {
		fillUpTable[models.Bass](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bass_step" {
		fillUpTable[models.Bass_step](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Beam" {
		fillUpTable[models.Beam](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Beat_repeat" {
		fillUpTable[models.Beat_repeat](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Beat_unit_tied" {
		fillUpTable[models.Beat_unit_tied](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Beater" {
		fillUpTable[models.Beater](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bend" {
		fillUpTable[models.Bend](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bookmark" {
		fillUpTable[models.Bookmark](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bracket" {
		fillUpTable[models.Bracket](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Breath_mark" {
		fillUpTable[models.Breath_mark](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Caesura" {
		fillUpTable[models.Caesura](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Cancel" {
		fillUpTable[models.Cancel](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Clef" {
		fillUpTable[models.Clef](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Coda" {
		fillUpTable[models.Coda](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Credit" {
		fillUpTable[models.Credit](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Dashes" {
		fillUpTable[models.Dashes](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Defaults" {
		fillUpTable[models.Defaults](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Degree" {
		fillUpTable[models.Degree](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Degree_alter" {
		fillUpTable[models.Degree_alter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Degree_type" {
		fillUpTable[models.Degree_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Degree_value" {
		fillUpTable[models.Degree_value](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Direction" {
		fillUpTable[models.Direction](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Direction_type" {
		fillUpTable[models.Direction_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Distance" {
		fillUpTable[models.Distance](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Double" {
		fillUpTable[models.Double](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Dynamics" {
		fillUpTable[models.Dynamics](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Effect" {
		fillUpTable[models.Effect](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Elision" {
		fillUpTable[models.Elision](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty" {
		fillUpTable[models.Empty](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_font" {
		fillUpTable[models.Empty_font](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_line" {
		fillUpTable[models.Empty_line](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_placement" {
		fillUpTable[models.Empty_placement](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_placement_smufl" {
		fillUpTable[models.Empty_placement_smufl](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_print_object_style_align" {
		fillUpTable[models.Empty_print_object_style_align](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_print_style" {
		fillUpTable[models.Empty_print_style](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_print_style_align" {
		fillUpTable[models.Empty_print_style_align](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_print_style_align_id" {
		fillUpTable[models.Empty_print_style_align_id](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Empty_trill_sound" {
		fillUpTable[models.Empty_trill_sound](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Encoding" {
		fillUpTable[models.Encoding](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Ending" {
		fillUpTable[models.Ending](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Extend" {
		fillUpTable[models.Extend](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Feature" {
		fillUpTable[models.Feature](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Fermata" {
		fillUpTable[models.Fermata](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Figure" {
		fillUpTable[models.Figure](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Figured_bass" {
		fillUpTable[models.Figured_bass](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Fingering" {
		fillUpTable[models.Fingering](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "First_fret" {
		fillUpTable[models.First_fret](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Foo" {
		fillUpTable[models.Foo](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "For_part" {
		fillUpTable[models.For_part](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Formatted_symbol" {
		fillUpTable[models.Formatted_symbol](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Formatted_symbol_id" {
		fillUpTable[models.Formatted_symbol_id](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Forward" {
		fillUpTable[models.Forward](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Frame" {
		fillUpTable[models.Frame](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Frame_note" {
		fillUpTable[models.Frame_note](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Fret" {
		fillUpTable[models.Fret](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Glass" {
		fillUpTable[models.Glass](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Glissando" {
		fillUpTable[models.Glissando](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Glyph" {
		fillUpTable[models.Glyph](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Grace" {
		fillUpTable[models.Grace](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Group_barline" {
		fillUpTable[models.Group_barline](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Group_symbol" {
		fillUpTable[models.Group_symbol](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Grouping" {
		fillUpTable[models.Grouping](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Hammer_on_pull_off" {
		fillUpTable[models.Hammer_on_pull_off](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Handbell" {
		fillUpTable[models.Handbell](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Harmon_closed" {
		fillUpTable[models.Harmon_closed](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Harmon_mute" {
		fillUpTable[models.Harmon_mute](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Harmonic" {
		fillUpTable[models.Harmonic](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Harmony" {
		fillUpTable[models.Harmony](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Harmony_alter" {
		fillUpTable[models.Harmony_alter](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Harp_pedals" {
		fillUpTable[models.Harp_pedals](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Heel_toe" {
		fillUpTable[models.Heel_toe](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Hole" {
		fillUpTable[models.Hole](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Hole_closed" {
		fillUpTable[models.Hole_closed](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Horizontal_turn" {
		fillUpTable[models.Horizontal_turn](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Identification" {
		fillUpTable[models.Identification](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Image" {
		fillUpTable[models.Image](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Instrument" {
		fillUpTable[models.Instrument](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Instrument_change" {
		fillUpTable[models.Instrument_change](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Instrument_link" {
		fillUpTable[models.Instrument_link](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Interchangeable" {
		fillUpTable[models.Interchangeable](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Inversion" {
		fillUpTable[models.Inversion](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Key" {
		fillUpTable[models.Key](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Key_accidental" {
		fillUpTable[models.Key_accidental](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Key_octave" {
		fillUpTable[models.Key_octave](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Kind" {
		fillUpTable[models.Kind](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Level" {
		fillUpTable[models.Level](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Line_detail" {
		fillUpTable[models.Line_detail](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Line_width" {
		fillUpTable[models.Line_width](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Listen" {
		fillUpTable[models.Listen](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Listening" {
		fillUpTable[models.Listening](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Lyric" {
		fillUpTable[models.Lyric](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Lyric_font" {
		fillUpTable[models.Lyric_font](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Lyric_language" {
		fillUpTable[models.Lyric_language](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Measure_layout" {
		fillUpTable[models.Measure_layout](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Measure_numbering" {
		fillUpTable[models.Measure_numbering](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Measure_repeat" {
		fillUpTable[models.Measure_repeat](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Measure_style" {
		fillUpTable[models.Measure_style](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Membrane" {
		fillUpTable[models.Membrane](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Metal" {
		fillUpTable[models.Metal](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Metronome" {
		fillUpTable[models.Metronome](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Metronome_beam" {
		fillUpTable[models.Metronome_beam](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Metronome_note" {
		fillUpTable[models.Metronome_note](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Metronome_tied" {
		fillUpTable[models.Metronome_tied](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Metronome_tuplet" {
		fillUpTable[models.Metronome_tuplet](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Midi_device" {
		fillUpTable[models.Midi_device](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Midi_instrument" {
		fillUpTable[models.Midi_instrument](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Miscellaneous" {
		fillUpTable[models.Miscellaneous](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Miscellaneous_field" {
		fillUpTable[models.Miscellaneous_field](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Mordent" {
		fillUpTable[models.Mordent](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Multiple_rest" {
		fillUpTable[models.Multiple_rest](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Name_display" {
		fillUpTable[models.Name_display](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Non_arpeggiate" {
		fillUpTable[models.Non_arpeggiate](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Notations" {
		fillUpTable[models.Notations](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Note" {
		fillUpTable[models.Note](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Note_size" {
		fillUpTable[models.Note_size](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Note_type" {
		fillUpTable[models.Note_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Notehead" {
		fillUpTable[models.Notehead](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Notehead_text" {
		fillUpTable[models.Notehead_text](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Numeral" {
		fillUpTable[models.Numeral](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Numeral_key" {
		fillUpTable[models.Numeral_key](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Numeral_root" {
		fillUpTable[models.Numeral_root](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Octave_shift" {
		fillUpTable[models.Octave_shift](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Offset" {
		fillUpTable[models.Offset](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Opus" {
		fillUpTable[models.Opus](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Ornaments" {
		fillUpTable[models.Ornaments](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Other_appearance" {
		fillUpTable[models.Other_appearance](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Other_listening" {
		fillUpTable[models.Other_listening](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Other_notation" {
		fillUpTable[models.Other_notation](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Other_play" {
		fillUpTable[models.Other_play](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Page_layout" {
		fillUpTable[models.Page_layout](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Page_margins" {
		fillUpTable[models.Page_margins](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Part_clef" {
		fillUpTable[models.Part_clef](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Part_group" {
		fillUpTable[models.Part_group](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Part_link" {
		fillUpTable[models.Part_link](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Part_list" {
		fillUpTable[models.Part_list](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Part_symbol" {
		fillUpTable[models.Part_symbol](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Part_transpose" {
		fillUpTable[models.Part_transpose](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Pedal" {
		fillUpTable[models.Pedal](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Pedal_tuning" {
		fillUpTable[models.Pedal_tuning](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Percussion" {
		fillUpTable[models.Percussion](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Pitch" {
		fillUpTable[models.Pitch](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Pitched" {
		fillUpTable[models.Pitched](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Play" {
		fillUpTable[models.Play](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Player" {
		fillUpTable[models.Player](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Principal_voice" {
		fillUpTable[models.Principal_voice](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Print" {
		fillUpTable[models.Print](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Release" {
		fillUpTable[models.Release](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Repeat" {
		fillUpTable[models.Repeat](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rest" {
		fillUpTable[models.Rest](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Root" {
		fillUpTable[models.Root](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Root_step" {
		fillUpTable[models.Root_step](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Scaling" {
		fillUpTable[models.Scaling](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Scordatura" {
		fillUpTable[models.Scordatura](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Score_instrument" {
		fillUpTable[models.Score_instrument](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Score_part" {
		fillUpTable[models.Score_part](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Score_partwise" {
		fillUpTable[models.Score_partwise](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Score_timewise" {
		fillUpTable[models.Score_timewise](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Segno" {
		fillUpTable[models.Segno](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Slash" {
		fillUpTable[models.Slash](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Slide" {
		fillUpTable[models.Slide](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Slur" {
		fillUpTable[models.Slur](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Sound" {
		fillUpTable[models.Sound](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Staff_details" {
		fillUpTable[models.Staff_details](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Staff_divide" {
		fillUpTable[models.Staff_divide](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Staff_layout" {
		fillUpTable[models.Staff_layout](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Staff_size" {
		fillUpTable[models.Staff_size](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Staff_tuning" {
		fillUpTable[models.Staff_tuning](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Stem" {
		fillUpTable[models.Stem](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Stick" {
		fillUpTable[models.Stick](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "String_mute" {
		fillUpTable[models.String_mute](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Strong_accent" {
		fillUpTable[models.Strong_accent](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Supports" {
		fillUpTable[models.Supports](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Swing" {
		fillUpTable[models.Swing](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Sync" {
		fillUpTable[models.Sync](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "System_dividers" {
		fillUpTable[models.System_dividers](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "System_layout" {
		fillUpTable[models.System_layout](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "System_margins" {
		fillUpTable[models.System_margins](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tap" {
		fillUpTable[models.Tap](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Technical" {
		fillUpTable[models.Technical](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Text_element_data" {
		fillUpTable[models.Text_element_data](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tie" {
		fillUpTable[models.Tie](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tied" {
		fillUpTable[models.Tied](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Time" {
		fillUpTable[models.Time](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Time_modification" {
		fillUpTable[models.Time_modification](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Timpani" {
		fillUpTable[models.Timpani](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Transpose" {
		fillUpTable[models.Transpose](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tremolo" {
		fillUpTable[models.Tremolo](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tuplet" {
		fillUpTable[models.Tuplet](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tuplet_dot" {
		fillUpTable[models.Tuplet_dot](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tuplet_number" {
		fillUpTable[models.Tuplet_number](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tuplet_portion" {
		fillUpTable[models.Tuplet_portion](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Tuplet_type" {
		fillUpTable[models.Tuplet_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Typed_text" {
		fillUpTable[models.Typed_text](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Unpitched" {
		fillUpTable[models.Unpitched](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Virtual_instrument" {
		fillUpTable[models.Virtual_instrument](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Wait" {
		fillUpTable[models.Wait](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Wavy_line" {
		fillUpTable[models.Wavy_line](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Wedge" {
		fillUpTable[models.Wedge](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Wood" {
		fillUpTable[models.Wood](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Work" {
		fillUpTable[models.Work](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
