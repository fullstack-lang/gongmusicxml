// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Accidental":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Accidental Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccidentalFormCallback(
			nil,
			probe,
			formGroup,
		)
		accidental := new(models.Accidental)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accidental, formGroup, probe)
	case "Accidental_mark":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Accidental_mark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_markFormCallback(
			nil,
			probe,
			formGroup,
		)
		accidental_mark := new(models.Accidental_mark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accidental_mark, formGroup, probe)
	case "Accidental_text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Accidental_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		accidental_text := new(models.Accidental_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accidental_text, formGroup, probe)
	case "Accord":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Accord Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccordFormCallback(
			nil,
			probe,
			formGroup,
		)
		accord := new(models.Accord)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accord, formGroup, probe)
	case "Accordion_registration":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Accordion_registration Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accordion_registrationFormCallback(
			nil,
			probe,
			formGroup,
		)
		accordion_registration := new(models.Accordion_registration)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accordion_registration, formGroup, probe)
	case "AnyType":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "AnyType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnyTypeFormCallback(
			nil,
			probe,
			formGroup,
		)
		anytype := new(models.AnyType)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(anytype, formGroup, probe)
	case "Appearance":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Appearance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AppearanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		appearance := new(models.Appearance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(appearance, formGroup, probe)
	case "Arpeggiate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Arpeggiate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArpeggiateFormCallback(
			nil,
			probe,
			formGroup,
		)
		arpeggiate := new(models.Arpeggiate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arpeggiate, formGroup, probe)
	case "Arrow":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Arrow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArrowFormCallback(
			nil,
			probe,
			formGroup,
		)
		arrow := new(models.Arrow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arrow, formGroup, probe)
	case "Articulations":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Articulations Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArticulationsFormCallback(
			nil,
			probe,
			formGroup,
		)
		articulations := new(models.Articulations)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(articulations, formGroup, probe)
	case "Assess":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Assess Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AssessFormCallback(
			nil,
			probe,
			formGroup,
		)
		assess := new(models.Assess)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(assess, formGroup, probe)
	case "Attributes":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Attributes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributesFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributes := new(models.Attributes)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributes, formGroup, probe)
	case "Backup":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Backup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BackupFormCallback(
			nil,
			probe,
			formGroup,
		)
		backup := new(models.Backup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(backup, formGroup, probe)
	case "Bar_style_color":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bar_style_color Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bar_style_colorFormCallback(
			nil,
			probe,
			formGroup,
		)
		bar_style_color := new(models.Bar_style_color)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bar_style_color, formGroup, probe)
	case "Barline":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Barline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarlineFormCallback(
			nil,
			probe,
			formGroup,
		)
		barline := new(models.Barline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(barline, formGroup, probe)
	case "Barre":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Barre Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarreFormCallback(
			nil,
			probe,
			formGroup,
		)
		barre := new(models.Barre)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(barre, formGroup, probe)
	case "Bass":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BassFormCallback(
			nil,
			probe,
			formGroup,
		)
		bass := new(models.Bass)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bass, formGroup, probe)
	case "Bass_step":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bass_step Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bass_stepFormCallback(
			nil,
			probe,
			formGroup,
		)
		bass_step := new(models.Bass_step)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bass_step, formGroup, probe)
	case "Beam":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Beam Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeamFormCallback(
			nil,
			probe,
			formGroup,
		)
		beam := new(models.Beam)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beam, formGroup, probe)
	case "Beat_repeat":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Beat_repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_repeatFormCallback(
			nil,
			probe,
			formGroup,
		)
		beat_repeat := new(models.Beat_repeat)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beat_repeat, formGroup, probe)
	case "Beat_unit_tied":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Beat_unit_tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_unit_tiedFormCallback(
			nil,
			probe,
			formGroup,
		)
		beat_unit_tied := new(models.Beat_unit_tied)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beat_unit_tied, formGroup, probe)
	case "Beater":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Beater Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeaterFormCallback(
			nil,
			probe,
			formGroup,
		)
		beater := new(models.Beater)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beater, formGroup, probe)
	case "Bend":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bend Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BendFormCallback(
			nil,
			probe,
			formGroup,
		)
		bend := new(models.Bend)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bend, formGroup, probe)
	case "Bookmark":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bookmark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookmarkFormCallback(
			nil,
			probe,
			formGroup,
		)
		bookmark := new(models.Bookmark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bookmark, formGroup, probe)
	case "Bracket":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Bracket Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BracketFormCallback(
			nil,
			probe,
			formGroup,
		)
		bracket := new(models.Bracket)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bracket, formGroup, probe)
	case "Breath_mark":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Breath_mark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Breath_markFormCallback(
			nil,
			probe,
			formGroup,
		)
		breath_mark := new(models.Breath_mark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(breath_mark, formGroup, probe)
	case "Caesura":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Caesura Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CaesuraFormCallback(
			nil,
			probe,
			formGroup,
		)
		caesura := new(models.Caesura)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(caesura, formGroup, probe)
	case "Cancel":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Cancel Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CancelFormCallback(
			nil,
			probe,
			formGroup,
		)
		cancel := new(models.Cancel)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cancel, formGroup, probe)
	case "Clef":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Clef Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClefFormCallback(
			nil,
			probe,
			formGroup,
		)
		clef := new(models.Clef)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(clef, formGroup, probe)
	case "Coda":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Coda Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CodaFormCallback(
			nil,
			probe,
			formGroup,
		)
		coda := new(models.Coda)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(coda, formGroup, probe)
	case "Credit":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Credit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			probe,
			formGroup,
		)
		credit := new(models.Credit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(credit, formGroup, probe)
	case "Dashes":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Dashes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DashesFormCallback(
			nil,
			probe,
			formGroup,
		)
		dashes := new(models.Dashes)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dashes, formGroup, probe)
	case "Defaults":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Defaults Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DefaultsFormCallback(
			nil,
			probe,
			formGroup,
		)
		defaults := new(models.Defaults)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(defaults, formGroup, probe)
	case "Degree":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Degree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DegreeFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree := new(models.Degree)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree, formGroup, probe)
	case "Degree_alter":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Degree_alter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_alterFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree_alter := new(models.Degree_alter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree_alter, formGroup, probe)
	case "Degree_type":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Degree_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree_type := new(models.Degree_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree_type, formGroup, probe)
	case "Degree_value":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Degree_value Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_valueFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree_value := new(models.Degree_value)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree_value, formGroup, probe)
	case "Direction":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Direction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DirectionFormCallback(
			nil,
			probe,
			formGroup,
		)
		direction := new(models.Direction)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(direction, formGroup, probe)
	case "Direction_type":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Direction_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Direction_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		direction_type := new(models.Direction_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(direction_type, formGroup, probe)
	case "Distance":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Distance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DistanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		distance := new(models.Distance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(distance, formGroup, probe)
	case "Double":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Double Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DoubleFormCallback(
			nil,
			probe,
			formGroup,
		)
		double := new(models.Double)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(double, formGroup, probe)
	case "Dynamics":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Dynamics Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DynamicsFormCallback(
			nil,
			probe,
			formGroup,
		)
		dynamics := new(models.Dynamics)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dynamics, formGroup, probe)
	case "Effect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Effect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EffectFormCallback(
			nil,
			probe,
			formGroup,
		)
		effect := new(models.Effect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(effect, formGroup, probe)
	case "Elision":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Elision Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElisionFormCallback(
			nil,
			probe,
			formGroup,
		)
		elision := new(models.Elision)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(elision, formGroup, probe)
	case "Empty":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EmptyFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty := new(models.Empty)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty, formGroup, probe)
	case "Empty_font":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_fontFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_font := new(models.Empty_font)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_font, formGroup, probe)
	case "Empty_line":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_lineFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_line := new(models.Empty_line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_line, formGroup, probe)
	case "Empty_placement":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_placement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placementFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_placement := new(models.Empty_placement)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_placement, formGroup, probe)
	case "Empty_placement_smufl":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_placement_smufl Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placement_smuflFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_placement_smufl := new(models.Empty_placement_smufl)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_placement_smufl, formGroup, probe)
	case "Empty_print_object_style_align":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_print_object_style_align Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_object_style_alignFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_object_style_align := new(models.Empty_print_object_style_align)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_object_style_align, formGroup, probe)
	case "Empty_print_style":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_print_style Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_styleFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_style := new(models.Empty_print_style)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_style, formGroup, probe)
	case "Empty_print_style_align":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_print_style_align Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_alignFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_style_align := new(models.Empty_print_style_align)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_style_align, formGroup, probe)
	case "Empty_print_style_align_id":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_print_style_align_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_align_idFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_style_align_id := new(models.Empty_print_style_align_id)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_style_align_id, formGroup, probe)
	case "Empty_trill_sound":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Empty_trill_sound Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_trill_soundFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_trill_sound := new(models.Empty_trill_sound)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_trill_sound, formGroup, probe)
	case "Encoding":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Encoding Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EncodingFormCallback(
			nil,
			probe,
			formGroup,
		)
		encoding := new(models.Encoding)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(encoding, formGroup, probe)
	case "Ending":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Ending Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndingFormCallback(
			nil,
			probe,
			formGroup,
		)
		ending := new(models.Ending)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ending, formGroup, probe)
	case "Extend":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Extend Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtendFormCallback(
			nil,
			probe,
			formGroup,
		)
		extend := new(models.Extend)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(extend, formGroup, probe)
	case "Feature":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Feature Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FeatureFormCallback(
			nil,
			probe,
			formGroup,
		)
		feature := new(models.Feature)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(feature, formGroup, probe)
	case "Fermata":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Fermata Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FermataFormCallback(
			nil,
			probe,
			formGroup,
		)
		fermata := new(models.Fermata)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fermata, formGroup, probe)
	case "Figure":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Figure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FigureFormCallback(
			nil,
			probe,
			formGroup,
		)
		figure := new(models.Figure)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(figure, formGroup, probe)
	case "Figured_bass":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Figured_bass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Figured_bassFormCallback(
			nil,
			probe,
			formGroup,
		)
		figured_bass := new(models.Figured_bass)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(figured_bass, formGroup, probe)
	case "Fingering":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Fingering Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FingeringFormCallback(
			nil,
			probe,
			formGroup,
		)
		fingering := new(models.Fingering)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fingering, formGroup, probe)
	case "First_fret":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "First_fret Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__First_fretFormCallback(
			nil,
			probe,
			formGroup,
		)
		first_fret := new(models.First_fret)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(first_fret, formGroup, probe)
	case "Foo":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Foo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			nil,
			probe,
			formGroup,
		)
		foo := new(models.Foo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(foo, formGroup, probe)
	case "For_part":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "For_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__For_partFormCallback(
			nil,
			probe,
			formGroup,
		)
		for_part := new(models.For_part)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(for_part, formGroup, probe)
	case "Formatted_symbol":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Formatted_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbolFormCallback(
			nil,
			probe,
			formGroup,
		)
		formatted_symbol := new(models.Formatted_symbol)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formatted_symbol, formGroup, probe)
	case "Formatted_symbol_id":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Formatted_symbol_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbol_idFormCallback(
			nil,
			probe,
			formGroup,
		)
		formatted_symbol_id := new(models.Formatted_symbol_id)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formatted_symbol_id, formGroup, probe)
	case "Forward":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Forward Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ForwardFormCallback(
			nil,
			probe,
			formGroup,
		)
		forward := new(models.Forward)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(forward, formGroup, probe)
	case "Frame":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Frame Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrameFormCallback(
			nil,
			probe,
			formGroup,
		)
		frame := new(models.Frame)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(frame, formGroup, probe)
	case "Frame_note":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Frame_note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Frame_noteFormCallback(
			nil,
			probe,
			formGroup,
		)
		frame_note := new(models.Frame_note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(frame_note, formGroup, probe)
	case "Fret":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Fret Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FretFormCallback(
			nil,
			probe,
			formGroup,
		)
		fret := new(models.Fret)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fret, formGroup, probe)
	case "Glass":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Glass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlassFormCallback(
			nil,
			probe,
			formGroup,
		)
		glass := new(models.Glass)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(glass, formGroup, probe)
	case "Glissando":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Glissando Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlissandoFormCallback(
			nil,
			probe,
			formGroup,
		)
		glissando := new(models.Glissando)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(glissando, formGroup, probe)
	case "Glyph":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Glyph Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlyphFormCallback(
			nil,
			probe,
			formGroup,
		)
		glyph := new(models.Glyph)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(glyph, formGroup, probe)
	case "Grace":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Grace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GraceFormCallback(
			nil,
			probe,
			formGroup,
		)
		grace := new(models.Grace)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(grace, formGroup, probe)
	case "Group_barline":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Group_barline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_barlineFormCallback(
			nil,
			probe,
			formGroup,
		)
		group_barline := new(models.Group_barline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group_barline, formGroup, probe)
	case "Group_symbol":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Group_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_symbolFormCallback(
			nil,
			probe,
			formGroup,
		)
		group_symbol := new(models.Group_symbol)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group_symbol, formGroup, probe)
	case "Grouping":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Grouping Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupingFormCallback(
			nil,
			probe,
			formGroup,
		)
		grouping := new(models.Grouping)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(grouping, formGroup, probe)
	case "Hammer_on_pull_off":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Hammer_on_pull_off Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hammer_on_pull_offFormCallback(
			nil,
			probe,
			formGroup,
		)
		hammer_on_pull_off := new(models.Hammer_on_pull_off)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hammer_on_pull_off, formGroup, probe)
	case "Handbell":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Handbell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HandbellFormCallback(
			nil,
			probe,
			formGroup,
		)
		handbell := new(models.Handbell)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(handbell, formGroup, probe)
	case "Harmon_closed":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Harmon_closed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_closedFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmon_closed := new(models.Harmon_closed)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmon_closed, formGroup, probe)
	case "Harmon_mute":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Harmon_mute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_muteFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmon_mute := new(models.Harmon_mute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmon_mute, formGroup, probe)
	case "Harmonic":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Harmonic Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonicFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmonic := new(models.Harmonic)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmonic, formGroup, probe)
	case "Harmony":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Harmony Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonyFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmony := new(models.Harmony)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmony, formGroup, probe)
	case "Harmony_alter":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Harmony_alter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmony_alterFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmony_alter := new(models.Harmony_alter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmony_alter, formGroup, probe)
	case "Harp_pedals":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Harp_pedals Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harp_pedalsFormCallback(
			nil,
			probe,
			formGroup,
		)
		harp_pedals := new(models.Harp_pedals)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harp_pedals, formGroup, probe)
	case "Heel_toe":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Heel_toe Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Heel_toeFormCallback(
			nil,
			probe,
			formGroup,
		)
		heel_toe := new(models.Heel_toe)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(heel_toe, formGroup, probe)
	case "Hole":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Hole Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HoleFormCallback(
			nil,
			probe,
			formGroup,
		)
		hole := new(models.Hole)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hole, formGroup, probe)
	case "Hole_closed":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Hole_closed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hole_closedFormCallback(
			nil,
			probe,
			formGroup,
		)
		hole_closed := new(models.Hole_closed)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hole_closed, formGroup, probe)
	case "Horizontal_turn":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Horizontal_turn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Horizontal_turnFormCallback(
			nil,
			probe,
			formGroup,
		)
		horizontal_turn := new(models.Horizontal_turn)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(horizontal_turn, formGroup, probe)
	case "Identification":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Identification Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__IdentificationFormCallback(
			nil,
			probe,
			formGroup,
		)
		identification := new(models.Identification)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(identification, formGroup, probe)
	case "Image":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Image Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ImageFormCallback(
			nil,
			probe,
			formGroup,
		)
		image := new(models.Image)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(image, formGroup, probe)
	case "Instrument":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InstrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		instrument := new(models.Instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(instrument, formGroup, probe)
	case "Instrument_change":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Instrument_change Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_changeFormCallback(
			nil,
			probe,
			formGroup,
		)
		instrument_change := new(models.Instrument_change)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(instrument_change, formGroup, probe)
	case "Instrument_link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Instrument_link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_linkFormCallback(
			nil,
			probe,
			formGroup,
		)
		instrument_link := new(models.Instrument_link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(instrument_link, formGroup, probe)
	case "Interchangeable":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Interchangeable Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InterchangeableFormCallback(
			nil,
			probe,
			formGroup,
		)
		interchangeable := new(models.Interchangeable)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(interchangeable, formGroup, probe)
	case "Inversion":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Inversion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InversionFormCallback(
			nil,
			probe,
			formGroup,
		)
		inversion := new(models.Inversion)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(inversion, formGroup, probe)
	case "Key":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			probe,
			formGroup,
		)
		key := new(models.Key)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key, formGroup, probe)
	case "Key_accidental":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Key_accidental Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_accidentalFormCallback(
			nil,
			probe,
			formGroup,
		)
		key_accidental := new(models.Key_accidental)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key_accidental, formGroup, probe)
	case "Key_octave":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Key_octave Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_octaveFormCallback(
			nil,
			probe,
			formGroup,
		)
		key_octave := new(models.Key_octave)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key_octave, formGroup, probe)
	case "Kind":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Kind Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KindFormCallback(
			nil,
			probe,
			formGroup,
		)
		kind := new(models.Kind)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(kind, formGroup, probe)
	case "Level":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Level Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LevelFormCallback(
			nil,
			probe,
			formGroup,
		)
		level := new(models.Level)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(level, formGroup, probe)
	case "Line_detail":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Line_detail Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_detailFormCallback(
			nil,
			probe,
			formGroup,
		)
		line_detail := new(models.Line_detail)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line_detail, formGroup, probe)
	case "Line_width":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Line_width Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_widthFormCallback(
			nil,
			probe,
			formGroup,
		)
		line_width := new(models.Line_width)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line_width, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "Listen":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Listen Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListenFormCallback(
			nil,
			probe,
			formGroup,
		)
		listen := new(models.Listen)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(listen, formGroup, probe)
	case "Listening":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Listening Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListeningFormCallback(
			nil,
			probe,
			formGroup,
		)
		listening := new(models.Listening)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(listening, formGroup, probe)
	case "Lyric":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lyric Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LyricFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric := new(models.Lyric)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric, formGroup, probe)
	case "Lyric_font":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lyric_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric_font := new(models.Lyric_font)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric_font, formGroup, probe)
	case "Lyric_language":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Lyric_language Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric_language := new(models.Lyric_language)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric_language, formGroup, probe)
	case "Measure_layout":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Measure_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_layout := new(models.Measure_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_layout, formGroup, probe)
	case "Measure_numbering":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Measure_numbering Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_numberingFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_numbering := new(models.Measure_numbering)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_numbering, formGroup, probe)
	case "Measure_repeat":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Measure_repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_repeatFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_repeat := new(models.Measure_repeat)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_repeat, formGroup, probe)
	case "Measure_style":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Measure_style Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_styleFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_style := new(models.Measure_style)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_style, formGroup, probe)
	case "Membrane":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Membrane Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MembraneFormCallback(
			nil,
			probe,
			formGroup,
		)
		membrane := new(models.Membrane)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(membrane, formGroup, probe)
	case "Metal":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Metal Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetalFormCallback(
			nil,
			probe,
			formGroup,
		)
		metal := new(models.Metal)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metal, formGroup, probe)
	case "Metronome":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Metronome Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetronomeFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome := new(models.Metronome)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome, formGroup, probe)
	case "Metronome_beam":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Metronome_beam Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_beamFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_beam := new(models.Metronome_beam)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_beam, formGroup, probe)
	case "Metronome_note":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Metronome_note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_noteFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_note := new(models.Metronome_note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_note, formGroup, probe)
	case "Metronome_tied":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Metronome_tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tiedFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_tied := new(models.Metronome_tied)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_tied, formGroup, probe)
	case "Metronome_tuplet":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Metronome_tuplet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tupletFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_tuplet := new(models.Metronome_tuplet)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_tuplet, formGroup, probe)
	case "Midi_device":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Midi_device Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_deviceFormCallback(
			nil,
			probe,
			formGroup,
		)
		midi_device := new(models.Midi_device)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(midi_device, formGroup, probe)
	case "Midi_instrument":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Midi_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_instrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		midi_instrument := new(models.Midi_instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(midi_instrument, formGroup, probe)
	case "Miscellaneous":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Miscellaneous Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MiscellaneousFormCallback(
			nil,
			probe,
			formGroup,
		)
		miscellaneous := new(models.Miscellaneous)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(miscellaneous, formGroup, probe)
	case "Miscellaneous_field":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Miscellaneous_field Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		miscellaneous_field := new(models.Miscellaneous_field)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(miscellaneous_field, formGroup, probe)
	case "Mordent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Mordent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MordentFormCallback(
			nil,
			probe,
			formGroup,
		)
		mordent := new(models.Mordent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mordent, formGroup, probe)
	case "Multiple_rest":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Multiple_rest Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Multiple_restFormCallback(
			nil,
			probe,
			formGroup,
		)
		multiple_rest := new(models.Multiple_rest)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(multiple_rest, formGroup, probe)
	case "Name_display":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Name_display Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Name_displayFormCallback(
			nil,
			probe,
			formGroup,
		)
		name_display := new(models.Name_display)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(name_display, formGroup, probe)
	case "Non_arpeggiate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Non_arpeggiate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Non_arpeggiateFormCallback(
			nil,
			probe,
			formGroup,
		)
		non_arpeggiate := new(models.Non_arpeggiate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(non_arpeggiate, formGroup, probe)
	case "Notations":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Notations Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotationsFormCallback(
			nil,
			probe,
			formGroup,
		)
		notations := new(models.Notations)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notations, formGroup, probe)
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "Note_size":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Note_size Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_sizeFormCallback(
			nil,
			probe,
			formGroup,
		)
		note_size := new(models.Note_size)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note_size, formGroup, probe)
	case "Note_type":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Note_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		note_type := new(models.Note_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note_type, formGroup, probe)
	case "Notehead":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Notehead Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteheadFormCallback(
			nil,
			probe,
			formGroup,
		)
		notehead := new(models.Notehead)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notehead, formGroup, probe)
	case "Notehead_text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Notehead_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Notehead_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		notehead_text := new(models.Notehead_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notehead_text, formGroup, probe)
	case "Numeral":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Numeral Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NumeralFormCallback(
			nil,
			probe,
			formGroup,
		)
		numeral := new(models.Numeral)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(numeral, formGroup, probe)
	case "Numeral_key":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Numeral_key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_keyFormCallback(
			nil,
			probe,
			formGroup,
		)
		numeral_key := new(models.Numeral_key)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(numeral_key, formGroup, probe)
	case "Numeral_root":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Numeral_root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_rootFormCallback(
			nil,
			probe,
			formGroup,
		)
		numeral_root := new(models.Numeral_root)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(numeral_root, formGroup, probe)
	case "Octave_shift":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Octave_shift Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Octave_shiftFormCallback(
			nil,
			probe,
			formGroup,
		)
		octave_shift := new(models.Octave_shift)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(octave_shift, formGroup, probe)
	case "Offset":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Offset Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OffsetFormCallback(
			nil,
			probe,
			formGroup,
		)
		offset := new(models.Offset)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(offset, formGroup, probe)
	case "Opus":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Opus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OpusFormCallback(
			nil,
			probe,
			formGroup,
		)
		opus := new(models.Opus)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(opus, formGroup, probe)
	case "Ornaments":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Ornaments Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OrnamentsFormCallback(
			nil,
			probe,
			formGroup,
		)
		ornaments := new(models.Ornaments)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ornaments, formGroup, probe)
	case "Other_appearance":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Other_appearance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_appearanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_appearance := new(models.Other_appearance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_appearance, formGroup, probe)
	case "Other_listening":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Other_listening Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_listeningFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_listening := new(models.Other_listening)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_listening, formGroup, probe)
	case "Other_notation":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Other_notation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_notationFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_notation := new(models.Other_notation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_notation, formGroup, probe)
	case "Other_play":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Other_play Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_playFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_play := new(models.Other_play)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_play, formGroup, probe)
	case "Page_layout":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Page_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		page_layout := new(models.Page_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(page_layout, formGroup, probe)
	case "Page_margins":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Page_margins Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_marginsFormCallback(
			nil,
			probe,
			formGroup,
		)
		page_margins := new(models.Page_margins)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(page_margins, formGroup, probe)
	case "Part_clef":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Part_clef Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_clefFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_clef := new(models.Part_clef)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_clef, formGroup, probe)
	case "Part_group":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Part_group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_groupFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_group := new(models.Part_group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_group, formGroup, probe)
	case "Part_link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Part_link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_linkFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_link := new(models.Part_link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_link, formGroup, probe)
	case "Part_list":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Part_list Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_listFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_list := new(models.Part_list)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_list, formGroup, probe)
	case "Part_symbol":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Part_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_symbolFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_symbol := new(models.Part_symbol)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_symbol, formGroup, probe)
	case "Part_transpose":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Part_transpose Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_transposeFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_transpose := new(models.Part_transpose)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_transpose, formGroup, probe)
	case "Pedal":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Pedal Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PedalFormCallback(
			nil,
			probe,
			formGroup,
		)
		pedal := new(models.Pedal)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pedal, formGroup, probe)
	case "Pedal_tuning":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Pedal_tuning Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Pedal_tuningFormCallback(
			nil,
			probe,
			formGroup,
		)
		pedal_tuning := new(models.Pedal_tuning)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pedal_tuning, formGroup, probe)
	case "Percussion":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Percussion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PercussionFormCallback(
			nil,
			probe,
			formGroup,
		)
		percussion := new(models.Percussion)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(percussion, formGroup, probe)
	case "Pitch":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Pitch Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchFormCallback(
			nil,
			probe,
			formGroup,
		)
		pitch := new(models.Pitch)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pitch, formGroup, probe)
	case "Pitched":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Pitched Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchedFormCallback(
			nil,
			probe,
			formGroup,
		)
		pitched := new(models.Pitched)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pitched, formGroup, probe)
	case "Play":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Play Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayFormCallback(
			nil,
			probe,
			formGroup,
		)
		play := new(models.Play)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(play, formGroup, probe)
	case "Player":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Player Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayerFormCallback(
			nil,
			probe,
			formGroup,
		)
		player := new(models.Player)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(player, formGroup, probe)
	case "Principal_voice":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Principal_voice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Principal_voiceFormCallback(
			nil,
			probe,
			formGroup,
		)
		principal_voice := new(models.Principal_voice)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(principal_voice, formGroup, probe)
	case "Print":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Print Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PrintFormCallback(
			nil,
			probe,
			formGroup,
		)
		print := new(models.Print)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(print, formGroup, probe)
	case "Release":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Release Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ReleaseFormCallback(
			nil,
			probe,
			formGroup,
		)
		release := new(models.Release)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(release, formGroup, probe)
	case "Repeat":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepeatFormCallback(
			nil,
			probe,
			formGroup,
		)
		repeat := new(models.Repeat)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(repeat, formGroup, probe)
	case "Rest":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Rest Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestFormCallback(
			nil,
			probe,
			formGroup,
		)
		rest := new(models.Rest)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rest, formGroup, probe)
	case "Root":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RootFormCallback(
			nil,
			probe,
			formGroup,
		)
		root := new(models.Root)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(root, formGroup, probe)
	case "Root_step":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Root_step Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Root_stepFormCallback(
			nil,
			probe,
			formGroup,
		)
		root_step := new(models.Root_step)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(root_step, formGroup, probe)
	case "Scaling":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Scaling Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScalingFormCallback(
			nil,
			probe,
			formGroup,
		)
		scaling := new(models.Scaling)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scaling, formGroup, probe)
	case "Scordatura":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Scordatura Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScordaturaFormCallback(
			nil,
			probe,
			formGroup,
		)
		scordatura := new(models.Scordatura)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scordatura, formGroup, probe)
	case "Score_instrument":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Score_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_instrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_instrument := new(models.Score_instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_instrument, formGroup, probe)
	case "Score_part":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Score_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_part := new(models.Score_part)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_part, formGroup, probe)
	case "Score_partwise":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Score_partwise Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partwiseFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_partwise := new(models.Score_partwise)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_partwise, formGroup, probe)
	case "Score_timewise":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Score_timewise Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_timewiseFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_timewise := new(models.Score_timewise)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_timewise, formGroup, probe)
	case "Segno":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Segno Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SegnoFormCallback(
			nil,
			probe,
			formGroup,
		)
		segno := new(models.Segno)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(segno, formGroup, probe)
	case "Slash":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Slash Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlashFormCallback(
			nil,
			probe,
			formGroup,
		)
		slash := new(models.Slash)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slash, formGroup, probe)
	case "Slide":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Slide Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlideFormCallback(
			nil,
			probe,
			formGroup,
		)
		slide := new(models.Slide)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slide, formGroup, probe)
	case "Slur":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Slur Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlurFormCallback(
			nil,
			probe,
			formGroup,
		)
		slur := new(models.Slur)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slur, formGroup, probe)
	case "Sound":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Sound Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SoundFormCallback(
			nil,
			probe,
			formGroup,
		)
		sound := new(models.Sound)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sound, formGroup, probe)
	case "Staff_details":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Staff_details Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_detailsFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_details := new(models.Staff_details)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_details, formGroup, probe)
	case "Staff_divide":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Staff_divide Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_divideFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_divide := new(models.Staff_divide)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_divide, formGroup, probe)
	case "Staff_layout":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Staff_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_layout := new(models.Staff_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_layout, formGroup, probe)
	case "Staff_size":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Staff_size Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_sizeFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_size := new(models.Staff_size)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_size, formGroup, probe)
	case "Staff_tuning":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Staff_tuning Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_tuningFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_tuning := new(models.Staff_tuning)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_tuning, formGroup, probe)
	case "Stem":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Stem Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StemFormCallback(
			nil,
			probe,
			formGroup,
		)
		stem := new(models.Stem)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stem, formGroup, probe)
	case "Stick":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Stick Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StickFormCallback(
			nil,
			probe,
			formGroup,
		)
		stick := new(models.Stick)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stick, formGroup, probe)
	case "String_mute":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "String_mute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__String_muteFormCallback(
			nil,
			probe,
			formGroup,
		)
		string_mute := new(models.String_mute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(string_mute, formGroup, probe)
	case "Strong_accent":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Strong_accent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Strong_accentFormCallback(
			nil,
			probe,
			formGroup,
		)
		strong_accent := new(models.Strong_accent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(strong_accent, formGroup, probe)
	case "Supports":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Supports Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SupportsFormCallback(
			nil,
			probe,
			formGroup,
		)
		supports := new(models.Supports)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(supports, formGroup, probe)
	case "Swing":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Swing Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SwingFormCallback(
			nil,
			probe,
			formGroup,
		)
		swing := new(models.Swing)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(swing, formGroup, probe)
	case "Sync":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Sync Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SyncFormCallback(
			nil,
			probe,
			formGroup,
		)
		sync := new(models.Sync)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sync, formGroup, probe)
	case "System_dividers":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "System_dividers Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_dividersFormCallback(
			nil,
			probe,
			formGroup,
		)
		system_dividers := new(models.System_dividers)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system_dividers, formGroup, probe)
	case "System_layout":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "System_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		system_layout := new(models.System_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system_layout, formGroup, probe)
	case "System_margins":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "System_margins Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_marginsFormCallback(
			nil,
			probe,
			formGroup,
		)
		system_margins := new(models.System_margins)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system_margins, formGroup, probe)
	case "Tap":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tap Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TapFormCallback(
			nil,
			probe,
			formGroup,
		)
		tap := new(models.Tap)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tap, formGroup, probe)
	case "Technical":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Technical Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TechnicalFormCallback(
			nil,
			probe,
			formGroup,
		)
		technical := new(models.Technical)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(technical, formGroup, probe)
	case "Text_element_data":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Text_element_data Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Text_element_dataFormCallback(
			nil,
			probe,
			formGroup,
		)
		text_element_data := new(models.Text_element_data)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(text_element_data, formGroup, probe)
	case "Tie":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tie Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TieFormCallback(
			nil,
			probe,
			formGroup,
		)
		tie := new(models.Tie)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tie, formGroup, probe)
	case "Tied":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TiedFormCallback(
			nil,
			probe,
			formGroup,
		)
		tied := new(models.Tied)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tied, formGroup, probe)
	case "Time":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Time Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimeFormCallback(
			nil,
			probe,
			formGroup,
		)
		time := new(models.Time)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(time, formGroup, probe)
	case "Time_modification":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Time_modification Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Time_modificationFormCallback(
			nil,
			probe,
			formGroup,
		)
		time_modification := new(models.Time_modification)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(time_modification, formGroup, probe)
	case "Timpani":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Timpani Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimpaniFormCallback(
			nil,
			probe,
			formGroup,
		)
		timpani := new(models.Timpani)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(timpani, formGroup, probe)
	case "Transpose":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Transpose Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransposeFormCallback(
			nil,
			probe,
			formGroup,
		)
		transpose := new(models.Transpose)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(transpose, formGroup, probe)
	case "Tremolo":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tremolo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TremoloFormCallback(
			nil,
			probe,
			formGroup,
		)
		tremolo := new(models.Tremolo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tremolo, formGroup, probe)
	case "Tuplet":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tuplet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TupletFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet := new(models.Tuplet)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet, formGroup, probe)
	case "Tuplet_dot":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tuplet_dot Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_dotFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_dot := new(models.Tuplet_dot)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_dot, formGroup, probe)
	case "Tuplet_number":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tuplet_number Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_numberFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_number := new(models.Tuplet_number)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_number, formGroup, probe)
	case "Tuplet_portion":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tuplet_portion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_portionFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_portion := new(models.Tuplet_portion)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_portion, formGroup, probe)
	case "Tuplet_type":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Tuplet_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_type := new(models.Tuplet_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_type, formGroup, probe)
	case "Typed_text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Typed_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Typed_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		typed_text := new(models.Typed_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(typed_text, formGroup, probe)
	case "Unpitched":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Unpitched Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnpitchedFormCallback(
			nil,
			probe,
			formGroup,
		)
		unpitched := new(models.Unpitched)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(unpitched, formGroup, probe)
	case "Virtual_instrument":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Virtual_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Virtual_instrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		virtual_instrument := new(models.Virtual_instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(virtual_instrument, formGroup, probe)
	case "Wait":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Wait Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WaitFormCallback(
			nil,
			probe,
			formGroup,
		)
		wait := new(models.Wait)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wait, formGroup, probe)
	case "Wavy_line":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Wavy_line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Wavy_lineFormCallback(
			nil,
			probe,
			formGroup,
		)
		wavy_line := new(models.Wavy_line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wavy_line, formGroup, probe)
	case "Wedge":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Wedge Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WedgeFormCallback(
			nil,
			probe,
			formGroup,
		)
		wedge := new(models.Wedge)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wedge, formGroup, probe)
	case "Wood":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Wood Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WoodFormCallback(
			nil,
			probe,
			formGroup,
		)
		wood := new(models.Wood)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wood, formGroup, probe)
	case "Work":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Work Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkFormCallback(
			nil,
			probe,
			formGroup,
		)
		work := new(models.Work)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(work, formGroup, probe)
	}
	formStage.Commit()
}
