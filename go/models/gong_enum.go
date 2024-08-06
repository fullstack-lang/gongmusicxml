// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for Above_below
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (above_below Above_below) ToString() (res string) {

	// migration of former implementation of enum
	switch above_below {
	// insertion code per enum code
	case Above_belowAbove:
		res = "above"
	case Above_belowBelow:
		res = "below"
	}
	return
}

func (above_below *Above_below) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "above":
		*above_below = Above_belowAbove
		return
	case "below":
		*above_below = Above_belowBelow
		return
	default:
		return errUnkownEnum
	}
}

func (above_below *Above_below) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Above_belowAbove":
		*above_below = Above_belowAbove
	case "Above_belowBelow":
		*above_below = Above_belowBelow
	default:
		return errUnkownEnum
	}
	return
}

func (above_below *Above_below) ToCodeString() (res string) {

	switch *above_below {
	// insertion code per enum code
	case Above_belowAbove:
		res = "Above_belowAbove"
	case Above_belowBelow:
		res = "Above_belowBelow"
	}
	return
}

func (above_below Above_below) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Above_belowAbove")
	res = append(res, "Above_belowBelow")

	return
}

func (above_below Above_below) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "above")
	res = append(res, "below")

	return
}

// Utility function for Accidental_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (accidental_value Accidental_value) ToString() (res string) {

	// migration of former implementation of enum
	switch accidental_value {
	// insertion code per enum code
	case Accidental_valueSharp:
		res = "sharp"
	case Accidental_valueNatural:
		res = "natural"
	case Accidental_valueFlat:
		res = "flat"
	case Accidental_valueDouble_sharp:
		res = "double-sharp"
	case Accidental_valueSharp_sharp:
		res = "sharp-sharp"
	case Accidental_valueFlat_flat:
		res = "flat-flat"
	case Accidental_valueNatural_sharp:
		res = "natural-sharp"
	case Accidental_valueNatural_flat:
		res = "natural-flat"
	case Accidental_valueQuarter_flat:
		res = "quarter-flat"
	case Accidental_valueQuarter_sharp:
		res = "quarter-sharp"
	case Accidental_valueThree_quarters_flat:
		res = "three-quarters-flat"
	case Accidental_valueThree_quarters_sharp:
		res = "three-quarters-sharp"
	case Accidental_valueSharp_down:
		res = "sharp-down"
	case Accidental_valueSharp_up:
		res = "sharp-up"
	case Accidental_valueNatural_down:
		res = "natural-down"
	case Accidental_valueNatural_up:
		res = "natural-up"
	case Accidental_valueFlat_down:
		res = "flat-down"
	case Accidental_valueFlat_up:
		res = "flat-up"
	case Accidental_valueDouble_sharp_down:
		res = "double-sharp-down"
	case Accidental_valueDouble_sharp_up:
		res = "double-sharp-up"
	case Accidental_valueFlat_flat_down:
		res = "flat-flat-down"
	case Accidental_valueFlat_flat_up:
		res = "flat-flat-up"
	case Accidental_valueArrow_down:
		res = "arrow-down"
	case Accidental_valueArrow_up:
		res = "arrow-up"
	case Accidental_valueTriple_sharp:
		res = "triple-sharp"
	case Accidental_valueTriple_flat:
		res = "triple-flat"
	case Accidental_valueSlash_quarter_sharp:
		res = "slash-quarter-sharp"
	case Accidental_valueSlash_sharp:
		res = "slash-sharp"
	case Accidental_valueSlash_flat:
		res = "slash-flat"
	case Accidental_valueDouble_slash_flat:
		res = "double-slash-flat"
	case Accidental_valueSharp_1:
		res = "sharp-1"
	case Accidental_valueSharp_2:
		res = "sharp-2"
	case Accidental_valueSharp_3:
		res = "sharp-3"
	case Accidental_valueSharp_5:
		res = "sharp-5"
	case Accidental_valueFlat_1:
		res = "flat-1"
	case Accidental_valueFlat_2:
		res = "flat-2"
	case Accidental_valueFlat_3:
		res = "flat-3"
	case Accidental_valueFlat_4:
		res = "flat-4"
	case Accidental_valueSori:
		res = "sori"
	case Accidental_valueKoron:
		res = "koron"
	case Accidental_valueOther:
		res = "other"
	}
	return
}

func (accidental_value *Accidental_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "sharp":
		*accidental_value = Accidental_valueSharp
		return
	case "natural":
		*accidental_value = Accidental_valueNatural
		return
	case "flat":
		*accidental_value = Accidental_valueFlat
		return
	case "double-sharp":
		*accidental_value = Accidental_valueDouble_sharp
		return
	case "sharp-sharp":
		*accidental_value = Accidental_valueSharp_sharp
		return
	case "flat-flat":
		*accidental_value = Accidental_valueFlat_flat
		return
	case "natural-sharp":
		*accidental_value = Accidental_valueNatural_sharp
		return
	case "natural-flat":
		*accidental_value = Accidental_valueNatural_flat
		return
	case "quarter-flat":
		*accidental_value = Accidental_valueQuarter_flat
		return
	case "quarter-sharp":
		*accidental_value = Accidental_valueQuarter_sharp
		return
	case "three-quarters-flat":
		*accidental_value = Accidental_valueThree_quarters_flat
		return
	case "three-quarters-sharp":
		*accidental_value = Accidental_valueThree_quarters_sharp
		return
	case "sharp-down":
		*accidental_value = Accidental_valueSharp_down
		return
	case "sharp-up":
		*accidental_value = Accidental_valueSharp_up
		return
	case "natural-down":
		*accidental_value = Accidental_valueNatural_down
		return
	case "natural-up":
		*accidental_value = Accidental_valueNatural_up
		return
	case "flat-down":
		*accidental_value = Accidental_valueFlat_down
		return
	case "flat-up":
		*accidental_value = Accidental_valueFlat_up
		return
	case "double-sharp-down":
		*accidental_value = Accidental_valueDouble_sharp_down
		return
	case "double-sharp-up":
		*accidental_value = Accidental_valueDouble_sharp_up
		return
	case "flat-flat-down":
		*accidental_value = Accidental_valueFlat_flat_down
		return
	case "flat-flat-up":
		*accidental_value = Accidental_valueFlat_flat_up
		return
	case "arrow-down":
		*accidental_value = Accidental_valueArrow_down
		return
	case "arrow-up":
		*accidental_value = Accidental_valueArrow_up
		return
	case "triple-sharp":
		*accidental_value = Accidental_valueTriple_sharp
		return
	case "triple-flat":
		*accidental_value = Accidental_valueTriple_flat
		return
	case "slash-quarter-sharp":
		*accidental_value = Accidental_valueSlash_quarter_sharp
		return
	case "slash-sharp":
		*accidental_value = Accidental_valueSlash_sharp
		return
	case "slash-flat":
		*accidental_value = Accidental_valueSlash_flat
		return
	case "double-slash-flat":
		*accidental_value = Accidental_valueDouble_slash_flat
		return
	case "sharp-1":
		*accidental_value = Accidental_valueSharp_1
		return
	case "sharp-2":
		*accidental_value = Accidental_valueSharp_2
		return
	case "sharp-3":
		*accidental_value = Accidental_valueSharp_3
		return
	case "sharp-5":
		*accidental_value = Accidental_valueSharp_5
		return
	case "flat-1":
		*accidental_value = Accidental_valueFlat_1
		return
	case "flat-2":
		*accidental_value = Accidental_valueFlat_2
		return
	case "flat-3":
		*accidental_value = Accidental_valueFlat_3
		return
	case "flat-4":
		*accidental_value = Accidental_valueFlat_4
		return
	case "sori":
		*accidental_value = Accidental_valueSori
		return
	case "koron":
		*accidental_value = Accidental_valueKoron
		return
	case "other":
		*accidental_value = Accidental_valueOther
		return
	default:
		return errUnkownEnum
	}
}

func (accidental_value *Accidental_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Accidental_valueSharp":
		*accidental_value = Accidental_valueSharp
	case "Accidental_valueNatural":
		*accidental_value = Accidental_valueNatural
	case "Accidental_valueFlat":
		*accidental_value = Accidental_valueFlat
	case "Accidental_valueDouble_sharp":
		*accidental_value = Accidental_valueDouble_sharp
	case "Accidental_valueSharp_sharp":
		*accidental_value = Accidental_valueSharp_sharp
	case "Accidental_valueFlat_flat":
		*accidental_value = Accidental_valueFlat_flat
	case "Accidental_valueNatural_sharp":
		*accidental_value = Accidental_valueNatural_sharp
	case "Accidental_valueNatural_flat":
		*accidental_value = Accidental_valueNatural_flat
	case "Accidental_valueQuarter_flat":
		*accidental_value = Accidental_valueQuarter_flat
	case "Accidental_valueQuarter_sharp":
		*accidental_value = Accidental_valueQuarter_sharp
	case "Accidental_valueThree_quarters_flat":
		*accidental_value = Accidental_valueThree_quarters_flat
	case "Accidental_valueThree_quarters_sharp":
		*accidental_value = Accidental_valueThree_quarters_sharp
	case "Accidental_valueSharp_down":
		*accidental_value = Accidental_valueSharp_down
	case "Accidental_valueSharp_up":
		*accidental_value = Accidental_valueSharp_up
	case "Accidental_valueNatural_down":
		*accidental_value = Accidental_valueNatural_down
	case "Accidental_valueNatural_up":
		*accidental_value = Accidental_valueNatural_up
	case "Accidental_valueFlat_down":
		*accidental_value = Accidental_valueFlat_down
	case "Accidental_valueFlat_up":
		*accidental_value = Accidental_valueFlat_up
	case "Accidental_valueDouble_sharp_down":
		*accidental_value = Accidental_valueDouble_sharp_down
	case "Accidental_valueDouble_sharp_up":
		*accidental_value = Accidental_valueDouble_sharp_up
	case "Accidental_valueFlat_flat_down":
		*accidental_value = Accidental_valueFlat_flat_down
	case "Accidental_valueFlat_flat_up":
		*accidental_value = Accidental_valueFlat_flat_up
	case "Accidental_valueArrow_down":
		*accidental_value = Accidental_valueArrow_down
	case "Accidental_valueArrow_up":
		*accidental_value = Accidental_valueArrow_up
	case "Accidental_valueTriple_sharp":
		*accidental_value = Accidental_valueTriple_sharp
	case "Accidental_valueTriple_flat":
		*accidental_value = Accidental_valueTriple_flat
	case "Accidental_valueSlash_quarter_sharp":
		*accidental_value = Accidental_valueSlash_quarter_sharp
	case "Accidental_valueSlash_sharp":
		*accidental_value = Accidental_valueSlash_sharp
	case "Accidental_valueSlash_flat":
		*accidental_value = Accidental_valueSlash_flat
	case "Accidental_valueDouble_slash_flat":
		*accidental_value = Accidental_valueDouble_slash_flat
	case "Accidental_valueSharp_1":
		*accidental_value = Accidental_valueSharp_1
	case "Accidental_valueSharp_2":
		*accidental_value = Accidental_valueSharp_2
	case "Accidental_valueSharp_3":
		*accidental_value = Accidental_valueSharp_3
	case "Accidental_valueSharp_5":
		*accidental_value = Accidental_valueSharp_5
	case "Accidental_valueFlat_1":
		*accidental_value = Accidental_valueFlat_1
	case "Accidental_valueFlat_2":
		*accidental_value = Accidental_valueFlat_2
	case "Accidental_valueFlat_3":
		*accidental_value = Accidental_valueFlat_3
	case "Accidental_valueFlat_4":
		*accidental_value = Accidental_valueFlat_4
	case "Accidental_valueSori":
		*accidental_value = Accidental_valueSori
	case "Accidental_valueKoron":
		*accidental_value = Accidental_valueKoron
	case "Accidental_valueOther":
		*accidental_value = Accidental_valueOther
	default:
		return errUnkownEnum
	}
	return
}

func (accidental_value *Accidental_value) ToCodeString() (res string) {

	switch *accidental_value {
	// insertion code per enum code
	case Accidental_valueSharp:
		res = "Accidental_valueSharp"
	case Accidental_valueNatural:
		res = "Accidental_valueNatural"
	case Accidental_valueFlat:
		res = "Accidental_valueFlat"
	case Accidental_valueDouble_sharp:
		res = "Accidental_valueDouble_sharp"
	case Accidental_valueSharp_sharp:
		res = "Accidental_valueSharp_sharp"
	case Accidental_valueFlat_flat:
		res = "Accidental_valueFlat_flat"
	case Accidental_valueNatural_sharp:
		res = "Accidental_valueNatural_sharp"
	case Accidental_valueNatural_flat:
		res = "Accidental_valueNatural_flat"
	case Accidental_valueQuarter_flat:
		res = "Accidental_valueQuarter_flat"
	case Accidental_valueQuarter_sharp:
		res = "Accidental_valueQuarter_sharp"
	case Accidental_valueThree_quarters_flat:
		res = "Accidental_valueThree_quarters_flat"
	case Accidental_valueThree_quarters_sharp:
		res = "Accidental_valueThree_quarters_sharp"
	case Accidental_valueSharp_down:
		res = "Accidental_valueSharp_down"
	case Accidental_valueSharp_up:
		res = "Accidental_valueSharp_up"
	case Accidental_valueNatural_down:
		res = "Accidental_valueNatural_down"
	case Accidental_valueNatural_up:
		res = "Accidental_valueNatural_up"
	case Accidental_valueFlat_down:
		res = "Accidental_valueFlat_down"
	case Accidental_valueFlat_up:
		res = "Accidental_valueFlat_up"
	case Accidental_valueDouble_sharp_down:
		res = "Accidental_valueDouble_sharp_down"
	case Accidental_valueDouble_sharp_up:
		res = "Accidental_valueDouble_sharp_up"
	case Accidental_valueFlat_flat_down:
		res = "Accidental_valueFlat_flat_down"
	case Accidental_valueFlat_flat_up:
		res = "Accidental_valueFlat_flat_up"
	case Accidental_valueArrow_down:
		res = "Accidental_valueArrow_down"
	case Accidental_valueArrow_up:
		res = "Accidental_valueArrow_up"
	case Accidental_valueTriple_sharp:
		res = "Accidental_valueTriple_sharp"
	case Accidental_valueTriple_flat:
		res = "Accidental_valueTriple_flat"
	case Accidental_valueSlash_quarter_sharp:
		res = "Accidental_valueSlash_quarter_sharp"
	case Accidental_valueSlash_sharp:
		res = "Accidental_valueSlash_sharp"
	case Accidental_valueSlash_flat:
		res = "Accidental_valueSlash_flat"
	case Accidental_valueDouble_slash_flat:
		res = "Accidental_valueDouble_slash_flat"
	case Accidental_valueSharp_1:
		res = "Accidental_valueSharp_1"
	case Accidental_valueSharp_2:
		res = "Accidental_valueSharp_2"
	case Accidental_valueSharp_3:
		res = "Accidental_valueSharp_3"
	case Accidental_valueSharp_5:
		res = "Accidental_valueSharp_5"
	case Accidental_valueFlat_1:
		res = "Accidental_valueFlat_1"
	case Accidental_valueFlat_2:
		res = "Accidental_valueFlat_2"
	case Accidental_valueFlat_3:
		res = "Accidental_valueFlat_3"
	case Accidental_valueFlat_4:
		res = "Accidental_valueFlat_4"
	case Accidental_valueSori:
		res = "Accidental_valueSori"
	case Accidental_valueKoron:
		res = "Accidental_valueKoron"
	case Accidental_valueOther:
		res = "Accidental_valueOther"
	}
	return
}

func (accidental_value Accidental_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Accidental_valueSharp")
	res = append(res, "Accidental_valueNatural")
	res = append(res, "Accidental_valueFlat")
	res = append(res, "Accidental_valueDouble_sharp")
	res = append(res, "Accidental_valueSharp_sharp")
	res = append(res, "Accidental_valueFlat_flat")
	res = append(res, "Accidental_valueNatural_sharp")
	res = append(res, "Accidental_valueNatural_flat")
	res = append(res, "Accidental_valueQuarter_flat")
	res = append(res, "Accidental_valueQuarter_sharp")
	res = append(res, "Accidental_valueThree_quarters_flat")
	res = append(res, "Accidental_valueThree_quarters_sharp")
	res = append(res, "Accidental_valueSharp_down")
	res = append(res, "Accidental_valueSharp_up")
	res = append(res, "Accidental_valueNatural_down")
	res = append(res, "Accidental_valueNatural_up")
	res = append(res, "Accidental_valueFlat_down")
	res = append(res, "Accidental_valueFlat_up")
	res = append(res, "Accidental_valueDouble_sharp_down")
	res = append(res, "Accidental_valueDouble_sharp_up")
	res = append(res, "Accidental_valueFlat_flat_down")
	res = append(res, "Accidental_valueFlat_flat_up")
	res = append(res, "Accidental_valueArrow_down")
	res = append(res, "Accidental_valueArrow_up")
	res = append(res, "Accidental_valueTriple_sharp")
	res = append(res, "Accidental_valueTriple_flat")
	res = append(res, "Accidental_valueSlash_quarter_sharp")
	res = append(res, "Accidental_valueSlash_sharp")
	res = append(res, "Accidental_valueSlash_flat")
	res = append(res, "Accidental_valueDouble_slash_flat")
	res = append(res, "Accidental_valueSharp_1")
	res = append(res, "Accidental_valueSharp_2")
	res = append(res, "Accidental_valueSharp_3")
	res = append(res, "Accidental_valueSharp_5")
	res = append(res, "Accidental_valueFlat_1")
	res = append(res, "Accidental_valueFlat_2")
	res = append(res, "Accidental_valueFlat_3")
	res = append(res, "Accidental_valueFlat_4")
	res = append(res, "Accidental_valueSori")
	res = append(res, "Accidental_valueKoron")
	res = append(res, "Accidental_valueOther")

	return
}

func (accidental_value Accidental_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "sharp")
	res = append(res, "natural")
	res = append(res, "flat")
	res = append(res, "double-sharp")
	res = append(res, "sharp-sharp")
	res = append(res, "flat-flat")
	res = append(res, "natural-sharp")
	res = append(res, "natural-flat")
	res = append(res, "quarter-flat")
	res = append(res, "quarter-sharp")
	res = append(res, "three-quarters-flat")
	res = append(res, "three-quarters-sharp")
	res = append(res, "sharp-down")
	res = append(res, "sharp-up")
	res = append(res, "natural-down")
	res = append(res, "natural-up")
	res = append(res, "flat-down")
	res = append(res, "flat-up")
	res = append(res, "double-sharp-down")
	res = append(res, "double-sharp-up")
	res = append(res, "flat-flat-down")
	res = append(res, "flat-flat-up")
	res = append(res, "arrow-down")
	res = append(res, "arrow-up")
	res = append(res, "triple-sharp")
	res = append(res, "triple-flat")
	res = append(res, "slash-quarter-sharp")
	res = append(res, "slash-sharp")
	res = append(res, "slash-flat")
	res = append(res, "double-slash-flat")
	res = append(res, "sharp-1")
	res = append(res, "sharp-2")
	res = append(res, "sharp-3")
	res = append(res, "sharp-5")
	res = append(res, "flat-1")
	res = append(res, "flat-2")
	res = append(res, "flat-3")
	res = append(res, "flat-4")
	res = append(res, "sori")
	res = append(res, "koron")
	res = append(res, "other")

	return
}

// Utility function for AnyURI
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (anyuri AnyURI) ToString() (res string) {

	// migration of former implementation of enum
	switch anyuri {
	// insertion code per enum code
	}
	return
}

func (anyuri *AnyURI) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (anyuri *AnyURI) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (anyuri *AnyURI) ToCodeString() (res string) {

	switch *anyuri {
	// insertion code per enum code
	}
	return
}

func (anyuri AnyURI) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (anyuri AnyURI) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Arrow_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (arrow_direction Arrow_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch arrow_direction {
	// insertion code per enum code
	case Arrow_directionLeft:
		res = "left"
	case Arrow_directionUp:
		res = "up"
	case Arrow_directionRight:
		res = "right"
	case Arrow_directionDown:
		res = "down"
	case Arrow_directionNorthwest:
		res = "northwest"
	case Arrow_directionNortheast:
		res = "northeast"
	case Arrow_directionSoutheast:
		res = "southeast"
	case Arrow_directionSouthwest:
		res = "southwest"
	case Arrow_directionLeftright:
		res = "left right"
	case Arrow_directionUpdown:
		res = "up down"
	case Arrow_directionNorthwestsoutheast:
		res = "northwest southeast"
	case Arrow_directionNortheastsouthwest:
		res = "northeast southwest"
	case Arrow_directionOther:
		res = "other"
	}
	return
}

func (arrow_direction *Arrow_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*arrow_direction = Arrow_directionLeft
		return
	case "up":
		*arrow_direction = Arrow_directionUp
		return
	case "right":
		*arrow_direction = Arrow_directionRight
		return
	case "down":
		*arrow_direction = Arrow_directionDown
		return
	case "northwest":
		*arrow_direction = Arrow_directionNorthwest
		return
	case "northeast":
		*arrow_direction = Arrow_directionNortheast
		return
	case "southeast":
		*arrow_direction = Arrow_directionSoutheast
		return
	case "southwest":
		*arrow_direction = Arrow_directionSouthwest
		return
	case "left right":
		*arrow_direction = Arrow_directionLeftright
		return
	case "up down":
		*arrow_direction = Arrow_directionUpdown
		return
	case "northwest southeast":
		*arrow_direction = Arrow_directionNorthwestsoutheast
		return
	case "northeast southwest":
		*arrow_direction = Arrow_directionNortheastsouthwest
		return
	case "other":
		*arrow_direction = Arrow_directionOther
		return
	default:
		return errUnkownEnum
	}
}

func (arrow_direction *Arrow_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Arrow_directionLeft":
		*arrow_direction = Arrow_directionLeft
	case "Arrow_directionUp":
		*arrow_direction = Arrow_directionUp
	case "Arrow_directionRight":
		*arrow_direction = Arrow_directionRight
	case "Arrow_directionDown":
		*arrow_direction = Arrow_directionDown
	case "Arrow_directionNorthwest":
		*arrow_direction = Arrow_directionNorthwest
	case "Arrow_directionNortheast":
		*arrow_direction = Arrow_directionNortheast
	case "Arrow_directionSoutheast":
		*arrow_direction = Arrow_directionSoutheast
	case "Arrow_directionSouthwest":
		*arrow_direction = Arrow_directionSouthwest
	case "Arrow_directionLeftright":
		*arrow_direction = Arrow_directionLeftright
	case "Arrow_directionUpdown":
		*arrow_direction = Arrow_directionUpdown
	case "Arrow_directionNorthwestsoutheast":
		*arrow_direction = Arrow_directionNorthwestsoutheast
	case "Arrow_directionNortheastsouthwest":
		*arrow_direction = Arrow_directionNortheastsouthwest
	case "Arrow_directionOther":
		*arrow_direction = Arrow_directionOther
	default:
		return errUnkownEnum
	}
	return
}

func (arrow_direction *Arrow_direction) ToCodeString() (res string) {

	switch *arrow_direction {
	// insertion code per enum code
	case Arrow_directionLeft:
		res = "Arrow_directionLeft"
	case Arrow_directionUp:
		res = "Arrow_directionUp"
	case Arrow_directionRight:
		res = "Arrow_directionRight"
	case Arrow_directionDown:
		res = "Arrow_directionDown"
	case Arrow_directionNorthwest:
		res = "Arrow_directionNorthwest"
	case Arrow_directionNortheast:
		res = "Arrow_directionNortheast"
	case Arrow_directionSoutheast:
		res = "Arrow_directionSoutheast"
	case Arrow_directionSouthwest:
		res = "Arrow_directionSouthwest"
	case Arrow_directionLeftright:
		res = "Arrow_directionLeftright"
	case Arrow_directionUpdown:
		res = "Arrow_directionUpdown"
	case Arrow_directionNorthwestsoutheast:
		res = "Arrow_directionNorthwestsoutheast"
	case Arrow_directionNortheastsouthwest:
		res = "Arrow_directionNortheastsouthwest"
	case Arrow_directionOther:
		res = "Arrow_directionOther"
	}
	return
}

func (arrow_direction Arrow_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Arrow_directionLeft")
	res = append(res, "Arrow_directionUp")
	res = append(res, "Arrow_directionRight")
	res = append(res, "Arrow_directionDown")
	res = append(res, "Arrow_directionNorthwest")
	res = append(res, "Arrow_directionNortheast")
	res = append(res, "Arrow_directionSoutheast")
	res = append(res, "Arrow_directionSouthwest")
	res = append(res, "Arrow_directionLeftright")
	res = append(res, "Arrow_directionUpdown")
	res = append(res, "Arrow_directionNorthwestsoutheast")
	res = append(res, "Arrow_directionNortheastsouthwest")
	res = append(res, "Arrow_directionOther")

	return
}

func (arrow_direction Arrow_direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "up")
	res = append(res, "right")
	res = append(res, "down")
	res = append(res, "northwest")
	res = append(res, "northeast")
	res = append(res, "southeast")
	res = append(res, "southwest")
	res = append(res, "left right")
	res = append(res, "up down")
	res = append(res, "northwest southeast")
	res = append(res, "northeast southwest")
	res = append(res, "other")

	return
}

// Utility function for Arrow_style
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (arrow_style Arrow_style) ToString() (res string) {

	// migration of former implementation of enum
	switch arrow_style {
	// insertion code per enum code
	case Arrow_styleSingle:
		res = "single"
	case Arrow_styleDouble:
		res = "double"
	case Arrow_styleFilled:
		res = "filled"
	case Arrow_styleHollow:
		res = "hollow"
	case Arrow_stylePaired:
		res = "paired"
	case Arrow_styleCombined:
		res = "combined"
	case Arrow_styleOther:
		res = "other"
	}
	return
}

func (arrow_style *Arrow_style) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "single":
		*arrow_style = Arrow_styleSingle
		return
	case "double":
		*arrow_style = Arrow_styleDouble
		return
	case "filled":
		*arrow_style = Arrow_styleFilled
		return
	case "hollow":
		*arrow_style = Arrow_styleHollow
		return
	case "paired":
		*arrow_style = Arrow_stylePaired
		return
	case "combined":
		*arrow_style = Arrow_styleCombined
		return
	case "other":
		*arrow_style = Arrow_styleOther
		return
	default:
		return errUnkownEnum
	}
}

func (arrow_style *Arrow_style) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Arrow_styleSingle":
		*arrow_style = Arrow_styleSingle
	case "Arrow_styleDouble":
		*arrow_style = Arrow_styleDouble
	case "Arrow_styleFilled":
		*arrow_style = Arrow_styleFilled
	case "Arrow_styleHollow":
		*arrow_style = Arrow_styleHollow
	case "Arrow_stylePaired":
		*arrow_style = Arrow_stylePaired
	case "Arrow_styleCombined":
		*arrow_style = Arrow_styleCombined
	case "Arrow_styleOther":
		*arrow_style = Arrow_styleOther
	default:
		return errUnkownEnum
	}
	return
}

func (arrow_style *Arrow_style) ToCodeString() (res string) {

	switch *arrow_style {
	// insertion code per enum code
	case Arrow_styleSingle:
		res = "Arrow_styleSingle"
	case Arrow_styleDouble:
		res = "Arrow_styleDouble"
	case Arrow_styleFilled:
		res = "Arrow_styleFilled"
	case Arrow_styleHollow:
		res = "Arrow_styleHollow"
	case Arrow_stylePaired:
		res = "Arrow_stylePaired"
	case Arrow_styleCombined:
		res = "Arrow_styleCombined"
	case Arrow_styleOther:
		res = "Arrow_styleOther"
	}
	return
}

func (arrow_style Arrow_style) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Arrow_styleSingle")
	res = append(res, "Arrow_styleDouble")
	res = append(res, "Arrow_styleFilled")
	res = append(res, "Arrow_styleHollow")
	res = append(res, "Arrow_stylePaired")
	res = append(res, "Arrow_styleCombined")
	res = append(res, "Arrow_styleOther")

	return
}

func (arrow_style Arrow_style) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "single")
	res = append(res, "double")
	res = append(res, "filled")
	res = append(res, "hollow")
	res = append(res, "paired")
	res = append(res, "combined")
	res = append(res, "other")

	return
}

// Utility function for Backward_forward
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (backward_forward Backward_forward) ToString() (res string) {

	// migration of former implementation of enum
	switch backward_forward {
	// insertion code per enum code
	case Backward_forwardBackward:
		res = "backward"
	case Backward_forwardForward:
		res = "forward"
	}
	return
}

func (backward_forward *Backward_forward) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "backward":
		*backward_forward = Backward_forwardBackward
		return
	case "forward":
		*backward_forward = Backward_forwardForward
		return
	default:
		return errUnkownEnum
	}
}

func (backward_forward *Backward_forward) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Backward_forwardBackward":
		*backward_forward = Backward_forwardBackward
	case "Backward_forwardForward":
		*backward_forward = Backward_forwardForward
	default:
		return errUnkownEnum
	}
	return
}

func (backward_forward *Backward_forward) ToCodeString() (res string) {

	switch *backward_forward {
	// insertion code per enum code
	case Backward_forwardBackward:
		res = "Backward_forwardBackward"
	case Backward_forwardForward:
		res = "Backward_forwardForward"
	}
	return
}

func (backward_forward Backward_forward) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Backward_forwardBackward")
	res = append(res, "Backward_forwardForward")

	return
}

func (backward_forward Backward_forward) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "backward")
	res = append(res, "forward")

	return
}

// Utility function for Bar_style
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (bar_style Bar_style) ToString() (res string) {

	// migration of former implementation of enum
	switch bar_style {
	// insertion code per enum code
	case Bar_styleRegular:
		res = "regular"
	case Bar_styleDotted:
		res = "dotted"
	case Bar_styleDashed:
		res = "dashed"
	case Bar_styleHeavy:
		res = "heavy"
	case Bar_styleLight_light:
		res = "light-light"
	case Bar_styleLight_heavy:
		res = "light-heavy"
	case Bar_styleHeavy_light:
		res = "heavy-light"
	case Bar_styleHeavy_heavy:
		res = "heavy-heavy"
	case Bar_styleTick:
		res = "tick"
	case Bar_styleShort:
		res = "short"
	case Bar_styleNone:
		res = "none"
	}
	return
}

func (bar_style *Bar_style) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "regular":
		*bar_style = Bar_styleRegular
		return
	case "dotted":
		*bar_style = Bar_styleDotted
		return
	case "dashed":
		*bar_style = Bar_styleDashed
		return
	case "heavy":
		*bar_style = Bar_styleHeavy
		return
	case "light-light":
		*bar_style = Bar_styleLight_light
		return
	case "light-heavy":
		*bar_style = Bar_styleLight_heavy
		return
	case "heavy-light":
		*bar_style = Bar_styleHeavy_light
		return
	case "heavy-heavy":
		*bar_style = Bar_styleHeavy_heavy
		return
	case "tick":
		*bar_style = Bar_styleTick
		return
	case "short":
		*bar_style = Bar_styleShort
		return
	case "none":
		*bar_style = Bar_styleNone
		return
	default:
		return errUnkownEnum
	}
}

func (bar_style *Bar_style) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Bar_styleRegular":
		*bar_style = Bar_styleRegular
	case "Bar_styleDotted":
		*bar_style = Bar_styleDotted
	case "Bar_styleDashed":
		*bar_style = Bar_styleDashed
	case "Bar_styleHeavy":
		*bar_style = Bar_styleHeavy
	case "Bar_styleLight_light":
		*bar_style = Bar_styleLight_light
	case "Bar_styleLight_heavy":
		*bar_style = Bar_styleLight_heavy
	case "Bar_styleHeavy_light":
		*bar_style = Bar_styleHeavy_light
	case "Bar_styleHeavy_heavy":
		*bar_style = Bar_styleHeavy_heavy
	case "Bar_styleTick":
		*bar_style = Bar_styleTick
	case "Bar_styleShort":
		*bar_style = Bar_styleShort
	case "Bar_styleNone":
		*bar_style = Bar_styleNone
	default:
		return errUnkownEnum
	}
	return
}

func (bar_style *Bar_style) ToCodeString() (res string) {

	switch *bar_style {
	// insertion code per enum code
	case Bar_styleRegular:
		res = "Bar_styleRegular"
	case Bar_styleDotted:
		res = "Bar_styleDotted"
	case Bar_styleDashed:
		res = "Bar_styleDashed"
	case Bar_styleHeavy:
		res = "Bar_styleHeavy"
	case Bar_styleLight_light:
		res = "Bar_styleLight_light"
	case Bar_styleLight_heavy:
		res = "Bar_styleLight_heavy"
	case Bar_styleHeavy_light:
		res = "Bar_styleHeavy_light"
	case Bar_styleHeavy_heavy:
		res = "Bar_styleHeavy_heavy"
	case Bar_styleTick:
		res = "Bar_styleTick"
	case Bar_styleShort:
		res = "Bar_styleShort"
	case Bar_styleNone:
		res = "Bar_styleNone"
	}
	return
}

func (bar_style Bar_style) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Bar_styleRegular")
	res = append(res, "Bar_styleDotted")
	res = append(res, "Bar_styleDashed")
	res = append(res, "Bar_styleHeavy")
	res = append(res, "Bar_styleLight_light")
	res = append(res, "Bar_styleLight_heavy")
	res = append(res, "Bar_styleHeavy_light")
	res = append(res, "Bar_styleHeavy_heavy")
	res = append(res, "Bar_styleTick")
	res = append(res, "Bar_styleShort")
	res = append(res, "Bar_styleNone")

	return
}

func (bar_style Bar_style) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "regular")
	res = append(res, "dotted")
	res = append(res, "dashed")
	res = append(res, "heavy")
	res = append(res, "light-light")
	res = append(res, "light-heavy")
	res = append(res, "heavy-light")
	res = append(res, "heavy-heavy")
	res = append(res, "tick")
	res = append(res, "short")
	res = append(res, "none")

	return
}

// Utility function for Beam_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (beam_value Beam_value) ToString() (res string) {

	// migration of former implementation of enum
	switch beam_value {
	// insertion code per enum code
	case Beam_valueBegin:
		res = "begin"
	case Beam_valueContinue_:
		res = "continue"
	case Beam_valueEnd:
		res = "end"
	case Beam_valueForwardhook:
		res = "forward hook"
	case Beam_valueBackwardhook:
		res = "backward hook"
	}
	return
}

func (beam_value *Beam_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "begin":
		*beam_value = Beam_valueBegin
		return
	case "continue":
		*beam_value = Beam_valueContinue_
		return
	case "end":
		*beam_value = Beam_valueEnd
		return
	case "forward hook":
		*beam_value = Beam_valueForwardhook
		return
	case "backward hook":
		*beam_value = Beam_valueBackwardhook
		return
	default:
		return errUnkownEnum
	}
}

func (beam_value *Beam_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Beam_valueBegin":
		*beam_value = Beam_valueBegin
	case "Beam_valueContinue_":
		*beam_value = Beam_valueContinue_
	case "Beam_valueEnd":
		*beam_value = Beam_valueEnd
	case "Beam_valueForwardhook":
		*beam_value = Beam_valueForwardhook
	case "Beam_valueBackwardhook":
		*beam_value = Beam_valueBackwardhook
	default:
		return errUnkownEnum
	}
	return
}

func (beam_value *Beam_value) ToCodeString() (res string) {

	switch *beam_value {
	// insertion code per enum code
	case Beam_valueBegin:
		res = "Beam_valueBegin"
	case Beam_valueContinue_:
		res = "Beam_valueContinue_"
	case Beam_valueEnd:
		res = "Beam_valueEnd"
	case Beam_valueForwardhook:
		res = "Beam_valueForwardhook"
	case Beam_valueBackwardhook:
		res = "Beam_valueBackwardhook"
	}
	return
}

