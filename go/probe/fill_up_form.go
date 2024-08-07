// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
	"github.com/fullstack-lang/gongmusicxml/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Accidental:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Accidental_mark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ornaments"
			rf.Fieldname = "Accidental_mark"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Ornaments),
					"Accidental_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Ornaments, *models.Accidental_mark](
					nil,
					"Accidental_mark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Accidental_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Accord:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Scordatura"
			rf.Fieldname = "Accord"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Scordatura),
					"Accord",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Scordatura, *models.Accord](
					nil,
					"Accord",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Accordion_registration:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Accordion_high", instanceWithInferedType.Accordion_high, formGroup, probe)
		AssociationFieldToForm("Accordion_low", instanceWithInferedType.Accordion_low, formGroup, probe)

	case *models.AnyType:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("InnerXML", instanceWithInferedType.InnerXML, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Appearance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Line_width", instanceWithInferedType, &instanceWithInferedType.Line_width, formGroup, probe)
		AssociationSliceToForm("Note_size", instanceWithInferedType, &instanceWithInferedType.Note_size, formGroup, probe)
		AssociationSliceToForm("Distance", instanceWithInferedType, &instanceWithInferedType.Distance, formGroup, probe)
		AssociationSliceToForm("Glyph", instanceWithInferedType, &instanceWithInferedType.Glyph, formGroup, probe)
		AssociationSliceToForm("Other_appearance", instanceWithInferedType, &instanceWithInferedType.Other_appearance, formGroup, probe)

	case *models.Arpeggiate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Arrow:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Articulations:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Accent", instanceWithInferedType.Accent, formGroup, probe)
		AssociationFieldToForm("Strong_accent", instanceWithInferedType.Strong_accent, formGroup, probe)
		AssociationFieldToForm("Staccato", instanceWithInferedType.Staccato, formGroup, probe)
		AssociationFieldToForm("Tenuto", instanceWithInferedType.Tenuto, formGroup, probe)
		AssociationFieldToForm("Detached_legato", instanceWithInferedType.Detached_legato, formGroup, probe)
		AssociationFieldToForm("Staccatissimo", instanceWithInferedType.Staccatissimo, formGroup, probe)
		AssociationFieldToForm("Spiccato", instanceWithInferedType.Spiccato, formGroup, probe)
		AssociationFieldToForm("Scoop", instanceWithInferedType.Scoop, formGroup, probe)
		AssociationFieldToForm("Plop", instanceWithInferedType.Plop, formGroup, probe)
		AssociationFieldToForm("Doit", instanceWithInferedType.Doit, formGroup, probe)
		AssociationFieldToForm("Falloff", instanceWithInferedType.Falloff, formGroup, probe)
		AssociationFieldToForm("Breath_mark", instanceWithInferedType.Breath_mark, formGroup, probe)
		AssociationFieldToForm("Caesura", instanceWithInferedType.Caesura, formGroup, probe)
		AssociationFieldToForm("Stress", instanceWithInferedType.Stress, formGroup, probe)
		AssociationFieldToForm("Unstress", instanceWithInferedType.Unstress, formGroup, probe)
		AssociationFieldToForm("Soft_accent", instanceWithInferedType.Soft_accent, formGroup, probe)

	case *models.Assess:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Attributes:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Key", instanceWithInferedType, &instanceWithInferedType.Key, formGroup, probe)
		AssociationFieldToForm("Part_symbol", instanceWithInferedType.Part_symbol, formGroup, probe)
		AssociationSliceToForm("Clef", instanceWithInferedType, &instanceWithInferedType.Clef, formGroup, probe)
		AssociationSliceToForm("Staff_details", instanceWithInferedType, &instanceWithInferedType.Staff_details, formGroup, probe)
		AssociationSliceToForm("Measure_style", instanceWithInferedType, &instanceWithInferedType.Measure_style, formGroup, probe)
		AssociationSliceToForm("Transpose", instanceWithInferedType, &instanceWithInferedType.Transpose, formGroup, probe)
		AssociationSliceToForm("For_part", instanceWithInferedType, &instanceWithInferedType.For_part, formGroup, probe)

	case *models.Backup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Bar_style_color:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Barline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bar_style", instanceWithInferedType.Bar_style, formGroup, probe)
		AssociationFieldToForm("Wavy_line", instanceWithInferedType.Wavy_line, formGroup, probe)
		AssociationFieldToForm("Fermata", instanceWithInferedType.Fermata, formGroup, probe)
		AssociationFieldToForm("Ending", instanceWithInferedType.Ending, formGroup, probe)
		AssociationFieldToForm("Repeat", instanceWithInferedType.Repeat, formGroup, probe)
		BasicFieldtoForm("Segno", instanceWithInferedType.Segno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Coda", instanceWithInferedType.Coda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Barre:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Bass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Bass_step", instanceWithInferedType.Bass_step, formGroup, probe)
		AssociationFieldToForm("Bass_alter", instanceWithInferedType.Bass_alter, formGroup, probe)

	case *models.Bass_step:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beam:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beat_repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beat_unit_tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Beater:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Bend:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Pre_bend", instanceWithInferedType.Pre_bend, formGroup, probe)
		AssociationFieldToForm("Release", instanceWithInferedType.Release, formGroup, probe)

	case *models.Bookmark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Bookmark"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Credit),
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Credit, *models.Bookmark](
					nil,
					"Bookmark",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Bracket:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Breath_mark:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Caesura:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Cancel:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Clef:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Clef"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Clef",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes, *models.Clef](
					nil,
					"Clef",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Coda:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Coda"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Coda",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type, *models.Coda](
					nil,
					"Coda",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Credit:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Link", instanceWithInferedType, &instanceWithInferedType.Link, formGroup, probe)
		AssociationSliceToForm("Bookmark", instanceWithInferedType, &instanceWithInferedType.Bookmark, formGroup, probe)
		AssociationFieldToForm("Credit_image", instanceWithInferedType.Credit_image, formGroup, probe)

	case *models.Dashes:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Defaults:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Scaling", instanceWithInferedType.Scaling, formGroup, probe)
		AssociationFieldToForm("Concert_score", instanceWithInferedType.Concert_score, formGroup, probe)
		AssociationFieldToForm("Appearance", instanceWithInferedType.Appearance, formGroup, probe)
		AssociationFieldToForm("Music_font", instanceWithInferedType.Music_font, formGroup, probe)
		AssociationFieldToForm("Word_font", instanceWithInferedType.Word_font, formGroup, probe)
		AssociationSliceToForm("Lyric_font", instanceWithInferedType, &instanceWithInferedType.Lyric_font, formGroup, probe)
		AssociationSliceToForm("Lyric_language", instanceWithInferedType, &instanceWithInferedType.Lyric_language, formGroup, probe)

	case *models.Degree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Degree_value", instanceWithInferedType.Degree_value, formGroup, probe)
		AssociationFieldToForm("Degree_alter", instanceWithInferedType.Degree_alter, formGroup, probe)
		AssociationFieldToForm("Degree_type", instanceWithInferedType.Degree_type, formGroup, probe)

	case *models.Degree_alter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Degree_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Degree_value:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Direction:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Direction_type", instanceWithInferedType, &instanceWithInferedType.Direction_type, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		AssociationFieldToForm("Sound", instanceWithInferedType.Sound, formGroup, probe)
		AssociationFieldToForm("Listening", instanceWithInferedType.Listening, formGroup, probe)

	case *models.Direction_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Segno", instanceWithInferedType, &instanceWithInferedType.Segno, formGroup, probe)
		AssociationSliceToForm("Coda", instanceWithInferedType, &instanceWithInferedType.Coda, formGroup, probe)
		AssociationFieldToForm("Wedge", instanceWithInferedType.Wedge, formGroup, probe)
		AssociationSliceToForm("Dynamics", instanceWithInferedType, &instanceWithInferedType.Dynamics, formGroup, probe)
		AssociationFieldToForm("Dashes", instanceWithInferedType.Dashes, formGroup, probe)
		AssociationFieldToForm("Bracket", instanceWithInferedType.Bracket, formGroup, probe)
		AssociationFieldToForm("Pedal", instanceWithInferedType.Pedal, formGroup, probe)
		AssociationFieldToForm("Metronome", instanceWithInferedType.Metronome, formGroup, probe)
		AssociationFieldToForm("Octave_shift", instanceWithInferedType.Octave_shift, formGroup, probe)
		AssociationFieldToForm("Harp_pedals", instanceWithInferedType.Harp_pedals, formGroup, probe)
		AssociationFieldToForm("Damp", instanceWithInferedType.Damp, formGroup, probe)
		AssociationFieldToForm("Damp_all", instanceWithInferedType.Damp_all, formGroup, probe)
		AssociationFieldToForm("Eyeglasses", instanceWithInferedType.Eyeglasses, formGroup, probe)
		AssociationFieldToForm("String_mute", instanceWithInferedType.String_mute, formGroup, probe)
		AssociationFieldToForm("Scordatura", instanceWithInferedType.Scordatura, formGroup, probe)
		AssociationFieldToForm("Image", instanceWithInferedType.Image, formGroup, probe)
		AssociationFieldToForm("Principal_voice", instanceWithInferedType.Principal_voice, formGroup, probe)
		AssociationSliceToForm("Percussion", instanceWithInferedType, &instanceWithInferedType.Percussion, formGroup, probe)
		AssociationFieldToForm("Accordion_registration", instanceWithInferedType.Accordion_registration, formGroup, probe)
		AssociationFieldToForm("Staff_divide", instanceWithInferedType.Staff_divide, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction"
			rf.Fieldname = "Direction_type"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction),
					"Direction_type",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction, *models.Direction_type](
					nil,
					"Direction_type",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Distance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Distance"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Distance",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance, *models.Distance](
					nil,
					"Distance",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Double:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Dynamics:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("P", instanceWithInferedType.P, formGroup, probe)
		AssociationFieldToForm("Pp", instanceWithInferedType.Pp, formGroup, probe)
		AssociationFieldToForm("Ppp", instanceWithInferedType.Ppp, formGroup, probe)
		AssociationFieldToForm("Pppp", instanceWithInferedType.Pppp, formGroup, probe)
		AssociationFieldToForm("Ppppp", instanceWithInferedType.Ppppp, formGroup, probe)
		AssociationFieldToForm("Pppppp", instanceWithInferedType.Pppppp, formGroup, probe)
		AssociationFieldToForm("F", instanceWithInferedType.F, formGroup, probe)
		AssociationFieldToForm("Ff", instanceWithInferedType.Ff, formGroup, probe)
		AssociationFieldToForm("Fff", instanceWithInferedType.Fff, formGroup, probe)
		AssociationFieldToForm("Ffff", instanceWithInferedType.Ffff, formGroup, probe)
		AssociationFieldToForm("Fffff", instanceWithInferedType.Fffff, formGroup, probe)
		AssociationFieldToForm("Ffffff", instanceWithInferedType.Ffffff, formGroup, probe)
		AssociationFieldToForm("Mp", instanceWithInferedType.Mp, formGroup, probe)
		AssociationFieldToForm("Mf", instanceWithInferedType.Mf, formGroup, probe)
		AssociationFieldToForm("Sf", instanceWithInferedType.Sf, formGroup, probe)
		AssociationFieldToForm("Sfp", instanceWithInferedType.Sfp, formGroup, probe)
		AssociationFieldToForm("Sfpp", instanceWithInferedType.Sfpp, formGroup, probe)
		AssociationFieldToForm("Fp", instanceWithInferedType.Fp, formGroup, probe)
		AssociationFieldToForm("Rf", instanceWithInferedType.Rf, formGroup, probe)
		AssociationFieldToForm("Rfz", instanceWithInferedType.Rfz, formGroup, probe)
		AssociationFieldToForm("Sfz", instanceWithInferedType.Sfz, formGroup, probe)
		AssociationFieldToForm("Sffz", instanceWithInferedType.Sffz, formGroup, probe)
		AssociationFieldToForm("Fz", instanceWithInferedType.Fz, formGroup, probe)
		AssociationFieldToForm("N", instanceWithInferedType.N, formGroup, probe)
		AssociationFieldToForm("Pf", instanceWithInferedType.Pf, formGroup, probe)
		AssociationFieldToForm("Sfzp", instanceWithInferedType.Sfzp, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Dynamics"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type, *models.Dynamics](
					nil,
					"Dynamics",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Effect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Elision:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome_note"
			rf.Fieldname = "Metronome_dot"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Metronome_note),
					"Metronome_dot",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Metronome_note, *models.Empty](
					nil,
					"Metronome_dot",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Empty_font:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_placement:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Dot"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Dot",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note, *models.Empty_placement](
					nil,
					"Dot",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Empty_placement_smufl:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_object_style_align:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_style:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_style_align:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_print_style_align_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Empty_trill_sound:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Encoding:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Encoder", instanceWithInferedType.Encoder, formGroup, probe)
		BasicFieldtoForm("Software", instanceWithInferedType.Software, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Encoding_description", instanceWithInferedType.Encoding_description, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Supports", instanceWithInferedType.Supports, formGroup, probe)

	case *models.Ending:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Extend:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Feature:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Grouping"
			rf.Fieldname = "Feature"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Grouping),
					"Feature",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Grouping, *models.Feature](
					nil,
					"Feature",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Fermata:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Figure:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Extend", instanceWithInferedType.Extend, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Figured_bass"
			rf.Fieldname = "Figure"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Figured_bass),
					"Figure",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Figured_bass, *models.Figure](
					nil,
					"Figure",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Figured_bass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Figure", instanceWithInferedType, &instanceWithInferedType.Figure, formGroup, probe)

	case *models.Fingering:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.First_fret:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Foo:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.For_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Part_clef", instanceWithInferedType.Part_clef, formGroup, probe)
		AssociationFieldToForm("Part_transpose", instanceWithInferedType.Part_transpose, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "For_part"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"For_part",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes, *models.For_part](
					nil,
					"For_part",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Formatted_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Formatted_symbol_id:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Forward:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Frame:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("First_fret", instanceWithInferedType.First_fret, formGroup, probe)
		AssociationSliceToForm("Frame_note", instanceWithInferedType, &instanceWithInferedType.Frame_note, formGroup, probe)
		BasicFieldtoForm("Unplayed", instanceWithInferedType.Unplayed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Frame_note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Astring", instanceWithInferedType.Astring, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Fret", instanceWithInferedType.Fret, formGroup, probe)
		AssociationFieldToForm("Fingering", instanceWithInferedType.Fingering, formGroup, probe)
		AssociationFieldToForm("Barre", instanceWithInferedType.Barre, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Frame"
			rf.Fieldname = "Frame_note"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Frame),
					"Frame_note",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Frame, *models.Frame_note](
					nil,
					"Frame_note",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Fret:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Glass:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Glissando:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Glyph:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Glyph"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Glyph",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance, *models.Glyph](
					nil,
					"Glyph",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Grace:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Group_barline:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Group_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Grouping:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Feature", instanceWithInferedType, &instanceWithInferedType.Feature, formGroup, probe)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Member_of", instanceWithInferedType.Member_of, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Hammer_on_pull_off:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Handbell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Harmon_closed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Harmon_mute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Harmon_closed", instanceWithInferedType.Harmon_closed, formGroup, probe)

	case *models.Harmonic:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Natural", instanceWithInferedType.Natural, formGroup, probe)
		AssociationFieldToForm("Artificial", instanceWithInferedType.Artificial, formGroup, probe)
		AssociationFieldToForm("Base_pitch", instanceWithInferedType.Base_pitch, formGroup, probe)
		AssociationFieldToForm("Touching_pitch", instanceWithInferedType.Touching_pitch, formGroup, probe)
		AssociationFieldToForm("Sounding_pitch", instanceWithInferedType.Sounding_pitch, formGroup, probe)

	case *models.Harmony:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Frame", instanceWithInferedType.Frame, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)

	case *models.Harmony_alter:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Harp_pedals:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Pedal_tuning", instanceWithInferedType, &instanceWithInferedType.Pedal_tuning, formGroup, probe)

	case *models.Heel_toe:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Hole:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Hole_type", instanceWithInferedType.Hole_type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Hole_closed", instanceWithInferedType.Hole_closed, formGroup, probe)
		BasicFieldtoForm("Hole_shape", instanceWithInferedType.Hole_shape, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Hole_closed:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Horizontal_turn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Identification:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Creator", instanceWithInferedType, &instanceWithInferedType.Creator, formGroup, probe)
		AssociationSliceToForm("Rights", instanceWithInferedType, &instanceWithInferedType.Rights, formGroup, probe)
		AssociationFieldToForm("Encoding", instanceWithInferedType.Encoding, formGroup, probe)
		BasicFieldtoForm("Source", instanceWithInferedType.Source, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Relation", instanceWithInferedType, &instanceWithInferedType.Relation, formGroup, probe)
		AssociationFieldToForm("Miscellaneous", instanceWithInferedType.Miscellaneous, formGroup, probe)

	case *models.Image:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Instrument"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note, *models.Instrument](
					nil,
					"Instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Instrument_change:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Instrument_link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Part_link"
			rf.Fieldname = "Instrument_link"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Part_link),
					"Instrument_link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Part_link, *models.Instrument_link](
					nil,
					"Instrument_link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Interchangeable:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Inversion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Key_octave", instanceWithInferedType, &instanceWithInferedType.Key_octave, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Key"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Key",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes, *models.Key](
					nil,
					"Key",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Key_accidental:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Key_octave:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Key"
			rf.Fieldname = "Key_octave"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Key),
					"Key_octave",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Key, *models.Key_octave](
					nil,
					"Key_octave",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Kind:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Level:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Line_detail:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Line_width:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Line_width"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Line_width",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance, *models.Line_width](
					nil,
					"Line_width",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Credit"
			rf.Fieldname = "Link"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Credit),
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Credit, *models.Link](
					nil,
					"Link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Listen:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Assess", instanceWithInferedType.Assess, formGroup, probe)
		AssociationFieldToForm("Wait", instanceWithInferedType.Wait, formGroup, probe)
		AssociationFieldToForm("Other_listen", instanceWithInferedType.Other_listen, formGroup, probe)

	case *models.Listening:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		AssociationFieldToForm("Sync", instanceWithInferedType.Sync, formGroup, probe)
		AssociationFieldToForm("Other_listening", instanceWithInferedType.Other_listening, formGroup, probe)

	case *models.Lyric:
		// insertion point
		AssociationFieldToForm("End_line", instanceWithInferedType.End_line, formGroup, probe)
		AssociationFieldToForm("End_paragraph", instanceWithInferedType.End_paragraph, formGroup, probe)
		AssociationFieldToForm("Extend", instanceWithInferedType.Extend, formGroup, probe)
		AssociationFieldToForm("Laughing", instanceWithInferedType.Laughing, formGroup, probe)
		AssociationFieldToForm("Humming", instanceWithInferedType.Humming, formGroup, probe)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Lyric"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Lyric",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note, *models.Lyric](
					nil,
					"Lyric",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lyric_font:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Lyric_font"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Defaults),
					"Lyric_font",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Defaults, *models.Lyric_font](
					nil,
					"Lyric_font",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Lyric_language:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EmptyString", instanceWithInferedType.EmptyString, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Defaults"
			rf.Fieldname = "Lyric_language"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Defaults),
					"Lyric_language",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Defaults, *models.Lyric_language](
					nil,
					"Lyric_language",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Measure_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Measure_numbering:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Measure_repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Measure_style:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Multiple_rest", instanceWithInferedType.Multiple_rest, formGroup, probe)
		AssociationFieldToForm("Measure_repeat", instanceWithInferedType.Measure_repeat, formGroup, probe)
		AssociationFieldToForm("Beat_repeat", instanceWithInferedType.Beat_repeat, formGroup, probe)
		AssociationFieldToForm("Slash", instanceWithInferedType.Slash, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Measure_style"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Measure_style",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes, *models.Measure_style](
					nil,
					"Measure_style",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Membrane:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Metal:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Metronome:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Metronome_beam:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Metronome_note"
			rf.Fieldname = "Metronome_beam"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Metronome_note),
					"Metronome_beam",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Metronome_note, *models.Metronome_beam](
					nil,
					"Metronome_beam",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Metronome_note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Metronome_dot", instanceWithInferedType, &instanceWithInferedType.Metronome_dot, formGroup, probe)
		AssociationSliceToForm("Metronome_beam", instanceWithInferedType, &instanceWithInferedType.Metronome_beam, formGroup, probe)
		AssociationFieldToForm("Metronome_tied", instanceWithInferedType.Metronome_tied, formGroup, probe)
		AssociationFieldToForm("Metronome_tuplet", instanceWithInferedType.Metronome_tuplet, formGroup, probe)

	case *models.Metronome_tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Metronome_tuplet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Midi_device:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Midi_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Midi_name", instanceWithInferedType.Midi_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Miscellaneous:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Miscellaneous_field", instanceWithInferedType, &instanceWithInferedType.Miscellaneous_field, formGroup, probe)

	case *models.Miscellaneous_field:
		// insertion point
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Miscellaneous"
			rf.Fieldname = "Miscellaneous_field"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Miscellaneous),
					"Miscellaneous_field",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Miscellaneous, *models.Miscellaneous_field](
					nil,
					"Miscellaneous_field",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Mordent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Multiple_rest:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Name_display:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Accidental_text", instanceWithInferedType.Accidental_text, formGroup, probe)

	case *models.Non_arpeggiate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Notations:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Tied", instanceWithInferedType.Tied, formGroup, probe)
		AssociationFieldToForm("Slur", instanceWithInferedType.Slur, formGroup, probe)
		AssociationFieldToForm("Tuplet", instanceWithInferedType.Tuplet, formGroup, probe)
		AssociationFieldToForm("Glissando", instanceWithInferedType.Glissando, formGroup, probe)
		AssociationFieldToForm("Slide", instanceWithInferedType.Slide, formGroup, probe)
		AssociationFieldToForm("Ornaments", instanceWithInferedType.Ornaments, formGroup, probe)
		AssociationFieldToForm("Technical", instanceWithInferedType.Technical, formGroup, probe)
		AssociationFieldToForm("Articulations", instanceWithInferedType.Articulations, formGroup, probe)
		AssociationFieldToForm("Dynamics", instanceWithInferedType.Dynamics, formGroup, probe)
		AssociationFieldToForm("Fermata", instanceWithInferedType.Fermata, formGroup, probe)
		AssociationFieldToForm("Arpeggiate", instanceWithInferedType.Arpeggiate, formGroup, probe)
		AssociationFieldToForm("Non_arpeggiate", instanceWithInferedType.Non_arpeggiate, formGroup, probe)
		AssociationFieldToForm("Accidental_mark", instanceWithInferedType.Accidental_mark, formGroup, probe)
		AssociationFieldToForm("Other_notation", instanceWithInferedType.Other_notation, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Note"
			rf.Fieldname = "Notations"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Note),
					"Notations",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Note, *models.Notations](
					nil,
					"Notations",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Note:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Instrument", instanceWithInferedType, &instanceWithInferedType.Instrument, formGroup, probe)
		AssociationFieldToForm("Type_", instanceWithInferedType.Type_, formGroup, probe)
		AssociationSliceToForm("Dot", instanceWithInferedType, &instanceWithInferedType.Dot, formGroup, probe)
		AssociationFieldToForm("Accidental", instanceWithInferedType.Accidental, formGroup, probe)
		AssociationFieldToForm("Time_modification", instanceWithInferedType.Time_modification, formGroup, probe)
		AssociationFieldToForm("Stem", instanceWithInferedType.Stem, formGroup, probe)
		AssociationFieldToForm("Notehead", instanceWithInferedType.Notehead, formGroup, probe)
		AssociationFieldToForm("Notehead_text", instanceWithInferedType.Notehead_text, formGroup, probe)
		AssociationFieldToForm("Beam", instanceWithInferedType.Beam, formGroup, probe)
		AssociationSliceToForm("Notations", instanceWithInferedType, &instanceWithInferedType.Notations, formGroup, probe)
		AssociationSliceToForm("Lyric", instanceWithInferedType, &instanceWithInferedType.Lyric, formGroup, probe)
		AssociationFieldToForm("Play", instanceWithInferedType.Play, formGroup, probe)
		AssociationFieldToForm("Listen", instanceWithInferedType.Listen, formGroup, probe)

	case *models.Note_size:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Note_size"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Note_size",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance, *models.Note_size](
					nil,
					"Note_size",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Note_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Notehead:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Notehead_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Accidental_text", instanceWithInferedType.Accidental_text, formGroup, probe)

	case *models.Numeral:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Numeral_root", instanceWithInferedType.Numeral_root, formGroup, probe)
		AssociationFieldToForm("Numeral_alter", instanceWithInferedType.Numeral_alter, formGroup, probe)
		AssociationFieldToForm("Numeral_key", instanceWithInferedType.Numeral_key, formGroup, probe)

	case *models.Numeral_key:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Numeral_root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Octave_shift:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Offset:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Opus:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Ornaments:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Accidental_mark", instanceWithInferedType, &instanceWithInferedType.Accidental_mark, formGroup, probe)
		AssociationFieldToForm("Trill_mark", instanceWithInferedType.Trill_mark, formGroup, probe)
		AssociationFieldToForm("Turn", instanceWithInferedType.Turn, formGroup, probe)
		AssociationFieldToForm("Delayed_turn", instanceWithInferedType.Delayed_turn, formGroup, probe)
		AssociationFieldToForm("Inverted_turn", instanceWithInferedType.Inverted_turn, formGroup, probe)
		AssociationFieldToForm("Delayed_inverted_turn", instanceWithInferedType.Delayed_inverted_turn, formGroup, probe)
		AssociationFieldToForm("Vertical_turn", instanceWithInferedType.Vertical_turn, formGroup, probe)
		AssociationFieldToForm("Inverted_vertical_turn", instanceWithInferedType.Inverted_vertical_turn, formGroup, probe)
		AssociationFieldToForm("Shake", instanceWithInferedType.Shake, formGroup, probe)
		AssociationFieldToForm("Wavy_line", instanceWithInferedType.Wavy_line, formGroup, probe)
		AssociationFieldToForm("Mordent", instanceWithInferedType.Mordent, formGroup, probe)
		AssociationFieldToForm("Inverted_mordent", instanceWithInferedType.Inverted_mordent, formGroup, probe)
		AssociationFieldToForm("Schleifer", instanceWithInferedType.Schleifer, formGroup, probe)
		AssociationFieldToForm("Tremolo", instanceWithInferedType.Tremolo, formGroup, probe)
		AssociationFieldToForm("Haydn", instanceWithInferedType.Haydn, formGroup, probe)

	case *models.Other_appearance:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Appearance"
			rf.Fieldname = "Other_appearance"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Appearance),
					"Other_appearance",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Appearance, *models.Other_appearance](
					nil,
					"Other_appearance",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Other_listening:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Other_notation:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Other_play:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Page_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Page_margins", instanceWithInferedType.Page_margins, formGroup, probe)

	case *models.Page_margins:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_clef:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Group_name_display", instanceWithInferedType.Group_name_display, formGroup, probe)
		AssociationFieldToForm("Group_abbreviation_display", instanceWithInferedType.Group_abbreviation_display, formGroup, probe)
		AssociationFieldToForm("Group_symbol", instanceWithInferedType.Group_symbol, formGroup, probe)
		AssociationFieldToForm("Group_barline", instanceWithInferedType.Group_barline, formGroup, probe)
		AssociationFieldToForm("Group_time", instanceWithInferedType.Group_time, formGroup, probe)
		BasicFieldtoForm("Number", instanceWithInferedType.Number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_link:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Instrument_link", instanceWithInferedType, &instanceWithInferedType.Instrument_link, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Part_link"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Part_link",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part, *models.Part_link](
					nil,
					"Part_link",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Part_list:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_symbol:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Part_transpose:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pedal:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pedal_tuning:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Harp_pedals"
			rf.Fieldname = "Pedal_tuning"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Harp_pedals),
					"Pedal_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Harp_pedals, *models.Pedal_tuning](
					nil,
					"Pedal_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Percussion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Glass", instanceWithInferedType.Glass, formGroup, probe)
		AssociationFieldToForm("Metal", instanceWithInferedType.Metal, formGroup, probe)
		AssociationFieldToForm("Wood", instanceWithInferedType.Wood, formGroup, probe)
		AssociationFieldToForm("Pitched", instanceWithInferedType.Pitched, formGroup, probe)
		AssociationFieldToForm("Membrane", instanceWithInferedType.Membrane, formGroup, probe)
		AssociationFieldToForm("Effect", instanceWithInferedType.Effect, formGroup, probe)
		AssociationFieldToForm("Timpani", instanceWithInferedType.Timpani, formGroup, probe)
		AssociationFieldToForm("Beater", instanceWithInferedType.Beater, formGroup, probe)
		AssociationFieldToForm("Stick", instanceWithInferedType.Stick, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Percussion"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Percussion",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type, *models.Percussion](
					nil,
					"Percussion",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Pitch:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Pitched:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Play:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Ipa", instanceWithInferedType.Ipa, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Other_play", instanceWithInferedType.Other_play, formGroup, probe)

	case *models.Player:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Player_name", instanceWithInferedType.Player_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Player"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Player",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part, *models.Player](
					nil,
					"Player",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Principal_voice:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Print:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Measure_layout", instanceWithInferedType.Measure_layout, formGroup, probe)
		AssociationFieldToForm("Measure_numbering", instanceWithInferedType.Measure_numbering, formGroup, probe)
		AssociationFieldToForm("Part_name_display", instanceWithInferedType.Part_name_display, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation_display", instanceWithInferedType.Part_abbreviation_display, formGroup, probe)

	case *models.Release:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Repeat:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Rest:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Root:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Root_step", instanceWithInferedType.Root_step, formGroup, probe)
		AssociationFieldToForm("Root_alter", instanceWithInferedType.Root_alter, formGroup, probe)

	case *models.Root_step:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Text", instanceWithInferedType.Text, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Scaling:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Scordatura:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Accord", instanceWithInferedType, &instanceWithInferedType.Accord, formGroup, probe)

	case *models.Score_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Instrument_name", instanceWithInferedType.Instrument_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Instrument_abbreviation", instanceWithInferedType.Instrument_abbreviation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Score_part"
			rf.Fieldname = "Score_instrument"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Score_part),
					"Score_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Score_part, *models.Score_instrument](
					nil,
					"Score_instrument",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Score_part:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Identification", instanceWithInferedType.Identification, formGroup, probe)
		AssociationSliceToForm("Part_link", instanceWithInferedType, &instanceWithInferedType.Part_link, formGroup, probe)
		AssociationFieldToForm("Part_name_display", instanceWithInferedType.Part_name_display, formGroup, probe)
		AssociationFieldToForm("Part_abbreviation_display", instanceWithInferedType.Part_abbreviation_display, formGroup, probe)
		AssociationSliceToForm("Score_instrument", instanceWithInferedType, &instanceWithInferedType.Score_instrument, formGroup, probe)
		AssociationSliceToForm("Player", instanceWithInferedType, &instanceWithInferedType.Player, formGroup, probe)

	case *models.Score_partwise:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Score_timewise:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Segno:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Direction_type"
			rf.Fieldname = "Segno"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Direction_type),
					"Segno",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Direction_type, *models.Segno](
					nil,
					"Segno",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Slash:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Slide:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Slur:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Sound:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Swing", instanceWithInferedType.Swing, formGroup, probe)
		AssociationFieldToForm("Offset", instanceWithInferedType.Offset, formGroup, probe)
		BasicFieldtoForm("Segno", instanceWithInferedType.Segno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Dalsegno", instanceWithInferedType.Dalsegno, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Coda", instanceWithInferedType.Coda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Tocoda", instanceWithInferedType.Tocoda, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Fine", instanceWithInferedType.Fine, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Staff_details:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Staff_tuning", instanceWithInferedType, &instanceWithInferedType.Staff_tuning, formGroup, probe)
		AssociationFieldToForm("Staff_size", instanceWithInferedType.Staff_size, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Staff_details"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Staff_details",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes, *models.Staff_details](
					nil,
					"Staff_details",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Staff_divide:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Staff_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Staff_size:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Staff_tuning:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Staff_details"
			rf.Fieldname = "Staff_tuning"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Staff_details),
					"Staff_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Staff_details, *models.Staff_tuning](
					nil,
					"Staff_tuning",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Stem:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Stick:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.String_mute:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Strong_accent:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Supports:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Swing:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Swing_style", instanceWithInferedType.Swing_style, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Straight", instanceWithInferedType.Straight, formGroup, probe)

	case *models.Sync:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.System_dividers:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Left_divider", instanceWithInferedType.Left_divider, formGroup, probe)
		AssociationFieldToForm("Right_divider", instanceWithInferedType.Right_divider, formGroup, probe)

	case *models.System_layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("System_margins", instanceWithInferedType.System_margins, formGroup, probe)
		AssociationFieldToForm("System_dividers", instanceWithInferedType.System_dividers, formGroup, probe)

	case *models.System_margins:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tap:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Technical:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Up_bow", instanceWithInferedType.Up_bow, formGroup, probe)
		AssociationFieldToForm("Down_bow", instanceWithInferedType.Down_bow, formGroup, probe)
		AssociationFieldToForm("Harmonic", instanceWithInferedType.Harmonic, formGroup, probe)
		AssociationFieldToForm("Open_string", instanceWithInferedType.Open_string, formGroup, probe)
		AssociationFieldToForm("Thumb_position", instanceWithInferedType.Thumb_position, formGroup, probe)
		AssociationFieldToForm("Fingering", instanceWithInferedType.Fingering, formGroup, probe)
		AssociationFieldToForm("Double_tongue", instanceWithInferedType.Double_tongue, formGroup, probe)
		AssociationFieldToForm("Triple_tongue", instanceWithInferedType.Triple_tongue, formGroup, probe)
		AssociationFieldToForm("Stopped", instanceWithInferedType.Stopped, formGroup, probe)
		AssociationFieldToForm("Snap_pizzicato", instanceWithInferedType.Snap_pizzicato, formGroup, probe)
		AssociationFieldToForm("Fret", instanceWithInferedType.Fret, formGroup, probe)
		BasicFieldtoForm("Astring", instanceWithInferedType.Astring, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Hammer_on", instanceWithInferedType.Hammer_on, formGroup, probe)
		AssociationFieldToForm("Pull_off", instanceWithInferedType.Pull_off, formGroup, probe)
		AssociationFieldToForm("Bend", instanceWithInferedType.Bend, formGroup, probe)
		AssociationFieldToForm("Tap", instanceWithInferedType.Tap, formGroup, probe)
		AssociationFieldToForm("Heel", instanceWithInferedType.Heel, formGroup, probe)
		AssociationFieldToForm("Toe", instanceWithInferedType.Toe, formGroup, probe)
		AssociationFieldToForm("Fingernails", instanceWithInferedType.Fingernails, formGroup, probe)
		AssociationFieldToForm("Hole", instanceWithInferedType.Hole, formGroup, probe)
		AssociationFieldToForm("Arrow", instanceWithInferedType.Arrow, formGroup, probe)
		AssociationFieldToForm("Handbell", instanceWithInferedType.Handbell, formGroup, probe)
		AssociationFieldToForm("Brass_bend", instanceWithInferedType.Brass_bend, formGroup, probe)
		AssociationFieldToForm("Flip", instanceWithInferedType.Flip, formGroup, probe)
		AssociationFieldToForm("Smear", instanceWithInferedType.Smear, formGroup, probe)
		AssociationFieldToForm("Open", instanceWithInferedType.Open, formGroup, probe)
		AssociationFieldToForm("Half_muted", instanceWithInferedType.Half_muted, formGroup, probe)
		AssociationFieldToForm("Harmon_mute", instanceWithInferedType.Harmon_mute, formGroup, probe)
		AssociationFieldToForm("Golpe", instanceWithInferedType.Golpe, formGroup, probe)

	case *models.Text_element_data:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("EmptyString", instanceWithInferedType.EmptyString, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tie:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tied:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Time:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Senza_misura", instanceWithInferedType.Senza_misura, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Time_modification:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Timpani:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Transpose:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Attributes"
			rf.Fieldname = "Transpose"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Attributes),
					"Transpose",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Attributes, *models.Transpose](
					nil,
					"Transpose",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tremolo:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tuplet:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Tuplet_actual", instanceWithInferedType.Tuplet_actual, formGroup, probe)
		AssociationFieldToForm("Tuplet_normal", instanceWithInferedType.Tuplet_normal, formGroup, probe)

	case *models.Tuplet_dot:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tuplet_portion"
			rf.Fieldname = "Tuplet_dot"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Tuplet_portion),
					"Tuplet_dot",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Tuplet_portion, *models.Tuplet_dot](
					nil,
					"Tuplet_dot",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Tuplet_number:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Tuplet_portion:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Tuplet_number", instanceWithInferedType.Tuplet_number, formGroup, probe)
		AssociationFieldToForm("Tuplet_type", instanceWithInferedType.Tuplet_type, formGroup, probe)
		AssociationSliceToForm("Tuplet_dot", instanceWithInferedType, &instanceWithInferedType.Tuplet_dot, formGroup, probe)

	case *models.Tuplet_type:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Typed_text:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Type", instanceWithInferedType.Type, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Creator"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Identification),
					"Creator",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Identification, *models.Typed_text](
					nil,
					"Creator",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Rights"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Identification),
					"Rights",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Identification, *models.Typed_text](
					nil,
					"Rights",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Identification"
			rf.Fieldname = "Relation"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Identification),
					"Relation",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Identification, *models.Typed_text](
					nil,
					"Relation",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Unpitched:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Virtual_instrument:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Virtual_library", instanceWithInferedType.Virtual_library, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Virtual_name", instanceWithInferedType.Virtual_name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Wait:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Wavy_line:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Wedge:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Wood:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Work:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Work_number", instanceWithInferedType.Work_number, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Work_title", instanceWithInferedType.Work_title, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Opus", instanceWithInferedType.Opus, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
