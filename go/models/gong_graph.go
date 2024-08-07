// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Accidental:
		ok = stage.IsStagedAccidental(target)

	case *Accidental_mark:
		ok = stage.IsStagedAccidental_mark(target)

	case *Accidental_text:
		ok = stage.IsStagedAccidental_text(target)

	case *Accord:
		ok = stage.IsStagedAccord(target)

	case *Accordion_registration:
		ok = stage.IsStagedAccordion_registration(target)

	case *AnyType:
		ok = stage.IsStagedAnyType(target)

	case *Appearance:
		ok = stage.IsStagedAppearance(target)

	case *Arpeggiate:
		ok = stage.IsStagedArpeggiate(target)

	case *Arrow:
		ok = stage.IsStagedArrow(target)

	case *Articulations:
		ok = stage.IsStagedArticulations(target)

	case *Assess:
		ok = stage.IsStagedAssess(target)

	case *Attributes:
		ok = stage.IsStagedAttributes(target)

	case *Backup:
		ok = stage.IsStagedBackup(target)

	case *Bar_style_color:
		ok = stage.IsStagedBar_style_color(target)

	case *Barline:
		ok = stage.IsStagedBarline(target)

	case *Barre:
		ok = stage.IsStagedBarre(target)

	case *Bass:
		ok = stage.IsStagedBass(target)

	case *Bass_step:
		ok = stage.IsStagedBass_step(target)

	case *Beam:
		ok = stage.IsStagedBeam(target)

	case *Beat_repeat:
		ok = stage.IsStagedBeat_repeat(target)

	case *Beat_unit_tied:
		ok = stage.IsStagedBeat_unit_tied(target)

	case *Beater:
		ok = stage.IsStagedBeater(target)

	case *Bend:
		ok = stage.IsStagedBend(target)

	case *Bookmark:
		ok = stage.IsStagedBookmark(target)

	case *Bracket:
		ok = stage.IsStagedBracket(target)

	case *Breath_mark:
		ok = stage.IsStagedBreath_mark(target)

	case *Caesura:
		ok = stage.IsStagedCaesura(target)

	case *Cancel:
		ok = stage.IsStagedCancel(target)

	case *Clef:
		ok = stage.IsStagedClef(target)

	case *Coda:
		ok = stage.IsStagedCoda(target)

	case *Credit:
		ok = stage.IsStagedCredit(target)

	case *Dashes:
		ok = stage.IsStagedDashes(target)

	case *Defaults:
		ok = stage.IsStagedDefaults(target)

	case *Degree:
		ok = stage.IsStagedDegree(target)

	case *Degree_alter:
		ok = stage.IsStagedDegree_alter(target)

	case *Degree_type:
		ok = stage.IsStagedDegree_type(target)

	case *Degree_value:
		ok = stage.IsStagedDegree_value(target)

	case *Direction:
		ok = stage.IsStagedDirection(target)

	case *Direction_type:
		ok = stage.IsStagedDirection_type(target)

	case *Distance:
		ok = stage.IsStagedDistance(target)

	case *Double:
		ok = stage.IsStagedDouble(target)

	case *Dynamics:
		ok = stage.IsStagedDynamics(target)

	case *Effect:
		ok = stage.IsStagedEffect(target)

	case *Elision:
		ok = stage.IsStagedElision(target)

	case *Empty:
		ok = stage.IsStagedEmpty(target)

	case *Empty_font:
		ok = stage.IsStagedEmpty_font(target)

	case *Empty_line:
		ok = stage.IsStagedEmpty_line(target)

	case *Empty_placement:
		ok = stage.IsStagedEmpty_placement(target)

	case *Empty_placement_smufl:
		ok = stage.IsStagedEmpty_placement_smufl(target)

	case *Empty_print_object_style_align:
		ok = stage.IsStagedEmpty_print_object_style_align(target)

	case *Empty_print_style:
		ok = stage.IsStagedEmpty_print_style(target)

	case *Empty_print_style_align:
		ok = stage.IsStagedEmpty_print_style_align(target)

	case *Empty_print_style_align_id:
		ok = stage.IsStagedEmpty_print_style_align_id(target)

	case *Empty_trill_sound:
		ok = stage.IsStagedEmpty_trill_sound(target)

	case *Encoding:
		ok = stage.IsStagedEncoding(target)

	case *Ending:
		ok = stage.IsStagedEnding(target)

	case *Extend:
		ok = stage.IsStagedExtend(target)

	case *Feature:
		ok = stage.IsStagedFeature(target)

	case *Fermata:
		ok = stage.IsStagedFermata(target)

	case *Figure:
		ok = stage.IsStagedFigure(target)

	case *Figured_bass:
		ok = stage.IsStagedFigured_bass(target)

	case *Fingering:
		ok = stage.IsStagedFingering(target)

	case *First_fret:
		ok = stage.IsStagedFirst_fret(target)

	case *Foo:
		ok = stage.IsStagedFoo(target)

	case *For_part:
		ok = stage.IsStagedFor_part(target)

	case *Formatted_symbol:
		ok = stage.IsStagedFormatted_symbol(target)

	case *Formatted_symbol_id:
		ok = stage.IsStagedFormatted_symbol_id(target)

	case *Forward:
		ok = stage.IsStagedForward(target)

	case *Frame:
		ok = stage.IsStagedFrame(target)

	case *Frame_note:
		ok = stage.IsStagedFrame_note(target)

	case *Fret:
		ok = stage.IsStagedFret(target)

	case *Glass:
		ok = stage.IsStagedGlass(target)

	case *Glissando:
		ok = stage.IsStagedGlissando(target)

	case *Glyph:
		ok = stage.IsStagedGlyph(target)

	case *Grace:
		ok = stage.IsStagedGrace(target)

	case *Group_barline:
		ok = stage.IsStagedGroup_barline(target)

	case *Group_symbol:
		ok = stage.IsStagedGroup_symbol(target)

	case *Grouping:
		ok = stage.IsStagedGrouping(target)

	case *Hammer_on_pull_off:
		ok = stage.IsStagedHammer_on_pull_off(target)

	case *Handbell:
		ok = stage.IsStagedHandbell(target)

	case *Harmon_closed:
		ok = stage.IsStagedHarmon_closed(target)

	case *Harmon_mute:
		ok = stage.IsStagedHarmon_mute(target)

	case *Harmonic:
		ok = stage.IsStagedHarmonic(target)

	case *Harmony:
		ok = stage.IsStagedHarmony(target)

	case *Harmony_alter:
		ok = stage.IsStagedHarmony_alter(target)

	case *Harp_pedals:
		ok = stage.IsStagedHarp_pedals(target)

	case *Heel_toe:
		ok = stage.IsStagedHeel_toe(target)

	case *Hole:
		ok = stage.IsStagedHole(target)

	case *Hole_closed:
		ok = stage.IsStagedHole_closed(target)

	case *Horizontal_turn:
		ok = stage.IsStagedHorizontal_turn(target)

	case *Identification:
		ok = stage.IsStagedIdentification(target)

	case *Image:
		ok = stage.IsStagedImage(target)

	case *Instrument:
		ok = stage.IsStagedInstrument(target)

	case *Instrument_change:
		ok = stage.IsStagedInstrument_change(target)

	case *Instrument_link:
		ok = stage.IsStagedInstrument_link(target)

	case *Interchangeable:
		ok = stage.IsStagedInterchangeable(target)

	case *Inversion:
		ok = stage.IsStagedInversion(target)

	case *Key:
		ok = stage.IsStagedKey(target)

	case *Key_accidental:
		ok = stage.IsStagedKey_accidental(target)

	case *Key_octave:
		ok = stage.IsStagedKey_octave(target)

	case *Kind:
		ok = stage.IsStagedKind(target)

	case *Level:
		ok = stage.IsStagedLevel(target)

	case *Line_detail:
		ok = stage.IsStagedLine_detail(target)

	case *Line_width:
		ok = stage.IsStagedLine_width(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *Listen:
		ok = stage.IsStagedListen(target)

	case *Listening:
		ok = stage.IsStagedListening(target)

	case *Lyric:
		ok = stage.IsStagedLyric(target)

	case *Lyric_font:
		ok = stage.IsStagedLyric_font(target)

	case *Lyric_language:
		ok = stage.IsStagedLyric_language(target)

	case *Measure_layout:
		ok = stage.IsStagedMeasure_layout(target)

	case *Measure_numbering:
		ok = stage.IsStagedMeasure_numbering(target)

	case *Measure_repeat:
		ok = stage.IsStagedMeasure_repeat(target)

	case *Measure_style:
		ok = stage.IsStagedMeasure_style(target)

	case *Membrane:
		ok = stage.IsStagedMembrane(target)

	case *Metal:
		ok = stage.IsStagedMetal(target)

	case *Metronome:
		ok = stage.IsStagedMetronome(target)

	case *Metronome_beam:
		ok = stage.IsStagedMetronome_beam(target)

	case *Metronome_note:
		ok = stage.IsStagedMetronome_note(target)

	case *Metronome_tied:
		ok = stage.IsStagedMetronome_tied(target)

	case *Metronome_tuplet:
		ok = stage.IsStagedMetronome_tuplet(target)

	case *Midi_device:
		ok = stage.IsStagedMidi_device(target)

	case *Midi_instrument:
		ok = stage.IsStagedMidi_instrument(target)

	case *Miscellaneous:
		ok = stage.IsStagedMiscellaneous(target)

	case *Miscellaneous_field:
		ok = stage.IsStagedMiscellaneous_field(target)

	case *Mordent:
		ok = stage.IsStagedMordent(target)

	case *Multiple_rest:
		ok = stage.IsStagedMultiple_rest(target)

	case *Name_display:
		ok = stage.IsStagedName_display(target)

	case *Non_arpeggiate:
		ok = stage.IsStagedNon_arpeggiate(target)

	case *Notations:
		ok = stage.IsStagedNotations(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *Note_size:
		ok = stage.IsStagedNote_size(target)

	case *Note_type:
		ok = stage.IsStagedNote_type(target)

	case *Notehead:
		ok = stage.IsStagedNotehead(target)

	case *Notehead_text:
		ok = stage.IsStagedNotehead_text(target)

	case *Numeral:
		ok = stage.IsStagedNumeral(target)

	case *Numeral_key:
		ok = stage.IsStagedNumeral_key(target)

	case *Numeral_root:
		ok = stage.IsStagedNumeral_root(target)

	case *Octave_shift:
		ok = stage.IsStagedOctave_shift(target)

	case *Offset:
		ok = stage.IsStagedOffset(target)

	case *Opus:
		ok = stage.IsStagedOpus(target)

	case *Ornaments:
		ok = stage.IsStagedOrnaments(target)

	case *Other_appearance:
		ok = stage.IsStagedOther_appearance(target)

	case *Other_listening:
		ok = stage.IsStagedOther_listening(target)

	case *Other_notation:
		ok = stage.IsStagedOther_notation(target)

	case *Other_play:
		ok = stage.IsStagedOther_play(target)

	case *Page_layout:
		ok = stage.IsStagedPage_layout(target)

	case *Page_margins:
		ok = stage.IsStagedPage_margins(target)

	case *Part_clef:
		ok = stage.IsStagedPart_clef(target)

	case *Part_group:
		ok = stage.IsStagedPart_group(target)

	case *Part_link:
		ok = stage.IsStagedPart_link(target)

	case *Part_list:
		ok = stage.IsStagedPart_list(target)

	case *Part_symbol:
		ok = stage.IsStagedPart_symbol(target)

	case *Part_transpose:
		ok = stage.IsStagedPart_transpose(target)

	case *Pedal:
		ok = stage.IsStagedPedal(target)

	case *Pedal_tuning:
		ok = stage.IsStagedPedal_tuning(target)

	case *Percussion:
		ok = stage.IsStagedPercussion(target)

	case *Pitch:
		ok = stage.IsStagedPitch(target)

	case *Pitched:
		ok = stage.IsStagedPitched(target)

	case *Play:
		ok = stage.IsStagedPlay(target)

	case *Player:
		ok = stage.IsStagedPlayer(target)

	case *Principal_voice:
		ok = stage.IsStagedPrincipal_voice(target)

	case *Print:
		ok = stage.IsStagedPrint(target)

	case *Release:
		ok = stage.IsStagedRelease(target)

	case *Repeat:
		ok = stage.IsStagedRepeat(target)

	case *Rest:
		ok = stage.IsStagedRest(target)

	case *Root:
		ok = stage.IsStagedRoot(target)

	case *Root_step:
		ok = stage.IsStagedRoot_step(target)

	case *Scaling:
		ok = stage.IsStagedScaling(target)

	case *Scordatura:
		ok = stage.IsStagedScordatura(target)

	case *Score_instrument:
		ok = stage.IsStagedScore_instrument(target)

	case *Score_part:
		ok = stage.IsStagedScore_part(target)

	case *Score_partwise:
		ok = stage.IsStagedScore_partwise(target)

	case *Score_timewise:
		ok = stage.IsStagedScore_timewise(target)

	case *Segno:
		ok = stage.IsStagedSegno(target)

	case *Slash:
		ok = stage.IsStagedSlash(target)

	case *Slide:
		ok = stage.IsStagedSlide(target)

	case *Slur:
		ok = stage.IsStagedSlur(target)

	case *Sound:
		ok = stage.IsStagedSound(target)

	case *Staff_details:
		ok = stage.IsStagedStaff_details(target)

	case *Staff_divide:
		ok = stage.IsStagedStaff_divide(target)

	case *Staff_layout:
		ok = stage.IsStagedStaff_layout(target)

	case *Staff_size:
		ok = stage.IsStagedStaff_size(target)

	case *Staff_tuning:
		ok = stage.IsStagedStaff_tuning(target)

	case *Stem:
		ok = stage.IsStagedStem(target)

	case *Stick:
		ok = stage.IsStagedStick(target)

	case *String_mute:
		ok = stage.IsStagedString_mute(target)

	case *Strong_accent:
		ok = stage.IsStagedStrong_accent(target)

	case *Supports:
		ok = stage.IsStagedSupports(target)

	case *Swing:
		ok = stage.IsStagedSwing(target)

	case *Sync:
		ok = stage.IsStagedSync(target)

	case *System_dividers:
		ok = stage.IsStagedSystem_dividers(target)

	case *System_layout:
		ok = stage.IsStagedSystem_layout(target)

	case *System_margins:
		ok = stage.IsStagedSystem_margins(target)

	case *Tap:
		ok = stage.IsStagedTap(target)

	case *Technical:
		ok = stage.IsStagedTechnical(target)

	case *Text_element_data:
		ok = stage.IsStagedText_element_data(target)

	case *Tie:
		ok = stage.IsStagedTie(target)

	case *Tied:
		ok = stage.IsStagedTied(target)

	case *Time:
		ok = stage.IsStagedTime(target)

	case *Time_modification:
		ok = stage.IsStagedTime_modification(target)

	case *Timpani:
		ok = stage.IsStagedTimpani(target)

	case *Transpose:
		ok = stage.IsStagedTranspose(target)

	case *Tremolo:
		ok = stage.IsStagedTremolo(target)

	case *Tuplet:
		ok = stage.IsStagedTuplet(target)

	case *Tuplet_dot:
		ok = stage.IsStagedTuplet_dot(target)

	case *Tuplet_number:
		ok = stage.IsStagedTuplet_number(target)

	case *Tuplet_portion:
		ok = stage.IsStagedTuplet_portion(target)

	case *Tuplet_type:
		ok = stage.IsStagedTuplet_type(target)

	case *Typed_text:
		ok = stage.IsStagedTyped_text(target)

	case *Unpitched:
		ok = stage.IsStagedUnpitched(target)

	case *Virtual_instrument:
		ok = stage.IsStagedVirtual_instrument(target)

	case *Wait:
		ok = stage.IsStagedWait(target)

	case *Wavy_line:
		ok = stage.IsStagedWavy_line(target)

	case *Wedge:
		ok = stage.IsStagedWedge(target)

	case *Wood:
		ok = stage.IsStagedWood(target)

	case *Work:
		ok = stage.IsStagedWork(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedAccidental(accidental *Accidental) (ok bool) {

	_, ok = stage.Accidentals[accidental]

	return
}

func (stage *StageStruct) IsStagedAccidental_mark(accidental_mark *Accidental_mark) (ok bool) {

	_, ok = stage.Accidental_marks[accidental_mark]

	return
}

func (stage *StageStruct) IsStagedAccidental_text(accidental_text *Accidental_text) (ok bool) {

	_, ok = stage.Accidental_texts[accidental_text]

	return
}

func (stage *StageStruct) IsStagedAccord(accord *Accord) (ok bool) {

	_, ok = stage.Accords[accord]

	return
}

func (stage *StageStruct) IsStagedAccordion_registration(accordion_registration *Accordion_registration) (ok bool) {

	_, ok = stage.Accordion_registrations[accordion_registration]

	return
}

func (stage *StageStruct) IsStagedAnyType(anytype *AnyType) (ok bool) {

	_, ok = stage.AnyTypes[anytype]

	return
}

func (stage *StageStruct) IsStagedAppearance(appearance *Appearance) (ok bool) {

	_, ok = stage.Appearances[appearance]

	return
}

func (stage *StageStruct) IsStagedArpeggiate(arpeggiate *Arpeggiate) (ok bool) {

	_, ok = stage.Arpeggiates[arpeggiate]

	return
}

func (stage *StageStruct) IsStagedArrow(arrow *Arrow) (ok bool) {

	_, ok = stage.Arrows[arrow]

	return
}

func (stage *StageStruct) IsStagedArticulations(articulations *Articulations) (ok bool) {

	_, ok = stage.Articulationss[articulations]

	return
}

func (stage *StageStruct) IsStagedAssess(assess *Assess) (ok bool) {

	_, ok = stage.Assesss[assess]

	return
}

func (stage *StageStruct) IsStagedAttributes(attributes *Attributes) (ok bool) {

	_, ok = stage.Attributess[attributes]

	return
}

func (stage *StageStruct) IsStagedBackup(backup *Backup) (ok bool) {

	_, ok = stage.Backups[backup]

	return
}

func (stage *StageStruct) IsStagedBar_style_color(bar_style_color *Bar_style_color) (ok bool) {

	_, ok = stage.Bar_style_colors[bar_style_color]

	return
}

func (stage *StageStruct) IsStagedBarline(barline *Barline) (ok bool) {

	_, ok = stage.Barlines[barline]

	return
}

func (stage *StageStruct) IsStagedBarre(barre *Barre) (ok bool) {

	_, ok = stage.Barres[barre]

	return
}

func (stage *StageStruct) IsStagedBass(bass *Bass) (ok bool) {

	_, ok = stage.Basss[bass]

	return
}

func (stage *StageStruct) IsStagedBass_step(bass_step *Bass_step) (ok bool) {

	_, ok = stage.Bass_steps[bass_step]

	return
}

func (stage *StageStruct) IsStagedBeam(beam *Beam) (ok bool) {

	_, ok = stage.Beams[beam]

	return
}

func (stage *StageStruct) IsStagedBeat_repeat(beat_repeat *Beat_repeat) (ok bool) {

	_, ok = stage.Beat_repeats[beat_repeat]

	return
}

func (stage *StageStruct) IsStagedBeat_unit_tied(beat_unit_tied *Beat_unit_tied) (ok bool) {

	_, ok = stage.Beat_unit_tieds[beat_unit_tied]

	return
}

func (stage *StageStruct) IsStagedBeater(beater *Beater) (ok bool) {

	_, ok = stage.Beaters[beater]

	return
}

func (stage *StageStruct) IsStagedBend(bend *Bend) (ok bool) {

	_, ok = stage.Bends[bend]

	return
}

func (stage *StageStruct) IsStagedBookmark(bookmark *Bookmark) (ok bool) {

	_, ok = stage.Bookmarks[bookmark]

	return
}

func (stage *StageStruct) IsStagedBracket(bracket *Bracket) (ok bool) {

	_, ok = stage.Brackets[bracket]

	return
}

func (stage *StageStruct) IsStagedBreath_mark(breath_mark *Breath_mark) (ok bool) {

	_, ok = stage.Breath_marks[breath_mark]

	return
}

func (stage *StageStruct) IsStagedCaesura(caesura *Caesura) (ok bool) {

	_, ok = stage.Caesuras[caesura]

	return
}

func (stage *StageStruct) IsStagedCancel(cancel *Cancel) (ok bool) {

	_, ok = stage.Cancels[cancel]

	return
}

func (stage *StageStruct) IsStagedClef(clef *Clef) (ok bool) {

	_, ok = stage.Clefs[clef]

	return
}

func (stage *StageStruct) IsStagedCoda(coda *Coda) (ok bool) {

	_, ok = stage.Codas[coda]

	return
}

func (stage *StageStruct) IsStagedCredit(credit *Credit) (ok bool) {

	_, ok = stage.Credits[credit]

	return
}

func (stage *StageStruct) IsStagedDashes(dashes *Dashes) (ok bool) {

	_, ok = stage.Dashess[dashes]

	return
}

func (stage *StageStruct) IsStagedDefaults(defaults *Defaults) (ok bool) {

	_, ok = stage.Defaultss[defaults]

	return
}

func (stage *StageStruct) IsStagedDegree(degree *Degree) (ok bool) {

	_, ok = stage.Degrees[degree]

	return
}

func (stage *StageStruct) IsStagedDegree_alter(degree_alter *Degree_alter) (ok bool) {

	_, ok = stage.Degree_alters[degree_alter]

	return
}

func (stage *StageStruct) IsStagedDegree_type(degree_type *Degree_type) (ok bool) {

	_, ok = stage.Degree_types[degree_type]

	return
}

func (stage *StageStruct) IsStagedDegree_value(degree_value *Degree_value) (ok bool) {

	_, ok = stage.Degree_values[degree_value]

	return
}

func (stage *StageStruct) IsStagedDirection(direction *Direction) (ok bool) {

	_, ok = stage.Directions[direction]

	return
}

func (stage *StageStruct) IsStagedDirection_type(direction_type *Direction_type) (ok bool) {

	_, ok = stage.Direction_types[direction_type]

	return
}

func (stage *StageStruct) IsStagedDistance(distance *Distance) (ok bool) {

	_, ok = stage.Distances[distance]

	return
}

func (stage *StageStruct) IsStagedDouble(double *Double) (ok bool) {

	_, ok = stage.Doubles[double]

	return
}

func (stage *StageStruct) IsStagedDynamics(dynamics *Dynamics) (ok bool) {

	_, ok = stage.Dynamicss[dynamics]

	return
}

func (stage *StageStruct) IsStagedEffect(effect *Effect) (ok bool) {

	_, ok = stage.Effects[effect]

	return
}

func (stage *StageStruct) IsStagedElision(elision *Elision) (ok bool) {

	_, ok = stage.Elisions[elision]

	return
}

func (stage *StageStruct) IsStagedEmpty(empty *Empty) (ok bool) {

	_, ok = stage.Emptys[empty]

	return
}

func (stage *StageStruct) IsStagedEmpty_font(empty_font *Empty_font) (ok bool) {

	_, ok = stage.Empty_fonts[empty_font]

	return
}

func (stage *StageStruct) IsStagedEmpty_line(empty_line *Empty_line) (ok bool) {

	_, ok = stage.Empty_lines[empty_line]

	return
}

func (stage *StageStruct) IsStagedEmpty_placement(empty_placement *Empty_placement) (ok bool) {

	_, ok = stage.Empty_placements[empty_placement]

	return
}

func (stage *StageStruct) IsStagedEmpty_placement_smufl(empty_placement_smufl *Empty_placement_smufl) (ok bool) {

	_, ok = stage.Empty_placement_smufls[empty_placement_smufl]

	return
}

func (stage *StageStruct) IsStagedEmpty_print_object_style_align(empty_print_object_style_align *Empty_print_object_style_align) (ok bool) {

	_, ok = stage.Empty_print_object_style_aligns[empty_print_object_style_align]

	return
}

func (stage *StageStruct) IsStagedEmpty_print_style(empty_print_style *Empty_print_style) (ok bool) {

	_, ok = stage.Empty_print_styles[empty_print_style]

	return
}

func (stage *StageStruct) IsStagedEmpty_print_style_align(empty_print_style_align *Empty_print_style_align) (ok bool) {

	_, ok = stage.Empty_print_style_aligns[empty_print_style_align]

	return
}

func (stage *StageStruct) IsStagedEmpty_print_style_align_id(empty_print_style_align_id *Empty_print_style_align_id) (ok bool) {

	_, ok = stage.Empty_print_style_align_ids[empty_print_style_align_id]

	return
}

func (stage *StageStruct) IsStagedEmpty_trill_sound(empty_trill_sound *Empty_trill_sound) (ok bool) {

	_, ok = stage.Empty_trill_sounds[empty_trill_sound]

	return
}

func (stage *StageStruct) IsStagedEncoding(encoding *Encoding) (ok bool) {

	_, ok = stage.Encodings[encoding]

	return
}

func (stage *StageStruct) IsStagedEnding(ending *Ending) (ok bool) {

	_, ok = stage.Endings[ending]

	return
}

func (stage *StageStruct) IsStagedExtend(extend *Extend) (ok bool) {

	_, ok = stage.Extends[extend]

	return
}

func (stage *StageStruct) IsStagedFeature(feature *Feature) (ok bool) {

	_, ok = stage.Features[feature]

	return
}

func (stage *StageStruct) IsStagedFermata(fermata *Fermata) (ok bool) {

	_, ok = stage.Fermatas[fermata]

	return
}

func (stage *StageStruct) IsStagedFigure(figure *Figure) (ok bool) {

	_, ok = stage.Figures[figure]

	return
}

func (stage *StageStruct) IsStagedFigured_bass(figured_bass *Figured_bass) (ok bool) {

	_, ok = stage.Figured_basss[figured_bass]

	return
}

func (stage *StageStruct) IsStagedFingering(fingering *Fingering) (ok bool) {

	_, ok = stage.Fingerings[fingering]

	return
}

func (stage *StageStruct) IsStagedFirst_fret(first_fret *First_fret) (ok bool) {

	_, ok = stage.First_frets[first_fret]

	return
}

func (stage *StageStruct) IsStagedFoo(foo *Foo) (ok bool) {

	_, ok = stage.Foos[foo]

	return
}

func (stage *StageStruct) IsStagedFor_part(for_part *For_part) (ok bool) {

	_, ok = stage.For_parts[for_part]

	return
}

func (stage *StageStruct) IsStagedFormatted_symbol(formatted_symbol *Formatted_symbol) (ok bool) {

	_, ok = stage.Formatted_symbols[formatted_symbol]

	return
}

func (stage *StageStruct) IsStagedFormatted_symbol_id(formatted_symbol_id *Formatted_symbol_id) (ok bool) {

	_, ok = stage.Formatted_symbol_ids[formatted_symbol_id]

	return
}

func (stage *StageStruct) IsStagedForward(forward *Forward) (ok bool) {

	_, ok = stage.Forwards[forward]

	return
}

func (stage *StageStruct) IsStagedFrame(frame *Frame) (ok bool) {

	_, ok = stage.Frames[frame]

	return
}

func (stage *StageStruct) IsStagedFrame_note(frame_note *Frame_note) (ok bool) {

	_, ok = stage.Frame_notes[frame_note]

	return
}

func (stage *StageStruct) IsStagedFret(fret *Fret) (ok bool) {

	_, ok = stage.Frets[fret]

	return
}

func (stage *StageStruct) IsStagedGlass(glass *Glass) (ok bool) {

	_, ok = stage.Glasss[glass]

	return
}

func (stage *StageStruct) IsStagedGlissando(glissando *Glissando) (ok bool) {

	_, ok = stage.Glissandos[glissando]

	return
}

func (stage *StageStruct) IsStagedGlyph(glyph *Glyph) (ok bool) {

	_, ok = stage.Glyphs[glyph]

	return
}

func (stage *StageStruct) IsStagedGrace(grace *Grace) (ok bool) {

	_, ok = stage.Graces[grace]

	return
}

func (stage *StageStruct) IsStagedGroup_barline(group_barline *Group_barline) (ok bool) {

	_, ok = stage.Group_barlines[group_barline]

	return
}

func (stage *StageStruct) IsStagedGroup_symbol(group_symbol *Group_symbol) (ok bool) {

	_, ok = stage.Group_symbols[group_symbol]

	return
}

func (stage *StageStruct) IsStagedGrouping(grouping *Grouping) (ok bool) {

	_, ok = stage.Groupings[grouping]

	return
}

func (stage *StageStruct) IsStagedHammer_on_pull_off(hammer_on_pull_off *Hammer_on_pull_off) (ok bool) {

	_, ok = stage.Hammer_on_pull_offs[hammer_on_pull_off]

	return
}

func (stage *StageStruct) IsStagedHandbell(handbell *Handbell) (ok bool) {

	_, ok = stage.Handbells[handbell]

	return
}

func (stage *StageStruct) IsStagedHarmon_closed(harmon_closed *Harmon_closed) (ok bool) {

	_, ok = stage.Harmon_closeds[harmon_closed]

	return
}

func (stage *StageStruct) IsStagedHarmon_mute(harmon_mute *Harmon_mute) (ok bool) {

	_, ok = stage.Harmon_mutes[harmon_mute]

	return
}

func (stage *StageStruct) IsStagedHarmonic(harmonic *Harmonic) (ok bool) {

	_, ok = stage.Harmonics[harmonic]

	return
}

func (stage *StageStruct) IsStagedHarmony(harmony *Harmony) (ok bool) {

	_, ok = stage.Harmonys[harmony]

	return
}

func (stage *StageStruct) IsStagedHarmony_alter(harmony_alter *Harmony_alter) (ok bool) {

	_, ok = stage.Harmony_alters[harmony_alter]

	return
}

func (stage *StageStruct) IsStagedHarp_pedals(harp_pedals *Harp_pedals) (ok bool) {

	_, ok = stage.Harp_pedalss[harp_pedals]

	return
}

func (stage *StageStruct) IsStagedHeel_toe(heel_toe *Heel_toe) (ok bool) {

	_, ok = stage.Heel_toes[heel_toe]

	return
}

func (stage *StageStruct) IsStagedHole(hole *Hole) (ok bool) {

	_, ok = stage.Holes[hole]

	return
}

func (stage *StageStruct) IsStagedHole_closed(hole_closed *Hole_closed) (ok bool) {

	_, ok = stage.Hole_closeds[hole_closed]

	return
}

func (stage *StageStruct) IsStagedHorizontal_turn(horizontal_turn *Horizontal_turn) (ok bool) {

	_, ok = stage.Horizontal_turns[horizontal_turn]

	return
}

func (stage *StageStruct) IsStagedIdentification(identification *Identification) (ok bool) {

	_, ok = stage.Identifications[identification]

	return
}

func (stage *StageStruct) IsStagedImage(image *Image) (ok bool) {

	_, ok = stage.Images[image]

	return
}

func (stage *StageStruct) IsStagedInstrument(instrument *Instrument) (ok bool) {

	_, ok = stage.Instruments[instrument]

	return
}

func (stage *StageStruct) IsStagedInstrument_change(instrument_change *Instrument_change) (ok bool) {

	_, ok = stage.Instrument_changes[instrument_change]

	return
}

func (stage *StageStruct) IsStagedInstrument_link(instrument_link *Instrument_link) (ok bool) {

	_, ok = stage.Instrument_links[instrument_link]

	return
}

func (stage *StageStruct) IsStagedInterchangeable(interchangeable *Interchangeable) (ok bool) {

	_, ok = stage.Interchangeables[interchangeable]

	return
}

func (stage *StageStruct) IsStagedInversion(inversion *Inversion) (ok bool) {

	_, ok = stage.Inversions[inversion]

	return
}

func (stage *StageStruct) IsStagedKey(key *Key) (ok bool) {

	_, ok = stage.Keys[key]

	return
}

func (stage *StageStruct) IsStagedKey_accidental(key_accidental *Key_accidental) (ok bool) {

	_, ok = stage.Key_accidentals[key_accidental]

	return
}

func (stage *StageStruct) IsStagedKey_octave(key_octave *Key_octave) (ok bool) {

	_, ok = stage.Key_octaves[key_octave]

	return
}

func (stage *StageStruct) IsStagedKind(kind *Kind) (ok bool) {

	_, ok = stage.Kinds[kind]

	return
}

func (stage *StageStruct) IsStagedLevel(level *Level) (ok bool) {

	_, ok = stage.Levels[level]

	return
}

func (stage *StageStruct) IsStagedLine_detail(line_detail *Line_detail) (ok bool) {

	_, ok = stage.Line_details[line_detail]

	return
}

func (stage *StageStruct) IsStagedLine_width(line_width *Line_width) (ok bool) {

	_, ok = stage.Line_widths[line_width]

	return
}

func (stage *StageStruct) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

	return
}

func (stage *StageStruct) IsStagedListen(listen *Listen) (ok bool) {

	_, ok = stage.Listens[listen]

	return
}

func (stage *StageStruct) IsStagedListening(listening *Listening) (ok bool) {

	_, ok = stage.Listenings[listening]

	return
}

func (stage *StageStruct) IsStagedLyric(lyric *Lyric) (ok bool) {

	_, ok = stage.Lyrics[lyric]

	return
}

func (stage *StageStruct) IsStagedLyric_font(lyric_font *Lyric_font) (ok bool) {

	_, ok = stage.Lyric_fonts[lyric_font]

	return
}

func (stage *StageStruct) IsStagedLyric_language(lyric_language *Lyric_language) (ok bool) {

	_, ok = stage.Lyric_languages[lyric_language]

	return
}

func (stage *StageStruct) IsStagedMeasure_layout(measure_layout *Measure_layout) (ok bool) {

	_, ok = stage.Measure_layouts[measure_layout]

	return
}

func (stage *StageStruct) IsStagedMeasure_numbering(measure_numbering *Measure_numbering) (ok bool) {

	_, ok = stage.Measure_numberings[measure_numbering]

	return
}

func (stage *StageStruct) IsStagedMeasure_repeat(measure_repeat *Measure_repeat) (ok bool) {

	_, ok = stage.Measure_repeats[measure_repeat]

	return
}

func (stage *StageStruct) IsStagedMeasure_style(measure_style *Measure_style) (ok bool) {

	_, ok = stage.Measure_styles[measure_style]

	return
}

func (stage *StageStruct) IsStagedMembrane(membrane *Membrane) (ok bool) {

	_, ok = stage.Membranes[membrane]

	return
}

func (stage *StageStruct) IsStagedMetal(metal *Metal) (ok bool) {

	_, ok = stage.Metals[metal]

	return
}

func (stage *StageStruct) IsStagedMetronome(metronome *Metronome) (ok bool) {

	_, ok = stage.Metronomes[metronome]

	return
}

func (stage *StageStruct) IsStagedMetronome_beam(metronome_beam *Metronome_beam) (ok bool) {

	_, ok = stage.Metronome_beams[metronome_beam]

	return
}

func (stage *StageStruct) IsStagedMetronome_note(metronome_note *Metronome_note) (ok bool) {

	_, ok = stage.Metronome_notes[metronome_note]

	return
}

func (stage *StageStruct) IsStagedMetronome_tied(metronome_tied *Metronome_tied) (ok bool) {

	_, ok = stage.Metronome_tieds[metronome_tied]

	return
}

func (stage *StageStruct) IsStagedMetronome_tuplet(metronome_tuplet *Metronome_tuplet) (ok bool) {

	_, ok = stage.Metronome_tuplets[metronome_tuplet]

	return
}

func (stage *StageStruct) IsStagedMidi_device(midi_device *Midi_device) (ok bool) {

	_, ok = stage.Midi_devices[midi_device]

	return
}

func (stage *StageStruct) IsStagedMidi_instrument(midi_instrument *Midi_instrument) (ok bool) {

	_, ok = stage.Midi_instruments[midi_instrument]

	return
}

func (stage *StageStruct) IsStagedMiscellaneous(miscellaneous *Miscellaneous) (ok bool) {

	_, ok = stage.Miscellaneouss[miscellaneous]

	return
}

func (stage *StageStruct) IsStagedMiscellaneous_field(miscellaneous_field *Miscellaneous_field) (ok bool) {

	_, ok = stage.Miscellaneous_fields[miscellaneous_field]

	return
}

func (stage *StageStruct) IsStagedMordent(mordent *Mordent) (ok bool) {

	_, ok = stage.Mordents[mordent]

	return
}

func (stage *StageStruct) IsStagedMultiple_rest(multiple_rest *Multiple_rest) (ok bool) {

	_, ok = stage.Multiple_rests[multiple_rest]

	return
}

func (stage *StageStruct) IsStagedName_display(name_display *Name_display) (ok bool) {

	_, ok = stage.Name_displays[name_display]

	return
}

func (stage *StageStruct) IsStagedNon_arpeggiate(non_arpeggiate *Non_arpeggiate) (ok bool) {

	_, ok = stage.Non_arpeggiates[non_arpeggiate]

	return
}

func (stage *StageStruct) IsStagedNotations(notations *Notations) (ok bool) {

	_, ok = stage.Notationss[notations]

	return
}

func (stage *StageStruct) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *StageStruct) IsStagedNote_size(note_size *Note_size) (ok bool) {

	_, ok = stage.Note_sizes[note_size]

	return
}

func (stage *StageStruct) IsStagedNote_type(note_type *Note_type) (ok bool) {

	_, ok = stage.Note_types[note_type]

	return
}

func (stage *StageStruct) IsStagedNotehead(notehead *Notehead) (ok bool) {

	_, ok = stage.Noteheads[notehead]

	return
}

func (stage *StageStruct) IsStagedNotehead_text(notehead_text *Notehead_text) (ok bool) {

	_, ok = stage.Notehead_texts[notehead_text]

	return
}

func (stage *StageStruct) IsStagedNumeral(numeral *Numeral) (ok bool) {

	_, ok = stage.Numerals[numeral]

	return
}

func (stage *StageStruct) IsStagedNumeral_key(numeral_key *Numeral_key) (ok bool) {

	_, ok = stage.Numeral_keys[numeral_key]

	return
}

func (stage *StageStruct) IsStagedNumeral_root(numeral_root *Numeral_root) (ok bool) {

	_, ok = stage.Numeral_roots[numeral_root]

	return
}

func (stage *StageStruct) IsStagedOctave_shift(octave_shift *Octave_shift) (ok bool) {

	_, ok = stage.Octave_shifts[octave_shift]

	return
}

func (stage *StageStruct) IsStagedOffset(offset *Offset) (ok bool) {

	_, ok = stage.Offsets[offset]

	return
}

func (stage *StageStruct) IsStagedOpus(opus *Opus) (ok bool) {

	_, ok = stage.Opuss[opus]

	return
}

func (stage *StageStruct) IsStagedOrnaments(ornaments *Ornaments) (ok bool) {

	_, ok = stage.Ornamentss[ornaments]

	return
}

func (stage *StageStruct) IsStagedOther_appearance(other_appearance *Other_appearance) (ok bool) {

	_, ok = stage.Other_appearances[other_appearance]

	return
}

func (stage *StageStruct) IsStagedOther_listening(other_listening *Other_listening) (ok bool) {

	_, ok = stage.Other_listenings[other_listening]

	return
}

func (stage *StageStruct) IsStagedOther_notation(other_notation *Other_notation) (ok bool) {

	_, ok = stage.Other_notations[other_notation]

	return
}

func (stage *StageStruct) IsStagedOther_play(other_play *Other_play) (ok bool) {

	_, ok = stage.Other_plays[other_play]

	return
}

func (stage *StageStruct) IsStagedPage_layout(page_layout *Page_layout) (ok bool) {

	_, ok = stage.Page_layouts[page_layout]

	return
}

func (stage *StageStruct) IsStagedPage_margins(page_margins *Page_margins) (ok bool) {

	_, ok = stage.Page_marginss[page_margins]

	return
}

func (stage *StageStruct) IsStagedPart_clef(part_clef *Part_clef) (ok bool) {

	_, ok = stage.Part_clefs[part_clef]

	return
}

func (stage *StageStruct) IsStagedPart_group(part_group *Part_group) (ok bool) {

	_, ok = stage.Part_groups[part_group]

	return
}

func (stage *StageStruct) IsStagedPart_link(part_link *Part_link) (ok bool) {

	_, ok = stage.Part_links[part_link]

	return
}

func (stage *StageStruct) IsStagedPart_list(part_list *Part_list) (ok bool) {

	_, ok = stage.Part_lists[part_list]

	return
}

func (stage *StageStruct) IsStagedPart_symbol(part_symbol *Part_symbol) (ok bool) {

	_, ok = stage.Part_symbols[part_symbol]

	return
}

func (stage *StageStruct) IsStagedPart_transpose(part_transpose *Part_transpose) (ok bool) {

	_, ok = stage.Part_transposes[part_transpose]

	return
}

func (stage *StageStruct) IsStagedPedal(pedal *Pedal) (ok bool) {

	_, ok = stage.Pedals[pedal]

	return
}

func (stage *StageStruct) IsStagedPedal_tuning(pedal_tuning *Pedal_tuning) (ok bool) {

	_, ok = stage.Pedal_tunings[pedal_tuning]

	return
}

func (stage *StageStruct) IsStagedPercussion(percussion *Percussion) (ok bool) {

	_, ok = stage.Percussions[percussion]

	return
}

func (stage *StageStruct) IsStagedPitch(pitch *Pitch) (ok bool) {

	_, ok = stage.Pitchs[pitch]

	return
}

func (stage *StageStruct) IsStagedPitched(pitched *Pitched) (ok bool) {

	_, ok = stage.Pitcheds[pitched]

	return
}

func (stage *StageStruct) IsStagedPlay(play *Play) (ok bool) {

	_, ok = stage.Plays[play]

	return
}

func (stage *StageStruct) IsStagedPlayer(player *Player) (ok bool) {

	_, ok = stage.Players[player]

	return
}

func (stage *StageStruct) IsStagedPrincipal_voice(principal_voice *Principal_voice) (ok bool) {

	_, ok = stage.Principal_voices[principal_voice]

	return
}

func (stage *StageStruct) IsStagedPrint(print *Print) (ok bool) {

	_, ok = stage.Prints[print]

	return
}

func (stage *StageStruct) IsStagedRelease(release *Release) (ok bool) {

	_, ok = stage.Releases[release]

	return
}

func (stage *StageStruct) IsStagedRepeat(repeat *Repeat) (ok bool) {

	_, ok = stage.Repeats[repeat]

	return
}

func (stage *StageStruct) IsStagedRest(rest *Rest) (ok bool) {

	_, ok = stage.Rests[rest]

	return
}

func (stage *StageStruct) IsStagedRoot(root *Root) (ok bool) {

	_, ok = stage.Roots[root]

	return
}

func (stage *StageStruct) IsStagedRoot_step(root_step *Root_step) (ok bool) {

	_, ok = stage.Root_steps[root_step]

	return
}

func (stage *StageStruct) IsStagedScaling(scaling *Scaling) (ok bool) {

	_, ok = stage.Scalings[scaling]

	return
}

func (stage *StageStruct) IsStagedScordatura(scordatura *Scordatura) (ok bool) {

	_, ok = stage.Scordaturas[scordatura]

	return
}

func (stage *StageStruct) IsStagedScore_instrument(score_instrument *Score_instrument) (ok bool) {

	_, ok = stage.Score_instruments[score_instrument]

	return
}

func (stage *StageStruct) IsStagedScore_part(score_part *Score_part) (ok bool) {

	_, ok = stage.Score_parts[score_part]

	return
}

func (stage *StageStruct) IsStagedScore_partwise(score_partwise *Score_partwise) (ok bool) {

	_, ok = stage.Score_partwises[score_partwise]

	return
}

func (stage *StageStruct) IsStagedScore_timewise(score_timewise *Score_timewise) (ok bool) {

	_, ok = stage.Score_timewises[score_timewise]

	return
}

func (stage *StageStruct) IsStagedSegno(segno *Segno) (ok bool) {

	_, ok = stage.Segnos[segno]

	return
}

func (stage *StageStruct) IsStagedSlash(slash *Slash) (ok bool) {

	_, ok = stage.Slashs[slash]

	return
}

func (stage *StageStruct) IsStagedSlide(slide *Slide) (ok bool) {

	_, ok = stage.Slides[slide]

	return
}

func (stage *StageStruct) IsStagedSlur(slur *Slur) (ok bool) {

	_, ok = stage.Slurs[slur]

	return
}

func (stage *StageStruct) IsStagedSound(sound *Sound) (ok bool) {

	_, ok = stage.Sounds[sound]

	return
}

func (stage *StageStruct) IsStagedStaff_details(staff_details *Staff_details) (ok bool) {

	_, ok = stage.Staff_detailss[staff_details]

	return
}

func (stage *StageStruct) IsStagedStaff_divide(staff_divide *Staff_divide) (ok bool) {

	_, ok = stage.Staff_divides[staff_divide]

	return
}

func (stage *StageStruct) IsStagedStaff_layout(staff_layout *Staff_layout) (ok bool) {

	_, ok = stage.Staff_layouts[staff_layout]

	return
}

func (stage *StageStruct) IsStagedStaff_size(staff_size *Staff_size) (ok bool) {

	_, ok = stage.Staff_sizes[staff_size]

	return
}

func (stage *StageStruct) IsStagedStaff_tuning(staff_tuning *Staff_tuning) (ok bool) {

	_, ok = stage.Staff_tunings[staff_tuning]

	return
}

func (stage *StageStruct) IsStagedStem(stem *Stem) (ok bool) {

	_, ok = stage.Stems[stem]

	return
}

func (stage *StageStruct) IsStagedStick(stick *Stick) (ok bool) {

	_, ok = stage.Sticks[stick]

	return
}

func (stage *StageStruct) IsStagedString_mute(string_mute *String_mute) (ok bool) {

	_, ok = stage.String_mutes[string_mute]

	return
}

func (stage *StageStruct) IsStagedStrong_accent(strong_accent *Strong_accent) (ok bool) {

	_, ok = stage.Strong_accents[strong_accent]

	return
}

func (stage *StageStruct) IsStagedSupports(supports *Supports) (ok bool) {

	_, ok = stage.Supportss[supports]

	return
}

func (stage *StageStruct) IsStagedSwing(swing *Swing) (ok bool) {

	_, ok = stage.Swings[swing]

	return
}

func (stage *StageStruct) IsStagedSync(sync *Sync) (ok bool) {

	_, ok = stage.Syncs[sync]

	return
}

func (stage *StageStruct) IsStagedSystem_dividers(system_dividers *System_dividers) (ok bool) {

	_, ok = stage.System_dividerss[system_dividers]

	return
}

func (stage *StageStruct) IsStagedSystem_layout(system_layout *System_layout) (ok bool) {

	_, ok = stage.System_layouts[system_layout]

	return
}

func (stage *StageStruct) IsStagedSystem_margins(system_margins *System_margins) (ok bool) {

	_, ok = stage.System_marginss[system_margins]

	return
}

func (stage *StageStruct) IsStagedTap(tap *Tap) (ok bool) {

	_, ok = stage.Taps[tap]

	return
}

func (stage *StageStruct) IsStagedTechnical(technical *Technical) (ok bool) {

	_, ok = stage.Technicals[technical]

	return
}

func (stage *StageStruct) IsStagedText_element_data(text_element_data *Text_element_data) (ok bool) {

	_, ok = stage.Text_element_datas[text_element_data]

	return
}

func (stage *StageStruct) IsStagedTie(tie *Tie) (ok bool) {

	_, ok = stage.Ties[tie]

	return
}

func (stage *StageStruct) IsStagedTied(tied *Tied) (ok bool) {

	_, ok = stage.Tieds[tied]

	return
}

func (stage *StageStruct) IsStagedTime(time *Time) (ok bool) {

	_, ok = stage.Times[time]

	return
}

func (stage *StageStruct) IsStagedTime_modification(time_modification *Time_modification) (ok bool) {

	_, ok = stage.Time_modifications[time_modification]

	return
}

func (stage *StageStruct) IsStagedTimpani(timpani *Timpani) (ok bool) {

	_, ok = stage.Timpanis[timpani]

	return
}

func (stage *StageStruct) IsStagedTranspose(transpose *Transpose) (ok bool) {

	_, ok = stage.Transposes[transpose]

	return
}

func (stage *StageStruct) IsStagedTremolo(tremolo *Tremolo) (ok bool) {

	_, ok = stage.Tremolos[tremolo]

	return
}

func (stage *StageStruct) IsStagedTuplet(tuplet *Tuplet) (ok bool) {

	_, ok = stage.Tuplets[tuplet]

	return
}

func (stage *StageStruct) IsStagedTuplet_dot(tuplet_dot *Tuplet_dot) (ok bool) {

	_, ok = stage.Tuplet_dots[tuplet_dot]

	return
}

func (stage *StageStruct) IsStagedTuplet_number(tuplet_number *Tuplet_number) (ok bool) {

	_, ok = stage.Tuplet_numbers[tuplet_number]

	return
}

func (stage *StageStruct) IsStagedTuplet_portion(tuplet_portion *Tuplet_portion) (ok bool) {

	_, ok = stage.Tuplet_portions[tuplet_portion]

	return
}

func (stage *StageStruct) IsStagedTuplet_type(tuplet_type *Tuplet_type) (ok bool) {

	_, ok = stage.Tuplet_types[tuplet_type]

	return
}

func (stage *StageStruct) IsStagedTyped_text(typed_text *Typed_text) (ok bool) {

	_, ok = stage.Typed_texts[typed_text]

	return
}

func (stage *StageStruct) IsStagedUnpitched(unpitched *Unpitched) (ok bool) {

	_, ok = stage.Unpitcheds[unpitched]

	return
}

func (stage *StageStruct) IsStagedVirtual_instrument(virtual_instrument *Virtual_instrument) (ok bool) {

	_, ok = stage.Virtual_instruments[virtual_instrument]

	return
}

func (stage *StageStruct) IsStagedWait(wait *Wait) (ok bool) {

	_, ok = stage.Waits[wait]

	return
}

func (stage *StageStruct) IsStagedWavy_line(wavy_line *Wavy_line) (ok bool) {

	_, ok = stage.Wavy_lines[wavy_line]

	return
}

func (stage *StageStruct) IsStagedWedge(wedge *Wedge) (ok bool) {

	_, ok = stage.Wedges[wedge]

	return
}

func (stage *StageStruct) IsStagedWood(wood *Wood) (ok bool) {

	_, ok = stage.Woods[wood]

	return
}

func (stage *StageStruct) IsStagedWork(work *Work) (ok bool) {

	_, ok = stage.Works[work]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Accidental:
		stage.StageBranchAccidental(target)

	case *Accidental_mark:
		stage.StageBranchAccidental_mark(target)

	case *Accidental_text:
		stage.StageBranchAccidental_text(target)

	case *Accord:
		stage.StageBranchAccord(target)

	case *Accordion_registration:
		stage.StageBranchAccordion_registration(target)

	case *AnyType:
		stage.StageBranchAnyType(target)

	case *Appearance:
		stage.StageBranchAppearance(target)

	case *Arpeggiate:
		stage.StageBranchArpeggiate(target)

	case *Arrow:
		stage.StageBranchArrow(target)

	case *Articulations:
		stage.StageBranchArticulations(target)

	case *Assess:
		stage.StageBranchAssess(target)

	case *Attributes:
		stage.StageBranchAttributes(target)

	case *Backup:
		stage.StageBranchBackup(target)

	case *Bar_style_color:
		stage.StageBranchBar_style_color(target)

	case *Barline:
		stage.StageBranchBarline(target)

	case *Barre:
		stage.StageBranchBarre(target)

	case *Bass:
		stage.StageBranchBass(target)

	case *Bass_step:
		stage.StageBranchBass_step(target)

	case *Beam:
		stage.StageBranchBeam(target)

	case *Beat_repeat:
		stage.StageBranchBeat_repeat(target)

	case *Beat_unit_tied:
		stage.StageBranchBeat_unit_tied(target)

	case *Beater:
		stage.StageBranchBeater(target)

	case *Bend:
		stage.StageBranchBend(target)

	case *Bookmark:
		stage.StageBranchBookmark(target)

	case *Bracket:
		stage.StageBranchBracket(target)

	case *Breath_mark:
		stage.StageBranchBreath_mark(target)

	case *Caesura:
		stage.StageBranchCaesura(target)

	case *Cancel:
		stage.StageBranchCancel(target)

	case *Clef:
		stage.StageBranchClef(target)

	case *Coda:
		stage.StageBranchCoda(target)

	case *Credit:
		stage.StageBranchCredit(target)

	case *Dashes:
		stage.StageBranchDashes(target)

	case *Defaults:
		stage.StageBranchDefaults(target)

	case *Degree:
		stage.StageBranchDegree(target)

	case *Degree_alter:
		stage.StageBranchDegree_alter(target)

	case *Degree_type:
		stage.StageBranchDegree_type(target)

	case *Degree_value:
		stage.StageBranchDegree_value(target)

	case *Direction:
		stage.StageBranchDirection(target)

	case *Direction_type:
		stage.StageBranchDirection_type(target)

	case *Distance:
		stage.StageBranchDistance(target)

	case *Double:
		stage.StageBranchDouble(target)

	case *Dynamics:
		stage.StageBranchDynamics(target)

	case *Effect:
		stage.StageBranchEffect(target)

	case *Elision:
		stage.StageBranchElision(target)

	case *Empty:
		stage.StageBranchEmpty(target)

	case *Empty_font:
		stage.StageBranchEmpty_font(target)

	case *Empty_line:
		stage.StageBranchEmpty_line(target)

	case *Empty_placement:
		stage.StageBranchEmpty_placement(target)

	case *Empty_placement_smufl:
		stage.StageBranchEmpty_placement_smufl(target)

	case *Empty_print_object_style_align:
		stage.StageBranchEmpty_print_object_style_align(target)

	case *Empty_print_style:
		stage.StageBranchEmpty_print_style(target)

	case *Empty_print_style_align:
		stage.StageBranchEmpty_print_style_align(target)

	case *Empty_print_style_align_id:
		stage.StageBranchEmpty_print_style_align_id(target)

	case *Empty_trill_sound:
		stage.StageBranchEmpty_trill_sound(target)

	case *Encoding:
		stage.StageBranchEncoding(target)

	case *Ending:
		stage.StageBranchEnding(target)

	case *Extend:
		stage.StageBranchExtend(target)

	case *Feature:
		stage.StageBranchFeature(target)

	case *Fermata:
		stage.StageBranchFermata(target)

	case *Figure:
		stage.StageBranchFigure(target)

	case *Figured_bass:
		stage.StageBranchFigured_bass(target)

	case *Fingering:
		stage.StageBranchFingering(target)

	case *First_fret:
		stage.StageBranchFirst_fret(target)

	case *Foo:
		stage.StageBranchFoo(target)

	case *For_part:
		stage.StageBranchFor_part(target)

	case *Formatted_symbol:
		stage.StageBranchFormatted_symbol(target)

	case *Formatted_symbol_id:
		stage.StageBranchFormatted_symbol_id(target)

	case *Forward:
		stage.StageBranchForward(target)

	case *Frame:
		stage.StageBranchFrame(target)

	case *Frame_note:
		stage.StageBranchFrame_note(target)

	case *Fret:
		stage.StageBranchFret(target)

	case *Glass:
		stage.StageBranchGlass(target)

	case *Glissando:
		stage.StageBranchGlissando(target)

	case *Glyph:
		stage.StageBranchGlyph(target)

	case *Grace:
		stage.StageBranchGrace(target)

	case *Group_barline:
		stage.StageBranchGroup_barline(target)

	case *Group_symbol:
		stage.StageBranchGroup_symbol(target)

	case *Grouping:
		stage.StageBranchGrouping(target)

	case *Hammer_on_pull_off:
		stage.StageBranchHammer_on_pull_off(target)

	case *Handbell:
		stage.StageBranchHandbell(target)

	case *Harmon_closed:
		stage.StageBranchHarmon_closed(target)

	case *Harmon_mute:
		stage.StageBranchHarmon_mute(target)

	case *Harmonic:
		stage.StageBranchHarmonic(target)

	case *Harmony:
		stage.StageBranchHarmony(target)

	case *Harmony_alter:
		stage.StageBranchHarmony_alter(target)

	case *Harp_pedals:
		stage.StageBranchHarp_pedals(target)

	case *Heel_toe:
		stage.StageBranchHeel_toe(target)

	case *Hole:
		stage.StageBranchHole(target)

	case *Hole_closed:
		stage.StageBranchHole_closed(target)

	case *Horizontal_turn:
		stage.StageBranchHorizontal_turn(target)

	case *Identification:
		stage.StageBranchIdentification(target)

	case *Image:
		stage.StageBranchImage(target)

	case *Instrument:
		stage.StageBranchInstrument(target)

	case *Instrument_change:
		stage.StageBranchInstrument_change(target)

	case *Instrument_link:
		stage.StageBranchInstrument_link(target)

	case *Interchangeable:
		stage.StageBranchInterchangeable(target)

	case *Inversion:
		stage.StageBranchInversion(target)

	case *Key:
		stage.StageBranchKey(target)

	case *Key_accidental:
		stage.StageBranchKey_accidental(target)

	case *Key_octave:
		stage.StageBranchKey_octave(target)

	case *Kind:
		stage.StageBranchKind(target)

	case *Level:
		stage.StageBranchLevel(target)

	case *Line_detail:
		stage.StageBranchLine_detail(target)

	case *Line_width:
		stage.StageBranchLine_width(target)

	case *Link:
		stage.StageBranchLink(target)

	case *Listen:
		stage.StageBranchListen(target)

	case *Listening:
		stage.StageBranchListening(target)

	case *Lyric:
		stage.StageBranchLyric(target)

	case *Lyric_font:
		stage.StageBranchLyric_font(target)

	case *Lyric_language:
		stage.StageBranchLyric_language(target)

	case *Measure_layout:
		stage.StageBranchMeasure_layout(target)

	case *Measure_numbering:
		stage.StageBranchMeasure_numbering(target)

	case *Measure_repeat:
		stage.StageBranchMeasure_repeat(target)

	case *Measure_style:
		stage.StageBranchMeasure_style(target)

	case *Membrane:
		stage.StageBranchMembrane(target)

	case *Metal:
		stage.StageBranchMetal(target)

	case *Metronome:
		stage.StageBranchMetronome(target)

	case *Metronome_beam:
		stage.StageBranchMetronome_beam(target)

	case *Metronome_note:
		stage.StageBranchMetronome_note(target)

	case *Metronome_tied:
		stage.StageBranchMetronome_tied(target)

	case *Metronome_tuplet:
		stage.StageBranchMetronome_tuplet(target)

	case *Midi_device:
		stage.StageBranchMidi_device(target)

	case *Midi_instrument:
		stage.StageBranchMidi_instrument(target)

	case *Miscellaneous:
		stage.StageBranchMiscellaneous(target)

	case *Miscellaneous_field:
		stage.StageBranchMiscellaneous_field(target)

	case *Mordent:
		stage.StageBranchMordent(target)

	case *Multiple_rest:
		stage.StageBranchMultiple_rest(target)

	case *Name_display:
		stage.StageBranchName_display(target)

	case *Non_arpeggiate:
		stage.StageBranchNon_arpeggiate(target)

	case *Notations:
		stage.StageBranchNotations(target)

	case *Note:
		stage.StageBranchNote(target)

	case *Note_size:
		stage.StageBranchNote_size(target)

	case *Note_type:
		stage.StageBranchNote_type(target)

	case *Notehead:
		stage.StageBranchNotehead(target)

	case *Notehead_text:
		stage.StageBranchNotehead_text(target)

	case *Numeral:
		stage.StageBranchNumeral(target)

	case *Numeral_key:
		stage.StageBranchNumeral_key(target)

	case *Numeral_root:
		stage.StageBranchNumeral_root(target)

	case *Octave_shift:
		stage.StageBranchOctave_shift(target)

	case *Offset:
		stage.StageBranchOffset(target)

	case *Opus:
		stage.StageBranchOpus(target)

	case *Ornaments:
		stage.StageBranchOrnaments(target)

	case *Other_appearance:
		stage.StageBranchOther_appearance(target)

	case *Other_listening:
		stage.StageBranchOther_listening(target)

	case *Other_notation:
		stage.StageBranchOther_notation(target)

	case *Other_play:
		stage.StageBranchOther_play(target)

	case *Page_layout:
		stage.StageBranchPage_layout(target)

	case *Page_margins:
		stage.StageBranchPage_margins(target)

	case *Part_clef:
		stage.StageBranchPart_clef(target)

	case *Part_group:
		stage.StageBranchPart_group(target)

	case *Part_link:
		stage.StageBranchPart_link(target)

	case *Part_list:
		stage.StageBranchPart_list(target)

	case *Part_symbol:
		stage.StageBranchPart_symbol(target)

	case *Part_transpose:
		stage.StageBranchPart_transpose(target)

	case *Pedal:
		stage.StageBranchPedal(target)

	case *Pedal_tuning:
		stage.StageBranchPedal_tuning(target)

	case *Percussion:
		stage.StageBranchPercussion(target)

	case *Pitch:
		stage.StageBranchPitch(target)

	case *Pitched:
		stage.StageBranchPitched(target)

	case *Play:
		stage.StageBranchPlay(target)

	case *Player:
		stage.StageBranchPlayer(target)

	case *Principal_voice:
		stage.StageBranchPrincipal_voice(target)

	case *Print:
		stage.StageBranchPrint(target)

	case *Release:
		stage.StageBranchRelease(target)

	case *Repeat:
		stage.StageBranchRepeat(target)

	case *Rest:
		stage.StageBranchRest(target)

	case *Root:
		stage.StageBranchRoot(target)

	case *Root_step:
		stage.StageBranchRoot_step(target)

	case *Scaling:
		stage.StageBranchScaling(target)

	case *Scordatura:
		stage.StageBranchScordatura(target)

	case *Score_instrument:
		stage.StageBranchScore_instrument(target)

	case *Score_part:
		stage.StageBranchScore_part(target)

	case *Score_partwise:
		stage.StageBranchScore_partwise(target)

	case *Score_timewise:
		stage.StageBranchScore_timewise(target)

	case *Segno:
		stage.StageBranchSegno(target)

	case *Slash:
		stage.StageBranchSlash(target)

	case *Slide:
		stage.StageBranchSlide(target)

	case *Slur:
		stage.StageBranchSlur(target)

	case *Sound:
		stage.StageBranchSound(target)

	case *Staff_details:
		stage.StageBranchStaff_details(target)

	case *Staff_divide:
		stage.StageBranchStaff_divide(target)

	case *Staff_layout:
		stage.StageBranchStaff_layout(target)

	case *Staff_size:
		stage.StageBranchStaff_size(target)

	case *Staff_tuning:
		stage.StageBranchStaff_tuning(target)

	case *Stem:
		stage.StageBranchStem(target)

	case *Stick:
		stage.StageBranchStick(target)

	case *String_mute:
		stage.StageBranchString_mute(target)

	case *Strong_accent:
		stage.StageBranchStrong_accent(target)

	case *Supports:
		stage.StageBranchSupports(target)

	case *Swing:
		stage.StageBranchSwing(target)

	case *Sync:
		stage.StageBranchSync(target)

	case *System_dividers:
		stage.StageBranchSystem_dividers(target)

	case *System_layout:
		stage.StageBranchSystem_layout(target)

	case *System_margins:
		stage.StageBranchSystem_margins(target)

	case *Tap:
		stage.StageBranchTap(target)

	case *Technical:
		stage.StageBranchTechnical(target)

	case *Text_element_data:
		stage.StageBranchText_element_data(target)

	case *Tie:
		stage.StageBranchTie(target)

	case *Tied:
		stage.StageBranchTied(target)

	case *Time:
		stage.StageBranchTime(target)

	case *Time_modification:
		stage.StageBranchTime_modification(target)

	case *Timpani:
		stage.StageBranchTimpani(target)

	case *Transpose:
		stage.StageBranchTranspose(target)

	case *Tremolo:
		stage.StageBranchTremolo(target)

	case *Tuplet:
		stage.StageBranchTuplet(target)

	case *Tuplet_dot:
		stage.StageBranchTuplet_dot(target)

	case *Tuplet_number:
		stage.StageBranchTuplet_number(target)

	case *Tuplet_portion:
		stage.StageBranchTuplet_portion(target)

	case *Tuplet_type:
		stage.StageBranchTuplet_type(target)

	case *Typed_text:
		stage.StageBranchTyped_text(target)

	case *Unpitched:
		stage.StageBranchUnpitched(target)

	case *Virtual_instrument:
		stage.StageBranchVirtual_instrument(target)

	case *Wait:
		stage.StageBranchWait(target)

	case *Wavy_line:
		stage.StageBranchWavy_line(target)

	case *Wedge:
		stage.StageBranchWedge(target)

	case *Wood:
		stage.StageBranchWood(target)

	case *Work:
		stage.StageBranchWork(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchAccidental(accidental *Accidental) {

	// check if instance is already staged
	if IsStaged(stage, accidental) {
		return
	}

	accidental.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAccidental_mark(accidental_mark *Accidental_mark) {

	// check if instance is already staged
	if IsStaged(stage, accidental_mark) {
		return
	}

	accidental_mark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAccidental_text(accidental_text *Accidental_text) {

	// check if instance is already staged
	if IsStaged(stage, accidental_text) {
		return
	}

	accidental_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAccord(accord *Accord) {

	// check if instance is already staged
	if IsStaged(stage, accord) {
		return
	}

	accord.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAccordion_registration(accordion_registration *Accordion_registration) {

	// check if instance is already staged
	if IsStaged(stage, accordion_registration) {
		return
	}

	accordion_registration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if accordion_registration.Accordion_high != nil {
		StageBranch(stage, accordion_registration.Accordion_high)
	}
	if accordion_registration.Accordion_low != nil {
		StageBranch(stage, accordion_registration.Accordion_low)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAnyType(anytype *AnyType) {

	// check if instance is already staged
	if IsStaged(stage, anytype) {
		return
	}

	anytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAppearance(appearance *Appearance) {

	// check if instance is already staged
	if IsStaged(stage, appearance) {
		return
	}

	appearance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_width := range appearance.Line_width {
		StageBranch(stage, _line_width)
	}
	for _, _note_size := range appearance.Note_size {
		StageBranch(stage, _note_size)
	}
	for _, _distance := range appearance.Distance {
		StageBranch(stage, _distance)
	}
	for _, _glyph := range appearance.Glyph {
		StageBranch(stage, _glyph)
	}
	for _, _other_appearance := range appearance.Other_appearance {
		StageBranch(stage, _other_appearance)
	}

}

func (stage *StageStruct) StageBranchArpeggiate(arpeggiate *Arpeggiate) {

	// check if instance is already staged
	if IsStaged(stage, arpeggiate) {
		return
	}

	arpeggiate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchArrow(arrow *Arrow) {

	// check if instance is already staged
	if IsStaged(stage, arrow) {
		return
	}

	arrow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchArticulations(articulations *Articulations) {

	// check if instance is already staged
	if IsStaged(stage, articulations) {
		return
	}

	articulations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if articulations.Accent != nil {
		StageBranch(stage, articulations.Accent)
	}
	if articulations.Strong_accent != nil {
		StageBranch(stage, articulations.Strong_accent)
	}
	if articulations.Staccato != nil {
		StageBranch(stage, articulations.Staccato)
	}
	if articulations.Tenuto != nil {
		StageBranch(stage, articulations.Tenuto)
	}
	if articulations.Detached_legato != nil {
		StageBranch(stage, articulations.Detached_legato)
	}
	if articulations.Staccatissimo != nil {
		StageBranch(stage, articulations.Staccatissimo)
	}
	if articulations.Spiccato != nil {
		StageBranch(stage, articulations.Spiccato)
	}
	if articulations.Scoop != nil {
		StageBranch(stage, articulations.Scoop)
	}
	if articulations.Plop != nil {
		StageBranch(stage, articulations.Plop)
	}
	if articulations.Doit != nil {
		StageBranch(stage, articulations.Doit)
	}
	if articulations.Falloff != nil {
		StageBranch(stage, articulations.Falloff)
	}
	if articulations.Breath_mark != nil {
		StageBranch(stage, articulations.Breath_mark)
	}
	if articulations.Caesura != nil {
		StageBranch(stage, articulations.Caesura)
	}
	if articulations.Stress != nil {
		StageBranch(stage, articulations.Stress)
	}
	if articulations.Unstress != nil {
		StageBranch(stage, articulations.Unstress)
	}
	if articulations.Soft_accent != nil {
		StageBranch(stage, articulations.Soft_accent)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAssess(assess *Assess) {

	// check if instance is already staged
	if IsStaged(stage, assess) {
		return
	}

	assess.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchAttributes(attributes *Attributes) {

	// check if instance is already staged
	if IsStaged(stage, attributes) {
		return
	}

	attributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributes.Part_symbol != nil {
		StageBranch(stage, attributes.Part_symbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributes.Key {
		StageBranch(stage, _key)
	}
	for _, _clef := range attributes.Clef {
		StageBranch(stage, _clef)
	}
	for _, _staff_details := range attributes.Staff_details {
		StageBranch(stage, _staff_details)
	}
	for _, _measure_style := range attributes.Measure_style {
		StageBranch(stage, _measure_style)
	}
	for _, _transpose := range attributes.Transpose {
		StageBranch(stage, _transpose)
	}
	for _, _for_part := range attributes.For_part {
		StageBranch(stage, _for_part)
	}

}

func (stage *StageStruct) StageBranchBackup(backup *Backup) {

	// check if instance is already staged
	if IsStaged(stage, backup) {
		return
	}

	backup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBar_style_color(bar_style_color *Bar_style_color) {

	// check if instance is already staged
	if IsStaged(stage, bar_style_color) {
		return
	}

	bar_style_color.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBarline(barline *Barline) {

	// check if instance is already staged
	if IsStaged(stage, barline) {
		return
	}

	barline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if barline.Bar_style != nil {
		StageBranch(stage, barline.Bar_style)
	}
	if barline.Wavy_line != nil {
		StageBranch(stage, barline.Wavy_line)
	}
	if barline.Fermata != nil {
		StageBranch(stage, barline.Fermata)
	}
	if barline.Ending != nil {
		StageBranch(stage, barline.Ending)
	}
	if barline.Repeat != nil {
		StageBranch(stage, barline.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBarre(barre *Barre) {

	// check if instance is already staged
	if IsStaged(stage, barre) {
		return
	}

	barre.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBass(bass *Bass) {

	// check if instance is already staged
	if IsStaged(stage, bass) {
		return
	}

	bass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bass.Bass_step != nil {
		StageBranch(stage, bass.Bass_step)
	}
	if bass.Bass_alter != nil {
		StageBranch(stage, bass.Bass_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBass_step(bass_step *Bass_step) {

	// check if instance is already staged
	if IsStaged(stage, bass_step) {
		return
	}

	bass_step.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBeam(beam *Beam) {

	// check if instance is already staged
	if IsStaged(stage, beam) {
		return
	}

	beam.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBeat_repeat(beat_repeat *Beat_repeat) {

	// check if instance is already staged
	if IsStaged(stage, beat_repeat) {
		return
	}

	beat_repeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBeat_unit_tied(beat_unit_tied *Beat_unit_tied) {

	// check if instance is already staged
	if IsStaged(stage, beat_unit_tied) {
		return
	}

	beat_unit_tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBeater(beater *Beater) {

	// check if instance is already staged
	if IsStaged(stage, beater) {
		return
	}

	beater.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBend(bend *Bend) {

	// check if instance is already staged
	if IsStaged(stage, bend) {
		return
	}

	bend.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bend.Pre_bend != nil {
		StageBranch(stage, bend.Pre_bend)
	}
	if bend.Release != nil {
		StageBranch(stage, bend.Release)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBookmark(bookmark *Bookmark) {

	// check if instance is already staged
	if IsStaged(stage, bookmark) {
		return
	}

	bookmark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBracket(bracket *Bracket) {

	// check if instance is already staged
	if IsStaged(stage, bracket) {
		return
	}

	bracket.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchBreath_mark(breath_mark *Breath_mark) {

	// check if instance is already staged
	if IsStaged(stage, breath_mark) {
		return
	}

	breath_mark.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCaesura(caesura *Caesura) {

	// check if instance is already staged
	if IsStaged(stage, caesura) {
		return
	}

	caesura.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCancel(cancel *Cancel) {

	// check if instance is already staged
	if IsStaged(stage, cancel) {
		return
	}

	cancel.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchClef(clef *Clef) {

	// check if instance is already staged
	if IsStaged(stage, clef) {
		return
	}

	clef.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCoda(coda *Coda) {

	// check if instance is already staged
	if IsStaged(stage, coda) {
		return
	}

	coda.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if IsStaged(stage, credit) {
		return
	}

	credit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if credit.Credit_image != nil {
		StageBranch(stage, credit.Credit_image)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		StageBranch(stage, _link)
	}
	for _, _bookmark := range credit.Bookmark {
		StageBranch(stage, _bookmark)
	}

}

func (stage *StageStruct) StageBranchDashes(dashes *Dashes) {

	// check if instance is already staged
	if IsStaged(stage, dashes) {
		return
	}

	dashes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDefaults(defaults *Defaults) {

	// check if instance is already staged
	if IsStaged(stage, defaults) {
		return
	}

	defaults.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaults.Scaling != nil {
		StageBranch(stage, defaults.Scaling)
	}
	if defaults.Concert_score != nil {
		StageBranch(stage, defaults.Concert_score)
	}
	if defaults.Appearance != nil {
		StageBranch(stage, defaults.Appearance)
	}
	if defaults.Music_font != nil {
		StageBranch(stage, defaults.Music_font)
	}
	if defaults.Word_font != nil {
		StageBranch(stage, defaults.Word_font)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lyric_font := range defaults.Lyric_font {
		StageBranch(stage, _lyric_font)
	}
	for _, _lyric_language := range defaults.Lyric_language {
		StageBranch(stage, _lyric_language)
	}

}

func (stage *StageStruct) StageBranchDegree(degree *Degree) {

	// check if instance is already staged
	if IsStaged(stage, degree) {
		return
	}

	degree.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if degree.Degree_value != nil {
		StageBranch(stage, degree.Degree_value)
	}
	if degree.Degree_alter != nil {
		StageBranch(stage, degree.Degree_alter)
	}
	if degree.Degree_type != nil {
		StageBranch(stage, degree.Degree_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDegree_alter(degree_alter *Degree_alter) {

	// check if instance is already staged
	if IsStaged(stage, degree_alter) {
		return
	}

	degree_alter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDegree_type(degree_type *Degree_type) {

	// check if instance is already staged
	if IsStaged(stage, degree_type) {
		return
	}

	degree_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDegree_value(degree_value *Degree_value) {

	// check if instance is already staged
	if IsStaged(stage, degree_value) {
		return
	}

	degree_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDirection(direction *Direction) {

	// check if instance is already staged
	if IsStaged(stage, direction) {
		return
	}

	direction.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction.Offset != nil {
		StageBranch(stage, direction.Offset)
	}
	if direction.Sound != nil {
		StageBranch(stage, direction.Sound)
	}
	if direction.Listening != nil {
		StageBranch(stage, direction.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _direction_type := range direction.Direction_type {
		StageBranch(stage, _direction_type)
	}

}

func (stage *StageStruct) StageBranchDirection_type(direction_type *Direction_type) {

	// check if instance is already staged
	if IsStaged(stage, direction_type) {
		return
	}

	direction_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction_type.Wedge != nil {
		StageBranch(stage, direction_type.Wedge)
	}
	if direction_type.Dashes != nil {
		StageBranch(stage, direction_type.Dashes)
	}
	if direction_type.Bracket != nil {
		StageBranch(stage, direction_type.Bracket)
	}
	if direction_type.Pedal != nil {
		StageBranch(stage, direction_type.Pedal)
	}
	if direction_type.Metronome != nil {
		StageBranch(stage, direction_type.Metronome)
	}
	if direction_type.Octave_shift != nil {
		StageBranch(stage, direction_type.Octave_shift)
	}
	if direction_type.Harp_pedals != nil {
		StageBranch(stage, direction_type.Harp_pedals)
	}
	if direction_type.Damp != nil {
		StageBranch(stage, direction_type.Damp)
	}
	if direction_type.Damp_all != nil {
		StageBranch(stage, direction_type.Damp_all)
	}
	if direction_type.Eyeglasses != nil {
		StageBranch(stage, direction_type.Eyeglasses)
	}
	if direction_type.String_mute != nil {
		StageBranch(stage, direction_type.String_mute)
	}
	if direction_type.Scordatura != nil {
		StageBranch(stage, direction_type.Scordatura)
	}
	if direction_type.Image != nil {
		StageBranch(stage, direction_type.Image)
	}
	if direction_type.Principal_voice != nil {
		StageBranch(stage, direction_type.Principal_voice)
	}
	if direction_type.Accordion_registration != nil {
		StageBranch(stage, direction_type.Accordion_registration)
	}
	if direction_type.Staff_divide != nil {
		StageBranch(stage, direction_type.Staff_divide)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _segno := range direction_type.Segno {
		StageBranch(stage, _segno)
	}
	for _, _coda := range direction_type.Coda {
		StageBranch(stage, _coda)
	}
	for _, _dynamics := range direction_type.Dynamics {
		StageBranch(stage, _dynamics)
	}
	for _, _percussion := range direction_type.Percussion {
		StageBranch(stage, _percussion)
	}

}

func (stage *StageStruct) StageBranchDistance(distance *Distance) {

	// check if instance is already staged
	if IsStaged(stage, distance) {
		return
	}

	distance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDouble(double *Double) {

	// check if instance is already staged
	if IsStaged(stage, double) {
		return
	}

	double.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDynamics(dynamics *Dynamics) {

	// check if instance is already staged
	if IsStaged(stage, dynamics) {
		return
	}

	dynamics.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dynamics.P != nil {
		StageBranch(stage, dynamics.P)
	}
	if dynamics.Pp != nil {
		StageBranch(stage, dynamics.Pp)
	}
	if dynamics.Ppp != nil {
		StageBranch(stage, dynamics.Ppp)
	}
	if dynamics.Pppp != nil {
		StageBranch(stage, dynamics.Pppp)
	}
	if dynamics.Ppppp != nil {
		StageBranch(stage, dynamics.Ppppp)
	}
	if dynamics.Pppppp != nil {
		StageBranch(stage, dynamics.Pppppp)
	}
	if dynamics.F != nil {
		StageBranch(stage, dynamics.F)
	}
	if dynamics.Ff != nil {
		StageBranch(stage, dynamics.Ff)
	}
	if dynamics.Fff != nil {
		StageBranch(stage, dynamics.Fff)
	}
	if dynamics.Ffff != nil {
		StageBranch(stage, dynamics.Ffff)
	}
	if dynamics.Fffff != nil {
		StageBranch(stage, dynamics.Fffff)
	}
	if dynamics.Ffffff != nil {
		StageBranch(stage, dynamics.Ffffff)
	}
	if dynamics.Mp != nil {
		StageBranch(stage, dynamics.Mp)
	}
	if dynamics.Mf != nil {
		StageBranch(stage, dynamics.Mf)
	}
	if dynamics.Sf != nil {
		StageBranch(stage, dynamics.Sf)
	}
	if dynamics.Sfp != nil {
		StageBranch(stage, dynamics.Sfp)
	}
	if dynamics.Sfpp != nil {
		StageBranch(stage, dynamics.Sfpp)
	}
	if dynamics.Fp != nil {
		StageBranch(stage, dynamics.Fp)
	}
	if dynamics.Rf != nil {
		StageBranch(stage, dynamics.Rf)
	}
	if dynamics.Rfz != nil {
		StageBranch(stage, dynamics.Rfz)
	}
	if dynamics.Sfz != nil {
		StageBranch(stage, dynamics.Sfz)
	}
	if dynamics.Sffz != nil {
		StageBranch(stage, dynamics.Sffz)
	}
	if dynamics.Fz != nil {
		StageBranch(stage, dynamics.Fz)
	}
	if dynamics.N != nil {
		StageBranch(stage, dynamics.N)
	}
	if dynamics.Pf != nil {
		StageBranch(stage, dynamics.Pf)
	}
	if dynamics.Sfzp != nil {
		StageBranch(stage, dynamics.Sfzp)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEffect(effect *Effect) {

	// check if instance is already staged
	if IsStaged(stage, effect) {
		return
	}

	effect.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchElision(elision *Elision) {

	// check if instance is already staged
	if IsStaged(stage, elision) {
		return
	}

	elision.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty(empty *Empty) {

	// check if instance is already staged
	if IsStaged(stage, empty) {
		return
	}

	empty.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_font(empty_font *Empty_font) {

	// check if instance is already staged
	if IsStaged(stage, empty_font) {
		return
	}

	empty_font.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_line(empty_line *Empty_line) {

	// check if instance is already staged
	if IsStaged(stage, empty_line) {
		return
	}

	empty_line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_placement(empty_placement *Empty_placement) {

	// check if instance is already staged
	if IsStaged(stage, empty_placement) {
		return
	}

	empty_placement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_placement_smufl(empty_placement_smufl *Empty_placement_smufl) {

	// check if instance is already staged
	if IsStaged(stage, empty_placement_smufl) {
		return
	}

	empty_placement_smufl.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_print_object_style_align(empty_print_object_style_align *Empty_print_object_style_align) {

	// check if instance is already staged
	if IsStaged(stage, empty_print_object_style_align) {
		return
	}

	empty_print_object_style_align.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_print_style(empty_print_style *Empty_print_style) {

	// check if instance is already staged
	if IsStaged(stage, empty_print_style) {
		return
	}

	empty_print_style.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_print_style_align(empty_print_style_align *Empty_print_style_align) {

	// check if instance is already staged
	if IsStaged(stage, empty_print_style_align) {
		return
	}

	empty_print_style_align.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_print_style_align_id(empty_print_style_align_id *Empty_print_style_align_id) {

	// check if instance is already staged
	if IsStaged(stage, empty_print_style_align_id) {
		return
	}

	empty_print_style_align_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEmpty_trill_sound(empty_trill_sound *Empty_trill_sound) {

	// check if instance is already staged
	if IsStaged(stage, empty_trill_sound) {
		return
	}

	empty_trill_sound.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEncoding(encoding *Encoding) {

	// check if instance is already staged
	if IsStaged(stage, encoding) {
		return
	}

	encoding.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if encoding.Encoder != nil {
		StageBranch(stage, encoding.Encoder)
	}
	if encoding.Supports != nil {
		StageBranch(stage, encoding.Supports)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchEnding(ending *Ending) {

	// check if instance is already staged
	if IsStaged(stage, ending) {
		return
	}

	ending.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchExtend(extend *Extend) {

	// check if instance is already staged
	if IsStaged(stage, extend) {
		return
	}

	extend.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFeature(feature *Feature) {

	// check if instance is already staged
	if IsStaged(stage, feature) {
		return
	}

	feature.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFermata(fermata *Fermata) {

	// check if instance is already staged
	if IsStaged(stage, fermata) {
		return
	}

	fermata.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFigure(figure *Figure) {

	// check if instance is already staged
	if IsStaged(stage, figure) {
		return
	}

	figure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figure.Extend != nil {
		StageBranch(stage, figure.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFigured_bass(figured_bass *Figured_bass) {

	// check if instance is already staged
	if IsStaged(stage, figured_bass) {
		return
	}

	figured_bass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figured_bass.Figure {
		StageBranch(stage, _figure)
	}

}

func (stage *StageStruct) StageBranchFingering(fingering *Fingering) {

	// check if instance is already staged
	if IsStaged(stage, fingering) {
		return
	}

	fingering.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFirst_fret(first_fret *First_fret) {

	// check if instance is already staged
	if IsStaged(stage, first_fret) {
		return
	}

	first_fret.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFoo(foo *Foo) {

	// check if instance is already staged
	if IsStaged(stage, foo) {
		return
	}

	foo.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFor_part(for_part *For_part) {

	// check if instance is already staged
	if IsStaged(stage, for_part) {
		return
	}

	for_part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if for_part.Part_clef != nil {
		StageBranch(stage, for_part.Part_clef)
	}
	if for_part.Part_transpose != nil {
		StageBranch(stage, for_part.Part_transpose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormatted_symbol(formatted_symbol *Formatted_symbol) {

	// check if instance is already staged
	if IsStaged(stage, formatted_symbol) {
		return
	}

	formatted_symbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFormatted_symbol_id(formatted_symbol_id *Formatted_symbol_id) {

	// check if instance is already staged
	if IsStaged(stage, formatted_symbol_id) {
		return
	}

	formatted_symbol_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchForward(forward *Forward) {

	// check if instance is already staged
	if IsStaged(stage, forward) {
		return
	}

	forward.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFrame(frame *Frame) {

	// check if instance is already staged
	if IsStaged(stage, frame) {
		return
	}

	frame.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame.First_fret != nil {
		StageBranch(stage, frame.First_fret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frame_note := range frame.Frame_note {
		StageBranch(stage, _frame_note)
	}

}

func (stage *StageStruct) StageBranchFrame_note(frame_note *Frame_note) {

	// check if instance is already staged
	if IsStaged(stage, frame_note) {
		return
	}

	frame_note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame_note.Fret != nil {
		StageBranch(stage, frame_note.Fret)
	}
	if frame_note.Fingering != nil {
		StageBranch(stage, frame_note.Fingering)
	}
	if frame_note.Barre != nil {
		StageBranch(stage, frame_note.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchFret(fret *Fret) {

	// check if instance is already staged
	if IsStaged(stage, fret) {
		return
	}

	fret.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGlass(glass *Glass) {

	// check if instance is already staged
	if IsStaged(stage, glass) {
		return
	}

	glass.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGlissando(glissando *Glissando) {

	// check if instance is already staged
	if IsStaged(stage, glissando) {
		return
	}

	glissando.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGlyph(glyph *Glyph) {

	// check if instance is already staged
	if IsStaged(stage, glyph) {
		return
	}

	glyph.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGrace(grace *Grace) {

	// check if instance is already staged
	if IsStaged(stage, grace) {
		return
	}

	grace.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGroup_barline(group_barline *Group_barline) {

	// check if instance is already staged
	if IsStaged(stage, group_barline) {
		return
	}

	group_barline.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGroup_symbol(group_symbol *Group_symbol) {

	// check if instance is already staged
	if IsStaged(stage, group_symbol) {
		return
	}

	group_symbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGrouping(grouping *Grouping) {

	// check if instance is already staged
	if IsStaged(stage, grouping) {
		return
	}

	grouping.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range grouping.Feature {
		StageBranch(stage, _feature)
	}

}

func (stage *StageStruct) StageBranchHammer_on_pull_off(hammer_on_pull_off *Hammer_on_pull_off) {

	// check if instance is already staged
	if IsStaged(stage, hammer_on_pull_off) {
		return
	}

	hammer_on_pull_off.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHandbell(handbell *Handbell) {

	// check if instance is already staged
	if IsStaged(stage, handbell) {
		return
	}

	handbell.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHarmon_closed(harmon_closed *Harmon_closed) {

	// check if instance is already staged
	if IsStaged(stage, harmon_closed) {
		return
	}

	harmon_closed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHarmon_mute(harmon_mute *Harmon_mute) {

	// check if instance is already staged
	if IsStaged(stage, harmon_mute) {
		return
	}

	harmon_mute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmon_mute.Harmon_closed != nil {
		StageBranch(stage, harmon_mute.Harmon_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHarmonic(harmonic *Harmonic) {

	// check if instance is already staged
	if IsStaged(stage, harmonic) {
		return
	}

	harmonic.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmonic.Natural != nil {
		StageBranch(stage, harmonic.Natural)
	}
	if harmonic.Artificial != nil {
		StageBranch(stage, harmonic.Artificial)
	}
	if harmonic.Base_pitch != nil {
		StageBranch(stage, harmonic.Base_pitch)
	}
	if harmonic.Touching_pitch != nil {
		StageBranch(stage, harmonic.Touching_pitch)
	}
	if harmonic.Sounding_pitch != nil {
		StageBranch(stage, harmonic.Sounding_pitch)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHarmony(harmony *Harmony) {

	// check if instance is already staged
	if IsStaged(stage, harmony) {
		return
	}

	harmony.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmony.Frame != nil {
		StageBranch(stage, harmony.Frame)
	}
	if harmony.Offset != nil {
		StageBranch(stage, harmony.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHarmony_alter(harmony_alter *Harmony_alter) {

	// check if instance is already staged
	if IsStaged(stage, harmony_alter) {
		return
	}

	harmony_alter.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHarp_pedals(harp_pedals *Harp_pedals) {

	// check if instance is already staged
	if IsStaged(stage, harp_pedals) {
		return
	}

	harp_pedals.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
		StageBranch(stage, _pedal_tuning)
	}

}

func (stage *StageStruct) StageBranchHeel_toe(heel_toe *Heel_toe) {

	// check if instance is already staged
	if IsStaged(stage, heel_toe) {
		return
	}

	heel_toe.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHole(hole *Hole) {

	// check if instance is already staged
	if IsStaged(stage, hole) {
		return
	}

	hole.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if hole.Hole_closed != nil {
		StageBranch(stage, hole.Hole_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHole_closed(hole_closed *Hole_closed) {

	// check if instance is already staged
	if IsStaged(stage, hole_closed) {
		return
	}

	hole_closed.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchHorizontal_turn(horizontal_turn *Horizontal_turn) {

	// check if instance is already staged
	if IsStaged(stage, horizontal_turn) {
		return
	}

	horizontal_turn.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchIdentification(identification *Identification) {

	// check if instance is already staged
	if IsStaged(stage, identification) {
		return
	}

	identification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if identification.Encoding != nil {
		StageBranch(stage, identification.Encoding)
	}
	if identification.Miscellaneous != nil {
		StageBranch(stage, identification.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range identification.Creator {
		StageBranch(stage, _typed_text)
	}
	for _, _typed_text := range identification.Rights {
		StageBranch(stage, _typed_text)
	}
	for _, _typed_text := range identification.Relation {
		StageBranch(stage, _typed_text)
	}

}

func (stage *StageStruct) StageBranchImage(image *Image) {

	// check if instance is already staged
	if IsStaged(stage, image) {
		return
	}

	image.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchInstrument(instrument *Instrument) {

	// check if instance is already staged
	if IsStaged(stage, instrument) {
		return
	}

	instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchInstrument_change(instrument_change *Instrument_change) {

	// check if instance is already staged
	if IsStaged(stage, instrument_change) {
		return
	}

	instrument_change.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchInstrument_link(instrument_link *Instrument_link) {

	// check if instance is already staged
	if IsStaged(stage, instrument_link) {
		return
	}

	instrument_link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchInterchangeable(interchangeable *Interchangeable) {

	// check if instance is already staged
	if IsStaged(stage, interchangeable) {
		return
	}

	interchangeable.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchInversion(inversion *Inversion) {

	// check if instance is already staged
	if IsStaged(stage, inversion) {
		return
	}

	inversion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchKey(key *Key) {

	// check if instance is already staged
	if IsStaged(stage, key) {
		return
	}

	key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key_octave := range key.Key_octave {
		StageBranch(stage, _key_octave)
	}

}

func (stage *StageStruct) StageBranchKey_accidental(key_accidental *Key_accidental) {

	// check if instance is already staged
	if IsStaged(stage, key_accidental) {
		return
	}

	key_accidental.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchKey_octave(key_octave *Key_octave) {

	// check if instance is already staged
	if IsStaged(stage, key_octave) {
		return
	}

	key_octave.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchKind(kind *Kind) {

	// check if instance is already staged
	if IsStaged(stage, kind) {
		return
	}

	kind.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLevel(level *Level) {

	// check if instance is already staged
	if IsStaged(stage, level) {
		return
	}

	level.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLine_detail(line_detail *Line_detail) {

	// check if instance is already staged
	if IsStaged(stage, line_detail) {
		return
	}

	line_detail.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLine_width(line_width *Line_width) {

	// check if instance is already staged
	if IsStaged(stage, line_width) {
		return
	}

	line_width.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchListen(listen *Listen) {

	// check if instance is already staged
	if IsStaged(stage, listen) {
		return
	}

	listen.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listen.Assess != nil {
		StageBranch(stage, listen.Assess)
	}
	if listen.Wait != nil {
		StageBranch(stage, listen.Wait)
	}
	if listen.Other_listen != nil {
		StageBranch(stage, listen.Other_listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchListening(listening *Listening) {

	// check if instance is already staged
	if IsStaged(stage, listening) {
		return
	}

	listening.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listening.Offset != nil {
		StageBranch(stage, listening.Offset)
	}
	if listening.Sync != nil {
		StageBranch(stage, listening.Sync)
	}
	if listening.Other_listening != nil {
		StageBranch(stage, listening.Other_listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLyric(lyric *Lyric) {

	// check if instance is already staged
	if IsStaged(stage, lyric) {
		return
	}

	lyric.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lyric.End_line != nil {
		StageBranch(stage, lyric.End_line)
	}
	if lyric.End_paragraph != nil {
		StageBranch(stage, lyric.End_paragraph)
	}
	if lyric.Extend != nil {
		StageBranch(stage, lyric.Extend)
	}
	if lyric.Laughing != nil {
		StageBranch(stage, lyric.Laughing)
	}
	if lyric.Humming != nil {
		StageBranch(stage, lyric.Humming)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLyric_font(lyric_font *Lyric_font) {

	// check if instance is already staged
	if IsStaged(stage, lyric_font) {
		return
	}

	lyric_font.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLyric_language(lyric_language *Lyric_language) {

	// check if instance is already staged
	if IsStaged(stage, lyric_language) {
		return
	}

	lyric_language.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMeasure_layout(measure_layout *Measure_layout) {

	// check if instance is already staged
	if IsStaged(stage, measure_layout) {
		return
	}

	measure_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMeasure_numbering(measure_numbering *Measure_numbering) {

	// check if instance is already staged
	if IsStaged(stage, measure_numbering) {
		return
	}

	measure_numbering.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMeasure_repeat(measure_repeat *Measure_repeat) {

	// check if instance is already staged
	if IsStaged(stage, measure_repeat) {
		return
	}

	measure_repeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMeasure_style(measure_style *Measure_style) {

	// check if instance is already staged
	if IsStaged(stage, measure_style) {
		return
	}

	measure_style.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if measure_style.Multiple_rest != nil {
		StageBranch(stage, measure_style.Multiple_rest)
	}
	if measure_style.Measure_repeat != nil {
		StageBranch(stage, measure_style.Measure_repeat)
	}
	if measure_style.Beat_repeat != nil {
		StageBranch(stage, measure_style.Beat_repeat)
	}
	if measure_style.Slash != nil {
		StageBranch(stage, measure_style.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMembrane(membrane *Membrane) {

	// check if instance is already staged
	if IsStaged(stage, membrane) {
		return
	}

	membrane.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMetal(metal *Metal) {

	// check if instance is already staged
	if IsStaged(stage, metal) {
		return
	}

	metal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMetronome(metronome *Metronome) {

	// check if instance is already staged
	if IsStaged(stage, metronome) {
		return
	}

	metronome.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMetronome_beam(metronome_beam *Metronome_beam) {

	// check if instance is already staged
	if IsStaged(stage, metronome_beam) {
		return
	}

	metronome_beam.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMetronome_note(metronome_note *Metronome_note) {

	// check if instance is already staged
	if IsStaged(stage, metronome_note) {
		return
	}

	metronome_note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome_note.Metronome_tied != nil {
		StageBranch(stage, metronome_note.Metronome_tied)
	}
	if metronome_note.Metronome_tuplet != nil {
		StageBranch(stage, metronome_note.Metronome_tuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty := range metronome_note.Metronome_dot {
		StageBranch(stage, _empty)
	}
	for _, _metronome_beam := range metronome_note.Metronome_beam {
		StageBranch(stage, _metronome_beam)
	}

}

func (stage *StageStruct) StageBranchMetronome_tied(metronome_tied *Metronome_tied) {

	// check if instance is already staged
	if IsStaged(stage, metronome_tied) {
		return
	}

	metronome_tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMetronome_tuplet(metronome_tuplet *Metronome_tuplet) {

	// check if instance is already staged
	if IsStaged(stage, metronome_tuplet) {
		return
	}

	metronome_tuplet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMidi_device(midi_device *Midi_device) {

	// check if instance is already staged
	if IsStaged(stage, midi_device) {
		return
	}

	midi_device.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMidi_instrument(midi_instrument *Midi_instrument) {

	// check if instance is already staged
	if IsStaged(stage, midi_instrument) {
		return
	}

	midi_instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMiscellaneous(miscellaneous *Miscellaneous) {

	// check if instance is already staged
	if IsStaged(stage, miscellaneous) {
		return
	}

	miscellaneous.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
		StageBranch(stage, _miscellaneous_field)
	}

}

func (stage *StageStruct) StageBranchMiscellaneous_field(miscellaneous_field *Miscellaneous_field) {

	// check if instance is already staged
	if IsStaged(stage, miscellaneous_field) {
		return
	}

	miscellaneous_field.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMordent(mordent *Mordent) {

	// check if instance is already staged
	if IsStaged(stage, mordent) {
		return
	}

	mordent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMultiple_rest(multiple_rest *Multiple_rest) {

	// check if instance is already staged
	if IsStaged(stage, multiple_rest) {
		return
	}

	multiple_rest.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchName_display(name_display *Name_display) {

	// check if instance is already staged
	if IsStaged(stage, name_display) {
		return
	}

	name_display.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if name_display.Accidental_text != nil {
		StageBranch(stage, name_display.Accidental_text)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNon_arpeggiate(non_arpeggiate *Non_arpeggiate) {

	// check if instance is already staged
	if IsStaged(stage, non_arpeggiate) {
		return
	}

	non_arpeggiate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNotations(notations *Notations) {

	// check if instance is already staged
	if IsStaged(stage, notations) {
		return
	}

	notations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notations.Tied != nil {
		StageBranch(stage, notations.Tied)
	}
	if notations.Slur != nil {
		StageBranch(stage, notations.Slur)
	}
	if notations.Tuplet != nil {
		StageBranch(stage, notations.Tuplet)
	}
	if notations.Glissando != nil {
		StageBranch(stage, notations.Glissando)
	}
	if notations.Slide != nil {
		StageBranch(stage, notations.Slide)
	}
	if notations.Ornaments != nil {
		StageBranch(stage, notations.Ornaments)
	}
	if notations.Technical != nil {
		StageBranch(stage, notations.Technical)
	}
	if notations.Articulations != nil {
		StageBranch(stage, notations.Articulations)
	}
	if notations.Dynamics != nil {
		StageBranch(stage, notations.Dynamics)
	}
	if notations.Fermata != nil {
		StageBranch(stage, notations.Fermata)
	}
	if notations.Arpeggiate != nil {
		StageBranch(stage, notations.Arpeggiate)
	}
	if notations.Non_arpeggiate != nil {
		StageBranch(stage, notations.Non_arpeggiate)
	}
	if notations.Accidental_mark != nil {
		StageBranch(stage, notations.Accidental_mark)
	}
	if notations.Other_notation != nil {
		StageBranch(stage, notations.Other_notation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.Type_ != nil {
		StageBranch(stage, note.Type_)
	}
	if note.Accidental != nil {
		StageBranch(stage, note.Accidental)
	}
	if note.Time_modification != nil {
		StageBranch(stage, note.Time_modification)
	}
	if note.Stem != nil {
		StageBranch(stage, note.Stem)
	}
	if note.Notehead != nil {
		StageBranch(stage, note.Notehead)
	}
	if note.Notehead_text != nil {
		StageBranch(stage, note.Notehead_text)
	}
	if note.Beam != nil {
		StageBranch(stage, note.Beam)
	}
	if note.Play != nil {
		StageBranch(stage, note.Play)
	}
	if note.Listen != nil {
		StageBranch(stage, note.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument := range note.Instrument {
		StageBranch(stage, _instrument)
	}
	for _, _empty_placement := range note.Dot {
		StageBranch(stage, _empty_placement)
	}
	for _, _notations := range note.Notations {
		StageBranch(stage, _notations)
	}
	for _, _lyric := range note.Lyric {
		StageBranch(stage, _lyric)
	}

}

func (stage *StageStruct) StageBranchNote_size(note_size *Note_size) {

	// check if instance is already staged
	if IsStaged(stage, note_size) {
		return
	}

	note_size.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNote_type(note_type *Note_type) {

	// check if instance is already staged
	if IsStaged(stage, note_type) {
		return
	}

	note_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNotehead(notehead *Notehead) {

	// check if instance is already staged
	if IsStaged(stage, notehead) {
		return
	}

	notehead.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNotehead_text(notehead_text *Notehead_text) {

	// check if instance is already staged
	if IsStaged(stage, notehead_text) {
		return
	}

	notehead_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notehead_text.Accidental_text != nil {
		StageBranch(stage, notehead_text.Accidental_text)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNumeral(numeral *Numeral) {

	// check if instance is already staged
	if IsStaged(stage, numeral) {
		return
	}

	numeral.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if numeral.Numeral_root != nil {
		StageBranch(stage, numeral.Numeral_root)
	}
	if numeral.Numeral_alter != nil {
		StageBranch(stage, numeral.Numeral_alter)
	}
	if numeral.Numeral_key != nil {
		StageBranch(stage, numeral.Numeral_key)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNumeral_key(numeral_key *Numeral_key) {

	// check if instance is already staged
	if IsStaged(stage, numeral_key) {
		return
	}

	numeral_key.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNumeral_root(numeral_root *Numeral_root) {

	// check if instance is already staged
	if IsStaged(stage, numeral_root) {
		return
	}

	numeral_root.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOctave_shift(octave_shift *Octave_shift) {

	// check if instance is already staged
	if IsStaged(stage, octave_shift) {
		return
	}

	octave_shift.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOffset(offset *Offset) {

	// check if instance is already staged
	if IsStaged(stage, offset) {
		return
	}

	offset.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOpus(opus *Opus) {

	// check if instance is already staged
	if IsStaged(stage, opus) {
		return
	}

	opus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOrnaments(ornaments *Ornaments) {

	// check if instance is already staged
	if IsStaged(stage, ornaments) {
		return
	}

	ornaments.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ornaments.Trill_mark != nil {
		StageBranch(stage, ornaments.Trill_mark)
	}
	if ornaments.Turn != nil {
		StageBranch(stage, ornaments.Turn)
	}
	if ornaments.Delayed_turn != nil {
		StageBranch(stage, ornaments.Delayed_turn)
	}
	if ornaments.Inverted_turn != nil {
		StageBranch(stage, ornaments.Inverted_turn)
	}
	if ornaments.Delayed_inverted_turn != nil {
		StageBranch(stage, ornaments.Delayed_inverted_turn)
	}
	if ornaments.Vertical_turn != nil {
		StageBranch(stage, ornaments.Vertical_turn)
	}
	if ornaments.Inverted_vertical_turn != nil {
		StageBranch(stage, ornaments.Inverted_vertical_turn)
	}
	if ornaments.Shake != nil {
		StageBranch(stage, ornaments.Shake)
	}
	if ornaments.Wavy_line != nil {
		StageBranch(stage, ornaments.Wavy_line)
	}
	if ornaments.Mordent != nil {
		StageBranch(stage, ornaments.Mordent)
	}
	if ornaments.Inverted_mordent != nil {
		StageBranch(stage, ornaments.Inverted_mordent)
	}
	if ornaments.Schleifer != nil {
		StageBranch(stage, ornaments.Schleifer)
	}
	if ornaments.Tremolo != nil {
		StageBranch(stage, ornaments.Tremolo)
	}
	if ornaments.Haydn != nil {
		StageBranch(stage, ornaments.Haydn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accidental_mark := range ornaments.Accidental_mark {
		StageBranch(stage, _accidental_mark)
	}

}

func (stage *StageStruct) StageBranchOther_appearance(other_appearance *Other_appearance) {

	// check if instance is already staged
	if IsStaged(stage, other_appearance) {
		return
	}

	other_appearance.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOther_listening(other_listening *Other_listening) {

	// check if instance is already staged
	if IsStaged(stage, other_listening) {
		return
	}

	other_listening.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOther_notation(other_notation *Other_notation) {

	// check if instance is already staged
	if IsStaged(stage, other_notation) {
		return
	}

	other_notation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchOther_play(other_play *Other_play) {

	// check if instance is already staged
	if IsStaged(stage, other_play) {
		return
	}

	other_play.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPage_layout(page_layout *Page_layout) {

	// check if instance is already staged
	if IsStaged(stage, page_layout) {
		return
	}

	page_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if page_layout.Page_margins != nil {
		StageBranch(stage, page_layout.Page_margins)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPage_margins(page_margins *Page_margins) {

	// check if instance is already staged
	if IsStaged(stage, page_margins) {
		return
	}

	page_margins.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPart_clef(part_clef *Part_clef) {

	// check if instance is already staged
	if IsStaged(stage, part_clef) {
		return
	}

	part_clef.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPart_group(part_group *Part_group) {

	// check if instance is already staged
	if IsStaged(stage, part_group) {
		return
	}

	part_group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part_group.Group_name_display != nil {
		StageBranch(stage, part_group.Group_name_display)
	}
	if part_group.Group_abbreviation_display != nil {
		StageBranch(stage, part_group.Group_abbreviation_display)
	}
	if part_group.Group_symbol != nil {
		StageBranch(stage, part_group.Group_symbol)
	}
	if part_group.Group_barline != nil {
		StageBranch(stage, part_group.Group_barline)
	}
	if part_group.Group_time != nil {
		StageBranch(stage, part_group.Group_time)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPart_link(part_link *Part_link) {

	// check if instance is already staged
	if IsStaged(stage, part_link) {
		return
	}

	part_link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_link := range part_link.Instrument_link {
		StageBranch(stage, _instrument_link)
	}

}

func (stage *StageStruct) StageBranchPart_list(part_list *Part_list) {

	// check if instance is already staged
	if IsStaged(stage, part_list) {
		return
	}

	part_list.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPart_symbol(part_symbol *Part_symbol) {

	// check if instance is already staged
	if IsStaged(stage, part_symbol) {
		return
	}

	part_symbol.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPart_transpose(part_transpose *Part_transpose) {

	// check if instance is already staged
	if IsStaged(stage, part_transpose) {
		return
	}

	part_transpose.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPedal(pedal *Pedal) {

	// check if instance is already staged
	if IsStaged(stage, pedal) {
		return
	}

	pedal.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPedal_tuning(pedal_tuning *Pedal_tuning) {

	// check if instance is already staged
	if IsStaged(stage, pedal_tuning) {
		return
	}

	pedal_tuning.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPercussion(percussion *Percussion) {

	// check if instance is already staged
	if IsStaged(stage, percussion) {
		return
	}

	percussion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if percussion.Glass != nil {
		StageBranch(stage, percussion.Glass)
	}
	if percussion.Metal != nil {
		StageBranch(stage, percussion.Metal)
	}
	if percussion.Wood != nil {
		StageBranch(stage, percussion.Wood)
	}
	if percussion.Pitched != nil {
		StageBranch(stage, percussion.Pitched)
	}
	if percussion.Membrane != nil {
		StageBranch(stage, percussion.Membrane)
	}
	if percussion.Effect != nil {
		StageBranch(stage, percussion.Effect)
	}
	if percussion.Timpani != nil {
		StageBranch(stage, percussion.Timpani)
	}
	if percussion.Beater != nil {
		StageBranch(stage, percussion.Beater)
	}
	if percussion.Stick != nil {
		StageBranch(stage, percussion.Stick)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPitch(pitch *Pitch) {

	// check if instance is already staged
	if IsStaged(stage, pitch) {
		return
	}

	pitch.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPitched(pitched *Pitched) {

	// check if instance is already staged
	if IsStaged(stage, pitched) {
		return
	}

	pitched.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPlay(play *Play) {

	// check if instance is already staged
	if IsStaged(stage, play) {
		return
	}

	play.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if play.Other_play != nil {
		StageBranch(stage, play.Other_play)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPlayer(player *Player) {

	// check if instance is already staged
	if IsStaged(stage, player) {
		return
	}

	player.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPrincipal_voice(principal_voice *Principal_voice) {

	// check if instance is already staged
	if IsStaged(stage, principal_voice) {
		return
	}

	principal_voice.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPrint(print *Print) {

	// check if instance is already staged
	if IsStaged(stage, print) {
		return
	}

	print.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if print.Measure_layout != nil {
		StageBranch(stage, print.Measure_layout)
	}
	if print.Measure_numbering != nil {
		StageBranch(stage, print.Measure_numbering)
	}
	if print.Part_name_display != nil {
		StageBranch(stage, print.Part_name_display)
	}
	if print.Part_abbreviation_display != nil {
		StageBranch(stage, print.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRelease(release *Release) {

	// check if instance is already staged
	if IsStaged(stage, release) {
		return
	}

	release.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRepeat(repeat *Repeat) {

	// check if instance is already staged
	if IsStaged(stage, repeat) {
		return
	}

	repeat.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRest(rest *Rest) {

	// check if instance is already staged
	if IsStaged(stage, rest) {
		return
	}

	rest.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRoot(root *Root) {

	// check if instance is already staged
	if IsStaged(stage, root) {
		return
	}

	root.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if root.Root_step != nil {
		StageBranch(stage, root.Root_step)
	}
	if root.Root_alter != nil {
		StageBranch(stage, root.Root_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchRoot_step(root_step *Root_step) {

	// check if instance is already staged
	if IsStaged(stage, root_step) {
		return
	}

	root_step.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchScaling(scaling *Scaling) {

	// check if instance is already staged
	if IsStaged(stage, scaling) {
		return
	}

	scaling.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchScordatura(scordatura *Scordatura) {

	// check if instance is already staged
	if IsStaged(stage, scordatura) {
		return
	}

	scordatura.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordatura.Accord {
		StageBranch(stage, _accord)
	}

}

func (stage *StageStruct) StageBranchScore_instrument(score_instrument *Score_instrument) {

	// check if instance is already staged
	if IsStaged(stage, score_instrument) {
		return
	}

	score_instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchScore_part(score_part *Score_part) {

	// check if instance is already staged
	if IsStaged(stage, score_part) {
		return
	}

	score_part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_part.Identification != nil {
		StageBranch(stage, score_part.Identification)
	}
	if score_part.Part_name_display != nil {
		StageBranch(stage, score_part.Part_name_display)
	}
	if score_part.Part_abbreviation_display != nil {
		StageBranch(stage, score_part.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part_link := range score_part.Part_link {
		StageBranch(stage, _part_link)
	}
	for _, _score_instrument := range score_part.Score_instrument {
		StageBranch(stage, _score_instrument)
	}
	for _, _player := range score_part.Player {
		StageBranch(stage, _player)
	}

}

func (stage *StageStruct) StageBranchScore_partwise(score_partwise *Score_partwise) {

	// check if instance is already staged
	if IsStaged(stage, score_partwise) {
		return
	}

	score_partwise.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchScore_timewise(score_timewise *Score_timewise) {

	// check if instance is already staged
	if IsStaged(stage, score_timewise) {
		return
	}

	score_timewise.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSegno(segno *Segno) {

	// check if instance is already staged
	if IsStaged(stage, segno) {
		return
	}

	segno.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSlash(slash *Slash) {

	// check if instance is already staged
	if IsStaged(stage, slash) {
		return
	}

	slash.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSlide(slide *Slide) {

	// check if instance is already staged
	if IsStaged(stage, slide) {
		return
	}

	slide.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSlur(slur *Slur) {

	// check if instance is already staged
	if IsStaged(stage, slur) {
		return
	}

	slur.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSound(sound *Sound) {

	// check if instance is already staged
	if IsStaged(stage, sound) {
		return
	}

	sound.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sound.Swing != nil {
		StageBranch(stage, sound.Swing)
	}
	if sound.Offset != nil {
		StageBranch(stage, sound.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStaff_details(staff_details *Staff_details) {

	// check if instance is already staged
	if IsStaged(stage, staff_details) {
		return
	}

	staff_details.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if staff_details.Staff_size != nil {
		StageBranch(stage, staff_details.Staff_size)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_tuning := range staff_details.Staff_tuning {
		StageBranch(stage, _staff_tuning)
	}

}

func (stage *StageStruct) StageBranchStaff_divide(staff_divide *Staff_divide) {

	// check if instance is already staged
	if IsStaged(stage, staff_divide) {
		return
	}

	staff_divide.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStaff_layout(staff_layout *Staff_layout) {

	// check if instance is already staged
	if IsStaged(stage, staff_layout) {
		return
	}

	staff_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStaff_size(staff_size *Staff_size) {

	// check if instance is already staged
	if IsStaged(stage, staff_size) {
		return
	}

	staff_size.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStaff_tuning(staff_tuning *Staff_tuning) {

	// check if instance is already staged
	if IsStaged(stage, staff_tuning) {
		return
	}

	staff_tuning.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStem(stem *Stem) {

	// check if instance is already staged
	if IsStaged(stage, stem) {
		return
	}

	stem.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStick(stick *Stick) {

	// check if instance is already staged
	if IsStaged(stage, stick) {
		return
	}

	stick.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchString_mute(string_mute *String_mute) {

	// check if instance is already staged
	if IsStaged(stage, string_mute) {
		return
	}

	string_mute.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchStrong_accent(strong_accent *Strong_accent) {

	// check if instance is already staged
	if IsStaged(stage, strong_accent) {
		return
	}

	strong_accent.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSupports(supports *Supports) {

	// check if instance is already staged
	if IsStaged(stage, supports) {
		return
	}

	supports.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSwing(swing *Swing) {

	// check if instance is already staged
	if IsStaged(stage, swing) {
		return
	}

	swing.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if swing.Straight != nil {
		StageBranch(stage, swing.Straight)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSync(sync *Sync) {

	// check if instance is already staged
	if IsStaged(stage, sync) {
		return
	}

	sync.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSystem_dividers(system_dividers *System_dividers) {

	// check if instance is already staged
	if IsStaged(stage, system_dividers) {
		return
	}

	system_dividers.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_dividers.Left_divider != nil {
		StageBranch(stage, system_dividers.Left_divider)
	}
	if system_dividers.Right_divider != nil {
		StageBranch(stage, system_dividers.Right_divider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSystem_layout(system_layout *System_layout) {

	// check if instance is already staged
	if IsStaged(stage, system_layout) {
		return
	}

	system_layout.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_layout.System_margins != nil {
		StageBranch(stage, system_layout.System_margins)
	}
	if system_layout.System_dividers != nil {
		StageBranch(stage, system_layout.System_dividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSystem_margins(system_margins *System_margins) {

	// check if instance is already staged
	if IsStaged(stage, system_margins) {
		return
	}

	system_margins.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTap(tap *Tap) {

	// check if instance is already staged
	if IsStaged(stage, tap) {
		return
	}

	tap.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTechnical(technical *Technical) {

	// check if instance is already staged
	if IsStaged(stage, technical) {
		return
	}

	technical.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if technical.Up_bow != nil {
		StageBranch(stage, technical.Up_bow)
	}
	if technical.Down_bow != nil {
		StageBranch(stage, technical.Down_bow)
	}
	if technical.Harmonic != nil {
		StageBranch(stage, technical.Harmonic)
	}
	if technical.Open_string != nil {
		StageBranch(stage, technical.Open_string)
	}
	if technical.Thumb_position != nil {
		StageBranch(stage, technical.Thumb_position)
	}
	if technical.Fingering != nil {
		StageBranch(stage, technical.Fingering)
	}
	if technical.Double_tongue != nil {
		StageBranch(stage, technical.Double_tongue)
	}
	if technical.Triple_tongue != nil {
		StageBranch(stage, technical.Triple_tongue)
	}
	if technical.Stopped != nil {
		StageBranch(stage, technical.Stopped)
	}
	if technical.Snap_pizzicato != nil {
		StageBranch(stage, technical.Snap_pizzicato)
	}
	if technical.Fret != nil {
		StageBranch(stage, technical.Fret)
	}
	if technical.Hammer_on != nil {
		StageBranch(stage, technical.Hammer_on)
	}
	if technical.Pull_off != nil {
		StageBranch(stage, technical.Pull_off)
	}
	if technical.Bend != nil {
		StageBranch(stage, technical.Bend)
	}
	if technical.Tap != nil {
		StageBranch(stage, technical.Tap)
	}
	if technical.Heel != nil {
		StageBranch(stage, technical.Heel)
	}
	if technical.Toe != nil {
		StageBranch(stage, technical.Toe)
	}
	if technical.Fingernails != nil {
		StageBranch(stage, technical.Fingernails)
	}
	if technical.Hole != nil {
		StageBranch(stage, technical.Hole)
	}
	if technical.Arrow != nil {
		StageBranch(stage, technical.Arrow)
	}
	if technical.Handbell != nil {
		StageBranch(stage, technical.Handbell)
	}
	if technical.Brass_bend != nil {
		StageBranch(stage, technical.Brass_bend)
	}
	if technical.Flip != nil {
		StageBranch(stage, technical.Flip)
	}
	if technical.Smear != nil {
		StageBranch(stage, technical.Smear)
	}
	if technical.Open != nil {
		StageBranch(stage, technical.Open)
	}
	if technical.Half_muted != nil {
		StageBranch(stage, technical.Half_muted)
	}
	if technical.Harmon_mute != nil {
		StageBranch(stage, technical.Harmon_mute)
	}
	if technical.Golpe != nil {
		StageBranch(stage, technical.Golpe)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchText_element_data(text_element_data *Text_element_data) {

	// check if instance is already staged
	if IsStaged(stage, text_element_data) {
		return
	}

	text_element_data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTie(tie *Tie) {

	// check if instance is already staged
	if IsStaged(stage, tie) {
		return
	}

	tie.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTied(tied *Tied) {

	// check if instance is already staged
	if IsStaged(stage, tied) {
		return
	}

	tied.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTime(time *Time) {

	// check if instance is already staged
	if IsStaged(stage, time) {
		return
	}

	time.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTime_modification(time_modification *Time_modification) {

	// check if instance is already staged
	if IsStaged(stage, time_modification) {
		return
	}

	time_modification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTimpani(timpani *Timpani) {

	// check if instance is already staged
	if IsStaged(stage, timpani) {
		return
	}

	timpani.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTranspose(transpose *Transpose) {

	// check if instance is already staged
	if IsStaged(stage, transpose) {
		return
	}

	transpose.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTremolo(tremolo *Tremolo) {

	// check if instance is already staged
	if IsStaged(stage, tremolo) {
		return
	}

	tremolo.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTuplet(tuplet *Tuplet) {

	// check if instance is already staged
	if IsStaged(stage, tuplet) {
		return
	}

	tuplet.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet.Tuplet_actual != nil {
		StageBranch(stage, tuplet.Tuplet_actual)
	}
	if tuplet.Tuplet_normal != nil {
		StageBranch(stage, tuplet.Tuplet_normal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTuplet_dot(tuplet_dot *Tuplet_dot) {

	// check if instance is already staged
	if IsStaged(stage, tuplet_dot) {
		return
	}

	tuplet_dot.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTuplet_number(tuplet_number *Tuplet_number) {

	// check if instance is already staged
	if IsStaged(stage, tuplet_number) {
		return
	}

	tuplet_number.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTuplet_portion(tuplet_portion *Tuplet_portion) {

	// check if instance is already staged
	if IsStaged(stage, tuplet_portion) {
		return
	}

	tuplet_portion.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet_portion.Tuplet_number != nil {
		StageBranch(stage, tuplet_portion.Tuplet_number)
	}
	if tuplet_portion.Tuplet_type != nil {
		StageBranch(stage, tuplet_portion.Tuplet_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
		StageBranch(stage, _tuplet_dot)
	}

}

func (stage *StageStruct) StageBranchTuplet_type(tuplet_type *Tuplet_type) {

	// check if instance is already staged
	if IsStaged(stage, tuplet_type) {
		return
	}

	tuplet_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchTyped_text(typed_text *Typed_text) {

	// check if instance is already staged
	if IsStaged(stage, typed_text) {
		return
	}

	typed_text.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchUnpitched(unpitched *Unpitched) {

	// check if instance is already staged
	if IsStaged(stage, unpitched) {
		return
	}

	unpitched.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchVirtual_instrument(virtual_instrument *Virtual_instrument) {

	// check if instance is already staged
	if IsStaged(stage, virtual_instrument) {
		return
	}

	virtual_instrument.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWait(wait *Wait) {

	// check if instance is already staged
	if IsStaged(stage, wait) {
		return
	}

	wait.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWavy_line(wavy_line *Wavy_line) {

	// check if instance is already staged
	if IsStaged(stage, wavy_line) {
		return
	}

	wavy_line.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWedge(wedge *Wedge) {

	// check if instance is already staged
	if IsStaged(stage, wedge) {
		return
	}

	wedge.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWood(wood *Wood) {

	// check if instance is already staged
	if IsStaged(stage, wood) {
		return
	}

	wood.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchWork(work *Work) {

	// check if instance is already staged
	if IsStaged(stage, work) {
		return
	}

	work.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if work.Opus != nil {
		StageBranch(stage, work.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Accidental:
		toT := CopyBranchAccidental(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accidental_mark:
		toT := CopyBranchAccidental_mark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accidental_text:
		toT := CopyBranchAccidental_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accord:
		toT := CopyBranchAccord(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Accordion_registration:
		toT := CopyBranchAccordion_registration(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AnyType:
		toT := CopyBranchAnyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Appearance:
		toT := CopyBranchAppearance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Arpeggiate:
		toT := CopyBranchArpeggiate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Arrow:
		toT := CopyBranchArrow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Articulations:
		toT := CopyBranchArticulations(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Assess:
		toT := CopyBranchAssess(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Attributes:
		toT := CopyBranchAttributes(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Backup:
		toT := CopyBranchBackup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bar_style_color:
		toT := CopyBranchBar_style_color(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Barline:
		toT := CopyBranchBarline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Barre:
		toT := CopyBranchBarre(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bass:
		toT := CopyBranchBass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bass_step:
		toT := CopyBranchBass_step(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beam:
		toT := CopyBranchBeam(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beat_repeat:
		toT := CopyBranchBeat_repeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beat_unit_tied:
		toT := CopyBranchBeat_unit_tied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Beater:
		toT := CopyBranchBeater(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bend:
		toT := CopyBranchBend(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bookmark:
		toT := CopyBranchBookmark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Bracket:
		toT := CopyBranchBracket(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Breath_mark:
		toT := CopyBranchBreath_mark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Caesura:
		toT := CopyBranchCaesura(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Cancel:
		toT := CopyBranchCancel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Clef:
		toT := CopyBranchClef(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Coda:
		toT := CopyBranchCoda(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Credit:
		toT := CopyBranchCredit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dashes:
		toT := CopyBranchDashes(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Defaults:
		toT := CopyBranchDefaults(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree:
		toT := CopyBranchDegree(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree_alter:
		toT := CopyBranchDegree_alter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree_type:
		toT := CopyBranchDegree_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Degree_value:
		toT := CopyBranchDegree_value(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Direction:
		toT := CopyBranchDirection(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Direction_type:
		toT := CopyBranchDirection_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Distance:
		toT := CopyBranchDistance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Double:
		toT := CopyBranchDouble(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Dynamics:
		toT := CopyBranchDynamics(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Effect:
		toT := CopyBranchEffect(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Elision:
		toT := CopyBranchElision(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty:
		toT := CopyBranchEmpty(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_font:
		toT := CopyBranchEmpty_font(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_line:
		toT := CopyBranchEmpty_line(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_placement:
		toT := CopyBranchEmpty_placement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_placement_smufl:
		toT := CopyBranchEmpty_placement_smufl(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_object_style_align:
		toT := CopyBranchEmpty_print_object_style_align(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_style:
		toT := CopyBranchEmpty_print_style(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_style_align:
		toT := CopyBranchEmpty_print_style_align(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_print_style_align_id:
		toT := CopyBranchEmpty_print_style_align_id(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Empty_trill_sound:
		toT := CopyBranchEmpty_trill_sound(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Encoding:
		toT := CopyBranchEncoding(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ending:
		toT := CopyBranchEnding(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Extend:
		toT := CopyBranchExtend(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Feature:
		toT := CopyBranchFeature(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fermata:
		toT := CopyBranchFermata(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Figure:
		toT := CopyBranchFigure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Figured_bass:
		toT := CopyBranchFigured_bass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fingering:
		toT := CopyBranchFingering(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *First_fret:
		toT := CopyBranchFirst_fret(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Foo:
		toT := CopyBranchFoo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *For_part:
		toT := CopyBranchFor_part(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Formatted_symbol:
		toT := CopyBranchFormatted_symbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Formatted_symbol_id:
		toT := CopyBranchFormatted_symbol_id(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Forward:
		toT := CopyBranchForward(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Frame:
		toT := CopyBranchFrame(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Frame_note:
		toT := CopyBranchFrame_note(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Fret:
		toT := CopyBranchFret(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glass:
		toT := CopyBranchGlass(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glissando:
		toT := CopyBranchGlissando(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Glyph:
		toT := CopyBranchGlyph(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Grace:
		toT := CopyBranchGrace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group_barline:
		toT := CopyBranchGroup_barline(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Group_symbol:
		toT := CopyBranchGroup_symbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Grouping:
		toT := CopyBranchGrouping(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hammer_on_pull_off:
		toT := CopyBranchHammer_on_pull_off(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Handbell:
		toT := CopyBranchHandbell(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmon_closed:
		toT := CopyBranchHarmon_closed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmon_mute:
		toT := CopyBranchHarmon_mute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmonic:
		toT := CopyBranchHarmonic(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmony:
		toT := CopyBranchHarmony(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harmony_alter:
		toT := CopyBranchHarmony_alter(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Harp_pedals:
		toT := CopyBranchHarp_pedals(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Heel_toe:
		toT := CopyBranchHeel_toe(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hole:
		toT := CopyBranchHole(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Hole_closed:
		toT := CopyBranchHole_closed(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Horizontal_turn:
		toT := CopyBranchHorizontal_turn(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Identification:
		toT := CopyBranchIdentification(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Image:
		toT := CopyBranchImage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument:
		toT := CopyBranchInstrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument_change:
		toT := CopyBranchInstrument_change(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Instrument_link:
		toT := CopyBranchInstrument_link(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Interchangeable:
		toT := CopyBranchInterchangeable(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Inversion:
		toT := CopyBranchInversion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key:
		toT := CopyBranchKey(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key_accidental:
		toT := CopyBranchKey_accidental(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Key_octave:
		toT := CopyBranchKey_octave(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Kind:
		toT := CopyBranchKind(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Level:
		toT := CopyBranchLevel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line_detail:
		toT := CopyBranchLine_detail(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Line_width:
		toT := CopyBranchLine_width(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := CopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Listen:
		toT := CopyBranchListen(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Listening:
		toT := CopyBranchListening(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric:
		toT := CopyBranchLyric(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric_font:
		toT := CopyBranchLyric_font(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Lyric_language:
		toT := CopyBranchLyric_language(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_layout:
		toT := CopyBranchMeasure_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_numbering:
		toT := CopyBranchMeasure_numbering(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_repeat:
		toT := CopyBranchMeasure_repeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Measure_style:
		toT := CopyBranchMeasure_style(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Membrane:
		toT := CopyBranchMembrane(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metal:
		toT := CopyBranchMetal(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome:
		toT := CopyBranchMetronome(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_beam:
		toT := CopyBranchMetronome_beam(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_note:
		toT := CopyBranchMetronome_note(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_tied:
		toT := CopyBranchMetronome_tied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Metronome_tuplet:
		toT := CopyBranchMetronome_tuplet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Midi_device:
		toT := CopyBranchMidi_device(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Midi_instrument:
		toT := CopyBranchMidi_instrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Miscellaneous:
		toT := CopyBranchMiscellaneous(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Miscellaneous_field:
		toT := CopyBranchMiscellaneous_field(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Mordent:
		toT := CopyBranchMordent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Multiple_rest:
		toT := CopyBranchMultiple_rest(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Name_display:
		toT := CopyBranchName_display(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Non_arpeggiate:
		toT := CopyBranchNon_arpeggiate(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notations:
		toT := CopyBranchNotations(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note_size:
		toT := CopyBranchNote_size(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note_type:
		toT := CopyBranchNote_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notehead:
		toT := CopyBranchNotehead(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Notehead_text:
		toT := CopyBranchNotehead_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral:
		toT := CopyBranchNumeral(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral_key:
		toT := CopyBranchNumeral_key(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Numeral_root:
		toT := CopyBranchNumeral_root(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Octave_shift:
		toT := CopyBranchOctave_shift(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Offset:
		toT := CopyBranchOffset(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Opus:
		toT := CopyBranchOpus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Ornaments:
		toT := CopyBranchOrnaments(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_appearance:
		toT := CopyBranchOther_appearance(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_listening:
		toT := CopyBranchOther_listening(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_notation:
		toT := CopyBranchOther_notation(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Other_play:
		toT := CopyBranchOther_play(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Page_layout:
		toT := CopyBranchPage_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Page_margins:
		toT := CopyBranchPage_margins(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_clef:
		toT := CopyBranchPart_clef(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_group:
		toT := CopyBranchPart_group(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_link:
		toT := CopyBranchPart_link(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_list:
		toT := CopyBranchPart_list(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_symbol:
		toT := CopyBranchPart_symbol(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part_transpose:
		toT := CopyBranchPart_transpose(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pedal:
		toT := CopyBranchPedal(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pedal_tuning:
		toT := CopyBranchPedal_tuning(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Percussion:
		toT := CopyBranchPercussion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pitch:
		toT := CopyBranchPitch(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Pitched:
		toT := CopyBranchPitched(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Play:
		toT := CopyBranchPlay(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Player:
		toT := CopyBranchPlayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Principal_voice:
		toT := CopyBranchPrincipal_voice(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Print:
		toT := CopyBranchPrint(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Release:
		toT := CopyBranchRelease(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Repeat:
		toT := CopyBranchRepeat(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Rest:
		toT := CopyBranchRest(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Root:
		toT := CopyBranchRoot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Root_step:
		toT := CopyBranchRoot_step(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scaling:
		toT := CopyBranchScaling(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Scordatura:
		toT := CopyBranchScordatura(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_instrument:
		toT := CopyBranchScore_instrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_part:
		toT := CopyBranchScore_part(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_partwise:
		toT := CopyBranchScore_partwise(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Score_timewise:
		toT := CopyBranchScore_timewise(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Segno:
		toT := CopyBranchSegno(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slash:
		toT := CopyBranchSlash(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slide:
		toT := CopyBranchSlide(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Slur:
		toT := CopyBranchSlur(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sound:
		toT := CopyBranchSound(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_details:
		toT := CopyBranchStaff_details(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_divide:
		toT := CopyBranchStaff_divide(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_layout:
		toT := CopyBranchStaff_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_size:
		toT := CopyBranchStaff_size(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Staff_tuning:
		toT := CopyBranchStaff_tuning(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stem:
		toT := CopyBranchStem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Stick:
		toT := CopyBranchStick(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *String_mute:
		toT := CopyBranchString_mute(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Strong_accent:
		toT := CopyBranchStrong_accent(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Supports:
		toT := CopyBranchSupports(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Swing:
		toT := CopyBranchSwing(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Sync:
		toT := CopyBranchSync(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System_dividers:
		toT := CopyBranchSystem_dividers(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System_layout:
		toT := CopyBranchSystem_layout(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System_margins:
		toT := CopyBranchSystem_margins(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tap:
		toT := CopyBranchTap(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Technical:
		toT := CopyBranchTechnical(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Text_element_data:
		toT := CopyBranchText_element_data(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tie:
		toT := CopyBranchTie(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tied:
		toT := CopyBranchTied(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Time:
		toT := CopyBranchTime(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Time_modification:
		toT := CopyBranchTime_modification(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Timpani:
		toT := CopyBranchTimpani(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Transpose:
		toT := CopyBranchTranspose(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tremolo:
		toT := CopyBranchTremolo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet:
		toT := CopyBranchTuplet(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_dot:
		toT := CopyBranchTuplet_dot(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_number:
		toT := CopyBranchTuplet_number(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_portion:
		toT := CopyBranchTuplet_portion(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Tuplet_type:
		toT := CopyBranchTuplet_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Typed_text:
		toT := CopyBranchTyped_text(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Unpitched:
		toT := CopyBranchUnpitched(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Virtual_instrument:
		toT := CopyBranchVirtual_instrument(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wait:
		toT := CopyBranchWait(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wavy_line:
		toT := CopyBranchWavy_line(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wedge:
		toT := CopyBranchWedge(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Wood:
		toT := CopyBranchWood(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Work:
		toT := CopyBranchWork(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAccidental(mapOrigCopy map[any]any, accidentalFrom *Accidental) (accidentalTo *Accidental) {

	// accidentalFrom has already been copied
	if _accidentalTo, ok := mapOrigCopy[accidentalFrom]; ok {
		accidentalTo = _accidentalTo.(*Accidental)
		return
	}

	accidentalTo = new(Accidental)
	mapOrigCopy[accidentalFrom] = accidentalTo
	accidentalFrom.CopyBasicFields(accidentalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAccidental_mark(mapOrigCopy map[any]any, accidental_markFrom *Accidental_mark) (accidental_markTo *Accidental_mark) {

	// accidental_markFrom has already been copied
	if _accidental_markTo, ok := mapOrigCopy[accidental_markFrom]; ok {
		accidental_markTo = _accidental_markTo.(*Accidental_mark)
		return
	}

	accidental_markTo = new(Accidental_mark)
	mapOrigCopy[accidental_markFrom] = accidental_markTo
	accidental_markFrom.CopyBasicFields(accidental_markTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAccidental_text(mapOrigCopy map[any]any, accidental_textFrom *Accidental_text) (accidental_textTo *Accidental_text) {

	// accidental_textFrom has already been copied
	if _accidental_textTo, ok := mapOrigCopy[accidental_textFrom]; ok {
		accidental_textTo = _accidental_textTo.(*Accidental_text)
		return
	}

	accidental_textTo = new(Accidental_text)
	mapOrigCopy[accidental_textFrom] = accidental_textTo
	accidental_textFrom.CopyBasicFields(accidental_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAccord(mapOrigCopy map[any]any, accordFrom *Accord) (accordTo *Accord) {

	// accordFrom has already been copied
	if _accordTo, ok := mapOrigCopy[accordFrom]; ok {
		accordTo = _accordTo.(*Accord)
		return
	}

	accordTo = new(Accord)
	mapOrigCopy[accordFrom] = accordTo
	accordFrom.CopyBasicFields(accordTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAccordion_registration(mapOrigCopy map[any]any, accordion_registrationFrom *Accordion_registration) (accordion_registrationTo *Accordion_registration) {

	// accordion_registrationFrom has already been copied
	if _accordion_registrationTo, ok := mapOrigCopy[accordion_registrationFrom]; ok {
		accordion_registrationTo = _accordion_registrationTo.(*Accordion_registration)
		return
	}

	accordion_registrationTo = new(Accordion_registration)
	mapOrigCopy[accordion_registrationFrom] = accordion_registrationTo
	accordion_registrationFrom.CopyBasicFields(accordion_registrationTo)

	//insertion point for the staging of instances referenced by pointers
	if accordion_registrationFrom.Accordion_high != nil {
		accordion_registrationTo.Accordion_high = CopyBranchEmpty(mapOrigCopy, accordion_registrationFrom.Accordion_high)
	}
	if accordion_registrationFrom.Accordion_low != nil {
		accordion_registrationTo.Accordion_low = CopyBranchEmpty(mapOrigCopy, accordion_registrationFrom.Accordion_low)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAnyType(mapOrigCopy map[any]any, anytypeFrom *AnyType) (anytypeTo *AnyType) {

	// anytypeFrom has already been copied
	if _anytypeTo, ok := mapOrigCopy[anytypeFrom]; ok {
		anytypeTo = _anytypeTo.(*AnyType)
		return
	}

	anytypeTo = new(AnyType)
	mapOrigCopy[anytypeFrom] = anytypeTo
	anytypeFrom.CopyBasicFields(anytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAppearance(mapOrigCopy map[any]any, appearanceFrom *Appearance) (appearanceTo *Appearance) {

	// appearanceFrom has already been copied
	if _appearanceTo, ok := mapOrigCopy[appearanceFrom]; ok {
		appearanceTo = _appearanceTo.(*Appearance)
		return
	}

	appearanceTo = new(Appearance)
	mapOrigCopy[appearanceFrom] = appearanceTo
	appearanceFrom.CopyBasicFields(appearanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_width := range appearanceFrom.Line_width {
		appearanceTo.Line_width = append(appearanceTo.Line_width, CopyBranchLine_width(mapOrigCopy, _line_width))
	}
	for _, _note_size := range appearanceFrom.Note_size {
		appearanceTo.Note_size = append(appearanceTo.Note_size, CopyBranchNote_size(mapOrigCopy, _note_size))
	}
	for _, _distance := range appearanceFrom.Distance {
		appearanceTo.Distance = append(appearanceTo.Distance, CopyBranchDistance(mapOrigCopy, _distance))
	}
	for _, _glyph := range appearanceFrom.Glyph {
		appearanceTo.Glyph = append(appearanceTo.Glyph, CopyBranchGlyph(mapOrigCopy, _glyph))
	}
	for _, _other_appearance := range appearanceFrom.Other_appearance {
		appearanceTo.Other_appearance = append(appearanceTo.Other_appearance, CopyBranchOther_appearance(mapOrigCopy, _other_appearance))
	}

	return
}

func CopyBranchArpeggiate(mapOrigCopy map[any]any, arpeggiateFrom *Arpeggiate) (arpeggiateTo *Arpeggiate) {

	// arpeggiateFrom has already been copied
	if _arpeggiateTo, ok := mapOrigCopy[arpeggiateFrom]; ok {
		arpeggiateTo = _arpeggiateTo.(*Arpeggiate)
		return
	}

	arpeggiateTo = new(Arpeggiate)
	mapOrigCopy[arpeggiateFrom] = arpeggiateTo
	arpeggiateFrom.CopyBasicFields(arpeggiateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArrow(mapOrigCopy map[any]any, arrowFrom *Arrow) (arrowTo *Arrow) {

	// arrowFrom has already been copied
	if _arrowTo, ok := mapOrigCopy[arrowFrom]; ok {
		arrowTo = _arrowTo.(*Arrow)
		return
	}

	arrowTo = new(Arrow)
	mapOrigCopy[arrowFrom] = arrowTo
	arrowFrom.CopyBasicFields(arrowTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArticulations(mapOrigCopy map[any]any, articulationsFrom *Articulations) (articulationsTo *Articulations) {

	// articulationsFrom has already been copied
	if _articulationsTo, ok := mapOrigCopy[articulationsFrom]; ok {
		articulationsTo = _articulationsTo.(*Articulations)
		return
	}

	articulationsTo = new(Articulations)
	mapOrigCopy[articulationsFrom] = articulationsTo
	articulationsFrom.CopyBasicFields(articulationsTo)

	//insertion point for the staging of instances referenced by pointers
	if articulationsFrom.Accent != nil {
		articulationsTo.Accent = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Accent)
	}
	if articulationsFrom.Strong_accent != nil {
		articulationsTo.Strong_accent = CopyBranchStrong_accent(mapOrigCopy, articulationsFrom.Strong_accent)
	}
	if articulationsFrom.Staccato != nil {
		articulationsTo.Staccato = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Staccato)
	}
	if articulationsFrom.Tenuto != nil {
		articulationsTo.Tenuto = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Tenuto)
	}
	if articulationsFrom.Detached_legato != nil {
		articulationsTo.Detached_legato = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Detached_legato)
	}
	if articulationsFrom.Staccatissimo != nil {
		articulationsTo.Staccatissimo = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Staccatissimo)
	}
	if articulationsFrom.Spiccato != nil {
		articulationsTo.Spiccato = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Spiccato)
	}
	if articulationsFrom.Scoop != nil {
		articulationsTo.Scoop = CopyBranchEmpty_line(mapOrigCopy, articulationsFrom.Scoop)
	}
	if articulationsFrom.Plop != nil {
		articulationsTo.Plop = CopyBranchEmpty_line(mapOrigCopy, articulationsFrom.Plop)
	}
	if articulationsFrom.Doit != nil {
		articulationsTo.Doit = CopyBranchEmpty_line(mapOrigCopy, articulationsFrom.Doit)
	}
	if articulationsFrom.Falloff != nil {
		articulationsTo.Falloff = CopyBranchEmpty_line(mapOrigCopy, articulationsFrom.Falloff)
	}
	if articulationsFrom.Breath_mark != nil {
		articulationsTo.Breath_mark = CopyBranchBreath_mark(mapOrigCopy, articulationsFrom.Breath_mark)
	}
	if articulationsFrom.Caesura != nil {
		articulationsTo.Caesura = CopyBranchCaesura(mapOrigCopy, articulationsFrom.Caesura)
	}
	if articulationsFrom.Stress != nil {
		articulationsTo.Stress = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Stress)
	}
	if articulationsFrom.Unstress != nil {
		articulationsTo.Unstress = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Unstress)
	}
	if articulationsFrom.Soft_accent != nil {
		articulationsTo.Soft_accent = CopyBranchEmpty_placement(mapOrigCopy, articulationsFrom.Soft_accent)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAssess(mapOrigCopy map[any]any, assessFrom *Assess) (assessTo *Assess) {

	// assessFrom has already been copied
	if _assessTo, ok := mapOrigCopy[assessFrom]; ok {
		assessTo = _assessTo.(*Assess)
		return
	}

	assessTo = new(Assess)
	mapOrigCopy[assessFrom] = assessTo
	assessFrom.CopyBasicFields(assessTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAttributes(mapOrigCopy map[any]any, attributesFrom *Attributes) (attributesTo *Attributes) {

	// attributesFrom has already been copied
	if _attributesTo, ok := mapOrigCopy[attributesFrom]; ok {
		attributesTo = _attributesTo.(*Attributes)
		return
	}

	attributesTo = new(Attributes)
	mapOrigCopy[attributesFrom] = attributesTo
	attributesFrom.CopyBasicFields(attributesTo)

	//insertion point for the staging of instances referenced by pointers
	if attributesFrom.Part_symbol != nil {
		attributesTo.Part_symbol = CopyBranchPart_symbol(mapOrigCopy, attributesFrom.Part_symbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributesFrom.Key {
		attributesTo.Key = append(attributesTo.Key, CopyBranchKey(mapOrigCopy, _key))
	}
	for _, _clef := range attributesFrom.Clef {
		attributesTo.Clef = append(attributesTo.Clef, CopyBranchClef(mapOrigCopy, _clef))
	}
	for _, _staff_details := range attributesFrom.Staff_details {
		attributesTo.Staff_details = append(attributesTo.Staff_details, CopyBranchStaff_details(mapOrigCopy, _staff_details))
	}
	for _, _measure_style := range attributesFrom.Measure_style {
		attributesTo.Measure_style = append(attributesTo.Measure_style, CopyBranchMeasure_style(mapOrigCopy, _measure_style))
	}
	for _, _transpose := range attributesFrom.Transpose {
		attributesTo.Transpose = append(attributesTo.Transpose, CopyBranchTranspose(mapOrigCopy, _transpose))
	}
	for _, _for_part := range attributesFrom.For_part {
		attributesTo.For_part = append(attributesTo.For_part, CopyBranchFor_part(mapOrigCopy, _for_part))
	}

	return
}

func CopyBranchBackup(mapOrigCopy map[any]any, backupFrom *Backup) (backupTo *Backup) {

	// backupFrom has already been copied
	if _backupTo, ok := mapOrigCopy[backupFrom]; ok {
		backupTo = _backupTo.(*Backup)
		return
	}

	backupTo = new(Backup)
	mapOrigCopy[backupFrom] = backupTo
	backupFrom.CopyBasicFields(backupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBar_style_color(mapOrigCopy map[any]any, bar_style_colorFrom *Bar_style_color) (bar_style_colorTo *Bar_style_color) {

	// bar_style_colorFrom has already been copied
	if _bar_style_colorTo, ok := mapOrigCopy[bar_style_colorFrom]; ok {
		bar_style_colorTo = _bar_style_colorTo.(*Bar_style_color)
		return
	}

	bar_style_colorTo = new(Bar_style_color)
	mapOrigCopy[bar_style_colorFrom] = bar_style_colorTo
	bar_style_colorFrom.CopyBasicFields(bar_style_colorTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBarline(mapOrigCopy map[any]any, barlineFrom *Barline) (barlineTo *Barline) {

	// barlineFrom has already been copied
	if _barlineTo, ok := mapOrigCopy[barlineFrom]; ok {
		barlineTo = _barlineTo.(*Barline)
		return
	}

	barlineTo = new(Barline)
	mapOrigCopy[barlineFrom] = barlineTo
	barlineFrom.CopyBasicFields(barlineTo)

	//insertion point for the staging of instances referenced by pointers
	if barlineFrom.Bar_style != nil {
		barlineTo.Bar_style = CopyBranchBar_style_color(mapOrigCopy, barlineFrom.Bar_style)
	}
	if barlineFrom.Wavy_line != nil {
		barlineTo.Wavy_line = CopyBranchWavy_line(mapOrigCopy, barlineFrom.Wavy_line)
	}
	if barlineFrom.Fermata != nil {
		barlineTo.Fermata = CopyBranchFermata(mapOrigCopy, barlineFrom.Fermata)
	}
	if barlineFrom.Ending != nil {
		barlineTo.Ending = CopyBranchEnding(mapOrigCopy, barlineFrom.Ending)
	}
	if barlineFrom.Repeat != nil {
		barlineTo.Repeat = CopyBranchRepeat(mapOrigCopy, barlineFrom.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBarre(mapOrigCopy map[any]any, barreFrom *Barre) (barreTo *Barre) {

	// barreFrom has already been copied
	if _barreTo, ok := mapOrigCopy[barreFrom]; ok {
		barreTo = _barreTo.(*Barre)
		return
	}

	barreTo = new(Barre)
	mapOrigCopy[barreFrom] = barreTo
	barreFrom.CopyBasicFields(barreTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBass(mapOrigCopy map[any]any, bassFrom *Bass) (bassTo *Bass) {

	// bassFrom has already been copied
	if _bassTo, ok := mapOrigCopy[bassFrom]; ok {
		bassTo = _bassTo.(*Bass)
		return
	}

	bassTo = new(Bass)
	mapOrigCopy[bassFrom] = bassTo
	bassFrom.CopyBasicFields(bassTo)

	//insertion point for the staging of instances referenced by pointers
	if bassFrom.Bass_step != nil {
		bassTo.Bass_step = CopyBranchBass_step(mapOrigCopy, bassFrom.Bass_step)
	}
	if bassFrom.Bass_alter != nil {
		bassTo.Bass_alter = CopyBranchHarmony_alter(mapOrigCopy, bassFrom.Bass_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBass_step(mapOrigCopy map[any]any, bass_stepFrom *Bass_step) (bass_stepTo *Bass_step) {

	// bass_stepFrom has already been copied
	if _bass_stepTo, ok := mapOrigCopy[bass_stepFrom]; ok {
		bass_stepTo = _bass_stepTo.(*Bass_step)
		return
	}

	bass_stepTo = new(Bass_step)
	mapOrigCopy[bass_stepFrom] = bass_stepTo
	bass_stepFrom.CopyBasicFields(bass_stepTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBeam(mapOrigCopy map[any]any, beamFrom *Beam) (beamTo *Beam) {

	// beamFrom has already been copied
	if _beamTo, ok := mapOrigCopy[beamFrom]; ok {
		beamTo = _beamTo.(*Beam)
		return
	}

	beamTo = new(Beam)
	mapOrigCopy[beamFrom] = beamTo
	beamFrom.CopyBasicFields(beamTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBeat_repeat(mapOrigCopy map[any]any, beat_repeatFrom *Beat_repeat) (beat_repeatTo *Beat_repeat) {

	// beat_repeatFrom has already been copied
	if _beat_repeatTo, ok := mapOrigCopy[beat_repeatFrom]; ok {
		beat_repeatTo = _beat_repeatTo.(*Beat_repeat)
		return
	}

	beat_repeatTo = new(Beat_repeat)
	mapOrigCopy[beat_repeatFrom] = beat_repeatTo
	beat_repeatFrom.CopyBasicFields(beat_repeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBeat_unit_tied(mapOrigCopy map[any]any, beat_unit_tiedFrom *Beat_unit_tied) (beat_unit_tiedTo *Beat_unit_tied) {

	// beat_unit_tiedFrom has already been copied
	if _beat_unit_tiedTo, ok := mapOrigCopy[beat_unit_tiedFrom]; ok {
		beat_unit_tiedTo = _beat_unit_tiedTo.(*Beat_unit_tied)
		return
	}

	beat_unit_tiedTo = new(Beat_unit_tied)
	mapOrigCopy[beat_unit_tiedFrom] = beat_unit_tiedTo
	beat_unit_tiedFrom.CopyBasicFields(beat_unit_tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBeater(mapOrigCopy map[any]any, beaterFrom *Beater) (beaterTo *Beater) {

	// beaterFrom has already been copied
	if _beaterTo, ok := mapOrigCopy[beaterFrom]; ok {
		beaterTo = _beaterTo.(*Beater)
		return
	}

	beaterTo = new(Beater)
	mapOrigCopy[beaterFrom] = beaterTo
	beaterFrom.CopyBasicFields(beaterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBend(mapOrigCopy map[any]any, bendFrom *Bend) (bendTo *Bend) {

	// bendFrom has already been copied
	if _bendTo, ok := mapOrigCopy[bendFrom]; ok {
		bendTo = _bendTo.(*Bend)
		return
	}

	bendTo = new(Bend)
	mapOrigCopy[bendFrom] = bendTo
	bendFrom.CopyBasicFields(bendTo)

	//insertion point for the staging of instances referenced by pointers
	if bendFrom.Pre_bend != nil {
		bendTo.Pre_bend = CopyBranchEmpty(mapOrigCopy, bendFrom.Pre_bend)
	}
	if bendFrom.Release != nil {
		bendTo.Release = CopyBranchRelease(mapOrigCopy, bendFrom.Release)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBookmark(mapOrigCopy map[any]any, bookmarkFrom *Bookmark) (bookmarkTo *Bookmark) {

	// bookmarkFrom has already been copied
	if _bookmarkTo, ok := mapOrigCopy[bookmarkFrom]; ok {
		bookmarkTo = _bookmarkTo.(*Bookmark)
		return
	}

	bookmarkTo = new(Bookmark)
	mapOrigCopy[bookmarkFrom] = bookmarkTo
	bookmarkFrom.CopyBasicFields(bookmarkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBracket(mapOrigCopy map[any]any, bracketFrom *Bracket) (bracketTo *Bracket) {

	// bracketFrom has already been copied
	if _bracketTo, ok := mapOrigCopy[bracketFrom]; ok {
		bracketTo = _bracketTo.(*Bracket)
		return
	}

	bracketTo = new(Bracket)
	mapOrigCopy[bracketFrom] = bracketTo
	bracketFrom.CopyBasicFields(bracketTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchBreath_mark(mapOrigCopy map[any]any, breath_markFrom *Breath_mark) (breath_markTo *Breath_mark) {

	// breath_markFrom has already been copied
	if _breath_markTo, ok := mapOrigCopy[breath_markFrom]; ok {
		breath_markTo = _breath_markTo.(*Breath_mark)
		return
	}

	breath_markTo = new(Breath_mark)
	mapOrigCopy[breath_markFrom] = breath_markTo
	breath_markFrom.CopyBasicFields(breath_markTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCaesura(mapOrigCopy map[any]any, caesuraFrom *Caesura) (caesuraTo *Caesura) {

	// caesuraFrom has already been copied
	if _caesuraTo, ok := mapOrigCopy[caesuraFrom]; ok {
		caesuraTo = _caesuraTo.(*Caesura)
		return
	}

	caesuraTo = new(Caesura)
	mapOrigCopy[caesuraFrom] = caesuraTo
	caesuraFrom.CopyBasicFields(caesuraTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCancel(mapOrigCopy map[any]any, cancelFrom *Cancel) (cancelTo *Cancel) {

	// cancelFrom has already been copied
	if _cancelTo, ok := mapOrigCopy[cancelFrom]; ok {
		cancelTo = _cancelTo.(*Cancel)
		return
	}

	cancelTo = new(Cancel)
	mapOrigCopy[cancelFrom] = cancelTo
	cancelFrom.CopyBasicFields(cancelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchClef(mapOrigCopy map[any]any, clefFrom *Clef) (clefTo *Clef) {

	// clefFrom has already been copied
	if _clefTo, ok := mapOrigCopy[clefFrom]; ok {
		clefTo = _clefTo.(*Clef)
		return
	}

	clefTo = new(Clef)
	mapOrigCopy[clefFrom] = clefTo
	clefFrom.CopyBasicFields(clefTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCoda(mapOrigCopy map[any]any, codaFrom *Coda) (codaTo *Coda) {

	// codaFrom has already been copied
	if _codaTo, ok := mapOrigCopy[codaFrom]; ok {
		codaTo = _codaTo.(*Coda)
		return
	}

	codaTo = new(Coda)
	mapOrigCopy[codaFrom] = codaTo
	codaFrom.CopyBasicFields(codaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {

	// creditFrom has already been copied
	if _creditTo, ok := mapOrigCopy[creditFrom]; ok {
		creditTo = _creditTo.(*Credit)
		return
	}

	creditTo = new(Credit)
	mapOrigCopy[creditFrom] = creditTo
	creditFrom.CopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers
	if creditFrom.Credit_image != nil {
		creditTo.Credit_image = CopyBranchImage(mapOrigCopy, creditFrom.Credit_image)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range creditFrom.Link {
		creditTo.Link = append(creditTo.Link, CopyBranchLink(mapOrigCopy, _link))
	}
	for _, _bookmark := range creditFrom.Bookmark {
		creditTo.Bookmark = append(creditTo.Bookmark, CopyBranchBookmark(mapOrigCopy, _bookmark))
	}

	return
}

func CopyBranchDashes(mapOrigCopy map[any]any, dashesFrom *Dashes) (dashesTo *Dashes) {

	// dashesFrom has already been copied
	if _dashesTo, ok := mapOrigCopy[dashesFrom]; ok {
		dashesTo = _dashesTo.(*Dashes)
		return
	}

	dashesTo = new(Dashes)
	mapOrigCopy[dashesFrom] = dashesTo
	dashesFrom.CopyBasicFields(dashesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDefaults(mapOrigCopy map[any]any, defaultsFrom *Defaults) (defaultsTo *Defaults) {

	// defaultsFrom has already been copied
	if _defaultsTo, ok := mapOrigCopy[defaultsFrom]; ok {
		defaultsTo = _defaultsTo.(*Defaults)
		return
	}

	defaultsTo = new(Defaults)
	mapOrigCopy[defaultsFrom] = defaultsTo
	defaultsFrom.CopyBasicFields(defaultsTo)

	//insertion point for the staging of instances referenced by pointers
	if defaultsFrom.Scaling != nil {
		defaultsTo.Scaling = CopyBranchScaling(mapOrigCopy, defaultsFrom.Scaling)
	}
	if defaultsFrom.Concert_score != nil {
		defaultsTo.Concert_score = CopyBranchEmpty(mapOrigCopy, defaultsFrom.Concert_score)
	}
	if defaultsFrom.Appearance != nil {
		defaultsTo.Appearance = CopyBranchAppearance(mapOrigCopy, defaultsFrom.Appearance)
	}
	if defaultsFrom.Music_font != nil {
		defaultsTo.Music_font = CopyBranchEmpty_font(mapOrigCopy, defaultsFrom.Music_font)
	}
	if defaultsFrom.Word_font != nil {
		defaultsTo.Word_font = CopyBranchEmpty_font(mapOrigCopy, defaultsFrom.Word_font)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lyric_font := range defaultsFrom.Lyric_font {
		defaultsTo.Lyric_font = append(defaultsTo.Lyric_font, CopyBranchLyric_font(mapOrigCopy, _lyric_font))
	}
	for _, _lyric_language := range defaultsFrom.Lyric_language {
		defaultsTo.Lyric_language = append(defaultsTo.Lyric_language, CopyBranchLyric_language(mapOrigCopy, _lyric_language))
	}

	return
}

func CopyBranchDegree(mapOrigCopy map[any]any, degreeFrom *Degree) (degreeTo *Degree) {

	// degreeFrom has already been copied
	if _degreeTo, ok := mapOrigCopy[degreeFrom]; ok {
		degreeTo = _degreeTo.(*Degree)
		return
	}

	degreeTo = new(Degree)
	mapOrigCopy[degreeFrom] = degreeTo
	degreeFrom.CopyBasicFields(degreeTo)

	//insertion point for the staging of instances referenced by pointers
	if degreeFrom.Degree_value != nil {
		degreeTo.Degree_value = CopyBranchDegree_value(mapOrigCopy, degreeFrom.Degree_value)
	}
	if degreeFrom.Degree_alter != nil {
		degreeTo.Degree_alter = CopyBranchDegree_alter(mapOrigCopy, degreeFrom.Degree_alter)
	}
	if degreeFrom.Degree_type != nil {
		degreeTo.Degree_type = CopyBranchDegree_type(mapOrigCopy, degreeFrom.Degree_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDegree_alter(mapOrigCopy map[any]any, degree_alterFrom *Degree_alter) (degree_alterTo *Degree_alter) {

	// degree_alterFrom has already been copied
	if _degree_alterTo, ok := mapOrigCopy[degree_alterFrom]; ok {
		degree_alterTo = _degree_alterTo.(*Degree_alter)
		return
	}

	degree_alterTo = new(Degree_alter)
	mapOrigCopy[degree_alterFrom] = degree_alterTo
	degree_alterFrom.CopyBasicFields(degree_alterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDegree_type(mapOrigCopy map[any]any, degree_typeFrom *Degree_type) (degree_typeTo *Degree_type) {

	// degree_typeFrom has already been copied
	if _degree_typeTo, ok := mapOrigCopy[degree_typeFrom]; ok {
		degree_typeTo = _degree_typeTo.(*Degree_type)
		return
	}

	degree_typeTo = new(Degree_type)
	mapOrigCopy[degree_typeFrom] = degree_typeTo
	degree_typeFrom.CopyBasicFields(degree_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDegree_value(mapOrigCopy map[any]any, degree_valueFrom *Degree_value) (degree_valueTo *Degree_value) {

	// degree_valueFrom has already been copied
	if _degree_valueTo, ok := mapOrigCopy[degree_valueFrom]; ok {
		degree_valueTo = _degree_valueTo.(*Degree_value)
		return
	}

	degree_valueTo = new(Degree_value)
	mapOrigCopy[degree_valueFrom] = degree_valueTo
	degree_valueFrom.CopyBasicFields(degree_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDirection(mapOrigCopy map[any]any, directionFrom *Direction) (directionTo *Direction) {

	// directionFrom has already been copied
	if _directionTo, ok := mapOrigCopy[directionFrom]; ok {
		directionTo = _directionTo.(*Direction)
		return
	}

	directionTo = new(Direction)
	mapOrigCopy[directionFrom] = directionTo
	directionFrom.CopyBasicFields(directionTo)

	//insertion point for the staging of instances referenced by pointers
	if directionFrom.Offset != nil {
		directionTo.Offset = CopyBranchOffset(mapOrigCopy, directionFrom.Offset)
	}
	if directionFrom.Sound != nil {
		directionTo.Sound = CopyBranchSound(mapOrigCopy, directionFrom.Sound)
	}
	if directionFrom.Listening != nil {
		directionTo.Listening = CopyBranchListening(mapOrigCopy, directionFrom.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _direction_type := range directionFrom.Direction_type {
		directionTo.Direction_type = append(directionTo.Direction_type, CopyBranchDirection_type(mapOrigCopy, _direction_type))
	}

	return
}

func CopyBranchDirection_type(mapOrigCopy map[any]any, direction_typeFrom *Direction_type) (direction_typeTo *Direction_type) {

	// direction_typeFrom has already been copied
	if _direction_typeTo, ok := mapOrigCopy[direction_typeFrom]; ok {
		direction_typeTo = _direction_typeTo.(*Direction_type)
		return
	}

	direction_typeTo = new(Direction_type)
	mapOrigCopy[direction_typeFrom] = direction_typeTo
	direction_typeFrom.CopyBasicFields(direction_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if direction_typeFrom.Wedge != nil {
		direction_typeTo.Wedge = CopyBranchWedge(mapOrigCopy, direction_typeFrom.Wedge)
	}
	if direction_typeFrom.Dashes != nil {
		direction_typeTo.Dashes = CopyBranchDashes(mapOrigCopy, direction_typeFrom.Dashes)
	}
	if direction_typeFrom.Bracket != nil {
		direction_typeTo.Bracket = CopyBranchBracket(mapOrigCopy, direction_typeFrom.Bracket)
	}
	if direction_typeFrom.Pedal != nil {
		direction_typeTo.Pedal = CopyBranchPedal(mapOrigCopy, direction_typeFrom.Pedal)
	}
	if direction_typeFrom.Metronome != nil {
		direction_typeTo.Metronome = CopyBranchMetronome(mapOrigCopy, direction_typeFrom.Metronome)
	}
	if direction_typeFrom.Octave_shift != nil {
		direction_typeTo.Octave_shift = CopyBranchOctave_shift(mapOrigCopy, direction_typeFrom.Octave_shift)
	}
	if direction_typeFrom.Harp_pedals != nil {
		direction_typeTo.Harp_pedals = CopyBranchHarp_pedals(mapOrigCopy, direction_typeFrom.Harp_pedals)
	}
	if direction_typeFrom.Damp != nil {
		direction_typeTo.Damp = CopyBranchEmpty_print_style_align_id(mapOrigCopy, direction_typeFrom.Damp)
	}
	if direction_typeFrom.Damp_all != nil {
		direction_typeTo.Damp_all = CopyBranchEmpty_print_style_align_id(mapOrigCopy, direction_typeFrom.Damp_all)
	}
	if direction_typeFrom.Eyeglasses != nil {
		direction_typeTo.Eyeglasses = CopyBranchEmpty_print_style_align_id(mapOrigCopy, direction_typeFrom.Eyeglasses)
	}
	if direction_typeFrom.String_mute != nil {
		direction_typeTo.String_mute = CopyBranchString_mute(mapOrigCopy, direction_typeFrom.String_mute)
	}
	if direction_typeFrom.Scordatura != nil {
		direction_typeTo.Scordatura = CopyBranchScordatura(mapOrigCopy, direction_typeFrom.Scordatura)
	}
	if direction_typeFrom.Image != nil {
		direction_typeTo.Image = CopyBranchImage(mapOrigCopy, direction_typeFrom.Image)
	}
	if direction_typeFrom.Principal_voice != nil {
		direction_typeTo.Principal_voice = CopyBranchPrincipal_voice(mapOrigCopy, direction_typeFrom.Principal_voice)
	}
	if direction_typeFrom.Accordion_registration != nil {
		direction_typeTo.Accordion_registration = CopyBranchAccordion_registration(mapOrigCopy, direction_typeFrom.Accordion_registration)
	}
	if direction_typeFrom.Staff_divide != nil {
		direction_typeTo.Staff_divide = CopyBranchStaff_divide(mapOrigCopy, direction_typeFrom.Staff_divide)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _segno := range direction_typeFrom.Segno {
		direction_typeTo.Segno = append(direction_typeTo.Segno, CopyBranchSegno(mapOrigCopy, _segno))
	}
	for _, _coda := range direction_typeFrom.Coda {
		direction_typeTo.Coda = append(direction_typeTo.Coda, CopyBranchCoda(mapOrigCopy, _coda))
	}
	for _, _dynamics := range direction_typeFrom.Dynamics {
		direction_typeTo.Dynamics = append(direction_typeTo.Dynamics, CopyBranchDynamics(mapOrigCopy, _dynamics))
	}
	for _, _percussion := range direction_typeFrom.Percussion {
		direction_typeTo.Percussion = append(direction_typeTo.Percussion, CopyBranchPercussion(mapOrigCopy, _percussion))
	}

	return
}

func CopyBranchDistance(mapOrigCopy map[any]any, distanceFrom *Distance) (distanceTo *Distance) {

	// distanceFrom has already been copied
	if _distanceTo, ok := mapOrigCopy[distanceFrom]; ok {
		distanceTo = _distanceTo.(*Distance)
		return
	}

	distanceTo = new(Distance)
	mapOrigCopy[distanceFrom] = distanceTo
	distanceFrom.CopyBasicFields(distanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDouble(mapOrigCopy map[any]any, doubleFrom *Double) (doubleTo *Double) {

	// doubleFrom has already been copied
	if _doubleTo, ok := mapOrigCopy[doubleFrom]; ok {
		doubleTo = _doubleTo.(*Double)
		return
	}

	doubleTo = new(Double)
	mapOrigCopy[doubleFrom] = doubleTo
	doubleFrom.CopyBasicFields(doubleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDynamics(mapOrigCopy map[any]any, dynamicsFrom *Dynamics) (dynamicsTo *Dynamics) {

	// dynamicsFrom has already been copied
	if _dynamicsTo, ok := mapOrigCopy[dynamicsFrom]; ok {
		dynamicsTo = _dynamicsTo.(*Dynamics)
		return
	}

	dynamicsTo = new(Dynamics)
	mapOrigCopy[dynamicsFrom] = dynamicsTo
	dynamicsFrom.CopyBasicFields(dynamicsTo)

	//insertion point for the staging of instances referenced by pointers
	if dynamicsFrom.P != nil {
		dynamicsTo.P = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.P)
	}
	if dynamicsFrom.Pp != nil {
		dynamicsTo.Pp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Pp)
	}
	if dynamicsFrom.Ppp != nil {
		dynamicsTo.Ppp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Ppp)
	}
	if dynamicsFrom.Pppp != nil {
		dynamicsTo.Pppp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Pppp)
	}
	if dynamicsFrom.Ppppp != nil {
		dynamicsTo.Ppppp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Ppppp)
	}
	if dynamicsFrom.Pppppp != nil {
		dynamicsTo.Pppppp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Pppppp)
	}
	if dynamicsFrom.F != nil {
		dynamicsTo.F = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.F)
	}
	if dynamicsFrom.Ff != nil {
		dynamicsTo.Ff = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Ff)
	}
	if dynamicsFrom.Fff != nil {
		dynamicsTo.Fff = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Fff)
	}
	if dynamicsFrom.Ffff != nil {
		dynamicsTo.Ffff = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Ffff)
	}
	if dynamicsFrom.Fffff != nil {
		dynamicsTo.Fffff = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Fffff)
	}
	if dynamicsFrom.Ffffff != nil {
		dynamicsTo.Ffffff = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Ffffff)
	}
	if dynamicsFrom.Mp != nil {
		dynamicsTo.Mp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Mp)
	}
	if dynamicsFrom.Mf != nil {
		dynamicsTo.Mf = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Mf)
	}
	if dynamicsFrom.Sf != nil {
		dynamicsTo.Sf = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Sf)
	}
	if dynamicsFrom.Sfp != nil {
		dynamicsTo.Sfp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Sfp)
	}
	if dynamicsFrom.Sfpp != nil {
		dynamicsTo.Sfpp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Sfpp)
	}
	if dynamicsFrom.Fp != nil {
		dynamicsTo.Fp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Fp)
	}
	if dynamicsFrom.Rf != nil {
		dynamicsTo.Rf = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Rf)
	}
	if dynamicsFrom.Rfz != nil {
		dynamicsTo.Rfz = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Rfz)
	}
	if dynamicsFrom.Sfz != nil {
		dynamicsTo.Sfz = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Sfz)
	}
	if dynamicsFrom.Sffz != nil {
		dynamicsTo.Sffz = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Sffz)
	}
	if dynamicsFrom.Fz != nil {
		dynamicsTo.Fz = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Fz)
	}
	if dynamicsFrom.N != nil {
		dynamicsTo.N = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.N)
	}
	if dynamicsFrom.Pf != nil {
		dynamicsTo.Pf = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Pf)
	}
	if dynamicsFrom.Sfzp != nil {
		dynamicsTo.Sfzp = CopyBranchEmpty(mapOrigCopy, dynamicsFrom.Sfzp)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEffect(mapOrigCopy map[any]any, effectFrom *Effect) (effectTo *Effect) {

	// effectFrom has already been copied
	if _effectTo, ok := mapOrigCopy[effectFrom]; ok {
		effectTo = _effectTo.(*Effect)
		return
	}

	effectTo = new(Effect)
	mapOrigCopy[effectFrom] = effectTo
	effectFrom.CopyBasicFields(effectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchElision(mapOrigCopy map[any]any, elisionFrom *Elision) (elisionTo *Elision) {

	// elisionFrom has already been copied
	if _elisionTo, ok := mapOrigCopy[elisionFrom]; ok {
		elisionTo = _elisionTo.(*Elision)
		return
	}

	elisionTo = new(Elision)
	mapOrigCopy[elisionFrom] = elisionTo
	elisionFrom.CopyBasicFields(elisionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty(mapOrigCopy map[any]any, emptyFrom *Empty) (emptyTo *Empty) {

	// emptyFrom has already been copied
	if _emptyTo, ok := mapOrigCopy[emptyFrom]; ok {
		emptyTo = _emptyTo.(*Empty)
		return
	}

	emptyTo = new(Empty)
	mapOrigCopy[emptyFrom] = emptyTo
	emptyFrom.CopyBasicFields(emptyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_font(mapOrigCopy map[any]any, empty_fontFrom *Empty_font) (empty_fontTo *Empty_font) {

	// empty_fontFrom has already been copied
	if _empty_fontTo, ok := mapOrigCopy[empty_fontFrom]; ok {
		empty_fontTo = _empty_fontTo.(*Empty_font)
		return
	}

	empty_fontTo = new(Empty_font)
	mapOrigCopy[empty_fontFrom] = empty_fontTo
	empty_fontFrom.CopyBasicFields(empty_fontTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_line(mapOrigCopy map[any]any, empty_lineFrom *Empty_line) (empty_lineTo *Empty_line) {

	// empty_lineFrom has already been copied
	if _empty_lineTo, ok := mapOrigCopy[empty_lineFrom]; ok {
		empty_lineTo = _empty_lineTo.(*Empty_line)
		return
	}

	empty_lineTo = new(Empty_line)
	mapOrigCopy[empty_lineFrom] = empty_lineTo
	empty_lineFrom.CopyBasicFields(empty_lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_placement(mapOrigCopy map[any]any, empty_placementFrom *Empty_placement) (empty_placementTo *Empty_placement) {

	// empty_placementFrom has already been copied
	if _empty_placementTo, ok := mapOrigCopy[empty_placementFrom]; ok {
		empty_placementTo = _empty_placementTo.(*Empty_placement)
		return
	}

	empty_placementTo = new(Empty_placement)
	mapOrigCopy[empty_placementFrom] = empty_placementTo
	empty_placementFrom.CopyBasicFields(empty_placementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_placement_smufl(mapOrigCopy map[any]any, empty_placement_smuflFrom *Empty_placement_smufl) (empty_placement_smuflTo *Empty_placement_smufl) {

	// empty_placement_smuflFrom has already been copied
	if _empty_placement_smuflTo, ok := mapOrigCopy[empty_placement_smuflFrom]; ok {
		empty_placement_smuflTo = _empty_placement_smuflTo.(*Empty_placement_smufl)
		return
	}

	empty_placement_smuflTo = new(Empty_placement_smufl)
	mapOrigCopy[empty_placement_smuflFrom] = empty_placement_smuflTo
	empty_placement_smuflFrom.CopyBasicFields(empty_placement_smuflTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_print_object_style_align(mapOrigCopy map[any]any, empty_print_object_style_alignFrom *Empty_print_object_style_align) (empty_print_object_style_alignTo *Empty_print_object_style_align) {

	// empty_print_object_style_alignFrom has already been copied
	if _empty_print_object_style_alignTo, ok := mapOrigCopy[empty_print_object_style_alignFrom]; ok {
		empty_print_object_style_alignTo = _empty_print_object_style_alignTo.(*Empty_print_object_style_align)
		return
	}

	empty_print_object_style_alignTo = new(Empty_print_object_style_align)
	mapOrigCopy[empty_print_object_style_alignFrom] = empty_print_object_style_alignTo
	empty_print_object_style_alignFrom.CopyBasicFields(empty_print_object_style_alignTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_print_style(mapOrigCopy map[any]any, empty_print_styleFrom *Empty_print_style) (empty_print_styleTo *Empty_print_style) {

	// empty_print_styleFrom has already been copied
	if _empty_print_styleTo, ok := mapOrigCopy[empty_print_styleFrom]; ok {
		empty_print_styleTo = _empty_print_styleTo.(*Empty_print_style)
		return
	}

	empty_print_styleTo = new(Empty_print_style)
	mapOrigCopy[empty_print_styleFrom] = empty_print_styleTo
	empty_print_styleFrom.CopyBasicFields(empty_print_styleTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_print_style_align(mapOrigCopy map[any]any, empty_print_style_alignFrom *Empty_print_style_align) (empty_print_style_alignTo *Empty_print_style_align) {

	// empty_print_style_alignFrom has already been copied
	if _empty_print_style_alignTo, ok := mapOrigCopy[empty_print_style_alignFrom]; ok {
		empty_print_style_alignTo = _empty_print_style_alignTo.(*Empty_print_style_align)
		return
	}

	empty_print_style_alignTo = new(Empty_print_style_align)
	mapOrigCopy[empty_print_style_alignFrom] = empty_print_style_alignTo
	empty_print_style_alignFrom.CopyBasicFields(empty_print_style_alignTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_print_style_align_id(mapOrigCopy map[any]any, empty_print_style_align_idFrom *Empty_print_style_align_id) (empty_print_style_align_idTo *Empty_print_style_align_id) {

	// empty_print_style_align_idFrom has already been copied
	if _empty_print_style_align_idTo, ok := mapOrigCopy[empty_print_style_align_idFrom]; ok {
		empty_print_style_align_idTo = _empty_print_style_align_idTo.(*Empty_print_style_align_id)
		return
	}

	empty_print_style_align_idTo = new(Empty_print_style_align_id)
	mapOrigCopy[empty_print_style_align_idFrom] = empty_print_style_align_idTo
	empty_print_style_align_idFrom.CopyBasicFields(empty_print_style_align_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEmpty_trill_sound(mapOrigCopy map[any]any, empty_trill_soundFrom *Empty_trill_sound) (empty_trill_soundTo *Empty_trill_sound) {

	// empty_trill_soundFrom has already been copied
	if _empty_trill_soundTo, ok := mapOrigCopy[empty_trill_soundFrom]; ok {
		empty_trill_soundTo = _empty_trill_soundTo.(*Empty_trill_sound)
		return
	}

	empty_trill_soundTo = new(Empty_trill_sound)
	mapOrigCopy[empty_trill_soundFrom] = empty_trill_soundTo
	empty_trill_soundFrom.CopyBasicFields(empty_trill_soundTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEncoding(mapOrigCopy map[any]any, encodingFrom *Encoding) (encodingTo *Encoding) {

	// encodingFrom has already been copied
	if _encodingTo, ok := mapOrigCopy[encodingFrom]; ok {
		encodingTo = _encodingTo.(*Encoding)
		return
	}

	encodingTo = new(Encoding)
	mapOrigCopy[encodingFrom] = encodingTo
	encodingFrom.CopyBasicFields(encodingTo)

	//insertion point for the staging of instances referenced by pointers
	if encodingFrom.Encoder != nil {
		encodingTo.Encoder = CopyBranchTyped_text(mapOrigCopy, encodingFrom.Encoder)
	}
	if encodingFrom.Supports != nil {
		encodingTo.Supports = CopyBranchSupports(mapOrigCopy, encodingFrom.Supports)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchEnding(mapOrigCopy map[any]any, endingFrom *Ending) (endingTo *Ending) {

	// endingFrom has already been copied
	if _endingTo, ok := mapOrigCopy[endingFrom]; ok {
		endingTo = _endingTo.(*Ending)
		return
	}

	endingTo = new(Ending)
	mapOrigCopy[endingFrom] = endingTo
	endingFrom.CopyBasicFields(endingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchExtend(mapOrigCopy map[any]any, extendFrom *Extend) (extendTo *Extend) {

	// extendFrom has already been copied
	if _extendTo, ok := mapOrigCopy[extendFrom]; ok {
		extendTo = _extendTo.(*Extend)
		return
	}

	extendTo = new(Extend)
	mapOrigCopy[extendFrom] = extendTo
	extendFrom.CopyBasicFields(extendTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFeature(mapOrigCopy map[any]any, featureFrom *Feature) (featureTo *Feature) {

	// featureFrom has already been copied
	if _featureTo, ok := mapOrigCopy[featureFrom]; ok {
		featureTo = _featureTo.(*Feature)
		return
	}

	featureTo = new(Feature)
	mapOrigCopy[featureFrom] = featureTo
	featureFrom.CopyBasicFields(featureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFermata(mapOrigCopy map[any]any, fermataFrom *Fermata) (fermataTo *Fermata) {

	// fermataFrom has already been copied
	if _fermataTo, ok := mapOrigCopy[fermataFrom]; ok {
		fermataTo = _fermataTo.(*Fermata)
		return
	}

	fermataTo = new(Fermata)
	mapOrigCopy[fermataFrom] = fermataTo
	fermataFrom.CopyBasicFields(fermataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFigure(mapOrigCopy map[any]any, figureFrom *Figure) (figureTo *Figure) {

	// figureFrom has already been copied
	if _figureTo, ok := mapOrigCopy[figureFrom]; ok {
		figureTo = _figureTo.(*Figure)
		return
	}

	figureTo = new(Figure)
	mapOrigCopy[figureFrom] = figureTo
	figureFrom.CopyBasicFields(figureTo)

	//insertion point for the staging of instances referenced by pointers
	if figureFrom.Extend != nil {
		figureTo.Extend = CopyBranchExtend(mapOrigCopy, figureFrom.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFigured_bass(mapOrigCopy map[any]any, figured_bassFrom *Figured_bass) (figured_bassTo *Figured_bass) {

	// figured_bassFrom has already been copied
	if _figured_bassTo, ok := mapOrigCopy[figured_bassFrom]; ok {
		figured_bassTo = _figured_bassTo.(*Figured_bass)
		return
	}

	figured_bassTo = new(Figured_bass)
	mapOrigCopy[figured_bassFrom] = figured_bassTo
	figured_bassFrom.CopyBasicFields(figured_bassTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figured_bassFrom.Figure {
		figured_bassTo.Figure = append(figured_bassTo.Figure, CopyBranchFigure(mapOrigCopy, _figure))
	}

	return
}

func CopyBranchFingering(mapOrigCopy map[any]any, fingeringFrom *Fingering) (fingeringTo *Fingering) {

	// fingeringFrom has already been copied
	if _fingeringTo, ok := mapOrigCopy[fingeringFrom]; ok {
		fingeringTo = _fingeringTo.(*Fingering)
		return
	}

	fingeringTo = new(Fingering)
	mapOrigCopy[fingeringFrom] = fingeringTo
	fingeringFrom.CopyBasicFields(fingeringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFirst_fret(mapOrigCopy map[any]any, first_fretFrom *First_fret) (first_fretTo *First_fret) {

	// first_fretFrom has already been copied
	if _first_fretTo, ok := mapOrigCopy[first_fretFrom]; ok {
		first_fretTo = _first_fretTo.(*First_fret)
		return
	}

	first_fretTo = new(First_fret)
	mapOrigCopy[first_fretFrom] = first_fretTo
	first_fretFrom.CopyBasicFields(first_fretTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFoo(mapOrigCopy map[any]any, fooFrom *Foo) (fooTo *Foo) {

	// fooFrom has already been copied
	if _fooTo, ok := mapOrigCopy[fooFrom]; ok {
		fooTo = _fooTo.(*Foo)
		return
	}

	fooTo = new(Foo)
	mapOrigCopy[fooFrom] = fooTo
	fooFrom.CopyBasicFields(fooTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFor_part(mapOrigCopy map[any]any, for_partFrom *For_part) (for_partTo *For_part) {

	// for_partFrom has already been copied
	if _for_partTo, ok := mapOrigCopy[for_partFrom]; ok {
		for_partTo = _for_partTo.(*For_part)
		return
	}

	for_partTo = new(For_part)
	mapOrigCopy[for_partFrom] = for_partTo
	for_partFrom.CopyBasicFields(for_partTo)

	//insertion point for the staging of instances referenced by pointers
	if for_partFrom.Part_clef != nil {
		for_partTo.Part_clef = CopyBranchPart_clef(mapOrigCopy, for_partFrom.Part_clef)
	}
	if for_partFrom.Part_transpose != nil {
		for_partTo.Part_transpose = CopyBranchPart_transpose(mapOrigCopy, for_partFrom.Part_transpose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormatted_symbol(mapOrigCopy map[any]any, formatted_symbolFrom *Formatted_symbol) (formatted_symbolTo *Formatted_symbol) {

	// formatted_symbolFrom has already been copied
	if _formatted_symbolTo, ok := mapOrigCopy[formatted_symbolFrom]; ok {
		formatted_symbolTo = _formatted_symbolTo.(*Formatted_symbol)
		return
	}

	formatted_symbolTo = new(Formatted_symbol)
	mapOrigCopy[formatted_symbolFrom] = formatted_symbolTo
	formatted_symbolFrom.CopyBasicFields(formatted_symbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFormatted_symbol_id(mapOrigCopy map[any]any, formatted_symbol_idFrom *Formatted_symbol_id) (formatted_symbol_idTo *Formatted_symbol_id) {

	// formatted_symbol_idFrom has already been copied
	if _formatted_symbol_idTo, ok := mapOrigCopy[formatted_symbol_idFrom]; ok {
		formatted_symbol_idTo = _formatted_symbol_idTo.(*Formatted_symbol_id)
		return
	}

	formatted_symbol_idTo = new(Formatted_symbol_id)
	mapOrigCopy[formatted_symbol_idFrom] = formatted_symbol_idTo
	formatted_symbol_idFrom.CopyBasicFields(formatted_symbol_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchForward(mapOrigCopy map[any]any, forwardFrom *Forward) (forwardTo *Forward) {

	// forwardFrom has already been copied
	if _forwardTo, ok := mapOrigCopy[forwardFrom]; ok {
		forwardTo = _forwardTo.(*Forward)
		return
	}

	forwardTo = new(Forward)
	mapOrigCopy[forwardFrom] = forwardTo
	forwardFrom.CopyBasicFields(forwardTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFrame(mapOrigCopy map[any]any, frameFrom *Frame) (frameTo *Frame) {

	// frameFrom has already been copied
	if _frameTo, ok := mapOrigCopy[frameFrom]; ok {
		frameTo = _frameTo.(*Frame)
		return
	}

	frameTo = new(Frame)
	mapOrigCopy[frameFrom] = frameTo
	frameFrom.CopyBasicFields(frameTo)

	//insertion point for the staging of instances referenced by pointers
	if frameFrom.First_fret != nil {
		frameTo.First_fret = CopyBranchFirst_fret(mapOrigCopy, frameFrom.First_fret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frame_note := range frameFrom.Frame_note {
		frameTo.Frame_note = append(frameTo.Frame_note, CopyBranchFrame_note(mapOrigCopy, _frame_note))
	}

	return
}

func CopyBranchFrame_note(mapOrigCopy map[any]any, frame_noteFrom *Frame_note) (frame_noteTo *Frame_note) {

	// frame_noteFrom has already been copied
	if _frame_noteTo, ok := mapOrigCopy[frame_noteFrom]; ok {
		frame_noteTo = _frame_noteTo.(*Frame_note)
		return
	}

	frame_noteTo = new(Frame_note)
	mapOrigCopy[frame_noteFrom] = frame_noteTo
	frame_noteFrom.CopyBasicFields(frame_noteTo)

	//insertion point for the staging of instances referenced by pointers
	if frame_noteFrom.Fret != nil {
		frame_noteTo.Fret = CopyBranchFret(mapOrigCopy, frame_noteFrom.Fret)
	}
	if frame_noteFrom.Fingering != nil {
		frame_noteTo.Fingering = CopyBranchFingering(mapOrigCopy, frame_noteFrom.Fingering)
	}
	if frame_noteFrom.Barre != nil {
		frame_noteTo.Barre = CopyBranchBarre(mapOrigCopy, frame_noteFrom.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchFret(mapOrigCopy map[any]any, fretFrom *Fret) (fretTo *Fret) {

	// fretFrom has already been copied
	if _fretTo, ok := mapOrigCopy[fretFrom]; ok {
		fretTo = _fretTo.(*Fret)
		return
	}

	fretTo = new(Fret)
	mapOrigCopy[fretFrom] = fretTo
	fretFrom.CopyBasicFields(fretTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGlass(mapOrigCopy map[any]any, glassFrom *Glass) (glassTo *Glass) {

	// glassFrom has already been copied
	if _glassTo, ok := mapOrigCopy[glassFrom]; ok {
		glassTo = _glassTo.(*Glass)
		return
	}

	glassTo = new(Glass)
	mapOrigCopy[glassFrom] = glassTo
	glassFrom.CopyBasicFields(glassTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGlissando(mapOrigCopy map[any]any, glissandoFrom *Glissando) (glissandoTo *Glissando) {

	// glissandoFrom has already been copied
	if _glissandoTo, ok := mapOrigCopy[glissandoFrom]; ok {
		glissandoTo = _glissandoTo.(*Glissando)
		return
	}

	glissandoTo = new(Glissando)
	mapOrigCopy[glissandoFrom] = glissandoTo
	glissandoFrom.CopyBasicFields(glissandoTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGlyph(mapOrigCopy map[any]any, glyphFrom *Glyph) (glyphTo *Glyph) {

	// glyphFrom has already been copied
	if _glyphTo, ok := mapOrigCopy[glyphFrom]; ok {
		glyphTo = _glyphTo.(*Glyph)
		return
	}

	glyphTo = new(Glyph)
	mapOrigCopy[glyphFrom] = glyphTo
	glyphFrom.CopyBasicFields(glyphTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrace(mapOrigCopy map[any]any, graceFrom *Grace) (graceTo *Grace) {

	// graceFrom has already been copied
	if _graceTo, ok := mapOrigCopy[graceFrom]; ok {
		graceTo = _graceTo.(*Grace)
		return
	}

	graceTo = new(Grace)
	mapOrigCopy[graceFrom] = graceTo
	graceFrom.CopyBasicFields(graceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGroup_barline(mapOrigCopy map[any]any, group_barlineFrom *Group_barline) (group_barlineTo *Group_barline) {

	// group_barlineFrom has already been copied
	if _group_barlineTo, ok := mapOrigCopy[group_barlineFrom]; ok {
		group_barlineTo = _group_barlineTo.(*Group_barline)
		return
	}

	group_barlineTo = new(Group_barline)
	mapOrigCopy[group_barlineFrom] = group_barlineTo
	group_barlineFrom.CopyBasicFields(group_barlineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGroup_symbol(mapOrigCopy map[any]any, group_symbolFrom *Group_symbol) (group_symbolTo *Group_symbol) {

	// group_symbolFrom has already been copied
	if _group_symbolTo, ok := mapOrigCopy[group_symbolFrom]; ok {
		group_symbolTo = _group_symbolTo.(*Group_symbol)
		return
	}

	group_symbolTo = new(Group_symbol)
	mapOrigCopy[group_symbolFrom] = group_symbolTo
	group_symbolFrom.CopyBasicFields(group_symbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrouping(mapOrigCopy map[any]any, groupingFrom *Grouping) (groupingTo *Grouping) {

	// groupingFrom has already been copied
	if _groupingTo, ok := mapOrigCopy[groupingFrom]; ok {
		groupingTo = _groupingTo.(*Grouping)
		return
	}

	groupingTo = new(Grouping)
	mapOrigCopy[groupingFrom] = groupingTo
	groupingFrom.CopyBasicFields(groupingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range groupingFrom.Feature {
		groupingTo.Feature = append(groupingTo.Feature, CopyBranchFeature(mapOrigCopy, _feature))
	}

	return
}

func CopyBranchHammer_on_pull_off(mapOrigCopy map[any]any, hammer_on_pull_offFrom *Hammer_on_pull_off) (hammer_on_pull_offTo *Hammer_on_pull_off) {

	// hammer_on_pull_offFrom has already been copied
	if _hammer_on_pull_offTo, ok := mapOrigCopy[hammer_on_pull_offFrom]; ok {
		hammer_on_pull_offTo = _hammer_on_pull_offTo.(*Hammer_on_pull_off)
		return
	}

	hammer_on_pull_offTo = new(Hammer_on_pull_off)
	mapOrigCopy[hammer_on_pull_offFrom] = hammer_on_pull_offTo
	hammer_on_pull_offFrom.CopyBasicFields(hammer_on_pull_offTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHandbell(mapOrigCopy map[any]any, handbellFrom *Handbell) (handbellTo *Handbell) {

	// handbellFrom has already been copied
	if _handbellTo, ok := mapOrigCopy[handbellFrom]; ok {
		handbellTo = _handbellTo.(*Handbell)
		return
	}

	handbellTo = new(Handbell)
	mapOrigCopy[handbellFrom] = handbellTo
	handbellFrom.CopyBasicFields(handbellTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHarmon_closed(mapOrigCopy map[any]any, harmon_closedFrom *Harmon_closed) (harmon_closedTo *Harmon_closed) {

	// harmon_closedFrom has already been copied
	if _harmon_closedTo, ok := mapOrigCopy[harmon_closedFrom]; ok {
		harmon_closedTo = _harmon_closedTo.(*Harmon_closed)
		return
	}

	harmon_closedTo = new(Harmon_closed)
	mapOrigCopy[harmon_closedFrom] = harmon_closedTo
	harmon_closedFrom.CopyBasicFields(harmon_closedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHarmon_mute(mapOrigCopy map[any]any, harmon_muteFrom *Harmon_mute) (harmon_muteTo *Harmon_mute) {

	// harmon_muteFrom has already been copied
	if _harmon_muteTo, ok := mapOrigCopy[harmon_muteFrom]; ok {
		harmon_muteTo = _harmon_muteTo.(*Harmon_mute)
		return
	}

	harmon_muteTo = new(Harmon_mute)
	mapOrigCopy[harmon_muteFrom] = harmon_muteTo
	harmon_muteFrom.CopyBasicFields(harmon_muteTo)

	//insertion point for the staging of instances referenced by pointers
	if harmon_muteFrom.Harmon_closed != nil {
		harmon_muteTo.Harmon_closed = CopyBranchHarmon_closed(mapOrigCopy, harmon_muteFrom.Harmon_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHarmonic(mapOrigCopy map[any]any, harmonicFrom *Harmonic) (harmonicTo *Harmonic) {

	// harmonicFrom has already been copied
	if _harmonicTo, ok := mapOrigCopy[harmonicFrom]; ok {
		harmonicTo = _harmonicTo.(*Harmonic)
		return
	}

	harmonicTo = new(Harmonic)
	mapOrigCopy[harmonicFrom] = harmonicTo
	harmonicFrom.CopyBasicFields(harmonicTo)

	//insertion point for the staging of instances referenced by pointers
	if harmonicFrom.Natural != nil {
		harmonicTo.Natural = CopyBranchEmpty(mapOrigCopy, harmonicFrom.Natural)
	}
	if harmonicFrom.Artificial != nil {
		harmonicTo.Artificial = CopyBranchEmpty(mapOrigCopy, harmonicFrom.Artificial)
	}
	if harmonicFrom.Base_pitch != nil {
		harmonicTo.Base_pitch = CopyBranchEmpty(mapOrigCopy, harmonicFrom.Base_pitch)
	}
	if harmonicFrom.Touching_pitch != nil {
		harmonicTo.Touching_pitch = CopyBranchEmpty(mapOrigCopy, harmonicFrom.Touching_pitch)
	}
	if harmonicFrom.Sounding_pitch != nil {
		harmonicTo.Sounding_pitch = CopyBranchEmpty(mapOrigCopy, harmonicFrom.Sounding_pitch)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHarmony(mapOrigCopy map[any]any, harmonyFrom *Harmony) (harmonyTo *Harmony) {

	// harmonyFrom has already been copied
	if _harmonyTo, ok := mapOrigCopy[harmonyFrom]; ok {
		harmonyTo = _harmonyTo.(*Harmony)
		return
	}

	harmonyTo = new(Harmony)
	mapOrigCopy[harmonyFrom] = harmonyTo
	harmonyFrom.CopyBasicFields(harmonyTo)

	//insertion point for the staging of instances referenced by pointers
	if harmonyFrom.Frame != nil {
		harmonyTo.Frame = CopyBranchFrame(mapOrigCopy, harmonyFrom.Frame)
	}
	if harmonyFrom.Offset != nil {
		harmonyTo.Offset = CopyBranchOffset(mapOrigCopy, harmonyFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHarmony_alter(mapOrigCopy map[any]any, harmony_alterFrom *Harmony_alter) (harmony_alterTo *Harmony_alter) {

	// harmony_alterFrom has already been copied
	if _harmony_alterTo, ok := mapOrigCopy[harmony_alterFrom]; ok {
		harmony_alterTo = _harmony_alterTo.(*Harmony_alter)
		return
	}

	harmony_alterTo = new(Harmony_alter)
	mapOrigCopy[harmony_alterFrom] = harmony_alterTo
	harmony_alterFrom.CopyBasicFields(harmony_alterTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHarp_pedals(mapOrigCopy map[any]any, harp_pedalsFrom *Harp_pedals) (harp_pedalsTo *Harp_pedals) {

	// harp_pedalsFrom has already been copied
	if _harp_pedalsTo, ok := mapOrigCopy[harp_pedalsFrom]; ok {
		harp_pedalsTo = _harp_pedalsTo.(*Harp_pedals)
		return
	}

	harp_pedalsTo = new(Harp_pedals)
	mapOrigCopy[harp_pedalsFrom] = harp_pedalsTo
	harp_pedalsFrom.CopyBasicFields(harp_pedalsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedal_tuning := range harp_pedalsFrom.Pedal_tuning {
		harp_pedalsTo.Pedal_tuning = append(harp_pedalsTo.Pedal_tuning, CopyBranchPedal_tuning(mapOrigCopy, _pedal_tuning))
	}

	return
}

func CopyBranchHeel_toe(mapOrigCopy map[any]any, heel_toeFrom *Heel_toe) (heel_toeTo *Heel_toe) {

	// heel_toeFrom has already been copied
	if _heel_toeTo, ok := mapOrigCopy[heel_toeFrom]; ok {
		heel_toeTo = _heel_toeTo.(*Heel_toe)
		return
	}

	heel_toeTo = new(Heel_toe)
	mapOrigCopy[heel_toeFrom] = heel_toeTo
	heel_toeFrom.CopyBasicFields(heel_toeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHole(mapOrigCopy map[any]any, holeFrom *Hole) (holeTo *Hole) {

	// holeFrom has already been copied
	if _holeTo, ok := mapOrigCopy[holeFrom]; ok {
		holeTo = _holeTo.(*Hole)
		return
	}

	holeTo = new(Hole)
	mapOrigCopy[holeFrom] = holeTo
	holeFrom.CopyBasicFields(holeTo)

	//insertion point for the staging of instances referenced by pointers
	if holeFrom.Hole_closed != nil {
		holeTo.Hole_closed = CopyBranchHole_closed(mapOrigCopy, holeFrom.Hole_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHole_closed(mapOrigCopy map[any]any, hole_closedFrom *Hole_closed) (hole_closedTo *Hole_closed) {

	// hole_closedFrom has already been copied
	if _hole_closedTo, ok := mapOrigCopy[hole_closedFrom]; ok {
		hole_closedTo = _hole_closedTo.(*Hole_closed)
		return
	}

	hole_closedTo = new(Hole_closed)
	mapOrigCopy[hole_closedFrom] = hole_closedTo
	hole_closedFrom.CopyBasicFields(hole_closedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchHorizontal_turn(mapOrigCopy map[any]any, horizontal_turnFrom *Horizontal_turn) (horizontal_turnTo *Horizontal_turn) {

	// horizontal_turnFrom has already been copied
	if _horizontal_turnTo, ok := mapOrigCopy[horizontal_turnFrom]; ok {
		horizontal_turnTo = _horizontal_turnTo.(*Horizontal_turn)
		return
	}

	horizontal_turnTo = new(Horizontal_turn)
	mapOrigCopy[horizontal_turnFrom] = horizontal_turnTo
	horizontal_turnFrom.CopyBasicFields(horizontal_turnTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchIdentification(mapOrigCopy map[any]any, identificationFrom *Identification) (identificationTo *Identification) {

	// identificationFrom has already been copied
	if _identificationTo, ok := mapOrigCopy[identificationFrom]; ok {
		identificationTo = _identificationTo.(*Identification)
		return
	}

	identificationTo = new(Identification)
	mapOrigCopy[identificationFrom] = identificationTo
	identificationFrom.CopyBasicFields(identificationTo)

	//insertion point for the staging of instances referenced by pointers
	if identificationFrom.Encoding != nil {
		identificationTo.Encoding = CopyBranchEncoding(mapOrigCopy, identificationFrom.Encoding)
	}
	if identificationFrom.Miscellaneous != nil {
		identificationTo.Miscellaneous = CopyBranchMiscellaneous(mapOrigCopy, identificationFrom.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range identificationFrom.Creator {
		identificationTo.Creator = append(identificationTo.Creator, CopyBranchTyped_text(mapOrigCopy, _typed_text))
	}
	for _, _typed_text := range identificationFrom.Rights {
		identificationTo.Rights = append(identificationTo.Rights, CopyBranchTyped_text(mapOrigCopy, _typed_text))
	}
	for _, _typed_text := range identificationFrom.Relation {
		identificationTo.Relation = append(identificationTo.Relation, CopyBranchTyped_text(mapOrigCopy, _typed_text))
	}

	return
}

func CopyBranchImage(mapOrigCopy map[any]any, imageFrom *Image) (imageTo *Image) {

	// imageFrom has already been copied
	if _imageTo, ok := mapOrigCopy[imageFrom]; ok {
		imageTo = _imageTo.(*Image)
		return
	}

	imageTo = new(Image)
	mapOrigCopy[imageFrom] = imageTo
	imageFrom.CopyBasicFields(imageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInstrument(mapOrigCopy map[any]any, instrumentFrom *Instrument) (instrumentTo *Instrument) {

	// instrumentFrom has already been copied
	if _instrumentTo, ok := mapOrigCopy[instrumentFrom]; ok {
		instrumentTo = _instrumentTo.(*Instrument)
		return
	}

	instrumentTo = new(Instrument)
	mapOrigCopy[instrumentFrom] = instrumentTo
	instrumentFrom.CopyBasicFields(instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInstrument_change(mapOrigCopy map[any]any, instrument_changeFrom *Instrument_change) (instrument_changeTo *Instrument_change) {

	// instrument_changeFrom has already been copied
	if _instrument_changeTo, ok := mapOrigCopy[instrument_changeFrom]; ok {
		instrument_changeTo = _instrument_changeTo.(*Instrument_change)
		return
	}

	instrument_changeTo = new(Instrument_change)
	mapOrigCopy[instrument_changeFrom] = instrument_changeTo
	instrument_changeFrom.CopyBasicFields(instrument_changeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInstrument_link(mapOrigCopy map[any]any, instrument_linkFrom *Instrument_link) (instrument_linkTo *Instrument_link) {

	// instrument_linkFrom has already been copied
	if _instrument_linkTo, ok := mapOrigCopy[instrument_linkFrom]; ok {
		instrument_linkTo = _instrument_linkTo.(*Instrument_link)
		return
	}

	instrument_linkTo = new(Instrument_link)
	mapOrigCopy[instrument_linkFrom] = instrument_linkTo
	instrument_linkFrom.CopyBasicFields(instrument_linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInterchangeable(mapOrigCopy map[any]any, interchangeableFrom *Interchangeable) (interchangeableTo *Interchangeable) {

	// interchangeableFrom has already been copied
	if _interchangeableTo, ok := mapOrigCopy[interchangeableFrom]; ok {
		interchangeableTo = _interchangeableTo.(*Interchangeable)
		return
	}

	interchangeableTo = new(Interchangeable)
	mapOrigCopy[interchangeableFrom] = interchangeableTo
	interchangeableFrom.CopyBasicFields(interchangeableTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInversion(mapOrigCopy map[any]any, inversionFrom *Inversion) (inversionTo *Inversion) {

	// inversionFrom has already been copied
	if _inversionTo, ok := mapOrigCopy[inversionFrom]; ok {
		inversionTo = _inversionTo.(*Inversion)
		return
	}

	inversionTo = new(Inversion)
	mapOrigCopy[inversionFrom] = inversionTo
	inversionFrom.CopyBasicFields(inversionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKey(mapOrigCopy map[any]any, keyFrom *Key) (keyTo *Key) {

	// keyFrom has already been copied
	if _keyTo, ok := mapOrigCopy[keyFrom]; ok {
		keyTo = _keyTo.(*Key)
		return
	}

	keyTo = new(Key)
	mapOrigCopy[keyFrom] = keyTo
	keyFrom.CopyBasicFields(keyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key_octave := range keyFrom.Key_octave {
		keyTo.Key_octave = append(keyTo.Key_octave, CopyBranchKey_octave(mapOrigCopy, _key_octave))
	}

	return
}

func CopyBranchKey_accidental(mapOrigCopy map[any]any, key_accidentalFrom *Key_accidental) (key_accidentalTo *Key_accidental) {

	// key_accidentalFrom has already been copied
	if _key_accidentalTo, ok := mapOrigCopy[key_accidentalFrom]; ok {
		key_accidentalTo = _key_accidentalTo.(*Key_accidental)
		return
	}

	key_accidentalTo = new(Key_accidental)
	mapOrigCopy[key_accidentalFrom] = key_accidentalTo
	key_accidentalFrom.CopyBasicFields(key_accidentalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKey_octave(mapOrigCopy map[any]any, key_octaveFrom *Key_octave) (key_octaveTo *Key_octave) {

	// key_octaveFrom has already been copied
	if _key_octaveTo, ok := mapOrigCopy[key_octaveFrom]; ok {
		key_octaveTo = _key_octaveTo.(*Key_octave)
		return
	}

	key_octaveTo = new(Key_octave)
	mapOrigCopy[key_octaveFrom] = key_octaveTo
	key_octaveFrom.CopyBasicFields(key_octaveTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchKind(mapOrigCopy map[any]any, kindFrom *Kind) (kindTo *Kind) {

	// kindFrom has already been copied
	if _kindTo, ok := mapOrigCopy[kindFrom]; ok {
		kindTo = _kindTo.(*Kind)
		return
	}

	kindTo = new(Kind)
	mapOrigCopy[kindFrom] = kindTo
	kindFrom.CopyBasicFields(kindTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLevel(mapOrigCopy map[any]any, levelFrom *Level) (levelTo *Level) {

	// levelFrom has already been copied
	if _levelTo, ok := mapOrigCopy[levelFrom]; ok {
		levelTo = _levelTo.(*Level)
		return
	}

	levelTo = new(Level)
	mapOrigCopy[levelFrom] = levelTo
	levelFrom.CopyBasicFields(levelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLine_detail(mapOrigCopy map[any]any, line_detailFrom *Line_detail) (line_detailTo *Line_detail) {

	// line_detailFrom has already been copied
	if _line_detailTo, ok := mapOrigCopy[line_detailFrom]; ok {
		line_detailTo = _line_detailTo.(*Line_detail)
		return
	}

	line_detailTo = new(Line_detail)
	mapOrigCopy[line_detailFrom] = line_detailTo
	line_detailFrom.CopyBasicFields(line_detailTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLine_width(mapOrigCopy map[any]any, line_widthFrom *Line_width) (line_widthTo *Line_width) {

	// line_widthFrom has already been copied
	if _line_widthTo, ok := mapOrigCopy[line_widthFrom]; ok {
		line_widthTo = _line_widthTo.(*Line_width)
		return
	}

	line_widthTo = new(Line_width)
	mapOrigCopy[line_widthFrom] = line_widthTo
	line_widthFrom.CopyBasicFields(line_widthTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
	linkFrom.CopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchListen(mapOrigCopy map[any]any, listenFrom *Listen) (listenTo *Listen) {

	// listenFrom has already been copied
	if _listenTo, ok := mapOrigCopy[listenFrom]; ok {
		listenTo = _listenTo.(*Listen)
		return
	}

	listenTo = new(Listen)
	mapOrigCopy[listenFrom] = listenTo
	listenFrom.CopyBasicFields(listenTo)

	//insertion point for the staging of instances referenced by pointers
	if listenFrom.Assess != nil {
		listenTo.Assess = CopyBranchAssess(mapOrigCopy, listenFrom.Assess)
	}
	if listenFrom.Wait != nil {
		listenTo.Wait = CopyBranchWait(mapOrigCopy, listenFrom.Wait)
	}
	if listenFrom.Other_listen != nil {
		listenTo.Other_listen = CopyBranchOther_listening(mapOrigCopy, listenFrom.Other_listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchListening(mapOrigCopy map[any]any, listeningFrom *Listening) (listeningTo *Listening) {

	// listeningFrom has already been copied
	if _listeningTo, ok := mapOrigCopy[listeningFrom]; ok {
		listeningTo = _listeningTo.(*Listening)
		return
	}

	listeningTo = new(Listening)
	mapOrigCopy[listeningFrom] = listeningTo
	listeningFrom.CopyBasicFields(listeningTo)

	//insertion point for the staging of instances referenced by pointers
	if listeningFrom.Offset != nil {
		listeningTo.Offset = CopyBranchOffset(mapOrigCopy, listeningFrom.Offset)
	}
	if listeningFrom.Sync != nil {
		listeningTo.Sync = CopyBranchSync(mapOrigCopy, listeningFrom.Sync)
	}
	if listeningFrom.Other_listening != nil {
		listeningTo.Other_listening = CopyBranchOther_listening(mapOrigCopy, listeningFrom.Other_listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLyric(mapOrigCopy map[any]any, lyricFrom *Lyric) (lyricTo *Lyric) {

	// lyricFrom has already been copied
	if _lyricTo, ok := mapOrigCopy[lyricFrom]; ok {
		lyricTo = _lyricTo.(*Lyric)
		return
	}

	lyricTo = new(Lyric)
	mapOrigCopy[lyricFrom] = lyricTo
	lyricFrom.CopyBasicFields(lyricTo)

	//insertion point for the staging of instances referenced by pointers
	if lyricFrom.End_line != nil {
		lyricTo.End_line = CopyBranchEmpty(mapOrigCopy, lyricFrom.End_line)
	}
	if lyricFrom.End_paragraph != nil {
		lyricTo.End_paragraph = CopyBranchEmpty(mapOrigCopy, lyricFrom.End_paragraph)
	}
	if lyricFrom.Extend != nil {
		lyricTo.Extend = CopyBranchExtend(mapOrigCopy, lyricFrom.Extend)
	}
	if lyricFrom.Laughing != nil {
		lyricTo.Laughing = CopyBranchEmpty(mapOrigCopy, lyricFrom.Laughing)
	}
	if lyricFrom.Humming != nil {
		lyricTo.Humming = CopyBranchEmpty(mapOrigCopy, lyricFrom.Humming)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLyric_font(mapOrigCopy map[any]any, lyric_fontFrom *Lyric_font) (lyric_fontTo *Lyric_font) {

	// lyric_fontFrom has already been copied
	if _lyric_fontTo, ok := mapOrigCopy[lyric_fontFrom]; ok {
		lyric_fontTo = _lyric_fontTo.(*Lyric_font)
		return
	}

	lyric_fontTo = new(Lyric_font)
	mapOrigCopy[lyric_fontFrom] = lyric_fontTo
	lyric_fontFrom.CopyBasicFields(lyric_fontTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLyric_language(mapOrigCopy map[any]any, lyric_languageFrom *Lyric_language) (lyric_languageTo *Lyric_language) {

	// lyric_languageFrom has already been copied
	if _lyric_languageTo, ok := mapOrigCopy[lyric_languageFrom]; ok {
		lyric_languageTo = _lyric_languageTo.(*Lyric_language)
		return
	}

	lyric_languageTo = new(Lyric_language)
	mapOrigCopy[lyric_languageFrom] = lyric_languageTo
	lyric_languageFrom.CopyBasicFields(lyric_languageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMeasure_layout(mapOrigCopy map[any]any, measure_layoutFrom *Measure_layout) (measure_layoutTo *Measure_layout) {

	// measure_layoutFrom has already been copied
	if _measure_layoutTo, ok := mapOrigCopy[measure_layoutFrom]; ok {
		measure_layoutTo = _measure_layoutTo.(*Measure_layout)
		return
	}

	measure_layoutTo = new(Measure_layout)
	mapOrigCopy[measure_layoutFrom] = measure_layoutTo
	measure_layoutFrom.CopyBasicFields(measure_layoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMeasure_numbering(mapOrigCopy map[any]any, measure_numberingFrom *Measure_numbering) (measure_numberingTo *Measure_numbering) {

	// measure_numberingFrom has already been copied
	if _measure_numberingTo, ok := mapOrigCopy[measure_numberingFrom]; ok {
		measure_numberingTo = _measure_numberingTo.(*Measure_numbering)
		return
	}

	measure_numberingTo = new(Measure_numbering)
	mapOrigCopy[measure_numberingFrom] = measure_numberingTo
	measure_numberingFrom.CopyBasicFields(measure_numberingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMeasure_repeat(mapOrigCopy map[any]any, measure_repeatFrom *Measure_repeat) (measure_repeatTo *Measure_repeat) {

	// measure_repeatFrom has already been copied
	if _measure_repeatTo, ok := mapOrigCopy[measure_repeatFrom]; ok {
		measure_repeatTo = _measure_repeatTo.(*Measure_repeat)
		return
	}

	measure_repeatTo = new(Measure_repeat)
	mapOrigCopy[measure_repeatFrom] = measure_repeatTo
	measure_repeatFrom.CopyBasicFields(measure_repeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMeasure_style(mapOrigCopy map[any]any, measure_styleFrom *Measure_style) (measure_styleTo *Measure_style) {

	// measure_styleFrom has already been copied
	if _measure_styleTo, ok := mapOrigCopy[measure_styleFrom]; ok {
		measure_styleTo = _measure_styleTo.(*Measure_style)
		return
	}

	measure_styleTo = new(Measure_style)
	mapOrigCopy[measure_styleFrom] = measure_styleTo
	measure_styleFrom.CopyBasicFields(measure_styleTo)

	//insertion point for the staging of instances referenced by pointers
	if measure_styleFrom.Multiple_rest != nil {
		measure_styleTo.Multiple_rest = CopyBranchMultiple_rest(mapOrigCopy, measure_styleFrom.Multiple_rest)
	}
	if measure_styleFrom.Measure_repeat != nil {
		measure_styleTo.Measure_repeat = CopyBranchMeasure_repeat(mapOrigCopy, measure_styleFrom.Measure_repeat)
	}
	if measure_styleFrom.Beat_repeat != nil {
		measure_styleTo.Beat_repeat = CopyBranchBeat_repeat(mapOrigCopy, measure_styleFrom.Beat_repeat)
	}
	if measure_styleFrom.Slash != nil {
		measure_styleTo.Slash = CopyBranchSlash(mapOrigCopy, measure_styleFrom.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMembrane(mapOrigCopy map[any]any, membraneFrom *Membrane) (membraneTo *Membrane) {

	// membraneFrom has already been copied
	if _membraneTo, ok := mapOrigCopy[membraneFrom]; ok {
		membraneTo = _membraneTo.(*Membrane)
		return
	}

	membraneTo = new(Membrane)
	mapOrigCopy[membraneFrom] = membraneTo
	membraneFrom.CopyBasicFields(membraneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMetal(mapOrigCopy map[any]any, metalFrom *Metal) (metalTo *Metal) {

	// metalFrom has already been copied
	if _metalTo, ok := mapOrigCopy[metalFrom]; ok {
		metalTo = _metalTo.(*Metal)
		return
	}

	metalTo = new(Metal)
	mapOrigCopy[metalFrom] = metalTo
	metalFrom.CopyBasicFields(metalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMetronome(mapOrigCopy map[any]any, metronomeFrom *Metronome) (metronomeTo *Metronome) {

	// metronomeFrom has already been copied
	if _metronomeTo, ok := mapOrigCopy[metronomeFrom]; ok {
		metronomeTo = _metronomeTo.(*Metronome)
		return
	}

	metronomeTo = new(Metronome)
	mapOrigCopy[metronomeFrom] = metronomeTo
	metronomeFrom.CopyBasicFields(metronomeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMetronome_beam(mapOrigCopy map[any]any, metronome_beamFrom *Metronome_beam) (metronome_beamTo *Metronome_beam) {

	// metronome_beamFrom has already been copied
	if _metronome_beamTo, ok := mapOrigCopy[metronome_beamFrom]; ok {
		metronome_beamTo = _metronome_beamTo.(*Metronome_beam)
		return
	}

	metronome_beamTo = new(Metronome_beam)
	mapOrigCopy[metronome_beamFrom] = metronome_beamTo
	metronome_beamFrom.CopyBasicFields(metronome_beamTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMetronome_note(mapOrigCopy map[any]any, metronome_noteFrom *Metronome_note) (metronome_noteTo *Metronome_note) {

	// metronome_noteFrom has already been copied
	if _metronome_noteTo, ok := mapOrigCopy[metronome_noteFrom]; ok {
		metronome_noteTo = _metronome_noteTo.(*Metronome_note)
		return
	}

	metronome_noteTo = new(Metronome_note)
	mapOrigCopy[metronome_noteFrom] = metronome_noteTo
	metronome_noteFrom.CopyBasicFields(metronome_noteTo)

	//insertion point for the staging of instances referenced by pointers
	if metronome_noteFrom.Metronome_tied != nil {
		metronome_noteTo.Metronome_tied = CopyBranchMetronome_tied(mapOrigCopy, metronome_noteFrom.Metronome_tied)
	}
	if metronome_noteFrom.Metronome_tuplet != nil {
		metronome_noteTo.Metronome_tuplet = CopyBranchMetronome_tuplet(mapOrigCopy, metronome_noteFrom.Metronome_tuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty := range metronome_noteFrom.Metronome_dot {
		metronome_noteTo.Metronome_dot = append(metronome_noteTo.Metronome_dot, CopyBranchEmpty(mapOrigCopy, _empty))
	}
	for _, _metronome_beam := range metronome_noteFrom.Metronome_beam {
		metronome_noteTo.Metronome_beam = append(metronome_noteTo.Metronome_beam, CopyBranchMetronome_beam(mapOrigCopy, _metronome_beam))
	}

	return
}

func CopyBranchMetronome_tied(mapOrigCopy map[any]any, metronome_tiedFrom *Metronome_tied) (metronome_tiedTo *Metronome_tied) {

	// metronome_tiedFrom has already been copied
	if _metronome_tiedTo, ok := mapOrigCopy[metronome_tiedFrom]; ok {
		metronome_tiedTo = _metronome_tiedTo.(*Metronome_tied)
		return
	}

	metronome_tiedTo = new(Metronome_tied)
	mapOrigCopy[metronome_tiedFrom] = metronome_tiedTo
	metronome_tiedFrom.CopyBasicFields(metronome_tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMetronome_tuplet(mapOrigCopy map[any]any, metronome_tupletFrom *Metronome_tuplet) (metronome_tupletTo *Metronome_tuplet) {

	// metronome_tupletFrom has already been copied
	if _metronome_tupletTo, ok := mapOrigCopy[metronome_tupletFrom]; ok {
		metronome_tupletTo = _metronome_tupletTo.(*Metronome_tuplet)
		return
	}

	metronome_tupletTo = new(Metronome_tuplet)
	mapOrigCopy[metronome_tupletFrom] = metronome_tupletTo
	metronome_tupletFrom.CopyBasicFields(metronome_tupletTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMidi_device(mapOrigCopy map[any]any, midi_deviceFrom *Midi_device) (midi_deviceTo *Midi_device) {

	// midi_deviceFrom has already been copied
	if _midi_deviceTo, ok := mapOrigCopy[midi_deviceFrom]; ok {
		midi_deviceTo = _midi_deviceTo.(*Midi_device)
		return
	}

	midi_deviceTo = new(Midi_device)
	mapOrigCopy[midi_deviceFrom] = midi_deviceTo
	midi_deviceFrom.CopyBasicFields(midi_deviceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMidi_instrument(mapOrigCopy map[any]any, midi_instrumentFrom *Midi_instrument) (midi_instrumentTo *Midi_instrument) {

	// midi_instrumentFrom has already been copied
	if _midi_instrumentTo, ok := mapOrigCopy[midi_instrumentFrom]; ok {
		midi_instrumentTo = _midi_instrumentTo.(*Midi_instrument)
		return
	}

	midi_instrumentTo = new(Midi_instrument)
	mapOrigCopy[midi_instrumentFrom] = midi_instrumentTo
	midi_instrumentFrom.CopyBasicFields(midi_instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMiscellaneous(mapOrigCopy map[any]any, miscellaneousFrom *Miscellaneous) (miscellaneousTo *Miscellaneous) {

	// miscellaneousFrom has already been copied
	if _miscellaneousTo, ok := mapOrigCopy[miscellaneousFrom]; ok {
		miscellaneousTo = _miscellaneousTo.(*Miscellaneous)
		return
	}

	miscellaneousTo = new(Miscellaneous)
	mapOrigCopy[miscellaneousFrom] = miscellaneousTo
	miscellaneousFrom.CopyBasicFields(miscellaneousTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneous_field := range miscellaneousFrom.Miscellaneous_field {
		miscellaneousTo.Miscellaneous_field = append(miscellaneousTo.Miscellaneous_field, CopyBranchMiscellaneous_field(mapOrigCopy, _miscellaneous_field))
	}

	return
}

func CopyBranchMiscellaneous_field(mapOrigCopy map[any]any, miscellaneous_fieldFrom *Miscellaneous_field) (miscellaneous_fieldTo *Miscellaneous_field) {

	// miscellaneous_fieldFrom has already been copied
	if _miscellaneous_fieldTo, ok := mapOrigCopy[miscellaneous_fieldFrom]; ok {
		miscellaneous_fieldTo = _miscellaneous_fieldTo.(*Miscellaneous_field)
		return
	}

	miscellaneous_fieldTo = new(Miscellaneous_field)
	mapOrigCopy[miscellaneous_fieldFrom] = miscellaneous_fieldTo
	miscellaneous_fieldFrom.CopyBasicFields(miscellaneous_fieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMordent(mapOrigCopy map[any]any, mordentFrom *Mordent) (mordentTo *Mordent) {

	// mordentFrom has already been copied
	if _mordentTo, ok := mapOrigCopy[mordentFrom]; ok {
		mordentTo = _mordentTo.(*Mordent)
		return
	}

	mordentTo = new(Mordent)
	mapOrigCopy[mordentFrom] = mordentTo
	mordentFrom.CopyBasicFields(mordentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMultiple_rest(mapOrigCopy map[any]any, multiple_restFrom *Multiple_rest) (multiple_restTo *Multiple_rest) {

	// multiple_restFrom has already been copied
	if _multiple_restTo, ok := mapOrigCopy[multiple_restFrom]; ok {
		multiple_restTo = _multiple_restTo.(*Multiple_rest)
		return
	}

	multiple_restTo = new(Multiple_rest)
	mapOrigCopy[multiple_restFrom] = multiple_restTo
	multiple_restFrom.CopyBasicFields(multiple_restTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchName_display(mapOrigCopy map[any]any, name_displayFrom *Name_display) (name_displayTo *Name_display) {

	// name_displayFrom has already been copied
	if _name_displayTo, ok := mapOrigCopy[name_displayFrom]; ok {
		name_displayTo = _name_displayTo.(*Name_display)
		return
	}

	name_displayTo = new(Name_display)
	mapOrigCopy[name_displayFrom] = name_displayTo
	name_displayFrom.CopyBasicFields(name_displayTo)

	//insertion point for the staging of instances referenced by pointers
	if name_displayFrom.Accidental_text != nil {
		name_displayTo.Accidental_text = CopyBranchAccidental_text(mapOrigCopy, name_displayFrom.Accidental_text)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNon_arpeggiate(mapOrigCopy map[any]any, non_arpeggiateFrom *Non_arpeggiate) (non_arpeggiateTo *Non_arpeggiate) {

	// non_arpeggiateFrom has already been copied
	if _non_arpeggiateTo, ok := mapOrigCopy[non_arpeggiateFrom]; ok {
		non_arpeggiateTo = _non_arpeggiateTo.(*Non_arpeggiate)
		return
	}

	non_arpeggiateTo = new(Non_arpeggiate)
	mapOrigCopy[non_arpeggiateFrom] = non_arpeggiateTo
	non_arpeggiateFrom.CopyBasicFields(non_arpeggiateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNotations(mapOrigCopy map[any]any, notationsFrom *Notations) (notationsTo *Notations) {

	// notationsFrom has already been copied
	if _notationsTo, ok := mapOrigCopy[notationsFrom]; ok {
		notationsTo = _notationsTo.(*Notations)
		return
	}

	notationsTo = new(Notations)
	mapOrigCopy[notationsFrom] = notationsTo
	notationsFrom.CopyBasicFields(notationsTo)

	//insertion point for the staging of instances referenced by pointers
	if notationsFrom.Tied != nil {
		notationsTo.Tied = CopyBranchTied(mapOrigCopy, notationsFrom.Tied)
	}
	if notationsFrom.Slur != nil {
		notationsTo.Slur = CopyBranchSlur(mapOrigCopy, notationsFrom.Slur)
	}
	if notationsFrom.Tuplet != nil {
		notationsTo.Tuplet = CopyBranchTuplet(mapOrigCopy, notationsFrom.Tuplet)
	}
	if notationsFrom.Glissando != nil {
		notationsTo.Glissando = CopyBranchGlissando(mapOrigCopy, notationsFrom.Glissando)
	}
	if notationsFrom.Slide != nil {
		notationsTo.Slide = CopyBranchSlide(mapOrigCopy, notationsFrom.Slide)
	}
	if notationsFrom.Ornaments != nil {
		notationsTo.Ornaments = CopyBranchOrnaments(mapOrigCopy, notationsFrom.Ornaments)
	}
	if notationsFrom.Technical != nil {
		notationsTo.Technical = CopyBranchTechnical(mapOrigCopy, notationsFrom.Technical)
	}
	if notationsFrom.Articulations != nil {
		notationsTo.Articulations = CopyBranchArticulations(mapOrigCopy, notationsFrom.Articulations)
	}
	if notationsFrom.Dynamics != nil {
		notationsTo.Dynamics = CopyBranchDynamics(mapOrigCopy, notationsFrom.Dynamics)
	}
	if notationsFrom.Fermata != nil {
		notationsTo.Fermata = CopyBranchFermata(mapOrigCopy, notationsFrom.Fermata)
	}
	if notationsFrom.Arpeggiate != nil {
		notationsTo.Arpeggiate = CopyBranchArpeggiate(mapOrigCopy, notationsFrom.Arpeggiate)
	}
	if notationsFrom.Non_arpeggiate != nil {
		notationsTo.Non_arpeggiate = CopyBranchNon_arpeggiate(mapOrigCopy, notationsFrom.Non_arpeggiate)
	}
	if notationsFrom.Accidental_mark != nil {
		notationsTo.Accidental_mark = CopyBranchAccidental_mark(mapOrigCopy, notationsFrom.Accidental_mark)
	}
	if notationsFrom.Other_notation != nil {
		notationsTo.Other_notation = CopyBranchOther_notation(mapOrigCopy, notationsFrom.Other_notation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers
	if noteFrom.Type_ != nil {
		noteTo.Type_ = CopyBranchNote_type(mapOrigCopy, noteFrom.Type_)
	}
	if noteFrom.Accidental != nil {
		noteTo.Accidental = CopyBranchAccidental(mapOrigCopy, noteFrom.Accidental)
	}
	if noteFrom.Time_modification != nil {
		noteTo.Time_modification = CopyBranchTime_modification(mapOrigCopy, noteFrom.Time_modification)
	}
	if noteFrom.Stem != nil {
		noteTo.Stem = CopyBranchStem(mapOrigCopy, noteFrom.Stem)
	}
	if noteFrom.Notehead != nil {
		noteTo.Notehead = CopyBranchNotehead(mapOrigCopy, noteFrom.Notehead)
	}
	if noteFrom.Notehead_text != nil {
		noteTo.Notehead_text = CopyBranchNotehead_text(mapOrigCopy, noteFrom.Notehead_text)
	}
	if noteFrom.Beam != nil {
		noteTo.Beam = CopyBranchBeam(mapOrigCopy, noteFrom.Beam)
	}
	if noteFrom.Play != nil {
		noteTo.Play = CopyBranchPlay(mapOrigCopy, noteFrom.Play)
	}
	if noteFrom.Listen != nil {
		noteTo.Listen = CopyBranchListen(mapOrigCopy, noteFrom.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument := range noteFrom.Instrument {
		noteTo.Instrument = append(noteTo.Instrument, CopyBranchInstrument(mapOrigCopy, _instrument))
	}
	for _, _empty_placement := range noteFrom.Dot {
		noteTo.Dot = append(noteTo.Dot, CopyBranchEmpty_placement(mapOrigCopy, _empty_placement))
	}
	for _, _notations := range noteFrom.Notations {
		noteTo.Notations = append(noteTo.Notations, CopyBranchNotations(mapOrigCopy, _notations))
	}
	for _, _lyric := range noteFrom.Lyric {
		noteTo.Lyric = append(noteTo.Lyric, CopyBranchLyric(mapOrigCopy, _lyric))
	}

	return
}

func CopyBranchNote_size(mapOrigCopy map[any]any, note_sizeFrom *Note_size) (note_sizeTo *Note_size) {

	// note_sizeFrom has already been copied
	if _note_sizeTo, ok := mapOrigCopy[note_sizeFrom]; ok {
		note_sizeTo = _note_sizeTo.(*Note_size)
		return
	}

	note_sizeTo = new(Note_size)
	mapOrigCopy[note_sizeFrom] = note_sizeTo
	note_sizeFrom.CopyBasicFields(note_sizeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNote_type(mapOrigCopy map[any]any, note_typeFrom *Note_type) (note_typeTo *Note_type) {

	// note_typeFrom has already been copied
	if _note_typeTo, ok := mapOrigCopy[note_typeFrom]; ok {
		note_typeTo = _note_typeTo.(*Note_type)
		return
	}

	note_typeTo = new(Note_type)
	mapOrigCopy[note_typeFrom] = note_typeTo
	note_typeFrom.CopyBasicFields(note_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNotehead(mapOrigCopy map[any]any, noteheadFrom *Notehead) (noteheadTo *Notehead) {

	// noteheadFrom has already been copied
	if _noteheadTo, ok := mapOrigCopy[noteheadFrom]; ok {
		noteheadTo = _noteheadTo.(*Notehead)
		return
	}

	noteheadTo = new(Notehead)
	mapOrigCopy[noteheadFrom] = noteheadTo
	noteheadFrom.CopyBasicFields(noteheadTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNotehead_text(mapOrigCopy map[any]any, notehead_textFrom *Notehead_text) (notehead_textTo *Notehead_text) {

	// notehead_textFrom has already been copied
	if _notehead_textTo, ok := mapOrigCopy[notehead_textFrom]; ok {
		notehead_textTo = _notehead_textTo.(*Notehead_text)
		return
	}

	notehead_textTo = new(Notehead_text)
	mapOrigCopy[notehead_textFrom] = notehead_textTo
	notehead_textFrom.CopyBasicFields(notehead_textTo)

	//insertion point for the staging of instances referenced by pointers
	if notehead_textFrom.Accidental_text != nil {
		notehead_textTo.Accidental_text = CopyBranchAccidental_text(mapOrigCopy, notehead_textFrom.Accidental_text)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNumeral(mapOrigCopy map[any]any, numeralFrom *Numeral) (numeralTo *Numeral) {

	// numeralFrom has already been copied
	if _numeralTo, ok := mapOrigCopy[numeralFrom]; ok {
		numeralTo = _numeralTo.(*Numeral)
		return
	}

	numeralTo = new(Numeral)
	mapOrigCopy[numeralFrom] = numeralTo
	numeralFrom.CopyBasicFields(numeralTo)

	//insertion point for the staging of instances referenced by pointers
	if numeralFrom.Numeral_root != nil {
		numeralTo.Numeral_root = CopyBranchNumeral_root(mapOrigCopy, numeralFrom.Numeral_root)
	}
	if numeralFrom.Numeral_alter != nil {
		numeralTo.Numeral_alter = CopyBranchHarmony_alter(mapOrigCopy, numeralFrom.Numeral_alter)
	}
	if numeralFrom.Numeral_key != nil {
		numeralTo.Numeral_key = CopyBranchNumeral_key(mapOrigCopy, numeralFrom.Numeral_key)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNumeral_key(mapOrigCopy map[any]any, numeral_keyFrom *Numeral_key) (numeral_keyTo *Numeral_key) {

	// numeral_keyFrom has already been copied
	if _numeral_keyTo, ok := mapOrigCopy[numeral_keyFrom]; ok {
		numeral_keyTo = _numeral_keyTo.(*Numeral_key)
		return
	}

	numeral_keyTo = new(Numeral_key)
	mapOrigCopy[numeral_keyFrom] = numeral_keyTo
	numeral_keyFrom.CopyBasicFields(numeral_keyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNumeral_root(mapOrigCopy map[any]any, numeral_rootFrom *Numeral_root) (numeral_rootTo *Numeral_root) {

	// numeral_rootFrom has already been copied
	if _numeral_rootTo, ok := mapOrigCopy[numeral_rootFrom]; ok {
		numeral_rootTo = _numeral_rootTo.(*Numeral_root)
		return
	}

	numeral_rootTo = new(Numeral_root)
	mapOrigCopy[numeral_rootFrom] = numeral_rootTo
	numeral_rootFrom.CopyBasicFields(numeral_rootTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOctave_shift(mapOrigCopy map[any]any, octave_shiftFrom *Octave_shift) (octave_shiftTo *Octave_shift) {

	// octave_shiftFrom has already been copied
	if _octave_shiftTo, ok := mapOrigCopy[octave_shiftFrom]; ok {
		octave_shiftTo = _octave_shiftTo.(*Octave_shift)
		return
	}

	octave_shiftTo = new(Octave_shift)
	mapOrigCopy[octave_shiftFrom] = octave_shiftTo
	octave_shiftFrom.CopyBasicFields(octave_shiftTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOffset(mapOrigCopy map[any]any, offsetFrom *Offset) (offsetTo *Offset) {

	// offsetFrom has already been copied
	if _offsetTo, ok := mapOrigCopy[offsetFrom]; ok {
		offsetTo = _offsetTo.(*Offset)
		return
	}

	offsetTo = new(Offset)
	mapOrigCopy[offsetFrom] = offsetTo
	offsetFrom.CopyBasicFields(offsetTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOpus(mapOrigCopy map[any]any, opusFrom *Opus) (opusTo *Opus) {

	// opusFrom has already been copied
	if _opusTo, ok := mapOrigCopy[opusFrom]; ok {
		opusTo = _opusTo.(*Opus)
		return
	}

	opusTo = new(Opus)
	mapOrigCopy[opusFrom] = opusTo
	opusFrom.CopyBasicFields(opusTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOrnaments(mapOrigCopy map[any]any, ornamentsFrom *Ornaments) (ornamentsTo *Ornaments) {

	// ornamentsFrom has already been copied
	if _ornamentsTo, ok := mapOrigCopy[ornamentsFrom]; ok {
		ornamentsTo = _ornamentsTo.(*Ornaments)
		return
	}

	ornamentsTo = new(Ornaments)
	mapOrigCopy[ornamentsFrom] = ornamentsTo
	ornamentsFrom.CopyBasicFields(ornamentsTo)

	//insertion point for the staging of instances referenced by pointers
	if ornamentsFrom.Trill_mark != nil {
		ornamentsTo.Trill_mark = CopyBranchEmpty_trill_sound(mapOrigCopy, ornamentsFrom.Trill_mark)
	}
	if ornamentsFrom.Turn != nil {
		ornamentsTo.Turn = CopyBranchHorizontal_turn(mapOrigCopy, ornamentsFrom.Turn)
	}
	if ornamentsFrom.Delayed_turn != nil {
		ornamentsTo.Delayed_turn = CopyBranchHorizontal_turn(mapOrigCopy, ornamentsFrom.Delayed_turn)
	}
	if ornamentsFrom.Inverted_turn != nil {
		ornamentsTo.Inverted_turn = CopyBranchHorizontal_turn(mapOrigCopy, ornamentsFrom.Inverted_turn)
	}
	if ornamentsFrom.Delayed_inverted_turn != nil {
		ornamentsTo.Delayed_inverted_turn = CopyBranchHorizontal_turn(mapOrigCopy, ornamentsFrom.Delayed_inverted_turn)
	}
	if ornamentsFrom.Vertical_turn != nil {
		ornamentsTo.Vertical_turn = CopyBranchEmpty_trill_sound(mapOrigCopy, ornamentsFrom.Vertical_turn)
	}
	if ornamentsFrom.Inverted_vertical_turn != nil {
		ornamentsTo.Inverted_vertical_turn = CopyBranchEmpty_trill_sound(mapOrigCopy, ornamentsFrom.Inverted_vertical_turn)
	}
	if ornamentsFrom.Shake != nil {
		ornamentsTo.Shake = CopyBranchEmpty_trill_sound(mapOrigCopy, ornamentsFrom.Shake)
	}
	if ornamentsFrom.Wavy_line != nil {
		ornamentsTo.Wavy_line = CopyBranchWavy_line(mapOrigCopy, ornamentsFrom.Wavy_line)
	}
	if ornamentsFrom.Mordent != nil {
		ornamentsTo.Mordent = CopyBranchMordent(mapOrigCopy, ornamentsFrom.Mordent)
	}
	if ornamentsFrom.Inverted_mordent != nil {
		ornamentsTo.Inverted_mordent = CopyBranchMordent(mapOrigCopy, ornamentsFrom.Inverted_mordent)
	}
	if ornamentsFrom.Schleifer != nil {
		ornamentsTo.Schleifer = CopyBranchEmpty_placement(mapOrigCopy, ornamentsFrom.Schleifer)
	}
	if ornamentsFrom.Tremolo != nil {
		ornamentsTo.Tremolo = CopyBranchTremolo(mapOrigCopy, ornamentsFrom.Tremolo)
	}
	if ornamentsFrom.Haydn != nil {
		ornamentsTo.Haydn = CopyBranchEmpty_trill_sound(mapOrigCopy, ornamentsFrom.Haydn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accidental_mark := range ornamentsFrom.Accidental_mark {
		ornamentsTo.Accidental_mark = append(ornamentsTo.Accidental_mark, CopyBranchAccidental_mark(mapOrigCopy, _accidental_mark))
	}

	return
}

func CopyBranchOther_appearance(mapOrigCopy map[any]any, other_appearanceFrom *Other_appearance) (other_appearanceTo *Other_appearance) {

	// other_appearanceFrom has already been copied
	if _other_appearanceTo, ok := mapOrigCopy[other_appearanceFrom]; ok {
		other_appearanceTo = _other_appearanceTo.(*Other_appearance)
		return
	}

	other_appearanceTo = new(Other_appearance)
	mapOrigCopy[other_appearanceFrom] = other_appearanceTo
	other_appearanceFrom.CopyBasicFields(other_appearanceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOther_listening(mapOrigCopy map[any]any, other_listeningFrom *Other_listening) (other_listeningTo *Other_listening) {

	// other_listeningFrom has already been copied
	if _other_listeningTo, ok := mapOrigCopy[other_listeningFrom]; ok {
		other_listeningTo = _other_listeningTo.(*Other_listening)
		return
	}

	other_listeningTo = new(Other_listening)
	mapOrigCopy[other_listeningFrom] = other_listeningTo
	other_listeningFrom.CopyBasicFields(other_listeningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOther_notation(mapOrigCopy map[any]any, other_notationFrom *Other_notation) (other_notationTo *Other_notation) {

	// other_notationFrom has already been copied
	if _other_notationTo, ok := mapOrigCopy[other_notationFrom]; ok {
		other_notationTo = _other_notationTo.(*Other_notation)
		return
	}

	other_notationTo = new(Other_notation)
	mapOrigCopy[other_notationFrom] = other_notationTo
	other_notationFrom.CopyBasicFields(other_notationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchOther_play(mapOrigCopy map[any]any, other_playFrom *Other_play) (other_playTo *Other_play) {

	// other_playFrom has already been copied
	if _other_playTo, ok := mapOrigCopy[other_playFrom]; ok {
		other_playTo = _other_playTo.(*Other_play)
		return
	}

	other_playTo = new(Other_play)
	mapOrigCopy[other_playFrom] = other_playTo
	other_playFrom.CopyBasicFields(other_playTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPage_layout(mapOrigCopy map[any]any, page_layoutFrom *Page_layout) (page_layoutTo *Page_layout) {

	// page_layoutFrom has already been copied
	if _page_layoutTo, ok := mapOrigCopy[page_layoutFrom]; ok {
		page_layoutTo = _page_layoutTo.(*Page_layout)
		return
	}

	page_layoutTo = new(Page_layout)
	mapOrigCopy[page_layoutFrom] = page_layoutTo
	page_layoutFrom.CopyBasicFields(page_layoutTo)

	//insertion point for the staging of instances referenced by pointers
	if page_layoutFrom.Page_margins != nil {
		page_layoutTo.Page_margins = CopyBranchPage_margins(mapOrigCopy, page_layoutFrom.Page_margins)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPage_margins(mapOrigCopy map[any]any, page_marginsFrom *Page_margins) (page_marginsTo *Page_margins) {

	// page_marginsFrom has already been copied
	if _page_marginsTo, ok := mapOrigCopy[page_marginsFrom]; ok {
		page_marginsTo = _page_marginsTo.(*Page_margins)
		return
	}

	page_marginsTo = new(Page_margins)
	mapOrigCopy[page_marginsFrom] = page_marginsTo
	page_marginsFrom.CopyBasicFields(page_marginsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPart_clef(mapOrigCopy map[any]any, part_clefFrom *Part_clef) (part_clefTo *Part_clef) {

	// part_clefFrom has already been copied
	if _part_clefTo, ok := mapOrigCopy[part_clefFrom]; ok {
		part_clefTo = _part_clefTo.(*Part_clef)
		return
	}

	part_clefTo = new(Part_clef)
	mapOrigCopy[part_clefFrom] = part_clefTo
	part_clefFrom.CopyBasicFields(part_clefTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPart_group(mapOrigCopy map[any]any, part_groupFrom *Part_group) (part_groupTo *Part_group) {

	// part_groupFrom has already been copied
	if _part_groupTo, ok := mapOrigCopy[part_groupFrom]; ok {
		part_groupTo = _part_groupTo.(*Part_group)
		return
	}

	part_groupTo = new(Part_group)
	mapOrigCopy[part_groupFrom] = part_groupTo
	part_groupFrom.CopyBasicFields(part_groupTo)

	//insertion point for the staging of instances referenced by pointers
	if part_groupFrom.Group_name_display != nil {
		part_groupTo.Group_name_display = CopyBranchName_display(mapOrigCopy, part_groupFrom.Group_name_display)
	}
	if part_groupFrom.Group_abbreviation_display != nil {
		part_groupTo.Group_abbreviation_display = CopyBranchName_display(mapOrigCopy, part_groupFrom.Group_abbreviation_display)
	}
	if part_groupFrom.Group_symbol != nil {
		part_groupTo.Group_symbol = CopyBranchGroup_symbol(mapOrigCopy, part_groupFrom.Group_symbol)
	}
	if part_groupFrom.Group_barline != nil {
		part_groupTo.Group_barline = CopyBranchGroup_barline(mapOrigCopy, part_groupFrom.Group_barline)
	}
	if part_groupFrom.Group_time != nil {
		part_groupTo.Group_time = CopyBranchEmpty(mapOrigCopy, part_groupFrom.Group_time)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPart_link(mapOrigCopy map[any]any, part_linkFrom *Part_link) (part_linkTo *Part_link) {

	// part_linkFrom has already been copied
	if _part_linkTo, ok := mapOrigCopy[part_linkFrom]; ok {
		part_linkTo = _part_linkTo.(*Part_link)
		return
	}

	part_linkTo = new(Part_link)
	mapOrigCopy[part_linkFrom] = part_linkTo
	part_linkFrom.CopyBasicFields(part_linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_link := range part_linkFrom.Instrument_link {
		part_linkTo.Instrument_link = append(part_linkTo.Instrument_link, CopyBranchInstrument_link(mapOrigCopy, _instrument_link))
	}

	return
}

func CopyBranchPart_list(mapOrigCopy map[any]any, part_listFrom *Part_list) (part_listTo *Part_list) {

	// part_listFrom has already been copied
	if _part_listTo, ok := mapOrigCopy[part_listFrom]; ok {
		part_listTo = _part_listTo.(*Part_list)
		return
	}

	part_listTo = new(Part_list)
	mapOrigCopy[part_listFrom] = part_listTo
	part_listFrom.CopyBasicFields(part_listTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPart_symbol(mapOrigCopy map[any]any, part_symbolFrom *Part_symbol) (part_symbolTo *Part_symbol) {

	// part_symbolFrom has already been copied
	if _part_symbolTo, ok := mapOrigCopy[part_symbolFrom]; ok {
		part_symbolTo = _part_symbolTo.(*Part_symbol)
		return
	}

	part_symbolTo = new(Part_symbol)
	mapOrigCopy[part_symbolFrom] = part_symbolTo
	part_symbolFrom.CopyBasicFields(part_symbolTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPart_transpose(mapOrigCopy map[any]any, part_transposeFrom *Part_transpose) (part_transposeTo *Part_transpose) {

	// part_transposeFrom has already been copied
	if _part_transposeTo, ok := mapOrigCopy[part_transposeFrom]; ok {
		part_transposeTo = _part_transposeTo.(*Part_transpose)
		return
	}

	part_transposeTo = new(Part_transpose)
	mapOrigCopy[part_transposeFrom] = part_transposeTo
	part_transposeFrom.CopyBasicFields(part_transposeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPedal(mapOrigCopy map[any]any, pedalFrom *Pedal) (pedalTo *Pedal) {

	// pedalFrom has already been copied
	if _pedalTo, ok := mapOrigCopy[pedalFrom]; ok {
		pedalTo = _pedalTo.(*Pedal)
		return
	}

	pedalTo = new(Pedal)
	mapOrigCopy[pedalFrom] = pedalTo
	pedalFrom.CopyBasicFields(pedalTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPedal_tuning(mapOrigCopy map[any]any, pedal_tuningFrom *Pedal_tuning) (pedal_tuningTo *Pedal_tuning) {

	// pedal_tuningFrom has already been copied
	if _pedal_tuningTo, ok := mapOrigCopy[pedal_tuningFrom]; ok {
		pedal_tuningTo = _pedal_tuningTo.(*Pedal_tuning)
		return
	}

	pedal_tuningTo = new(Pedal_tuning)
	mapOrigCopy[pedal_tuningFrom] = pedal_tuningTo
	pedal_tuningFrom.CopyBasicFields(pedal_tuningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPercussion(mapOrigCopy map[any]any, percussionFrom *Percussion) (percussionTo *Percussion) {

	// percussionFrom has already been copied
	if _percussionTo, ok := mapOrigCopy[percussionFrom]; ok {
		percussionTo = _percussionTo.(*Percussion)
		return
	}

	percussionTo = new(Percussion)
	mapOrigCopy[percussionFrom] = percussionTo
	percussionFrom.CopyBasicFields(percussionTo)

	//insertion point for the staging of instances referenced by pointers
	if percussionFrom.Glass != nil {
		percussionTo.Glass = CopyBranchGlass(mapOrigCopy, percussionFrom.Glass)
	}
	if percussionFrom.Metal != nil {
		percussionTo.Metal = CopyBranchMetal(mapOrigCopy, percussionFrom.Metal)
	}
	if percussionFrom.Wood != nil {
		percussionTo.Wood = CopyBranchWood(mapOrigCopy, percussionFrom.Wood)
	}
	if percussionFrom.Pitched != nil {
		percussionTo.Pitched = CopyBranchPitched(mapOrigCopy, percussionFrom.Pitched)
	}
	if percussionFrom.Membrane != nil {
		percussionTo.Membrane = CopyBranchMembrane(mapOrigCopy, percussionFrom.Membrane)
	}
	if percussionFrom.Effect != nil {
		percussionTo.Effect = CopyBranchEffect(mapOrigCopy, percussionFrom.Effect)
	}
	if percussionFrom.Timpani != nil {
		percussionTo.Timpani = CopyBranchTimpani(mapOrigCopy, percussionFrom.Timpani)
	}
	if percussionFrom.Beater != nil {
		percussionTo.Beater = CopyBranchBeater(mapOrigCopy, percussionFrom.Beater)
	}
	if percussionFrom.Stick != nil {
		percussionTo.Stick = CopyBranchStick(mapOrigCopy, percussionFrom.Stick)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPitch(mapOrigCopy map[any]any, pitchFrom *Pitch) (pitchTo *Pitch) {

	// pitchFrom has already been copied
	if _pitchTo, ok := mapOrigCopy[pitchFrom]; ok {
		pitchTo = _pitchTo.(*Pitch)
		return
	}

	pitchTo = new(Pitch)
	mapOrigCopy[pitchFrom] = pitchTo
	pitchFrom.CopyBasicFields(pitchTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPitched(mapOrigCopy map[any]any, pitchedFrom *Pitched) (pitchedTo *Pitched) {

	// pitchedFrom has already been copied
	if _pitchedTo, ok := mapOrigCopy[pitchedFrom]; ok {
		pitchedTo = _pitchedTo.(*Pitched)
		return
	}

	pitchedTo = new(Pitched)
	mapOrigCopy[pitchedFrom] = pitchedTo
	pitchedFrom.CopyBasicFields(pitchedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPlay(mapOrigCopy map[any]any, playFrom *Play) (playTo *Play) {

	// playFrom has already been copied
	if _playTo, ok := mapOrigCopy[playFrom]; ok {
		playTo = _playTo.(*Play)
		return
	}

	playTo = new(Play)
	mapOrigCopy[playFrom] = playTo
	playFrom.CopyBasicFields(playTo)

	//insertion point for the staging of instances referenced by pointers
	if playFrom.Other_play != nil {
		playTo.Other_play = CopyBranchOther_play(mapOrigCopy, playFrom.Other_play)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPlayer(mapOrigCopy map[any]any, playerFrom *Player) (playerTo *Player) {

	// playerFrom has already been copied
	if _playerTo, ok := mapOrigCopy[playerFrom]; ok {
		playerTo = _playerTo.(*Player)
		return
	}

	playerTo = new(Player)
	mapOrigCopy[playerFrom] = playerTo
	playerFrom.CopyBasicFields(playerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPrincipal_voice(mapOrigCopy map[any]any, principal_voiceFrom *Principal_voice) (principal_voiceTo *Principal_voice) {

	// principal_voiceFrom has already been copied
	if _principal_voiceTo, ok := mapOrigCopy[principal_voiceFrom]; ok {
		principal_voiceTo = _principal_voiceTo.(*Principal_voice)
		return
	}

	principal_voiceTo = new(Principal_voice)
	mapOrigCopy[principal_voiceFrom] = principal_voiceTo
	principal_voiceFrom.CopyBasicFields(principal_voiceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPrint(mapOrigCopy map[any]any, printFrom *Print) (printTo *Print) {

	// printFrom has already been copied
	if _printTo, ok := mapOrigCopy[printFrom]; ok {
		printTo = _printTo.(*Print)
		return
	}

	printTo = new(Print)
	mapOrigCopy[printFrom] = printTo
	printFrom.CopyBasicFields(printTo)

	//insertion point for the staging of instances referenced by pointers
	if printFrom.Measure_layout != nil {
		printTo.Measure_layout = CopyBranchMeasure_layout(mapOrigCopy, printFrom.Measure_layout)
	}
	if printFrom.Measure_numbering != nil {
		printTo.Measure_numbering = CopyBranchMeasure_numbering(mapOrigCopy, printFrom.Measure_numbering)
	}
	if printFrom.Part_name_display != nil {
		printTo.Part_name_display = CopyBranchName_display(mapOrigCopy, printFrom.Part_name_display)
	}
	if printFrom.Part_abbreviation_display != nil {
		printTo.Part_abbreviation_display = CopyBranchName_display(mapOrigCopy, printFrom.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRelease(mapOrigCopy map[any]any, releaseFrom *Release) (releaseTo *Release) {

	// releaseFrom has already been copied
	if _releaseTo, ok := mapOrigCopy[releaseFrom]; ok {
		releaseTo = _releaseTo.(*Release)
		return
	}

	releaseTo = new(Release)
	mapOrigCopy[releaseFrom] = releaseTo
	releaseFrom.CopyBasicFields(releaseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRepeat(mapOrigCopy map[any]any, repeatFrom *Repeat) (repeatTo *Repeat) {

	// repeatFrom has already been copied
	if _repeatTo, ok := mapOrigCopy[repeatFrom]; ok {
		repeatTo = _repeatTo.(*Repeat)
		return
	}

	repeatTo = new(Repeat)
	mapOrigCopy[repeatFrom] = repeatTo
	repeatFrom.CopyBasicFields(repeatTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRest(mapOrigCopy map[any]any, restFrom *Rest) (restTo *Rest) {

	// restFrom has already been copied
	if _restTo, ok := mapOrigCopy[restFrom]; ok {
		restTo = _restTo.(*Rest)
		return
	}

	restTo = new(Rest)
	mapOrigCopy[restFrom] = restTo
	restFrom.CopyBasicFields(restTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRoot(mapOrigCopy map[any]any, rootFrom *Root) (rootTo *Root) {

	// rootFrom has already been copied
	if _rootTo, ok := mapOrigCopy[rootFrom]; ok {
		rootTo = _rootTo.(*Root)
		return
	}

	rootTo = new(Root)
	mapOrigCopy[rootFrom] = rootTo
	rootFrom.CopyBasicFields(rootTo)

	//insertion point for the staging of instances referenced by pointers
	if rootFrom.Root_step != nil {
		rootTo.Root_step = CopyBranchRoot_step(mapOrigCopy, rootFrom.Root_step)
	}
	if rootFrom.Root_alter != nil {
		rootTo.Root_alter = CopyBranchHarmony_alter(mapOrigCopy, rootFrom.Root_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRoot_step(mapOrigCopy map[any]any, root_stepFrom *Root_step) (root_stepTo *Root_step) {

	// root_stepFrom has already been copied
	if _root_stepTo, ok := mapOrigCopy[root_stepFrom]; ok {
		root_stepTo = _root_stepTo.(*Root_step)
		return
	}

	root_stepTo = new(Root_step)
	mapOrigCopy[root_stepFrom] = root_stepTo
	root_stepFrom.CopyBasicFields(root_stepTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchScaling(mapOrigCopy map[any]any, scalingFrom *Scaling) (scalingTo *Scaling) {

	// scalingFrom has already been copied
	if _scalingTo, ok := mapOrigCopy[scalingFrom]; ok {
		scalingTo = _scalingTo.(*Scaling)
		return
	}

	scalingTo = new(Scaling)
	mapOrigCopy[scalingFrom] = scalingTo
	scalingFrom.CopyBasicFields(scalingTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchScordatura(mapOrigCopy map[any]any, scordaturaFrom *Scordatura) (scordaturaTo *Scordatura) {

	// scordaturaFrom has already been copied
	if _scordaturaTo, ok := mapOrigCopy[scordaturaFrom]; ok {
		scordaturaTo = _scordaturaTo.(*Scordatura)
		return
	}

	scordaturaTo = new(Scordatura)
	mapOrigCopy[scordaturaFrom] = scordaturaTo
	scordaturaFrom.CopyBasicFields(scordaturaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordaturaFrom.Accord {
		scordaturaTo.Accord = append(scordaturaTo.Accord, CopyBranchAccord(mapOrigCopy, _accord))
	}

	return
}

func CopyBranchScore_instrument(mapOrigCopy map[any]any, score_instrumentFrom *Score_instrument) (score_instrumentTo *Score_instrument) {

	// score_instrumentFrom has already been copied
	if _score_instrumentTo, ok := mapOrigCopy[score_instrumentFrom]; ok {
		score_instrumentTo = _score_instrumentTo.(*Score_instrument)
		return
	}

	score_instrumentTo = new(Score_instrument)
	mapOrigCopy[score_instrumentFrom] = score_instrumentTo
	score_instrumentFrom.CopyBasicFields(score_instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchScore_part(mapOrigCopy map[any]any, score_partFrom *Score_part) (score_partTo *Score_part) {

	// score_partFrom has already been copied
	if _score_partTo, ok := mapOrigCopy[score_partFrom]; ok {
		score_partTo = _score_partTo.(*Score_part)
		return
	}

	score_partTo = new(Score_part)
	mapOrigCopy[score_partFrom] = score_partTo
	score_partFrom.CopyBasicFields(score_partTo)

	//insertion point for the staging of instances referenced by pointers
	if score_partFrom.Identification != nil {
		score_partTo.Identification = CopyBranchIdentification(mapOrigCopy, score_partFrom.Identification)
	}
	if score_partFrom.Part_name_display != nil {
		score_partTo.Part_name_display = CopyBranchName_display(mapOrigCopy, score_partFrom.Part_name_display)
	}
	if score_partFrom.Part_abbreviation_display != nil {
		score_partTo.Part_abbreviation_display = CopyBranchName_display(mapOrigCopy, score_partFrom.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part_link := range score_partFrom.Part_link {
		score_partTo.Part_link = append(score_partTo.Part_link, CopyBranchPart_link(mapOrigCopy, _part_link))
	}
	for _, _score_instrument := range score_partFrom.Score_instrument {
		score_partTo.Score_instrument = append(score_partTo.Score_instrument, CopyBranchScore_instrument(mapOrigCopy, _score_instrument))
	}
	for _, _player := range score_partFrom.Player {
		score_partTo.Player = append(score_partTo.Player, CopyBranchPlayer(mapOrigCopy, _player))
	}

	return
}

func CopyBranchScore_partwise(mapOrigCopy map[any]any, score_partwiseFrom *Score_partwise) (score_partwiseTo *Score_partwise) {

	// score_partwiseFrom has already been copied
	if _score_partwiseTo, ok := mapOrigCopy[score_partwiseFrom]; ok {
		score_partwiseTo = _score_partwiseTo.(*Score_partwise)
		return
	}

	score_partwiseTo = new(Score_partwise)
	mapOrigCopy[score_partwiseFrom] = score_partwiseTo
	score_partwiseFrom.CopyBasicFields(score_partwiseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchScore_timewise(mapOrigCopy map[any]any, score_timewiseFrom *Score_timewise) (score_timewiseTo *Score_timewise) {

	// score_timewiseFrom has already been copied
	if _score_timewiseTo, ok := mapOrigCopy[score_timewiseFrom]; ok {
		score_timewiseTo = _score_timewiseTo.(*Score_timewise)
		return
	}

	score_timewiseTo = new(Score_timewise)
	mapOrigCopy[score_timewiseFrom] = score_timewiseTo
	score_timewiseFrom.CopyBasicFields(score_timewiseTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSegno(mapOrigCopy map[any]any, segnoFrom *Segno) (segnoTo *Segno) {

	// segnoFrom has already been copied
	if _segnoTo, ok := mapOrigCopy[segnoFrom]; ok {
		segnoTo = _segnoTo.(*Segno)
		return
	}

	segnoTo = new(Segno)
	mapOrigCopy[segnoFrom] = segnoTo
	segnoFrom.CopyBasicFields(segnoTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSlash(mapOrigCopy map[any]any, slashFrom *Slash) (slashTo *Slash) {

	// slashFrom has already been copied
	if _slashTo, ok := mapOrigCopy[slashFrom]; ok {
		slashTo = _slashTo.(*Slash)
		return
	}

	slashTo = new(Slash)
	mapOrigCopy[slashFrom] = slashTo
	slashFrom.CopyBasicFields(slashTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSlide(mapOrigCopy map[any]any, slideFrom *Slide) (slideTo *Slide) {

	// slideFrom has already been copied
	if _slideTo, ok := mapOrigCopy[slideFrom]; ok {
		slideTo = _slideTo.(*Slide)
		return
	}

	slideTo = new(Slide)
	mapOrigCopy[slideFrom] = slideTo
	slideFrom.CopyBasicFields(slideTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSlur(mapOrigCopy map[any]any, slurFrom *Slur) (slurTo *Slur) {

	// slurFrom has already been copied
	if _slurTo, ok := mapOrigCopy[slurFrom]; ok {
		slurTo = _slurTo.(*Slur)
		return
	}

	slurTo = new(Slur)
	mapOrigCopy[slurFrom] = slurTo
	slurFrom.CopyBasicFields(slurTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSound(mapOrigCopy map[any]any, soundFrom *Sound) (soundTo *Sound) {

	// soundFrom has already been copied
	if _soundTo, ok := mapOrigCopy[soundFrom]; ok {
		soundTo = _soundTo.(*Sound)
		return
	}

	soundTo = new(Sound)
	mapOrigCopy[soundFrom] = soundTo
	soundFrom.CopyBasicFields(soundTo)

	//insertion point for the staging of instances referenced by pointers
	if soundFrom.Swing != nil {
		soundTo.Swing = CopyBranchSwing(mapOrigCopy, soundFrom.Swing)
	}
	if soundFrom.Offset != nil {
		soundTo.Offset = CopyBranchOffset(mapOrigCopy, soundFrom.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStaff_details(mapOrigCopy map[any]any, staff_detailsFrom *Staff_details) (staff_detailsTo *Staff_details) {

	// staff_detailsFrom has already been copied
	if _staff_detailsTo, ok := mapOrigCopy[staff_detailsFrom]; ok {
		staff_detailsTo = _staff_detailsTo.(*Staff_details)
		return
	}

	staff_detailsTo = new(Staff_details)
	mapOrigCopy[staff_detailsFrom] = staff_detailsTo
	staff_detailsFrom.CopyBasicFields(staff_detailsTo)

	//insertion point for the staging of instances referenced by pointers
	if staff_detailsFrom.Staff_size != nil {
		staff_detailsTo.Staff_size = CopyBranchStaff_size(mapOrigCopy, staff_detailsFrom.Staff_size)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_tuning := range staff_detailsFrom.Staff_tuning {
		staff_detailsTo.Staff_tuning = append(staff_detailsTo.Staff_tuning, CopyBranchStaff_tuning(mapOrigCopy, _staff_tuning))
	}

	return
}

func CopyBranchStaff_divide(mapOrigCopy map[any]any, staff_divideFrom *Staff_divide) (staff_divideTo *Staff_divide) {

	// staff_divideFrom has already been copied
	if _staff_divideTo, ok := mapOrigCopy[staff_divideFrom]; ok {
		staff_divideTo = _staff_divideTo.(*Staff_divide)
		return
	}

	staff_divideTo = new(Staff_divide)
	mapOrigCopy[staff_divideFrom] = staff_divideTo
	staff_divideFrom.CopyBasicFields(staff_divideTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStaff_layout(mapOrigCopy map[any]any, staff_layoutFrom *Staff_layout) (staff_layoutTo *Staff_layout) {

	// staff_layoutFrom has already been copied
	if _staff_layoutTo, ok := mapOrigCopy[staff_layoutFrom]; ok {
		staff_layoutTo = _staff_layoutTo.(*Staff_layout)
		return
	}

	staff_layoutTo = new(Staff_layout)
	mapOrigCopy[staff_layoutFrom] = staff_layoutTo
	staff_layoutFrom.CopyBasicFields(staff_layoutTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStaff_size(mapOrigCopy map[any]any, staff_sizeFrom *Staff_size) (staff_sizeTo *Staff_size) {

	// staff_sizeFrom has already been copied
	if _staff_sizeTo, ok := mapOrigCopy[staff_sizeFrom]; ok {
		staff_sizeTo = _staff_sizeTo.(*Staff_size)
		return
	}

	staff_sizeTo = new(Staff_size)
	mapOrigCopy[staff_sizeFrom] = staff_sizeTo
	staff_sizeFrom.CopyBasicFields(staff_sizeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStaff_tuning(mapOrigCopy map[any]any, staff_tuningFrom *Staff_tuning) (staff_tuningTo *Staff_tuning) {

	// staff_tuningFrom has already been copied
	if _staff_tuningTo, ok := mapOrigCopy[staff_tuningFrom]; ok {
		staff_tuningTo = _staff_tuningTo.(*Staff_tuning)
		return
	}

	staff_tuningTo = new(Staff_tuning)
	mapOrigCopy[staff_tuningFrom] = staff_tuningTo
	staff_tuningFrom.CopyBasicFields(staff_tuningTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStem(mapOrigCopy map[any]any, stemFrom *Stem) (stemTo *Stem) {

	// stemFrom has already been copied
	if _stemTo, ok := mapOrigCopy[stemFrom]; ok {
		stemTo = _stemTo.(*Stem)
		return
	}

	stemTo = new(Stem)
	mapOrigCopy[stemFrom] = stemTo
	stemFrom.CopyBasicFields(stemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStick(mapOrigCopy map[any]any, stickFrom *Stick) (stickTo *Stick) {

	// stickFrom has already been copied
	if _stickTo, ok := mapOrigCopy[stickFrom]; ok {
		stickTo = _stickTo.(*Stick)
		return
	}

	stickTo = new(Stick)
	mapOrigCopy[stickFrom] = stickTo
	stickFrom.CopyBasicFields(stickTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchString_mute(mapOrigCopy map[any]any, string_muteFrom *String_mute) (string_muteTo *String_mute) {

	// string_muteFrom has already been copied
	if _string_muteTo, ok := mapOrigCopy[string_muteFrom]; ok {
		string_muteTo = _string_muteTo.(*String_mute)
		return
	}

	string_muteTo = new(String_mute)
	mapOrigCopy[string_muteFrom] = string_muteTo
	string_muteFrom.CopyBasicFields(string_muteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchStrong_accent(mapOrigCopy map[any]any, strong_accentFrom *Strong_accent) (strong_accentTo *Strong_accent) {

	// strong_accentFrom has already been copied
	if _strong_accentTo, ok := mapOrigCopy[strong_accentFrom]; ok {
		strong_accentTo = _strong_accentTo.(*Strong_accent)
		return
	}

	strong_accentTo = new(Strong_accent)
	mapOrigCopy[strong_accentFrom] = strong_accentTo
	strong_accentFrom.CopyBasicFields(strong_accentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSupports(mapOrigCopy map[any]any, supportsFrom *Supports) (supportsTo *Supports) {

	// supportsFrom has already been copied
	if _supportsTo, ok := mapOrigCopy[supportsFrom]; ok {
		supportsTo = _supportsTo.(*Supports)
		return
	}

	supportsTo = new(Supports)
	mapOrigCopy[supportsFrom] = supportsTo
	supportsFrom.CopyBasicFields(supportsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSwing(mapOrigCopy map[any]any, swingFrom *Swing) (swingTo *Swing) {

	// swingFrom has already been copied
	if _swingTo, ok := mapOrigCopy[swingFrom]; ok {
		swingTo = _swingTo.(*Swing)
		return
	}

	swingTo = new(Swing)
	mapOrigCopy[swingFrom] = swingTo
	swingFrom.CopyBasicFields(swingTo)

	//insertion point for the staging of instances referenced by pointers
	if swingFrom.Straight != nil {
		swingTo.Straight = CopyBranchEmpty(mapOrigCopy, swingFrom.Straight)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSync(mapOrigCopy map[any]any, syncFrom *Sync) (syncTo *Sync) {

	// syncFrom has already been copied
	if _syncTo, ok := mapOrigCopy[syncFrom]; ok {
		syncTo = _syncTo.(*Sync)
		return
	}

	syncTo = new(Sync)
	mapOrigCopy[syncFrom] = syncTo
	syncFrom.CopyBasicFields(syncTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSystem_dividers(mapOrigCopy map[any]any, system_dividersFrom *System_dividers) (system_dividersTo *System_dividers) {

	// system_dividersFrom has already been copied
	if _system_dividersTo, ok := mapOrigCopy[system_dividersFrom]; ok {
		system_dividersTo = _system_dividersTo.(*System_dividers)
		return
	}

	system_dividersTo = new(System_dividers)
	mapOrigCopy[system_dividersFrom] = system_dividersTo
	system_dividersFrom.CopyBasicFields(system_dividersTo)

	//insertion point for the staging of instances referenced by pointers
	if system_dividersFrom.Left_divider != nil {
		system_dividersTo.Left_divider = CopyBranchEmpty_print_object_style_align(mapOrigCopy, system_dividersFrom.Left_divider)
	}
	if system_dividersFrom.Right_divider != nil {
		system_dividersTo.Right_divider = CopyBranchEmpty_print_object_style_align(mapOrigCopy, system_dividersFrom.Right_divider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSystem_layout(mapOrigCopy map[any]any, system_layoutFrom *System_layout) (system_layoutTo *System_layout) {

	// system_layoutFrom has already been copied
	if _system_layoutTo, ok := mapOrigCopy[system_layoutFrom]; ok {
		system_layoutTo = _system_layoutTo.(*System_layout)
		return
	}

	system_layoutTo = new(System_layout)
	mapOrigCopy[system_layoutFrom] = system_layoutTo
	system_layoutFrom.CopyBasicFields(system_layoutTo)

	//insertion point for the staging of instances referenced by pointers
	if system_layoutFrom.System_margins != nil {
		system_layoutTo.System_margins = CopyBranchSystem_margins(mapOrigCopy, system_layoutFrom.System_margins)
	}
	if system_layoutFrom.System_dividers != nil {
		system_layoutTo.System_dividers = CopyBranchSystem_dividers(mapOrigCopy, system_layoutFrom.System_dividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSystem_margins(mapOrigCopy map[any]any, system_marginsFrom *System_margins) (system_marginsTo *System_margins) {

	// system_marginsFrom has already been copied
	if _system_marginsTo, ok := mapOrigCopy[system_marginsFrom]; ok {
		system_marginsTo = _system_marginsTo.(*System_margins)
		return
	}

	system_marginsTo = new(System_margins)
	mapOrigCopy[system_marginsFrom] = system_marginsTo
	system_marginsFrom.CopyBasicFields(system_marginsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTap(mapOrigCopy map[any]any, tapFrom *Tap) (tapTo *Tap) {

	// tapFrom has already been copied
	if _tapTo, ok := mapOrigCopy[tapFrom]; ok {
		tapTo = _tapTo.(*Tap)
		return
	}

	tapTo = new(Tap)
	mapOrigCopy[tapFrom] = tapTo
	tapFrom.CopyBasicFields(tapTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTechnical(mapOrigCopy map[any]any, technicalFrom *Technical) (technicalTo *Technical) {

	// technicalFrom has already been copied
	if _technicalTo, ok := mapOrigCopy[technicalFrom]; ok {
		technicalTo = _technicalTo.(*Technical)
		return
	}

	technicalTo = new(Technical)
	mapOrigCopy[technicalFrom] = technicalTo
	technicalFrom.CopyBasicFields(technicalTo)

	//insertion point for the staging of instances referenced by pointers
	if technicalFrom.Up_bow != nil {
		technicalTo.Up_bow = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Up_bow)
	}
	if technicalFrom.Down_bow != nil {
		technicalTo.Down_bow = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Down_bow)
	}
	if technicalFrom.Harmonic != nil {
		technicalTo.Harmonic = CopyBranchHarmonic(mapOrigCopy, technicalFrom.Harmonic)
	}
	if technicalFrom.Open_string != nil {
		technicalTo.Open_string = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Open_string)
	}
	if technicalFrom.Thumb_position != nil {
		technicalTo.Thumb_position = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Thumb_position)
	}
	if technicalFrom.Fingering != nil {
		technicalTo.Fingering = CopyBranchFingering(mapOrigCopy, technicalFrom.Fingering)
	}
	if technicalFrom.Double_tongue != nil {
		technicalTo.Double_tongue = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Double_tongue)
	}
	if technicalFrom.Triple_tongue != nil {
		technicalTo.Triple_tongue = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Triple_tongue)
	}
	if technicalFrom.Stopped != nil {
		technicalTo.Stopped = CopyBranchEmpty_placement_smufl(mapOrigCopy, technicalFrom.Stopped)
	}
	if technicalFrom.Snap_pizzicato != nil {
		technicalTo.Snap_pizzicato = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Snap_pizzicato)
	}
	if technicalFrom.Fret != nil {
		technicalTo.Fret = CopyBranchFret(mapOrigCopy, technicalFrom.Fret)
	}
	if technicalFrom.Hammer_on != nil {
		technicalTo.Hammer_on = CopyBranchHammer_on_pull_off(mapOrigCopy, technicalFrom.Hammer_on)
	}
	if technicalFrom.Pull_off != nil {
		technicalTo.Pull_off = CopyBranchHammer_on_pull_off(mapOrigCopy, technicalFrom.Pull_off)
	}
	if technicalFrom.Bend != nil {
		technicalTo.Bend = CopyBranchBend(mapOrigCopy, technicalFrom.Bend)
	}
	if technicalFrom.Tap != nil {
		technicalTo.Tap = CopyBranchTap(mapOrigCopy, technicalFrom.Tap)
	}
	if technicalFrom.Heel != nil {
		technicalTo.Heel = CopyBranchHeel_toe(mapOrigCopy, technicalFrom.Heel)
	}
	if technicalFrom.Toe != nil {
		technicalTo.Toe = CopyBranchHeel_toe(mapOrigCopy, technicalFrom.Toe)
	}
	if technicalFrom.Fingernails != nil {
		technicalTo.Fingernails = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Fingernails)
	}
	if technicalFrom.Hole != nil {
		technicalTo.Hole = CopyBranchHole(mapOrigCopy, technicalFrom.Hole)
	}
	if technicalFrom.Arrow != nil {
		technicalTo.Arrow = CopyBranchArrow(mapOrigCopy, technicalFrom.Arrow)
	}
	if technicalFrom.Handbell != nil {
		technicalTo.Handbell = CopyBranchHandbell(mapOrigCopy, technicalFrom.Handbell)
	}
	if technicalFrom.Brass_bend != nil {
		technicalTo.Brass_bend = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Brass_bend)
	}
	if technicalFrom.Flip != nil {
		technicalTo.Flip = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Flip)
	}
	if technicalFrom.Smear != nil {
		technicalTo.Smear = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Smear)
	}
	if technicalFrom.Open != nil {
		technicalTo.Open = CopyBranchEmpty_placement_smufl(mapOrigCopy, technicalFrom.Open)
	}
	if technicalFrom.Half_muted != nil {
		technicalTo.Half_muted = CopyBranchEmpty_placement_smufl(mapOrigCopy, technicalFrom.Half_muted)
	}
	if technicalFrom.Harmon_mute != nil {
		technicalTo.Harmon_mute = CopyBranchHarmon_mute(mapOrigCopy, technicalFrom.Harmon_mute)
	}
	if technicalFrom.Golpe != nil {
		technicalTo.Golpe = CopyBranchEmpty_placement(mapOrigCopy, technicalFrom.Golpe)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchText_element_data(mapOrigCopy map[any]any, text_element_dataFrom *Text_element_data) (text_element_dataTo *Text_element_data) {

	// text_element_dataFrom has already been copied
	if _text_element_dataTo, ok := mapOrigCopy[text_element_dataFrom]; ok {
		text_element_dataTo = _text_element_dataTo.(*Text_element_data)
		return
	}

	text_element_dataTo = new(Text_element_data)
	mapOrigCopy[text_element_dataFrom] = text_element_dataTo
	text_element_dataFrom.CopyBasicFields(text_element_dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTie(mapOrigCopy map[any]any, tieFrom *Tie) (tieTo *Tie) {

	// tieFrom has already been copied
	if _tieTo, ok := mapOrigCopy[tieFrom]; ok {
		tieTo = _tieTo.(*Tie)
		return
	}

	tieTo = new(Tie)
	mapOrigCopy[tieFrom] = tieTo
	tieFrom.CopyBasicFields(tieTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTied(mapOrigCopy map[any]any, tiedFrom *Tied) (tiedTo *Tied) {

	// tiedFrom has already been copied
	if _tiedTo, ok := mapOrigCopy[tiedFrom]; ok {
		tiedTo = _tiedTo.(*Tied)
		return
	}

	tiedTo = new(Tied)
	mapOrigCopy[tiedFrom] = tiedTo
	tiedFrom.CopyBasicFields(tiedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTime(mapOrigCopy map[any]any, timeFrom *Time) (timeTo *Time) {

	// timeFrom has already been copied
	if _timeTo, ok := mapOrigCopy[timeFrom]; ok {
		timeTo = _timeTo.(*Time)
		return
	}

	timeTo = new(Time)
	mapOrigCopy[timeFrom] = timeTo
	timeFrom.CopyBasicFields(timeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTime_modification(mapOrigCopy map[any]any, time_modificationFrom *Time_modification) (time_modificationTo *Time_modification) {

	// time_modificationFrom has already been copied
	if _time_modificationTo, ok := mapOrigCopy[time_modificationFrom]; ok {
		time_modificationTo = _time_modificationTo.(*Time_modification)
		return
	}

	time_modificationTo = new(Time_modification)
	mapOrigCopy[time_modificationFrom] = time_modificationTo
	time_modificationFrom.CopyBasicFields(time_modificationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTimpani(mapOrigCopy map[any]any, timpaniFrom *Timpani) (timpaniTo *Timpani) {

	// timpaniFrom has already been copied
	if _timpaniTo, ok := mapOrigCopy[timpaniFrom]; ok {
		timpaniTo = _timpaniTo.(*Timpani)
		return
	}

	timpaniTo = new(Timpani)
	mapOrigCopy[timpaniFrom] = timpaniTo
	timpaniFrom.CopyBasicFields(timpaniTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTranspose(mapOrigCopy map[any]any, transposeFrom *Transpose) (transposeTo *Transpose) {

	// transposeFrom has already been copied
	if _transposeTo, ok := mapOrigCopy[transposeFrom]; ok {
		transposeTo = _transposeTo.(*Transpose)
		return
	}

	transposeTo = new(Transpose)
	mapOrigCopy[transposeFrom] = transposeTo
	transposeFrom.CopyBasicFields(transposeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTremolo(mapOrigCopy map[any]any, tremoloFrom *Tremolo) (tremoloTo *Tremolo) {

	// tremoloFrom has already been copied
	if _tremoloTo, ok := mapOrigCopy[tremoloFrom]; ok {
		tremoloTo = _tremoloTo.(*Tremolo)
		return
	}

	tremoloTo = new(Tremolo)
	mapOrigCopy[tremoloFrom] = tremoloTo
	tremoloFrom.CopyBasicFields(tremoloTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTuplet(mapOrigCopy map[any]any, tupletFrom *Tuplet) (tupletTo *Tuplet) {

	// tupletFrom has already been copied
	if _tupletTo, ok := mapOrigCopy[tupletFrom]; ok {
		tupletTo = _tupletTo.(*Tuplet)
		return
	}

	tupletTo = new(Tuplet)
	mapOrigCopy[tupletFrom] = tupletTo
	tupletFrom.CopyBasicFields(tupletTo)

	//insertion point for the staging of instances referenced by pointers
	if tupletFrom.Tuplet_actual != nil {
		tupletTo.Tuplet_actual = CopyBranchTuplet_portion(mapOrigCopy, tupletFrom.Tuplet_actual)
	}
	if tupletFrom.Tuplet_normal != nil {
		tupletTo.Tuplet_normal = CopyBranchTuplet_portion(mapOrigCopy, tupletFrom.Tuplet_normal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTuplet_dot(mapOrigCopy map[any]any, tuplet_dotFrom *Tuplet_dot) (tuplet_dotTo *Tuplet_dot) {

	// tuplet_dotFrom has already been copied
	if _tuplet_dotTo, ok := mapOrigCopy[tuplet_dotFrom]; ok {
		tuplet_dotTo = _tuplet_dotTo.(*Tuplet_dot)
		return
	}

	tuplet_dotTo = new(Tuplet_dot)
	mapOrigCopy[tuplet_dotFrom] = tuplet_dotTo
	tuplet_dotFrom.CopyBasicFields(tuplet_dotTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTuplet_number(mapOrigCopy map[any]any, tuplet_numberFrom *Tuplet_number) (tuplet_numberTo *Tuplet_number) {

	// tuplet_numberFrom has already been copied
	if _tuplet_numberTo, ok := mapOrigCopy[tuplet_numberFrom]; ok {
		tuplet_numberTo = _tuplet_numberTo.(*Tuplet_number)
		return
	}

	tuplet_numberTo = new(Tuplet_number)
	mapOrigCopy[tuplet_numberFrom] = tuplet_numberTo
	tuplet_numberFrom.CopyBasicFields(tuplet_numberTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTuplet_portion(mapOrigCopy map[any]any, tuplet_portionFrom *Tuplet_portion) (tuplet_portionTo *Tuplet_portion) {

	// tuplet_portionFrom has already been copied
	if _tuplet_portionTo, ok := mapOrigCopy[tuplet_portionFrom]; ok {
		tuplet_portionTo = _tuplet_portionTo.(*Tuplet_portion)
		return
	}

	tuplet_portionTo = new(Tuplet_portion)
	mapOrigCopy[tuplet_portionFrom] = tuplet_portionTo
	tuplet_portionFrom.CopyBasicFields(tuplet_portionTo)

	//insertion point for the staging of instances referenced by pointers
	if tuplet_portionFrom.Tuplet_number != nil {
		tuplet_portionTo.Tuplet_number = CopyBranchTuplet_number(mapOrigCopy, tuplet_portionFrom.Tuplet_number)
	}
	if tuplet_portionFrom.Tuplet_type != nil {
		tuplet_portionTo.Tuplet_type = CopyBranchTuplet_type(mapOrigCopy, tuplet_portionFrom.Tuplet_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tuplet_dot := range tuplet_portionFrom.Tuplet_dot {
		tuplet_portionTo.Tuplet_dot = append(tuplet_portionTo.Tuplet_dot, CopyBranchTuplet_dot(mapOrigCopy, _tuplet_dot))
	}

	return
}

func CopyBranchTuplet_type(mapOrigCopy map[any]any, tuplet_typeFrom *Tuplet_type) (tuplet_typeTo *Tuplet_type) {

	// tuplet_typeFrom has already been copied
	if _tuplet_typeTo, ok := mapOrigCopy[tuplet_typeFrom]; ok {
		tuplet_typeTo = _tuplet_typeTo.(*Tuplet_type)
		return
	}

	tuplet_typeTo = new(Tuplet_type)
	mapOrigCopy[tuplet_typeFrom] = tuplet_typeTo
	tuplet_typeFrom.CopyBasicFields(tuplet_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTyped_text(mapOrigCopy map[any]any, typed_textFrom *Typed_text) (typed_textTo *Typed_text) {

	// typed_textFrom has already been copied
	if _typed_textTo, ok := mapOrigCopy[typed_textFrom]; ok {
		typed_textTo = _typed_textTo.(*Typed_text)
		return
	}

	typed_textTo = new(Typed_text)
	mapOrigCopy[typed_textFrom] = typed_textTo
	typed_textFrom.CopyBasicFields(typed_textTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchUnpitched(mapOrigCopy map[any]any, unpitchedFrom *Unpitched) (unpitchedTo *Unpitched) {

	// unpitchedFrom has already been copied
	if _unpitchedTo, ok := mapOrigCopy[unpitchedFrom]; ok {
		unpitchedTo = _unpitchedTo.(*Unpitched)
		return
	}

	unpitchedTo = new(Unpitched)
	mapOrigCopy[unpitchedFrom] = unpitchedTo
	unpitchedFrom.CopyBasicFields(unpitchedTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchVirtual_instrument(mapOrigCopy map[any]any, virtual_instrumentFrom *Virtual_instrument) (virtual_instrumentTo *Virtual_instrument) {

	// virtual_instrumentFrom has already been copied
	if _virtual_instrumentTo, ok := mapOrigCopy[virtual_instrumentFrom]; ok {
		virtual_instrumentTo = _virtual_instrumentTo.(*Virtual_instrument)
		return
	}

	virtual_instrumentTo = new(Virtual_instrument)
	mapOrigCopy[virtual_instrumentFrom] = virtual_instrumentTo
	virtual_instrumentFrom.CopyBasicFields(virtual_instrumentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWait(mapOrigCopy map[any]any, waitFrom *Wait) (waitTo *Wait) {

	// waitFrom has already been copied
	if _waitTo, ok := mapOrigCopy[waitFrom]; ok {
		waitTo = _waitTo.(*Wait)
		return
	}

	waitTo = new(Wait)
	mapOrigCopy[waitFrom] = waitTo
	waitFrom.CopyBasicFields(waitTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWavy_line(mapOrigCopy map[any]any, wavy_lineFrom *Wavy_line) (wavy_lineTo *Wavy_line) {

	// wavy_lineFrom has already been copied
	if _wavy_lineTo, ok := mapOrigCopy[wavy_lineFrom]; ok {
		wavy_lineTo = _wavy_lineTo.(*Wavy_line)
		return
	}

	wavy_lineTo = new(Wavy_line)
	mapOrigCopy[wavy_lineFrom] = wavy_lineTo
	wavy_lineFrom.CopyBasicFields(wavy_lineTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWedge(mapOrigCopy map[any]any, wedgeFrom *Wedge) (wedgeTo *Wedge) {

	// wedgeFrom has already been copied
	if _wedgeTo, ok := mapOrigCopy[wedgeFrom]; ok {
		wedgeTo = _wedgeTo.(*Wedge)
		return
	}

	wedgeTo = new(Wedge)
	mapOrigCopy[wedgeFrom] = wedgeTo
	wedgeFrom.CopyBasicFields(wedgeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWood(mapOrigCopy map[any]any, woodFrom *Wood) (woodTo *Wood) {

	// woodFrom has already been copied
	if _woodTo, ok := mapOrigCopy[woodFrom]; ok {
		woodTo = _woodTo.(*Wood)
		return
	}

	woodTo = new(Wood)
	mapOrigCopy[woodFrom] = woodTo
	woodFrom.CopyBasicFields(woodTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchWork(mapOrigCopy map[any]any, workFrom *Work) (workTo *Work) {

	// workFrom has already been copied
	if _workTo, ok := mapOrigCopy[workFrom]; ok {
		workTo = _workTo.(*Work)
		return
	}

	workTo = new(Work)
	mapOrigCopy[workFrom] = workTo
	workFrom.CopyBasicFields(workTo)

	//insertion point for the staging of instances referenced by pointers
	if workFrom.Opus != nil {
		workTo.Opus = CopyBranchOpus(mapOrigCopy, workFrom.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Accidental:
		stage.UnstageBranchAccidental(target)

	case *Accidental_mark:
		stage.UnstageBranchAccidental_mark(target)

	case *Accidental_text:
		stage.UnstageBranchAccidental_text(target)

	case *Accord:
		stage.UnstageBranchAccord(target)

	case *Accordion_registration:
		stage.UnstageBranchAccordion_registration(target)

	case *AnyType:
		stage.UnstageBranchAnyType(target)

	case *Appearance:
		stage.UnstageBranchAppearance(target)

	case *Arpeggiate:
		stage.UnstageBranchArpeggiate(target)

	case *Arrow:
		stage.UnstageBranchArrow(target)

	case *Articulations:
		stage.UnstageBranchArticulations(target)

	case *Assess:
		stage.UnstageBranchAssess(target)

	case *Attributes:
		stage.UnstageBranchAttributes(target)

	case *Backup:
		stage.UnstageBranchBackup(target)

	case *Bar_style_color:
		stage.UnstageBranchBar_style_color(target)

	case *Barline:
		stage.UnstageBranchBarline(target)

	case *Barre:
		stage.UnstageBranchBarre(target)

	case *Bass:
		stage.UnstageBranchBass(target)

	case *Bass_step:
		stage.UnstageBranchBass_step(target)

	case *Beam:
		stage.UnstageBranchBeam(target)

	case *Beat_repeat:
		stage.UnstageBranchBeat_repeat(target)

	case *Beat_unit_tied:
		stage.UnstageBranchBeat_unit_tied(target)

	case *Beater:
		stage.UnstageBranchBeater(target)

	case *Bend:
		stage.UnstageBranchBend(target)

	case *Bookmark:
		stage.UnstageBranchBookmark(target)

	case *Bracket:
		stage.UnstageBranchBracket(target)

	case *Breath_mark:
		stage.UnstageBranchBreath_mark(target)

	case *Caesura:
		stage.UnstageBranchCaesura(target)

	case *Cancel:
		stage.UnstageBranchCancel(target)

	case *Clef:
		stage.UnstageBranchClef(target)

	case *Coda:
		stage.UnstageBranchCoda(target)

	case *Credit:
		stage.UnstageBranchCredit(target)

	case *Dashes:
		stage.UnstageBranchDashes(target)

	case *Defaults:
		stage.UnstageBranchDefaults(target)

	case *Degree:
		stage.UnstageBranchDegree(target)

	case *Degree_alter:
		stage.UnstageBranchDegree_alter(target)

	case *Degree_type:
		stage.UnstageBranchDegree_type(target)

	case *Degree_value:
		stage.UnstageBranchDegree_value(target)

	case *Direction:
		stage.UnstageBranchDirection(target)

	case *Direction_type:
		stage.UnstageBranchDirection_type(target)

	case *Distance:
		stage.UnstageBranchDistance(target)

	case *Double:
		stage.UnstageBranchDouble(target)

	case *Dynamics:
		stage.UnstageBranchDynamics(target)

	case *Effect:
		stage.UnstageBranchEffect(target)

	case *Elision:
		stage.UnstageBranchElision(target)

	case *Empty:
		stage.UnstageBranchEmpty(target)

	case *Empty_font:
		stage.UnstageBranchEmpty_font(target)

	case *Empty_line:
		stage.UnstageBranchEmpty_line(target)

	case *Empty_placement:
		stage.UnstageBranchEmpty_placement(target)

	case *Empty_placement_smufl:
		stage.UnstageBranchEmpty_placement_smufl(target)

	case *Empty_print_object_style_align:
		stage.UnstageBranchEmpty_print_object_style_align(target)

	case *Empty_print_style:
		stage.UnstageBranchEmpty_print_style(target)

	case *Empty_print_style_align:
		stage.UnstageBranchEmpty_print_style_align(target)

	case *Empty_print_style_align_id:
		stage.UnstageBranchEmpty_print_style_align_id(target)

	case *Empty_trill_sound:
		stage.UnstageBranchEmpty_trill_sound(target)

	case *Encoding:
		stage.UnstageBranchEncoding(target)

	case *Ending:
		stage.UnstageBranchEnding(target)

	case *Extend:
		stage.UnstageBranchExtend(target)

	case *Feature:
		stage.UnstageBranchFeature(target)

	case *Fermata:
		stage.UnstageBranchFermata(target)

	case *Figure:
		stage.UnstageBranchFigure(target)

	case *Figured_bass:
		stage.UnstageBranchFigured_bass(target)

	case *Fingering:
		stage.UnstageBranchFingering(target)

	case *First_fret:
		stage.UnstageBranchFirst_fret(target)

	case *Foo:
		stage.UnstageBranchFoo(target)

	case *For_part:
		stage.UnstageBranchFor_part(target)

	case *Formatted_symbol:
		stage.UnstageBranchFormatted_symbol(target)

	case *Formatted_symbol_id:
		stage.UnstageBranchFormatted_symbol_id(target)

	case *Forward:
		stage.UnstageBranchForward(target)

	case *Frame:
		stage.UnstageBranchFrame(target)

	case *Frame_note:
		stage.UnstageBranchFrame_note(target)

	case *Fret:
		stage.UnstageBranchFret(target)

	case *Glass:
		stage.UnstageBranchGlass(target)

	case *Glissando:
		stage.UnstageBranchGlissando(target)

	case *Glyph:
		stage.UnstageBranchGlyph(target)

	case *Grace:
		stage.UnstageBranchGrace(target)

	case *Group_barline:
		stage.UnstageBranchGroup_barline(target)

	case *Group_symbol:
		stage.UnstageBranchGroup_symbol(target)

	case *Grouping:
		stage.UnstageBranchGrouping(target)

	case *Hammer_on_pull_off:
		stage.UnstageBranchHammer_on_pull_off(target)

	case *Handbell:
		stage.UnstageBranchHandbell(target)

	case *Harmon_closed:
		stage.UnstageBranchHarmon_closed(target)

	case *Harmon_mute:
		stage.UnstageBranchHarmon_mute(target)

	case *Harmonic:
		stage.UnstageBranchHarmonic(target)

	case *Harmony:
		stage.UnstageBranchHarmony(target)

	case *Harmony_alter:
		stage.UnstageBranchHarmony_alter(target)

	case *Harp_pedals:
		stage.UnstageBranchHarp_pedals(target)

	case *Heel_toe:
		stage.UnstageBranchHeel_toe(target)

	case *Hole:
		stage.UnstageBranchHole(target)

	case *Hole_closed:
		stage.UnstageBranchHole_closed(target)

	case *Horizontal_turn:
		stage.UnstageBranchHorizontal_turn(target)

	case *Identification:
		stage.UnstageBranchIdentification(target)

	case *Image:
		stage.UnstageBranchImage(target)

	case *Instrument:
		stage.UnstageBranchInstrument(target)

	case *Instrument_change:
		stage.UnstageBranchInstrument_change(target)

	case *Instrument_link:
		stage.UnstageBranchInstrument_link(target)

	case *Interchangeable:
		stage.UnstageBranchInterchangeable(target)

	case *Inversion:
		stage.UnstageBranchInversion(target)

	case *Key:
		stage.UnstageBranchKey(target)

	case *Key_accidental:
		stage.UnstageBranchKey_accidental(target)

	case *Key_octave:
		stage.UnstageBranchKey_octave(target)

	case *Kind:
		stage.UnstageBranchKind(target)

	case *Level:
		stage.UnstageBranchLevel(target)

	case *Line_detail:
		stage.UnstageBranchLine_detail(target)

	case *Line_width:
		stage.UnstageBranchLine_width(target)

	case *Link:
		stage.UnstageBranchLink(target)

	case *Listen:
		stage.UnstageBranchListen(target)

	case *Listening:
		stage.UnstageBranchListening(target)

	case *Lyric:
		stage.UnstageBranchLyric(target)

	case *Lyric_font:
		stage.UnstageBranchLyric_font(target)

	case *Lyric_language:
		stage.UnstageBranchLyric_language(target)

	case *Measure_layout:
		stage.UnstageBranchMeasure_layout(target)

	case *Measure_numbering:
		stage.UnstageBranchMeasure_numbering(target)

	case *Measure_repeat:
		stage.UnstageBranchMeasure_repeat(target)

	case *Measure_style:
		stage.UnstageBranchMeasure_style(target)

	case *Membrane:
		stage.UnstageBranchMembrane(target)

	case *Metal:
		stage.UnstageBranchMetal(target)

	case *Metronome:
		stage.UnstageBranchMetronome(target)

	case *Metronome_beam:
		stage.UnstageBranchMetronome_beam(target)

	case *Metronome_note:
		stage.UnstageBranchMetronome_note(target)

	case *Metronome_tied:
		stage.UnstageBranchMetronome_tied(target)

	case *Metronome_tuplet:
		stage.UnstageBranchMetronome_tuplet(target)

	case *Midi_device:
		stage.UnstageBranchMidi_device(target)

	case *Midi_instrument:
		stage.UnstageBranchMidi_instrument(target)

	case *Miscellaneous:
		stage.UnstageBranchMiscellaneous(target)

	case *Miscellaneous_field:
		stage.UnstageBranchMiscellaneous_field(target)

	case *Mordent:
		stage.UnstageBranchMordent(target)

	case *Multiple_rest:
		stage.UnstageBranchMultiple_rest(target)

	case *Name_display:
		stage.UnstageBranchName_display(target)

	case *Non_arpeggiate:
		stage.UnstageBranchNon_arpeggiate(target)

	case *Notations:
		stage.UnstageBranchNotations(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *Note_size:
		stage.UnstageBranchNote_size(target)

	case *Note_type:
		stage.UnstageBranchNote_type(target)

	case *Notehead:
		stage.UnstageBranchNotehead(target)

	case *Notehead_text:
		stage.UnstageBranchNotehead_text(target)

	case *Numeral:
		stage.UnstageBranchNumeral(target)

	case *Numeral_key:
		stage.UnstageBranchNumeral_key(target)

	case *Numeral_root:
		stage.UnstageBranchNumeral_root(target)

	case *Octave_shift:
		stage.UnstageBranchOctave_shift(target)

	case *Offset:
		stage.UnstageBranchOffset(target)

	case *Opus:
		stage.UnstageBranchOpus(target)

	case *Ornaments:
		stage.UnstageBranchOrnaments(target)

	case *Other_appearance:
		stage.UnstageBranchOther_appearance(target)

	case *Other_listening:
		stage.UnstageBranchOther_listening(target)

	case *Other_notation:
		stage.UnstageBranchOther_notation(target)

	case *Other_play:
		stage.UnstageBranchOther_play(target)

	case *Page_layout:
		stage.UnstageBranchPage_layout(target)

	case *Page_margins:
		stage.UnstageBranchPage_margins(target)

	case *Part_clef:
		stage.UnstageBranchPart_clef(target)

	case *Part_group:
		stage.UnstageBranchPart_group(target)

	case *Part_link:
		stage.UnstageBranchPart_link(target)

	case *Part_list:
		stage.UnstageBranchPart_list(target)

	case *Part_symbol:
		stage.UnstageBranchPart_symbol(target)

	case *Part_transpose:
		stage.UnstageBranchPart_transpose(target)

	case *Pedal:
		stage.UnstageBranchPedal(target)

	case *Pedal_tuning:
		stage.UnstageBranchPedal_tuning(target)

	case *Percussion:
		stage.UnstageBranchPercussion(target)

	case *Pitch:
		stage.UnstageBranchPitch(target)

	case *Pitched:
		stage.UnstageBranchPitched(target)

	case *Play:
		stage.UnstageBranchPlay(target)

	case *Player:
		stage.UnstageBranchPlayer(target)

	case *Principal_voice:
		stage.UnstageBranchPrincipal_voice(target)

	case *Print:
		stage.UnstageBranchPrint(target)

	case *Release:
		stage.UnstageBranchRelease(target)

	case *Repeat:
		stage.UnstageBranchRepeat(target)

	case *Rest:
		stage.UnstageBranchRest(target)

	case *Root:
		stage.UnstageBranchRoot(target)

	case *Root_step:
		stage.UnstageBranchRoot_step(target)

	case *Scaling:
		stage.UnstageBranchScaling(target)

	case *Scordatura:
		stage.UnstageBranchScordatura(target)

	case *Score_instrument:
		stage.UnstageBranchScore_instrument(target)

	case *Score_part:
		stage.UnstageBranchScore_part(target)

	case *Score_partwise:
		stage.UnstageBranchScore_partwise(target)

	case *Score_timewise:
		stage.UnstageBranchScore_timewise(target)

	case *Segno:
		stage.UnstageBranchSegno(target)

	case *Slash:
		stage.UnstageBranchSlash(target)

	case *Slide:
		stage.UnstageBranchSlide(target)

	case *Slur:
		stage.UnstageBranchSlur(target)

	case *Sound:
		stage.UnstageBranchSound(target)

	case *Staff_details:
		stage.UnstageBranchStaff_details(target)

	case *Staff_divide:
		stage.UnstageBranchStaff_divide(target)

	case *Staff_layout:
		stage.UnstageBranchStaff_layout(target)

	case *Staff_size:
		stage.UnstageBranchStaff_size(target)

	case *Staff_tuning:
		stage.UnstageBranchStaff_tuning(target)

	case *Stem:
		stage.UnstageBranchStem(target)

	case *Stick:
		stage.UnstageBranchStick(target)

	case *String_mute:
		stage.UnstageBranchString_mute(target)

	case *Strong_accent:
		stage.UnstageBranchStrong_accent(target)

	case *Supports:
		stage.UnstageBranchSupports(target)

	case *Swing:
		stage.UnstageBranchSwing(target)

	case *Sync:
		stage.UnstageBranchSync(target)

	case *System_dividers:
		stage.UnstageBranchSystem_dividers(target)

	case *System_layout:
		stage.UnstageBranchSystem_layout(target)

	case *System_margins:
		stage.UnstageBranchSystem_margins(target)

	case *Tap:
		stage.UnstageBranchTap(target)

	case *Technical:
		stage.UnstageBranchTechnical(target)

	case *Text_element_data:
		stage.UnstageBranchText_element_data(target)

	case *Tie:
		stage.UnstageBranchTie(target)

	case *Tied:
		stage.UnstageBranchTied(target)

	case *Time:
		stage.UnstageBranchTime(target)

	case *Time_modification:
		stage.UnstageBranchTime_modification(target)

	case *Timpani:
		stage.UnstageBranchTimpani(target)

	case *Transpose:
		stage.UnstageBranchTranspose(target)

	case *Tremolo:
		stage.UnstageBranchTremolo(target)

	case *Tuplet:
		stage.UnstageBranchTuplet(target)

	case *Tuplet_dot:
		stage.UnstageBranchTuplet_dot(target)

	case *Tuplet_number:
		stage.UnstageBranchTuplet_number(target)

	case *Tuplet_portion:
		stage.UnstageBranchTuplet_portion(target)

	case *Tuplet_type:
		stage.UnstageBranchTuplet_type(target)

	case *Typed_text:
		stage.UnstageBranchTyped_text(target)

	case *Unpitched:
		stage.UnstageBranchUnpitched(target)

	case *Virtual_instrument:
		stage.UnstageBranchVirtual_instrument(target)

	case *Wait:
		stage.UnstageBranchWait(target)

	case *Wavy_line:
		stage.UnstageBranchWavy_line(target)

	case *Wedge:
		stage.UnstageBranchWedge(target)

	case *Wood:
		stage.UnstageBranchWood(target)

	case *Work:
		stage.UnstageBranchWork(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchAccidental(accidental *Accidental) {

	// check if instance is already staged
	if !IsStaged(stage, accidental) {
		return
	}

	accidental.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAccidental_mark(accidental_mark *Accidental_mark) {

	// check if instance is already staged
	if !IsStaged(stage, accidental_mark) {
		return
	}

	accidental_mark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAccidental_text(accidental_text *Accidental_text) {

	// check if instance is already staged
	if !IsStaged(stage, accidental_text) {
		return
	}

	accidental_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAccord(accord *Accord) {

	// check if instance is already staged
	if !IsStaged(stage, accord) {
		return
	}

	accord.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAccordion_registration(accordion_registration *Accordion_registration) {

	// check if instance is already staged
	if !IsStaged(stage, accordion_registration) {
		return
	}

	accordion_registration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if accordion_registration.Accordion_high != nil {
		UnstageBranch(stage, accordion_registration.Accordion_high)
	}
	if accordion_registration.Accordion_low != nil {
		UnstageBranch(stage, accordion_registration.Accordion_low)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAnyType(anytype *AnyType) {

	// check if instance is already staged
	if !IsStaged(stage, anytype) {
		return
	}

	anytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAppearance(appearance *Appearance) {

	// check if instance is already staged
	if !IsStaged(stage, appearance) {
		return
	}

	appearance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _line_width := range appearance.Line_width {
		UnstageBranch(stage, _line_width)
	}
	for _, _note_size := range appearance.Note_size {
		UnstageBranch(stage, _note_size)
	}
	for _, _distance := range appearance.Distance {
		UnstageBranch(stage, _distance)
	}
	for _, _glyph := range appearance.Glyph {
		UnstageBranch(stage, _glyph)
	}
	for _, _other_appearance := range appearance.Other_appearance {
		UnstageBranch(stage, _other_appearance)
	}

}

func (stage *StageStruct) UnstageBranchArpeggiate(arpeggiate *Arpeggiate) {

	// check if instance is already staged
	if !IsStaged(stage, arpeggiate) {
		return
	}

	arpeggiate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchArrow(arrow *Arrow) {

	// check if instance is already staged
	if !IsStaged(stage, arrow) {
		return
	}

	arrow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchArticulations(articulations *Articulations) {

	// check if instance is already staged
	if !IsStaged(stage, articulations) {
		return
	}

	articulations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if articulations.Accent != nil {
		UnstageBranch(stage, articulations.Accent)
	}
	if articulations.Strong_accent != nil {
		UnstageBranch(stage, articulations.Strong_accent)
	}
	if articulations.Staccato != nil {
		UnstageBranch(stage, articulations.Staccato)
	}
	if articulations.Tenuto != nil {
		UnstageBranch(stage, articulations.Tenuto)
	}
	if articulations.Detached_legato != nil {
		UnstageBranch(stage, articulations.Detached_legato)
	}
	if articulations.Staccatissimo != nil {
		UnstageBranch(stage, articulations.Staccatissimo)
	}
	if articulations.Spiccato != nil {
		UnstageBranch(stage, articulations.Spiccato)
	}
	if articulations.Scoop != nil {
		UnstageBranch(stage, articulations.Scoop)
	}
	if articulations.Plop != nil {
		UnstageBranch(stage, articulations.Plop)
	}
	if articulations.Doit != nil {
		UnstageBranch(stage, articulations.Doit)
	}
	if articulations.Falloff != nil {
		UnstageBranch(stage, articulations.Falloff)
	}
	if articulations.Breath_mark != nil {
		UnstageBranch(stage, articulations.Breath_mark)
	}
	if articulations.Caesura != nil {
		UnstageBranch(stage, articulations.Caesura)
	}
	if articulations.Stress != nil {
		UnstageBranch(stage, articulations.Stress)
	}
	if articulations.Unstress != nil {
		UnstageBranch(stage, articulations.Unstress)
	}
	if articulations.Soft_accent != nil {
		UnstageBranch(stage, articulations.Soft_accent)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAssess(assess *Assess) {

	// check if instance is already staged
	if !IsStaged(stage, assess) {
		return
	}

	assess.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchAttributes(attributes *Attributes) {

	// check if instance is already staged
	if !IsStaged(stage, attributes) {
		return
	}

	attributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attributes.Part_symbol != nil {
		UnstageBranch(stage, attributes.Part_symbol)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key := range attributes.Key {
		UnstageBranch(stage, _key)
	}
	for _, _clef := range attributes.Clef {
		UnstageBranch(stage, _clef)
	}
	for _, _staff_details := range attributes.Staff_details {
		UnstageBranch(stage, _staff_details)
	}
	for _, _measure_style := range attributes.Measure_style {
		UnstageBranch(stage, _measure_style)
	}
	for _, _transpose := range attributes.Transpose {
		UnstageBranch(stage, _transpose)
	}
	for _, _for_part := range attributes.For_part {
		UnstageBranch(stage, _for_part)
	}

}

func (stage *StageStruct) UnstageBranchBackup(backup *Backup) {

	// check if instance is already staged
	if !IsStaged(stage, backup) {
		return
	}

	backup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBar_style_color(bar_style_color *Bar_style_color) {

	// check if instance is already staged
	if !IsStaged(stage, bar_style_color) {
		return
	}

	bar_style_color.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBarline(barline *Barline) {

	// check if instance is already staged
	if !IsStaged(stage, barline) {
		return
	}

	barline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if barline.Bar_style != nil {
		UnstageBranch(stage, barline.Bar_style)
	}
	if barline.Wavy_line != nil {
		UnstageBranch(stage, barline.Wavy_line)
	}
	if barline.Fermata != nil {
		UnstageBranch(stage, barline.Fermata)
	}
	if barline.Ending != nil {
		UnstageBranch(stage, barline.Ending)
	}
	if barline.Repeat != nil {
		UnstageBranch(stage, barline.Repeat)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBarre(barre *Barre) {

	// check if instance is already staged
	if !IsStaged(stage, barre) {
		return
	}

	barre.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBass(bass *Bass) {

	// check if instance is already staged
	if !IsStaged(stage, bass) {
		return
	}

	bass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bass.Bass_step != nil {
		UnstageBranch(stage, bass.Bass_step)
	}
	if bass.Bass_alter != nil {
		UnstageBranch(stage, bass.Bass_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBass_step(bass_step *Bass_step) {

	// check if instance is already staged
	if !IsStaged(stage, bass_step) {
		return
	}

	bass_step.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBeam(beam *Beam) {

	// check if instance is already staged
	if !IsStaged(stage, beam) {
		return
	}

	beam.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBeat_repeat(beat_repeat *Beat_repeat) {

	// check if instance is already staged
	if !IsStaged(stage, beat_repeat) {
		return
	}

	beat_repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBeat_unit_tied(beat_unit_tied *Beat_unit_tied) {

	// check if instance is already staged
	if !IsStaged(stage, beat_unit_tied) {
		return
	}

	beat_unit_tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBeater(beater *Beater) {

	// check if instance is already staged
	if !IsStaged(stage, beater) {
		return
	}

	beater.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBend(bend *Bend) {

	// check if instance is already staged
	if !IsStaged(stage, bend) {
		return
	}

	bend.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if bend.Pre_bend != nil {
		UnstageBranch(stage, bend.Pre_bend)
	}
	if bend.Release != nil {
		UnstageBranch(stage, bend.Release)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBookmark(bookmark *Bookmark) {

	// check if instance is already staged
	if !IsStaged(stage, bookmark) {
		return
	}

	bookmark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBracket(bracket *Bracket) {

	// check if instance is already staged
	if !IsStaged(stage, bracket) {
		return
	}

	bracket.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchBreath_mark(breath_mark *Breath_mark) {

	// check if instance is already staged
	if !IsStaged(stage, breath_mark) {
		return
	}

	breath_mark.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCaesura(caesura *Caesura) {

	// check if instance is already staged
	if !IsStaged(stage, caesura) {
		return
	}

	caesura.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCancel(cancel *Cancel) {

	// check if instance is already staged
	if !IsStaged(stage, cancel) {
		return
	}

	cancel.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchClef(clef *Clef) {

	// check if instance is already staged
	if !IsStaged(stage, clef) {
		return
	}

	clef.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCoda(coda *Coda) {

	// check if instance is already staged
	if !IsStaged(stage, coda) {
		return
	}

	coda.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if !IsStaged(stage, credit) {
		return
	}

	credit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if credit.Credit_image != nil {
		UnstageBranch(stage, credit.Credit_image)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		UnstageBranch(stage, _link)
	}
	for _, _bookmark := range credit.Bookmark {
		UnstageBranch(stage, _bookmark)
	}

}

func (stage *StageStruct) UnstageBranchDashes(dashes *Dashes) {

	// check if instance is already staged
	if !IsStaged(stage, dashes) {
		return
	}

	dashes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDefaults(defaults *Defaults) {

	// check if instance is already staged
	if !IsStaged(stage, defaults) {
		return
	}

	defaults.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if defaults.Scaling != nil {
		UnstageBranch(stage, defaults.Scaling)
	}
	if defaults.Concert_score != nil {
		UnstageBranch(stage, defaults.Concert_score)
	}
	if defaults.Appearance != nil {
		UnstageBranch(stage, defaults.Appearance)
	}
	if defaults.Music_font != nil {
		UnstageBranch(stage, defaults.Music_font)
	}
	if defaults.Word_font != nil {
		UnstageBranch(stage, defaults.Word_font)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _lyric_font := range defaults.Lyric_font {
		UnstageBranch(stage, _lyric_font)
	}
	for _, _lyric_language := range defaults.Lyric_language {
		UnstageBranch(stage, _lyric_language)
	}

}

func (stage *StageStruct) UnstageBranchDegree(degree *Degree) {

	// check if instance is already staged
	if !IsStaged(stage, degree) {
		return
	}

	degree.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if degree.Degree_value != nil {
		UnstageBranch(stage, degree.Degree_value)
	}
	if degree.Degree_alter != nil {
		UnstageBranch(stage, degree.Degree_alter)
	}
	if degree.Degree_type != nil {
		UnstageBranch(stage, degree.Degree_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDegree_alter(degree_alter *Degree_alter) {

	// check if instance is already staged
	if !IsStaged(stage, degree_alter) {
		return
	}

	degree_alter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDegree_type(degree_type *Degree_type) {

	// check if instance is already staged
	if !IsStaged(stage, degree_type) {
		return
	}

	degree_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDegree_value(degree_value *Degree_value) {

	// check if instance is already staged
	if !IsStaged(stage, degree_value) {
		return
	}

	degree_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDirection(direction *Direction) {

	// check if instance is already staged
	if !IsStaged(stage, direction) {
		return
	}

	direction.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction.Offset != nil {
		UnstageBranch(stage, direction.Offset)
	}
	if direction.Sound != nil {
		UnstageBranch(stage, direction.Sound)
	}
	if direction.Listening != nil {
		UnstageBranch(stage, direction.Listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _direction_type := range direction.Direction_type {
		UnstageBranch(stage, _direction_type)
	}

}

func (stage *StageStruct) UnstageBranchDirection_type(direction_type *Direction_type) {

	// check if instance is already staged
	if !IsStaged(stage, direction_type) {
		return
	}

	direction_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if direction_type.Wedge != nil {
		UnstageBranch(stage, direction_type.Wedge)
	}
	if direction_type.Dashes != nil {
		UnstageBranch(stage, direction_type.Dashes)
	}
	if direction_type.Bracket != nil {
		UnstageBranch(stage, direction_type.Bracket)
	}
	if direction_type.Pedal != nil {
		UnstageBranch(stage, direction_type.Pedal)
	}
	if direction_type.Metronome != nil {
		UnstageBranch(stage, direction_type.Metronome)
	}
	if direction_type.Octave_shift != nil {
		UnstageBranch(stage, direction_type.Octave_shift)
	}
	if direction_type.Harp_pedals != nil {
		UnstageBranch(stage, direction_type.Harp_pedals)
	}
	if direction_type.Damp != nil {
		UnstageBranch(stage, direction_type.Damp)
	}
	if direction_type.Damp_all != nil {
		UnstageBranch(stage, direction_type.Damp_all)
	}
	if direction_type.Eyeglasses != nil {
		UnstageBranch(stage, direction_type.Eyeglasses)
	}
	if direction_type.String_mute != nil {
		UnstageBranch(stage, direction_type.String_mute)
	}
	if direction_type.Scordatura != nil {
		UnstageBranch(stage, direction_type.Scordatura)
	}
	if direction_type.Image != nil {
		UnstageBranch(stage, direction_type.Image)
	}
	if direction_type.Principal_voice != nil {
		UnstageBranch(stage, direction_type.Principal_voice)
	}
	if direction_type.Accordion_registration != nil {
		UnstageBranch(stage, direction_type.Accordion_registration)
	}
	if direction_type.Staff_divide != nil {
		UnstageBranch(stage, direction_type.Staff_divide)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _segno := range direction_type.Segno {
		UnstageBranch(stage, _segno)
	}
	for _, _coda := range direction_type.Coda {
		UnstageBranch(stage, _coda)
	}
	for _, _dynamics := range direction_type.Dynamics {
		UnstageBranch(stage, _dynamics)
	}
	for _, _percussion := range direction_type.Percussion {
		UnstageBranch(stage, _percussion)
	}

}

func (stage *StageStruct) UnstageBranchDistance(distance *Distance) {

	// check if instance is already staged
	if !IsStaged(stage, distance) {
		return
	}

	distance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDouble(double *Double) {

	// check if instance is already staged
	if !IsStaged(stage, double) {
		return
	}

	double.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDynamics(dynamics *Dynamics) {

	// check if instance is already staged
	if !IsStaged(stage, dynamics) {
		return
	}

	dynamics.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dynamics.P != nil {
		UnstageBranch(stage, dynamics.P)
	}
	if dynamics.Pp != nil {
		UnstageBranch(stage, dynamics.Pp)
	}
	if dynamics.Ppp != nil {
		UnstageBranch(stage, dynamics.Ppp)
	}
	if dynamics.Pppp != nil {
		UnstageBranch(stage, dynamics.Pppp)
	}
	if dynamics.Ppppp != nil {
		UnstageBranch(stage, dynamics.Ppppp)
	}
	if dynamics.Pppppp != nil {
		UnstageBranch(stage, dynamics.Pppppp)
	}
	if dynamics.F != nil {
		UnstageBranch(stage, dynamics.F)
	}
	if dynamics.Ff != nil {
		UnstageBranch(stage, dynamics.Ff)
	}
	if dynamics.Fff != nil {
		UnstageBranch(stage, dynamics.Fff)
	}
	if dynamics.Ffff != nil {
		UnstageBranch(stage, dynamics.Ffff)
	}
	if dynamics.Fffff != nil {
		UnstageBranch(stage, dynamics.Fffff)
	}
	if dynamics.Ffffff != nil {
		UnstageBranch(stage, dynamics.Ffffff)
	}
	if dynamics.Mp != nil {
		UnstageBranch(stage, dynamics.Mp)
	}
	if dynamics.Mf != nil {
		UnstageBranch(stage, dynamics.Mf)
	}
	if dynamics.Sf != nil {
		UnstageBranch(stage, dynamics.Sf)
	}
	if dynamics.Sfp != nil {
		UnstageBranch(stage, dynamics.Sfp)
	}
	if dynamics.Sfpp != nil {
		UnstageBranch(stage, dynamics.Sfpp)
	}
	if dynamics.Fp != nil {
		UnstageBranch(stage, dynamics.Fp)
	}
	if dynamics.Rf != nil {
		UnstageBranch(stage, dynamics.Rf)
	}
	if dynamics.Rfz != nil {
		UnstageBranch(stage, dynamics.Rfz)
	}
	if dynamics.Sfz != nil {
		UnstageBranch(stage, dynamics.Sfz)
	}
	if dynamics.Sffz != nil {
		UnstageBranch(stage, dynamics.Sffz)
	}
	if dynamics.Fz != nil {
		UnstageBranch(stage, dynamics.Fz)
	}
	if dynamics.N != nil {
		UnstageBranch(stage, dynamics.N)
	}
	if dynamics.Pf != nil {
		UnstageBranch(stage, dynamics.Pf)
	}
	if dynamics.Sfzp != nil {
		UnstageBranch(stage, dynamics.Sfzp)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEffect(effect *Effect) {

	// check if instance is already staged
	if !IsStaged(stage, effect) {
		return
	}

	effect.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchElision(elision *Elision) {

	// check if instance is already staged
	if !IsStaged(stage, elision) {
		return
	}

	elision.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty(empty *Empty) {

	// check if instance is already staged
	if !IsStaged(stage, empty) {
		return
	}

	empty.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_font(empty_font *Empty_font) {

	// check if instance is already staged
	if !IsStaged(stage, empty_font) {
		return
	}

	empty_font.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_line(empty_line *Empty_line) {

	// check if instance is already staged
	if !IsStaged(stage, empty_line) {
		return
	}

	empty_line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_placement(empty_placement *Empty_placement) {

	// check if instance is already staged
	if !IsStaged(stage, empty_placement) {
		return
	}

	empty_placement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_placement_smufl(empty_placement_smufl *Empty_placement_smufl) {

	// check if instance is already staged
	if !IsStaged(stage, empty_placement_smufl) {
		return
	}

	empty_placement_smufl.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_print_object_style_align(empty_print_object_style_align *Empty_print_object_style_align) {

	// check if instance is already staged
	if !IsStaged(stage, empty_print_object_style_align) {
		return
	}

	empty_print_object_style_align.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_print_style(empty_print_style *Empty_print_style) {

	// check if instance is already staged
	if !IsStaged(stage, empty_print_style) {
		return
	}

	empty_print_style.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_print_style_align(empty_print_style_align *Empty_print_style_align) {

	// check if instance is already staged
	if !IsStaged(stage, empty_print_style_align) {
		return
	}

	empty_print_style_align.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_print_style_align_id(empty_print_style_align_id *Empty_print_style_align_id) {

	// check if instance is already staged
	if !IsStaged(stage, empty_print_style_align_id) {
		return
	}

	empty_print_style_align_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEmpty_trill_sound(empty_trill_sound *Empty_trill_sound) {

	// check if instance is already staged
	if !IsStaged(stage, empty_trill_sound) {
		return
	}

	empty_trill_sound.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEncoding(encoding *Encoding) {

	// check if instance is already staged
	if !IsStaged(stage, encoding) {
		return
	}

	encoding.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if encoding.Encoder != nil {
		UnstageBranch(stage, encoding.Encoder)
	}
	if encoding.Supports != nil {
		UnstageBranch(stage, encoding.Supports)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchEnding(ending *Ending) {

	// check if instance is already staged
	if !IsStaged(stage, ending) {
		return
	}

	ending.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchExtend(extend *Extend) {

	// check if instance is already staged
	if !IsStaged(stage, extend) {
		return
	}

	extend.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFeature(feature *Feature) {

	// check if instance is already staged
	if !IsStaged(stage, feature) {
		return
	}

	feature.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFermata(fermata *Fermata) {

	// check if instance is already staged
	if !IsStaged(stage, fermata) {
		return
	}

	fermata.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFigure(figure *Figure) {

	// check if instance is already staged
	if !IsStaged(stage, figure) {
		return
	}

	figure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if figure.Extend != nil {
		UnstageBranch(stage, figure.Extend)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFigured_bass(figured_bass *Figured_bass) {

	// check if instance is already staged
	if !IsStaged(stage, figured_bass) {
		return
	}

	figured_bass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _figure := range figured_bass.Figure {
		UnstageBranch(stage, _figure)
	}

}

func (stage *StageStruct) UnstageBranchFingering(fingering *Fingering) {

	// check if instance is already staged
	if !IsStaged(stage, fingering) {
		return
	}

	fingering.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFirst_fret(first_fret *First_fret) {

	// check if instance is already staged
	if !IsStaged(stage, first_fret) {
		return
	}

	first_fret.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFoo(foo *Foo) {

	// check if instance is already staged
	if !IsStaged(stage, foo) {
		return
	}

	foo.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFor_part(for_part *For_part) {

	// check if instance is already staged
	if !IsStaged(stage, for_part) {
		return
	}

	for_part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if for_part.Part_clef != nil {
		UnstageBranch(stage, for_part.Part_clef)
	}
	if for_part.Part_transpose != nil {
		UnstageBranch(stage, for_part.Part_transpose)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormatted_symbol(formatted_symbol *Formatted_symbol) {

	// check if instance is already staged
	if !IsStaged(stage, formatted_symbol) {
		return
	}

	formatted_symbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFormatted_symbol_id(formatted_symbol_id *Formatted_symbol_id) {

	// check if instance is already staged
	if !IsStaged(stage, formatted_symbol_id) {
		return
	}

	formatted_symbol_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchForward(forward *Forward) {

	// check if instance is already staged
	if !IsStaged(stage, forward) {
		return
	}

	forward.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFrame(frame *Frame) {

	// check if instance is already staged
	if !IsStaged(stage, frame) {
		return
	}

	frame.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame.First_fret != nil {
		UnstageBranch(stage, frame.First_fret)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _frame_note := range frame.Frame_note {
		UnstageBranch(stage, _frame_note)
	}

}

func (stage *StageStruct) UnstageBranchFrame_note(frame_note *Frame_note) {

	// check if instance is already staged
	if !IsStaged(stage, frame_note) {
		return
	}

	frame_note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if frame_note.Fret != nil {
		UnstageBranch(stage, frame_note.Fret)
	}
	if frame_note.Fingering != nil {
		UnstageBranch(stage, frame_note.Fingering)
	}
	if frame_note.Barre != nil {
		UnstageBranch(stage, frame_note.Barre)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchFret(fret *Fret) {

	// check if instance is already staged
	if !IsStaged(stage, fret) {
		return
	}

	fret.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGlass(glass *Glass) {

	// check if instance is already staged
	if !IsStaged(stage, glass) {
		return
	}

	glass.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGlissando(glissando *Glissando) {

	// check if instance is already staged
	if !IsStaged(stage, glissando) {
		return
	}

	glissando.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGlyph(glyph *Glyph) {

	// check if instance is already staged
	if !IsStaged(stage, glyph) {
		return
	}

	glyph.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGrace(grace *Grace) {

	// check if instance is already staged
	if !IsStaged(stage, grace) {
		return
	}

	grace.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGroup_barline(group_barline *Group_barline) {

	// check if instance is already staged
	if !IsStaged(stage, group_barline) {
		return
	}

	group_barline.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGroup_symbol(group_symbol *Group_symbol) {

	// check if instance is already staged
	if !IsStaged(stage, group_symbol) {
		return
	}

	group_symbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGrouping(grouping *Grouping) {

	// check if instance is already staged
	if !IsStaged(stage, grouping) {
		return
	}

	grouping.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _feature := range grouping.Feature {
		UnstageBranch(stage, _feature)
	}

}

func (stage *StageStruct) UnstageBranchHammer_on_pull_off(hammer_on_pull_off *Hammer_on_pull_off) {

	// check if instance is already staged
	if !IsStaged(stage, hammer_on_pull_off) {
		return
	}

	hammer_on_pull_off.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHandbell(handbell *Handbell) {

	// check if instance is already staged
	if !IsStaged(stage, handbell) {
		return
	}

	handbell.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHarmon_closed(harmon_closed *Harmon_closed) {

	// check if instance is already staged
	if !IsStaged(stage, harmon_closed) {
		return
	}

	harmon_closed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHarmon_mute(harmon_mute *Harmon_mute) {

	// check if instance is already staged
	if !IsStaged(stage, harmon_mute) {
		return
	}

	harmon_mute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmon_mute.Harmon_closed != nil {
		UnstageBranch(stage, harmon_mute.Harmon_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHarmonic(harmonic *Harmonic) {

	// check if instance is already staged
	if !IsStaged(stage, harmonic) {
		return
	}

	harmonic.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmonic.Natural != nil {
		UnstageBranch(stage, harmonic.Natural)
	}
	if harmonic.Artificial != nil {
		UnstageBranch(stage, harmonic.Artificial)
	}
	if harmonic.Base_pitch != nil {
		UnstageBranch(stage, harmonic.Base_pitch)
	}
	if harmonic.Touching_pitch != nil {
		UnstageBranch(stage, harmonic.Touching_pitch)
	}
	if harmonic.Sounding_pitch != nil {
		UnstageBranch(stage, harmonic.Sounding_pitch)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHarmony(harmony *Harmony) {

	// check if instance is already staged
	if !IsStaged(stage, harmony) {
		return
	}

	harmony.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if harmony.Frame != nil {
		UnstageBranch(stage, harmony.Frame)
	}
	if harmony.Offset != nil {
		UnstageBranch(stage, harmony.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHarmony_alter(harmony_alter *Harmony_alter) {

	// check if instance is already staged
	if !IsStaged(stage, harmony_alter) {
		return
	}

	harmony_alter.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHarp_pedals(harp_pedals *Harp_pedals) {

	// check if instance is already staged
	if !IsStaged(stage, harp_pedals) {
		return
	}

	harp_pedals.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
		UnstageBranch(stage, _pedal_tuning)
	}

}

func (stage *StageStruct) UnstageBranchHeel_toe(heel_toe *Heel_toe) {

	// check if instance is already staged
	if !IsStaged(stage, heel_toe) {
		return
	}

	heel_toe.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHole(hole *Hole) {

	// check if instance is already staged
	if !IsStaged(stage, hole) {
		return
	}

	hole.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if hole.Hole_closed != nil {
		UnstageBranch(stage, hole.Hole_closed)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHole_closed(hole_closed *Hole_closed) {

	// check if instance is already staged
	if !IsStaged(stage, hole_closed) {
		return
	}

	hole_closed.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchHorizontal_turn(horizontal_turn *Horizontal_turn) {

	// check if instance is already staged
	if !IsStaged(stage, horizontal_turn) {
		return
	}

	horizontal_turn.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchIdentification(identification *Identification) {

	// check if instance is already staged
	if !IsStaged(stage, identification) {
		return
	}

	identification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if identification.Encoding != nil {
		UnstageBranch(stage, identification.Encoding)
	}
	if identification.Miscellaneous != nil {
		UnstageBranch(stage, identification.Miscellaneous)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _typed_text := range identification.Creator {
		UnstageBranch(stage, _typed_text)
	}
	for _, _typed_text := range identification.Rights {
		UnstageBranch(stage, _typed_text)
	}
	for _, _typed_text := range identification.Relation {
		UnstageBranch(stage, _typed_text)
	}

}

func (stage *StageStruct) UnstageBranchImage(image *Image) {

	// check if instance is already staged
	if !IsStaged(stage, image) {
		return
	}

	image.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchInstrument(instrument *Instrument) {

	// check if instance is already staged
	if !IsStaged(stage, instrument) {
		return
	}

	instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchInstrument_change(instrument_change *Instrument_change) {

	// check if instance is already staged
	if !IsStaged(stage, instrument_change) {
		return
	}

	instrument_change.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchInstrument_link(instrument_link *Instrument_link) {

	// check if instance is already staged
	if !IsStaged(stage, instrument_link) {
		return
	}

	instrument_link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchInterchangeable(interchangeable *Interchangeable) {

	// check if instance is already staged
	if !IsStaged(stage, interchangeable) {
		return
	}

	interchangeable.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchInversion(inversion *Inversion) {

	// check if instance is already staged
	if !IsStaged(stage, inversion) {
		return
	}

	inversion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchKey(key *Key) {

	// check if instance is already staged
	if !IsStaged(stage, key) {
		return
	}

	key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _key_octave := range key.Key_octave {
		UnstageBranch(stage, _key_octave)
	}

}

func (stage *StageStruct) UnstageBranchKey_accidental(key_accidental *Key_accidental) {

	// check if instance is already staged
	if !IsStaged(stage, key_accidental) {
		return
	}

	key_accidental.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchKey_octave(key_octave *Key_octave) {

	// check if instance is already staged
	if !IsStaged(stage, key_octave) {
		return
	}

	key_octave.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchKind(kind *Kind) {

	// check if instance is already staged
	if !IsStaged(stage, kind) {
		return
	}

	kind.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLevel(level *Level) {

	// check if instance is already staged
	if !IsStaged(stage, level) {
		return
	}

	level.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLine_detail(line_detail *Line_detail) {

	// check if instance is already staged
	if !IsStaged(stage, line_detail) {
		return
	}

	line_detail.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLine_width(line_width *Line_width) {

	// check if instance is already staged
	if !IsStaged(stage, line_width) {
		return
	}

	line_width.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchListen(listen *Listen) {

	// check if instance is already staged
	if !IsStaged(stage, listen) {
		return
	}

	listen.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listen.Assess != nil {
		UnstageBranch(stage, listen.Assess)
	}
	if listen.Wait != nil {
		UnstageBranch(stage, listen.Wait)
	}
	if listen.Other_listen != nil {
		UnstageBranch(stage, listen.Other_listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchListening(listening *Listening) {

	// check if instance is already staged
	if !IsStaged(stage, listening) {
		return
	}

	listening.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if listening.Offset != nil {
		UnstageBranch(stage, listening.Offset)
	}
	if listening.Sync != nil {
		UnstageBranch(stage, listening.Sync)
	}
	if listening.Other_listening != nil {
		UnstageBranch(stage, listening.Other_listening)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLyric(lyric *Lyric) {

	// check if instance is already staged
	if !IsStaged(stage, lyric) {
		return
	}

	lyric.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if lyric.End_line != nil {
		UnstageBranch(stage, lyric.End_line)
	}
	if lyric.End_paragraph != nil {
		UnstageBranch(stage, lyric.End_paragraph)
	}
	if lyric.Extend != nil {
		UnstageBranch(stage, lyric.Extend)
	}
	if lyric.Laughing != nil {
		UnstageBranch(stage, lyric.Laughing)
	}
	if lyric.Humming != nil {
		UnstageBranch(stage, lyric.Humming)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLyric_font(lyric_font *Lyric_font) {

	// check if instance is already staged
	if !IsStaged(stage, lyric_font) {
		return
	}

	lyric_font.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLyric_language(lyric_language *Lyric_language) {

	// check if instance is already staged
	if !IsStaged(stage, lyric_language) {
		return
	}

	lyric_language.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMeasure_layout(measure_layout *Measure_layout) {

	// check if instance is already staged
	if !IsStaged(stage, measure_layout) {
		return
	}

	measure_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMeasure_numbering(measure_numbering *Measure_numbering) {

	// check if instance is already staged
	if !IsStaged(stage, measure_numbering) {
		return
	}

	measure_numbering.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMeasure_repeat(measure_repeat *Measure_repeat) {

	// check if instance is already staged
	if !IsStaged(stage, measure_repeat) {
		return
	}

	measure_repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMeasure_style(measure_style *Measure_style) {

	// check if instance is already staged
	if !IsStaged(stage, measure_style) {
		return
	}

	measure_style.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if measure_style.Multiple_rest != nil {
		UnstageBranch(stage, measure_style.Multiple_rest)
	}
	if measure_style.Measure_repeat != nil {
		UnstageBranch(stage, measure_style.Measure_repeat)
	}
	if measure_style.Beat_repeat != nil {
		UnstageBranch(stage, measure_style.Beat_repeat)
	}
	if measure_style.Slash != nil {
		UnstageBranch(stage, measure_style.Slash)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMembrane(membrane *Membrane) {

	// check if instance is already staged
	if !IsStaged(stage, membrane) {
		return
	}

	membrane.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMetal(metal *Metal) {

	// check if instance is already staged
	if !IsStaged(stage, metal) {
		return
	}

	metal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMetronome(metronome *Metronome) {

	// check if instance is already staged
	if !IsStaged(stage, metronome) {
		return
	}

	metronome.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMetronome_beam(metronome_beam *Metronome_beam) {

	// check if instance is already staged
	if !IsStaged(stage, metronome_beam) {
		return
	}

	metronome_beam.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMetronome_note(metronome_note *Metronome_note) {

	// check if instance is already staged
	if !IsStaged(stage, metronome_note) {
		return
	}

	metronome_note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if metronome_note.Metronome_tied != nil {
		UnstageBranch(stage, metronome_note.Metronome_tied)
	}
	if metronome_note.Metronome_tuplet != nil {
		UnstageBranch(stage, metronome_note.Metronome_tuplet)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _empty := range metronome_note.Metronome_dot {
		UnstageBranch(stage, _empty)
	}
	for _, _metronome_beam := range metronome_note.Metronome_beam {
		UnstageBranch(stage, _metronome_beam)
	}

}

func (stage *StageStruct) UnstageBranchMetronome_tied(metronome_tied *Metronome_tied) {

	// check if instance is already staged
	if !IsStaged(stage, metronome_tied) {
		return
	}

	metronome_tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMetronome_tuplet(metronome_tuplet *Metronome_tuplet) {

	// check if instance is already staged
	if !IsStaged(stage, metronome_tuplet) {
		return
	}

	metronome_tuplet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMidi_device(midi_device *Midi_device) {

	// check if instance is already staged
	if !IsStaged(stage, midi_device) {
		return
	}

	midi_device.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMidi_instrument(midi_instrument *Midi_instrument) {

	// check if instance is already staged
	if !IsStaged(stage, midi_instrument) {
		return
	}

	midi_instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMiscellaneous(miscellaneous *Miscellaneous) {

	// check if instance is already staged
	if !IsStaged(stage, miscellaneous) {
		return
	}

	miscellaneous.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
		UnstageBranch(stage, _miscellaneous_field)
	}

}

func (stage *StageStruct) UnstageBranchMiscellaneous_field(miscellaneous_field *Miscellaneous_field) {

	// check if instance is already staged
	if !IsStaged(stage, miscellaneous_field) {
		return
	}

	miscellaneous_field.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMordent(mordent *Mordent) {

	// check if instance is already staged
	if !IsStaged(stage, mordent) {
		return
	}

	mordent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMultiple_rest(multiple_rest *Multiple_rest) {

	// check if instance is already staged
	if !IsStaged(stage, multiple_rest) {
		return
	}

	multiple_rest.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchName_display(name_display *Name_display) {

	// check if instance is already staged
	if !IsStaged(stage, name_display) {
		return
	}

	name_display.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if name_display.Accidental_text != nil {
		UnstageBranch(stage, name_display.Accidental_text)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNon_arpeggiate(non_arpeggiate *Non_arpeggiate) {

	// check if instance is already staged
	if !IsStaged(stage, non_arpeggiate) {
		return
	}

	non_arpeggiate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNotations(notations *Notations) {

	// check if instance is already staged
	if !IsStaged(stage, notations) {
		return
	}

	notations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notations.Tied != nil {
		UnstageBranch(stage, notations.Tied)
	}
	if notations.Slur != nil {
		UnstageBranch(stage, notations.Slur)
	}
	if notations.Tuplet != nil {
		UnstageBranch(stage, notations.Tuplet)
	}
	if notations.Glissando != nil {
		UnstageBranch(stage, notations.Glissando)
	}
	if notations.Slide != nil {
		UnstageBranch(stage, notations.Slide)
	}
	if notations.Ornaments != nil {
		UnstageBranch(stage, notations.Ornaments)
	}
	if notations.Technical != nil {
		UnstageBranch(stage, notations.Technical)
	}
	if notations.Articulations != nil {
		UnstageBranch(stage, notations.Articulations)
	}
	if notations.Dynamics != nil {
		UnstageBranch(stage, notations.Dynamics)
	}
	if notations.Fermata != nil {
		UnstageBranch(stage, notations.Fermata)
	}
	if notations.Arpeggiate != nil {
		UnstageBranch(stage, notations.Arpeggiate)
	}
	if notations.Non_arpeggiate != nil {
		UnstageBranch(stage, notations.Non_arpeggiate)
	}
	if notations.Accidental_mark != nil {
		UnstageBranch(stage, notations.Accidental_mark)
	}
	if notations.Other_notation != nil {
		UnstageBranch(stage, notations.Other_notation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if note.Type_ != nil {
		UnstageBranch(stage, note.Type_)
	}
	if note.Accidental != nil {
		UnstageBranch(stage, note.Accidental)
	}
	if note.Time_modification != nil {
		UnstageBranch(stage, note.Time_modification)
	}
	if note.Stem != nil {
		UnstageBranch(stage, note.Stem)
	}
	if note.Notehead != nil {
		UnstageBranch(stage, note.Notehead)
	}
	if note.Notehead_text != nil {
		UnstageBranch(stage, note.Notehead_text)
	}
	if note.Beam != nil {
		UnstageBranch(stage, note.Beam)
	}
	if note.Play != nil {
		UnstageBranch(stage, note.Play)
	}
	if note.Listen != nil {
		UnstageBranch(stage, note.Listen)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument := range note.Instrument {
		UnstageBranch(stage, _instrument)
	}
	for _, _empty_placement := range note.Dot {
		UnstageBranch(stage, _empty_placement)
	}
	for _, _notations := range note.Notations {
		UnstageBranch(stage, _notations)
	}
	for _, _lyric := range note.Lyric {
		UnstageBranch(stage, _lyric)
	}

}

func (stage *StageStruct) UnstageBranchNote_size(note_size *Note_size) {

	// check if instance is already staged
	if !IsStaged(stage, note_size) {
		return
	}

	note_size.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNote_type(note_type *Note_type) {

	// check if instance is already staged
	if !IsStaged(stage, note_type) {
		return
	}

	note_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNotehead(notehead *Notehead) {

	// check if instance is already staged
	if !IsStaged(stage, notehead) {
		return
	}

	notehead.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNotehead_text(notehead_text *Notehead_text) {

	// check if instance is already staged
	if !IsStaged(stage, notehead_text) {
		return
	}

	notehead_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notehead_text.Accidental_text != nil {
		UnstageBranch(stage, notehead_text.Accidental_text)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNumeral(numeral *Numeral) {

	// check if instance is already staged
	if !IsStaged(stage, numeral) {
		return
	}

	numeral.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if numeral.Numeral_root != nil {
		UnstageBranch(stage, numeral.Numeral_root)
	}
	if numeral.Numeral_alter != nil {
		UnstageBranch(stage, numeral.Numeral_alter)
	}
	if numeral.Numeral_key != nil {
		UnstageBranch(stage, numeral.Numeral_key)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNumeral_key(numeral_key *Numeral_key) {

	// check if instance is already staged
	if !IsStaged(stage, numeral_key) {
		return
	}

	numeral_key.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNumeral_root(numeral_root *Numeral_root) {

	// check if instance is already staged
	if !IsStaged(stage, numeral_root) {
		return
	}

	numeral_root.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOctave_shift(octave_shift *Octave_shift) {

	// check if instance is already staged
	if !IsStaged(stage, octave_shift) {
		return
	}

	octave_shift.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOffset(offset *Offset) {

	// check if instance is already staged
	if !IsStaged(stage, offset) {
		return
	}

	offset.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOpus(opus *Opus) {

	// check if instance is already staged
	if !IsStaged(stage, opus) {
		return
	}

	opus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOrnaments(ornaments *Ornaments) {

	// check if instance is already staged
	if !IsStaged(stage, ornaments) {
		return
	}

	ornaments.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if ornaments.Trill_mark != nil {
		UnstageBranch(stage, ornaments.Trill_mark)
	}
	if ornaments.Turn != nil {
		UnstageBranch(stage, ornaments.Turn)
	}
	if ornaments.Delayed_turn != nil {
		UnstageBranch(stage, ornaments.Delayed_turn)
	}
	if ornaments.Inverted_turn != nil {
		UnstageBranch(stage, ornaments.Inverted_turn)
	}
	if ornaments.Delayed_inverted_turn != nil {
		UnstageBranch(stage, ornaments.Delayed_inverted_turn)
	}
	if ornaments.Vertical_turn != nil {
		UnstageBranch(stage, ornaments.Vertical_turn)
	}
	if ornaments.Inverted_vertical_turn != nil {
		UnstageBranch(stage, ornaments.Inverted_vertical_turn)
	}
	if ornaments.Shake != nil {
		UnstageBranch(stage, ornaments.Shake)
	}
	if ornaments.Wavy_line != nil {
		UnstageBranch(stage, ornaments.Wavy_line)
	}
	if ornaments.Mordent != nil {
		UnstageBranch(stage, ornaments.Mordent)
	}
	if ornaments.Inverted_mordent != nil {
		UnstageBranch(stage, ornaments.Inverted_mordent)
	}
	if ornaments.Schleifer != nil {
		UnstageBranch(stage, ornaments.Schleifer)
	}
	if ornaments.Tremolo != nil {
		UnstageBranch(stage, ornaments.Tremolo)
	}
	if ornaments.Haydn != nil {
		UnstageBranch(stage, ornaments.Haydn)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accidental_mark := range ornaments.Accidental_mark {
		UnstageBranch(stage, _accidental_mark)
	}

}

func (stage *StageStruct) UnstageBranchOther_appearance(other_appearance *Other_appearance) {

	// check if instance is already staged
	if !IsStaged(stage, other_appearance) {
		return
	}

	other_appearance.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOther_listening(other_listening *Other_listening) {

	// check if instance is already staged
	if !IsStaged(stage, other_listening) {
		return
	}

	other_listening.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOther_notation(other_notation *Other_notation) {

	// check if instance is already staged
	if !IsStaged(stage, other_notation) {
		return
	}

	other_notation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchOther_play(other_play *Other_play) {

	// check if instance is already staged
	if !IsStaged(stage, other_play) {
		return
	}

	other_play.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPage_layout(page_layout *Page_layout) {

	// check if instance is already staged
	if !IsStaged(stage, page_layout) {
		return
	}

	page_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if page_layout.Page_margins != nil {
		UnstageBranch(stage, page_layout.Page_margins)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPage_margins(page_margins *Page_margins) {

	// check if instance is already staged
	if !IsStaged(stage, page_margins) {
		return
	}

	page_margins.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPart_clef(part_clef *Part_clef) {

	// check if instance is already staged
	if !IsStaged(stage, part_clef) {
		return
	}

	part_clef.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPart_group(part_group *Part_group) {

	// check if instance is already staged
	if !IsStaged(stage, part_group) {
		return
	}

	part_group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part_group.Group_name_display != nil {
		UnstageBranch(stage, part_group.Group_name_display)
	}
	if part_group.Group_abbreviation_display != nil {
		UnstageBranch(stage, part_group.Group_abbreviation_display)
	}
	if part_group.Group_symbol != nil {
		UnstageBranch(stage, part_group.Group_symbol)
	}
	if part_group.Group_barline != nil {
		UnstageBranch(stage, part_group.Group_barline)
	}
	if part_group.Group_time != nil {
		UnstageBranch(stage, part_group.Group_time)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPart_link(part_link *Part_link) {

	// check if instance is already staged
	if !IsStaged(stage, part_link) {
		return
	}

	part_link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _instrument_link := range part_link.Instrument_link {
		UnstageBranch(stage, _instrument_link)
	}

}

func (stage *StageStruct) UnstageBranchPart_list(part_list *Part_list) {

	// check if instance is already staged
	if !IsStaged(stage, part_list) {
		return
	}

	part_list.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPart_symbol(part_symbol *Part_symbol) {

	// check if instance is already staged
	if !IsStaged(stage, part_symbol) {
		return
	}

	part_symbol.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPart_transpose(part_transpose *Part_transpose) {

	// check if instance is already staged
	if !IsStaged(stage, part_transpose) {
		return
	}

	part_transpose.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPedal(pedal *Pedal) {

	// check if instance is already staged
	if !IsStaged(stage, pedal) {
		return
	}

	pedal.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPedal_tuning(pedal_tuning *Pedal_tuning) {

	// check if instance is already staged
	if !IsStaged(stage, pedal_tuning) {
		return
	}

	pedal_tuning.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPercussion(percussion *Percussion) {

	// check if instance is already staged
	if !IsStaged(stage, percussion) {
		return
	}

	percussion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if percussion.Glass != nil {
		UnstageBranch(stage, percussion.Glass)
	}
	if percussion.Metal != nil {
		UnstageBranch(stage, percussion.Metal)
	}
	if percussion.Wood != nil {
		UnstageBranch(stage, percussion.Wood)
	}
	if percussion.Pitched != nil {
		UnstageBranch(stage, percussion.Pitched)
	}
	if percussion.Membrane != nil {
		UnstageBranch(stage, percussion.Membrane)
	}
	if percussion.Effect != nil {
		UnstageBranch(stage, percussion.Effect)
	}
	if percussion.Timpani != nil {
		UnstageBranch(stage, percussion.Timpani)
	}
	if percussion.Beater != nil {
		UnstageBranch(stage, percussion.Beater)
	}
	if percussion.Stick != nil {
		UnstageBranch(stage, percussion.Stick)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPitch(pitch *Pitch) {

	// check if instance is already staged
	if !IsStaged(stage, pitch) {
		return
	}

	pitch.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPitched(pitched *Pitched) {

	// check if instance is already staged
	if !IsStaged(stage, pitched) {
		return
	}

	pitched.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPlay(play *Play) {

	// check if instance is already staged
	if !IsStaged(stage, play) {
		return
	}

	play.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if play.Other_play != nil {
		UnstageBranch(stage, play.Other_play)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPlayer(player *Player) {

	// check if instance is already staged
	if !IsStaged(stage, player) {
		return
	}

	player.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPrincipal_voice(principal_voice *Principal_voice) {

	// check if instance is already staged
	if !IsStaged(stage, principal_voice) {
		return
	}

	principal_voice.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPrint(print *Print) {

	// check if instance is already staged
	if !IsStaged(stage, print) {
		return
	}

	print.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if print.Measure_layout != nil {
		UnstageBranch(stage, print.Measure_layout)
	}
	if print.Measure_numbering != nil {
		UnstageBranch(stage, print.Measure_numbering)
	}
	if print.Part_name_display != nil {
		UnstageBranch(stage, print.Part_name_display)
	}
	if print.Part_abbreviation_display != nil {
		UnstageBranch(stage, print.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRelease(release *Release) {

	// check if instance is already staged
	if !IsStaged(stage, release) {
		return
	}

	release.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRepeat(repeat *Repeat) {

	// check if instance is already staged
	if !IsStaged(stage, repeat) {
		return
	}

	repeat.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRest(rest *Rest) {

	// check if instance is already staged
	if !IsStaged(stage, rest) {
		return
	}

	rest.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRoot(root *Root) {

	// check if instance is already staged
	if !IsStaged(stage, root) {
		return
	}

	root.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if root.Root_step != nil {
		UnstageBranch(stage, root.Root_step)
	}
	if root.Root_alter != nil {
		UnstageBranch(stage, root.Root_alter)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchRoot_step(root_step *Root_step) {

	// check if instance is already staged
	if !IsStaged(stage, root_step) {
		return
	}

	root_step.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchScaling(scaling *Scaling) {

	// check if instance is already staged
	if !IsStaged(stage, scaling) {
		return
	}

	scaling.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchScordatura(scordatura *Scordatura) {

	// check if instance is already staged
	if !IsStaged(stage, scordatura) {
		return
	}

	scordatura.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _accord := range scordatura.Accord {
		UnstageBranch(stage, _accord)
	}

}

func (stage *StageStruct) UnstageBranchScore_instrument(score_instrument *Score_instrument) {

	// check if instance is already staged
	if !IsStaged(stage, score_instrument) {
		return
	}

	score_instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchScore_part(score_part *Score_part) {

	// check if instance is already staged
	if !IsStaged(stage, score_part) {
		return
	}

	score_part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if score_part.Identification != nil {
		UnstageBranch(stage, score_part.Identification)
	}
	if score_part.Part_name_display != nil {
		UnstageBranch(stage, score_part.Part_name_display)
	}
	if score_part.Part_abbreviation_display != nil {
		UnstageBranch(stage, score_part.Part_abbreviation_display)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part_link := range score_part.Part_link {
		UnstageBranch(stage, _part_link)
	}
	for _, _score_instrument := range score_part.Score_instrument {
		UnstageBranch(stage, _score_instrument)
	}
	for _, _player := range score_part.Player {
		UnstageBranch(stage, _player)
	}

}

func (stage *StageStruct) UnstageBranchScore_partwise(score_partwise *Score_partwise) {

	// check if instance is already staged
	if !IsStaged(stage, score_partwise) {
		return
	}

	score_partwise.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchScore_timewise(score_timewise *Score_timewise) {

	// check if instance is already staged
	if !IsStaged(stage, score_timewise) {
		return
	}

	score_timewise.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSegno(segno *Segno) {

	// check if instance is already staged
	if !IsStaged(stage, segno) {
		return
	}

	segno.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSlash(slash *Slash) {

	// check if instance is already staged
	if !IsStaged(stage, slash) {
		return
	}

	slash.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSlide(slide *Slide) {

	// check if instance is already staged
	if !IsStaged(stage, slide) {
		return
	}

	slide.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSlur(slur *Slur) {

	// check if instance is already staged
	if !IsStaged(stage, slur) {
		return
	}

	slur.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSound(sound *Sound) {

	// check if instance is already staged
	if !IsStaged(stage, sound) {
		return
	}

	sound.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sound.Swing != nil {
		UnstageBranch(stage, sound.Swing)
	}
	if sound.Offset != nil {
		UnstageBranch(stage, sound.Offset)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStaff_details(staff_details *Staff_details) {

	// check if instance is already staged
	if !IsStaged(stage, staff_details) {
		return
	}

	staff_details.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if staff_details.Staff_size != nil {
		UnstageBranch(stage, staff_details.Staff_size)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _staff_tuning := range staff_details.Staff_tuning {
		UnstageBranch(stage, _staff_tuning)
	}

}

func (stage *StageStruct) UnstageBranchStaff_divide(staff_divide *Staff_divide) {

	// check if instance is already staged
	if !IsStaged(stage, staff_divide) {
		return
	}

	staff_divide.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStaff_layout(staff_layout *Staff_layout) {

	// check if instance is already staged
	if !IsStaged(stage, staff_layout) {
		return
	}

	staff_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStaff_size(staff_size *Staff_size) {

	// check if instance is already staged
	if !IsStaged(stage, staff_size) {
		return
	}

	staff_size.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStaff_tuning(staff_tuning *Staff_tuning) {

	// check if instance is already staged
	if !IsStaged(stage, staff_tuning) {
		return
	}

	staff_tuning.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStem(stem *Stem) {

	// check if instance is already staged
	if !IsStaged(stage, stem) {
		return
	}

	stem.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStick(stick *Stick) {

	// check if instance is already staged
	if !IsStaged(stage, stick) {
		return
	}

	stick.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchString_mute(string_mute *String_mute) {

	// check if instance is already staged
	if !IsStaged(stage, string_mute) {
		return
	}

	string_mute.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchStrong_accent(strong_accent *Strong_accent) {

	// check if instance is already staged
	if !IsStaged(stage, strong_accent) {
		return
	}

	strong_accent.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSupports(supports *Supports) {

	// check if instance is already staged
	if !IsStaged(stage, supports) {
		return
	}

	supports.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSwing(swing *Swing) {

	// check if instance is already staged
	if !IsStaged(stage, swing) {
		return
	}

	swing.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if swing.Straight != nil {
		UnstageBranch(stage, swing.Straight)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSync(sync *Sync) {

	// check if instance is already staged
	if !IsStaged(stage, sync) {
		return
	}

	sync.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSystem_dividers(system_dividers *System_dividers) {

	// check if instance is already staged
	if !IsStaged(stage, system_dividers) {
		return
	}

	system_dividers.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_dividers.Left_divider != nil {
		UnstageBranch(stage, system_dividers.Left_divider)
	}
	if system_dividers.Right_divider != nil {
		UnstageBranch(stage, system_dividers.Right_divider)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSystem_layout(system_layout *System_layout) {

	// check if instance is already staged
	if !IsStaged(stage, system_layout) {
		return
	}

	system_layout.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if system_layout.System_margins != nil {
		UnstageBranch(stage, system_layout.System_margins)
	}
	if system_layout.System_dividers != nil {
		UnstageBranch(stage, system_layout.System_dividers)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSystem_margins(system_margins *System_margins) {

	// check if instance is already staged
	if !IsStaged(stage, system_margins) {
		return
	}

	system_margins.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTap(tap *Tap) {

	// check if instance is already staged
	if !IsStaged(stage, tap) {
		return
	}

	tap.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTechnical(technical *Technical) {

	// check if instance is already staged
	if !IsStaged(stage, technical) {
		return
	}

	technical.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if technical.Up_bow != nil {
		UnstageBranch(stage, technical.Up_bow)
	}
	if technical.Down_bow != nil {
		UnstageBranch(stage, technical.Down_bow)
	}
	if technical.Harmonic != nil {
		UnstageBranch(stage, technical.Harmonic)
	}
	if technical.Open_string != nil {
		UnstageBranch(stage, technical.Open_string)
	}
	if technical.Thumb_position != nil {
		UnstageBranch(stage, technical.Thumb_position)
	}
	if technical.Fingering != nil {
		UnstageBranch(stage, technical.Fingering)
	}
	if technical.Double_tongue != nil {
		UnstageBranch(stage, technical.Double_tongue)
	}
	if technical.Triple_tongue != nil {
		UnstageBranch(stage, technical.Triple_tongue)
	}
	if technical.Stopped != nil {
		UnstageBranch(stage, technical.Stopped)
	}
	if technical.Snap_pizzicato != nil {
		UnstageBranch(stage, technical.Snap_pizzicato)
	}
	if technical.Fret != nil {
		UnstageBranch(stage, technical.Fret)
	}
	if technical.Hammer_on != nil {
		UnstageBranch(stage, technical.Hammer_on)
	}
	if technical.Pull_off != nil {
		UnstageBranch(stage, technical.Pull_off)
	}
	if technical.Bend != nil {
		UnstageBranch(stage, technical.Bend)
	}
	if technical.Tap != nil {
		UnstageBranch(stage, technical.Tap)
	}
	if technical.Heel != nil {
		UnstageBranch(stage, technical.Heel)
	}
	if technical.Toe != nil {
		UnstageBranch(stage, technical.Toe)
	}
	if technical.Fingernails != nil {
		UnstageBranch(stage, technical.Fingernails)
	}
	if technical.Hole != nil {
		UnstageBranch(stage, technical.Hole)
	}
	if technical.Arrow != nil {
		UnstageBranch(stage, technical.Arrow)
	}
	if technical.Handbell != nil {
		UnstageBranch(stage, technical.Handbell)
	}
	if technical.Brass_bend != nil {
		UnstageBranch(stage, technical.Brass_bend)
	}
	if technical.Flip != nil {
		UnstageBranch(stage, technical.Flip)
	}
	if technical.Smear != nil {
		UnstageBranch(stage, technical.Smear)
	}
	if technical.Open != nil {
		UnstageBranch(stage, technical.Open)
	}
	if technical.Half_muted != nil {
		UnstageBranch(stage, technical.Half_muted)
	}
	if technical.Harmon_mute != nil {
		UnstageBranch(stage, technical.Harmon_mute)
	}
	if technical.Golpe != nil {
		UnstageBranch(stage, technical.Golpe)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchText_element_data(text_element_data *Text_element_data) {

	// check if instance is already staged
	if !IsStaged(stage, text_element_data) {
		return
	}

	text_element_data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTie(tie *Tie) {

	// check if instance is already staged
	if !IsStaged(stage, tie) {
		return
	}

	tie.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTied(tied *Tied) {

	// check if instance is already staged
	if !IsStaged(stage, tied) {
		return
	}

	tied.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTime(time *Time) {

	// check if instance is already staged
	if !IsStaged(stage, time) {
		return
	}

	time.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTime_modification(time_modification *Time_modification) {

	// check if instance is already staged
	if !IsStaged(stage, time_modification) {
		return
	}

	time_modification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTimpani(timpani *Timpani) {

	// check if instance is already staged
	if !IsStaged(stage, timpani) {
		return
	}

	timpani.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTranspose(transpose *Transpose) {

	// check if instance is already staged
	if !IsStaged(stage, transpose) {
		return
	}

	transpose.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTremolo(tremolo *Tremolo) {

	// check if instance is already staged
	if !IsStaged(stage, tremolo) {
		return
	}

	tremolo.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTuplet(tuplet *Tuplet) {

	// check if instance is already staged
	if !IsStaged(stage, tuplet) {
		return
	}

	tuplet.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet.Tuplet_actual != nil {
		UnstageBranch(stage, tuplet.Tuplet_actual)
	}
	if tuplet.Tuplet_normal != nil {
		UnstageBranch(stage, tuplet.Tuplet_normal)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTuplet_dot(tuplet_dot *Tuplet_dot) {

	// check if instance is already staged
	if !IsStaged(stage, tuplet_dot) {
		return
	}

	tuplet_dot.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTuplet_number(tuplet_number *Tuplet_number) {

	// check if instance is already staged
	if !IsStaged(stage, tuplet_number) {
		return
	}

	tuplet_number.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTuplet_portion(tuplet_portion *Tuplet_portion) {

	// check if instance is already staged
	if !IsStaged(stage, tuplet_portion) {
		return
	}

	tuplet_portion.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if tuplet_portion.Tuplet_number != nil {
		UnstageBranch(stage, tuplet_portion.Tuplet_number)
	}
	if tuplet_portion.Tuplet_type != nil {
		UnstageBranch(stage, tuplet_portion.Tuplet_type)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
		UnstageBranch(stage, _tuplet_dot)
	}

}

func (stage *StageStruct) UnstageBranchTuplet_type(tuplet_type *Tuplet_type) {

	// check if instance is already staged
	if !IsStaged(stage, tuplet_type) {
		return
	}

	tuplet_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchTyped_text(typed_text *Typed_text) {

	// check if instance is already staged
	if !IsStaged(stage, typed_text) {
		return
	}

	typed_text.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchUnpitched(unpitched *Unpitched) {

	// check if instance is already staged
	if !IsStaged(stage, unpitched) {
		return
	}

	unpitched.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchVirtual_instrument(virtual_instrument *Virtual_instrument) {

	// check if instance is already staged
	if !IsStaged(stage, virtual_instrument) {
		return
	}

	virtual_instrument.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWait(wait *Wait) {

	// check if instance is already staged
	if !IsStaged(stage, wait) {
		return
	}

	wait.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWavy_line(wavy_line *Wavy_line) {

	// check if instance is already staged
	if !IsStaged(stage, wavy_line) {
		return
	}

	wavy_line.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWedge(wedge *Wedge) {

	// check if instance is already staged
	if !IsStaged(stage, wedge) {
		return
	}

	wedge.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWood(wood *Wood) {

	// check if instance is already staged
	if !IsStaged(stage, wood) {
		return
	}

	wood.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchWork(work *Work) {

	// check if instance is already staged
	if !IsStaged(stage, work) {
		return
	}

	work.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if work.Opus != nil {
		UnstageBranch(stage, work.Opus)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}
