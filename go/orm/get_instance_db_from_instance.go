// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongmusicxml/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Accidental:
		accidentalInstance := any(concreteInstance).(*models.Accidental)
		ret2 := backRepo.BackRepoAccidental.GetAccidentalDBFromAccidentalPtr(accidentalInstance)
		ret = any(ret2).(*T2)
	case *models.Accidental_mark:
		accidental_markInstance := any(concreteInstance).(*models.Accidental_mark)
		ret2 := backRepo.BackRepoAccidental_mark.GetAccidental_markDBFromAccidental_markPtr(accidental_markInstance)
		ret = any(ret2).(*T2)
	case *models.Accidental_text:
		accidental_textInstance := any(concreteInstance).(*models.Accidental_text)
		ret2 := backRepo.BackRepoAccidental_text.GetAccidental_textDBFromAccidental_textPtr(accidental_textInstance)
		ret = any(ret2).(*T2)
	case *models.Accord:
		accordInstance := any(concreteInstance).(*models.Accord)
		ret2 := backRepo.BackRepoAccord.GetAccordDBFromAccordPtr(accordInstance)
		ret = any(ret2).(*T2)
	case *models.Accordion_registration:
		accordion_registrationInstance := any(concreteInstance).(*models.Accordion_registration)
		ret2 := backRepo.BackRepoAccordion_registration.GetAccordion_registrationDBFromAccordion_registrationPtr(accordion_registrationInstance)
		ret = any(ret2).(*T2)
	case *models.AnyType:
		anytypeInstance := any(concreteInstance).(*models.AnyType)
		ret2 := backRepo.BackRepoAnyType.GetAnyTypeDBFromAnyTypePtr(anytypeInstance)
		ret = any(ret2).(*T2)
	case *models.Appearance:
		appearanceInstance := any(concreteInstance).(*models.Appearance)
		ret2 := backRepo.BackRepoAppearance.GetAppearanceDBFromAppearancePtr(appearanceInstance)
		ret = any(ret2).(*T2)
	case *models.Arpeggiate:
		arpeggiateInstance := any(concreteInstance).(*models.Arpeggiate)
		ret2 := backRepo.BackRepoArpeggiate.GetArpeggiateDBFromArpeggiatePtr(arpeggiateInstance)
		ret = any(ret2).(*T2)
	case *models.Arrow:
		arrowInstance := any(concreteInstance).(*models.Arrow)
		ret2 := backRepo.BackRepoArrow.GetArrowDBFromArrowPtr(arrowInstance)
		ret = any(ret2).(*T2)
	case *models.Articulations:
		articulationsInstance := any(concreteInstance).(*models.Articulations)
		ret2 := backRepo.BackRepoArticulations.GetArticulationsDBFromArticulationsPtr(articulationsInstance)
		ret = any(ret2).(*T2)
	case *models.Assess:
		assessInstance := any(concreteInstance).(*models.Assess)
		ret2 := backRepo.BackRepoAssess.GetAssessDBFromAssessPtr(assessInstance)
		ret = any(ret2).(*T2)
	case *models.Attributes:
		attributesInstance := any(concreteInstance).(*models.Attributes)
		ret2 := backRepo.BackRepoAttributes.GetAttributesDBFromAttributesPtr(attributesInstance)
		ret = any(ret2).(*T2)
	case *models.Backup:
		backupInstance := any(concreteInstance).(*models.Backup)
		ret2 := backRepo.BackRepoBackup.GetBackupDBFromBackupPtr(backupInstance)
		ret = any(ret2).(*T2)
	case *models.Bar_style_color:
		bar_style_colorInstance := any(concreteInstance).(*models.Bar_style_color)
		ret2 := backRepo.BackRepoBar_style_color.GetBar_style_colorDBFromBar_style_colorPtr(bar_style_colorInstance)
		ret = any(ret2).(*T2)
	case *models.Barline:
		barlineInstance := any(concreteInstance).(*models.Barline)
		ret2 := backRepo.BackRepoBarline.GetBarlineDBFromBarlinePtr(barlineInstance)
		ret = any(ret2).(*T2)
	case *models.Barre:
		barreInstance := any(concreteInstance).(*models.Barre)
		ret2 := backRepo.BackRepoBarre.GetBarreDBFromBarrePtr(barreInstance)
		ret = any(ret2).(*T2)
	case *models.Bass:
		bassInstance := any(concreteInstance).(*models.Bass)
		ret2 := backRepo.BackRepoBass.GetBassDBFromBassPtr(bassInstance)
		ret = any(ret2).(*T2)
	case *models.Bass_step:
		bass_stepInstance := any(concreteInstance).(*models.Bass_step)
		ret2 := backRepo.BackRepoBass_step.GetBass_stepDBFromBass_stepPtr(bass_stepInstance)
		ret = any(ret2).(*T2)
	case *models.Beam:
		beamInstance := any(concreteInstance).(*models.Beam)
		ret2 := backRepo.BackRepoBeam.GetBeamDBFromBeamPtr(beamInstance)
		ret = any(ret2).(*T2)
	case *models.Beat_repeat:
		beat_repeatInstance := any(concreteInstance).(*models.Beat_repeat)
		ret2 := backRepo.BackRepoBeat_repeat.GetBeat_repeatDBFromBeat_repeatPtr(beat_repeatInstance)
		ret = any(ret2).(*T2)
	case *models.Beat_unit_tied:
		beat_unit_tiedInstance := any(concreteInstance).(*models.Beat_unit_tied)
		ret2 := backRepo.BackRepoBeat_unit_tied.GetBeat_unit_tiedDBFromBeat_unit_tiedPtr(beat_unit_tiedInstance)
		ret = any(ret2).(*T2)
	case *models.Beater:
		beaterInstance := any(concreteInstance).(*models.Beater)
		ret2 := backRepo.BackRepoBeater.GetBeaterDBFromBeaterPtr(beaterInstance)
		ret = any(ret2).(*T2)
	case *models.Bend:
		bendInstance := any(concreteInstance).(*models.Bend)
		ret2 := backRepo.BackRepoBend.GetBendDBFromBendPtr(bendInstance)
		ret = any(ret2).(*T2)
	case *models.Bookmark:
		bookmarkInstance := any(concreteInstance).(*models.Bookmark)
		ret2 := backRepo.BackRepoBookmark.GetBookmarkDBFromBookmarkPtr(bookmarkInstance)
		ret = any(ret2).(*T2)
	case *models.Bracket:
		bracketInstance := any(concreteInstance).(*models.Bracket)
		ret2 := backRepo.BackRepoBracket.GetBracketDBFromBracketPtr(bracketInstance)
		ret = any(ret2).(*T2)
	case *models.Breath_mark:
		breath_markInstance := any(concreteInstance).(*models.Breath_mark)
		ret2 := backRepo.BackRepoBreath_mark.GetBreath_markDBFromBreath_markPtr(breath_markInstance)
		ret = any(ret2).(*T2)
	case *models.Caesura:
		caesuraInstance := any(concreteInstance).(*models.Caesura)
		ret2 := backRepo.BackRepoCaesura.GetCaesuraDBFromCaesuraPtr(caesuraInstance)
		ret = any(ret2).(*T2)
	case *models.Cancel:
		cancelInstance := any(concreteInstance).(*models.Cancel)
		ret2 := backRepo.BackRepoCancel.GetCancelDBFromCancelPtr(cancelInstance)
		ret = any(ret2).(*T2)
	case *models.Clef:
		clefInstance := any(concreteInstance).(*models.Clef)
		ret2 := backRepo.BackRepoClef.GetClefDBFromClefPtr(clefInstance)
		ret = any(ret2).(*T2)
	case *models.Coda:
		codaInstance := any(concreteInstance).(*models.Coda)
		ret2 := backRepo.BackRepoCoda.GetCodaDBFromCodaPtr(codaInstance)
		ret = any(ret2).(*T2)
	case *models.Credit:
		creditInstance := any(concreteInstance).(*models.Credit)
		ret2 := backRepo.BackRepoCredit.GetCreditDBFromCreditPtr(creditInstance)
		ret = any(ret2).(*T2)
	case *models.Dashes:
		dashesInstance := any(concreteInstance).(*models.Dashes)
		ret2 := backRepo.BackRepoDashes.GetDashesDBFromDashesPtr(dashesInstance)
		ret = any(ret2).(*T2)
	case *models.Defaults:
		defaultsInstance := any(concreteInstance).(*models.Defaults)
		ret2 := backRepo.BackRepoDefaults.GetDefaultsDBFromDefaultsPtr(defaultsInstance)
		ret = any(ret2).(*T2)
	case *models.Degree:
		degreeInstance := any(concreteInstance).(*models.Degree)
		ret2 := backRepo.BackRepoDegree.GetDegreeDBFromDegreePtr(degreeInstance)
		ret = any(ret2).(*T2)
	case *models.Degree_alter:
		degree_alterInstance := any(concreteInstance).(*models.Degree_alter)
		ret2 := backRepo.BackRepoDegree_alter.GetDegree_alterDBFromDegree_alterPtr(degree_alterInstance)
		ret = any(ret2).(*T2)
	case *models.Degree_type:
		degree_typeInstance := any(concreteInstance).(*models.Degree_type)
		ret2 := backRepo.BackRepoDegree_type.GetDegree_typeDBFromDegree_typePtr(degree_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Degree_value:
		degree_valueInstance := any(concreteInstance).(*models.Degree_value)
		ret2 := backRepo.BackRepoDegree_value.GetDegree_valueDBFromDegree_valuePtr(degree_valueInstance)
		ret = any(ret2).(*T2)
	case *models.Direction:
		directionInstance := any(concreteInstance).(*models.Direction)
		ret2 := backRepo.BackRepoDirection.GetDirectionDBFromDirectionPtr(directionInstance)
		ret = any(ret2).(*T2)
	case *models.Direction_type:
		direction_typeInstance := any(concreteInstance).(*models.Direction_type)
		ret2 := backRepo.BackRepoDirection_type.GetDirection_typeDBFromDirection_typePtr(direction_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Distance:
		distanceInstance := any(concreteInstance).(*models.Distance)
		ret2 := backRepo.BackRepoDistance.GetDistanceDBFromDistancePtr(distanceInstance)
		ret = any(ret2).(*T2)
	case *models.Double:
		doubleInstance := any(concreteInstance).(*models.Double)
		ret2 := backRepo.BackRepoDouble.GetDoubleDBFromDoublePtr(doubleInstance)
		ret = any(ret2).(*T2)
	case *models.Dynamics:
		dynamicsInstance := any(concreteInstance).(*models.Dynamics)
		ret2 := backRepo.BackRepoDynamics.GetDynamicsDBFromDynamicsPtr(dynamicsInstance)
		ret = any(ret2).(*T2)
	case *models.Effect:
		effectInstance := any(concreteInstance).(*models.Effect)
		ret2 := backRepo.BackRepoEffect.GetEffectDBFromEffectPtr(effectInstance)
		ret = any(ret2).(*T2)
	case *models.Elision:
		elisionInstance := any(concreteInstance).(*models.Elision)
		ret2 := backRepo.BackRepoElision.GetElisionDBFromElisionPtr(elisionInstance)
		ret = any(ret2).(*T2)
	case *models.Empty:
		emptyInstance := any(concreteInstance).(*models.Empty)
		ret2 := backRepo.BackRepoEmpty.GetEmptyDBFromEmptyPtr(emptyInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_font:
		empty_fontInstance := any(concreteInstance).(*models.Empty_font)
		ret2 := backRepo.BackRepoEmpty_font.GetEmpty_fontDBFromEmpty_fontPtr(empty_fontInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_line:
		empty_lineInstance := any(concreteInstance).(*models.Empty_line)
		ret2 := backRepo.BackRepoEmpty_line.GetEmpty_lineDBFromEmpty_linePtr(empty_lineInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_placement:
		empty_placementInstance := any(concreteInstance).(*models.Empty_placement)
		ret2 := backRepo.BackRepoEmpty_placement.GetEmpty_placementDBFromEmpty_placementPtr(empty_placementInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_placement_smufl:
		empty_placement_smuflInstance := any(concreteInstance).(*models.Empty_placement_smufl)
		ret2 := backRepo.BackRepoEmpty_placement_smufl.GetEmpty_placement_smuflDBFromEmpty_placement_smuflPtr(empty_placement_smuflInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_print_object_style_align:
		empty_print_object_style_alignInstance := any(concreteInstance).(*models.Empty_print_object_style_align)
		ret2 := backRepo.BackRepoEmpty_print_object_style_align.GetEmpty_print_object_style_alignDBFromEmpty_print_object_style_alignPtr(empty_print_object_style_alignInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_print_style:
		empty_print_styleInstance := any(concreteInstance).(*models.Empty_print_style)
		ret2 := backRepo.BackRepoEmpty_print_style.GetEmpty_print_styleDBFromEmpty_print_stylePtr(empty_print_styleInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_print_style_align:
		empty_print_style_alignInstance := any(concreteInstance).(*models.Empty_print_style_align)
		ret2 := backRepo.BackRepoEmpty_print_style_align.GetEmpty_print_style_alignDBFromEmpty_print_style_alignPtr(empty_print_style_alignInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_print_style_align_id:
		empty_print_style_align_idInstance := any(concreteInstance).(*models.Empty_print_style_align_id)
		ret2 := backRepo.BackRepoEmpty_print_style_align_id.GetEmpty_print_style_align_idDBFromEmpty_print_style_align_idPtr(empty_print_style_align_idInstance)
		ret = any(ret2).(*T2)
	case *models.Empty_trill_sound:
		empty_trill_soundInstance := any(concreteInstance).(*models.Empty_trill_sound)
		ret2 := backRepo.BackRepoEmpty_trill_sound.GetEmpty_trill_soundDBFromEmpty_trill_soundPtr(empty_trill_soundInstance)
		ret = any(ret2).(*T2)
	case *models.Encoding:
		encodingInstance := any(concreteInstance).(*models.Encoding)
		ret2 := backRepo.BackRepoEncoding.GetEncodingDBFromEncodingPtr(encodingInstance)
		ret = any(ret2).(*T2)
	case *models.Ending:
		endingInstance := any(concreteInstance).(*models.Ending)
		ret2 := backRepo.BackRepoEnding.GetEndingDBFromEndingPtr(endingInstance)
		ret = any(ret2).(*T2)
	case *models.Extend:
		extendInstance := any(concreteInstance).(*models.Extend)
		ret2 := backRepo.BackRepoExtend.GetExtendDBFromExtendPtr(extendInstance)
		ret = any(ret2).(*T2)
	case *models.Feature:
		featureInstance := any(concreteInstance).(*models.Feature)
		ret2 := backRepo.BackRepoFeature.GetFeatureDBFromFeaturePtr(featureInstance)
		ret = any(ret2).(*T2)
	case *models.Fermata:
		fermataInstance := any(concreteInstance).(*models.Fermata)
		ret2 := backRepo.BackRepoFermata.GetFermataDBFromFermataPtr(fermataInstance)
		ret = any(ret2).(*T2)
	case *models.Figure:
		figureInstance := any(concreteInstance).(*models.Figure)
		ret2 := backRepo.BackRepoFigure.GetFigureDBFromFigurePtr(figureInstance)
		ret = any(ret2).(*T2)
	case *models.Figured_bass:
		figured_bassInstance := any(concreteInstance).(*models.Figured_bass)
		ret2 := backRepo.BackRepoFigured_bass.GetFigured_bassDBFromFigured_bassPtr(figured_bassInstance)
		ret = any(ret2).(*T2)
	case *models.Fingering:
		fingeringInstance := any(concreteInstance).(*models.Fingering)
		ret2 := backRepo.BackRepoFingering.GetFingeringDBFromFingeringPtr(fingeringInstance)
		ret = any(ret2).(*T2)
	case *models.First_fret:
		first_fretInstance := any(concreteInstance).(*models.First_fret)
		ret2 := backRepo.BackRepoFirst_fret.GetFirst_fretDBFromFirst_fretPtr(first_fretInstance)
		ret = any(ret2).(*T2)
	case *models.Foo:
		fooInstance := any(concreteInstance).(*models.Foo)
		ret2 := backRepo.BackRepoFoo.GetFooDBFromFooPtr(fooInstance)
		ret = any(ret2).(*T2)
	case *models.For_part:
		for_partInstance := any(concreteInstance).(*models.For_part)
		ret2 := backRepo.BackRepoFor_part.GetFor_partDBFromFor_partPtr(for_partInstance)
		ret = any(ret2).(*T2)
	case *models.Formatted_symbol:
		formatted_symbolInstance := any(concreteInstance).(*models.Formatted_symbol)
		ret2 := backRepo.BackRepoFormatted_symbol.GetFormatted_symbolDBFromFormatted_symbolPtr(formatted_symbolInstance)
		ret = any(ret2).(*T2)
	case *models.Formatted_symbol_id:
		formatted_symbol_idInstance := any(concreteInstance).(*models.Formatted_symbol_id)
		ret2 := backRepo.BackRepoFormatted_symbol_id.GetFormatted_symbol_idDBFromFormatted_symbol_idPtr(formatted_symbol_idInstance)
		ret = any(ret2).(*T2)
	case *models.Forward:
		forwardInstance := any(concreteInstance).(*models.Forward)
		ret2 := backRepo.BackRepoForward.GetForwardDBFromForwardPtr(forwardInstance)
		ret = any(ret2).(*T2)
	case *models.Frame:
		frameInstance := any(concreteInstance).(*models.Frame)
		ret2 := backRepo.BackRepoFrame.GetFrameDBFromFramePtr(frameInstance)
		ret = any(ret2).(*T2)
	case *models.Frame_note:
		frame_noteInstance := any(concreteInstance).(*models.Frame_note)
		ret2 := backRepo.BackRepoFrame_note.GetFrame_noteDBFromFrame_notePtr(frame_noteInstance)
		ret = any(ret2).(*T2)
	case *models.Fret:
		fretInstance := any(concreteInstance).(*models.Fret)
		ret2 := backRepo.BackRepoFret.GetFretDBFromFretPtr(fretInstance)
		ret = any(ret2).(*T2)
	case *models.Glass:
		glassInstance := any(concreteInstance).(*models.Glass)
		ret2 := backRepo.BackRepoGlass.GetGlassDBFromGlassPtr(glassInstance)
		ret = any(ret2).(*T2)
	case *models.Glissando:
		glissandoInstance := any(concreteInstance).(*models.Glissando)
		ret2 := backRepo.BackRepoGlissando.GetGlissandoDBFromGlissandoPtr(glissandoInstance)
		ret = any(ret2).(*T2)
	case *models.Glyph:
		glyphInstance := any(concreteInstance).(*models.Glyph)
		ret2 := backRepo.BackRepoGlyph.GetGlyphDBFromGlyphPtr(glyphInstance)
		ret = any(ret2).(*T2)
	case *models.Grace:
		graceInstance := any(concreteInstance).(*models.Grace)
		ret2 := backRepo.BackRepoGrace.GetGraceDBFromGracePtr(graceInstance)
		ret = any(ret2).(*T2)
	case *models.Group_barline:
		group_barlineInstance := any(concreteInstance).(*models.Group_barline)
		ret2 := backRepo.BackRepoGroup_barline.GetGroup_barlineDBFromGroup_barlinePtr(group_barlineInstance)
		ret = any(ret2).(*T2)
	case *models.Group_symbol:
		group_symbolInstance := any(concreteInstance).(*models.Group_symbol)
		ret2 := backRepo.BackRepoGroup_symbol.GetGroup_symbolDBFromGroup_symbolPtr(group_symbolInstance)
		ret = any(ret2).(*T2)
	case *models.Grouping:
		groupingInstance := any(concreteInstance).(*models.Grouping)
		ret2 := backRepo.BackRepoGrouping.GetGroupingDBFromGroupingPtr(groupingInstance)
		ret = any(ret2).(*T2)
	case *models.Hammer_on_pull_off:
		hammer_on_pull_offInstance := any(concreteInstance).(*models.Hammer_on_pull_off)
		ret2 := backRepo.BackRepoHammer_on_pull_off.GetHammer_on_pull_offDBFromHammer_on_pull_offPtr(hammer_on_pull_offInstance)
		ret = any(ret2).(*T2)
	case *models.Handbell:
		handbellInstance := any(concreteInstance).(*models.Handbell)
		ret2 := backRepo.BackRepoHandbell.GetHandbellDBFromHandbellPtr(handbellInstance)
		ret = any(ret2).(*T2)
	case *models.Harmon_closed:
		harmon_closedInstance := any(concreteInstance).(*models.Harmon_closed)
		ret2 := backRepo.BackRepoHarmon_closed.GetHarmon_closedDBFromHarmon_closedPtr(harmon_closedInstance)
		ret = any(ret2).(*T2)
	case *models.Harmon_mute:
		harmon_muteInstance := any(concreteInstance).(*models.Harmon_mute)
		ret2 := backRepo.BackRepoHarmon_mute.GetHarmon_muteDBFromHarmon_mutePtr(harmon_muteInstance)
		ret = any(ret2).(*T2)
	case *models.Harmonic:
		harmonicInstance := any(concreteInstance).(*models.Harmonic)
		ret2 := backRepo.BackRepoHarmonic.GetHarmonicDBFromHarmonicPtr(harmonicInstance)
		ret = any(ret2).(*T2)
	case *models.Harmony:
		harmonyInstance := any(concreteInstance).(*models.Harmony)
		ret2 := backRepo.BackRepoHarmony.GetHarmonyDBFromHarmonyPtr(harmonyInstance)
		ret = any(ret2).(*T2)
	case *models.Harmony_alter:
		harmony_alterInstance := any(concreteInstance).(*models.Harmony_alter)
		ret2 := backRepo.BackRepoHarmony_alter.GetHarmony_alterDBFromHarmony_alterPtr(harmony_alterInstance)
		ret = any(ret2).(*T2)
	case *models.Harp_pedals:
		harp_pedalsInstance := any(concreteInstance).(*models.Harp_pedals)
		ret2 := backRepo.BackRepoHarp_pedals.GetHarp_pedalsDBFromHarp_pedalsPtr(harp_pedalsInstance)
		ret = any(ret2).(*T2)
	case *models.Heel_toe:
		heel_toeInstance := any(concreteInstance).(*models.Heel_toe)
		ret2 := backRepo.BackRepoHeel_toe.GetHeel_toeDBFromHeel_toePtr(heel_toeInstance)
		ret = any(ret2).(*T2)
	case *models.Hole:
		holeInstance := any(concreteInstance).(*models.Hole)
		ret2 := backRepo.BackRepoHole.GetHoleDBFromHolePtr(holeInstance)
		ret = any(ret2).(*T2)
	case *models.Hole_closed:
		hole_closedInstance := any(concreteInstance).(*models.Hole_closed)
		ret2 := backRepo.BackRepoHole_closed.GetHole_closedDBFromHole_closedPtr(hole_closedInstance)
		ret = any(ret2).(*T2)
	case *models.Horizontal_turn:
		horizontal_turnInstance := any(concreteInstance).(*models.Horizontal_turn)
		ret2 := backRepo.BackRepoHorizontal_turn.GetHorizontal_turnDBFromHorizontal_turnPtr(horizontal_turnInstance)
		ret = any(ret2).(*T2)
	case *models.Identification:
		identificationInstance := any(concreteInstance).(*models.Identification)
		ret2 := backRepo.BackRepoIdentification.GetIdentificationDBFromIdentificationPtr(identificationInstance)
		ret = any(ret2).(*T2)
	case *models.Image:
		imageInstance := any(concreteInstance).(*models.Image)
		ret2 := backRepo.BackRepoImage.GetImageDBFromImagePtr(imageInstance)
		ret = any(ret2).(*T2)
	case *models.Instrument:
		instrumentInstance := any(concreteInstance).(*models.Instrument)
		ret2 := backRepo.BackRepoInstrument.GetInstrumentDBFromInstrumentPtr(instrumentInstance)
		ret = any(ret2).(*T2)
	case *models.Instrument_change:
		instrument_changeInstance := any(concreteInstance).(*models.Instrument_change)
		ret2 := backRepo.BackRepoInstrument_change.GetInstrument_changeDBFromInstrument_changePtr(instrument_changeInstance)
		ret = any(ret2).(*T2)
	case *models.Instrument_link:
		instrument_linkInstance := any(concreteInstance).(*models.Instrument_link)
		ret2 := backRepo.BackRepoInstrument_link.GetInstrument_linkDBFromInstrument_linkPtr(instrument_linkInstance)
		ret = any(ret2).(*T2)
	case *models.Interchangeable:
		interchangeableInstance := any(concreteInstance).(*models.Interchangeable)
		ret2 := backRepo.BackRepoInterchangeable.GetInterchangeableDBFromInterchangeablePtr(interchangeableInstance)
		ret = any(ret2).(*T2)
	case *models.Inversion:
		inversionInstance := any(concreteInstance).(*models.Inversion)
		ret2 := backRepo.BackRepoInversion.GetInversionDBFromInversionPtr(inversionInstance)
		ret = any(ret2).(*T2)
	case *models.Key:
		keyInstance := any(concreteInstance).(*models.Key)
		ret2 := backRepo.BackRepoKey.GetKeyDBFromKeyPtr(keyInstance)
		ret = any(ret2).(*T2)
	case *models.Key_accidental:
		key_accidentalInstance := any(concreteInstance).(*models.Key_accidental)
		ret2 := backRepo.BackRepoKey_accidental.GetKey_accidentalDBFromKey_accidentalPtr(key_accidentalInstance)
		ret = any(ret2).(*T2)
	case *models.Key_octave:
		key_octaveInstance := any(concreteInstance).(*models.Key_octave)
		ret2 := backRepo.BackRepoKey_octave.GetKey_octaveDBFromKey_octavePtr(key_octaveInstance)
		ret = any(ret2).(*T2)
	case *models.Kind:
		kindInstance := any(concreteInstance).(*models.Kind)
		ret2 := backRepo.BackRepoKind.GetKindDBFromKindPtr(kindInstance)
		ret = any(ret2).(*T2)
	case *models.Level:
		levelInstance := any(concreteInstance).(*models.Level)
		ret2 := backRepo.BackRepoLevel.GetLevelDBFromLevelPtr(levelInstance)
		ret = any(ret2).(*T2)
	case *models.Line_detail:
		line_detailInstance := any(concreteInstance).(*models.Line_detail)
		ret2 := backRepo.BackRepoLine_detail.GetLine_detailDBFromLine_detailPtr(line_detailInstance)
		ret = any(ret2).(*T2)
	case *models.Line_width:
		line_widthInstance := any(concreteInstance).(*models.Line_width)
		ret2 := backRepo.BackRepoLine_width.GetLine_widthDBFromLine_widthPtr(line_widthInstance)
		ret = any(ret2).(*T2)
	case *models.Link:
		linkInstance := any(concreteInstance).(*models.Link)
		ret2 := backRepo.BackRepoLink.GetLinkDBFromLinkPtr(linkInstance)
		ret = any(ret2).(*T2)
	case *models.Listen:
		listenInstance := any(concreteInstance).(*models.Listen)
		ret2 := backRepo.BackRepoListen.GetListenDBFromListenPtr(listenInstance)
		ret = any(ret2).(*T2)
	case *models.Listening:
		listeningInstance := any(concreteInstance).(*models.Listening)
		ret2 := backRepo.BackRepoListening.GetListeningDBFromListeningPtr(listeningInstance)
		ret = any(ret2).(*T2)
	case *models.Lyric:
		lyricInstance := any(concreteInstance).(*models.Lyric)
		ret2 := backRepo.BackRepoLyric.GetLyricDBFromLyricPtr(lyricInstance)
		ret = any(ret2).(*T2)
	case *models.Lyric_font:
		lyric_fontInstance := any(concreteInstance).(*models.Lyric_font)
		ret2 := backRepo.BackRepoLyric_font.GetLyric_fontDBFromLyric_fontPtr(lyric_fontInstance)
		ret = any(ret2).(*T2)
	case *models.Lyric_language:
		lyric_languageInstance := any(concreteInstance).(*models.Lyric_language)
		ret2 := backRepo.BackRepoLyric_language.GetLyric_languageDBFromLyric_languagePtr(lyric_languageInstance)
		ret = any(ret2).(*T2)
	case *models.Measure_layout:
		measure_layoutInstance := any(concreteInstance).(*models.Measure_layout)
		ret2 := backRepo.BackRepoMeasure_layout.GetMeasure_layoutDBFromMeasure_layoutPtr(measure_layoutInstance)
		ret = any(ret2).(*T2)
	case *models.Measure_numbering:
		measure_numberingInstance := any(concreteInstance).(*models.Measure_numbering)
		ret2 := backRepo.BackRepoMeasure_numbering.GetMeasure_numberingDBFromMeasure_numberingPtr(measure_numberingInstance)
		ret = any(ret2).(*T2)
	case *models.Measure_repeat:
		measure_repeatInstance := any(concreteInstance).(*models.Measure_repeat)
		ret2 := backRepo.BackRepoMeasure_repeat.GetMeasure_repeatDBFromMeasure_repeatPtr(measure_repeatInstance)
		ret = any(ret2).(*T2)
	case *models.Measure_style:
		measure_styleInstance := any(concreteInstance).(*models.Measure_style)
		ret2 := backRepo.BackRepoMeasure_style.GetMeasure_styleDBFromMeasure_stylePtr(measure_styleInstance)
		ret = any(ret2).(*T2)
	case *models.Membrane:
		membraneInstance := any(concreteInstance).(*models.Membrane)
		ret2 := backRepo.BackRepoMembrane.GetMembraneDBFromMembranePtr(membraneInstance)
		ret = any(ret2).(*T2)
	case *models.Metal:
		metalInstance := any(concreteInstance).(*models.Metal)
		ret2 := backRepo.BackRepoMetal.GetMetalDBFromMetalPtr(metalInstance)
		ret = any(ret2).(*T2)
	case *models.Metronome:
		metronomeInstance := any(concreteInstance).(*models.Metronome)
		ret2 := backRepo.BackRepoMetronome.GetMetronomeDBFromMetronomePtr(metronomeInstance)
		ret = any(ret2).(*T2)
	case *models.Metronome_beam:
		metronome_beamInstance := any(concreteInstance).(*models.Metronome_beam)
		ret2 := backRepo.BackRepoMetronome_beam.GetMetronome_beamDBFromMetronome_beamPtr(metronome_beamInstance)
		ret = any(ret2).(*T2)
	case *models.Metronome_note:
		metronome_noteInstance := any(concreteInstance).(*models.Metronome_note)
		ret2 := backRepo.BackRepoMetronome_note.GetMetronome_noteDBFromMetronome_notePtr(metronome_noteInstance)
		ret = any(ret2).(*T2)
	case *models.Metronome_tied:
		metronome_tiedInstance := any(concreteInstance).(*models.Metronome_tied)
		ret2 := backRepo.BackRepoMetronome_tied.GetMetronome_tiedDBFromMetronome_tiedPtr(metronome_tiedInstance)
		ret = any(ret2).(*T2)
	case *models.Metronome_tuplet:
		metronome_tupletInstance := any(concreteInstance).(*models.Metronome_tuplet)
		ret2 := backRepo.BackRepoMetronome_tuplet.GetMetronome_tupletDBFromMetronome_tupletPtr(metronome_tupletInstance)
		ret = any(ret2).(*T2)
	case *models.Midi_device:
		midi_deviceInstance := any(concreteInstance).(*models.Midi_device)
		ret2 := backRepo.BackRepoMidi_device.GetMidi_deviceDBFromMidi_devicePtr(midi_deviceInstance)
		ret = any(ret2).(*T2)
	case *models.Midi_instrument:
		midi_instrumentInstance := any(concreteInstance).(*models.Midi_instrument)
		ret2 := backRepo.BackRepoMidi_instrument.GetMidi_instrumentDBFromMidi_instrumentPtr(midi_instrumentInstance)
		ret = any(ret2).(*T2)
	case *models.Miscellaneous:
		miscellaneousInstance := any(concreteInstance).(*models.Miscellaneous)
		ret2 := backRepo.BackRepoMiscellaneous.GetMiscellaneousDBFromMiscellaneousPtr(miscellaneousInstance)
		ret = any(ret2).(*T2)
	case *models.Miscellaneous_field:
		miscellaneous_fieldInstance := any(concreteInstance).(*models.Miscellaneous_field)
		ret2 := backRepo.BackRepoMiscellaneous_field.GetMiscellaneous_fieldDBFromMiscellaneous_fieldPtr(miscellaneous_fieldInstance)
		ret = any(ret2).(*T2)
	case *models.Mordent:
		mordentInstance := any(concreteInstance).(*models.Mordent)
		ret2 := backRepo.BackRepoMordent.GetMordentDBFromMordentPtr(mordentInstance)
		ret = any(ret2).(*T2)
	case *models.Multiple_rest:
		multiple_restInstance := any(concreteInstance).(*models.Multiple_rest)
		ret2 := backRepo.BackRepoMultiple_rest.GetMultiple_restDBFromMultiple_restPtr(multiple_restInstance)
		ret = any(ret2).(*T2)
	case *models.Name_display:
		name_displayInstance := any(concreteInstance).(*models.Name_display)
		ret2 := backRepo.BackRepoName_display.GetName_displayDBFromName_displayPtr(name_displayInstance)
		ret = any(ret2).(*T2)
	case *models.Non_arpeggiate:
		non_arpeggiateInstance := any(concreteInstance).(*models.Non_arpeggiate)
		ret2 := backRepo.BackRepoNon_arpeggiate.GetNon_arpeggiateDBFromNon_arpeggiatePtr(non_arpeggiateInstance)
		ret = any(ret2).(*T2)
	case *models.Notations:
		notationsInstance := any(concreteInstance).(*models.Notations)
		ret2 := backRepo.BackRepoNotations.GetNotationsDBFromNotationsPtr(notationsInstance)
		ret = any(ret2).(*T2)
	case *models.Note:
		noteInstance := any(concreteInstance).(*models.Note)
		ret2 := backRepo.BackRepoNote.GetNoteDBFromNotePtr(noteInstance)
		ret = any(ret2).(*T2)
	case *models.Note_size:
		note_sizeInstance := any(concreteInstance).(*models.Note_size)
		ret2 := backRepo.BackRepoNote_size.GetNote_sizeDBFromNote_sizePtr(note_sizeInstance)
		ret = any(ret2).(*T2)
	case *models.Note_type:
		note_typeInstance := any(concreteInstance).(*models.Note_type)
		ret2 := backRepo.BackRepoNote_type.GetNote_typeDBFromNote_typePtr(note_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Notehead:
		noteheadInstance := any(concreteInstance).(*models.Notehead)
		ret2 := backRepo.BackRepoNotehead.GetNoteheadDBFromNoteheadPtr(noteheadInstance)
		ret = any(ret2).(*T2)
	case *models.Notehead_text:
		notehead_textInstance := any(concreteInstance).(*models.Notehead_text)
		ret2 := backRepo.BackRepoNotehead_text.GetNotehead_textDBFromNotehead_textPtr(notehead_textInstance)
		ret = any(ret2).(*T2)
	case *models.Numeral:
		numeralInstance := any(concreteInstance).(*models.Numeral)
		ret2 := backRepo.BackRepoNumeral.GetNumeralDBFromNumeralPtr(numeralInstance)
		ret = any(ret2).(*T2)
	case *models.Numeral_key:
		numeral_keyInstance := any(concreteInstance).(*models.Numeral_key)
		ret2 := backRepo.BackRepoNumeral_key.GetNumeral_keyDBFromNumeral_keyPtr(numeral_keyInstance)
		ret = any(ret2).(*T2)
	case *models.Numeral_root:
		numeral_rootInstance := any(concreteInstance).(*models.Numeral_root)
		ret2 := backRepo.BackRepoNumeral_root.GetNumeral_rootDBFromNumeral_rootPtr(numeral_rootInstance)
		ret = any(ret2).(*T2)
	case *models.Octave_shift:
		octave_shiftInstance := any(concreteInstance).(*models.Octave_shift)
		ret2 := backRepo.BackRepoOctave_shift.GetOctave_shiftDBFromOctave_shiftPtr(octave_shiftInstance)
		ret = any(ret2).(*T2)
	case *models.Offset:
		offsetInstance := any(concreteInstance).(*models.Offset)
		ret2 := backRepo.BackRepoOffset.GetOffsetDBFromOffsetPtr(offsetInstance)
		ret = any(ret2).(*T2)
	case *models.Opus:
		opusInstance := any(concreteInstance).(*models.Opus)
		ret2 := backRepo.BackRepoOpus.GetOpusDBFromOpusPtr(opusInstance)
		ret = any(ret2).(*T2)
	case *models.Ornaments:
		ornamentsInstance := any(concreteInstance).(*models.Ornaments)
		ret2 := backRepo.BackRepoOrnaments.GetOrnamentsDBFromOrnamentsPtr(ornamentsInstance)
		ret = any(ret2).(*T2)
	case *models.Other_appearance:
		other_appearanceInstance := any(concreteInstance).(*models.Other_appearance)
		ret2 := backRepo.BackRepoOther_appearance.GetOther_appearanceDBFromOther_appearancePtr(other_appearanceInstance)
		ret = any(ret2).(*T2)
	case *models.Other_listening:
		other_listeningInstance := any(concreteInstance).(*models.Other_listening)
		ret2 := backRepo.BackRepoOther_listening.GetOther_listeningDBFromOther_listeningPtr(other_listeningInstance)
		ret = any(ret2).(*T2)
	case *models.Other_notation:
		other_notationInstance := any(concreteInstance).(*models.Other_notation)
		ret2 := backRepo.BackRepoOther_notation.GetOther_notationDBFromOther_notationPtr(other_notationInstance)
		ret = any(ret2).(*T2)
	case *models.Other_play:
		other_playInstance := any(concreteInstance).(*models.Other_play)
		ret2 := backRepo.BackRepoOther_play.GetOther_playDBFromOther_playPtr(other_playInstance)
		ret = any(ret2).(*T2)
	case *models.Page_layout:
		page_layoutInstance := any(concreteInstance).(*models.Page_layout)
		ret2 := backRepo.BackRepoPage_layout.GetPage_layoutDBFromPage_layoutPtr(page_layoutInstance)
		ret = any(ret2).(*T2)
	case *models.Page_margins:
		page_marginsInstance := any(concreteInstance).(*models.Page_margins)
		ret2 := backRepo.BackRepoPage_margins.GetPage_marginsDBFromPage_marginsPtr(page_marginsInstance)
		ret = any(ret2).(*T2)
	case *models.Part_clef:
		part_clefInstance := any(concreteInstance).(*models.Part_clef)
		ret2 := backRepo.BackRepoPart_clef.GetPart_clefDBFromPart_clefPtr(part_clefInstance)
		ret = any(ret2).(*T2)
	case *models.Part_group:
		part_groupInstance := any(concreteInstance).(*models.Part_group)
		ret2 := backRepo.BackRepoPart_group.GetPart_groupDBFromPart_groupPtr(part_groupInstance)
		ret = any(ret2).(*T2)
	case *models.Part_link:
		part_linkInstance := any(concreteInstance).(*models.Part_link)
		ret2 := backRepo.BackRepoPart_link.GetPart_linkDBFromPart_linkPtr(part_linkInstance)
		ret = any(ret2).(*T2)
	case *models.Part_list:
		part_listInstance := any(concreteInstance).(*models.Part_list)
		ret2 := backRepo.BackRepoPart_list.GetPart_listDBFromPart_listPtr(part_listInstance)
		ret = any(ret2).(*T2)
	case *models.Part_symbol:
		part_symbolInstance := any(concreteInstance).(*models.Part_symbol)
		ret2 := backRepo.BackRepoPart_symbol.GetPart_symbolDBFromPart_symbolPtr(part_symbolInstance)
		ret = any(ret2).(*T2)
	case *models.Part_transpose:
		part_transposeInstance := any(concreteInstance).(*models.Part_transpose)
		ret2 := backRepo.BackRepoPart_transpose.GetPart_transposeDBFromPart_transposePtr(part_transposeInstance)
		ret = any(ret2).(*T2)
	case *models.Pedal:
		pedalInstance := any(concreteInstance).(*models.Pedal)
		ret2 := backRepo.BackRepoPedal.GetPedalDBFromPedalPtr(pedalInstance)
		ret = any(ret2).(*T2)
	case *models.Pedal_tuning:
		pedal_tuningInstance := any(concreteInstance).(*models.Pedal_tuning)
		ret2 := backRepo.BackRepoPedal_tuning.GetPedal_tuningDBFromPedal_tuningPtr(pedal_tuningInstance)
		ret = any(ret2).(*T2)
	case *models.Percussion:
		percussionInstance := any(concreteInstance).(*models.Percussion)
		ret2 := backRepo.BackRepoPercussion.GetPercussionDBFromPercussionPtr(percussionInstance)
		ret = any(ret2).(*T2)
	case *models.Pitch:
		pitchInstance := any(concreteInstance).(*models.Pitch)
		ret2 := backRepo.BackRepoPitch.GetPitchDBFromPitchPtr(pitchInstance)
		ret = any(ret2).(*T2)
	case *models.Pitched:
		pitchedInstance := any(concreteInstance).(*models.Pitched)
		ret2 := backRepo.BackRepoPitched.GetPitchedDBFromPitchedPtr(pitchedInstance)
		ret = any(ret2).(*T2)
	case *models.Play:
		playInstance := any(concreteInstance).(*models.Play)
		ret2 := backRepo.BackRepoPlay.GetPlayDBFromPlayPtr(playInstance)
		ret = any(ret2).(*T2)
	case *models.Player:
		playerInstance := any(concreteInstance).(*models.Player)
		ret2 := backRepo.BackRepoPlayer.GetPlayerDBFromPlayerPtr(playerInstance)
		ret = any(ret2).(*T2)
	case *models.Principal_voice:
		principal_voiceInstance := any(concreteInstance).(*models.Principal_voice)
		ret2 := backRepo.BackRepoPrincipal_voice.GetPrincipal_voiceDBFromPrincipal_voicePtr(principal_voiceInstance)
		ret = any(ret2).(*T2)
	case *models.Print:
		printInstance := any(concreteInstance).(*models.Print)
		ret2 := backRepo.BackRepoPrint.GetPrintDBFromPrintPtr(printInstance)
		ret = any(ret2).(*T2)
	case *models.Release:
		releaseInstance := any(concreteInstance).(*models.Release)
		ret2 := backRepo.BackRepoRelease.GetReleaseDBFromReleasePtr(releaseInstance)
		ret = any(ret2).(*T2)
	case *models.Repeat:
		repeatInstance := any(concreteInstance).(*models.Repeat)
		ret2 := backRepo.BackRepoRepeat.GetRepeatDBFromRepeatPtr(repeatInstance)
		ret = any(ret2).(*T2)
	case *models.Rest:
		restInstance := any(concreteInstance).(*models.Rest)
		ret2 := backRepo.BackRepoRest.GetRestDBFromRestPtr(restInstance)
		ret = any(ret2).(*T2)
	case *models.Root:
		rootInstance := any(concreteInstance).(*models.Root)
		ret2 := backRepo.BackRepoRoot.GetRootDBFromRootPtr(rootInstance)
		ret = any(ret2).(*T2)
	case *models.Root_step:
		root_stepInstance := any(concreteInstance).(*models.Root_step)
		ret2 := backRepo.BackRepoRoot_step.GetRoot_stepDBFromRoot_stepPtr(root_stepInstance)
		ret = any(ret2).(*T2)
	case *models.Scaling:
		scalingInstance := any(concreteInstance).(*models.Scaling)
		ret2 := backRepo.BackRepoScaling.GetScalingDBFromScalingPtr(scalingInstance)
		ret = any(ret2).(*T2)
	case *models.Scordatura:
		scordaturaInstance := any(concreteInstance).(*models.Scordatura)
		ret2 := backRepo.BackRepoScordatura.GetScordaturaDBFromScordaturaPtr(scordaturaInstance)
		ret = any(ret2).(*T2)
	case *models.Score_instrument:
		score_instrumentInstance := any(concreteInstance).(*models.Score_instrument)
		ret2 := backRepo.BackRepoScore_instrument.GetScore_instrumentDBFromScore_instrumentPtr(score_instrumentInstance)
		ret = any(ret2).(*T2)
	case *models.Score_part:
		score_partInstance := any(concreteInstance).(*models.Score_part)
		ret2 := backRepo.BackRepoScore_part.GetScore_partDBFromScore_partPtr(score_partInstance)
		ret = any(ret2).(*T2)
	case *models.Score_partwise:
		score_partwiseInstance := any(concreteInstance).(*models.Score_partwise)
		ret2 := backRepo.BackRepoScore_partwise.GetScore_partwiseDBFromScore_partwisePtr(score_partwiseInstance)
		ret = any(ret2).(*T2)
	case *models.Score_timewise:
		score_timewiseInstance := any(concreteInstance).(*models.Score_timewise)
		ret2 := backRepo.BackRepoScore_timewise.GetScore_timewiseDBFromScore_timewisePtr(score_timewiseInstance)
		ret = any(ret2).(*T2)
	case *models.Segno:
		segnoInstance := any(concreteInstance).(*models.Segno)
		ret2 := backRepo.BackRepoSegno.GetSegnoDBFromSegnoPtr(segnoInstance)
		ret = any(ret2).(*T2)
	case *models.Slash:
		slashInstance := any(concreteInstance).(*models.Slash)
		ret2 := backRepo.BackRepoSlash.GetSlashDBFromSlashPtr(slashInstance)
		ret = any(ret2).(*T2)
	case *models.Slide:
		slideInstance := any(concreteInstance).(*models.Slide)
		ret2 := backRepo.BackRepoSlide.GetSlideDBFromSlidePtr(slideInstance)
		ret = any(ret2).(*T2)
	case *models.Slur:
		slurInstance := any(concreteInstance).(*models.Slur)
		ret2 := backRepo.BackRepoSlur.GetSlurDBFromSlurPtr(slurInstance)
		ret = any(ret2).(*T2)
	case *models.Sound:
		soundInstance := any(concreteInstance).(*models.Sound)
		ret2 := backRepo.BackRepoSound.GetSoundDBFromSoundPtr(soundInstance)
		ret = any(ret2).(*T2)
	case *models.Staff_details:
		staff_detailsInstance := any(concreteInstance).(*models.Staff_details)
		ret2 := backRepo.BackRepoStaff_details.GetStaff_detailsDBFromStaff_detailsPtr(staff_detailsInstance)
		ret = any(ret2).(*T2)
	case *models.Staff_divide:
		staff_divideInstance := any(concreteInstance).(*models.Staff_divide)
		ret2 := backRepo.BackRepoStaff_divide.GetStaff_divideDBFromStaff_dividePtr(staff_divideInstance)
		ret = any(ret2).(*T2)
	case *models.Staff_layout:
		staff_layoutInstance := any(concreteInstance).(*models.Staff_layout)
		ret2 := backRepo.BackRepoStaff_layout.GetStaff_layoutDBFromStaff_layoutPtr(staff_layoutInstance)
		ret = any(ret2).(*T2)
	case *models.Staff_size:
		staff_sizeInstance := any(concreteInstance).(*models.Staff_size)
		ret2 := backRepo.BackRepoStaff_size.GetStaff_sizeDBFromStaff_sizePtr(staff_sizeInstance)
		ret = any(ret2).(*T2)
	case *models.Staff_tuning:
		staff_tuningInstance := any(concreteInstance).(*models.Staff_tuning)
		ret2 := backRepo.BackRepoStaff_tuning.GetStaff_tuningDBFromStaff_tuningPtr(staff_tuningInstance)
		ret = any(ret2).(*T2)
	case *models.Stem:
		stemInstance := any(concreteInstance).(*models.Stem)
		ret2 := backRepo.BackRepoStem.GetStemDBFromStemPtr(stemInstance)
		ret = any(ret2).(*T2)
	case *models.Stick:
		stickInstance := any(concreteInstance).(*models.Stick)
		ret2 := backRepo.BackRepoStick.GetStickDBFromStickPtr(stickInstance)
		ret = any(ret2).(*T2)
	case *models.String_mute:
		string_muteInstance := any(concreteInstance).(*models.String_mute)
		ret2 := backRepo.BackRepoString_mute.GetString_muteDBFromString_mutePtr(string_muteInstance)
		ret = any(ret2).(*T2)
	case *models.Strong_accent:
		strong_accentInstance := any(concreteInstance).(*models.Strong_accent)
		ret2 := backRepo.BackRepoStrong_accent.GetStrong_accentDBFromStrong_accentPtr(strong_accentInstance)
		ret = any(ret2).(*T2)
	case *models.Supports:
		supportsInstance := any(concreteInstance).(*models.Supports)
		ret2 := backRepo.BackRepoSupports.GetSupportsDBFromSupportsPtr(supportsInstance)
		ret = any(ret2).(*T2)
	case *models.Swing:
		swingInstance := any(concreteInstance).(*models.Swing)
		ret2 := backRepo.BackRepoSwing.GetSwingDBFromSwingPtr(swingInstance)
		ret = any(ret2).(*T2)
	case *models.Sync:
		syncInstance := any(concreteInstance).(*models.Sync)
		ret2 := backRepo.BackRepoSync.GetSyncDBFromSyncPtr(syncInstance)
		ret = any(ret2).(*T2)
	case *models.System_dividers:
		system_dividersInstance := any(concreteInstance).(*models.System_dividers)
		ret2 := backRepo.BackRepoSystem_dividers.GetSystem_dividersDBFromSystem_dividersPtr(system_dividersInstance)
		ret = any(ret2).(*T2)
	case *models.System_layout:
		system_layoutInstance := any(concreteInstance).(*models.System_layout)
		ret2 := backRepo.BackRepoSystem_layout.GetSystem_layoutDBFromSystem_layoutPtr(system_layoutInstance)
		ret = any(ret2).(*T2)
	case *models.System_margins:
		system_marginsInstance := any(concreteInstance).(*models.System_margins)
		ret2 := backRepo.BackRepoSystem_margins.GetSystem_marginsDBFromSystem_marginsPtr(system_marginsInstance)
		ret = any(ret2).(*T2)
	case *models.Tap:
		tapInstance := any(concreteInstance).(*models.Tap)
		ret2 := backRepo.BackRepoTap.GetTapDBFromTapPtr(tapInstance)
		ret = any(ret2).(*T2)
	case *models.Technical:
		technicalInstance := any(concreteInstance).(*models.Technical)
		ret2 := backRepo.BackRepoTechnical.GetTechnicalDBFromTechnicalPtr(technicalInstance)
		ret = any(ret2).(*T2)
	case *models.Text_element_data:
		text_element_dataInstance := any(concreteInstance).(*models.Text_element_data)
		ret2 := backRepo.BackRepoText_element_data.GetText_element_dataDBFromText_element_dataPtr(text_element_dataInstance)
		ret = any(ret2).(*T2)
	case *models.Tie:
		tieInstance := any(concreteInstance).(*models.Tie)
		ret2 := backRepo.BackRepoTie.GetTieDBFromTiePtr(tieInstance)
		ret = any(ret2).(*T2)
	case *models.Tied:
		tiedInstance := any(concreteInstance).(*models.Tied)
		ret2 := backRepo.BackRepoTied.GetTiedDBFromTiedPtr(tiedInstance)
		ret = any(ret2).(*T2)
	case *models.Time:
		timeInstance := any(concreteInstance).(*models.Time)
		ret2 := backRepo.BackRepoTime.GetTimeDBFromTimePtr(timeInstance)
		ret = any(ret2).(*T2)
	case *models.Time_modification:
		time_modificationInstance := any(concreteInstance).(*models.Time_modification)
		ret2 := backRepo.BackRepoTime_modification.GetTime_modificationDBFromTime_modificationPtr(time_modificationInstance)
		ret = any(ret2).(*T2)
	case *models.Timpani:
		timpaniInstance := any(concreteInstance).(*models.Timpani)
		ret2 := backRepo.BackRepoTimpani.GetTimpaniDBFromTimpaniPtr(timpaniInstance)
		ret = any(ret2).(*T2)
	case *models.Transpose:
		transposeInstance := any(concreteInstance).(*models.Transpose)
		ret2 := backRepo.BackRepoTranspose.GetTransposeDBFromTransposePtr(transposeInstance)
		ret = any(ret2).(*T2)
	case *models.Tremolo:
		tremoloInstance := any(concreteInstance).(*models.Tremolo)
		ret2 := backRepo.BackRepoTremolo.GetTremoloDBFromTremoloPtr(tremoloInstance)
		ret = any(ret2).(*T2)
	case *models.Tuplet:
		tupletInstance := any(concreteInstance).(*models.Tuplet)
		ret2 := backRepo.BackRepoTuplet.GetTupletDBFromTupletPtr(tupletInstance)
		ret = any(ret2).(*T2)
	case *models.Tuplet_dot:
		tuplet_dotInstance := any(concreteInstance).(*models.Tuplet_dot)
		ret2 := backRepo.BackRepoTuplet_dot.GetTuplet_dotDBFromTuplet_dotPtr(tuplet_dotInstance)
		ret = any(ret2).(*T2)
	case *models.Tuplet_number:
		tuplet_numberInstance := any(concreteInstance).(*models.Tuplet_number)
		ret2 := backRepo.BackRepoTuplet_number.GetTuplet_numberDBFromTuplet_numberPtr(tuplet_numberInstance)
		ret = any(ret2).(*T2)
	case *models.Tuplet_portion:
		tuplet_portionInstance := any(concreteInstance).(*models.Tuplet_portion)
		ret2 := backRepo.BackRepoTuplet_portion.GetTuplet_portionDBFromTuplet_portionPtr(tuplet_portionInstance)
		ret = any(ret2).(*T2)
	case *models.Tuplet_type:
		tuplet_typeInstance := any(concreteInstance).(*models.Tuplet_type)
		ret2 := backRepo.BackRepoTuplet_type.GetTuplet_typeDBFromTuplet_typePtr(tuplet_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Typed_text:
		typed_textInstance := any(concreteInstance).(*models.Typed_text)
		ret2 := backRepo.BackRepoTyped_text.GetTyped_textDBFromTyped_textPtr(typed_textInstance)
		ret = any(ret2).(*T2)
	case *models.Unpitched:
		unpitchedInstance := any(concreteInstance).(*models.Unpitched)
		ret2 := backRepo.BackRepoUnpitched.GetUnpitchedDBFromUnpitchedPtr(unpitchedInstance)
		ret = any(ret2).(*T2)
	case *models.Virtual_instrument:
		virtual_instrumentInstance := any(concreteInstance).(*models.Virtual_instrument)
		ret2 := backRepo.BackRepoVirtual_instrument.GetVirtual_instrumentDBFromVirtual_instrumentPtr(virtual_instrumentInstance)
		ret = any(ret2).(*T2)
	case *models.Wait:
		waitInstance := any(concreteInstance).(*models.Wait)
		ret2 := backRepo.BackRepoWait.GetWaitDBFromWaitPtr(waitInstance)
		ret = any(ret2).(*T2)
	case *models.Wavy_line:
		wavy_lineInstance := any(concreteInstance).(*models.Wavy_line)
		ret2 := backRepo.BackRepoWavy_line.GetWavy_lineDBFromWavy_linePtr(wavy_lineInstance)
		ret = any(ret2).(*T2)
	case *models.Wedge:
		wedgeInstance := any(concreteInstance).(*models.Wedge)
		ret2 := backRepo.BackRepoWedge.GetWedgeDBFromWedgePtr(wedgeInstance)
		ret = any(ret2).(*T2)
	case *models.Wood:
		woodInstance := any(concreteInstance).(*models.Wood)
		ret2 := backRepo.BackRepoWood.GetWoodDBFromWoodPtr(woodInstance)
		ret = any(ret2).(*T2)
	case *models.Work:
		workInstance := any(concreteInstance).(*models.Work)
		ret2 := backRepo.BackRepoWork.GetWorkDBFromWorkPtr(workInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Accidental:
		tmp := GetInstanceDBFromInstance[models.Accidental, AccidentalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accidental_mark:
		tmp := GetInstanceDBFromInstance[models.Accidental_mark, Accidental_markDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accidental_text:
		tmp := GetInstanceDBFromInstance[models.Accidental_text, Accidental_textDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accord:
		tmp := GetInstanceDBFromInstance[models.Accord, AccordDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accordion_registration:
		tmp := GetInstanceDBFromInstance[models.Accordion_registration, Accordion_registrationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AnyType:
		tmp := GetInstanceDBFromInstance[models.AnyType, AnyTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Appearance:
		tmp := GetInstanceDBFromInstance[models.Appearance, AppearanceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Arpeggiate:
		tmp := GetInstanceDBFromInstance[models.Arpeggiate, ArpeggiateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Articulations:
		tmp := GetInstanceDBFromInstance[models.Articulations, ArticulationsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Assess:
		tmp := GetInstanceDBFromInstance[models.Assess, AssessDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Attributes:
		tmp := GetInstanceDBFromInstance[models.Attributes, AttributesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Backup:
		tmp := GetInstanceDBFromInstance[models.Backup, BackupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bar_style_color:
		tmp := GetInstanceDBFromInstance[models.Bar_style_color, Bar_style_colorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Barline:
		tmp := GetInstanceDBFromInstance[models.Barline, BarlineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Barre:
		tmp := GetInstanceDBFromInstance[models.Barre, BarreDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bass:
		tmp := GetInstanceDBFromInstance[models.Bass, BassDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bass_step:
		tmp := GetInstanceDBFromInstance[models.Bass_step, Bass_stepDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beam:
		tmp := GetInstanceDBFromInstance[models.Beam, BeamDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beat_repeat:
		tmp := GetInstanceDBFromInstance[models.Beat_repeat, Beat_repeatDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beat_unit_tied:
		tmp := GetInstanceDBFromInstance[models.Beat_unit_tied, Beat_unit_tiedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beater:
		tmp := GetInstanceDBFromInstance[models.Beater, BeaterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bend:
		tmp := GetInstanceDBFromInstance[models.Bend, BendDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bookmark:
		tmp := GetInstanceDBFromInstance[models.Bookmark, BookmarkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bracket:
		tmp := GetInstanceDBFromInstance[models.Bracket, BracketDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Breath_mark:
		tmp := GetInstanceDBFromInstance[models.Breath_mark, Breath_markDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Caesura:
		tmp := GetInstanceDBFromInstance[models.Caesura, CaesuraDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Cancel:
		tmp := GetInstanceDBFromInstance[models.Cancel, CancelDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Clef:
		tmp := GetInstanceDBFromInstance[models.Clef, ClefDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Coda:
		tmp := GetInstanceDBFromInstance[models.Coda, CodaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Credit:
		tmp := GetInstanceDBFromInstance[models.Credit, CreditDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Dashes:
		tmp := GetInstanceDBFromInstance[models.Dashes, DashesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Defaults:
		tmp := GetInstanceDBFromInstance[models.Defaults, DefaultsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree:
		tmp := GetInstanceDBFromInstance[models.Degree, DegreeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree_alter:
		tmp := GetInstanceDBFromInstance[models.Degree_alter, Degree_alterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree_type:
		tmp := GetInstanceDBFromInstance[models.Degree_type, Degree_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree_value:
		tmp := GetInstanceDBFromInstance[models.Degree_value, Degree_valueDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Direction:
		tmp := GetInstanceDBFromInstance[models.Direction, DirectionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Direction_type:
		tmp := GetInstanceDBFromInstance[models.Direction_type, Direction_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Distance:
		tmp := GetInstanceDBFromInstance[models.Distance, DistanceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Double:
		tmp := GetInstanceDBFromInstance[models.Double, DoubleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Dynamics:
		tmp := GetInstanceDBFromInstance[models.Dynamics, DynamicsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Effect:
		tmp := GetInstanceDBFromInstance[models.Effect, EffectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Elision:
		tmp := GetInstanceDBFromInstance[models.Elision, ElisionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty:
		tmp := GetInstanceDBFromInstance[models.Empty, EmptyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_font:
		tmp := GetInstanceDBFromInstance[models.Empty_font, Empty_fontDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_line:
		tmp := GetInstanceDBFromInstance[models.Empty_line, Empty_lineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_placement:
		tmp := GetInstanceDBFromInstance[models.Empty_placement, Empty_placementDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_placement_smufl:
		tmp := GetInstanceDBFromInstance[models.Empty_placement_smufl, Empty_placement_smuflDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_object_style_align:
		tmp := GetInstanceDBFromInstance[models.Empty_print_object_style_align, Empty_print_object_style_alignDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_style:
		tmp := GetInstanceDBFromInstance[models.Empty_print_style, Empty_print_styleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_style_align:
		tmp := GetInstanceDBFromInstance[models.Empty_print_style_align, Empty_print_style_alignDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_style_align_id:
		tmp := GetInstanceDBFromInstance[models.Empty_print_style_align_id, Empty_print_style_align_idDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_trill_sound:
		tmp := GetInstanceDBFromInstance[models.Empty_trill_sound, Empty_trill_soundDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Encoding:
		tmp := GetInstanceDBFromInstance[models.Encoding, EncodingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Ending:
		tmp := GetInstanceDBFromInstance[models.Ending, EndingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Extend:
		tmp := GetInstanceDBFromInstance[models.Extend, ExtendDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Feature:
		tmp := GetInstanceDBFromInstance[models.Feature, FeatureDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Fermata:
		tmp := GetInstanceDBFromInstance[models.Fermata, FermataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Figure:
		tmp := GetInstanceDBFromInstance[models.Figure, FigureDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Figured_bass:
		tmp := GetInstanceDBFromInstance[models.Figured_bass, Figured_bassDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Fingering:
		tmp := GetInstanceDBFromInstance[models.Fingering, FingeringDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.First_fret:
		tmp := GetInstanceDBFromInstance[models.First_fret, First_fretDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Foo:
		tmp := GetInstanceDBFromInstance[models.Foo, FooDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.For_part:
		tmp := GetInstanceDBFromInstance[models.For_part, For_partDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Formatted_symbol:
		tmp := GetInstanceDBFromInstance[models.Formatted_symbol, Formatted_symbolDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Formatted_symbol_id:
		tmp := GetInstanceDBFromInstance[models.Formatted_symbol_id, Formatted_symbol_idDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Forward:
		tmp := GetInstanceDBFromInstance[models.Forward, ForwardDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Frame:
		tmp := GetInstanceDBFromInstance[models.Frame, FrameDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Frame_note:
		tmp := GetInstanceDBFromInstance[models.Frame_note, Frame_noteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Fret:
		tmp := GetInstanceDBFromInstance[models.Fret, FretDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Glass:
		tmp := GetInstanceDBFromInstance[models.Glass, GlassDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Glissando:
		tmp := GetInstanceDBFromInstance[models.Glissando, GlissandoDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Glyph:
		tmp := GetInstanceDBFromInstance[models.Glyph, GlyphDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Grace:
		tmp := GetInstanceDBFromInstance[models.Grace, GraceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group_barline:
		tmp := GetInstanceDBFromInstance[models.Group_barline, Group_barlineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group_symbol:
		tmp := GetInstanceDBFromInstance[models.Group_symbol, Group_symbolDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Grouping:
		tmp := GetInstanceDBFromInstance[models.Grouping, GroupingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hammer_on_pull_off:
		tmp := GetInstanceDBFromInstance[models.Hammer_on_pull_off, Hammer_on_pull_offDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Handbell:
		tmp := GetInstanceDBFromInstance[models.Handbell, HandbellDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmon_closed:
		tmp := GetInstanceDBFromInstance[models.Harmon_closed, Harmon_closedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmon_mute:
		tmp := GetInstanceDBFromInstance[models.Harmon_mute, Harmon_muteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmonic:
		tmp := GetInstanceDBFromInstance[models.Harmonic, HarmonicDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmony:
		tmp := GetInstanceDBFromInstance[models.Harmony, HarmonyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmony_alter:
		tmp := GetInstanceDBFromInstance[models.Harmony_alter, Harmony_alterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harp_pedals:
		tmp := GetInstanceDBFromInstance[models.Harp_pedals, Harp_pedalsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Heel_toe:
		tmp := GetInstanceDBFromInstance[models.Heel_toe, Heel_toeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hole:
		tmp := GetInstanceDBFromInstance[models.Hole, HoleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hole_closed:
		tmp := GetInstanceDBFromInstance[models.Hole_closed, Hole_closedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Horizontal_turn:
		tmp := GetInstanceDBFromInstance[models.Horizontal_turn, Horizontal_turnDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Identification:
		tmp := GetInstanceDBFromInstance[models.Identification, IdentificationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Image:
		tmp := GetInstanceDBFromInstance[models.Image, ImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Instrument:
		tmp := GetInstanceDBFromInstance[models.Instrument, InstrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Instrument_change:
		tmp := GetInstanceDBFromInstance[models.Instrument_change, Instrument_changeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Instrument_link:
		tmp := GetInstanceDBFromInstance[models.Instrument_link, Instrument_linkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Interchangeable:
		tmp := GetInstanceDBFromInstance[models.Interchangeable, InterchangeableDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Inversion:
		tmp := GetInstanceDBFromInstance[models.Inversion, InversionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key_accidental:
		tmp := GetInstanceDBFromInstance[models.Key_accidental, Key_accidentalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key_octave:
		tmp := GetInstanceDBFromInstance[models.Key_octave, Key_octaveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Kind:
		tmp := GetInstanceDBFromInstance[models.Kind, KindDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Level:
		tmp := GetInstanceDBFromInstance[models.Level, LevelDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line_detail:
		tmp := GetInstanceDBFromInstance[models.Line_detail, Line_detailDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line_width:
		tmp := GetInstanceDBFromInstance[models.Line_width, Line_widthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Listen:
		tmp := GetInstanceDBFromInstance[models.Listen, ListenDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Listening:
		tmp := GetInstanceDBFromInstance[models.Listening, ListeningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric:
		tmp := GetInstanceDBFromInstance[models.Lyric, LyricDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_font:
		tmp := GetInstanceDBFromInstance[models.Lyric_font, Lyric_fontDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_language:
		tmp := GetInstanceDBFromInstance[models.Lyric_language, Lyric_languageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_layout:
		tmp := GetInstanceDBFromInstance[models.Measure_layout, Measure_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_numbering:
		tmp := GetInstanceDBFromInstance[models.Measure_numbering, Measure_numberingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_repeat:
		tmp := GetInstanceDBFromInstance[models.Measure_repeat, Measure_repeatDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_style:
		tmp := GetInstanceDBFromInstance[models.Measure_style, Measure_styleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Membrane:
		tmp := GetInstanceDBFromInstance[models.Membrane, MembraneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metal:
		tmp := GetInstanceDBFromInstance[models.Metal, MetalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome:
		tmp := GetInstanceDBFromInstance[models.Metronome, MetronomeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_beam:
		tmp := GetInstanceDBFromInstance[models.Metronome_beam, Metronome_beamDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_note:
		tmp := GetInstanceDBFromInstance[models.Metronome_note, Metronome_noteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_tied:
		tmp := GetInstanceDBFromInstance[models.Metronome_tied, Metronome_tiedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_tuplet:
		tmp := GetInstanceDBFromInstance[models.Metronome_tuplet, Metronome_tupletDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Midi_device:
		tmp := GetInstanceDBFromInstance[models.Midi_device, Midi_deviceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Midi_instrument:
		tmp := GetInstanceDBFromInstance[models.Midi_instrument, Midi_instrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Miscellaneous:
		tmp := GetInstanceDBFromInstance[models.Miscellaneous, MiscellaneousDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Miscellaneous_field:
		tmp := GetInstanceDBFromInstance[models.Miscellaneous_field, Miscellaneous_fieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Mordent:
		tmp := GetInstanceDBFromInstance[models.Mordent, MordentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Multiple_rest:
		tmp := GetInstanceDBFromInstance[models.Multiple_rest, Multiple_restDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Name_display:
		tmp := GetInstanceDBFromInstance[models.Name_display, Name_displayDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Non_arpeggiate:
		tmp := GetInstanceDBFromInstance[models.Non_arpeggiate, Non_arpeggiateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Notations:
		tmp := GetInstanceDBFromInstance[models.Notations, NotationsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note:
		tmp := GetInstanceDBFromInstance[models.Note, NoteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note_size:
		tmp := GetInstanceDBFromInstance[models.Note_size, Note_sizeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note_type:
		tmp := GetInstanceDBFromInstance[models.Note_type, Note_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Notehead:
		tmp := GetInstanceDBFromInstance[models.Notehead, NoteheadDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Notehead_text:
		tmp := GetInstanceDBFromInstance[models.Notehead_text, Notehead_textDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Numeral:
		tmp := GetInstanceDBFromInstance[models.Numeral, NumeralDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Numeral_key:
		tmp := GetInstanceDBFromInstance[models.Numeral_key, Numeral_keyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Numeral_root:
		tmp := GetInstanceDBFromInstance[models.Numeral_root, Numeral_rootDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Octave_shift:
		tmp := GetInstanceDBFromInstance[models.Octave_shift, Octave_shiftDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Offset:
		tmp := GetInstanceDBFromInstance[models.Offset, OffsetDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Opus:
		tmp := GetInstanceDBFromInstance[models.Opus, OpusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Ornaments:
		tmp := GetInstanceDBFromInstance[models.Ornaments, OrnamentsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_appearance:
		tmp := GetInstanceDBFromInstance[models.Other_appearance, Other_appearanceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_listening:
		tmp := GetInstanceDBFromInstance[models.Other_listening, Other_listeningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_notation:
		tmp := GetInstanceDBFromInstance[models.Other_notation, Other_notationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_play:
		tmp := GetInstanceDBFromInstance[models.Other_play, Other_playDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Page_layout:
		tmp := GetInstanceDBFromInstance[models.Page_layout, Page_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Page_margins:
		tmp := GetInstanceDBFromInstance[models.Page_margins, Page_marginsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_clef:
		tmp := GetInstanceDBFromInstance[models.Part_clef, Part_clefDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_group:
		tmp := GetInstanceDBFromInstance[models.Part_group, Part_groupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_link:
		tmp := GetInstanceDBFromInstance[models.Part_link, Part_linkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_list:
		tmp := GetInstanceDBFromInstance[models.Part_list, Part_listDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_symbol:
		tmp := GetInstanceDBFromInstance[models.Part_symbol, Part_symbolDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_transpose:
		tmp := GetInstanceDBFromInstance[models.Part_transpose, Part_transposeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pedal:
		tmp := GetInstanceDBFromInstance[models.Pedal, PedalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pedal_tuning:
		tmp := GetInstanceDBFromInstance[models.Pedal_tuning, Pedal_tuningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Percussion:
		tmp := GetInstanceDBFromInstance[models.Percussion, PercussionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pitch:
		tmp := GetInstanceDBFromInstance[models.Pitch, PitchDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pitched:
		tmp := GetInstanceDBFromInstance[models.Pitched, PitchedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Play:
		tmp := GetInstanceDBFromInstance[models.Play, PlayDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Player:
		tmp := GetInstanceDBFromInstance[models.Player, PlayerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Principal_voice:
		tmp := GetInstanceDBFromInstance[models.Principal_voice, Principal_voiceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Print:
		tmp := GetInstanceDBFromInstance[models.Print, PrintDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Release:
		tmp := GetInstanceDBFromInstance[models.Release, ReleaseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Repeat:
		tmp := GetInstanceDBFromInstance[models.Repeat, RepeatDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rest:
		tmp := GetInstanceDBFromInstance[models.Rest, RestDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Root:
		tmp := GetInstanceDBFromInstance[models.Root, RootDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Root_step:
		tmp := GetInstanceDBFromInstance[models.Root_step, Root_stepDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scaling:
		tmp := GetInstanceDBFromInstance[models.Scaling, ScalingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scordatura:
		tmp := GetInstanceDBFromInstance[models.Scordatura, ScordaturaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_instrument:
		tmp := GetInstanceDBFromInstance[models.Score_instrument, Score_instrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_part:
		tmp := GetInstanceDBFromInstance[models.Score_part, Score_partDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_partwise:
		tmp := GetInstanceDBFromInstance[models.Score_partwise, Score_partwiseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_timewise:
		tmp := GetInstanceDBFromInstance[models.Score_timewise, Score_timewiseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Segno:
		tmp := GetInstanceDBFromInstance[models.Segno, SegnoDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slash:
		tmp := GetInstanceDBFromInstance[models.Slash, SlashDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slide:
		tmp := GetInstanceDBFromInstance[models.Slide, SlideDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slur:
		tmp := GetInstanceDBFromInstance[models.Slur, SlurDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Sound:
		tmp := GetInstanceDBFromInstance[models.Sound, SoundDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_details:
		tmp := GetInstanceDBFromInstance[models.Staff_details, Staff_detailsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_divide:
		tmp := GetInstanceDBFromInstance[models.Staff_divide, Staff_divideDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_layout:
		tmp := GetInstanceDBFromInstance[models.Staff_layout, Staff_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_size:
		tmp := GetInstanceDBFromInstance[models.Staff_size, Staff_sizeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_tuning:
		tmp := GetInstanceDBFromInstance[models.Staff_tuning, Staff_tuningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Stem:
		tmp := GetInstanceDBFromInstance[models.Stem, StemDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Stick:
		tmp := GetInstanceDBFromInstance[models.Stick, StickDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.String_mute:
		tmp := GetInstanceDBFromInstance[models.String_mute, String_muteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Strong_accent:
		tmp := GetInstanceDBFromInstance[models.Strong_accent, Strong_accentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Supports:
		tmp := GetInstanceDBFromInstance[models.Supports, SupportsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Swing:
		tmp := GetInstanceDBFromInstance[models.Swing, SwingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Sync:
		tmp := GetInstanceDBFromInstance[models.Sync, SyncDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.System_dividers:
		tmp := GetInstanceDBFromInstance[models.System_dividers, System_dividersDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.System_layout:
		tmp := GetInstanceDBFromInstance[models.System_layout, System_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.System_margins:
		tmp := GetInstanceDBFromInstance[models.System_margins, System_marginsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tap:
		tmp := GetInstanceDBFromInstance[models.Tap, TapDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Technical:
		tmp := GetInstanceDBFromInstance[models.Technical, TechnicalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Text_element_data:
		tmp := GetInstanceDBFromInstance[models.Text_element_data, Text_element_dataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tie:
		tmp := GetInstanceDBFromInstance[models.Tie, TieDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tied:
		tmp := GetInstanceDBFromInstance[models.Tied, TiedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Time:
		tmp := GetInstanceDBFromInstance[models.Time, TimeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Time_modification:
		tmp := GetInstanceDBFromInstance[models.Time_modification, Time_modificationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Timpani:
		tmp := GetInstanceDBFromInstance[models.Timpani, TimpaniDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Transpose:
		tmp := GetInstanceDBFromInstance[models.Transpose, TransposeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tremolo:
		tmp := GetInstanceDBFromInstance[models.Tremolo, TremoloDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet:
		tmp := GetInstanceDBFromInstance[models.Tuplet, TupletDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_dot:
		tmp := GetInstanceDBFromInstance[models.Tuplet_dot, Tuplet_dotDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_number:
		tmp := GetInstanceDBFromInstance[models.Tuplet_number, Tuplet_numberDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_portion:
		tmp := GetInstanceDBFromInstance[models.Tuplet_portion, Tuplet_portionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_type:
		tmp := GetInstanceDBFromInstance[models.Tuplet_type, Tuplet_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Typed_text:
		tmp := GetInstanceDBFromInstance[models.Typed_text, Typed_textDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Unpitched:
		tmp := GetInstanceDBFromInstance[models.Unpitched, UnpitchedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Virtual_instrument:
		tmp := GetInstanceDBFromInstance[models.Virtual_instrument, Virtual_instrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wait:
		tmp := GetInstanceDBFromInstance[models.Wait, WaitDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wavy_line:
		tmp := GetInstanceDBFromInstance[models.Wavy_line, Wavy_lineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wedge:
		tmp := GetInstanceDBFromInstance[models.Wedge, WedgeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wood:
		tmp := GetInstanceDBFromInstance[models.Wood, WoodDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Work:
		tmp := GetInstanceDBFromInstance[models.Work, WorkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Accidental:
		tmp := GetInstanceDBFromInstance[models.Accidental, AccidentalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accidental_mark:
		tmp := GetInstanceDBFromInstance[models.Accidental_mark, Accidental_markDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accidental_text:
		tmp := GetInstanceDBFromInstance[models.Accidental_text, Accidental_textDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accord:
		tmp := GetInstanceDBFromInstance[models.Accord, AccordDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Accordion_registration:
		tmp := GetInstanceDBFromInstance[models.Accordion_registration, Accordion_registrationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.AnyType:
		tmp := GetInstanceDBFromInstance[models.AnyType, AnyTypeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Appearance:
		tmp := GetInstanceDBFromInstance[models.Appearance, AppearanceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Arpeggiate:
		tmp := GetInstanceDBFromInstance[models.Arpeggiate, ArpeggiateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Articulations:
		tmp := GetInstanceDBFromInstance[models.Articulations, ArticulationsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Assess:
		tmp := GetInstanceDBFromInstance[models.Assess, AssessDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Attributes:
		tmp := GetInstanceDBFromInstance[models.Attributes, AttributesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Backup:
		tmp := GetInstanceDBFromInstance[models.Backup, BackupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bar_style_color:
		tmp := GetInstanceDBFromInstance[models.Bar_style_color, Bar_style_colorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Barline:
		tmp := GetInstanceDBFromInstance[models.Barline, BarlineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Barre:
		tmp := GetInstanceDBFromInstance[models.Barre, BarreDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bass:
		tmp := GetInstanceDBFromInstance[models.Bass, BassDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bass_step:
		tmp := GetInstanceDBFromInstance[models.Bass_step, Bass_stepDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beam:
		tmp := GetInstanceDBFromInstance[models.Beam, BeamDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beat_repeat:
		tmp := GetInstanceDBFromInstance[models.Beat_repeat, Beat_repeatDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beat_unit_tied:
		tmp := GetInstanceDBFromInstance[models.Beat_unit_tied, Beat_unit_tiedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Beater:
		tmp := GetInstanceDBFromInstance[models.Beater, BeaterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bend:
		tmp := GetInstanceDBFromInstance[models.Bend, BendDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bookmark:
		tmp := GetInstanceDBFromInstance[models.Bookmark, BookmarkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Bracket:
		tmp := GetInstanceDBFromInstance[models.Bracket, BracketDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Breath_mark:
		tmp := GetInstanceDBFromInstance[models.Breath_mark, Breath_markDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Caesura:
		tmp := GetInstanceDBFromInstance[models.Caesura, CaesuraDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Cancel:
		tmp := GetInstanceDBFromInstance[models.Cancel, CancelDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Clef:
		tmp := GetInstanceDBFromInstance[models.Clef, ClefDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Coda:
		tmp := GetInstanceDBFromInstance[models.Coda, CodaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Credit:
		tmp := GetInstanceDBFromInstance[models.Credit, CreditDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Dashes:
		tmp := GetInstanceDBFromInstance[models.Dashes, DashesDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Defaults:
		tmp := GetInstanceDBFromInstance[models.Defaults, DefaultsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree:
		tmp := GetInstanceDBFromInstance[models.Degree, DegreeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree_alter:
		tmp := GetInstanceDBFromInstance[models.Degree_alter, Degree_alterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree_type:
		tmp := GetInstanceDBFromInstance[models.Degree_type, Degree_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Degree_value:
		tmp := GetInstanceDBFromInstance[models.Degree_value, Degree_valueDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Direction:
		tmp := GetInstanceDBFromInstance[models.Direction, DirectionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Direction_type:
		tmp := GetInstanceDBFromInstance[models.Direction_type, Direction_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Distance:
		tmp := GetInstanceDBFromInstance[models.Distance, DistanceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Double:
		tmp := GetInstanceDBFromInstance[models.Double, DoubleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Dynamics:
		tmp := GetInstanceDBFromInstance[models.Dynamics, DynamicsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Effect:
		tmp := GetInstanceDBFromInstance[models.Effect, EffectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Elision:
		tmp := GetInstanceDBFromInstance[models.Elision, ElisionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty:
		tmp := GetInstanceDBFromInstance[models.Empty, EmptyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_font:
		tmp := GetInstanceDBFromInstance[models.Empty_font, Empty_fontDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_line:
		tmp := GetInstanceDBFromInstance[models.Empty_line, Empty_lineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_placement:
		tmp := GetInstanceDBFromInstance[models.Empty_placement, Empty_placementDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_placement_smufl:
		tmp := GetInstanceDBFromInstance[models.Empty_placement_smufl, Empty_placement_smuflDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_object_style_align:
		tmp := GetInstanceDBFromInstance[models.Empty_print_object_style_align, Empty_print_object_style_alignDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_style:
		tmp := GetInstanceDBFromInstance[models.Empty_print_style, Empty_print_styleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_style_align:
		tmp := GetInstanceDBFromInstance[models.Empty_print_style_align, Empty_print_style_alignDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_print_style_align_id:
		tmp := GetInstanceDBFromInstance[models.Empty_print_style_align_id, Empty_print_style_align_idDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Empty_trill_sound:
		tmp := GetInstanceDBFromInstance[models.Empty_trill_sound, Empty_trill_soundDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Encoding:
		tmp := GetInstanceDBFromInstance[models.Encoding, EncodingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Ending:
		tmp := GetInstanceDBFromInstance[models.Ending, EndingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Extend:
		tmp := GetInstanceDBFromInstance[models.Extend, ExtendDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Feature:
		tmp := GetInstanceDBFromInstance[models.Feature, FeatureDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Fermata:
		tmp := GetInstanceDBFromInstance[models.Fermata, FermataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Figure:
		tmp := GetInstanceDBFromInstance[models.Figure, FigureDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Figured_bass:
		tmp := GetInstanceDBFromInstance[models.Figured_bass, Figured_bassDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Fingering:
		tmp := GetInstanceDBFromInstance[models.Fingering, FingeringDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.First_fret:
		tmp := GetInstanceDBFromInstance[models.First_fret, First_fretDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Foo:
		tmp := GetInstanceDBFromInstance[models.Foo, FooDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.For_part:
		tmp := GetInstanceDBFromInstance[models.For_part, For_partDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Formatted_symbol:
		tmp := GetInstanceDBFromInstance[models.Formatted_symbol, Formatted_symbolDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Formatted_symbol_id:
		tmp := GetInstanceDBFromInstance[models.Formatted_symbol_id, Formatted_symbol_idDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Forward:
		tmp := GetInstanceDBFromInstance[models.Forward, ForwardDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Frame:
		tmp := GetInstanceDBFromInstance[models.Frame, FrameDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Frame_note:
		tmp := GetInstanceDBFromInstance[models.Frame_note, Frame_noteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Fret:
		tmp := GetInstanceDBFromInstance[models.Fret, FretDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Glass:
		tmp := GetInstanceDBFromInstance[models.Glass, GlassDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Glissando:
		tmp := GetInstanceDBFromInstance[models.Glissando, GlissandoDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Glyph:
		tmp := GetInstanceDBFromInstance[models.Glyph, GlyphDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Grace:
		tmp := GetInstanceDBFromInstance[models.Grace, GraceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group_barline:
		tmp := GetInstanceDBFromInstance[models.Group_barline, Group_barlineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Group_symbol:
		tmp := GetInstanceDBFromInstance[models.Group_symbol, Group_symbolDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Grouping:
		tmp := GetInstanceDBFromInstance[models.Grouping, GroupingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hammer_on_pull_off:
		tmp := GetInstanceDBFromInstance[models.Hammer_on_pull_off, Hammer_on_pull_offDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Handbell:
		tmp := GetInstanceDBFromInstance[models.Handbell, HandbellDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmon_closed:
		tmp := GetInstanceDBFromInstance[models.Harmon_closed, Harmon_closedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmon_mute:
		tmp := GetInstanceDBFromInstance[models.Harmon_mute, Harmon_muteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmonic:
		tmp := GetInstanceDBFromInstance[models.Harmonic, HarmonicDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmony:
		tmp := GetInstanceDBFromInstance[models.Harmony, HarmonyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harmony_alter:
		tmp := GetInstanceDBFromInstance[models.Harmony_alter, Harmony_alterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Harp_pedals:
		tmp := GetInstanceDBFromInstance[models.Harp_pedals, Harp_pedalsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Heel_toe:
		tmp := GetInstanceDBFromInstance[models.Heel_toe, Heel_toeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hole:
		tmp := GetInstanceDBFromInstance[models.Hole, HoleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hole_closed:
		tmp := GetInstanceDBFromInstance[models.Hole_closed, Hole_closedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Horizontal_turn:
		tmp := GetInstanceDBFromInstance[models.Horizontal_turn, Horizontal_turnDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Identification:
		tmp := GetInstanceDBFromInstance[models.Identification, IdentificationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Image:
		tmp := GetInstanceDBFromInstance[models.Image, ImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Instrument:
		tmp := GetInstanceDBFromInstance[models.Instrument, InstrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Instrument_change:
		tmp := GetInstanceDBFromInstance[models.Instrument_change, Instrument_changeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Instrument_link:
		tmp := GetInstanceDBFromInstance[models.Instrument_link, Instrument_linkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Interchangeable:
		tmp := GetInstanceDBFromInstance[models.Interchangeable, InterchangeableDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Inversion:
		tmp := GetInstanceDBFromInstance[models.Inversion, InversionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key:
		tmp := GetInstanceDBFromInstance[models.Key, KeyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key_accidental:
		tmp := GetInstanceDBFromInstance[models.Key_accidental, Key_accidentalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Key_octave:
		tmp := GetInstanceDBFromInstance[models.Key_octave, Key_octaveDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Kind:
		tmp := GetInstanceDBFromInstance[models.Kind, KindDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Level:
		tmp := GetInstanceDBFromInstance[models.Level, LevelDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line_detail:
		tmp := GetInstanceDBFromInstance[models.Line_detail, Line_detailDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Line_width:
		tmp := GetInstanceDBFromInstance[models.Line_width, Line_widthDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Link:
		tmp := GetInstanceDBFromInstance[models.Link, LinkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Listen:
		tmp := GetInstanceDBFromInstance[models.Listen, ListenDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Listening:
		tmp := GetInstanceDBFromInstance[models.Listening, ListeningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric:
		tmp := GetInstanceDBFromInstance[models.Lyric, LyricDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_font:
		tmp := GetInstanceDBFromInstance[models.Lyric_font, Lyric_fontDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Lyric_language:
		tmp := GetInstanceDBFromInstance[models.Lyric_language, Lyric_languageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_layout:
		tmp := GetInstanceDBFromInstance[models.Measure_layout, Measure_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_numbering:
		tmp := GetInstanceDBFromInstance[models.Measure_numbering, Measure_numberingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_repeat:
		tmp := GetInstanceDBFromInstance[models.Measure_repeat, Measure_repeatDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Measure_style:
		tmp := GetInstanceDBFromInstance[models.Measure_style, Measure_styleDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Membrane:
		tmp := GetInstanceDBFromInstance[models.Membrane, MembraneDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metal:
		tmp := GetInstanceDBFromInstance[models.Metal, MetalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome:
		tmp := GetInstanceDBFromInstance[models.Metronome, MetronomeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_beam:
		tmp := GetInstanceDBFromInstance[models.Metronome_beam, Metronome_beamDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_note:
		tmp := GetInstanceDBFromInstance[models.Metronome_note, Metronome_noteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_tied:
		tmp := GetInstanceDBFromInstance[models.Metronome_tied, Metronome_tiedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Metronome_tuplet:
		tmp := GetInstanceDBFromInstance[models.Metronome_tuplet, Metronome_tupletDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Midi_device:
		tmp := GetInstanceDBFromInstance[models.Midi_device, Midi_deviceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Midi_instrument:
		tmp := GetInstanceDBFromInstance[models.Midi_instrument, Midi_instrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Miscellaneous:
		tmp := GetInstanceDBFromInstance[models.Miscellaneous, MiscellaneousDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Miscellaneous_field:
		tmp := GetInstanceDBFromInstance[models.Miscellaneous_field, Miscellaneous_fieldDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Mordent:
		tmp := GetInstanceDBFromInstance[models.Mordent, MordentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Multiple_rest:
		tmp := GetInstanceDBFromInstance[models.Multiple_rest, Multiple_restDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Name_display:
		tmp := GetInstanceDBFromInstance[models.Name_display, Name_displayDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Non_arpeggiate:
		tmp := GetInstanceDBFromInstance[models.Non_arpeggiate, Non_arpeggiateDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Notations:
		tmp := GetInstanceDBFromInstance[models.Notations, NotationsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note:
		tmp := GetInstanceDBFromInstance[models.Note, NoteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note_size:
		tmp := GetInstanceDBFromInstance[models.Note_size, Note_sizeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Note_type:
		tmp := GetInstanceDBFromInstance[models.Note_type, Note_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Notehead:
		tmp := GetInstanceDBFromInstance[models.Notehead, NoteheadDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Notehead_text:
		tmp := GetInstanceDBFromInstance[models.Notehead_text, Notehead_textDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Numeral:
		tmp := GetInstanceDBFromInstance[models.Numeral, NumeralDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Numeral_key:
		tmp := GetInstanceDBFromInstance[models.Numeral_key, Numeral_keyDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Numeral_root:
		tmp := GetInstanceDBFromInstance[models.Numeral_root, Numeral_rootDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Octave_shift:
		tmp := GetInstanceDBFromInstance[models.Octave_shift, Octave_shiftDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Offset:
		tmp := GetInstanceDBFromInstance[models.Offset, OffsetDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Opus:
		tmp := GetInstanceDBFromInstance[models.Opus, OpusDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Ornaments:
		tmp := GetInstanceDBFromInstance[models.Ornaments, OrnamentsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_appearance:
		tmp := GetInstanceDBFromInstance[models.Other_appearance, Other_appearanceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_listening:
		tmp := GetInstanceDBFromInstance[models.Other_listening, Other_listeningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_notation:
		tmp := GetInstanceDBFromInstance[models.Other_notation, Other_notationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Other_play:
		tmp := GetInstanceDBFromInstance[models.Other_play, Other_playDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Page_layout:
		tmp := GetInstanceDBFromInstance[models.Page_layout, Page_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Page_margins:
		tmp := GetInstanceDBFromInstance[models.Page_margins, Page_marginsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_clef:
		tmp := GetInstanceDBFromInstance[models.Part_clef, Part_clefDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_group:
		tmp := GetInstanceDBFromInstance[models.Part_group, Part_groupDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_link:
		tmp := GetInstanceDBFromInstance[models.Part_link, Part_linkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_list:
		tmp := GetInstanceDBFromInstance[models.Part_list, Part_listDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_symbol:
		tmp := GetInstanceDBFromInstance[models.Part_symbol, Part_symbolDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Part_transpose:
		tmp := GetInstanceDBFromInstance[models.Part_transpose, Part_transposeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pedal:
		tmp := GetInstanceDBFromInstance[models.Pedal, PedalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pedal_tuning:
		tmp := GetInstanceDBFromInstance[models.Pedal_tuning, Pedal_tuningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Percussion:
		tmp := GetInstanceDBFromInstance[models.Percussion, PercussionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pitch:
		tmp := GetInstanceDBFromInstance[models.Pitch, PitchDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Pitched:
		tmp := GetInstanceDBFromInstance[models.Pitched, PitchedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Play:
		tmp := GetInstanceDBFromInstance[models.Play, PlayDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Player:
		tmp := GetInstanceDBFromInstance[models.Player, PlayerDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Principal_voice:
		tmp := GetInstanceDBFromInstance[models.Principal_voice, Principal_voiceDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Print:
		tmp := GetInstanceDBFromInstance[models.Print, PrintDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Release:
		tmp := GetInstanceDBFromInstance[models.Release, ReleaseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Repeat:
		tmp := GetInstanceDBFromInstance[models.Repeat, RepeatDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Rest:
		tmp := GetInstanceDBFromInstance[models.Rest, RestDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Root:
		tmp := GetInstanceDBFromInstance[models.Root, RootDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Root_step:
		tmp := GetInstanceDBFromInstance[models.Root_step, Root_stepDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scaling:
		tmp := GetInstanceDBFromInstance[models.Scaling, ScalingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Scordatura:
		tmp := GetInstanceDBFromInstance[models.Scordatura, ScordaturaDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_instrument:
		tmp := GetInstanceDBFromInstance[models.Score_instrument, Score_instrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_part:
		tmp := GetInstanceDBFromInstance[models.Score_part, Score_partDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_partwise:
		tmp := GetInstanceDBFromInstance[models.Score_partwise, Score_partwiseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Score_timewise:
		tmp := GetInstanceDBFromInstance[models.Score_timewise, Score_timewiseDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Segno:
		tmp := GetInstanceDBFromInstance[models.Segno, SegnoDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slash:
		tmp := GetInstanceDBFromInstance[models.Slash, SlashDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slide:
		tmp := GetInstanceDBFromInstance[models.Slide, SlideDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Slur:
		tmp := GetInstanceDBFromInstance[models.Slur, SlurDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Sound:
		tmp := GetInstanceDBFromInstance[models.Sound, SoundDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_details:
		tmp := GetInstanceDBFromInstance[models.Staff_details, Staff_detailsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_divide:
		tmp := GetInstanceDBFromInstance[models.Staff_divide, Staff_divideDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_layout:
		tmp := GetInstanceDBFromInstance[models.Staff_layout, Staff_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_size:
		tmp := GetInstanceDBFromInstance[models.Staff_size, Staff_sizeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Staff_tuning:
		tmp := GetInstanceDBFromInstance[models.Staff_tuning, Staff_tuningDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Stem:
		tmp := GetInstanceDBFromInstance[models.Stem, StemDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Stick:
		tmp := GetInstanceDBFromInstance[models.Stick, StickDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.String_mute:
		tmp := GetInstanceDBFromInstance[models.String_mute, String_muteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Strong_accent:
		tmp := GetInstanceDBFromInstance[models.Strong_accent, Strong_accentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Supports:
		tmp := GetInstanceDBFromInstance[models.Supports, SupportsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Swing:
		tmp := GetInstanceDBFromInstance[models.Swing, SwingDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Sync:
		tmp := GetInstanceDBFromInstance[models.Sync, SyncDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.System_dividers:
		tmp := GetInstanceDBFromInstance[models.System_dividers, System_dividersDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.System_layout:
		tmp := GetInstanceDBFromInstance[models.System_layout, System_layoutDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.System_margins:
		tmp := GetInstanceDBFromInstance[models.System_margins, System_marginsDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tap:
		tmp := GetInstanceDBFromInstance[models.Tap, TapDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Technical:
		tmp := GetInstanceDBFromInstance[models.Technical, TechnicalDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Text_element_data:
		tmp := GetInstanceDBFromInstance[models.Text_element_data, Text_element_dataDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tie:
		tmp := GetInstanceDBFromInstance[models.Tie, TieDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tied:
		tmp := GetInstanceDBFromInstance[models.Tied, TiedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Time:
		tmp := GetInstanceDBFromInstance[models.Time, TimeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Time_modification:
		tmp := GetInstanceDBFromInstance[models.Time_modification, Time_modificationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Timpani:
		tmp := GetInstanceDBFromInstance[models.Timpani, TimpaniDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Transpose:
		tmp := GetInstanceDBFromInstance[models.Transpose, TransposeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tremolo:
		tmp := GetInstanceDBFromInstance[models.Tremolo, TremoloDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet:
		tmp := GetInstanceDBFromInstance[models.Tuplet, TupletDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_dot:
		tmp := GetInstanceDBFromInstance[models.Tuplet_dot, Tuplet_dotDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_number:
		tmp := GetInstanceDBFromInstance[models.Tuplet_number, Tuplet_numberDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_portion:
		tmp := GetInstanceDBFromInstance[models.Tuplet_portion, Tuplet_portionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Tuplet_type:
		tmp := GetInstanceDBFromInstance[models.Tuplet_type, Tuplet_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Typed_text:
		tmp := GetInstanceDBFromInstance[models.Typed_text, Typed_textDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Unpitched:
		tmp := GetInstanceDBFromInstance[models.Unpitched, UnpitchedDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Virtual_instrument:
		tmp := GetInstanceDBFromInstance[models.Virtual_instrument, Virtual_instrumentDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wait:
		tmp := GetInstanceDBFromInstance[models.Wait, WaitDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wavy_line:
		tmp := GetInstanceDBFromInstance[models.Wavy_line, Wavy_lineDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wedge:
		tmp := GetInstanceDBFromInstance[models.Wedge, WedgeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Wood:
		tmp := GetInstanceDBFromInstance[models.Wood, WoodDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Work:
		tmp := GetInstanceDBFromInstance[models.Work, WorkDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