func (beam_value Beam_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Beam_valueBegin")
	res = append(res, "Beam_valueContinue_")
	res = append(res, "Beam_valueEnd")
	res = append(res, "Beam_valueForwardhook")
	res = append(res, "Beam_valueBackwardhook")

	return
}

func (beam_value Beam_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "begin")
	res = append(res, "continue")
	res = append(res, "end")
	res = append(res, "forward hook")
	res = append(res, "backward hook")

	return
}

// Utility function for Beater_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (beater_value Beater_value) ToString() (res string) {

	// migration of former implementation of enum
	switch beater_value {
	// insertion code per enum code
	case Beater_valueBow:
		res = "bow"
	case Beater_valueChimehammer:
		res = "chime hammer"
	case Beater_valueCoin:
		res = "coin"
	case Beater_valueDrumstick:
		res = "drum stick"
	case Beater_valueFinger:
		res = "finger"
	case Beater_valueFingernail:
		res = "fingernail"
	case Beater_valueFist:
		res = "fist"
	case Beater_valueGuiroscraper:
		res = "guiro scraper"
	case Beater_valueHammer:
		res = "hammer"
	case Beater_valueHand:
		res = "hand"
	case Beater_valueJazzstick:
		res = "jazz stick"
	case Beater_valueKnittingneedle:
		res = "knitting needle"
	case Beater_valueMetalhammer:
		res = "metal hammer"
	case Beater_valueSlidebrushongong:
		res = "slide brush on gong"
	case Beater_valueSnarestick:
		res = "snare stick"
	case Beater_valueSpoonmallet:
		res = "spoon mallet"
	case Beater_valueSuperball:
		res = "superball"
	case Beater_valueTrianglebeater:
		res = "triangle beater"
	case Beater_valueTrianglebeaterplain:
		res = "triangle beater plain"
	case Beater_valueWirebrush:
		res = "wire brush"
	}
	return
}

func (beater_value *Beater_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bow":
		*beater_value = Beater_valueBow
		return
	case "chime hammer":
		*beater_value = Beater_valueChimehammer
		return
	case "coin":
		*beater_value = Beater_valueCoin
		return
	case "drum stick":
		*beater_value = Beater_valueDrumstick
		return
	case "finger":
		*beater_value = Beater_valueFinger
		return
	case "fingernail":
		*beater_value = Beater_valueFingernail
		return
	case "fist":
		*beater_value = Beater_valueFist
		return
	case "guiro scraper":
		*beater_value = Beater_valueGuiroscraper
		return
	case "hammer":
		*beater_value = Beater_valueHammer
		return
	case "hand":
		*beater_value = Beater_valueHand
		return
	case "jazz stick":
		*beater_value = Beater_valueJazzstick
		return
	case "knitting needle":
		*beater_value = Beater_valueKnittingneedle
		return
	case "metal hammer":
		*beater_value = Beater_valueMetalhammer
		return
	case "slide brush on gong":
		*beater_value = Beater_valueSlidebrushongong
		return
	case "snare stick":
		*beater_value = Beater_valueSnarestick
		return
	case "spoon mallet":
		*beater_value = Beater_valueSpoonmallet
		return
	case "superball":
		*beater_value = Beater_valueSuperball
		return
	case "triangle beater":
		*beater_value = Beater_valueTrianglebeater
		return
	case "triangle beater plain":
		*beater_value = Beater_valueTrianglebeaterplain
		return
	case "wire brush":
		*beater_value = Beater_valueWirebrush
		return
	default:
		return errUnkownEnum
	}
}

func (beater_value *Beater_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Beater_valueBow":
		*beater_value = Beater_valueBow
	case "Beater_valueChimehammer":
		*beater_value = Beater_valueChimehammer
	case "Beater_valueCoin":
		*beater_value = Beater_valueCoin
	case "Beater_valueDrumstick":
		*beater_value = Beater_valueDrumstick
	case "Beater_valueFinger":
		*beater_value = Beater_valueFinger
	case "Beater_valueFingernail":
		*beater_value = Beater_valueFingernail
	case "Beater_valueFist":
		*beater_value = Beater_valueFist
	case "Beater_valueGuiroscraper":
		*beater_value = Beater_valueGuiroscraper
	case "Beater_valueHammer":
		*beater_value = Beater_valueHammer
	case "Beater_valueHand":
		*beater_value = Beater_valueHand
	case "Beater_valueJazzstick":
		*beater_value = Beater_valueJazzstick
	case "Beater_valueKnittingneedle":
		*beater_value = Beater_valueKnittingneedle
	case "Beater_valueMetalhammer":
		*beater_value = Beater_valueMetalhammer
	case "Beater_valueSlidebrushongong":
		*beater_value = Beater_valueSlidebrushongong
	case "Beater_valueSnarestick":
		*beater_value = Beater_valueSnarestick
	case "Beater_valueSpoonmallet":
		*beater_value = Beater_valueSpoonmallet
	case "Beater_valueSuperball":
		*beater_value = Beater_valueSuperball
	case "Beater_valueTrianglebeater":
		*beater_value = Beater_valueTrianglebeater
	case "Beater_valueTrianglebeaterplain":
		*beater_value = Beater_valueTrianglebeaterplain
	case "Beater_valueWirebrush":
		*beater_value = Beater_valueWirebrush
	default:
		return errUnkownEnum
	}
	return
}

func (beater_value *Beater_value) ToCodeString() (res string) {

	switch *beater_value {
	// insertion code per enum code
	case Beater_valueBow:
		res = "Beater_valueBow"
	case Beater_valueChimehammer:
		res = "Beater_valueChimehammer"
	case Beater_valueCoin:
		res = "Beater_valueCoin"
	case Beater_valueDrumstick:
		res = "Beater_valueDrumstick"
	case Beater_valueFinger:
		res = "Beater_valueFinger"
	case Beater_valueFingernail:
		res = "Beater_valueFingernail"
	case Beater_valueFist:
		res = "Beater_valueFist"
	case Beater_valueGuiroscraper:
		res = "Beater_valueGuiroscraper"
	case Beater_valueHammer:
		res = "Beater_valueHammer"
	case Beater_valueHand:
		res = "Beater_valueHand"
	case Beater_valueJazzstick:
		res = "Beater_valueJazzstick"
	case Beater_valueKnittingneedle:
		res = "Beater_valueKnittingneedle"
	case Beater_valueMetalhammer:
		res = "Beater_valueMetalhammer"
	case Beater_valueSlidebrushongong:
		res = "Beater_valueSlidebrushongong"
	case Beater_valueSnarestick:
		res = "Beater_valueSnarestick"
	case Beater_valueSpoonmallet:
		res = "Beater_valueSpoonmallet"
	case Beater_valueSuperball:
		res = "Beater_valueSuperball"
	case Beater_valueTrianglebeater:
		res = "Beater_valueTrianglebeater"
	case Beater_valueTrianglebeaterplain:
		res = "Beater_valueTrianglebeaterplain"
	case Beater_valueWirebrush:
		res = "Beater_valueWirebrush"
	}
	return
}

func (beater_value Beater_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Beater_valueBow")
	res = append(res, "Beater_valueChimehammer")
	res = append(res, "Beater_valueCoin")
	res = append(res, "Beater_valueDrumstick")
	res = append(res, "Beater_valueFinger")
	res = append(res, "Beater_valueFingernail")
	res = append(res, "Beater_valueFist")
	res = append(res, "Beater_valueGuiroscraper")
	res = append(res, "Beater_valueHammer")
	res = append(res, "Beater_valueHand")
	res = append(res, "Beater_valueJazzstick")
	res = append(res, "Beater_valueKnittingneedle")
	res = append(res, "Beater_valueMetalhammer")
	res = append(res, "Beater_valueSlidebrushongong")
	res = append(res, "Beater_valueSnarestick")
	res = append(res, "Beater_valueSpoonmallet")
	res = append(res, "Beater_valueSuperball")
	res = append(res, "Beater_valueTrianglebeater")
	res = append(res, "Beater_valueTrianglebeaterplain")
	res = append(res, "Beater_valueWirebrush")

	return
}

func (beater_value Beater_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bow")
	res = append(res, "chime hammer")
	res = append(res, "coin")
	res = append(res, "drum stick")
	res = append(res, "finger")
	res = append(res, "fingernail")
	res = append(res, "fist")
	res = append(res, "guiro scraper")
	res = append(res, "hammer")
	res = append(res, "hand")
	res = append(res, "jazz stick")
	res = append(res, "knitting needle")
	res = append(res, "metal hammer")
	res = append(res, "slide brush on gong")
	res = append(res, "snare stick")
	res = append(res, "spoon mallet")
	res = append(res, "superball")
	res = append(res, "triangle beater")
	res = append(res, "triangle beater plain")
	res = append(res, "wire brush")

	return
}

// Utility function for Bend_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (bend_shape Bend_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch bend_shape {
	// insertion code per enum code
	case Bend_shapeAngled:
		res = "angled"
	case Bend_shapeCurved:
		res = "curved"
	}
	return
}

func (bend_shape *Bend_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "angled":
		*bend_shape = Bend_shapeAngled
		return
	case "curved":
		*bend_shape = Bend_shapeCurved
		return
	default:
		return errUnkownEnum
	}
}

func (bend_shape *Bend_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Bend_shapeAngled":
		*bend_shape = Bend_shapeAngled
	case "Bend_shapeCurved":
		*bend_shape = Bend_shapeCurved
	default:
		return errUnkownEnum
	}
	return
}

func (bend_shape *Bend_shape) ToCodeString() (res string) {

	switch *bend_shape {
	// insertion code per enum code
	case Bend_shapeAngled:
		res = "Bend_shapeAngled"
	case Bend_shapeCurved:
		res = "Bend_shapeCurved"
	}
	return
}

func (bend_shape Bend_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Bend_shapeAngled")
	res = append(res, "Bend_shapeCurved")

	return
}

func (bend_shape Bend_shape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "angled")
	res = append(res, "curved")

	return
}

// Utility function for Breath_mark_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (breath_mark_value Breath_mark_value) ToString() (res string) {

	// migration of former implementation of enum
	switch breath_mark_value {
	// insertion code per enum code
	case Breath_mark_valueEmptyString:
		res = ""
	case Breath_mark_valueComma:
		res = "comma"
	case Breath_mark_valueTick:
		res = "tick"
	case Breath_mark_valueUpbow:
		res = "upbow"
	case Breath_mark_valueSalzedo:
		res = "salzedo"
	}
	return
}

func (breath_mark_value *Breath_mark_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "":
		*breath_mark_value = Breath_mark_valueEmptyString
		return
	case "comma":
		*breath_mark_value = Breath_mark_valueComma
		return
	case "tick":
		*breath_mark_value = Breath_mark_valueTick
		return
	case "upbow":
		*breath_mark_value = Breath_mark_valueUpbow
		return
	case "salzedo":
		*breath_mark_value = Breath_mark_valueSalzedo
		return
	default:
		return errUnkownEnum
	}
}

func (breath_mark_value *Breath_mark_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Breath_mark_valueEmptyString":
		*breath_mark_value = Breath_mark_valueEmptyString
	case "Breath_mark_valueComma":
		*breath_mark_value = Breath_mark_valueComma
	case "Breath_mark_valueTick":
		*breath_mark_value = Breath_mark_valueTick
	case "Breath_mark_valueUpbow":
		*breath_mark_value = Breath_mark_valueUpbow
	case "Breath_mark_valueSalzedo":
		*breath_mark_value = Breath_mark_valueSalzedo
	default:
		return errUnkownEnum
	}
	return
}

func (breath_mark_value *Breath_mark_value) ToCodeString() (res string) {

	switch *breath_mark_value {
	// insertion code per enum code
	case Breath_mark_valueEmptyString:
		res = "Breath_mark_valueEmptyString"
	case Breath_mark_valueComma:
		res = "Breath_mark_valueComma"
	case Breath_mark_valueTick:
		res = "Breath_mark_valueTick"
	case Breath_mark_valueUpbow:
		res = "Breath_mark_valueUpbow"
	case Breath_mark_valueSalzedo:
		res = "Breath_mark_valueSalzedo"
	}
	return
}

func (breath_mark_value Breath_mark_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Breath_mark_valueEmptyString")
	res = append(res, "Breath_mark_valueComma")
	res = append(res, "Breath_mark_valueTick")
	res = append(res, "Breath_mark_valueUpbow")
	res = append(res, "Breath_mark_valueSalzedo")

	return
}

func (breath_mark_value Breath_mark_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "")
	res = append(res, "comma")
	res = append(res, "tick")
	res = append(res, "upbow")
	res = append(res, "salzedo")

	return
}

// Utility function for Caesura_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (caesura_value Caesura_value) ToString() (res string) {

	// migration of former implementation of enum
	switch caesura_value {
	// insertion code per enum code
	case Caesura_valueNormal:
		res = "normal"
	case Caesura_valueThick:
		res = "thick"
	case Caesura_valueShort:
		res = "short"
	case Caesura_valueCurved:
		res = "curved"
	case Caesura_valueSingle:
		res = "single"
	case Caesura_valueEmptyString:
		res = ""
	}
	return
}

func (caesura_value *Caesura_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*caesura_value = Caesura_valueNormal
		return
	case "thick":
		*caesura_value = Caesura_valueThick
		return
	case "short":
		*caesura_value = Caesura_valueShort
		return
	case "curved":
		*caesura_value = Caesura_valueCurved
		return
	case "single":
		*caesura_value = Caesura_valueSingle
		return
	case "":
		*caesura_value = Caesura_valueEmptyString
		return
	default:
		return errUnkownEnum
	}
}

func (caesura_value *Caesura_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Caesura_valueNormal":
		*caesura_value = Caesura_valueNormal
	case "Caesura_valueThick":
		*caesura_value = Caesura_valueThick
	case "Caesura_valueShort":
		*caesura_value = Caesura_valueShort
	case "Caesura_valueCurved":
		*caesura_value = Caesura_valueCurved
	case "Caesura_valueSingle":
		*caesura_value = Caesura_valueSingle
	case "Caesura_valueEmptyString":
		*caesura_value = Caesura_valueEmptyString
	default:
		return errUnkownEnum
	}
	return
}

func (caesura_value *Caesura_value) ToCodeString() (res string) {

	switch *caesura_value {
	// insertion code per enum code
	case Caesura_valueNormal:
		res = "Caesura_valueNormal"
	case Caesura_valueThick:
		res = "Caesura_valueThick"
	case Caesura_valueShort:
		res = "Caesura_valueShort"
	case Caesura_valueCurved:
		res = "Caesura_valueCurved"
	case Caesura_valueSingle:
		res = "Caesura_valueSingle"
	case Caesura_valueEmptyString:
		res = "Caesura_valueEmptyString"
	}
	return
}

func (caesura_value Caesura_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Caesura_valueNormal")
	res = append(res, "Caesura_valueThick")
	res = append(res, "Caesura_valueShort")
	res = append(res, "Caesura_valueCurved")
	res = append(res, "Caesura_valueSingle")
	res = append(res, "Caesura_valueEmptyString")

	return
}

func (caesura_value Caesura_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "thick")
	res = append(res, "short")
	res = append(res, "curved")
	res = append(res, "single")
	res = append(res, "")

	return
}

// Utility function for Cancel_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (cancel_location Cancel_location) ToString() (res string) {

	// migration of former implementation of enum
	switch cancel_location {
	// insertion code per enum code
	case Cancel_locationLeft:
		res = "left"
	case Cancel_locationRight:
		res = "right"
	case Cancel_locationBefore_barline:
		res = "before-barline"
	}
	return
}

func (cancel_location *Cancel_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*cancel_location = Cancel_locationLeft
		return
	case "right":
		*cancel_location = Cancel_locationRight
		return
	case "before-barline":
		*cancel_location = Cancel_locationBefore_barline
		return
	default:
		return errUnkownEnum
	}
}

func (cancel_location *Cancel_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Cancel_locationLeft":
		*cancel_location = Cancel_locationLeft
	case "Cancel_locationRight":
		*cancel_location = Cancel_locationRight
	case "Cancel_locationBefore_barline":
		*cancel_location = Cancel_locationBefore_barline
	default:
		return errUnkownEnum
	}
	return
}

func (cancel_location *Cancel_location) ToCodeString() (res string) {

	switch *cancel_location {
	// insertion code per enum code
	case Cancel_locationLeft:
		res = "Cancel_locationLeft"
	case Cancel_locationRight:
		res = "Cancel_locationRight"
	case Cancel_locationBefore_barline:
		res = "Cancel_locationBefore_barline"
	}
	return
}

func (cancel_location Cancel_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Cancel_locationLeft")
	res = append(res, "Cancel_locationRight")
	res = append(res, "Cancel_locationBefore_barline")

	return
}

func (cancel_location Cancel_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")
	res = append(res, "before-barline")

	return
}

// Utility function for Circular_arrow
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (circular_arrow Circular_arrow) ToString() (res string) {

	// migration of former implementation of enum
	switch circular_arrow {
	// insertion code per enum code
	case Circular_arrowClockwise:
		res = "clockwise"
	case Circular_arrowAnticlockwise:
		res = "anticlockwise"
	}
	return
}

func (circular_arrow *Circular_arrow) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "clockwise":
		*circular_arrow = Circular_arrowClockwise
		return
	case "anticlockwise":
		*circular_arrow = Circular_arrowAnticlockwise
		return
	default:
		return errUnkownEnum
	}
}

func (circular_arrow *Circular_arrow) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Circular_arrowClockwise":
		*circular_arrow = Circular_arrowClockwise
	case "Circular_arrowAnticlockwise":
		*circular_arrow = Circular_arrowAnticlockwise
	default:
		return errUnkownEnum
	}
	return
}

func (circular_arrow *Circular_arrow) ToCodeString() (res string) {

	switch *circular_arrow {
	// insertion code per enum code
	case Circular_arrowClockwise:
		res = "Circular_arrowClockwise"
	case Circular_arrowAnticlockwise:
		res = "Circular_arrowAnticlockwise"
	}
	return
}

func (circular_arrow Circular_arrow) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Circular_arrowClockwise")
	res = append(res, "Circular_arrowAnticlockwise")

	return
}

func (circular_arrow Circular_arrow) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "clockwise")
	res = append(res, "anticlockwise")

	return
}

// Utility function for Clef_sign
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (clef_sign Clef_sign) ToString() (res string) {

	// migration of former implementation of enum
	switch clef_sign {
	// insertion code per enum code
	case Clef_signG:
		res = "G"
	case Clef_signF:
		res = "F"
	case Clef_signC:
		res = "C"
	case Clef_signPercussion:
		res = "percussion"
	case Clef_signTAB:
		res = "TAB"
	case Clef_signJianpu:
		res = "jianpu"
	case Clef_signNone:
		res = "none"
	}
	return
}

func (clef_sign *Clef_sign) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "G":
		*clef_sign = Clef_signG
		return
	case "F":
		*clef_sign = Clef_signF
		return
	case "C":
		*clef_sign = Clef_signC
		return
	case "percussion":
		*clef_sign = Clef_signPercussion
		return
	case "TAB":
		*clef_sign = Clef_signTAB
		return
	case "jianpu":
		*clef_sign = Clef_signJianpu
		return
	case "none":
		*clef_sign = Clef_signNone
		return
	default:
		return errUnkownEnum
	}
}

func (clef_sign *Clef_sign) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Clef_signG":
		*clef_sign = Clef_signG
	case "Clef_signF":
		*clef_sign = Clef_signF
	case "Clef_signC":
		*clef_sign = Clef_signC
	case "Clef_signPercussion":
		*clef_sign = Clef_signPercussion
	case "Clef_signTAB":
		*clef_sign = Clef_signTAB
	case "Clef_signJianpu":
		*clef_sign = Clef_signJianpu
	case "Clef_signNone":
		*clef_sign = Clef_signNone
	default:
		return errUnkownEnum
	}
	return
}

func (clef_sign *Clef_sign) ToCodeString() (res string) {

	switch *clef_sign {
	// insertion code per enum code
	case Clef_signG:
		res = "Clef_signG"
	case Clef_signF:
		res = "Clef_signF"
	case Clef_signC:
		res = "Clef_signC"
	case Clef_signPercussion:
		res = "Clef_signPercussion"
	case Clef_signTAB:
		res = "Clef_signTAB"
	case Clef_signJianpu:
		res = "Clef_signJianpu"
	case Clef_signNone:
		res = "Clef_signNone"
	}
	return
}

func (clef_sign Clef_sign) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Clef_signG")
	res = append(res, "Clef_signF")
	res = append(res, "Clef_signC")
	res = append(res, "Clef_signPercussion")
	res = append(res, "Clef_signTAB")
	res = append(res, "Clef_signJianpu")
	res = append(res, "Clef_signNone")

	return
}

func (clef_sign Clef_sign) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "G")
	res = append(res, "F")
	res = append(res, "C")
	res = append(res, "percussion")
	res = append(res, "TAB")
	res = append(res, "jianpu")
	res = append(res, "none")

	return
}

// Utility function for Color
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (color Color) ToString() (res string) {

	// migration of former implementation of enum
	switch color {
	// insertion code per enum code
	}
	return
}

func (color *Color) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (color *Color) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (color *Color) ToCodeString() (res string) {

	switch *color {
	// insertion code per enum code
	}
	return
}

func (color Color) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (color Color) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Comma_separated_text
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (comma_separated_text Comma_separated_text) ToString() (res string) {

	// migration of former implementation of enum
	switch comma_separated_text {
	// insertion code per enum code
	}
	return
}

func (comma_separated_text *Comma_separated_text) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (comma_separated_text *Comma_separated_text) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (comma_separated_text *Comma_separated_text) ToCodeString() (res string) {

	switch *comma_separated_text {
	// insertion code per enum code
	}
	return
}

func (comma_separated_text Comma_separated_text) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (comma_separated_text Comma_separated_text) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Css_font_size
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (css_font_size Css_font_size) ToString() (res string) {

	// migration of former implementation of enum
	switch css_font_size {
	// insertion code per enum code
	case Css_font_sizeXx_small:
		res = "xx-small"
	case Css_font_sizeX_small:
		res = "x-small"
	case Css_font_sizeSmall:
		res = "small"
	case Css_font_sizeMedium:
		res = "medium"
	case Css_font_sizeLarge:
		res = "large"
	case Css_font_sizeX_large:
		res = "x-large"
	case Css_font_sizeXx_large:
		res = "xx-large"
	}
	return
}

func (css_font_size *Css_font_size) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "xx-small":
		*css_font_size = Css_font_sizeXx_small
		return
	case "x-small":
		*css_font_size = Css_font_sizeX_small
		return
	case "small":
		*css_font_size = Css_font_sizeSmall
		return
	case "medium":
		*css_font_size = Css_font_sizeMedium
		return
	case "large":
		*css_font_size = Css_font_sizeLarge
		return
	case "x-large":
		*css_font_size = Css_font_sizeX_large
		return
	case "xx-large":
		*css_font_size = Css_font_sizeXx_large
		return
	default:
		return errUnkownEnum
	}
}

func (css_font_size *Css_font_size) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Css_font_sizeXx_small":
		*css_font_size = Css_font_sizeXx_small
	case "Css_font_sizeX_small":
		*css_font_size = Css_font_sizeX_small
	case "Css_font_sizeSmall":
		*css_font_size = Css_font_sizeSmall
	case "Css_font_sizeMedium":
		*css_font_size = Css_font_sizeMedium
	case "Css_font_sizeLarge":
		*css_font_size = Css_font_sizeLarge
	case "Css_font_sizeX_large":
		*css_font_size = Css_font_sizeX_large
	case "Css_font_sizeXx_large":
		*css_font_size = Css_font_sizeXx_large
	default:
		return errUnkownEnum
	}
	return
}

func (css_font_size *Css_font_size) ToCodeString() (res string) {

	switch *css_font_size {
	// insertion code per enum code
	case Css_font_sizeXx_small:
		res = "Css_font_sizeXx_small"
	case Css_font_sizeX_small:
		res = "Css_font_sizeX_small"
	case Css_font_sizeSmall:
		res = "Css_font_sizeSmall"
	case Css_font_sizeMedium:
		res = "Css_font_sizeMedium"
	case Css_font_sizeLarge:
		res = "Css_font_sizeLarge"
	case Css_font_sizeX_large:
		res = "Css_font_sizeX_large"
	case Css_font_sizeXx_large:
		res = "Css_font_sizeXx_large"
	}
	return
}

func (css_font_size Css_font_size) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Css_font_sizeXx_small")
	res = append(res, "Css_font_sizeX_small")
	res = append(res, "Css_font_sizeSmall")
	res = append(res, "Css_font_sizeMedium")
	res = append(res, "Css_font_sizeLarge")
	res = append(res, "Css_font_sizeX_large")
	res = append(res, "Css_font_sizeXx_large")

	return
}

func (css_font_size Css_font_size) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "xx-small")
	res = append(res, "x-small")
	res = append(res, "small")
	res = append(res, "medium")
	res = append(res, "large")
	res = append(res, "x-large")
	res = append(res, "xx-large")

	return
}

// Utility function for Degree_symbol_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (degree_symbol_value Degree_symbol_value) ToString() (res string) {

	// migration of former implementation of enum
	switch degree_symbol_value {
	// insertion code per enum code
	case Degree_symbol_valueMajor:
		res = "major"
	case Degree_symbol_valueMinor:
		res = "minor"
	case Degree_symbol_valueAugmented:
		res = "augmented"
	case Degree_symbol_valueDiminished:
		res = "diminished"
	case Degree_symbol_valueHalf_diminished:
		res = "half-diminished"
	}
	return
}

func (degree_symbol_value *Degree_symbol_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*degree_symbol_value = Degree_symbol_valueMajor
		return
	case "minor":
		*degree_symbol_value = Degree_symbol_valueMinor
		return
	case "augmented":
		*degree_symbol_value = Degree_symbol_valueAugmented
		return
	case "diminished":
		*degree_symbol_value = Degree_symbol_valueDiminished
		return
	case "half-diminished":
		*degree_symbol_value = Degree_symbol_valueHalf_diminished
		return
	default:
		return errUnkownEnum
	}
}

func (degree_symbol_value *Degree_symbol_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Degree_symbol_valueMajor":
		*degree_symbol_value = Degree_symbol_valueMajor
	case "Degree_symbol_valueMinor":
		*degree_symbol_value = Degree_symbol_valueMinor
	case "Degree_symbol_valueAugmented":
		*degree_symbol_value = Degree_symbol_valueAugmented
	case "Degree_symbol_valueDiminished":
		*degree_symbol_value = Degree_symbol_valueDiminished
	case "Degree_symbol_valueHalf_diminished":
		*degree_symbol_value = Degree_symbol_valueHalf_diminished
	default:
		return errUnkownEnum
	}
	return
}

func (degree_symbol_value *Degree_symbol_value) ToCodeString() (res string) {

	switch *degree_symbol_value {
	// insertion code per enum code
	case Degree_symbol_valueMajor:
		res = "Degree_symbol_valueMajor"
	case Degree_symbol_valueMinor:
		res = "Degree_symbol_valueMinor"
	case Degree_symbol_valueAugmented:
		res = "Degree_symbol_valueAugmented"
	case Degree_symbol_valueDiminished:
		res = "Degree_symbol_valueDiminished"
	case Degree_symbol_valueHalf_diminished:
		res = "Degree_symbol_valueHalf_diminished"
	}
	return
}

func (degree_symbol_value Degree_symbol_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Degree_symbol_valueMajor")
	res = append(res, "Degree_symbol_valueMinor")
	res = append(res, "Degree_symbol_valueAugmented")
	res = append(res, "Degree_symbol_valueDiminished")
	res = append(res, "Degree_symbol_valueHalf_diminished")

	return
}

func (degree_symbol_value Degree_symbol_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "augmented")
	res = append(res, "diminished")
	res = append(res, "half-diminished")

	return
}

// Utility function for Degree_type_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (degree_type_value Degree_type_value) ToString() (res string) {

	// migration of former implementation of enum
	switch degree_type_value {
	// insertion code per enum code
	case Degree_type_valueAdd:
		res = "add"
	case Degree_type_valueAlter:
		res = "alter"
	case Degree_type_valueSubtract:
		res = "subtract"
	}
	return
}

func (degree_type_value *Degree_type_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "add":
		*degree_type_value = Degree_type_valueAdd
		return
	case "alter":
		*degree_type_value = Degree_type_valueAlter
		return
	case "subtract":
		*degree_type_value = Degree_type_valueSubtract
		return
	default:
		return errUnkownEnum
	}
}

func (degree_type_value *Degree_type_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Degree_type_valueAdd":
		*degree_type_value = Degree_type_valueAdd
	case "Degree_type_valueAlter":
		*degree_type_value = Degree_type_valueAlter
	case "Degree_type_valueSubtract":
		*degree_type_value = Degree_type_valueSubtract
	default:
		return errUnkownEnum
	}
	return
}

func (degree_type_value *Degree_type_value) ToCodeString() (res string) {

	switch *degree_type_value {
	// insertion code per enum code
	case Degree_type_valueAdd:
		res = "Degree_type_valueAdd"
	case Degree_type_valueAlter:
		res = "Degree_type_valueAlter"
	case Degree_type_valueSubtract:
		res = "Degree_type_valueSubtract"
	}
	return
}

func (degree_type_value Degree_type_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Degree_type_valueAdd")
	res = append(res, "Degree_type_valueAlter")
	res = append(res, "Degree_type_valueSubtract")

	return
}

func (degree_type_value Degree_type_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "add")
	res = append(res, "alter")
	res = append(res, "subtract")

	return
}

// Utility function for Distance_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (distance_type Distance_type) ToString() (res string) {

	// migration of former implementation of enum
	switch distance_type {
	// insertion code per enum code
	}
	return
}

func (distance_type *Distance_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (distance_type *Distance_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (distance_type *Distance_type) ToCodeString() (res string) {

	switch *distance_type {
	// insertion code per enum code
	}
	return
}

func (distance_type Distance_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (distance_type Distance_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Effect_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (effect_value Effect_value) ToString() (res string) {

	// migration of former implementation of enum
	switch effect_value {
	// insertion code per enum code
	case Effect_valueAnvil:
		res = "anvil"
	case Effect_valueAutohorn:
		res = "auto horn"
	case Effect_valueBirdwhistle:
		res = "bird whistle"
	case Effect_valueCannon:
		res = "cannon"
	case Effect_valueDuckcall:
		res = "duck call"
	case Effect_valueGunshot:
		res = "gun shot"
	case Effect_valueKlaxonhorn:
		res = "klaxon horn"
	case Effect_valueLionsroar:
		res = "lions roar"
	case Effect_valueLotusflute:
		res = "lotus flute"
	case Effect_valueMegaphone:
		res = "megaphone"
	case Effect_valuePolicewhistle:
		res = "police whistle"
	case Effect_valueSiren:
		res = "siren"
	case Effect_valueSlidewhistle:
		res = "slide whistle"
	case Effect_valueThundersheet:
		res = "thunder sheet"
	case Effect_valueWindmachine:
		res = "wind machine"
	case Effect_valueWindwhistle:
		res = "wind whistle"
	}
	return
}

func (effect_value *Effect_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "anvil":
		*effect_value = Effect_valueAnvil
		return
	case "auto horn":
		*effect_value = Effect_valueAutohorn
		return
	case "bird whistle":
		*effect_value = Effect_valueBirdwhistle
		return
	case "cannon":
		*effect_value = Effect_valueCannon
		return
	case "duck call":
		*effect_value = Effect_valueDuckcall
		return
	case "gun shot":
		*effect_value = Effect_valueGunshot
		return
	case "klaxon horn":
		*effect_value = Effect_valueKlaxonhorn
		return
	case "lions roar":
		*effect_value = Effect_valueLionsroar
		return
	case "lotus flute":
		*effect_value = Effect_valueLotusflute
		return
	case "megaphone":
		*effect_value = Effect_valueMegaphone
		return
	case "police whistle":
		*effect_value = Effect_valuePolicewhistle
		return
	case "siren":
		*effect_value = Effect_valueSiren
		return
	case "slide whistle":
		*effect_value = Effect_valueSlidewhistle
		return
	case "thunder sheet":
		*effect_value = Effect_valueThundersheet
		return
	case "wind machine":
		*effect_value = Effect_valueWindmachine
		return
	case "wind whistle":
		*effect_value = Effect_valueWindwhistle
		return
	default:
		return errUnkownEnum
	}
}

func (effect_value *Effect_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Effect_valueAnvil":
		*effect_value = Effect_valueAnvil
	case "Effect_valueAutohorn":
		*effect_value = Effect_valueAutohorn
	case "Effect_valueBirdwhistle":
		*effect_value = Effect_valueBirdwhistle
	case "Effect_valueCannon":
		*effect_value = Effect_valueCannon
	case "Effect_valueDuckcall":
		*effect_value = Effect_valueDuckcall
	case "Effect_valueGunshot":
		*effect_value = Effect_valueGunshot
	case "Effect_valueKlaxonhorn":
		*effect_value = Effect_valueKlaxonhorn
	case "Effect_valueLionsroar":
		*effect_value = Effect_valueLionsroar
	case "Effect_valueLotusflute":
		*effect_value = Effect_valueLotusflute
	case "Effect_valueMegaphone":
		*effect_value = Effect_valueMegaphone
	case "Effect_valuePolicewhistle":
		*effect_value = Effect_valuePolicewhistle
	case "Effect_valueSiren":
		*effect_value = Effect_valueSiren
	case "Effect_valueSlidewhistle":
		*effect_value = Effect_valueSlidewhistle
	case "Effect_valueThundersheet":
		*effect_value = Effect_valueThundersheet
	case "Effect_valueWindmachine":
		*effect_value = Effect_valueWindmachine
	case "Effect_valueWindwhistle":
		*effect_value = Effect_valueWindwhistle
	default:
		return errUnkownEnum
	}
	return
}

func (effect_value *Effect_value) ToCodeString() (res string) {

	switch *effect_value {
	// insertion code per enum code
	case Effect_valueAnvil:
		res = "Effect_valueAnvil"
	case Effect_valueAutohorn:
		res = "Effect_valueAutohorn"
	case Effect_valueBirdwhistle:
		res = "Effect_valueBirdwhistle"
	case Effect_valueCannon:
		res = "Effect_valueCannon"
	case Effect_valueDuckcall:
		res = "Effect_valueDuckcall"
	case Effect_valueGunshot:
		res = "Effect_valueGunshot"
	case Effect_valueKlaxonhorn:
		res = "Effect_valueKlaxonhorn"
	case Effect_valueLionsroar:
		res = "Effect_valueLionsroar"
	case Effect_valueLotusflute:
		res = "Effect_valueLotusflute"
	case Effect_valueMegaphone:
		res = "Effect_valueMegaphone"
	case Effect_valuePolicewhistle:
		res = "Effect_valuePolicewhistle"
	case Effect_valueSiren:
		res = "Effect_valueSiren"
	case Effect_valueSlidewhistle:
		res = "Effect_valueSlidewhistle"
	case Effect_valueThundersheet:
		res = "Effect_valueThundersheet"
	case Effect_valueWindmachine:
		res = "Effect_valueWindmachine"
	case Effect_valueWindwhistle:
		res = "Effect_valueWindwhistle"
	}
	return
}

func (effect_value Effect_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Effect_valueAnvil")
	res = append(res, "Effect_valueAutohorn")
	res = append(res, "Effect_valueBirdwhistle")
	res = append(res, "Effect_valueCannon")
	res = append(res, "Effect_valueDuckcall")
	res = append(res, "Effect_valueGunshot")
	res = append(res, "Effect_valueKlaxonhorn")
	res = append(res, "Effect_valueLionsroar")
	res = append(res, "Effect_valueLotusflute")
	res = append(res, "Effect_valueMegaphone")
	res = append(res, "Effect_valuePolicewhistle")
	res = append(res, "Effect_valueSiren")
	res = append(res, "Effect_valueSlidewhistle")
	res = append(res, "Effect_valueThundersheet")
	res = append(res, "Effect_valueWindmachine")
	res = append(res, "Effect_valueWindwhistle")

	return
}

func (effect_value Effect_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "anvil")
	res = append(res, "auto horn")
	res = append(res, "bird whistle")
	res = append(res, "cannon")
	res = append(res, "duck call")
	res = append(res, "gun shot")
	res = append(res, "klaxon horn")
	res = append(res, "lions roar")
	res = append(res, "lotus flute")
	res = append(res, "megaphone")
	res = append(res, "police whistle")
	res = append(res, "siren")
	res = append(res, "slide whistle")
	res = append(res, "thunder sheet")
	res = append(res, "wind machine")
	res = append(res, "wind whistle")

	return
}

// Utility function for Enclosure_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enclosure_shape Enclosure_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch enclosure_shape {
	// insertion code per enum code
	case Enclosure_shapeRectangle:
		res = "rectangle"
	case Enclosure_shapeSquare:
		res = "square"
	case Enclosure_shapeOval:
		res = "oval"
	case Enclosure_shapeCircle:
		res = "circle"
	case Enclosure_shapeBracket:
		res = "bracket"
	case Enclosure_shapeInverted_bracket:
		res = "inverted-bracket"
	case Enclosure_shapeTriangle:
		res = "triangle"
	case Enclosure_shapeDiamond:
		res = "diamond"
	case Enclosure_shapePentagon:
		res = "pentagon"
	case Enclosure_shapeHexagon:
		res = "hexagon"
	case Enclosure_shapeHeptagon:
		res = "heptagon"
	case Enclosure_shapeOctagon:
		res = "octagon"
	case Enclosure_shapeNonagon:
		res = "nonagon"
	case Enclosure_shapeDecagon:
		res = "decagon"
	case Enclosure_shapeNone:
		res = "none"
	}
	return
}

