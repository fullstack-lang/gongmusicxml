// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Accidental:
		// insertion point per field

	case *Accidental_mark:
		// insertion point per field

	case *Accidental_text:
		// insertion point per field

	case *Accord:
		// insertion point per field

	case *Accordion_registration:
		// insertion point per field

	case *AnyType:
		// insertion point per field

	case *Appearance:
		// insertion point per field
		if fieldName == "Line_width" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Appearance) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Appearance)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Line_width).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Line_width = _inferedTypeInstance.Line_width[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Line_width =
								append(_inferedTypeInstance.Line_width, any(fieldInstance).(*Line_width))
						}
					}
				}
			}
		}
		if fieldName == "Note_size" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Appearance) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Appearance)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Note_size).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Note_size = _inferedTypeInstance.Note_size[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Note_size =
								append(_inferedTypeInstance.Note_size, any(fieldInstance).(*Note_size))
						}
					}
				}
			}
		}
		if fieldName == "Distance" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Appearance) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Appearance)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Distance).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Distance = _inferedTypeInstance.Distance[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Distance =
								append(_inferedTypeInstance.Distance, any(fieldInstance).(*Distance))
						}
					}
				}
			}
		}
		if fieldName == "Glyph" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Appearance) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Appearance)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Glyph).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Glyph = _inferedTypeInstance.Glyph[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Glyph =
								append(_inferedTypeInstance.Glyph, any(fieldInstance).(*Glyph))
						}
					}
				}
			}
		}
		if fieldName == "Other_appearance" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Appearance) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Appearance)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Other_appearance).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Other_appearance = _inferedTypeInstance.Other_appearance[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Other_appearance =
								append(_inferedTypeInstance.Other_appearance, any(fieldInstance).(*Other_appearance))
						}
					}
				}
			}
		}

	case *Arpeggiate:
		// insertion point per field

	case *Arrow:
		// insertion point per field

	case *Articulations:
		// insertion point per field

	case *Assess:
		// insertion point per field

	case *Attributes:
		// insertion point per field
		if fieldName == "Key" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Attributes) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Attributes)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Key).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Key = _inferedTypeInstance.Key[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Key =
								append(_inferedTypeInstance.Key, any(fieldInstance).(*Key))
						}
					}
				}
			}
		}
		if fieldName == "Clef" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Attributes) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Attributes)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Clef).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Clef = _inferedTypeInstance.Clef[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Clef =
								append(_inferedTypeInstance.Clef, any(fieldInstance).(*Clef))
						}
					}
				}
			}
		}
		if fieldName == "Staff_details" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Attributes) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Attributes)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Staff_details).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Staff_details = _inferedTypeInstance.Staff_details[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Staff_details =
								append(_inferedTypeInstance.Staff_details, any(fieldInstance).(*Staff_details))
						}
					}
				}
			}
		}
		if fieldName == "Measure_style" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Attributes) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Attributes)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Measure_style).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Measure_style = _inferedTypeInstance.Measure_style[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Measure_style =
								append(_inferedTypeInstance.Measure_style, any(fieldInstance).(*Measure_style))
						}
					}
				}
			}
		}
		if fieldName == "Transpose" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Attributes) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Attributes)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Transpose).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Transpose = _inferedTypeInstance.Transpose[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Transpose =
								append(_inferedTypeInstance.Transpose, any(fieldInstance).(*Transpose))
						}
					}
				}
			}
		}
		if fieldName == "For_part" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Attributes) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Attributes)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.For_part).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.For_part = _inferedTypeInstance.For_part[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.For_part =
								append(_inferedTypeInstance.For_part, any(fieldInstance).(*For_part))
						}
					}
				}
			}
		}

	case *Backup:
		// insertion point per field

	case *Bar_style_color:
		// insertion point per field

	case *Barline:
		// insertion point per field

	case *Barre:
		// insertion point per field

	case *Bass:
		// insertion point per field

	case *Bass_step:
		// insertion point per field

	case *Beam:
		// insertion point per field

	case *Beat_repeat:
		// insertion point per field

	case *Beat_unit_tied:
		// insertion point per field

	case *Beater:
		// insertion point per field

	case *Bend:
		// insertion point per field

	case *Bookmark:
		// insertion point per field

	case *Bracket:
		// insertion point per field

	case *Breath_mark:
		// insertion point per field

	case *Caesura:
		// insertion point per field

	case *Cancel:
		// insertion point per field

	case *Clef:
		// insertion point per field

	case *Coda:
		// insertion point per field

	case *Credit:
		// insertion point per field
		if fieldName == "Link" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Credit) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Credit)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Link).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Link = _inferedTypeInstance.Link[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Link =
								append(_inferedTypeInstance.Link, any(fieldInstance).(*Link))
						}
					}
				}
			}
		}
		if fieldName == "Bookmark" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Credit) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Credit)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Bookmark).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Bookmark = _inferedTypeInstance.Bookmark[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Bookmark =
								append(_inferedTypeInstance.Bookmark, any(fieldInstance).(*Bookmark))
						}
					}
				}
			}
		}

	case *Dashes:
		// insertion point per field

	case *Defaults:
		// insertion point per field
		if fieldName == "Lyric_font" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Defaults) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Defaults)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lyric_font).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lyric_font = _inferedTypeInstance.Lyric_font[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lyric_font =
								append(_inferedTypeInstance.Lyric_font, any(fieldInstance).(*Lyric_font))
						}
					}
				}
			}
		}
		if fieldName == "Lyric_language" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Defaults) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Defaults)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lyric_language).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lyric_language = _inferedTypeInstance.Lyric_language[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lyric_language =
								append(_inferedTypeInstance.Lyric_language, any(fieldInstance).(*Lyric_language))
						}
					}
				}
			}
		}

	case *Degree:
		// insertion point per field

	case *Degree_alter:
		// insertion point per field

	case *Degree_type:
		// insertion point per field

	case *Degree_value:
		// insertion point per field

	case *Direction:
		// insertion point per field
		if fieldName == "Direction_type" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Direction) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Direction)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Direction_type).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Direction_type = _inferedTypeInstance.Direction_type[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Direction_type =
								append(_inferedTypeInstance.Direction_type, any(fieldInstance).(*Direction_type))
						}
					}
				}
			}
		}

	case *Direction_type:
		// insertion point per field
		if fieldName == "Segno" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Direction_type) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Direction_type)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Segno).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Segno = _inferedTypeInstance.Segno[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Segno =
								append(_inferedTypeInstance.Segno, any(fieldInstance).(*Segno))
						}
					}
				}
			}
		}
		if fieldName == "Coda" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Direction_type) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Direction_type)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Coda).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Coda = _inferedTypeInstance.Coda[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Coda =
								append(_inferedTypeInstance.Coda, any(fieldInstance).(*Coda))
						}
					}
				}
			}
		}
		if fieldName == "Dynamics" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Direction_type) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Direction_type)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Dynamics).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Dynamics = _inferedTypeInstance.Dynamics[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Dynamics =
								append(_inferedTypeInstance.Dynamics, any(fieldInstance).(*Dynamics))
						}
					}
				}
			}
		}
		if fieldName == "Percussion" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Direction_type) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Direction_type)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Percussion).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Percussion = _inferedTypeInstance.Percussion[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Percussion =
								append(_inferedTypeInstance.Percussion, any(fieldInstance).(*Percussion))
						}
					}
				}
			}
		}

	case *Distance:
		// insertion point per field

	case *Double:
		// insertion point per field

	case *Dynamics:
		// insertion point per field

	case *Effect:
		// insertion point per field

	case *Elision:
		// insertion point per field

	case *Empty:
		// insertion point per field

	case *Empty_font:
		// insertion point per field

	case *Empty_line:
		// insertion point per field

	case *Empty_placement:
		// insertion point per field

	case *Empty_placement_smufl:
		// insertion point per field

	case *Empty_print_object_style_align:
		// insertion point per field

	case *Empty_print_style:
		// insertion point per field

	case *Empty_print_style_align:
		// insertion point per field

	case *Empty_print_style_align_id:
		// insertion point per field

	case *Empty_trill_sound:
		// insertion point per field

	case *Encoding:
		// insertion point per field

	case *Ending:
		// insertion point per field

	case *Extend:
		// insertion point per field

	case *Feature:
		// insertion point per field

	case *Fermata:
		// insertion point per field

	case *Figure:
		// insertion point per field

	case *Figured_bass:
		// insertion point per field
		if fieldName == "Figure" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Figured_bass) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Figured_bass)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Figure).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Figure = _inferedTypeInstance.Figure[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Figure =
								append(_inferedTypeInstance.Figure, any(fieldInstance).(*Figure))
						}
					}
				}
			}
		}

	case *Fingering:
		// insertion point per field

	case *First_fret:
		// insertion point per field

	case *Foo:
		// insertion point per field

	case *For_part:
		// insertion point per field

	case *Formatted_symbol:
		// insertion point per field

	case *Formatted_symbol_id:
		// insertion point per field

	case *Forward:
		// insertion point per field

	case *Frame:
		// insertion point per field
		if fieldName == "Frame_note" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Frame) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Frame)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Frame_note).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Frame_note = _inferedTypeInstance.Frame_note[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Frame_note =
								append(_inferedTypeInstance.Frame_note, any(fieldInstance).(*Frame_note))
						}
					}
				}
			}
		}

	case *Frame_note:
		// insertion point per field

	case *Fret:
		// insertion point per field

	case *Glass:
		// insertion point per field

	case *Glissando:
		// insertion point per field

	case *Glyph:
		// insertion point per field

	case *Grace:
		// insertion point per field

	case *Group_barline:
		// insertion point per field

	case *Group_symbol:
		// insertion point per field

	case *Grouping:
		// insertion point per field
		if fieldName == "Feature" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Grouping) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Grouping)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Feature).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Feature = _inferedTypeInstance.Feature[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Feature =
								append(_inferedTypeInstance.Feature, any(fieldInstance).(*Feature))
						}
					}
				}
			}
		}

	case *Hammer_on_pull_off:
		// insertion point per field

	case *Handbell:
		// insertion point per field

	case *Harmon_closed:
		// insertion point per field

	case *Harmon_mute:
		// insertion point per field

	case *Harmonic:
		// insertion point per field

	case *Harmony:
		// insertion point per field

	case *Harmony_alter:
		// insertion point per field

	case *Harp_pedals:
		// insertion point per field
		if fieldName == "Pedal_tuning" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Harp_pedals) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Harp_pedals)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Pedal_tuning).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Pedal_tuning = _inferedTypeInstance.Pedal_tuning[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Pedal_tuning =
								append(_inferedTypeInstance.Pedal_tuning, any(fieldInstance).(*Pedal_tuning))
						}
					}
				}
			}
		}

	case *Heel_toe:
		// insertion point per field

	case *Hole:
		// insertion point per field

	case *Hole_closed:
		// insertion point per field

	case *Horizontal_turn:
		// insertion point per field

	case *Identification:
		// insertion point per field
		if fieldName == "Creator" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Identification) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Identification)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Creator).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Creator = _inferedTypeInstance.Creator[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Creator =
								append(_inferedTypeInstance.Creator, any(fieldInstance).(*Typed_text))
						}
					}
				}
			}
		}
		if fieldName == "Rights" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Identification) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Identification)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Rights).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Rights = _inferedTypeInstance.Rights[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Rights =
								append(_inferedTypeInstance.Rights, any(fieldInstance).(*Typed_text))
						}
					}
				}
			}
		}
		if fieldName == "Relation" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Identification) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Identification)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Relation).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Relation = _inferedTypeInstance.Relation[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Relation =
								append(_inferedTypeInstance.Relation, any(fieldInstance).(*Typed_text))
						}
					}
				}
			}
		}

	case *Image:
		// insertion point per field

	case *Instrument:
		// insertion point per field

	case *Instrument_change:
		// insertion point per field

	case *Instrument_link:
		// insertion point per field

	case *Interchangeable:
		// insertion point per field

	case *Inversion:
		// insertion point per field

	case *Key:
		// insertion point per field
		if fieldName == "Key_octave" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Key) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Key)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Key_octave).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Key_octave = _inferedTypeInstance.Key_octave[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Key_octave =
								append(_inferedTypeInstance.Key_octave, any(fieldInstance).(*Key_octave))
						}
					}
				}
			}
		}

	case *Key_accidental:
		// insertion point per field

	case *Key_octave:
		// insertion point per field

	case *Kind:
		// insertion point per field

	case *Level:
		// insertion point per field

	case *Line_detail:
		// insertion point per field

	case *Line_width:
		// insertion point per field

	case *Link:
		// insertion point per field

	case *Listen:
		// insertion point per field

	case *Listening:
		// insertion point per field

	case *Lyric:
		// insertion point per field

	case *Lyric_font:
		// insertion point per field

	case *Lyric_language:
		// insertion point per field

	case *Measure_layout:
		// insertion point per field

	case *Measure_numbering:
		// insertion point per field

	case *Measure_repeat:
		// insertion point per field

	case *Measure_style:
		// insertion point per field

	case *Membrane:
		// insertion point per field

	case *Metal:
		// insertion point per field

	case *Metronome:
		// insertion point per field

	case *Metronome_beam:
		// insertion point per field

	case *Metronome_note:
		// insertion point per field
		if fieldName == "Metronome_dot" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Metronome_note) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Metronome_note)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Metronome_dot).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Metronome_dot = _inferedTypeInstance.Metronome_dot[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Metronome_dot =
								append(_inferedTypeInstance.Metronome_dot, any(fieldInstance).(*Empty))
						}
					}
				}
			}
		}
		if fieldName == "Metronome_beam" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Metronome_note) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Metronome_note)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Metronome_beam).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Metronome_beam = _inferedTypeInstance.Metronome_beam[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Metronome_beam =
								append(_inferedTypeInstance.Metronome_beam, any(fieldInstance).(*Metronome_beam))
						}
					}
				}
			}
		}

	case *Metronome_tied:
		// insertion point per field

	case *Metronome_tuplet:
		// insertion point per field

	case *Midi_device:
		// insertion point per field

	case *Midi_instrument:
		// insertion point per field

	case *Miscellaneous:
		// insertion point per field
		if fieldName == "Miscellaneous_field" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Miscellaneous) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Miscellaneous)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Miscellaneous_field).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Miscellaneous_field = _inferedTypeInstance.Miscellaneous_field[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Miscellaneous_field =
								append(_inferedTypeInstance.Miscellaneous_field, any(fieldInstance).(*Miscellaneous_field))
						}
					}
				}
			}
		}

	case *Miscellaneous_field:
		// insertion point per field

	case *Mordent:
		// insertion point per field

	case *Multiple_rest:
		// insertion point per field

	case *Name_display:
		// insertion point per field

	case *Non_arpeggiate:
		// insertion point per field

	case *Notations:
		// insertion point per field

	case *Note:
		// insertion point per field
		if fieldName == "Instrument" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Note) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Note)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Instrument).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Instrument = _inferedTypeInstance.Instrument[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Instrument =
								append(_inferedTypeInstance.Instrument, any(fieldInstance).(*Instrument))
						}
					}
				}
			}
		}
		if fieldName == "Dot" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Note) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Note)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Dot).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Dot = _inferedTypeInstance.Dot[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Dot =
								append(_inferedTypeInstance.Dot, any(fieldInstance).(*Empty_placement))
						}
					}
				}
			}
		}
		if fieldName == "Notations" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Note) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Note)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Notations).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Notations = _inferedTypeInstance.Notations[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Notations =
								append(_inferedTypeInstance.Notations, any(fieldInstance).(*Notations))
						}
					}
				}
			}
		}
		if fieldName == "Lyric" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Note) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Note)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lyric).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lyric = _inferedTypeInstance.Lyric[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lyric =
								append(_inferedTypeInstance.Lyric, any(fieldInstance).(*Lyric))
						}
					}
				}
			}
		}

	case *Note_size:
		// insertion point per field

	case *Note_type:
		// insertion point per field

	case *Notehead:
		// insertion point per field

	case *Notehead_text:
		// insertion point per field

	case *Numeral:
		// insertion point per field

	case *Numeral_key:
		// insertion point per field

	case *Numeral_root:
		// insertion point per field

	case *Octave_shift:
		// insertion point per field

	case *Offset:
		// insertion point per field

	case *Opus:
		// insertion point per field

	case *Ornaments:
		// insertion point per field
		if fieldName == "Accidental_mark" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Ornaments) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Ornaments)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Accidental_mark).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Accidental_mark = _inferedTypeInstance.Accidental_mark[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Accidental_mark =
								append(_inferedTypeInstance.Accidental_mark, any(fieldInstance).(*Accidental_mark))
						}
					}
				}
			}
		}

	case *Other_appearance:
		// insertion point per field

	case *Other_listening:
		// insertion point per field

	case *Other_notation:
		// insertion point per field

	case *Other_play:
		// insertion point per field

	case *Page_layout:
		// insertion point per field

	case *Page_margins:
		// insertion point per field

	case *Part_clef:
		// insertion point per field

	case *Part_group:
		// insertion point per field

	case *Part_link:
		// insertion point per field
		if fieldName == "Instrument_link" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Part_link) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Part_link)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Instrument_link).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Instrument_link = _inferedTypeInstance.Instrument_link[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Instrument_link =
								append(_inferedTypeInstance.Instrument_link, any(fieldInstance).(*Instrument_link))
						}
					}
				}
			}
		}

	case *Part_list:
		// insertion point per field

	case *Part_symbol:
		// insertion point per field

	case *Part_transpose:
		// insertion point per field

	case *Pedal:
		// insertion point per field

	case *Pedal_tuning:
		// insertion point per field

	case *Percussion:
		// insertion point per field

	case *Pitch:
		// insertion point per field

	case *Pitched:
		// insertion point per field

	case *Play:
		// insertion point per field

	case *Player:
		// insertion point per field

	case *Principal_voice:
		// insertion point per field

	case *Print:
		// insertion point per field

	case *Release:
		// insertion point per field

	case *Repeat:
		// insertion point per field

	case *Rest:
		// insertion point per field

	case *Root:
		// insertion point per field

	case *Root_step:
		// insertion point per field

	case *Scaling:
		// insertion point per field

	case *Scordatura:
		// insertion point per field
		if fieldName == "Accord" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Scordatura) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Scordatura)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Accord).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Accord = _inferedTypeInstance.Accord[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Accord =
								append(_inferedTypeInstance.Accord, any(fieldInstance).(*Accord))
						}
					}
				}
			}
		}

	case *Score_instrument:
		// insertion point per field

	case *Score_part:
		// insertion point per field
		if fieldName == "Part_link" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Score_part) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Score_part)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Part_link).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Part_link = _inferedTypeInstance.Part_link[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Part_link =
								append(_inferedTypeInstance.Part_link, any(fieldInstance).(*Part_link))
						}
					}
				}
			}
		}
		if fieldName == "Score_instrument" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Score_part) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Score_part)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Score_instrument).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Score_instrument = _inferedTypeInstance.Score_instrument[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Score_instrument =
								append(_inferedTypeInstance.Score_instrument, any(fieldInstance).(*Score_instrument))
						}
					}
				}
			}
		}
		if fieldName == "Player" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Score_part) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Score_part)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Player).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Player = _inferedTypeInstance.Player[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Player =
								append(_inferedTypeInstance.Player, any(fieldInstance).(*Player))
						}
					}
				}
			}
		}

	case *Score_partwise:
		// insertion point per field

	case *Score_timewise:
		// insertion point per field

	case *Segno:
		// insertion point per field

	case *Slash:
		// insertion point per field

	case *Slide:
		// insertion point per field

	case *Slur:
		// insertion point per field

	case *Sound:
		// insertion point per field

	case *Staff_details:
		// insertion point per field
		if fieldName == "Staff_tuning" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Staff_details) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Staff_details)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Staff_tuning).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Staff_tuning = _inferedTypeInstance.Staff_tuning[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Staff_tuning =
								append(_inferedTypeInstance.Staff_tuning, any(fieldInstance).(*Staff_tuning))
						}
					}
				}
			}
		}

	case *Staff_divide:
		// insertion point per field

	case *Staff_layout:
		// insertion point per field

	case *Staff_size:
		// insertion point per field

	case *Staff_tuning:
		// insertion point per field

	case *Stem:
		// insertion point per field

	case *Stick:
		// insertion point per field

	case *String_mute:
		// insertion point per field

	case *Strong_accent:
		// insertion point per field

	case *Supports:
		// insertion point per field

	case *Swing:
		// insertion point per field

	case *Sync:
		// insertion point per field

	case *System_dividers:
		// insertion point per field

	case *System_layout:
		// insertion point per field

	case *System_margins:
		// insertion point per field

	case *Tap:
		// insertion point per field

	case *Technical:
		// insertion point per field

	case *Text_element_data:
		// insertion point per field

	case *Tie:
		// insertion point per field

	case *Tied:
		// insertion point per field

	case *Time:
		// insertion point per field

	case *Time_modification:
		// insertion point per field

	case *Timpani:
		// insertion point per field

	case *Transpose:
		// insertion point per field

	case *Tremolo:
		// insertion point per field

	case *Tuplet:
		// insertion point per field

	case *Tuplet_dot:
		// insertion point per field

	case *Tuplet_number:
		// insertion point per field

	case *Tuplet_portion:
		// insertion point per field
		if fieldName == "Tuplet_dot" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Tuplet_portion) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Tuplet_portion)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Tuplet_dot).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Tuplet_dot = _inferedTypeInstance.Tuplet_dot[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Tuplet_dot =
								append(_inferedTypeInstance.Tuplet_dot, any(fieldInstance).(*Tuplet_dot))
						}
					}
				}
			}
		}

	case *Tuplet_type:
		// insertion point per field

	case *Typed_text:
		// insertion point per field

	case *Unpitched:
		// insertion point per field

	case *Virtual_instrument:
		// insertion point per field

	case *Wait:
		// insertion point per field

	case *Wavy_line:
		// insertion point per field

	case *Wedge:
		// insertion point per field

	case *Wood:
		// insertion point per field

	case *Work:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Accidental
	// insertion point per field

	// Compute reverse map for named struct Accidental_mark
	// insertion point per field

	// Compute reverse map for named struct Accidental_text
	// insertion point per field

	// Compute reverse map for named struct Accord
	// insertion point per field

	// Compute reverse map for named struct Accordion_registration
	// insertion point per field

	// Compute reverse map for named struct AnyType
	// insertion point per field

	// Compute reverse map for named struct Appearance
	// insertion point per field
	clear(stage.Appearance_Line_width_reverseMap)
	stage.Appearance_Line_width_reverseMap = make(map[*Line_width]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _line_width := range appearance.Line_width {
			stage.Appearance_Line_width_reverseMap[_line_width] = appearance
		}
	}
	clear(stage.Appearance_Note_size_reverseMap)
	stage.Appearance_Note_size_reverseMap = make(map[*Note_size]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _note_size := range appearance.Note_size {
			stage.Appearance_Note_size_reverseMap[_note_size] = appearance
		}
	}
	clear(stage.Appearance_Distance_reverseMap)
	stage.Appearance_Distance_reverseMap = make(map[*Distance]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _distance := range appearance.Distance {
			stage.Appearance_Distance_reverseMap[_distance] = appearance
		}
	}
	clear(stage.Appearance_Glyph_reverseMap)
	stage.Appearance_Glyph_reverseMap = make(map[*Glyph]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _glyph := range appearance.Glyph {
			stage.Appearance_Glyph_reverseMap[_glyph] = appearance
		}
	}
	clear(stage.Appearance_Other_appearance_reverseMap)
	stage.Appearance_Other_appearance_reverseMap = make(map[*Other_appearance]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _other_appearance := range appearance.Other_appearance {
			stage.Appearance_Other_appearance_reverseMap[_other_appearance] = appearance
		}
	}

	// Compute reverse map for named struct Arpeggiate
	// insertion point per field

	// Compute reverse map for named struct Arrow
	// insertion point per field

	// Compute reverse map for named struct Articulations
	// insertion point per field

	// Compute reverse map for named struct Assess
	// insertion point per field

	// Compute reverse map for named struct Attributes
	// insertion point per field
	clear(stage.Attributes_Key_reverseMap)
	stage.Attributes_Key_reverseMap = make(map[*Key]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _key := range attributes.Key {
			stage.Attributes_Key_reverseMap[_key] = attributes
		}
	}
	clear(stage.Attributes_Clef_reverseMap)
	stage.Attributes_Clef_reverseMap = make(map[*Clef]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _clef := range attributes.Clef {
			stage.Attributes_Clef_reverseMap[_clef] = attributes
		}
	}
	clear(stage.Attributes_Staff_details_reverseMap)
	stage.Attributes_Staff_details_reverseMap = make(map[*Staff_details]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _staff_details := range attributes.Staff_details {
			stage.Attributes_Staff_details_reverseMap[_staff_details] = attributes
		}
	}
	clear(stage.Attributes_Measure_style_reverseMap)
	stage.Attributes_Measure_style_reverseMap = make(map[*Measure_style]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _measure_style := range attributes.Measure_style {
			stage.Attributes_Measure_style_reverseMap[_measure_style] = attributes
		}
	}
	clear(stage.Attributes_Transpose_reverseMap)
	stage.Attributes_Transpose_reverseMap = make(map[*Transpose]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _transpose := range attributes.Transpose {
			stage.Attributes_Transpose_reverseMap[_transpose] = attributes
		}
	}
	clear(stage.Attributes_For_part_reverseMap)
	stage.Attributes_For_part_reverseMap = make(map[*For_part]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _for_part := range attributes.For_part {
			stage.Attributes_For_part_reverseMap[_for_part] = attributes
		}
	}

	// Compute reverse map for named struct Backup
	// insertion point per field

	// Compute reverse map for named struct Bar_style_color
	// insertion point per field

	// Compute reverse map for named struct Barline
	// insertion point per field

	// Compute reverse map for named struct Barre
	// insertion point per field

	// Compute reverse map for named struct Bass
	// insertion point per field

	// Compute reverse map for named struct Bass_step
	// insertion point per field

	// Compute reverse map for named struct Beam
	// insertion point per field

	// Compute reverse map for named struct Beat_repeat
	// insertion point per field

	// Compute reverse map for named struct Beat_unit_tied
	// insertion point per field

	// Compute reverse map for named struct Beater
	// insertion point per field

	// Compute reverse map for named struct Bend
	// insertion point per field

	// Compute reverse map for named struct Bookmark
	// insertion point per field

	// Compute reverse map for named struct Bracket
	// insertion point per field

	// Compute reverse map for named struct Breath_mark
	// insertion point per field

	// Compute reverse map for named struct Caesura
	// insertion point per field

	// Compute reverse map for named struct Cancel
	// insertion point per field

	// Compute reverse map for named struct Clef
	// insertion point per field

	// Compute reverse map for named struct Coda
	// insertion point per field

	// Compute reverse map for named struct Credit
	// insertion point per field
	clear(stage.Credit_Link_reverseMap)
	stage.Credit_Link_reverseMap = make(map[*Link]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _link := range credit.Link {
			stage.Credit_Link_reverseMap[_link] = credit
		}
	}
	clear(stage.Credit_Bookmark_reverseMap)
	stage.Credit_Bookmark_reverseMap = make(map[*Bookmark]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _bookmark := range credit.Bookmark {
			stage.Credit_Bookmark_reverseMap[_bookmark] = credit
		}
	}

	// Compute reverse map for named struct Dashes
	// insertion point per field

	// Compute reverse map for named struct Defaults
	// insertion point per field
	clear(stage.Defaults_Lyric_font_reverseMap)
	stage.Defaults_Lyric_font_reverseMap = make(map[*Lyric_font]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _lyric_font := range defaults.Lyric_font {
			stage.Defaults_Lyric_font_reverseMap[_lyric_font] = defaults
		}
	}
	clear(stage.Defaults_Lyric_language_reverseMap)
	stage.Defaults_Lyric_language_reverseMap = make(map[*Lyric_language]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _lyric_language := range defaults.Lyric_language {
			stage.Defaults_Lyric_language_reverseMap[_lyric_language] = defaults
		}
	}

	// Compute reverse map for named struct Degree
	// insertion point per field

	// Compute reverse map for named struct Degree_alter
	// insertion point per field

	// Compute reverse map for named struct Degree_type
	// insertion point per field

	// Compute reverse map for named struct Degree_value
	// insertion point per field

	// Compute reverse map for named struct Direction
	// insertion point per field
	clear(stage.Direction_Direction_type_reverseMap)
	stage.Direction_Direction_type_reverseMap = make(map[*Direction_type]*Direction)
	for direction := range stage.Directions {
		_ = direction
		for _, _direction_type := range direction.Direction_type {
			stage.Direction_Direction_type_reverseMap[_direction_type] = direction
		}
	}

	// Compute reverse map for named struct Direction_type
	// insertion point per field
	clear(stage.Direction_type_Segno_reverseMap)
	stage.Direction_type_Segno_reverseMap = make(map[*Segno]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _segno := range direction_type.Segno {
			stage.Direction_type_Segno_reverseMap[_segno] = direction_type
		}
	}
	clear(stage.Direction_type_Coda_reverseMap)
	stage.Direction_type_Coda_reverseMap = make(map[*Coda]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _coda := range direction_type.Coda {
			stage.Direction_type_Coda_reverseMap[_coda] = direction_type
		}
	}
	clear(stage.Direction_type_Dynamics_reverseMap)
	stage.Direction_type_Dynamics_reverseMap = make(map[*Dynamics]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _dynamics := range direction_type.Dynamics {
			stage.Direction_type_Dynamics_reverseMap[_dynamics] = direction_type
		}
	}
	clear(stage.Direction_type_Percussion_reverseMap)
	stage.Direction_type_Percussion_reverseMap = make(map[*Percussion]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _percussion := range direction_type.Percussion {
			stage.Direction_type_Percussion_reverseMap[_percussion] = direction_type
		}
	}

	// Compute reverse map for named struct Distance
	// insertion point per field

	// Compute reverse map for named struct Double
	// insertion point per field

	// Compute reverse map for named struct Dynamics
	// insertion point per field

	// Compute reverse map for named struct Effect
	// insertion point per field

	// Compute reverse map for named struct Elision
	// insertion point per field

	// Compute reverse map for named struct Empty
	// insertion point per field

	// Compute reverse map for named struct Empty_font
	// insertion point per field

	// Compute reverse map for named struct Empty_line
	// insertion point per field

	// Compute reverse map for named struct Empty_placement
	// insertion point per field

	// Compute reverse map for named struct Empty_placement_smufl
	// insertion point per field

	// Compute reverse map for named struct Empty_print_object_style_align
	// insertion point per field

	// Compute reverse map for named struct Empty_print_style
	// insertion point per field

	// Compute reverse map for named struct Empty_print_style_align
	// insertion point per field

	// Compute reverse map for named struct Empty_print_style_align_id
	// insertion point per field

	// Compute reverse map for named struct Empty_trill_sound
	// insertion point per field

	// Compute reverse map for named struct Encoding
	// insertion point per field

	// Compute reverse map for named struct Ending
	// insertion point per field

	// Compute reverse map for named struct Extend
	// insertion point per field

	// Compute reverse map for named struct Feature
	// insertion point per field

	// Compute reverse map for named struct Fermata
	// insertion point per field

	// Compute reverse map for named struct Figure
	// insertion point per field

	// Compute reverse map for named struct Figured_bass
	// insertion point per field
	clear(stage.Figured_bass_Figure_reverseMap)
	stage.Figured_bass_Figure_reverseMap = make(map[*Figure]*Figured_bass)
	for figured_bass := range stage.Figured_basss {
		_ = figured_bass
		for _, _figure := range figured_bass.Figure {
			stage.Figured_bass_Figure_reverseMap[_figure] = figured_bass
		}
	}

	// Compute reverse map for named struct Fingering
	// insertion point per field

	// Compute reverse map for named struct First_fret
	// insertion point per field

	// Compute reverse map for named struct Foo
	// insertion point per field

	// Compute reverse map for named struct For_part
	// insertion point per field

	// Compute reverse map for named struct Formatted_symbol
	// insertion point per field

	// Compute reverse map for named struct Formatted_symbol_id
	// insertion point per field

	// Compute reverse map for named struct Forward
	// insertion point per field

	// Compute reverse map for named struct Frame
	// insertion point per field
	clear(stage.Frame_Frame_note_reverseMap)
	stage.Frame_Frame_note_reverseMap = make(map[*Frame_note]*Frame)
	for frame := range stage.Frames {
		_ = frame
		for _, _frame_note := range frame.Frame_note {
			stage.Frame_Frame_note_reverseMap[_frame_note] = frame
		}
	}

	// Compute reverse map for named struct Frame_note
	// insertion point per field

	// Compute reverse map for named struct Fret
	// insertion point per field

	// Compute reverse map for named struct Glass
	// insertion point per field

	// Compute reverse map for named struct Glissando
	// insertion point per field

	// Compute reverse map for named struct Glyph
	// insertion point per field

	// Compute reverse map for named struct Grace
	// insertion point per field

	// Compute reverse map for named struct Group_barline
	// insertion point per field

	// Compute reverse map for named struct Group_symbol
	// insertion point per field

	// Compute reverse map for named struct Grouping
	// insertion point per field
	clear(stage.Grouping_Feature_reverseMap)
	stage.Grouping_Feature_reverseMap = make(map[*Feature]*Grouping)
	for grouping := range stage.Groupings {
		_ = grouping
		for _, _feature := range grouping.Feature {
			stage.Grouping_Feature_reverseMap[_feature] = grouping
		}
	}

	// Compute reverse map for named struct Hammer_on_pull_off
	// insertion point per field

	// Compute reverse map for named struct Handbell
	// insertion point per field

	// Compute reverse map for named struct Harmon_closed
	// insertion point per field

	// Compute reverse map for named struct Harmon_mute
	// insertion point per field

	// Compute reverse map for named struct Harmonic
	// insertion point per field

	// Compute reverse map for named struct Harmony
	// insertion point per field

	// Compute reverse map for named struct Harmony_alter
	// insertion point per field

	// Compute reverse map for named struct Harp_pedals
	// insertion point per field
	clear(stage.Harp_pedals_Pedal_tuning_reverseMap)
	stage.Harp_pedals_Pedal_tuning_reverseMap = make(map[*Pedal_tuning]*Harp_pedals)
	for harp_pedals := range stage.Harp_pedalss {
		_ = harp_pedals
		for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
			stage.Harp_pedals_Pedal_tuning_reverseMap[_pedal_tuning] = harp_pedals
		}
	}

	// Compute reverse map for named struct Heel_toe
	// insertion point per field

	// Compute reverse map for named struct Hole
	// insertion point per field

	// Compute reverse map for named struct Hole_closed
	// insertion point per field

	// Compute reverse map for named struct Horizontal_turn
	// insertion point per field

	// Compute reverse map for named struct Identification
	// insertion point per field
	clear(stage.Identification_Creator_reverseMap)
	stage.Identification_Creator_reverseMap = make(map[*Typed_text]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typed_text := range identification.Creator {
			stage.Identification_Creator_reverseMap[_typed_text] = identification
		}
	}
	clear(stage.Identification_Rights_reverseMap)
	stage.Identification_Rights_reverseMap = make(map[*Typed_text]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typed_text := range identification.Rights {
			stage.Identification_Rights_reverseMap[_typed_text] = identification
		}
	}
	clear(stage.Identification_Relation_reverseMap)
	stage.Identification_Relation_reverseMap = make(map[*Typed_text]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typed_text := range identification.Relation {
			stage.Identification_Relation_reverseMap[_typed_text] = identification
		}
	}

	// Compute reverse map for named struct Image
	// insertion point per field

	// Compute reverse map for named struct Instrument
	// insertion point per field

	// Compute reverse map for named struct Instrument_change
	// insertion point per field

	// Compute reverse map for named struct Instrument_link
	// insertion point per field

	// Compute reverse map for named struct Interchangeable
	// insertion point per field

	// Compute reverse map for named struct Inversion
	// insertion point per field

	// Compute reverse map for named struct Key
	// insertion point per field
	clear(stage.Key_Key_octave_reverseMap)
	stage.Key_Key_octave_reverseMap = make(map[*Key_octave]*Key)
	for key := range stage.Keys {
		_ = key
		for _, _key_octave := range key.Key_octave {
			stage.Key_Key_octave_reverseMap[_key_octave] = key
		}
	}

	// Compute reverse map for named struct Key_accidental
	// insertion point per field

	// Compute reverse map for named struct Key_octave
	// insertion point per field

	// Compute reverse map for named struct Kind
	// insertion point per field

	// Compute reverse map for named struct Level
	// insertion point per field

	// Compute reverse map for named struct Line_detail
	// insertion point per field

	// Compute reverse map for named struct Line_width
	// insertion point per field

	// Compute reverse map for named struct Link
	// insertion point per field

	// Compute reverse map for named struct Listen
	// insertion point per field

	// Compute reverse map for named struct Listening
	// insertion point per field

	// Compute reverse map for named struct Lyric
	// insertion point per field

	// Compute reverse map for named struct Lyric_font
	// insertion point per field

	// Compute reverse map for named struct Lyric_language
	// insertion point per field

	// Compute reverse map for named struct Measure_layout
	// insertion point per field

	// Compute reverse map for named struct Measure_numbering
	// insertion point per field

	// Compute reverse map for named struct Measure_repeat
	// insertion point per field

	// Compute reverse map for named struct Measure_style
	// insertion point per field

	// Compute reverse map for named struct Membrane
	// insertion point per field

	// Compute reverse map for named struct Metal
	// insertion point per field

	// Compute reverse map for named struct Metronome
	// insertion point per field

	// Compute reverse map for named struct Metronome_beam
	// insertion point per field

	// Compute reverse map for named struct Metronome_note
	// insertion point per field
	clear(stage.Metronome_note_Metronome_dot_reverseMap)
	stage.Metronome_note_Metronome_dot_reverseMap = make(map[*Empty]*Metronome_note)
	for metronome_note := range stage.Metronome_notes {
		_ = metronome_note
		for _, _empty := range metronome_note.Metronome_dot {
			stage.Metronome_note_Metronome_dot_reverseMap[_empty] = metronome_note
		}
	}
	clear(stage.Metronome_note_Metronome_beam_reverseMap)
	stage.Metronome_note_Metronome_beam_reverseMap = make(map[*Metronome_beam]*Metronome_note)
	for metronome_note := range stage.Metronome_notes {
		_ = metronome_note
		for _, _metronome_beam := range metronome_note.Metronome_beam {
			stage.Metronome_note_Metronome_beam_reverseMap[_metronome_beam] = metronome_note
		}
	}

	// Compute reverse map for named struct Metronome_tied
	// insertion point per field

	// Compute reverse map for named struct Metronome_tuplet
	// insertion point per field

	// Compute reverse map for named struct Midi_device
	// insertion point per field

	// Compute reverse map for named struct Midi_instrument
	// insertion point per field

	// Compute reverse map for named struct Miscellaneous
	// insertion point per field
	clear(stage.Miscellaneous_Miscellaneous_field_reverseMap)
	stage.Miscellaneous_Miscellaneous_field_reverseMap = make(map[*Miscellaneous_field]*Miscellaneous)
	for miscellaneous := range stage.Miscellaneouss {
		_ = miscellaneous
		for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
			stage.Miscellaneous_Miscellaneous_field_reverseMap[_miscellaneous_field] = miscellaneous
		}
	}

	// Compute reverse map for named struct Miscellaneous_field
	// insertion point per field

	// Compute reverse map for named struct Mordent
	// insertion point per field

	// Compute reverse map for named struct Multiple_rest
	// insertion point per field

	// Compute reverse map for named struct Name_display
	// insertion point per field

	// Compute reverse map for named struct Non_arpeggiate
	// insertion point per field

	// Compute reverse map for named struct Notations
	// insertion point per field

	// Compute reverse map for named struct Note
	// insertion point per field
	clear(stage.Note_Instrument_reverseMap)
	stage.Note_Instrument_reverseMap = make(map[*Instrument]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _instrument := range note.Instrument {
			stage.Note_Instrument_reverseMap[_instrument] = note
		}
	}
	clear(stage.Note_Dot_reverseMap)
	stage.Note_Dot_reverseMap = make(map[*Empty_placement]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _empty_placement := range note.Dot {
			stage.Note_Dot_reverseMap[_empty_placement] = note
		}
	}
	clear(stage.Note_Notations_reverseMap)
	stage.Note_Notations_reverseMap = make(map[*Notations]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _notations := range note.Notations {
			stage.Note_Notations_reverseMap[_notations] = note
		}
	}
	clear(stage.Note_Lyric_reverseMap)
	stage.Note_Lyric_reverseMap = make(map[*Lyric]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _lyric := range note.Lyric {
			stage.Note_Lyric_reverseMap[_lyric] = note
		}
	}

	// Compute reverse map for named struct Note_size
	// insertion point per field

	// Compute reverse map for named struct Note_type
	// insertion point per field

	// Compute reverse map for named struct Notehead
	// insertion point per field

	// Compute reverse map for named struct Notehead_text
	// insertion point per field

	// Compute reverse map for named struct Numeral
	// insertion point per field

	// Compute reverse map for named struct Numeral_key
	// insertion point per field

	// Compute reverse map for named struct Numeral_root
	// insertion point per field

	// Compute reverse map for named struct Octave_shift
	// insertion point per field

	// Compute reverse map for named struct Offset
	// insertion point per field

	// Compute reverse map for named struct Opus
	// insertion point per field

	// Compute reverse map for named struct Ornaments
	// insertion point per field
	clear(stage.Ornaments_Accidental_mark_reverseMap)
	stage.Ornaments_Accidental_mark_reverseMap = make(map[*Accidental_mark]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _accidental_mark := range ornaments.Accidental_mark {
			stage.Ornaments_Accidental_mark_reverseMap[_accidental_mark] = ornaments
		}
	}

	// Compute reverse map for named struct Other_appearance
	// insertion point per field

	// Compute reverse map for named struct Other_listening
	// insertion point per field

	// Compute reverse map for named struct Other_notation
	// insertion point per field

	// Compute reverse map for named struct Other_play
	// insertion point per field

	// Compute reverse map for named struct Page_layout
	// insertion point per field

	// Compute reverse map for named struct Page_margins
	// insertion point per field

	// Compute reverse map for named struct Part_clef
	// insertion point per field

	// Compute reverse map for named struct Part_group
	// insertion point per field

	// Compute reverse map for named struct Part_link
	// insertion point per field
	clear(stage.Part_link_Instrument_link_reverseMap)
	stage.Part_link_Instrument_link_reverseMap = make(map[*Instrument_link]*Part_link)
	for part_link := range stage.Part_links {
		_ = part_link
		for _, _instrument_link := range part_link.Instrument_link {
			stage.Part_link_Instrument_link_reverseMap[_instrument_link] = part_link
		}
	}

	// Compute reverse map for named struct Part_list
	// insertion point per field

	// Compute reverse map for named struct Part_symbol
	// insertion point per field

	// Compute reverse map for named struct Part_transpose
	// insertion point per field

	// Compute reverse map for named struct Pedal
	// insertion point per field

	// Compute reverse map for named struct Pedal_tuning
	// insertion point per field

	// Compute reverse map for named struct Percussion
	// insertion point per field

	// Compute reverse map for named struct Pitch
	// insertion point per field

	// Compute reverse map for named struct Pitched
	// insertion point per field

	// Compute reverse map for named struct Play
	// insertion point per field

	// Compute reverse map for named struct Player
	// insertion point per field

	// Compute reverse map for named struct Principal_voice
	// insertion point per field

	// Compute reverse map for named struct Print
	// insertion point per field

	// Compute reverse map for named struct Release
	// insertion point per field

	// Compute reverse map for named struct Repeat
	// insertion point per field

	// Compute reverse map for named struct Rest
	// insertion point per field

	// Compute reverse map for named struct Root
	// insertion point per field

	// Compute reverse map for named struct Root_step
	// insertion point per field

	// Compute reverse map for named struct Scaling
	// insertion point per field

	// Compute reverse map for named struct Scordatura
	// insertion point per field
	clear(stage.Scordatura_Accord_reverseMap)
	stage.Scordatura_Accord_reverseMap = make(map[*Accord]*Scordatura)
	for scordatura := range stage.Scordaturas {
		_ = scordatura
		for _, _accord := range scordatura.Accord {
			stage.Scordatura_Accord_reverseMap[_accord] = scordatura
		}
	}

	// Compute reverse map for named struct Score_instrument
	// insertion point per field

	// Compute reverse map for named struct Score_part
	// insertion point per field
	clear(stage.Score_part_Part_link_reverseMap)
	stage.Score_part_Part_link_reverseMap = make(map[*Part_link]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _part_link := range score_part.Part_link {
			stage.Score_part_Part_link_reverseMap[_part_link] = score_part
		}
	}
	clear(stage.Score_part_Score_instrument_reverseMap)
	stage.Score_part_Score_instrument_reverseMap = make(map[*Score_instrument]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _score_instrument := range score_part.Score_instrument {
			stage.Score_part_Score_instrument_reverseMap[_score_instrument] = score_part
		}
	}
	clear(stage.Score_part_Player_reverseMap)
	stage.Score_part_Player_reverseMap = make(map[*Player]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _player := range score_part.Player {
			stage.Score_part_Player_reverseMap[_player] = score_part
		}
	}

	// Compute reverse map for named struct Score_partwise
	// insertion point per field

	// Compute reverse map for named struct Score_timewise
	// insertion point per field

	// Compute reverse map for named struct Segno
	// insertion point per field

	// Compute reverse map for named struct Slash
	// insertion point per field

	// Compute reverse map for named struct Slide
	// insertion point per field

	// Compute reverse map for named struct Slur
	// insertion point per field

	// Compute reverse map for named struct Sound
	// insertion point per field

	// Compute reverse map for named struct Staff_details
	// insertion point per field
	clear(stage.Staff_details_Staff_tuning_reverseMap)
	stage.Staff_details_Staff_tuning_reverseMap = make(map[*Staff_tuning]*Staff_details)
	for staff_details := range stage.Staff_detailss {
		_ = staff_details
		for _, _staff_tuning := range staff_details.Staff_tuning {
			stage.Staff_details_Staff_tuning_reverseMap[_staff_tuning] = staff_details
		}
	}

	// Compute reverse map for named struct Staff_divide
	// insertion point per field

	// Compute reverse map for named struct Staff_layout
	// insertion point per field

	// Compute reverse map for named struct Staff_size
	// insertion point per field

	// Compute reverse map for named struct Staff_tuning
	// insertion point per field

	// Compute reverse map for named struct Stem
	// insertion point per field

	// Compute reverse map for named struct Stick
	// insertion point per field

	// Compute reverse map for named struct String_mute
	// insertion point per field

	// Compute reverse map for named struct Strong_accent
	// insertion point per field

	// Compute reverse map for named struct Supports
	// insertion point per field

	// Compute reverse map for named struct Swing
	// insertion point per field

	// Compute reverse map for named struct Sync
	// insertion point per field

	// Compute reverse map for named struct System_dividers
	// insertion point per field

	// Compute reverse map for named struct System_layout
	// insertion point per field

	// Compute reverse map for named struct System_margins
	// insertion point per field

	// Compute reverse map for named struct Tap
	// insertion point per field

	// Compute reverse map for named struct Technical
	// insertion point per field

	// Compute reverse map for named struct Text_element_data
	// insertion point per field

	// Compute reverse map for named struct Tie
	// insertion point per field

	// Compute reverse map for named struct Tied
	// insertion point per field

	// Compute reverse map for named struct Time
	// insertion point per field

	// Compute reverse map for named struct Time_modification
	// insertion point per field

	// Compute reverse map for named struct Timpani
	// insertion point per field

	// Compute reverse map for named struct Transpose
	// insertion point per field

	// Compute reverse map for named struct Tremolo
	// insertion point per field

	// Compute reverse map for named struct Tuplet
	// insertion point per field

	// Compute reverse map for named struct Tuplet_dot
	// insertion point per field

	// Compute reverse map for named struct Tuplet_number
	// insertion point per field

	// Compute reverse map for named struct Tuplet_portion
	// insertion point per field
	clear(stage.Tuplet_portion_Tuplet_dot_reverseMap)
	stage.Tuplet_portion_Tuplet_dot_reverseMap = make(map[*Tuplet_dot]*Tuplet_portion)
	for tuplet_portion := range stage.Tuplet_portions {
		_ = tuplet_portion
		for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
			stage.Tuplet_portion_Tuplet_dot_reverseMap[_tuplet_dot] = tuplet_portion
		}
	}

	// Compute reverse map for named struct Tuplet_type
	// insertion point per field

	// Compute reverse map for named struct Typed_text
	// insertion point per field

	// Compute reverse map for named struct Unpitched
	// insertion point per field

	// Compute reverse map for named struct Virtual_instrument
	// insertion point per field

	// Compute reverse map for named struct Wait
	// insertion point per field

	// Compute reverse map for named struct Wavy_line
	// insertion point per field

	// Compute reverse map for named struct Wedge
	// insertion point per field

	// Compute reverse map for named struct Wood
	// insertion point per field

	// Compute reverse map for named struct Work
	// insertion point per field

}
