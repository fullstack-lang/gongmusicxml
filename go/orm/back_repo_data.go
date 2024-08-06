// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	BookmarkAPIs []*BookmarkAPI

	FooAPIs []*FooAPI

	LinkAPIs []*LinkAPI

	LyricAPIs []*LyricAPI

	Lyric_fontAPIs []*Lyric_fontAPI

	Lyric_languageAPIs []*Lyric_languageAPI

	Miscellaneous_fieldAPIs []*Miscellaneous_fieldAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, bookmarkDB := range backRepo.BackRepoBookmark.Map_BookmarkDBID_BookmarkDB {

		var bookmarkAPI BookmarkAPI
		bookmarkAPI.ID = bookmarkDB.ID
		bookmarkAPI.BookmarkPointersEncoding = bookmarkDB.BookmarkPointersEncoding
		bookmarkDB.CopyBasicFieldsToBookmark_WOP(&bookmarkAPI.Bookmark_WOP)

		backRepoData.BookmarkAPIs = append(backRepoData.BookmarkAPIs, &bookmarkAPI)
	}

	for _, fooDB := range backRepo.BackRepoFoo.Map_FooDBID_FooDB {

		var fooAPI FooAPI
		fooAPI.ID = fooDB.ID
		fooAPI.FooPointersEncoding = fooDB.FooPointersEncoding
		fooDB.CopyBasicFieldsToFoo_WOP(&fooAPI.Foo_WOP)

		backRepoData.FooAPIs = append(backRepoData.FooAPIs, &fooAPI)
	}

	for _, linkDB := range backRepo.BackRepoLink.Map_LinkDBID_LinkDB {

		var linkAPI LinkAPI
		linkAPI.ID = linkDB.ID
		linkAPI.LinkPointersEncoding = linkDB.LinkPointersEncoding
		linkDB.CopyBasicFieldsToLink_WOP(&linkAPI.Link_WOP)

		backRepoData.LinkAPIs = append(backRepoData.LinkAPIs, &linkAPI)
	}

	for _, lyricDB := range backRepo.BackRepoLyric.Map_LyricDBID_LyricDB {

		var lyricAPI LyricAPI
		lyricAPI.ID = lyricDB.ID
		lyricAPI.LyricPointersEncoding = lyricDB.LyricPointersEncoding
		lyricDB.CopyBasicFieldsToLyric_WOP(&lyricAPI.Lyric_WOP)

		backRepoData.LyricAPIs = append(backRepoData.LyricAPIs, &lyricAPI)
	}

	for _, lyric_fontDB := range backRepo.BackRepoLyric_font.Map_Lyric_fontDBID_Lyric_fontDB {

		var lyric_fontAPI Lyric_fontAPI
		lyric_fontAPI.ID = lyric_fontDB.ID
		lyric_fontAPI.Lyric_fontPointersEncoding = lyric_fontDB.Lyric_fontPointersEncoding
		lyric_fontDB.CopyBasicFieldsToLyric_font_WOP(&lyric_fontAPI.Lyric_font_WOP)

		backRepoData.Lyric_fontAPIs = append(backRepoData.Lyric_fontAPIs, &lyric_fontAPI)
	}

	for _, lyric_languageDB := range backRepo.BackRepoLyric_language.Map_Lyric_languageDBID_Lyric_languageDB {

		var lyric_languageAPI Lyric_languageAPI
		lyric_languageAPI.ID = lyric_languageDB.ID
		lyric_languageAPI.Lyric_languagePointersEncoding = lyric_languageDB.Lyric_languagePointersEncoding
		lyric_languageDB.CopyBasicFieldsToLyric_language_WOP(&lyric_languageAPI.Lyric_language_WOP)

		backRepoData.Lyric_languageAPIs = append(backRepoData.Lyric_languageAPIs, &lyric_languageAPI)
	}

	for _, miscellaneous_fieldDB := range backRepo.BackRepoMiscellaneous_field.Map_Miscellaneous_fieldDBID_Miscellaneous_fieldDB {

		var miscellaneous_fieldAPI Miscellaneous_fieldAPI
		miscellaneous_fieldAPI.ID = miscellaneous_fieldDB.ID
		miscellaneous_fieldAPI.Miscellaneous_fieldPointersEncoding = miscellaneous_fieldDB.Miscellaneous_fieldPointersEncoding
		miscellaneous_fieldDB.CopyBasicFieldsToMiscellaneous_field_WOP(&miscellaneous_fieldAPI.Miscellaneous_field_WOP)

		backRepoData.Miscellaneous_fieldAPIs = append(backRepoData.Miscellaneous_fieldAPIs, &miscellaneous_fieldAPI)
	}

}