func (enclosure_shape *Enclosure_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "rectangle":
		*enclosure_shape = Enclosure_shapeRectangle
		return
	case "square":
		*enclosure_shape = Enclosure_shapeSquare
		return
	case "oval":
		*enclosure_shape = Enclosure_shapeOval
		return
	case "circle":
		*enclosure_shape = Enclosure_shapeCircle
		return
	case "bracket":
		*enclosure_shape = Enclosure_shapeBracket
		return
	case "inverted-bracket":
		*enclosure_shape = Enclosure_shapeInverted_bracket
		return
	case "triangle":
		*enclosure_shape = Enclosure_shapeTriangle
		return
	case "diamond":
		*enclosure_shape = Enclosure_shapeDiamond
		return
	case "pentagon":
		*enclosure_shape = Enclosure_shapePentagon
		return
	case "hexagon":
		*enclosure_shape = Enclosure_shapeHexagon
		return
	case "heptagon":
		*enclosure_shape = Enclosure_shapeHeptagon
		return
	case "octagon":
		*enclosure_shape = Enclosure_shapeOctagon
		return
	case "nonagon":
		*enclosure_shape = Enclosure_shapeNonagon
		return
	case "decagon":
		*enclosure_shape = Enclosure_shapeDecagon
		return
	case "none":
		*enclosure_shape = Enclosure_shapeNone
		return
	default:
		return errUnkownEnum
	}
}

func (enclosure_shape *Enclosure_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Enclosure_shapeRectangle":
		*enclosure_shape = Enclosure_shapeRectangle
	case "Enclosure_shapeSquare":
		*enclosure_shape = Enclosure_shapeSquare
	case "Enclosure_shapeOval":
		*enclosure_shape = Enclosure_shapeOval
	case "Enclosure_shapeCircle":
		*enclosure_shape = Enclosure_shapeCircle
	case "Enclosure_shapeBracket":
		*enclosure_shape = Enclosure_shapeBracket
	case "Enclosure_shapeInverted_bracket":
		*enclosure_shape = Enclosure_shapeInverted_bracket
	case "Enclosure_shapeTriangle":
		*enclosure_shape = Enclosure_shapeTriangle
	case "Enclosure_shapeDiamond":
		*enclosure_shape = Enclosure_shapeDiamond
	case "Enclosure_shapePentagon":
		*enclosure_shape = Enclosure_shapePentagon
	case "Enclosure_shapeHexagon":
		*enclosure_shape = Enclosure_shapeHexagon
	case "Enclosure_shapeHeptagon":
		*enclosure_shape = Enclosure_shapeHeptagon
	case "Enclosure_shapeOctagon":
		*enclosure_shape = Enclosure_shapeOctagon
	case "Enclosure_shapeNonagon":
		*enclosure_shape = Enclosure_shapeNonagon
	case "Enclosure_shapeDecagon":
		*enclosure_shape = Enclosure_shapeDecagon
	case "Enclosure_shapeNone":
		*enclosure_shape = Enclosure_shapeNone
	default:
		return errUnkownEnum
	}
	return
}

func (enclosure_shape *Enclosure_shape) ToCodeString() (res string) {

	switch *enclosure_shape {
	// insertion code per enum code
	case Enclosure_shapeRectangle:
		res = "Enclosure_shapeRectangle"
	case Enclosure_shapeSquare:
		res = "Enclosure_shapeSquare"
	case Enclosure_shapeOval:
		res = "Enclosure_shapeOval"
	case Enclosure_shapeCircle:
		res = "Enclosure_shapeCircle"
	case Enclosure_shapeBracket:
		res = "Enclosure_shapeBracket"
	case Enclosure_shapeInverted_bracket:
		res = "Enclosure_shapeInverted_bracket"
	case Enclosure_shapeTriangle:
		res = "Enclosure_shapeTriangle"
	case Enclosure_shapeDiamond:
		res = "Enclosure_shapeDiamond"
	case Enclosure_shapePentagon:
		res = "Enclosure_shapePentagon"
	case Enclosure_shapeHexagon:
		res = "Enclosure_shapeHexagon"
	case Enclosure_shapeHeptagon:
		res = "Enclosure_shapeHeptagon"
	case Enclosure_shapeOctagon:
		res = "Enclosure_shapeOctagon"
	case Enclosure_shapeNonagon:
		res = "Enclosure_shapeNonagon"
	case Enclosure_shapeDecagon:
		res = "Enclosure_shapeDecagon"
	case Enclosure_shapeNone:
		res = "Enclosure_shapeNone"
	}
	return
}

func (enclosure_shape Enclosure_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Enclosure_shapeRectangle")
	res = append(res, "Enclosure_shapeSquare")
	res = append(res, "Enclosure_shapeOval")
	res = append(res, "Enclosure_shapeCircle")
	res = append(res, "Enclosure_shapeBracket")
	res = append(res, "Enclosure_shapeInverted_bracket")
	res = append(res, "Enclosure_shapeTriangle")
	res = append(res, "Enclosure_shapeDiamond")
	res = append(res, "Enclosure_shapePentagon")
	res = append(res, "Enclosure_shapeHexagon")
	res = append(res, "Enclosure_shapeHeptagon")
	res = append(res, "Enclosure_shapeOctagon")
	res = append(res, "Enclosure_shapeNonagon")
	res = append(res, "Enclosure_shapeDecagon")
	res = append(res, "Enclosure_shapeNone")

	return
}

func (enclosure_shape Enclosure_shape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "rectangle")
	res = append(res, "square")
	res = append(res, "oval")
	res = append(res, "circle")
	res = append(res, "bracket")
	res = append(res, "inverted-bracket")
	res = append(res, "triangle")
	res = append(res, "diamond")
	res = append(res, "pentagon")
	res = append(res, "hexagon")
	res = append(res, "heptagon")
	res = append(res, "octagon")
	res = append(res, "nonagon")
	res = append(res, "decagon")
	res = append(res, "none")

	return
}

// Utility function for Ending_number
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (ending_number Ending_number) ToString() (res string) {

	// migration of former implementation of enum
	switch ending_number {
	// insertion code per enum code
	}
	return
}

func (ending_number *Ending_number) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (ending_number *Ending_number) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (ending_number *Ending_number) ToCodeString() (res string) {

	switch *ending_number {
	// insertion code per enum code
	}
	return
}

func (ending_number Ending_number) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (ending_number Ending_number) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Fan
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (fan Fan) ToString() (res string) {

	// migration of former implementation of enum
	switch fan {
	// insertion code per enum code
	case FanAccel:
		res = "accel"
	case FanRit:
		res = "rit"
	case FanNone:
		res = "none"
	}
	return
}

func (fan *Fan) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "accel":
		*fan = FanAccel
		return
	case "rit":
		*fan = FanRit
		return
	case "none":
		*fan = FanNone
		return
	default:
		return errUnkownEnum
	}
}

func (fan *Fan) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FanAccel":
		*fan = FanAccel
	case "FanRit":
		*fan = FanRit
	case "FanNone":
		*fan = FanNone
	default:
		return errUnkownEnum
	}
	return
}

func (fan *Fan) ToCodeString() (res string) {

	switch *fan {
	// insertion code per enum code
	case FanAccel:
		res = "FanAccel"
	case FanRit:
		res = "FanRit"
	case FanNone:
		res = "FanNone"
	}
	return
}

func (fan Fan) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FanAccel")
	res = append(res, "FanRit")
	res = append(res, "FanNone")

	return
}

func (fan Fan) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "accel")
	res = append(res, "rit")
	res = append(res, "none")

	return
}

// Utility function for Fermata_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (fermata_shape Fermata_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch fermata_shape {
	// insertion code per enum code
	case Fermata_shapeNormal:
		res = "normal"
	case Fermata_shapeAngled:
		res = "angled"
	case Fermata_shapeSquare:
		res = "square"
	case Fermata_shapeDouble_angled:
		res = "double-angled"
	case Fermata_shapeDouble_square:
		res = "double-square"
	case Fermata_shapeDouble_dot:
		res = "double-dot"
	case Fermata_shapeHalf_curve:
		res = "half-curve"
	case Fermata_shapeCurlew:
		res = "curlew"
	case Fermata_shapeEmptyString:
		res = ""
	}
	return
}

func (fermata_shape *Fermata_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*fermata_shape = Fermata_shapeNormal
		return
	case "angled":
		*fermata_shape = Fermata_shapeAngled
		return
	case "square":
		*fermata_shape = Fermata_shapeSquare
		return
	case "double-angled":
		*fermata_shape = Fermata_shapeDouble_angled
		return
	case "double-square":
		*fermata_shape = Fermata_shapeDouble_square
		return
	case "double-dot":
		*fermata_shape = Fermata_shapeDouble_dot
		return
	case "half-curve":
		*fermata_shape = Fermata_shapeHalf_curve
		return
	case "curlew":
		*fermata_shape = Fermata_shapeCurlew
		return
	case "":
		*fermata_shape = Fermata_shapeEmptyString
		return
	default:
		return errUnkownEnum
	}
}

func (fermata_shape *Fermata_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Fermata_shapeNormal":
		*fermata_shape = Fermata_shapeNormal
	case "Fermata_shapeAngled":
		*fermata_shape = Fermata_shapeAngled
	case "Fermata_shapeSquare":
		*fermata_shape = Fermata_shapeSquare
	case "Fermata_shapeDouble_angled":
		*fermata_shape = Fermata_shapeDouble_angled
	case "Fermata_shapeDouble_square":
		*fermata_shape = Fermata_shapeDouble_square
	case "Fermata_shapeDouble_dot":
		*fermata_shape = Fermata_shapeDouble_dot
	case "Fermata_shapeHalf_curve":
		*fermata_shape = Fermata_shapeHalf_curve
	case "Fermata_shapeCurlew":
		*fermata_shape = Fermata_shapeCurlew
	case "Fermata_shapeEmptyString":
		*fermata_shape = Fermata_shapeEmptyString
	default:
		return errUnkownEnum
	}
	return
}

func (fermata_shape *Fermata_shape) ToCodeString() (res string) {

	switch *fermata_shape {
	// insertion code per enum code
	case Fermata_shapeNormal:
		res = "Fermata_shapeNormal"
	case Fermata_shapeAngled:
		res = "Fermata_shapeAngled"
	case Fermata_shapeSquare:
		res = "Fermata_shapeSquare"
	case Fermata_shapeDouble_angled:
		res = "Fermata_shapeDouble_angled"
	case Fermata_shapeDouble_square:
		res = "Fermata_shapeDouble_square"
	case Fermata_shapeDouble_dot:
		res = "Fermata_shapeDouble_dot"
	case Fermata_shapeHalf_curve:
		res = "Fermata_shapeHalf_curve"
	case Fermata_shapeCurlew:
		res = "Fermata_shapeCurlew"
	case Fermata_shapeEmptyString:
		res = "Fermata_shapeEmptyString"
	}
	return
}

func (fermata_shape Fermata_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Fermata_shapeNormal")
	res = append(res, "Fermata_shapeAngled")
	res = append(res, "Fermata_shapeSquare")
	res = append(res, "Fermata_shapeDouble_angled")
	res = append(res, "Fermata_shapeDouble_square")
	res = append(res, "Fermata_shapeDouble_dot")
	res = append(res, "Fermata_shapeHalf_curve")
	res = append(res, "Fermata_shapeCurlew")
	res = append(res, "Fermata_shapeEmptyString")

	return
}

func (fermata_shape Fermata_shape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "angled")
	res = append(res, "square")
	res = append(res, "double-angled")
	res = append(res, "double-square")
	res = append(res, "double-dot")
	res = append(res, "half-curve")
	res = append(res, "curlew")
	res = append(res, "")

	return
}

// Utility function for Font_size
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (font_size Font_size) ToString() (res string) {

	// migration of former implementation of enum
	switch font_size {
	// insertion code per enum code
	}
	return
}

func (font_size *Font_size) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (font_size *Font_size) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (font_size *Font_size) ToCodeString() (res string) {

	switch *font_size {
	// insertion code per enum code
	}
	return
}

func (font_size Font_size) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (font_size Font_size) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Font_style
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (font_style Font_style) ToString() (res string) {

	// migration of former implementation of enum
	switch font_style {
	// insertion code per enum code
	case Font_styleNormal:
		res = "normal"
	case Font_styleItalic:
		res = "italic"
	}
	return
}

func (font_style *Font_style) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*font_style = Font_styleNormal
		return
	case "italic":
		*font_style = Font_styleItalic
		return
	default:
		return errUnkownEnum
	}
}

func (font_style *Font_style) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Font_styleNormal":
		*font_style = Font_styleNormal
	case "Font_styleItalic":
		*font_style = Font_styleItalic
	default:
		return errUnkownEnum
	}
	return
}

func (font_style *Font_style) ToCodeString() (res string) {

	switch *font_style {
	// insertion code per enum code
	case Font_styleNormal:
		res = "Font_styleNormal"
	case Font_styleItalic:
		res = "Font_styleItalic"
	}
	return
}

func (font_style Font_style) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Font_styleNormal")
	res = append(res, "Font_styleItalic")

	return
}

func (font_style Font_style) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "italic")

	return
}

// Utility function for Font_weight
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (font_weight Font_weight) ToString() (res string) {

	// migration of former implementation of enum
	switch font_weight {
	// insertion code per enum code
	case Font_weightNormal:
		res = "normal"
	case Font_weightBold:
		res = "bold"
	}
	return
}

func (font_weight *Font_weight) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "normal":
		*font_weight = Font_weightNormal
		return
	case "bold":
		*font_weight = Font_weightBold
		return
	default:
		return errUnkownEnum
	}
}

func (font_weight *Font_weight) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Font_weightNormal":
		*font_weight = Font_weightNormal
	case "Font_weightBold":
		*font_weight = Font_weightBold
	default:
		return errUnkownEnum
	}
	return
}

func (font_weight *Font_weight) ToCodeString() (res string) {

	switch *font_weight {
	// insertion code per enum code
	case Font_weightNormal:
		res = "Font_weightNormal"
	case Font_weightBold:
		res = "Font_weightBold"
	}
	return
}

func (font_weight Font_weight) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Font_weightNormal")
	res = append(res, "Font_weightBold")

	return
}

func (font_weight Font_weight) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "normal")
	res = append(res, "bold")

	return
}

// Utility function for Glass_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (glass_value Glass_value) ToString() (res string) {

	// migration of former implementation of enum
	switch glass_value {
	// insertion code per enum code
	case Glass_valueGlassharmonica:
		res = "glass harmonica"
	case Glass_valueGlassharp:
		res = "glass harp"
	case Glass_valueWindchimes:
		res = "wind chimes"
	}
	return
}

func (glass_value *Glass_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "glass harmonica":
		*glass_value = Glass_valueGlassharmonica
		return
	case "glass harp":
		*glass_value = Glass_valueGlassharp
		return
	case "wind chimes":
		*glass_value = Glass_valueWindchimes
		return
	default:
		return errUnkownEnum
	}
}

func (glass_value *Glass_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Glass_valueGlassharmonica":
		*glass_value = Glass_valueGlassharmonica
	case "Glass_valueGlassharp":
		*glass_value = Glass_valueGlassharp
	case "Glass_valueWindchimes":
		*glass_value = Glass_valueWindchimes
	default:
		return errUnkownEnum
	}
	return
}

func (glass_value *Glass_value) ToCodeString() (res string) {

	switch *glass_value {
	// insertion code per enum code
	case Glass_valueGlassharmonica:
		res = "Glass_valueGlassharmonica"
	case Glass_valueGlassharp:
		res = "Glass_valueGlassharp"
	case Glass_valueWindchimes:
		res = "Glass_valueWindchimes"
	}
	return
}

func (glass_value Glass_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Glass_valueGlassharmonica")
	res = append(res, "Glass_valueGlassharp")
	res = append(res, "Glass_valueWindchimes")

	return
}

func (glass_value Glass_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "glass harmonica")
	res = append(res, "glass harp")
	res = append(res, "wind chimes")

	return
}

// Utility function for Glyph_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (glyph_type Glyph_type) ToString() (res string) {

	// migration of former implementation of enum
	switch glyph_type {
	// insertion code per enum code
	}
	return
}

func (glyph_type *Glyph_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (glyph_type *Glyph_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (glyph_type *Glyph_type) ToCodeString() (res string) {

	switch *glyph_type {
	// insertion code per enum code
	}
	return
}

func (glyph_type Glyph_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (glyph_type Glyph_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Group_barline_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (group_barline_value Group_barline_value) ToString() (res string) {

	// migration of former implementation of enum
	switch group_barline_value {
	// insertion code per enum code
	case Group_barline_valueYes:
		res = "yes"
	case Group_barline_valueNo:
		res = "no"
	case Group_barline_valueMensurstrich:
		res = "Mensurstrich"
	}
	return
}

func (group_barline_value *Group_barline_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*group_barline_value = Group_barline_valueYes
		return
	case "no":
		*group_barline_value = Group_barline_valueNo
		return
	case "Mensurstrich":
		*group_barline_value = Group_barline_valueMensurstrich
		return
	default:
		return errUnkownEnum
	}
}

func (group_barline_value *Group_barline_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Group_barline_valueYes":
		*group_barline_value = Group_barline_valueYes
	case "Group_barline_valueNo":
		*group_barline_value = Group_barline_valueNo
	case "Group_barline_valueMensurstrich":
		*group_barline_value = Group_barline_valueMensurstrich
	default:
		return errUnkownEnum
	}
	return
}

func (group_barline_value *Group_barline_value) ToCodeString() (res string) {

	switch *group_barline_value {
	// insertion code per enum code
	case Group_barline_valueYes:
		res = "Group_barline_valueYes"
	case Group_barline_valueNo:
		res = "Group_barline_valueNo"
	case Group_barline_valueMensurstrich:
		res = "Group_barline_valueMensurstrich"
	}
	return
}

func (group_barline_value Group_barline_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Group_barline_valueYes")
	res = append(res, "Group_barline_valueNo")
	res = append(res, "Group_barline_valueMensurstrich")

	return
}

func (group_barline_value Group_barline_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "Mensurstrich")

	return
}

// Utility function for Group_symbol_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (group_symbol_value Group_symbol_value) ToString() (res string) {

	// migration of former implementation of enum
	switch group_symbol_value {
	// insertion code per enum code
	case Group_symbol_valueNone:
		res = "none"
	case Group_symbol_valueBrace:
		res = "brace"
	case Group_symbol_valueLine:
		res = "line"
	case Group_symbol_valueBracket:
		res = "bracket"
	case Group_symbol_valueSquare:
		res = "square"
	}
	return
}

func (group_symbol_value *Group_symbol_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*group_symbol_value = Group_symbol_valueNone
		return
	case "brace":
		*group_symbol_value = Group_symbol_valueBrace
		return
	case "line":
		*group_symbol_value = Group_symbol_valueLine
		return
	case "bracket":
		*group_symbol_value = Group_symbol_valueBracket
		return
	case "square":
		*group_symbol_value = Group_symbol_valueSquare
		return
	default:
		return errUnkownEnum
	}
}

func (group_symbol_value *Group_symbol_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Group_symbol_valueNone":
		*group_symbol_value = Group_symbol_valueNone
	case "Group_symbol_valueBrace":
		*group_symbol_value = Group_symbol_valueBrace
	case "Group_symbol_valueLine":
		*group_symbol_value = Group_symbol_valueLine
	case "Group_symbol_valueBracket":
		*group_symbol_value = Group_symbol_valueBracket
	case "Group_symbol_valueSquare":
		*group_symbol_value = Group_symbol_valueSquare
	default:
		return errUnkownEnum
	}
	return
}

func (group_symbol_value *Group_symbol_value) ToCodeString() (res string) {

	switch *group_symbol_value {
	// insertion code per enum code
	case Group_symbol_valueNone:
		res = "Group_symbol_valueNone"
	case Group_symbol_valueBrace:
		res = "Group_symbol_valueBrace"
	case Group_symbol_valueLine:
		res = "Group_symbol_valueLine"
	case Group_symbol_valueBracket:
		res = "Group_symbol_valueBracket"
	case Group_symbol_valueSquare:
		res = "Group_symbol_valueSquare"
	}
	return
}

func (group_symbol_value Group_symbol_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Group_symbol_valueNone")
	res = append(res, "Group_symbol_valueBrace")
	res = append(res, "Group_symbol_valueLine")
	res = append(res, "Group_symbol_valueBracket")
	res = append(res, "Group_symbol_valueSquare")

	return
}

func (group_symbol_value Group_symbol_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "brace")
	res = append(res, "line")
	res = append(res, "bracket")
	res = append(res, "square")

	return
}

// Utility function for Handbell_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (handbell_value Handbell_value) ToString() (res string) {

	// migration of former implementation of enum
	switch handbell_value {
	// insertion code per enum code
	case Handbell_valueBelltree:
		res = "belltree"
	case Handbell_valueDamp:
		res = "damp"
	case Handbell_valueEcho:
		res = "echo"
	case Handbell_valueGyro:
		res = "gyro"
	case Handbell_valueHandmartellato:
		res = "hand martellato"
	case Handbell_valueMalletlift:
		res = "mallet lift"
	case Handbell_valueMallettable:
		res = "mallet table"
	case Handbell_valueMartellato:
		res = "martellato"
	case Handbell_valueMartellatolift:
		res = "martellato lift"
	case Handbell_valueMutedmartellato:
		res = "muted martellato"
	case Handbell_valuePlucklift:
		res = "pluck lift"
	case Handbell_valueSwing:
		res = "swing"
	}
	return
}

func (handbell_value *Handbell_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "belltree":
		*handbell_value = Handbell_valueBelltree
		return
	case "damp":
		*handbell_value = Handbell_valueDamp
		return
	case "echo":
		*handbell_value = Handbell_valueEcho
		return
	case "gyro":
		*handbell_value = Handbell_valueGyro
		return
	case "hand martellato":
		*handbell_value = Handbell_valueHandmartellato
		return
	case "mallet lift":
		*handbell_value = Handbell_valueMalletlift
		return
	case "mallet table":
		*handbell_value = Handbell_valueMallettable
		return
	case "martellato":
		*handbell_value = Handbell_valueMartellato
		return
	case "martellato lift":
		*handbell_value = Handbell_valueMartellatolift
		return
	case "muted martellato":
		*handbell_value = Handbell_valueMutedmartellato
		return
	case "pluck lift":
		*handbell_value = Handbell_valuePlucklift
		return
	case "swing":
		*handbell_value = Handbell_valueSwing
		return
	default:
		return errUnkownEnum
	}
}

func (handbell_value *Handbell_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Handbell_valueBelltree":
		*handbell_value = Handbell_valueBelltree
	case "Handbell_valueDamp":
		*handbell_value = Handbell_valueDamp
	case "Handbell_valueEcho":
		*handbell_value = Handbell_valueEcho
	case "Handbell_valueGyro":
		*handbell_value = Handbell_valueGyro
	case "Handbell_valueHandmartellato":
		*handbell_value = Handbell_valueHandmartellato
	case "Handbell_valueMalletlift":
		*handbell_value = Handbell_valueMalletlift
	case "Handbell_valueMallettable":
		*handbell_value = Handbell_valueMallettable
	case "Handbell_valueMartellato":
		*handbell_value = Handbell_valueMartellato
	case "Handbell_valueMartellatolift":
		*handbell_value = Handbell_valueMartellatolift
	case "Handbell_valueMutedmartellato":
		*handbell_value = Handbell_valueMutedmartellato
	case "Handbell_valuePlucklift":
		*handbell_value = Handbell_valuePlucklift
	case "Handbell_valueSwing":
		*handbell_value = Handbell_valueSwing
	default:
		return errUnkownEnum
	}
	return
}

func (handbell_value *Handbell_value) ToCodeString() (res string) {

	switch *handbell_value {
	// insertion code per enum code
	case Handbell_valueBelltree:
		res = "Handbell_valueBelltree"
	case Handbell_valueDamp:
		res = "Handbell_valueDamp"
	case Handbell_valueEcho:
		res = "Handbell_valueEcho"
	case Handbell_valueGyro:
		res = "Handbell_valueGyro"
	case Handbell_valueHandmartellato:
		res = "Handbell_valueHandmartellato"
	case Handbell_valueMalletlift:
		res = "Handbell_valueMalletlift"
	case Handbell_valueMallettable:
		res = "Handbell_valueMallettable"
	case Handbell_valueMartellato:
		res = "Handbell_valueMartellato"
	case Handbell_valueMartellatolift:
		res = "Handbell_valueMartellatolift"
	case Handbell_valueMutedmartellato:
		res = "Handbell_valueMutedmartellato"
	case Handbell_valuePlucklift:
		res = "Handbell_valuePlucklift"
	case Handbell_valueSwing:
		res = "Handbell_valueSwing"
	}
	return
}

func (handbell_value Handbell_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Handbell_valueBelltree")
	res = append(res, "Handbell_valueDamp")
	res = append(res, "Handbell_valueEcho")
	res = append(res, "Handbell_valueGyro")
	res = append(res, "Handbell_valueHandmartellato")
	res = append(res, "Handbell_valueMalletlift")
	res = append(res, "Handbell_valueMallettable")
	res = append(res, "Handbell_valueMartellato")
	res = append(res, "Handbell_valueMartellatolift")
	res = append(res, "Handbell_valueMutedmartellato")
	res = append(res, "Handbell_valuePlucklift")
	res = append(res, "Handbell_valueSwing")

	return
}

func (handbell_value Handbell_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "belltree")
	res = append(res, "damp")
	res = append(res, "echo")
	res = append(res, "gyro")
	res = append(res, "hand martellato")
	res = append(res, "mallet lift")
	res = append(res, "mallet table")
	res = append(res, "martellato")
	res = append(res, "martellato lift")
	res = append(res, "muted martellato")
	res = append(res, "pluck lift")
	res = append(res, "swing")

	return
}

// Utility function for Harmon_closed_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (harmon_closed_location Harmon_closed_location) ToString() (res string) {

	// migration of former implementation of enum
	switch harmon_closed_location {
	// insertion code per enum code
	case Harmon_closed_locationRight:
		res = "right"
	case Harmon_closed_locationBottom:
		res = "bottom"
	case Harmon_closed_locationLeft:
		res = "left"
	case Harmon_closed_locationTop:
		res = "top"
	}
	return
}

func (harmon_closed_location *Harmon_closed_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*harmon_closed_location = Harmon_closed_locationRight
		return
	case "bottom":
		*harmon_closed_location = Harmon_closed_locationBottom
		return
	case "left":
		*harmon_closed_location = Harmon_closed_locationLeft
		return
	case "top":
		*harmon_closed_location = Harmon_closed_locationTop
		return
	default:
		return errUnkownEnum
	}
}

func (harmon_closed_location *Harmon_closed_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Harmon_closed_locationRight":
		*harmon_closed_location = Harmon_closed_locationRight
	case "Harmon_closed_locationBottom":
		*harmon_closed_location = Harmon_closed_locationBottom
	case "Harmon_closed_locationLeft":
		*harmon_closed_location = Harmon_closed_locationLeft
	case "Harmon_closed_locationTop":
		*harmon_closed_location = Harmon_closed_locationTop
	default:
		return errUnkownEnum
	}
	return
}

func (harmon_closed_location *Harmon_closed_location) ToCodeString() (res string) {

	switch *harmon_closed_location {
	// insertion code per enum code
	case Harmon_closed_locationRight:
		res = "Harmon_closed_locationRight"
	case Harmon_closed_locationBottom:
		res = "Harmon_closed_locationBottom"
	case Harmon_closed_locationLeft:
		res = "Harmon_closed_locationLeft"
	case Harmon_closed_locationTop:
		res = "Harmon_closed_locationTop"
	}
	return
}

func (harmon_closed_location Harmon_closed_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Harmon_closed_locationRight")
	res = append(res, "Harmon_closed_locationBottom")
	res = append(res, "Harmon_closed_locationLeft")
	res = append(res, "Harmon_closed_locationTop")

	return
}

func (harmon_closed_location Harmon_closed_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "top")

	return
}

// Utility function for Harmon_closed_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (harmon_closed_value Harmon_closed_value) ToString() (res string) {

	// migration of former implementation of enum
	switch harmon_closed_value {
	// insertion code per enum code
	case Harmon_closed_valueYes:
		res = "yes"
	case Harmon_closed_valueNo:
		res = "no"
	case Harmon_closed_valueHalf:
		res = "half"
	}
	return
}

func (harmon_closed_value *Harmon_closed_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*harmon_closed_value = Harmon_closed_valueYes
		return
	case "no":
		*harmon_closed_value = Harmon_closed_valueNo
		return
	case "half":
		*harmon_closed_value = Harmon_closed_valueHalf
		return
	default:
		return errUnkownEnum
	}
}

func (harmon_closed_value *Harmon_closed_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Harmon_closed_valueYes":
		*harmon_closed_value = Harmon_closed_valueYes
	case "Harmon_closed_valueNo":
		*harmon_closed_value = Harmon_closed_valueNo
	case "Harmon_closed_valueHalf":
		*harmon_closed_value = Harmon_closed_valueHalf
	default:
		return errUnkownEnum
	}
	return
}

func (harmon_closed_value *Harmon_closed_value) ToCodeString() (res string) {

	switch *harmon_closed_value {
	// insertion code per enum code
	case Harmon_closed_valueYes:
		res = "Harmon_closed_valueYes"
	case Harmon_closed_valueNo:
		res = "Harmon_closed_valueNo"
	case Harmon_closed_valueHalf:
		res = "Harmon_closed_valueHalf"
	}
	return
}

func (harmon_closed_value Harmon_closed_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Harmon_closed_valueYes")
	res = append(res, "Harmon_closed_valueNo")
	res = append(res, "Harmon_closed_valueHalf")

	return
}

func (harmon_closed_value Harmon_closed_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "half")

	return
}

// Utility function for Harmony_arrangement
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (harmony_arrangement Harmony_arrangement) ToString() (res string) {

	// migration of former implementation of enum
	switch harmony_arrangement {
	// insertion code per enum code
	case Harmony_arrangementVertical:
		res = "vertical"
	case Harmony_arrangementHorizontal:
		res = "horizontal"
	case Harmony_arrangementDiagonal:
		res = "diagonal"
	}
	return
}

func (harmony_arrangement *Harmony_arrangement) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "vertical":
		*harmony_arrangement = Harmony_arrangementVertical
		return
	case "horizontal":
		*harmony_arrangement = Harmony_arrangementHorizontal
		return
	case "diagonal":
		*harmony_arrangement = Harmony_arrangementDiagonal
		return
	default:
		return errUnkownEnum
	}
}

func (harmony_arrangement *Harmony_arrangement) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Harmony_arrangementVertical":
		*harmony_arrangement = Harmony_arrangementVertical
	case "Harmony_arrangementHorizontal":
		*harmony_arrangement = Harmony_arrangementHorizontal
	case "Harmony_arrangementDiagonal":
		*harmony_arrangement = Harmony_arrangementDiagonal
	default:
		return errUnkownEnum
	}
	return
}

func (harmony_arrangement *Harmony_arrangement) ToCodeString() (res string) {

	switch *harmony_arrangement {
	// insertion code per enum code
	case Harmony_arrangementVertical:
		res = "Harmony_arrangementVertical"
	case Harmony_arrangementHorizontal:
		res = "Harmony_arrangementHorizontal"
	case Harmony_arrangementDiagonal:
		res = "Harmony_arrangementDiagonal"
	}
	return
}

func (harmony_arrangement Harmony_arrangement) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Harmony_arrangementVertical")
	res = append(res, "Harmony_arrangementHorizontal")
	res = append(res, "Harmony_arrangementDiagonal")

	return
}

func (harmony_arrangement Harmony_arrangement) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "vertical")
	res = append(res, "horizontal")
	res = append(res, "diagonal")

	return
}

// Utility function for Harmony_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (harmony_type Harmony_type) ToString() (res string) {

	// migration of former implementation of enum
	switch harmony_type {
	// insertion code per enum code
	case Harmony_typeExplicit:
		res = "explicit"
	case Harmony_typeImplied:
		res = "implied"
	case Harmony_typeAlternate:
		res = "alternate"
	}
	return
}

func (harmony_type *Harmony_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "explicit":
		*harmony_type = Harmony_typeExplicit
		return
	case "implied":
		*harmony_type = Harmony_typeImplied
		return
	case "alternate":
		*harmony_type = Harmony_typeAlternate
		return
	default:
		return errUnkownEnum
	}
}

func (harmony_type *Harmony_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Harmony_typeExplicit":
		*harmony_type = Harmony_typeExplicit
	case "Harmony_typeImplied":
		*harmony_type = Harmony_typeImplied
	case "Harmony_typeAlternate":
		*harmony_type = Harmony_typeAlternate
	default:
		return errUnkownEnum
	}
	return
}

func (harmony_type *Harmony_type) ToCodeString() (res string) {

	switch *harmony_type {
	// insertion code per enum code
	case Harmony_typeExplicit:
		res = "Harmony_typeExplicit"
	case Harmony_typeImplied:
		res = "Harmony_typeImplied"
	case Harmony_typeAlternate:
		res = "Harmony_typeAlternate"
	}
	return
}

func (harmony_type Harmony_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Harmony_typeExplicit")
	res = append(res, "Harmony_typeImplied")
	res = append(res, "Harmony_typeAlternate")

	return
}

func (harmony_type Harmony_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "explicit")
	res = append(res, "implied")
	res = append(res, "alternate")

	return
}

// Utility function for Hole_closed_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (hole_closed_location Hole_closed_location) ToString() (res string) {

	// migration of former implementation of enum
	switch hole_closed_location {
	// insertion code per enum code
	case Hole_closed_locationRight:
		res = "right"
	case Hole_closed_locationBottom:
		res = "bottom"
	case Hole_closed_locationLeft:
		res = "left"
	case Hole_closed_locationTop:
		res = "top"
	}
	return
}

