// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Bookmark:
		ok = stage.IsStagedBookmark(target)

	case *Foo:
		ok = stage.IsStagedFoo(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *Lyric:
		ok = stage.IsStagedLyric(target)

	case *Lyric_font:
		ok = stage.IsStagedLyric_font(target)

	case *Lyric_language:
		ok = stage.IsStagedLyric_language(target)

	case *Miscellaneous_field:
		ok = stage.IsStagedMiscellaneous_field(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedBookmark(bookmark *Bookmark) (ok bool) {

	_, ok = stage.Bookmarks[bookmark]

	return
}

func (stage *StageStruct) IsStagedFoo(foo *Foo) (ok bool) {

	_, ok = stage.Foos[foo]

	return
}

func (stage *StageStruct) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

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

func (stage *StageStruct) IsStagedMiscellaneous_field(miscellaneous_field *Miscellaneous_field) (ok bool) {

	_, ok = stage.Miscellaneous_fields[miscellaneous_field]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Bookmark:
		stage.StageBranchBookmark(target)

	case *Foo:
		stage.StageBranchFoo(target)

	case *Link:
		stage.StageBranchLink(target)

	case *Lyric:
		stage.StageBranchLyric(target)

	case *Lyric_font:
		stage.StageBranchLyric_font(target)

	case *Lyric_language:
		stage.StageBranchLyric_language(target)

	case *Miscellaneous_field:
		stage.StageBranchMiscellaneous_field(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchBookmark(bookmark *Bookmark) {

	// check if instance is already staged
	if IsStaged(stage, bookmark) {
		return
	}

	bookmark.Stage(stage)

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

func (stage *StageStruct) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchLyric(lyric *Lyric) {

	// check if instance is already staged
	if IsStaged(stage, lyric) {
		return
	}

	lyric.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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

func (stage *StageStruct) StageBranchMiscellaneous_field(miscellaneous_field *Miscellaneous_field) {

	// check if instance is already staged
	if IsStaged(stage, miscellaneous_field) {
		return
	}

	miscellaneous_field.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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
	case *Bookmark:
		toT := CopyBranchBookmark(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Foo:
		toT := CopyBranchFoo(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := CopyBranchLink(mapOrigCopy, fromT)
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

	case *Miscellaneous_field:
		toT := CopyBranchMiscellaneous_field(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
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

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Bookmark:
		stage.UnstageBranchBookmark(target)

	case *Foo:
		stage.UnstageBranchFoo(target)

	case *Link:
		stage.UnstageBranchLink(target)

	case *Lyric:
		stage.UnstageBranchLyric(target)

	case *Lyric_font:
		stage.UnstageBranchLyric_font(target)

	case *Lyric_language:
		stage.UnstageBranchLyric_language(target)

	case *Miscellaneous_field:
		stage.UnstageBranchMiscellaneous_field(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchBookmark(bookmark *Bookmark) {

	// check if instance is already staged
	if !IsStaged(stage, bookmark) {
		return
	}

	bookmark.Unstage(stage)

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

func (stage *StageStruct) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchLyric(lyric *Lyric) {

	// check if instance is already staged
	if !IsStaged(stage, lyric) {
		return
	}

	lyric.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

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

func (stage *StageStruct) UnstageBranchMiscellaneous_field(miscellaneous_field *Miscellaneous_field) {

	// check if instance is already staged
	if !IsStaged(stage, miscellaneous_field) {
		return
	}

	miscellaneous_field.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
