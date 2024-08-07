// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongmusicxml/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	FillUpNamedFormFromGongstruct[T](instance, probe, formStage, gongtable.FormGroupDefaultName.ToString())

}

func FillUpNamedFormFromGongstruct[T models.Gongstruct](instance *T, probe *Probe, formStage *gongtable.StageStruct, formName string) {

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Accidental:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Accidental Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccidentalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accidental_mark:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Accidental_mark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_markFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accidental_text:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Accidental_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accord:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Accord Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccordFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Accordion_registration:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Accordion_registration Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accordion_registrationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.AnyType:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "AnyType Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AnyTypeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Appearance:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Appearance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AppearanceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Arpeggiate:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Arpeggiate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArpeggiateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Arrow:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Arrow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArrowFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Articulations:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Articulations Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArticulationsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Assess:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Assess Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AssessFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Attributes:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Attributes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Backup:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Backup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BackupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bar_style_color:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bar_style_color Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bar_style_colorFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Barline:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Barline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarlineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Barre:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Barre Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarreFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bass:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BassFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bass_step:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bass_step Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bass_stepFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beam:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Beam Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeamFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beat_repeat:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Beat_repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_repeatFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beat_unit_tied:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Beat_unit_tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_unit_tiedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Beater:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Beater Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeaterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bend:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bend Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BendFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bookmark:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bookmark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookmarkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Bracket:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Bracket Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BracketFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Breath_mark:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Breath_mark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Breath_markFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Caesura:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Caesura Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CaesuraFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Cancel:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Cancel Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CancelFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Clef:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Clef Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClefFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Coda:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Coda Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CodaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Credit:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Credit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CreditFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Dashes:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Dashes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DashesFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Defaults:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Defaults Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DefaultsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Degree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DegreeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree_alter:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Degree_alter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_alterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Degree_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Degree_value:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Degree_value Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_valueFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Direction:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Direction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DirectionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Direction_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Direction_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Direction_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Distance:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Distance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DistanceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Double:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Double Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DoubleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Dynamics:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Dynamics Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DynamicsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Effect:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Effect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EffectFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Elision:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Elision Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElisionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EmptyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_font:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_fontFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_line:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_lineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_placement:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_placement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placementFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_placement_smufl:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_placement_smufl Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placement_smuflFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_object_style_align:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_print_object_style_align Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_object_style_alignFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_style:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_print_style Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_styleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_style_align:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_print_style_align Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_alignFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_print_style_align_id:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_print_style_align_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_align_idFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Empty_trill_sound:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Empty_trill_sound Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_trill_soundFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Encoding:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Encoding Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EncodingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Ending:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Ending Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Extend:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Extend Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtendFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Feature:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Feature Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FeatureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fermata:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Fermata Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FermataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Figure:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Figure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FigureFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Figured_bass:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Figured_bass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Figured_bassFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fingering:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Fingering Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FingeringFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.First_fret:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "First_fret Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__First_fretFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Foo:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Foo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FooFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.For_part:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "For_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__For_partFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Formatted_symbol:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Formatted_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbolFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Formatted_symbol_id:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Formatted_symbol_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbol_idFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Forward:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Forward Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ForwardFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Frame:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Frame Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrameFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Frame_note:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Frame_note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Frame_noteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Fret:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Fret Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FretFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Glass:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Glass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlassFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Glissando:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Glissando Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlissandoFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Glyph:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Glyph Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlyphFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Grace:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Grace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GraceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group_barline:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Group_barline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_barlineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Group_symbol:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Group_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_symbolFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Grouping:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Grouping Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Hammer_on_pull_off:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Hammer_on_pull_off Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hammer_on_pull_offFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Handbell:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Handbell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HandbellFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmon_closed:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Harmon_closed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_closedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmon_mute:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Harmon_mute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_muteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmonic:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Harmonic Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonicFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmony:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Harmony Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harmony_alter:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Harmony_alter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmony_alterFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Harp_pedals:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Harp_pedals Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harp_pedalsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Heel_toe:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Heel_toe Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Heel_toeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Hole:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Hole Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HoleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Hole_closed:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Hole_closed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hole_closedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Horizontal_turn:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Horizontal_turn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Horizontal_turnFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Identification:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Identification Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__IdentificationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Image:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Image Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ImageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Instrument:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InstrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Instrument_change:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Instrument_change Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_changeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Instrument_link:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Instrument_link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_linkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Interchangeable:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Interchangeable Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InterchangeableFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Inversion:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Inversion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InversionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Key:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Key_accidental:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Key_accidental Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_accidentalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Key_octave:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Key_octave Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_octaveFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Kind:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Kind Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KindFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Level:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Level Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LevelFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line_detail:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Line_detail Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_detailFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Line_width:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Line_width Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_widthFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Link:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Listen:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Listen Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListenFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Listening:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Listening Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListeningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lyric Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LyricFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric_font:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lyric_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Lyric_language:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Lyric_language Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_layout:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Measure_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_numbering:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Measure_numbering Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_numberingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_repeat:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Measure_repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_repeatFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Measure_style:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Measure_style Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_styleFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Membrane:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Membrane Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MembraneFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metal:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Metal Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Metronome Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetronomeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_beam:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Metronome_beam Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_beamFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_note:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Metronome_note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_noteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_tied:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Metronome_tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tiedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Metronome_tuplet:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Metronome_tuplet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tupletFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Midi_device:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Midi_device Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_deviceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Midi_instrument:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Midi_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_instrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Miscellaneous:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Miscellaneous Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MiscellaneousFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Miscellaneous_field:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Miscellaneous_field Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Mordent:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Mordent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MordentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Multiple_rest:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Multiple_rest Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Multiple_restFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Name_display:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Name_display Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Name_displayFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Non_arpeggiate:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Non_arpeggiate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Non_arpeggiateFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Notations:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Notations Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotationsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note_size:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Note_size Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_sizeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Note_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Note_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Notehead:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Notehead Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteheadFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Notehead_text:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Notehead_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Notehead_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Numeral:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Numeral Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NumeralFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Numeral_key:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Numeral_key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_keyFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Numeral_root:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Numeral_root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_rootFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Octave_shift:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Octave_shift Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Octave_shiftFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Offset:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Offset Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OffsetFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Opus:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Opus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OpusFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Ornaments:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Ornaments Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OrnamentsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_appearance:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Other_appearance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_appearanceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_listening:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Other_listening Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_listeningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_notation:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Other_notation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_notationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Other_play:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Other_play Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_playFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Page_layout:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Page_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Page_margins:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Page_margins Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_marginsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_clef:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Part_clef Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_clefFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_group:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Part_group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_groupFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_link:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Part_link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_linkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_list:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Part_list Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_listFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_symbol:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Part_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_symbolFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Part_transpose:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Part_transpose Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_transposeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pedal:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Pedal Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PedalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pedal_tuning:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Pedal_tuning Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Pedal_tuningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Percussion:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Percussion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PercussionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pitch:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Pitch Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Pitched:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Pitched Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Play:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Play Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Player:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Player Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayerFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Principal_voice:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Principal_voice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Principal_voiceFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Print:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Print Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PrintFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Release:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Release Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ReleaseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Repeat:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepeatFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Rest:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Rest Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Root:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RootFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Root_step:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Root_step Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Root_stepFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scaling:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Scaling Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScalingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Scordatura:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Scordatura Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScordaturaFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_instrument:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Score_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_instrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_part:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Score_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_partwise:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Score_partwise Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partwiseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Score_timewise:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Score_timewise Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_timewiseFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Segno:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Segno Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SegnoFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slash:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Slash Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlashFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slide:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Slide Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlideFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Slur:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Slur Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlurFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Sound:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Sound Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SoundFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_details:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Staff_details Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_detailsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_divide:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Staff_divide Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_divideFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_layout:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Staff_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_size:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Staff_size Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_sizeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Staff_tuning:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Staff_tuning Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_tuningFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Stem:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Stem Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StemFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Stick:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Stick Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StickFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.String_mute:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "String_mute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__String_muteFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Strong_accent:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Strong_accent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Strong_accentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Supports:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Supports Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SupportsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Swing:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Swing Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SwingFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Sync:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Sync Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SyncFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System_dividers:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "System_dividers Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_dividersFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System_layout:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "System_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_layoutFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.System_margins:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "System_margins Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_marginsFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tap:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tap Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TapFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Technical:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Technical Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TechnicalFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Text_element_data:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Text_element_data Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Text_element_dataFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tie:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tie Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TieFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tied:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TiedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Time:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Time Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Time_modification:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Time_modification Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Time_modificationFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Timpani:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Timpani Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimpaniFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Transpose:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Transpose Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransposeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tremolo:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tremolo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TremoloFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tuplet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TupletFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_dot:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tuplet_dot Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_dotFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_number:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tuplet_number Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_numberFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_portion:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tuplet_portion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_portionFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Tuplet_type:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Tuplet_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_typeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Typed_text:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Typed_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Typed_textFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Unpitched:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Unpitched Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnpitchedFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Virtual_instrument:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Virtual_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Virtual_instrumentFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wait:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Wait Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WaitFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wavy_line:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Wavy_line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Wavy_lineFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wedge:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Wedge Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WedgeFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Wood:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Wood Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WoodFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	case *models.Work:
		formGroup := (&gongtable.FormGroup{
			Name:  formName,
			Label: "Work Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkFormCallback(
			instancesTyped,
			probe,
			formGroup,
		)
		formGroup.HasSuppressButton = true
		FillUpForm(instancesTyped, formGroup, probe)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