func (hole_closed_location *Hole_closed_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*hole_closed_location = Hole_closed_locationRight
		return
	case "bottom":
		*hole_closed_location = Hole_closed_locationBottom
		return
	case "left":
		*hole_closed_location = Hole_closed_locationLeft
		return
	case "top":
		*hole_closed_location = Hole_closed_locationTop
		return
	default:
		return errUnkownEnum
	}
}

func (hole_closed_location *Hole_closed_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Hole_closed_locationRight":
		*hole_closed_location = Hole_closed_locationRight
	case "Hole_closed_locationBottom":
		*hole_closed_location = Hole_closed_locationBottom
	case "Hole_closed_locationLeft":
		*hole_closed_location = Hole_closed_locationLeft
	case "Hole_closed_locationTop":
		*hole_closed_location = Hole_closed_locationTop
	default:
		return errUnkownEnum
	}
	return
}

func (hole_closed_location *Hole_closed_location) ToCodeString() (res string) {

	switch *hole_closed_location {
	// insertion code per enum code
	case Hole_closed_locationRight:
		res = "Hole_closed_locationRight"
	case Hole_closed_locationBottom:
		res = "Hole_closed_locationBottom"
	case Hole_closed_locationLeft:
		res = "Hole_closed_locationLeft"
	case Hole_closed_locationTop:
		res = "Hole_closed_locationTop"
	}
	return
}

func (hole_closed_location Hole_closed_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Hole_closed_locationRight")
	res = append(res, "Hole_closed_locationBottom")
	res = append(res, "Hole_closed_locationLeft")
	res = append(res, "Hole_closed_locationTop")

	return
}

func (hole_closed_location Hole_closed_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "bottom")
	res = append(res, "left")
	res = append(res, "top")

	return
}

// Utility function for Hole_closed_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (hole_closed_value Hole_closed_value) ToString() (res string) {

	// migration of former implementation of enum
	switch hole_closed_value {
	// insertion code per enum code
	case Hole_closed_valueYes:
		res = "yes"
	case Hole_closed_valueNo:
		res = "no"
	case Hole_closed_valueHalf:
		res = "half"
	}
	return
}

func (hole_closed_value *Hole_closed_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*hole_closed_value = Hole_closed_valueYes
		return
	case "no":
		*hole_closed_value = Hole_closed_valueNo
		return
	case "half":
		*hole_closed_value = Hole_closed_valueHalf
		return
	default:
		return errUnkownEnum
	}
}

func (hole_closed_value *Hole_closed_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Hole_closed_valueYes":
		*hole_closed_value = Hole_closed_valueYes
	case "Hole_closed_valueNo":
		*hole_closed_value = Hole_closed_valueNo
	case "Hole_closed_valueHalf":
		*hole_closed_value = Hole_closed_valueHalf
	default:
		return errUnkownEnum
	}
	return
}

func (hole_closed_value *Hole_closed_value) ToCodeString() (res string) {

	switch *hole_closed_value {
	// insertion code per enum code
	case Hole_closed_valueYes:
		res = "Hole_closed_valueYes"
	case Hole_closed_valueNo:
		res = "Hole_closed_valueNo"
	case Hole_closed_valueHalf:
		res = "Hole_closed_valueHalf"
	}
	return
}

func (hole_closed_value Hole_closed_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Hole_closed_valueYes")
	res = append(res, "Hole_closed_valueNo")
	res = append(res, "Hole_closed_valueHalf")

	return
}

func (hole_closed_value Hole_closed_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")
	res = append(res, "half")

	return
}

// Utility function for Kind_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (kind_value Kind_value) ToString() (res string) {

	// migration of former implementation of enum
	switch kind_value {
	// insertion code per enum code
	case Kind_valueMajor:
		res = "major"
	case Kind_valueMinor:
		res = "minor"
	case Kind_valueAugmented:
		res = "augmented"
	case Kind_valueDiminished:
		res = "diminished"
	case Kind_valueDominant:
		res = "dominant"
	case Kind_valueMajor_seventh:
		res = "major-seventh"
	case Kind_valueMinor_seventh:
		res = "minor-seventh"
	case Kind_valueDiminished_seventh:
		res = "diminished-seventh"
	case Kind_valueAugmented_seventh:
		res = "augmented-seventh"
	case Kind_valueHalf_diminished:
		res = "half-diminished"
	case Kind_valueMajor_minor:
		res = "major-minor"
	case Kind_valueMajor_sixth:
		res = "major-sixth"
	case Kind_valueMinor_sixth:
		res = "minor-sixth"
	case Kind_valueDominant_ninth:
		res = "dominant-ninth"
	case Kind_valueMajor_ninth:
		res = "major-ninth"
	case Kind_valueMinor_ninth:
		res = "minor-ninth"
	case Kind_valueDominant_11th:
		res = "dominant-11th"
	case Kind_valueMajor_11th:
		res = "major-11th"
	case Kind_valueMinor_11th:
		res = "minor-11th"
	case Kind_valueDominant_13th:
		res = "dominant-13th"
	case Kind_valueMajor_13th:
		res = "major-13th"
	case Kind_valueMinor_13th:
		res = "minor-13th"
	case Kind_valueSuspended_second:
		res = "suspended-second"
	case Kind_valueSuspended_fourth:
		res = "suspended-fourth"
	case Kind_valueNeapolitan:
		res = "Neapolitan"
	case Kind_valueItalian:
		res = "Italian"
	case Kind_valueFrench:
		res = "French"
	case Kind_valueGerman:
		res = "German"
	case Kind_valuePedal:
		res = "pedal"
	case Kind_valuePower:
		res = "power"
	case Kind_valueTristan:
		res = "Tristan"
	case Kind_valueOther:
		res = "other"
	case Kind_valueNone:
		res = "none"
	}
	return
}

func (kind_value *Kind_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*kind_value = Kind_valueMajor
		return
	case "minor":
		*kind_value = Kind_valueMinor
		return
	case "augmented":
		*kind_value = Kind_valueAugmented
		return
	case "diminished":
		*kind_value = Kind_valueDiminished
		return
	case "dominant":
		*kind_value = Kind_valueDominant
		return
	case "major-seventh":
		*kind_value = Kind_valueMajor_seventh
		return
	case "minor-seventh":
		*kind_value = Kind_valueMinor_seventh
		return
	case "diminished-seventh":
		*kind_value = Kind_valueDiminished_seventh
		return
	case "augmented-seventh":
		*kind_value = Kind_valueAugmented_seventh
		return
	case "half-diminished":
		*kind_value = Kind_valueHalf_diminished
		return
	case "major-minor":
		*kind_value = Kind_valueMajor_minor
		return
	case "major-sixth":
		*kind_value = Kind_valueMajor_sixth
		return
	case "minor-sixth":
		*kind_value = Kind_valueMinor_sixth
		return
	case "dominant-ninth":
		*kind_value = Kind_valueDominant_ninth
		return
	case "major-ninth":
		*kind_value = Kind_valueMajor_ninth
		return
	case "minor-ninth":
		*kind_value = Kind_valueMinor_ninth
		return
	case "dominant-11th":
		*kind_value = Kind_valueDominant_11th
		return
	case "major-11th":
		*kind_value = Kind_valueMajor_11th
		return
	case "minor-11th":
		*kind_value = Kind_valueMinor_11th
		return
	case "dominant-13th":
		*kind_value = Kind_valueDominant_13th
		return
	case "major-13th":
		*kind_value = Kind_valueMajor_13th
		return
	case "minor-13th":
		*kind_value = Kind_valueMinor_13th
		return
	case "suspended-second":
		*kind_value = Kind_valueSuspended_second
		return
	case "suspended-fourth":
		*kind_value = Kind_valueSuspended_fourth
		return
	case "Neapolitan":
		*kind_value = Kind_valueNeapolitan
		return
	case "Italian":
		*kind_value = Kind_valueItalian
		return
	case "French":
		*kind_value = Kind_valueFrench
		return
	case "German":
		*kind_value = Kind_valueGerman
		return
	case "pedal":
		*kind_value = Kind_valuePedal
		return
	case "power":
		*kind_value = Kind_valuePower
		return
	case "Tristan":
		*kind_value = Kind_valueTristan
		return
	case "other":
		*kind_value = Kind_valueOther
		return
	case "none":
		*kind_value = Kind_valueNone
		return
	default:
		return errUnkownEnum
	}
}

func (kind_value *Kind_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Kind_valueMajor":
		*kind_value = Kind_valueMajor
	case "Kind_valueMinor":
		*kind_value = Kind_valueMinor
	case "Kind_valueAugmented":
		*kind_value = Kind_valueAugmented
	case "Kind_valueDiminished":
		*kind_value = Kind_valueDiminished
	case "Kind_valueDominant":
		*kind_value = Kind_valueDominant
	case "Kind_valueMajor_seventh":
		*kind_value = Kind_valueMajor_seventh
	case "Kind_valueMinor_seventh":
		*kind_value = Kind_valueMinor_seventh
	case "Kind_valueDiminished_seventh":
		*kind_value = Kind_valueDiminished_seventh
	case "Kind_valueAugmented_seventh":
		*kind_value = Kind_valueAugmented_seventh
	case "Kind_valueHalf_diminished":
		*kind_value = Kind_valueHalf_diminished
	case "Kind_valueMajor_minor":
		*kind_value = Kind_valueMajor_minor
	case "Kind_valueMajor_sixth":
		*kind_value = Kind_valueMajor_sixth
	case "Kind_valueMinor_sixth":
		*kind_value = Kind_valueMinor_sixth
	case "Kind_valueDominant_ninth":
		*kind_value = Kind_valueDominant_ninth
	case "Kind_valueMajor_ninth":
		*kind_value = Kind_valueMajor_ninth
	case "Kind_valueMinor_ninth":
		*kind_value = Kind_valueMinor_ninth
	case "Kind_valueDominant_11th":
		*kind_value = Kind_valueDominant_11th
	case "Kind_valueMajor_11th":
		*kind_value = Kind_valueMajor_11th
	case "Kind_valueMinor_11th":
		*kind_value = Kind_valueMinor_11th
	case "Kind_valueDominant_13th":
		*kind_value = Kind_valueDominant_13th
	case "Kind_valueMajor_13th":
		*kind_value = Kind_valueMajor_13th
	case "Kind_valueMinor_13th":
		*kind_value = Kind_valueMinor_13th
	case "Kind_valueSuspended_second":
		*kind_value = Kind_valueSuspended_second
	case "Kind_valueSuspended_fourth":
		*kind_value = Kind_valueSuspended_fourth
	case "Kind_valueNeapolitan":
		*kind_value = Kind_valueNeapolitan
	case "Kind_valueItalian":
		*kind_value = Kind_valueItalian
	case "Kind_valueFrench":
		*kind_value = Kind_valueFrench
	case "Kind_valueGerman":
		*kind_value = Kind_valueGerman
	case "Kind_valuePedal":
		*kind_value = Kind_valuePedal
	case "Kind_valuePower":
		*kind_value = Kind_valuePower
	case "Kind_valueTristan":
		*kind_value = Kind_valueTristan
	case "Kind_valueOther":
		*kind_value = Kind_valueOther
	case "Kind_valueNone":
		*kind_value = Kind_valueNone
	default:
		return errUnkownEnum
	}
	return
}

func (kind_value *Kind_value) ToCodeString() (res string) {

	switch *kind_value {
	// insertion code per enum code
	case Kind_valueMajor:
		res = "Kind_valueMajor"
	case Kind_valueMinor:
		res = "Kind_valueMinor"
	case Kind_valueAugmented:
		res = "Kind_valueAugmented"
	case Kind_valueDiminished:
		res = "Kind_valueDiminished"
	case Kind_valueDominant:
		res = "Kind_valueDominant"
	case Kind_valueMajor_seventh:
		res = "Kind_valueMajor_seventh"
	case Kind_valueMinor_seventh:
		res = "Kind_valueMinor_seventh"
	case Kind_valueDiminished_seventh:
		res = "Kind_valueDiminished_seventh"
	case Kind_valueAugmented_seventh:
		res = "Kind_valueAugmented_seventh"
	case Kind_valueHalf_diminished:
		res = "Kind_valueHalf_diminished"
	case Kind_valueMajor_minor:
		res = "Kind_valueMajor_minor"
	case Kind_valueMajor_sixth:
		res = "Kind_valueMajor_sixth"
	case Kind_valueMinor_sixth:
		res = "Kind_valueMinor_sixth"
	case Kind_valueDominant_ninth:
		res = "Kind_valueDominant_ninth"
	case Kind_valueMajor_ninth:
		res = "Kind_valueMajor_ninth"
	case Kind_valueMinor_ninth:
		res = "Kind_valueMinor_ninth"
	case Kind_valueDominant_11th:
		res = "Kind_valueDominant_11th"
	case Kind_valueMajor_11th:
		res = "Kind_valueMajor_11th"
	case Kind_valueMinor_11th:
		res = "Kind_valueMinor_11th"
	case Kind_valueDominant_13th:
		res = "Kind_valueDominant_13th"
	case Kind_valueMajor_13th:
		res = "Kind_valueMajor_13th"
	case Kind_valueMinor_13th:
		res = "Kind_valueMinor_13th"
	case Kind_valueSuspended_second:
		res = "Kind_valueSuspended_second"
	case Kind_valueSuspended_fourth:
		res = "Kind_valueSuspended_fourth"
	case Kind_valueNeapolitan:
		res = "Kind_valueNeapolitan"
	case Kind_valueItalian:
		res = "Kind_valueItalian"
	case Kind_valueFrench:
		res = "Kind_valueFrench"
	case Kind_valueGerman:
		res = "Kind_valueGerman"
	case Kind_valuePedal:
		res = "Kind_valuePedal"
	case Kind_valuePower:
		res = "Kind_valuePower"
	case Kind_valueTristan:
		res = "Kind_valueTristan"
	case Kind_valueOther:
		res = "Kind_valueOther"
	case Kind_valueNone:
		res = "Kind_valueNone"
	}
	return
}

func (kind_value Kind_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Kind_valueMajor")
	res = append(res, "Kind_valueMinor")
	res = append(res, "Kind_valueAugmented")
	res = append(res, "Kind_valueDiminished")
	res = append(res, "Kind_valueDominant")
	res = append(res, "Kind_valueMajor_seventh")
	res = append(res, "Kind_valueMinor_seventh")
	res = append(res, "Kind_valueDiminished_seventh")
	res = append(res, "Kind_valueAugmented_seventh")
	res = append(res, "Kind_valueHalf_diminished")
	res = append(res, "Kind_valueMajor_minor")
	res = append(res, "Kind_valueMajor_sixth")
	res = append(res, "Kind_valueMinor_sixth")
	res = append(res, "Kind_valueDominant_ninth")
	res = append(res, "Kind_valueMajor_ninth")
	res = append(res, "Kind_valueMinor_ninth")
	res = append(res, "Kind_valueDominant_11th")
	res = append(res, "Kind_valueMajor_11th")
	res = append(res, "Kind_valueMinor_11th")
	res = append(res, "Kind_valueDominant_13th")
	res = append(res, "Kind_valueMajor_13th")
	res = append(res, "Kind_valueMinor_13th")
	res = append(res, "Kind_valueSuspended_second")
	res = append(res, "Kind_valueSuspended_fourth")
	res = append(res, "Kind_valueNeapolitan")
	res = append(res, "Kind_valueItalian")
	res = append(res, "Kind_valueFrench")
	res = append(res, "Kind_valueGerman")
	res = append(res, "Kind_valuePedal")
	res = append(res, "Kind_valuePower")
	res = append(res, "Kind_valueTristan")
	res = append(res, "Kind_valueOther")
	res = append(res, "Kind_valueNone")

	return
}

func (kind_value Kind_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "augmented")
	res = append(res, "diminished")
	res = append(res, "dominant")
	res = append(res, "major-seventh")
	res = append(res, "minor-seventh")
	res = append(res, "diminished-seventh")
	res = append(res, "augmented-seventh")
	res = append(res, "half-diminished")
	res = append(res, "major-minor")
	res = append(res, "major-sixth")
	res = append(res, "minor-sixth")
	res = append(res, "dominant-ninth")
	res = append(res, "major-ninth")
	res = append(res, "minor-ninth")
	res = append(res, "dominant-11th")
	res = append(res, "major-11th")
	res = append(res, "minor-11th")
	res = append(res, "dominant-13th")
	res = append(res, "major-13th")
	res = append(res, "minor-13th")
	res = append(res, "suspended-second")
	res = append(res, "suspended-fourth")
	res = append(res, "Neapolitan")
	res = append(res, "Italian")
	res = append(res, "French")
	res = append(res, "German")
	res = append(res, "pedal")
	res = append(res, "power")
	res = append(res, "Tristan")
	res = append(res, "other")
	res = append(res, "none")

	return
}

// Utility function for Left_center_right
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (left_center_right Left_center_right) ToString() (res string) {

	// migration of former implementation of enum
	switch left_center_right {
	// insertion code per enum code
	case Left_center_rightLeft:
		res = "left"
	case Left_center_rightCenter:
		res = "center"
	case Left_center_rightRight:
		res = "right"
	}
	return
}

func (left_center_right *Left_center_right) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*left_center_right = Left_center_rightLeft
		return
	case "center":
		*left_center_right = Left_center_rightCenter
		return
	case "right":
		*left_center_right = Left_center_rightRight
		return
	default:
		return errUnkownEnum
	}
}

func (left_center_right *Left_center_right) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Left_center_rightLeft":
		*left_center_right = Left_center_rightLeft
	case "Left_center_rightCenter":
		*left_center_right = Left_center_rightCenter
	case "Left_center_rightRight":
		*left_center_right = Left_center_rightRight
	default:
		return errUnkownEnum
	}
	return
}

func (left_center_right *Left_center_right) ToCodeString() (res string) {

	switch *left_center_right {
	// insertion code per enum code
	case Left_center_rightLeft:
		res = "Left_center_rightLeft"
	case Left_center_rightCenter:
		res = "Left_center_rightCenter"
	case Left_center_rightRight:
		res = "Left_center_rightRight"
	}
	return
}

func (left_center_right Left_center_right) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Left_center_rightLeft")
	res = append(res, "Left_center_rightCenter")
	res = append(res, "Left_center_rightRight")

	return
}

func (left_center_right Left_center_right) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "center")
	res = append(res, "right")

	return
}

// Utility function for Left_right
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (left_right Left_right) ToString() (res string) {

	// migration of former implementation of enum
	switch left_right {
	// insertion code per enum code
	case Left_rightLeft:
		res = "left"
	case Left_rightRight:
		res = "right"
	}
	return
}

func (left_right *Left_right) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*left_right = Left_rightLeft
		return
	case "right":
		*left_right = Left_rightRight
		return
	default:
		return errUnkownEnum
	}
}

func (left_right *Left_right) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Left_rightLeft":
		*left_right = Left_rightLeft
	case "Left_rightRight":
		*left_right = Left_rightRight
	default:
		return errUnkownEnum
	}
	return
}

func (left_right *Left_right) ToCodeString() (res string) {

	switch *left_right {
	// insertion code per enum code
	case Left_rightLeft:
		res = "Left_rightLeft"
	case Left_rightRight:
		res = "Left_rightRight"
	}
	return
}

func (left_right Left_right) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Left_rightLeft")
	res = append(res, "Left_rightRight")

	return
}

func (left_right Left_right) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")

	return
}

// Utility function for Line_end
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (line_end Line_end) ToString() (res string) {

	// migration of former implementation of enum
	switch line_end {
	// insertion code per enum code
	case Line_endUp:
		res = "up"
	case Line_endDown:
		res = "down"
	case Line_endBoth:
		res = "both"
	case Line_endArrow:
		res = "arrow"
	case Line_endNone:
		res = "none"
	}
	return
}

func (line_end *Line_end) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*line_end = Line_endUp
		return
	case "down":
		*line_end = Line_endDown
		return
	case "both":
		*line_end = Line_endBoth
		return
	case "arrow":
		*line_end = Line_endArrow
		return
	case "none":
		*line_end = Line_endNone
		return
	default:
		return errUnkownEnum
	}
}

func (line_end *Line_end) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Line_endUp":
		*line_end = Line_endUp
	case "Line_endDown":
		*line_end = Line_endDown
	case "Line_endBoth":
		*line_end = Line_endBoth
	case "Line_endArrow":
		*line_end = Line_endArrow
	case "Line_endNone":
		*line_end = Line_endNone
	default:
		return errUnkownEnum
	}
	return
}

func (line_end *Line_end) ToCodeString() (res string) {

	switch *line_end {
	// insertion code per enum code
	case Line_endUp:
		res = "Line_endUp"
	case Line_endDown:
		res = "Line_endDown"
	case Line_endBoth:
		res = "Line_endBoth"
	case Line_endArrow:
		res = "Line_endArrow"
	case Line_endNone:
		res = "Line_endNone"
	}
	return
}

func (line_end Line_end) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Line_endUp")
	res = append(res, "Line_endDown")
	res = append(res, "Line_endBoth")
	res = append(res, "Line_endArrow")
	res = append(res, "Line_endNone")

	return
}

func (line_end Line_end) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "both")
	res = append(res, "arrow")
	res = append(res, "none")

	return
}

// Utility function for Line_length
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (line_length Line_length) ToString() (res string) {

	// migration of former implementation of enum
	switch line_length {
	// insertion code per enum code
	case Line_lengthShort:
		res = "short"
	case Line_lengthMedium:
		res = "medium"
	case Line_lengthLong:
		res = "long"
	}
	return
}

func (line_length *Line_length) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "short":
		*line_length = Line_lengthShort
		return
	case "medium":
		*line_length = Line_lengthMedium
		return
	case "long":
		*line_length = Line_lengthLong
		return
	default:
		return errUnkownEnum
	}
}

func (line_length *Line_length) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Line_lengthShort":
		*line_length = Line_lengthShort
	case "Line_lengthMedium":
		*line_length = Line_lengthMedium
	case "Line_lengthLong":
		*line_length = Line_lengthLong
	default:
		return errUnkownEnum
	}
	return
}

func (line_length *Line_length) ToCodeString() (res string) {

	switch *line_length {
	// insertion code per enum code
	case Line_lengthShort:
		res = "Line_lengthShort"
	case Line_lengthMedium:
		res = "Line_lengthMedium"
	case Line_lengthLong:
		res = "Line_lengthLong"
	}
	return
}

func (line_length Line_length) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Line_lengthShort")
	res = append(res, "Line_lengthMedium")
	res = append(res, "Line_lengthLong")

	return
}

func (line_length Line_length) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "short")
	res = append(res, "medium")
	res = append(res, "long")

	return
}

// Utility function for Line_shape
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (line_shape Line_shape) ToString() (res string) {

	// migration of former implementation of enum
	switch line_shape {
	// insertion code per enum code
	case Line_shapeStraight:
		res = "straight"
	case Line_shapeCurved:
		res = "curved"
	}
	return
}

func (line_shape *Line_shape) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "straight":
		*line_shape = Line_shapeStraight
		return
	case "curved":
		*line_shape = Line_shapeCurved
		return
	default:
		return errUnkownEnum
	}
}

func (line_shape *Line_shape) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Line_shapeStraight":
		*line_shape = Line_shapeStraight
	case "Line_shapeCurved":
		*line_shape = Line_shapeCurved
	default:
		return errUnkownEnum
	}
	return
}

func (line_shape *Line_shape) ToCodeString() (res string) {

	switch *line_shape {
	// insertion code per enum code
	case Line_shapeStraight:
		res = "Line_shapeStraight"
	case Line_shapeCurved:
		res = "Line_shapeCurved"
	}
	return
}

func (line_shape Line_shape) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Line_shapeStraight")
	res = append(res, "Line_shapeCurved")

	return
}

func (line_shape Line_shape) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "straight")
	res = append(res, "curved")

	return
}

// Utility function for Line_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (line_type Line_type) ToString() (res string) {

	// migration of former implementation of enum
	switch line_type {
	// insertion code per enum code
	case Line_typeSolid:
		res = "solid"
	case Line_typeDashed:
		res = "dashed"
	case Line_typeDotted:
		res = "dotted"
	case Line_typeWavy:
		res = "wavy"
	}
	return
}

func (line_type *Line_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "solid":
		*line_type = Line_typeSolid
		return
	case "dashed":
		*line_type = Line_typeDashed
		return
	case "dotted":
		*line_type = Line_typeDotted
		return
	case "wavy":
		*line_type = Line_typeWavy
		return
	default:
		return errUnkownEnum
	}
}

func (line_type *Line_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Line_typeSolid":
		*line_type = Line_typeSolid
	case "Line_typeDashed":
		*line_type = Line_typeDashed
	case "Line_typeDotted":
		*line_type = Line_typeDotted
	case "Line_typeWavy":
		*line_type = Line_typeWavy
	default:
		return errUnkownEnum
	}
	return
}

func (line_type *Line_type) ToCodeString() (res string) {

	switch *line_type {
	// insertion code per enum code
	case Line_typeSolid:
		res = "Line_typeSolid"
	case Line_typeDashed:
		res = "Line_typeDashed"
	case Line_typeDotted:
		res = "Line_typeDotted"
	case Line_typeWavy:
		res = "Line_typeWavy"
	}
	return
}

func (line_type Line_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Line_typeSolid")
	res = append(res, "Line_typeDashed")
	res = append(res, "Line_typeDotted")
	res = append(res, "Line_typeWavy")

	return
}

func (line_type Line_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "solid")
	res = append(res, "dashed")
	res = append(res, "dotted")
	res = append(res, "wavy")

	return
}

// Utility function for Line_width_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (line_width_type Line_width_type) ToString() (res string) {

	// migration of former implementation of enum
	switch line_width_type {
	// insertion code per enum code
	}
	return
}

func (line_width_type *Line_width_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (line_width_type *Line_width_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (line_width_type *Line_width_type) ToCodeString() (res string) {

	switch *line_width_type {
	// insertion code per enum code
	}
	return
}

func (line_width_type Line_width_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (line_width_type Line_width_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Margin_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (margin_type Margin_type) ToString() (res string) {

	// migration of former implementation of enum
	switch margin_type {
	// insertion code per enum code
	case Margin_typeOdd:
		res = "odd"
	case Margin_typeEven:
		res = "even"
	case Margin_typeBoth:
		res = "both"
	}
	return
}

func (margin_type *Margin_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "odd":
		*margin_type = Margin_typeOdd
		return
	case "even":
		*margin_type = Margin_typeEven
		return
	case "both":
		*margin_type = Margin_typeBoth
		return
	default:
		return errUnkownEnum
	}
}

func (margin_type *Margin_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Margin_typeOdd":
		*margin_type = Margin_typeOdd
	case "Margin_typeEven":
		*margin_type = Margin_typeEven
	case "Margin_typeBoth":
		*margin_type = Margin_typeBoth
	default:
		return errUnkownEnum
	}
	return
}

func (margin_type *Margin_type) ToCodeString() (res string) {

	switch *margin_type {
	// insertion code per enum code
	case Margin_typeOdd:
		res = "Margin_typeOdd"
	case Margin_typeEven:
		res = "Margin_typeEven"
	case Margin_typeBoth:
		res = "Margin_typeBoth"
	}
	return
}

func (margin_type Margin_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Margin_typeOdd")
	res = append(res, "Margin_typeEven")
	res = append(res, "Margin_typeBoth")

	return
}

func (margin_type Margin_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "odd")
	res = append(res, "even")
	res = append(res, "both")

	return
}

// Utility function for Measure_numbering_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (measure_numbering_value Measure_numbering_value) ToString() (res string) {

	// migration of former implementation of enum
	switch measure_numbering_value {
	// insertion code per enum code
	case Measure_numbering_valueNone:
		res = "none"
	case Measure_numbering_valueMeasure:
		res = "measure"
	case Measure_numbering_valueSystem:
		res = "system"
	}
	return
}

func (measure_numbering_value *Measure_numbering_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*measure_numbering_value = Measure_numbering_valueNone
		return
	case "measure":
		*measure_numbering_value = Measure_numbering_valueMeasure
		return
	case "system":
		*measure_numbering_value = Measure_numbering_valueSystem
		return
	default:
		return errUnkownEnum
	}
}

func (measure_numbering_value *Measure_numbering_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Measure_numbering_valueNone":
		*measure_numbering_value = Measure_numbering_valueNone
	case "Measure_numbering_valueMeasure":
		*measure_numbering_value = Measure_numbering_valueMeasure
	case "Measure_numbering_valueSystem":
		*measure_numbering_value = Measure_numbering_valueSystem
	default:
		return errUnkownEnum
	}
	return
}

func (measure_numbering_value *Measure_numbering_value) ToCodeString() (res string) {

	switch *measure_numbering_value {
	// insertion code per enum code
	case Measure_numbering_valueNone:
		res = "Measure_numbering_valueNone"
	case Measure_numbering_valueMeasure:
		res = "Measure_numbering_valueMeasure"
	case Measure_numbering_valueSystem:
		res = "Measure_numbering_valueSystem"
	}
	return
}

func (measure_numbering_value Measure_numbering_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Measure_numbering_valueNone")
	res = append(res, "Measure_numbering_valueMeasure")
	res = append(res, "Measure_numbering_valueSystem")

	return
}

func (measure_numbering_value Measure_numbering_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "measure")
	res = append(res, "system")

	return
}

// Utility function for Measure_text
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (measure_text Measure_text) ToString() (res string) {

	// migration of former implementation of enum
	switch measure_text {
	// insertion code per enum code
	}
	return
}

func (measure_text *Measure_text) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (measure_text *Measure_text) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (measure_text *Measure_text) ToCodeString() (res string) {

	switch *measure_text {
	// insertion code per enum code
	}
	return
}

func (measure_text Measure_text) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (measure_text Measure_text) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Membrane_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (membrane_value Membrane_value) ToString() (res string) {

	// migration of former implementation of enum
	switch membrane_value {
	// insertion code per enum code
	case Membrane_valueBassdrum:
		res = "bass drum"
	case Membrane_valueBassdrumonside:
		res = "bass drum on side"
	case Membrane_valueBongos:
		res = "bongos"
	case Membrane_valueChinesetomtom:
		res = "Chinese tomtom"
	case Membrane_valueCongadrum:
		res = "conga drum"
	case Membrane_valueCuica:
		res = "cuica"
	case Membrane_valueGobletdrum:
		res = "goblet drum"
	case Membrane_valueIndo_Americantomtom:
		res = "Indo-American tomtom"
	case Membrane_valueJapanesetomtom:
		res = "Japanese tomtom"
	case Membrane_valueMilitarydrum:
		res = "military drum"
	case Membrane_valueSnaredrum:
		res = "snare drum"
	case Membrane_valueSnaredrumsnaresoff:
		res = "snare drum snares off"
	case Membrane_valueTabla:
		res = "tabla"
	case Membrane_valueTambourine:
		res = "tambourine"
	case Membrane_valueTenordrum:
		res = "tenor drum"
	case Membrane_valueTimbales:
		res = "timbales"
	case Membrane_valueTomtom:
		res = "tomtom"
	}
	return
}

func (membrane_value *Membrane_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bass drum":
		*membrane_value = Membrane_valueBassdrum
		return
	case "bass drum on side":
		*membrane_value = Membrane_valueBassdrumonside
		return
	case "bongos":
		*membrane_value = Membrane_valueBongos
		return
	case "Chinese tomtom":
		*membrane_value = Membrane_valueChinesetomtom
		return
	case "conga drum":
		*membrane_value = Membrane_valueCongadrum
		return
	case "cuica":
		*membrane_value = Membrane_valueCuica
		return
	case "goblet drum":
		*membrane_value = Membrane_valueGobletdrum
		return
	case "Indo-American tomtom":
		*membrane_value = Membrane_valueIndo_Americantomtom
		return
	case "Japanese tomtom":
		*membrane_value = Membrane_valueJapanesetomtom
		return
	case "military drum":
		*membrane_value = Membrane_valueMilitarydrum
		return
	case "snare drum":
		*membrane_value = Membrane_valueSnaredrum
		return
	case "snare drum snares off":
		*membrane_value = Membrane_valueSnaredrumsnaresoff
		return
	case "tabla":
		*membrane_value = Membrane_valueTabla
		return
	case "tambourine":
		*membrane_value = Membrane_valueTambourine
		return
	case "tenor drum":
		*membrane_value = Membrane_valueTenordrum
		return
	case "timbales":
		*membrane_value = Membrane_valueTimbales
		return
	case "tomtom":
		*membrane_value = Membrane_valueTomtom
		return
	default:
		return errUnkownEnum
	}
}

func (membrane_value *Membrane_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Membrane_valueBassdrum":
		*membrane_value = Membrane_valueBassdrum
	case "Membrane_valueBassdrumonside":
		*membrane_value = Membrane_valueBassdrumonside
	case "Membrane_valueBongos":
		*membrane_value = Membrane_valueBongos
	case "Membrane_valueChinesetomtom":
		*membrane_value = Membrane_valueChinesetomtom
	case "Membrane_valueCongadrum":
		*membrane_value = Membrane_valueCongadrum
	case "Membrane_valueCuica":
		*membrane_value = Membrane_valueCuica
	case "Membrane_valueGobletdrum":
		*membrane_value = Membrane_valueGobletdrum
	case "Membrane_valueIndo_Americantomtom":
		*membrane_value = Membrane_valueIndo_Americantomtom
	case "Membrane_valueJapanesetomtom":
		*membrane_value = Membrane_valueJapanesetomtom
	case "Membrane_valueMilitarydrum":
		*membrane_value = Membrane_valueMilitarydrum
	case "Membrane_valueSnaredrum":
		*membrane_value = Membrane_valueSnaredrum
	case "Membrane_valueSnaredrumsnaresoff":
		*membrane_value = Membrane_valueSnaredrumsnaresoff
	case "Membrane_valueTabla":
		*membrane_value = Membrane_valueTabla
	case "Membrane_valueTambourine":
		*membrane_value = Membrane_valueTambourine
	case "Membrane_valueTenordrum":
		*membrane_value = Membrane_valueTenordrum
	case "Membrane_valueTimbales":
		*membrane_value = Membrane_valueTimbales
	case "Membrane_valueTomtom":
		*membrane_value = Membrane_valueTomtom
	default:
		return errUnkownEnum
	}
	return
}

func (membrane_value *Membrane_value) ToCodeString() (res string) {

	switch *membrane_value {
	// insertion code per enum code
	case Membrane_valueBassdrum:
		res = "Membrane_valueBassdrum"
	case Membrane_valueBassdrumonside:
		res = "Membrane_valueBassdrumonside"
	case Membrane_valueBongos:
		res = "Membrane_valueBongos"
	case Membrane_valueChinesetomtom:
		res = "Membrane_valueChinesetomtom"
	case Membrane_valueCongadrum:
		res = "Membrane_valueCongadrum"
	case Membrane_valueCuica:
		res = "Membrane_valueCuica"
	case Membrane_valueGobletdrum:
		res = "Membrane_valueGobletdrum"
	case Membrane_valueIndo_Americantomtom:
		res = "Membrane_valueIndo_Americantomtom"
	case Membrane_valueJapanesetomtom:
		res = "Membrane_valueJapanesetomtom"
	case Membrane_valueMilitarydrum:
		res = "Membrane_valueMilitarydrum"
	case Membrane_valueSnaredrum:
		res = "Membrane_valueSnaredrum"
	case Membrane_valueSnaredrumsnaresoff:
		res = "Membrane_valueSnaredrumsnaresoff"
	case Membrane_valueTabla:
		res = "Membrane_valueTabla"
	case Membrane_valueTambourine:
		res = "Membrane_valueTambourine"
	case Membrane_valueTenordrum:
		res = "Membrane_valueTenordrum"
	case Membrane_valueTimbales:
		res = "Membrane_valueTimbales"
	case Membrane_valueTomtom:
		res = "Membrane_valueTomtom"
	}
	return
}

