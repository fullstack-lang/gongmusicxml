// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Bookmark:
		if stage.OnAfterBookmarkCreateCallback != nil {
			stage.OnAfterBookmarkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Foo:
		if stage.OnAfterFooCreateCallback != nil {
			stage.OnAfterFooCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lyric:
		if stage.OnAfterLyricCreateCallback != nil {
			stage.OnAfterLyricCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lyric_font:
		if stage.OnAfterLyric_fontCreateCallback != nil {
			stage.OnAfterLyric_fontCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lyric_language:
		if stage.OnAfterLyric_languageCreateCallback != nil {
			stage.OnAfterLyric_languageCreateCallback.OnAfterCreate(stage, target)
		}
	case *Miscellaneous_field:
		if stage.OnAfterMiscellaneous_fieldCreateCallback != nil {
			stage.OnAfterMiscellaneous_fieldCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Bookmark:
		newTarget := any(new).(*Bookmark)
		if stage.OnAfterBookmarkUpdateCallback != nil {
			stage.OnAfterBookmarkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Foo:
		newTarget := any(new).(*Foo)
		if stage.OnAfterFooUpdateCallback != nil {
			stage.OnAfterFooUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lyric:
		newTarget := any(new).(*Lyric)
		if stage.OnAfterLyricUpdateCallback != nil {
			stage.OnAfterLyricUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lyric_font:
		newTarget := any(new).(*Lyric_font)
		if stage.OnAfterLyric_fontUpdateCallback != nil {
			stage.OnAfterLyric_fontUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lyric_language:
		newTarget := any(new).(*Lyric_language)
		if stage.OnAfterLyric_languageUpdateCallback != nil {
			stage.OnAfterLyric_languageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Miscellaneous_field:
		newTarget := any(new).(*Miscellaneous_field)
		if stage.OnAfterMiscellaneous_fieldUpdateCallback != nil {
			stage.OnAfterMiscellaneous_fieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Bookmark:
		if stage.OnAfterBookmarkDeleteCallback != nil {
			staged := any(staged).(*Bookmark)
			stage.OnAfterBookmarkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Foo:
		if stage.OnAfterFooDeleteCallback != nil {
			staged := any(staged).(*Foo)
			stage.OnAfterFooDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lyric:
		if stage.OnAfterLyricDeleteCallback != nil {
			staged := any(staged).(*Lyric)
			stage.OnAfterLyricDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lyric_font:
		if stage.OnAfterLyric_fontDeleteCallback != nil {
			staged := any(staged).(*Lyric_font)
			stage.OnAfterLyric_fontDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lyric_language:
		if stage.OnAfterLyric_languageDeleteCallback != nil {
			staged := any(staged).(*Lyric_language)
			stage.OnAfterLyric_languageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Miscellaneous_field:
		if stage.OnAfterMiscellaneous_fieldDeleteCallback != nil {
			staged := any(staged).(*Miscellaneous_field)
			stage.OnAfterMiscellaneous_fieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Bookmark:
		if stage.OnAfterBookmarkReadCallback != nil {
			stage.OnAfterBookmarkReadCallback.OnAfterRead(stage, target)
		}
	case *Foo:
		if stage.OnAfterFooReadCallback != nil {
			stage.OnAfterFooReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	case *Lyric:
		if stage.OnAfterLyricReadCallback != nil {
			stage.OnAfterLyricReadCallback.OnAfterRead(stage, target)
		}
	case *Lyric_font:
		if stage.OnAfterLyric_fontReadCallback != nil {
			stage.OnAfterLyric_fontReadCallback.OnAfterRead(stage, target)
		}
	case *Lyric_language:
		if stage.OnAfterLyric_languageReadCallback != nil {
			stage.OnAfterLyric_languageReadCallback.OnAfterRead(stage, target)
		}
	case *Miscellaneous_field:
		if stage.OnAfterMiscellaneous_fieldReadCallback != nil {
			stage.OnAfterMiscellaneous_fieldReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bookmark:
		stage.OnAfterBookmarkUpdateCallback = any(callback).(OnAfterUpdateInterface[Bookmark])
	
	case *Foo:
		stage.OnAfterFooUpdateCallback = any(callback).(OnAfterUpdateInterface[Foo])
	
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	
	case *Lyric:
		stage.OnAfterLyricUpdateCallback = any(callback).(OnAfterUpdateInterface[Lyric])
	
	case *Lyric_font:
		stage.OnAfterLyric_fontUpdateCallback = any(callback).(OnAfterUpdateInterface[Lyric_font])
	
	case *Lyric_language:
		stage.OnAfterLyric_languageUpdateCallback = any(callback).(OnAfterUpdateInterface[Lyric_language])
	
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldUpdateCallback = any(callback).(OnAfterUpdateInterface[Miscellaneous_field])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bookmark:
		stage.OnAfterBookmarkCreateCallback = any(callback).(OnAfterCreateInterface[Bookmark])
	
	case *Foo:
		stage.OnAfterFooCreateCallback = any(callback).(OnAfterCreateInterface[Foo])
	
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	
	case *Lyric:
		stage.OnAfterLyricCreateCallback = any(callback).(OnAfterCreateInterface[Lyric])
	
	case *Lyric_font:
		stage.OnAfterLyric_fontCreateCallback = any(callback).(OnAfterCreateInterface[Lyric_font])
	
	case *Lyric_language:
		stage.OnAfterLyric_languageCreateCallback = any(callback).(OnAfterCreateInterface[Lyric_language])
	
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldCreateCallback = any(callback).(OnAfterCreateInterface[Miscellaneous_field])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bookmark:
		stage.OnAfterBookmarkDeleteCallback = any(callback).(OnAfterDeleteInterface[Bookmark])
	
	case *Foo:
		stage.OnAfterFooDeleteCallback = any(callback).(OnAfterDeleteInterface[Foo])
	
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	
	case *Lyric:
		stage.OnAfterLyricDeleteCallback = any(callback).(OnAfterDeleteInterface[Lyric])
	
	case *Lyric_font:
		stage.OnAfterLyric_fontDeleteCallback = any(callback).(OnAfterDeleteInterface[Lyric_font])
	
	case *Lyric_language:
		stage.OnAfterLyric_languageDeleteCallback = any(callback).(OnAfterDeleteInterface[Lyric_language])
	
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldDeleteCallback = any(callback).(OnAfterDeleteInterface[Miscellaneous_field])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Bookmark:
		stage.OnAfterBookmarkReadCallback = any(callback).(OnAfterReadInterface[Bookmark])
	
	case *Foo:
		stage.OnAfterFooReadCallback = any(callback).(OnAfterReadInterface[Foo])
	
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	
	case *Lyric:
		stage.OnAfterLyricReadCallback = any(callback).(OnAfterReadInterface[Lyric])
	
	case *Lyric_font:
		stage.OnAfterLyric_fontReadCallback = any(callback).(OnAfterReadInterface[Lyric_font])
	
	case *Lyric_language:
		stage.OnAfterLyric_languageReadCallback = any(callback).(OnAfterReadInterface[Lyric_language])
	
	case *Miscellaneous_field:
		stage.OnAfterMiscellaneous_fieldReadCallback = any(callback).(OnAfterReadInterface[Miscellaneous_field])
	
	}
}
