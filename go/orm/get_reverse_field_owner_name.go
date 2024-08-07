// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Accidental:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Accidental_mark:
		switch reverseField.GongstructName {
		// insertion point
		case "Ornaments":
			switch reverseField.Fieldname {
			case "Accidental_mark":
				if _ornaments, ok := stage.Ornaments_Accidental_mark_reverseMap[inst]; ok {
					res = _ornaments.Name
				}
			}
		}

	case *models.Accidental_text:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Accord:
		switch reverseField.GongstructName {
		// insertion point
		case "Scordatura":
			switch reverseField.Fieldname {
			case "Accord":
				if _scordatura, ok := stage.Scordatura_Accord_reverseMap[inst]; ok {
					res = _scordatura.Name
				}
			}
		}

	case *models.Accordion_registration:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.AnyType:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Appearance:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Arpeggiate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Arrow:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Articulations:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Assess:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Attributes:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Backup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bar_style_color:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Barline:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Barre:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bass:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bass_step:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beam:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beat_repeat:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beat_unit_tied:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beater:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bend:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bookmark:
		switch reverseField.GongstructName {
		// insertion point
		case "Credit":
			switch reverseField.Fieldname {
			case "Bookmark":
				if _credit, ok := stage.Credit_Bookmark_reverseMap[inst]; ok {
					res = _credit.Name
				}
			}
		}

	case *models.Bracket:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Breath_mark:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Caesura:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Cancel:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Clef:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Clef":
				if _attributes, ok := stage.Attributes_Clef_reverseMap[inst]; ok {
					res = _attributes.Name
				}
			}
		}

	case *models.Coda:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Coda":
				if _direction_type, ok := stage.Direction_type_Coda_reverseMap[inst]; ok {
					res = _direction_type.Name
				}
			}
		}

	case *models.Credit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Dashes:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Defaults:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree_alter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree_value:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Direction:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Direction_type:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction":
			switch reverseField.Fieldname {
			case "Direction_type":
				if _direction, ok := stage.Direction_Direction_type_reverseMap[inst]; ok {
					res = _direction.Name
				}
			}
		}

	case *models.Distance:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Distance":
				if _appearance, ok := stage.Appearance_Distance_reverseMap[inst]; ok {
					res = _appearance.Name
				}
			}
		}

	case *models.Double:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Dynamics:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Dynamics":
				if _direction_type, ok := stage.Direction_type_Dynamics_reverseMap[inst]; ok {
					res = _direction_type.Name
				}
			}
		}

	case *models.Effect:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Elision:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty:
		switch reverseField.GongstructName {
		// insertion point
		case "Metronome_note":
			switch reverseField.Fieldname {
			case "Metronome_dot":
				if _metronome_note, ok := stage.Metronome_note_Metronome_dot_reverseMap[inst]; ok {
					res = _metronome_note.Name
				}
			}
		}

	case *models.Empty_font:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_line:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_placement:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Dot":
				if _note, ok := stage.Note_Dot_reverseMap[inst]; ok {
					res = _note.Name
				}
			}
		}

	case *models.Empty_placement_smufl:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_object_style_align:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_style:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_style_align:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_style_align_id:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_trill_sound:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Encoding:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Ending:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Extend:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Feature:
		switch reverseField.GongstructName {
		// insertion point
		case "Grouping":
			switch reverseField.Fieldname {
			case "Feature":
				if _grouping, ok := stage.Grouping_Feature_reverseMap[inst]; ok {
					res = _grouping.Name
				}
			}
		}

	case *models.Fermata:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Figure:
		switch reverseField.GongstructName {
		// insertion point
		case "Figured_bass":
			switch reverseField.Fieldname {
			case "Figure":
				if _figured_bass, ok := stage.Figured_bass_Figure_reverseMap[inst]; ok {
					res = _figured_bass.Name
				}
			}
		}

	case *models.Figured_bass:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Fingering:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.First_fret:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Foo:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.For_part:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "For_part":
				if _attributes, ok := stage.Attributes_For_part_reverseMap[inst]; ok {
					res = _attributes.Name
				}
			}
		}

	case *models.Formatted_symbol:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Formatted_symbol_id:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Forward:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Frame:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Frame_note:
		switch reverseField.GongstructName {
		// insertion point
		case "Frame":
			switch reverseField.Fieldname {
			case "Frame_note":
				if _frame, ok := stage.Frame_Frame_note_reverseMap[inst]; ok {
					res = _frame.Name
				}
			}
		}

	case *models.Fret:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Glass:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Glissando:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Glyph:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Glyph":
				if _appearance, ok := stage.Appearance_Glyph_reverseMap[inst]; ok {
					res = _appearance.Name
				}
			}
		}

	case *models.Grace:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group_barline:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group_symbol:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Grouping:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hammer_on_pull_off:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Handbell:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmon_closed:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmon_mute:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmonic:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmony:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmony_alter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harp_pedals:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Heel_toe:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hole:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hole_closed:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Horizontal_turn:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Identification:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Image:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Instrument:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Instrument":
				if _note, ok := stage.Note_Instrument_reverseMap[inst]; ok {
					res = _note.Name
				}
			}
		}

	case *models.Instrument_change:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Instrument_link:
		switch reverseField.GongstructName {
		// insertion point
		case "Part_link":
			switch reverseField.Fieldname {
			case "Instrument_link":
				if _part_link, ok := stage.Part_link_Instrument_link_reverseMap[inst]; ok {
					res = _part_link.Name
				}
			}
		}

	case *models.Interchangeable:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Inversion:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Key":
				if _attributes, ok := stage.Attributes_Key_reverseMap[inst]; ok {
					res = _attributes.Name
				}
			}
		}

	case *models.Key_accidental:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key_octave:
		switch reverseField.GongstructName {
		// insertion point
		case "Key":
			switch reverseField.Fieldname {
			case "Key_octave":
				if _key, ok := stage.Key_Key_octave_reverseMap[inst]; ok {
					res = _key.Name
				}
			}
		}

	case *models.Kind:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Level:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Line_detail:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Line_width:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Line_width":
				if _appearance, ok := stage.Appearance_Line_width_reverseMap[inst]; ok {
					res = _appearance.Name
				}
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Credit":
			switch reverseField.Fieldname {
			case "Link":
				if _credit, ok := stage.Credit_Link_reverseMap[inst]; ok {
					res = _credit.Name
				}
			}
		}

	case *models.Listen:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Listening:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Lyric":
				if _note, ok := stage.Note_Lyric_reverseMap[inst]; ok {
					res = _note.Name
				}
			}
		}

	case *models.Lyric_font:
		switch reverseField.GongstructName {
		// insertion point
		case "Defaults":
			switch reverseField.Fieldname {
			case "Lyric_font":
				if _defaults, ok := stage.Defaults_Lyric_font_reverseMap[inst]; ok {
					res = _defaults.Name
				}
			}
		}

	case *models.Lyric_language:
		switch reverseField.GongstructName {
		// insertion point
		case "Defaults":
			switch reverseField.Fieldname {
			case "Lyric_language":
				if _defaults, ok := stage.Defaults_Lyric_language_reverseMap[inst]; ok {
					res = _defaults.Name
				}
			}
		}

	case *models.Measure_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Measure_numbering:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Measure_repeat:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Measure_style:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Measure_style":
				if _attributes, ok := stage.Attributes_Measure_style_reverseMap[inst]; ok {
					res = _attributes.Name
				}
			}
		}

	case *models.Membrane:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metal:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome_beam:
		switch reverseField.GongstructName {
		// insertion point
		case "Metronome_note":
			switch reverseField.Fieldname {
			case "Metronome_beam":
				if _metronome_note, ok := stage.Metronome_note_Metronome_beam_reverseMap[inst]; ok {
					res = _metronome_note.Name
				}
			}
		}

	case *models.Metronome_note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome_tied:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome_tuplet:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Midi_device:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Midi_instrument:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Miscellaneous:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Miscellaneous_field:
		switch reverseField.GongstructName {
		// insertion point
		case "Miscellaneous":
			switch reverseField.Fieldname {
			case "Miscellaneous_field":
				if _miscellaneous, ok := stage.Miscellaneous_Miscellaneous_field_reverseMap[inst]; ok {
					res = _miscellaneous.Name
				}
			}
		}

	case *models.Mordent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Multiple_rest:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Name_display:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Non_arpeggiate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Notations:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Notations":
				if _note, ok := stage.Note_Notations_reverseMap[inst]; ok {
					res = _note.Name
				}
			}
		}

	case *models.Note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Note_size:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Note_size":
				if _appearance, ok := stage.Appearance_Note_size_reverseMap[inst]; ok {
					res = _appearance.Name
				}
			}
		}

	case *models.Note_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Notehead:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Notehead_text:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Numeral:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Numeral_key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Numeral_root:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Octave_shift:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Offset:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Opus:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Ornaments:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Other_appearance:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Other_appearance":
				if _appearance, ok := stage.Appearance_Other_appearance_reverseMap[inst]; ok {
					res = _appearance.Name
				}
			}
		}

	case *models.Other_listening:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Other_notation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Other_play:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Page_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Page_margins:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_clef:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_group:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_link:
		switch reverseField.GongstructName {
		// insertion point
		case "Score_part":
			switch reverseField.Fieldname {
			case "Part_link":
				if _score_part, ok := stage.Score_part_Part_link_reverseMap[inst]; ok {
					res = _score_part.Name
				}
			}
		}

	case *models.Part_list:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_symbol:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_transpose:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pedal:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pedal_tuning:
		switch reverseField.GongstructName {
		// insertion point
		case "Harp_pedals":
			switch reverseField.Fieldname {
			case "Pedal_tuning":
				if _harp_pedals, ok := stage.Harp_pedals_Pedal_tuning_reverseMap[inst]; ok {
					res = _harp_pedals.Name
				}
			}
		}

	case *models.Percussion:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Percussion":
				if _direction_type, ok := stage.Direction_type_Percussion_reverseMap[inst]; ok {
					res = _direction_type.Name
				}
			}
		}

	case *models.Pitch:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pitched:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Play:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Player:
		switch reverseField.GongstructName {
		// insertion point
		case "Score_part":
			switch reverseField.Fieldname {
			case "Player":
				if _score_part, ok := stage.Score_part_Player_reverseMap[inst]; ok {
					res = _score_part.Name
				}
			}
		}

	case *models.Principal_voice:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Print:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Release:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Repeat:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Rest:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Root:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Root_step:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scaling:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scordatura:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Score_instrument:
		switch reverseField.GongstructName {
		// insertion point
		case "Score_part":
			switch reverseField.Fieldname {
			case "Score_instrument":
				if _score_part, ok := stage.Score_part_Score_instrument_reverseMap[inst]; ok {
					res = _score_part.Name
				}
			}
		}

	case *models.Score_part:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Score_partwise:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Score_timewise:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Segno:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Segno":
				if _direction_type, ok := stage.Direction_type_Segno_reverseMap[inst]; ok {
					res = _direction_type.Name
				}
			}
		}

	case *models.Slash:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Slide:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Slur:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Sound:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_details:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Staff_details":
				if _attributes, ok := stage.Attributes_Staff_details_reverseMap[inst]; ok {
					res = _attributes.Name
				}
			}
		}

	case *models.Staff_divide:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_size:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_tuning:
		switch reverseField.GongstructName {
		// insertion point
		case "Staff_details":
			switch reverseField.Fieldname {
			case "Staff_tuning":
				if _staff_details, ok := stage.Staff_details_Staff_tuning_reverseMap[inst]; ok {
					res = _staff_details.Name
				}
			}
		}

	case *models.Stem:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Stick:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.String_mute:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Strong_accent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Supports:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Swing:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Sync:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.System_dividers:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.System_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.System_margins:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tap:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Technical:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Text_element_data:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tie:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tied:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Time:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Time_modification:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Timpani:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Transpose:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Transpose":
				if _attributes, ok := stage.Attributes_Transpose_reverseMap[inst]; ok {
					res = _attributes.Name
				}
			}
		}

	case *models.Tremolo:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet_dot:
		switch reverseField.GongstructName {
		// insertion point
		case "Tuplet_portion":
			switch reverseField.Fieldname {
			case "Tuplet_dot":
				if _tuplet_portion, ok := stage.Tuplet_portion_Tuplet_dot_reverseMap[inst]; ok {
					res = _tuplet_portion.Name
				}
			}
		}

	case *models.Tuplet_number:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet_portion:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Typed_text:
		switch reverseField.GongstructName {
		// insertion point
		case "Identification":
			switch reverseField.Fieldname {
			case "Creator":
				if _identification, ok := stage.Identification_Creator_reverseMap[inst]; ok {
					res = _identification.Name
				}
			case "Rights":
				if _identification, ok := stage.Identification_Rights_reverseMap[inst]; ok {
					res = _identification.Name
				}
			case "Relation":
				if _identification, ok := stage.Identification_Relation_reverseMap[inst]; ok {
					res = _identification.Name
				}
			}
		}

	case *models.Unpitched:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Virtual_instrument:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wait:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wavy_line:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wedge:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wood:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Work:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Accidental:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Accidental_mark:
		switch reverseField.GongstructName {
		// insertion point
		case "Ornaments":
			switch reverseField.Fieldname {
			case "Accidental_mark":
				res = stage.Ornaments_Accidental_mark_reverseMap[inst]
			}
		}

	case *models.Accidental_text:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Accord:
		switch reverseField.GongstructName {
		// insertion point
		case "Scordatura":
			switch reverseField.Fieldname {
			case "Accord":
				res = stage.Scordatura_Accord_reverseMap[inst]
			}
		}

	case *models.Accordion_registration:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.AnyType:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Appearance:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Arpeggiate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Arrow:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Articulations:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Assess:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Attributes:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Backup:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bar_style_color:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Barline:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Barre:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bass:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bass_step:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beam:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beat_repeat:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beat_unit_tied:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Beater:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bend:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Bookmark:
		switch reverseField.GongstructName {
		// insertion point
		case "Credit":
			switch reverseField.Fieldname {
			case "Bookmark":
				res = stage.Credit_Bookmark_reverseMap[inst]
			}
		}

	case *models.Bracket:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Breath_mark:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Caesura:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Cancel:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Clef:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Clef":
				res = stage.Attributes_Clef_reverseMap[inst]
			}
		}

	case *models.Coda:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Coda":
				res = stage.Direction_type_Coda_reverseMap[inst]
			}
		}

	case *models.Credit:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Dashes:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Defaults:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree_alter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Degree_value:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Direction:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Direction_type:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction":
			switch reverseField.Fieldname {
			case "Direction_type":
				res = stage.Direction_Direction_type_reverseMap[inst]
			}
		}

	case *models.Distance:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Distance":
				res = stage.Appearance_Distance_reverseMap[inst]
			}
		}

	case *models.Double:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Dynamics:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Dynamics":
				res = stage.Direction_type_Dynamics_reverseMap[inst]
			}
		}

	case *models.Effect:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Elision:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty:
		switch reverseField.GongstructName {
		// insertion point
		case "Metronome_note":
			switch reverseField.Fieldname {
			case "Metronome_dot":
				res = stage.Metronome_note_Metronome_dot_reverseMap[inst]
			}
		}

	case *models.Empty_font:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_line:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_placement:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Dot":
				res = stage.Note_Dot_reverseMap[inst]
			}
		}

	case *models.Empty_placement_smufl:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_object_style_align:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_style:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_style_align:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_print_style_align_id:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Empty_trill_sound:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Encoding:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Ending:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Extend:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Feature:
		switch reverseField.GongstructName {
		// insertion point
		case "Grouping":
			switch reverseField.Fieldname {
			case "Feature":
				res = stage.Grouping_Feature_reverseMap[inst]
			}
		}

	case *models.Fermata:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Figure:
		switch reverseField.GongstructName {
		// insertion point
		case "Figured_bass":
			switch reverseField.Fieldname {
			case "Figure":
				res = stage.Figured_bass_Figure_reverseMap[inst]
			}
		}

	case *models.Figured_bass:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Fingering:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.First_fret:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Foo:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.For_part:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "For_part":
				res = stage.Attributes_For_part_reverseMap[inst]
			}
		}

	case *models.Formatted_symbol:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Formatted_symbol_id:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Forward:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Frame:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Frame_note:
		switch reverseField.GongstructName {
		// insertion point
		case "Frame":
			switch reverseField.Fieldname {
			case "Frame_note":
				res = stage.Frame_Frame_note_reverseMap[inst]
			}
		}

	case *models.Fret:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Glass:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Glissando:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Glyph:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Glyph":
				res = stage.Appearance_Glyph_reverseMap[inst]
			}
		}

	case *models.Grace:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group_barline:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Group_symbol:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Grouping:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hammer_on_pull_off:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Handbell:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmon_closed:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmon_mute:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmonic:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmony:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harmony_alter:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Harp_pedals:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Heel_toe:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hole:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Hole_closed:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Horizontal_turn:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Identification:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Image:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Instrument:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Instrument":
				res = stage.Note_Instrument_reverseMap[inst]
			}
		}

	case *models.Instrument_change:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Instrument_link:
		switch reverseField.GongstructName {
		// insertion point
		case "Part_link":
			switch reverseField.Fieldname {
			case "Instrument_link":
				res = stage.Part_link_Instrument_link_reverseMap[inst]
			}
		}

	case *models.Interchangeable:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Inversion:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Key":
				res = stage.Attributes_Key_reverseMap[inst]
			}
		}

	case *models.Key_accidental:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Key_octave:
		switch reverseField.GongstructName {
		// insertion point
		case "Key":
			switch reverseField.Fieldname {
			case "Key_octave":
				res = stage.Key_Key_octave_reverseMap[inst]
			}
		}

	case *models.Kind:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Level:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Line_detail:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Line_width:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Line_width":
				res = stage.Appearance_Line_width_reverseMap[inst]
			}
		}

	case *models.Link:
		switch reverseField.GongstructName {
		// insertion point
		case "Credit":
			switch reverseField.Fieldname {
			case "Link":
				res = stage.Credit_Link_reverseMap[inst]
			}
		}

	case *models.Listen:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Listening:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Lyric:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Lyric":
				res = stage.Note_Lyric_reverseMap[inst]
			}
		}

	case *models.Lyric_font:
		switch reverseField.GongstructName {
		// insertion point
		case "Defaults":
			switch reverseField.Fieldname {
			case "Lyric_font":
				res = stage.Defaults_Lyric_font_reverseMap[inst]
			}
		}

	case *models.Lyric_language:
		switch reverseField.GongstructName {
		// insertion point
		case "Defaults":
			switch reverseField.Fieldname {
			case "Lyric_language":
				res = stage.Defaults_Lyric_language_reverseMap[inst]
			}
		}

	case *models.Measure_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Measure_numbering:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Measure_repeat:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Measure_style:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Measure_style":
				res = stage.Attributes_Measure_style_reverseMap[inst]
			}
		}

	case *models.Membrane:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metal:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome_beam:
		switch reverseField.GongstructName {
		// insertion point
		case "Metronome_note":
			switch reverseField.Fieldname {
			case "Metronome_beam":
				res = stage.Metronome_note_Metronome_beam_reverseMap[inst]
			}
		}

	case *models.Metronome_note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome_tied:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Metronome_tuplet:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Midi_device:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Midi_instrument:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Miscellaneous:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Miscellaneous_field:
		switch reverseField.GongstructName {
		// insertion point
		case "Miscellaneous":
			switch reverseField.Fieldname {
			case "Miscellaneous_field":
				res = stage.Miscellaneous_Miscellaneous_field_reverseMap[inst]
			}
		}

	case *models.Mordent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Multiple_rest:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Name_display:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Non_arpeggiate:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Notations:
		switch reverseField.GongstructName {
		// insertion point
		case "Note":
			switch reverseField.Fieldname {
			case "Notations":
				res = stage.Note_Notations_reverseMap[inst]
			}
		}

	case *models.Note:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Note_size:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Note_size":
				res = stage.Appearance_Note_size_reverseMap[inst]
			}
		}

	case *models.Note_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Notehead:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Notehead_text:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Numeral:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Numeral_key:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Numeral_root:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Octave_shift:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Offset:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Opus:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Ornaments:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Other_appearance:
		switch reverseField.GongstructName {
		// insertion point
		case "Appearance":
			switch reverseField.Fieldname {
			case "Other_appearance":
				res = stage.Appearance_Other_appearance_reverseMap[inst]
			}
		}

	case *models.Other_listening:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Other_notation:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Other_play:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Page_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Page_margins:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_clef:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_group:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_link:
		switch reverseField.GongstructName {
		// insertion point
		case "Score_part":
			switch reverseField.Fieldname {
			case "Part_link":
				res = stage.Score_part_Part_link_reverseMap[inst]
			}
		}

	case *models.Part_list:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_symbol:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Part_transpose:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pedal:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pedal_tuning:
		switch reverseField.GongstructName {
		// insertion point
		case "Harp_pedals":
			switch reverseField.Fieldname {
			case "Pedal_tuning":
				res = stage.Harp_pedals_Pedal_tuning_reverseMap[inst]
			}
		}

	case *models.Percussion:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Percussion":
				res = stage.Direction_type_Percussion_reverseMap[inst]
			}
		}

	case *models.Pitch:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Pitched:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Play:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Player:
		switch reverseField.GongstructName {
		// insertion point
		case "Score_part":
			switch reverseField.Fieldname {
			case "Player":
				res = stage.Score_part_Player_reverseMap[inst]
			}
		}

	case *models.Principal_voice:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Print:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Release:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Repeat:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Rest:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Root:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Root_step:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scaling:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Scordatura:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Score_instrument:
		switch reverseField.GongstructName {
		// insertion point
		case "Score_part":
			switch reverseField.Fieldname {
			case "Score_instrument":
				res = stage.Score_part_Score_instrument_reverseMap[inst]
			}
		}

	case *models.Score_part:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Score_partwise:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Score_timewise:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Segno:
		switch reverseField.GongstructName {
		// insertion point
		case "Direction_type":
			switch reverseField.Fieldname {
			case "Segno":
				res = stage.Direction_type_Segno_reverseMap[inst]
			}
		}

	case *models.Slash:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Slide:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Slur:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Sound:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_details:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Staff_details":
				res = stage.Attributes_Staff_details_reverseMap[inst]
			}
		}

	case *models.Staff_divide:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_size:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Staff_tuning:
		switch reverseField.GongstructName {
		// insertion point
		case "Staff_details":
			switch reverseField.Fieldname {
			case "Staff_tuning":
				res = stage.Staff_details_Staff_tuning_reverseMap[inst]
			}
		}

	case *models.Stem:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Stick:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.String_mute:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Strong_accent:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Supports:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Swing:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Sync:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.System_dividers:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.System_layout:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.System_margins:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tap:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Technical:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Text_element_data:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tie:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tied:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Time:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Time_modification:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Timpani:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Transpose:
		switch reverseField.GongstructName {
		// insertion point
		case "Attributes":
			switch reverseField.Fieldname {
			case "Transpose":
				res = stage.Attributes_Transpose_reverseMap[inst]
			}
		}

	case *models.Tremolo:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet_dot:
		switch reverseField.GongstructName {
		// insertion point
		case "Tuplet_portion":
			switch reverseField.Fieldname {
			case "Tuplet_dot":
				res = stage.Tuplet_portion_Tuplet_dot_reverseMap[inst]
			}
		}

	case *models.Tuplet_number:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet_portion:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Tuplet_type:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Typed_text:
		switch reverseField.GongstructName {
		// insertion point
		case "Identification":
			switch reverseField.Fieldname {
			case "Creator":
				res = stage.Identification_Creator_reverseMap[inst]
			case "Rights":
				res = stage.Identification_Rights_reverseMap[inst]
			case "Relation":
				res = stage.Identification_Relation_reverseMap[inst]
			}
		}

	case *models.Unpitched:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Virtual_instrument:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wait:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wavy_line:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wedge:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Wood:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Work:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