func (membrane_value Membrane_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Membrane_valueBassdrum")
	res = append(res, "Membrane_valueBassdrumonside")
	res = append(res, "Membrane_valueBongos")
	res = append(res, "Membrane_valueChinesetomtom")
	res = append(res, "Membrane_valueCongadrum")
	res = append(res, "Membrane_valueCuica")
	res = append(res, "Membrane_valueGobletdrum")
	res = append(res, "Membrane_valueIndo_Americantomtom")
	res = append(res, "Membrane_valueJapanesetomtom")
	res = append(res, "Membrane_valueMilitarydrum")
	res = append(res, "Membrane_valueSnaredrum")
	res = append(res, "Membrane_valueSnaredrumsnaresoff")
	res = append(res, "Membrane_valueTabla")
	res = append(res, "Membrane_valueTambourine")
	res = append(res, "Membrane_valueTenordrum")
	res = append(res, "Membrane_valueTimbales")
	res = append(res, "Membrane_valueTomtom")

	return
}

func (membrane_value Membrane_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bass drum")
	res = append(res, "bass drum on side")
	res = append(res, "bongos")
	res = append(res, "Chinese tomtom")
	res = append(res, "conga drum")
	res = append(res, "cuica")
	res = append(res, "goblet drum")
	res = append(res, "Indo-American tomtom")
	res = append(res, "Japanese tomtom")
	res = append(res, "military drum")
	res = append(res, "snare drum")
	res = append(res, "snare drum snares off")
	res = append(res, "tabla")
	res = append(res, "tambourine")
	res = append(res, "tenor drum")
	res = append(res, "timbales")
	res = append(res, "tomtom")

	return
}

// Utility function for Metal_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (metal_value Metal_value) ToString() (res string) {

	// migration of former implementation of enum
	switch metal_value {
	// insertion code per enum code
	case Metal_valueAgogo:
		res = "agogo"
	case Metal_valueAlmglocken:
		res = "almglocken"
	case Metal_valueBell:
		res = "bell"
	case Metal_valueBellplate:
		res = "bell plate"
	case Metal_valueBelltree:
		res = "bell tree"
	case Metal_valueBrakedrum:
		res = "brake drum"
	case Metal_valueCencerro:
		res = "cencerro"
	case Metal_valueChainrattle:
		res = "chain rattle"
	case Metal_valueChinesecymbal:
		res = "Chinese cymbal"
	case Metal_valueCowbell:
		res = "cowbell"
	case Metal_valueCrashcymbals:
		res = "crash cymbals"
	case Metal_valueCrotale:
		res = "crotale"
	case Metal_valueCymbaltongs:
		res = "cymbal tongs"
	case Metal_valueDomedgong:
		res = "domed gong"
	case Metal_valueFingercymbals:
		res = "finger cymbals"
	case Metal_valueFlexatone:
		res = "flexatone"
	case Metal_valueGong:
		res = "gong"
	case Metal_valueHi_hat:
		res = "hi-hat"
	case Metal_valueHigh_hatcymbals:
		res = "high-hat cymbals"
	case Metal_valueHandbell:
		res = "handbell"
	case Metal_valueJawharp:
		res = "jaw harp"
	case Metal_valueJinglebells:
		res = "jingle bells"
	case Metal_valueMusicalsaw:
		res = "musical saw"
	case Metal_valueShellbells:
		res = "shell bells"
	case Metal_valueSistrum:
		res = "sistrum"
	case Metal_valueSizzlecymbal:
		res = "sizzle cymbal"
	case Metal_valueSleighbells:
		res = "sleigh bells"
	case Metal_valueSuspendedcymbal:
		res = "suspended cymbal"
	case Metal_valueTamtam:
		res = "tam tam"
	case Metal_valueTamtamwithbeater:
		res = "tam tam with beater"
	case Metal_valueTriangle:
		res = "triangle"
	case Metal_valueVietnamesehat:
		res = "Vietnamese hat"
	}
	return
}

func (metal_value *Metal_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "agogo":
		*metal_value = Metal_valueAgogo
		return
	case "almglocken":
		*metal_value = Metal_valueAlmglocken
		return
	case "bell":
		*metal_value = Metal_valueBell
		return
	case "bell plate":
		*metal_value = Metal_valueBellplate
		return
	case "bell tree":
		*metal_value = Metal_valueBelltree
		return
	case "brake drum":
		*metal_value = Metal_valueBrakedrum
		return
	case "cencerro":
		*metal_value = Metal_valueCencerro
		return
	case "chain rattle":
		*metal_value = Metal_valueChainrattle
		return
	case "Chinese cymbal":
		*metal_value = Metal_valueChinesecymbal
		return
	case "cowbell":
		*metal_value = Metal_valueCowbell
		return
	case "crash cymbals":
		*metal_value = Metal_valueCrashcymbals
		return
	case "crotale":
		*metal_value = Metal_valueCrotale
		return
	case "cymbal tongs":
		*metal_value = Metal_valueCymbaltongs
		return
	case "domed gong":
		*metal_value = Metal_valueDomedgong
		return
	case "finger cymbals":
		*metal_value = Metal_valueFingercymbals
		return
	case "flexatone":
		*metal_value = Metal_valueFlexatone
		return
	case "gong":
		*metal_value = Metal_valueGong
		return
	case "hi-hat":
		*metal_value = Metal_valueHi_hat
		return
	case "high-hat cymbals":
		*metal_value = Metal_valueHigh_hatcymbals
		return
	case "handbell":
		*metal_value = Metal_valueHandbell
		return
	case "jaw harp":
		*metal_value = Metal_valueJawharp
		return
	case "jingle bells":
		*metal_value = Metal_valueJinglebells
		return
	case "musical saw":
		*metal_value = Metal_valueMusicalsaw
		return
	case "shell bells":
		*metal_value = Metal_valueShellbells
		return
	case "sistrum":
		*metal_value = Metal_valueSistrum
		return
	case "sizzle cymbal":
		*metal_value = Metal_valueSizzlecymbal
		return
	case "sleigh bells":
		*metal_value = Metal_valueSleighbells
		return
	case "suspended cymbal":
		*metal_value = Metal_valueSuspendedcymbal
		return
	case "tam tam":
		*metal_value = Metal_valueTamtam
		return
	case "tam tam with beater":
		*metal_value = Metal_valueTamtamwithbeater
		return
	case "triangle":
		*metal_value = Metal_valueTriangle
		return
	case "Vietnamese hat":
		*metal_value = Metal_valueVietnamesehat
		return
	default:
		return errUnkownEnum
	}
}

func (metal_value *Metal_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Metal_valueAgogo":
		*metal_value = Metal_valueAgogo
	case "Metal_valueAlmglocken":
		*metal_value = Metal_valueAlmglocken
	case "Metal_valueBell":
		*metal_value = Metal_valueBell
	case "Metal_valueBellplate":
		*metal_value = Metal_valueBellplate
	case "Metal_valueBelltree":
		*metal_value = Metal_valueBelltree
	case "Metal_valueBrakedrum":
		*metal_value = Metal_valueBrakedrum
	case "Metal_valueCencerro":
		*metal_value = Metal_valueCencerro
	case "Metal_valueChainrattle":
		*metal_value = Metal_valueChainrattle
	case "Metal_valueChinesecymbal":
		*metal_value = Metal_valueChinesecymbal
	case "Metal_valueCowbell":
		*metal_value = Metal_valueCowbell
	case "Metal_valueCrashcymbals":
		*metal_value = Metal_valueCrashcymbals
	case "Metal_valueCrotale":
		*metal_value = Metal_valueCrotale
	case "Metal_valueCymbaltongs":
		*metal_value = Metal_valueCymbaltongs
	case "Metal_valueDomedgong":
		*metal_value = Metal_valueDomedgong
	case "Metal_valueFingercymbals":
		*metal_value = Metal_valueFingercymbals
	case "Metal_valueFlexatone":
		*metal_value = Metal_valueFlexatone
	case "Metal_valueGong":
		*metal_value = Metal_valueGong
	case "Metal_valueHi_hat":
		*metal_value = Metal_valueHi_hat
	case "Metal_valueHigh_hatcymbals":
		*metal_value = Metal_valueHigh_hatcymbals
	case "Metal_valueHandbell":
		*metal_value = Metal_valueHandbell
	case "Metal_valueJawharp":
		*metal_value = Metal_valueJawharp
	case "Metal_valueJinglebells":
		*metal_value = Metal_valueJinglebells
	case "Metal_valueMusicalsaw":
		*metal_value = Metal_valueMusicalsaw
	case "Metal_valueShellbells":
		*metal_value = Metal_valueShellbells
	case "Metal_valueSistrum":
		*metal_value = Metal_valueSistrum
	case "Metal_valueSizzlecymbal":
		*metal_value = Metal_valueSizzlecymbal
	case "Metal_valueSleighbells":
		*metal_value = Metal_valueSleighbells
	case "Metal_valueSuspendedcymbal":
		*metal_value = Metal_valueSuspendedcymbal
	case "Metal_valueTamtam":
		*metal_value = Metal_valueTamtam
	case "Metal_valueTamtamwithbeater":
		*metal_value = Metal_valueTamtamwithbeater
	case "Metal_valueTriangle":
		*metal_value = Metal_valueTriangle
	case "Metal_valueVietnamesehat":
		*metal_value = Metal_valueVietnamesehat
	default:
		return errUnkownEnum
	}
	return
}

func (metal_value *Metal_value) ToCodeString() (res string) {

	switch *metal_value {
	// insertion code per enum code
	case Metal_valueAgogo:
		res = "Metal_valueAgogo"
	case Metal_valueAlmglocken:
		res = "Metal_valueAlmglocken"
	case Metal_valueBell:
		res = "Metal_valueBell"
	case Metal_valueBellplate:
		res = "Metal_valueBellplate"
	case Metal_valueBelltree:
		res = "Metal_valueBelltree"
	case Metal_valueBrakedrum:
		res = "Metal_valueBrakedrum"
	case Metal_valueCencerro:
		res = "Metal_valueCencerro"
	case Metal_valueChainrattle:
		res = "Metal_valueChainrattle"
	case Metal_valueChinesecymbal:
		res = "Metal_valueChinesecymbal"
	case Metal_valueCowbell:
		res = "Metal_valueCowbell"
	case Metal_valueCrashcymbals:
		res = "Metal_valueCrashcymbals"
	case Metal_valueCrotale:
		res = "Metal_valueCrotale"
	case Metal_valueCymbaltongs:
		res = "Metal_valueCymbaltongs"
	case Metal_valueDomedgong:
		res = "Metal_valueDomedgong"
	case Metal_valueFingercymbals:
		res = "Metal_valueFingercymbals"
	case Metal_valueFlexatone:
		res = "Metal_valueFlexatone"
	case Metal_valueGong:
		res = "Metal_valueGong"
	case Metal_valueHi_hat:
		res = "Metal_valueHi_hat"
	case Metal_valueHigh_hatcymbals:
		res = "Metal_valueHigh_hatcymbals"
	case Metal_valueHandbell:
		res = "Metal_valueHandbell"
	case Metal_valueJawharp:
		res = "Metal_valueJawharp"
	case Metal_valueJinglebells:
		res = "Metal_valueJinglebells"
	case Metal_valueMusicalsaw:
		res = "Metal_valueMusicalsaw"
	case Metal_valueShellbells:
		res = "Metal_valueShellbells"
	case Metal_valueSistrum:
		res = "Metal_valueSistrum"
	case Metal_valueSizzlecymbal:
		res = "Metal_valueSizzlecymbal"
	case Metal_valueSleighbells:
		res = "Metal_valueSleighbells"
	case Metal_valueSuspendedcymbal:
		res = "Metal_valueSuspendedcymbal"
	case Metal_valueTamtam:
		res = "Metal_valueTamtam"
	case Metal_valueTamtamwithbeater:
		res = "Metal_valueTamtamwithbeater"
	case Metal_valueTriangle:
		res = "Metal_valueTriangle"
	case Metal_valueVietnamesehat:
		res = "Metal_valueVietnamesehat"
	}
	return
}

func (metal_value Metal_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Metal_valueAgogo")
	res = append(res, "Metal_valueAlmglocken")
	res = append(res, "Metal_valueBell")
	res = append(res, "Metal_valueBellplate")
	res = append(res, "Metal_valueBelltree")
	res = append(res, "Metal_valueBrakedrum")
	res = append(res, "Metal_valueCencerro")
	res = append(res, "Metal_valueChainrattle")
	res = append(res, "Metal_valueChinesecymbal")
	res = append(res, "Metal_valueCowbell")
	res = append(res, "Metal_valueCrashcymbals")
	res = append(res, "Metal_valueCrotale")
	res = append(res, "Metal_valueCymbaltongs")
	res = append(res, "Metal_valueDomedgong")
	res = append(res, "Metal_valueFingercymbals")
	res = append(res, "Metal_valueFlexatone")
	res = append(res, "Metal_valueGong")
	res = append(res, "Metal_valueHi_hat")
	res = append(res, "Metal_valueHigh_hatcymbals")
	res = append(res, "Metal_valueHandbell")
	res = append(res, "Metal_valueJawharp")
	res = append(res, "Metal_valueJinglebells")
	res = append(res, "Metal_valueMusicalsaw")
	res = append(res, "Metal_valueShellbells")
	res = append(res, "Metal_valueSistrum")
	res = append(res, "Metal_valueSizzlecymbal")
	res = append(res, "Metal_valueSleighbells")
	res = append(res, "Metal_valueSuspendedcymbal")
	res = append(res, "Metal_valueTamtam")
	res = append(res, "Metal_valueTamtamwithbeater")
	res = append(res, "Metal_valueTriangle")
	res = append(res, "Metal_valueVietnamesehat")

	return
}

func (metal_value Metal_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "agogo")
	res = append(res, "almglocken")
	res = append(res, "bell")
	res = append(res, "bell plate")
	res = append(res, "bell tree")
	res = append(res, "brake drum")
	res = append(res, "cencerro")
	res = append(res, "chain rattle")
	res = append(res, "Chinese cymbal")
	res = append(res, "cowbell")
	res = append(res, "crash cymbals")
	res = append(res, "crotale")
	res = append(res, "cymbal tongs")
	res = append(res, "domed gong")
	res = append(res, "finger cymbals")
	res = append(res, "flexatone")
	res = append(res, "gong")
	res = append(res, "hi-hat")
	res = append(res, "high-hat cymbals")
	res = append(res, "handbell")
	res = append(res, "jaw harp")
	res = append(res, "jingle bells")
	res = append(res, "musical saw")
	res = append(res, "shell bells")
	res = append(res, "sistrum")
	res = append(res, "sizzle cymbal")
	res = append(res, "sleigh bells")
	res = append(res, "suspended cymbal")
	res = append(res, "tam tam")
	res = append(res, "tam tam with beater")
	res = append(res, "triangle")
	res = append(res, "Vietnamese hat")

	return
}

// Utility function for Mode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (mode Mode) ToString() (res string) {

	// migration of former implementation of enum
	switch mode {
	// insertion code per enum code
	}
	return
}

func (mode *Mode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (mode *Mode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (mode *Mode) ToCodeString() (res string) {

	switch *mode {
	// insertion code per enum code
	}
	return
}

func (mode Mode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (mode Mode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Mute
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (mute Mute) ToString() (res string) {

	// migration of former implementation of enum
	switch mute {
	// insertion code per enum code
	case MuteOn:
		res = "on"
	case MuteOff:
		res = "off"
	case MuteStraight:
		res = "straight"
	case MuteCup:
		res = "cup"
	case MuteHarmon_no_stem:
		res = "harmon-no-stem"
	case MuteHarmon_stem:
		res = "harmon-stem"
	case MuteBucket:
		res = "bucket"
	case MutePlunger:
		res = "plunger"
	case MuteHat:
		res = "hat"
	case MuteSolotone:
		res = "solotone"
	case MutePractice:
		res = "practice"
	case MuteStop_mute:
		res = "stop-mute"
	case MuteStop_hand:
		res = "stop-hand"
	case MuteEcho:
		res = "echo"
	case MutePalm:
		res = "palm"
	}
	return
}

func (mute *Mute) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "on":
		*mute = MuteOn
		return
	case "off":
		*mute = MuteOff
		return
	case "straight":
		*mute = MuteStraight
		return
	case "cup":
		*mute = MuteCup
		return
	case "harmon-no-stem":
		*mute = MuteHarmon_no_stem
		return
	case "harmon-stem":
		*mute = MuteHarmon_stem
		return
	case "bucket":
		*mute = MuteBucket
		return
	case "plunger":
		*mute = MutePlunger
		return
	case "hat":
		*mute = MuteHat
		return
	case "solotone":
		*mute = MuteSolotone
		return
	case "practice":
		*mute = MutePractice
		return
	case "stop-mute":
		*mute = MuteStop_mute
		return
	case "stop-hand":
		*mute = MuteStop_hand
		return
	case "echo":
		*mute = MuteEcho
		return
	case "palm":
		*mute = MutePalm
		return
	default:
		return errUnkownEnum
	}
}

func (mute *Mute) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "MuteOn":
		*mute = MuteOn
	case "MuteOff":
		*mute = MuteOff
	case "MuteStraight":
		*mute = MuteStraight
	case "MuteCup":
		*mute = MuteCup
	case "MuteHarmon_no_stem":
		*mute = MuteHarmon_no_stem
	case "MuteHarmon_stem":
		*mute = MuteHarmon_stem
	case "MuteBucket":
		*mute = MuteBucket
	case "MutePlunger":
		*mute = MutePlunger
	case "MuteHat":
		*mute = MuteHat
	case "MuteSolotone":
		*mute = MuteSolotone
	case "MutePractice":
		*mute = MutePractice
	case "MuteStop_mute":
		*mute = MuteStop_mute
	case "MuteStop_hand":
		*mute = MuteStop_hand
	case "MuteEcho":
		*mute = MuteEcho
	case "MutePalm":
		*mute = MutePalm
	default:
		return errUnkownEnum
	}
	return
}

func (mute *Mute) ToCodeString() (res string) {

	switch *mute {
	// insertion code per enum code
	case MuteOn:
		res = "MuteOn"
	case MuteOff:
		res = "MuteOff"
	case MuteStraight:
		res = "MuteStraight"
	case MuteCup:
		res = "MuteCup"
	case MuteHarmon_no_stem:
		res = "MuteHarmon_no_stem"
	case MuteHarmon_stem:
		res = "MuteHarmon_stem"
	case MuteBucket:
		res = "MuteBucket"
	case MutePlunger:
		res = "MutePlunger"
	case MuteHat:
		res = "MuteHat"
	case MuteSolotone:
		res = "MuteSolotone"
	case MutePractice:
		res = "MutePractice"
	case MuteStop_mute:
		res = "MuteStop_mute"
	case MuteStop_hand:
		res = "MuteStop_hand"
	case MuteEcho:
		res = "MuteEcho"
	case MutePalm:
		res = "MutePalm"
	}
	return
}

func (mute Mute) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "MuteOn")
	res = append(res, "MuteOff")
	res = append(res, "MuteStraight")
	res = append(res, "MuteCup")
	res = append(res, "MuteHarmon_no_stem")
	res = append(res, "MuteHarmon_stem")
	res = append(res, "MuteBucket")
	res = append(res, "MutePlunger")
	res = append(res, "MuteHat")
	res = append(res, "MuteSolotone")
	res = append(res, "MutePractice")
	res = append(res, "MuteStop_mute")
	res = append(res, "MuteStop_hand")
	res = append(res, "MuteEcho")
	res = append(res, "MutePalm")

	return
}

func (mute Mute) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "on")
	res = append(res, "off")
	res = append(res, "straight")
	res = append(res, "cup")
	res = append(res, "harmon-no-stem")
	res = append(res, "harmon-stem")
	res = append(res, "bucket")
	res = append(res, "plunger")
	res = append(res, "hat")
	res = append(res, "solotone")
	res = append(res, "practice")
	res = append(res, "stop-mute")
	res = append(res, "stop-hand")
	res = append(res, "echo")
	res = append(res, "palm")

	return
}

// Utility function for NCName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (ncname NCName) ToString() (res string) {

	// migration of former implementation of enum
	switch ncname {
	// insertion code per enum code
	}
	return
}

func (ncname *NCName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (ncname *NCName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (ncname *NCName) ToCodeString() (res string) {

	switch *ncname {
	// insertion code per enum code
	}
	return
}

func (ncname NCName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (ncname NCName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for NMTOKEN
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (nmtoken NMTOKEN) ToString() (res string) {

	// migration of former implementation of enum
	switch nmtoken {
	// insertion code per enum code
	}
	return
}

func (nmtoken *NMTOKEN) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (nmtoken *NMTOKEN) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (nmtoken *NMTOKEN) ToCodeString() (res string) {

	switch *nmtoken {
	// insertion code per enum code
	}
	return
}

func (nmtoken NMTOKEN) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (nmtoken NMTOKEN) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Note_size_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (note_size_type Note_size_type) ToString() (res string) {

	// migration of former implementation of enum
	switch note_size_type {
	// insertion code per enum code
	case Note_size_typeCue:
		res = "cue"
	case Note_size_typeGrace:
		res = "grace"
	case Note_size_typeGrace_cue:
		res = "grace-cue"
	case Note_size_typeLarge:
		res = "large"
	}
	return
}

func (note_size_type *Note_size_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "cue":
		*note_size_type = Note_size_typeCue
		return
	case "grace":
		*note_size_type = Note_size_typeGrace
		return
	case "grace-cue":
		*note_size_type = Note_size_typeGrace_cue
		return
	case "large":
		*note_size_type = Note_size_typeLarge
		return
	default:
		return errUnkownEnum
	}
}

func (note_size_type *Note_size_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Note_size_typeCue":
		*note_size_type = Note_size_typeCue
	case "Note_size_typeGrace":
		*note_size_type = Note_size_typeGrace
	case "Note_size_typeGrace_cue":
		*note_size_type = Note_size_typeGrace_cue
	case "Note_size_typeLarge":
		*note_size_type = Note_size_typeLarge
	default:
		return errUnkownEnum
	}
	return
}

func (note_size_type *Note_size_type) ToCodeString() (res string) {

	switch *note_size_type {
	// insertion code per enum code
	case Note_size_typeCue:
		res = "Note_size_typeCue"
	case Note_size_typeGrace:
		res = "Note_size_typeGrace"
	case Note_size_typeGrace_cue:
		res = "Note_size_typeGrace_cue"
	case Note_size_typeLarge:
		res = "Note_size_typeLarge"
	}
	return
}

func (note_size_type Note_size_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Note_size_typeCue")
	res = append(res, "Note_size_typeGrace")
	res = append(res, "Note_size_typeGrace_cue")
	res = append(res, "Note_size_typeLarge")

	return
}

func (note_size_type Note_size_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "cue")
	res = append(res, "grace")
	res = append(res, "grace-cue")
	res = append(res, "large")

	return
}

// Utility function for Note_type_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (note_type_value Note_type_value) ToString() (res string) {

	// migration of former implementation of enum
	switch note_type_value {
	// insertion code per enum code
	case Note_type_value1024th:
		res = "1024th"
	case Note_type_value512th:
		res = "512th"
	case Note_type_value256th:
		res = "256th"
	case Note_type_value128th:
		res = "128th"
	case Note_type_value64th:
		res = "64th"
	case Note_type_value32nd:
		res = "32nd"
	case Note_type_value16th:
		res = "16th"
	case Note_type_valueEighth:
		res = "eighth"
	case Note_type_valueQuarter:
		res = "quarter"
	case Note_type_valueHalf:
		res = "half"
	case Note_type_valueWhole:
		res = "whole"
	case Note_type_valueBreve:
		res = "breve"
	case Note_type_valueLong:
		res = "long"
	case Note_type_valueMaxima:
		res = "maxima"
	}
	return
}

func (note_type_value *Note_type_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "1024th":
		*note_type_value = Note_type_value1024th
		return
	case "512th":
		*note_type_value = Note_type_value512th
		return
	case "256th":
		*note_type_value = Note_type_value256th
		return
	case "128th":
		*note_type_value = Note_type_value128th
		return
	case "64th":
		*note_type_value = Note_type_value64th
		return
	case "32nd":
		*note_type_value = Note_type_value32nd
		return
	case "16th":
		*note_type_value = Note_type_value16th
		return
	case "eighth":
		*note_type_value = Note_type_valueEighth
		return
	case "quarter":
		*note_type_value = Note_type_valueQuarter
		return
	case "half":
		*note_type_value = Note_type_valueHalf
		return
	case "whole":
		*note_type_value = Note_type_valueWhole
		return
	case "breve":
		*note_type_value = Note_type_valueBreve
		return
	case "long":
		*note_type_value = Note_type_valueLong
		return
	case "maxima":
		*note_type_value = Note_type_valueMaxima
		return
	default:
		return errUnkownEnum
	}
}

func (note_type_value *Note_type_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Note_type_value1024th":
		*note_type_value = Note_type_value1024th
	case "Note_type_value512th":
		*note_type_value = Note_type_value512th
	case "Note_type_value256th":
		*note_type_value = Note_type_value256th
	case "Note_type_value128th":
		*note_type_value = Note_type_value128th
	case "Note_type_value64th":
		*note_type_value = Note_type_value64th
	case "Note_type_value32nd":
		*note_type_value = Note_type_value32nd
	case "Note_type_value16th":
		*note_type_value = Note_type_value16th
	case "Note_type_valueEighth":
		*note_type_value = Note_type_valueEighth
	case "Note_type_valueQuarter":
		*note_type_value = Note_type_valueQuarter
	case "Note_type_valueHalf":
		*note_type_value = Note_type_valueHalf
	case "Note_type_valueWhole":
		*note_type_value = Note_type_valueWhole
	case "Note_type_valueBreve":
		*note_type_value = Note_type_valueBreve
	case "Note_type_valueLong":
		*note_type_value = Note_type_valueLong
	case "Note_type_valueMaxima":
		*note_type_value = Note_type_valueMaxima
	default:
		return errUnkownEnum
	}
	return
}

func (note_type_value *Note_type_value) ToCodeString() (res string) {

	switch *note_type_value {
	// insertion code per enum code
	case Note_type_value1024th:
		res = "Note_type_value1024th"
	case Note_type_value512th:
		res = "Note_type_value512th"
	case Note_type_value256th:
		res = "Note_type_value256th"
	case Note_type_value128th:
		res = "Note_type_value128th"
	case Note_type_value64th:
		res = "Note_type_value64th"
	case Note_type_value32nd:
		res = "Note_type_value32nd"
	case Note_type_value16th:
		res = "Note_type_value16th"
	case Note_type_valueEighth:
		res = "Note_type_valueEighth"
	case Note_type_valueQuarter:
		res = "Note_type_valueQuarter"
	case Note_type_valueHalf:
		res = "Note_type_valueHalf"
	case Note_type_valueWhole:
		res = "Note_type_valueWhole"
	case Note_type_valueBreve:
		res = "Note_type_valueBreve"
	case Note_type_valueLong:
		res = "Note_type_valueLong"
	case Note_type_valueMaxima:
		res = "Note_type_valueMaxima"
	}
	return
}

func (note_type_value Note_type_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Note_type_value1024th")
	res = append(res, "Note_type_value512th")
	res = append(res, "Note_type_value256th")
	res = append(res, "Note_type_value128th")
	res = append(res, "Note_type_value64th")
	res = append(res, "Note_type_value32nd")
	res = append(res, "Note_type_value16th")
	res = append(res, "Note_type_valueEighth")
	res = append(res, "Note_type_valueQuarter")
	res = append(res, "Note_type_valueHalf")
	res = append(res, "Note_type_valueWhole")
	res = append(res, "Note_type_valueBreve")
	res = append(res, "Note_type_valueLong")
	res = append(res, "Note_type_valueMaxima")

	return
}

func (note_type_value Note_type_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "1024th")
	res = append(res, "512th")
	res = append(res, "256th")
	res = append(res, "128th")
	res = append(res, "64th")
	res = append(res, "32nd")
	res = append(res, "16th")
	res = append(res, "eighth")
	res = append(res, "quarter")
	res = append(res, "half")
	res = append(res, "whole")
	res = append(res, "breve")
	res = append(res, "long")
	res = append(res, "maxima")

	return
}

// Utility function for Notehead_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (notehead_value Notehead_value) ToString() (res string) {

	// migration of former implementation of enum
	switch notehead_value {
	// insertion code per enum code
	case Notehead_valueSlash:
		res = "slash"
	case Notehead_valueTriangle:
		res = "triangle"
	case Notehead_valueDiamond:
		res = "diamond"
	case Notehead_valueSquare:
		res = "square"
	case Notehead_valueCross:
		res = "cross"
	case Notehead_valueX:
		res = "x"
	case Notehead_valueCircle_x:
		res = "circle-x"
	case Notehead_valueInvertedtriangle:
		res = "inverted triangle"
	case Notehead_valueArrowdown:
		res = "arrow down"
	case Notehead_valueArrowup:
		res = "arrow up"
	case Notehead_valueCircled:
		res = "circled"
	case Notehead_valueSlashed:
		res = "slashed"
	case Notehead_valueBackslashed:
		res = "back slashed"
	case Notehead_valueNormal:
		res = "normal"
	case Notehead_valueCluster:
		res = "cluster"
	case Notehead_valueCircledot:
		res = "circle dot"
	case Notehead_valueLefttriangle:
		res = "left triangle"
	case Notehead_valueRectangle:
		res = "rectangle"
	case Notehead_valueNone:
		res = "none"
	case Notehead_valueDo:
		res = "do"
	case Notehead_valueRe:
		res = "re"
	case Notehead_valueMi:
		res = "mi"
	case Notehead_valueFa:
		res = "fa"
	case Notehead_valueFaup:
		res = "fa up"
	case Notehead_valueSo:
		res = "so"
	case Notehead_valueLa:
		res = "la"
	case Notehead_valueTi:
		res = "ti"
	case Notehead_valueOther:
		res = "other"
	}
	return
}

func (notehead_value *Notehead_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "slash":
		*notehead_value = Notehead_valueSlash
		return
	case "triangle":
		*notehead_value = Notehead_valueTriangle
		return
	case "diamond":
		*notehead_value = Notehead_valueDiamond
		return
	case "square":
		*notehead_value = Notehead_valueSquare
		return
	case "cross":
		*notehead_value = Notehead_valueCross
		return
	case "x":
		*notehead_value = Notehead_valueX
		return
	case "circle-x":
		*notehead_value = Notehead_valueCircle_x
		return
	case "inverted triangle":
		*notehead_value = Notehead_valueInvertedtriangle
		return
	case "arrow down":
		*notehead_value = Notehead_valueArrowdown
		return
	case "arrow up":
		*notehead_value = Notehead_valueArrowup
		return
	case "circled":
		*notehead_value = Notehead_valueCircled
		return
	case "slashed":
		*notehead_value = Notehead_valueSlashed
		return
	case "back slashed":
		*notehead_value = Notehead_valueBackslashed
		return
	case "normal":
		*notehead_value = Notehead_valueNormal
		return
	case "cluster":
		*notehead_value = Notehead_valueCluster
		return
	case "circle dot":
		*notehead_value = Notehead_valueCircledot
		return
	case "left triangle":
		*notehead_value = Notehead_valueLefttriangle
		return
	case "rectangle":
		*notehead_value = Notehead_valueRectangle
		return
	case "none":
		*notehead_value = Notehead_valueNone
		return
	case "do":
		*notehead_value = Notehead_valueDo
		return
	case "re":
		*notehead_value = Notehead_valueRe
		return
	case "mi":
		*notehead_value = Notehead_valueMi
		return
	case "fa":
		*notehead_value = Notehead_valueFa
		return
	case "fa up":
		*notehead_value = Notehead_valueFaup
		return
	case "so":
		*notehead_value = Notehead_valueSo
		return
	case "la":
		*notehead_value = Notehead_valueLa
		return
	case "ti":
		*notehead_value = Notehead_valueTi
		return
	case "other":
		*notehead_value = Notehead_valueOther
		return
	default:
		return errUnkownEnum
	}
}

func (notehead_value *Notehead_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Notehead_valueSlash":
		*notehead_value = Notehead_valueSlash
	case "Notehead_valueTriangle":
		*notehead_value = Notehead_valueTriangle
	case "Notehead_valueDiamond":
		*notehead_value = Notehead_valueDiamond
	case "Notehead_valueSquare":
		*notehead_value = Notehead_valueSquare
	case "Notehead_valueCross":
		*notehead_value = Notehead_valueCross
	case "Notehead_valueX":
		*notehead_value = Notehead_valueX
	case "Notehead_valueCircle_x":
		*notehead_value = Notehead_valueCircle_x
	case "Notehead_valueInvertedtriangle":
		*notehead_value = Notehead_valueInvertedtriangle
	case "Notehead_valueArrowdown":
		*notehead_value = Notehead_valueArrowdown
	case "Notehead_valueArrowup":
		*notehead_value = Notehead_valueArrowup
	case "Notehead_valueCircled":
		*notehead_value = Notehead_valueCircled
	case "Notehead_valueSlashed":
		*notehead_value = Notehead_valueSlashed
	case "Notehead_valueBackslashed":
		*notehead_value = Notehead_valueBackslashed
	case "Notehead_valueNormal":
		*notehead_value = Notehead_valueNormal
	case "Notehead_valueCluster":
		*notehead_value = Notehead_valueCluster
	case "Notehead_valueCircledot":
		*notehead_value = Notehead_valueCircledot
	case "Notehead_valueLefttriangle":
		*notehead_value = Notehead_valueLefttriangle
	case "Notehead_valueRectangle":
		*notehead_value = Notehead_valueRectangle
	case "Notehead_valueNone":
		*notehead_value = Notehead_valueNone
	case "Notehead_valueDo":
		*notehead_value = Notehead_valueDo
	case "Notehead_valueRe":
		*notehead_value = Notehead_valueRe
	case "Notehead_valueMi":
		*notehead_value = Notehead_valueMi
	case "Notehead_valueFa":
		*notehead_value = Notehead_valueFa
	case "Notehead_valueFaup":
		*notehead_value = Notehead_valueFaup
	case "Notehead_valueSo":
		*notehead_value = Notehead_valueSo
	case "Notehead_valueLa":
		*notehead_value = Notehead_valueLa
	case "Notehead_valueTi":
		*notehead_value = Notehead_valueTi
	case "Notehead_valueOther":
		*notehead_value = Notehead_valueOther
	default:
		return errUnkownEnum
	}
	return
}

func (notehead_value *Notehead_value) ToCodeString() (res string) {

	switch *notehead_value {
	// insertion code per enum code
	case Notehead_valueSlash:
		res = "Notehead_valueSlash"
	case Notehead_valueTriangle:
		res = "Notehead_valueTriangle"
	case Notehead_valueDiamond:
		res = "Notehead_valueDiamond"
	case Notehead_valueSquare:
		res = "Notehead_valueSquare"
	case Notehead_valueCross:
		res = "Notehead_valueCross"
	case Notehead_valueX:
		res = "Notehead_valueX"
	case Notehead_valueCircle_x:
		res = "Notehead_valueCircle_x"
	case Notehead_valueInvertedtriangle:
		res = "Notehead_valueInvertedtriangle"
	case Notehead_valueArrowdown:
		res = "Notehead_valueArrowdown"
	case Notehead_valueArrowup:
		res = "Notehead_valueArrowup"
	case Notehead_valueCircled:
		res = "Notehead_valueCircled"
	case Notehead_valueSlashed:
		res = "Notehead_valueSlashed"
	case Notehead_valueBackslashed:
		res = "Notehead_valueBackslashed"
	case Notehead_valueNormal:
		res = "Notehead_valueNormal"
	case Notehead_valueCluster:
		res = "Notehead_valueCluster"
	case Notehead_valueCircledot:
		res = "Notehead_valueCircledot"
	case Notehead_valueLefttriangle:
		res = "Notehead_valueLefttriangle"
	case Notehead_valueRectangle:
		res = "Notehead_valueRectangle"
	case Notehead_valueNone:
		res = "Notehead_valueNone"
	case Notehead_valueDo:
		res = "Notehead_valueDo"
	case Notehead_valueRe:
		res = "Notehead_valueRe"
	case Notehead_valueMi:
		res = "Notehead_valueMi"
	case Notehead_valueFa:
		res = "Notehead_valueFa"
	case Notehead_valueFaup:
		res = "Notehead_valueFaup"
	case Notehead_valueSo:
		res = "Notehead_valueSo"
	case Notehead_valueLa:
		res = "Notehead_valueLa"
	case Notehead_valueTi:
		res = "Notehead_valueTi"
	case Notehead_valueOther:
		res = "Notehead_valueOther"
	}
	return
}

func (notehead_value Notehead_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Notehead_valueSlash")
	res = append(res, "Notehead_valueTriangle")
	res = append(res, "Notehead_valueDiamond")
	res = append(res, "Notehead_valueSquare")
	res = append(res, "Notehead_valueCross")
	res = append(res, "Notehead_valueX")
	res = append(res, "Notehead_valueCircle_x")
	res = append(res, "Notehead_valueInvertedtriangle")
	res = append(res, "Notehead_valueArrowdown")
	res = append(res, "Notehead_valueArrowup")
	res = append(res, "Notehead_valueCircled")
	res = append(res, "Notehead_valueSlashed")
	res = append(res, "Notehead_valueBackslashed")
	res = append(res, "Notehead_valueNormal")
	res = append(res, "Notehead_valueCluster")
	res = append(res, "Notehead_valueCircledot")
	res = append(res, "Notehead_valueLefttriangle")
	res = append(res, "Notehead_valueRectangle")
	res = append(res, "Notehead_valueNone")
	res = append(res, "Notehead_valueDo")
	res = append(res, "Notehead_valueRe")
	res = append(res, "Notehead_valueMi")
	res = append(res, "Notehead_valueFa")
	res = append(res, "Notehead_valueFaup")
	res = append(res, "Notehead_valueSo")
	res = append(res, "Notehead_valueLa")
	res = append(res, "Notehead_valueTi")
	res = append(res, "Notehead_valueOther")

	return
}

func (notehead_value Notehead_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "slash")
	res = append(res, "triangle")
	res = append(res, "diamond")
	res = append(res, "square")
	res = append(res, "cross")
	res = append(res, "x")
	res = append(res, "circle-x")
	res = append(res, "inverted triangle")
	res = append(res, "arrow down")
	res = append(res, "arrow up")
	res = append(res, "circled")
	res = append(res, "slashed")
	res = append(res, "back slashed")
	res = append(res, "normal")
	res = append(res, "cluster")
	res = append(res, "circle dot")
	res = append(res, "left triangle")
	res = append(res, "rectangle")
	res = append(res, "none")
	res = append(res, "do")
	res = append(res, "re")
	res = append(res, "mi")
	res = append(res, "fa")
	res = append(res, "fa up")
	res = append(res, "so")
	res = append(res, "la")
	res = append(res, "ti")
	res = append(res, "other")

	return
}

// Utility function for Number_or_normal
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (number_or_normal Number_or_normal) ToString() (res string) {

	// migration of former implementation of enum
	switch number_or_normal {
	// insertion code per enum code
	}
	return
}

func (number_or_normal *Number_or_normal) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (number_or_normal *Number_or_normal) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (number_or_normal *Number_or_normal) ToCodeString() (res string) {

	switch *number_or_normal {
	// insertion code per enum code
	}
	return
}

func (number_or_normal Number_or_normal) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (number_or_normal Number_or_normal) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Numeral_mode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (numeral_mode Numeral_mode) ToString() (res string) {

	// migration of former implementation of enum
	switch numeral_mode {
	// insertion code per enum code
	case Numeral_modeMajor:
		res = "major"
	case Numeral_modeMinor:
		res = "minor"
	case Numeral_modeNaturalminor:
		res = "natural minor"
	case Numeral_modeMelodicminor:
		res = "melodic minor"
	case Numeral_modeHarmonicminor:
		res = "harmonic minor"
	}
	return
}

func (numeral_mode *Numeral_mode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "major":
		*numeral_mode = Numeral_modeMajor
		return
	case "minor":
		*numeral_mode = Numeral_modeMinor
		return
	case "natural minor":
		*numeral_mode = Numeral_modeNaturalminor
		return
	case "melodic minor":
		*numeral_mode = Numeral_modeMelodicminor
		return
	case "harmonic minor":
		*numeral_mode = Numeral_modeHarmonicminor
		return
	default:
		return errUnkownEnum
	}
}

func (numeral_mode *Numeral_mode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Numeral_modeMajor":
		*numeral_mode = Numeral_modeMajor
	case "Numeral_modeMinor":
		*numeral_mode = Numeral_modeMinor
	case "Numeral_modeNaturalminor":
		*numeral_mode = Numeral_modeNaturalminor
	case "Numeral_modeMelodicminor":
		*numeral_mode = Numeral_modeMelodicminor
	case "Numeral_modeHarmonicminor":
		*numeral_mode = Numeral_modeHarmonicminor
	default:
		return errUnkownEnum
	}
	return
}

func (numeral_mode *Numeral_mode) ToCodeString() (res string) {

	switch *numeral_mode {
	// insertion code per enum code
	case Numeral_modeMajor:
		res = "Numeral_modeMajor"
	case Numeral_modeMinor:
		res = "Numeral_modeMinor"
	case Numeral_modeNaturalminor:
		res = "Numeral_modeNaturalminor"
	case Numeral_modeMelodicminor:
		res = "Numeral_modeMelodicminor"
	case Numeral_modeHarmonicminor:
		res = "Numeral_modeHarmonicminor"
	}
	return
}

func (numeral_mode Numeral_mode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Numeral_modeMajor")
	res = append(res, "Numeral_modeMinor")
	res = append(res, "Numeral_modeNaturalminor")
	res = append(res, "Numeral_modeMelodicminor")
	res = append(res, "Numeral_modeHarmonicminor")

	return
}

func (numeral_mode Numeral_mode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "major")
	res = append(res, "minor")
	res = append(res, "natural minor")
	res = append(res, "melodic minor")
	res = append(res, "harmonic minor")

	return
}

// Utility function for On_off
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (on_off On_off) ToString() (res string) {

	// migration of former implementation of enum
	switch on_off {
	// insertion code per enum code
	case On_offOn:
		res = "on"
	case On_offOff:
		res = "off"
	}
	return
}

func (on_off *On_off) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "on":
		*on_off = On_offOn
		return
	case "off":
		*on_off = On_offOff
		return
	default:
		return errUnkownEnum
	}
}

func (on_off *On_off) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "On_offOn":
		*on_off = On_offOn
	case "On_offOff":
		*on_off = On_offOff
	default:
		return errUnkownEnum
	}
	return
}

func (on_off *On_off) ToCodeString() (res string) {

	switch *on_off {
	// insertion code per enum code
	case On_offOn:
		res = "On_offOn"
	case On_offOff:
		res = "On_offOff"
	}
	return
}

func (on_off On_off) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "On_offOn")
	res = append(res, "On_offOff")

	return
}

func (on_off On_off) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "on")
	res = append(res, "off")

	return
}

// Utility function for Over_under
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (over_under Over_under) ToString() (res string) {

	// migration of former implementation of enum
	switch over_under {
	// insertion code per enum code
	case Over_underOver:
		res = "over"
	case Over_underUnder:
		res = "under"
	}
	return
}

func (over_under *Over_under) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "over":
		*over_under = Over_underOver
		return
	case "under":
		*over_under = Over_underUnder
		return
	default:
		return errUnkownEnum
	}
}

func (over_under *Over_under) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Over_underOver":
		*over_under = Over_underOver
	case "Over_underUnder":
		*over_under = Over_underUnder
	default:
		return errUnkownEnum
	}
	return
}

func (over_under *Over_under) ToCodeString() (res string) {

	switch *over_under {
	// insertion code per enum code
	case Over_underOver:
		res = "Over_underOver"
	case Over_underUnder:
		res = "Over_underUnder"
	}
	return
}

func (over_under Over_under) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Over_underOver")
	res = append(res, "Over_underUnder")

	return
}

func (over_under Over_under) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "over")
	res = append(res, "under")

	return
}

// Utility function for Pedal_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (pedal_type Pedal_type) ToString() (res string) {

	// migration of former implementation of enum
	switch pedal_type {
	// insertion code per enum code
	case Pedal_typeStart:
		res = "start"
	case Pedal_typeStop:
		res = "stop"
	case Pedal_typeSostenuto:
		res = "sostenuto"
	case Pedal_typeChange:
		res = "change"
	case Pedal_typeContinue_:
		res = "continue"
	case Pedal_typeDiscontinue:
		res = "discontinue"
	case Pedal_typeResume:
		res = "resume"
	}
	return
}

func (pedal_type *Pedal_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*pedal_type = Pedal_typeStart
		return
	case "stop":
		*pedal_type = Pedal_typeStop
		return
	case "sostenuto":
		*pedal_type = Pedal_typeSostenuto
		return
	case "change":
		*pedal_type = Pedal_typeChange
		return
	case "continue":
		*pedal_type = Pedal_typeContinue_
		return
	case "discontinue":
		*pedal_type = Pedal_typeDiscontinue
		return
	case "resume":
		*pedal_type = Pedal_typeResume
		return
	default:
		return errUnkownEnum
	}
}

func (pedal_type *Pedal_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Pedal_typeStart":
		*pedal_type = Pedal_typeStart
	case "Pedal_typeStop":
		*pedal_type = Pedal_typeStop
	case "Pedal_typeSostenuto":
		*pedal_type = Pedal_typeSostenuto
	case "Pedal_typeChange":
		*pedal_type = Pedal_typeChange
	case "Pedal_typeContinue_":
		*pedal_type = Pedal_typeContinue_
	case "Pedal_typeDiscontinue":
		*pedal_type = Pedal_typeDiscontinue
	case "Pedal_typeResume":
		*pedal_type = Pedal_typeResume
	default:
		return errUnkownEnum
	}
	return
}

func (pedal_type *Pedal_type) ToCodeString() (res string) {

	switch *pedal_type {
	// insertion code per enum code
	case Pedal_typeStart:
		res = "Pedal_typeStart"
	case Pedal_typeStop:
		res = "Pedal_typeStop"
	case Pedal_typeSostenuto:
		res = "Pedal_typeSostenuto"
	case Pedal_typeChange:
		res = "Pedal_typeChange"
	case Pedal_typeContinue_:
		res = "Pedal_typeContinue_"
	case Pedal_typeDiscontinue:
		res = "Pedal_typeDiscontinue"
	case Pedal_typeResume:
		res = "Pedal_typeResume"
	}
	return
}

func (pedal_type Pedal_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Pedal_typeStart")
	res = append(res, "Pedal_typeStop")
	res = append(res, "Pedal_typeSostenuto")
	res = append(res, "Pedal_typeChange")
	res = append(res, "Pedal_typeContinue_")
	res = append(res, "Pedal_typeDiscontinue")
	res = append(res, "Pedal_typeResume")

	return
}

func (pedal_type Pedal_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "sostenuto")
	res = append(res, "change")
	res = append(res, "continue")
	res = append(res, "discontinue")
	res = append(res, "resume")

	return
}

// Utility function for Pitched_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (pitched_value Pitched_value) ToString() (res string) {

	// migration of former implementation of enum
	switch pitched_value {
	// insertion code per enum code
	case Pitched_valueCelesta:
		res = "celesta"
	case Pitched_valueChimes:
		res = "chimes"
	case Pitched_valueGlockenspiel:
		res = "glockenspiel"
	case Pitched_valueLithophone:
		res = "lithophone"
	case Pitched_valueMallet:
		res = "mallet"
	case Pitched_valueMarimba:
		res = "marimba"
	case Pitched_valueSteeldrums:
		res = "steel drums"
	case Pitched_valueTubaphone:
		res = "tubaphone"
	case Pitched_valueTubularchimes:
		res = "tubular chimes"
	case Pitched_valueVibraphone:
		res = "vibraphone"
	case Pitched_valueXylophone:
		res = "xylophone"
	}
	return
}

func (pitched_value *Pitched_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "celesta":
		*pitched_value = Pitched_valueCelesta
		return
	case "chimes":
		*pitched_value = Pitched_valueChimes
		return
	case "glockenspiel":
		*pitched_value = Pitched_valueGlockenspiel
		return
	case "lithophone":
		*pitched_value = Pitched_valueLithophone
		return
	case "mallet":
		*pitched_value = Pitched_valueMallet
		return
	case "marimba":
		*pitched_value = Pitched_valueMarimba
		return
	case "steel drums":
		*pitched_value = Pitched_valueSteeldrums
		return
	case "tubaphone":
		*pitched_value = Pitched_valueTubaphone
		return
	case "tubular chimes":
		*pitched_value = Pitched_valueTubularchimes
		return
	case "vibraphone":
		*pitched_value = Pitched_valueVibraphone
		return
	case "xylophone":
		*pitched_value = Pitched_valueXylophone
		return
	default:
		return errUnkownEnum
	}
}

func (pitched_value *Pitched_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Pitched_valueCelesta":
		*pitched_value = Pitched_valueCelesta
	case "Pitched_valueChimes":
		*pitched_value = Pitched_valueChimes
	case "Pitched_valueGlockenspiel":
		*pitched_value = Pitched_valueGlockenspiel
	case "Pitched_valueLithophone":
		*pitched_value = Pitched_valueLithophone
	case "Pitched_valueMallet":
		*pitched_value = Pitched_valueMallet
	case "Pitched_valueMarimba":
		*pitched_value = Pitched_valueMarimba
	case "Pitched_valueSteeldrums":
		*pitched_value = Pitched_valueSteeldrums
	case "Pitched_valueTubaphone":
		*pitched_value = Pitched_valueTubaphone
	case "Pitched_valueTubularchimes":
		*pitched_value = Pitched_valueTubularchimes
	case "Pitched_valueVibraphone":
		*pitched_value = Pitched_valueVibraphone
	case "Pitched_valueXylophone":
		*pitched_value = Pitched_valueXylophone
	default:
		return errUnkownEnum
	}
	return
}

func (pitched_value *Pitched_value) ToCodeString() (res string) {

	switch *pitched_value {
	// insertion code per enum code
	case Pitched_valueCelesta:
		res = "Pitched_valueCelesta"
	case Pitched_valueChimes:
		res = "Pitched_valueChimes"
	case Pitched_valueGlockenspiel:
		res = "Pitched_valueGlockenspiel"
	case Pitched_valueLithophone:
		res = "Pitched_valueLithophone"
	case Pitched_valueMallet:
		res = "Pitched_valueMallet"
	case Pitched_valueMarimba:
		res = "Pitched_valueMarimba"
	case Pitched_valueSteeldrums:
		res = "Pitched_valueSteeldrums"
	case Pitched_valueTubaphone:
		res = "Pitched_valueTubaphone"
	case Pitched_valueTubularchimes:
		res = "Pitched_valueTubularchimes"
	case Pitched_valueVibraphone:
		res = "Pitched_valueVibraphone"
	case Pitched_valueXylophone:
		res = "Pitched_valueXylophone"
	}
	return
}

func (pitched_value Pitched_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Pitched_valueCelesta")
	res = append(res, "Pitched_valueChimes")
	res = append(res, "Pitched_valueGlockenspiel")
	res = append(res, "Pitched_valueLithophone")
	res = append(res, "Pitched_valueMallet")
	res = append(res, "Pitched_valueMarimba")
	res = append(res, "Pitched_valueSteeldrums")
	res = append(res, "Pitched_valueTubaphone")
	res = append(res, "Pitched_valueTubularchimes")
	res = append(res, "Pitched_valueVibraphone")
	res = append(res, "Pitched_valueXylophone")

	return
}

func (pitched_value Pitched_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "celesta")
	res = append(res, "chimes")
	res = append(res, "glockenspiel")
	res = append(res, "lithophone")
	res = append(res, "mallet")
	res = append(res, "marimba")
	res = append(res, "steel drums")
	res = append(res, "tubaphone")
	res = append(res, "tubular chimes")
	res = append(res, "vibraphone")
	res = append(res, "xylophone")

	return
}

// Utility function for PositiveInteger
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (positiveinteger PositiveInteger) ToInt() (res int) {

	// migration of former implementation of enum
	switch positiveinteger {
	// insertion code per enum code
	}
	return
}

func (positiveinteger *PositiveInteger) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (positiveinteger *PositiveInteger) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (positiveinteger *PositiveInteger) ToCodeString() (res string) {

	switch *positiveinteger {
	// insertion code per enum code
	}
	return
}

func (positiveinteger PositiveInteger) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (positiveinteger PositiveInteger) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code

	return
}

// Utility function for Positive_integer_or_empty
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (positive_integer_or_empty Positive_integer_or_empty) ToString() (res string) {

	// migration of former implementation of enum
	switch positive_integer_or_empty {
	// insertion code per enum code
	}
	return
}

func (positive_integer_or_empty *Positive_integer_or_empty) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (positive_integer_or_empty *Positive_integer_or_empty) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (positive_integer_or_empty *Positive_integer_or_empty) ToCodeString() (res string) {

	switch *positive_integer_or_empty {
	// insertion code per enum code
	}
	return
}

func (positive_integer_or_empty Positive_integer_or_empty) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (positive_integer_or_empty Positive_integer_or_empty) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Principal_voice_symbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (principal_voice_symbol Principal_voice_symbol) ToString() (res string) {

	// migration of former implementation of enum
	switch principal_voice_symbol {
	// insertion code per enum code
	case Principal_voice_symbolHauptstimme:
		res = "Hauptstimme"
	case Principal_voice_symbolNebenstimme:
		res = "Nebenstimme"
	case Principal_voice_symbolPlain:
		res = "plain"
	case Principal_voice_symbolNone:
		res = "none"
	}
	return
}

func (principal_voice_symbol *Principal_voice_symbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Hauptstimme":
		*principal_voice_symbol = Principal_voice_symbolHauptstimme
		return
	case "Nebenstimme":
		*principal_voice_symbol = Principal_voice_symbolNebenstimme
		return
	case "plain":
		*principal_voice_symbol = Principal_voice_symbolPlain
		return
	case "none":
		*principal_voice_symbol = Principal_voice_symbolNone
		return
	default:
		return errUnkownEnum
	}
}

func (principal_voice_symbol *Principal_voice_symbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Principal_voice_symbolHauptstimme":
		*principal_voice_symbol = Principal_voice_symbolHauptstimme
	case "Principal_voice_symbolNebenstimme":
		*principal_voice_symbol = Principal_voice_symbolNebenstimme
	case "Principal_voice_symbolPlain":
		*principal_voice_symbol = Principal_voice_symbolPlain
	case "Principal_voice_symbolNone":
		*principal_voice_symbol = Principal_voice_symbolNone
	default:
		return errUnkownEnum
	}
	return
}

func (principal_voice_symbol *Principal_voice_symbol) ToCodeString() (res string) {

	switch *principal_voice_symbol {
	// insertion code per enum code
	case Principal_voice_symbolHauptstimme:
		res = "Principal_voice_symbolHauptstimme"
	case Principal_voice_symbolNebenstimme:
		res = "Principal_voice_symbolNebenstimme"
	case Principal_voice_symbolPlain:
		res = "Principal_voice_symbolPlain"
	case Principal_voice_symbolNone:
		res = "Principal_voice_symbolNone"
	}
	return
}

func (principal_voice_symbol Principal_voice_symbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Principal_voice_symbolHauptstimme")
	res = append(res, "Principal_voice_symbolNebenstimme")
	res = append(res, "Principal_voice_symbolPlain")
	res = append(res, "Principal_voice_symbolNone")

	return
}

func (principal_voice_symbol Principal_voice_symbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Hauptstimme")
	res = append(res, "Nebenstimme")
	res = append(res, "plain")
	res = append(res, "none")

	return
}

// Utility function for Right_left_middle
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (right_left_middle Right_left_middle) ToString() (res string) {

	// migration of former implementation of enum
	switch right_left_middle {
	// insertion code per enum code
	case Right_left_middleRight:
		res = "right"
	case Right_left_middleLeft:
		res = "left"
	case Right_left_middleMiddle:
		res = "middle"
	}
	return
}

func (right_left_middle *Right_left_middle) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "right":
		*right_left_middle = Right_left_middleRight
		return
	case "left":
		*right_left_middle = Right_left_middleLeft
		return
	case "middle":
		*right_left_middle = Right_left_middleMiddle
		return
	default:
		return errUnkownEnum
	}
}

func (right_left_middle *Right_left_middle) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Right_left_middleRight":
		*right_left_middle = Right_left_middleRight
	case "Right_left_middleLeft":
		*right_left_middle = Right_left_middleLeft
	case "Right_left_middleMiddle":
		*right_left_middle = Right_left_middleMiddle
	default:
		return errUnkownEnum
	}
	return
}

func (right_left_middle *Right_left_middle) ToCodeString() (res string) {

	switch *right_left_middle {
	// insertion code per enum code
	case Right_left_middleRight:
		res = "Right_left_middleRight"
	case Right_left_middleLeft:
		res = "Right_left_middleLeft"
	case Right_left_middleMiddle:
		res = "Right_left_middleMiddle"
	}
	return
}

func (right_left_middle Right_left_middle) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Right_left_middleRight")
	res = append(res, "Right_left_middleLeft")
	res = append(res, "Right_left_middleMiddle")

	return
}

func (right_left_middle Right_left_middle) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "right")
	res = append(res, "left")
	res = append(res, "middle")

	return
}

// Utility function for Semi_pitched
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (semi_pitched Semi_pitched) ToString() (res string) {

	// migration of former implementation of enum
	switch semi_pitched {
	// insertion code per enum code
	case Semi_pitchedHigh:
		res = "high"
	case Semi_pitchedMedium_high:
		res = "medium-high"
	case Semi_pitchedMedium:
		res = "medium"
	case Semi_pitchedMedium_low:
		res = "medium-low"
	case Semi_pitchedLow:
		res = "low"
	case Semi_pitchedVery_low:
		res = "very-low"
	}
	return
}

func (semi_pitched *Semi_pitched) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "high":
		*semi_pitched = Semi_pitchedHigh
		return
	case "medium-high":
		*semi_pitched = Semi_pitchedMedium_high
		return
	case "medium":
		*semi_pitched = Semi_pitchedMedium
		return
	case "medium-low":
		*semi_pitched = Semi_pitchedMedium_low
		return
	case "low":
		*semi_pitched = Semi_pitchedLow
		return
	case "very-low":
		*semi_pitched = Semi_pitchedVery_low
		return
	default:
		return errUnkownEnum
	}
}

func (semi_pitched *Semi_pitched) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Semi_pitchedHigh":
		*semi_pitched = Semi_pitchedHigh
	case "Semi_pitchedMedium_high":
		*semi_pitched = Semi_pitchedMedium_high
	case "Semi_pitchedMedium":
		*semi_pitched = Semi_pitchedMedium
	case "Semi_pitchedMedium_low":
		*semi_pitched = Semi_pitchedMedium_low
	case "Semi_pitchedLow":
		*semi_pitched = Semi_pitchedLow
	case "Semi_pitchedVery_low":
		*semi_pitched = Semi_pitchedVery_low
	default:
		return errUnkownEnum
	}
	return
}

func (semi_pitched *Semi_pitched) ToCodeString() (res string) {

	switch *semi_pitched {
	// insertion code per enum code
	case Semi_pitchedHigh:
		res = "Semi_pitchedHigh"
	case Semi_pitchedMedium_high:
		res = "Semi_pitchedMedium_high"
	case Semi_pitchedMedium:
		res = "Semi_pitchedMedium"
	case Semi_pitchedMedium_low:
		res = "Semi_pitchedMedium_low"
	case Semi_pitchedLow:
		res = "Semi_pitchedLow"
	case Semi_pitchedVery_low:
		res = "Semi_pitchedVery_low"
	}
	return
}

func (semi_pitched Semi_pitched) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Semi_pitchedHigh")
	res = append(res, "Semi_pitchedMedium_high")
	res = append(res, "Semi_pitchedMedium")
	res = append(res, "Semi_pitchedMedium_low")
	res = append(res, "Semi_pitchedLow")
	res = append(res, "Semi_pitchedVery_low")

	return
}

func (semi_pitched Semi_pitched) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "high")
	res = append(res, "medium-high")
	res = append(res, "medium")
	res = append(res, "medium-low")
	res = append(res, "low")
	res = append(res, "very-low")

	return
}

// Utility function for Show_frets
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (show_frets Show_frets) ToString() (res string) {

	// migration of former implementation of enum
	switch show_frets {
	// insertion code per enum code
	case Show_fretsNumbers:
		res = "numbers"
	case Show_fretsLetters:
		res = "letters"
	}
	return
}

func (show_frets *Show_frets) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "numbers":
		*show_frets = Show_fretsNumbers
		return
	case "letters":
		*show_frets = Show_fretsLetters
		return
	default:
		return errUnkownEnum
	}
}

func (show_frets *Show_frets) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Show_fretsNumbers":
		*show_frets = Show_fretsNumbers
	case "Show_fretsLetters":
		*show_frets = Show_fretsLetters
	default:
		return errUnkownEnum
	}
	return
}

func (show_frets *Show_frets) ToCodeString() (res string) {

	switch *show_frets {
	// insertion code per enum code
	case Show_fretsNumbers:
		res = "Show_fretsNumbers"
	case Show_fretsLetters:
		res = "Show_fretsLetters"
	}
	return
}

func (show_frets Show_frets) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Show_fretsNumbers")
	res = append(res, "Show_fretsLetters")

	return
}

func (show_frets Show_frets) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "numbers")
	res = append(res, "letters")

	return
}

// Utility function for Show_tuplet
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (show_tuplet Show_tuplet) ToString() (res string) {

	// migration of former implementation of enum
	switch show_tuplet {
	// insertion code per enum code
	case Show_tupletActual:
		res = "actual"
	case Show_tupletBoth:
		res = "both"
	case Show_tupletNone:
		res = "none"
	}
	return
}

func (show_tuplet *Show_tuplet) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "actual":
		*show_tuplet = Show_tupletActual
		return
	case "both":
		*show_tuplet = Show_tupletBoth
		return
	case "none":
		*show_tuplet = Show_tupletNone
		return
	default:
		return errUnkownEnum
	}
}

func (show_tuplet *Show_tuplet) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Show_tupletActual":
		*show_tuplet = Show_tupletActual
	case "Show_tupletBoth":
		*show_tuplet = Show_tupletBoth
	case "Show_tupletNone":
		*show_tuplet = Show_tupletNone
	default:
		return errUnkownEnum
	}
	return
}

func (show_tuplet *Show_tuplet) ToCodeString() (res string) {

	switch *show_tuplet {
	// insertion code per enum code
	case Show_tupletActual:
		res = "Show_tupletActual"
	case Show_tupletBoth:
		res = "Show_tupletBoth"
	case Show_tupletNone:
		res = "Show_tupletNone"
	}
	return
}

func (show_tuplet Show_tuplet) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Show_tupletActual")
	res = append(res, "Show_tupletBoth")
	res = append(res, "Show_tupletNone")

	return
}

func (show_tuplet Show_tuplet) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "actual")
	res = append(res, "both")
	res = append(res, "none")

	return
}

// Utility function for Staff_divide_symbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (staff_divide_symbol Staff_divide_symbol) ToString() (res string) {

	// migration of former implementation of enum
	switch staff_divide_symbol {
	// insertion code per enum code
	case Staff_divide_symbolDown:
		res = "down"
	case Staff_divide_symbolUp:
		res = "up"
	case Staff_divide_symbolUp_down:
		res = "up-down"
	}
	return
}

func (staff_divide_symbol *Staff_divide_symbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "down":
		*staff_divide_symbol = Staff_divide_symbolDown
		return
	case "up":
		*staff_divide_symbol = Staff_divide_symbolUp
		return
	case "up-down":
		*staff_divide_symbol = Staff_divide_symbolUp_down
		return
	default:
		return errUnkownEnum
	}
}

func (staff_divide_symbol *Staff_divide_symbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Staff_divide_symbolDown":
		*staff_divide_symbol = Staff_divide_symbolDown
	case "Staff_divide_symbolUp":
		*staff_divide_symbol = Staff_divide_symbolUp
	case "Staff_divide_symbolUp_down":
		*staff_divide_symbol = Staff_divide_symbolUp_down
	default:
		return errUnkownEnum
	}
	return
}

func (staff_divide_symbol *Staff_divide_symbol) ToCodeString() (res string) {

	switch *staff_divide_symbol {
	// insertion code per enum code
	case Staff_divide_symbolDown:
		res = "Staff_divide_symbolDown"
	case Staff_divide_symbolUp:
		res = "Staff_divide_symbolUp"
	case Staff_divide_symbolUp_down:
		res = "Staff_divide_symbolUp_down"
	}
	return
}

func (staff_divide_symbol Staff_divide_symbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Staff_divide_symbolDown")
	res = append(res, "Staff_divide_symbolUp")
	res = append(res, "Staff_divide_symbolUp_down")

	return
}

func (staff_divide_symbol Staff_divide_symbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "down")
	res = append(res, "up")
	res = append(res, "up-down")

	return
}

// Utility function for Staff_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (staff_type Staff_type) ToString() (res string) {

	// migration of former implementation of enum
	switch staff_type {
	// insertion code per enum code
	case Staff_typeOssia:
		res = "ossia"
	case Staff_typeEditorial:
		res = "editorial"
	case Staff_typeCue:
		res = "cue"
	case Staff_typeAlternate:
		res = "alternate"
	case Staff_typeRegular:
		res = "regular"
	}
	return
}

func (staff_type *Staff_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ossia":
		*staff_type = Staff_typeOssia
		return
	case "editorial":
		*staff_type = Staff_typeEditorial
		return
	case "cue":
		*staff_type = Staff_typeCue
		return
	case "alternate":
		*staff_type = Staff_typeAlternate
		return
	case "regular":
		*staff_type = Staff_typeRegular
		return
	default:
		return errUnkownEnum
	}
}

func (staff_type *Staff_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Staff_typeOssia":
		*staff_type = Staff_typeOssia
	case "Staff_typeEditorial":
		*staff_type = Staff_typeEditorial
	case "Staff_typeCue":
		*staff_type = Staff_typeCue
	case "Staff_typeAlternate":
		*staff_type = Staff_typeAlternate
	case "Staff_typeRegular":
		*staff_type = Staff_typeRegular
	default:
		return errUnkownEnum
	}
	return
}

func (staff_type *Staff_type) ToCodeString() (res string) {

	switch *staff_type {
	// insertion code per enum code
	case Staff_typeOssia:
		res = "Staff_typeOssia"
	case Staff_typeEditorial:
		res = "Staff_typeEditorial"
	case Staff_typeCue:
		res = "Staff_typeCue"
	case Staff_typeAlternate:
		res = "Staff_typeAlternate"
	case Staff_typeRegular:
		res = "Staff_typeRegular"
	}
	return
}

func (staff_type Staff_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Staff_typeOssia")
	res = append(res, "Staff_typeEditorial")
	res = append(res, "Staff_typeCue")
	res = append(res, "Staff_typeAlternate")
	res = append(res, "Staff_typeRegular")

	return
}

func (staff_type Staff_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ossia")
	res = append(res, "editorial")
	res = append(res, "cue")
	res = append(res, "alternate")
	res = append(res, "regular")

	return
}

// Utility function for Start_note
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_note Start_note) ToString() (res string) {

	// migration of former implementation of enum
	switch start_note {
	// insertion code per enum code
	case Start_noteUpper:
		res = "upper"
	case Start_noteMain:
		res = "main"
	case Start_noteBelow:
		res = "below"
	}
	return
}

func (start_note *Start_note) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "upper":
		*start_note = Start_noteUpper
		return
	case "main":
		*start_note = Start_noteMain
		return
	case "below":
		*start_note = Start_noteBelow
		return
	default:
		return errUnkownEnum
	}
}

func (start_note *Start_note) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Start_noteUpper":
		*start_note = Start_noteUpper
	case "Start_noteMain":
		*start_note = Start_noteMain
	case "Start_noteBelow":
		*start_note = Start_noteBelow
	default:
		return errUnkownEnum
	}
	return
}

func (start_note *Start_note) ToCodeString() (res string) {

	switch *start_note {
	// insertion code per enum code
	case Start_noteUpper:
		res = "Start_noteUpper"
	case Start_noteMain:
		res = "Start_noteMain"
	case Start_noteBelow:
		res = "Start_noteBelow"
	}
	return
}

func (start_note Start_note) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Start_noteUpper")
	res = append(res, "Start_noteMain")
	res = append(res, "Start_noteBelow")

	return
}

func (start_note Start_note) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "upper")
	res = append(res, "main")
	res = append(res, "below")

	return
}

// Utility function for Start_stop
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_stop Start_stop) ToString() (res string) {

	// migration of former implementation of enum
	switch start_stop {
	// insertion code per enum code
	case Start_stopStart:
		res = "start"
	case Start_stopStop:
		res = "stop"
	}
	return
}

func (start_stop *Start_stop) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*start_stop = Start_stopStart
		return
	case "stop":
		*start_stop = Start_stopStop
		return
	default:
		return errUnkownEnum
	}
}

func (start_stop *Start_stop) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Start_stopStart":
		*start_stop = Start_stopStart
	case "Start_stopStop":
		*start_stop = Start_stopStop
	default:
		return errUnkownEnum
	}
	return
}

func (start_stop *Start_stop) ToCodeString() (res string) {

	switch *start_stop {
	// insertion code per enum code
	case Start_stopStart:
		res = "Start_stopStart"
	case Start_stopStop:
		res = "Start_stopStop"
	}
	return
}

func (start_stop Start_stop) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Start_stopStart")
	res = append(res, "Start_stopStop")

	return
}

func (start_stop Start_stop) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")

	return
}

// Utility function for Start_stop_change_continue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_stop_change_continue Start_stop_change_continue) ToString() (res string) {

	// migration of former implementation of enum
	switch start_stop_change_continue {
	// insertion code per enum code
	case Start_stop_change_continueStart:
		res = "start"
	case Start_stop_change_continueStop:
		res = "stop"
	case Start_stop_change_continueChange:
		res = "change"
	case Start_stop_change_continueContinue_:
		res = "continue"
	}
	return
}

func (start_stop_change_continue *Start_stop_change_continue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*start_stop_change_continue = Start_stop_change_continueStart
		return
	case "stop":
		*start_stop_change_continue = Start_stop_change_continueStop
		return
	case "change":
		*start_stop_change_continue = Start_stop_change_continueChange
		return
	case "continue":
		*start_stop_change_continue = Start_stop_change_continueContinue_
		return
	default:
		return errUnkownEnum
	}
}

func (start_stop_change_continue *Start_stop_change_continue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Start_stop_change_continueStart":
		*start_stop_change_continue = Start_stop_change_continueStart
	case "Start_stop_change_continueStop":
		*start_stop_change_continue = Start_stop_change_continueStop
	case "Start_stop_change_continueChange":
		*start_stop_change_continue = Start_stop_change_continueChange
	case "Start_stop_change_continueContinue_":
		*start_stop_change_continue = Start_stop_change_continueContinue_
	default:
		return errUnkownEnum
	}
	return
}

func (start_stop_change_continue *Start_stop_change_continue) ToCodeString() (res string) {

	switch *start_stop_change_continue {
	// insertion code per enum code
	case Start_stop_change_continueStart:
		res = "Start_stop_change_continueStart"
	case Start_stop_change_continueStop:
		res = "Start_stop_change_continueStop"
	case Start_stop_change_continueChange:
		res = "Start_stop_change_continueChange"
	case Start_stop_change_continueContinue_:
		res = "Start_stop_change_continueContinue_"
	}
	return
}

func (start_stop_change_continue Start_stop_change_continue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Start_stop_change_continueStart")
	res = append(res, "Start_stop_change_continueStop")
	res = append(res, "Start_stop_change_continueChange")
	res = append(res, "Start_stop_change_continueContinue_")

	return
}

func (start_stop_change_continue Start_stop_change_continue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "change")
	res = append(res, "continue")

	return
}

// Utility function for Start_stop_continue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_stop_continue Start_stop_continue) ToString() (res string) {

	// migration of former implementation of enum
	switch start_stop_continue {
	// insertion code per enum code
	case Start_stop_continueStart:
		res = "start"
	case Start_stop_continueStop:
		res = "stop"
	case Start_stop_continueContinue_:
		res = "continue"
	}
	return
}

func (start_stop_continue *Start_stop_continue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*start_stop_continue = Start_stop_continueStart
		return
	case "stop":
		*start_stop_continue = Start_stop_continueStop
		return
	case "continue":
		*start_stop_continue = Start_stop_continueContinue_
		return
	default:
		return errUnkownEnum
	}
}

func (start_stop_continue *Start_stop_continue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Start_stop_continueStart":
		*start_stop_continue = Start_stop_continueStart
	case "Start_stop_continueStop":
		*start_stop_continue = Start_stop_continueStop
	case "Start_stop_continueContinue_":
		*start_stop_continue = Start_stop_continueContinue_
	default:
		return errUnkownEnum
	}
	return
}

func (start_stop_continue *Start_stop_continue) ToCodeString() (res string) {

	switch *start_stop_continue {
	// insertion code per enum code
	case Start_stop_continueStart:
		res = "Start_stop_continueStart"
	case Start_stop_continueStop:
		res = "Start_stop_continueStop"
	case Start_stop_continueContinue_:
		res = "Start_stop_continueContinue_"
	}
	return
}

func (start_stop_continue Start_stop_continue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Start_stop_continueStart")
	res = append(res, "Start_stop_continueStop")
	res = append(res, "Start_stop_continueContinue_")

	return
}

func (start_stop_continue Start_stop_continue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for Start_stop_discontinue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_stop_discontinue Start_stop_discontinue) ToString() (res string) {

	// migration of former implementation of enum
	switch start_stop_discontinue {
	// insertion code per enum code
	case Start_stop_discontinueStart:
		res = "start"
	case Start_stop_discontinueStop:
		res = "stop"
	case Start_stop_discontinueDiscontinue:
		res = "discontinue"
	}
	return
}

func (start_stop_discontinue *Start_stop_discontinue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*start_stop_discontinue = Start_stop_discontinueStart
		return
	case "stop":
		*start_stop_discontinue = Start_stop_discontinueStop
		return
	case "discontinue":
		*start_stop_discontinue = Start_stop_discontinueDiscontinue
		return
	default:
		return errUnkownEnum
	}
}

func (start_stop_discontinue *Start_stop_discontinue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Start_stop_discontinueStart":
		*start_stop_discontinue = Start_stop_discontinueStart
	case "Start_stop_discontinueStop":
		*start_stop_discontinue = Start_stop_discontinueStop
	case "Start_stop_discontinueDiscontinue":
		*start_stop_discontinue = Start_stop_discontinueDiscontinue
	default:
		return errUnkownEnum
	}
	return
}

func (start_stop_discontinue *Start_stop_discontinue) ToCodeString() (res string) {

	switch *start_stop_discontinue {
	// insertion code per enum code
	case Start_stop_discontinueStart:
		res = "Start_stop_discontinueStart"
	case Start_stop_discontinueStop:
		res = "Start_stop_discontinueStop"
	case Start_stop_discontinueDiscontinue:
		res = "Start_stop_discontinueDiscontinue"
	}
	return
}

func (start_stop_discontinue Start_stop_discontinue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Start_stop_discontinueStart")
	res = append(res, "Start_stop_discontinueStop")
	res = append(res, "Start_stop_discontinueDiscontinue")

	return
}

func (start_stop_discontinue Start_stop_discontinue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "discontinue")

	return
}

// Utility function for Start_stop_single
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (start_stop_single Start_stop_single) ToString() (res string) {

	// migration of former implementation of enum
	switch start_stop_single {
	// insertion code per enum code
	case Start_stop_singleStart:
		res = "start"
	case Start_stop_singleStop:
		res = "stop"
	case Start_stop_singleSingle:
		res = "single"
	}
	return
}

func (start_stop_single *Start_stop_single) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*start_stop_single = Start_stop_singleStart
		return
	case "stop":
		*start_stop_single = Start_stop_singleStop
		return
	case "single":
		*start_stop_single = Start_stop_singleSingle
		return
	default:
		return errUnkownEnum
	}
}

func (start_stop_single *Start_stop_single) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Start_stop_singleStart":
		*start_stop_single = Start_stop_singleStart
	case "Start_stop_singleStop":
		*start_stop_single = Start_stop_singleStop
	case "Start_stop_singleSingle":
		*start_stop_single = Start_stop_singleSingle
	default:
		return errUnkownEnum
	}
	return
}

func (start_stop_single *Start_stop_single) ToCodeString() (res string) {

	switch *start_stop_single {
	// insertion code per enum code
	case Start_stop_singleStart:
		res = "Start_stop_singleStart"
	case Start_stop_singleStop:
		res = "Start_stop_singleStop"
	case Start_stop_singleSingle:
		res = "Start_stop_singleSingle"
	}
	return
}

func (start_stop_single Start_stop_single) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Start_stop_singleStart")
	res = append(res, "Start_stop_singleStop")
	res = append(res, "Start_stop_singleSingle")

	return
}

func (start_stop_single Start_stop_single) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "single")

	return
}

// Utility function for Stem_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stem_value Stem_value) ToString() (res string) {

	// migration of former implementation of enum
	switch stem_value {
	// insertion code per enum code
	case Stem_valueDown:
		res = "down"
	case Stem_valueUp:
		res = "up"
	case Stem_valueDouble:
		res = "double"
	case Stem_valueNone:
		res = "none"
	}
	return
}

func (stem_value *Stem_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "down":
		*stem_value = Stem_valueDown
		return
	case "up":
		*stem_value = Stem_valueUp
		return
	case "double":
		*stem_value = Stem_valueDouble
		return
	case "none":
		*stem_value = Stem_valueNone
		return
	default:
		return errUnkownEnum
	}
}

func (stem_value *Stem_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Stem_valueDown":
		*stem_value = Stem_valueDown
	case "Stem_valueUp":
		*stem_value = Stem_valueUp
	case "Stem_valueDouble":
		*stem_value = Stem_valueDouble
	case "Stem_valueNone":
		*stem_value = Stem_valueNone
	default:
		return errUnkownEnum
	}
	return
}

func (stem_value *Stem_value) ToCodeString() (res string) {

	switch *stem_value {
	// insertion code per enum code
	case Stem_valueDown:
		res = "Stem_valueDown"
	case Stem_valueUp:
		res = "Stem_valueUp"
	case Stem_valueDouble:
		res = "Stem_valueDouble"
	case Stem_valueNone:
		res = "Stem_valueNone"
	}
	return
}

func (stem_value Stem_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Stem_valueDown")
	res = append(res, "Stem_valueUp")
	res = append(res, "Stem_valueDouble")
	res = append(res, "Stem_valueNone")

	return
}

func (stem_value Stem_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "down")
	res = append(res, "up")
	res = append(res, "double")
	res = append(res, "none")

	return
}

// Utility function for Step
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (step Step) ToString() (res string) {

	// migration of former implementation of enum
	switch step {
	// insertion code per enum code
	case StepA:
		res = "A"
	case StepB:
		res = "B"
	case StepC:
		res = "C"
	case StepD:
		res = "D"
	case StepE:
		res = "E"
	case StepF:
		res = "F"
	case StepG:
		res = "G"
	}
	return
}

func (step *Step) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "A":
		*step = StepA
		return
	case "B":
		*step = StepB
		return
	case "C":
		*step = StepC
		return
	case "D":
		*step = StepD
		return
	case "E":
		*step = StepE
		return
	case "F":
		*step = StepF
		return
	case "G":
		*step = StepG
		return
	default:
		return errUnkownEnum
	}
}

func (step *Step) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "StepA":
		*step = StepA
	case "StepB":
		*step = StepB
	case "StepC":
		*step = StepC
	case "StepD":
		*step = StepD
	case "StepE":
		*step = StepE
	case "StepF":
		*step = StepF
	case "StepG":
		*step = StepG
	default:
		return errUnkownEnum
	}
	return
}

func (step *Step) ToCodeString() (res string) {

	switch *step {
	// insertion code per enum code
	case StepA:
		res = "StepA"
	case StepB:
		res = "StepB"
	case StepC:
		res = "StepC"
	case StepD:
		res = "StepD"
	case StepE:
		res = "StepE"
	case StepF:
		res = "StepF"
	case StepG:
		res = "StepG"
	}
	return
}

func (step Step) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "StepA")
	res = append(res, "StepB")
	res = append(res, "StepC")
	res = append(res, "StepD")
	res = append(res, "StepE")
	res = append(res, "StepF")
	res = append(res, "StepG")

	return
}

func (step Step) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "A")
	res = append(res, "B")
	res = append(res, "C")
	res = append(res, "D")
	res = append(res, "E")
	res = append(res, "F")
	res = append(res, "G")

	return
}

// Utility function for Stick_location
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stick_location Stick_location) ToString() (res string) {

	// migration of former implementation of enum
	switch stick_location {
	// insertion code per enum code
	case Stick_locationCenter:
		res = "center"
	case Stick_locationRim:
		res = "rim"
	case Stick_locationCymbalbell:
		res = "cymbal bell"
	case Stick_locationCymbaledge:
		res = "cymbal edge"
	}
	return
}

func (stick_location *Stick_location) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "center":
		*stick_location = Stick_locationCenter
		return
	case "rim":
		*stick_location = Stick_locationRim
		return
	case "cymbal bell":
		*stick_location = Stick_locationCymbalbell
		return
	case "cymbal edge":
		*stick_location = Stick_locationCymbaledge
		return
	default:
		return errUnkownEnum
	}
}

func (stick_location *Stick_location) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Stick_locationCenter":
		*stick_location = Stick_locationCenter
	case "Stick_locationRim":
		*stick_location = Stick_locationRim
	case "Stick_locationCymbalbell":
		*stick_location = Stick_locationCymbalbell
	case "Stick_locationCymbaledge":
		*stick_location = Stick_locationCymbaledge
	default:
		return errUnkownEnum
	}
	return
}

func (stick_location *Stick_location) ToCodeString() (res string) {

	switch *stick_location {
	// insertion code per enum code
	case Stick_locationCenter:
		res = "Stick_locationCenter"
	case Stick_locationRim:
		res = "Stick_locationRim"
	case Stick_locationCymbalbell:
		res = "Stick_locationCymbalbell"
	case Stick_locationCymbaledge:
		res = "Stick_locationCymbaledge"
	}
	return
}

func (stick_location Stick_location) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Stick_locationCenter")
	res = append(res, "Stick_locationRim")
	res = append(res, "Stick_locationCymbalbell")
	res = append(res, "Stick_locationCymbaledge")

	return
}

func (stick_location Stick_location) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "center")
	res = append(res, "rim")
	res = append(res, "cymbal bell")
	res = append(res, "cymbal edge")

	return
}

// Utility function for Stick_material
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stick_material Stick_material) ToString() (res string) {

	// migration of former implementation of enum
	switch stick_material {
	// insertion code per enum code
	case Stick_materialSoft:
		res = "soft"
	case Stick_materialMedium:
		res = "medium"
	case Stick_materialHard:
		res = "hard"
	case Stick_materialShaded:
		res = "shaded"
	case Stick_materialX:
		res = "x"
	}
	return
}

func (stick_material *Stick_material) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "soft":
		*stick_material = Stick_materialSoft
		return
	case "medium":
		*stick_material = Stick_materialMedium
		return
	case "hard":
		*stick_material = Stick_materialHard
		return
	case "shaded":
		*stick_material = Stick_materialShaded
		return
	case "x":
		*stick_material = Stick_materialX
		return
	default:
		return errUnkownEnum
	}
}

func (stick_material *Stick_material) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Stick_materialSoft":
		*stick_material = Stick_materialSoft
	case "Stick_materialMedium":
		*stick_material = Stick_materialMedium
	case "Stick_materialHard":
		*stick_material = Stick_materialHard
	case "Stick_materialShaded":
		*stick_material = Stick_materialShaded
	case "Stick_materialX":
		*stick_material = Stick_materialX
	default:
		return errUnkownEnum
	}
	return
}

func (stick_material *Stick_material) ToCodeString() (res string) {

	switch *stick_material {
	// insertion code per enum code
	case Stick_materialSoft:
		res = "Stick_materialSoft"
	case Stick_materialMedium:
		res = "Stick_materialMedium"
	case Stick_materialHard:
		res = "Stick_materialHard"
	case Stick_materialShaded:
		res = "Stick_materialShaded"
	case Stick_materialX:
		res = "Stick_materialX"
	}
	return
}

func (stick_material Stick_material) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Stick_materialSoft")
	res = append(res, "Stick_materialMedium")
	res = append(res, "Stick_materialHard")
	res = append(res, "Stick_materialShaded")
	res = append(res, "Stick_materialX")

	return
}

func (stick_material Stick_material) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "soft")
	res = append(res, "medium")
	res = append(res, "hard")
	res = append(res, "shaded")
	res = append(res, "x")

	return
}

// Utility function for Stick_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stick_type Stick_type) ToString() (res string) {

	// migration of former implementation of enum
	switch stick_type {
	// insertion code per enum code
	case Stick_typeBassdrum:
		res = "bass drum"
	case Stick_typeDoublebassdrum:
		res = "double bass drum"
	case Stick_typeGlockenspiel:
		res = "glockenspiel"
	case Stick_typeGum:
		res = "gum"
	case Stick_typeHammer:
		res = "hammer"
	case Stick_typeSuperball:
		res = "superball"
	case Stick_typeTimpani:
		res = "timpani"
	case Stick_typeWound:
		res = "wound"
	case Stick_typeXylophone:
		res = "xylophone"
	case Stick_typeYarn:
		res = "yarn"
	}
	return
}

func (stick_type *Stick_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bass drum":
		*stick_type = Stick_typeBassdrum
		return
	case "double bass drum":
		*stick_type = Stick_typeDoublebassdrum
		return
	case "glockenspiel":
		*stick_type = Stick_typeGlockenspiel
		return
	case "gum":
		*stick_type = Stick_typeGum
		return
	case "hammer":
		*stick_type = Stick_typeHammer
		return
	case "superball":
		*stick_type = Stick_typeSuperball
		return
	case "timpani":
		*stick_type = Stick_typeTimpani
		return
	case "wound":
		*stick_type = Stick_typeWound
		return
	case "xylophone":
		*stick_type = Stick_typeXylophone
		return
	case "yarn":
		*stick_type = Stick_typeYarn
		return
	default:
		return errUnkownEnum
	}
}

func (stick_type *Stick_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Stick_typeBassdrum":
		*stick_type = Stick_typeBassdrum
	case "Stick_typeDoublebassdrum":
		*stick_type = Stick_typeDoublebassdrum
	case "Stick_typeGlockenspiel":
		*stick_type = Stick_typeGlockenspiel
	case "Stick_typeGum":
		*stick_type = Stick_typeGum
	case "Stick_typeHammer":
		*stick_type = Stick_typeHammer
	case "Stick_typeSuperball":
		*stick_type = Stick_typeSuperball
	case "Stick_typeTimpani":
		*stick_type = Stick_typeTimpani
	case "Stick_typeWound":
		*stick_type = Stick_typeWound
	case "Stick_typeXylophone":
		*stick_type = Stick_typeXylophone
	case "Stick_typeYarn":
		*stick_type = Stick_typeYarn
	default:
		return errUnkownEnum
	}
	return
}

func (stick_type *Stick_type) ToCodeString() (res string) {

	switch *stick_type {
	// insertion code per enum code
	case Stick_typeBassdrum:
		res = "Stick_typeBassdrum"
	case Stick_typeDoublebassdrum:
		res = "Stick_typeDoublebassdrum"
	case Stick_typeGlockenspiel:
		res = "Stick_typeGlockenspiel"
	case Stick_typeGum:
		res = "Stick_typeGum"
	case Stick_typeHammer:
		res = "Stick_typeHammer"
	case Stick_typeSuperball:
		res = "Stick_typeSuperball"
	case Stick_typeTimpani:
		res = "Stick_typeTimpani"
	case Stick_typeWound:
		res = "Stick_typeWound"
	case Stick_typeXylophone:
		res = "Stick_typeXylophone"
	case Stick_typeYarn:
		res = "Stick_typeYarn"
	}
	return
}

func (stick_type Stick_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Stick_typeBassdrum")
	res = append(res, "Stick_typeDoublebassdrum")
	res = append(res, "Stick_typeGlockenspiel")
	res = append(res, "Stick_typeGum")
	res = append(res, "Stick_typeHammer")
	res = append(res, "Stick_typeSuperball")
	res = append(res, "Stick_typeTimpani")
	res = append(res, "Stick_typeWound")
	res = append(res, "Stick_typeXylophone")
	res = append(res, "Stick_typeYarn")

	return
}

func (stick_type Stick_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bass drum")
	res = append(res, "double bass drum")
	res = append(res, "glockenspiel")
	res = append(res, "gum")
	res = append(res, "hammer")
	res = append(res, "superball")
	res = append(res, "timpani")
	res = append(res, "wound")
	res = append(res, "xylophone")
	res = append(res, "yarn")

	return
}

// Utility function for Syllabic
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (syllabic Syllabic) ToString() (res string) {

	// migration of former implementation of enum
	switch syllabic {
	// insertion code per enum code
	case SyllabicSingle:
		res = "single"
	case SyllabicBegin:
		res = "begin"
	case SyllabicEnd:
		res = "end"
	case SyllabicMiddle:
		res = "middle"
	}
	return
}

func (syllabic *Syllabic) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "single":
		*syllabic = SyllabicSingle
		return
	case "begin":
		*syllabic = SyllabicBegin
		return
	case "end":
		*syllabic = SyllabicEnd
		return
	case "middle":
		*syllabic = SyllabicMiddle
		return
	default:
		return errUnkownEnum
	}
}

func (syllabic *Syllabic) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SyllabicSingle":
		*syllabic = SyllabicSingle
	case "SyllabicBegin":
		*syllabic = SyllabicBegin
	case "SyllabicEnd":
		*syllabic = SyllabicEnd
	case "SyllabicMiddle":
		*syllabic = SyllabicMiddle
	default:
		return errUnkownEnum
	}
	return
}

func (syllabic *Syllabic) ToCodeString() (res string) {

	switch *syllabic {
	// insertion code per enum code
	case SyllabicSingle:
		res = "SyllabicSingle"
	case SyllabicBegin:
		res = "SyllabicBegin"
	case SyllabicEnd:
		res = "SyllabicEnd"
	case SyllabicMiddle:
		res = "SyllabicMiddle"
	}
	return
}

func (syllabic Syllabic) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SyllabicSingle")
	res = append(res, "SyllabicBegin")
	res = append(res, "SyllabicEnd")
	res = append(res, "SyllabicMiddle")

	return
}

func (syllabic Syllabic) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "single")
	res = append(res, "begin")
	res = append(res, "end")
	res = append(res, "middle")

	return
}

// Utility function for Symbol_size
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (symbol_size Symbol_size) ToString() (res string) {

	// migration of former implementation of enum
	switch symbol_size {
	// insertion code per enum code
	case Symbol_sizeFull:
		res = "full"
	case Symbol_sizeCue:
		res = "cue"
	case Symbol_sizeGrace_cue:
		res = "grace-cue"
	case Symbol_sizeLarge:
		res = "large"
	}
	return
}

func (symbol_size *Symbol_size) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "full":
		*symbol_size = Symbol_sizeFull
		return
	case "cue":
		*symbol_size = Symbol_sizeCue
		return
	case "grace-cue":
		*symbol_size = Symbol_sizeGrace_cue
		return
	case "large":
		*symbol_size = Symbol_sizeLarge
		return
	default:
		return errUnkownEnum
	}
}

func (symbol_size *Symbol_size) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Symbol_sizeFull":
		*symbol_size = Symbol_sizeFull
	case "Symbol_sizeCue":
		*symbol_size = Symbol_sizeCue
	case "Symbol_sizeGrace_cue":
		*symbol_size = Symbol_sizeGrace_cue
	case "Symbol_sizeLarge":
		*symbol_size = Symbol_sizeLarge
	default:
		return errUnkownEnum
	}
	return
}

func (symbol_size *Symbol_size) ToCodeString() (res string) {

	switch *symbol_size {
	// insertion code per enum code
	case Symbol_sizeFull:
		res = "Symbol_sizeFull"
	case Symbol_sizeCue:
		res = "Symbol_sizeCue"
	case Symbol_sizeGrace_cue:
		res = "Symbol_sizeGrace_cue"
	case Symbol_sizeLarge:
		res = "Symbol_sizeLarge"
	}
	return
}

func (symbol_size Symbol_size) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Symbol_sizeFull")
	res = append(res, "Symbol_sizeCue")
	res = append(res, "Symbol_sizeGrace_cue")
	res = append(res, "Symbol_sizeLarge")

	return
}

func (symbol_size Symbol_size) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "full")
	res = append(res, "cue")
	res = append(res, "grace-cue")
	res = append(res, "large")

	return
}

// Utility function for Sync_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (sync_type Sync_type) ToString() (res string) {

	// migration of former implementation of enum
	switch sync_type {
	// insertion code per enum code
	case Sync_typeNone:
		res = "none"
	case Sync_typeTempo:
		res = "tempo"
	case Sync_typeMostly_tempo:
		res = "mostly-tempo"
	case Sync_typeMostly_event:
		res = "mostly-event"
	case Sync_typeEvent:
		res = "event"
	case Sync_typeAlways_event:
		res = "always-event"
	}
	return
}

func (sync_type *Sync_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*sync_type = Sync_typeNone
		return
	case "tempo":
		*sync_type = Sync_typeTempo
		return
	case "mostly-tempo":
		*sync_type = Sync_typeMostly_tempo
		return
	case "mostly-event":
		*sync_type = Sync_typeMostly_event
		return
	case "event":
		*sync_type = Sync_typeEvent
		return
	case "always-event":
		*sync_type = Sync_typeAlways_event
		return
	default:
		return errUnkownEnum
	}
}

func (sync_type *Sync_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Sync_typeNone":
		*sync_type = Sync_typeNone
	case "Sync_typeTempo":
		*sync_type = Sync_typeTempo
	case "Sync_typeMostly_tempo":
		*sync_type = Sync_typeMostly_tempo
	case "Sync_typeMostly_event":
		*sync_type = Sync_typeMostly_event
	case "Sync_typeEvent":
		*sync_type = Sync_typeEvent
	case "Sync_typeAlways_event":
		*sync_type = Sync_typeAlways_event
	default:
		return errUnkownEnum
	}
	return
}

func (sync_type *Sync_type) ToCodeString() (res string) {

	switch *sync_type {
	// insertion code per enum code
	case Sync_typeNone:
		res = "Sync_typeNone"
	case Sync_typeTempo:
		res = "Sync_typeTempo"
	case Sync_typeMostly_tempo:
		res = "Sync_typeMostly_tempo"
	case Sync_typeMostly_event:
		res = "Sync_typeMostly_event"
	case Sync_typeEvent:
		res = "Sync_typeEvent"
	case Sync_typeAlways_event:
		res = "Sync_typeAlways_event"
	}
	return
}

func (sync_type Sync_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Sync_typeNone")
	res = append(res, "Sync_typeTempo")
	res = append(res, "Sync_typeMostly_tempo")
	res = append(res, "Sync_typeMostly_event")
	res = append(res, "Sync_typeEvent")
	res = append(res, "Sync_typeAlways_event")

	return
}

func (sync_type Sync_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "tempo")
	res = append(res, "mostly-tempo")
	res = append(res, "mostly-event")
	res = append(res, "event")
	res = append(res, "always-event")

	return
}

// Utility function for System_relation_number
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (system_relation_number System_relation_number) ToString() (res string) {

	// migration of former implementation of enum
	switch system_relation_number {
	// insertion code per enum code
	case System_relation_numberOnly_top:
		res = "only-top"
	case System_relation_numberOnly_bottom:
		res = "only-bottom"
	case System_relation_numberAlso_top:
		res = "also-top"
	case System_relation_numberAlso_bottom:
		res = "also-bottom"
	case System_relation_numberNone:
		res = "none"
	}
	return
}

func (system_relation_number *System_relation_number) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "only-top":
		*system_relation_number = System_relation_numberOnly_top
		return
	case "only-bottom":
		*system_relation_number = System_relation_numberOnly_bottom
		return
	case "also-top":
		*system_relation_number = System_relation_numberAlso_top
		return
	case "also-bottom":
		*system_relation_number = System_relation_numberAlso_bottom
		return
	case "none":
		*system_relation_number = System_relation_numberNone
		return
	default:
		return errUnkownEnum
	}
}

func (system_relation_number *System_relation_number) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "System_relation_numberOnly_top":
		*system_relation_number = System_relation_numberOnly_top
	case "System_relation_numberOnly_bottom":
		*system_relation_number = System_relation_numberOnly_bottom
	case "System_relation_numberAlso_top":
		*system_relation_number = System_relation_numberAlso_top
	case "System_relation_numberAlso_bottom":
		*system_relation_number = System_relation_numberAlso_bottom
	case "System_relation_numberNone":
		*system_relation_number = System_relation_numberNone
	default:
		return errUnkownEnum
	}
	return
}

func (system_relation_number *System_relation_number) ToCodeString() (res string) {

	switch *system_relation_number {
	// insertion code per enum code
	case System_relation_numberOnly_top:
		res = "System_relation_numberOnly_top"
	case System_relation_numberOnly_bottom:
		res = "System_relation_numberOnly_bottom"
	case System_relation_numberAlso_top:
		res = "System_relation_numberAlso_top"
	case System_relation_numberAlso_bottom:
		res = "System_relation_numberAlso_bottom"
	case System_relation_numberNone:
		res = "System_relation_numberNone"
	}
	return
}

func (system_relation_number System_relation_number) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "System_relation_numberOnly_top")
	res = append(res, "System_relation_numberOnly_bottom")
	res = append(res, "System_relation_numberAlso_top")
	res = append(res, "System_relation_numberAlso_bottom")
	res = append(res, "System_relation_numberNone")

	return
}

func (system_relation_number System_relation_number) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "only-top")
	res = append(res, "only-bottom")
	res = append(res, "also-top")
	res = append(res, "also-bottom")
	res = append(res, "none")

	return
}

// Utility function for Tap_hand
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tap_hand Tap_hand) ToString() (res string) {

	// migration of former implementation of enum
	switch tap_hand {
	// insertion code per enum code
	case Tap_handLeft:
		res = "left"
	case Tap_handRight:
		res = "right"
	}
	return
}

func (tap_hand *Tap_hand) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "left":
		*tap_hand = Tap_handLeft
		return
	case "right":
		*tap_hand = Tap_handRight
		return
	default:
		return errUnkownEnum
	}
}

func (tap_hand *Tap_hand) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tap_handLeft":
		*tap_hand = Tap_handLeft
	case "Tap_handRight":
		*tap_hand = Tap_handRight
	default:
		return errUnkownEnum
	}
	return
}

func (tap_hand *Tap_hand) ToCodeString() (res string) {

	switch *tap_hand {
	// insertion code per enum code
	case Tap_handLeft:
		res = "Tap_handLeft"
	case Tap_handRight:
		res = "Tap_handRight"
	}
	return
}

func (tap_hand Tap_hand) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tap_handLeft")
	res = append(res, "Tap_handRight")

	return
}

func (tap_hand Tap_hand) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "left")
	res = append(res, "right")

	return
}

// Utility function for Text_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (text_direction Text_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch text_direction {
	// insertion code per enum code
	case Text_directionLtr:
		res = "ltr"
	case Text_directionRtl:
		res = "rtl"
	case Text_directionLro:
		res = "lro"
	case Text_directionRlo:
		res = "rlo"
	}
	return
}

func (text_direction *Text_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ltr":
		*text_direction = Text_directionLtr
		return
	case "rtl":
		*text_direction = Text_directionRtl
		return
	case "lro":
		*text_direction = Text_directionLro
		return
	case "rlo":
		*text_direction = Text_directionRlo
		return
	default:
		return errUnkownEnum
	}
}

func (text_direction *Text_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Text_directionLtr":
		*text_direction = Text_directionLtr
	case "Text_directionRtl":
		*text_direction = Text_directionRtl
	case "Text_directionLro":
		*text_direction = Text_directionLro
	case "Text_directionRlo":
		*text_direction = Text_directionRlo
	default:
		return errUnkownEnum
	}
	return
}

func (text_direction *Text_direction) ToCodeString() (res string) {

	switch *text_direction {
	// insertion code per enum code
	case Text_directionLtr:
		res = "Text_directionLtr"
	case Text_directionRtl:
		res = "Text_directionRtl"
	case Text_directionLro:
		res = "Text_directionLro"
	case Text_directionRlo:
		res = "Text_directionRlo"
	}
	return
}

func (text_direction Text_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Text_directionLtr")
	res = append(res, "Text_directionRtl")
	res = append(res, "Text_directionLro")
	res = append(res, "Text_directionRlo")

	return
}

func (text_direction Text_direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ltr")
	res = append(res, "rtl")
	res = append(res, "lro")
	res = append(res, "rlo")

	return
}

// Utility function for Tied_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tied_type Tied_type) ToString() (res string) {

	// migration of former implementation of enum
	switch tied_type {
	// insertion code per enum code
	case Tied_typeStart:
		res = "start"
	case Tied_typeStop:
		res = "stop"
	case Tied_typeContinue_:
		res = "continue"
	case Tied_typeLet_ring:
		res = "let-ring"
	}
	return
}

func (tied_type *Tied_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*tied_type = Tied_typeStart
		return
	case "stop":
		*tied_type = Tied_typeStop
		return
	case "continue":
		*tied_type = Tied_typeContinue_
		return
	case "let-ring":
		*tied_type = Tied_typeLet_ring
		return
	default:
		return errUnkownEnum
	}
}

func (tied_type *Tied_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tied_typeStart":
		*tied_type = Tied_typeStart
	case "Tied_typeStop":
		*tied_type = Tied_typeStop
	case "Tied_typeContinue_":
		*tied_type = Tied_typeContinue_
	case "Tied_typeLet_ring":
		*tied_type = Tied_typeLet_ring
	default:
		return errUnkownEnum
	}
	return
}

func (tied_type *Tied_type) ToCodeString() (res string) {

	switch *tied_type {
	// insertion code per enum code
	case Tied_typeStart:
		res = "Tied_typeStart"
	case Tied_typeStop:
		res = "Tied_typeStop"
	case Tied_typeContinue_:
		res = "Tied_typeContinue_"
	case Tied_typeLet_ring:
		res = "Tied_typeLet_ring"
	}
	return
}

func (tied_type Tied_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tied_typeStart")
	res = append(res, "Tied_typeStop")
	res = append(res, "Tied_typeContinue_")
	res = append(res, "Tied_typeLet_ring")

	return
}

func (tied_type Tied_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "continue")
	res = append(res, "let-ring")

	return
}

// Utility function for Time_only
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (time_only Time_only) ToString() (res string) {

	// migration of former implementation of enum
	switch time_only {
	// insertion code per enum code
	}
	return
}

func (time_only *Time_only) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (time_only *Time_only) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (time_only *Time_only) ToCodeString() (res string) {

	switch *time_only {
	// insertion code per enum code
	}
	return
}

func (time_only Time_only) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (time_only Time_only) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// Utility function for Time_relation
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (time_relation Time_relation) ToString() (res string) {

	// migration of former implementation of enum
	switch time_relation {
	// insertion code per enum code
	case Time_relationParentheses:
		res = "parentheses"
	case Time_relationBracket:
		res = "bracket"
	case Time_relationEquals:
		res = "equals"
	case Time_relationSlash:
		res = "slash"
	case Time_relationSpace:
		res = "space"
	case Time_relationHyphen:
		res = "hyphen"
	}
	return
}

func (time_relation *Time_relation) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "parentheses":
		*time_relation = Time_relationParentheses
		return
	case "bracket":
		*time_relation = Time_relationBracket
		return
	case "equals":
		*time_relation = Time_relationEquals
		return
	case "slash":
		*time_relation = Time_relationSlash
		return
	case "space":
		*time_relation = Time_relationSpace
		return
	case "hyphen":
		*time_relation = Time_relationHyphen
		return
	default:
		return errUnkownEnum
	}
}

func (time_relation *Time_relation) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Time_relationParentheses":
		*time_relation = Time_relationParentheses
	case "Time_relationBracket":
		*time_relation = Time_relationBracket
	case "Time_relationEquals":
		*time_relation = Time_relationEquals
	case "Time_relationSlash":
		*time_relation = Time_relationSlash
	case "Time_relationSpace":
		*time_relation = Time_relationSpace
	case "Time_relationHyphen":
		*time_relation = Time_relationHyphen
	default:
		return errUnkownEnum
	}
	return
}

func (time_relation *Time_relation) ToCodeString() (res string) {

	switch *time_relation {
	// insertion code per enum code
	case Time_relationParentheses:
		res = "Time_relationParentheses"
	case Time_relationBracket:
		res = "Time_relationBracket"
	case Time_relationEquals:
		res = "Time_relationEquals"
	case Time_relationSlash:
		res = "Time_relationSlash"
	case Time_relationSpace:
		res = "Time_relationSpace"
	case Time_relationHyphen:
		res = "Time_relationHyphen"
	}
	return
}

func (time_relation Time_relation) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Time_relationParentheses")
	res = append(res, "Time_relationBracket")
	res = append(res, "Time_relationEquals")
	res = append(res, "Time_relationSlash")
	res = append(res, "Time_relationSpace")
	res = append(res, "Time_relationHyphen")

	return
}

func (time_relation Time_relation) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "parentheses")
	res = append(res, "bracket")
	res = append(res, "equals")
	res = append(res, "slash")
	res = append(res, "space")
	res = append(res, "hyphen")

	return
}

// Utility function for Time_separator
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (time_separator Time_separator) ToString() (res string) {

	// migration of former implementation of enum
	switch time_separator {
	// insertion code per enum code
	case Time_separatorNone:
		res = "none"
	case Time_separatorHorizontal:
		res = "horizontal"
	case Time_separatorDiagonal:
		res = "diagonal"
	case Time_separatorVertical:
		res = "vertical"
	case Time_separatorAdjacent:
		res = "adjacent"
	}
	return
}

func (time_separator *Time_separator) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*time_separator = Time_separatorNone
		return
	case "horizontal":
		*time_separator = Time_separatorHorizontal
		return
	case "diagonal":
		*time_separator = Time_separatorDiagonal
		return
	case "vertical":
		*time_separator = Time_separatorVertical
		return
	case "adjacent":
		*time_separator = Time_separatorAdjacent
		return
	default:
		return errUnkownEnum
	}
}

func (time_separator *Time_separator) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Time_separatorNone":
		*time_separator = Time_separatorNone
	case "Time_separatorHorizontal":
		*time_separator = Time_separatorHorizontal
	case "Time_separatorDiagonal":
		*time_separator = Time_separatorDiagonal
	case "Time_separatorVertical":
		*time_separator = Time_separatorVertical
	case "Time_separatorAdjacent":
		*time_separator = Time_separatorAdjacent
	default:
		return errUnkownEnum
	}
	return
}

func (time_separator *Time_separator) ToCodeString() (res string) {

	switch *time_separator {
	// insertion code per enum code
	case Time_separatorNone:
		res = "Time_separatorNone"
	case Time_separatorHorizontal:
		res = "Time_separatorHorizontal"
	case Time_separatorDiagonal:
		res = "Time_separatorDiagonal"
	case Time_separatorVertical:
		res = "Time_separatorVertical"
	case Time_separatorAdjacent:
		res = "Time_separatorAdjacent"
	}
	return
}

func (time_separator Time_separator) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Time_separatorNone")
	res = append(res, "Time_separatorHorizontal")
	res = append(res, "Time_separatorDiagonal")
	res = append(res, "Time_separatorVertical")
	res = append(res, "Time_separatorAdjacent")

	return
}

func (time_separator Time_separator) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "horizontal")
	res = append(res, "diagonal")
	res = append(res, "vertical")
	res = append(res, "adjacent")

	return
}

// Utility function for Time_symbol
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (time_symbol Time_symbol) ToString() (res string) {

	// migration of former implementation of enum
	switch time_symbol {
	// insertion code per enum code
	case Time_symbolCommon:
		res = "common"
	case Time_symbolCut:
		res = "cut"
	case Time_symbolSingle_number:
		res = "single-number"
	case Time_symbolNote:
		res = "note"
	case Time_symbolDotted_note:
		res = "dotted-note"
	case Time_symbolNormal:
		res = "normal"
	}
	return
}

func (time_symbol *Time_symbol) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "common":
		*time_symbol = Time_symbolCommon
		return
	case "cut":
		*time_symbol = Time_symbolCut
		return
	case "single-number":
		*time_symbol = Time_symbolSingle_number
		return
	case "note":
		*time_symbol = Time_symbolNote
		return
	case "dotted-note":
		*time_symbol = Time_symbolDotted_note
		return
	case "normal":
		*time_symbol = Time_symbolNormal
		return
	default:
		return errUnkownEnum
	}
}

func (time_symbol *Time_symbol) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Time_symbolCommon":
		*time_symbol = Time_symbolCommon
	case "Time_symbolCut":
		*time_symbol = Time_symbolCut
	case "Time_symbolSingle_number":
		*time_symbol = Time_symbolSingle_number
	case "Time_symbolNote":
		*time_symbol = Time_symbolNote
	case "Time_symbolDotted_note":
		*time_symbol = Time_symbolDotted_note
	case "Time_symbolNormal":
		*time_symbol = Time_symbolNormal
	default:
		return errUnkownEnum
	}
	return
}

func (time_symbol *Time_symbol) ToCodeString() (res string) {

	switch *time_symbol {
	// insertion code per enum code
	case Time_symbolCommon:
		res = "Time_symbolCommon"
	case Time_symbolCut:
		res = "Time_symbolCut"
	case Time_symbolSingle_number:
		res = "Time_symbolSingle_number"
	case Time_symbolNote:
		res = "Time_symbolNote"
	case Time_symbolDotted_note:
		res = "Time_symbolDotted_note"
	case Time_symbolNormal:
		res = "Time_symbolNormal"
	}
	return
}

func (time_symbol Time_symbol) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Time_symbolCommon")
	res = append(res, "Time_symbolCut")
	res = append(res, "Time_symbolSingle_number")
	res = append(res, "Time_symbolNote")
	res = append(res, "Time_symbolDotted_note")
	res = append(res, "Time_symbolNormal")

	return
}

func (time_symbol Time_symbol) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "common")
	res = append(res, "cut")
	res = append(res, "single-number")
	res = append(res, "note")
	res = append(res, "dotted-note")
	res = append(res, "normal")

	return
}

// Utility function for Tip_direction
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tip_direction Tip_direction) ToString() (res string) {

	// migration of former implementation of enum
	switch tip_direction {
	// insertion code per enum code
	case Tip_directionUp:
		res = "up"
	case Tip_directionDown:
		res = "down"
	case Tip_directionLeft:
		res = "left"
	case Tip_directionRight:
		res = "right"
	case Tip_directionNorthwest:
		res = "northwest"
	case Tip_directionNortheast:
		res = "northeast"
	case Tip_directionSoutheast:
		res = "southeast"
	case Tip_directionSouthwest:
		res = "southwest"
	}
	return
}

func (tip_direction *Tip_direction) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*tip_direction = Tip_directionUp
		return
	case "down":
		*tip_direction = Tip_directionDown
		return
	case "left":
		*tip_direction = Tip_directionLeft
		return
	case "right":
		*tip_direction = Tip_directionRight
		return
	case "northwest":
		*tip_direction = Tip_directionNorthwest
		return
	case "northeast":
		*tip_direction = Tip_directionNortheast
		return
	case "southeast":
		*tip_direction = Tip_directionSoutheast
		return
	case "southwest":
		*tip_direction = Tip_directionSouthwest
		return
	default:
		return errUnkownEnum
	}
}

func (tip_direction *Tip_direction) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tip_directionUp":
		*tip_direction = Tip_directionUp
	case "Tip_directionDown":
		*tip_direction = Tip_directionDown
	case "Tip_directionLeft":
		*tip_direction = Tip_directionLeft
	case "Tip_directionRight":
		*tip_direction = Tip_directionRight
	case "Tip_directionNorthwest":
		*tip_direction = Tip_directionNorthwest
	case "Tip_directionNortheast":
		*tip_direction = Tip_directionNortheast
	case "Tip_directionSoutheast":
		*tip_direction = Tip_directionSoutheast
	case "Tip_directionSouthwest":
		*tip_direction = Tip_directionSouthwest
	default:
		return errUnkownEnum
	}
	return
}

func (tip_direction *Tip_direction) ToCodeString() (res string) {

	switch *tip_direction {
	// insertion code per enum code
	case Tip_directionUp:
		res = "Tip_directionUp"
	case Tip_directionDown:
		res = "Tip_directionDown"
	case Tip_directionLeft:
		res = "Tip_directionLeft"
	case Tip_directionRight:
		res = "Tip_directionRight"
	case Tip_directionNorthwest:
		res = "Tip_directionNorthwest"
	case Tip_directionNortheast:
		res = "Tip_directionNortheast"
	case Tip_directionSoutheast:
		res = "Tip_directionSoutheast"
	case Tip_directionSouthwest:
		res = "Tip_directionSouthwest"
	}
	return
}

func (tip_direction Tip_direction) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tip_directionUp")
	res = append(res, "Tip_directionDown")
	res = append(res, "Tip_directionLeft")
	res = append(res, "Tip_directionRight")
	res = append(res, "Tip_directionNorthwest")
	res = append(res, "Tip_directionNortheast")
	res = append(res, "Tip_directionSoutheast")
	res = append(res, "Tip_directionSouthwest")

	return
}

func (tip_direction Tip_direction) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "left")
	res = append(res, "right")
	res = append(res, "northwest")
	res = append(res, "northeast")
	res = append(res, "southeast")
	res = append(res, "southwest")

	return
}

// Utility function for Top_bottom
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (top_bottom Top_bottom) ToString() (res string) {

	// migration of former implementation of enum
	switch top_bottom {
	// insertion code per enum code
	case Top_bottomTop:
		res = "top"
	case Top_bottomBottom:
		res = "bottom"
	}
	return
}

func (top_bottom *Top_bottom) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*top_bottom = Top_bottomTop
		return
	case "bottom":
		*top_bottom = Top_bottomBottom
		return
	default:
		return errUnkownEnum
	}
}

func (top_bottom *Top_bottom) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Top_bottomTop":
		*top_bottom = Top_bottomTop
	case "Top_bottomBottom":
		*top_bottom = Top_bottomBottom
	default:
		return errUnkownEnum
	}
	return
}

func (top_bottom *Top_bottom) ToCodeString() (res string) {

	switch *top_bottom {
	// insertion code per enum code
	case Top_bottomTop:
		res = "Top_bottomTop"
	case Top_bottomBottom:
		res = "Top_bottomBottom"
	}
	return
}

func (top_bottom Top_bottom) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Top_bottomTop")
	res = append(res, "Top_bottomBottom")

	return
}

func (top_bottom Top_bottom) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "bottom")

	return
}

// Utility function for Tremolo_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (tremolo_type Tremolo_type) ToString() (res string) {

	// migration of former implementation of enum
	switch tremolo_type {
	// insertion code per enum code
	case Tremolo_typeStart:
		res = "start"
	case Tremolo_typeStop:
		res = "stop"
	case Tremolo_typeSingle:
		res = "single"
	case Tremolo_typeUnmeasured:
		res = "unmeasured"
	}
	return
}

func (tremolo_type *Tremolo_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*tremolo_type = Tremolo_typeStart
		return
	case "stop":
		*tremolo_type = Tremolo_typeStop
		return
	case "single":
		*tremolo_type = Tremolo_typeSingle
		return
	case "unmeasured":
		*tremolo_type = Tremolo_typeUnmeasured
		return
	default:
		return errUnkownEnum
	}
}

func (tremolo_type *Tremolo_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Tremolo_typeStart":
		*tremolo_type = Tremolo_typeStart
	case "Tremolo_typeStop":
		*tremolo_type = Tremolo_typeStop
	case "Tremolo_typeSingle":
		*tremolo_type = Tremolo_typeSingle
	case "Tremolo_typeUnmeasured":
		*tremolo_type = Tremolo_typeUnmeasured
	default:
		return errUnkownEnum
	}
	return
}

func (tremolo_type *Tremolo_type) ToCodeString() (res string) {

	switch *tremolo_type {
	// insertion code per enum code
	case Tremolo_typeStart:
		res = "Tremolo_typeStart"
	case Tremolo_typeStop:
		res = "Tremolo_typeStop"
	case Tremolo_typeSingle:
		res = "Tremolo_typeSingle"
	case Tremolo_typeUnmeasured:
		res = "Tremolo_typeUnmeasured"
	}
	return
}

func (tremolo_type Tremolo_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Tremolo_typeStart")
	res = append(res, "Tremolo_typeStop")
	res = append(res, "Tremolo_typeSingle")
	res = append(res, "Tremolo_typeUnmeasured")

	return
}

func (tremolo_type Tremolo_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "stop")
	res = append(res, "single")
	res = append(res, "unmeasured")

	return
}

// Utility function for Trill_step
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (trill_step Trill_step) ToString() (res string) {

	// migration of former implementation of enum
	switch trill_step {
	// insertion code per enum code
	case Trill_stepWhole:
		res = "whole"
	case Trill_stepHalf:
		res = "half"
	case Trill_stepUnison:
		res = "unison"
	}
	return
}

func (trill_step *Trill_step) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "whole":
		*trill_step = Trill_stepWhole
		return
	case "half":
		*trill_step = Trill_stepHalf
		return
	case "unison":
		*trill_step = Trill_stepUnison
		return
	default:
		return errUnkownEnum
	}
}

func (trill_step *Trill_step) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Trill_stepWhole":
		*trill_step = Trill_stepWhole
	case "Trill_stepHalf":
		*trill_step = Trill_stepHalf
	case "Trill_stepUnison":
		*trill_step = Trill_stepUnison
	default:
		return errUnkownEnum
	}
	return
}

func (trill_step *Trill_step) ToCodeString() (res string) {

	switch *trill_step {
	// insertion code per enum code
	case Trill_stepWhole:
		res = "Trill_stepWhole"
	case Trill_stepHalf:
		res = "Trill_stepHalf"
	case Trill_stepUnison:
		res = "Trill_stepUnison"
	}
	return
}

func (trill_step Trill_step) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Trill_stepWhole")
	res = append(res, "Trill_stepHalf")
	res = append(res, "Trill_stepUnison")

	return
}

func (trill_step Trill_step) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "whole")
	res = append(res, "half")
	res = append(res, "unison")

	return
}

// Utility function for Two_note_turn
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (two_note_turn Two_note_turn) ToString() (res string) {

	// migration of former implementation of enum
	switch two_note_turn {
	// insertion code per enum code
	case Two_note_turnWhole:
		res = "whole"
	case Two_note_turnHalf:
		res = "half"
	case Two_note_turnNone:
		res = "none"
	}
	return
}

func (two_note_turn *Two_note_turn) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "whole":
		*two_note_turn = Two_note_turnWhole
		return
	case "half":
		*two_note_turn = Two_note_turnHalf
		return
	case "none":
		*two_note_turn = Two_note_turnNone
		return
	default:
		return errUnkownEnum
	}
}

func (two_note_turn *Two_note_turn) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Two_note_turnWhole":
		*two_note_turn = Two_note_turnWhole
	case "Two_note_turnHalf":
		*two_note_turn = Two_note_turnHalf
	case "Two_note_turnNone":
		*two_note_turn = Two_note_turnNone
	default:
		return errUnkownEnum
	}
	return
}

func (two_note_turn *Two_note_turn) ToCodeString() (res string) {

	switch *two_note_turn {
	// insertion code per enum code
	case Two_note_turnWhole:
		res = "Two_note_turnWhole"
	case Two_note_turnHalf:
		res = "Two_note_turnHalf"
	case Two_note_turnNone:
		res = "Two_note_turnNone"
	}
	return
}

func (two_note_turn Two_note_turn) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Two_note_turnWhole")
	res = append(res, "Two_note_turnHalf")
	res = append(res, "Two_note_turnNone")

	return
}

func (two_note_turn Two_note_turn) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "whole")
	res = append(res, "half")
	res = append(res, "none")

	return
}

// Utility function for Up_down
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (up_down Up_down) ToString() (res string) {

	// migration of former implementation of enum
	switch up_down {
	// insertion code per enum code
	case Up_downUp:
		res = "up"
	case Up_downDown:
		res = "down"
	}
	return
}

func (up_down *Up_down) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*up_down = Up_downUp
		return
	case "down":
		*up_down = Up_downDown
		return
	default:
		return errUnkownEnum
	}
}

func (up_down *Up_down) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Up_downUp":
		*up_down = Up_downUp
	case "Up_downDown":
		*up_down = Up_downDown
	default:
		return errUnkownEnum
	}
	return
}

func (up_down *Up_down) ToCodeString() (res string) {

	switch *up_down {
	// insertion code per enum code
	case Up_downUp:
		res = "Up_downUp"
	case Up_downDown:
		res = "Up_downDown"
	}
	return
}

func (up_down Up_down) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Up_downUp")
	res = append(res, "Up_downDown")

	return
}

func (up_down Up_down) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")

	return
}

// Utility function for Up_down_stop_continue
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (up_down_stop_continue Up_down_stop_continue) ToString() (res string) {

	// migration of former implementation of enum
	switch up_down_stop_continue {
	// insertion code per enum code
	case Up_down_stop_continueUp:
		res = "up"
	case Up_down_stop_continueDown:
		res = "down"
	case Up_down_stop_continueStop:
		res = "stop"
	case Up_down_stop_continueContinue_:
		res = "continue"
	}
	return
}

func (up_down_stop_continue *Up_down_stop_continue) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "up":
		*up_down_stop_continue = Up_down_stop_continueUp
		return
	case "down":
		*up_down_stop_continue = Up_down_stop_continueDown
		return
	case "stop":
		*up_down_stop_continue = Up_down_stop_continueStop
		return
	case "continue":
		*up_down_stop_continue = Up_down_stop_continueContinue_
		return
	default:
		return errUnkownEnum
	}
}

func (up_down_stop_continue *Up_down_stop_continue) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Up_down_stop_continueUp":
		*up_down_stop_continue = Up_down_stop_continueUp
	case "Up_down_stop_continueDown":
		*up_down_stop_continue = Up_down_stop_continueDown
	case "Up_down_stop_continueStop":
		*up_down_stop_continue = Up_down_stop_continueStop
	case "Up_down_stop_continueContinue_":
		*up_down_stop_continue = Up_down_stop_continueContinue_
	default:
		return errUnkownEnum
	}
	return
}

func (up_down_stop_continue *Up_down_stop_continue) ToCodeString() (res string) {

	switch *up_down_stop_continue {
	// insertion code per enum code
	case Up_down_stop_continueUp:
		res = "Up_down_stop_continueUp"
	case Up_down_stop_continueDown:
		res = "Up_down_stop_continueDown"
	case Up_down_stop_continueStop:
		res = "Up_down_stop_continueStop"
	case Up_down_stop_continueContinue_:
		res = "Up_down_stop_continueContinue_"
	}
	return
}

func (up_down_stop_continue Up_down_stop_continue) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Up_down_stop_continueUp")
	res = append(res, "Up_down_stop_continueDown")
	res = append(res, "Up_down_stop_continueStop")
	res = append(res, "Up_down_stop_continueContinue_")

	return
}

func (up_down_stop_continue Up_down_stop_continue) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "up")
	res = append(res, "down")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for Upright_inverted
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (upright_inverted Upright_inverted) ToString() (res string) {

	// migration of former implementation of enum
	switch upright_inverted {
	// insertion code per enum code
	case Upright_invertedUpright:
		res = "upright"
	case Upright_invertedInverted:
		res = "inverted"
	}
	return
}

func (upright_inverted *Upright_inverted) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "upright":
		*upright_inverted = Upright_invertedUpright
		return
	case "inverted":
		*upright_inverted = Upright_invertedInverted
		return
	default:
		return errUnkownEnum
	}
}

func (upright_inverted *Upright_inverted) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Upright_invertedUpright":
		*upright_inverted = Upright_invertedUpright
	case "Upright_invertedInverted":
		*upright_inverted = Upright_invertedInverted
	default:
		return errUnkownEnum
	}
	return
}

func (upright_inverted *Upright_inverted) ToCodeString() (res string) {

	switch *upright_inverted {
	// insertion code per enum code
	case Upright_invertedUpright:
		res = "Upright_invertedUpright"
	case Upright_invertedInverted:
		res = "Upright_invertedInverted"
	}
	return
}

func (upright_inverted Upright_inverted) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Upright_invertedUpright")
	res = append(res, "Upright_invertedInverted")

	return
}

func (upright_inverted Upright_inverted) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "upright")
	res = append(res, "inverted")

	return
}

// Utility function for Valign
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (valign Valign) ToString() (res string) {

	// migration of former implementation of enum
	switch valign {
	// insertion code per enum code
	case ValignTop:
		res = "top"
	case ValignMiddle:
		res = "middle"
	case ValignBottom:
		res = "bottom"
	case ValignBaseline:
		res = "baseline"
	}
	return
}

func (valign *Valign) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*valign = ValignTop
		return
	case "middle":
		*valign = ValignMiddle
		return
	case "bottom":
		*valign = ValignBottom
		return
	case "baseline":
		*valign = ValignBaseline
		return
	default:
		return errUnkownEnum
	}
}

func (valign *Valign) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ValignTop":
		*valign = ValignTop
	case "ValignMiddle":
		*valign = ValignMiddle
	case "ValignBottom":
		*valign = ValignBottom
	case "ValignBaseline":
		*valign = ValignBaseline
	default:
		return errUnkownEnum
	}
	return
}

func (valign *Valign) ToCodeString() (res string) {

	switch *valign {
	// insertion code per enum code
	case ValignTop:
		res = "ValignTop"
	case ValignMiddle:
		res = "ValignMiddle"
	case ValignBottom:
		res = "ValignBottom"
	case ValignBaseline:
		res = "ValignBaseline"
	}
	return
}

func (valign Valign) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ValignTop")
	res = append(res, "ValignMiddle")
	res = append(res, "ValignBottom")
	res = append(res, "ValignBaseline")

	return
}

func (valign Valign) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "middle")
	res = append(res, "bottom")
	res = append(res, "baseline")

	return
}

// Utility function for Valign_image
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (valign_image Valign_image) ToString() (res string) {

	// migration of former implementation of enum
	switch valign_image {
	// insertion code per enum code
	case Valign_imageTop:
		res = "top"
	case Valign_imageMiddle:
		res = "middle"
	case Valign_imageBottom:
		res = "bottom"
	}
	return
}

func (valign_image *Valign_image) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "top":
		*valign_image = Valign_imageTop
		return
	case "middle":
		*valign_image = Valign_imageMiddle
		return
	case "bottom":
		*valign_image = Valign_imageBottom
		return
	default:
		return errUnkownEnum
	}
}

func (valign_image *Valign_image) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Valign_imageTop":
		*valign_image = Valign_imageTop
	case "Valign_imageMiddle":
		*valign_image = Valign_imageMiddle
	case "Valign_imageBottom":
		*valign_image = Valign_imageBottom
	default:
		return errUnkownEnum
	}
	return
}

func (valign_image *Valign_image) ToCodeString() (res string) {

	switch *valign_image {
	// insertion code per enum code
	case Valign_imageTop:
		res = "Valign_imageTop"
	case Valign_imageMiddle:
		res = "Valign_imageMiddle"
	case Valign_imageBottom:
		res = "Valign_imageBottom"
	}
	return
}

func (valign_image Valign_image) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Valign_imageTop")
	res = append(res, "Valign_imageMiddle")
	res = append(res, "Valign_imageBottom")

	return
}

func (valign_image Valign_image) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "top")
	res = append(res, "middle")
	res = append(res, "bottom")

	return
}

// Utility function for Wedge_type
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (wedge_type Wedge_type) ToString() (res string) {

	// migration of former implementation of enum
	switch wedge_type {
	// insertion code per enum code
	case Wedge_typeCrescendo:
		res = "crescendo"
	case Wedge_typeDiminuendo:
		res = "diminuendo"
	case Wedge_typeStop:
		res = "stop"
	case Wedge_typeContinue_:
		res = "continue"
	}
	return
}

func (wedge_type *Wedge_type) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "crescendo":
		*wedge_type = Wedge_typeCrescendo
		return
	case "diminuendo":
		*wedge_type = Wedge_typeDiminuendo
		return
	case "stop":
		*wedge_type = Wedge_typeStop
		return
	case "continue":
		*wedge_type = Wedge_typeContinue_
		return
	default:
		return errUnkownEnum
	}
}

func (wedge_type *Wedge_type) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Wedge_typeCrescendo":
		*wedge_type = Wedge_typeCrescendo
	case "Wedge_typeDiminuendo":
		*wedge_type = Wedge_typeDiminuendo
	case "Wedge_typeStop":
		*wedge_type = Wedge_typeStop
	case "Wedge_typeContinue_":
		*wedge_type = Wedge_typeContinue_
	default:
		return errUnkownEnum
	}
	return
}

func (wedge_type *Wedge_type) ToCodeString() (res string) {

	switch *wedge_type {
	// insertion code per enum code
	case Wedge_typeCrescendo:
		res = "Wedge_typeCrescendo"
	case Wedge_typeDiminuendo:
		res = "Wedge_typeDiminuendo"
	case Wedge_typeStop:
		res = "Wedge_typeStop"
	case Wedge_typeContinue_:
		res = "Wedge_typeContinue_"
	}
	return
}

func (wedge_type Wedge_type) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Wedge_typeCrescendo")
	res = append(res, "Wedge_typeDiminuendo")
	res = append(res, "Wedge_typeStop")
	res = append(res, "Wedge_typeContinue_")

	return
}

func (wedge_type Wedge_type) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "crescendo")
	res = append(res, "diminuendo")
	res = append(res, "stop")
	res = append(res, "continue")

	return
}

// Utility function for Winged
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (winged Winged) ToString() (res string) {

	// migration of former implementation of enum
	switch winged {
	// insertion code per enum code
	case WingedNone:
		res = "none"
	case WingedStraight:
		res = "straight"
	case WingedCurved:
		res = "curved"
	case WingedDouble_straight:
		res = "double-straight"
	case WingedDouble_curved:
		res = "double-curved"
	}
	return
}

func (winged *Winged) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "none":
		*winged = WingedNone
		return
	case "straight":
		*winged = WingedStraight
		return
	case "curved":
		*winged = WingedCurved
		return
	case "double-straight":
		*winged = WingedDouble_straight
		return
	case "double-curved":
		*winged = WingedDouble_curved
		return
	default:
		return errUnkownEnum
	}
}

func (winged *Winged) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "WingedNone":
		*winged = WingedNone
	case "WingedStraight":
		*winged = WingedStraight
	case "WingedCurved":
		*winged = WingedCurved
	case "WingedDouble_straight":
		*winged = WingedDouble_straight
	case "WingedDouble_curved":
		*winged = WingedDouble_curved
	default:
		return errUnkownEnum
	}
	return
}

func (winged *Winged) ToCodeString() (res string) {

	switch *winged {
	// insertion code per enum code
	case WingedNone:
		res = "WingedNone"
	case WingedStraight:
		res = "WingedStraight"
	case WingedCurved:
		res = "WingedCurved"
	case WingedDouble_straight:
		res = "WingedDouble_straight"
	case WingedDouble_curved:
		res = "WingedDouble_curved"
	}
	return
}

func (winged Winged) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "WingedNone")
	res = append(res, "WingedStraight")
	res = append(res, "WingedCurved")
	res = append(res, "WingedDouble_straight")
	res = append(res, "WingedDouble_curved")

	return
}

func (winged Winged) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "none")
	res = append(res, "straight")
	res = append(res, "curved")
	res = append(res, "double-straight")
	res = append(res, "double-curved")

	return
}

// Utility function for Wood_value
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (wood_value Wood_value) ToString() (res string) {

	// migration of former implementation of enum
	switch wood_value {
	// insertion code per enum code
	case Wood_valueBambooscraper:
		res = "bamboo scraper"
	case Wood_valueBoardclapper:
		res = "board clapper"
	case Wood_valueCabasa:
		res = "cabasa"
	case Wood_valueCastanets:
		res = "castanets"
	case Wood_valueCastanetswithhandle:
		res = "castanets with handle"
	case Wood_valueClaves:
		res = "claves"
	case Wood_valueFootballrattle:
		res = "football rattle"
	case Wood_valueGuiro:
		res = "guiro"
	case Wood_valueLogdrum:
		res = "log drum"
	case Wood_valueMaraca:
		res = "maraca"
	case Wood_valueMaracas:
		res = "maracas"
	case Wood_valueQuijada:
		res = "quijada"
	case Wood_valueRainstick:
		res = "rainstick"
	case Wood_valueRatchet:
		res = "ratchet"
	case Wood_valueReco_reco:
		res = "reco-reco"
	case Wood_valueSandpaperblocks:
		res = "sandpaper blocks"
	case Wood_valueSlitdrum:
		res = "slit drum"
	case Wood_valueTempleblock:
		res = "temple block"
	case Wood_valueVibraslap:
		res = "vibraslap"
	case Wood_valueWhip:
		res = "whip"
	case Wood_valueWoodblock:
		res = "wood block"
	}
	return
}

func (wood_value *Wood_value) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "bamboo scraper":
		*wood_value = Wood_valueBambooscraper
		return
	case "board clapper":
		*wood_value = Wood_valueBoardclapper
		return
	case "cabasa":
		*wood_value = Wood_valueCabasa
		return
	case "castanets":
		*wood_value = Wood_valueCastanets
		return
	case "castanets with handle":
		*wood_value = Wood_valueCastanetswithhandle
		return
	case "claves":
		*wood_value = Wood_valueClaves
		return
	case "football rattle":
		*wood_value = Wood_valueFootballrattle
		return
	case "guiro":
		*wood_value = Wood_valueGuiro
		return
	case "log drum":
		*wood_value = Wood_valueLogdrum
		return
	case "maraca":
		*wood_value = Wood_valueMaraca
		return
	case "maracas":
		*wood_value = Wood_valueMaracas
		return
	case "quijada":
		*wood_value = Wood_valueQuijada
		return
	case "rainstick":
		*wood_value = Wood_valueRainstick
		return
	case "ratchet":
		*wood_value = Wood_valueRatchet
		return
	case "reco-reco":
		*wood_value = Wood_valueReco_reco
		return
	case "sandpaper blocks":
		*wood_value = Wood_valueSandpaperblocks
		return
	case "slit drum":
		*wood_value = Wood_valueSlitdrum
		return
	case "temple block":
		*wood_value = Wood_valueTempleblock
		return
	case "vibraslap":
		*wood_value = Wood_valueVibraslap
		return
	case "whip":
		*wood_value = Wood_valueWhip
		return
	case "wood block":
		*wood_value = Wood_valueWoodblock
		return
	default:
		return errUnkownEnum
	}
}

func (wood_value *Wood_value) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Wood_valueBambooscraper":
		*wood_value = Wood_valueBambooscraper
	case "Wood_valueBoardclapper":
		*wood_value = Wood_valueBoardclapper
	case "Wood_valueCabasa":
		*wood_value = Wood_valueCabasa
	case "Wood_valueCastanets":
		*wood_value = Wood_valueCastanets
	case "Wood_valueCastanetswithhandle":
		*wood_value = Wood_valueCastanetswithhandle
	case "Wood_valueClaves":
		*wood_value = Wood_valueClaves
	case "Wood_valueFootballrattle":
		*wood_value = Wood_valueFootballrattle
	case "Wood_valueGuiro":
		*wood_value = Wood_valueGuiro
	case "Wood_valueLogdrum":
		*wood_value = Wood_valueLogdrum
	case "Wood_valueMaraca":
		*wood_value = Wood_valueMaraca
	case "Wood_valueMaracas":
		*wood_value = Wood_valueMaracas
	case "Wood_valueQuijada":
		*wood_value = Wood_valueQuijada
	case "Wood_valueRainstick":
		*wood_value = Wood_valueRainstick
	case "Wood_valueRatchet":
		*wood_value = Wood_valueRatchet
	case "Wood_valueReco_reco":
		*wood_value = Wood_valueReco_reco
	case "Wood_valueSandpaperblocks":
		*wood_value = Wood_valueSandpaperblocks
	case "Wood_valueSlitdrum":
		*wood_value = Wood_valueSlitdrum
	case "Wood_valueTempleblock":
		*wood_value = Wood_valueTempleblock
	case "Wood_valueVibraslap":
		*wood_value = Wood_valueVibraslap
	case "Wood_valueWhip":
		*wood_value = Wood_valueWhip
	case "Wood_valueWoodblock":
		*wood_value = Wood_valueWoodblock
	default:
		return errUnkownEnum
	}
	return
}

func (wood_value *Wood_value) ToCodeString() (res string) {

	switch *wood_value {
	// insertion code per enum code
	case Wood_valueBambooscraper:
		res = "Wood_valueBambooscraper"
	case Wood_valueBoardclapper:
		res = "Wood_valueBoardclapper"
	case Wood_valueCabasa:
		res = "Wood_valueCabasa"
	case Wood_valueCastanets:
		res = "Wood_valueCastanets"
	case Wood_valueCastanetswithhandle:
		res = "Wood_valueCastanetswithhandle"
	case Wood_valueClaves:
		res = "Wood_valueClaves"
	case Wood_valueFootballrattle:
		res = "Wood_valueFootballrattle"
	case Wood_valueGuiro:
		res = "Wood_valueGuiro"
	case Wood_valueLogdrum:
		res = "Wood_valueLogdrum"
	case Wood_valueMaraca:
		res = "Wood_valueMaraca"
	case Wood_valueMaracas:
		res = "Wood_valueMaracas"
	case Wood_valueQuijada:
		res = "Wood_valueQuijada"
	case Wood_valueRainstick:
		res = "Wood_valueRainstick"
	case Wood_valueRatchet:
		res = "Wood_valueRatchet"
	case Wood_valueReco_reco:
		res = "Wood_valueReco_reco"
	case Wood_valueSandpaperblocks:
		res = "Wood_valueSandpaperblocks"
	case Wood_valueSlitdrum:
		res = "Wood_valueSlitdrum"
	case Wood_valueTempleblock:
		res = "Wood_valueTempleblock"
	case Wood_valueVibraslap:
		res = "Wood_valueVibraslap"
	case Wood_valueWhip:
		res = "Wood_valueWhip"
	case Wood_valueWoodblock:
		res = "Wood_valueWoodblock"
	}
	return
}

func (wood_value Wood_value) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Wood_valueBambooscraper")
	res = append(res, "Wood_valueBoardclapper")
	res = append(res, "Wood_valueCabasa")
	res = append(res, "Wood_valueCastanets")
	res = append(res, "Wood_valueCastanetswithhandle")
	res = append(res, "Wood_valueClaves")
	res = append(res, "Wood_valueFootballrattle")
	res = append(res, "Wood_valueGuiro")
	res = append(res, "Wood_valueLogdrum")
	res = append(res, "Wood_valueMaraca")
	res = append(res, "Wood_valueMaracas")
	res = append(res, "Wood_valueQuijada")
	res = append(res, "Wood_valueRainstick")
	res = append(res, "Wood_valueRatchet")
	res = append(res, "Wood_valueReco_reco")
	res = append(res, "Wood_valueSandpaperblocks")
	res = append(res, "Wood_valueSlitdrum")
	res = append(res, "Wood_valueTempleblock")
	res = append(res, "Wood_valueVibraslap")
	res = append(res, "Wood_valueWhip")
	res = append(res, "Wood_valueWoodblock")

	return
}

func (wood_value Wood_value) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "bamboo scraper")
	res = append(res, "board clapper")
	res = append(res, "cabasa")
	res = append(res, "castanets")
	res = append(res, "castanets with handle")
	res = append(res, "claves")
	res = append(res, "football rattle")
	res = append(res, "guiro")
	res = append(res, "log drum")
	res = append(res, "maraca")
	res = append(res, "maracas")
	res = append(res, "quijada")
	res = append(res, "rainstick")
	res = append(res, "ratchet")
	res = append(res, "reco-reco")
	res = append(res, "sandpaper blocks")
	res = append(res, "slit drum")
	res = append(res, "temple block")
	res = append(res, "vibraslap")
	res = append(res, "whip")
	res = append(res, "wood block")

	return
}

// Utility function for Yes_no
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (yes_no Yes_no) ToString() (res string) {

	// migration of former implementation of enum
	switch yes_no {
	// insertion code per enum code
	case Yes_noYes:
		res = "yes"
	case Yes_noNo:
		res = "no"
	}
	return
}

func (yes_no *Yes_no) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "yes":
		*yes_no = Yes_noYes
		return
	case "no":
		*yes_no = Yes_noNo
		return
	default:
		return errUnkownEnum
	}
}

func (yes_no *Yes_no) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Yes_noYes":
		*yes_no = Yes_noYes
	case "Yes_noNo":
		*yes_no = Yes_noNo
	default:
		return errUnkownEnum
	}
	return
}

func (yes_no *Yes_no) ToCodeString() (res string) {

	switch *yes_no {
	// insertion code per enum code
	case Yes_noYes:
		res = "Yes_noYes"
	case Yes_noNo:
		res = "Yes_noNo"
	}
	return
}

func (yes_no Yes_no) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Yes_noYes")
	res = append(res, "Yes_noNo")

	return
}

func (yes_no Yes_no) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "yes")
	res = append(res, "no")

	return
}

// Utility function for Yes_no_number
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (yes_no_number Yes_no_number) ToString() (res string) {

	// migration of former implementation of enum
	switch yes_no_number {
	// insertion code per enum code
	}
	return
}

func (yes_no_number *Yes_no_number) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
}

func (yes_no_number *Yes_no_number) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	default:
		return errUnkownEnum
	}
	return
}

func (yes_no_number *Yes_no_number) ToCodeString() (res string) {

	switch *yes_no_number {
	// insertion code per enum code
	}
	return
}

func (yes_no_number Yes_no_number) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

func (yes_no_number Yes_no_number) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | PositiveInteger
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*PositiveInteger
	FromCodeString(input string) (err error)
}

// Last line of the template
